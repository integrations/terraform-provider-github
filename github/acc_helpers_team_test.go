package github

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

type createTestTeamOptionsFunc func(*github.NewTeam)

func withNewTeamParent(parentID int64) createTestTeamOptionsFunc {
	return func(team *github.NewTeam) {
		team.ParentTeamID = &parentID
	}
}

func mustCreateTestTeam(t *testing.T, f ...createTestTeamOptionsFunc) *github.Team {
	t.Helper()

	randomID := acctest.RandString(testRandomIDLength)
	name := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

	req := &github.NewTeam{
		Name:    name,
		Privacy: new("closed"),
	}

	for _, fn := range f {
		fn(req)
	}

	team, _, err := testAccConf.meta.v3client.Teams.CreateTeam(t.Context(), testAccConf.meta.name, *req)
	if err != nil {
		t.Fatalf("failed to create test team: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Teams.DeleteTeamByID(context.Background(), testAccConf.meta.id, team.GetID()); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test team %s: %v", name, err)
		}
	})

	return team
}

func mustRenameTeam(t *testing.T, team *github.Team, newName string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Teams.EditTeamBySlug(t.Context(), testAccConf.meta.name, team.GetSlug(), github.NewTeam{Name: newName}, false)
	if err != nil {
		t.Fatalf("failed to rename test team %s to %s: %v", team.GetName(), newName, err)
	}
}

func mustDeleteTeam(t *testing.T, team *github.Team) {
	t.Helper()

	if _, err := testAccConf.meta.v3client.Teams.DeleteTeamBySlug(context.Background(), testAccConf.meta.name, team.GetSlug()); err != nil {
		if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
			return
		}
		t.Fatalf("failed to delete test team %s: %v", team.GetName(), err)
	}
}

func mustAddTeamMember(t *testing.T, team *github.Team, username string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Teams.AddTeamMembershipBySlug(t.Context(), testAccConf.meta.name, team.GetSlug(), username, &github.TeamAddTeamMembershipOptions{Role: "member"})
	if err != nil {
		t.Fatalf("failed to add member %s to test team %s: %v", username, team.GetName(), err)
	}
}

func mustAddTeamMaintainer(t *testing.T, team *github.Team, username string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Teams.AddTeamMembershipBySlug(t.Context(), testAccConf.meta.name, team.GetSlug(), username, &github.TeamAddTeamMembershipOptions{Role: "maintainer"})
	if err != nil {
		t.Fatalf("failed to add member %s to test team %s: %v", username, team.GetName(), err)
	}
}

func mustAssignOrganizationRoleToTeam(t *testing.T, team *github.Team, roleID int64) {
	t.Helper()

	_, err := testAccConf.meta.v3client.Organizations.AssignOrgRoleToTeam(t.Context(), testAccConf.meta.name, team.GetSlug(), roleID)
	if err != nil {
		t.Fatalf("failed to assign organization role %d to team %s: %v", roleID, team.GetName(), err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Organizations.RemoveOrgRoleFromTeam(context.Background(), testAccConf.meta.name, team.GetSlug(), roleID); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to unassign organization role %d from team %s: %v", roleID, team.GetName(), err)
		}
	})
}

func mustAddRepositoryToTeam(t *testing.T, team *github.Team, repo *github.Repository) {
	t.Helper()

	_, err := testAccConf.meta.v3client.Teams.AddTeamRepoByID(t.Context(), testAccConf.meta.id, team.GetID(), testAccConf.meta.name, repo.GetName(), &github.TeamAddTeamRepoOptions{Permission: "pull"})
	if err != nil {
		t.Fatalf("failed to add team %s to test repository %s: %v", team.GetName(), repo.GetName(), err)
	}
}
