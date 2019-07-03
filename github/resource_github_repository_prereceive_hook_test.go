package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubRepositoryPreReceiveHook_basic(t *testing.T) {
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	enforcement := "enabled"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryPreReceiveHookConfig_basic(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGitHubPreReceiveHookEnforcement("github_repository_prereceive_hook.foo", enforcement),
					resource.TestCheckResourceAttr(
						"github_repository_prereceive_hook.foo", "enforcement", enforcement),
					resource.TestCheckResourceAttrSet(
						"github_repository_prereceive_hook.foo", "id"),
				),
			},
		},
	})
}

func testAccCheckGitHubPreReceiveHookEnforcement(n string, enforcement string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No hook ID is set")
		}

		id := strings.Split(rs.Primary.ID, "/")
		orgName, repoName, i := id[0], id[1], id[2]

		hookID, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return err
		}

		org := testAccProvider.Meta().(*Organization)
		client := org.client

		hook, _, err := client.Repositories.GetPreReceiveHook(context.TODO(), orgName, repoName, hookID)

		if err != nil {
			return err
		}

		if *hook.Enforcement != enforcement {
			return fmt.Errorf("Enforcement set to %s instead of %s", *hook.Enforcement, enforcement)
		}

		return nil
	}
}

func testAccGithubRepositoryPreReceiveHookConfig_basic(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "foo-%s"
  description  = "Terraform acceptance tests"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues    = true
  has_wiki      = true
  has_downloads = true
}

resource "github_repository_prereceive_hook" "foo" {
  name = "require-code-review"
  repository = "${github_repository.foo.name}"
  enforcement = "enabled"
}
`, randString)
}
