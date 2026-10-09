package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrganizationBlock() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrganizationBlockCreate,
		ReadContext:   resourceOrganizationBlockRead,
		DeleteContext: resourceOrganizationBlockDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: diffETag,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the user to block.",
			},

			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the organization block.",
			},
		},
	}
}

func resourceOrganizationBlockCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	username := d.Get("username").(string)

	_, err = client.Organizations.BlockUser(ctx, orgName, username)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(username)

	return resourceOrganizationBlockRead(ctx, d, meta)
}

func resourceOrganizationBlockRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	username := d.Id()

	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	blocked, resp, err := client.Organizations.IsBlocked(ctx, orgName, username)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			// not sure if this will ever be hit, I imagine just returns false?
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing organization block from state because it no longer exists in GitHub", map[string]any{"owner": orgName, "resource_id": d.Id()})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if !blocked {
		d.SetId("")
		return nil
	}

	if err = d.Set("username", username); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceOrganizationBlockDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	username := d.Id()
	ctx = context.WithValue(ctx, ctxId, d.Id())

	_, err := client.Organizations.UnblockUser(ctx, orgName, username)
	return diag.FromErr(err)
}
