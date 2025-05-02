package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseIpAllowListEntry_basic(t *testing.T) {
	t.Skip("Acceptance test requires a real GitHub Enterprise environment")

	resourceName := "github_enterprise_ip_allow_list_entry.test"
	enterpriseSlug := "test-enterprise"
	ip := "192.168.1.0/24"
	name := "Test Entry"
	isActive := true

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckEnterprise(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubEnterpriseIpAllowListEntryConfig(enterpriseSlug, ip, name, isActive),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "enterprise_slug", enterpriseSlug),
					resource.TestCheckResourceAttr(resourceName, "ip", ip),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "is_active", fmt.Sprintf("%t", isActive)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubEnterpriseIpAllowListEntry_update(t *testing.T) {
	t.Skip("Acceptance test requires a real GitHub Enterprise environment")

	resourceName := "github_enterprise_ip_allow_list_entry.test"
	enterpriseSlug := "test-enterprise"
	ip := "192.168.1.0/24"
	name := "Test Entry"
	isActive := true

	updatedIP := "10.0.0.0/16"
	updatedName := "Updated Entry"
	updatedIsActive := false

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckEnterprise(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubEnterpriseIpAllowListEntryConfig(enterpriseSlug, ip, name, isActive),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "enterprise_slug", enterpriseSlug),
					resource.TestCheckResourceAttr(resourceName, "ip", ip),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "is_active", fmt.Sprintf("%t", isActive)),
				),
			},
			{
				Config: testAccGithubEnterpriseIpAllowListEntryConfig(enterpriseSlug, updatedIP, updatedName, updatedIsActive),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "enterprise_slug", enterpriseSlug),
					resource.TestCheckResourceAttr(resourceName, "ip", updatedIP),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "is_active", fmt.Sprintf("%t", updatedIsActive)),
				),
			},
		},
	})
}

func testAccGithubEnterpriseIpAllowListEntryConfig(enterpriseSlug, ip, name string, isActive bool) string {
	return fmt.Sprintf(`
resource "github_enterprise_ip_allow_list_entry" "test" {
  enterprise_slug = "%s"
  ip              = "%s"
  name            = "%s"
  is_active       = %t
}
`, enterpriseSlug, ip, name, isActive)
}

func testAccPreCheckEnterprise(t *testing.T) {
	if v := testAccProvider.Meta().(*Owner).name; v == "" {
		t.Fatal("The GITHUB_ENTERPRISE_SLUG environment variable must be set for enterprise tests")
	}
}
