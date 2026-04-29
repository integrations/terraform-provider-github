package github

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-github/v85/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubRepositoryFile(t *testing.T) {
	t.Run("creates and manages files", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-file-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name                 = "%s"
				auto_init            = true
				vulnerability_alerts = true
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"
			}
		`, repoName)
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository_file.test", "content", "bar"),
			resource.TestCheckResourceAttr("github_repository_file.test", "sha", "ba0e162e1c47469e3fe4b393a8bf8c569f302116"),
			resource.TestCheckResourceAttr("github_repository_file.test", "ref", "main"),
			resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_author"),
			resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_email"),
			resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_message"),
			resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_sha"),
			resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch"),
			resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_branch"),
			resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_sha"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
	t.Run("validates_commit_email_must_be_specified_if_commit_author_is_specified", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%srepo-file-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name                 = "%s"
				auto_init            = true
				vulnerability_alerts = true
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_author  = "Terraform User"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("all of `commit_author,commit_email` must be specified"),
				},
			},
		})
	})

	t.Run("validates_commit_author_must_be_specified_if_commit_email_is_specified", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%srepo-file-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name                 = "%s"
				auto_init            = true
				vulnerability_alerts = true
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_email   = "terraform@example.com"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("all of `commit_author,commit_email` must be specified"),
				},
			},
		})
	})

	t.Run("can be configured to overwrite files on create", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-file-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name                 = "%s"
			  auto_init            = true
	              vulnerability_alerts = true
			}

			resource "github_repository_file" "test" {
				repository          = github_repository.test.name
				branch              = "main"
				file                = "README.md"
				content             = "overwritten"
				overwrite_on_create = false
				commit_message      = "Managed by Terraform"
				commit_author       = "Terraform User"
				commit_email        = "terraform@example.com"
			}

		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_file.test", "content",
				"overwritten",
			),
			resource.TestCheckResourceAttr(
				"github_repository_file.test", "sha",
				"67c1a95c2d9bb138aefeaebb319cca82e531736b",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_file.test", "commit_author",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_file.test", "commit_email",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_file.test", "commit_message",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_file.test", "commit_sha",
			),
			resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch"),
			resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_branch"),
			resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_sha"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`refusing to overwrite existing file`),
				},
				{
					Config: strings.Replace(config,
						"overwrite_on_create = false",
						"overwrite_on_create = true", 1),
					Check: check,
				},
			},
		})
	})

	t.Run("creates and manages files on default branch if branch is omitted", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-file-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name                 = "%s"
				auto_init            = true
				vulnerability_alerts = true
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_file.test", "content",
				"bar",
			),
			resource.TestCheckResourceAttr(
				"github_repository_file.test", "sha",
				"ba0e162e1c47469e3fe4b393a8bf8c569f302116",
			),
			resource.TestCheckResourceAttr(
				"github_repository_file.test", "ref",
				"main",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_file.test", "commit_author",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_file.test", "commit_email",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_file.test", "commit_message",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_file.test", "commit_sha",
			),
			resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch"),
			resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_branch"),
			resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_sha"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("creates and manages files on auto created branch if branch does not exist", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-file-%s", testResourcePrefix, randomID)
		config := `
			resource "github_repository" "test" {
				name                 = "%s"
				auto_init            = true
				vulnerability_alerts = true
			}

			resource "github_repository_file" "test" {
				repository        = github_repository.test.name
				branch            = "does/not/exist"
				file              = "test"
				content           = "bar"
				commit_message    = "Managed by Terraform"
				commit_author     = "Terraform User"
				commit_email      = "terraform@example.com"
				autocreate_branch = %t
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, repoName, false),
					ExpectError: regexp.MustCompile(`unexpected status code: 404 Not Found`),
				},
				{
					Config: fmt.Sprintf(config, repoName, true),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_file.test", "content", "bar"),
						resource.TestCheckResourceAttr("github_repository_file.test", "sha", "ba0e162e1c47469e3fe4b393a8bf8c569f302116"),
						resource.TestCheckResourceAttr("github_repository_file.test", "ref", "does/not/exist"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_author"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_email"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_message"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_sha"),
						resource.TestCheckResourceAttr("github_repository_file.test", "autocreate_branch", "true"),
						resource.TestCheckResourceAttr("github_repository_file.test", "autocreate_branch_source_branch", "main"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "autocreate_branch_source_sha"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "repository_id"),
					),
				},
			},
		})
	})

	t.Run("can delete files from archived repositories without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-file-arch-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}

			resource "github_repository_file" "test" {
				repository = github_repository.test.name
				branch = "main"
				file = "archived-test.md"
				content = "# Test file for archived repo"
				commit_message = "Add test file"
				commit_author = "Terraform User"
				commit_email = "terraform@example.com"
			}
		`, repoName)

		archivedConfig := strings.Replace(config,
			`auto_init = true`,
			`auto_init = true
				archived = true`, 1)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_file.test", "file", "archived-test.md"),
					),
				},
				{
					Config: archivedConfig,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository.test", "archived", "true"),
					),
				},
				// This step should succeed - the file should be removed from state
				// without trying to actually delete it from the archived repo
				{
					Config: fmt.Sprintf(`
							resource "github_repository" "test" {
								name = "%s"
								auto_init = true
								archived = true
							}
						`, repoName),
				},
			},
		})
	})
	t.Run("imports_files_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sfile-import-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name                 = "%s"
				auto_init            = true
				vulnerability_alerts = true
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_file.test", "content", "bar"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "sha"),
						resource.TestCheckResourceAttr("github_repository_file.test", "ref", "main"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_author"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_email"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_message"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_sha"),
						resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch"),
						resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_branch"),
						resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_sha"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "repository_id"),
					),
				},
				{
					ResourceName:            "github_repository_file.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"commit_author", "commit_email"}, // For some reason `d` doesn't contain the commit author and email when importing.
				},
			},
		})
	})
	t.Run("imports_files_with_branch_in_id_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sfile-import-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name                 = "%s"
				auto_init            = true
				vulnerability_alerts = true
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_file.test", "content", "bar"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "sha"),
						resource.TestCheckResourceAttr("github_repository_file.test", "ref", "main"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_author"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_email"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_message"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "commit_sha"),
						resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch"),
						resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_branch"),
						resource.TestCheckNoResourceAttr("github_repository_file.test", "autocreate_branch_source_sha"),
						resource.TestCheckResourceAttrSet("github_repository_file.test", "repository_id"),
					),
				},
				{
					ResourceName:            "github_repository_file.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"commit_author", "commit_email"}, // For some reason `d` doesn't contain the commit author and email when importing.
				},
			},
		})
	})

	t.Run("produces a signed commit when deleting a managed file", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-file-signed-del-%s", testResourcePrefix, randomID)
		filename := "signed-delete-test.md"

		withFile := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_file" "test" {
				repository = github_repository.test.name
				branch     = "main"
				file       = "%s"
				content    = "signed delete test\n"
			}
		`, repoName, filename)

		withoutFile := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: withFile,
					Check: resource.TestCheckResourceAttr(
						"github_repository_file.test", "file", filename,
					),
				},
				{
					Config: withoutFile,
					Check:  testAccCheckLatestCommitIsSigned(repoName, "main", filename),
				},
			},
		})
	})
}

// testAccCheckLatestCommitIsSigned asserts that the HEAD of the given branch is
// a signed commit (commit.verification.verified == true) and that its message
// references the expected file path. It is used to guard against regressing the
// delete path off the GraphQL createCommitOnBranch mutation, since the REST
// Contents API produces unsigned commits on DELETE.
func testAccCheckLatestCommitIsSigned(repoName, branch, expectedPathInMessage string) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		meta, err := getTestMeta()
		if err != nil {
			return err
		}

		ctx := context.Background()
		commits, _, err := meta.v3client.Repositories.ListCommits(ctx, meta.name, repoName, &github.CommitsListOptions{
			SHA: branch,
			ListOptions: github.ListOptions{
				PerPage: 1,
			},
		})
		if err != nil {
			return fmt.Errorf("listing commits on %s/%s@%s: %w", meta.name, repoName, branch, err)
		}
		if len(commits) == 0 {
			return fmt.Errorf("no commits found on %s/%s@%s", meta.name, repoName, branch)
		}

		head := commits[0]
		msg := head.GetCommit().GetMessage()
		if !strings.Contains(msg, expectedPathInMessage) {
			return fmt.Errorf(
				"expected HEAD commit on %s/%s@%s to reference %q, got message %q (sha %s)",
				meta.name, repoName, branch, expectedPathInMessage, msg, head.GetSHA(),
			)
		}

		verification := head.GetCommit().GetVerification()
		if !verification.GetVerified() {
			return fmt.Errorf(
				"expected HEAD commit on %s/%s@%s (sha %s, message %q) to be signed, got verified=false reason=%q",
				meta.name, repoName, branch, head.GetSHA(), msg, verification.GetReason(),
			)
		}

		return nil
	}
}
