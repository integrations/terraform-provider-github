package cli

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/fakeprovider"
	"github.com/github/terraform-provider-tester/provider"
)

// ── parseEnvFile ─────────────────────────────────────────────────────────────

func TestParseEnvFileBasic(t *testing.T) {
	input := `
# full-line comment
KEY1=value1
KEY2 = value2  # trailing comment
export KEY3=value3
KEY4=
KEY5=hello world
`
	got, err := parseEnvFile(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	cases := map[string]string{
		"KEY1": "value1",
		"KEY2": "value2",
		"KEY3": "value3",
		"KEY4": "",
		"KEY5": "hello world",
	}
	for k, want := range cases {
		if got[k] != want {
			t.Errorf("KEY %s: got %q, want %q", k, got[k], want)
		}
	}
	// comment line should not be a key
	for k := range got {
		if strings.HasPrefix(k, "#") {
			t.Errorf("comment parsed as key: %q", k)
		}
	}
}

func TestParseEnvFileSingleQuoted(t *testing.T) {
	input := "KEY='it is literal $VAR \\n no expand'\n"
	got, err := parseEnvFile(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := `it is literal $VAR \n no expand`
	if got["KEY"] != want {
		t.Errorf("got %q, want %q", got["KEY"], want)
	}
}

func TestParseEnvFileDoubleQuoted(t *testing.T) {
	input := `KEY="line1\nline2\ttab\"quote\\back"`
	got, err := parseEnvFile(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "line1\nline2\ttab\"quote\\back"
	if got["KEY"] != want {
		t.Errorf("got %q, want %q", got["KEY"], want)
	}
}

func TestParseEnvFileDoubleQuotedNoShellExpansion(t *testing.T) {
	// $VAR in a double-quoted value must be preserved literally, not expanded.
	input := `KEY="hello $USER world"`
	got, err := parseEnvFile(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "hello $USER world"
	if got["KEY"] != want {
		t.Errorf("got %q, want %q", got["KEY"], want)
	}
}

func TestParseEnvFileCRLF(t *testing.T) {
	input := "KEY1=a\r\nKEY2=b\r\n"
	got, err := parseEnvFile(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got["KEY1"] != "a" || got["KEY2"] != "b" {
		t.Errorf("CRLF: got %v", got)
	}
}

func TestParseEnvFileMalformedLineReturnsError(t *testing.T) {
	input := "GOODKEY=ok\nNO_EQUALS_HERE\nANOTHER=val\n"
	_, err := parseEnvFile(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for malformed line, got nil")
	}
	if !strings.Contains(err.Error(), "line 2") {
		t.Errorf("error should mention line 2: %v", err)
	}
}

func TestParseEnvFileInvalidKey(t *testing.T) {
	input := "123INVALID=value\n"
	_, err := parseEnvFile(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for invalid key, got nil")
	}
	if !strings.Contains(err.Error(), "line 1") {
		t.Errorf("error should mention line 1: %v", err)
	}
}

func TestParseEnvFileUnterminatedSingleQuote(t *testing.T) {
	input := "KEY='unterminated\n"
	_, err := parseEnvFile(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for unterminated single quote")
	}
}

func TestParseEnvFileUnterminatedDoubleQuote(t *testing.T) {
	input := `KEY="unterminated`
	_, err := parseEnvFile(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for unterminated double quote")
	}
}

func TestParseEnvFileDollarVarUnquotedLiteral(t *testing.T) {
	// $VAR in an unquoted value must be preserved literally.
	input := "KEY=$NOTEXPANDED\n"
	got, err := parseEnvFile(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got["KEY"] != "$NOTEXPANDED" {
		t.Errorf("got %q, want %q", got["KEY"], "$NOTEXPANDED")
	}
}

func TestParseEnvFileExportWithSpaces(t *testing.T) {
	input := "export   KEY=value\n"
	got, err := parseEnvFile(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got["KEY"] != "value" {
		t.Errorf("got %q, want %q", got["KEY"], "value")
	}
}

func TestParseEnvFileLargeLine(t *testing.T) {
	// 256 KB value: well above the default 64 KB scanner limit, below 1 MB cap.
	bigVal := strings.Repeat("x", 256*1024)
	input := "BIG=" + bigVal + "\n"
	got, err := parseEnvFile(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error for 256 KB value: %v", err)
	}
	if got["BIG"] != bigVal {
		t.Errorf("BIG length = %d, want %d", len(got["BIG"]), len(bigVal))
	}

	// Value just under the 1 MB cap must also parse.
	nearCapVal := strings.Repeat("y", maxEnvLineBytes-len("BIG=")-2)
	input2 := "BIG=" + nearCapVal + "\n"
	got2, err2 := parseEnvFile(strings.NewReader(input2))
	if err2 != nil {
		t.Fatalf("unexpected error for near-cap value: %v", err2)
	}
	if got2["BIG"] != nearCapVal {
		t.Errorf("near-cap BIG length = %d, want %d", len(got2["BIG"]), len(nearCapVal))
	}
}

// ── resolveEnvFile ────────────────────────────────────────────────────────────

// mapEnv returns injectable getenv/setenv over a local map for hermetic tests.
func mapEnv(initial map[string]string) (func(string) string, func(string, string) error, *map[string]string) {
	m := make(map[string]string, len(initial))
	for k, v := range initial {
		m[k] = v
	}
	get := func(k string) string { return m[k] }
	set := func(k, v string) error { m[k] = v; return nil }
	return get, set, &m
}

func getenvFromMap(values map[string]string) func(string) string {
	return func(key string) string {
		return values[key]
	}
}

// accModes mirrors the provider's real auth mode names for resolveEnvFile tests.
var accModes = []string{"anonymous", "individual", "organization", "team", "enterprise"}

func TestResolveEnvFileModeCaseInsensitive(t *testing.T) {
	// A PULSAR_<mode>_ prefix must match the active mode case-insensitively, so
	// the lowercase form documented in older examples still resolves.
	for _, key := range []string{
		"PULSAR_organization_GITHUB_OWNER",
		"PULSAR_Organization_GITHUB_OWNER",
		"PULSAR_ORGANIZATION_GITHUB_OWNER",
	} {
		vars := map[string]string{key: "acme-test"}
		get, set, m := mapEnv(nil)
		if _, err := resolveEnvFile(vars, "organization", accModes, get, set); err != nil {
			t.Fatalf("%s: unexpected error: %v", key, err)
		}
		if (*m)["GITHUB_OWNER"] != "acme-test" {
			t.Errorf("%s: GITHUB_OWNER = %q, want acme-test", key, (*m)["GITHUB_OWNER"])
		}
	}
}

func TestResolveEnvFileNonModePulsarKeyStaysUsable(t *testing.T) {
	// PULSAR_ names whose first segment is not a real mode (behavior vars) must
	// remain settable through the file under their full name.
	vars := map[string]string{
		"PULSAR_FORCE_TTY": "1",
		"PULSAR_NO_TUI":    "1",
	}
	get, set, m := mapEnv(nil)
	if _, err := resolveEnvFile(vars, "anonymous", accModes, get, set); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if (*m)["PULSAR_FORCE_TTY"] != "1" {
		t.Errorf("PULSAR_FORCE_TTY = %q, want 1", (*m)["PULSAR_FORCE_TTY"])
	}
	if (*m)["PULSAR_NO_TUI"] != "1" {
		t.Errorf("PULSAR_NO_TUI = %q, want 1", (*m)["PULSAR_NO_TUI"])
	}
}

func TestResolveEnvFileUnknownModePrefixKeptLiteral(t *testing.T) {
	// An unknown mode segment must not be promoted to a bare key; it is kept
	// under its literal PULSAR_ name rather than silently dropped.
	vars := map[string]string{"PULSAR_BOGUS_GITHUB_OWNER": "x"}
	get, set, m := mapEnv(nil)
	if _, err := resolveEnvFile(vars, "organization", accModes, get, set); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if (*m)["GITHUB_OWNER"] != "" {
		t.Errorf("GITHUB_OWNER = %q, want unset for unknown mode prefix", (*m)["GITHUB_OWNER"])
	}
	if (*m)["PULSAR_BOGUS_GITHUB_OWNER"] != "x" {
		t.Errorf("PULSAR_BOGUS_GITHUB_OWNER = %q, want x (kept literal)", (*m)["PULSAR_BOGUS_GITHUB_OWNER"])
	}
}

func TestResolveEnvFileNoOverride(t *testing.T) {
	vars := map[string]string{
		"GITHUB_TOKEN": "from_file",
		"GITHUB_OWNER": "from_file_owner",
	}
	// Real env already has GITHUB_TOKEN set.
	get, set, m := mapEnv(map[string]string{"GITHUB_TOKEN": "real_token"})

	n, err := resolveEnvFile(vars, "anonymous", accModes, get, set)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// GITHUB_TOKEN must not be overridden.
	if (*m)["GITHUB_TOKEN"] != "real_token" {
		t.Errorf("GITHUB_TOKEN overridden: got %q, want %q", (*m)["GITHUB_TOKEN"], "real_token")
	}
	// GITHUB_OWNER was unset, so file value should apply.
	if (*m)["GITHUB_OWNER"] != "from_file_owner" {
		t.Errorf("GITHUB_OWNER not set: got %q, want %q", (*m)["GITHUB_OWNER"], "from_file_owner")
	}
	// loaded count: only GITHUB_OWNER was newly set.
	if n != 1 {
		t.Errorf("loaded = %d, want 1", n)
	}
}

func TestResolveEnvFileModeScoped(t *testing.T) {
	vars := map[string]string{
		"PULSAR_INDIVIDUAL_GITHUB_OWNER":   "alice",
		"PULSAR_ORGANIZATION_GITHUB_OWNER": "acme-test",
		"GITHUB_TOKEN":                     "shared_token",
	}

	// individual mode: GITHUB_OWNER should become "alice"
	{
		get, set, m := mapEnv(nil)
		n, err := resolveEnvFile(vars, "individual", accModes, get, set)
		if err != nil {
			t.Fatalf("individual mode error: %v", err)
		}
		if (*m)["GITHUB_OWNER"] != "alice" {
			t.Errorf("individual: GITHUB_OWNER = %q, want alice", (*m)["GITHUB_OWNER"])
		}
		if (*m)["GITHUB_TOKEN"] != "shared_token" {
			t.Errorf("individual: GITHUB_TOKEN = %q, want shared_token", (*m)["GITHUB_TOKEN"])
		}
		// PULSAR_ prefix keys should NOT be set as bare PULSAR_ keys.
		if _, ok := (*m)["PULSAR_INDIVIDUAL_GITHUB_OWNER"]; ok {
			t.Error("PULSAR_INDIVIDUAL_GITHUB_OWNER should not be set as a bare key")
		}
		_ = n
	}

	// organization mode: GITHUB_OWNER should become "acme-test"
	{
		get, set, m := mapEnv(nil)
		_, err := resolveEnvFile(vars, "organization", accModes, get, set)
		if err != nil {
			t.Fatalf("organization mode error: %v", err)
		}
		if (*m)["GITHUB_OWNER"] != "acme-test" {
			t.Errorf("organization: GITHUB_OWNER = %q, want acme-test", (*m)["GITHUB_OWNER"])
		}
		if (*m)["GITHUB_TOKEN"] != "shared_token" {
			t.Errorf("organization: GITHUB_TOKEN = %q, want shared_token", (*m)["GITHUB_TOKEN"])
		}
	}
}

func TestResolveEnvFileModeScopedWinsOverPlain(t *testing.T) {
	// Both a plain and a mode-scoped key for GITHUB_OWNER exist.
	// Mode-scoped must win.
	vars := map[string]string{
		"GITHUB_OWNER":                   "plain_owner",
		"PULSAR_INDIVIDUAL_GITHUB_OWNER": "scoped_owner",
	}
	get, set, m := mapEnv(nil)
	_, err := resolveEnvFile(vars, "individual", accModes, get, set)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if (*m)["GITHUB_OWNER"] != "scoped_owner" {
		t.Errorf("GITHUB_OWNER = %q, want scoped_owner (mode-scoped must win)", (*m)["GITHUB_OWNER"])
	}
}

func TestResolveEnvFileSharedKeyAllModes(t *testing.T) {
	vars := map[string]string{"GITHUB_TOKEN": "tok123"}
	for _, mode := range []string{"individual", "organization", "team", "enterprise", "anonymous"} {
		get, set, m := mapEnv(nil)
		if _, err := resolveEnvFile(vars, mode, accModes, get, set); err != nil {
			t.Fatalf("mode %s: %v", mode, err)
		}
		if (*m)["GITHUB_TOKEN"] != "tok123" {
			t.Errorf("mode %s: GITHUB_TOKEN = %q, want tok123", mode, (*m)["GITHUB_TOKEN"])
		}
	}
}

func TestResolveEnvFileIgnoresOtherModePrefixes(t *testing.T) {
	// PULSAR_ORGANIZATION_ keys should NOT be applied when mode is "individual".
	vars := map[string]string{
		"PULSAR_ORGANIZATION_GITHUB_OWNER": "org_owner",
	}
	get, set, m := mapEnv(nil)
	_, err := resolveEnvFile(vars, "individual", accModes, get, set)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if (*m)["GITHUB_OWNER"] != "" {
		t.Errorf("GITHUB_OWNER should be unset for individual mode, got %q", (*m)["GITHUB_OWNER"])
	}
}

// ── integration: runWithDeps wires env-file loading ───────────────────────────

func TestEnvFileAutoDetectionPrecedence(t *testing.T) {
	for _, tc := range []struct {
		name           string
		explicit       bool
		envOverride    bool
		cwdFile        bool
		userConfigFile bool
		homeFile       bool
		wantToken      string
	}{
		{
			name:           "cwd file wins over user config and home",
			cwdFile:        true,
			userConfigFile: true,
			homeFile:       true,
			wantToken:      "from_cwd",
		},
		{
			name:           "user config file used when cwd file is absent",
			userConfigFile: true,
			homeFile:       true,
			wantToken:      "from_user_config",
		},
		{
			name:      "home file used when cwd and user config are absent",
			homeFile:  true,
			wantToken: "from_home",
		},
		{
			name:           "explicit flag wins over all auto locations",
			explicit:       true,
			envOverride:    true,
			cwdFile:        true,
			userConfigFile: true,
			homeFile:       true,
			wantToken:      "from_explicit",
		},
		{
			name:           "env override wins over auto locations",
			envOverride:    true,
			cwdFile:        true,
			userConfigFile: true,
			homeFile:       true,
			wantToken:      "from_env_override",
		},
		{
			name:      "no auto file is not an error",
			wantToken: "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("GITHUB_TOKEN", "")
			t.Setenv("GH_TEST_AUTH_MODE", "")

			root := t.TempDir()
			t.Setenv("PULSAR_ENV_FILE", filepath.Join(root, "process-env-must-not-be-read.env"))
			userConfigDir := filepath.Join(t.TempDir(), "config")
			userHomeDir := filepath.Join(t.TempDir(), "home")
			writeEnv := func(path, token string) {
				t.Helper()
				if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(path, []byte("GITHUB_TOKEN="+token+"\n"), 0o600); err != nil {
					t.Fatal(err)
				}
			}

			explicitFile := filepath.Join(root, "explicit.env")
			envOverrideFile := filepath.Join(root, "env-override.env")
			cwdFile := filepath.Join(root, ".pulsar.env")
			userConfigFile := filepath.Join(userConfigDir, "terraform-provider-tester", ".pulsar.env")
			homeFile := filepath.Join(userHomeDir, ".pulsar.env")

			if tc.explicit {
				writeEnv(explicitFile, "from_explicit")
			}
			if tc.envOverride {
				writeEnv(envOverrideFile, "from_env_override")
			}
			if tc.cwdFile {
				writeEnv(cwdFile, "from_cwd")
			}
			if tc.userConfigFile {
				writeEnv(userConfigFile, "from_user_config")
			}
			if tc.homeFile {
				writeEnv(homeFile, "from_home")
			}

			getenv := func(key string) string {
				switch key {
				case "PULSAR_ENV_FILE":
					if tc.envOverride {
						return envOverrideFile
					}
					return ""
				case "GITHUB_TOKEN", "GH_TEST_AUTH_MODE":
					return ""
				default:
					return ""
				}
			}
			fr := &fakeRunner{result: engine.RunResult{}}
			pf := &fakeprovider.Fake{
				NameVal:  "test",
				Packages: []string{"./..."},
				Pattern:  "^TestAcc",
				ModesVal: testModes,
				EnvByMode: map[string][]provider.EnvVar{
					"anonymous": {{Key: "GITHUB_TOKEN"}},
				},
				RequirementsFn: allowAllRequirements,
				PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
					return provider.PreflightReport{Mode: m}
				},
			}
			d := deps{
				provider:      pf,
				newRunner:     func(_ io.Writer) testRunner { return fr },
				list:          stubList([]string{"TestAccA"}),
				cwd:           func() (string, error) { return root, nil },
				userConfigDir: func() (string, error) { return userConfigDir, nil },
				userHomeDir:   func() (string, error) { return userHomeDir, nil },
				getenv:        getenv,
			}

			args := []string{"run", "--repo-root", root}
			if tc.explicit {
				args = []string{"run", "--env-file", explicitFile, "--repo-root", root}
			}
			var out, errOut strings.Builder
			code := runWithDeps(args, &out, &errOut, d)
			if code == 2 {
				t.Fatalf("exit code 2: stderr=%s", errOut.String())
			}
			if tc.wantToken == "" {
				for _, val := range []string{"from_explicit", "from_env_override", "from_cwd", "from_user_config", "from_home"} {
					if extraEnvContains(fr.spec.ExtraEnv, "GITHUB_TOKEN="+val) {
						t.Fatalf("ExtraEnv unexpectedly loaded %s: %v", val, fr.spec.ExtraEnv)
					}
				}
				return
			}
			want := "GITHUB_TOKEN=" + tc.wantToken
			if !extraEnvContains(fr.spec.ExtraEnv, want) {
				t.Errorf("ExtraEnv missing %s; got: %v", want, fr.spec.ExtraEnv)
			}
		})
	}
}

func TestEnvFileFromEnvMissingIsError(t *testing.T) {
	root := t.TempDir()
	missing := filepath.Join(root, "missing.env")
	processEnvPath := filepath.Join(root, "process-env-must-not-be-read.env")
	t.Setenv("PULSAR_ENV_FILE", processEnvPath)
	d := deps{
		provider:      &fakeprovider.Fake{},
		list:          stubList(nil),
		cwd:           func() (string, error) { return root, nil },
		userConfigDir: func() (string, error) { return filepath.Join(root, "config"), nil },
		userHomeDir:   func() (string, error) { return filepath.Join(root, "home"), nil },
		getenv: func(key string) string {
			if key == "PULSAR_ENV_FILE" {
				return missing
			}
			return ""
		},
	}

	var out, errOut strings.Builder
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)
	if code != 2 {
		t.Fatalf("expected exit 2 for missing PULSAR_ENV_FILE path, got %d", code)
	}
	if !strings.Contains(errOut.String(), "env-file:") {
		t.Fatalf("expected env-file error, got: %s", errOut.String())
	}
	if !strings.Contains(errOut.String(), filepath.Base(missing)) {
		t.Fatalf("expected error for injected env file path, got: %s", errOut.String())
	}
	if strings.Contains(errOut.String(), filepath.Base(processEnvPath)) {
		t.Fatalf("process PULSAR_ENV_FILE path was used: %s", errOut.String())
	}
}

func TestRunWithNilGetenvIgnoresProcessEnvFile(t *testing.T) {
	root := t.TempDir()
	processEnvPath := filepath.Join(root, "process-env-must-not-be-read.env")
	t.Setenv("PULSAR_ENV_FILE", processEnvPath)

	fr := &fakeRunner{result: engine.RunResult{}}
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut strings.Builder
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)
	if code == 2 {
		t.Fatalf("nil deps.getenv should ignore process PULSAR_ENV_FILE; stderr=%s", errOut.String())
	}
}

func TestDefaultDepsObservesProcessEnvFile(t *testing.T) {
	root := t.TempDir()
	envFilePath := filepath.Join(root, "controlled.env")
	t.Setenv("PULSAR_ENV_FILE", envFilePath)

	if got := resolveAutoEnvFile(defaultDeps()); got != envFilePath {
		t.Fatalf("resolveAutoEnvFile(defaultDeps()) = %q, want %q", got, envFilePath)
	}
}

// TestEnvFileLoadedBeforeRun verifies that --env-file values are visible to
// buildExtraEnv and that the --env-file flag is stripped from subcommand args.
func TestEnvFileLoadedBeforeRun(t *testing.T) {
	t.Setenv("GITHUB_TOKEN", "")

	root := t.TempDir()
	envFile := filepath.Join(root, "test.env")
	if err := os.WriteFile(envFile, []byte("GITHUB_TOKEN=ghp_fake_from_file\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		EnvByMode: map[string][]provider.EnvVar{
			"anonymous": {{Key: "GITHUB_TOKEN"}},
		},
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
		getenv:    getenvFromMap(map[string]string{"GITHUB_TOKEN": ""}),
	}

	var out, errOut strings.Builder
	code := runWithDeps([]string{"run", "--env-file", envFile, "--repo-root", root}, &out, &errOut, d)
	// Must not be exit 2 (flag parse error) since --env-file should be stripped.
	if code == 2 {
		t.Fatalf("exit code 2 (flag error): stderr=%s", errOut.String())
	}
	if !extraEnvContains(fr.spec.ExtraEnv, "GITHUB_TOKEN=ghp_fake_from_file") {
		t.Errorf("ExtraEnv missing GITHUB_TOKEN=ghp_fake_from_file; got: %v", fr.spec.ExtraEnv)
	}
	// --env-file must not appear in errOut as an unknown flag.
	if strings.Contains(errOut.String(), "flag provided but not defined") {
		t.Errorf("--env-file was not stripped; errOut: %s", errOut.String())
	}
}

// TestEnvFileEqualSignForm verifies --env-file=<path> syntax works.
func TestEnvFileEqualSignForm(t *testing.T) {
	t.Setenv("GITHUB_TOKEN", "")

	root := t.TempDir()
	envFile := filepath.Join(root, "eq.env")
	if err := os.WriteFile(envFile, []byte("GITHUB_TOKEN=ghp_eq_form\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		EnvByMode: map[string][]provider.EnvVar{
			"anonymous": {{Key: "GITHUB_TOKEN"}},
		},
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
		getenv:    getenvFromMap(map[string]string{"GITHUB_TOKEN": ""}),
	}

	var out, errOut strings.Builder
	arg := "--env-file=" + envFile
	code := runWithDeps([]string{"run", arg, "--repo-root", root}, &out, &errOut, d)
	if code == 2 {
		t.Fatalf("exit code 2: stderr=%s", errOut.String())
	}
	if !extraEnvContains(fr.spec.ExtraEnv, "GITHUB_TOKEN=ghp_eq_form") {
		t.Errorf("ExtraEnv missing GITHUB_TOKEN=ghp_eq_form; got: %v", fr.spec.ExtraEnv)
	}
}

// TestEnvFileMissingExplicitIsError verifies exit 2 when --env-file points to
// a nonexistent file.
func TestEnvFileMissingExplicitIsError(t *testing.T) {
	root := t.TempDir()
	missing := filepath.Join(root, "missing-explicit.env")
	d := deps{
		provider:      &fakeprovider.Fake{},
		list:          stubList(nil),
		cwd:           func() (string, error) { return root, nil },
		userConfigDir: func() (string, error) { return filepath.Join(root, "config"), nil },
		userHomeDir:   func() (string, error) { return filepath.Join(root, "home"), nil },
	}
	var out, errOut strings.Builder
	code := runWithDeps([]string{"run", "--env-file", missing, "--repo-root", root}, &out, &errOut, d)
	if code != 2 {
		t.Errorf("expected exit 2 for missing explicit --env-file, got %d", code)
	}
	if !strings.Contains(errOut.String(), "env-file:") {
		t.Fatalf("expected env-file error, got: %s", errOut.String())
	}
}

// TestEnvFileDanglingFlagIsError verifies exit 2 when --env-file is the last
// argument with no following path.
func TestEnvFileDanglingFlagIsError(t *testing.T) {
	d := deps{
		provider: &fakeprovider.Fake{},
		list:     stubList(nil),
		cwd:      func() (string, error) { return t.TempDir(), nil },
	}
	var out, errOut strings.Builder
	code := runWithDeps([]string{"run", "--env-file"}, &out, &errOut, d)
	if code != 2 {
		t.Errorf("expected exit 2 for dangling --env-file, got %d", code)
	}
	if !strings.Contains(errOut.String(), "--env-file requires a path") {
		t.Errorf("expected error message in stderr, got: %s", errOut.String())
	}
}

// TestEnvFileModeScoped verifies mode-scoped keys resolve correctly through
// the full runWithDeps path when --mode is given.
func TestEnvFileModeScoped(t *testing.T) {
	t.Setenv("GITHUB_OWNER", "")

	root := t.TempDir()
	envFile := filepath.Join(root, "multi.env")
	content := "PULSAR_INDIVIDUAL_GITHUB_OWNER=alice\nPULSAR_ORGANIZATION_GITHUB_OWNER=acme-test\n"
	if err := os.WriteFile(envFile, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		mode      string
		wantOwner string
	}{
		{"individual", "alice"},
		{"organization", "acme-test"},
	} {
		t.Run(tc.mode, func(t *testing.T) {
			t.Setenv("GITHUB_OWNER", "")

			fr := &fakeRunner{result: engine.RunResult{}}
			pf := &fakeprovider.Fake{
				NameVal:  "test",
				Packages: []string{"./..."},
				Pattern:  "^TestAcc",
				ModesVal: []provider.Mode{
					{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"},
					{Name: "team"}, {Name: "enterprise"},
				},
				EnvByMode: map[string][]provider.EnvVar{
					tc.mode: {{Key: "GITHUB_OWNER"}},
				},
				RequirementsFn: allowAllRequirements,
				PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
					return provider.PreflightReport{Mode: m}
				},
			}
			d := deps{
				provider:  pf,
				newRunner: func(_ io.Writer) testRunner { return fr },
				list:      stubList([]string{"TestAccA"}),
				cwd:       func() (string, error) { return root, nil },
				getenv:    getenvFromMap(map[string]string{"GITHUB_OWNER": ""}),
			}

			var out, errOut strings.Builder
			code := runWithDeps([]string{
				"run", "--env-file", envFile, "--mode", tc.mode, "--repo-root", root,
			}, &out, &errOut, d)
			if code == 2 {
				t.Fatalf("exit code 2: stderr=%s", errOut.String())
			}
			want := "GITHUB_OWNER=" + tc.wantOwner
			if !extraEnvContains(fr.spec.ExtraEnv, want) {
				t.Errorf("ExtraEnv missing %s; got: %v", want, fr.spec.ExtraEnv)
			}
		})
	}
}

func TestEnvFileScopingPrefersShellModeOverFileMode(t *testing.T) {
	// When both the shell and the file set GH_TEST_AUTH_MODE, the run resolves
	// the shell mode because resolveEnvFile never overrides a non-empty shell
	// value. The PULSAR_<mode>_ scoping must follow the shell mode too, otherwise
	// the file would promote keys for a mode the run does not actually use.
	t.Setenv("GH_TEST_AUTH_MODE", "organization")
	t.Setenv("GITHUB_OWNER", "")

	root := t.TempDir()
	envFile := filepath.Join(root, "multi.env")
	content := "GH_TEST_AUTH_MODE=team\n" +
		"PULSAR_TEAM_GITHUB_OWNER=teamowner\n" +
		"PULSAR_ORGANIZATION_GITHUB_OWNER=orgowner\n"
	if err := os.WriteFile(envFile, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		ModesVal: []provider.Mode{
			{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"},
			{Name: "team"}, {Name: "enterprise"},
		},
		EnvByMode: map[string][]provider.EnvVar{
			"organization": {{Key: "GITHUB_OWNER"}},
		},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
		getenv: getenvFromMap(map[string]string{
			"GH_TEST_AUTH_MODE": "organization",
			"GITHUB_OWNER":      "",
		}),
	}

	var out, errOut strings.Builder
	// No --mode flag: the run mode comes from the shell GH_TEST_AUTH_MODE.
	code := runWithDeps([]string{
		"run", "--env-file", envFile, "--repo-root", root,
	}, &out, &errOut, d)
	if code == 2 {
		t.Fatalf("exit code 2: stderr=%s", errOut.String())
	}
	if !extraEnvContains(fr.spec.ExtraEnv, "GITHUB_OWNER=orgowner") {
		t.Errorf("scoping should follow shell mode organization (GITHUB_OWNER=orgowner); got: %v", fr.spec.ExtraEnv)
	}
	if extraEnvContains(fr.spec.ExtraEnv, "GITHUB_OWNER=teamowner") {
		t.Errorf("scoping wrongly followed file mode team (GITHUB_OWNER=teamowner); got: %v", fr.spec.ExtraEnv)
	}
}

func TestRetryEnvFileScopingUsesPersistedPlanModeWithoutFlag(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, ".git"), []byte("gitdir: /tmp/fake\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	cwd := filepath.Join(root, "nested", "repo")
	if err := os.MkdirAll(cwd, 0o755); err != nil {
		t.Fatal(err)
	}
	state := engine.State{
		Provider: "test",
		Mode:     "organization",
		Plan:     planWithEligible("organization", []string{"TestAccA"}),
		Results: []engine.PersistResult{{
			Test:    "TestAccA",
			Status:  "fail",
			Package: "./github",
		}},
	}
	if err := state.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	envFile := filepath.Join(root, "retry.env")
	if err := os.WriteFile(envFile, []byte(
		"GITHUB_OWNER=plain-owner\nPULSAR_ORGANIZATION_GITHUB_OWNER=org-owner\n",
	), 0o600); err != nil {
		t.Fatal(err)
	}

	getenv, setenv, values := mapEnv(nil)
	fr := &fakeRunner{result: engine.RunResult{}}
	d := deps{
		provider: &fakeprovider.Fake{
			Packages: []string{"./..."},
			Pattern:  "^TestAcc",
			ModesVal: testModes,
			EnvByMode: map[string][]provider.EnvVar{
				"organization": {{Key: "GITHUB_OWNER"}},
			},
		},
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return cwd, nil },
		getenv:    getenv,
		setenv:    setenv,
	}

	var out, errOut strings.Builder
	code := runWithDeps([]string{"retry", "--failed", "--env-file", envFile}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if (*values)["GITHUB_OWNER"] != "org-owner" {
		t.Fatalf("GITHUB_OWNER = %q, want organization-scoped env value", (*values)["GITHUB_OWNER"])
	}
	if !extraEnvContains(fr.spec.ExtraEnv, "GITHUB_OWNER=org-owner") {
		t.Fatalf("ExtraEnv = %v, want GITHUB_OWNER=org-owner", fr.spec.ExtraEnv)
	}
	if extraEnvContains(fr.spec.ExtraEnv, "GITHUB_OWNER=plain-owner") {
		t.Fatalf("ExtraEnv = %v, must not keep plain-owner when persisted mode is organization", fr.spec.ExtraEnv)
	}
}

func TestSweepEnvFileScopingUsesPersistedCleanupModeFromRepoRoot(t *testing.T) {
	root := t.TempDir()
	wantResources := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-repo", URL: "https://example.test/repo"},
	}
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
		Orphans: &engine.OrphanAccounting{
			Mode:          "organization",
			New:           wantResources,
			CleanupStatus: engine.CleanupComplete,
		},
	})

	envFile := filepath.Join(root, "sweep.env")
	if err := os.WriteFile(envFile, []byte(
		"GITHUB_OWNER=plain-owner\nPULSAR_ORGANIZATION_GITHUB_OWNER=org-owner\n",
	), 0o600); err != nil {
		t.Fatal(err)
	}

	getenv, setenv, values := mapEnv(nil)
	sweepCalls := 0
	d := deps{
		provider: &fakeprovider.Fake{
			ModesVal: testModes,
			SweepFn: func(_ context.Context, mode string, opts provider.SweepOpts) error {
				sweepCalls++
				if mode != "organization" {
					t.Fatalf("Sweep mode = %q, want organization", mode)
				}
				if (*values)["GITHUB_OWNER"] != "org-owner" {
					t.Fatalf("GITHUB_OWNER = %q at sweep time, want org-owner", (*values)["GITHUB_OWNER"])
				}
				if !reflect.DeepEqual(opts.Resources, wantResources) {
					t.Fatalf("Sweep resources = %#v, want %#v", opts.Resources, wantResources)
				}
				return nil
			},
		},
		cwd:    func() (string, error) { return t.TempDir(), nil },
		getenv: getenv,
		setenv: setenv,
	}

	var out, errOut strings.Builder
	code := runWithDeps([]string{
		"sweep", "--env-file", envFile, "--repo-root", root, "--run-delta", "--confirm",
	}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if sweepCalls != 1 {
		t.Fatalf("Sweep calls = %d, want 1", sweepCalls)
	}
	persisted, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if persisted.Orphans == nil || !reflect.DeepEqual(persisted.Orphans.New, wantResources) {
		t.Fatalf("persisted Orphans.New = %#v, want %#v", persisted.Orphans, wantResources)
	}
}
