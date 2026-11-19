package github

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/crypto/nacl/box"
)

func TestAgentsSecretEncryption(t *testing.T) {
	t.Parallel()
	publicKey, privateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/test/repository/agents/secrets/public-key" || r.Header.Get("X-GitHub-Api-Version") != "2026-03-10" {
			t.Errorf("unexpected agent key request: %s, version %s", r.URL.Path, r.Header.Get("X-GitHub-Api-Version"))
		}
		if err := json.NewEncoder(w).Encode(map[string]string{"key_id": "agent-key", "key": base64.StdEncoding.EncodeToString(publicKey[:])}); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	meta := &Owner{name: "test", v3client: mustCreateTestGitHubClient(t, server.URL)}
	// Include empty plaintext: GetOk must not accidentally exclude a valid empty secret.
	for _, plaintext := range []string{"", "secret"} {
		d := schema.TestResourceDataRaw(t, resourceGithubAgentsSecret().Schema, map[string]any{"value": plaintext})
		request, err := agentSecretValue(t.Context(), d, meta, "repository")
		if err != nil {
			t.Fatal(err)
		}
		ciphertext, err := base64.StdEncoding.DecodeString(request.EncryptedValue)
		if err != nil {
			t.Fatal(err)
		}
		decrypted, ok := box.OpenAnonymous(nil, ciphertext, publicKey, privateKey)
		if !ok || string(decrypted) != plaintext || request.KeyID != "agent-key" {
			t.Fatal("agent secret did not round-trip with its public key")
		}
	}
	d := schema.TestResourceDataRaw(t, resourceGithubAgentsSecret().Schema, map[string]any{"value_encrypted": "c2VjcmV0", "key_id": "external-key"})
	request, err := agentSecretValue(t.Context(), d, &Owner{}, "repository")
	if err != nil || request.KeyID != "external-key" || request.EncryptedValue != "c2VjcmV0" {
		t.Fatalf("pre-encrypted input was not preserved: %v", err)
	}
}

func TestAgentsSecretMissing(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNotFound) }))
	defer server.Close()
	meta := &Owner{name: "test", IsOrganization: true, v3client: mustCreateTestGitHubClient(t, server.URL), maxPerPage: 100}
	for _, resource := range []*schema.Resource{resourceGithubAgentsSecret(), resourceGithubAgentsOrganizationSecret(), resourceGithubAgentsOrganizationSecretRepositories(), resourceGithubAgentsOrganizationSecretRepository()} {
		d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"secret_name": "TEST", "repository": "repository", "repository_id": 123})
		d.SetId("TEST")
		if diagnostics := resource.ReadContext(t.Context(), d, meta); diagnostics.HasError() || d.Id() != "" {
			t.Fatalf("read did not remove missing resource: %v", diagnostics)
		}
		if diagnostics := resource.DeleteContext(t.Context(), d, meta); diagnostics.HasError() {
			t.Fatalf("delete did not tolerate missing resource: %v", diagnostics)
		}
	}
}

func agentSecretsPaginationTestOwner(t *testing.T) *Owner {
	t.Helper()
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("page") == "2" {
			if _, err := fmt.Fprint(w, `{"secrets":[{"name":"SECOND"}],"repositories":[{"id":2}]}`); err != nil {
				t.Error(err)
			}
			return
		}
		w.Header().Set("Link", fmt.Sprintf(`<%s%s?per_page=1&page=2>; rel="next"`, server.URL, r.URL.Path))
		if _, err := fmt.Fprint(w, `{"secrets":[{"name":"FIRST"}],"repositories":[{"id":1}]}`); err != nil {
			t.Error(err)
		}
	}))
	t.Cleanup(server.Close)
	return &Owner{name: "test", IsOrganization: true, maxPerPage: 1, v3client: mustCreateTestGitHubClient(t, server.URL)}
}

func TestAgentsSecretReadError(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusForbidden) }))
	defer server.Close()
	meta := &Owner{name: "test", IsOrganization: true, v3client: mustCreateTestGitHubClient(t, server.URL), maxPerPage: 100}
	for _, resource := range []*schema.Resource{resourceGithubAgentsSecret(), resourceGithubAgentsOrganizationSecret(), resourceGithubAgentsOrganizationSecretRepositories(), resourceGithubAgentsOrganizationSecretRepository()} {
		d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"secret_name": "TEST", "repository": "repository", "repository_id": 123})
		d.SetId("TEST")
		if diagnostics := resource.ReadContext(t.Context(), d, meta); !diagnostics.HasError() || d.Id() != "TEST" {
			t.Fatalf("read swallowed an authorization error or discarded state: %v", diagnostics)
		}
	}
}

func TestAgentsSecretAccessDeleteConflict(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name       string
		visibility string
		lookupCode int
		wantError  bool
	}{
		{name: "all", visibility: "all", lookupCode: http.StatusOK},
		{name: "private", visibility: "private", lookupCode: http.StatusOK},
		{name: "selected", visibility: "selected", lookupCode: http.StatusOK, wantError: true},
		{name: "deleted", lookupCode: http.StatusNotFound},
		{name: "forbidden", lookupCode: http.StatusForbidden, wantError: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					w.WriteHeader(http.StatusConflict)
					return
				}
				w.WriteHeader(test.lookupCode)
				if test.lookupCode == http.StatusOK {
					if err := json.NewEncoder(w).Encode(map[string]string{"visibility": test.visibility}); err != nil {
						t.Error(err)
					}
				}
			}))
			defer server.Close()
			meta := &Owner{name: "test", IsOrganization: true, v3client: mustCreateTestGitHubClient(t, server.URL)}
			for _, resource := range []*schema.Resource{resourceGithubAgentsOrganizationSecretRepositories(), resourceGithubAgentsOrganizationSecretRepository()} {
				d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"secret_name": "TEST", "repository_id": 123})
				if diagnostics := resource.DeleteContext(t.Context(), d, meta); diagnostics.HasError() != test.wantError {
					t.Fatalf("unexpected delete result for %s: %v", test.name, diagnostics)
				}
			}
		})
	}
}
