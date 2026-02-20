package github

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubDependabotOrganizationSecretRepository(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		secretName := fmt.Sprintf("test_%s", randomID)
		secretValue := base64.StdEncoding.EncodeToString([]byte("foo"))
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_dependabot_organization_secret" "test" {
	secret_name     = "%s"
	encrypted_value = "%s"
	visibility      = "selected"
}

resource "github_repository" "test" {
	name       = "%s"
	visibility = "public"
}

resource "github_dependabot_organization_secret_repository" "test" {
	secret_name   = github_dependabot_organization_secret.test.secret_name
	repository_id = github_repository.test.repo_id
}
`, secretName, secretValue, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_dependabot_organization_secret_repository.test", "secret_name", "github_dependabot_organization_secret.test", "secret_name"),
						resource.TestCheckResourceAttrPair("github_dependabot_organization_secret_repository.test", "repository_id", "github_repository.test", "repo_id"),
					),
				},
			},
		})
	})
}
