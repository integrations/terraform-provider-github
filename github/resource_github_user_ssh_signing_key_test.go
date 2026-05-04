package github

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"golang.org/x/crypto/ssh"
)

func TestAccGithubUserSshSigningKey(t *testing.T) {
	t.Run("creates and destroys a user SSH signing key without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		name := fmt.Sprintf(`%s-%s`, testResourcePrefix, randomID)
		testKey := newTestSigningKey()

		config := fmt.Sprintf(`
			resource "github_user_ssh_signing_key" "test" {
				title = "%[1]s"
				key   = "%[2]s"
			}
		`, name, testKey)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_user_ssh_signing_key.test", tfjsonpath.New("title"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("github_user_ssh_signing_key.test", tfjsonpath.New("key"), knownvalue.StringExact(testKey)),
					},
				},
			},
		})
	})

	t.Run("imports an individual account SSH signing key without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		name := fmt.Sprintf(`%s-%s`, testResourcePrefix, randomID)
		testKey := newTestSigningKey()

		config := fmt.Sprintf(`
			resource "github_user_ssh_signing_key" "test" {
				title = "%[1]s"
				key   = "%[2]s"
			}
		`, name, testKey)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_user_ssh_signing_key.test", tfjsonpath.New("title"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue("github_user_ssh_signing_key.test", tfjsonpath.New("key"), knownvalue.StringExact(testKey)),
					},
				},
				{
					ResourceName:      "github_user_ssh_signing_key.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}

func newTestSigningKey() string {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	publicKey, _ := ssh.NewPublicKey(&privateKey.PublicKey)
	return strings.TrimRight(string(ssh.MarshalAuthorizedKey(publicKey)), "\n")
}
