package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubBranchDataSource_basic(t *testing.T) {

	var (
		name = "main"
		repo = "test-repo"
		rn   = "data.github_branch." + name
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubBranchDataSourceConfigDefault(name, repo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "repository", repo),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttrSet(rn, "ref"),
					resource.TestCheckResourceAttrSet(rn, "sha"),
				),
			},
			{
				Config: testAccCheckGithubBranchDataSourceConfigSpecified(name, repo, "test-branch"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "repository", repo),
					resource.TestCheckResourceAttr(rn, "branch", "test-branch"),
					resource.TestCheckResourceAttrSet(rn, "ref"),
					resource.TestCheckResourceAttrSet(rn, "sha"),
				),
			},
		},
	})

}

func testAccCheckGithubBranchDataSourceConfigDefault(name, repo string) string {
	return fmt.Sprintf(`
data "github_branch" "%s" {
  repository = "%s"
}
`, name, repo)
}

func testAccCheckGithubBranchDataSourceConfigSpecified(name, repo, branch string) string {
	return fmt.Sprintf(`
data "github_branch" "%s" {
  repository = "%s"
  branch = "%s"
}
`, name, repo, branch)
}
