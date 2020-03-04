package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubBranch_basic(t *testing.T) {

	var (
		reference github.Reference

		name   = "test"
		repo   = "test-repo"
		branch = "test-branch-" + acctest.RandString(5)
		rn     = "github_branch." + name
		id     = repo + ":" + branch
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubBranchConfig(name, repo, branch),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubBranchExists(rn, id, &reference),
					resource.TestCheckResourceAttr(rn, "repository", repo),
					resource.TestCheckResourceAttr(rn, "branch", branch),
					resource.TestCheckResourceAttr(rn, "source_branch", "master"),
					resource.TestCheckResourceAttrSet(rn, "source_sha"),
					resource.TestCheckResourceAttrSet(rn, "etag"),
					resource.TestCheckResourceAttrSet(rn, "ref"),
					resource.TestCheckResourceAttrSet(rn, "sha"),
				),
			},
			{
				Config: testAccGithubBranchConfig(name, repo, branch),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubBranchExists(rn, id, &reference),
					resource.TestCheckResourceAttr(rn, "repository", repo),
					resource.TestCheckResourceAttr(rn, "branch", branch),
					resource.TestCheckResourceAttr(rn, "source_branch", "master"),
					resource.TestCheckResourceAttrSet(rn, "source_sha"),
					resource.TestCheckResourceAttrSet(rn, "etag"),
					resource.TestCheckResourceAttrSet(rn, "ref"),
					resource.TestCheckResourceAttrSet(rn, "sha"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"source_branch",
					"source_sha",
					"sha",
				},
			},
		},
	})
}

func testAccGithubBranchConfig(name, repo, branch string) string {
	return fmt.Sprintf(`
resource "github_branch" "%s" {
  repository = "%s"
  branch     = "%s"
}
`, name, repo, branch)
}

func testAccGithubBranchDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_branch" {
			continue
		}

		conn := testAccProvider.Meta().(*Organization).client
		orgName := testAccProvider.Meta().(*Organization).name
		repoName, branchName, err := parseTwoPartID(rs.Primary.ID, "repository", "branch")
		if err != nil {
			return err
		}

		ref, resp, err := conn.Git.GetRef(context.TODO(), orgName, repoName, branchName)
		if err == nil {
			if ref != nil {
				return fmt.Errorf("Repository branch still exists: %s/%s (%s)",
					orgName, repoName, branchName)
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccCheckGithubBranchExists(n, id string, reference *github.Reference) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID != id {
			return fmt.Errorf("Expected ID to be %v, got %v", id, rs.Primary.ID)
		}

		conn := testAccProvider.Meta().(*Organization).client
		orgName := testAccProvider.Meta().(*Organization).name
		repoName, branchName, err := parseTwoPartID(rs.Primary.ID, "repository", "branch")
		if err != nil {
			return err
		}

		ref, _, err := conn.Git.GetRef(context.TODO(), orgName, repoName, "refs/heads/"+branchName)
		if err != nil {
			return err
		}

		*reference = *ref
		return nil
	}
}
