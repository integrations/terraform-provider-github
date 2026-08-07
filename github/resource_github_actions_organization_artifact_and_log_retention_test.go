package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// mustRestoreOrganizationArtifactAndLogRetention records the organization's current retention
// period and restores it once the test finishes. Destroying the resource resets to GitHub's
// default rather than to whatever the organization had configured beforehand.
func mustRestoreOrganizationArtifactAndLogRetention(t *testing.T) {
	t.Helper()

	client := testAccConf.meta.v3client
	orgName := testAccConf.meta.name

	original, _, err := client.Actions.GetArtifactAndLogRetentionPeriodInOrganization(t.Context(), orgName)
	if err != nil {
		t.Fatalf("failed to read organization artifact and log retention: %v", err)
	}

	t.Cleanup(func() {
		if _, err := client.Actions.UpdateArtifactAndLogRetentionPeriodInOrganization(context.Background(), orgName, github.ArtifactPeriodOpt{
			Days: new(original.GetDays()),
		}); err != nil {
			t.Errorf("failed to restore organization artifact and log retention to %d: %v", original.GetDays(), err)
		}
	})
}

func TestAccGithubActionsOrganizationArtifactAndLogRetention(t *testing.T) {
	// IMPORTANT: Do not run these tests in parallel as they modify the organization state.

	t.Run("full_lifecycle", func(t *testing.T) {
		config := `
resource "github_actions_organization_artifact_and_log_retention" "test" {
  days = %d
}
`

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessHasOrgs(t)
				mustRestoreOrganizationArtifactAndLogRetention(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, 7),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_artifact_and_log_retention.test", tfjsonpath.New("days"), knownvalue.Int64Exact(7)),
						statecheck.ExpectKnownValue("github_actions_organization_artifact_and_log_retention.test", tfjsonpath.New("maximum_allowed_days"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, 30),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_organization_artifact_and_log_retention.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_artifact_and_log_retention.test", tfjsonpath.New("days"), knownvalue.Int64Exact(30)),
					},
				},
				{
					ResourceName:      "github_actions_organization_artifact_and_log_retention.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
