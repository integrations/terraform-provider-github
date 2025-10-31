package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseSettings(t *testing.T) {
	t.Skip("This test requires enterprise access and should only be run manually")

	if testEnterprise == "" {
		t.Skip("Skipping because `GITHUB_TEST_ENTERPRISE` is not set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubEnterpriseSettingsConfig(testEnterprise),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("github_enterprise_settings.test", "enterprise_slug", testEnterprise),
					resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_enabled_organizations", "all"),
					resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_allowed_actions", "all"),
					resource.TestCheckResourceAttr("github_enterprise_settings.test", "default_workflow_permissions", "read"),
					resource.TestCheckResourceAttr("github_enterprise_settings.test", "can_approve_pull_request_reviews", "false"),
				),
			},
		},
	})
}

func testAccGithubEnterpriseSettingsConfig(enterprise string) string {
	return fmt.Sprintf(`
resource "github_enterprise_settings" "test" {
  enterprise_slug = "%s"
  
  actions_enabled_organizations = "all"
  actions_allowed_actions = "all"
  
  default_workflow_permissions = "read"
  can_approve_pull_request_reviews = false
}
`, enterprise)
}
