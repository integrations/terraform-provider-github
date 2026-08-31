package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"sync"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

const sweepTestOwner = "testowner"

func reposJSON(names ...string) []byte {
	type repoJSON struct {
		Name    string `json:"name"`
		HTMLURL string `json:"html_url"`
	}
	repos := make([]repoJSON, len(names))
	for i, n := range names {
		repos[i] = repoJSON{Name: n, HTMLURL: "https://github.com/" + sweepTestOwner + "/" + n}
	}
	b, _ := json.Marshal(repos)
	return b
}

func teamsJSON(slugs ...string) []byte {
	type teamJSON struct {
		Slug    string `json:"slug"`
		HTMLURL string `json:"html_url"`
	}
	teams := make([]teamJSON, len(slugs))
	for i, s := range slugs {
		teams[i] = teamJSON{Slug: s, HTMLURL: "https://github.com/orgs/" + sweepTestOwner + "/teams/" + s}
	}
	b, _ := json.Marshal(teams)
	return b
}

// writeJSON writes a JSON payload to the response with Content-Type set.
func writeJSON(w http.ResponseWriter, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload) //nolint:errcheck
}

// TestOrphansListsOnlyPrefixed verifies that only tf-acc-test-* repos are returned.
// No DELETE handler is registered - a DELETE 404s, which would propagate as an error
// and cause the test to fail, proving Orphans is read-only (invariant #3).
func TestOrphansListsOnlyPrefixed(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, reposJSON("tf-acc-test-a", "real-repo", "tf-acc-test-b"))
	})
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, teamsJSON())
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	resources, err := orphansWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true)
	if err != nil {
		t.Fatalf("orphansWithClient error: %v", err)
	}
	var repos []provider.Resource
	for _, r := range resources {
		if r.Kind == "repository" {
			repos = append(repos, r)
		}
	}
	if len(repos) != 2 {
		t.Fatalf("expected 2 prefixed repos, got %d: %v", len(repos), repos)
	}
	for _, r := range repos {
		if r.Name != "tf-acc-test-a" && r.Name != "tf-acc-test-b" {
			t.Errorf("unexpected repo returned: %q", r.Name)
		}
		if r.URL == "" {
			t.Errorf("repo %q has empty URL", r.Name)
		}
	}
}

// TestOrphansPaginates verifies that all prefixed repos across 2 pages are returned.
func TestOrphansPaginates(t *testing.T) {
	var mu sync.Mutex
	pagesHit := map[string]int{}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, r *http.Request) {
		pg := r.URL.Query().Get("page")
		mu.Lock()
		pagesHit[pg]++
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		if pg == "2" {
			w.Write(reposJSON("tf-acc-test-page2")) //nolint:errcheck
		} else {
			linkHeader := fmt.Sprintf(`<%s/orgs/%s/repos?page=2>; rel="next"`,
				"http://"+r.Host, sweepTestOwner)
			w.Header().Set("Link", linkHeader)
			w.Write(reposJSON("tf-acc-test-page1")) //nolint:errcheck
		}
	})
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, teamsJSON())
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	resources, err := orphansWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true)
	if err != nil {
		t.Fatalf("orphansWithClient error: %v", err)
	}

	names := map[string]bool{}
	for _, r := range resources {
		if r.Kind == "repository" {
			names[r.Name] = true
		}
	}
	if !names["tf-acc-test-page1"] {
		t.Error("missing tf-acc-test-page1 from page 1")
	}
	if !names["tf-acc-test-page2"] {
		t.Error("missing tf-acc-test-page2 from page 2")
	}
	if len(names) != 2 {
		t.Errorf("expected 2 repos, got %d: %v", len(names), names)
	}
}

// TestOrphansOrgIncludesTeams verifies that org mode returns prefixed teams as Kind "team".
func TestOrphansOrgIncludesTeams(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, reposJSON())
	})
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, teamsJSON("tf-acc-test-team", "real-team"))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	resources, err := orphansWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true)
	if err != nil {
		t.Fatalf("orphansWithClient error: %v", err)
	}
	var teams []provider.Resource
	for _, r := range resources {
		if r.Kind == "team" {
			teams = append(teams, r)
		}
	}
	if len(teams) != 1 {
		t.Fatalf("expected 1 prefixed team, got %d: %v", len(teams), teams)
	}
	if teams[0].Name != "tf-acc-test-team" {
		t.Errorf("team name = %q, want %q", teams[0].Name, "tf-acc-test-team")
	}
	if teams[0].URL == "" {
		t.Error("team URL should not be empty")
	}
}

// TestOrphansUserModeNoTeams verifies that org=false never calls the teams endpoint (invariant #3).
func TestOrphansUserModeNoTeams(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, reposJSON("tf-acc-test-user-repo"))
	})
	mux.HandleFunc("/orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("teams endpoint must not be called in user mode; got %s %s", r.Method, r.URL.Path)
		http.Error(w, "unexpected teams call", http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	resources, err := orphansWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, false)
	if err != nil {
		t.Fatalf("orphansWithClient error: %v", err)
	}
	for _, r := range resources {
		if r.Kind == "team" {
			t.Errorf("user mode must not return team resources; got: %v", r)
		}
	}
}

// TestSweepRefusesWithoutConfirm is invariant #2: Confirm=false returns an error and
// the test server records zero requests.
func TestSweepRefusesWithoutConfirm(t *testing.T) {
	var mu sync.Mutex
	var hitPaths []string

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		hitPaths = append(hitPaths, r.Method+" "+r.URL.Path)
		mu.Unlock()
		http.Error(w, "unexpected request", http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true,
		provider.SweepOpts{Targets: []string{"repositories", "teams"}, Confirm: false})
	if err == nil {
		t.Fatal("expected an error when Confirm=false, got nil")
	}

	mu.Lock()
	n := len(hitPaths)
	captured := append([]string(nil), hitPaths...)
	mu.Unlock()
	if n != 0 {
		t.Errorf("expected zero HTTP requests when Confirm=false, got %d: %v", n, captured)
	}
}

// TestSweepDeletesOnlyPrefixed is invariant #1: only tf-acc-test-* repos are deleted.
func TestSweepDeletesOnlyPrefixed(t *testing.T) {
	var mu sync.Mutex
	var deletedRepos []string

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, reposJSON("tf-acc-test-a", "production-database", "tf-acc-test-b"))
	})
	mux.HandleFunc("DELETE /repos/"+sweepTestOwner+"/{name}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		deletedRepos = append(deletedRepos, r.PathValue("name"))
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true,
		provider.SweepOpts{Targets: []string{"repositories"}, Confirm: true})
	if err != nil {
		t.Fatalf("sweepWithClient error: %v", err)
	}

	mu.Lock()
	got := append([]string(nil), deletedRepos...)
	mu.Unlock()
	sort.Strings(got)

	want := []string{"tf-acc-test-a", "tf-acc-test-b"}
	if len(got) != len(want) {
		t.Fatalf("deleted repos = %v, want %v", got, want)
	}
	for i, g := range got {
		if g != want[i] {
			t.Errorf("deleted[%d] = %q, want %q", i, g, want[i])
		}
	}
}

func TestSweepExplicitModeIgnoresAmbientEnvironment(t *testing.T) {
	t.Setenv("GH_TEST_AUTH_MODE", "organization")

	var mu sync.Mutex
	var deletedRepos []string

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, reposJSON("tf-acc-test-user-repo"))
	})
	mux.HandleFunc("DELETE /repos/"+sweepTestOwner+"/{name}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		deletedRepos = append(deletedRepos, r.PathValue("name"))
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("/orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("organization repos endpoint must not be called for explicit individual mode: %s %s", r.Method, r.URL.Path)
		http.Error(w, "unexpected organization repos call", http.StatusInternalServerError)
	})
	mux.HandleFunc("/orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("teams endpoint must not be called for explicit individual mode: %s %s", r.Method, r.URL.Path)
		http.Error(w, "unexpected teams call", http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, isOrgMode("individual"), provider.SweepOpts{
		Targets: []string{"repositories", "teams"},
		Confirm: true,
	})
	if err != nil {
		t.Fatalf("sweepWithClient error: %v", err)
	}

	mu.Lock()
	got := append([]string(nil), deletedRepos...)
	mu.Unlock()
	if len(got) != 1 || got[0] != "tf-acc-test-user-repo" {
		t.Fatalf("deleted repos = %v, want [tf-acc-test-user-repo]", got)
	}
}

// TestSweepCollectsBeforeDeletingRepos proves sweep does not delete while
// walking page-numbered results. GitHub list pagination can shift after a
// delete; collecting all names first prevents page 2 from skipping forward.
func TestSweepCollectsBeforeDeletingRepos(t *testing.T) {
	var mu sync.Mutex
	var deletedRepos []string

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		deletedBeforeList := len(deletedRepos)
		mu.Unlock()
		if deletedBeforeList != 0 {
			t.Errorf("sweep deleted %d repo(s) before finishing list pagination", deletedBeforeList)
		}
		if r.URL.Query().Get("page") == "2" {
			writeJSON(w, reposJSON("tf-acc-test-page2"))
			return
		}
		linkHeader := fmt.Sprintf(`<%s/orgs/%s/repos?page=2>; rel="next"`,
			"http://"+r.Host, sweepTestOwner)
		w.Header().Set("Link", linkHeader)
		writeJSON(w, reposJSON("tf-acc-test-page1"))
	})
	mux.HandleFunc("DELETE /repos/"+sweepTestOwner+"/{name}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		deletedRepos = append(deletedRepos, r.PathValue("name"))
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true,
		provider.SweepOpts{Targets: []string{"repositories"}, Confirm: true})
	if err != nil {
		t.Fatalf("sweepWithClient error: %v", err)
	}

	mu.Lock()
	got := append([]string(nil), deletedRepos...)
	mu.Unlock()
	sort.Strings(got)
	want := []string{"tf-acc-test-page1", "tf-acc-test-page2"}
	if len(got) != len(want) {
		t.Fatalf("deleted repos = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("deleted repos = %v, want %v", got, want)
		}
	}
}

// TestSweepPrefixSafetyProductionName explicitly asserts that a repo named
// "production-database" is never deleted (invariant #1).
func TestSweepPrefixSafetyProductionName(t *testing.T) {
	var mu sync.Mutex
	var deletedRepos []string

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, reposJSON("tf-acc-test-foo", "production-database"))
	})
	mux.HandleFunc("DELETE /repos/"+sweepTestOwner+"/{name}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		deletedRepos = append(deletedRepos, r.PathValue("name"))
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	if err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true,
		provider.SweepOpts{Targets: []string{"repositories"}, Confirm: true}); err != nil {
		t.Fatalf("sweepWithClient error: %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	for _, name := range deletedRepos {
		if name == "production-database" {
			t.Error("production-database must NEVER be deleted")
		}
	}
}

// TestSweepTeamsOnlyInOrgMode verifies teams are deleted in org mode but skipped in user mode.
func TestSweepTeamsOnlyInOrgMode(t *testing.T) {
	t.Run("org_mode", func(t *testing.T) {
		var mu sync.Mutex
		var deletedSlugs []string

		mux := http.NewServeMux()
		mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, _ *http.Request) {
			writeJSON(w, teamsJSON("tf-acc-test-team-x", "real-team"))
		})
		mux.HandleFunc("DELETE /orgs/"+sweepTestOwner+"/teams/{slug}", func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			deletedSlugs = append(deletedSlugs, r.PathValue("slug"))
			mu.Unlock()
			w.WriteHeader(http.StatusNoContent)
		})
		srv := httptest.NewServer(mux)
		defer srv.Close()

		if err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true,
			provider.SweepOpts{Targets: []string{"teams"}, Confirm: true}); err != nil {
			t.Fatalf("sweepWithClient error: %v", err)
		}

		mu.Lock()
		got := append([]string(nil), deletedSlugs...)
		mu.Unlock()
		if len(got) != 1 || got[0] != "tf-acc-test-team-x" {
			t.Errorf("deleted teams = %v, want [tf-acc-test-team-x]", got)
		}
	})

	t.Run("user_mode", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("teams endpoint must not be called in user mode; method=%s", r.Method)
			http.Error(w, "unexpected", http.StatusInternalServerError)
		})
		srv := httptest.NewServer(mux)
		defer srv.Close()

		if err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, false,
			provider.SweepOpts{Targets: []string{"teams"}, Confirm: true}); err != nil {
			t.Fatalf("sweepWithClient user mode returned unexpected error: %v", err)
		}
	})
}

// TestSweepEmptyTargetsNoOp verifies that empty Targets with Confirm=true results in zero HTTP requests.
func TestSweepEmptyTargetsNoOp(t *testing.T) {
	var mu sync.Mutex
	var hitPaths []string

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		hitPaths = append(hitPaths, r.Method+" "+r.URL.Path)
		mu.Unlock()
		http.Error(w, "unexpected request", http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true,
		provider.SweepOpts{Targets: []string{}, Confirm: true})
	if err != nil {
		t.Fatalf("expected nil error for empty targets, got: %v", err)
	}

	mu.Lock()
	n := len(hitPaths)
	captured := append([]string(nil), hitPaths...)
	mu.Unlock()
	if n != 0 {
		t.Errorf("expected zero HTTP requests for empty targets, got %d: %v", n, captured)
	}
}

func TestSweepExactResourcesDeletesOnlySnapshotWithoutListing(t *testing.T) {
	var mu sync.Mutex
	var requests []string

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("exact sweep must not list repositories; got %s %s", r.Method, r.URL.String())
		http.Error(w, "unexpected list", http.StatusInternalServerError)
	})
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("exact sweep must not list teams; got %s %s", r.Method, r.URL.String())
		http.Error(w, "unexpected list", http.StatusInternalServerError)
	})
	mux.HandleFunc("DELETE /repos/"+sweepTestOwner+"/{name}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		requests = append(requests, r.Method+" "+r.URL.Path)
		mu.Unlock()
		if r.PathValue("name") == "tf-acc-test-c" {
			t.Error("exact sweep deleted unreviewed repository C")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("DELETE /orgs/"+sweepTestOwner+"/teams/{slug}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		requests = append(requests, r.Method+" "+r.URL.Path)
		mu.Unlock()
		if r.PathValue("slug") == "tf-acc-test-c" {
			t.Error("exact sweep deleted unreviewed team C")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true,
		provider.SweepOpts{
			Targets: []string{"repositories", "teams"},
			Confirm: true,
			ExactResources: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-a"},
				{Kind: "team", Name: "tf-acc-test-b"},
				{Kind: "repository", Name: "tf-acc-test-a"},
			},
		})
	if err != nil {
		t.Fatalf("sweepWithClient error: %v", err)
	}

	mu.Lock()
	got := append([]string(nil), requests...)
	mu.Unlock()
	sort.Strings(got)
	want := []string{
		"DELETE /orgs/" + sweepTestOwner + "/teams/tf-acc-test-b",
		"DELETE /repos/" + sweepTestOwner + "/tf-acc-test-a",
	}
	if len(got) != len(want) {
		t.Fatalf("requests = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("requests = %v, want %v", got, want)
		}
	}
}

func TestSweepExactResourcesDeletesOnlyRequestedKindAndName(t *testing.T) {
	var mu sync.Mutex
	var deletedRepos []string
	var deletedTeams []string

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, reposJSON("tf-acc-test-shared", "tf-acc-test-other-repo"))
	})
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, teamsJSON("tf-acc-test-shared", "tf-acc-test-team-only"))
	})
	mux.HandleFunc("DELETE /repos/"+sweepTestOwner+"/{name}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		deletedRepos = append(deletedRepos, r.PathValue("name"))
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("DELETE /orgs/"+sweepTestOwner+"/teams/{slug}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		deletedTeams = append(deletedTeams, r.PathValue("slug"))
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true, provider.SweepOpts{
		Targets: []string{"repositories", "teams"},
		Confirm: true,
		Resources: []provider.Resource{
			{Kind: "repository", Name: "tf-acc-test-shared", URL: "https://ignored.example/repo"},
			{Kind: "team", Name: "tf-acc-test-team-only", URL: "https://ignored.example/team"},
		},
	})
	if err != nil {
		t.Fatalf("sweepWithClient error: %v", err)
	}

	mu.Lock()
	gotRepos := append([]string(nil), deletedRepos...)
	gotTeams := append([]string(nil), deletedTeams...)
	mu.Unlock()
	if !reflect.DeepEqual(gotRepos, []string{"tf-acc-test-shared"}) {
		t.Fatalf("deleted repos = %v, want [tf-acc-test-shared]", gotRepos)
	}
	if !reflect.DeepEqual(gotTeams, []string{"tf-acc-test-team-only"}) {
		t.Fatalf("deleted teams = %v, want [tf-acc-test-team-only]", gotTeams)
	}
}

func TestSweepExactResourcesValidatesBeforeDeleting(t *testing.T) {
	for _, tc := range []struct {
		name      string
		org       bool
		targets   []string
		resources []provider.Resource
	}{
		{
			name:    "unprefixed repository",
			org:     true,
			targets: []string{"repositories"},
			resources: []provider.Resource{
				{Kind: "repository", Name: "production-database"},
			},
		},
		{
			name:    "unsupported kind",
			org:     true,
			targets: []string{"repositories"},
			resources: []provider.Resource{
				{Kind: "branch", Name: "tf-acc-test-branch"},
			},
		},
		{
			name:    "repository target disabled",
			org:     true,
			targets: []string{"teams"},
			resources: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-repo"},
			},
		},
		{
			name:    "team target disabled",
			org:     true,
			targets: []string{"repositories"},
			resources: []provider.Resource{
				{Kind: "team", Name: "tf-acc-test-team"},
			},
		},
		{
			name:    "team unsupported in user mode",
			org:     false,
			targets: []string{"teams"},
			resources: []provider.Resource{
				{Kind: "team", Name: "tf-acc-test-team"},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var mu sync.Mutex
			var hitPaths []string

			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				mu.Lock()
				hitPaths = append(hitPaths, r.Method+" "+r.URL.Path)
				mu.Unlock()
				http.Error(w, "unexpected request", http.StatusInternalServerError)
			})
			srv := httptest.NewServer(mux)
			defer srv.Close()

			err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, tc.org,
				provider.SweepOpts{Targets: tc.targets, Confirm: true, ExactResources: tc.resources})
			if err == nil {
				t.Fatal("expected validation error, got nil")
			}

			mu.Lock()
			n := len(hitPaths)
			captured := append([]string(nil), hitPaths...)
			mu.Unlock()
			if n != 0 {
				t.Fatalf("expected zero HTTP requests before validation error, got %d: %v", n, captured)
			}
		})
	}
}

func TestSweepExactResourcesEmptyNoOp(t *testing.T) {
	var mu sync.Mutex
	var hitPaths []string

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		hitPaths = append(hitPaths, r.Method+" "+r.URL.Path)
		mu.Unlock()
		http.Error(w, "unexpected request", http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true,
		provider.SweepOpts{Targets: []string{"repositories"}, Confirm: true, ExactResources: []provider.Resource{}})
	if err != nil {
		t.Fatalf("expected nil error for empty exact resources, got: %v", err)
	}

	mu.Lock()
	n := len(hitPaths)
	captured := append([]string(nil), hitPaths...)
	mu.Unlock()
	if n != 0 {
		t.Errorf("expected zero HTTP requests for empty exact resources, got %d: %v", n, captured)
	}
}

func TestSweepNonNilEmptyResourcesDeletesNothing(t *testing.T) {
	var mu sync.Mutex
	var hitPaths []string

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		hitPaths = append(hitPaths, r.Method+" "+r.URL.Path)
		mu.Unlock()
		http.Error(w, "unexpected request", http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true, provider.SweepOpts{
		Targets:   []string{"repositories", "teams"},
		Confirm:   true,
		Resources: []provider.Resource{},
	})
	if err != nil {
		t.Fatalf("sweepWithClient error: %v", err)
	}

	mu.Lock()
	got := append([]string(nil), hitPaths...)
	mu.Unlock()
	if len(got) != 0 {
		t.Fatalf("requests = %v, want none", got)
	}
}

func TestSweepNilExactResourcesUsesBroadList(t *testing.T) {
	var mu sync.Mutex
	var requests []string

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		requests = append(requests, r.Method+" "+r.URL.Path)
		mu.Unlock()
		writeJSON(w, reposJSON("tf-acc-test-a", "tf-acc-test-c"))
	})
	mux.HandleFunc("DELETE /repos/"+sweepTestOwner+"/{name}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		requests = append(requests, r.Method+" "+r.URL.Path)
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	err := sweepWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, true,
		provider.SweepOpts{Targets: []string{"repositories"}, Confirm: true})
	if err != nil {
		t.Fatalf("sweepWithClient error: %v", err)
	}

	mu.Lock()
	got := append([]string(nil), requests...)
	mu.Unlock()
	sort.Strings(got)
	want := []string{
		"DELETE /repos/" + sweepTestOwner + "/tf-acc-test-a",
		"DELETE /repos/" + sweepTestOwner + "/tf-acc-test-c",
		"GET /orgs/" + sweepTestOwner + "/repos",
	}
	if len(got) != len(want) {
		t.Fatalf("requests = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("requests = %v, want %v", got, want)
		}
	}
}
