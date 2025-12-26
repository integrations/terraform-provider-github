package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseCostCenterDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}
	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}
	testEnterpriseCostCenterOrganization := os.Getenv("ENTERPRISE_TEST_ORGANIZATION")
	testEnterpriseCostCenterRepository := os.Getenv("ENTERPRISE_TEST_REPOSITORY")
	testEnterpriseCostCenterUsers := os.Getenv("ENTERPRISE_TEST_USERS")

	if testEnterpriseCostCenterOrganization == "" {
		t.Skip("Skipping because `ENTERPRISE_TEST_ORGANIZATION` is not set")
	}
	if testEnterpriseCostCenterRepository == "" {
		t.Skip("Skipping because `ENTERPRISE_TEST_REPOSITORY` is not set")
	}
	if testEnterpriseCostCenterUsers == "" {
		t.Skip("Skipping because `ENTERPRISE_TEST_USERS` is not set")
	}

	users := splitCommaSeparated(testEnterpriseCostCenterUsers)
	if len(users) == 0 {
		t.Skip("Skipping because `ENTERPRISE_TEST_USERS` must contain at least one username")
	}

	userList := fmt.Sprintf("%q", users[0])

	config := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_cost_center" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			name            = "tf-acc-test-%s"

			users         = [%s]
			organizations = [%q]
			repositories  = [%q]
		}

		data "github_enterprise_cost_center" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			cost_center_id  = github_enterprise_cost_center.test.id
		}
	`, testEnterprise, randomID, userList, testEnterpriseCostCenterOrganization, testEnterpriseCostCenterRepository)

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrPair("data.github_enterprise_cost_center.test", "cost_center_id", "github_enterprise_cost_center.test", "id"),
		resource.TestCheckResourceAttrPair("data.github_enterprise_cost_center.test", "name", "github_enterprise_cost_center.test", "name"),
		resource.TestCheckResourceAttr("data.github_enterprise_cost_center.test", "state", "active"),
		resource.TestCheckResourceAttr("data.github_enterprise_cost_center.test", "organizations.#", "1"),
		resource.TestCheckTypeSetElemAttr("data.github_enterprise_cost_center.test", "organizations.*", testEnterpriseCostCenterOrganization),
		resource.TestCheckResourceAttr("data.github_enterprise_cost_center.test", "repositories.#", "1"),
		resource.TestCheckTypeSetElemAttr("data.github_enterprise_cost_center.test", "repositories.*", testEnterpriseCostCenterRepository),
		resource.TestCheckResourceAttr("data.github_enterprise_cost_center.test", "users.#", "1"),
		resource.TestCheckTypeSetElemAttr("data.github_enterprise_cost_center.test", "users.*", users[0]),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessMode(t, enterprise) },
		Providers: testAccProviders,
		Steps:     []resource.TestStep{{Config: config, Check: check}},
	})
}
