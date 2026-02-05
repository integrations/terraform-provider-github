package github

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func getResourceAttr(name, key string, target *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s", name)
		}

		attr, ok := res.Primary.Attributes[key]
		if !ok {
			return fmt.Errorf("attribute not found: %s", key)
		}

		*target = attr
		return nil
	}
}
