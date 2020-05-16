package github

import (
	"context"
	"log"
	"os"
	"regexp"
	"strings"

	"encoding/base64"
	"fmt"
	"testing"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// The authenticated user's name used for commits should be exported as GITHUB_TEST_USER_NAME
var userName string = os.Getenv("GITHUB_TEST_USER_NAME")

// The authenticated user's email address used for commits should be exported as GITHUB_TEST_USER_EMAIL
var userEmail string = os.Getenv("GITHUB_TEST_USER_EMAIL")

func init() {
	resource.AddTestSweepers("github_repository_file", &resource.Sweeper{
		Name: "github_repository_file",
		F:    testSweepRepositoryFiles,
	})

}

func testSweepRepositoryFiles(region string) error {
	meta, err := sharedConfigForRegion(region)
	if err != nil {
		return err
	}

	if err := testSweepDeleteRepositoryFiles(meta, "master"); err != nil {
		return err
	}

	if err := testSweepDeleteRepositoryFiles(meta, "test-branch"); err != nil {
		return err
	}

	return nil
}

func testSweepDeleteRepositoryFiles(meta interface{}, branch string) error {
	client := meta.(*Organization).v3client
	org := meta.(*Organization).name

	_, files, _, err := client.Repositories.GetContents(
		context.TODO(), org, "test-repo", "", &github.RepositoryContentGetOptions{Ref: branch})
	if err != nil {
		return err
	}

	for _, f := range files {
		if name := f.GetName(); strings.HasPrefix(name, "tf-acc-") {
			log.Printf("Deleting repository file: %s, repo: %s/test-repo, branch: %s", name, org, branch)
			opts := &github.RepositoryContentFileOptions{Branch: github.String(branch)}
			if _, _, err := client.Repositories.DeleteFile(context.TODO(), org, "test-repo", name, opts); err != nil {
				return err
			}
		}
	}

	return nil
}

func TestAccGithubRepositoryFile_basic(t *testing.T) {
	if userName == "" {
		t.Skip("This test requires you to set the test user's name (set it by exporting GITHUB_TEST_USER_NAME)")
	}

	if userEmail == "" {
		t.Skip("This test requires you to set the test user's email address (set it by exporting GITHUB_TEST_USER_EMAIL)")
	}

	var content github.RepositoryContent
	var commit github.RepositoryCommit

	rn := "github_repository_file.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	path := fmt.Sprintf("tf-acc-test-file-%s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryFileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryFileConfig(
					path, "Terraform acceptance test file"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryFileExists(rn, path, "master", &content, &commit),
					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file")) + "\n",
					}),
					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
						Branch:        "master",
						CommitAuthor:  userName,
						CommitEmail:   userEmail,
						CommitMessage: fmt.Sprintf("Add %s", path),
						Filename:      path,
					}),
					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "file", path),
					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file"),
					resource.TestCheckResourceAttr(rn, "commit_author", userName),
					resource.TestCheckResourceAttr(rn, "commit_email", userEmail),
					resource.TestCheckResourceAttr(rn, "commit_message", fmt.Sprintf("Add %s", path)),
				),
			},
			{
				Config: testAccGithubRepositoryFileConfig(
					path, "Terraform acceptance test file updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryFileExists(rn, path, "master", &content, &commit),
					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file updated")) + "\n",
					}),
					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
						Branch:        "master",
						CommitAuthor:  userName,
						CommitEmail:   userEmail,
						CommitMessage: fmt.Sprintf("Update %s", path),
						Filename:      path,
					}),
					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "file", path),
					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file updated"),
					resource.TestCheckResourceAttr(rn, "commit_author", userName),
					resource.TestCheckResourceAttr(rn, "commit_email", userEmail),
					resource.TestCheckResourceAttr(rn, "commit_message", fmt.Sprintf("Update %s", path)),
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

func TestAccGithubRepositoryFile_branch(t *testing.T) {
	if userName == "" {
		t.Skip("This test requires you to set the test user's name (set it by exporting GITHUB_TEST_USER_NAME)")
	}

	if userEmail == "" {
		t.Skip("This test requires you to set the test user's email address (set it by exporting GITHUB_TEST_USER_EMAIL)")
	}

	var content github.RepositoryContent
	var commit github.RepositoryCommit

	rn := "github_repository_file.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	path := fmt.Sprintf("tf-acc-test-file-%s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryFileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryFileBranchConfig(
					path, "Terraform acceptance test file"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryFileExists(rn, path, "test-branch", &content, &commit),
					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file")) + "\n",
					}),
					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
						Branch:        "test-branch",
						CommitAuthor:  userName,
						CommitEmail:   userEmail,
						CommitMessage: fmt.Sprintf("Add %s", path),
						Filename:      path,
					}),
					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
					resource.TestCheckResourceAttr(rn, "branch", "test-branch"),
					resource.TestCheckResourceAttr(rn, "file", path),
					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file"),
					resource.TestCheckResourceAttr(rn, "commit_author", userName),
					resource.TestCheckResourceAttr(rn, "commit_email", userEmail),
					resource.TestCheckResourceAttr(rn, "commit_message", fmt.Sprintf("Add %s", path)),
				),
			},
			{
				Config: testAccGithubRepositoryFileBranchConfig(
					path, "Terraform acceptance test file updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryFileExists(rn, path, "test-branch", &content, &commit),
					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file updated")) + "\n",
					}),
					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
						Branch:        "test-branch",
						CommitAuthor:  userName,
						CommitEmail:   userEmail,
						CommitMessage: fmt.Sprintf("Update %s", path),
						Filename:      path,
					}),
					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
					resource.TestCheckResourceAttr(rn, "branch", "test-branch"),
					resource.TestCheckResourceAttr(rn, "file", path),
					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file updated"),
					resource.TestCheckResourceAttr(rn, "commit_author", userName),
					resource.TestCheckResourceAttr(rn, "commit_email", userEmail),
					resource.TestCheckResourceAttr(rn, "commit_message", fmt.Sprintf("Update %s", path)),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateId:     fmt.Sprintf("test-repo/%s:test-branch", path),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubRepositoryFile_committer(t *testing.T) {
	var content github.RepositoryContent
	var commit github.RepositoryCommit

	rn := "github_repository_file.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	path := fmt.Sprintf("tf-acc-test-file-%s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryFileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryFileCommitterConfig(
					path, "Terraform acceptance test file"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryFileExists(rn, path, "master", &content, &commit),
					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file")) + "\n",
					}),
					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
						Branch:        "master",
						CommitAuthor:  "Terraform User",
						CommitEmail:   "terraform@example.com",
						CommitMessage: "Managed by Terraform",
						Filename:      path,
					}),
					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "file", path),
					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file"),
					resource.TestCheckResourceAttr(rn, "commit_author", "Terraform User"),
					resource.TestCheckResourceAttr(rn, "commit_email", "terraform@example.com"),
					resource.TestCheckResourceAttr(rn, "commit_message", "Managed by Terraform"),
				),
			},
			{
				Config: testAccGithubRepositoryFileCommitterConfig(
					path, "Terraform acceptance test file updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryFileExists(rn, path, "master", &content, &commit),
					testAccCheckGithubRepositoryFileAttributes(&content, &testAccGithubRepositoryFileExpectedAttributes{
						Content: base64.StdEncoding.EncodeToString([]byte("Terraform acceptance test file updated")) + "\n",
					}),
					testAccCheckGithubRepositoryFileCommitAttributes(&commit, &testAccGithubRepositoryFileExpectedCommitAttributes{
						Branch:        "master",
						CommitAuthor:  "Terraform User",
						CommitEmail:   "terraform@example.com",
						CommitMessage: "Managed by Terraform",
						Filename:      path,
					}),
					resource.TestCheckResourceAttr(rn, "repository", "test-repo"),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "file", path),
					resource.TestCheckResourceAttr(rn, "content", "Terraform acceptance test file updated"),
					resource.TestCheckResourceAttr(rn, "commit_author", "Terraform User"),
					resource.TestCheckResourceAttr(rn, "commit_email", "terraform@example.com"),
					resource.TestCheckResourceAttr(rn, "commit_message", "Managed by Terraform"),
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

func TestAccGithubRepositoryFile_overwrite(t *testing.T) {
	var content github.RepositoryContent
	var commit github.RepositoryCommit

	rn := "github_repository_file.foo"
	path := ".gitignore"
	fileContent := "**/.terraform/*"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	repo := fmt.Sprintf("tf-acc-test-%s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryFileDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccGithubRepositoryFileOverwriteDisabledConfig(randString),
				ExpectError: regexp.MustCompile(`Refusing to overwrite existing file`),
			},
			{
				Config: testAccGithubRepositoryFileOverwriteEnabledConfig(randString, path, fileContent),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryFileExistsWithRepo(rn, path, "master", repo, &content, &commit),
					testAccCheckGithubRepositoryFileAttributes(
						&content, &testAccGithubRepositoryFileExpectedAttributes{
							Content: base64.StdEncoding.EncodeToString([]byte(fileContent)) + "\n",
						},
					),
				),
			},
		},
	})
}

func testAccCheckGithubRepositoryFileExists(n, path, branch string, content *github.RepositoryContent, commit *github.RepositoryCommit) resource.TestCheckFunc {
	hardCodedRepoName := "test-repo"
	return testAccCheckGithubRepositoryFileExistsWithRepo(n, path, branch, hardCodedRepoName, content, commit)
}

func testAccCheckGithubRepositoryFileExistsWithRepo(n, path, branch, repo string, content *github.RepositoryContent, commit *github.RepositoryCommit) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No repository file path set")
		}

		conn := testAccProvider.Meta().(*Organization).v3client
		org := testAccProvider.Meta().(*Organization).name

		opts := &github.RepositoryContentGetOptions{Ref: branch}
		gotContent, _, _, err := conn.Repositories.GetContents(context.TODO(), org, repo, path, opts)
		if err != nil {
			return err
		}

		gotCommit, err := getFileCommit(conn, org, "test-repo", path, branch)
		if err != nil {
			return err
		}

		*content = *gotContent
		*commit = *gotCommit

		return nil
	}
}

type testAccGithubRepositoryFileExpectedAttributes struct {
	Content string
}

func testAccCheckGithubRepositoryFileAttributes(content *github.RepositoryContent, want *testAccGithubRepositoryFileExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if *content.Content != want.Content {
			return fmt.Errorf("got content %q; want %q", *content.Content, want.Content)
		}

		return nil
	}
}

type testAccGithubRepositoryFileExpectedCommitAttributes struct {
	Branch        string
	CommitAuthor  string
	CommitEmail   string
	CommitMessage string
	Filename      string
}

func testAccCheckGithubRepositoryFileCommitAttributes(commit *github.RepositoryCommit, want *testAccGithubRepositoryFileExpectedCommitAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if name := commit.GetCommit().GetCommitter().GetName(); name != want.CommitAuthor {
			return fmt.Errorf("got committer author name %q; want %q", name, want.CommitAuthor)
		}

		if email := commit.GetCommit().GetCommitter().GetEmail(); email != want.CommitEmail {
			return fmt.Errorf("got committer author email %q; want %q", email, want.CommitEmail)
		}

		if message := commit.GetCommit().GetMessage(); message != want.CommitMessage {
			return fmt.Errorf("got commit message %q; want %q", message, want.CommitMessage)
		}

		if len(commit.Files) != 1 {
			return fmt.Errorf("got multiple files in commit (%q); expected 1", len(commit.Files))
		}

		file := commit.Files[0]
		if filename := file.GetFilename(); filename != want.Filename {
			return fmt.Errorf("got filename %q; want %q", filename, want.Filename)
		}

		return nil
	}
}

func testAccCheckGithubRepositoryFileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client
	org := testAccProvider.Meta().(*Organization).name

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_repository_file" {
			continue
		}

		repo, file := splitRepoFilePath(rs.Primary.ID)
		opts := &github.RepositoryContentGetOptions{Ref: rs.Primary.Attributes["branch"]}

		fc, _, resp, err := conn.Repositories.GetContents(context.TODO(), org, repo, file, opts)
		if err == nil {
			if fc != nil {
				return fmt.Errorf("Repository file %s/%s/%s still exists", org, repo, file)
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccGithubRepositoryFileConfig(file, content string) string {
	return fmt.Sprintf(`
resource "github_repository_file" "foo" {
  repository = "test-repo"
  file       = "%s"
  content    = "%s"
}
`, file, content)
}

func testAccGithubRepositoryFileBranchConfig(file, content string) string {
	return fmt.Sprintf(`
resource "github_repository_file" "foo" {
  repository = "test-repo"
  branch     = "test-branch"
  file       = "%s"
  content    = "%s"
}
`, file, content)
}

func testAccGithubRepositoryFileCommitterConfig(file, content string) string {
	return fmt.Sprintf(`
resource "github_repository_file" "foo" {
  repository     = "test-repo"
  file           = "%s"
  content        = "%s"
  commit_message = "Managed by Terraform"
  commit_author  = "Terraform User"
  commit_email   = "terraform@example.com"
}
`, file, content)
}

func testAccGithubRepositoryFileOverwriteDisabledConfig(randString string) string {

	owner := testOrganization
	repository := os.Getenv("GITHUB_TEMPLATE_REPOSITORY")

	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  description  = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

	template {
		owner = "%s"
		repository = "%s"
	}

  private            = false
  has_issues         = false
  has_wiki           = false

}

resource "github_repository_file" "foo" {
  repository     = github_repository.foo.name
  file           = ".gitignore"
  content        = "**/.terraform/*"
  commit_message = "Managed by Terraform"
  commit_author  = "Terraform User"
  commit_email   = "terraform@example.com"
}
`, randString, randString, owner, repository)
}

func testAccGithubRepositoryFileOverwriteEnabledConfig(randString, path, content string) string {

	owner := testOrganization
	templateRepository := os.Getenv("GITHUB_TEMPLATE_REPOSITORY")

	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "tf-acc-test-%s"
  description  = "Terraform acceptance tests %s"
  homepage_url = "http://example.com/"

	template {
		owner = "%s"
		repository = "%s"
	}

  private            = false
  has_issues         = false
  has_wiki           = false
}

resource "github_repository_file" "foo" {
  repository     = github_repository.foo.name
  file           = "%s"
  content        = "%s"
  commit_message = "Managed by Terraform"
  commit_author  = "Terraform User"
  commit_email   = "terraform@example.com"
  overwrite      = true
}
`, randString, randString, owner, templateRepository, path, content)
}
