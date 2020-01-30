package github

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubTeam_basic(t *testing.T) {
	var team github.Team

	rn := "github_team.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-acc-test-%s", randString)
	updatedName := fmt.Sprintf("tf-acc-test-updated-%s", randString)
	description := "Terraform acc test group"
	updatedDescription := "Terraform acc test group - updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamExists(rn, &team),
					testAccCheckGithubTeamAttributes(&team, name, description, nil),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "description", description),
					resource.TestCheckResourceAttr(rn, "privacy", "secret"),
					resource.TestCheckNoResourceAttr(rn, "parent_team_id"),
					resource.TestCheckResourceAttr(rn, "ldap_dn", ""),
					resource.TestCheckResourceAttr(rn, "slug", name),
				),
			},
			{
				Config: testAccGithubTeamUpdateConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamExists(rn, &team),
					testAccCheckGithubTeamAttributes(&team, updatedName, updatedDescription, nil),
					resource.TestCheckResourceAttr(rn, "name", updatedName),
					resource.TestCheckResourceAttr(rn, "description", updatedDescription),
					resource.TestCheckResourceAttr(rn, "privacy", "closed"),
					resource.TestCheckNoResourceAttr(rn, "parent_team_id"),
					resource.TestCheckResourceAttr(rn, "ldap_dn", ""),
					resource.TestCheckResourceAttr(rn, "slug", updatedName),
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

func TestAccGithubTeam_slug(t *testing.T) {
	var team github.Team

	rn := "github_team.foo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("TF Acc Test %s", randString)
	description := "Terraform acc test group"
	expectedSlug := fmt.Sprintf("tf-acc-test-%s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamExists(rn, &team),
					testAccCheckGithubTeamAttributes(&team, name, description, nil),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "description", description),
					resource.TestCheckResourceAttr(rn, "privacy", "secret"),
					resource.TestCheckNoResourceAttr(rn, "parent_team_id"),
					resource.TestCheckResourceAttr(rn, "ldap_dn", ""),
					resource.TestCheckResourceAttr(rn, "slug", expectedSlug),
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

func TestAccGithubTeam_hierarchical(t *testing.T) {
	var parent, child github.Team

	rn := "github_team.parent"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	parentName := fmt.Sprintf("tf-acc-parent-%s", randString)
	childName := fmt.Sprintf("tf-acc-child-%s", randString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamHierarchicalConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamExists(rn, &parent),
					testAccCheckGithubTeamAttributes(&parent, parentName, "Terraform acc test parent team", nil),
					testAccCheckGithubTeamExists("github_team.child", &child),
					testAccCheckGithubTeamAttributes(&child, childName, "Terraform acc test child team", &parent),
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

func testAccCheckGithubTeamExists(n string, team *github.Team) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Team ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).v3client
		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return unconvertibleIdErr(rs.Primary.ID, err)
		}

		githubTeam, _, err := conn.Teams.GetTeam(context.TODO(), id)
		if err != nil {
			return err
		}
		*team = *githubTeam
		return nil
	}
}

func testAccCheckGithubTeamAttributes(team *github.Team, name, description string, parentTeam *github.Team) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *team.Name != name {
			return fmt.Errorf("Team name does not match: %s, %s", *team.Name, name)
		}

		if *team.Description != description {
			return fmt.Errorf("Team description does not match: %s, %s", *team.Description, description)
		}

		if parentTeam == nil && team.Parent != nil {
			return fmt.Errorf("Team parent ID was expected to be empty, but was %d", team.Parent.GetID())
		} else if parentTeam != nil && team.Parent == nil {
			return fmt.Errorf("Team parent ID was expected to be %d, but was not present", parentTeam.GetID())
		} else if parentTeam != nil && team.Parent.GetID() != parentTeam.GetID() {
			return fmt.Errorf("Team parent ID does not match: %d, %d", team.Parent.GetID(), parentTeam.GetID())
		}

		return nil
	}
}

func testAccCheckGithubTeamDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_team" {
			continue
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return unconvertibleIdErr(rs.Primary.ID, err)
		}

		team, resp, err := conn.Teams.GetTeam(context.TODO(), id)
		if err == nil {
			teamId := strconv.FormatInt(*team.ID, 10)
			if team != nil && teamId == rs.Primary.ID {
				return fmt.Errorf("Team still exists")
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccGithubTeamConfig(teamName string) string {
	return fmt.Sprintf(`
resource "github_team" "foo" {
  name        = "%s"
  description = "Terraform acc test group"
  privacy     = "secret"
}
`, teamName)
}

func testAccGithubTeamUpdateConfig(randString string) string {
	return fmt.Sprintf(`
resource "github_team" "foo" {
  name        = "tf-acc-test-updated-%s"
  description = "Terraform acc test group - updated"
  privacy     = "closed"
}
`, randString)
}

func testAccGithubTeamHierarchicalConfig(randString string) string {
	return fmt.Sprintf(`
resource "github_team" "parent" {
  name        = "tf-acc-parent-%s"
  description = "Terraform acc test parent team"
  privacy     = "closed"
}
resource "github_team" "child" {
  name           = "tf-acc-child-%s"
  description    = "Terraform acc test child team"
  privacy        = "closed"
  parent_team_id = "${github_team.parent.id}"
}
`, randString, randString)
}
