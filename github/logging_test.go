package github

import (
	"bytes"
	"context"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"maps"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	goGithub "github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func TestProviderStructuredLogging(t *testing.T) {
	t.Parallel()

	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range files {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}

		t.Run(path, func(t *testing.T) {
			t.Parallel()

			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, path, nil, 0)
			if err != nil {
				t.Fatal(err)
			}

			for _, imp := range file.Imports {
				if imp.Path.Value == `"log"` {
					t.Error("provider logging must use tflog instead of the standard log package")
				}
			}

			ast.Inspect(file, func(node ast.Node) bool {
				call, ok := node.(*ast.CallExpr)
				if !ok {
					return true
				}
				selector, ok := call.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				pkg, ok := selector.X.(*ast.Ident)
				if !ok || pkg.Name != "tflog" {
					return true
				}
				switch selector.Sel.Name {
				case "Trace", "Debug", "Info", "Warn", "Error":
				default:
					return true
				}

				position := fset.Position(call.Pos())
				if len(call.Args) < 2 {
					t.Errorf("%s: logging requires a context and message", position)
					return true
				}
				message, ok := call.Args[1].(*ast.BasicLit)
				if !ok || message.Kind != token.STRING {
					t.Errorf("%s: use a static message and structured fields", position)
					return true
				}
				text, err := strconv.Unquote(message.Value)
				if err != nil {
					t.Fatal(err)
				}
				if strings.HasPrefix(text, "[") {
					t.Errorf("%s: the logger supplies the level, do not embed it in the message", position)
				}
				if len(call.Args) > 2 && strings.Contains(path, "migration") {
					ast.Inspect(call.Args[2], func(node ast.Node) bool {
						value := node
						if field, ok := node.(*ast.KeyValueExpr); ok {
							if key, ok := field.Key.(*ast.BasicLit); ok && key.Value == `"field_count"` {
								t.Errorf("%s: log the fields changed by a migration instead of a total field count", position)
							}
							value = field.Value
						}
						if ident, ok := value.(*ast.Ident); ok && (node == call.Args[2] || value != node) {
							switch ident.Name {
							case "rawState", "state", "migratedState":
								t.Errorf("%s: do not log complete migration state", position)
							}
						}
						return true
					})
				}
				return true
			})
		})
	}
}

func TestResourceStructuredLogging(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		resource *schema.Resource
		config   map[string]any
		id       string
		level    string
		message  string
		fields   map[string]any
	}{
		"Codespaces secret removed": {
			resource: resourceGithubCodespacesSecret(),
			config:   map[string]any{"repository": "repo", "secret_name": "EXAMPLE"},
			id:       "repo:EXAMPLE",
			level:    "warn",
			message:  "Removing Codespaces secret from state because it no longer exists in GitHub",
			fields:   map[string]any{"secret_id": "repo:EXAMPLE", "owner": "owner", "repository": "repo", "secret_name": "EXAMPLE"},
		},
		"branch data source missing": {
			resource: dataSourceGithubBranch(),
			config:   map[string]any{"repository": "repo", "branch": "main"},
			id:       "repo:main",
			level:    "debug",
			message:  "GitHub branch not found",
			fields:   map[string]any{"owner": "owner", "repository": "repo", "branch": "refs/heads/main"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var output bytes.Buffer
			ctx := tflogtest.RootLogger(t.Context(), &output)
			ctx = tflog.SetField(ctx, "correlation_id", "test-operation")
			ctx = tflog.SetField(ctx, "additional_context", "extra-field")
			tflog.Trace(ctx, "Starting test operation")
			client, err := goGithub.NewClient(goGithub.WithTransport(&mockRoundTripper{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusNotFound, Body: http.NoBody, Request: req}, nil
				},
			}))
			if err != nil {
				t.Fatal(err)
			}
			data := schema.TestResourceDataRaw(t, tc.resource.Schema, tc.config)
			data.SetId(tc.id)

			if diags := tc.resource.ReadContext(ctx, data, &Owner{name: "owner", v3client: client}); diags.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}
			if data.Id() != "" {
				t.Fatal("missing resource was not removed from state")
			}

			entries, err := tflogtest.MultilineJSONDecode(&output)
			if err != nil {
				t.Fatal(err)
			}
			expected := map[string]any{
				"@level":         tc.level,
				"@message":       tc.message,
				"@module":        "provider",
				"correlation_id": "test-operation",
			}
			maps.Copy(expected, tc.fields)
			if findLogEntry(entries, expected) == nil {
				t.Fatalf("no log entry contains expected fields %v; got %v", expected, entries)
			}
		})
	}
}

func TestCodespacesSecretContext(t *testing.T) {
	t.Parallel()

	resource := resourceGithubCodespacesSecret()
	for name, callback := range map[string]func(context.Context, *schema.ResourceData, any) diag.Diagnostics{
		"create": resource.CreateContext,
		"read":   resource.ReadContext,
		"delete": resource.DeleteContext,
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			type contextKey struct{}
			ctx, cancel := context.WithTimeout(t.Context(), time.Minute)
			defer cancel()
			ctx = context.WithValue(ctx, contextKey{}, "operation")
			wantErr := errors.New("stop request")
			requests := 0
			client, err := goGithub.NewClient(goGithub.WithTransport(&mockRoundTripper{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					requests++
					if req.Context().Value(contextKey{}) != "operation" {
						t.Error("request lost the callback context values")
					}
					deadline, ok := req.Context().Deadline()
					wantDeadline, _ := ctx.Deadline()
					if !ok || !deadline.Equal(wantDeadline) {
						t.Error("request lost the callback deadline")
					}
					if name == "delete" && req.Context().Value(ctxId) != "repo:EXAMPLE" {
						t.Error("request lost its resource ID")
					}
					return nil, wantErr
				},
			}))
			if err != nil {
				t.Fatal(err)
			}
			data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
				"repository":      "repo",
				"secret_name":     "EXAMPLE",
				"encrypted_value": "test-ciphertext",
			})
			data.SetId("repo:EXAMPLE")

			diags := callback(ctx, data, &Owner{name: "owner", v3client: client})
			if !diags.HasError() || !strings.Contains(diags[0].Summary, wantErr.Error()) {
				t.Fatalf("API error was not returned as diagnostics: %v", diags)
			}
			if requests != 1 {
				t.Fatalf("expected one API request, got %d", requests)
			}
		})
	}
}

func TestResourceErrorStructuredLogging(t *testing.T) {
	t.Parallel()

	for _, status := range []int{http.StatusNotModified, http.StatusNotFound, http.StatusInternalServerError} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			t.Parallel()

			var output bytes.Buffer
			ctx := tflogtest.RootLogger(t.Context(), &output)
			ctx = tflog.SetField(ctx, "additional_context", "extra-field")
			tflog.Trace(ctx, "Starting test operation")
			data := schema.TestResourceDataRaw(t, resourceGithubRepositoryFile().Schema, nil)
			data.SetId("repo:file")
			apiErr := &goGithub.ErrorResponse{Response: &http.Response{StatusCode: status}}

			outputBefore := output.Len()
			err := deleteResourceOn404AndSwallow304OtherwiseReturnError(ctx, apiErr, data, "repository file", map[string]any{"repository": "repo"})
			if status == http.StatusInternalServerError {
				if !errors.Is(err, apiErr) || output.Len() != outputBefore || data.Id() != "repo:file" {
					t.Fatal("unexpected errors must be preserved without changing state or logging removal")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if (data.Id() == "") != (status == http.StatusNotFound) {
				t.Fatal("only a missing resource should be removed from state")
			}
			entries, err := tflogtest.MultilineJSONDecode(&output)
			if err != nil {
				t.Fatal(err)
			}
			expected := map[string]any{
				"@level":        "info",
				"resource_type": "repository file",
				"resource_id":   "repo:file",
				"repository":    "repo",
			}
			if findLogEntry(entries, expected) == nil {
				t.Fatalf("no log entry contains expected fields %v; got %v", expected, entries)
			}
		})
	}
}

func TestMigrationLoggingOmitsSecrets(t *testing.T) {
	t.Parallel()

	for name, tc := range map[string]struct {
		upgrade schema.StateUpgradeFunc
		state   map[string]any
		before  map[string]any
		after   map[string]any
		logs    int
	}{
		"repository webhook": {
			upgrade: resourceGithubRepositoryWebhookInstanceStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "configuration_map_size": "4", "configuration_content_type": "json", "configuration_insecure_ssl": "0"},
			after:   map[string]any{"repository": "repo", "configuration_block_count": "1", "configuration_content_type": "json", "configuration_insecure_ssl": "0"},
		},
		"organization webhook": {
			upgrade: resourceGithubOrganizationWebhookInstanceStateUpgradeV0,
			before:  map[string]any{"configuration_map_size": "4", "configuration_content_type": "json", "configuration_insecure_ssl": "0"},
			after:   map[string]any{"configuration_block_count": "1", "configuration_content_type": "json", "configuration_insecure_ssl": "0"},
		},
		"actions secret v0": {
			upgrade: resourceGithubActionsSecretStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "secret_name": "EXAMPLE", "destroy_on_drift": nil},
			after:   map[string]any{"repository": "repo", "secret_name": "EXAMPLE", "destroy_on_drift": true},
		},
		"actions secret v1": {
			upgrade: resourceGithubActionsSecretStateUpgradeV1,
			before:  map[string]any{"repository": "repo", "secret_name": "EXAMPLE", "repository_id": nil},
			after:   map[string]any{"repository": "repo", "secret_name": "EXAMPLE", "repository_id": float64(123)},
		},
		"actions environment secret": {
			upgrade: resourceGithubActionsEnvironmentSecretStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "environment": "production", "secret_name": "EXAMPLE", "repository_id": nil, "resource_id": nil},
			after:   map[string]any{"repository": "repo", "environment": "production", "secret_name": "EXAMPLE", "repository_id": float64(123), "resource_id": "repo:production:EXAMPLE"},
		},
		"actions organization secret": {
			upgrade: resourceGithubActionsOrganizationSecretStateUpgradeV0,
			before:  map[string]any{"secret_name": "EXAMPLE", "destroy_on_drift": nil},
			after:   map[string]any{"secret_name": "EXAMPLE", "destroy_on_drift": true},
		},
		"dependabot secret": {
			upgrade: resourceGithubDependabotSecretStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "secret_name": "EXAMPLE", "repository_id": nil},
			after:   map[string]any{"repository": "repo", "secret_name": "EXAMPLE", "repository_id": float64(123)},
		},
		"repository file": {
			upgrade: resourceGithubRepositoryFileStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "repository_id": nil, "file": "config.txt", "branch": nil, "resource_id": nil},
			after:   map[string]any{"repository": "repo", "repository_id": float64(123), "file": "config.txt", "branch": "main", "resource_id": "repo:config.txt:main"},
		},
		"repository": {
			upgrade: resourceGithubRepositoryInstanceStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "branches_count": "1"},
			after:   map[string]any{"repository": "repo", "branches_count": nil},
		},
		"actions variable": {
			upgrade: resourceGithubActionsVariableStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "variable_name": "EXAMPLE", "repository_id": nil},
			after:   map[string]any{"repository": "repo", "variable_name": "EXAMPLE", "repository_id": float64(123)},
		},
		"actions environment variable": {
			upgrade: resourceGithubActionsEnvironmentVariableStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "environment": "production", "variable_name": "EXAMPLE", "repository_id": nil, "resource_id": nil},
			after:   map[string]any{"repository": "repo", "environment": "production", "variable_name": "EXAMPLE", "repository_id": float64(123), "resource_id": "repo:production:EXAMPLE"},
		},
		"repository environment": {
			upgrade: resourceGithubRepositoryEnvironmentStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "environment": "production", "repository_id": nil},
			after:   map[string]any{"repository": "repo", "environment": "production", "repository_id": float64(123)},
		},
		"repository environment deployment policy": {
			upgrade: resourceGithubRepositoryEnvironmentDeploymentPolicyStateUpgradeV0,
			state:   map[string]any{"id": "repo:production:7"},
			before:  map[string]any{"repository": "repo", "repository_id": nil, "environment": "production", "policy_id": nil, "resource_id": "repo:production:7"},
			after:   map[string]any{"repository": "repo", "repository_id": float64(123), "environment": "production", "policy_id": float64(7), "resource_id": "repo:production:7"},
		},
		"repository collaborators v0": {
			upgrade: resourceGithubRepositoryCollaboratorsStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "repository_id": nil, "resource_id": nil},
			after:   map[string]any{"repository": "repo", "repository_id": float64(123), "resource_id": "123"},
		},
		"repository collaborators v1": {
			upgrade: resourceGithubRepositoryCollaboratorsStateUpgradeV1,
			before:  map[string]any{"repository": "repo", "owner_configured": nil},
			after:   map[string]any{"repository": "repo", "owner_configured": false},
		},
		"branch default": {
			upgrade: resourceGithubBranchDefaultStateUpgradeV0,
			state:   map[string]any{"branch": "main"},
			before:  map[string]any{"repository": "repo", "branch": "main", "repository_id": nil},
			after:   map[string]any{"repository": "repo", "branch": "main", "repository_id": float64(123)},
			logs:    4,
		},
		"repository custom property": {
			upgrade: resourceGithubRepositoryCustomPropertyStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "repository_id": nil},
			after:   map[string]any{"repository": "repo", "repository_id": float64(123)},
			logs:    4,
		},
		"release": {
			upgrade: resourceGithubReleaseStateUpgradeV0,
			before:  map[string]any{"repository": "repo", "repository_id": nil},
			after:   map[string]any{"repository": "repo", "repository_id": float64(123)},
			logs:    4,
		},
		"EMU group mapping v1": {
			upgrade: resourceGithubEMUGroupMappingStateUpgradeV1,
			state:   map[string]any{"id": "12:demo-team:34", "team_id": 12, "group_id": 34},
			before:  map[string]any{"resource_id": "12:demo-team:34", "team_id": float64(12), "group_id": float64(34)},
			after:   map[string]any{"resource_id": "34:12", "team_id": float64(12), "group_id": float64(34)},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			const secret = "test-sensitive-value-not-for-logging"
			state := map[string]any{
				"name":                       "repo",
				"repository":                 "repo",
				"environment":                "production",
				"secret_name":                "EXAMPLE",
				"variable_name":              "EXAMPLE",
				"file":                       "config.txt",
				"content":                    secret,
				"value":                      secret,
				"configuration.%":            "4",
				"configuration.content_type": "json",
				"configuration.insecure_ssl": "0",
				"configuration.secret":       secret,
				"configuration.url":          "https://example.com/" + secret,
				"plaintext_value":            secret,
				"encrypted_value":            secret,
				"branches.#":                 "1",
				"branches.0.name":            "main",
				"branches.0.protected":       "false",
			}
			maps.Copy(state, tc.state)
			var output bytes.Buffer
			ctx := loggingTestContext(t, &output)
			client, err := goGithub.NewClient(goGithub.WithTransport(&mockRoundTripper{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					assertLoggingRequestContext(t, req, ctx)
					if req.Method != http.MethodGet || req.URL.Path != "/repos/owner/repo" {
						t.Fatalf("unexpected migration request: %s %s", req.Method, req.URL.Path)
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"id":123,"default_branch":"main"}`)),
						Request:    req,
					}, nil
				},
			}))
			if err != nil {
				t.Fatal(err)
			}
			migratedState, err := tc.upgrade(ctx, state, &Owner{name: "owner", IsOrganization: true, v3client: client})
			if err != nil {
				t.Fatal(err)
			}
			if migratedState["plaintext_value"] != secret || migratedState["encrypted_value"] != secret || migratedState["content"] != secret || migratedState["value"] != secret {
				t.Fatal("migration did not preserve sensitive state values")
			}
			if strings.Contains(output.String(), secret) {
				t.Fatal("migration logged sensitive state values")
			}
			entries, err := tflogtest.MultilineJSONDecode(&output)
			if err != nil {
				t.Fatal(err)
			}
			wantLogs := tc.logs
			if wantLogs == 0 {
				wantLogs = 2
			}
			if len(entries) != wantLogs {
				t.Fatalf("expected %d migration logs, got %d", wantLogs, len(entries))
			}
			level := "debug"
			if name == "EMU group mapping v1" {
				level = "trace"
			}
			for i, fields := range []map[string]any{tc.before, tc.after} {
				expected := map[string]any{
					"@level":         level,
					"@module":        "provider",
					"correlation_id": "test-operation",
				}
				maps.Copy(expected, fields)
				entry := entries[0]
				if i == 1 {
					entry = entries[len(entries)-1]
				}
				if findLogEntry([]map[string]any{entry}, expected) == nil {
					t.Fatalf("migration log is missing expected fields %v; got %v", expected, entry)
				}
			}
		})
	}
}

func TestRateLimitTransportStructuredLogging(t *testing.T) {
	t.Parallel()

	for _, primary := range []bool{false, true} {
		t.Run(strconv.FormatBool(primary), func(t *testing.T) {
			t.Parallel()

			var output bytes.Buffer
			ctx := tflogtest.RootLogger(t.Context(), &output)
			ctx = tflog.SetField(ctx, "correlation_id", "rate-limit-test")
			ctx = tflog.SetField(ctx, "additional_context", "extra-field")
			tflog.Trace(ctx, "Starting test operation")
			requests := 0
			transport := NewRateLimitTransport(&mockRoundTripper{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					requests++
					if requests > 1 {
						return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody, Request: req}, nil
					}
					headers := http.Header{}
					if primary {
						headers.Set("X-RateLimit-Limit", "60")
						headers.Set("X-RateLimit-Remaining", "0")
						headers.Set("X-RateLimit-Reset", strconv.FormatInt(time.Now().Add(-time.Second).Unix(), 10))
					} else {
						headers.Set("Retry-After", "0")
					}
					body := io.NopCloser(strings.NewReader(`{"documentation_url":"https://docs.github.com/rest/using-the-rest-api/rate-limits-for-the-rest-api#about-secondary-rate-limits"}`))
					return &http.Response{StatusCode: http.StatusForbidden, Header: headers, Body: body, Request: req}, nil
				},
			})
			transport.nextRequestDelay = time.Nanosecond
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/repos/owner/repo", nil)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := transport.RoundTrip(req)
			if err != nil {
				t.Fatal(err)
			}
			if err := resp.Body.Close(); err != nil {
				t.Fatal(err)
			}
			if requests != 2 || resp.StatusCode != http.StatusOK {
				t.Fatal("transport did not retry the rate-limited request")
			}
			entries, err := tflogtest.MultilineJSONDecode(&output)
			if err != nil {
				t.Fatal(err)
			}
			expectedDelay := map[string]any{
				"@level":         "debug",
				"@message":       "Sleeping between operations",
				"correlation_id": "rate-limit-test",
				"delay":          time.Nanosecond.String(),
			}
			if findLogEntry(entries, expectedDelay) == nil {
				t.Fatalf("no delay log contains expected fields %v; got %v", expectedDelay, entries)
			}
			for _, entry := range entries {
				if entry["correlation_id"] != "rate-limit-test" {
					t.Fatal("transport lost the request logger's fields")
				}
			}
			expectedRateLimit := map[string]any{
				"@level":         "warn",
				"@message":       "Abuse detection mechanism triggered, sleeping before retrying",
				"correlation_id": "rate-limit-test",
			}
			if primary {
				expectedRateLimit["@message"] = "Rate limit reached, sleeping before retrying"
				expectedRateLimit["rate_limit"] = float64(60)
			}
			rateLimitEntry := findLogEntry(entries, expectedRateLimit)
			if rateLimitEntry == nil {
				t.Fatalf("no rate limit log contains expected fields %v; got %v", expectedRateLimit, entries)
			}
			if _, ok := rateLimitEntry["retry_after"]; !ok {
				t.Fatal("rate limit log is missing the retry delay")
			}
			if primary && rateLimitEntry["reset_at"] == nil {
				t.Fatal("primary rate limit log is missing the reset time")
			}
		})
	}
}

func findLogEntry(entries []map[string]any, expected map[string]any) map[string]any {
	for _, entry := range entries {
		matches := true
		for key, want := range expected {
			got, ok := entry[key]
			if !ok || !cmp.Equal(want, got) {
				matches = false
				break
			}
		}
		if matches {
			return entry
		}
	}
	return nil
}

func TestFindLogEntry(t *testing.T) {
	t.Parallel()

	entry := map[string]any{
		"@level":      "warn",
		"resource_id": "repo",
		"extra":       "additional field",
		"optional":    nil,
	}
	entries := []map[string]any{
		{"@level": "trace", "resource_id": "repo"},
		{"@level": "warn", "resource_id": "other"},
		entry,
	}
	tests := map[string]struct {
		expected map[string]any
		matches  bool
	}{
		"allows extra fields and unrelated entries": {
			expected: map[string]any{"@level": "warn", "resource_id": "repo"},
			matches:  true,
		},
		"rejects incorrect level": {
			expected: map[string]any{"@level": "error", "resource_id": "repo"},
		},
		"rejects incorrect identifier": {
			expected: map[string]any{"@level": "warn", "resource_id": "missing"},
		},
		"requires field presence even when nil": {
			expected: map[string]any{"@level": "warn", "missing": nil},
		},
		"accepts present nil field": {
			expected: map[string]any{"@level": "warn", "optional": nil},
			matches:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := findLogEntry(entries, tc.expected)
			if !tc.matches {
				if got != nil {
					t.Fatalf("unexpected matching entry: %v", got)
				}
				return
			}
			if diff := cmp.Diff(entry, got); diff != "" {
				t.Fatalf("incorrect matching entry (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGraphQLHelperContext(t *testing.T) {
	t.Parallel()

	type contextKey struct{}
	ctx := context.WithValue(t.Context(), contextKey{}, "graphql-operation")
	requests := 0
	client := githubv4.NewClient(&http.Client{Transport: &mockRoundTripper{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			requests++
			if req.Context().Value(contextKey{}) != "graphql-operation" {
				t.Error("GraphQL helper lost the callback context")
			}
			return nil, errors.New("stop request")
		},
	}})
	meta := &Owner{name: "owner", v4client: client}

	if _, err := getActorIds(ctx, []string{"owner/team"}, meta); err == nil {
		t.Fatal("actor lookup did not return the API error")
	}
	if _, err := getRepositoryID(ctx, "repo-name", meta); err == nil {
		t.Fatal("repository lookup did not return the API error")
	}
	if requests != 3 {
		t.Fatalf("expected actor, repository node, and repository name requests, got %d", requests)
	}
}
