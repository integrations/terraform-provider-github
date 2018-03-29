package github

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccGithubIpRangesDataSource_existing(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
				data "github_ip_ranges" "test" {}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "hooks.#"),
					resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "git.#"),
					resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "pages.#"),
				),
			},
		},
	})
}
