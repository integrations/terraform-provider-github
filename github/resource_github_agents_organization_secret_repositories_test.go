package github

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubAgentsOrganizationSecretRepositories(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)
	repo := mustCreateTestRepository(t)
	other := mustCreateTestRepository(t)
	name := strings.ToUpper(strings.ReplaceAll(testResourcePrefix, "-", "_") + acctest.RandString(testRandomIDLength))
	config := func(value, ids string) string {
		return fmt.Sprintf(`
resource "github_agents_organization_secret" "test" {
  secret_name = %q
  value = %q
  visibility = "selected"
}
resource "github_agents_organization_secret_repositories" "test" {
  secret_name = github_agents_organization_secret.test.secret_name
  selected_repository_ids = [%s]
}`, name, value, ids)
	}
	checks := []statecheck.StateCheck{statecheck.ExpectKnownValue("github_agents_organization_secret_repositories.test", tfjsonpath.New("selected_repository_ids"), knownvalue.SetExact([]knownvalue.Check{knownvalue.Int64Exact(repo.GetID()), knownvalue.Int64Exact(other.GetID())}))}
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{Config: config("initial", fmt.Sprintf("%d", repo.GetID()))},
			{ResourceName: "github_agents_organization_secret_repositories.test", ImportState: true, ImportStateVerify: true},
			{Config: config("initial", fmt.Sprintf("%d, %d", repo.GetID(), other.GetID())), ConfigStateChecks: checks},
			{Config: config("rotated", fmt.Sprintf("%d, %d", repo.GetID(), other.GetID())), ConfigStateChecks: checks},
			// A post-rotation plan catches access erased by the secret PUT, before refresh can repair it.
			{Config: config("rotated", fmt.Sprintf("%d, %d", repo.GetID(), other.GetID())), PlanOnly: true},
			{Config: config("rotated", ""), ConfigStateChecks: []statecheck.StateCheck{statecheck.ExpectKnownValue("github_agents_organization_secret_repositories.test", tfjsonpath.New("selected_repository_ids"), knownvalue.SetExact([]knownvalue.Check{}))}},
			{
				Config: fmt.Sprintf(`resource "github_agents_organization_secret" "test" {
  secret_name = %q
  value = "rotated"
  visibility = "selected"
}`, name),
				ConfigStateChecks: []statecheck.StateCheck{statecheck.ExpectKnownValue("github_agents_organization_secret.test", tfjsonpath.New("secret_name"), knownvalue.StringExact(name))},
			},
		},
	})
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
