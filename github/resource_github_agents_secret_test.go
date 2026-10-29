package github

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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

func TestAccGithubAgentsSecret(t *testing.T) {
	t.Parallel()
	skipUnauthenticated(t)

	for _, encrypted := range []bool{false, true} {
		t.Run(fmt.Sprintf("encrypted_%t", encrypted), func(t *testing.T) {
			t.Parallel()
			repo := mustCreateTestRepository(t)
			name := "TEST"
			key, _, err := testAccConf.meta.v3client.Agents.GetRepoPublicKey(t.Context(), testAccConf.meta.name, repo.GetName())
			if err != nil {
				t.Fatal(err)
			}
			config := func(repository, value string) string {
				attribute := fmt.Sprintf("value = %q", value)
				if encrypted {
					ciphertext, err := encryptPlaintext(value, key.GetKey())
					if err != nil {
						t.Fatal(err)
					}
					attribute = fmt.Sprintf("key_id = %q\nvalue_encrypted = %q", key.GetKeyID(), base64.StdEncoding.EncodeToString(ciphertext))
				}
				return fmt.Sprintf(`
resource "github_agents_secret" "test" {
  repository = %q
  secret_name = %q
  %s
}
data "github_agents_public_key" "test" { repository = github_agents_secret.test.repository }
data "github_agents_secrets" "test" {
  full_name = %q
  depends_on = [github_agents_secret.test]
}
`, repository, name, attribute, testAccConf.meta.name+"/"+repository)
			}
			first := config(repo.GetName(), "initial")
			second := config(repo.GetName(), "rotated")
			ignored := strings.Replace(second, "  secret_name =", "  lifecycle { ignore_changes = [updated_at] }\n  secret_name =", 1)
			writeExternally := func() {
				// GitHub timestamps have second precision; ensure the external write is newer.
				time.Sleep(time.Second)
				ciphertext, err := encryptPlaintext("external", key.GetKey())
				if err != nil {
					t.Fatal(err)
				}
				_, err = testAccConf.meta.v3client.Agents.CreateOrUpdateRepoSecret(t.Context(), testAccConf.meta.name, repo.GetName(), name, github.SecretRequest{KeyID: key.GetKeyID(), EncryptedValue: base64.StdEncoding.EncodeToString(ciphertext)})
				if err != nil {
					t.Fatal(err)
				}
			}
			ignore := "value"
			if encrypted {
				ignore = "value_encrypted"
			}
			resource.Test(t, resource.TestCase{
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config: first,
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue("github_agents_secret.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
							statecheck.ExpectKnownValue("github_agents_secret.test", tfjsonpath.New("key_id"), knownvalue.StringExact(key.GetKeyID())),
							statecheck.ExpectKnownValue("data.github_agents_public_key.test", tfjsonpath.New("key"), knownvalue.StringExact(key.GetKey())),
							statecheck.ExpectKnownValue("data.github_agents_secrets.test", tfjsonpath.New("secrets"), knownvalue.ListExact([]knownvalue.Check{
								knownvalue.MapExact(map[string]knownvalue.Check{"name": knownvalue.StringExact(name), "created_at": knownvalue.NotNull(), "updated_at": knownvalue.NotNull()}),
							})),
						},
					},
					{ResourceName: "github_agents_secret.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"key_id", ignore}},
					{Config: first, PlanOnly: true},
					{Config: second, ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_secret.test", plancheck.ResourceActionUpdate)}}},
					{
						PreConfig:        writeExternally,
						Config:           second,
						ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_secret.test", plancheck.ResourceActionUpdate)}},
					},
					{Config: ignored},
					{PreConfig: writeExternally, Config: ignored, ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_secret.test", plancheck.ResourceActionNoop)}}},
				},
			})
		})
	}

	t.Run("repository_changes", func(t *testing.T) {
		t.Parallel()
		repo := mustCreateTestRepository(t)
		other := mustCreateTestRepository(t)
		renamed := repo.GetName() + "-renamed"
		config := func(name string) string {
			return fmt.Sprintf(`resource "github_agents_secret" "test" {
  repository = %q
  secret_name = "TEST"
  value = "placeholder"
}`, name)
		}
		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: config(repo.GetName())},
				{PreConfig: func() { mustRenameTestRepository(t, repo, renamed) }, Config: config(renamed), ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_secret.test", plancheck.ResourceActionUpdate)}}},
				{Config: config(other.GetName()), ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_secret.test", plancheck.ResourceActionReplace)}}},
			},
		})
	})
}
