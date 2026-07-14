package github

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// checkCollectionItemAbsent checks that a collection attribute does not contain an item with a specific field value.
func checkCollectionItemAbsent(resourceName, collectionAttr, field, forbidden string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		prefix := collectionAttr + "."
		suffix := "." + field

		for k, v := range rs.Primary.Attributes {
			if !strings.HasPrefix(k, prefix) {
				continue
			}
			// skip collection metadata
			if strings.HasSuffix(k, ".#") || strings.HasSuffix(k, ".%") {
				continue
			}
			if strings.HasSuffix(k, suffix) && v == forbidden {
				return fmt.Errorf("%s contains forbidden %s=%q (key %s)", collectionAttr, field, forbidden, k)
			}
		}

		return nil
	}
}
