package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccGithubReleaseDataSource_fetchByLatestNoReleaseReturnsError(t *testing.T) {
	repo := "nonExistantRepo"
	owner := "no-user"
	retrieveBy := "latest"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
		},
	})
}

// func TestAccGithubReleaseDataSource_latestExisting(t *testing.T) {

// }

// func TestAccGithubReleaseDataSource_fetchByIdWithNoIdReturnsError(t *testing.T) {

// }

// func TestAccGithubReleaseDataSource_fetchByIdExisting(t *testing.T) {

// }

// func TestAccGithubReleaseDataSource_fetchByTagNoTagReturnsError(t *testing.T) {

// }

// func TestAccGithubReleaseDataSource_fetchByTagExisting(t *testing.T) {

// }

// func TestAccGithubReleaseDataSource_invalidRetrieveMethodReturnsError(t *testing.T) {

// }

func testAccCheckGithubReleaseDataSourceConfig(repo string, owner string, retrieveBy string) string {
	return fmt.Sprintf(`
data "github_release" "test" {
	repository = "%s"
	owner = "%s"
	retrieve_by = "%s"
}
`, repo, owner, retrieveBy)
}
