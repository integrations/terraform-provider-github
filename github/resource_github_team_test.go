package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubTeam(t *testing.T) {
	t.Run("creates a team configured with defaults", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
resource "github_team" "test" {
	name = "%s"
}
`, teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_team.test", "slug"),
						resource.TestCheckResourceAttr("github_team.test", "privacy", "secret"),
						resource.TestCheckResourceAttr("github_team.test", "notification_setting", "notifications_enabled"),
					),
				},
			},
		})
	})

	t.Run("creates a team configured with alternatives", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testResourceName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
resource "github_team" "test" {
	name                 = "%s"
	privacy              = "closed"
	notification_setting = "notifications_disabled"
}
`, testResourceName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_team.test", "slug"),
						resource.TestCheckResourceAttr("github_team.test", "privacy", "closed"),
						resource.TestCheckResourceAttr("github_team.test", "notification_setting", "notifications_disabled"),
					),
				},
			},
		})
	})

	t.Run("creates a hierarchy of teams", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		team01Name := fmt.Sprintf("%steam-hier01-%s", testResourcePrefix, randomID)
		team02Name := fmt.Sprintf("%steam-hier02-%s", testResourcePrefix, randomID)
		team03Name := fmt.Sprintf("%steam-hier03-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_team" "team01" {
				name        = "%s"
				description = "Terraform acc test team01a"
				privacy     = "closed"
			}

			resource "github_team" "team02" {
				name           = "%s"
				description    = "Terraform acc test team02a"
				privacy        = "closed"
				parent_team_id = "${github_team.team01.id}"
			}

			resource "github_team" "team03" {
				name           = "%s"
				description    = "Terraform acc test team03a"
				privacy        = "closed"
				parent_team_id = "${github_team.team02.slug}"
			}
		`, team01Name, team02Name, team03Name)

		config2 := fmt.Sprintf(`
			resource "github_team" "team01" {
				name        = "%s"
				description = "Terraform acc test team01b"
				privacy     = "closed"
			}

			resource "github_team" "team02" {
				name           = "%s"
				description    = "Terraform acc test team02b"
				privacy        = "closed"
			}

			resource "github_team" "team03" {
				name           = "%s"
				description    = "Terraform acc test team03b"
				privacy        = "closed"
			}
		`, team01Name, team02Name, team03Name)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_team.team02", "parent_team_id"),
			resource.TestCheckResourceAttrSet("github_team.team03", "parent_team_id"),
		)

		check2 := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("github_team.team02", "parent_team_id", ""),
			resource.TestCheckResourceAttr("github_team.team03", "parent_team_id", ""),
			resource.TestCheckResourceAttr("github_team.team02", "parent_team_read_id", ""),
			resource.TestCheckResourceAttr("github_team.team03", "parent_team_read_id", ""),
			resource.TestCheckResourceAttr("github_team.team02", "parent_team_read_slug", ""),
			resource.TestCheckResourceAttr("github_team.team03", "parent_team_read_slug", ""),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					Config: config2,
					Check:  check2,
				},
			},
		})
	})

	t.Run("creates a team and removes the default maintainer", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-no-maint-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name         = "%s"
				create_default_maintainer = false
			}
		`, teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_team.test", "members_count", "0"),
					),
				},
			},
		})
	})

	t.Run("updates_slug", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-slug-%s", testResourcePrefix, randomID)
		teamNameUpdated := fmt.Sprintf("%s-updated", teamName)
		config := `
resource "github_team" "test" {
	name         = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, teamName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_team.test", "slug", teamName),
					),
				},
				{
					Config: fmt.Sprintf(config, teamNameUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_team.other", "description", teamNameUpdated),
					),
				},
			},
		})
	})
}
