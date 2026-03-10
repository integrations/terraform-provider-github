package a

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return resourceRead(ctx, d, meta) // want "L003: function resourceRead is a resource CRUD function \\(ReadContext\\) and must not be called directly"
}

func resourceRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return resourceRead(ctx, d, meta) // want "L003: function resourceRead is a resource CRUD function \\(ReadContext\\) and must not be called directly"
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func resourceImportState(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	return nil, nil
}

func helperFunc() {}

func f() {
	_ = schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceImportState,
		},
	}

	resourceCreate(nil, nil, nil)             // want "L003: function resourceCreate is a resource CRUD function \\(CreateContext\\) and must not be called directly"
	resourceRead(nil, nil, nil)               // want "L003: function resourceRead is a resource CRUD function \\(ReadContext\\) and must not be called directly"
	resourceUpdate(nil, nil, nil)             // want "L003: function resourceUpdate is a resource CRUD function \\(UpdateContext\\) and must not be called directly"
	resourceDelete(nil, nil, nil)             // want "L003: function resourceDelete is a resource CRUD function \\(DeleteContext\\) and must not be called directly"
	_, _ = resourceImportState(nil, nil, nil) // want "L003: function resourceImportState is a resource CRUD function \\(Importer.StateContext\\) and must not be called directly"

	// Calling a non-CRUD function is not a violation
	helperFunc()
}
