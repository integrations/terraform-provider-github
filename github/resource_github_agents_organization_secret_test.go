package github

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

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
	if diagnostics := resourceGithubAgentsOrganizationSecretCreateOrUpdate(t.Context(), d, meta); diagnostics.HasError() {
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
	if diagnostics := resourceGithubAgentsOrganizationSecretCreateOrUpdate(t.Context(), d, meta); !diagnostics.HasError() {
		t.Fatal("rotation ignored repository access lookup failure")
	}
}

func TestAccGithubAgentsOrganizationSecret(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)

	for _, encrypted := range []bool{false, true} {
		t.Run(fmt.Sprintf("encrypted_%t", encrypted), func(t *testing.T) {
			t.Parallel()
			name := strings.ToUpper(strings.ReplaceAll(testResourcePrefix, "-", "_") + acctest.RandString(testRandomIDLength))
			key, _, err := testAccConf.meta.v3client.Agents.GetOrgPublicKey(t.Context(), testAccConf.meta.name)
			if err != nil {
				t.Fatal(err)
			}
			config := func(value, visibility string) string {
				attribute := fmt.Sprintf("value = %q", value)
				if encrypted {
					ciphertext, err := encryptPlaintext(value, key.GetKey())
					if err != nil {
						t.Fatal(err)
					}
					attribute = fmt.Sprintf("key_id = %q\nvalue_encrypted = %q", key.GetKeyID(), base64.StdEncoding.EncodeToString(ciphertext))
				}
				return fmt.Sprintf(`
resource "github_agents_organization_secret" "test" {
  secret_name = %q
  visibility = %q
  %s
}
data "github_agents_organization_public_key" "test" {}
data "github_agents_organization_secrets" "test" { depends_on = [github_agents_organization_secret.test] }
output "test_secret_metadata" {
  value = [for secret in data.github_agents_organization_secrets.test.secrets : secret if secret.name == github_agents_organization_secret.test.secret_name]
}
`, name, visibility, attribute)
			}
			initial := config("initial", "all")
			selected := config("rotated", "selected")
			ignored := strings.Replace(selected, "  secret_name =", "  lifecycle { ignore_changes = [updated_at] }\n  secret_name =", 1)
			writeExternally := func() {
				time.Sleep(time.Second)
				ciphertext, err := encryptPlaintext("external", key.GetKey())
				if err != nil {
					t.Fatal(err)
				}
				_, err = testAccConf.meta.v3client.Agents.CreateOrUpdateOrgSecret(t.Context(), testAccConf.meta.name, name, github.SecretOrgRequest{KeyID: key.GetKeyID(), EncryptedValue: base64.StdEncoding.EncodeToString(ciphertext), Visibility: "selected"})
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
					{Config: initial, ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_agents_organization_public_key.test", tfjsonpath.New("key"), knownvalue.StringExact(key.GetKey())),
						statecheck.ExpectKnownOutputValue("test_secret_metadata", knownvalue.ListExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{"name": knownvalue.StringExact(name), "visibility": knownvalue.StringExact("all"), "created_at": knownvalue.NotNull(), "updated_at": knownvalue.NotNull()}),
						})),
						statecheck.ExpectKnownValue("github_agents_organization_secret.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
					}},
					{ResourceName: "github_agents_organization_secret.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"key_id", ignore}},
					{Config: initial, PlanOnly: true},
					{Config: config("rotated", "private"), ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_organization_secret.test", plancheck.ResourceActionUpdate)}}},
					{Config: selected},
					{PreConfig: writeExternally, Config: selected, ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_organization_secret.test", plancheck.ResourceActionUpdate)}}},
					{Config: ignored},
					{PreConfig: writeExternally, Config: ignored, ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("github_agents_organization_secret.test", plancheck.ResourceActionNoop)}}},
				},
			})
		})
	}
}
