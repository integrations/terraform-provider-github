package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationRepositoryCustomPropertyDataSource(t *testing.T) {
	const dataAddr = "data.github_organization_repository_custom_property.test"

	t.Run("reads a property created by the resource", func(t *testing.T) {
		config := `
		resource "github_organization_repository_custom_property" "test" {
			property_name  = "tf-acc-test-ds"
			value_type     = "single_select"
			description    = "tf-acc-test data source"
			allowed_values = ["a", "b"]
		}

		data "github_organization_repository_custom_property" "test" {
			property_name = github_organization_repository_custom_property.test.property_name
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(dataAddr, tfjsonpath.New("value_type"), knownvalue.StringExact("single_select")),
						statecheck.ExpectKnownValue(dataAddr, tfjsonpath.New("description"), knownvalue.StringExact("tf-acc-test data source")),
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
