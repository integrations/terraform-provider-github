package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccGithubCollaboratorsDataSource_basic(t *testing.T) {
	repoName := fmt.Sprintf("tf-acc-test-collab-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubCollaboratorsDataSourceConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_collaborators.test", "collaborator.#"),
					resource.TestCheckResourceAttr("data.github_collaborators.test", "affiliation", "all"),
				),
			},
		},
	})
}

func testAccCheckGithubCollaboratorsDataSourceConfig(repo string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name = "%s"
}

data "github_collaborators" "test" {
  owner      = "%s"
  repository = "${github_repository.test.name}"
}
`, repo, testOrganization)
}
