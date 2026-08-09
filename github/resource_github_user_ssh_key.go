package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubUserSshKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubUserSshKeyCreate,
		ReadContext:   resourceGithubUserSshKeyRead,
		DeleteContext: resourceGithubUserSshKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubUserSshKeyImport,
		},

		Description: "Manages a SSH key for the authenticated user.",

		SchemaVersion: 1,
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
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the SSH key.",
			},
			"key_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The unique identifier of the SSH key.",
			},
			"etag": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "An etag representing the SSH key.",
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					return true
				},
				DiffSuppressOnRefresh: true,
			},
		},

		StateUpgraders: []schema.StateUpgrader{
			{
				Version: 0,
				Type:    resourceGithubUserSshKeyV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubUserSshKeyStateUpgradeV0,
			},
		},
	}
}

func resourceGithubUserSshKeyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	title := d.Get("title").(string)
	key := d.Get("key").(string)

	userKey, resp, err := client.Users.CreateKey(ctx, &github.Key{
		Title: new(title),
		Key:   new(key),
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
	if err = d.Set("url", userKey.GetURL()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("title", userKey.GetTitle()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubUserSshKeyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	keyID := int64(d.Get("key_id").(int))
	userKey, resp, err := client.Users.GetKey(ctx, keyID)
	if err != nil {
		return diag.FromErr(deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "user SSH key (%d)", keyID))
	}

	// set computed fields
	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("url", userKey.GetURL()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubUserSshKeyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	keyID := int64(d.Get("key_id").(int))

	_, err := client.Users.DeleteKey(ctx, keyID)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubUserSshKeyImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	client := meta.(*Owner).v3client

	keyID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid SSH key ID format: %w", err)
	}

	key, _, err := client.Users.GetKey(ctx, keyID)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("SSH key with ID %d not found", keyID)
		}
		return nil, err
	}

	if err = d.Set("key_id", key.GetID()); err != nil {
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
