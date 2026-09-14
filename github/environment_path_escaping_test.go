package github

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Test_environmentNamePathEscaping guards against escaping an environment name
// on both sides of the go-github call. go-github escapes it for us, so a
// provider that escapes first turns "a b" into "a%2520b" and requests a path
// that does not exist.
func Test_environmentNamePathEscaping(t *testing.T) {
	t.Parallel()

	// Escaping this name is not a no-op, so escaping it twice is visible.
	const envName = "review app/1"

	for _, tt := range []struct {
		name     string
		resource func() *schema.Resource
		read     schema.ReadContextFunc
		state    map[string]any
		wantPath string
	}{
		{
			name:     "github_repository_environment",
			resource: resourceGithubRepositoryEnvironment,
			read:     resourceGithubRepositoryEnvironmentRead,
			state: map[string]any{
				"repository":  "test-repo",
				"environment": envName,
			},
			wantPath: "/repos/test-owner/test-repo/environments/" + url.PathEscape(envName),
		},
		{
			name:     "github_actions_environment_variable",
			resource: resourceGithubActionsEnvironmentVariable,
			read:     resourceGithubActionsEnvironmentVariableRead,
			state: map[string]any{
				"repository":    "test-repo",
				"environment":   envName,
				"variable_name": "TEST_VARIABLE",
			},
			wantPath: "/repos/test-owner/test-repo/environments/" + url.PathEscape(envName) + "/variables/TEST_VARIABLE",
		},
		{
			name:     "github_actions_environment_secret",
			resource: resourceGithubActionsEnvironmentSecret,
			read:     resourceGithubActionsEnvironmentSecretRead,
			state: map[string]any{
				"repository":  "test-repo",
				"environment": envName,
				"secret_name": "TEST_SECRET",
			},
			wantPath: "/repos/test-owner/test-repo/environments/" + url.PathEscape(envName) + "/secrets/TEST_SECRET",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Answering 404 lets each read take its "gone from GitHub" path,
			// so the test only depends on the request, not on response shape.
			var gotPath string
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotPath = r.URL.EscapedPath()
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"message": "Not Found"}`))
			}))
			t.Cleanup(ts.Close)

			baseURL := ts.URL + "/"

			client, err := github.NewClient(github.WithURLs(&baseURL, nil))
			if err != nil {
				t.Fatalf("failed to create test client: %v", err)
			}

			d := schema.TestResourceDataRaw(t, tt.resource().Schema, tt.state)
			d.SetId("placeholder")

			if diags := tt.read(context.Background(), d, &Owner{name: "test-owner", v3client: client}); diags.HasError() {
				t.Fatalf("read returned an error: %v", diags)
			}

			if gotPath != tt.wantPath {
				t.Errorf("request path was %q, want %q", gotPath, tt.wantPath)
			}
		})
	}
}
