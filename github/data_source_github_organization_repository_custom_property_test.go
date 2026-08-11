package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

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
					},
				},
			},
		})
	})
}
