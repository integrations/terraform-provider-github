package github

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubUserSshKey_basic(t *testing.T) {
	var key github.Key
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	title := fmt.Sprintf("tf-acc-test-%s", randString)
	keyRe := regexp.MustCompile("^ecdsa-sha2-nistp384 ")
	urlRe := regexp.MustCompile("^https://api.github.com/[a-z0-9]+/keys/")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubUserSshKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubUserSshKeyConfig(title),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubUserSshKeyExists("github_user_ssh_key.test", &key),
					resource.TestCheckResourceAttr("github_user_ssh_key.test", "title", title),
					resource.TestMatchResourceAttr("github_user_ssh_key.test", "key", keyRe),
					resource.TestMatchResourceAttr("github_user_ssh_key.test", "url", urlRe),
				),
			},
		},
	})
}

func TestAccGithubUserSshKey_importBasic(t *testing.T) {
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	title := fmt.Sprintf("tf-acc-test-%s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubUserSshKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubUserSshKeyConfig(title),
			},
			{
				ResourceName:      "github_user_ssh_key.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckGithubUserSshKeyExists(n string, key *github.Key) resource.TestCheckFunc {
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
		receivedKey, _, err := org.client.Users.GetKey(context.TODO(), id)
		if err != nil {
			return err
		}
		*key = *receivedKey
		return nil
	}
}

func testAccCheckGithubUserSshKeyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_user_ssh_key" {
			continue
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return unconvertibleIdErr(rs.Primary.ID, err)
		}

		_, resp, err := conn.Users.GetKey(context.TODO(), id)
		if err == nil {
			return fmt.Errorf("SSH key %s still exists", rs.Primary.ID)
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccGithubUserSshKeyConfig(title string) string {
	return fmt.Sprintf(`
resource "github_user_ssh_key" "test" {
  title = "%s"
  key = "${tls_private_key.test.public_key_openssh}"
}

resource "tls_private_key" "test" {
  algorithm   = "ECDSA"
  ecdsa_curve = "P384"
}
`, title)
}
