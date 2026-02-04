package github

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGithubActionsOrganizationVariable(t *testing.T) {
	t.Run("create_update_visibility_all", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		varName := fmt.Sprintf("test_%s", randomID)
		value := "foo"
		valueUpdated := "bar"

		config := `
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "%s"
	visibility    = "all"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, varName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "value", value),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, varName, valueUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "value", valueUpdated),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_visibility_private", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		varName := fmt.Sprintf("test_%s", randomID)
		value := "foo"
		valueUpdated := "bar"

		config := `
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "%s"
	visibility    = "private"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, varName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "value", value),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "visibility", "private"),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, varName, valueUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "value", valueUpdated),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "visibility", "private"),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_visibility_selected", func(t *testing.T) {
		repoName0 := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		repoName1 := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		varName := fmt.Sprintf("test_%s", randomID)
		value := "foo"
		valueUpdated := "bar"

		config := `
resource "github_repository" "test_0" {
	name = "%s"
}

resource "github_repository" "test_1" {
	name = "%s"
}

resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "%s"
	visibility    = "selected"

	selected_repository_ids = [github_repository.test_%s.repo_id]
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName0, repoName1, varName, value, "0"),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "value", value),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "visibility", "selected"),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "selected_repository_ids.#", "1"),
						resource.TestCheckResourceAttrPair("github_actions_organization_variable.test", "selected_repository_ids.0", "github_repository.test_0", "repo_id"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName0, repoName1, varName, valueUpdated, "1"),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "value", valueUpdated),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "visibility", "selected"),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "selected_repository_ids.#", "1"),
						resource.TestCheckResourceAttrPair("github_actions_organization_variable.test", "selected_repository_ids.0", "github_repository.test_1", "repo_id"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_change_visibility", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		varName := fmt.Sprintf("test_%s", randomID)
		value := "foo"
		visibility := "all"
		valueUpdated := "bar"
		visibilityUpdated := "private"

		config := `
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "%s"
	visibility    = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, varName, value, visibility),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "value", value),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "visibility", visibility),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, varName, valueUpdated, visibilityUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "variable_name", varName),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "value", valueUpdated),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "visibility", visibilityUpdated),
						resource.TestCheckResourceAttr("github_actions_organization_variable.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_organization_variable.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("destroy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		varName := fmt.Sprintf("test_%s", randomID)

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "foo"
	visibility    = "all"
}
`, varName)

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
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		varName := fmt.Sprintf("test_%s", randomID)

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "foo"
	visibility    = "all"
}
`, varName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      "github_actions_organization_variable.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("error_on_invalid_selected_repository_ids", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		varName := fmt.Sprintf("test_%s", randomID)

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "foo"
	visibility    = "all"

	selected_repository_ids = [123456]
}
`, varName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("cannot use selected_repository_ids without visibility being set to selected"),
				},
			},
		})
	})

	t.Run("error_on_existing", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		varName := fmt.Sprintf("test_%s", randomID)

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "foo"
	visibility    = "all"
}
`, varName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
`,
					Check: func(*terraform.State) error {
						meta, err := getTestMeta()
						if err != nil {
							return err
						}
						client := meta.v3client
						owner := meta.name
						ctx := context.Background()

						_, err = client.Actions.CreateOrgVariable(ctx, owner, &github.ActionsVariable{
							Name:       varName,
							Value:      "test",
							Visibility: github.Ptr("all"),
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

						_, err = client.Actions.DeleteOrgVariable(ctx, owner, varName)
						return err
					},
				},
			},
		})
	})
}
