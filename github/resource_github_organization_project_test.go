package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubOrganizationProject_basic(t *testing.T) {
	var project github.Project

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubOrganizationProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubOrganizationProjectConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubOrganizationProjectExists("github_organization_project.test", &project),
					testAccCheckGithubOrganizationProjectAttributes(&project, &testAccGithubOrganizationProjectExpectedAttributes{
						Name: "test-project",
						Body: "this is a test project",
					}),
				),
			},
		},
	})
}

func TestAccGithubOrganizationProject_importBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubOrganizationProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubOrganizationProjectConfig,
			},
			{
				ResourceName:      "github_organization_project.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGithubOrganizationProjectDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_organization_project" {
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
				return fmt.Errorf("Organization project still exists")
			}
		}
		if res.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccCheckGithubOrganizationProjectExists(n string, project *github.Project) resource.TestCheckFunc {
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

type testAccGithubOrganizationProjectExpectedAttributes struct {
	Name string
	Body string
}

func testAccCheckGithubOrganizationProjectAttributes(project *github.Project, want *testAccGithubOrganizationProjectExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if *project.Name != want.Name {
			return fmt.Errorf("got project %q; want %q", *project.Name, want.Name)
		}
		if *project.Body != want.Body {
			return fmt.Errorf("got project %q; want %q", *project.Body, want.Body)
		}
		if !strings.HasPrefix(*project.URL, "https://") {
			return fmt.Errorf("got http URL %q; want to start with 'https://'", *project.URL)
		}

		return nil
	}
}

const testAccGithubOrganizationProjectConfig = `
resource "github_organization_project" "test" {
  name = "test-project"
  body = "this is a test project"
}
`
