package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationIpAllowListEntry_basic(t *testing.T) {
	t.Skip("Acceptance test requires a real GitHub organization")

	resourceName := "github_organization_ip_allow_list_entry.test"
	orgName := "test-organization"
	ip := "192.168.1.0/24"
	name := "Test Entry"
	isActive := true

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckOrg(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubOrganizationIpAllowListEntryConfig(orgName, ip, name, isActive),
				Check: resource.ComposeTestCheckFunc(
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

func TestAccGithubOrganizationIpAllowListEntry_update(t *testing.T) {
	t.Skip("Acceptance test requires a real GitHub organization")

	resourceName := "github_organization_ip_allow_list_entry.test"
	orgName := "test-organization"
	ip := "192.168.1.0/24"
	name := "Test Entry"
	isActive := true

	updatedIP := "10.0.0.0/16"
	updatedName := "Updated Entry"
	updatedIsActive := false

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckOrg(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubOrganizationIpAllowListEntryConfig(orgName, ip, name, isActive),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ip", ip),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "is_active", fmt.Sprintf("%t", isActive)),
				),
			},
			{
				Config: testAccGithubOrganizationIpAllowListEntryConfig(orgName, updatedIP, updatedName, updatedIsActive),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ip", updatedIP),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "is_active", fmt.Sprintf("%t", updatedIsActive)),
				),
			},
		},
	})
}

func testAccGithubOrganizationIpAllowListEntryConfig(orgName, ip, name string, isActive bool) string {
	return fmt.Sprintf(`
provider "github" {
  owner = "%s"
}

resource "github_organization_ip_allow_list_entry" "test" {
  ip        = "%s"
  name      = "%s"
  is_active = %t
}
`, orgName, ip, name, isActive)
}

func testAccPreCheckOrg(t *testing.T) {
	if v := testAccProvider.Meta().(*Owner).name; v == "" {
		t.Fatal("The GITHUB_OWNER environment variable must be set for acceptance tests")
	}
}
