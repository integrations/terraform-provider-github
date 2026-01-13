package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationSecurityManager(t *testing.T) {
	t.Run("adds team as security manager", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-sec-mgr-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			resource "github_organization_security_manager" "test" {
				team_slug = github_team.test.slug
			}
		`, teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_team.test", "ID", "github_organization_security_manager.test", "ID"),
						resource.TestCheckResourceAttrPair("github_team.test", "slug", "github_organization_security_manager.test", "team_slug"),
						resource.TestCheckResourceAttr("github_organization_security_manager.test", "team_slug", teamName),
					),
				},
			},
		})
	})

	t.Run("handles team name changes", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-sec-mgr-%s", testResourcePrefix, randomID)
		teamNameUpdated := fmt.Sprintf("%s-updated", teamName)
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			resource "github_organization_security_manager" "test" {
				team_slug = github_team.test.slug
			}
		`, teamName)

		configUpdated := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			resource "github_organization_security_manager" "test" {
				team_slug = github_team.test.slug
			}
		`, teamNameUpdated)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_team.test", "ID", "github_organization_security_manager.test", "ID"),
						resource.TestCheckResourceAttrPair("github_team.test", "slug", "github_organization_security_manager.test", "team_slug"),
						resource.TestCheckResourceAttr("github_organization_security_manager.test", "team_slug", teamName),
					),
				},
				{
					Config: configUpdated,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_team.test", "ID", "github_organization_security_manager.test", "ID"),
						resource.TestCheckResourceAttrPair("github_team.test", "slug", "github_organization_security_manager.test", "team_slug"),
						resource.TestCheckResourceAttr("github_organization_security_manager.test", "team_slug", teamNameUpdated),
					),
				},
			},
		})
	})
}
