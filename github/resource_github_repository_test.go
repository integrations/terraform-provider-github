package github

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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

	client := meta.(*Organization).client

	repos, _, err := client.Repositories.List(context.TODO(), meta.(*Organization).name, nil)
	if err != nil {
		return err
	}

	for _, r := range repos {
		if strings.HasPrefix(*r.Name, "tf-acc-") || strings.HasPrefix(*r.Name, "foo-") {
			log.Printf("Destroying Repository %s", *r.Name)

			if _, err := client.Repositories.Delete(context.TODO(), meta.(*Organization).name, *r.Name); err != nil {
				return err
			}
		}
	}

	return nil
}

func TestAccGithubRepository_basic(t *testing.T) {
	var repo github.Repository
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
						HasProjects:      false,
						DefaultBranch:    "master",
						Archived:         false,
					}),
				),
			},
			{
				Config: testAccGithubRepositoryUpdateConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
					testAccCheckGithubRepositoryAttributes(&repo, &testAccGithubRepositoryExpectedAttributes{
						Name:             name,
						Description:      "Updated " + description,
						Homepage:         "http://example.com/",
						AllowMergeCommit: false,
						AllowSquashMerge: true,
						AllowRebaseMerge: true,
						DefaultBranch:    "master",
						HasProjects:      false,
						Archived:         false,
					}),
				),
			},
		},
	})
}

func TestAccGithubRepository_archive(t *testing.T) {
	var repo github.Repository
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryArchivedConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
		},
	})
}

func TestAccGithubRepository_archiveUpdate(t *testing.T) {
	var repo github.Repository
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
		},
	})
}

func TestAccGithubRepository_importBasic(t *testing.T) {
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfig(randString),
			},
			{
				ResourceName:      "github_repository.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubRepository_defaultBranch(t *testing.T) {
	var repo github.Repository
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfigDefaultBranch(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
					testAccCreateRepositoryBranch("foo", *repo.Name)
				},
				Config: testAccGithubRepositoryUpdateConfigDefaultBranch(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
		},
	})
}

func TestAccGithubRepository_templates(t *testing.T) {
	var repo github.Repository
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfigTemplates(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
		},
	})
}

func TestAccGithubRepository_topics(t *testing.T) {
	var repo github.Repository
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfigTopics(randString, `"topic1", "topic2"`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
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
		},
	})
}

func TestAccGithubRepository_autoInitForceNew(t *testing.T) {
	var repo github.Repository
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryConfigAutoInitForceNew(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
					resource.TestCheckResourceAttr("github_repository.foo", "name", name),
					resource.TestCheckResourceAttr("github_repository.foo", "auto_init", "false"),
				),
			},
			{
				Config: testAccGithubRepositoryConfigAutoInitForceNewUpdate(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryExists("github_repository.foo", &repo),
					resource.TestCheckResourceAttr("github_repository.foo", "name", name),
					resource.TestCheckResourceAttr("github_repository.foo", "auto_init", "true"),
					resource.TestCheckResourceAttr("github_repository.foo", "license_template", "mpl-2.0"),
					resource.TestCheckResourceAttr("github_repository.foo", "gitignore_template", "Go"),
				),
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
		conn := org.client
		gotRepo, _, err := conn.Repositories.Get(context.TODO(), org.name, repoName)
		if err != nil {
			return err
		}
		*repo = *gotRepo
		return nil
	}
}

type testAccGithubRepositoryExpectedAttributes struct {
	Name              string
	Description       string
	Homepage          string
	Private           bool
	HasDownloads      bool
	HasIssues         bool
	HasProjects       bool
	HasWiki           bool
	AllowMergeCommit  bool
	AllowSquashMerge  bool
	AllowRebaseMerge  bool
	AutoInit          bool
	DefaultBranch     string
	LicenseTemplate   string
	GitignoreTemplate string
	Archived          bool
	Topics            []string
}

func testAccCheckGithubRepositoryAttributes(repo *github.Repository, want *testAccGithubRepositoryExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if *repo.Name != want.Name {
			return fmt.Errorf("got repo %q; want %q", *repo.Name, want.Name)
		}
		if *repo.Description != want.Description {
			return fmt.Errorf("got description %q; want %q", *repo.Description, want.Description)
		}
		if *repo.Homepage != want.Homepage {
			return fmt.Errorf("got homepage URL %q; want %q", *repo.Homepage, want.Homepage)
		}
		if *repo.Private != want.Private {
			return fmt.Errorf("got private %#v; want %#v", *repo.Private, want.Private)
		}
		if *repo.HasIssues != want.HasIssues {
			return fmt.Errorf("got has issues %#v; want %#v", *repo.HasIssues, want.HasIssues)
		}
		if *repo.HasWiki != want.HasWiki {
			return fmt.Errorf("got has wiki %#v; want %#v", *repo.HasWiki, want.HasWiki)
		}
		if *repo.AllowMergeCommit != want.AllowMergeCommit {
			return fmt.Errorf("got allow merge commit %#v; want %#v", *repo.AllowMergeCommit, want.AllowMergeCommit)
		}
		if *repo.AllowSquashMerge != want.AllowSquashMerge {
			return fmt.Errorf("got allow squash merge %#v; want %#v", *repo.AllowSquashMerge, want.AllowSquashMerge)
		}
		if *repo.AllowRebaseMerge != want.AllowRebaseMerge {
			return fmt.Errorf("got allow rebase merge %#v; want %#v", *repo.AllowRebaseMerge, want.AllowRebaseMerge)
		}
		if *repo.HasDownloads != want.HasDownloads {
			return fmt.Errorf("got has downloads %#v; want %#v", *repo.HasDownloads, want.HasDownloads)
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
		if *repo.DefaultBranch != want.DefaultBranch {
			return fmt.Errorf("got default branch %q; want %q", *repo.DefaultBranch, want.DefaultBranch)
		}

		if repo.AutoInit != nil {
			if *repo.AutoInit != want.AutoInit {
				return fmt.Errorf("got auto init %t; want %t", *repo.AutoInit, want.AutoInit)
			}
		}

		if repo.GitignoreTemplate != nil {
			if *repo.GitignoreTemplate != want.GitignoreTemplate {
				return fmt.Errorf("got gitignore_template %q; want %q", *repo.GitignoreTemplate, want.GitignoreTemplate)
			}
		}

		if repo.LicenseTemplate != nil {
			if *repo.LicenseTemplate != want.LicenseTemplate {
				return fmt.Errorf("got license_template %q; want %q", *repo.LicenseTemplate, want.LicenseTemplate)
			}
		}

		// For the rest of these, we just want to make sure they've been
		// populated with something that seems somewhat reasonable.
		if !strings.HasSuffix(*repo.FullName, "/"+want.Name) {
			return fmt.Errorf("got full name %q; want to end with '/%s'", *repo.FullName, want.Name)
		}
		if !strings.HasSuffix(*repo.CloneURL, "/"+want.Name+".git") {
			return fmt.Errorf("got Clone URL %q; want to end with '/%s.git'", *repo.CloneURL, want.Name)
		}
		if !strings.HasPrefix(*repo.CloneURL, "https://") {
			return fmt.Errorf("got Clone URL %q; want to start with 'https://'", *repo.CloneURL)
		}
		if !strings.HasSuffix(*repo.HTMLURL, "/"+want.Name) {
			return fmt.Errorf("got HTML URL %q; want to end with '%s'", *repo.HTMLURL, want.Name)
		}
		if !strings.HasSuffix(*repo.SSHURL, "/"+want.Name+".git") {
			return fmt.Errorf("got SSH URL %q; want to end with '/%s.git'", *repo.SSHURL, want.Name)
		}
		if !strings.HasPrefix(*repo.SSHURL, "git@github.com:") {
			return fmt.Errorf("got SSH URL %q; want to start with 'git@github.com:'", *repo.SSHURL)
		}
		if !strings.HasSuffix(*repo.GitURL, "/"+want.Name+".git") {
			return fmt.Errorf("got git URL %q; want to end with '/%s.git'", *repo.GitURL, want.Name)
		}
		if !strings.HasPrefix(*repo.GitURL, "git://") {
			return fmt.Errorf("got git URL %q; want to start with 'git://'", *repo.GitURL)
		}
		if !strings.HasSuffix(*repo.SVNURL, "/"+want.Name) {
			return fmt.Errorf("got svn URL %q; want to end with '/%s'", *repo.SVNURL, want.Name)
		}
		if !strings.HasPrefix(*repo.SVNURL, "https://") {
			return fmt.Errorf("got svn URL %q; want to start with 'https://'", *repo.SVNURL)
		}

		return nil
	}
}

func testAccCheckGithubRepositoryDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client
	orgName := testAccProvider.Meta().(*Organization).name

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_repository" {
			continue
		}

		gotRepo, resp, err := conn.Repositories.Get(context.TODO(), orgName, rs.Primary.ID)
		if err == nil {
			if gotRepo != nil && *gotRepo.Name == rs.Primary.ID {
				return fmt.Errorf("Repository %s/%s still exists", orgName, *gotRepo.Name)
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
	org := os.Getenv("GITHUB_ORGANIZATION")
	token := os.Getenv("GITHUB_TOKEN")

	config := Config{
		Token:        token,
		Organization: org,
	}

	c, err := config.Client()
	if err != nil {
		return fmt.Errorf("Error creating github client: %s", err)
	}
	client := c.(*Organization).client

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
  name = "tf-acc-test-%s"
  description = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues = true
  has_wiki = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads = true
  auto_init = false
}
`, randString, randString)
}

func testAccGithubRepositoryUpdateConfig(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name = "tf-acc-test-%s"
  description = "Updated Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues = false
  has_wiki = false
  allow_merge_commit = false
  allow_squash_merge = true
  allow_rebase_merge = true
  has_downloads = false
}
`, randString, randString)
}

func testAccGithubRepositoryArchivedConfig(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name = "tf-acc-test-%s"
  description = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues = true
  has_wiki = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads = true
  archived = true
}
`, randString, randString)
}

func testAccGithubRepositoryConfigDefaultBranch(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name = "tf-acc-test-%s"
  description = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues = true
  has_wiki = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads = true
  auto_init = true
}
`, randString, randString)
}

func testAccGithubRepositoryUpdateConfigDefaultBranch(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name = "tf-acc-test-%s"
  description = "Updated Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues = true
  has_wiki = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads = true
  auto_init = true
  default_branch = "foo"
}
`, randString, randString)
}

func testAccGithubRepositoryConfigTemplates(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name = "tf-acc-test-%s"
  description = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues = true
  has_wiki = true
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads = true

  license_template = "ms-pl"
  gitignore_template = "C++"
}
`, randString, randString)
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
  name = "tf-acc-test-%s"
  auto_init = false
}
`, randString)
}

func testAccGithubRepositoryConfigAutoInitForceNewUpdate(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name = "tf-acc-test-%s"
  auto_init = true
  license_template = "mpl-2.0"
  gitignore_template = "Go"
}

resource "github_branch_protection" "repo_name_master" {
  repository = "${github_repository.foo.name}"
  branch = "master"
}
`, randString)
}
