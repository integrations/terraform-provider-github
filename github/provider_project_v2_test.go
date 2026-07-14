package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestProjectV2ResourcesValidate(t *testing.T) {
	t.Parallel()
	resources := map[string]*schema.Resource{
		"github_project":                  resourceGithubProject(),
		"github_project_field":            resourceGithubProjectField(),
		"github_project_item":             resourceGithubProjectItem(),
		"github_project_item_field_value": resourceGithubProjectItemFieldValue(),
		"github_project_repository":       resourceGithubProjectRepository(),
		"github_team_project":             resourceGithubTeamProject(),
	}
	for name, resource := range resources {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if len(resource.Schema) == 0 {
				t.Fatal("resource has no schema")
			}
			if NewProvider("test", "none")().ResourcesMap[name] == nil {
				t.Fatalf("resource is not registered in the provider")
			}
		})
	}
}

func TestValidateProjectV2Date(t *testing.T) {
	t.Parallel()
	if _, errors := validateProjectV2Date("2026-07-14", "date"); len(errors) != 0 {
		t.Fatalf("valid date was rejected: %v", errors)
	}
	if _, errors := validateProjectV2Date("2026-07-14T00:00:00Z", "date"); len(errors) == 0 {
		t.Fatal("RFC3339 timestamp was accepted as a Projects V2 date")
	}
}
