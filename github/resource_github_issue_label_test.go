package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubIssueLabel_basic(t *testing.T) {
	var label github.Label

	rn := "github_issue_label.test"
	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-%s", rString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubIssueLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubIssueLabelConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists(rn, &label),
					testAccCheckGithubIssueLabelAttributes(&label, "foo", "000000"),
				),
			},
			{
				Config: testAccGithubIssueLabelUpdateConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists(rn, &label),
					testAccCheckGithubIssueLabelAttributes(&label, "bar", "FFFFFF"),
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

func TestAccGithubIssueLabel_existingLabel(t *testing.T) {
	var label github.Label

	rn := "github_issue_label.test"
	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-%s", rString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubIssueLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGitHubIssueLabelExistsConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists(rn, &label),
					testAccCheckGithubIssueLabelAttributes(&label, "enhancement", "FF00FF"),
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

func TestAccGithubIssueLabel_description(t *testing.T) {
	var label github.Label

	rn := "github_issue_label.test"
	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-desc-%s", rString)
	description := "Terraform Acceptance Test"
	updatedDescription := "Terraform Acceptance Test Updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubIssueLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubIssueLabelConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists(rn, &label),
					resource.TestCheckResourceAttr(rn, "description", ""),
				),
			},
			{
				Config: testAccGithubIssueLabelConfig_description(repoName, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists(rn, &label),
					resource.TestCheckResourceAttr(rn, "description", description),
				),
			},
			{
				Config: testAccGithubIssueLabelConfig_description(repoName, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists(rn, &label),
					resource.TestCheckResourceAttr(rn, "description", updatedDescription),
				),
			},
			{
				Config: testAccGithubIssueLabelConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists(rn, &label),
					resource.TestCheckResourceAttr(rn, "description", ""),
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

func testAccCheckGithubIssueLabelExists(n string, label *github.Label) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No issue label ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).v3client
		orgName := testAccProvider.Meta().(*Organization).name
		repoName, name, err := parseTwoPartID(rs.Primary.ID, "repository", "name")
		if err != nil {
			return err
		}

		githubLabel, _, err := conn.Issues.GetLabel(context.TODO(),
			orgName, repoName, name)
		if err != nil {
			return err
		}

		*label = *githubLabel
		return nil
	}
}

func testAccCheckGithubIssueLabelAttributes(label *github.Label, name, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *label.Name != name {
			return fmt.Errorf("Issue label name does not match: %s, %s", *label.Name, name)
		}

		if *label.Color != color {
			return fmt.Errorf("Issue label color does not match: %s, %s", *label.Color, color)
		}

		return nil
	}
}

func testAccGithubIssueLabelDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_issue_label" {
			continue
		}

		orgName := testAccProvider.Meta().(*Organization).name
		repoName, name, err := parseTwoPartID(rs.Primary.ID, "repository", "name")
		if err != nil {
			return err
		}

		label, res, err := conn.Issues.GetLabel(context.TODO(),
			orgName, repoName, name)
		if err == nil {
			if label != nil &&
				buildTwoPartID(label.Name, label.Color) == rs.Primary.ID {
				return fmt.Errorf("Issue label still exists")
			}
		}
		if res.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccGithubIssueLabelConfig(repoName string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name = "%s"
}

resource "github_issue_label" "test" {
  repository = "${github_repository.test.name}"
  name       = "foo"
  color      = "000000"
}
`, repoName)
}

func testAccGithubIssueLabelUpdateConfig(repoName string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name = "%s"
}

resource "github_issue_label" "test" {
  repository = "${github_repository.test.name}"
  name       = "bar"
  color      = "FFFFFF"
}
`, repoName)
}

func testAccGitHubIssueLabelExistsConfig(repoName string) string {
	return fmt.Sprintf(`
// Create a repository which has the default labels
resource "github_repository" "test" {
  name = "%s"
}

resource "github_issue_label" "test" {
  repository = "${github_repository.test.name}"
  name       = "enhancement" // Important! This is a pre-created label
  color      = "FF00FF"
}
`, repoName)
}

func testAccGithubIssueLabelConfig_description(repoName, description string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name = "%s"
}

resource "github_issue_label" "test" {
  repository  = "${github_repository.test.name}"
  name        = "foo"
  color       = "000000"
  description = "%s"
}
`, repoName, description)
}
