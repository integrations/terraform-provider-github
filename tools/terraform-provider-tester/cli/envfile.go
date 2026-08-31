package cli

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

const maxEnvLineBytes = 1 << 20

var keyRE = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

// parseEnvFile parses an env file from r and returns a map of key->value pairs.
// Each line is KEY=VALUE. Blank lines and # comments are ignored. An optional
// leading "export " is stripped. Values may be unquoted, single-quoted
// (literal, no expansion), or double-quoted (only \n, \t, \", \\ escapes;
// $VAR is never expanded). Malformed lines return an error citing the line number.
func parseEnvFile(r io.Reader) (map[string]string, error) {
	m := make(map[string]string)
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), maxEnvLineBytes)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimRight(scanner.Text(), "\r")

		trimmed := strings.TrimLeft(line, " \t")
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		rest := trimmed
		if strings.HasPrefix(rest, "export ") {
			rest = strings.TrimLeft(rest[7:], " \t")
		}

		eqIdx := strings.IndexByte(rest, '=')
		if eqIdx < 0 {
			return nil, fmt.Errorf("env file line %d: missing '='", lineNum)
		}

		key := strings.TrimRight(rest[:eqIdx], " \t")
		if !keyRE.MatchString(key) {
			return nil, fmt.Errorf("env file line %d: invalid key %q", lineNum, key)
		}

		val, err := parseValue(strings.TrimLeft(rest[eqIdx+1:], " \t"), lineNum)
		if err != nil {
			return nil, err
		}
		m[key] = val
	}
	return m, scanner.Err()
}

func parseValue(raw string, lineNum int) (string, error) {
	if strings.HasPrefix(raw, "'") {
		end := strings.Index(raw[1:], "'")
		if end < 0 {
			return "", fmt.Errorf("env file line %d: unterminated single quote", lineNum)
		}
		return raw[1 : end+1], nil
	}
	if strings.HasPrefix(raw, "\"") {
		return parseDoubleQuoted(raw[1:], lineNum)
	}
	// Unquoted: strip trailing comment (unquoted #) and trim trailing whitespace.
	return strings.TrimRight(stripTrailingComment(raw), " \t"), nil
}

func stripTrailingComment(s string) string {
	for i, c := range s {
		if c == '#' {
			return s[:i]
		}
	}
	return s
}

func parseDoubleQuoted(s string, lineNum int) (string, error) {
	var b strings.Builder
	i := 0
	for i < len(s) {
		c := s[i]
		if c == '"' {
			return b.String(), nil
		}
		if c == '\\' {
			i++
			if i >= len(s) {
				return "", fmt.Errorf("env file line %d: trailing backslash in double-quoted value", lineNum)
			}
			switch s[i] {
			case 'n':
				b.WriteByte('\n')
			case 't':
				b.WriteByte('\t')
			case '"':
				b.WriteByte('"')
			case '\\':
				b.WriteByte('\\')
			default:
				return "", fmt.Errorf("env file line %d: unsupported escape \\%c in double-quoted value", lineNum, s[i])
			}
		} else {
			b.WriteByte(c)
		}
		i++
	}
	return "", fmt.Errorf("env file line %d: unterminated double quote", lineNum)
}

// resolveEnvFile applies vars to the environment through setenv, respecting
// precedence: real env > mode-scoped PULSAR_<MODE>_KEY > plain KEY. A key
// already set in the real env is never overridden. Returns the count of
// distinct bare keys newly written.
//
// knownModes lists the real auth mode names. The PULSAR_<mode>_ prefix is
// matched case-insensitively against them, so PULSAR_organization_GITHUB_OWNER
// and PULSAR_ORGANIZATION_GITHUB_OWNER both scope to organization mode. Only
// real mode names are reserved, so other PULSAR_* names (for example
// PULSAR_FORCE_TTY) stay settable as plain variables through the same file.
//
// Precedence is enforced by snapshotting original emptiness before any setenv
// calls, so a mode-scoped key always wins over the plain key it shadows even
// though they both target the same bare name.
func resolveEnvFile(vars map[string]string, mode string, knownModes []string, getenv func(string) string, setenv func(string, string) error) (int, error) {
	modeSet := make(map[string]bool, len(knownModes))
	for _, m := range knownModes {
		modeSet[strings.ToLower(m)] = true
	}
	lowerMode := strings.ToLower(mode)

	// Collect all bare keys we will attempt to set so we can snapshot
	// original emptiness before any mutations.
	type entry struct{ bare, val string }
	var plain, scoped []entry
	for k, v := range vars {
		if rest, isPulsar := strings.CutPrefix(k, "PULSAR_"); isPulsar {
			// PULSAR_<modePart>_<bare>: the mode part runs up to the first
			// underscore, and only real mode names are reserved.
			if modePart, bare, hasBare := strings.Cut(rest, "_"); hasBare && modeSet[strings.ToLower(modePart)] {
				if strings.ToLower(modePart) == lowerMode && keyRE.MatchString(bare) {
					scoped = append(scoped, entry{bare, v})
				}
				// A known but non-active mode is intentionally ignored.
				continue
			}
			// Not a real mode prefix: fall through and keep the full PULSAR_
			// name so the rest of that namespace stays usable.
		}
		plain = append(plain, entry{k, v})
	}

	// Snapshot original emptiness before any setenv calls.
	originalEmpty := make(map[string]bool, len(plain)+len(scoped))
	for _, e := range plain {
		originalEmpty[e.bare] = getenv(e.bare) == ""
	}
	for _, e := range scoped {
		originalEmpty[e.bare] = getenv(e.bare) == ""
	}

	set := make(map[string]bool, len(plain)+len(scoped))

	// First pass: plain keys (only when originally unset).
	for _, e := range plain {
		if !originalEmpty[e.bare] {
			continue
		}
		if err := setenv(e.bare, e.val); err != nil {
			return len(set), fmt.Errorf("setting %s: %w", e.bare, err)
		}
		set[e.bare] = true
	}

	// Second pass: mode-scoped keys win over plain (still guard on original env).
	for _, e := range scoped {
		if !originalEmpty[e.bare] {
			continue
		}
		if err := setenv(e.bare, e.val); err != nil {
			return len(set), fmt.Errorf("setting %s: %w", e.bare, err)
		}
		set[e.bare] = true
	}

	return len(set), nil
}
