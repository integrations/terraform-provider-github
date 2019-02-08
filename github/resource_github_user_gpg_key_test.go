package github

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubUserGpgKey_basic(t *testing.T) {
	var key github.GPGKey
	keyRe := regexp.MustCompile("^-----BEGIN PGP PUBLIC KEY BLOCK-----")
	pubKeyPath := filepath.Join("test-fixtures", "gpg-pubkey.asc")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubUserGpgKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubUserGpgKeyConfig(pubKeyPath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubUserGpgKeyExists("github_user_gpg_key.test", &key),
					resource.TestMatchResourceAttr("github_user_gpg_key.test", "armored_public_key", keyRe),
					resource.TestCheckResourceAttr("github_user_gpg_key.test", "key_id", "AC541D2D1709CD33"),
				),
			},
		},
	})
}

func testAccCheckGithubUserGpgKeyExists(n string, key *github.GPGKey) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return unconvertibleIdErr(rs.Primary.ID, err)
		}

		org := testAccProvider.Meta().(*Organization)
		receivedKey, _, err := org.client.Users.GetGPGKey(context.TODO(), id)
		if err != nil {
			return err
		}
		*key = *receivedKey
		return nil
	}
}

func testAccCheckGithubUserGpgKeyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_user_gpg_key" {
			continue
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return unconvertibleIdErr(rs.Primary.ID, err)
		}

		_, resp, err := conn.Users.GetGPGKey(context.TODO(), id)
		if err == nil {
			return fmt.Errorf("GPG key %s still exists", rs.Primary.ID)
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccGithubUserGpgKeyConfig(pubKeyPath string) string {
	return fmt.Sprintf(`
resource "github_user_gpg_key" "test" {
  armored_public_key = "${file("%s")}"
}
`, pubKeyPath)
}
