package github

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"golang.org/x/crypto/nacl/box"
)

func TestAgentsSecretLifecycle(t *testing.T) {
	t.Parallel()
	publicKey, _, err := box.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	var mutex sync.Mutex
	var updated int
	var exists bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()
		switch r.URL.Path {
		case "/repos/test/repository":
			if err := json.NewEncoder(w).Encode(map[string]int{"id": 123}); err != nil {
				t.Error(err)
			}
		case "/repos/test/repository/agents/secrets/public-key":
			if err := json.NewEncoder(w).Encode(map[string]string{"key_id": "key", "key": base64.StdEncoding.EncodeToString(publicKey[:])}); err != nil {
				t.Error(err)
			}
		case "/repos/test/repository/agents/secrets/TEST":
			switch r.Method {
			case http.MethodPut:
				exists = true
				updated++
				w.WriteHeader(http.StatusNoContent)
			case http.MethodDelete:
				exists = false
				w.WriteHeader(http.StatusNoContent)
			default:
				if !exists {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				if err := json.NewEncoder(w).Encode(map[string]string{"name": "TEST", "created_at": "2026-01-01T00:00:00Z", "updated_at": time.Date(2026, 1, 1, 0, 0, updated, 0, time.UTC).Format(time.RFC3339)}); err != nil {
					t.Error(err)
				}
			}
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	meta := &Owner{name: "test", v3client: mustCreateTestGitHubClient(t, server.URL)}
	config := `resource "github_agents_secret" "test" {
  repository = "repository"
  secret_name = "TEST"
  value = "initial"
}`
	rotated := strings.Replace(config, "initial", "rotated", 1)
	ignored := strings.Replace(rotated, "  value", "  lifecycle { ignore_changes = [updated_at] }\n  value", 1)
	externalWrite := func() { mutex.Lock(); updated++; mutex.Unlock() }
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"github": func() (*schema.Provider, error) {
				provider := NewProvider("test", "none")()
				provider.ConfigureContextFunc = func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) { return meta, nil }
				return provider, nil
			},
		},
		Steps: []resource.TestStep{
			{Config: config},
			{ResourceName: "github_agents_secret.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"value", "key_id"}},
			{Config: config, PlanOnly: true},
			{Config: rotated, ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_secret.test", plancheck.ResourceActionUpdate)}}},
			{Config: rotated, PreConfig: externalWrite, ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_secret.test", plancheck.ResourceActionUpdate)}}},
			{Config: ignored},
			{Config: ignored, PreConfig: externalWrite, ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_secret.test", plancheck.ResourceActionNoop)}}},
		},
	})
	mutex.Lock()
	defer mutex.Unlock()
	if exists {
		t.Fatal("Terraform destroy left the agent secret behind")
	}
}

func TestAgentsSecretValidation(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name         string
		attributes   string
		errorPattern string
	}{
		{name: "missing_value", errorPattern: "one of.*value.*must be specified"},
		{name: "both_values", attributes: `value = "test"
value_encrypted = "c2VjcmV0"
key_id = "key"`, errorPattern: "only one of|conflicts"},
		{name: "missing_key", attributes: `value_encrypted = "c2VjcmV0"`, errorPattern: "key_id"},
		{name: "invalid_base64", attributes: `value_encrypted = "!"
key_id = "key"`, errorPattern: "base64"},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			resource.UnitTest(t, resource.TestCase{
				ProviderFactories: map[string]func() (*schema.Provider, error){
					"github": func() (*schema.Provider, error) {
						provider := NewProvider("test", "none")()
						provider.ConfigureContextFunc = nil
						return provider, nil
					},
				},
				Steps: []resource.TestStep{{
					Config: fmt.Sprintf(`resource "github_agents_secret" "test" {
repository = "repository"
secret_name = "TEST"
%s
}`, test.attributes),
					PlanOnly:    true,
					ExpectError: regexp.MustCompile(test.errorPattern),
				}},
			})
		})
	}
}

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

func TestAgentsOrganizationSecretRotation(t *testing.T) {
	t.Parallel()
	ids := []int64{123, 456}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/orgs/test/agents/secrets/TEST/repositories" {
			repositories := []map[string]int64{}
			for _, id := range ids {
				repositories = append(repositories, map[string]int64{"id": id})
			}
			if err := json.NewEncoder(w).Encode(map[string]any{"repositories": repositories}); err != nil {
				t.Error(err)
			}
			return
		}
		if r.URL.Path != "/orgs/test/agents/secrets/TEST" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method == http.MethodPut {
			var body struct {
				IDs []int64 `json:"selected_repository_ids"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			// GitHub replaces access with an empty list when the field is omitted.
			ids = body.IDs
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if _, err := fmt.Fprint(w, `{"name":"TEST","visibility":"selected","created_at":"2026-01-01T00:00:00Z","updated_at":"2026-01-02T00:00:00Z"}`); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	meta := &Owner{name: "test", IsOrganization: true, maxPerPage: 100, v3client: mustCreateTestGitHubClient(t, server.URL)}
	d := schema.TestResourceDataRaw(t, resourceGithubAgentsOrganizationSecret().Schema, map[string]any{"secret_name": "TEST", "visibility": "selected", "value_encrypted": "c2VjcmV0", "key_id": "key"})
	d.SetId("TEST")
	if diagnostics := resourceGithubAgentsOrganizationSecretUpdate(t.Context(), d, meta); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	observedIDs, _ := d.Get("selected_repository_ids").(*schema.Set)
	if !observedIDs.Equal(schema.NewSet(schema.HashInt, []any{123, 456})) {
		t.Fatal("rotation removed live repository access not present in Terraform state")
	}
	baseline := d.Get("updated_at")
	if err := d.Set("updated_at", "2026-01-01T00:00:00Z"); err != nil {
		t.Fatal(err)
	}
	if diagnostics := resourceGithubAgentsOrganizationSecretRead(t.Context(), d, meta); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	if d.Get("updated_at") == baseline || d.Get("remote_updated_at") != baseline {
		t.Fatal("refresh lost the provider timestamp baseline")
	}
}

func TestAgentsOrganizationSecretAccessDrift(t *testing.T) {
	t.Parallel()
	for _, ids := range [][]int{{123}, {}} {
		t.Run(fmt.Sprint(ids), func(t *testing.T) {
			t.Parallel()
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasSuffix(r.URL.Path, "/repositories") {
					repositories := []map[string]int{}
					for _, id := range ids {
						repositories = append(repositories, map[string]int{"id": id})
					}
					if err := json.NewEncoder(w).Encode(map[string]any{"repositories": repositories}); err != nil {
						t.Error(err)
					}
					return
				}
				if _, err := fmt.Fprint(w, `{"name":"TEST","visibility":"selected","created_at":"2026-01-01T00:00:00Z","updated_at":"2026-01-02T00:00:00Z"}`); err != nil {
					t.Error(err)
				}
			}))
			defer server.Close()
			meta := &Owner{name: "test", IsOrganization: true, maxPerPage: 100, v3client: mustCreateTestGitHubClient(t, server.URL)}
			previousIDs := []any{}
			if len(ids) == 0 {
				previousIDs = append(previousIDs, 123)
			}
			d := schema.TestResourceDataRaw(t, resourceGithubAgentsOrganizationSecret().Schema, map[string]any{
				"secret_name": "TEST", "visibility": "selected", "value": "placeholder",
				"updated_at": "2026-01-01 00:00:00 +0000 UTC", "selected_repository_ids": previousIDs,
			})
			d.SetId("TEST")
			if diagnostics := resourceGithubAgentsOrganizationSecretRead(t.Context(), d, meta); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			if d.Get("updated_at") != d.Get("remote_updated_at") {
				t.Fatal("repository access change was mistaken for secret value drift")
			}
			// Once access is unchanged, a later timestamp must still detect value drift.
			if err := d.Set("updated_at", "2026-01-01 00:00:00 +0000 UTC"); err != nil {
				t.Fatal(err)
			}
			if diagnostics := resourceGithubAgentsOrganizationSecretRead(t.Context(), d, meta); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			if d.Get("updated_at") == d.Get("remote_updated_at") {
				t.Fatal("value drift with unchanged access was ignored")
			}
		})
	}
}

func TestAgentsOrganizationSecretRotationAccessError(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			t.Error("rotation must not proceed when repository access cannot be read")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if strings.HasSuffix(r.URL.Path, "/repositories") {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if _, err := fmt.Fprint(w, `{"name":"TEST","visibility":"selected"}`); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	meta := &Owner{name: "test", IsOrganization: true, maxPerPage: 100, v3client: mustCreateTestGitHubClient(t, server.URL)}
	d := schema.TestResourceDataRaw(t, resourceGithubAgentsOrganizationSecret().Schema, map[string]any{"secret_name": "TEST", "visibility": "selected", "value_encrypted": "c2VjcmV0", "key_id": "key"})
	d.SetId("TEST")
	if diagnostics := resourceGithubAgentsOrganizationSecretUpdate(t.Context(), d, meta); !diagnostics.HasError() {
		t.Fatal("rotation ignored repository access lookup failure")
	}
}

func TestAgentsSecretsPagination(t *testing.T) {
	t.Parallel()
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
	defer server.Close()
	meta := &Owner{name: "test", IsOrganization: true, maxPerPage: 1, v3client: mustCreateTestGitHubClient(t, server.URL)}
	for _, source := range []*schema.Resource{dataSourceGithubAgentsSecrets(), dataSourceGithubAgentsOrganizationSecrets()} {
		d := schema.TestResourceDataRaw(t, source.Schema, map[string]any{"name": "repository"})
		if diagnostics := source.ReadContext(t.Context(), d, meta); diagnostics.HasError() {
			t.Fatal(diagnostics)
		}
		secrets, _ := d.Get("secrets").([]any)
		if len(secrets) != 2 {
			t.Fatal("secret listing did not follow pagination")
		}
	}
	d := schema.TestResourceDataRaw(t, resourceGithubAgentsOrganizationSecretRepositories().Schema, map[string]any{"secret_name": "TEST"})
	d.SetId("TEST")
	if diagnostics := resourceGithubAgentsOrganizationSecretRepositoriesRead(t.Context(), d, meta); diagnostics.HasError() {
		t.Fatalf("repository listing did not follow pagination: %v", diagnostics)
	}
	ids, _ := d.Get("selected_repository_ids").(*schema.Set)
	if ids.Len() != 2 {
		t.Fatal("repository listing did not follow pagination")
	}
	association := schema.TestResourceDataRaw(t, resourceGithubAgentsOrganizationSecretRepository().Schema, map[string]any{"secret_name": "TEST", "repository_id": 2})
	association.SetId("TEST:2")
	if diagnostics := resourceGithubAgentsOrganizationSecretRepositoryRead(t.Context(), association, meta); diagnostics.HasError() || association.Id() == "" {
		t.Fatalf("association on second page was lost: %v", diagnostics)
	}
}

func TestAgentsSecretAccessWrites(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name       string
		resource   *schema.Resource
		attributes map[string]any
		path       string
	}{
		{name: "complete_set", resource: resourceGithubAgentsOrganizationSecretRepositories(), attributes: map[string]any{"secret_name": "TEST", "selected_repository_ids": []any{123}}, path: "/orgs/test/agents/secrets/TEST/repositories"},
		{name: "individual", resource: resourceGithubAgentsOrganizationSecretRepository(), attributes: map[string]any{"secret_name": "TEST", "repository_id": 123}, path: "/orgs/test/agents/secrets/TEST/repositories/123"},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			var methods []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != test.path {
					t.Errorf("access write targeted wrong entity: %s", r.URL.Path)
				}
				methods = append(methods, r.Method)
				if test.name == "complete_set" {
					var body struct {
						IDs []int64 `json:"selected_repository_ids"`
					}
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Error(err)
					}
					if len(methods) == 1 && (len(body.IDs) != 1 || body.IDs[0] != 123) {
						t.Error("create did not set the selected repositories")
					}
					if len(methods) == 2 && (body.IDs == nil || len(body.IDs) != 0) {
						t.Error("delete must send an empty array, not null or omit access")
					}
				}
				w.WriteHeader(http.StatusNoContent)
			}))
			defer server.Close()
			meta := &Owner{name: "test", IsOrganization: true, v3client: mustCreateTestGitHubClient(t, server.URL)}
			d := schema.TestResourceDataRaw(t, test.resource.Schema, test.attributes)
			if diagnostics := test.resource.CreateContext(t.Context(), d, meta); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			if diagnostics := test.resource.DeleteContext(t.Context(), d, meta); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			deleteMethod := http.MethodPut
			if test.name == "individual" {
				deleteMethod = http.MethodDelete
			}
			if len(methods) != 2 || methods[0] != http.MethodPut || methods[1] != deleteMethod {
				t.Fatalf("unexpected access operations: %v", methods)
			}
		})
	}
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
