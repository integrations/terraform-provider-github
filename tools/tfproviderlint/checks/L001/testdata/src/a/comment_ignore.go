package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore() {
	// lintignore:L001
	_ = schema.Schema{
		Type: schema.TypeMap,
		ValidateFunc: func(v any, k string) (ws []string, es []error) {
			return nil, nil
		},
	}

	// lintignore:L001
	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeMap,
			ValidateFunc: func(v any, k string) (ws []string, es []error) {
				return nil, nil
			},
		},
	}
}
