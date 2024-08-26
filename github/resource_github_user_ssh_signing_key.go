package github

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v83/github"
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
			StateContext: resourceGithubUserSshSigningKeyImport,
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
			},
			"key_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The unique identifier of the SSH signing key.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubUserSshSigningKeyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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

	d.SetId(strconv.FormatInt(userKey.GetID(), 10))

	if err = d.Set("key_id", userKey.GetID()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("title", userKey.GetTitle()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubUserSshSigningKeyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	keyID := d.Get("key_id").(int64)
	_, _, err := client.Users.GetSSHSigningKey(ctx, keyID)
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
	return nil
}

func resourceGithubUserSshSigningKeyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	keyID := d.Get("key_id").(int64)
	resp, err := client.Users.DeleteSSHSigningKey(ctx, keyID)
	if resp.StatusCode == http.StatusNotFound {
		return nil
	}
	return diag.FromErr(err)
}

func resourceGithubUserSshSigningKeyImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	client := meta.(*Owner).v3client

	keyID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid SSH signing key ID format: %v", err)
	}

	key, resp, err := client.Users.GetSSHSigningKey(ctx, keyID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				return nil, fmt.Errorf("SSH signing key with ID %d not found", keyID)
			}
		}
		return nil, err
	}

	d.SetId(strconv.FormatInt(key.GetID(), 10))

	if err = d.Set("key_id", key.GetID()); err != nil {
		return nil, err
	}
	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return nil, err
	}
	if err = d.Set("title", key.GetTitle()); err != nil {
		return nil, err
	}
	if err = d.Set("key", key.GetKey()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
