package github

import (
	"context"
	"fmt"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubRepositoryProject_basic(t *testing.T) {
	randRepoName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	var project github.Project

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubRepositoryProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryProjectConfig(randRepoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryProjectExists("github_repository_project.test", &project),
					testAccCheckGithubRepositoryProjectAttributes(&project, &testAccGithubRepositoryProjectExpectedAttributes{
						Name:       "test-project",
						Repository: randRepoName,
						Body:       "this is a test project",
					}),
				),
			},
		},
	})
}

func TestAccGithubRepositoryProject_importBasic(t *testing.T) {
	randRepoName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubRepositoryProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryProjectConfig(randRepoName),
			},
			{
				ResourceName:        "github_repository_project.test",
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: randRepoName + `/`,
			},
		},
	})
}

func testAccGithubRepositoryProjectDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_repository_project" {
			continue
		}

		projectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		project, res, err := conn.Projects.GetProject(context.TODO(), projectID)
		if err == nil {
			if project != nil &&
				project.GetID() == projectID {
				return fmt.Errorf("Repository project still exists")
			}
		}
		if res.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccCheckGithubRepositoryProjectExists(n string, project *github.Project) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		projectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		conn := testAccProvider.Meta().(*Organization).client
		gotProject, _, err := conn.Projects.GetProject(context.TODO(), projectID)
		if err != nil {
			return err
		}
		*project = *gotProject
		return nil
	}
}

type testAccGithubRepositoryProjectExpectedAttributes struct {
	Name       string
	Repository string
	Body       string
}

func testAccCheckGithubRepositoryProjectAttributes(project *github.Project, want *testAccGithubRepositoryProjectExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if *project.Name != want.Name {
			return fmt.Errorf("got project %q; want %q", *project.Name, want.Name)
		}
		if got := path.Base(project.GetOwnerURL()); got != want.Repository {
			return fmt.Errorf("got project %q; want %q", got, want.Repository)
		}
		if *project.Body != want.Body {
			return fmt.Errorf("got project n%q; want %q", *project.Body, want.Body)
		}
		if !strings.HasPrefix(*project.URL, "https://") {
			return fmt.Errorf("got http URL %q; want to start with 'https://'", *project.URL)
		}

		return nil
	}
}

func testAccGithubRepositoryProjectConfig(repoName string) string {
	return fmt.Sprintf(`
resource "github_repository" "foo" {
  name         = "%[1]s"
  description  = "Terraform acceptance tests"
  homepage_url = "http://example.com/"

  # So that acceptance tests can be run in a github organization
  # with no billing
  private = false

  has_projects  = true
  has_issues    = true
  has_wiki      = true
  has_downloads = true
}

resource "github_repository_project" "test" {
  depends_on = ["github_repository.foo"]

  name       = "test-project"
  repository = "%[1]s"
  body       = "this is a test project"
}`, repoName)
}
