package github

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"maps"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	goGithub "github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestOrganizationSettingsHTTPLogging(t *testing.T) {
	t.Parallel()

	resource := resourceGithubOrganizationSettings()
	for name, callback := range map[string]func(context.Context, *schema.ResourceData, any) diag.Diagnostics{
		"create": resource.CreateContext,
		"update": resource.UpdateContext,
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			config := map[string]any{
				"billing_email":                                                "billing@example.com",
				"name":                                                         "Some Fantastic",
				"has_organization_projects":                                    true,
				"has_repository_projects":                                      true,
				"default_repository_permission":                                "read",
				"members_can_create_repositories":                              true,
				"members_can_create_private_repositories":                      true,
				"members_can_create_public_repositories":                       true,
				"members_can_create_internal_repositories":                     true,
				"members_can_create_pages":                                     true,
				"members_can_create_public_pages":                              true,
				"members_can_create_private_pages":                             true,
				"members_can_fork_private_repositories":                        true,
				"web_commit_signoff_required":                                  true,
				"advanced_security_enabled_for_new_repositories":               true,
				"dependabot_alerts_enabled_for_new_repositories":               true,
				"dependabot_security_updates_enabled_for_new_repositories":     true,
				"dependency_graph_enabled_for_new_repositories":                true,
				"secret_scanning_enabled_for_new_repositories":                 true,
				"secret_scanning_push_protection_enabled_for_new_repositories": true,
			}
			wantSettings := maps.Clone(config)

			response := maps.Clone(wantSettings)
			response["id"] = 12345
			response["login"] = "some-fantastic"
			response["plan"] = map[string]any{"name": "enterprise"}
			responseBody, err := json.Marshal(response)
			if err != nil {
				t.Fatal(err)
			}
			var output bytes.Buffer
			ctx := loggingTestContext(t, &output)
			var requests []string
			client, err := goGithub.NewClient(goGithub.WithTransport(logging.NewLoggingHTTPTransport(&mockRoundTripper{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					assertLoggingRequestContext(t, req, ctx)
					if req.URL.Host != "api.github.com" || req.URL.Path != "/orgs/some-fantastic" {
						t.Fatalf("unexpected API URL: %s", req.URL)
					}
					requests = append(requests, req.Method+" "+req.URL.Path)
					if req.Method == http.MethodPatch {
						var settings map[string]any
						if err := json.NewDecoder(req.Body).Decode(&settings); err != nil {
							t.Fatal(err)
						}
						if diff := cmp.Diff(wantSettings, settings); diff != "" {
							t.Fatalf("unexpected settings payload (-want +got):\n%s", diff)
						}
					}
					return &http.Response{
						StatusCode:    http.StatusOK,
						Status:        "200 OK",
						Proto:         "HTTP/1.1",
						ProtoMajor:    1,
						ProtoMinor:    1,
						Header:        http.Header{"Content-Type": {"application/json"}},
						ContentLength: int64(len(responseBody)),
						Body:          io.NopCloser(bytes.NewReader(responseBody)),
						Request:       req,
					}, nil
				},
			})))
			if err != nil {
				t.Fatal(err)
			}
			data := schema.TestResourceDataRaw(t, resource.Schema, config)
			if name == "update" {
				data.SetId("12345")
			}
			meta := &Owner{name: "some-fantastic", IsOrganization: true, v3client: client}
			if diags := callback(ctx, data, meta); diags.HasError() {
				t.Fatalf("organization settings operation failed: %v", diags)
			}
			if data.Id() != "12345" || data.Get("name") != "Some Fantastic" {
				t.Fatal("organization settings operation did not populate state")
			}
			wantRequests := []string{"GET /orgs/some-fantastic", "PATCH /orgs/some-fantastic", "GET /orgs/some-fantastic"}
			if diff := cmp.Diff(wantRequests, requests); diff != "" {
				t.Fatalf("unexpected API requests (-want +got):\n%s", diff)
			}
			entries, err := tflogtest.MultilineJSONDecode(&output)
			if err != nil {
				t.Fatal(err)
			}
			if len(entries) != 7 {
				t.Fatalf("expected six HTTP logs and one summary, got %d logs", len(entries))
			}
			for _, entry := range entries {
				if entry["correlation_id"] != "test-operation" {
					t.Fatal("organization settings log lost the operation context")
				}
				for key := range entry {
					if strings.HasPrefix(key, "settings_") {
						t.Fatalf("redundant per-setting log remains: %s", key)
					}
				}
			}
			if findLogEntry(entries, map[string]any{
				"@message":      "Built organization settings",
				"owner":         "some-fantastic",
				"is_enterprise": true,
			}) == nil {
				t.Fatal("organization settings summary log is missing")
			}
			requestLog := findLogEntry(entries, map[string]any{
				"@message":           "Sending HTTP Request",
				"tf_http_req_method": "PATCH",
				"tf_http_req_uri":    "/orgs/some-fantastic",
			})
			if requestLog == nil {
				t.Fatal("organization settings HTTP request log is missing")
			}
			loggedBody, ok := requestLog["tf_http_req_body"].(string)
			if !ok {
				t.Fatal("HTTP request log does not contain a body")
			}
			var loggedSettings map[string]any
			if err := json.Unmarshal([]byte(loggedBody), &loggedSettings); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(wantSettings, loggedSettings); diff != "" {
				t.Fatalf("HTTP log does not contain the complete settings payload (-want +got):\n%s", diff)
			}
		})
	}
}
