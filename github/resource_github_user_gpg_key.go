package github

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubUserGpgKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubUserGpgKeyCreate,
		ReadContext:   resourceGithubUserGpgKeyRead,
		DeleteContext: resourceGithubUserGpgKeyDelete,

		Schema: map[string]*schema.Schema{
			"armored_public_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Your public GPG key, generated in ASCII-armored format.",
			},
			"key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The key ID of the GPG key.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubUserGpgKeyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	pubKey := d.Get("armored_public_key").(string)

	key, _, err := client.Users.CreateGPGKey(ctx, pubKey)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(key.GetID(), 10))

	return resourceGithubUserGpgKeyRead(ctx, d, meta)
}

func resourceGithubUserGpgKeyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}

	key, _, err := client.Users.GetGPGKey(ctx, id)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing user GPG key %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("key_id", key.GetKeyID()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubUserGpgKeyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}

	_, err = client.Users.DeleteGPGKey(ctx, id)

	return diag.FromErr(err)
}
