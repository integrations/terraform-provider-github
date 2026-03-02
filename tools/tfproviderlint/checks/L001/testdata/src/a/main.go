package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "L001: schema should not configure ValidateFunc, replace it with ValidateDiagFunc"
		Type: schema.TypeMap,
		ValidateFunc: func(v any, k string) (ws []string, es []error) {
			return nil, nil
		},
	}

	_ = schema.Schema{
		Type: schema.TypeMap,
	}

	_ = map[string]*schema.Schema{
		"name": { // want "L001: schema should not configure ValidateFunc, replace it with ValidateDiagFunc"
			Type: schema.TypeMap,
			ValidateFunc: func(v any, k string) (ws []string, es []error) {
				return nil, nil
			},
		},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Type: schema.TypeMap,
		},
	}
}
