package a

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ignoredCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func ignoredRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func fcommentignore() {
	_ = schema.Resource{
		CreateContext: ignoredCreate,
		ReadContext:   ignoredRead,
	}

	// lintignore:L003
	ignoredRead(nil, nil, nil)
}
