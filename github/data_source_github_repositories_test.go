package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccGithubRepositoriesDataSource_basic(t *testing.T) {
	query := "org:hashicorp terraform"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubRepositoriesDataSourceConfig(query),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_repositories.test", "full_names.#"),
					resource.TestCheckResourceAttr("data.github_repositories.test", "full_names.3450805659", "hashicorp/terraform"),
					resource.TestCheckResourceAttrSet("data.github_repositories.test", "names.#"),
					resource.TestCheckResourceAttr("data.github_repositories.test", "names.535570215", "terraform"),
				),
			},
		},
	})
}

func TestAccGithubRepositoriesDataSource_noMatch(t *testing.T) {
	query := "klsafj_23434_doesnt_exist"
	resource.Test(t, resource.TestCase{
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
