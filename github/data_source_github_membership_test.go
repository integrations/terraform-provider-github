package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubMembershipDataSource(t *testing.T) {
	t.Parallel()

	t.Run("queries the membership for a user in a specified organization", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			data "github_membership" "test" {
				username = "%s"
				organization = "%s"
			}
		`, testAccConf.testOrgUser1, testAccConf.owner)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_membership.test", "username", testAccConf.testOrgUser1),
			resource.TestCheckResourceAttrSet("data.github_membership.test", "role"),
			resource.TestCheckResourceAttrSet("data.github_membership.test", "etag"),
			resource.TestCheckResourceAttrSet("data.github_membership.test", "state"),
			resource.TestCheckResourceAttrSet("data.github_membership.test", "user_id"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t); skipUnlessHasOrgUser1(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("errors when querying with non-existent user", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			data "github_membership" "test" {
				username = "!%s"
				organization = "%s"
			}
		`, testAccConf.testOrgUser1, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Not Found`),
				},
			},
		})
	})

	t.Run("queries the membership for a user by user_id", func(t *testing.T) {
		ctx := t.Context()

		meta, err := getTestMeta()
		if err != nil {
			t.Fatalf("failed to get test meta: %s", err)
		}

		ghUser, _, err := meta.v3client.Users.Get(ctx, testAccConf.testOrgUser)
		if err != nil {
			t.Fatalf("failed to resolve org user id: %s", err)
		}

		config := fmt.Sprintf(`
			data "github_membership" "test" {
				user_id      = %d
				organization = "%s"
			}
		`, ghUser.GetID(), testAccConf.owner)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_membership.test", "username", testAccConf.testOrgUser),
			resource.TestCheckResourceAttr("data.github_membership.test", "user_id", fmt.Sprintf("%d", ghUser.GetID())),
			resource.TestCheckResourceAttrSet("data.github_membership.test", "role"),
			resource.TestCheckResourceAttrSet("data.github_membership.test", "etag"),
			resource.TestCheckResourceAttrSet("data.github_membership.test", "state"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("errors when querying with non-existent user_id", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_membership" "test" {
				user_id      = 999999999999
				organization = "%s"
			}
		`, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
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
		config := fmt.Sprintf(`
			data "github_membership" "test" {
				organization = "%s"
			}
		`, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
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
			data "github_membership" "test" {
				username     = "%s"
				user_id      = 1
				organization = "%s"
			}
		`, testAccConf.testOrgUser, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
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
