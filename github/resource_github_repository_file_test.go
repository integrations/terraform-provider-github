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

func TestAccGithubRepositoryFile(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("creates_updates_and_imports_files", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_file" "test" {
  repository     = "%s"
  branch         = "%s"
  file           = "test"
  content        = "%%s"
  commit_message = "Managed by Terraform"
  commit_author  = "Terraform User"
  commit_email   = "terraform@example.com"
}
`, repo.GetName(), repo.GetDefaultBranch())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "bar"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("id"), knownvalue.StringExact(fmt.Sprintf("%s:test:%s", repo.GetName(), repo.GetDefaultBranch()))),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("ref"), knownvalue.StringExact(repo.GetDefaultBranch())),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("sha"), knownvalue.StringExact("ba0e162e1c47469e3fe4b393a8bf8c569f302116")),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("commit_sha"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("autocreate_branch"), knownvalue.Null()),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("autocreate_branch_source_branch"), knownvalue.Null()),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("autocreate_branch_source_sha"), knownvalue.Null()),
					},
				},
				{
					Config: fmt.Sprintf(config, "baz"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_file.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("sha"), knownvalue.StringExact("3f9538666251333f5fa519e01eb267d371ca9c78")),
					},
				},
				{
					ResourceName:            "github_repository_file.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"commit_author", "commit_email"}, // The author and email are only read into state when they are configured.
				},
			},
		})
	})

	t.Run("creates_and_imports_files_on_the_default_branch_if_branch_is_omitted", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_file" "test" {
  repository = "%s"
  file       = "test"
  content    = "bar"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("id"), knownvalue.StringExact(fmt.Sprintf("%s:test:%s", repo.GetName(), repo.GetDefaultBranch()))),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("branch"), knownvalue.StringExact(repo.GetDefaultBranch())),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("ref"), knownvalue.StringExact(repo.GetDefaultBranch())),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("sha"), knownvalue.StringExact("ba0e162e1c47469e3fe4b393a8bf8c569f302116")),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("commit_sha"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:      "github_repository_file.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
				{
					// An ID without a branch resolves to the default branch.
					ResourceName:      "github_repository_file.test",
					ImportState:       true,
					ImportStateId:     fmt.Sprintf("%s:test:", repo.GetName()),
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("requires_commit_author_and_commit_email_together", func(t *testing.T) {
		t.Parallel()

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
resource "github_repository_file" "test" {
  repository     = "test"
  file           = "test"
  content        = "bar"
  commit_message = "Managed by Terraform"
  commit_author  = "Terraform User"
}
`,
					ExpectError: regexp.MustCompile("all of `commit_author,commit_email` must be specified"),
				},
				{
					Config: `
resource "github_repository_file" "test" {
  repository     = "test"
  file           = "test"
  content        = "bar"
  commit_message = "Managed by Terraform"
  commit_email   = "terraform@example.com"
}
`,
					ExpectError: regexp.MustCompile("all of `commit_author,commit_email` must be specified"),
				},
			},
		})
	})

	t.Run("can_be_configured_to_overwrite_files_on_create", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		mustCreateTestRepositoryFile(t, repo, "test", "foo")

		config := fmt.Sprintf(`
resource "github_repository_file" "test" {
  repository          = "%s"
  file                = "test"
  content             = "overwritten"
  overwrite_on_create = %%t
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, false),
					ExpectError: regexp.MustCompile(`refusing to overwrite existing file`),
				},
				{
					Config: fmt.Sprintf(config, true),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("sha"), knownvalue.StringExact("67c1a95c2d9bb138aefeaebb319cca82e531736b")),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("commit_sha"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("creates_files_on_an_auto_created_branch_if_the_branch_does_not_exist", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_file" "test" {
  repository        = "%s"
  branch            = "does/not/exist"
  file              = "test"
  content           = "bar"
  autocreate_branch = %%t
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, false),
					ExpectError: regexp.MustCompile(`unexpected status code: 404 Not Found`),
				},
				{
					Config: fmt.Sprintf(config, true),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("ref"), knownvalue.StringExact("does/not/exist")),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("sha"), knownvalue.StringExact("ba0e162e1c47469e3fe4b393a8bf8c569f302116")),
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("autocreate_branch_source_sha"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("can_delete_files_from_archived_repositories_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_file" "test" {
  repository = "%s"
  file       = "archived-test.md"
  content    = "# Test file for archived repo"
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
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					// The file is removed from state without trying to delete it from the archived repository.
					Config: "# no resources",
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_file.test", plancheck.ResourceActionDestroy),
						},
					},
				},
			},
		})
	})

	t.Run("supports_colons_in_the_file_path", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		filePathWithColon := "repro/example:one.yaml"

		config := fmt.Sprintf(`
resource "github_repository_file" "test" {
  repository = "%s"
  file       = "%s"
  content    = "bar"
}
`, repo.GetName(), filePathWithColon)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_file.test", tfjsonpath.New("id"), knownvalue.StringExact(fmt.Sprintf("%s:%s:%s", repo.GetName(), escapeIDPart(filePathWithColon), repo.GetDefaultBranch()))),
					},
				},
				{
					ResourceName:      "github_repository_file.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
