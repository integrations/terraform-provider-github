package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryAutolinkReference() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryAutolinkReferenceCreate,
		UpdateContext: resourceGithubRepositoryAutolinkReferenceUpdate,
		ReadContext:   resourceGithubRepositoryAutolinkReferenceRead,
		DeleteContext: resourceGithubRepositoryAutolinkReferenceDelete,

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
				parts := strings.Split(d.Id(), "/")
				if len(parts) != 2 {
					return nil, fmt.Errorf("invalid ID specified: supplied ID must be written as <repository>/<autolink_reference_id>")
				}

				repository := parts[0]
				id := parts[1]

				client := meta.(*Owner).v3client
				owner := meta.(*Owner).name

				// If the second part of the provided ID isn't an integer, assume that the
				// caller provided the key prefix for the autolink reference, and look up
				// the autolink by the key prefix.

				_, err := strconv.Atoi(id)
				if err != nil {
					autolink, err := getAutolinkByKeyPrefix(ctx, client, owner, repository, id)
					if err != nil {
						return nil, err
					}

					id = strconv.FormatInt(*autolink.ID, 10)
				}

				d.SetId(id)

				repo, _, err := client.Repositories.Get(ctx, owner, repository)
				if err != nil {
					return nil, fmt.Errorf("failed to retrieve repository %s: %w", repository, err)
				}

				if err = d.Set("repository", repository); err != nil {
					return nil, err
				}
				if err = d.Set("repository_id", int(repo.GetID())); err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},

		CustomizeDiff: diffRepository,

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository name",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the GitHub repository.",
			},
			"key_prefix": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit",
			},
			"target_url_template": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringMatch(regexp.MustCompile(`^http[s]?:\/\/[a-z0-9-.]*(:[0-9]+)?\/.*?<num>.*?$`), "must be a valid URL and contain <num> token")),
			},
			"is_alphanumeric": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubRepositoryAutolinkReferenceCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	keyPrefix := d.Get("key_prefix").(string)
	targetURLTemplate := d.Get("target_url_template").(string)
	isAlphanumeric := d.Get("is_alphanumeric").(bool)

	opts := &github.AutolinkOptions{
		KeyPrefix:      &keyPrefix,
		URLTemplate:    &targetURLTemplate,
		IsAlphanumeric: &isAlphanumeric,
	}

	autolinkRef, _, err := client.Repositories.AddAutolink(ctx, owner, repoName, opts)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(autolinkRef.GetID(), 10))

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryAutolinkReferenceRead(ctx, d, m)
}

func resourceGithubRepositoryAutolinkReferenceRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	autolinkRefID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}

	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	autolinkRef, _, err := client.Repositories.GetAutolink(ctx, owner, repoName, autolinkRefID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Autolink reference not found, removing from state.", map[string]any{
					"owner":      owner,
					"repository": repoName,
				})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(autolinkRef.GetID(), 10))
	if err = d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("key_prefix", autolinkRef.KeyPrefix); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("target_url_template", autolinkRef.URLTemplate); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("is_alphanumeric", autolinkRef.IsAlphanumeric); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryAutolinkReferenceUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	tflog.Warn(ctx, "Update function of autolink reference. This should not be called. But it's necessary when 'repository' doesn't have `ForceNew`", map[string]any{
		"repository":    d.Get("repository"),
		"repository_id": d.Get("repository_id"),
		"id":            d.Id(),
	})
	return nil
}

func resourceGithubRepositoryAutolinkReferenceDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	autolinkRefID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}

	_, err = client.Repositories.DeleteAutolink(ctx, owner, repoName, autolinkRefID)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
