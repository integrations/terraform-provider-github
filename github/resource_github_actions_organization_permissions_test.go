package github

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsOrganizationPermissions(t *testing.T) {
	// IMPORTANT: Do not run these tests in parallel as they modify the organization state.

	skipUnlessHasOrgs(t)

	t.Run("full_lifecycle", func(t *testing.T) {
		repo := mustCreateTestRepository(t)

		configNone := `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "none"
}
`

		configMinimal := `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "%s"

  allowed_actions = "%s"
}
`

		configFull := `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "selected"

  enabled_repositories_config {
    repository_ids = [%v]
  }

  allowed_actions = "selected"

  allowed_actions_config {
    github_owned_allowed = true
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }

  sha_pinning_required = true
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configMinimal, "all", "all"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_permissions.test", tfjsonpath.New("allowed_actions_config"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("github_actions_organization_permissions.test", tfjsonpath.New("enabled_repositories_config"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("github_actions_organization_permissions.test", tfjsonpath.New("sha_pinning_required"), knownvalue.NotNull()),
					},
				},
				{
					Config: configNone,
				},
				{
					Config: fmt.Sprintf(configMinimal, "all", "local_only"),
				},
				{
					Config: fmt.Sprintf(configFull, ""),
				},
				{
					Config: fmt.Sprintf(configFull, repo.GetID()),
				},
				{
					ResourceName:      "github_actions_organization_permissions.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("selected_repos_unmanaged", func(t *testing.T) {
		repo := mustCreateTestRepository(t)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "selected"

  enabled_repositories_config {
    repository_ids = [%v]
  }
}`, repo.GetID()),
				},
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "selected"
}`,
					Check: func(state *terraform.State) error {
						// Need to wait for GitHub to catch up
						time.Sleep(5 * time.Second)

						found := false
						for _, err := range testAccConf.meta.v3client.Actions.ListEnabledReposInOrgIter(t.Context(), testAccConf.meta.name, &github.ListOptions{PerPage: testAccConf.meta.maxPerPage}) {
							if err != nil {
								return err
							}

							found = true
							break
						}

						if !found {
							return fmt.Errorf("seleced repositories no longer set")
						}

						return nil
					},
				},
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "selected"

  enabled_repositories_config {
    repository_ids = []
  }
}`,
					ExpectNonEmptyPlan: true,
					Check: func(state *terraform.State) error {
						// Need to wait for GitHub to catch up
						time.Sleep(5 * time.Second)

						found := false
						for _, err := range testAccConf.meta.v3client.Actions.ListEnabledReposInOrgIter(t.Context(), testAccConf.meta.name, &github.ListOptions{PerPage: testAccConf.meta.maxPerPage}) {
							if err != nil {
								return err
							}

							found = true
							break
						}

						if found {
							return fmt.Errorf("seleced repositories should be empty")
						}

						return nil
					},
				},
			},
		})
	})

	t.Run("selected_allowed_actions_unmanaged", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
	resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  allowed_actions = "selected"
	}`,
				},
			},
		})
	})

	t.Run("disabled_actions", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "none"

  allowed_actions = "all"
}`,
					ExpectError: regexp.MustCompile(regexp.QuoteMeta("allowed_actions is only valid when enabled_repositories is set to all or selected")),
				},
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "none"
}`,
				},
			},
		})
	})

	t.Run("sha_pinning_can_be_disabled", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  allowed_actions = "all"

  sha_pinning_required = true
}`,
				},
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  allowed_actions = "all"

  sha_pinning_required = false
}`,
				},
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  allowed_actions = "all"

  sha_pinning_required = true
}`,
				},
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  allowed_actions = "all"
}`,
				},
			},
		})
	})

	t.Run("sha_pinning_can_be_enabled", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  allowed_actions = "all"
}`,
				},
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  allowed_actions = "all"

  sha_pinning_required = true
}`,
				},
			},
		})
	})

	t.Run("invalid_repository_config_set", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
			resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  enabled_repositories_config {
    repository_ids = []
  }

  allowed_actions = "all"
}`,
					ExpectError: regexp.MustCompile(regexp.QuoteMeta("enabled_repositories_config is only valid when enabled_repositories is set to selected")),
				},
			},
		})
	})

	t.Run("invalid_actions_config_set", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
			resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  allowed_actions = "all"

  allowed_actions_config {
    github_owned_allowed = true
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }
}`,
					ExpectError: regexp.MustCompile(regexp.QuoteMeta("allowed_actions_config is only valid when allowed_actions is set to selected")),
				},
			},
		})
	})

	t.Run("invalid_actions_config_missing", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
resource "github_actions_organization_permissions" "test" {
  enabled_repositories = "all"

  allowed_actions = "selected"
}`,
					ExpectError: regexp.MustCompile(regexp.QuoteMeta("allowed_actions_config is required when allowed_actions is set to selected")),
				},
			},
		})
	})
}
