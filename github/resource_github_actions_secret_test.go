package github

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubActionsSecret_basic(t *testing.T) {
	repo := os.Getenv("GITHUB_TEMPLATE_REPOSITORY")

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

func testAccGithubActionsSecretFullConfig(repoName, plaintext string) string {

	// Take resources from other tests to avoid manual creation of secrets / repos
	githubPKData := testAccCheckGithubActionsPublicKeyDataSourceConfig(repoName)
	githubActionsSecretResource := testAccGithubActionsSecretConfig(repoName, plaintext)

	return fmt.Sprintf("%s%s", githubPKData, githubActionsSecretResource)
}

func testAccGithubActionsSecretConfig(repo, plaintext string) string {
	return fmt.Sprintf(`
resource "github_actions_secret" "test_secret" {
  repository       = "%s"
  secret_name      = "test_secret_name"
  plaintext_value  = "%s"
}
`, repo, plaintext)
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

		org := testAccProvider.Meta().(*Organization)
		conn := org.v3client
		_, _, err := conn.Actions.GetSecret(context.TODO(), org.name, repoName, secretName)
		if err != nil {
			t.Log("Failed to get secret")
			return err
		}

		return nil
	}
}

func testAccCheckGithubActionsSecretDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Organization).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_actions_secret" {
			continue
		}
		owner := testAccProvider.Meta().(*Organization).name
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
