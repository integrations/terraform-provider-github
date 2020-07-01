package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoriesDataSource_basic(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	query := "org:hashicorp repository:terraform"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubRepositoriesDataSourceConfig(query),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.github_repositories.test", "full_names.0", regexp.MustCompile(`^hashicorp`)),
					resource.TestMatchResourceAttr("data.github_repositories.test", "names.0", regexp.MustCompile(`^terraform`)),
					resource.TestCheckResourceAttr("data.github_repositories.test", "sort", "updated"),
				),
			},
		},
	})
}
func TestAccGithubRepositoriesDataSource_Sort(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubRepositoriesDataSourceConfigWithSort("org:hashicorp repository:terraform", "updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.github_repositories.test", "full_names.0", regexp.MustCompile(`^hashicorp`)),
					resource.TestMatchResourceAttr("data.github_repositories.test", "names.0", regexp.MustCompile(`^terraform`)),
					resource.TestCheckResourceAttr("data.github_repositories.test", "sort", "updated"),
				),
			},
			{
				Config: testAccCheckGithubRepositoriesDataSourceConfigWithSort("org:hashicorp language:go", "stars"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.github_repositories.test", "full_names.0", "hashicorp/terraform"),
					resource.TestCheckResourceAttr("data.github_repositories.test", "names.0", "terraform"),
					resource.TestCheckResourceAttr("data.github_repositories.test", "sort", "stars"),
				),
			},
		},
	})
}

func TestAccGithubRepositoriesDataSource_noMatch(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	query := "klsafj_23434_doesnt_exist"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubRepositoriesDataSourceConfig(query),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.github_repositories.test", "full_names.#", "0"),
					resource.TestCheckResourceAttr("data.github_repositories.test", "names.#", "0"),
				),
			},
		},
	})
}

func testAccCheckGithubRepositoriesDataSourceConfig(query string) string {
	return fmt.Sprintf(`
data "github_repositories" "test" {
  query = "%s"
}
`, query)
}

func testAccCheckGithubRepositoriesDataSourceConfigWithSort(query, sort string) string {
	return fmt.Sprintf(`
data "github_repositories" "test" {
  query = "%s"
  sort  = "%s"
}
`, query, sort)
}
