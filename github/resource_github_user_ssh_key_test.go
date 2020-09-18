package github

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"golang.org/x/crypto/ssh"
)

func TestAccGithubUserSshKey(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	testKey := newTestKey()

	t.Run("creates and destroys a user SSH key without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_user_ssh_key" "test" {
				title = "tf-acc-test-%s"
				key   = "%s"
			}
		`, randomID, testKey)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_user_ssh_key.test", "title",
				regexp.MustCompile(randomID),
			),
			resource.TestMatchResourceAttr(
				"github_user_ssh_key.test", "key",
				regexp.MustCompile("^ssh-rsa "),
			),
			resource.TestMatchResourceAttr(
				"github_user_ssh_key.test", "url",
				regexp.MustCompile("^https://api.github.com/[a-z0-9]+/keys/"),
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("imports an individual account SSH key without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_user_ssh_key" "test" {
				title = "tf-acc-test-%s"
				key   = "%s"
			}
		`, randomID, testKey)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_user_ssh_key.test", "title"),
			resource.TestCheckResourceAttrSet("github_user_ssh_key.test", "key"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:      "github_user_ssh_key.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})
}

func newTestKey() string {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	publicKey, _ := ssh.NewPublicKey(&privateKey.PublicKey)
	testKey := strings.TrimRight(string(ssh.MarshalAuthorizedKey(publicKey)), "\n")
	return testKey
}
