package github

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestDataSourceGithubOrganizationRepositoryCustomPropertyRejectsEmptyResponseName(t *testing.T) {
	t.Parallel()

	server := githubApiMock([]*mockResponse{{
		ExpectedUri:    "/orgs/test-org/properties/schema/test-property",
		ExpectedMethod: http.MethodGet,
		StatusCode:     http.StatusOK,
		ResponseBody:   `{"value_type":"string"}`,
	}})
	defer server.Close()

	meta := &Owner{
		name:           "test-org",
		v3client:       mustCreateTestGitHubClient(t, server.URL),
		IsOrganization: true,
	}
	data := schema.TestResourceDataRaw(t, dataSourceGithubOrganizationRepositoryCustomProperty().Schema, map[string]any{
		"property_name": "test-property",
	})

	diags := dataSourceGithubOrganizationRepositoryCustomPropertyRead(t.Context(), data, meta)
	if !diags.HasError() || !strings.Contains(diags[0].Summary+diags[0].Detail, "empty name") {
		t.Fatalf("Read() diagnostics = %v, want an empty-name response error", diags)
	}
}

func TestAccGithubOrganizationRepositoryCustomPropertyDataSource(t *testing.T) {
	const dataAddr = "data.github_organization_repository_custom_property.test"
	t.Parallel()

	t.Run("reads a property created by the fixture", func(t *testing.T) {
		t.Parallel()

		property := mustCreateTestOrganizationRepositoryCustomProperty(t, "single_select", []string{"a", "b"})
		config := fmt.Sprintf(`
data "github_organization_repository_custom_property" "test" {
  property_name = %q
}
`, property.GetPropertyName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(dataAddr, tfjsonpath.New("property_name"), knownvalue.StringExact(property.GetPropertyName())),
						statecheck.ExpectKnownValue(dataAddr, tfjsonpath.New("value_type"), knownvalue.StringExact("single_select")),
						statecheck.ExpectKnownValue(dataAddr, tfjsonpath.New("allowed_values"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("a"),
							knownvalue.StringExact("b"),
						})),
						// The fixture sets no default, so the list must come back empty
						// rather than as a phantom single empty string.
						statecheck.ExpectKnownValue(dataAddr, tfjsonpath.New("default_value"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})
}
