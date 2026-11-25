package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryFileDataSource(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("reads_a_file_with_a_branch_name_provided_without_erroring", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		file := mustCreateTestRepositoryFile(t, repo, "test", "bar")

		config := fmt.Sprintf(`
data "github_repository_file" "test" {
  repository = "%s"
  branch     = "%s"
  file       = "test"
}
`, repo.GetName(), repo.GetDefaultBranch())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("id"), knownvalue.StringExact(fmt.Sprintf("%s/test", repo.GetName()))),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("ref"), knownvalue.StringExact(repo.GetDefaultBranch())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("content"), knownvalue.StringExact("bar")),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("sha"), knownvalue.StringExact(file.Content.GetSHA())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_sha"), knownvalue.StringExact(file.GetSHA())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_message"), knownvalue.StringExact(file.GetMessage())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_author"), knownvalue.StringExact(file.Committer.GetName())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_email"), knownvalue.StringExact(file.Committer.GetEmail())),
					},
				},
			},
		})
	})

	t.Run("reads_a_file_from_a_short_repo_name_without_erroring", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		file := mustCreateTestRepositoryFile(t, repo, "test", "bar")

		config := fmt.Sprintf(`
data "github_repository_file" "test" {
  repository = "%s"
  branch     = "%s"
  file       = "test"
}
`, repo.GetName(), repo.GetDefaultBranch())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("id"), knownvalue.StringExact(fmt.Sprintf("%s/test", repo.GetName()))),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("ref"), knownvalue.StringExact(repo.GetDefaultBranch())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("content"), knownvalue.StringExact("bar")),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("sha"), knownvalue.StringExact(file.Content.GetSHA())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_sha"), knownvalue.StringExact(file.GetSHA())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_message"), knownvalue.StringExact(file.GetMessage())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_author"), knownvalue.StringExact(file.Committer.GetName())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_email"), knownvalue.StringExact(file.Committer.GetEmail())),
					},
				},
			},
		})
	})

	t.Run("reads_a_file_without_providing_a_branch_name", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		file := mustCreateTestRepositoryFile(t, repo, "test", "bar")

		config := fmt.Sprintf(`
data "github_repository_file" "test" {
  repository = "%s"
  file       = "test"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("id"), knownvalue.StringExact(fmt.Sprintf("%s/test", repo.GetName()))),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("branch"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("ref"), knownvalue.StringExact(repo.GetDefaultBranch())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("content"), knownvalue.StringExact("bar")),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("sha"), knownvalue.StringExact(file.Content.GetSHA())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_sha"), knownvalue.StringExact(file.GetSHA())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_message"), knownvalue.StringExact(file.GetMessage())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_author"), knownvalue.StringExact(file.Committer.GetName())),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_email"), knownvalue.StringExact(file.Committer.GetEmail())),
					},
				},
			},
		})
	})

	t.Run("reads_a_non_existent_file_without_an_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_repository_file" "test" {
  repository = "%s"
  file       = "test"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
			},
		})
	})

	t.Run("reads_a_directory_without_erroring", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_repository_file" "test" {
  repository = "%s"
  file       = "."
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("id"), knownvalue.StringExact(fmt.Sprintf("%s/.", repo.GetName()))),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("ref"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("content"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("sha"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_sha"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_message"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_author"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_repository_file.test", tfjsonpath.New("commit_email"), knownvalue.Null()),
					},
				},
			},
		})
	})
}
