package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsRepositoryArtifactAndLogRetention(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_actions_repository_artifact_and_log_retention" "test" {
  repository = "%s"
  days       = %%d
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, 7),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_repository_artifact_and_log_retention.test", tfjsonpath.New("days"), knownvalue.Int64Exact(7)),
						statecheck.ExpectKnownValue("github_actions_repository_artifact_and_log_retention.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_repository_artifact_and_log_retention.test", tfjsonpath.New("maximum_allowed_days"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, 30),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_repository_artifact_and_log_retention.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_repository_artifact_and_log_retention.test", tfjsonpath.New("days"), knownvalue.Int64Exact(30)),
					},
				},
				{
					ResourceName:      "github_actions_repository_artifact_and_log_retention.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("rejects days outside the allowed range", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_actions_repository_artifact_and_log_retention" "test" {
  repository = "%s"
  days       = 0
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`expected days to be in the range \(1 - 400\)`),
				},
			},
		})
	})

	t.Run("rejects days above the allowed ceiling", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		// The organization caps what a repository may set, and the API rejects anything
		// higher with a 409 rather than silently clamping it. The ceiling depends on the
		// test organization, so it is read rather than assumed.
		retention, _, err := testAccConf.meta.v3client.Repositories.GetArtifactAndLogRetentionPeriod(t.Context(), testAccConf.meta.name, repo.GetName())
		if err != nil {
			t.Fatalf("failed to read artifact and log retention: %v", err)
		}
		if retention.GetMaximumAllowedDays() >= 400 {
			t.Skip("organization allows the schema maximum, so no API-rejected value exists")
		}

		config := fmt.Sprintf(`
resource "github_actions_repository_artifact_and_log_retention" "test" {
  repository = "%s"
  days       = %d
}
`, repo.GetName(), retention.GetMaximumAllowedDays()+1)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Artifact and log retention settings are limited`),
				},
			},
		})
	})
}
