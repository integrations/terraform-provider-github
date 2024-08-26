package github

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_user_ssh_signing_key.test",
				"title",
				regexp.MustCompile(randomID),
			),
			resource.TestMatchResourceAttr(
				"github_user_ssh_signing_key.test",
				"key",
				regexp.MustCompile("^ssh-rsa "),
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
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

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_user_ssh_signing_key.test", "title"),
			resource.TestCheckResourceAttrSet("github_user_ssh_signing_key.test", "key"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
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
