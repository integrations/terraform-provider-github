package github

import (
	"context"
	"fmt"
	"log"
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
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository.test", "default_branch",
					"main",
				),
			),
			// FIXME: Deferred until https://github.com/integrations/terraform-provider-github/issues/513
			// > Cannot update default branch for an empty repository. Please init the repository and push first
			// "after": resource.ComposeTestCheckFunc(
			// 	resource.TestCheckResourceAttr(
			// 		"github_repository.test", "default_branch",
			// 		"default",
			// 	),
			// ),
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
					// {
					// 	Config: strings.Replace(config,
					// 		`default_branch = "main"`,
					// 		`default_branch = "default"`, 1),
					// 	Check: checks["after"],
					// },
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
