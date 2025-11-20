package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseDataSource(t *testing.T) {
	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}

	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}

	config := fmt.Sprintf(`
			data "github_enterprise" "test" {
				slug = "%s"
			}
		`,
		testEnterprise,
	)

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("data.github_enterprise.test", "slug", testEnterprise),
		resource.TestCheckResourceAttrSet("data.github_enterprise.test", "name"),
		resource.TestCheckResourceAttrSet("data.github_enterprise.test", "created_at"),
		resource.TestCheckResourceAttrSet("data.github_enterprise.test", "url"),
	)

	resource.Test(
		t,
		resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, enterprise) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		},
	)
}
