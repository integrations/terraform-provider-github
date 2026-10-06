package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAgentsOrganizationSecretRepositoriesDeleteEmptyList(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			IDs []int64 `json:"selected_repository_ids"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		if body.IDs == nil || len(body.IDs) != 0 {
			t.Error("deleting repository access must send [], not null")
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()
	meta := &Owner{name: "test", IsOrganization: true, v3client: mustCreateTestGitHubClient(t, server.URL)}
	d := schema.TestResourceDataRaw(t, resourceGithubAgentsOrganizationSecretRepositories().Schema, map[string]any{"secret_name": "TEST"})
	if diagnostics := resourceGithubAgentsOrganizationSecretRepositoriesDelete(t.Context(), d, meta); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
}

func TestAgentsOrganizationSecretRepositoriesPagination(t *testing.T) {
	t.Parallel()
	meta := agentSecretsPaginationTestOwner(t)
	d := schema.TestResourceDataRaw(t, resourceGithubAgentsOrganizationSecretRepositories().Schema, map[string]any{"secret_name": "TEST"})
	d.SetId("TEST")
	if diagnostics := resourceGithubAgentsOrganizationSecretRepositoriesRead(t.Context(), d, meta); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	ids, _ := d.Get("selected_repository_ids").(*schema.Set)
	if ids.Len() != 2 {
		t.Fatal("repository listing did not follow pagination")
	}
}

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
