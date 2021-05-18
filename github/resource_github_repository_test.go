package github

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositories(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates repositories without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {

			  name         = "tf-acc-test-create-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"

			  has_issues         = true
			  has_wiki           = true
			  has_downloads      = true
			  allow_merge_commit = true
			  allow_squash_merge = false
			  allow_rebase_merge = false
			  auto_init          = false

			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository.test", "has_issues",
				"true",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("updates a repositories name without error", func(t *testing.T) {

		oldName := fmt.Sprintf(`tf-acc-test-rename-%[1]s`, randomID)
		newName := fmt.Sprintf(`%[1]s-renamed`, oldName)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "%[1]s"
			  description  = "Terraform acceptance tests %[2]s"
			}
		`, oldName, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "name",
					oldName,
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "name",
					newName,
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						// Rename the repo to something else
						Config: strings.Replace(
							config,
							oldName,
							newName, 1),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("imports repositories without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "tf-acc-test-import-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"
				auto_init 	 = false
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_repository.test", "name"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:      "github_repository.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("archives repositories without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "tf-acc-test-archive-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"
				archived     = false
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "archived",
					"false",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "archived",
					"true",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: strings.Replace(config,
							`archived     = false`,
							`archived     = true`, 1),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("manages the project feature for a repository", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "tf-acc-test-project-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"
				has_projects = false
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "has_projects",
					"false",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "has_projects",
					"true",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: strings.Replace(config,
							`has_projects = false`,
							`has_projects = true`, 1),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("manages the default branch feature for a repository", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name           = "tf-acc-test-branch-%[1]s"
			  description    = "Terraform acceptance tests %[1]s"
			  default_branch = "main"
			  auto_init      = true
			}

			resource "github_branch" "default" {
			  repository = github_repository.test.name
			  branch     = "default"
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "default_branch",
					"main",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "default_branch",
					"default",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					// Test changing default_branch
					{
						Config: strings.Replace(config,
							`default_branch = "main"`,
							`default_branch = "default"`, 1),
						Check: checks["after"],
					},
					// Test changing default_branch back to main again
					{
						Config: config,
						Check:  checks["before"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("allows setting default_branch on an empty repository", func(t *testing.T) {

		// Although default_branch is deprecated, for backwards compatibility
		// we allow setting it to "main".

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name           = "tf-acc-test-empty-%[1]s"
			  description    = "Terraform acceptance tests %[1]s"
			  default_branch = "main"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository.test", "default_branch",
				"main",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					// Test creation with default_branch set
					{
						Config: config,
						Check:  check,
					},
					// Test that changing another property does not try to set
					// default_branch (which would crash).
					{
						Config: strings.Replace(config,
							`acceptance tests`,
							`acceptance test`, 1),
						Check: check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("manages the license and gitignore feature for a repository", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name           = "tf-acc-test-license-%[1]s"
				description    = "Terraform acceptance tests %[1]s"
				license_template   = "ms-pl"
				gitignore_template = "C++"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository.test", "license_template",
				"ms-pl",
			),
			resource.TestCheckResourceAttr(
				"github_repository.test", "gitignore_template",
				"C++",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("configures topics for a repository", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "tf-acc-test-topic-%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics			= ["terraform", "testing"]
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository.test", "topics.#",
				"2",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("creates a repository using a template", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "tf-acc-test-template-%s"
				description = "Terraform acceptance tests %[1]s"

				template {
					owner = "%s"
					repository = "%s"
				}

			}
		`, randomID, testOrganization, "terraform-template-module")

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository.test", "is_template",
				"false",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("archives repositories on destroy", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name               = "tf-acc-test-destroy-%[1]s"
				auto_init          = true
				archive_on_destroy = true
				archived           = false
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "archived",
					"false",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "archived",
					"true",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: strings.Replace(config,
							`archived           = false`,
							`archived           = true`, 1),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("configures vulnerability alerts", func(t *testing.T) {

		t.Run("for a public repository", func(t *testing.T) {

			config := fmt.Sprintf(`
				resource "github_repository" "test" {
					name       = "tf-acc-test-pub-vuln-%s"
					visibility = "public"
				}
			`, randomID)

			checks := map[string]resource.TestCheckFunc{
				"before": resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"github_repository.test", "vulnerability_alerts",
						"false",
					),
				),
				"after": resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"github_repository.test", "vulnerability_alerts",
						"true",
					),
					resource.TestCheckResourceAttr(
						"github_repository.test", "visibility",
						"public",
					),
				),
			}

			testCase := func(t *testing.T, mode string) {
				resource.Test(t, resource.TestCase{
					PreCheck:  func() { skipUnlessMode(t, mode) },
					Providers: testAccProviders,
					Steps: []resource.TestStep{
						{
							Config: config,
							Check:  checks["before"],
						},
						{
							Config: strings.Replace(config,
								`}`,
								"vulnerability_alerts = true\n}", 1),
							Check: checks["after"],
						},
					},
				})
			}

			t.Run("with an anonymous account", func(t *testing.T) {
				t.Skip("anonymous account not supported for this operation")
			})

			t.Run("with an individual account", func(t *testing.T) {
				testCase(t, individual)
			})

			t.Run("with an organization account", func(t *testing.T) {
				testCase(t, organization)
			})
		})

		t.Run("for a private repository", func(t *testing.T) {

			config := fmt.Sprintf(`
				resource "github_repository" "test" {
					name       = "tf-acc-test-prv-vuln-%s"
					visibility = "private"
				}
			`, randomID)

			checks := map[string]resource.TestCheckFunc{
				"before": resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"github_repository.test", "vulnerability_alerts",
						"false",
					),
				),
				"after": resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"github_repository.test", "vulnerability_alerts",
						"true",
					),
					resource.TestCheckResourceAttr(
						"github_repository.test", "visibility",
						"private",
					),
				),
			}

			testCase := func(t *testing.T, mode string) {
				resource.Test(t, resource.TestCase{
					PreCheck:  func() { skipUnlessMode(t, mode) },
					Providers: testAccProviders,
					Steps: []resource.TestStep{
						{
							Config: config,
							Check:  checks["before"],
						},
						{
							Config: strings.Replace(config,
								`}`,
								"vulnerability_alerts = true\n}", 1),
							Check: checks["after"],
						},
					},
				})
			}

			t.Run("with an anonymous account", func(t *testing.T) {
				t.Skip("anonymous account not supported for this operation")
			})

			t.Run("with an individual account", func(t *testing.T) {
				testCase(t, individual)
			})

			t.Run("with an organization account", func(t *testing.T) {
				testCase(t, organization)
			})
		})

	})
}
func TestAccGithubRepositoryPages(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("manages the pages feature for a repository", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-%s"
				auto_init    = true
				pages {
					source {
						branch = "main"
					}
				}
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository.test", "pages.0.source.0.branch",
				"main",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

}

func TestAccGithubRepositoryVisibility(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	for _, visibility := range []string{"private", "internal"} {
		t.Run(fmt.Sprintf("creates repos with %s visibility", visibility), func(t *testing.T) {

			config := fmt.Sprintf(`
				resource "github_repository" "%[1]s" {
					name       = "tf-acc-test-visibility-%[1]s-%[2]s"
					visibility = "%[1]s"
				}
			`, visibility, randomID)

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					fmt.Sprintf("github_repository.%s", visibility), "visibility",
					visibility,
				),
			)

			testCase := func(t *testing.T, mode string) {
				resource.Test(t, resource.TestCase{
					PreCheck:  func() { skipUnlessMode(t, mode) },
					Providers: testAccProviders,
					Steps: []resource.TestStep{
						{
							Config: config,
							Check:  check,
						},
					},
				})
			}

			t.Run("with an anonymous account", func(t *testing.T) {
				t.Skip("anonymous account not supported for this operation")
			})

			t.Run("with an individual account", func(t *testing.T) {
				testCase(t, individual)
			})

			t.Run("with an organization account", func(t *testing.T) {
				testCase(t, organization)
			})
		})
	}

	t.Run("updates repos to private visibility", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "public" {
				name       = "tf-acc-test-visibility-public-%s"
				visibility = "public"
				vulnerability_alerts = false
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.public", "visibility",
					"public",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.public", "visibility",
					"private",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: reconfigureVisibility(config, "private"),
						Check:  checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("updates repos to public visibility", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "tf-acc-test-prv-vuln-%s"
				visibility = "private"
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "vulnerability_alerts",
					"false",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "vulnerability_alerts",
					"true",
				),
				resource.TestCheckResourceAttr(
					"github_repository.test", "visibility",
					"private",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: strings.Replace(config,
							`}`,
							"vulnerability_alerts = true\n}", 1),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("updates repos to internal visibility", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "tf-acc-test-prv-vuln-%s"
				visibility = "private"
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "vulnerability_alerts",
					"false",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "vulnerability_alerts",
					"true",
				),
				resource.TestCheckResourceAttr(
					"github_repository.test", "visibility",
					"private",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: strings.Replace(config,
							`}`,
							"vulnerability_alerts = true\n}", 1),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("sets private visibility for repositories created by a template", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "private" {
				name       = "tf-acc-test-visibility-private-%s"
				visibility = "private"
				template {
					owner      = "%s"
					repository = "%s"
				}
			}
		`, randomID, testOrganization, "terraform-template-module")

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository.private", "visibility",
				"private",
			),
			resource.TestCheckResourceAttr(
				"github_repository.private", "private",
				"true",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

}

func testSweepRepositories(region string) error {
	meta, err := sharedConfigForRegion(region)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client

	repos, _, err := client.Repositories.List(context.TODO(), meta.(*Owner).name, nil)
	if err != nil {
		return err
	}

	for _, r := range repos {
		if name := r.GetName(); strings.HasPrefix(name, "tf-acc-") || strings.HasPrefix(name, "foo-") {
			log.Printf("Destroying Repository %s", name)

			if _, err := client.Repositories.Delete(context.TODO(), meta.(*Owner).name, name); err != nil {
				return err
			}
		}
	}

	return nil
}

func init() {
	resource.AddTestSweepers("github_repository", &resource.Sweeper{
		Name: "github_repository",
		F:    testSweepRepositories,
	})
}

func reconfigureVisibility(config, visibility string) string {
	re := regexp.MustCompile(`visibility = "(.*)"`)
	newConfig := re.ReplaceAllString(
		config,
		fmt.Sprintf(`visibility = "%s"`, visibility),
	)
	return newConfig
}
