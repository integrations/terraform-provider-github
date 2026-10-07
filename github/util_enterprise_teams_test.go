package github

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	for _, id := range []string{"only-one-part", "ent/team/user/extra", "/team/user", "ent//user", "ent/team/", " /team/user", "ent/ /user", "ent/team/ "} {
		t.Run(id, func(t *testing.T) {
			if _, _, _, err := parseEnterpriseTeamMembershipID(id); err == nil {
				t.Fatal("accepted invalid membership ID")
			}
		})
	}
}

func TestEnterpriseTeamImportValidation(t *testing.T) {
	r := resourceGithubEnterpriseTeam()
	for _, id := range []string{"ent", "ent/42/extra", "/42", " /42", "ent/", "ent/0", "ent/-1", "ent/not-a-number", "ent/999999999999999999999999"} {
		t.Run(id, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{})
			d.SetId(id)
			if _, err := r.Importer.StateContext(t.Context(), d, nil); err == nil {
				t.Fatal("accepted invalid team import")
			}
			if d.Id() != id || d.Get("enterprise_slug") != "" {
				t.Fatal("invalid import modified state")
			}
		})
	}
	t.Run("valid", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{})
		d.SetId("ent/42")
		if _, err := r.Importer.StateContext(t.Context(), d, nil); err != nil {
			t.Fatal(err)
		}
		if d.Id() != "42" || d.Get("enterprise_slug") != "ent" {
			t.Fatalf("incorrect imported identity: %q %v", d.Id(), d.Get("enterprise_slug"))
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

func TestFindEnterpriseTeamByIDPagination(t *testing.T) {
	for _, tc := range []struct {
		name      string
		foundPage int
		fail      bool
	}{
		{name: "first page", foundPage: 1},
		{name: "second page", foundPage: 2},
		{name: "not found"},
		{name: "second page error", fail: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
				calls++
				if r.Method != http.MethodGet || r.URL.Path != "/enterprises/ent/teams" || r.URL.Query().Get("per_page") != "100" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL)
				}
				page := 1
				if r.URL.Query().Get("page") == "2" {
					page = 2
					if tc.fail {
						http.Error(w, "listing unavailable", http.StatusInternalServerError)
						return
					}
				} else {
					w.Header().Set("Link", fmt.Sprintf("<http://%s/enterprises/ent/teams?page=2>; rel=\"next\"", r.Host))
				}
				id := 99
				if page == tc.foundPage {
					id = 42
				}
				fmt.Fprintf(w, `[{"id":%d,"slug":"ent:team"}]`, id)
			})
			team, err := findEnterpriseTeamByID(owner, t.Context(), "ent", 42)
			if tc.fail {
				ghErr, ok := errors.AsType[*github.ErrorResponse](err)
				if !ok || ghErr.Response == nil || ghErr.Response.StatusCode != http.StatusInternalServerError || team != nil {
					t.Fatalf("expected listing error and no team, got %v, %v", team, err)
				}
			} else if err != nil {
				t.Fatal(err)
			} else if tc.foundPage == 0 {
				if team != nil {
					t.Fatalf("found an unrelated team: %v", team)
				}
			} else if team == nil || team.ID != 42 {
				t.Fatalf("lost matching team: %v", team)
			}
			wantCalls := 2
			if tc.foundPage == 1 {
				wantCalls = 1
			}
			if calls != wantCalls {
				t.Fatalf("listing calls = %d, want %d", calls, wantCalls)
			}
		})
	}
}

func TestListAllEnterpriseTeamOrganizationsPagination(t *testing.T) {
	for _, fail := range []bool{false, true} {
		t.Run(fmt.Sprintf("error=%v", fail), func(t *testing.T) {
			calls := 0
			owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
				calls++
				if r.Method != http.MethodGet || r.URL.Path != "/enterprises/ent/teams/42/organizations" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL)
				}
				if r.URL.Query().Get("page") == "2" {
					if fail {
						http.Error(w, "listing unavailable", http.StatusInternalServerError)
						return
					}
					fmt.Fprint(w, `[{"login":"org-b"}]`)
					return
				}
				w.Header().Set("Link", fmt.Sprintf("<http://%s/enterprises/ent/teams/42/organizations?page=2>; rel=\"next\"", r.Host))
				fmt.Fprint(w, `[{"login":"org-a"}]`)
			})
			orgs, err := listAllEnterpriseTeamOrganizations(owner, t.Context(), "ent", "42")
			if fail {
				if err == nil || orgs != nil {
					t.Fatalf("returned partial assignments: %v, %v", orgs, err)
				}
			} else if err != nil || !slices.Equal(organizationSlugs(orgs), []string{"org-a", "org-b"}) {
				t.Fatalf("incomplete assignments: %v, %v", orgs, err)
			}
			if calls != 2 {
				t.Fatalf("listing calls = %d, want 2", calls)
			}
		})
	}
}
