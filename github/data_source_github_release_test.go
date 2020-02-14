package github

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccGithubReleaseDataSource_fetchByLatestNoReleaseReturnsError(t *testing.T) {
	repo := "nonExistentRepo"
	owner := "no-user"
	retrieveBy := "latest"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, "", 0),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
		},
	})
}

func TestAccGithubReleaseDataSource_latestExisting(t *testing.T) {
	repo := "terraform"
	owner := "hashicorp"
	retrieveBy := "latest"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, "", 0),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.github_release.test", "url", regexp.MustCompile(`hashicorp/terraform`)),
					resource.TestMatchResourceAttr("data.github_release.test", "tarball_url", regexp.MustCompile(`hashicorp/terraform/tarball`)),
				),
			},
		},
	})

}

func TestAccGithubReleaseDataSource_fetchByIdWithNoIdReturnsError(t *testing.T) {
	retrieveBy := "id"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubReleaseDataSourceConfig("", "", retrieveBy, "", 0),
				ExpectError: regexp.MustCompile("release_id` must be set when `retrieve_by` = `id`"),
			},
		},
	})
}

func TestAccGithubReleaseDataSource_fetchByIdExisting(t *testing.T) {
	repo := "terraform"
	owner := "hashicorp"
	retrieveBy := "id"
	id := int64(23055013)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, "", id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.github_release.test", "release_id", strconv.FormatInt(id, 10)),
					resource.TestMatchResourceAttr("data.github_release.test", "url", regexp.MustCompile(`hashicorp/terraform`)),
					resource.TestMatchResourceAttr("data.github_release.test", "tarball_url", regexp.MustCompile(`hashicorp/terraform/tarball`)),
				),
			},
		},
	})
}

func TestAccGithubReleaseDataSource_fetchByTagNoTagReturnsError(t *testing.T) {
	repo := "terraform"
	owner := "hashicorp"
	retrieveBy := "tag"
	id := int64(23055013)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, "", id),
				ExpectError: regexp.MustCompile("`release_tag` must be set when `retrieve_by` = `tag`"),
			},
		},
	})
}

func TestAccGithubReleaseDataSource_fetchByTagExisting(t *testing.T) {
	repo := "terraform"
	owner := "hashicorp"
	retrieveBy := "tag"
	tag := "v0.12.20"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, tag, 0),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.github_release.test", "release_tag", tag),
					resource.TestMatchResourceAttr("data.github_release.test", "url", regexp.MustCompile(`hashicorp/terraform`)),
					resource.TestMatchResourceAttr("data.github_release.test", "tarball_url", regexp.MustCompile(`hashicorp/terraform/tarball`)),
				),
			},
		},
	})
}

func TestAccGithubReleaseDataSource_invalidRetrieveMethodReturnsError(t *testing.T) {
	retrieveBy := "not valid"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubReleaseDataSourceConfig("", "", retrieveBy, "", 0),
				ExpectError: regexp.MustCompile("one of: `latest`, `id`, `tag` must be set for `retrieve_by`"),
			},
		},
	})

}

func testAccCheckGithubReleaseDataSourceConfig(repo, owner, retrieveBy, tag string, id int64) string {
	return fmt.Sprintf(`
data "github_release" "test" {
	repository = "%s"
	owner = "%s"
	retrieve_by = "%s"
	release_tag = "%s"
	release_id = %d
}
`, repo, owner, retrieveBy, tag, id)
}
