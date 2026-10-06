package github

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

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
						statecheck.ExpectKnownValue("data.github_agents_organization_secrets.test", tfjsonpath.New("secrets"), knownvalue.NotNull()),
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
