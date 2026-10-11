package github

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryDeployKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryDeployKeyCreate,
		ReadContext:   resourceGithubRepositoryDeployKeyRead,
		DeleteContext: resourceGithubRepositoryDeployKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: diffETag,

		// Deploy keys are defined immutable in the API. Updating results in force new.
		Schema: map[string]*schema.Schema{
			"key": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: suppressDeployKeyDiff,
				Description:      "A SSH key.",
			},
			"read_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "A boolean qualifying the key to be either read only or read/write.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the GitHub repository.",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "A title.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the deploy key.",
			},
		},
	}
}

func resourceGithubRepositoryDeployKeyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	repoName := d.Get("repository").(string)
	key := d.Get("key").(string)
	title := d.Get("title").(string)
	readOnly := d.Get("read_only").(bool)
	owner := meta.(*Owner).name

	resultKey, _, err := client.Repositories.CreateKey(ctx, owner, repoName, github.CreateDeployKeyRequest{
		Key:      key,
		Title:    new(title),
		ReadOnly: new(readOnly),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	id := strconv.FormatInt(resultKey.GetID(), 10)

	d.SetId(buildTwoPartID(repoName, id))

	return resourceGithubRepositoryDeployKeyRead(ctx, d, meta)
}

func resourceGithubRepositoryDeployKeyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName, idString, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(idString, err))
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	key, resp, err := client.Repositories.GetKey(ctx, owner, repoName, id)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing repository deploy key from state because it no longer exists in GitHub", map[string]any{"resource_id": d.Id()})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("key", key.GetKey()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("read_only", key.GetReadOnly()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("title", key.GetTitle()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryDeployKeyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName, idString, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(idString, err))
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())

	_, err = client.Repositories.DeleteKey(ctx, owner, repoName, id)
	return diag.FromErr(handleArchivedRepoDelete(ctx, err, "repository deploy key", idString, owner, repoName))
}

func suppressDeployKeyDiff(k, oldV, newV string, d *schema.ResourceData) bool {
	newV = strings.TrimSpace(newV)
	keyRe := regexp.MustCompile(`^([a-z0-9-]+ [^\s]+)( [^\s]+)?$`)
	newTrimmed := keyRe.ReplaceAllString(newV, "$1")

	return oldV == newTrimmed
}
