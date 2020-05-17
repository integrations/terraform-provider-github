package github

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsPublicKeyDataSource_noMatchReturnsError(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	repo := "non-existent"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubActionsPublicKeyDataSourceConfig(repo),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
		},
	})
}

func TestAccCheckGithubActionsPublicKeyDataSource_existing(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	repo := os.Getenv("GITHUB_TEMPLATE_REPOSITORY")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubActionsPublicKeyDataSourceConfig(repo),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_actions_public_key.test_pk", "key"),
					resource.TestCheckResourceAttrSet("data.github_actions_public_key.test_pk", "key_id"),
				),
			},
		},
	})
}

func testAccCheckGithubActionsPublicKeyDataSourceConfig(repo string) string {
	return fmt.Sprintf(`
data "github_actions_public_key" "test_pk" {
  repository = "%s"
}
`, repo)
}
