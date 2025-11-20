package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccGithubRepositoryEtagPresent tests that etag field is populated.
func TestAccGithubRepositoryEtagPresent(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	repoName := fmt.Sprintf("tf-acc-test-etag-%s", randomID)

	config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name      = "%s"
			auto_init = true
		}
	`, repoName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("github_repository.test", "name", repoName),
					resource.TestCheckResourceAttrSet("github_repository.test", "etag"),
				),
			},
		},
	})
}

// TestAccGithubRepositoryEtagNoDiff tests that re-running the same config shows no changes.
func TestAccGithubRepositoryEtagNoDiff(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	repoName := fmt.Sprintf("tf-acc-test-etag-nodiff-%s", randomID)

	config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name        = "%s"
			description = "Test repository for etag diff suppression"
			auto_init   = true
		}
	`, repoName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("github_repository.test", "name", repoName),
					resource.TestCheckResourceAttrSet("github_repository.test", "etag"),
				),
			},
			{
				// Re-run the same config - should not show any changes - etag diff suppression
				Config:   config,
				PlanOnly: true,
			},
		},
	})
}
