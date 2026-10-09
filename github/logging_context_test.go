package github

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	goGithub "github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

type loggingTestContextKey struct{}

func loggingTestContext(t *testing.T, output io.Writer) context.Context {
	t.Helper()

	ctx, cancel := context.WithTimeout(t.Context(), time.Minute)
	t.Cleanup(cancel)
	ctx = context.WithValue(ctx, loggingTestContextKey{}, "test-operation")
	return tflog.SetField(tflogtest.RootLogger(ctx, output), "correlation_id", "test-operation")
}

func assertLoggingRequestContext(t *testing.T, req *http.Request, ctx context.Context) {
	t.Helper()

	if req.Context().Value(loggingTestContextKey{}) != ctx.Value(loggingTestContextKey{}) {
		t.Error("API request lost the operation context value")
	}
	deadline, ok := req.Context().Deadline()
	wantDeadline, _ := ctx.Deadline()
	if !ok || !deadline.Equal(wantDeadline) {
		t.Error("API request lost the operation deadline")
	}
}

func TestCodespacesSecretLifecycleContext(t *testing.T) {
	t.Parallel()

	const secretPath = "/repos/owner/repo/codespaces/secrets/EXAMPLE"
	tests := map[string][]string{
		"create": {
			"GET /repos/owner/repo/codespaces/secrets/public-key",
			"PUT " + secretPath,
			"GET " + secretPath,
		},
		"import": {"GET " + secretPath},
		"delete": {"DELETE " + secretPath},
	}
	for name, wantRequests := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			const ciphertext = "test-sensitive-ciphertext"
			var output bytes.Buffer
			ctx := loggingTestContext(t, &output)
			var requests []string
			client, err := goGithub.NewClient(goGithub.WithTransport(&mockRoundTripper{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					assertLoggingRequestContext(t, req, ctx)
					requests = append(requests, req.Method+" "+req.URL.Path)
					status := http.StatusOK
					body := `{"name":"EXAMPLE","created_at":"2026-01-01T00:00:00Z","updated_at":"2026-01-02T00:00:00Z"}`
					switch {
					case req.URL.Path == "/repos/owner/repo/codespaces/secrets/public-key":
						body = `{"key_id":"test-key","key":"test-public-key"}`
					case req.Method == http.MethodPut:
						var payload goGithub.EncryptedSecret
						if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
							t.Fatal(err)
						}
						if payload.KeyID != "test-key" || payload.EncryptedValue != ciphertext {
							t.Fatal("create request lost the key ID or encrypted value")
						}
						status = http.StatusCreated
						body = ""
					case req.Method == http.MethodDelete:
						if req.Context().Value(ctxId) != "repo:EXAMPLE" {
							t.Error("delete request lost its resource ID")
						}
						status = http.StatusNoContent
						body = ""
					}
					return &http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader(body)), Request: req}, nil
				},
			}))
			if err != nil {
				t.Fatal(err)
			}
			resource := resourceGithubCodespacesSecret()
			data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
				"repository":      "repo",
				"secret_name":     "EXAMPLE",
				"encrypted_value": ciphertext,
			})
			meta := &Owner{name: "owner", v3client: client}
			switch name {
			case "create":
				if diags := resource.CreateContext(ctx, data, meta); diags.HasError() {
					t.Fatalf("create failed: %v", diags)
				}
				if data.Get("encrypted_value") != ciphertext {
					t.Fatal("create did not preserve the encrypted value")
				}
			case "import":
				data = schema.TestResourceDataRaw(t, resource.Schema, nil)
				data.SetId("repo/EXAMPLE")
				imported, err := resource.Importer.StateContext(ctx, data, meta)
				if err != nil {
					t.Fatal(err)
				}
				if len(imported) != 1 || imported[0] != data || data.Get("repository") != "repo" || data.Get("secret_name") != "EXAMPLE" {
					t.Fatal("import did not populate the resource identity")
				}
				if data.Get("encrypted_value") != "" || data.Get("plaintext_value") != "" {
					t.Fatal("import populated a secret value")
				}
			case "delete":
				data.SetId("repo:EXAMPLE")
				if diags := resource.DeleteContext(ctx, data, meta); diags.HasError() {
					t.Fatalf("delete failed: %v", diags)
				}
				entries, err := tflogtest.MultilineJSONDecode(&output)
				if err != nil {
					t.Fatal(err)
				}
				expected := map[string]any{
					"@level":         "debug",
					"@message":       "Deleting secret",
					"correlation_id": "test-operation",
					"secret_id":      "repo:EXAMPLE",
					"owner":          "owner",
					"repository":     "repo",
					"secret_name":    "EXAMPLE",
				}
				if findLogEntry(entries, expected) == nil {
					t.Fatalf("delete log is missing expected fields %v; got %v", expected, entries)
				}
			}
			if diff := cmp.Diff(wantRequests, requests); diff != "" {
				t.Fatalf("unexpected API requests (-want +got):\n%s", diff)
			}
			if data.Id() != "repo:EXAMPLE" {
				t.Fatalf("unexpected resource ID: %q", data.Id())
			}
			if name != "delete" && (data.Get("created_at") == "" || data.Get("updated_at") == "") {
				t.Fatal("secret timestamps were not populated")
			}
			if strings.Contains(output.String(), ciphertext) {
				t.Fatal("operation logged the secret value")
			}
		})
	}
}

func TestCodespacesSecretDriftStructuredLogging(t *testing.T) {
	t.Parallel()

	for _, drifted := range []bool{false, true} {
		name := "unchanged"
		if drifted {
			name = "externally updated"
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			const plaintext = "test-sensitive-plaintext"
			var output bytes.Buffer
			ctx := loggingTestContext(t, &output)
			updatedAt := time.Date(2026, time.January, 2, 0, 0, 0, 0, time.UTC).String()
			client, err := goGithub.NewClient(goGithub.WithTransport(&mockRoundTripper{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					assertLoggingRequestContext(t, req, ctx)
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"name":"EXAMPLE","created_at":"2026-01-01T00:00:00Z","updated_at":"2026-01-02T00:00:00Z"}`)),
						Request:    req,
					}, nil
				},
			}))
			if err != nil {
				t.Fatal(err)
			}
			if drifted {
				updatedAt = time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC).String()
			}
			resource := resourceGithubCodespacesSecret()
			data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
				"repository":      "repo",
				"secret_name":     "EXAMPLE",
				"plaintext_value": plaintext,
				"updated_at":      updatedAt,
			})
			data.SetId("repo:EXAMPLE")
			if diags := resource.ReadContext(ctx, data, &Owner{name: "owner", v3client: client}); diags.HasError() {
				t.Fatalf("read failed: %v", diags)
			}
			if data.Get("plaintext_value") != plaintext || strings.Contains(output.String(), plaintext) {
				t.Fatal("read must preserve the secret without logging it")
			}
			if !drifted {
				if data.Id() != "repo:EXAMPLE" || output.Len() != 0 {
					t.Fatal("unchanged secret must retain its ID without a drift warning")
				}
				return
			}
			if data.Id() != "" {
				t.Fatal("drifted secret was not removed from state")
			}
			entries, err := tflogtest.MultilineJSONDecode(&output)
			if err != nil {
				t.Fatal(err)
			}
			expected := map[string]any{
				"@level":         "warn",
				"@message":       "Secret has been externally updated in GitHub",
				"correlation_id": "test-operation",
				"secret_id":      "repo:EXAMPLE",
				"owner":          "owner",
				"repository":     "repo",
				"secret_name":    "EXAMPLE",
			}
			if findLogEntry(entries, expected) == nil {
				t.Fatalf("drift log is missing expected fields %v; got %v", expected, entries)
			}
		})
	}
}

func TestGraphQLActorStructuredLogging(t *testing.T) {
	t.Parallel()

	var output bytes.Buffer
	ctx := loggingTestContext(t, &output)
	requests := 0
	client := githubv4.NewClient(&http.Client{Transport: &mockRoundTripper{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			assertLoggingRequestContext(t, req, ctx)
			requests++
			var payload struct {
				Variables map[string]any `json:"variables"`
			}
			if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
				t.Fatal(err)
			}
			var body string
			switch {
			case payload.Variables["slug"] == "team" && payload.Variables["organization"] == "owner":
				body = `{"data":{"organization":{"team":{"id":"TEAM_node"}}}}`
			case payload.Variables["user"] == "octocat":
				body = `{"data":{"user":{"id":"USER_node"}}}`
			default:
				t.Fatalf("unexpected actor query variables: %v", payload.Variables)
			}
			return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body)), Request: req}, nil
		},
	}})
	actors := []string{"owner/team", "/octocat", "EXISTING_node"}
	wantIDs := []string{"TEAM_node", "USER_node", "EXISTING_node"}
	ids, err := getActorIds(ctx, actors, &Owner{name: "owner", v4client: client})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantIDs, ids); diff != "" {
		t.Fatalf("unexpected actor IDs (-want +got):\n%s", diff)
	}
	if requests != 2 {
		t.Fatalf("expected only team and user requests, got %d", requests)
	}
	entries, err := tflogtest.MultilineJSONDecode(&output)
	if err != nil {
		t.Fatal(err)
	}
	for i, actor := range actors {
		messages := []string{"Retrieved node ID for actor"}
		switch i {
		case 0:
			messages = append(messages, "Retrieved node ID for team")
		case 1:
			messages = append(messages, "Retrieved node ID for user")
		}
		for _, message := range messages {
			expected := map[string]any{
				"@level":         "debug",
				"@message":       message,
				"@module":        "provider",
				"correlation_id": "test-operation",
				"actor":          actor,
				"node_id":        wantIDs[i],
			}
			if findLogEntry(entries, expected) == nil {
				t.Fatalf("actor log is missing expected fields %v; got %v", expected, entries)
			}
		}
	}
}

func TestGraphQLBranchProtectionPaginationContext(t *testing.T) {
	t.Parallel()

	var output bytes.Buffer
	ctx := loggingTestContext(t, &output)
	requests := 0
	client := githubv4.NewClient(&http.Client{Transport: &mockRoundTripper{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			assertLoggingRequestContext(t, req, ctx)
			requests++
			var payload struct {
				Variables map[string]any `json:"variables"`
			}
			if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
				t.Fatal(err)
			}
			if payload.Variables["id"] != "REPO_node" || payload.Variables["first"] != float64(1) {
				t.Fatalf("unexpected branch protection query variables: %v", payload.Variables)
			}
			var body string
			switch requests {
			case 1:
				if payload.Variables["cursor"] != nil {
					t.Fatal("first page must not have a cursor")
				}
				body = `{"data":{"node":{"id":"REPO_node","branchProtectionRules":{"nodes":[{"id":"RULE_other","pattern":"other"}],"pageInfo":{"hasNextPage":true,"endCursor":"next-page"}}}}}`
			case 2:
				if payload.Variables["cursor"] != "next-page" {
					t.Fatal("second page lost the pagination cursor")
				}
				body = `{"data":{"node":{"id":"REPO_node","branchProtectionRules":{"nodes":[{"id":"RULE_main","pattern":"main"}],"pageInfo":{"hasNextPage":false,"endCursor":"last-page"}}}}}`
			default:
				t.Fatal("unexpected extra branch protection request")
			}
			return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body)), Request: req}, nil
		},
	}})
	id, err := getBranchProtectionID(ctx, "REPO_node", "main", &Owner{v4client: client, maxPerPage: 1})
	if err != nil {
		t.Fatal(err)
	}
	if id != "RULE_main" || requests != 2 {
		t.Fatalf("expected rule from second page; got ID %v after %d requests", id, requests)
	}
}
