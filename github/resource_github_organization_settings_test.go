package github

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

// testResourceDataWithConfig builds ResourceData carrying both the attributes
// d.Get reads and the raw configuration isConfigured reads. Attributes absent
// from raw are recorded as null in the configuration, which is how an unset
// attribute is told apart from one explicitly configured as false, and carry
// their schema Default in the attributes, which is what the planned state
// holds during a real create and what d.GetOk answers from.
//
// schema.TestResourceDataRaw cannot be used here: it leaves the raw
// configuration null, and that is the one distinction these cases exercise.
func testResourceDataWithConfig(t *testing.T, raw map[string]any) *schema.ResourceData {
	t.Helper()

	res := resourceGithubOrganizationSettings()
	attributes := map[string]string{}
	config := map[string]cty.Value{}

	for name, attrSchema := range res.Schema {
		value, configured := raw[name]

		switch attrSchema.Type {
		case schema.TypeBool:
			if !configured {
				config[name] = cty.NullVal(cty.Bool)
				if def, ok := attrSchema.Default.(bool); ok {
					attributes[name] = strconv.FormatBool(def)
				}

				continue
			}

			b, ok := value.(bool)
			if !ok {
				t.Fatalf("attribute %q is a bool in the schema but %T in the test case", name, value)
			}
			attributes[name] = strconv.FormatBool(b)
			config[name] = cty.BoolVal(b)
		case schema.TypeString:
			if !configured {
				config[name] = cty.NullVal(cty.String)
				if def, ok := attrSchema.Default.(string); ok {
					attributes[name] = def
				}

				continue
			}

			s, ok := value.(string)
			if !ok {
				t.Fatalf("attribute %q is a string in the schema but %T in the test case", name, value)
			}
			attributes[name] = s
			config[name] = cty.StringVal(s)
		default:
			t.Fatalf("attribute %q has unsupported type %s; extend this helper", name, attrSchema.Type)
		}
	}

	return res.Data(&terraform.InstanceState{
		Attributes: attributes,
		RawConfig:  cty.ObjectVal(config),
	})
}

// createBaseline is the payload a configuration setting nothing but
// billing_email produces. Attributes whose schema Default is a non-zero value
// are reported by d.GetOk even when the practitioner omits them, so they are
// sent on every create, exactly as #2807 did. Attributes defaulting to false
// or "" stay out unless configured, which keeps the request narrow enough to
// avoid #2305.
func createBaseline() *github.Organization {
	return &github.Organization{
		BillingEmail:                 new("org@example.com"),
		HasOrganizationProjects:      new(true),
		HasRepositoryProjects:        new(true),
		DefaultRepoPermission:        new("read"),
		MembersCanCreateRepos:        new(true),
		MembersCanCreatePrivateRepos: new(true),
		MembersCanCreatePublicRepos:  new(true),
		MembersCanCreatePages:        new(true),
		MembersCanCreatePublicPages:  new(true),
		MembersCanCreatePrivatePages: new(true),
	}
}

// updateBaseline is the update-path counterpart of createBaseline. The update
// payload is driven by d.HasChange, and under schema.TestResourceDataRaw an
// attribute left out of raw reads as the zero value from state but as its
// schema default from d.Get, so every attribute whose default is non-zero
// looks changed. These entries are that artifact, not a behaviour the
// provider intends.
func updateBaseline() *github.Organization {
	return &github.Organization{
		BillingEmail:                 new("org@example.com"),
		HasOrganizationProjects:      new(true),
		HasRepositoryProjects:        new(true),
		DefaultRepoPermission:        new("read"),
		MembersCanCreateRepos:        new(true),
		MembersCanCreatePrivateRepos: new(true),
		MembersCanCreatePublicRepos:  new(true),
		MembersCanCreatePages:        new(true),
		MembersCanCreatePublicPages:  new(true),
		MembersCanCreatePrivatePages: new(true),
	}
}

func Test_organizationSettingsForCreate(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name         string
		raw          map[string]any
		isEnterprise bool
		want         *github.Organization
	}{
		{
			// The bug reported in #3493: a boolean configured as false used to
			// be dropped from the create payload, so the API applied its own
			// default and the first apply silently produced the opposite of the
			// configuration. Both attributes here are explicitly false and both
			// must reach the payload.
			name: "create_includes_explicitly_false_booleans",
			raw: map[string]any{
				"billing_email":               "org@example.com",
				"has_organization_projects":   false,
				"web_commit_signoff_required": false,
			},
			want: func() *github.Organization {
				want := createBaseline()
				want.HasOrganizationProjects = new(false)
				want.WebCommitSignoffRequired = new(false)
				return want
			}(),
		},
		{
			// An attribute the user never wrote is still sent when its schema
			// default is a definite true, as #2807 already did. Leaving it out
			// would let the org keep a differing value, so the first plan after
			// create would not be clean.
			name: "create_includes_unconfigured_default_true_boolean",
			raw: map[string]any{
				"billing_email": "org@example.com",
			},
			want: createBaseline(),
		},
		{
			name: "create_includes_explicitly_true_booleans",
			raw: map[string]any{
				"billing_email":               "org@example.com",
				"web_commit_signoff_required": true,
			},
			want: func() *github.Organization {
				want := createBaseline()
				want.WebCommitSignoffRequired = new(true)
				return want
			}(),
		},
		{
			name: "create_omits_unconfigured_optional_string",
			raw: map[string]any{
				"billing_email": "org@example.com",
			},
			want: createBaseline(),
		},
		{
			name: "enterprise_create_includes_internal_repositories_boolean",
			raw: map[string]any{
				"billing_email": "org@example.com",
				"members_can_create_internal_repositories": true,
			},
			isEnterprise: true,
			want: func() *github.Organization {
				want := createBaseline()
				want.MembersCanCreateInternalRepos = new(true)
				return want
			}(),
		},
		{
			// The enterprise-only attribute must never reach the payload for a
			// non-enterprise organization, even when it is configured.
			name: "non_enterprise_create_omits_internal_repositories_boolean",
			raw: map[string]any{
				"billing_email": "org@example.com",
				"members_can_create_internal_repositories": true,
			},
			want: createBaseline(),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := testResourceDataWithConfig(t, tt.raw)

			got := organizationSettingsForCreate(d, tt.isEnterprise)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("organizationSettingsForCreate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_organizationSettingsForUpdate(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name         string
		raw          map[string]any
		isEnterprise bool
		want         *github.Organization
	}{
		{
			// Only changed attributes are sent, which is what keeps
			// unconfigured attributes out of the request (#2305). The
			// attributes that do survive here are an artifact of
			// TestResourceDataRaw: one left out of raw reads as the zero value
			// from state but as its schema default from d.Get, so every
			// attribute defaulting to a non-zero value looks changed.
			name: "update_omits_unchanged_booleans",
			raw: map[string]any{
				"billing_email":             "org@example.com",
				"has_organization_projects": false,
			},
			want: func() *github.Organization {
				want := updateBaseline()
				want.HasOrganizationProjects = nil
				return want
			}(),
		},
		{
			// A configured optional string is a change against empty state, so
			// it is included; the create path reaches the same result through
			// GetOk instead.
			name: "update_includes_changed_optional_string",
			raw: map[string]any{
				"billing_email": "org@example.com",
				"description":   "example description",
			},
			want: func() *github.Organization {
				want := updateBaseline()
				want.Description = new("example description")
				return want
			}(),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, resourceGithubOrganizationSettings().Schema, tt.raw)
			d.SetId("example-org")

			got := organizationSettingsForUpdate(d, tt.isEnterprise)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("organizationSettingsForUpdate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
