package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	resource.AddTestSweepers("github_fork", &resource.Sweeper{
		Name: "github_fork",
		F:    testSweepForkedRepositories,
	})
}

func deleteRepos(client *github.Client, repos []*github.Repository, owner string) string {
	errors := make([]string, 0)
	for _, r := range repos {
		if name := r.GetName(); strings.HasPrefix(name, "tf-acc-") || strings.HasPrefix(name, "foo-") {
			log.Printf("Destroying Repository %s", name)

			if _, err := client.Repositories.Delete(context.TODO(), owner, name); err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	return strings.Join(errors, "\n")
}

func testSweepForkedRepositories(region string) error {
	meta, err := sharedConfigForRegion(region)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: maxPerPage}}
	for {
		repos, resp, err := client.Repositories.ListByOrg(context.TODO(), meta.(*Organization).name, opts)
		if err != nil {
			return err
		}

		res := deleteRepos(client, repos, meta.(*Organization).name)
		if res != "" {
			return errors.New(res)
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	listOpts := &github.RepositoryListOptions{ListOptions: opts.ListOptions}
	for {
		repos, resp, err := client.Repositories.List(context.TODO(), "", listOpts)
		if err != nil {
			return err
		}

		res := deleteRepos(client, repos, testUser)
		if res != "" {
			return errors.New(res)
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil
}

func TestAccGithubFork_basic_user(t *testing.T) {
	var fork github.Repository

	rn := "github_fork.test"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("Terraform acceptance tests %s", randString)
	updatedName := acctest.RandomWithPrefix("tf-acc-test-update")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubForkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubForkUserConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubForkExists(rn, &fork),
					testAccCheckGithubForkAttributes(&fork, &testAccGithubForkExpectedAttributes{
						Name:                name,
						FullName:            fmt.Sprintf("%s/%s", testUser, name),
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
				Config: testAccGithubForkUserUpdateConfig(randString, updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubForkExists(rn, &fork),
					testAccCheckGithubForkAttributes(&fork, &testAccGithubForkExpectedAttributes{
						Name:                updatedName,
						FullName:            fmt.Sprintf("%s/%s", testUser, updatedName),
						Description:         "Test description updated",
						Homepage:            "http://example-test.com/",
						HasIssues:           true,
						HasWiki:             true,
						IsTemplate:          true,
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
					"fork_from_owner", "fork_from_repository",
				},
			},
		},
	})
}

func TestAccGithubFork_basic_org(t *testing.T) {
	var fork github.Repository

	rn := "github_fork.test"
	owner := "google"
	repo := "go-github"
	description := "Go library for accessing the GitHub API"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	updatedRepoName := fmt.Sprintf("tf-acc-test-%s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubForkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubForkOrgConfig(owner, repo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubForkExists(rn, &fork),
					testAccCheckGithubForkAttributes(&fork, &testAccGithubForkExpectedAttributes{
						Name:                repo,
						FullName:            fmt.Sprintf("%s/%s", testOrganization, repo),
						Description:         description,
						Homepage:            "https://pkg.go.dev/github.com/google/go-github/v31/github",
						HasIssues:           true,
						HasWiki:             false,
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
				Config: testAccGithubForkOrgUpdateConfig(owner, repo, updatedRepoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubForkExists(rn, &fork),
					testAccCheckGithubForkAttributes(&fork, &testAccGithubForkExpectedAttributes{
						Name:                updatedRepoName,
						FullName:            fmt.Sprintf("%s/%s", testOrganization, updatedRepoName),
						Description:         "Test description updated",
						Homepage:            "http://example.com/",
						HasIssues:           true,
						HasWiki:             false,
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
					"fork_from_owner", "fork_from_repository", "fork_into_organization",
				},
			},
		},
	})
}

func testAccCheckGithubForkExists(n string, repo *github.Repository) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No fork name is set")
		}

		forkParts := strings.Split(rs.Primary.ID, "/")
		owner := forkParts[0]
		repoName := forkParts[1]
		conn := testAccProvider.Meta().(*Organization).v3client
		gotRepo, _, err := conn.Repositories.Get(context.TODO(), owner, repoName)
		if err != nil {
			return err
		}
		*repo = *gotRepo
		return nil
	}
}

type testAccGithubForkExpectedAttributes struct {
	Name                string
	FullName            string
	Description         string
	Homepage            string
	Private             bool
	HasDownloads        bool
	HasIssues           bool
	HasProjects         bool
	HasWiki             bool
	IsTemplate          bool
	LicenseTemplate     string
	GitignoreTemplate   string
	AllowMergeCommit    bool
	AllowSquashMerge    bool
	AllowRebaseMerge    bool
	DeleteBranchOnMerge bool
	DefaultBranch       string
	Archived            bool
	Topics              []string
}

func testAccCheckGithubForkAttributes(repo *github.Repository, want *testAccGithubForkExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if name := repo.GetName(); name != want.Name {
			return fmt.Errorf("got repo %q; want %q", name, want.Name)
		}
		if fullName := repo.GetFullName(); fullName != want.FullName {
			return fmt.Errorf("got repo %q; want %q", fullName, want.FullName)
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

func testAccCheckGithubForkDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_fork" {
			continue
		}
		forkName := strings.Split(rs.Primary.ID, "/")
		owner := forkName[0]
		repoName := forkName[1]
		gotRepo, resp, err := conn.Repositories.Get(context.TODO(), owner, repoName)
		if err == nil {
			if gotRepo != nil && gotRepo.GetFullName() == rs.Primary.ID {
				return fmt.Errorf("Fork %s still exists", gotRepo.GetFullName())
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccGithubRepositoryResourceConfig(randString string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
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
  auto_init          = true
}
`, randString, randString)
}

func testAccGithubForkUserConfig(randString string) string {
	return testAccGithubRepositoryResourceConfig(randString) + `
resource "github_fork" "test" {
  fork_from_owner      = split("/", github_repository.test.full_name)[0]
  fork_from_repository = github_repository.test.name

  description  = github_repository.test.description
  homepage_url = github_repository.test.homepage_url

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = github_repository.test.has_issues
  has_wiki           = github_repository.test.has_wiki
  is_template        = github_repository.test.is_template
  allow_merge_commit = github_repository.test.allow_merge_commit
  allow_squash_merge = github_repository.test.allow_squash_merge
  allow_rebase_merge = github_repository.test.allow_rebase_merge
  has_downloads      = github_repository.test.has_downloads
}
`
}

func testAccGithubForkUserUpdateConfig(randString, name string) string {
	return testAccGithubRepositoryResourceConfig(randString) + fmt.Sprintf(`
resource "github_fork" "test" {
  name = "%s"
  fork_from_owner      = split("/", github_repository.test.full_name)[0]
  fork_from_repository = github_repository.test.name

  description  = "Test description updated"
  homepage_url = "http://example-test.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = github_repository.test.has_issues
  has_wiki           = github_repository.test.has_wiki
  is_template        = true
  allow_merge_commit = github_repository.test.allow_merge_commit
  allow_squash_merge = github_repository.test.allow_squash_merge
  allow_rebase_merge = github_repository.test.allow_rebase_merge
  has_downloads      = github_repository.test.has_downloads
}
`, name)
}

func testAccGithubForkOrgConfig(owner, repoName string) string {
	return fmt.Sprintf(`
data "github_repository" "test" {
  full_name = "%s/%s"
}

resource "github_fork" "test" {
  fork_from_owner        = split("/", data.github_repository.test.full_name)[0]
  fork_from_repository   = data.github_repository.test.name
  fork_into_organization = "%s"

  description  = data.github_repository.test.description
  homepage_url = data.github_repository.test.homepage_url

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = data.github_repository.test.has_issues
  has_wiki           = data.github_repository.test.has_wiki
  is_template        = false
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads      = data.github_repository.test.has_downloads
}
`, owner, repoName, testOrganization)
}

func testAccGithubForkOrgUpdateConfig(owner, repoName, updatedRepoName string) string {
	return fmt.Sprintf(`
data "github_repository" "test" {
  full_name = "%s/%s"
}

resource "github_fork" "test" {
  name                   = "%s"
  fork_from_owner        = split("/", data.github_repository.test.full_name)[0]
  fork_from_repository   = data.github_repository.test.name
  fork_into_organization = "%s"

  description  = "Test description updated"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_issues         = data.github_repository.test.has_issues
  has_wiki           = data.github_repository.test.has_wiki
  is_template        = false
  allow_merge_commit = true
  allow_squash_merge = false
  allow_rebase_merge = false
  has_downloads      = data.github_repository.test.has_downloads
}
`, owner, repoName, updatedRepoName, testOrganization)
}
