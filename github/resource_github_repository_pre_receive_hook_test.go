package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryPreReceiveHook_basic(t *testing.T) {
	skipUnlessEnterpriseServer(t)

	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryPreReceiveHookConfig(randString, "enabled"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"github_repository_pre_receive_hook.test", "enforcement", "enabled"),
					resource.TestCheckResourceAttrSet(
						"github_repository_pre_receive_hook.test", "id"),
					resource.TestCheckResourceAttrSet(
						"github_repository_pre_receive_hook.test", "configuration_url"),
				),
			},
			{
				Config: testAccGithubRepositoryPreReceiveHookConfig(randString, "disabled"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"github_repository_pre_receive_hook.test", "enforcement", "disabled"),
					resource.TestCheckResourceAttrSet(
						"github_repository_pre_receive_hook.test", "id"),
					resource.TestCheckResourceAttrSet(
						"github_repository_pre_receive_hook.test", "configuration_url"),
				),
			},
		},
	})
}

func testAccGithubRepositoryPreReceiveHookConfig(randString string, enforcement string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name         = "foo-%[1]s"
  description  = "Terraform acceptance tests"
}

resource "github_repository_pre_receive_hook" "test" {
  name = "%[2]s"
  repository = github_repository.test.name
  enforcement = "%[3]s"
}
`, randString, testPreReceiveHookFunc(), enforcement)
}
