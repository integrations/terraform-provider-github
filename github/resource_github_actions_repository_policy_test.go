package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestParseActionsRepositoryPolicyID(t *testing.T) {
	repo, id, err := parseActionsRepositoryPolicyID("some-repo:42")
	if err != nil {
		t.Fatal(err)
	}
	if repo != "some-repo" || id != 42 {
		t.Errorf("got (%q, %d), want (\"some-repo\", 42)", repo, id)
	}

	if _, _, err := parseActionsRepositoryPolicyID("some-repo"); err == nil {
		t.Error("expected error for missing policy ID")
	}
	if _, _, err := parseActionsRepositoryPolicyID("some-repo:abc"); err == nil {
		t.Error("expected error for non-numeric policy ID")
	}
}

func TestAccGithubActionsRepositoryPolicy(t *testing.T) {
	t.Run("creates, updates and imports a repository Actions policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-policy-%s", testResourcePrefix, randomID)

		config := `
resource "github_repository" "test" {
	name       = "%s"
	visibility = "public"
}

resource "github_actions_repository_policy" "test" {
	repository  = github_repository.test.name
	name        = "tf-acc-policy"
	enforcement = "active"

	workflow_path {
		include = [".github/workflows/triage.yml"]
	}

	allowed_events = [%s]
}
`
		resourceName := "github_actions_repository_policy.test"

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, `"pull_request", "pull_request_target"`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_events"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.StringExact("pull_request"),
							knownvalue.StringExact("pull_request_target"),
						})),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("workflow_path").AtSliceIndex(0).AtMapKey("include"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact(".github/workflows/triage.yml"),
						})),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, `"pull_request_target"`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allowed_events"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.StringExact("pull_request_target"),
						})),
					},
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
