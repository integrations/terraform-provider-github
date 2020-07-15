package github

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubActionsSecret_basic(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	repo := acctest.RandomWithPrefix("tf-acc-test")

	secretResourceName := "github_actions_secret.test_secret"
	secretValue := "super_secret_value"
	updatedSecretValue := "updated_super_secret_value"
	t.Log(testAccGithubActionsSecretFullConfig(repo, secretValue))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubActionsSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubActionsSecretFullConfig(repo, secretValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubActionsSecretExists(secretResourceName, "test_secret_name", t),
					resource.TestCheckResourceAttr("github_actions_secret.test_secret", "plaintext_value", secretValue),
					resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "created_at"),
					resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "updated_at"),
				),
			},
			{
				Config: testAccGithubActionsSecretFullConfig(repo, updatedSecretValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubActionsSecretExists(secretResourceName, "test_secret_name", t),
					resource.TestCheckResourceAttr("github_actions_secret.test_secret", "plaintext_value", updatedSecretValue),
					resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "created_at"),
					resource.TestCheckResourceAttrSet("github_actions_secret.test_secret", "updated_at"),
				),
			},
		},
	})
}

func TestAccGithubActionsSecret_disappears(t *testing.T) {
	repo := acctest.RandomWithPrefix("tf-acc-test")
	secretResourceName := "github_actions_secret.test_secret"
	secretValue := "super_secret_value"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubActionsSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubActionsSecretFullConfig(repo, secretValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubActionsSecretExists(secretResourceName, "test_secret_name", t),
					testAccCheckGithubActionsSecretDisappears(secretResourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGithubActionsSecretFullConfig(repoName, plaintext string) string {
	// To allow tests to run in parallel and prevent re-using resources defined across the
	// codebase, we create a repository resource and define it's actions public key here
	// alongside the new actions secret resource
	return fmt.Sprintf(`
data "github_actions_public_key" "test_pk" {
  repository = github_repository.test.name
}

resource "github_repository" "test" {
  name = "%s"
}

resource "github_actions_secret" "test_secret" {
  repository       = github_repository.test.name
  secret_name      = "test_secret_name"
  plaintext_value  = "%s"
}
`, repoName, plaintext)
}

func testAccCheckGithubActionsSecretExists(resourceName, secretName string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actualResource, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not Found: %s", resourceName)
		}

		repoName := actualResource.Primary.Attributes["repository"]
		if repoName == "" {
			return fmt.Errorf("no repo name is set")
		}

		org := testAccProvider.Meta().(*Owner)
		conn := org.v3client
		_, _, err := conn.Actions.GetSecret(context.TODO(), org.name, repoName, secretName)
		if err != nil {
			t.Log("Failed to get secret")
			return err
		}

		return nil
	}
}

func testAccCheckGithubActionsSecretDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}
		conn := testAccProvider.Meta().(*Owner).v3client
		owner := testAccProvider.Meta().(*Owner).name
		repoName, secretName, err := parseTwoPartID(rs.Primary.ID, "repository", "secret_name")
		if err != nil {
			return err
		}
		_, err = conn.Actions.DeleteSecret(context.TODO(), owner, repoName, secretName)
		return err
	}
}

func testAccCheckGithubActionsSecretDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Owner).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_actions_secret" {
			continue
		}
		owner := testAccProvider.Meta().(*Owner).name
		repoName := rs.Primary.Attributes["repository"]
		secretName := rs.Primary.Attributes["secret_name"]

		gotSecret, resp, err := client.Actions.GetSecret(context.TODO(), owner, repoName, secretName)
		if err == nil && gotSecret != nil && gotSecret.Name == secretName {
			return fmt.Errorf("secret %s still exists", rs.Primary.ID)
		}
		if resp != nil && resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}
