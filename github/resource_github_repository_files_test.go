package github

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubRepositoryFiles(t *testing.T) {
	t.Run("creates multiple files in a single commit", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-files-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_files" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				commit_message = "Initial batch from terraform"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"

				file {
					path    = "a.txt"
					content = "alpha"
				}
				file {
					path    = "nested/b.txt"
					content = "bravo"
				}
				file {
					path    = "nested/deeper/c.txt"
					content = "charlie"
				}
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_files.test", "file.#", "3"),
						resource.TestCheckResourceAttrSet("github_repository_files.test", "commit_sha"),
						resource.TestCheckResourceAttrSet("github_repository_files.test", "tree_sha"),
						resource.TestCheckResourceAttrSet("github_repository_files.test", "repository_id"),
						resource.TestCheckResourceAttr("github_repository_files.test", "ref", "refs/heads/main"),
						resource.TestCheckResourceAttr("github_repository_files.test", "commit_message", "Initial batch from terraform"),
						checkCommitHasFileCount("github_repository_files.test", repoName, 3),
						checkBlobContent("github_repository_files.test", repoName, "main", "a.txt", "alpha"),
						checkBlobContent("github_repository_files.test", repoName, "main", "nested/deeper/c.txt", "charlie"),
						checkEveryFileHasSHA("github_repository_files.test"),
					),
				},
			},
		})
	})

	t.Run("applies add, modify, and delete in a single commit", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-files-%s", testResourcePrefix, randomID)

		base := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_files" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				commit_message = "initial"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"

				file {
					path    = "keep.txt"
					content = "keep"
				}
				file {
					path    = "modify.txt"
					content = "original"
				}
				file {
					path    = "remove.txt"
					content = "to be removed"
				}
			}
		`, repoName)

		updated := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_files" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				commit_message = "follow-up batch"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"

				file {
					path    = "keep.txt"
					content = "keep"
				}
				file {
					path    = "modify.txt"
					content = "edited"
				}
				file {
					path    = "added.txt"
					content = "newly added"
				}
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: base,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_files.test", "file.#", "3"),
					),
				},
				{
					Config: updated,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_files.test", "file.#", "3"),
						resource.TestCheckResourceAttr("github_repository_files.test", "commit_message", "follow-up batch"),
						checkCommitHasFileCount("github_repository_files.test", repoName, 3),
						checkBlobContent("github_repository_files.test", repoName, "main", "modify.txt", "edited"),
						checkBlobContent("github_repository_files.test", repoName, "main", "added.txt", "newly added"),
						checkPathAbsent(repoName, "main", "remove.txt"),
					),
				},
			},
		})
	})

	t.Run("leaves untracked files untouched", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-files-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_files" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"

				file {
					path    = "managed.txt"
					content = "managed"
				}
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_files.test", "file.#", "1"),
						// README.md is auto-created by auto_init=true and is not managed by this resource;
						// it must survive the commit.
						checkBlobContent("github_repository_files.test", repoName, "main", "README.md", ""),
					),
				},
			},
		})
	})

	t.Run("validates commit_email is required when commit_author is set", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-files-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_files" "test" {
				repository    = github_repository.test.name
				branch        = "main"
				commit_author = "Terraform User"

				file {
					path    = "a.txt"
					content = "alpha"
				}
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

	t.Run("defaults to repository default branch when branch is omitted", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-files-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_files" "test" {
				repository    = github_repository.test.name
				commit_author = "Terraform User"
				commit_email  = "terraform@example.com"

				file {
					path    = "a.txt"
					content = "alpha"
				}
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_files.test", "branch", "main"),
					),
				},
			},
		})
	})
}

func checkCommitHasFileCount(resourceName, repo string, want int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}
		commitSHA := rs.Primary.Attributes["commit_sha"]
		if commitSHA == "" {
			return fmt.Errorf("commit_sha is empty in state for %s", resourceName)
		}
		meta, err := getTestMeta()
		if err != nil {
			return err
		}
		ctx := context.Background()
		commit, _, err := meta.v3client.Repositories.GetCommit(ctx, meta.name, repo, commitSHA, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch commit %s: %w", commitSHA, err)
		}
		if got := len(commit.Files); got != want {
			return fmt.Errorf("commit %s touched %d files; want %d", commitSHA, got, want)
		}
		return nil
	}
}

func checkBlobContent(resourceName, repo, branch, path, wantContent string) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		meta, err := getTestMeta()
		if err != nil {
			return err
		}
		ctx := context.Background()
		fc, _, _, err := meta.v3client.Repositories.GetContents(ctx, meta.name, repo, path, &github.RepositoryContentGetOptions{Ref: branch})
		if err != nil {
			return fmt.Errorf("failed to fetch %s on %s: %w", path, branch, err)
		}
		if fc == nil {
			return fmt.Errorf("file %s not found on branch %s", path, branch)
		}
		if wantContent == "" {
			return nil
		}
		got, err := fc.GetContent()
		if err != nil {
			return fmt.Errorf("failed to decode content for %s: %w", path, err)
		}
		if got != wantContent {
			return fmt.Errorf("content of %s = %q; want %q", path, got, wantContent)
		}
		return nil
	}
}

func checkEveryFileHasSHA(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}
		count := 0
		for k := range rs.Primary.Attributes {
			// file.<hash>.sha attributes are produced by SDKv2 set serialization
			if !regexpFileSha.MatchString(k) {
				continue
			}
			count++
			if v := rs.Primary.Attributes[k]; v == "" {
				return fmt.Errorf("expected %s to have a non-empty SHA in state", k)
			}
		}
		if count == 0 {
			return fmt.Errorf("expected at least one file.*.sha attribute in state")
		}
		return nil
	}
}

var regexpFileSha = regexp.MustCompile(`^file\.\d+\.sha$`)

func checkPathAbsent(repo, branch, path string) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		meta, err := getTestMeta()
		if err != nil {
			return err
		}
		ctx := context.Background()
		fc, _, _, err := meta.v3client.Repositories.GetContents(ctx, meta.name, repo, path, &github.RepositoryContentGetOptions{Ref: branch})
		if err == nil && fc != nil {
			return fmt.Errorf("expected %s to be absent from %s; found content", path, branch)
		}
		return nil
	}
}
