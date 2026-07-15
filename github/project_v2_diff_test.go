package github

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	goGithub "github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestProjectV2IdentityDiffPreservesRenames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		resource   *schema.Resource
		state      map[string]string
		config     map[string]any
		requestURI string
		response   func(id int) string
	}{
		{
			name: "project owner", resource: resourceGithubProject(),
			state:      map[string]string{"owner": "old-owner", "owner_type": projectV2OwnerOrganization, "owner_id": "101", "title": "Planning"},
			config:     map[string]any{"owner": "new-owner", "owner_type": projectV2OwnerOrganization, "title": "Planning"},
			requestURI: "/users/new-owner", response: func(id int) string {
				return fmt.Sprintf(`{"id":%d,"login":"new-owner","type":"Organization"}`, id)
			},
		},
		{
			name: "repository", resource: resourceGithubProjectRepository(),
			state:      map[string]string{"project_id": "PVT_1", "repository_owner": "old-owner", "repository": "old-name", "repository_id": "101"},
			config:     map[string]any{"project_id": "PVT_1", "repository_owner": "new-owner", "repository": "new-name"},
			requestURI: "/repos/new-owner/new-name", response: func(id int) string {
				return fmt.Sprintf(`{"id":%d,"name":"new-name","owner":{"login":"new-owner"}}`, id)
			},
		},
		{
			name: "team", resource: resourceGithubTeamProject(),
			state:      map[string]string{"project_id": "PVT_1", "organization": "old-org", "team_slug": "old-slug", "team_id": "101"},
			config:     map[string]any{"project_id": "PVT_1", "organization": "new-org", "team_slug": "new-slug"},
			requestURI: "/orgs/new-org/teams/new-slug", response: func(id int) string {
				return fmt.Sprintf(`{"id":%d,"slug":"new-slug","organization":{"login":"new-org"}}`, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assertProjectV2IdentityDiff(t, test.resource, test.state, test.config, test.requestURI, test.response(101), false)
			assertProjectV2IdentityDiff(t, test.resource, test.state, test.config, test.requestURI, test.response(202), true)
		})
	}
}

func assertProjectV2IdentityDiff(t *testing.T, resource *schema.Resource, state map[string]string, config map[string]any, requestURI, response string, wantRequiresNew bool) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.URL.Path != requestURI {
			t.Fatalf("unexpected REST request path %q, want %q", request.URL.Path, requestURI)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, response)
	}))
	t.Cleanup(server.Close)

	baseURL := server.URL + "/"
	client, err := goGithub.NewClient(goGithub.WithHTTPClient(server.Client()), goGithub.WithURLs(&baseURL, nil))
	if err != nil {
		t.Fatalf("creating test GitHub client: %v", err)
	}

	diff, err := resource.Diff(t.Context(), &terraform.InstanceState{ID: "existing", Attributes: state}, terraform.NewResourceConfigRaw(config), &Owner{name: "old-owner", v3client: client, IsOrganization: true})
	if err != nil {
		t.Fatalf("calculating identity diff: %v", err)
	}
	if got := diff.RequiresNew(); got != wantRequiresNew {
		t.Fatalf("unexpected replacement decision: got %v, want %v; diff=%#v", got, wantRequiresNew, diff)
	}
}
