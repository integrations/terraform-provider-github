package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate(t *testing.T) {
	t.Parallel()

	t.Run("creates repository oidc subject claim customization template without error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-oidc-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name = "%s"
			visibility = "private"
		}

		resource "github_actions_repository_oidc_subject_claim_customization_template" "test" {
			repository = github_repository.test.name
			use_default = false
			include_claim_keys = ["repo", "context", "job_workflow_ref"]
		}`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"use_default", "false",
			),
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"include_claim_keys.#", "3",
			),
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"include_claim_keys.0", "repo",
			),
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"include_claim_keys.1", "context",
			),
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"include_claim_keys.2", "job_workflow_ref",
			),
		)
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates repository oidc subject claim customization template without error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-oidc-%s", testResourcePrefix, randomID)
		configTemplate := `
		resource "github_repository" "test" {
			name = "%s"
			visibility = "private"
		}

		resource "github_actions_repository_oidc_subject_claim_customization_template" "test" {
			repository = github_repository.test.name
			use_default = %t
			include_claim_keys = %s
		}`

		claims := `["repository_owner_id", "run_id", "workflow"]`
		updatedClaims := `["actor", "actor_id", "head_ref", "repository"]`

		resetToDefaultConfigTemplate := `
		resource "github_repository" "test" {
			name = "%s"
			visibility = "private"
		}

		resource "github_actions_repository_oidc_subject_claim_customization_template" "test" {
			repository = github_repository.test.name
			use_default = true
		}
`

		configs := map[string]string{
			"before": fmt.Sprintf(configTemplate, repoName, false, claims),

			"after": fmt.Sprintf(configTemplate, repoName, false, updatedClaims),

			"reset_to_default": fmt.Sprintf(resetToDefaultConfigTemplate, repoName),
		}
		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"use_default", "false",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.#", "3",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.0", "repository_owner_id",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.1", "run_id",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.2", "workflow",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"use_default", "false",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.#", "4",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.0", "actor",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.1", "actor_id",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.2", "head_ref",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.3", "repository",
				),
			),
			"reset_to_default": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"use_default", "true",
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_oidc_subject_claim_customization_template.test",
					"include_claim_keys.#", "0",
				),
			),
		}
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
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
				{
					Config: configs["reset_to_default"],
					Check:  checks["reset_to_default"],
				},
			},
		})
	})

	t.Run("imports repository oidc subject claim customization template without error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-oidc-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name = "%s"
			visibility = "private"
		}
		resource "github_actions_repository_oidc_subject_claim_customization_template" "test" {
			repository = github_repository.test.name
			use_default = false
			include_claim_keys = ["repository_owner_id", "run_id", "workflow"]
		}`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"use_default", "false",
			),
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"include_claim_keys.#", "3",
			),
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"include_claim_keys.0", "repository_owner_id",
			),
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"include_claim_keys.1", "run_id",
			),
			resource.TestCheckResourceAttr(
				"github_actions_repository_oidc_subject_claim_customization_template.test",
				"include_claim_keys.2", "workflow",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:      "github_actions_repository_oidc_subject_claim_customization_template.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("updates_renamed_repo", func(t *testing.T) {
		t.Parallel()
		skipUnauthenticated(t)

		repo := mustCreateTestRepository(t)
		newRepoName := fmt.Sprintf("%s-updated", repo.GetName())

		config := `
resource "github_actions_repository_oidc_subject_claim_customization_template" "test" {
  repository         = "%s"
  use_default        = false
  include_claim_keys = ["repo", "context"]
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName()),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_repository_oidc_subject_claim_customization_template.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
					},
				},
				{
					PreConfig: func() {
						mustRenameTestRepository(t, repo, newRepoName)
					},
					Config: fmt.Sprintf(config, newRepoName),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_repository_oidc_subject_claim_customization_template.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_repository_oidc_subject_claim_customization_template.test", tfjsonpath.New("repository"), knownvalue.StringExact(newRepoName)),
						statecheck.ExpectKnownValue("github_actions_repository_oidc_subject_claim_customization_template.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
					},
				},
			},
		})
	})

	t.Run("recreates_changed_repo", func(t *testing.T) {
		t.Parallel()
		skipUnauthenticated(t)

		repo := mustCreateTestRepository(t)
		repo2 := mustCreateTestRepository(t)

		config := `
resource "github_actions_repository_oidc_subject_claim_customization_template" "test" {
  repository         = "%s"
  use_default        = false
  include_claim_keys = ["repo", "context"]
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName()),
				},
				{
					Config: fmt.Sprintf(config, repo2.GetName()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_repository_oidc_subject_claim_customization_template.test", plancheck.ResourceActionReplace),
						},
					},
				},
			},
		})
	})
}
