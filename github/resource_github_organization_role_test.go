package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationRole(t *testing.T) {
	t.Run("can create an empty organization role", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		name := fmt.Sprintf("tf-acc-org-role-%s", randomID)
		config := fmt.Sprintf(`
			resource "github_organization_role" "test" {
				name        = "%s"
				permissions = []
			}
		`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, enterprise) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_organization_role.test", "id"),
						resource.TestCheckResourceAttrSet("github_organization_role.test", "role_id"),
						resource.TestCheckResourceAttr("github_organization_role.test", "name", name),
						resource.TestCheckResourceAttr("github_organization_role.test", "base_role", "none"),
						resource.TestCheckNoResourceAttr("github_organization_role.test", "permissions"),
						resource.TestCheckResourceAttr("github_organization_role.test", "permissions.#", "0"),
					),
				},
			},
		})
	})

	t.Run("can create an empty organization role with a base role", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		name := fmt.Sprintf("tf-acc-org-role-%s", randomID)
		baseRole := "read"

		config := fmt.Sprintf(`
			resource "github_organization_role" "test" {
				name        = "%s"
				base_role   = "%s"
				permissions = []
			}
		`, name, baseRole)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, enterprise) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_organization_role.test", "id"),
						resource.TestCheckResourceAttr("github_organization_role.test", "name", name),
						resource.TestCheckResourceAttr("github_organization_role.test", "base_role", baseRole),
						resource.TestCheckNoResourceAttr("github_organization_role.test", "permissions"),
						resource.TestCheckResourceAttr("github_organization_role.test", "permissions.#", "0"),
					),
				},
			},
		})
	})

	t.Run("can create an organization role", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		name := fmt.Sprintf("tf-acc-org-role-%s", randomID)
		baseRole := "none"
		permission0 := "read_organization_actions_usage_metrics"

		config := fmt.Sprintf(`
			resource "github_organization_role" "test" {
				name        = "%s"
				base_role   = "%s"
				permissions = [
				  "%s"
				]
			}
		`, name, baseRole, permission0)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, enterprise) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_organization_role.test", "id"),
						resource.TestCheckResourceAttr("github_organization_role.test", "name", name),
						resource.TestCheckResourceAttr("github_organization_role.test", "base_role", baseRole),
						resource.TestCheckResourceAttrSet("github_organization_role.test", "permissions.#"),
						resource.TestCheckResourceAttr("github_organization_role.test", "permissions.#", "1"),
						resource.TestCheckResourceAttr("github_organization_role.test", "permissions.0", permission0),
					),
				},
			},
		})
	})

	t.Run("can create an organization role with repo permissions", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		name := fmt.Sprintf("tf-acc-org-role-%s", randomID)
		description := "This is a test org role."
		baseRole := "write"
		permission0 := "read_audit_logs"
		config := fmt.Sprintf(`
			resource "github_organization_role" "test" {
				name        = "%s"
				description = "%s"
				base_role   = "%s"
				permissions = [
				"%s"
				]
			}
		`, name, description, baseRole, permission0)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, enterprise) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_organization_role.test", "id"),
						resource.TestCheckResourceAttrSet("github_organization_role.test", "role_id"),
						resource.TestCheckResourceAttr("github_organization_role.test", "name", name),
						resource.TestCheckResourceAttr("github_organization_role.test", "description", description),
						resource.TestCheckResourceAttr("github_organization_role.test", "base_role", baseRole),
						resource.TestCheckResourceAttrSet("github_organization_role.test", "permissions.#"),
						resource.TestCheckResourceAttr("github_organization_role.test", "permissions.#", "1"),
						resource.TestCheckResourceAttr("github_organization_role.test", "permissions.0", permission0),
					),
				},
			},
		})
	})

	t.Run("can create an organization role with org and repo permissions", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		name := fmt.Sprintf("tf-acc-org-role-%s", randomID)
		description := "This is a test org role."
		baseRole := "write"
		permission0 := "read_organization_actions_usage_metrics"
		permission1 := "read_audit_logs"
		config := fmt.Sprintf(`
			resource "github_organization_role" "test" {
				name        = "%s"
				description = "%s"
				base_role   = "%s"
				permissions = [
				"%s",
				"%s"
				]
			}
		`, name, description, baseRole, permission0, permission1)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, enterprise) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_organization_role.test", "id"),
						resource.TestCheckResourceAttrSet("github_organization_role.test", "role_id"),
						resource.TestCheckResourceAttr("github_organization_role.test", "name", name),
						resource.TestCheckResourceAttr("github_organization_role.test", "description", description),
						resource.TestCheckResourceAttr("github_organization_role.test", "base_role", baseRole),
						resource.TestCheckResourceAttrSet("github_organization_role.test", "permissions.#"),
						resource.TestCheckResourceAttr("github_organization_role.test", "permissions.#", "2"),
						resource.TestCheckTypeSetElemAttr("github_organization_role.test", "permissions.*", permission0),
						resource.TestCheckTypeSetElemAttr("github_organization_role.test", "permissions.*", permission1),
					),
				},
			},
		})
	})
}
