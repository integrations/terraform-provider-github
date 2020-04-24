package github

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubTeamRepository_basic(t *testing.T) {
	var repository github.Repository

	rn := "github_team_repository.test_team_test_repo"
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	repoName := fmt.Sprintf("tf-acc-test-team-%s", acctest.RandString(5))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamRepositoryConfig(randString, repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamRepositoryExists(rn, &repository),
					testAccCheckGithubTeamRepositoryRoleState("pull", &repository),
				),
			},
			{
				Config: testAccGithubTeamRepositoryUpdateConfig(randString, repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamRepositoryExists(rn, &repository),
					testAccCheckGithubTeamRepositoryRoleState("push", &repository),
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

func TestAccCheckGetPermissions(t *testing.T) {
	pullMap := map[string]bool{"pull": true, "triage": false, "push": false, "maintain": false, "admin": false}
	triageMap := map[string]bool{"pull": false, "triage": true, "push": false, "maintain": false, "admin": false}
	pushMap := map[string]bool{"pull": true, "triage": false, "push": true, "maintain": false, "admin": false}
	maintainMap := map[string]bool{"pull": false, "triage": false, "push": false, "maintain": true, "admin": false}
	adminMap := map[string]bool{"pull": true, "triage": false, "push": true, "maintain": false, "admin": true}
	errorMap := map[string]bool{"pull": false, "triage": false, "push": false, "maintain": false, "admin": false}

	pull, _ := getRepoPermission(pullMap)
	if pull != "pull" {
		t.Fatalf("Expected pull permission, actual: %s", pull)
	}

	triage, _ := getRepoPermission(triageMap)
	if triage != "triage" {
		t.Fatalf("Expected triage permission, actual: %s", triage)
	}

	push, _ := getRepoPermission(pushMap)
	if push != "push" {
		t.Fatalf("Expected push permission, actual: %s", push)
	}

	maintain, _ := getRepoPermission(maintainMap)
	if maintain != "maintain" {
		t.Fatalf("Expected maintain permission, actual: %s", maintain)
	}

	admin, _ := getRepoPermission(adminMap)
	if admin != "admin" {
		t.Fatalf("Expected admin permission, actual: %s", admin)
	}

	errPerm, err := getRepoPermission(errorMap)
	if err == nil {
		t.Fatalf("Expected an error getting permissions, actual: %v", errPerm)
	}
}

func testAccCheckGithubTeamRepositoryRoleState(role string, repository *github.Repository) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceRole, err := getRepoPermission(repository.GetPermissions())
		if err != nil {
			return err
		}

		if resourceRole != role {
			return fmt.Errorf("Team repository role %v in resource does match expected state of %v", resourceRole, role)
		}
		return nil
	}
}

func testAccCheckGithubTeamRepositoryExists(n string, repository *github.Repository) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No team repository ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).v3client
		orgId := testAccProvider.Meta().(*Organization).id
		orgName := testAccProvider.Meta().(*Organization).name

		teamIdString, repoName, err := parseTwoPartID(rs.Primary.ID, "team_id", "repository")
		if err != nil {
			return err
		}
		teamId, err := strconv.ParseInt(teamIdString, 10, 64)
		if err != nil {
			return unconvertibleIdErr(teamIdString, err)
		}

		var repo *github.Repository
		if testAccProvider.Meta().(*Organization).isEnterprise {
			repo, _, err = IsEnterpriseTeamRepoByID(context.TODO(),
				conn,
				teamId,
				orgName,
				repoName)
		} else {
			repo, _, err = conn.Teams.IsTeamRepoByID(context.TODO(),
				orgId,
				teamId,
				orgName,
				repoName)
		}

		if err != nil {
			return err
		}
		*repository = *repo
		return nil
	}
}

func testAccCheckGithubTeamRepositoryDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client
	orgId := testAccProvider.Meta().(*Organization).id
	orgName := testAccProvider.Meta().(*Organization).name

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_team_repository" {
			continue
		}

		teamIdString, repoName, err := parseTwoPartID(rs.Primary.ID, "team_id", "repository")
		if err != nil {
			return err
		}
		teamId, err := strconv.ParseInt(teamIdString, 10, 64)
		if err != nil {
			return unconvertibleIdErr(teamIdString, err)
		}

		var repo *github.Repository
		var resp *github.Response
		if testAccProvider.Meta().(*Organization).isEnterprise {
			repo, resp, err = IsEnterpriseTeamRepoByID(context.TODO(),
				conn,
				teamId,
				orgName,
				repoName)
		} else {
			repo, resp, err = conn.Teams.IsTeamRepoByID(context.TODO(),
				orgId,
				teamId,
				orgName,
				repoName)
		}

		if err == nil {
			if repo != nil &&
				buildTwoPartID(teamIdString, repo.GetName()) == rs.Primary.ID {
				return fmt.Errorf("Team repository still exists")
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccGithubTeamRepositoryConfig(randString, repoName string) string {
	return fmt.Sprintf(`
resource "github_team" "test_team" {
  name        = "tf-acc-test-team-repo-%s"
  description = "Terraform acc test group"
}

resource "github_repository" "test" {
  name = "%s"
}

resource "github_team_repository" "test_team_test_repo" {
  team_id    = "${github_team.test_team.id}"
  repository = "${github_repository.test.name}"
  permission = "pull"
}
`, randString, repoName)
}

func testAccGithubTeamRepositoryUpdateConfig(randString, repoName string) string {
	return fmt.Sprintf(`
resource "github_team" "test_team" {
  name        = "tf-acc-test-team-repo-%s"
  description = "Terraform acc test group"
}

resource "github_repository" "test" {
  name = "%s"
}

resource "github_team_repository" "test_team_test_repo" {
  team_id    = "${github_team.test_team.id}"
  repository = "${github_repository.test.name}"
  permission = "push"
}
`, randString, repoName)
}
