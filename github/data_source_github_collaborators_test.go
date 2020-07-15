package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubCollaboratorsDataSource_basic(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	dsn := "data.github_collaborators.test"
	repoName := fmt.Sprintf("tf-acc-test-collab-%s", acctest.RandString(5))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubCollaboratorsDataSourceConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "collaborator.#"),
					resource.TestCheckResourceAttr(dsn, "affiliation", "all"),
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
`, repo, testOwner)
}
