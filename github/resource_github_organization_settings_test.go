package github

import (
	"fmt"
	"testing"

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

		wantBillingEmail                  *string
		wantCompany                       *string
		wantHasOrganizationProjects       *bool
		wantWebCommitSignoffRequired      *bool
		wantMembersCanCreateInternalRepos *bool
		wantOmitted                       []string
	}{
		{
			// Regression test for the create path dropping booleans configured as
			// false, which only corrected itself on a second apply through the
			// update/HasChange path.
			name: "create_includes_explicitly_false_booleans",
			raw: map[string]any{
				"billing_email":               "org@example.com",
				"has_organization_projects":   false,
				"web_commit_signoff_required": false,
			},
			wantBillingEmail:             new("org@example.com"),
			wantHasOrganizationProjects:  new(false),
			wantWebCommitSignoffRequired: new(false),
		},
		{
			name: "create_includes_true_booleans",
			raw: map[string]any{
				"billing_email":               "org@example.com",
				"web_commit_signoff_required": true,
			},
			wantBillingEmail:             new("org@example.com"),
			wantWebCommitSignoffRequired: new(true),
		},
		{
			// Only booleans are included unconditionally on create; an
			// unconfigured optional string stays omitted.
			name: "create_omits_unconfigured_optional_string",
			raw: map[string]any{
				"billing_email": "org@example.com",
			},
			wantBillingEmail: new("org@example.com"),
			wantOmitted:      []string{"company"},
		},
		{
			name: "update_omits_unchanged_booleans",
			raw: map[string]any{
				"billing_email":             "org@example.com",
				"has_organization_projects": false,
			},
			id:          "example-org",
			wantOmitted: []string{"has_organization_projects"},
		},
		{
			name: "enterprise_create_includes_internal_repositories_boolean",
			raw: map[string]any{
				"billing_email": "org@example.com",
				"members_can_create_internal_repositories": false,
			},
			isEnterprise:                      true,
			wantMembersCanCreateInternalRepos: new(false),
		},
		{
			name: "non_enterprise_create_omits_internal_repositories_boolean",
			raw: map[string]any{
				"billing_email": "org@example.com",
				"members_can_create_internal_repositories": false,
			},
			wantOmitted: []string{"members_can_create_internal_repositories"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, resourceGithubOrganizationSettings().Schema, tt.raw)
			if tt.id != "" {
				d.SetId(tt.id)
			}

			got := buildOrganizationSettings(d, tt.isEnterprise)

			assertStringPtr(t, "billing_email", got.BillingEmail, tt.wantBillingEmail)
			assertStringPtr(t, "company", got.Company, tt.wantCompany)
			assertBoolPtr(t, "has_organization_projects", got.HasOrganizationProjects, tt.wantHasOrganizationProjects)
			assertBoolPtr(t, "web_commit_signoff_required", got.WebCommitSignoffRequired, tt.wantWebCommitSignoffRequired)
			assertBoolPtr(t, "members_can_create_internal_repositories", got.MembersCanCreateInternalRepos, tt.wantMembersCanCreateInternalRepos)

			omitted := map[string]bool{
				"billing_email":                            got.BillingEmail == nil,
				"company":                                  got.Company == nil,
				"has_organization_projects":                got.HasOrganizationProjects == nil,
				"web_commit_signoff_required":              got.WebCommitSignoffRequired == nil,
				"members_can_create_internal_repositories": got.MembersCanCreateInternalRepos == nil,
			}
			for _, field := range tt.wantOmitted {
				isOmitted, ok := omitted[field]
				if !ok {
					t.Fatalf("wantOmitted references unknown field %q", field)
				}
				if !isOmitted {
					t.Errorf("%s should have been omitted from the payload", field)
				}
			}
		})
	}
}

// assertBoolPtr checks a boolean field of the built payload. A nil want means
// the field is not asserted by this test case; use wantOmitted to require that
// a field is absent.
func assertBoolPtr(t *testing.T, field string, got, want *bool) {
	t.Helper()

	if want == nil {
		return
	}
	if got == nil {
		t.Errorf("%s was omitted from the payload, want %v", field, *want)
		return
	}
	if *got != *want {
		t.Errorf("%s = %v, want %v", field, *got, *want)
	}
}

// assertStringPtr is the string counterpart of assertBoolPtr.
func assertStringPtr(t *testing.T, field string, got, want *string) {
	t.Helper()

	if want == nil {
		return
	}
	if got == nil {
		t.Errorf("%s was omitted from the payload, want %q", field, *want)
		return
	}
	if *got != *want {
		t.Errorf("%s = %q, want %q", field, *got, *want)
	}
}
