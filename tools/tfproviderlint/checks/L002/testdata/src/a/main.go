package a

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func f() {
	// Test for CreateContext
	_ = schema.Resource{
		Create: func(d *schema.ResourceData, meta any) error { // want "L002: resource should not configure Create, replace it with CreateContext"
			return nil
		},
	}

	// Test for ReadContext
	_ = schema.Resource{
		Read: func(d *schema.ResourceData, meta any) error { // want "L002: resource should not configure Read, replace it with ReadContext"
			return nil
		},
	}

	// Test for UpdateContext
	_ = schema.Resource{
		Update: func(d *schema.ResourceData, meta any) error { // want "L002: resource should not configure Update, replace it with UpdateContext"
			return nil
		},
	}

	// Test for DeleteContext
	_ = schema.Resource{
		Delete: func(d *schema.ResourceData, meta any) error { // want "L002: resource should not configure Delete, replace it with DeleteContext"
			return nil
		},
	}

	// Test for Importer.StateContext
	_ = schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough, // want "L002: resource should not configure Importer.State, replace it with Importer.StateContext"
		},
	}

	_ = schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
			return nil
		},
	}

	_ = schema.Resource{
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
			return nil
		},
	}

	_ = schema.Resource{
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
			return nil
		},
	}

	_ = schema.Resource{
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
