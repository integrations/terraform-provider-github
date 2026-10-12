package github

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestAccGithubIssueLabels(t *testing.T) {
	t.Parallel()
	skipUnauthenticated(t)

	t.Run("authoritatively_overtakes_existing_labels", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		time.Sleep(2 * time.Second)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				// 1. Check all the labels created by default are removed when the resource is added
				{
					Config: fmt.Sprintf(`
resource "github_issue_labels" "test" {
  repository = "%s"
}
`, repo.GetName()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_issue_labels.test", plancheck.ResourceActionCreate),
						},
					},
				},
				// 2. Check a label can be created
				{
					Config: fmt.Sprintf(`
resource "github_issue_labels" "test" {
  repository = "%s"

  label {
    name        = "foo"
    color       = "000000"
    description = "foo"
  }
}
`, repo.GetName()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_issue_labels.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				// 3. Check a label can be recreated
				{
					Config: fmt.Sprintf(`
resource "github_issue_labels" "test" {
  repository = "%s"

  label {
    name        = "Foo"
    color       = "000000"
    description = "foo"
  }
}
`, repo.GetName()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_issue_labels.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				// 4. Check multiple labels can be created
				{
					Config: fmt.Sprintf(`
resource "github_issue_labels" "test" {
  repository = "%s"

  label {
    name        = "Foo"
    color       = "000000"
    description = "foo"
  }

  label {
    name        = "bar"
    color       = "000000"
    description = "bar"
  }

  label {
    name        = "baz"
    color       = "000000"
    description = "baz"
  }
}
`, repo.GetName()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_issue_labels.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				// 5. Check the labels are destroyed along with the resource
				{
					Config: `
removed {
  from = github_issue_labels.test

  lifecycle {
    destroy = true
  }
}
`,
					ConfigStateChecks: []statecheck.StateCheck{
						ExpectFunc(func(ctx context.Context, _ statecheck.CheckStateRequest) error {
							labels, _, err := testAccConf.meta.v3client.Issues.ListLabels(ctx, testAccConf.meta.name, repo.GetName(), &github.ListOptions{PerPage: testAccConf.meta.maxPerPage})
							if err != nil {
								return err
							}

							if len(labels) != 0 {
								return fmt.Errorf("expected no labels, got %d", len(labels))
							}

							return nil
						}),
					},
				},
			},
		})
	})

	t.Run("can_delete_labels_from_archived_repositories_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		time.Sleep(2 * time.Second)

		config := fmt.Sprintf(`
resource "github_issue_labels" "test" {
  repository = "%s"

  label {
    name        = "archived-label-1"
    color       = "ff0000"
    description = "First test label"
  }

  label {
    name        = "archived-label-2"
    color       = "00ff00"
    description = "Second test label"
  }
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					PreConfig: func() {
						mustArchiveTestRepository(t, repo)
					},
					Config: config,
				},
				// This step should succeed - the labels should be removed from state
				// without trying to actually delete them from the archived repo
				{
					Config: `
removed {
  from = github_issue_labels.test

  lifecycle {
    destroy = true
  }
}
`,
				},
			},
		})
	})
}
