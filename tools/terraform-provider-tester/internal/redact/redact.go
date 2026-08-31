package redact

import (
	"sort"
	"strings"
)

const marker = "***REDACTED***"

// Redactor replaces known secret substrings with a fixed marker.
type Redactor struct {
	secrets []string
}

// New builds a Redactor from secret VALUES (not env keys). Empty or
// single-character values are ignored so we never blanket-redact output.
//
// Secrets are deduplicated and sorted by descending length so that when one
// secret is a substring of another (e.g. a token and that same token plus a
// suffix), String always masks the LONGEST match first. Replacing a shorter
// secret first would otherwise leave an exposed fragment of the longer, real
// secret in the output.
func New(values []string) *Redactor {
	seen := make(map[string]struct{}, len(values))
	var s []string
	for _, v := range values {
		if len(v) <= 1 {
			continue
		}
		if _, dup := seen[v]; dup {
			continue
		}
		seen[v] = struct{}{}
		s = append(s, v)
	}
	sort.SliceStable(s, func(i, j int) bool {
		return len(s[i]) > len(s[j])
	})
	return &Redactor{secrets: s}
}

func (r *Redactor) String(in string) string {
	out := in
	for _, sec := range r.secrets {
		out = strings.ReplaceAll(out, sec, marker)
	}
	return redactPatterns(out)
}

func (r *Redactor) Lines(in []string) []string {
	if len(in) == 0 {
		return nil
	}
	out := strings.SplitAfter(r.String(joinWithLineBoundaries(in)), "\n")
	if len(out) > 0 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}
	return out
}

func joinWithLineBoundaries(in []string) string {
	var b strings.Builder
	for i, line := range in {
		b.WriteString(line)
		if i+1 < len(in) && !strings.HasSuffix(line, "\n") && !strings.HasSuffix(line, "\r") {
			b.WriteByte('\n')
		}
	}
	return b.String()
}
