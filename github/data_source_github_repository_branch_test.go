package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccGithubRepositoryBranchDataSource_basic(t *testing.T) {
	name := "test"
	repo := "tf-acc-test-repo-branch-" + acctest.RandString(5)

	rn := "data.github_repository_branch." + name

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubRepositoryBranchDataSourceConfig(name, repo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "repository", repo),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttrSet(rn, "ref"),
					resource.TestCheckResourceAttrSet(rn, "sha"),
				),
			},
		},
	})
}

func testAccCheckGithubRepositoryBranchDataSourceConfig(name, repo string) string {
	return fmt.Sprintf(`
resource "github_repository" "%s" {
  name        = "%s"
  description = "Terraform Acceptance Test"
  auto_init   = true
}

data "github_repository_branch" "%s" {
  repository = github_repository.%s.name
}
`, name, repo, name, name)
}
