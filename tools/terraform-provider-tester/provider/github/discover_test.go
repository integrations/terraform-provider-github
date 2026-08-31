package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

// orgsJSON encodes a slice of org login strings as go-github expects.
func orgsJSON(logins ...string) []byte {
	type orgItem struct {
		Login string `json:"login"`
	}
	items := make([]orgItem, len(logins))
	for i, l := range logins {
		items[i] = orgItem{Login: l}
	}
	b, _ := json.Marshal(items)
	return b
}

// enterprisesGQLResp builds a minimal GraphQL response body.
func enterprisesGQLResp(slugs, names []string) []byte {
	type node struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	}
	nodes := make([]node, len(slugs))
	for i := range slugs {
		nodes[i] = node{Slug: slugs[i], Name: names[i]}
	}
	payload := map[string]any{
		"data": map[string]any{
			"viewer": map[string]any{
				"enterprises": map[string]any{
					"nodes": nodes,
				},
			},
		},
	}
	b, _ := json.Marshal(payload)
	return b
}

// TestDiscoverOrgs verifies that GET /user/orgs is called and org logins are returned.
func TestDiscoverOrgs(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(orgsJSON("acme", "widgets")) //nolint:errcheck
	})
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(enterprisesGQLResp(nil, nil)) //nolint:errcheck
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	res, err := discoverWithClient(
		context.Background(),
		newClient(srv.URL+"/", "ghp_fake"),
		credInfo{},
		srv.URL+"/",
		"ghp_fake",
		provider.DiscoverOpts{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.Orgs) != 2 {
		t.Fatalf("expected 2 orgs, got %d: %v", len(res.Orgs), res.Orgs)
	}
	orgsSet := map[string]bool{}
	for _, o := range res.Orgs {
		orgsSet[o] = true
	}
	for _, want := range []string{"acme", "widgets"} {
		if !orgsSet[want] {
			t.Errorf("expected org %q in result; got %v", want, res.Orgs)
		}
	}
}

// TestDiscoverOrgsPagination verifies that GET /user/orgs paginates across two pages.
func TestDiscoverOrgsPagination(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("page") == "2" {
			w.Write(orgsJSON("page2org")) //nolint:errcheck
			return
		}
		w.Header().Set("Link", fmt.Sprintf(`<%s/user/orgs?page=2>; rel="next"`, "http://"+r.Host))
		w.Write(orgsJSON("page1org")) //nolint:errcheck
	})
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(enterprisesGQLResp(nil, nil)) //nolint:errcheck
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	res, err := discoverWithClient(
		context.Background(),
		newClient(srv.URL+"/", "ghp_fake"),
		credInfo{},
		srv.URL+"/",
		"ghp_fake",
		provider.DiscoverOpts{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.Orgs) != 2 {
		t.Fatalf("expected 2 orgs across pages, got %d: %v", len(res.Orgs), res.Orgs)
	}
	found := map[string]bool{}
	for _, o := range res.Orgs {
		found[o] = true
	}
	for _, want := range []string{"page1org", "page2org"} {
		if !found[want] {
			t.Errorf("missing org %q after pagination; got %v", want, res.Orgs)
		}
	}
}

// TestDiscoverEnterprises verifies that the GraphQL enterprises query is parsed.
func TestDiscoverEnterprises(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(orgsJSON()) //nolint:errcheck
	})
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "bearer ghp_fake" {
			t.Errorf("Authorization header = %q, want %q", got, "bearer ghp_fake")
		}
		var reqBody struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("failed to decode GraphQL request body: %v", err)
		} else if want := "query{viewer{enterprises(first:100){nodes{slug name}}}}"; reqBody.Query != want {
			t.Errorf("GraphQL query = %q, want %q", reqBody.Query, want)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(enterprisesGQLResp([]string{"e1"}, []string{"E One"})) //nolint:errcheck
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	res, err := discoverWithClient(
		context.Background(),
		newClient(srv.URL+"/", "ghp_fake"),
		credInfo{},
		srv.URL+"/",
		"ghp_fake",
		provider.DiscoverOpts{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.Enterprises) != 1 {
		t.Fatalf("expected 1 enterprise, got %d: %v", len(res.Enterprises), res.Enterprises)
	}
	if res.Enterprises[0].Slug != "e1" {
		t.Errorf("enterprise slug = %q, want %q", res.Enterprises[0].Slug, "e1")
	}
	if res.Enterprises[0].Name != "E One" {
		t.Errorf("enterprise name = %q, want %q", res.Enterprises[0].Name, "E One")
	}
}

// TestDiscoverTemplateRepos verifies that --org lists only template repos.
func TestDiscoverTemplateRepos(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(orgsJSON()) //nolint:errcheck
	})
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(enterprisesGQLResp(nil, nil)) //nolint:errcheck
	})
	mux.HandleFunc("/orgs/acme/repos", func(w http.ResponseWriter, _ *http.Request) {
		type repoItem struct {
			Name       string `json:"name"`
			IsTemplate bool   `json:"is_template"`
		}
		repos := []repoItem{
			{Name: "tmpl-a", IsTemplate: true},
			{Name: "normal-b", IsTemplate: false},
			{Name: "tmpl-c", IsTemplate: true},
		}
		w.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(repos)
		w.Write(b) //nolint:errcheck
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	res, err := discoverWithClient(
		context.Background(),
		newClient(srv.URL+"/", "ghp_fake"),
		credInfo{},
		srv.URL+"/",
		"ghp_fake",
		provider.DiscoverOpts{Org: "acme"},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.TemplateRepos) != 2 {
		t.Fatalf("expected 2 template repos, got %d: %v", len(res.TemplateRepos), res.TemplateRepos)
	}
	found := map[string]bool{}
	for _, r := range res.TemplateRepos {
		found[r] = true
	}
	for _, want := range []string{"tmpl-a", "tmpl-c"} {
		if !found[want] {
			t.Errorf("missing template repo %q; got %v", want, res.TemplateRepos)
		}
	}
	if found["normal-b"] {
		t.Error("non-template repo should not appear in TemplateRepos")
	}
}

// TestDiscoverOrgsDegradation403 verifies that a 403 on /user/orgs appends a note but does not fail.
func TestDiscoverOrgsDegradation403(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, `{"message":"Forbidden"}`, http.StatusForbidden)
	})
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(enterprisesGQLResp(nil, nil)) //nolint:errcheck
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	res, err := discoverWithClient(
		context.Background(),
		newClient(srv.URL+"/", "ghp_fake"),
		credInfo{},
		srv.URL+"/",
		"ghp_fake",
		provider.DiscoverOpts{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.Orgs) != 0 {
		t.Errorf("expected no orgs on 403, got %v", res.Orgs)
	}
	if len(res.Notes) == 0 {
		t.Error("expected a note when orgs call returned 403")
	}
	found := false
	for _, n := range res.Notes {
		if len(n) >= len("orgs unverified:") && n[:len("orgs unverified:")] == "orgs unverified:" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a note starting with 'orgs unverified:', got %v", res.Notes)
	}
}

// TestDiscoverEnterprisesDegradation401 verifies that a 401 on /graphql appends a note but does not fail.
func TestDiscoverEnterprisesDegradation401(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(orgsJSON("someorg")) //nolint:errcheck
	})
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, `{"message":"Unauthorized"}`, http.StatusUnauthorized)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	res, err := discoverWithClient(
		context.Background(),
		newClient(srv.URL+"/", "ghp_fake"),
		credInfo{},
		srv.URL+"/",
		"ghp_fake",
		provider.DiscoverOpts{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.Orgs) != 1 || res.Orgs[0] != "someorg" {
		t.Errorf("expected orgs=[someorg], got %v", res.Orgs)
	}
	if len(res.Enterprises) != 0 {
		t.Errorf("expected no enterprises on 401, got %v", res.Enterprises)
	}
	found := false
	for _, n := range res.Notes {
		if len(n) >= len("enterprises unverified:") && n[:len("enterprises unverified:")] == "enterprises unverified:" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a note starting with 'enterprises unverified:', got %v", res.Notes)
	}
}

// TestDiscoverOrgsPaginationBounded verifies that the orgs loop stops after
// at most maxDiscoverPages iterations even when the server always signals
// another page. This is a regression test for the unbounded-loop hang.
func TestDiscoverOrgsPaginationBounded(t *testing.T) {
	var count atomic.Int32

	mux := http.NewServeMux()
	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, r *http.Request) {
		n := int(count.Add(1))
		w.Header().Set("Content-Type", "application/json")
		if n <= maxDiscoverPages {
			// Always advertise another page so an unbounded loop never exits.
			w.Header().Set("Link", fmt.Sprintf(`<%s/user/orgs?page=%d>; rel="next"`, "http://"+r.Host, n+1))
		}
		w.Write(orgsJSON(fmt.Sprintf("org%d", n))) //nolint:errcheck
	})
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(enterprisesGQLResp(nil, nil)) //nolint:errcheck
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	_, err := discoverWithClient(
		context.Background(),
		newClient(srv.URL+"/", "ghp_fake"),
		credInfo{},
		srv.URL+"/",
		"ghp_fake",
		provider.DiscoverOpts{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := int(count.Load())
	if got > maxDiscoverPages {
		t.Fatalf("orgs pagination made %d requests, want <= %d", got, maxDiscoverPages)
	}
}

// TestDiscoverEnterprisesGQLErrors verifies that a GraphQL HTTP 200 response
// with a top-level errors array appends a note and leaves Enterprises empty.
func TestDiscoverEnterprisesGQLErrors(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(orgsJSON()) //nolint:errcheck
	})
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"errors":[{"message":"scope 'read:enterprise' required"}]}`)) //nolint:errcheck
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	res, err := discoverWithClient(
		context.Background(),
		newClient(srv.URL+"/", "ghp_fake"),
		credInfo{},
		srv.URL+"/",
		"ghp_fake",
		provider.DiscoverOpts{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.Enterprises) != 0 {
		t.Errorf("expected no enterprises on GQL error response, got %v", res.Enterprises)
	}
	wantPrefix := "enterprises unverified:"
	wantMsg := "scope 'read:enterprise' required"
	found := false
	for _, n := range res.Notes {
		if strings.HasPrefix(n, wantPrefix) && strings.Contains(n, wantMsg) {
			found = true
		}
	}
	if !found {
		t.Errorf("expected note with prefix %q containing %q, got %v", wantPrefix, wantMsg, res.Notes)
	}
}

// TestDiscoverOrgsPaginationTruncationNote verifies that when the orgs loop
// hits the page cap while more pages remain, a note is appended.
func TestDiscoverOrgsPaginationTruncationNote(t *testing.T) {
	var count atomic.Int32

	mux := http.NewServeMux()
	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, r *http.Request) {
		n := int(count.Add(1))
		w.Header().Set("Content-Type", "application/json")
		// Always advertise another page so the loop must stop at the cap.
		w.Header().Set("Link", fmt.Sprintf(`<%s/user/orgs?page=%d>; rel="next"`, "http://"+r.Host, n+1))
		w.Write(orgsJSON(fmt.Sprintf("org%d", n))) //nolint:errcheck
	})
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(enterprisesGQLResp(nil, nil)) //nolint:errcheck
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	res, err := discoverWithClient(
		context.Background(),
		newClient(srv.URL+"/", "ghp_fake"),
		credInfo{},
		srv.URL+"/",
		"ghp_fake",
		provider.DiscoverOpts{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := int(count.Load())
	if got > maxDiscoverPages {
		t.Fatalf("orgs pagination made %d requests, want <= %d", got, maxDiscoverPages)
	}
	found := false
	for _, n := range res.Notes {
		if strings.HasPrefix(n, "orgs may be incomplete:") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected note starting with 'orgs may be incomplete:', got %v", res.Notes)
	}
}
