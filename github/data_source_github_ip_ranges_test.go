package github

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubIpRangesDataSource_basic(t *testing.T) {
	slug := "non-existing"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubTeamDataSourceConfig(slug),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubIpRangesSetsValue(),
				),
			},
		},
	})
}

func testAccCheckGithubIpRangesSetsValue() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}
