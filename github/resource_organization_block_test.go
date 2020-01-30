package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccOrganizationBlock_basic(t *testing.T) {
	rn := "github_organization_block.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccOrganizationBlockDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationBlockConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOrganizationBlockExists(rn),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccOrganizationBlockDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client
	orgName := testAccProvider.Meta().(*Organization).name

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_organization_block" {
			continue
		}

		username := rs.Primary.ID

		res, err := conn.Organizations.UnblockUser(context.TODO(), orgName, username)
		if res.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccCheckOrganizationBlockExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		username := rs.Primary.ID
		conn := testAccProvider.Meta().(*Organization).v3client
		orgName := testAccProvider.Meta().(*Organization).name

		blocked, _, err := conn.Organizations.IsBlocked(context.TODO(), orgName, username)
		if err != nil {
			return err
		}
		if !blocked {
			return fmt.Errorf("not blocked: %s %s", orgName, username)
		}
		return nil
	}
}

const testAccOrganizationBlockConfig = `
resource "github_organization_block" "test" {
  username = "cgriggs01"
}
`
