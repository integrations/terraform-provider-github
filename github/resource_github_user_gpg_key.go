package github

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubUserGpgKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubUserGpgKeyCreate,
		Read:   resourceGithubUserGpgKeyRead,
		Delete: resourceGithubUserGpgKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: diffETag,

		Schema: map[string]*schema.Schema{
			"armored_public_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Your public GPG key, generated in ASCII-armored format.",
				DiffSuppressFunc: func(k, oldV, newV string, d *schema.ResourceData) bool {
					return normalizeArmoredKey(oldV) == normalizeArmoredKey(newV)
				},
			},
			"key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The key ID of the GPG key.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the GPG key.",
			},
		},
	}
}

func resourceGithubUserGpgKeyCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client

	pubKey := d.Get("armored_public_key").(string)
	ctx := context.Background()

	key, _, err := client.Users.CreateGPGKey(ctx, pubKey)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(key.GetID(), 10))

	return resourceGithubUserGpgKeyRead(d, meta)
}

func resourceGithubUserGpgKeyRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	key, _, err := client.Users.GetGPGKey(ctx, id)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
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
		return err
	}

	if err = d.Set("key_id", key.GetKeyID()); err != nil {
		return err
	}

	// An imported key has no armored_public_key until it is read from GitHub,
	// which returns the key as it was uploaded.
	if d.Get("armored_public_key").(string) == "" {
		if err = d.Set("armored_public_key", key.GetRawKey()); err != nil {
			return err
		}
	}

	return nil
}

// normalizeArmoredKey ignores line endings and surrounding whitespace, which
// GitHub keeps as uploaded.
func normalizeArmoredKey(key string) string {
	return strings.ReplaceAll(strings.TrimSpace(key), "\r\n", "\n")
}

func resourceGithubUserGpgKeyDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, err = client.Users.DeleteGPGKey(ctx, id)

	return err
}
