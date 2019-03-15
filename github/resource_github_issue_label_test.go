package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubIssueLabel_basic(t *testing.T) {
	var label github.Label

	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-%s", rString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubIssueLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubIssueLabelConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists("github_issue_label.test", &label),
					testAccCheckGithubIssueLabelAttributes(&label, "foo", "000000"),
				),
			},
			{
				Config: testAccGithubIssueLabelUpdateConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists("github_issue_label.test", &label),
					testAccCheckGithubIssueLabelAttributes(&label, "bar", "FFFFFF"),
				),
			},
		},
	})
}

func TestAccGithubIssueLabel_disappears(t *testing.T) {
	var label github.Label

	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-%s", rString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubIssueLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubIssueLabelConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists("github_issue_label.test", &label),
					testAccCheckGithubIssueLabelDisappears("github_issue_label.test"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccGithubIssueLabel_many(t *testing.T) {
	// TF parallelism is the default of 10, so this should be set to more than that
	// to ensure you test the case of multiple batches based on timing.
	// TODO: this should probably also test larger than the page size,
	// but its currently 100, so that may exhaust the API in testing.
	const numberToTest = 20

	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-%s", rString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubIssueLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubIssueLabelManyConfig(repoName, numberToTest),
			},
			{
				Config: testAccGithubIssueLabelManyConfig(repoName, numberToTest+1),
			},
			{
				Config: testAccGithubIssueLabelManyConfig(repoName, numberToTest-1),
			},
		},
	})
}

func TestAccGithubIssueLabel_nameExists(t *testing.T) {
	var label github.Label

	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-%s", rString)
	labelName := fmt.Sprintf("foo-%s", rString)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(s *terraform.State) error {
			conn := testAccProvider.Meta().(*Organization).client
			orgName := testAccProvider.Meta().(*Organization).name

			// ignore errors
			_, _ = conn.Repositories.Delete(context.TODO(), orgName, repoName)

			return testAccGithubIssueLabelDestroy(s)
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(*Organization).client
					orgName := testAccProvider.Meta().(*Organization).name

					_, _, err := conn.Repositories.Create(context.TODO(), orgName, &github.Repository{
						Name: &repoName,
					})
					if err != nil {
						t.Fatal(err)
					}

					_, _, err = conn.Issues.CreateLabel(context.TODO(), orgName, repoName, &github.Label{
						Name: &labelName,
					})
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccGitHubIssueLabelNameExistsConfig(repoName, labelName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists("github_issue_label.test", &label),
					testAccCheckGithubIssueLabelAttributes(&label, labelName, "FF00FF"),
				),
			},
		},
	})
}

func TestAccGithubIssueLabel_initialLabel(t *testing.T) {
	var label github.Label

	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-%s", rString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubIssueLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGitHubIssueLabelInitialConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists("github_issue_label.test", &label),
					testAccCheckGithubIssueLabelAttributes(&label, "enhancement", "FF00FF"),
				),
			},
		},
	})
}

func TestAccGithubIssueLabel_importBasic(t *testing.T) {
	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-%s", rString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubIssueLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubIssueLabelConfig(repoName),
			},
			{
				ResourceName:      "github_issue_label.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubIssueLabel_description(t *testing.T) {
	var label github.Label

	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-issue-label-desc-%s", rString)
	description := "Terraform Acceptance Test"
	updatedDescription := "Terraform Acceptance Test Updated"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubIssueLabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubIssueLabelConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists("github_issue_label.test", &label),
					resource.TestCheckResourceAttr("github_issue_label.test", "description", ""),
				),
			},
			{
				Config: testAccGithubIssueLabelConfig_description(repoName, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists("github_issue_label.test", &label),
					resource.TestCheckResourceAttr("github_issue_label.test", "description", description),
				),
			},
			{
				Config: testAccGithubIssueLabelConfig_description(repoName, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists("github_issue_label.test", &label),
					resource.TestCheckResourceAttr("github_issue_label.test", "description", updatedDescription),
				),
			},
			{
				Config: testAccGithubIssueLabelConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIssueLabelExists("github_issue_label.test", &label),
					resource.TestCheckResourceAttr("github_issue_label.test", "description", ""),
				),
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

		conn := testAccProvider.Meta().(*Organization).client
		orgName := testAccProvider.Meta().(*Organization).name
		repoName, name, err := parseTwoPartID(rs.Primary.ID)
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

func testAccCheckGithubIssueLabelDisappears(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No issue label ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).client
		orgName := testAccProvider.Meta().(*Organization).name
		repoName, name, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = conn.Issues.DeleteLabel(context.TODO(), orgName, repoName, name)
		return err
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
	conn := testAccProvider.Meta().(*Organization).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_issue_label" {
			continue
		}

		orgName := testAccProvider.Meta().(*Organization).name
		repoName, name, err := parseTwoPartID(rs.Primary.ID)
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

func testAccGitHubIssueLabelNameExistsConfig(repoName, labelName string) string {
	return fmt.Sprintf(`
resource "github_issue_label" "test" {
  repository = "%s"
  name       = "%s"
  color      = "FF00FF"
}
`, repoName, labelName)
}

func testAccGithubIssueLabelManyConfig(repoName string, count int) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name = "%s"
}

resource "github_issue_label" "test" {
  repository = "${github_repository.test.name}"
  name       = "foo${count.index}"
  color      = "000000"

  count = "%d"
}
`, repoName, count)
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

func testAccGitHubIssueLabelInitialConfig(repoName string) string {
	return fmt.Sprintf(`
// Create a repository which has the default labels
resource "github_repository" "test" {
  name = "%s"
}

resource "github_issue_label" "test" {
  repository = "${github_repository.test.name}"
  name       = "enhancement" // Important! This is an initial / pre-created label
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
  repository = "${github_repository.test.name}"
  name       = "foo"
  color      = "000000"
  description = "%s"
}
`, repoName, description)
}
