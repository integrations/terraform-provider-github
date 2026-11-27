package github

import (
	"fmt"
	"testing"

	"github.com/google/go-github/v93/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryPullRequestDataSource(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("manages_the_pull_request_lifecycle", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		branch := mustCreateTestRepositoryBranch(t, repo)
		mustCreateTestRepositoryFile(t, repo, "test", "bar", func(opts *github.RepositoryContentFileOptions) {
			opts.Branch = &branch
		})
		pr := mustCreateTestRepositoryPullRequest(t, repo, branch)

		config := fmt.Sprintf(`
data "github_repository_pull_request" "test" {
  base_repository = "%s"
  number          = %d
}
`, repo.GetName(), pr.GetNumber())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("base_ref"), knownvalue.StringExact(repo.GetDefaultBranch())),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("base_sha"), knownvalue.StringExact(pr.GetBase().GetSHA())),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("head_ref"), knownvalue.StringExact(branch)),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("head_sha"), knownvalue.StringExact(pr.GetHead().GetSHA())),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("head_repository"), knownvalue.StringExact(repo.GetName())),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("head_owner"), knownvalue.StringExact(testAccConf.meta.name)),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("title"), knownvalue.StringExact(pr.GetTitle())),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("body"), knownvalue.StringExact(pr.GetBody())),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("state"), knownvalue.StringExact("open")),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("draft"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("maintainer_can_modify"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("labels"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("opened_by"), knownvalue.StringExact(pr.GetUser().GetLogin())),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("opened_at"), knownvalue.Int64Exact(pr.GetCreatedAt().Unix())),
						statecheck.ExpectKnownValue("data.github_repository_pull_request.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
					},
				},
			},
		})
	})
}
