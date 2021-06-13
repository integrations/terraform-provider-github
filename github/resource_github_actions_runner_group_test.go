package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-github/v35/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubActionsRunnerGroup_all(t *testing.T) {
	// ???
	// resource_github_actions_runner_group_test.go:19: Skipping because GITHUB_OWNER is a user, not an organization.
	// if err := testAccCheckOrganization(); err != nil {
	// t.Skipf("Skipping because %s.", err.Error())
	// }

	var runnerGroup github.RunnerGroup

	var testAccGithubActionsRunnerGroupConfigAll = `
resource "github_actions_runner_group" "test_all" {
  name = "test-runner-group-all"
  visibility = "all"
}
`
	rn := "github_actions_runner_group.test_all"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubActionsRunnerGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubActionsRunnerGroupConfigAll,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubActionsRunnerGroupExists(rn, &runnerGroup),
					testAccCheckGithubActionsRunnerGroupAttributes(&runnerGroup, &testAccGithubActionsRunnerGroupExpectedAttributes{
						Name:                     "test-runner-group-all",
						Visibility:               "all",
						Default:                  false,
						AllowsPublicRepositories: false,
						RunnersURL:               fmt.Sprintf(`https://api.github.com/orgs/octo-org/actions/runner_groups/%d/runners`, runnerGroup.ID),
						SelectedRepositoriesURL:  fmt.Sprintf(`https://api.github.com/orgs/octo-org/actions/runner_groups/%d/repositories`, runnerGroup.ID),
					}),
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

func TestAccGithubActionsRunnerGroup_private(t *testing.T) {
	// ???
	// resource_github_actions_runner_group_test.go:19: Skipping because GITHUB_OWNER is a user, not an organization.
	// if err := testAccCheckOrganization(); err != nil {
	// t.Skipf("Skipping because %s.", err.Error())
	// }

	var runnerGroup github.RunnerGroup

	rn := "github_actions_runner_group.test_private"
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	var testAccGithubActionsRunnerGroupConfigPrivate = fmt.Sprintf(`
resource "github_repository" "test" {
  name = "tf-acc-test-%s"
  visibility = "private"
}
resource "github_actions_runner_group" "test_private" {
  name = "test-runner-group-private"
  visibility = "private"
}
`, randomID)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubActionsRunnerGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubActionsRunnerGroupConfigPrivate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubActionsRunnerGroupExists(rn, &runnerGroup),
					testAccCheckGithubActionsRunnerGroupAttributes(&runnerGroup, &testAccGithubActionsRunnerGroupExpectedAttributes{
						Name:                     "test-runner-group-private",
						Visibility:               "private",
						Default:                  false,
						AllowsPublicRepositories: false,
						RunnersURL:               fmt.Sprintf(`https://api.github.com/orgs/octo-org/actions/runner_groups/%d/runners`, runnerGroup.ID),
						SelectedRepositoriesURL:  fmt.Sprintf(`https://api.github.com/orgs/octo-org/actions/runner_groups/%d/repositories`, runnerGroup.ID),
					}),
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

func TestAccGithubActionsRunnerGroup_selected(t *testing.T) {
	// ???
	// resource_github_actions_runner_group_test.go:19: Skipping because GITHUB_OWNER is a user, not an organization.
	// if err := testAccCheckOrganization(); err != nil {
	// t.Skipf("Skipping because %s.", err.Error())
	// }

	var runnerGroup github.RunnerGroup

	rn := "github_actions_runner_group.test_selected"
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	var testAccGithubActionsRunnerGroupConfigSelected = fmt.Sprintf(`
resource "github_repository" "test" {
  name = "tf-acc-test-%s"
  auto_init = true
}

resource "github_actions_runner_group" "test_selected" {
  name = "test-runner-group-selected"
  visibility = "selected"
  selected_repository_ids = [github_repository.test.repo_id]
}
`, randomID)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubActionsRunnerGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubActionsRunnerGroupConfigSelected,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubActionsRunnerGroupExists(rn, &runnerGroup),
					testAccCheckGithubActionsRunnerGroupAttributes(&runnerGroup, &testAccGithubActionsRunnerGroupExpectedAttributes{
						Name:                     "test-runner-group-selected",
						Visibility:               "selected",
						Default:                  false,
						AllowsPublicRepositories: false,
						RunnersURL:               fmt.Sprintf(`https://api.github.com/orgs/octo-org/actions/runner_groups/%d/runners`, runnerGroup.ID),
						SelectedRepositoriesURL:  fmt.Sprintf(`https://api.github.com/orgs/octo-org/actions/runner_groups/%d/repositories`, runnerGroup.ID),
					}),
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

func testAccGithubActionsRunnerGroupDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Owner).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_actions_runner_group" {
			continue
		}

		runnerGroupID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		orgName := testAccProvider.Meta().(*Owner).name
		runnerGroup, res, err := conn.Actions.GetOrganizationRunnerGroup(context.TODO(), orgName, runnerGroupID)
		if err == nil {
			if runnerGroup != nil &&
				runnerGroup.GetID() == runnerGroupID {
				return fmt.Errorf("Organization runner group still exists")
			}
		}
		if res.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccCheckGithubActionsRunnerGroupExists(n string, runnerGroup *github.RunnerGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		runnerGroupID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		conn := testAccProvider.Meta().(*Owner).v3client
		orgName := testAccProvider.Meta().(*Owner).name
		gotRunnerGroup, _, err := conn.Actions.GetOrganizationRunnerGroup(context.TODO(), orgName, runnerGroupID)
		if err != nil {
			return err
		}
		*runnerGroup = *gotRunnerGroup
		return nil
	}
}

type testAccGithubActionsRunnerGroupExpectedAttributes struct {
	AllowsPublicRepositories bool
	Default                  bool
	ID                       int64
	Inherited                bool
	Name                     string
	Runners                  []int64
	RunnersURL               string
	SelectedRepositoriesURL  string
	SelectedRepositoryIDs    []int64
	Visibility               string
}

func testAccCheckGithubActionsRunnerGroupAttributes(runnerGroup *github.RunnerGroup, want *testAccGithubActionsRunnerGroupExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if name := runnerGroup.GetName(); name != want.Name {
			return fmt.Errorf("got runnerGroup name %q; want %q", name, want.Name)
		}
		if visibility := runnerGroup.GetVisibility(); visibility != want.Visibility {
			return fmt.Errorf("got runnerGroup visibility %q; want %q", visibility, want.Visibility)
		}
		if inherited := runnerGroup.GetInherited(); inherited != want.Inherited {
			return fmt.Errorf("got runnerGroup inherited %t; want %t", inherited, want.Inherited)
		}
		if URL := runnerGroup.GetRunnersURL(); !strings.HasPrefix(URL, "https://") {
			return fmt.Errorf("got runners URL %q; want to start with 'https://'", URL)
		}
		if isDefault := runnerGroup.GetDefault(); isDefault != want.Default {
			return fmt.Errorf("got runnerGroup default %t; want %t", isDefault, want.Default)
		}

		return nil
	}
}
