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

func TestAccDataSourceGithubRepository(t *testing.T) {
	t.Parallel()

	t.Run("queries_a_public_repository_without_error", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
data "github_repository" "test" {
  full_name = "%s/%s"
}
`, testAccConf.testPublicRepositoryOwner, testAccConf.testPublicRepository)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("repo_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("sets_all_computed_attributes", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
data "github_repository" "test" {
	full_name = "%s/%s"
}
`, testAccConf.testPublicRepositoryOwner, testAccConf.testPublicRepository)
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("full_name"), knownvalue.StringExact(fmt.Sprintf("%s/%s", testAccConf.testPublicRepositoryOwner, testAccConf.testPublicRepository))),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("name"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("repo_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("node_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("private"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("visibility"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("has_issues"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("has_discussions"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("has_projects"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("has_downloads"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("has_wiki"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("is_template"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("fork"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("allow_merge_commit"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("allow_squash_merge"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("allow_rebase_merge"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("allow_auto_merge"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("allow_update_branch"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("allow_forking"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("squash_merge_commit_title"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("squash_merge_commit_message"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("merge_commit_title"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("merge_commit_message"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("default_branch"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("primary_language"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("archived"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("delete_branch_on_merge"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("web_commit_signoff_required"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("html_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("ssh_clone_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("svn_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("git_clone_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("http_clone_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("topics"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("template"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("repository_license"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("pages"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("queries_a_repository_belonging_to_owner_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_repository" "test" {
  name = "%s"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("repo_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("full_name"), knownvalue.StringExact(fmt.Sprintf("%s/%s", testAccConf.owner, repo.GetName()))),
					},
				},
			},
		})
	})

	t.Run("queries_a_repository_with_pages_configured", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ds-pages-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "%s"
				auto_init    = true
				pages {
					source {
						branch = "main"
					}
				}
			}

			data "github_repository" "test" {
				name = github_repository.test.name
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("pages").AtSliceIndex(0).AtMapKey("source").AtSliceIndex(0).AtMapKey("branch"), knownvalue.StringExact("main")),
					},
				},
			},
		})
	})

	t.Run("checks_defaults_on_a_new_repository", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ds-defaults-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
  name         = "%s"
  auto_init    = true
}

data "github_repository" "test" {
  name = github_repository.test.name
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("description"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("homepage_url"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("pages"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("repository_license"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("topics"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("template"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("queries_a_public_repository_that_is_a_template", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			data "github_repository" "test" {
				full_name = "%s/%s"
			}
		`, testAccConf.testPublicTemplateRepositoryOwner, testAccConf.testPublicTemplateRepository)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("is_template"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("queries_an_org_repository_that_is_a_template", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ds-defaults-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
  name         = "%s"
  auto_init    = true
  is_template = true
}

data "github_repository" "test" {
  name = github_repository.test.name
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("is_template"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("queries_a_repository_that_was_generated_from_a_template", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ds-template-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
				template {
					owner      = "%s"
					repository = "%s"
				}
			}

			data "github_repository" "test" {
				name = github_repository.test.name
			}
		`, repoName, testAccConf.testPublicTemplateRepositoryOwner, testAccConf.testPublicTemplateRepository)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("template").AtSliceIndex(0).AtMapKey("owner"), knownvalue.StringExact(testAccConf.testPublicTemplateRepositoryOwner)),
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("template").AtSliceIndex(0).AtMapKey("repository"), knownvalue.StringExact(testAccConf.testPublicTemplateRepository)),
					},
				},
			},
		})
	})

	t.Run("queries_a_repository_that_has_no_primary_language", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ds-nolang-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
			}

			data "github_repository" "test" {
				name = github_repository.test.name
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("primary_language"), knownvalue.StringExact("")),
					},
				},
			},
		})
	})

	t.Run("queries_a_repository_that_has_go_as_primary_language", func(t *testing.T) {
		t.Parallel()

		config := `
			data "github_repository" "test" {
				full_name = "integrations/terraform-provider-github"
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("primary_language"), knownvalue.StringExact("Go")),
					},
				},
			},
		})
	})

	t.Run("queries_a_repository_that_has_a_license", func(t *testing.T) {
		t.Parallel()

		config := `
data "github_repository" "test" {
  full_name = "integrations/terraform-provider-github"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository.test", tfjsonpath.New("repository_license").AtSliceIndex(0).AtMapKey("license").AtSliceIndex(0).AtMapKey("key"), knownvalue.StringExact("mit")),
					},
				},
			},
		})
	})
}
