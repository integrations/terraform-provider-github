package github

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

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
