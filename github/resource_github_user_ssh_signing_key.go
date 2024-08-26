package github

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubUserSshSigningKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubUserSshSigningKeyCreate,
		ReadContext:   resourceGithubUserSshSigningKeyRead,
		DeleteContext: resourceGithubUserSshSigningKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "A descriptive name for the new key.",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The public SSH key to add to your GitHub account.",
				DiffSuppressFunc: func(k, oldV, newV string, d *schema.ResourceData) bool {
					newTrimmed := strings.TrimSpace(newV)
					return oldV == newTrimmed
				},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubUserSshSigningKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Owner).v3client

	title := d.Get("title").(string)
	key := d.Get("key").(string)

	userKey, resp, err := client.Users.CreateSSHSigningKey(ctx, &github.Key{
		Title: github.Ptr(title),
		Key:   github.Ptr(key),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(*userKey.ID, 10))

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("title", userKey.GetTitle()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("key", userKey.GetKey()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubUserSshSigningKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Owner).v3client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("failed to convert ID %s: %v", d.Id(), err)
	}
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	key, resp, err := client.Users.GetSSHSigningKey(ctx, id)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, fmt.Sprintf("Removing user SSH key %s from state because it no longer exists in GitHub", d.Id()), map[string]any{
					"ssh_signing_key_id": d.Id(),
				})
				d.SetId("")
				return nil
			}
		}
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("title", key.GetTitle()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("key", key.GetKey()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubUserSshSigningKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Owner).v3client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("failed to convert ID %s: %v", d.Id(), err)
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())

	resp, err := client.Users.DeleteSSHSigningKey(ctx, id)
	if resp.StatusCode == http.StatusNotFound {
		return nil
	}
	return diag.FromErr(err)
}
