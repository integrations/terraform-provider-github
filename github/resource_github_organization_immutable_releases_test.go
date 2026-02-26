package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationImmutableReleases(t *testing.T) {
	t.Run("test setting enforced_repositories to all", func(t *testing.T) {
		config := `
			resource "github_organization_immutable_releases" "test" {
				enforced_repositories = "all"
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_immutable_releases.test", "enforced_repositories", "all",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("test setting enforced_repositories to none", func(t *testing.T) {
		config := `
			resource "github_organization_immutable_releases" "test" {
				enforced_repositories = "none"
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_immutable_releases.test", "enforced_repositories", "none",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("test setting enforced_repositories to selected with repository IDs", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-immutable-rel-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics      = ["terraform", "testing"]
			}

			resource "github_organization_immutable_releases" "test" {
				enforced_repositories  = "selected"
				selected_repository_ids = [github_repository.test.repo_id]
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_immutable_releases.test", "enforced_repositories", "selected",
			),
			resource.TestCheckResourceAttr(
				"github_organization_immutable_releases.test", "selected_repository_ids.#", "1",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("test import of organization immutable releases", func(t *testing.T) {
		config := `
			resource "github_organization_immutable_releases" "test" {
				enforced_repositories = "all"
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_immutable_releases.test", "enforced_repositories", "all",
			),
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
					ResourceName:      "github_organization_immutable_releases.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("test updating enforced_repositories from all to selected", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-immutable-rel-%s", testResourcePrefix, randomID)

		configAll := `
			resource "github_organization_immutable_releases" "test" {
				enforced_repositories = "all"
			}
		`

		configSelected := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics      = ["terraform", "testing"]
			}

			resource "github_organization_immutable_releases" "test" {
				enforced_repositories  = "selected"
				selected_repository_ids = [github_repository.test.repo_id]
			}
		`, repoName)

		checkAll := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_immutable_releases.test", "enforced_repositories", "all",
			),
		)

		checkSelected := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_immutable_releases.test", "enforced_repositories", "selected",
			),
			resource.TestCheckResourceAttr(
				"github_organization_immutable_releases.test", "selected_repository_ids.#", "1",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configAll,
					Check:  checkAll,
				},
				{
					Config: configSelected,
					Check:  checkSelected,
				},
			},
		})
	})
}
