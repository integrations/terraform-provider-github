package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsOrganizationVariableRepository(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		variableName := fmt.Sprintf("test_%s", randomID)
		variableValue := "foo"
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "%s"
	visibility    = "selected"
}

resource "github_repository" "test" {
	name       = "%s"
	visibility = "public"
}

resource "github_actions_organization_variable_repository" "test" {
	variable_name   = github_actions_organization_variable.test.variable_name
	repository_id = github_repository.test.repo_id
}
`, variableName, variableValue, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_organization_variable_repository.test", "variable_name", "github_actions_organization_variable.test", "variable_name"),
						resource.TestCheckResourceAttrPair("github_actions_organization_variable_repository.test", "repository_id", "github_repository.test", "repo_id"),
					),
				},
			},
		})
	})
}
