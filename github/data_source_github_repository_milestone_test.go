package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryMilestoneDataSource(t *testing.T) {
	t.Parallel()

	t.Run("queries a repository milestone", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		title := "v1.0.0"
		description := "General Availability"
		dueDate := "2020-11-22"
		state := "closed"

		config := fmt.Sprintf(`
resource "github_repository_milestone" "test" {
  owner       = "%s"
  repository  = "%s"
  title       = "%s"
  description = "%s"
  due_date    = "%s"
  state       = "%s"
}

data "github_repository_milestone" "test" {
	owner      = github_repository_milestone.test.owner
	repository = github_repository_milestone.test.repository
	number     = github_repository_milestone.test.number
}
`, repo.GetOwner().GetLogin(), repo.GetName(), title, description, dueDate, state)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository_milestone.test", tfjsonpath.New("title"), knownvalue.StringExact(title)),
						statecheck.ExpectKnownValue("data.github_repository_milestone.test", tfjsonpath.New("description"), knownvalue.StringExact(description)),
						statecheck.ExpectKnownValue("data.github_repository_milestone.test", tfjsonpath.New("due_date"), knownvalue.StringExact(dueDate)),
						statecheck.ExpectKnownValue("data.github_repository_milestone.test", tfjsonpath.New("state"), knownvalue.StringExact(state)),
					},
				},
			},
		})
	})
}
