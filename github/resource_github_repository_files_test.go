package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryFiles(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("with_files", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		configInvalid := fmt.Sprintf(`
resource "github_repository_files" "test" {
  repository    = "%s"
  commit_author = "Terraform User"

  file {
    path    = "a.txt"
    content = "alpha"
  }
}
`, repo.GetName())

		config := fmt.Sprintf(`
resource "github_repository_files" "test" {
  repository    = "%s"
  commit_author = "Terraform User"
  commit_email  = "terraform@example.com"

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

data "github_repository_file" "readme" {
  repository = github_repository_files.test.repository
  branch     = github_repository_files.test.branch
  file       = "README.md"
}
`, repo.GetName())

		configEdited := fmt.Sprintf(`
resource "github_repository_files" "test" {
  repository    = "%s"
  commit_author = "Terraform User"
  commit_email  = "terraform@example.com"

  file {
    path    = "a.txt"
    content = "alpha"
  }
  file {
    path    = "nested/b.txt"
    content = "bravo edited"
  }
  file {
    path    = "nested/deeper/c.txt"
    content = "charlie"
  }
}
`, repo.GetName())

		configDuplicate := fmt.Sprintf(`
resource "github_repository_files" "test" {
  repository = "%s"

  file {
    path    = "a.txt"
    content = "alpha"
  }
  file {
    path    = "a.txt"
    content = "alpha again"
  }
}
`, repo.GetName())

		configUpdated := fmt.Sprintf(`
resource "github_repository_files" "test" {
  repository     = "%s"
  commit_message = "follow-up batch"
  commit_author  = "Terraform User"
  commit_email   = "terraform@example.com"

  file {
    path    = "a.txt"
    content = "alpha"
  }
  file {
    path    = "nested/b.txt"
    content = "bravo edited"
  }
  file {
    path    = "added.txt"
    content = "delta"
  }
}
`, repo.GetName())

		commitSHA := statecheck.CompareValue(compare.ValuesDiffer())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      configInvalid,
					PlanOnly:    true,
					ExpectError: regexp.MustCompile("all of `commit_author,commit_email` must be specified"),
				},
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("id"), knownvalue.StringRegexp(regexp.MustCompile(fmt.Sprintf(`^%s:%s:[0-9a-f]{40}$`, repo.GetName(), repo.GetDefaultBranch())))),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("branch"), knownvalue.StringExact(repo.GetDefaultBranch())),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("ref"), knownvalue.StringExact("refs/heads/"+repo.GetDefaultBranch())),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("commit_message"), knownvalue.StringExact("Terraform: 3 added")),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("tree_sha"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("file"), knownvalue.SetSizeExact(3)),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("file"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"path":    knownvalue.StringExact("nested/deeper/c.txt"),
								"content": knownvalue.StringExact("charlie"),
								"sha":     knownvalue.NotNull(),
							}),
						})),
						statecheck.ExpectKnownValue("data.github_repository_file.readme", tfjsonpath.New("content"), knownvalue.NotNull()),
						commitSHA.AddStateValue("github_repository_files.test", tfjsonpath.New("commit_sha")),
					},
				},
				{
					Config: configEdited,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_files.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("commit_message"), knownvalue.StringExact("Terraform: 1 updated")),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("file"), knownvalue.SetSizeExact(3)),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("file"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"path":    knownvalue.StringExact("nested/b.txt"),
								"content": knownvalue.StringExact("bravo edited"),
								"sha":     knownvalue.NotNull(),
							}),
						})),
						commitSHA.AddStateValue("github_repository_files.test", tfjsonpath.New("commit_sha")),
					},
				},
				{
					Config:      configDuplicate,
					PlanOnly:    true,
					ExpectError: regexp.MustCompile(`file path "a.txt" is declared more than once`),
				},
				{
					Config: configUpdated,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_files.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("commit_message"), knownvalue.StringExact("follow-up batch")),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("file"), knownvalue.SetSizeExact(3)),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("file"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"path":    knownvalue.StringExact("added.txt"),
								"content": knownvalue.StringExact("delta"),
								"sha":     knownvalue.NotNull(),
							}),
						})),
						commitSHA.AddStateValue("github_repository_files.test", tfjsonpath.New("commit_sha")),
					},
				},
				{
					ResourceName:            "github_repository_files.test",
					ImportState:             true,
					ImportStateId:           fmt.Sprintf("%s:%s", repo.GetName(), repo.GetDefaultBranch()),
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"file", "commit_message", "commit_author", "commit_email"},
				},
			},
		})
	})

	t.Run("with_branch", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		branch := mustCreateTestBranch(t, repo)

		config := fmt.Sprintf(`
resource "github_repository_files" "test" {
  repository = "%s"
  branch     = "%s"

  file {
    path    = "a.txt"
    content = "alpha"
  }
}
`, repo.GetName(), branch)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("branch"), knownvalue.StringExact(branch)),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("ref"), knownvalue.StringExact("refs/heads/"+branch)),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("commit_message"), knownvalue.StringExact("Terraform: 1 added")),
					},
				},
			},
		})
	})

	t.Run("adopts_matching_files_without_a_commit", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		existing := mustCreateRepositoryFile(t, repo, "a.txt", "alpha")

		config := fmt.Sprintf(`
resource "github_repository_files" "test" {
  repository = "%s"

  file {
    path    = "a.txt"
    content = "alpha"
  }
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("commit_sha"), knownvalue.StringExact(existing.GetSHA())),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("file"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"path":    knownvalue.StringExact("a.txt"),
								"content": knownvalue.StringExact("alpha"),
								"sha":     knownvalue.StringExact(existing.GetContent().GetSHA()),
							}),
						})),
					},
				},
			},
		})
	})

	t.Run("with_out_of_band_commit", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_files" "test" {
  repository = "%s"

  file {
    path    = "a.txt"
    content = "alpha"
  }
}
`, repo.GetName())

		commitSHA := statecheck.CompareValue(compare.ValuesSame())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						commitSHA.AddStateValue("github_repository_files.test", tfjsonpath.New("commit_sha")),
					},
				},
				{
					PreConfig:    func() { mustDeleteRepositoryFile(t, repo, "README.md") },
					RefreshState: true,
					RefreshPlanChecks: resource.RefreshPlanChecks{
						PostRefresh: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						commitSHA.AddStateValue("github_repository_files.test", tfjsonpath.New("commit_sha")),
					},
				},
			},
		})
	})

	t.Run("with_drift", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_files" "test" {
  repository = "%s"

  file {
    path    = "a.txt"
    content = "alpha"
  }
  file {
    path    = "b.txt"
    content = "bravo"
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
					PreConfig:          func() { mustDeleteRepositoryFile(t, repo, "a.txt") },
					RefreshState:       true,
					ExpectNonEmptyPlan: true,
					RefreshPlanChecks: resource.RefreshPlanChecks{
						PostRefresh: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_files.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_files.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("commit_message"), knownvalue.StringExact("Terraform: 1 added")),
						statecheck.ExpectKnownValue("github_repository_files.test", tfjsonpath.New("file"), knownvalue.SetSizeExact(2)),
					},
				},
			},
		})
	})
}

func TestGithubRepositoryFilesDiff(t *testing.T) {
	t.Parallel()

	file := func(path cty.Value, content string) cty.Value {
		return cty.ObjectVal(map[string]cty.Value{"path": path, "content": cty.StringVal(content), "sha": cty.NullVal(cty.String)})
	}
	unknown := cty.UnknownVal(cty.String)

	for name, tt := range map[string]struct {
		files   []cty.Value
		wantErr bool
	}{
		"distinct paths":  {files: []cty.Value{file(cty.StringVal("a.txt"), "x"), file(cty.StringVal("b.txt"), "x")}},
		"duplicate paths": {files: []cty.Value{file(cty.StringVal("a.txt"), "x"), file(cty.StringVal("a.txt"), "y")}, wantErr: true},
		"unknown paths":   {files: []cty.Value{file(unknown, "x"), file(unknown, "y")}},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			rawConfig := cty.ObjectVal(map[string]cty.Value{"repository": cty.StringVal("example"), "file": cty.SetVal(tt.files)})
			config := terraform.NewResourceConfigShimmed(rawConfig, resourceGithubRepositoryFiles().CoreConfigSchema())

			_, err := resourceGithubRepositoryFiles().Diff(t.Context(), &terraform.InstanceState{RawConfig: rawConfig}, config, nil)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error to be %v, got %v", tt.wantErr, err)
			}
		})
	}
}
