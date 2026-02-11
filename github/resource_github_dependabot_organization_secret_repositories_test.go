package github

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubDependabotOrganizationSecretRepositories(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		secretName := fmt.Sprintf("test_%s", randomID)
		secretValue := base64.StdEncoding.EncodeToString([]byte("foo"))
		repoName0 := fmt.Sprintf("%s%s-0", testResourcePrefix, randomID)
		repoName1 := fmt.Sprintf("%s%s-1", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_dependabot_organization_secret" "test" {
	secret_name     = "%s"
	encrypted_value = "%s"
	visibility      = "selected"
}

resource "github_repository" "test_0" {
	name       = "%s"
	visibility = "public"
}

resource "github_repository" "test_1" {
	name       = "%s"
	visibility = "public"
}

resource "github_dependabot_organization_secret_repositories" "test" {
	secret_name = github_dependabot_organization_secret.test.secret_name
	selected_repository_ids = [
		github_repository.test_0.repo_id,
		github_repository.test_1.repo_id
	]
}
`, secretName, secretValue, repoName0, repoName1)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_dependabot_organization_secret_repositories.test", "secret_name", "github_dependabot_organization_secret.test", "secret_name"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret_repositories.test", "selected_repository_ids.#", "2"),
					),
				},
			},
		})
	})
}
