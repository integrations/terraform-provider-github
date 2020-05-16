package github

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	resource.AddTestSweepers("github_repository", &resource.Sweeper{
		Name: "github_repository",
		F:    testSweepRepositories,
	})

}

func testSweepRepositories(region string) error {
	meta, err := sharedConfigForRegion(region)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client

	repos, _, err := client.Repositories.List(context.TODO(), meta.(*Organization).name, nil)
	if err != nil {
		return err
	}

	for _, r := range repos {
		if name := r.GetName(); strings.HasPrefix(name, "tf-acc-") || strings.HasPrefix(name, "foo-") {
			log.Printf("Destroying Repository %s", name)

			if _, err := client.Repositories.Delete(context.TODO(), meta.(*Organization).name, name); err != nil {
				return err
			}
		}
	}

	return nil
}

func TestAccGithubRepository_basic(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:                name,
						Description:         description,
						Homepage:            "http://example.com/",
						HasIssues:           true,
						HasWiki:             true,
						IsTemplate:          false,
						AllowMergeCommit:    true,
						AllowSquashMerge:    false,
						AllowRebaseMerge:    false,
						DeleteBranchOnMerge: false,
						HasDownloads:        true,
						HasProjects:         false,
						DefaultBranch:       "master",
						Archived:            false,
					}),
				),
			},
			{
				Config: testAccGithubRepositoryUpdateConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:             name,
						Description:      "Updated " + description,
						Homepage:         "http://example.com/",
						AllowMergeCommit: false,
						AllowSquashMerge: true,
						AllowRebaseMerge: true,
						IsTemplate:       true,
						DefaultBranch:    "master",
						HasProjects:      false,
						Archived:         false,
					}),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init",
				},
			},
		},
	})
}

func TestAccGithubRepository_archive(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryArchivedConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:             name,
						Description:      description,
						Homepage:         "http://example.com/",
						HasIssues:        true,
						HasWiki:          true,
						AllowMergeCommit: true,
						AllowSquashMerge: false,
						AllowRebaseMerge: false,
						HasDownloads:     true,
						DefaultBranch:    "master",
						Archived:         true,
					}),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init",
				},
			},
		},
	})
}

func TestAccGithubRepository_archiveUpdate(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:             name,
						Description:      description,
						Homepage:         "http://example.com/",
						HasIssues:        true,
						HasWiki:          true,
						AllowMergeCommit: true,
						AllowSquashMerge: false,
						AllowRebaseMerge: false,
						HasDownloads:     true,
						DefaultBranch:    "master",
						Archived:         false,
					}),
				),
			},
			{
				Config: testAccGithubRepositoryArchivedConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:             name,
						Description:      description,
						Homepage:         "http://example.com/",
						HasIssues:        true,
						HasWiki:          true,
						AllowMergeCommit: true,
						AllowSquashMerge: false,
						AllowRebaseMerge: false,
						HasDownloads:     true,
						DefaultBranch:    "master",
						Archived:         true,
					}),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubRepository_hasProjects(t *testing.T) {
	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfigHasProjects(randString),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init",
				},
			},
		},
	})
}

func TestAccGithubRepository_defaultBranch(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfigDefaultBranch(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:             name,
						Description:      description,
						Homepage:         "http://example.com/",
						HasIssues:        true,
						HasWiki:          true,
						AllowMergeCommit: true,
						AutoInit:         true,
						AllowSquashMerge: false,
						AllowRebaseMerge: false,
						HasDownloads:     true,
						DefaultBranch:    "master",
						Archived:         false,
					}),
				),
			},
			{
				PreConfig: func() {
					if err := testAccCreateRepositoryBranch("foo", *repo.Name); err != nil {
						panic(err.Error())
					}
				},
				Config: testAccGithubRepositoryUpdateConfigDefaultBranch(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:             name,
						Description:      "Updated " + description,
						Homepage:         "http://example.com/",
						AutoInit:         true,
						HasIssues:        true,
						HasWiki:          true,
						AllowMergeCommit: true,
						AllowSquashMerge: false,
						AllowRebaseMerge: false,
						HasDownloads:     true,
						DefaultBranch:    "foo",
						Archived:         false,
					}),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init",
				},
			},
		},
	})
}

func TestAccGithubRepository_templates(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfigTemplates(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:              name,
						Description:       description,
						Homepage:          "http://example.com/",
						HasIssues:         true,
						HasWiki:           true,
						AllowMergeCommit:  true,
						AutoInit:          true,
						AllowSquashMerge:  false,
						AllowRebaseMerge:  false,
						HasDownloads:      true,
						DefaultBranch:     "master",
						LicenseTemplate:   "ms-pl",
						GitignoreTemplate: "C++",
						Archived:          false,
					}),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init", "gitignore_template", "license_template",
				},
			},
		},
	})
}

func TestAccGithubRepository_topics(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccGithubRepositoryConfigTopics(randString, `"TOPIC"`),
				ExpectError: regexp.MustCompile(`must include only lowercase alphanumeric characters or hyphens and cannot start with a hyphen`),
			},
			{
				Config:      testAccGithubRepositoryConfigTopics(randString, `"-topic"`),
				ExpectError: regexp.MustCompile(`must include only lowercase alphanumeric characters or hyphens and cannot start with a hyphen`),
			},
			{
				Config:      testAccGithubRepositoryConfigTopics(randString, `"töpic"`),
				ExpectError: regexp.MustCompile(`must include only lowercase alphanumeric characters or hyphens and cannot start with a hyphen`),
			},
			{
				Config: testAccGithubRepositoryConfigTopics(randString, `"topic1", "topic2"`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:        name,
						Description: description,
						Homepage:    "http://example.com/",
						Topics:      []string{"topic2", "topic1"},

						// non-zero defaults
						DefaultBranch:    "master",
						AllowMergeCommit: true,
						AllowSquashMerge: true,
						AllowRebaseMerge: true,
					}),
				),
			},
			{
				Config: testAccGithubRepositoryConfigTopics(randString, `"topic1", "topic2", "topic3"`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:        name,
						Description: description,
						Homepage:    "http://example.com/",
						Topics:      []string{"topic1", "topic2", "topic3"},

						// non-zero defaults
						DefaultBranch:    "master",
						AllowMergeCommit: true,
						AllowSquashMerge: true,
						AllowRebaseMerge: true,
					}),
				),
			},
			{
				Config: testAccGithubRepositoryConfigTopics(randString, ``),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:        name,
						Description: description,
						Homepage:    "http://example.com/",
						Topics:      []string{},

						// non-zero defaults
						DefaultBranch:    "master",
						AllowMergeCommit: true,
						AllowSquashMerge: true,
						AllowRebaseMerge: true,
					}),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init",
				},
			},
		},
	})
}

func TestAccGithubRepository_autoInitForceNew(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfigAutoInitForceNew(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "auto_init", "false"),
				),
			},
			{
				Config: testAccGithubRepositoryConfigAutoInitForceNewUpdate(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "auto_init", "true"),
					resource.TestCheckResourceAttr(rn, "license_template", "mpl-2.0"),
					resource.TestCheckResourceAttr(rn, "gitignore_template", "Go"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init", "license_template", "gitignore_template",
				},
			},
		},
	})
}

func TestAccGithubRepository_createFromTemplate(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryCreateFromTemplate(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryTemplateRepoAttribute(rn, &repo),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init",
				},
			},
		},
	})
}

func TestAccGithubRepository_createFromForkForUser(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)
	updatedDescription := "Terraform acceptance tests updated"
	homepage := "http://example.com/"
	updatedHomepage := "http://example-new.com/"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryCreateFromForkForUser(randString, description, homepage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:                name,
						Description:         description,
						Homepage:            homepage,
						HasIssues:           true,
						HasWiki:             true,
						IsTemplate:          false,
						AllowMergeCommit:    true,
						AllowSquashMerge:    false,
						AllowRebaseMerge:    false,
						DeleteBranchOnMerge: false,
						HasDownloads:        true,
						HasProjects:         false,
						DefaultBranch:       "master",
						Archived:            false,
					}),
				),
			},
			{
				Config: testAccGithubRepositoryCreateFromForkForUser(randString, updatedDescription, updatedHomepage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:                name,
						Description:         updatedDescription,
						Homepage:            updatedHomepage,
						HasIssues:           true,
						HasWiki:             true,
						IsTemplate:          false,
						AllowMergeCommit:    true,
						AllowSquashMerge:    false,
						AllowRebaseMerge:    false,
						DeleteBranchOnMerge: false,
						HasDownloads:        true,
						HasProjects:         false,
						DefaultBranch:       "master",
						Archived:            false,
					}),
				),
			},

			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init", "fork_from",
				},
			},
		},
	})
}
func TestAccGithubRepository_createFromForkForOrg(t *testing.T) {
	var repo github.Repository

	rn := "github_repository.foo"
	name := "go-github"
	description := "Terraform acceptance tests updated"
	homepage := "http://example.com/"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryCreateFromForkForOrg(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:                name,
						HasIssues:           false,
						HasWiki:             false,
						IsTemplate:          false,
						AllowMergeCommit:    true,
						AllowSquashMerge:    true,
						AllowRebaseMerge:    true,
						DeleteBranchOnMerge: false,
						HasDownloads:        false,
						HasProjects:         false,
						DefaultBranch:       "master",
						Archived:            false,
					}),
				),
			},
			{
				Config: testAccGithubRepositoryCreateFromForkForOrgUpdate(description, homepage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists(rn, &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:                name,
						Description:         description,
						Homepage:            homepage,
						HasIssues:           false,
						HasWiki:             false,
						IsTemplate:          false,
						AllowMergeCommit:    true,
						AllowSquashMerge:    true,
						AllowRebaseMerge:    true,
						DeleteBranchOnMerge: false,
						HasDownloads:        false,
						HasProjects:         false,
						DefaultBranch:       "master",
						Archived:            false,
					}),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init", "fork_from", "fork_into_organization",
				},
			},
		},
	})
}

func testAccCheckGithubRepositoryExists(n string, repo *github.Repository) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		repoName := rs.Primary.ID
		if repoName == "" {
			return fmt.Errorf("No repository name is set")
		}

		org := testAccProvider.Meta().(*Organization)
		owner := org.name
		conn := org.v3client
		if parsedID := strings.Split(repoName, "/"); len(parsedID) == 2 {
			owner = parsedID[0]
			repoName = parsedID[1]
		}
		gotRepo, _, err := conn.Repositories.Get(context.TODO(), owner, repoName)
		if err != nil {
			return err
		}
		*repo = *gotRepo
		return nil
	}
}

func testAccCheckGithubRepositoryTemplateRepoAttribute(n string, repo *github.Repository) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if templateRepository := repo.GetTemplateRepository(); templateRepository.GetIsTemplate() != true {
			return fmt.Errorf("got repo %q; want %q", templateRepository, repo)
		}

		return nil
	}
}

type testAccGithubRepositoryExpectedAttributes struct {
	Name                string
	Description         string
	Homepage            string
	Private             bool
	HasDownloads        bool
	HasIssues           bool
	HasProjects         bool
	HasWiki             bool
	IsTemplate          bool
	AllowMergeCommit    bool
	AllowSquashMerge    bool
	AllowRebaseMerge    bool
	DeleteBranchOnMerge bool
	AutoInit            bool
	DefaultBranch       string
	LicenseTemplate     string
	GitignoreTemplate   string
	Archived            bool
	Topics              []string
}

func testAccCheckGithubRepositoryAttributes(repo *github.Repository, want *testAccGithubRepositoryExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if name := repo.GetName(); name != want.Name {
			return fmt.Errorf("got repo %q; want %q", name, want.Name)
		}
		if description := repo.GetDescription(); description != want.Description {
			return fmt.Errorf("got description %q; want %q", description, want.Description)
		}
		if homepage := repo.GetHomepage(); homepage != want.Homepage {
			return fmt.Errorf("got homepage URL %q; want %q", homepage, want.Homepage)
		}
		if private := repo.GetPrivate(); private != want.Private {
			return fmt.Errorf("got private %#v; want %#v", private, want.Private)
		}
		if hasIssues := repo.GetHasIssues(); hasIssues != want.HasIssues {
			return fmt.Errorf("got has issues %#v; want %#v", hasIssues, want.HasIssues)
		}
		if hasProjects := repo.GetHasProjects(); hasProjects != want.HasProjects {
			return fmt.Errorf("got has projects %#v; want %#v", hasProjects, want.HasProjects)
		}
		if hasWiki := repo.GetHasWiki(); hasWiki != want.HasWiki {
			return fmt.Errorf("got has wiki %#v; want %#v", hasWiki, want.HasWiki)
		}
		if isTemplate := repo.GetIsTemplate(); isTemplate != want.IsTemplate {
			return fmt.Errorf("got has IsTemplate %#v; want %#v", isTemplate, want.IsTemplate)
		}
		if allowMergeCommit := repo.GetAllowMergeCommit(); allowMergeCommit != want.AllowMergeCommit {
			return fmt.Errorf("got allow merge commit %#v; want %#v", allowMergeCommit, want.AllowMergeCommit)
		}
		if allowSquashMerge := repo.GetAllowSquashMerge(); allowSquashMerge != want.AllowSquashMerge {
			return fmt.Errorf("got allow squash merge %#v; want %#v", allowSquashMerge, want.AllowSquashMerge)
		}
		if allowRebaseMerge := repo.GetAllowRebaseMerge(); allowRebaseMerge != want.AllowRebaseMerge {
			return fmt.Errorf("got allow rebase merge %#v; want %#v", allowRebaseMerge, want.AllowRebaseMerge)
		}
		if hasDownloads := repo.GetHasDownloads(); hasDownloads != want.HasDownloads {
			return fmt.Errorf("got has downloads %#v; want %#v", hasDownloads, want.HasDownloads)
		}
		if len(want.Topics) != len(repo.Topics) {
			return fmt.Errorf("got topics %#v; want %#v", repo.Topics, want.Topics)
		}
		sort.Strings(repo.Topics)
		sort.Strings(want.Topics)
		for i := range want.Topics {
			if repo.Topics[i] != want.Topics[i] {
				return fmt.Errorf("got topics %#v; want %#v", repo.Topics, want.Topics)
			}
		}
		if defaultBranch := repo.GetDefaultBranch(); defaultBranch != want.DefaultBranch {
			return fmt.Errorf("got default branch %q; want %q", defaultBranch, want.DefaultBranch)
		}

		if autoInit := repo.GetAutoInit(); repo.AutoInit != nil {
			if autoInit != want.AutoInit {
				return fmt.Errorf("got auto init %t; want %t", autoInit, want.AutoInit)
			}
		}

		if gitignoreTemplate := repo.GetGitignoreTemplate(); repo.GitignoreTemplate != nil {
			if gitignoreTemplate != want.GitignoreTemplate {
				return fmt.Errorf("got gitignore_template %q; want %q", gitignoreTemplate, want.GitignoreTemplate)
			}
		}

		if licenseTemplate := repo.GetLicenseTemplate(); repo.LicenseTemplate != nil {
			if licenseTemplate != want.LicenseTemplate {
				return fmt.Errorf("got license_template %q; want %q", licenseTemplate, want.LicenseTemplate)
			}
		}

		// For the rest of these, we just want to make sure they've been
		// populated with something that seems somewhat reasonable.
		if fullName := repo.GetFullName(); !strings.HasSuffix(fullName, "/"+want.Name) {
			return fmt.Errorf("got full name %q; want to end with '/%s'", fullName, want.Name)
		}
		if cloneURL := repo.GetCloneURL(); !strings.HasSuffix(cloneURL, "/"+want.Name+".git") {
			return fmt.Errorf("got Clone URL %q; want to end with '/%s.git'", cloneURL, want.Name)
		}
		if cloneURL := repo.GetCloneURL(); !strings.HasPrefix(cloneURL, "https://") {
			return fmt.Errorf("got Clone URL %q; want to start with 'https://'", cloneURL)
		}
		if HTMLURL := repo.GetHTMLURL(); !strings.HasSuffix(HTMLURL, "/"+want.Name) {
			return fmt.Errorf("got HTML URL %q; want to end with '%s'", HTMLURL, want.Name)
		}
		if SSHURL := repo.GetSSHURL(); !strings.HasSuffix(SSHURL, "/"+want.Name+".git") {
			return fmt.Errorf("got SSH URL %q; want to end with '/%s.git'", SSHURL, want.Name)
		}
		if SSHURL := repo.GetSSHURL(); !strings.HasPrefix(SSHURL, "git@github.com:") {
			return fmt.Errorf("got SSH URL %q; want to start with 'git@github.com:'", SSHURL)
		}
		if gitURL := repo.GetGitURL(); !strings.HasSuffix(gitURL, "/"+want.Name+".git") {
			return fmt.Errorf("got git URL %q; want to end with '/%s.git'", gitURL, want.Name)
		}
		if gitURL := repo.GetGitURL(); !strings.HasPrefix(gitURL, "git://") {
			return fmt.Errorf("got git URL %q; want to start with 'git://'", gitURL)
		}
		if SVNURL := repo.GetSVNURL(); !strings.HasSuffix(SVNURL, "/"+want.Name) {
			return fmt.Errorf("got svn URL %q; want to end with '/%s'", SVNURL, want.Name)
		}
		if SVNURL := repo.GetSVNURL(); !strings.HasPrefix(SVNURL, "https://") {
			return fmt.Errorf("got svn URL %q; want to start with 'https://'", SVNURL)
		}

		return nil
	}
}

func testAccCheckGithubRepositoryDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client
	orgName := testAccProvider.Meta().(*Organization).name

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_repository" {
			continue
		}

		gotRepo, resp, err := conn.Repositories.Get(context.TODO(), orgName, rs.Primary.ID)
		if err == nil {
			if name := gotRepo.GetName(); gotRepo != nil && name == rs.Primary.ID {
				return fmt.Errorf("Repository %s/%s still exists", orgName, name)
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccCreateRepositoryBranch(branch, repository string) error {
	baseURL := os.Getenv("GITHUB_BASE_URL")
	org := os.Getenv("GITHUB_ORGANIZATION")
	token := os.Getenv("GITHUB_TOKEN")

	config := Config{
		BaseURL:      baseURL,
		Token:        token,
		Organization: org,
	}

	c, err := config.Clients()
	if err != nil {
		return fmt.Errorf("Error creating github client: %s", err)
	}
	client := c.(*Organization).v3client

	refs, _, err := client.Git.GetRefs(context.TODO(), org, repository, "heads")
	if err != nil {
		return fmt.Errorf("Error getting reference commit: %s", err)
	}
	ref := refs[0]

	newRef := &github.Reference{
		Ref: github.String(fmt.Sprintf("refs/heads/%s", branch)),
		Object: &github.GitObject{
			SHA: ref.Object.SHA,
		},
	}

	_, _, err = client.Git.CreateRef(context.TODO(), org, repository, newRef)
	if err != nil {
		return fmt.Errorf("Error creating git reference: %s", err)
	}

	return nil
}

func testAccGithubRepositoryConfig(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  description  = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = true
  has_wiki           = true
  is_template        = false
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads      = true
  auto_init          = false
}
`, randString, randString)
}

func testAccGithubRepositoryConfigHasProjects(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  has_projects = true
}
`, randString)
}

func testAccGithubRepositoryUpdateConfig(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  description  = "Updated Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = false
  has_wiki           = false
  is_template        = true
  allow_merge_commit = false
  allow_squash_merge = true
  allow_rebase_merge = true
  has_downloads      = false
}
`, randString, randString)
}

func testAccGithubRepositoryArchivedConfig(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  description  = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = true
  has_wiki           = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads      = true
  archived           = true
}
`, randString, randString)
}

func testAccGithubRepositoryConfigDefaultBranch(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  description  = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = true
  has_wiki           = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads      = true
  auto_init          = true
  default_branch     = "master"
}
`, randString, randString)
}

func testAccGithubRepositoryUpdateConfigDefaultBranch(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  description  = "Updated Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = true
  has_wiki           = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads      = true
  auto_init          = true
  default_branch     = "foo"
}
`, randString, randString)
}

func testAccGithubRepositoryConfigTemplates(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  description  = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = true
  has_wiki           = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads      = true

  license_template   = "ms-pl"
  gitignore_template = "C++"
}
`, randString, randString)
}

func testAccGithubRepositoryCreateFromTemplate(randString string) string {

	owner := os.Getenv("GITHUB_ORGANIZATION")
	repository := os.Getenv("GITHUB_TEMPLATE_REPOSITORY")

	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  description  = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  template {
    owner      = "%s"
    repository = "%s"
  }

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = true
  has_wiki           = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads      = true

}
`, randString, randString, owner, repository)
}

func testAccGithubRepositoryCreateFromForkForUser(name, description, homepage string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name         = "tf-acc-test-%s"
  description  = "%s"
  homepage_url = "%s"


  private     = false
  has_issues  = true
  has_wiki    = true
  is_template = false

  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads      = true
  auto_init          = true

}

resource "github_repository" "foo" {
  name                 = github_repository.test.name
  description          = github_repository.test.description
  homepage_url         = github_repository.test.homepage_url
  fork_from_repository = github_repository.test.full_name
  has_issues           = github_repository.test.has_issues
  has_wiki             = github_repository.test.has_wiki
  is_template          = github_repository.test.is_template

  allow_merge_commit = github_repository.test.allow_merge_commit
  allow_squash_merge = github_repository.test.allow_squash_merge
  allow_rebase_merge = github_repository.test.allow_rebase_merge
  has_downloads      = github_repository.test.has_downloads
  auto_init          = github_repository.test.auto_init

}
`, name, description, homepage)
}

func testAccGithubRepositoryCreateFromForkForOrg() string {
	return fmt.Sprintf(`
data "github_repositories" "test" {
  query = "repo:google/go-github"
}

resource "github_repository" "foo" {
  name                   = data.github_repositories.test.names[0]
  fork_from_repository   = data.github_repositories.test.full_names[0]
  fork_into_organization = "%s"
}
`, testOrganization)
}

func testAccGithubRepositoryCreateFromForkForOrgUpdate(description, homepage string) string {
	return fmt.Sprintf(`
data "github_repositories" "test" {
  query = "repo:google/go-github"
}

resource "github_repository" "foo" {
  name                   = data.github_repositories.test.names[0]
  description            = "%s"
  homepage_url           = "%s"
  fork_from_repository   = data.github_repositories.test.full_names[0]
  fork_into_organization = "%s"
}
`, description, homepage, testOrganization)
}

func testAccGithubRepositoryConfigTopics(randString string, topicList string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name = "tf-acc-test-%s"
  description = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  topics = [%s]
}
`, randString, randString, topicList)
}

func testAccGithubRepositoryConfigAutoInitForceNew(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name      = "tf-acc-test-%s"
  auto_init = false
}
`, randString)
}

func testAccGithubRepositoryConfigAutoInitForceNewUpdate(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name               = "tf-acc-test-%s"
  auto_init          = true
  license_template   = "mpl-2.0"
  gitignore_template = "Go"
}

resource "github_branch_protection" "repo_name_master" {
  repository = "${github_repository.foo.name}"
  branch     = "master"
}
`, randString)
}
