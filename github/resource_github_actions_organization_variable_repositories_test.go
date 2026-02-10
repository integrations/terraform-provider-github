package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsOrganizationVariableRepositories(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		variableName := fmt.Sprintf("test_%s", randomID)
		variableValue := "foo"
		repoName0 := fmt.Sprintf("%s%s-0", testResourcePrefix, randomID)
		repoName1 := fmt.Sprintf("%s%s-1", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "%s"
	visibility    = "selected"
}

resource "github_repository" "test_0" {
	name       = "%s"
	visibility = "public"
}

resource "github_repository" "test_1" {
	name       = "%s"
	visibility = "public"
}

resource "github_actions_organization_variable_repositories" "test" {
	variable_name = github_actions_organization_variable.test.variable_name
	selected_repository_ids = [
		github_repository.test_0.repo_id,
		github_repository.test_1.repo_id
	]
}
`, variableName, variableValue, repoName0, repoName1)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_organization_variable_repositories.test", "variable_name", "github_actions_organization_variable.test", "variable_name"),
						resource.TestCheckResourceAttr("github_actions_organization_variable_repositories.test", "selected_repository_ids.#", "2"),
					),
				},
			},
		})
	})
}
