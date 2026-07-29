package github

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubOrganizationSettings(t *testing.T) {
	// IMPORTANT: Do not run these tests in parallel as they modify the organization state.

	t.Skip("TODO: Make this test cleanup correctly")

	t.Run("creates organization settings without error", func(t *testing.T) {
		config := `
		resource "github_organization_settings" "test" {
			billing_email = "test@example.com"
			company = "Test Company"
			blog = "https://example.com"
			email = "test@example.com"
			twitter_username = "Test"
			location = "Test Location"
			name = "Test Name"
			description = "Test Description"
			has_organization_projects = true
			has_repository_projects = true
			default_repository_permission = "read"
			members_can_create_repositories = true
			members_can_create_public_repositories = true
			members_can_create_private_repositories = true
			members_can_create_internal_repositories = false
			members_can_create_pages = true
			members_can_create_public_pages = true
			members_can_create_private_pages = true
			members_can_fork_private_repositories = true
			web_commit_signoff_required = true
			advanced_security_enabled_for_new_repositories = false
			  dependabot_alerts_enabled_for_new_repositories=  false
			dependabot_security_updates_enabled_for_new_repositories = false
			dependency_graph_enabled_for_new_repositories = false
			secret_scanning_enabled_for_new_repositories = false
			secret_scanning_push_protection_enabled_for_new_repositories = false
		  }`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"billing_email", "test@example.com",
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

	t.Run("updates organization settings without error", func(t *testing.T) {
		billingEmail := "test1@example.com"
		company := "Test Company"
		blog := "https://test.com"
		updatedBillingEmail := "test2@example.com"
		updatedCompany := "Test Company 2"
		updatedBlog := "https://test2.com"

		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_organization_settings" "test" {
				billing_email = "%s"
				company = "%s"
				blog = "%s"
				}`, billingEmail, company, blog),

			"after": fmt.Sprintf(`
			resource "github_organization_settings" "test" {
				billing_email = "%s"
				company = "%s"
				blog = "%s"
				}`, updatedBillingEmail, updatedCompany, updatedBlog),
		}
		checks := map[string]resource.TestCheckFunc{
			"before": resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"billing_email", billingEmail,
			),
			"after": resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"billing_email", updatedBillingEmail,
			),
		}
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configs["before"],
					Check:  checks["before"],
				},
				{
					Config: configs["after"],
					Check:  checks["after"],
				},
			},
		})
	})

	t.Run("imports organization settings without error", func(t *testing.T) {
		billingEmail := "test@example.com"
		company := "Test Company"
		blog := "https://example.com"

		config := fmt.Sprintf(`
		resource "github_organization_settings" "test" {
			billing_email = "%s"
			company = "%s"
			blog = "%s"
			}`, billingEmail, company, blog)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"billing_email", billingEmail,
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
					ResourceName:      "github_organization_settings.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("handles boolean false values correctly", func(t *testing.T) {
		config := `
		resource "github_organization_settings" "test" {
			billing_email = "test@example.com"
			members_can_create_private_repositories = false
			members_can_create_internal_repositories = false
			members_can_fork_private_repositories = false
			web_commit_signoff_required = false
			advanced_security_enabled_for_new_repositories = false
			dependabot_alerts_enabled_for_new_repositories = false
			dependabot_security_updates_enabled_for_new_repositories = false
			dependency_graph_enabled_for_new_repositories = false
			secret_scanning_enabled_for_new_repositories = false
			secret_scanning_push_protection_enabled_for_new_repositories = false
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"billing_email", "test@example.com",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"members_can_create_private_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"members_can_create_internal_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"members_can_fork_private_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"web_commit_signoff_required", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"advanced_security_enabled_for_new_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"dependabot_alerts_enabled_for_new_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"dependabot_security_updates_enabled_for_new_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"dependency_graph_enabled_for_new_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"secret_scanning_enabled_for_new_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"secret_scanning_push_protection_enabled_for_new_repositories", "false",
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

	t.Run("handles mixed boolean values correctly", func(t *testing.T) {
		config := `
		resource "github_organization_settings" "test" {
			billing_email = "test@example.com"
			members_can_create_private_repositories = false
			members_can_create_internal_repositories = true
			members_can_fork_private_repositories = false
			web_commit_signoff_required = true
			advanced_security_enabled_for_new_repositories = false
			dependabot_alerts_enabled_for_new_repositories = true
			dependabot_security_updates_enabled_for_new_repositories = false
			dependency_graph_enabled_for_new_repositories = true
			secret_scanning_enabled_for_new_repositories = false
			secret_scanning_push_protection_enabled_for_new_repositories = true
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"billing_email", "test@example.com",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"members_can_create_private_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"members_can_create_internal_repositories", "true",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"members_can_fork_private_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"web_commit_signoff_required", "true",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"advanced_security_enabled_for_new_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"dependabot_alerts_enabled_for_new_repositories", "true",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"dependabot_security_updates_enabled_for_new_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"dependency_graph_enabled_for_new_repositories", "true",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"secret_scanning_enabled_for_new_repositories", "false",
			),
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"secret_scanning_push_protection_enabled_for_new_repositories", "true",
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

	t.Run("handles minimal configuration without errors", func(t *testing.T) {
		config := `
		resource "github_organization_settings" "test" {
			billing_email = "test@example.com"
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_settings.test",
				"billing_email", "test@example.com",
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

	t.Run("comprehensive parameter testing", func(t *testing.T) {
		t.Run("test all string fields", func(t *testing.T) {
			config := `
			resource "github_organization_settings" "test" {
				billing_email = "test@example.com"
				company = "Test Company"
				email = "contact@test.com"
				twitter_username = "testorg"
				location = "Test City, Country"
				name = "Test Organization"
				description = "Test organization description"
				blog = "https://test.com/blog"
			}`

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_organization_settings.test", "billing_email", "test@example.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "company", "Test Company"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "email", "contact@test.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "twitter_username", "testorg"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "location", "Test City, Country"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "name", "Test Organization"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "description", "Test organization description"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "blog", "https://test.com/blog"),
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

		t.Run("test all security boolean fields", func(t *testing.T) {
			config := `
			resource "github_organization_settings" "test" {
				billing_email = "test@example.com"
				advanced_security_enabled_for_new_repositories = true
				dependabot_alerts_enabled_for_new_repositories = true
				dependabot_security_updates_enabled_for_new_repositories = true
				dependency_graph_enabled_for_new_repositories = true
				secret_scanning_enabled_for_new_repositories = true
				secret_scanning_push_protection_enabled_for_new_repositories = true
			}`

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_organization_settings.test", "billing_email", "test@example.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "advanced_security_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "dependabot_alerts_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "dependabot_security_updates_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "dependency_graph_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "secret_scanning_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "true"),
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

		t.Run("test repository creation fields", func(t *testing.T) {
			config := `
			resource "github_organization_settings" "test" {
				billing_email = "test@example.com"
				members_can_create_private_repositories = true
				members_can_create_internal_repositories = true
				members_can_create_pages = true
				members_can_create_public_pages = true
				members_can_create_private_pages = true
			}`

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_organization_settings.test", "billing_email", "test@example.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_private_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_internal_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_pages", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_public_pages", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_private_pages", "true"),
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

		t.Run("test other boolean fields", func(t *testing.T) {
			config := `
			resource "github_organization_settings" "test" {
				billing_email = "test@example.com"
				web_commit_signoff_required = true
				has_organization_projects = true
				has_repository_projects = true
			}`

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_organization_settings.test", "billing_email", "test@example.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "web_commit_signoff_required", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "has_organization_projects", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "has_repository_projects", "true"),
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

		t.Run("test enum fields", func(t *testing.T) {
			config := `
			resource "github_organization_settings" "test" {
				billing_email = "test@example.com"
				default_repository_permission = "write"
			}`

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_organization_settings.test", "billing_email", "test@example.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "default_repository_permission", "write"),
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

		t.Run("test comprehensive configuration", func(t *testing.T) {
			config := `
			resource "github_organization_settings" "test" {
				billing_email = "test@example.com"
				company = "Test Company"
				email = "contact@test.com"
				twitter_username = "testorg"
				location = "Test City, Country"
				name = "Test Organization"
				description = "Test organization description"
				blog = "https://test.com/blog"

				advanced_security_enabled_for_new_repositories = true
				dependabot_alerts_enabled_for_new_repositories = true
				dependabot_security_updates_enabled_for_new_repositories = true
				dependency_graph_enabled_for_new_repositories = true
				secret_scanning_enabled_for_new_repositories = true
				secret_scanning_push_protection_enabled_for_new_repositories = true

				members_can_create_private_repositories = true
				members_can_create_internal_repositories = true
				members_can_create_pages = true
				members_can_create_public_pages = true
				members_can_create_private_pages = true

				web_commit_signoff_required = true
				default_repository_permission = "write"
			}`

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_organization_settings.test", "billing_email", "test@example.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "company", "Test Company"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "email", "contact@test.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "twitter_username", "testorg"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "location", "Test City, Country"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "name", "Test Organization"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "description", "Test organization description"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "blog", "https://test.com/blog"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "advanced_security_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "dependabot_alerts_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "dependabot_security_updates_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "dependency_graph_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "secret_scanning_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_private_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_internal_repositories", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_pages", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_public_pages", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_private_pages", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "web_commit_signoff_required", "true"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "default_repository_permission", "write"),
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

		t.Run("test boolean false values for all fields", func(t *testing.T) {
			config := `
			resource "github_organization_settings" "test" {
				billing_email = "test@example.com"
				advanced_security_enabled_for_new_repositories = false
				dependabot_alerts_enabled_for_new_repositories = false
				dependabot_security_updates_enabled_for_new_repositories = false
				dependency_graph_enabled_for_new_repositories = false
				secret_scanning_enabled_for_new_repositories = false
				secret_scanning_push_protection_enabled_for_new_repositories = false
				members_can_create_private_repositories = false
				members_can_create_internal_repositories = false
				members_can_create_pages = false
				members_can_create_public_pages = false
				members_can_create_private_pages = false
				web_commit_signoff_required = false
			}`

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_organization_settings.test", "billing_email", "test@example.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "advanced_security_enabled_for_new_repositories", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "dependabot_alerts_enabled_for_new_repositories", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "dependabot_security_updates_enabled_for_new_repositories", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "dependency_graph_enabled_for_new_repositories", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "secret_scanning_enabled_for_new_repositories", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_private_repositories", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_internal_repositories", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_pages", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_public_pages", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "members_can_create_private_pages", "false"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "web_commit_signoff_required", "false"),
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

		t.Run("test enum field variations", func(t *testing.T) {
			config := `
			resource "github_organization_settings" "test" {
				billing_email = "test@example.com"
				default_repository_permission = "admin"
			}`

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_organization_settings.test", "billing_email", "test@example.com"),
				resource.TestCheckResourceAttr("github_organization_settings.test", "default_repository_permission", "admin"),
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
	})
}

func Test_buildOrganizationSettings(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name         string
		raw          map[string]any
		id           string // a non-empty ID exercises the update path
		isEnterprise bool
		want         *github.Organization
	}{
		{
			// Regression test for the create path dropping booleans configured as
			// false, which only corrected itself on a second apply through the
			// update/HasChange path. Every boolean carries a definite value on
			// create, so the full payload is asserted here.
			name: "create_includes_explicitly_false_booleans",
			raw: map[string]any{
				"billing_email":               "org@example.com",
				"has_organization_projects":   false,
				"web_commit_signoff_required": false,
			},
			want: &github.Organization{
				BillingEmail:                                   new("org@example.com"),
				HasOrganizationProjects:                        new(false),
				HasRepositoryProjects:                          new(true),
				DefaultRepoPermission:                          new("read"),
				MembersCanCreateRepos:                          new(true),
				MembersCanCreatePrivateRepos:                   new(true),
				MembersCanCreatePublicRepos:                    new(true),
				MembersCanCreatePages:                          new(true),
				MembersCanCreatePublicPages:                    new(true),
				MembersCanCreatePrivatePages:                   new(true),
				MembersCanForkPrivateRepos:                     new(false),
				WebCommitSignoffRequired:                       new(false),
				AdvancedSecurityEnabledForNewRepos:             new(false),
				DependabotAlertsEnabledForNewRepos:             new(false),
				DependabotSecurityUpdatesEnabledForNewRepos:    new(false),
				DependencyGraphEnabledForNewRepos:              new(false),
				SecretScanningEnabledForNewRepos:               new(false),
				SecretScanningPushProtectionEnabledForNewRepos: new(false),
			},
		},
		{
			name: "create_includes_true_booleans",
			raw: map[string]any{
				"billing_email":               "org@example.com",
				"web_commit_signoff_required": true,
			},
			want: &github.Organization{
				BillingEmail:                                   new("org@example.com"),
				HasOrganizationProjects:                        new(true),
				HasRepositoryProjects:                          new(true),
				DefaultRepoPermission:                          new("read"),
				MembersCanCreateRepos:                          new(true),
				MembersCanCreatePrivateRepos:                   new(true),
				MembersCanCreatePublicRepos:                    new(true),
				MembersCanCreatePages:                          new(true),
				MembersCanCreatePublicPages:                    new(true),
				MembersCanCreatePrivatePages:                   new(true),
				MembersCanForkPrivateRepos:                     new(false),
				WebCommitSignoffRequired:                       new(true),
				AdvancedSecurityEnabledForNewRepos:             new(false),
				DependabotAlertsEnabledForNewRepos:             new(false),
				DependabotSecurityUpdatesEnabledForNewRepos:    new(false),
				DependencyGraphEnabledForNewRepos:              new(false),
				SecretScanningEnabledForNewRepos:               new(false),
				SecretScanningPushProtectionEnabledForNewRepos: new(false),
			},
		},
		{
			// Only booleans are included unconditionally on create; an
			// unconfigured optional string stays absent from the payload.
			name: "create_omits_unconfigured_optional_string",
			raw: map[string]any{
				"billing_email": "org@example.com",
			},
			want: &github.Organization{
				BillingEmail:                                   new("org@example.com"),
				HasOrganizationProjects:                        new(true),
				HasRepositoryProjects:                          new(true),
				DefaultRepoPermission:                          new("read"),
				MembersCanCreateRepos:                          new(true),
				MembersCanCreatePrivateRepos:                   new(true),
				MembersCanCreatePublicRepos:                    new(true),
				MembersCanCreatePages:                          new(true),
				MembersCanCreatePublicPages:                    new(true),
				MembersCanCreatePrivatePages:                   new(true),
				MembersCanForkPrivateRepos:                     new(false),
				WebCommitSignoffRequired:                       new(false),
				AdvancedSecurityEnabledForNewRepos:             new(false),
				DependabotAlertsEnabledForNewRepos:             new(false),
				DependabotSecurityUpdatesEnabledForNewRepos:    new(false),
				DependencyGraphEnabledForNewRepos:              new(false),
				SecretScanningEnabledForNewRepos:               new(false),
				SecretScanningPushProtectionEnabledForNewRepos: new(false),
			},
		},
		{
			// On update a field is only sent when it has changed, so the
			// explicitly configured has_organization_projects is dropped. The
			// attributes that do survive are an artifact of TestResourceDataRaw:
			// one left out of raw reads as the zero value from state but as its
			// schema default from d.Get, so every attribute defaulting to a
			// non-zero value (true, "read") looks changed here.
			name: "update_omits_unchanged_booleans",
			raw: map[string]any{
				"billing_email":             "org@example.com",
				"has_organization_projects": false,
			},
			id: "example-org",
			want: &github.Organization{
				BillingEmail:                 new("org@example.com"),
				HasRepositoryProjects:        new(true),
				DefaultRepoPermission:        new("read"),
				MembersCanCreateRepos:        new(true),
				MembersCanCreatePrivateRepos: new(true),
				MembersCanCreatePublicRepos:  new(true),
				MembersCanCreatePages:        new(true),
				MembersCanCreatePublicPages:  new(true),
				MembersCanCreatePrivatePages: new(true),
			},
		},
		{
			name: "enterprise_create_includes_internal_repositories_boolean",
			raw: map[string]any{
				"billing_email": "org@example.com",
				"members_can_create_internal_repositories": false,
			},
			isEnterprise: true,
			want: &github.Organization{
				BillingEmail:                                   new("org@example.com"),
				HasOrganizationProjects:                        new(true),
				HasRepositoryProjects:                          new(true),
				DefaultRepoPermission:                          new("read"),
				MembersCanCreateRepos:                          new(true),
				MembersCanCreateInternalRepos:                  new(false),
				MembersCanCreatePrivateRepos:                   new(true),
				MembersCanCreatePublicRepos:                    new(true),
				MembersCanCreatePages:                          new(true),
				MembersCanCreatePublicPages:                    new(true),
				MembersCanCreatePrivatePages:                   new(true),
				MembersCanForkPrivateRepos:                     new(false),
				WebCommitSignoffRequired:                       new(false),
				AdvancedSecurityEnabledForNewRepos:             new(false),
				DependabotAlertsEnabledForNewRepos:             new(false),
				DependabotSecurityUpdatesEnabledForNewRepos:    new(false),
				DependencyGraphEnabledForNewRepos:              new(false),
				SecretScanningEnabledForNewRepos:               new(false),
				SecretScanningPushProtectionEnabledForNewRepos: new(false),
			},
		},
		{
			// The enterprise-only attribute must never reach the payload for a
			// non-enterprise organization, even when it is configured.
			name: "non_enterprise_create_omits_internal_repositories_boolean",
			raw: map[string]any{
				"billing_email": "org@example.com",
				"members_can_create_internal_repositories": false,
			},
			want: &github.Organization{
				BillingEmail:                                   new("org@example.com"),
				HasOrganizationProjects:                        new(true),
				HasRepositoryProjects:                          new(true),
				DefaultRepoPermission:                          new("read"),
				MembersCanCreateRepos:                          new(true),
				MembersCanCreatePrivateRepos:                   new(true),
				MembersCanCreatePublicRepos:                    new(true),
				MembersCanCreatePages:                          new(true),
				MembersCanCreatePublicPages:                    new(true),
				MembersCanCreatePrivatePages:                   new(true),
				MembersCanForkPrivateRepos:                     new(false),
				WebCommitSignoffRequired:                       new(false),
				AdvancedSecurityEnabledForNewRepos:             new(false),
				DependabotAlertsEnabledForNewRepos:             new(false),
				DependabotSecurityUpdatesEnabledForNewRepos:    new(false),
				DependencyGraphEnabledForNewRepos:              new(false),
				SecretScanningEnabledForNewRepos:               new(false),
				SecretScanningPushProtectionEnabledForNewRepos: new(false),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, resourceGithubOrganizationSettings().Schema, tt.raw)
			if tt.id != "" {
				d.SetId(tt.id)
			}

			got := buildOrganizationSettings(d, tt.isEnterprise)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("buildOrganizationSettings() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
