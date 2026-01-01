package github

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGithubEnterpriseCostCenter(t *testing.T) {
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
	if len(users) < 2 {
		t.Skip("Skipping because `ENTERPRISE_TEST_USERS` must contain at least two usernames")
	}

	usersBefore := fmt.Sprintf("%q, %q", users[0], users[1])
	usersAfter := fmt.Sprintf("%q", users[0])

	configBefore := fmt.Sprintf(`
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
	`, testEnterprise, randomID, usersBefore, testEnterpriseCostCenterOrganization, testEnterpriseCostCenterRepository)

	configAfter := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_cost_center" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			name            = "tf-acc-test-updated-%s"

			users         = [%s]
			organizations = []
			repositories  = []
		}
	`, testEnterprise, randomID, usersAfter)

	configEmpty := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_cost_center" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			name            = "tf-acc-test-%s"

			users         = []
			organizations = []
			repositories  = []
		}
	`, testEnterprise, randomID)

	checkBefore := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "enterprise_slug", testEnterprise),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "name", fmt.Sprintf("tf-acc-test-%s", randomID)),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "state", "active"),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "organizations.#", "1"),
		resource.TestCheckTypeSetElemAttr("github_enterprise_cost_center.test", "organizations.*", testEnterpriseCostCenterOrganization),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "repositories.#", "1"),
		resource.TestCheckTypeSetElemAttr("github_enterprise_cost_center.test", "repositories.*", testEnterpriseCostCenterRepository),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "users.#", "2"),
		resource.TestCheckTypeSetElemAttr("github_enterprise_cost_center.test", "users.*", users[0]),
		resource.TestCheckTypeSetElemAttr("github_enterprise_cost_center.test", "users.*", users[1]),
	)

	checkAfter := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "name", fmt.Sprintf("tf-acc-test-updated-%s", randomID)),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "state", "active"),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "organizations.#", "0"),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "repositories.#", "0"),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "users.#", "1"),
		resource.TestCheckTypeSetElemAttr("github_enterprise_cost_center.test", "users.*", users[0]),
	)

	checkEmpty := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "organizations.#", "0"),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "repositories.#", "0"),
		resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "users.#", "0"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessMode(t, enterprise) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configBefore,
				Check:  checkBefore,
			},
			{
				Config: configAfter,
				Check:  checkAfter,
			},
			{
				Config: configEmpty,
				Check:  checkEmpty,
			},
			{
				ResourceName:      "github_enterprise_cost_center.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["github_enterprise_cost_center.test"]
					if !ok {
						return "", fmt.Errorf("resource not found in state")
					}
					return fmt.Sprintf("%s/%s", testEnterprise, rs.Primary.ID), nil
				},
			},
		},
	})
}

func splitCommaSeparated(v string) []string {
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	return out
}
