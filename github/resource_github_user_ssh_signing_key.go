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

func resourceGithubUserSshSigningKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubUserSshSigningKeyCreate,
		ReadContext:   resourceGithubUserSshSigningKeyRead,
		DeleteContext: resourceGithubUserSshSigningKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubUserSshSigningKeyImport,
		},

		Description: "Resource to manage a SSH signing key for the authenticated user.",

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
				Description: "The public SSH signing key to add to your GitHub account.",
			},
			"key_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The unique identifier of the SSH signing key.",
			},
			"etag": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "An etag representing the SSH signing key.",
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					return true
				},
				DiffSuppressOnRefresh: true,
			},
		},
	}
}

func resourceGithubUserSshSigningKeyCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	title, _ := d.Get("title").(string)
	key, _ := d.Get("key").(string)

	userKey, resp, err := client.Users.CreateSSHSigningKey(ctx, &github.Key{
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
	if err = d.Set("title", userKey.GetTitle()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubUserSshSigningKeyRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	keyIDInt, _ := d.Get("key_id").(int)
	keyID := int64(keyIDInt)

	_, resp, err := client.Users.GetSSHSigningKey(ctx, keyID)
	if err != nil {
		return diag.FromErr(deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "user SSH signing key (%d)", keyID))
	}

	// set computed fields
	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubUserSshSigningKeyDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	keyIDInt, _ := d.Get("key_id").(int)
	keyID := int64(keyIDInt)

	_, err := client.Users.DeleteSSHSigningKey(ctx, keyID)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubUserSshSigningKeyImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client

	keyID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid SSH signing key ID format: %w", err)
	}

	key, _, err := client.Users.GetSSHSigningKey(ctx, keyID)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("SSH signing key with ID %d not found", keyID)
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
