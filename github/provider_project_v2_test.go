package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
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
			assertProjectV2SchemaContract(t, name, resource.Schema)
			if NewProvider("test", "none")().ResourcesMap[name] == nil {
				t.Fatalf("resource is not registered in the provider")
			}
		})
	}
}

func assertProjectV2SchemaContract(t *testing.T, prefix string, schemas map[string]*schema.Schema) {
	t.Helper()
	for name, value := range schemas {
		path := fmt.Sprintf("%s.%s", prefix, name)
		if strings.TrimSpace(value.Description) == "" {
			t.Errorf("%s has no description", path)
		}
		if value.ValidateFunc != nil {
			t.Errorf("%s uses deprecated ValidateFunc", path)
		}
		if nested, ok := value.Elem.(*schema.Resource); ok {
			assertProjectV2SchemaContract(t, path, nested.Schema)
		}
	}
}

func TestValidateProjectV2Date(t *testing.T) {
	t.Parallel()
	path := cty.GetAttrPath("date")
	if diagnostics := validateProjectV2Date("2026-07-14", path); diagnostics.HasError() {
		t.Fatalf("valid date was rejected: %v", diagnostics)
	}
	if diagnostics := validateProjectV2Date("2026-07-14T00:00:00Z", path); !diagnostics.HasError() {
		t.Fatal("RFC3339 timestamp was accepted as a Projects V2 date")
	}
}
