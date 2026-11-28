package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubUserSshKey(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		title := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		updatedTitle := fmt.Sprintf("%s-updated", title)
		key := mustNewSshPublicKey(t)
		updatedKey := mustNewSshPublicKey(t)

		config := `
resource "github_user_ssh_key" "test" {
  title = "%s"
  key   = "%s"
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, title, key),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_user_ssh_key.test", tfjsonpath.New("title"), knownvalue.StringExact(title)),
						statecheck.ExpectKnownValue("github_user_ssh_key.test", tfjsonpath.New("key"), knownvalue.StringExact(key)),
						statecheck.ExpectKnownValue("github_user_ssh_key.test", tfjsonpath.New("key_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_user_ssh_key.test", tfjsonpath.New("url"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, updatedTitle, key),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_user_ssh_key.test", plancheck.ResourceActionReplace),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_user_ssh_key.test", tfjsonpath.New("title"), knownvalue.StringExact(updatedTitle)),
					},
				},
				{
					Config: fmt.Sprintf(config, updatedTitle, updatedKey),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_user_ssh_key.test", plancheck.ResourceActionReplace),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_user_ssh_key.test", tfjsonpath.New("key"), knownvalue.StringExact(updatedKey)),
					},
				},
				{
					ResourceName:      "github_user_ssh_key.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
