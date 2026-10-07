package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func enterpriseTeamTestOwner(t *testing.T, handler http.HandlerFunc) *Owner {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	return &Owner{v3client: mustCreateTestGitHubClient(t, server.URL+"/"), maxPerPage: 100}
}

func TestEnterpriseTeamReadRejectsReusedSlug(t *testing.T) {
	owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/enterprises/ent/teams/ent:old":
			fmt.Fprint(w, `{"id":99,"slug":"ent:old","name":"Other team"}`)
		case "/enterprises/ent/teams":
			fmt.Fprint(w, `[{"id":42,"slug":"ent:renamed","name":"Renamed"}]`)
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL)
			http.NotFound(w, r)
		}
	})
	d := schema.TestResourceDataRaw(t, resourceGithubEnterpriseTeam().Schema, map[string]any{
		"enterprise_slug": "ent", "name": "Old", "slug": "ent:old", "team_id": 42,
		"description": "Old description", "group_id": "old-group", "organization_selection_type": "all",
	})
	d.SetId("42")
	if diags := resourceGithubEnterpriseTeamRead(t.Context(), d, owner); diags.HasError() {
		t.Fatal(diags)
	}
	if d.Id() != "42" || d.Get("team_id") != 42 || d.Get("slug") != "ent:renamed" {
		t.Fatalf("adopted wrong team: id=%s team_id=%v slug=%v", d.Id(), d.Get("team_id"), d.Get("slug"))
	}
	if d.Get("description") != "" || d.Get("group_id") != "" || d.Get("organization_selection_type") != "disabled" {
		t.Fatal("read did not clear missing attributes or restore default selection")
	}
}

// Regression: with a stored numeric identity, Read must address the
// organization-assignment API directly by team ID — no GetTeam and no
// ListTeams ID->slug resolution — and must keep the resource ID stable even
// when the slug embedded in it is stale after an out-of-band rename.
func TestEnterpriseTeamOrganizationsReadFollowsNumericID(t *testing.T) {
	lookups := 0
	owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
		// The read must hit exactly this numeric-ID endpoint.
		if r.URL.Path == "/enterprises/ent/teams/42/organizations" && r.Method == http.MethodGet {
			fmt.Fprint(w, `[{"login":"org-a"}]`)
			return
		}
		// Any other request (team lookup, slug endpoint, ListTeams) is forbidden.
		lookups++
		t.Errorf("unexpected request: %s %s", r.Method, r.URL)
		http.NotFound(w, r)
	})
	d := schema.TestResourceDataRaw(t, resourceGithubEnterpriseTeamOrganizations().Schema, map[string]any{
		"enterprise_slug": "ent", "team_id": 42, "organization_slugs": []any{"org-a"},
	})
	d.SetId("ent/ent:old")
	if diags := resourceGithubEnterpriseTeamOrganizationsRead(t.Context(), d, owner); diags.HasError() {
		t.Fatal(diags)
	}
	if lookups != 0 {
		t.Fatalf("performed %d team lookups", lookups)
	}
	if d.Id() != "ent/ent:old" || d.Get("resolved_team_id") != 42 {
		t.Fatalf("lost assignment identity: %q, %v", d.Id(), d.Get("resolved_team_id"))
	}
}

func TestEnterpriseTeamUpdateClearsDescription(t *testing.T) {
	resource := resourceGithubEnterpriseTeam()
	old := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"enterprise_slug": "ent", "name": "Team", "slug": "ent:team", "team_id": 42,
		"description": "Remove this", "group_id": "keep-group",
	})
	old.SetId("42")
	config := terraform.NewResourceConfigRaw(map[string]any{
		"enterprise_slug": "ent", "name": "Team", "group_id": "keep-group",
	})
	diff, err := resource.Diff(t.Context(), old.State(), config, nil)
	if err != nil {
		t.Fatal(err)
	}
	d := resource.Data(old.State())
	// Apply through the SDK so HasChange sees the real old/new values.
	owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprint(w, `{"id":42,"slug":"ent:team","name":"Team"}`)
			return
		}
		var body map[string]json.RawMessage
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		if string(body["description"]) != `""` {
			t.Errorf("description not explicitly cleared: %s", body["description"])
		}
		if _, present := body["group_id"]; present {
			t.Error("unchanged IdP mapping should not be resubmitted")
		}
		fmt.Fprint(w, `{"id":42,"slug":"ent:team","name":"Team","description":"","group_id":"keep-group"}`)
	})
	_, diags := resource.Apply(t.Context(), d.State(), diff, owner)
	if diags.HasError() {
		t.Fatal(diags)
	}
}

func TestEnterpriseTeamInputValidation(t *testing.T) {
	for name, resource := range map[string]*schema.Resource{
		"membership":    resourceGithubEnterpriseTeamMembership(),
		"organizations": resourceGithubEnterpriseTeamOrganizations(),
	} {
		t.Run(name, func(t *testing.T) {
			for _, id := range []int{0, -1} {
				config := map[string]any{"enterprise_slug": "ent", "team_id": id}
				if name == "membership" {
					config["username"] = "user"
				} else {
					config["organization_slugs"] = []any{"org"}
				}
				if diags := resource.Validate(terraform.NewResourceConfigRaw(config)); !diags.HasError() {
					t.Errorf("accepted invalid team_id %d", id)
				}
			}
		})
	}
	for _, slug := range []string{"", "   "} {
		config := terraform.NewResourceConfigRaw(map[string]any{
			"enterprise_slug": "ent", "team_id": 42, "organization_slugs": []any{slug},
		})
		if diags := resourceGithubEnterpriseTeamOrganizations().Validate(config); !diags.HasError() {
			t.Errorf("accepted invalid organization slug %q", slug)
		}
	}
}

func TestEnterpriseTeamGroupRemovalPlansReplacement(t *testing.T) {
	resource := resourceGithubEnterpriseTeam()
	for _, tc := range []struct {
		name, old, next string
		replace         bool
	}{
		{"remove", "group-a", "", true},
		{"change", "group-a", "group-b", false},
		{"add", "", "group-a", false},
		{"unchanged", "group-a", "group-a", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			old := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
				"enterprise_slug": "ent", "name": "Team", "group_id": tc.old,
				"slug": "ent:team", "team_id": 42,
			})
			old.SetId("42")
			config := map[string]any{"enterprise_slug": "ent", "name": "Team"}
			if tc.next != "" {
				config["group_id"] = tc.next
			}
			diff, err := resource.Diff(t.Context(), old.State(), terraform.NewResourceConfigRaw(config), nil)
			if err != nil {
				t.Fatal(err)
			}
			if got := diff != nil && diff.RequiresNew(); got != tc.replace {
				t.Fatalf("replacement = %v, want %v", got, tc.replace)
			}
		})
	}
}

// Regression: dependent resources must use the stored numeric identity for
// Read/Delete without any team lookup (GetTeam or ListTeams), addressing the
// API directly at /teams/{id}/..., and must never rewrite the configured
// team_slug or the composite resource ID during those operations.
func TestEnterpriseTeamDependentsRetainIdentity(t *testing.T) {
	for name, resource := range map[string]*schema.Resource{
		"membership":    resourceGithubEnterpriseTeamMembership(),
		"organizations": resourceGithubEnterpriseTeamOrganizations(),
	} {
		for _, selector := range []string{"team_slug", "team_id"} {
			for _, operation := range []string{"read", "delete"} {
				t.Run(name+"/"+selector+"/"+operation, func(t *testing.T) {
					mutations := 0
					owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
						switch r.URL.Path {
						// All API calls must target the numeric team ID directly.
						case "/enterprises/ent/teams/42/memberships/user":
							if r.Method == http.MethodDelete {
								mutations++
								w.WriteHeader(http.StatusNoContent)
							} else {
								fmt.Fprint(w, `{"id":7,"login":"user"}`)
							}
						case "/enterprises/ent/teams/42/organizations":
							fmt.Fprint(w, `[{"login":"org-a"}]`)
						case "/enterprises/ent/teams/42/organizations/remove":
							mutations++
							fmt.Fprint(w, `[]`)
						default:
							// Any GetTeam/ListTeams call is a regression: the
							// stored numeric identity removes the need to resolve
							// slugs for these operations.
							t.Errorf("unexpected request: %s %s", r.Method, r.URL)
							http.NotFound(w, r)
						}
					})
					config := map[string]any{"enterprise_slug": "ent", "resolved_team_id": 42}
					if selector == "team_id" {
						config[selector] = 42
					} else {
						config[selector] = "ent:old"
					}
					id := "ent/ent:old"
					if name == "membership" {
						config["username"] = "user"
						id += "/user"
					} else {
						config["organization_slugs"] = []any{"org-a"}
					}
					d := schema.TestResourceDataRaw(t, resource.Schema, config)
					d.SetId(id)
					if operation == "read" {
						if diags := resource.ReadContext(t.Context(), d, owner); diags.HasError() {
							t.Fatal(diags)
						}
						if d.Id() != id || d.Get("resolved_team_id") != 42 {
							t.Fatalf("lost identity: %s, %v", d.Id(), d.Get("resolved_team_id"))
						}
						if selector == "team_slug" && d.Get("team_slug") != "ent:old" {
							t.Fatal("rewrote configured team_slug")
						}
						if selector == "team_id" && d.Get("team_slug") != "" {
							t.Fatal("populated conflicting selector")
						}
					} else {
						if diags := resource.DeleteContext(t.Context(), d, owner); diags.HasError() {
							t.Fatal(diags)
						}
						if mutations != 1 {
							t.Fatalf("expected one deletion, got %d", mutations)
						}
					}
				})
			}
		}
	}
}

// Regression: import and legacy state (no numeric identity yet) bootstrap the
// team ID with a single GetTeam on the ID slug — never a ListTeams scan — and
// keep the slug-based ID, resolved_team_id, and configured selector intact.
func TestEnterpriseTeamDependentsImportAndLegacyRefresh(t *testing.T) {
	for name, resource := range map[string]*schema.Resource{
		"membership":    resourceGithubEnterpriseTeamMembership(),
		"organizations": resourceGithubEnterpriseTeamOrganizations(),
	} {
		for _, legacy := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/legacy=%v", name, legacy), func(t *testing.T) {
				owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
					switch r.URL.Path {
					case "/enterprises/ent/teams/ent:team":
						fmt.Fprint(w, `{"id":42,"slug":"ent:team","name":"Team"}`)
					case "/enterprises/ent/teams/42/memberships/user":
						fmt.Fprint(w, `{"id":7,"login":"user"}`)
					case "/enterprises/ent/teams/42/organizations":
						fmt.Fprint(w, `[{"login":"org-a"}]`)
					default:
						t.Errorf("unexpected request: %s", r.URL)
						http.NotFound(w, r)
					}
				})
				config := map[string]any{}
				if legacy {
					config["enterprise_slug"] = "ent"
					config["team_slug"] = "ent:team"
				}
				d := schema.TestResourceDataRaw(t, resource.Schema, config)
				id := "ent/ent:team"
				if name == "membership" {
					id += "/user"
				}
				d.SetId(id)
				if !legacy {
					if _, err := resource.Importer.StateContext(t.Context(), d, owner); err != nil {
						t.Fatal(err)
					}
				}
				if diags := resource.ReadContext(t.Context(), d, owner); diags.HasError() {
					t.Fatal(diags)
				}
				// The composite ID keeps the slug selector; only resolved_team_id
				// records the numeric identity discovered by the single GetTeam.
				if d.Id() != id || d.Get("resolved_team_id") != 42 || d.Get("team_slug") != "ent:team" || d.Get("enterprise_slug") != "ent" {
					t.Fatalf("incomplete imported state: %#v", d.State().Attributes)
				}
			})
		}
	}
}

// Regression: a team_slug-configured assignment resource must survive a team
// rename in place — the old state, resource ID, and config selector ent:old
// transition to the renamed team's current slug ent:renamed without
// replacement, the update verifies the new slug resolves to the managed
// team's stable numeric ID before any mutation, and the assignment delta is
// applied against that ID rather than the slug (no ListTeams resolution).
func TestEnterpriseTeamOrganizationsTeamSlugRenameUpdatesInPlace(t *testing.T) {
	resource := resourceGithubEnterpriseTeamOrganizations()
	// State predates the rename: ID and configured slug still say ent:old,
	// resolved_team_id preserves the stable numeric identity.
	old := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"enterprise_slug": "ent", "team_slug": "ent:old", "resolved_team_id": 42,
		"organization_slugs": []any{"org-a", "org-b"},
	})
	old.SetId("ent/ent:old")
	config := terraform.NewResourceConfigRaw(map[string]any{
		"enterprise_slug": "ent", "team_slug": "ent:renamed", "organization_slugs": []any{"org-b", "org-c"},
	})
	diff, err := resource.Diff(t.Context(), old.State(), config, nil)
	if err != nil {
		t.Fatal(err)
	}
	if diff == nil || diff.Empty() {
		t.Fatal("expected nonempty diff for renamed team_slug and organization delta")
	}
	if diff.RequiresNew() {
		t.Fatal("team_slug change forced replacement; renamed team would be destroyed and recreation would fail")
	}
	mutations := 0
	owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/enterprises/ent/teams/ent:renamed":
			// The guard probes the configured (renamed) slug and must see the
			// managed team's numeric ID before any mutation.
			fmt.Fprint(w, `{"id":42,"slug":"ent:renamed"}`)
		case "/enterprises/ent/teams/42/organizations/add", "/enterprises/ent/teams/42/organizations/remove":
			var body struct {
				Slugs []string `json:"organization_slugs"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			want := "org-c"
			if r.URL.Path == "/enterprises/ent/teams/42/organizations/remove" {
				want = "org-a"
			}
			if len(body.Slugs) != 1 || body.Slugs[0] != want {
				t.Errorf("incorrect delta: %v, want %s", body.Slugs, want)
			}
			mutations++
			fmt.Fprint(w, `[]`)
		default:
			t.Errorf("unexpected request (no ListTeams fallback allowed): %s %s", r.Method, r.URL)
			http.NotFound(w, r)
		}
	})
	state, diags := resource.Apply(t.Context(), old.State(), diff, owner)
	if diags.HasError() {
		t.Fatal(diags)
	}
	if mutations != 2 {
		t.Fatalf("got %d mutations, want 2", mutations)
	}
	if state.ID != "ent/ent:renamed" {
		t.Fatalf("update did not adopt the renamed slug in the resource ID: %s", state.ID)
	}
	if got := state.Attributes["team_slug"]; got != "ent:renamed" {
		t.Fatalf("configured team_slug not reflected in state: %q", got)
	}
	if got := state.Attributes["resolved_team_id"]; got != "42" {
		t.Fatalf("lost numeric identity: %q", got)
	}
}

// Regression: a configured team_slug that no longer exists must fail the
// update before any add/remove mutation instead of being treated as a rename
// hint, and the failure must never fall back to a ListTeams ID->slug scan.
// Both a stale unchanged slug alongside an org delta and a config transition
// to the stale slug must be rejected the same way.
func TestEnterpriseTeamOrganizationsRejectsNonexistentSlug(t *testing.T) {
	resource := resourceGithubEnterpriseTeamOrganizations()
	for name, tc := range map[string]struct {
		oldSlug, nextSlug string
	}{
		"stale slug with org delta":      {oldSlug: "ent:gone", nextSlug: "ent:gone"},
		"stale slug after config rename": {oldSlug: "ent:old", nextSlug: "ent:gone"},
	} {
		t.Run(name, func(t *testing.T) {
			old := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
				"enterprise_slug": "ent", "team_slug": tc.oldSlug, "resolved_team_id": 42,
				"organization_slugs": []any{"org-a"},
			})
			old.SetId("ent/" + tc.oldSlug)
			config := terraform.NewResourceConfigRaw(map[string]any{
				"enterprise_slug": "ent", "team_slug": tc.nextSlug, "organization_slugs": []any{"org-b"},
			})
			diff, err := resource.Diff(t.Context(), old.State(), config, nil)
			if err != nil {
				t.Fatal(err)
			}
			mutations := 0
			owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					mutations++
					t.Errorf("mutated on rejected slug: %s %s", r.Method, r.URL)
				}
				switch r.URL.Path {
				case "/enterprises/ent/teams/ent:gone":
					http.NotFound(w, r)
				default:
					t.Errorf("unexpected request (no ListTeams fallback allowed): %s %s", r.Method, r.URL)
					http.NotFound(w, r)
				}
			})
			state, diags := resource.Apply(t.Context(), old.State(), diff, owner)
			if !diags.HasError() {
				t.Fatal("expected error for nonexistent configured team_slug")
			}
			if mutations != 0 {
				t.Fatalf("got %d mutations, want 0", mutations)
			}
			if state == nil {
				t.Fatal("failed update lost resource state")
			}
			if state.ID != old.Id() || state.Attributes["resolved_team_id"] != "42" {
				t.Fatalf("failed update changed identity: %#v", state)
			}
		})
	}
}

// Regression: changing team_slug to a different existing team must be rejected
// with an error diagnostic and zero add/remove mutations, instead of silently
// applying the delta to the previously-resolved team.
func TestEnterpriseTeamOrganizationsRejectTeamSwitch(t *testing.T) {
	mutations := 0
	owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			mutations++
			t.Errorf("mutated team on rejected switch: %s %s", r.Method, r.URL)
		}
		switch r.URL.Path {
		case "/enterprises/ent/teams/ent:a":
			fmt.Fprint(w, `{"id":42,"slug":"ent:a"}`)
		case "/enterprises/ent/teams/ent:b":
			fmt.Fprint(w, `{"id":99,"slug":"ent:b"}`)
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL)
			http.NotFound(w, r)
		}
	})
	d := schema.TestResourceDataRaw(t, resourceGithubEnterpriseTeamOrganizations().Schema, map[string]any{
		"enterprise_slug": "ent", "team_slug": "ent:b", "resolved_team_id": 42,
		"organization_slugs": []any{"org-a"},
	})
	d.SetId("ent/ent:a")
	diags := resourceGithubEnterpriseTeamOrganizationsUpdate(t.Context(), d, owner)
	if !diags.HasError() {
		t.Fatal("expected error when team_slug switches to a different existing team")
	}
	if mutations != 0 {
		t.Fatalf("got %d mutations, want 0", mutations)
	}
}

// Regression: any failure to resolve the configured team_slug during an
// update must surface as an error diagnostic with zero add/remove mutations
// instead of silently proceeding against the managed team. Since the update
// now mutates against the stored numeric identity, a 404 probe result is an
// invalid configured slug — not a stale rename to recover from — and is
// rejected too (covered by TestEnterpriseTeamOrganizationsRejectsNonexistentSlug);
// this test asserts the non-404 failures (403, 5xx) surface verbatim, with the
// diagnostics preserving the underlying HTTP status.
func TestEnterpriseTeamOrganizationsProbeErrorsAreFatal(t *testing.T) {
	for _, code := range []int{http.StatusForbidden, http.StatusInternalServerError} {
		t.Run(fmt.Sprintf("%d", code), func(t *testing.T) {
			mutations := 0
			owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					mutations++
					t.Errorf("mutated team after probe failure: %s %s", r.Method, r.URL)
					return
				}
				switch r.URL.Path {
				case "/enterprises/ent/teams/ent:a":
					fmt.Fprint(w, `{"id":42,"slug":"ent:a"}`)
				case "/enterprises/ent/teams/ent:b":
					w.WriteHeader(code)
					fmt.Fprint(w, `{"message":"error"}`)
				default:
					t.Errorf("unexpected request: %s %s", r.Method, r.URL)
					http.NotFound(w, r)
				}
			})
			d := schema.TestResourceDataRaw(t, resourceGithubEnterpriseTeamOrganizations().Schema, map[string]any{
				"enterprise_slug": "ent", "team_slug": "ent:b", "resolved_team_id": 42,
				"organization_slugs": []any{"org-a"},
			})
			d.SetId("ent/ent:a")
			diags := resourceGithubEnterpriseTeamOrganizationsUpdate(t.Context(), d, owner)
			if !diags.HasError() {
				t.Fatalf("expected error diagnostics for probe HTTP %d, got %v", code, diags)
			}
			// The raw probe failure must not be swallowed, mislabeled, or
			// rewritten: the diagnostics must preserve the underlying HTTP
			// status and must not claim the team does not exist.
			preserved, mislabeled := false, false
			for _, pdiag := range diags {
				if strings.Contains(pdiag.Summary, fmt.Sprintf("%d", code)) {
					preserved = true
				}
				if strings.Contains(pdiag.Summary, "does not exist in enterprise") {
					mislabeled = true
				}
			}
			if !preserved {
				t.Fatalf("probe error diagnostics lost HTTP %d: %v", code, diags)
			}
			if mislabeled {
				t.Fatalf("probe error diagnostics mislabel HTTP %d as nonexistent: %v", code, diags)
			}
			if mutations != 0 {
				t.Fatalf("got %d mutations, want 0", mutations)
			}
		})
	}
}

// Regression: a team_id-configured assignment resource must apply its delta
// against the stored numeric identity with no team lookup at all (the ID's
// slug may be stale after a rename) and must keep the stored resource ID.
func TestEnterpriseTeamOrganizationsUpdateDeltaAfterRename(t *testing.T) {
	resource := resourceGithubEnterpriseTeamOrganizations()
	old := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"enterprise_slug": "ent", "team_id": 42, "resolved_team_id": 42, "organization_slugs": []any{"org-a", "org-b"},
	})
	old.SetId("ent/ent:old")
	config := terraform.NewResourceConfigRaw(map[string]any{
		"enterprise_slug": "ent", "team_id": 42, "organization_slugs": []any{"org-b", "org-c"},
	})
	diff, err := resource.Diff(t.Context(), old.State(), config, nil)
	if err != nil {
		t.Fatal(err)
	}
	mutations := 0
	owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/enterprises/ent/teams/42/organizations/add", "/enterprises/ent/teams/42/organizations/remove":
			var body struct {
				Slugs []string `json:"organization_slugs"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			want := "org-c"
			if r.URL.Path == "/enterprises/ent/teams/42/organizations/remove" {
				want = "org-a"
			}
			if len(body.Slugs) != 1 || body.Slugs[0] != want {
				t.Errorf("incorrect delta: %v, want %s", body.Slugs, want)
			}
			mutations++
			fmt.Fprint(w, `[]`)
		default:
			t.Errorf("unexpected request (no team lookups allowed): %s", r.URL)
			http.NotFound(w, r)
		}
	})
	state, diags := resource.Apply(t.Context(), old.State(), diff, owner)
	if diags.HasError() {
		t.Fatal(diags)
	}
	if mutations != 2 {
		t.Fatalf("got %d mutations, want 2", mutations)
	}
	if state.ID != "ent/ent:old" {
		t.Fatalf("update must keep the stored ID: %s", state.ID)
	}
}

func TestEnterpriseTeamMissingAndAPIError(t *testing.T) {
	for name, resource := range map[string]*schema.Resource{
		"team":          resourceGithubEnterpriseTeam(),
		"membership":    resourceGithubEnterpriseTeamMembership(),
		"organizations": resourceGithubEnterpriseTeamOrganizations(),
	} {
		for _, code := range []int{http.StatusNotFound, http.StatusForbidden, http.StatusInternalServerError} {
			t.Run(fmt.Sprintf("%s/%d", name, code), func(t *testing.T) {
				writes := 0
				owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
					if r.Method != http.MethodGet {
						writes++
					}
					if code == http.StatusNotFound && r.URL.Path == "/enterprises/ent/teams" {
						fmt.Fprint(w, `[]`)
						return
					}
					w.WriteHeader(code)
					fmt.Fprint(w, `{"message":"error"}`)
				})
				config := map[string]any{"enterprise_slug": "ent", "team_id": 42}
				id := "42"
				if name == "team" {
					config["slug"] = "ent:old"
					config["name"] = "Old"
				} else {
					config["resolved_team_id"] = 42
					id = "ent/ent:old"
					if name == "membership" {
						config["username"] = "user"
						id += "/user"
					} else {
						config["organization_slugs"] = []any{"org-a"}
					}
				}
				d := schema.TestResourceDataRaw(t, resource.Schema, config)
				d.SetId(id)
				diags := resource.ReadContext(t.Context(), d, owner)
				if code == http.StatusNotFound {
					if diags.HasError() || d.Id() != "" {
						t.Fatalf("missing resource retained: %s %v", d.Id(), diags)
					}
				} else if !diags.HasError() || d.Id() != id {
					t.Fatalf("API error lost or state cleared: %s %v", d.Id(), diags)
				}
				d.SetId(id)
				diags = resource.DeleteContext(t.Context(), d, owner)
				if diags.HasError() != (code != http.StatusNotFound) {
					t.Fatalf("wrong delete error: %v", diags)
				}
				// Membership delete addresses the stored numeric identity
				// directly — the DELETE itself verifies reachability. Team and
				// organizations deletes must fail before any mutation.
				expectedWrites := 0
				if name == "membership" {
					expectedWrites = 1
				}
				if writes != expectedWrites {
					t.Fatalf("got %d writes, want %d", writes, expectedWrites)
				}
			})
		}
	}
}

func TestEnterpriseTeamMutationsRejectReusedSlug(t *testing.T) {
	for _, operation := range []string{"update", "delete"} {
		t.Run(operation, func(t *testing.T) {
			mutations := 0
			owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/enterprises/ent/teams/ent:old":
					if r.Method != http.MethodGet {
						t.Errorf("mutated unrelated team: %s", r.Method)
					}
					fmt.Fprint(w, `{"id":99,"slug":"ent:old"}`)
				case "/enterprises/ent/teams":
					if r.URL.Query().Get("page") == "2" {
						fmt.Fprint(w, `[{"id":42,"slug":"ent:renamed"}]`)
						return
					}
					w.Header().Set("Link", fmt.Sprintf("<http://%s/enterprises/ent/teams?page=2>; rel=\"next\"", r.Host))
					fmt.Fprint(w, `[{"id":99,"slug":"ent:old"}]`)
				case "/enterprises/ent/teams/ent:renamed":
					mutations++
					if operation == "delete" {
						w.WriteHeader(http.StatusNoContent)
					} else {
						fmt.Fprint(w, `{"id":42,"slug":"ent:renamed"}`)
					}
				default:
					t.Errorf("unexpected request: %s", r.URL)
					http.NotFound(w, r)
				}
			})
			r := resourceGithubEnterpriseTeam()
			d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
				"enterprise_slug": "ent", "name": "Renamed", "slug": "ent:old", "team_id": 42,
			})
			d.SetId("42")
			if operation == "delete" {
				if diags := r.DeleteContext(t.Context(), d, owner); diags.HasError() {
					t.Fatal(diags)
				}
			} else {
				if diags := r.UpdateContext(t.Context(), d, owner); diags.HasError() {
					t.Fatal(diags)
				}
			}
			if mutations != 1 {
				t.Fatalf("got %d writes to intended team, want 1", mutations)
			}
		})
	}
}

// Regression: creates must persist the numeric identity, and create-time team
// resolution must be minimal: the team_id selector needs no team lookup at all
// (the dependent endpoints accept the numeric ID), while the team_slug
// selector resolves with exactly one direct GetTeam — never a ListTeams scan.
func TestEnterpriseTeamDependentsCreateStoresIdentity(t *testing.T) {
	for name, resource := range map[string]*schema.Resource{
		"membership":    resourceGithubEnterpriseTeamMembership(),
		"organizations": resourceGithubEnterpriseTeamOrganizations(),
	} {
		for _, selector := range []string{"team_slug", "team_id"} {
			t.Run(name+"/"+selector, func(t *testing.T) {
				getTeams := 0
				owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path == "/enterprises/ent/teams/ent:team" {
						getTeams++
						fmt.Fprint(w, `{"id":42,"slug":"ent:team"}`)
						return
					}
					// Dependent mutations must use the exact numeric endpoint;
					// anything else (including ListTeams) is unexpected.
					switch r.URL.Path {
					case "/enterprises/ent/teams/42/memberships/user":
						fmt.Fprint(w, `{"id":7,"login":"user"}`)
					case "/enterprises/ent/teams/42/organizations", "/enterprises/ent/teams/42/organizations/add":
						fmt.Fprint(w, `[]`)
					default:
						t.Errorf("unexpected request: %s %s", r.Method, r.URL)
						http.NotFound(w, r)
					}
				})
				config := map[string]any{"enterprise_slug": "ent"}
				if selector == "team_id" {
					config[selector] = 42
				} else {
					config[selector] = "ent:team"
				}
				if name == "membership" {
					config["username"] = "user"
				} else {
					config["organization_slugs"] = []any{"org-a"}
				}
				d := schema.TestResourceDataRaw(t, resource.Schema, config)
				if diags := resource.CreateContext(t.Context(), d, owner); diags.HasError() {
					t.Fatal(diags)
				}
				if selector == "team_slug" && getTeams != 1 {
					t.Fatalf("team_slug create must use exactly one GetTeam, got %d", getTeams)
				}
				if selector == "team_id" && getTeams != 0 {
					t.Fatalf("team_id create must not look the team up, got %d GetTeam calls", getTeams)
				}
				wantID := "ent/ent:team"
				if selector == "team_id" {
					wantID = "ent/42"
				}
				if name == "membership" {
					wantID += "/user"
				}
				if d.Id() != wantID || d.Get("resolved_team_id") != 42 {
					t.Fatalf("missing persisted identity: id=%q resolved=%v", d.Id(), d.Get("resolved_team_id"))
				}
			})
		}
	}
}

func TestEnterpriseTeamOrganizationsCaseInsensitive(t *testing.T) {
	for _, tc := range []struct {
		name          string
		before, after []any
		writes        int
	}{
		{"case only", []any{"org-a"}, []any{"Org-A"}, 0},
		{"duplicates", []any{"org-a"}, []any{"org-a", "Org-A"}, 0},
		{"legacy casing", []any{"Org-A"}, []any{"org-a"}, 0},
		{"real delta", []any{"Org-A", "org-b"}, []any{"org-a", "Org-C"}, 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := resourceGithubEnterpriseTeamOrganizations()
			legacy := resourceGithubEnterpriseTeamOrganizations()
			legacy.Schema["organization_slugs"].Set = schema.HashString
			element, _ := legacy.Schema["organization_slugs"].Elem.(*schema.Schema)
			element.StateFunc = nil
			old := schema.TestResourceDataRaw(t, legacy.Schema, map[string]any{
				"enterprise_slug": "ent", "team_id": 42, "resolved_team_id": 42, "organization_slugs": tc.before,
			})
			old.SetId("ent/ent:team")
			config := terraform.NewResourceConfigRaw(map[string]any{
				"enterprise_slug": "ent", "team_id": 42, "organization_slugs": tc.after,
			})
			diff, err := r.Diff(t.Context(), old.State(), config, nil)
			if err != nil {
				t.Fatal(err)
			}
			writes := 0
			owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, req *http.Request) {
				switch req.URL.Path {
				case "/enterprises/ent/teams/42/organizations":
					if tc.writes == 2 {
						fmt.Fprint(w, `[{"login":"Org-A"},{"login":"Org-C"}]`)
					} else {
						fmt.Fprint(w, `[{"login":"Org-A"}]`)
					}
				case "/enterprises/ent/teams/42/organizations/add", "/enterprises/ent/teams/42/organizations/remove":
					writes++
					var body struct {
						Slugs []string `json:"organization_slugs"`
					}
					if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
						t.Error(err)
					}
					want := "org-c"
					if req.URL.Path == "/enterprises/ent/teams/42/organizations/remove" {
						want = "org-b"
					}
					if tc.writes == 0 || len(body.Slugs) != 1 || body.Slugs[0] != want {
						t.Errorf("unexpected mutation: %s %v", req.URL.Path, body.Slugs)
					}
					w.WriteHeader(http.StatusNoContent)
				default:
					t.Errorf("unexpected request %s", req.URL)
					http.NotFound(w, req)
				}
			})
			state := old.State()
			if diff != nil && !diff.Empty() {
				var diags diag.Diagnostics
				state, diags = r.Apply(t.Context(), state, diff, owner)
				if diags.HasError() {
					t.Fatal(diags)
				}
			}
			if writes != tc.writes {
				t.Fatalf("writes = %d, want %d", writes, tc.writes)
			}
			refreshed := r.Data(state)
			if diags := r.ReadContext(t.Context(), refreshed, owner); diags.HasError() {
				t.Fatal(diags)
			}
			diff, err = r.Diff(t.Context(), refreshed.State(), config, owner)
			if err != nil {
				t.Fatal(err)
			}
			if diff != nil && !diff.Empty() {
				t.Fatalf("nonempty plan after refresh: %#v", diff.Attributes)
			}
		})
	}
}

func TestEnterpriseTeamDependentImportSelectors(t *testing.T) {
	for name, r := range map[string]*schema.Resource{
		"membership": resourceGithubEnterpriseTeamMembership(), "organizations": resourceGithubEnterpriseTeamOrganizations(),
	} {
		for _, numeric := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/numeric=%v", name, numeric), func(t *testing.T) {
				owner := enterpriseTeamTestOwner(t, func(w http.ResponseWriter, req *http.Request) {
					switch req.URL.Path {
					case "/enterprises/ent/teams/ent:team":
						fmt.Fprint(w, `{"id":42,"slug":"ent:team"}`)
					case "/enterprises/ent/teams/42/organizations":
						fmt.Fprint(w, `[{"login":"org-a"}]`)
					case "/enterprises/ent/teams/42/memberships/user":
						fmt.Fprint(w, `{"id":7,"login":"user"}`)
					default:
						// Numeric imports must not list teams to resolve a slug;
						// slug imports bootstrap with a single GetTeam.
						t.Errorf("unexpected request: %s", req.URL)
						http.NotFound(w, req)
					}
				})
				selector := "ent:team"
				config := map[string]any{"enterprise_slug": "ent"}
				if numeric {
					selector = "42"
					config["team_id"] = 42
				} else {
					config["team_slug"] = selector
				}
				id := "ent/" + selector
				// The composite ID preserves the imported selector.
				expectedID := id
				if name == "membership" {
					id += "/user"
					expectedID = id
					config["username"] = "user"
				} else {
					config["organization_slugs"] = []any{"org-a"}
				}
				d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{})
				d.SetId(id)
				if _, err := r.Importer.StateContext(t.Context(), d, owner); err != nil {
					t.Fatal(err)
				}
				if diags := r.ReadContext(t.Context(), d, owner); diags.HasError() {
					t.Fatal(diags)
				}
				if d.Id() != expectedID || d.Get("resolved_team_id") != 42 {
					t.Fatalf("incorrect imported state: %#v", d.State().Attributes)
				}
				diff, err := r.Diff(t.Context(), d.State(), terraform.NewResourceConfigRaw(config), owner)
				if err != nil {
					t.Fatal(err)
				}
				if diff != nil && !diff.Empty() {
					t.Fatalf("import requires changes: %#v", diff.Attributes)
				}
			})
		}
		for _, selector := range []string{"0", "-1", "999999999999999999999999", ""} {
			t.Run(name+"/invalid/"+selector, func(t *testing.T) {
				d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{})
				id := "ent/" + selector
				if name == "membership" {
					id += "/user"
				}
				d.SetId(id)
				if _, err := r.Importer.StateContext(t.Context(), d, nil); err == nil {
					t.Fatalf("accepted invalid import %q", id)
				}
			})
		}
	}
}
