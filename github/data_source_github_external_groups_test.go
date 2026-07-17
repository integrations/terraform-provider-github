package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubExternalGroups(t *testing.T) {
	t.Parallel()

	if testAccConf.authMode != enterprise && !testAccConf.enterpriseIsEMU {
		t.Skip("Skipping as test mode is not enterprise using EMU")
	}

	t.Run("queries_all_external_groups", func(t *testing.T) {
		t.Parallel()

		if testAccConf.testExternalGroup1ID == 0 {
			t.Skip("Skipping as no external groups are configured for the test organization")
		}

		groupID := int32(testAccConf.testExternalGroup1ID)

		config := `
data "github_external_groups" "example" {}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_external_groups.example", tfjsonpath.New("groups"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapPartial(map[string]knownvalue.Check{
								"group_id": knownvalue.Int32Exact(groupID),
							}),
						})),
					},
				},
			},
		})
	})

	t.Run("queries_external_groups_with_filter", func(t *testing.T) {
		t.Parallel()

		if testAccConf.testExternalGroup1ID == 0 || testAccConf.testExternalGroup1DisplayName == "" || testAccConf.testExternalGroup2ID == 0 {
			t.Skip("Skipping as no external groups are configured for the test organization")
		}

		groupID := int32(testAccConf.testExternalGroup1ID)

		config := fmt.Sprintf(`
data "github_external_groups" "example" {
  display_name_filter = "%s"
}
`, testAccConf.testExternalGroup1DisplayName)

		// TODO: Use new not in set matcher when available from #3537
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_external_groups.example", tfjsonpath.New("groups"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.MapPartial(map[string]knownvalue.Check{
								"group_id": knownvalue.Int32Exact(groupID),
							}),
						})),
					},
				},
			},
		})
	})
}
