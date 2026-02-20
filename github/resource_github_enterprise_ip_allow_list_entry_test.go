package github

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubEnterpriseIpAllowListEntry(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		resourceName := "github_enterprise_ip_allow_list_entry.test"
		ip := "192.168.1.0/24"
		name := "Test Entry"
		isActive := true

		config := `
resource "github_enterprise_ip_allow_list_entry" "test" {
	enterprise_slug = "%s"
	ip              = "%s"
	name            = "%s"
	is_active       = %t
}
`

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessEnterprise(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testAccConf.enterpriseSlug, ip, name, isActive),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "enterprise_slug", testAccConf.enterpriseSlug),
						resource.TestCheckResourceAttr(resourceName, "ip", ip),
						resource.TestCheckResourceAttr(resourceName, "name", name),
						resource.TestCheckResourceAttr(resourceName, "is_active", strconv.FormatBool(isActive)),
					),
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}

func TestAccGithubEnterpriseIpAllowListEntry_update(t *testing.T) {
	t.Run("update", func(t *testing.T) {
		resourceName := "github_enterprise_ip_allow_list_entry.test"
		ip := "192.168.1.0/24"
		name := "Test Entry"
		isActive := true

		updatedIP := "10.0.0.0/16"
		updatedName := "Updated Entry"
		updatedIsActive := false

		config := `
	resource "github_enterprise_ip_allow_list_entry" "test" {
		enterprise_slug = "%s"
		ip              = "%s"
		name            = "%s"
		is_active       = %t
	}
	`

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessEnterprise(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testAccConf.enterpriseSlug, ip, name, isActive),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "enterprise_slug", testAccConf.enterpriseSlug),
						resource.TestCheckResourceAttr(resourceName, "ip", ip),
						resource.TestCheckResourceAttr(resourceName, "name", name),
						resource.TestCheckResourceAttr(resourceName, "is_active", fmt.Sprintf("%t", isActive)),
					),
				},
				{
					Config: fmt.Sprintf(config, testAccConf.enterpriseSlug, updatedIP, updatedName, updatedIsActive),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "enterprise_slug", testAccConf.enterpriseSlug),
						resource.TestCheckResourceAttr(resourceName, "ip", updatedIP),
						resource.TestCheckResourceAttr(resourceName, "name", updatedName),
						resource.TestCheckResourceAttr(resourceName, "is_active", fmt.Sprintf("%t", updatedIsActive)),
					),
				},
			},
		})
	})
}
