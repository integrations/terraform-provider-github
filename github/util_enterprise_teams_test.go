package github

import (
	"testing"
)

func TestBuildEnterpriseTeamMembershipID(t *testing.T) {
	got := buildEnterpriseTeamMembershipID("my-enterprise", "ent:my-team", "testuser")
	want := "my-enterprise/ent:my-team/testuser"
	if got != want {
		t.Fatalf("buildEnterpriseTeamMembershipID() = %q, want %q", got, want)
	}
}

func TestParseEnterpriseTeamMembershipID(t *testing.T) {
	t.Run("parses valid ID with slug containing ':'", func(t *testing.T) {
		enterprise, teamSlug, username, err := parseEnterpriseTeamMembershipID("my-enterprise/ent:my-team/testuser")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if enterprise != "my-enterprise" || teamSlug != "ent:my-team" || username != "testuser" {
			t.Fatalf("got (%q, %q, %q), want (my-enterprise, ent:my-team, testuser)", enterprise, teamSlug, username)
		}
	})

	t.Run("parses ID with slashes in username", func(t *testing.T) {
		enterprise, teamSlug, username, err := parseEnterpriseTeamMembershipID("ent/team/user/extra")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if enterprise != "ent" {
			t.Fatalf("enterprise = %q, want %q", enterprise, "ent")
		}
		if teamSlug != "team" {
			t.Fatalf("teamSlug = %q, want %q", teamSlug, "team")
		}
		if username != "user/extra" {
			t.Fatalf("username = %q, want %q", username, "user/extra")
		}
	})

	t.Run("returns error for invalid format", func(t *testing.T) {
		if _, _, _, err := parseEnterpriseTeamMembershipID("only-one-part"); err == nil {
			t.Fatal("expected error for invalid ID format, got nil")
		}
	})
}

func TestBuildEnterpriseTeamOrganizationsID(t *testing.T) {
	got := buildEnterpriseTeamOrganizationsID("my-enterprise", "ent:my-team")
	want := "my-enterprise/ent:my-team"
	if got != want {
		t.Fatalf("buildEnterpriseTeamOrganizationsID() = %q, want %q", got, want)
	}
}

func TestParseEnterpriseTeamOrganizationsID(t *testing.T) {
	t.Run("parses valid ID with slug containing ':'", func(t *testing.T) {
		enterprise, teamSlug, err := parseEnterpriseTeamOrganizationsID("my-enterprise/ent:my-team")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if enterprise != "my-enterprise" || teamSlug != "ent:my-team" {
			t.Fatalf("got (%q, %q), want (my-enterprise, ent:my-team)", enterprise, teamSlug)
		}
	})

	t.Run("returns error for invalid format", func(t *testing.T) {
		if _, _, err := parseEnterpriseTeamOrganizationsID("no-slash-here"); err == nil {
			t.Fatal("expected error for invalid ID format, got nil")
		}
	})
}
