package github

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestDataSourceGithubOrganizationRepositoryCustomPropertiesRead(t *testing.T) {
	t.Parallel()

	server := githubApiMock([]*mockResponse{{
		ExpectedUri:    "/orgs/test-org/properties/schema",
		ExpectedMethod: http.MethodGet,
		StatusCode:     http.StatusOK,
		ResponseBody: `[
  {
    "property_name": "environment",
    "value_type": "single_select",
    "required": false,
    "default_value": "production",
    "description": "Deployment environment",
    "allowed_values": ["development", "production"],
    "values_editable_by": "org_actors"
  },
  {
    "property_name": "teams",
    "value_type": "multi_select",
    "required": false,
    "default_value": ["backend", "frontend"],
    "allowed_values": ["backend", "frontend"],
    "values_editable_by": "org_actors"
  },
  {
    "property_name": "owner",
    "value_type": "string",
    "default_value": "platform",
    "allowed_values": ["ignored"]
  }
]`,
	}})
	defer server.Close()

	meta := &Owner{
		name:           "test-org",
		v3client:       mustCreateTestGitHubClient(t, server.URL),
		IsOrganization: true,
	}
	data := schema.TestResourceDataRaw(t, dataSourceGithubOrganizationRepositoryCustomProperties().Schema, nil)

	diags := dataSourceGithubOrganizationRepositoryCustomPropertiesRead(t.Context(), data, meta)
	if diags.HasError() {
		t.Fatalf("Read() diagnostics = %v", diags)
	}
	if data.Id() != "test-org" {
		t.Fatalf("Read() id = %q, want test-org", data.Id())
	}

	propertySet, ok := data.Get("properties").(*schema.Set)
	if !ok {
		t.Fatalf("properties has type %T, want *schema.Set", data.Get("properties"))
	}
	if propertySet.Len() != 3 {
		t.Fatalf("properties has %d items, want 3", propertySet.Len())
	}

	propertiesByName := make(map[string]map[string]any, propertySet.Len())
	for _, value := range propertySet.List() {
		property, ok := value.(map[string]any)
		if !ok {
			t.Fatalf("property has type %T, want map[string]any", value)
		}
		propertiesByName[property["property_name"].(string)] = property
	}

	assertList := func(name, attribute string, want []any) {
		t.Helper()
		got, ok := propertiesByName[name][attribute].([]any)
		if !ok || !reflect.DeepEqual(got, want) {
			t.Errorf("properties[%q].%s = %v, want %v", name, attribute, propertiesByName[name][attribute], want)
		}
	}
	assertList("environment", "default_value", []any{"production"})
	assertList("environment", "allowed_values", []any{"development", "production"})
	assertList("teams", "default_value", []any{"backend", "frontend"})
	assertList("owner", "allowed_values", []any{})
}

func TestAccDataSourceGithubOrganizationRepositoryCustomProperties(t *testing.T) {
	t.Parallel()
	if testAccConf == nil {
		t.Skip("requires TF_ACC=1")
	}
	skipUnlessHasOrgs(t)

	property := mustCreateTestOrganizationRepositoryCustomProperty(t, "single_select", []string{"a", "b"})

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessHasOrgs(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "github_organization_repository_custom_properties" "test" {}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.github_organization_repository_custom_properties.test", tfjsonpath.New("properties"), knownvalue.SetPartial([]knownvalue.Check{
						knownvalue.MapPartial(map[string]knownvalue.Check{
							"property_name":  knownvalue.StringExact(property.GetPropertyName()),
							"value_type":     knownvalue.StringExact("single_select"),
							"allowed_values": knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("a"), knownvalue.StringExact("b")}),
							"default_value":  knownvalue.ListSizeExact(0),
						}),
					})),
				},
			},
		},
	})
}
