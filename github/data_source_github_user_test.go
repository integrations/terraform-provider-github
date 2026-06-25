package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubUserDataSource(t *testing.T) {
	if len(testAccConf.testExternalUser1) == 0 {
		t.Skip("No external user provided")
	}

	t.Run("queries an existing individual account without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_user" "test" {
				username = "%s"
			}
		`, testAccConf.testExternalUser1)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_user.test", "login"),
			resource.TestCheckResourceAttrSet("data.github_user.test", "id"),
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

	t.Run("errors when querying a non-existing individual account", func(t *testing.T) {
		config := fmt.Sprintf(`
				data "github_user" "test" {
					username = "!%s"
				}
			`, testAccConf.testExternalUser1)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Not Found`),
				},
			},
		})
	})

	t.Run("queries an existing individual account by user_id", func(t *testing.T) {
		ctx := t.Context()

		meta, err := getTestMeta()
		if err != nil {
			t.Fatalf("failed to get test meta: %s", err)
		}

		ghUser, _, err := meta.v3client.Users.Get(ctx, testAccConf.testExternalUser)
		if err != nil {
			t.Fatalf("failed to resolve external user id: %s", err)
		}

		config := fmt.Sprintf(`
			data "github_user" "test" {
				user_id = %d
			}
		`, ghUser.GetID())

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_user.test", "login", testAccConf.testExternalUser),
			resource.TestCheckResourceAttr("data.github_user.test", "username", testAccConf.testExternalUser),
			resource.TestCheckResourceAttr("data.github_user.test", "id", fmt.Sprintf("%d", ghUser.GetID())),
			resource.TestCheckResourceAttr("data.github_user.test", "user_id", fmt.Sprintf("%d", ghUser.GetID())),
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

	t.Run("errors when querying a non-existing user_id", func(t *testing.T) {
		config := `
			data "github_user" "test" {
				user_id = 999999999999
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Not Found`),
				},
			},
		})
	})

	t.Run("errors when neither username nor user_id is provided", func(t *testing.T) {
		config := `
			data "github_user" "test" {}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`one of (\x60username\x60,\x60user_id\x60|\x60user_id\x60,\x60username\x60) must be specified`),
				},
			},
		})
	})

	t.Run("errors when both username and user_id are provided", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_user" "test" {
				username = "%s"
				user_id  = 1
			}
		`, testAccConf.testExternalUser)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`only one of (\x60user_id\x60,\x60username\x60|\x60username\x60,\x60user_id\x60) can be specified`),
				},
			},
		})
	})
}
