package github

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAgentsOrganizationSecretRepositoryPagination(t *testing.T) {
	t.Parallel()
	meta := agentSecretsPaginationTestOwner(t)
	d := schema.TestResourceDataRaw(t, resourceGithubAgentsOrganizationSecretRepository().Schema, map[string]any{"secret_name": "TEST", "repository_id": 2})
	d.SetId("TEST:2")
	if diagnostics := resourceGithubAgentsOrganizationSecretRepositoryRead(t.Context(), d, meta); diagnostics.HasError() || d.Id() == "" {
		t.Fatalf("association on second page was lost: %v", diagnostics)
	}
}

func TestAccGithubAgentsOrganizationSecretRepository(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)
	repo := mustCreateTestRepository(t)
	other := mustCreateTestRepository(t)
	name := strings.ToUpper(strings.ReplaceAll(testResourcePrefix, "-", "_") + acctest.RandString(testRandomIDLength))
	config := func(value string, id int64) string {
		return fmt.Sprintf(`
resource "github_agents_organization_secret" "test" {
  secret_name = %q
  value = %q
  visibility = "selected"
}
resource "github_agents_organization_secret_repository" "test" {
  secret_name = github_agents_organization_secret.test.secret_name
  repository_id = %d
}`, name, value, id)
	}
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{Config: config("initial", repo.GetID())},
			{ResourceName: "github_agents_organization_secret_repository.test", ImportState: true, ImportStateVerify: true},
			{Config: config("rotated", repo.GetID())},
			{Config: config("rotated", repo.GetID()), PlanOnly: true},
			{Config: config("rotated", other.GetID()), ConfigStateChecks: []statecheck.StateCheck{statecheck.ExpectKnownValue("github_agents_organization_secret_repository.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(other.GetID()))}},
		},
		CheckDestroy: func(_ *terraform.State) error {
			_, _, err := testAccConf.meta.v3client.Agents.GetOrgSecret(t.Context(), testAccConf.meta.name, name)
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == 404 {
				return nil
			}
			if err == nil {
				return fmt.Errorf("agent secret still exists after destroy")
			}
			return fmt.Errorf("agent secret still exists after destroy: %w", err)
		},
	})
}
