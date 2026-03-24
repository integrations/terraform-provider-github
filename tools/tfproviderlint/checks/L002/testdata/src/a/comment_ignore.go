package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore() {
	_ = schema.Resource{
		// lintignore:L002
		Create: func(d *schema.ResourceData, meta any) error {
			return nil
		},
	}

	_ = schema.Resource{
		// lintignore:L002
		Read: func(d *schema.ResourceData, meta any) error {
			return nil
		},
	}

	_ = schema.Resource{
		// lintignore:L002
		Update: func(d *schema.ResourceData, meta any) error {
			return nil
		},
	}

	_ = schema.Resource{
		// lintignore:L002
		Delete: func(d *schema.ResourceData, meta any) error {
			return nil
		},
	}

	_ = schema.Resource{
		Importer: &schema.ResourceImporter{
			// lintignore:L002
			State: schema.ImportStatePassthrough,
		},
	}
}
