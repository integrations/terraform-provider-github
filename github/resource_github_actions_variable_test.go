package github

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsVariable(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		varName := "test"
		value := "foo"

		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_variable" "test" {
	repository    = github_repository.test.name
	variable_name = "%s"
	value         = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, varName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_variable.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_variable.test", "value", value),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		varName := "test"
		value := "foo"
		valueUpdated := "bar"

		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_variable" "test" {
	repository    = github_repository.test.name
	variable_name = "%s"
	value         = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, varName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_variable.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_variable.test", "value", value),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, varName, valueUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_variable.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_variable.test", "value", valueUpdated),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("update_renamed_repo", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		updatedRepoName := fmt.Sprintf("%s%s-updated", testResourcePrefix, randomID)

		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_variable" "test" {
	repository    = github_repository.test.name
	variable_name = "test"
	value         = "test"
}
`

		var beforeCreatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "updated_at"),
						func(s *terraform.State) error {
							beforeCreatedAt = s.RootModule().Resources["github_actions_variable.test"].Primary.Attributes["created_at"]
							return nil
						},
					),
				},
				{
					Config: fmt.Sprintf(config, updatedRepoName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "updated_at"),
						func(s *terraform.State) error {
							afterCreatedAt := s.RootModule().Resources["github_actions_variable.test"].Primary.Attributes["created_at"]

							if afterCreatedAt != beforeCreatedAt {
								return fmt.Errorf("expected resource to not be recreated, but created_at has been modified: %s", beforeCreatedAt)
							}
							return nil
						},
					),
				},
			},
		})
	})

	t.Run("recreate_changed_repo", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		repoName2 := fmt.Sprintf("%supdated-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository" "test2" {
	name = "%s"
}

resource "github_actions_variable" "test" {
	repository    = github_repository.test.name
	variable_name = "test_variable"
	value         = "test"
}
`, repoName, repoName2)

		configUpdated := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository" "test2" {
	name = "%s"
}

resource "github_actions_variable" "test" {
	repository    = github_repository.test2.name
	variable_name = "test_variable"
	value         = "test"
}
`, repoName, repoName2)

		var beforeCreatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "updated_at"),
						func(s *terraform.State) error {
							beforeCreatedAt = s.RootModule().Resources["github_actions_variable.test"].Primary.Attributes["created_at"]
							return nil
						},
					),
				},
				{
					Config: configUpdated,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_variable.test", "updated_at"),
						func(s *terraform.State) error {
							afterCreatedAt := s.RootModule().Resources["github_actions_variable.test"].Primary.Attributes["created_at"]

							if afterCreatedAt == beforeCreatedAt {
								return fmt.Errorf("expected resource to be recreated, but created_at has not been modified: %s", beforeCreatedAt)
							}
							return nil
						},
					),
				},
			},
		})
	})

	t.Run("destroy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_variable" "test" {
	repository    = github_repository.test.name
	variable_name = "test"
	value         = "foo"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_variable" "test" {
	repository    = github_repository.test.name
	variable_name = "test"
	value         = "foo"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      "github_actions_variable.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("error_on_existing", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		varName := "test"

		baseConfig := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}
`, repoName)

		config := fmt.Sprintf(`
%s

resource "github_actions_variable" "test" {
	repository    = github_repository.test.name
	variable_name = "%s"
	value         = "test"
}
`, baseConfig, varName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: baseConfig,
					Check: func(*terraform.State) error {
						meta, err := getTestMeta()
						if err != nil {
							return err
						}
						client := meta.v3client
						owner := meta.name
						ctx := context.Background()

						_, err = client.Actions.CreateRepoVariable(ctx, owner, repoName, &github.ActionsVariable{
							Name:  varName,
							Value: "test",
						})
						return err
					},
				},
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Variable already exists`),
					Check: func(*terraform.State) error {
						meta, err := getTestMeta()
						if err != nil {
							return err
						}
						client := meta.v3client
						owner := meta.name
						ctx := context.Background()

						_, err = client.Actions.DeleteRepoVariable(ctx, owner, repoName, varName)
						return err
					},
				},
			},
		})
	})
}
