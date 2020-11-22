package github

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryMilestoneDataSource_noMatchReturnsError(t *testing.T) {
	repo := "nonExistentRepo"
	owner := "no-user"
	number := "1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubRepositoryMilestoneDataSourceNonExistentConfig(repo, owner, number),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
		},
	})
}

func TestAccGithubRepositoryMilestoneDataSource_existing(t *testing.T) {
	repo := acctest.RandomWithPrefix("tf-acc-test")
	title := acctest.RandomWithPrefix("ms")
	description := acctest.RandomWithPrefix("tf-acc-test-desc")
	dueDate := time.Now().UTC().Format(layoutISO)

	rn := "github_repository_milestone.test"
	dataSource := "data.github_repository_milestone.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubRepositoryMilestoneDataSourceConfig(repo, title, description, dueDate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSource, "title", rn, "title"),
					resource.TestCheckResourceAttrPair(dataSource, "description", rn, "description"),
					resource.TestCheckResourceAttrPair(dataSource, "due_date", rn, "due_date"),
					resource.TestCheckResourceAttrPair(dataSource, "state", rn, "state"),
					resource.TestCheckResourceAttrPair(dataSource, "number", rn, "number"),
					resource.TestCheckResourceAttrPair(dataSource, "owner", rn, "owner"),
					resource.TestCheckResourceAttrPair(dataSource, "repository", rn, "repository"),
				),
			},
		},
	})
}

func testAccCheckGithubRepositoryMilestoneDataSourceNonExistentConfig(owner, repo, number string) string {
	return fmt.Sprintf(`
data "github_repository_milestone" "test" {
	owner = "%s"
	repository = "%s"
    number = "%s"
}
`, owner, repo, number)
}

func testAccCheckGithubRepositoryMilestoneDataSourceConfig(repo, title, description, dueDate string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
    name = "%s"
}

resource "github_repository_milestone" "test" {
	owner = split("/", "${github_repository.test.full_name}")[0]
	repository = github_repository.test.name
    title = "%s"
    description = "%s"
    due_date = "%s"
}

data "github_repository_milestone" "test" {
	owner = github_repository_milestone.test.owner
	repository = github_repository_milestone.test.repository
    number = github_repository_milestone.test.number
}
`, repo, title, description, dueDate)
}
