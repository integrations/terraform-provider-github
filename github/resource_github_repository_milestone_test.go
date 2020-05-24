package github

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubRepositoryMilestone_basic(t *testing.T) {
	var milestone github.Milestone
	title := acctest.RandomWithPrefix("tf-acc-test")
	repoName := acctest.RandomWithPrefix("tf-acc-test-repo")
	rn := "github_repository_milestone.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryMilestoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryMilestoneConfig(repoName, title),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryMilestoneExists(rn, &milestone),
					resource.TestCheckResourceAttr(rn, "title", title),
					resource.TestCheckResourceAttr(rn, "description", ""),
					resource.TestCheckResourceAttr(rn, "state", "open"),
					resource.TestCheckResourceAttr(rn, "number", "1"),
				),
			},
			{
				ResourceName:      rn,
				ImportStateIdFunc: testAccGithubRepositoryMilestoneImportStateIdFunc(rn),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubRepositoryMilestone_update(t *testing.T) {
	var milestone github.Milestone
	title := acctest.RandomWithPrefix("tf-acc-test")
	repoName := acctest.RandomWithPrefix("tf-acc-test-repo")
	dueDate := time.Now().UTC().Format(layoutISO)
	description := acctest.RandomWithPrefix("tf-acc-test-desc")
	rn := "github_repository_milestone.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryMilestoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryMilestoneConfig(repoName, title),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryMilestoneExists(rn, &milestone),
					resource.TestCheckResourceAttr(rn, "title", title),
					resource.TestCheckResourceAttr(rn, "description", ""),
					resource.TestCheckResourceAttr(rn, "state", "open"),
					resource.TestCheckResourceAttr(rn, "number", "1"),
				),
			},
			{
				ResourceName:      rn,
				ImportStateIdFunc: testAccGithubRepositoryMilestoneImportStateIdFunc(rn),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGithubRepositoryMilestoneConfigUpdate(repoName, title, description, dueDate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryMilestoneExists(rn, &milestone),
					resource.TestCheckResourceAttr(rn, "title", title),
					resource.TestCheckResourceAttr(rn, "description", description),
					resource.TestCheckResourceAttr(rn, "due_date", dueDate),
					resource.TestCheckResourceAttr(rn, "state", "open"),
					resource.TestCheckResourceAttr(rn, "number", "1"),
				),
			},
		},
	})
}

func TestAccGithubRepositoryMilestone_multiple(t *testing.T) {
	var milestone1, milestone2 github.Milestone
	titleOne := acctest.RandomWithPrefix("tf-acc-test-ms")
	titleTwo := acctest.RandomWithPrefix("tf-acc-test-ms-two")

	repoName := acctest.RandomWithPrefix("tf-acc-test-repo")
	dueDate := time.Now().UTC().Format(layoutISO)
	descriptionOne := acctest.RandomWithPrefix("tf-acc-test-desc")
	descriptionTwo := acctest.RandomWithPrefix("tf-acc-test-desc-two")

	rn := "github_repository_milestone.test"
	rnTwo := "github_repository_milestone.test_two"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryMilestoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryMultipleMilestoneConfig(repoName, titleOne, descriptionOne, titleTwo, descriptionTwo, dueDate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryMilestoneExists(rn, &milestone1),
					resource.TestCheckResourceAttr(rn, "title", titleOne),
					resource.TestCheckResourceAttr(rn, "description", descriptionOne),
					resource.TestCheckResourceAttr(rn, "due_date", dueDate),
					resource.TestCheckResourceAttr(rn, "state", "closed"),
					resource.TestCheckResourceAttrSet(rn, "number"),
					testAccCheckGithubRepositoryMilestoneExists(rnTwo, &milestone2),
					resource.TestCheckResourceAttr(rnTwo, "title", titleTwo),
					resource.TestCheckResourceAttr(rnTwo, "description", descriptionTwo),
					resource.TestCheckResourceAttr(rn, "due_date", dueDate),
					resource.TestCheckResourceAttr(rnTwo, "state", "open"),
					resource.TestCheckResourceAttrSet(rnTwo, "number"),
				),
			},
			{
				ResourceName:      rn,
				ImportStateIdFunc: testAccGithubRepositoryMilestoneImportStateIdFunc(rn),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      rnTwo,
				ImportStateIdFunc: testAccGithubRepositoryMilestoneImportStateIdFunc(rnTwo),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGithubRepositoryMilestoneConfigUpdate(repoName, titleOne, descriptionOne, dueDate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryMilestoneExists(rn, &milestone1),
					resource.TestCheckResourceAttr(rn, "title", titleOne),
					resource.TestCheckResourceAttr(rn, "description", descriptionOne),
					resource.TestCheckResourceAttr(rn, "due_date", dueDate),
					resource.TestCheckResourceAttr(rn, "state", "open"),
					resource.TestCheckResourceAttrSet(rn, "number"),
				),
			},
		},
	})
}

func TestAccGithubRepositoryMilestone_disappears(t *testing.T) {
	var milestone github.Milestone

	title := acctest.RandomWithPrefix("tf-acc-test")
	repoName := acctest.RandomWithPrefix("tf-acc-test-repo")
	rn := "github_repository_milestone.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryMilestoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryMilestoneConfig(title, repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryMilestoneExists(rn, &milestone),
					testAccCheckGithubRepositoryMilestoneDisappears(rn),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckGithubRepositoryMilestoneExists(rn string, milestone *github.Milestone) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("not Found: %s", rn)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Milestone ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).v3client
		number, err := parseMilestoneNumber(rs.Primary.ID)
		if err != nil {
			return err
		}
		parts := strings.Split(rs.Primary.ID, "/")
		owner := parts[0]
		repoName := parts[1]

		m, _, err := conn.Issues.GetMilestone(context.TODO(), owner, repoName, number)
		if err != nil {
			return err
		}
		*milestone = *m
		return nil
	}
}

func testAccCheckGithubRepositoryMilestoneDisappears(rn string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("Not found: %s", rn)
		}
		conn := testAccProvider.Meta().(*Organization).v3client
		number, err := parseMilestoneNumber(rs.Primary.ID)
		if err != nil {
			return err
		}
		parts := strings.Split(rs.Primary.ID, "/")
		owner := parts[0]
		repoName := parts[1]

		_, err = conn.Issues.DeleteMilestone(context.TODO(), owner, repoName, number)
		return err
	}
}

func testAccCheckGithubRepositoryMilestoneDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_repository_milestone" {
			continue
		}

		number, err := parseMilestoneNumber(rs.Primary.ID)
		if err != nil {
			return err
		}
		parts := strings.Split(rs.Primary.ID, "/")
		owner := parts[0]
		repoName := parts[1]

		milestone, resp, err := conn.Issues.GetMilestone(context.TODO(), owner, repoName, number)
		if err != nil {
			if resp.StatusCode != 404 {
				return err
			}
		}
		if milestone != nil {
			return fmt.Errorf("milestone still exists")
		}

		return nil
	}

	return nil
}

func testAccGithubRepositoryMilestoneImportStateIdFunc(rn string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return "", fmt.Errorf("Not found: %s", rn)
		}

		return fmt.Sprintf("%s/%s/%s", rs.Primary.Attributes["owner"], rs.Primary.Attributes["repository"], rs.Primary.Attributes["number"]), nil
	}
}

func testAccGithubRepositoryMilestoneConfig(repoName, title string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_milestone" "test" {
	owner = split("/", "${github_repository.test.full_name}")[0]
	repository = github_repository.test.name
    title = "%s"
}
`, repoName, title)
}

func testAccGithubRepositoryMilestoneConfigUpdate(repoName, title, description, dueDate string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_milestone" "test" {
	owner = split("/", "${github_repository.test.full_name}")[0]
	repository = github_repository.test.name
    title = "%s"
    description = "%s"
    due_date = "%s"
}
`, repoName, title, description, dueDate)
}

func testAccGithubRepositoryMultipleMilestoneConfig(repoName, titleOne, descriptionOne, titleTwo, descriptionTwo, dueDate string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_milestone" "test" {
	owner = split("/", "${github_repository.test.full_name}")[0]
	repository = github_repository.test.name
    title = "%s"
    description = "%s"
    due_date = "%s"
    state = "closed"
}

resource "github_repository_milestone" "test_two" {
	owner = split("/", "${github_repository.test.full_name}")[0]
	repository = github_repository.test.name
    title = "%s"
    description = "%s"
    due_date = "%s"
    state = "open"
}
`, repoName, titleOne, descriptionOne, dueDate, titleTwo, descriptionTwo, dueDate)
}
