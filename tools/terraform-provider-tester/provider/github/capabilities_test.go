package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

// countingMux wraps an http.ServeMux and records how many times each
// registered pattern was hit, so capability-probe tests can assert exact
// request counts: dedup must collapse repeats to one request, and capabilities
// nobody asked for must make zero.
type countingMux struct {
	mu     sync.Mutex
	counts map[string]int
	mux    *http.ServeMux
}

func newCountingMux() *countingMux {
	return &countingMux{counts: make(map[string]int), mux: http.NewServeMux()}
}

func (c *countingMux) handle(pattern string, fn func(http.ResponseWriter, *http.Request)) {
	c.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c.mu.Lock()
		c.counts[pattern]++
		c.mu.Unlock()
		fn(w, r)
	})
}

func (c *countingMux) count(pattern string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[pattern]
}

// capabilityTestServer wires a countingMux serving all three registered
// capability endpoints (each a happy-path 200) and returns it alongside the
// underlying httptest server.
func capabilityTestServer(t *testing.T, owner, templateRepo string) (*countingMux, *httptest.Server) {
	t.Helper()
	cm := newCountingMux()
	cm.handle("/user/codespaces/secrets/public-key", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"key_id": "1", "key": "abc"})
	})
	cm.handle("/orgs/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"login": owner})
	})
	cm.handle(fmt.Sprintf("/repos/%s/%s", owner, templateRepo), func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"is_template": true})
	})
	srv := httptest.NewServer(cm.mux)
	t.Cleanup(srv.Close)
	return cm, srv
}

// TestCapabilityProbesAllRegisteredHappyPath verifies that every registered
// capability probe runs exactly once, in request order, when all three are
// requested and all three endpoints succeed.
func TestCapabilityProbesAllRegisteredHappyPath(t *testing.T) {
	const owner, templateRepo = "myorg", "my-template"
	cm, srv := capabilityTestServer(t, owner, templateRepo)

	client := newClient(srv.URL+"/", "ghp_fake")
	names := []string{"codespaces-user-secrets", "organization", "template-repository"}
	checks := probeCapabilities(context.Background(), client, names, capabilityInput{Owner: owner, TemplateRepo: templateRepo})

	if len(checks) != 3 {
		t.Fatalf("probeCapabilities returned %d checks, want 3: %+v", len(checks), checks)
	}
	wantNames := []string{"capability/codespaces-user-secrets", "capability/organization", "capability/template-repository"} // gitleaks:allow -- check names, not secrets
	for i, want := range wantNames {
		if checks[i].Name != want {
			t.Errorf("checks[%d].Name = %q, want %q", i, checks[i].Name, want)
		}
		if checks[i].Status != provider.CheckOK {
			t.Errorf("checks[%d] (%s) status = %v, want CheckOK (detail: %s)", i, checks[i].Name, checks[i].Status, checks[i].Detail)
		}
		if checks[i].Detail == "" {
			t.Errorf("checks[%d] (%s) has empty Detail", i, checks[i].Name)
		}
	}
	for _, path := range []string{"/user/codespaces/secrets/public-key", "/orgs/", fmt.Sprintf("/repos/%s/%s", owner, templateRepo)} {
		if got := cm.count(path); got != 1 {
			t.Errorf("request count for %s = %d, want 1", path, got)
		}
	}
}

// TestCapabilityProbesDeduplicateRepeatedNames verifies that requesting the same
// capability name multiple times still makes exactly one HTTP request and
// returns exactly one Check.
func TestCapabilityProbesDeduplicateRepeatedNames(t *testing.T) {
	const owner = "myorg"
	cm, srv := capabilityTestServer(t, owner, "unused-template")

	client := newClient(srv.URL+"/", "ghp_fake")
	names := []string{"organization", "organization", "organization"}
	checks := probeCapabilities(context.Background(), client, names, capabilityInput{Owner: owner})

	if len(checks) != 1 {
		t.Fatalf("probeCapabilities returned %d checks for 3 duplicate names, want 1: %+v", len(checks), checks)
	}
	if checks[0].Name != "capability/organization" {
		t.Errorf("checks[0].Name = %q, want %q", checks[0].Name, "capability/organization")
	}
	if got := cm.count("/orgs/"); got != 1 {
		t.Errorf("request count for /orgs/ = %d, want 1 (duplicate names must collapse to one request)", got)
	}
}

// TestCapabilityProbesSkipUnrequestedCapabilities verifies that capabilities
// nobody asked for make zero requests, even though they are registered.
func TestCapabilityProbesSkipUnrequestedCapabilities(t *testing.T) {
	const owner, templateRepo = "myorg", "my-template"
	cm, srv := capabilityTestServer(t, owner, templateRepo)

	client := newClient(srv.URL+"/", "ghp_fake")
	// Only "organization" is requested; codespaces-user-secrets and
	// template-repository must make zero requests.
	checks := probeCapabilities(context.Background(), client, []string{"organization"}, capabilityInput{Owner: owner, TemplateRepo: templateRepo})

	if len(checks) != 1 {
		t.Fatalf("probeCapabilities returned %d checks, want 1: %+v", len(checks), checks)
	}
	for _, path := range []string{"/user/codespaces/secrets/public-key", fmt.Sprintf("/repos/%s/%s", owner, templateRepo)} {
		if got := cm.count(path); got != 0 {
			t.Errorf("request count for unrequested capability path %s = %d, want 0", path, got)
		}
	}
	if got := cm.count("/orgs/"); got != 1 {
		t.Errorf("request count for /orgs/ = %d, want 1", got)
	}
}

// TestCapabilityProbesEmptyNamesMakeNoRequests verifies that an empty/nil names
// slice produces no checks and no HTTP requests at all.
func TestCapabilityProbesEmptyNamesMakeNoRequests(t *testing.T) {
	const owner = "myorg"
	cm, srv := capabilityTestServer(t, owner, "my-template")

	client := newClient(srv.URL+"/", "ghp_fake")
	checks := probeCapabilities(context.Background(), client, nil, capabilityInput{Owner: owner})

	if len(checks) != 0 {
		t.Fatalf("probeCapabilities returned %d checks for nil names, want 0: %+v", len(checks), checks)
	}
	for _, path := range []string{"/user/codespaces/secrets/public-key", "/orgs/", "/repos/myorg/my-template"} {
		if got := cm.count(path); got != 0 {
			t.Errorf("request count for %s = %d, want 0", path, got)
		}
	}
}

// TestCapabilityProbesUnknownNameFailsClosedWithoutRequests verifies that a
// capability name absent from the registry produces a CheckFail (never a
// silent success) and makes no HTTP request at all.
func TestCapabilityProbesUnknownNameFailsClosedWithoutRequests(t *testing.T) {
	const owner = "myorg"
	cm, srv := capabilityTestServer(t, owner, "my-template")

	client := newClient(srv.URL+"/", "ghp_fake")
	checks := probeCapabilities(context.Background(), client, []string{"totally-unregistered-capability"}, capabilityInput{Owner: owner})

	if len(checks) != 1 {
		t.Fatalf("probeCapabilities returned %d checks, want 1: %+v", len(checks), checks)
	}
	c := checks[0]
	if c.Name != "capability/totally-unregistered-capability" {
		t.Errorf("Name = %q, want %q", c.Name, "capability/totally-unregistered-capability")
	}
	if c.Status != provider.CheckFail {
		t.Errorf("Status = %v, want CheckFail (unknown capability must fail closed, never succeed)", c.Status)
	}
	if c.Detail == "" {
		t.Error("expected non-empty Detail for unknown capability")
	}
	if c.Fix == "" {
		t.Error("expected non-empty Fix for unknown capability")
	}
	for _, path := range []string{"/user/codespaces/secrets/public-key", "/orgs/", "/repos/myorg/my-template"} {
		if got := cm.count(path); got != 0 {
			t.Errorf("unknown capability must not make any HTTP request, got %d for %s", got, path)
		}
	}
}

// TestCapabilityProbesFailureIncludesDetailAndFix verifies that every
// registered capability's failure path carries a non-empty problem Detail and
// a non-empty, actionable Fix.
func TestCapabilityProbesFailureIncludesDetailAndFix(t *testing.T) {
	cases := []string{"codespaces-user-secrets", "organization", "template-repository"}

	for _, capName := range cases {
		t.Run(capName, func(t *testing.T) {
			const owner, templateRepo = "myorg", "my-template"
			mux := http.NewServeMux()
			mux.HandleFunc("/user/codespaces/secrets/public-key", func(w http.ResponseWriter, _ *http.Request) {
				http.Error(w, `{"message":"Forbidden"}`, http.StatusForbidden)
			})
			mux.HandleFunc("/orgs/", func(w http.ResponseWriter, _ *http.Request) {
				http.Error(w, `{"message":"Not Found"}`, http.StatusNotFound)
			})
			mux.HandleFunc(fmt.Sprintf("/repos/%s/%s", owner, templateRepo), func(w http.ResponseWriter, _ *http.Request) {
				http.Error(w, `{"message":"Not Found"}`, http.StatusNotFound)
			})
			srv := httptest.NewServer(mux)
			t.Cleanup(srv.Close)

			client := newClient(srv.URL+"/", "ghp_fake")
			checks := probeCapabilities(context.Background(), client, []string{capName}, capabilityInput{Owner: owner, TemplateRepo: templateRepo})
			if len(checks) != 1 {
				t.Fatalf("probeCapabilities returned %d checks, want 1: %+v", len(checks), checks)
			}
			c := checks[0]
			if c.Status != provider.CheckFail {
				t.Errorf("status = %v, want CheckFail", c.Status)
			}
			if c.Detail == "" {
				t.Error("expected non-empty Detail")
			}
			if c.Fix == "" {
				t.Error("expected non-empty Fix")
			}
		})
	}
}

// TestCapabilityProbeTemplateRepositoryNotMarkedAsTemplateFails verifies that
// an existing-but-not-a-template repo fails with a Fix naming the remediation,
// distinct from the "repo does not exist" failure case above.
func TestCapabilityProbeTemplateRepositoryNotMarkedAsTemplateFails(t *testing.T) {
	const owner, templateRepo = "myorg", "not-a-template"
	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("/repos/%s/%s", owner, templateRepo), func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"is_template": false})
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)

	client := newClient(srv.URL+"/", "ghp_fake")
	checks := probeCapabilities(context.Background(), client, []string{"template-repository"}, capabilityInput{Owner: owner, TemplateRepo: templateRepo})
	if len(checks) != 1 {
		t.Fatalf("probeCapabilities returned %d checks, want 1", len(checks))
	}
	c := checks[0]
	if c.Status != provider.CheckFail {
		t.Errorf("status = %v, want CheckFail", c.Status)
	}
	if !strings.Contains(c.Fix, "mark the repo as a template") {
		t.Errorf("Fix = %q, expected it to mention marking as template", c.Fix)
	}
}
