package github

import (
	"testing"
)

func TestBuildEnterpriseTeamMembershipID(t *testing.T) {
	t.Run("builds correct ID format", func(t *testing.T) {
		got := buildEnterpriseTeamMembershipID("my-enterprise", "ent:my-team", "testuser")
		want := "my-enterprise/ent:my-team/testuser"
		if got != want {
			t.Fatalf("buildEnterpriseTeamMembershipID() = %q, want %q", got, want)
		}
	})

	t.Run("handles empty strings", func(t *testing.T) {
		got := buildEnterpriseTeamMembershipID("", "", "")
		want := "//"
		if got != want {
			t.Fatalf("buildEnterpriseTeamMembershipID() = %q, want %q", got, want)
		}
	})
}

func TestParseEnterpriseTeamMembershipID(t *testing.T) {
	t.Run("parses valid ID", func(t *testing.T) {
		enterprise, teamSlug, username, err := parseEnterpriseTeamMembershipID("my-enterprise/ent:my-team/testuser")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if enterprise != "my-enterprise" {
			t.Fatalf("enterprise = %q, want %q", enterprise, "my-enterprise")
		}
		if teamSlug != "ent:my-team" {
			t.Fatalf("teamSlug = %q, want %q", teamSlug, "ent:my-team")
		}
		if username != "testuser" {
			t.Fatalf("username = %q, want %q", username, "testuser")
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
		_, _, _, err := parseEnterpriseTeamMembershipID("only-one-part")
		if err == nil {
			t.Fatal("expected error for invalid ID format, got nil")
		}
	})

	t.Run("returns error for empty string", func(t *testing.T) {
		_, _, _, err := parseEnterpriseTeamMembershipID("")
		if err == nil {
			t.Fatal("expected error for empty ID, got nil")
		}
	})

	t.Run("roundtrips with build function", func(t *testing.T) {
		id := buildEnterpriseTeamMembershipID("enterprise", "ent:team-slug", "user123")
		enterprise, teamSlug, username, err := parseEnterpriseTeamMembershipID(id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if enterprise != "enterprise" || teamSlug != "ent:team-slug" || username != "user123" {
			t.Fatalf("roundtrip failed: got (%q, %q, %q)", enterprise, teamSlug, username)
		}
	})
}

func TestBuildEnterpriseTeamOrganizationsID(t *testing.T) {
	t.Run("builds correct ID format", func(t *testing.T) {
		got := buildEnterpriseTeamOrganizationsID("my-enterprise", "ent:my-team")
		want := "my-enterprise/ent:my-team"
		if got != want {
			t.Fatalf("buildEnterpriseTeamOrganizationsID() = %q, want %q", got, want)
		}
	})

	t.Run("handles empty strings", func(t *testing.T) {
		got := buildEnterpriseTeamOrganizationsID("", "")
		want := "/"
		if got != want {
			t.Fatalf("buildEnterpriseTeamOrganizationsID() = %q, want %q", got, want)
		}
	})
}

func TestParseEnterpriseTeamOrganizationsID(t *testing.T) {
	t.Run("parses valid ID", func(t *testing.T) {
		enterprise, teamSlug, err := parseEnterpriseTeamOrganizationsID("my-enterprise/ent:my-team")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if enterprise != "my-enterprise" {
			t.Fatalf("enterprise = %q, want %q", enterprise, "my-enterprise")
		}
		if teamSlug != "ent:my-team" {
			t.Fatalf("teamSlug = %q, want %q", teamSlug, "ent:my-team")
		}
	})

	t.Run("returns error for invalid format", func(t *testing.T) {
		_, _, err := parseEnterpriseTeamOrganizationsID("no-slash-here")
		if err == nil {
			t.Fatal("expected error for invalid ID format, got nil")
		}
	})

	t.Run("returns error for empty string", func(t *testing.T) {
		_, _, err := parseEnterpriseTeamOrganizationsID("")
		if err == nil {
			t.Fatal("expected error for empty ID, got nil")
		}
	})

	t.Run("roundtrips with build function", func(t *testing.T) {
		id := buildEnterpriseTeamOrganizationsID("enterprise", "ent:team-slug")
		enterprise, teamSlug, err := parseEnterpriseTeamOrganizationsID(id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if enterprise != "enterprise" || teamSlug != "ent:team-slug" {
			t.Fatalf("roundtrip failed: got (%q, %q)", enterprise, teamSlug)
		}
	})
}
