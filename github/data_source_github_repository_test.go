package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryDataSource_fullName_noMatchReturnsError(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	fullName := "klsafj_23434_doesnt_exist/not-exists"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubRepositoryDataSourceConfig_fullName(fullName),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
		},
	})
}

func TestAccGithubRepositoryDataSource_name_noMatchReturnsError(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	name := "not-exists"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubRepositoryDataSourceConfig_name(name),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
		},
	})
}

func TestAccGithubRepositoryDataSource_fullName_existing(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	fullName := testOwner + "/test-repo"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubRepositoryDataSourceConfig_fullName(fullName),
				Check:  testRepoCheck(),
			},
		},
	})
}

func TestAccGithubRepositoryDataSource_name_existing(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	name := "test-repo"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubRepositoryDataSourceConfig_name(name),
				Check:  testRepoCheck(),
			},
		},
	})
}

func testRepoCheck() resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.github_repository.test", "id", "test-repo"),
		resource.TestCheckResourceAttr("data.github_repository.test", "name", "test-repo"),
		resource.TestCheckResourceAttr("data.github_repository.test", "private", "false"),
		resource.TestCheckResourceAttr("data.github_repository.test", "description", "Test description, used in GitHub Terraform provider acceptance test."),
		resource.TestCheckResourceAttr("data.github_repository.test", "homepage_url", "http://www.example.com"),
		resource.TestCheckResourceAttr("data.github_repository.test", "has_issues", "true"),
		resource.TestCheckResourceAttr("data.github_repository.test", "has_wiki", "true"),
		resource.TestCheckResourceAttr("data.github_repository.test", "allow_merge_commit", "true"),
		resource.TestCheckResourceAttr("data.github_repository.test", "allow_squash_merge", "true"),
		resource.TestCheckResourceAttr("data.github_repository.test", "allow_rebase_merge", "true"),
		resource.TestCheckResourceAttr("data.github_repository.test", "has_downloads", "true"),
		resource.TestCheckResourceAttr("data.github_repository.test", "full_name", testOwner+"/test-repo"),
		resource.TestCheckResourceAttr("data.github_repository.test", "default_branch", "master"),
		resource.TestCheckResourceAttr("data.github_repository.test", "html_url", "https://github.com/"+testOwner+"/test-repo"),
		resource.TestCheckResourceAttr("data.github_repository.test", "ssh_clone_url", "git@github.com:"+testOwner+"/test-repo.git"),
		resource.TestCheckResourceAttr("data.github_repository.test", "svn_url", "https://github.com/"+testOwner+"/test-repo"),
		resource.TestCheckResourceAttr("data.github_repository.test", "git_clone_url", "git://github.com/"+testOwner+"/test-repo.git"),
		resource.TestCheckResourceAttr("data.github_repository.test", "http_clone_url", "https://github.com/"+testOwner+"/test-repo.git"),
		resource.TestCheckResourceAttr("data.github_repository.test", "archived", "false"),
		resource.TestCheckResourceAttr("data.github_repository.test", "topics.#", "2"),
		resource.TestCheckResourceAttr("data.github_repository.test", "topics.0", "second-test-topic"),
		resource.TestCheckResourceAttr("data.github_repository.test", "topics.1", "test-topic"),
	)
}

func testAccCheckGithubRepositoryDataSourceConfig_fullName(fullName string) string {
	return fmt.Sprintf(`
data "github_repository" "test" {
  full_name = "%s"
}
`, fullName)
}

func testAccCheckGithubRepositoryDataSourceConfig_name(name string) string {
	return fmt.Sprintf(`
data "github_repository" "test" {
  name = "%s"
}
`, name)
}
