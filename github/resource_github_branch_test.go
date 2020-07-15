package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubBranch_basic(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	var (
		reference github.Reference

		name   = "basic"
		repo   = "test-repo"
		branch = "test-branch-" + acctest.RandString(5)
		ref    = "refs/heads/" + branch
		rn     = "github_branch." + name
		id     = repo + ":" + branch
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubBranchDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubBranchConfig(name, repo, branch),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubBranchExists(rn, id, &reference),
					resource.TestCheckResourceAttr(rn, "repository", repo),
					resource.TestCheckResourceAttr(rn, "branch", branch),
					resource.TestCheckResourceAttr(rn, "source_branch", "master"),
					resource.TestCheckResourceAttrSet(rn, "source_sha"),
					resource.TestCheckResourceAttrSet(rn, "etag"),
					resource.TestCheckResourceAttr(rn, "ref", ref),
					resource.TestCheckResourceAttrSet(rn, "sha"),
				),
			},
			{
				Config: testAccCheckGithubBranchConfig(name, repo, branch),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubBranchExists(rn, id, &reference),
					resource.TestCheckResourceAttr(rn, "repository", repo),
					resource.TestCheckResourceAttr(rn, "branch", branch),
					resource.TestCheckResourceAttr(rn, "source_branch", "master"),
					resource.TestCheckResourceAttrSet(rn, "source_sha"),
					resource.TestCheckResourceAttrSet(rn, "etag"),
					resource.TestCheckResourceAttr(rn, "ref", ref),
					resource.TestCheckResourceAttrSet(rn, "sha"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateId:     fmt.Sprintf("%s:%s", repo, branch),
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"source_sha",
				},
			},
		},
	})
}
func TestAccGithubBranch_withSourceBranch(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	var (
		reference github.Reference

		name         = "withSourceBranch"
		repo         = "test-repo"
		sourceBranch = "test-branch"
		branch       = "test-branch-" + acctest.RandString(5)
		ref          = "refs/heads/" + branch
		rn           = "github_branch." + name
		id           = repo + ":" + branch
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubBranchDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubBranchConfigWithSourceBranch(name, repo, sourceBranch, branch),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubBranchExists(rn, id, &reference),
					resource.TestCheckResourceAttr(rn, "repository", repo),
					resource.TestCheckResourceAttr(rn, "branch", branch),
					resource.TestCheckResourceAttr(rn, "source_branch", sourceBranch),
					resource.TestCheckResourceAttrSet(rn, "source_sha"),
					resource.TestCheckResourceAttrSet(rn, "etag"),
					resource.TestCheckResourceAttr(rn, "ref", ref),
					resource.TestCheckResourceAttrSet(rn, "sha"),
				),
			},
			{
				Config: testAccCheckGithubBranchConfigWithSourceBranch(name, repo, sourceBranch, branch),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubBranchExists(rn, id, &reference),
					resource.TestCheckResourceAttr(rn, "repository", repo),
					resource.TestCheckResourceAttr(rn, "branch", branch),
					resource.TestCheckResourceAttr(rn, "source_branch", sourceBranch),
					resource.TestCheckResourceAttrSet(rn, "source_sha"),
					resource.TestCheckResourceAttrSet(rn, "etag"),
					resource.TestCheckResourceAttr(rn, "ref", ref),
					resource.TestCheckResourceAttrSet(rn, "sha"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateId:     fmt.Sprintf("%s:%s:%s", repo, branch, sourceBranch),
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"source_sha",
				},
			},
		},
	})
}

func testAccCheckGithubBranchConfig(name, repo, branch string) string {
	return fmt.Sprintf(`
resource "github_branch" "%s" {
  repository = "%s"
  branch     = "%s"
}
`, name, repo, branch)
}

func testAccCheckGithubBranchConfigWithSourceBranch(name, repo, sourceBranch, branch string) string {
	return fmt.Sprintf(`
resource "github_branch" "%s" {
  repository    = "%s"
  source_branch = "%s"
  branch        = "%s"
}
`, name, repo, sourceBranch, branch)
}

func testAccCheckGithubBranchDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_branch" {
			continue
		}

		conn := testAccProvider.Meta().(*Owner).v3client
		orgName := testAccProvider.Meta().(*Owner).name
		repoName, branchName, err := parseTwoPartID(rs.Primary.ID, "repository", "branch")
		if err != nil {
			return err
		}

		ref, resp, err := conn.Git.GetRef(context.TODO(), orgName, repoName, branchName)
		if err == nil {
			if ref != nil {
				return fmt.Errorf("Repository branch still exists %s/%s (%s)",
					orgName, repoName, branchName)
			}
		}
		if resp.StatusCode != 404 {
			return fmt.Errorf("Error destroying branch %s/%s (%s)",
				orgName, repoName, branchName)
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

		conn := testAccProvider.Meta().(*Owner).v3client
		orgName := testAccProvider.Meta().(*Owner).name
		repoName, branchName, err := parseTwoPartID(rs.Primary.ID, "repository", "branch")
		if err != nil {
			return err
		}

		branchRefName := "refs/heads/" + branchName
		ref, _, err := conn.Git.GetRef(context.TODO(), orgName, repoName, branchRefName)
		if err != nil {
			return fmt.Errorf("Error querying GitHub branch reference %s/%s (%s): %s",
				orgName, repoName, branchRefName, err)
		}

		*reference = *ref
		return nil
	}
}
