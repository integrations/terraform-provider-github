package github

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryDeployKeysDataSource(t *testing.T) {
	t.Run("manages deploy keys", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		keyName := fmt.Sprintf("%s_rsa", randomID)
		cmd := exec.Command("bash", "-c", fmt.Sprintf("ssh-keygen -t rsa -b 4096 -C test@example.com -N '' -f test-fixtures/%s>/dev/null <<< y >/dev/null", keyName))
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		keyPath := strings.ReplaceAll(filepath.Join("test-fixtures", fmt.Sprintf("%s.pub", keyName)), "\\", "/")

		repoName := fmt.Sprintf("tf-acc-test-deploy-keys-%s", acctest.RandString(5))
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name      = "%s"
			auto_init = true
		}

		resource "github_repository_deploy_key" "test" {
			repository = github_repository.test.name
			key        = file("%s")
			title      = "title1"
			read_only  = true
		}

		data "github_repository_deploy_keys" "test" {
			repository = github_repository.test.name

			depends_on = [github_repository_deploy_key.test]
		}
		`, repoName, keyPath)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_repository_deploy_keys.test", "keys.#"),
						resource.TestCheckResourceAttr("data.github_repository_deploy_keys.test", "keys.#", "1"),
						resource.TestCheckResourceAttr("data.github_repository_deploy_keys.test", "keys.0.title", "title1"),
						resource.TestCheckResourceAttrSet("data.github_repository_deploy_keys.test", "keys.0.id"),
						resource.TestCheckResourceAttr("data.github_repository_deploy_keys.test", "keys.0.verified", "true"),
					),
				},
			},
		})
	})
}
