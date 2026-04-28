package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/google/go-github/v85/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubIssueLabel() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubIssueLabelCreateOrUpdate,
		ReadContext:   resourceGithubIssueLabelRead,
		UpdateContext: resourceGithubIssueLabelCreateOrUpdate,
		DeleteContext: resourceGithubIssueLabelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub repository.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the label.",
			},
			"color": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A 6 character hex code, without the leading '#', identifying the color of the label.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A short description of the label.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to the issue label.",
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					return true
				},
				DiffSuppressOnRefresh: true,
			},
		},
	}
}

// resourceGithubIssueLabelCreateOrUpdate idempotently creates or updates an
// issue label. Issue labels are keyed off of their "name", so pre-existing
// issue labels result in a 422 HTTP error if they exist outside of Terraform.
// Normally this would not be an issue, except new repositories are created with
// a "default" set of labels, and those labels easily conflict with custom ones.
//
// This function will first check if the label exists, and then issue an update,
// otherwise it will create. This is also advantageous in that we get to use the
// same function for two schema funcs.

func resourceGithubIssueLabelCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	name := d.Get("name").(string)
	color := d.Get("color").(string)

	label := &github.Label{
		Name:  new(name),
		Color: new(color),
	}
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	// Pull out the original name. If we already have a resource, this is the
	// parsed ID. If not, it's the value given to the resource.
	var originalName string
	if d.Id() == "" {
		originalName = name
	} else {
		var err error
		_, originalName, err = parseTwoPartID(d.Id(), "repository", "name")
		if err != nil {
			return diag.FromErr(err)
		}
	}

	existing, resp, err := client.Issues.GetLabel(ctx,
		orgName, repoName, originalName)
	if err != nil && resp.StatusCode != http.StatusNotFound {
		return diag.FromErr(err)
	}

	if existing != nil {
		label.Description = new(d.Get("description").(string))

		// Pull out the original name. If we already have a resource, this is the
		// parsed ID. If not, it's the value given to the resource.
		var originalName string
		if d.Id() == "" {
			originalName = name
		} else {
			var err error
			_, originalName, err = parseTwoPartID(d.Id(), "repository", "name")
			if err != nil {
				return diag.FromErr(err)
			}
		}

		_, _, err := client.Issues.EditLabel(ctx,
			orgName, repoName, originalName, label)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		if v, ok := d.GetOk("description"); ok {
			label.Description = new(v.(string))
		}

		_, _, err := client.Issues.CreateLabel(ctx,
			orgName, repoName, label)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(buildTwoPartID(repoName, name))

	return resourceGithubIssueLabelRead(ctx, d, meta)
}

func resourceGithubIssueLabelRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	repoName, name, err := parseTwoPartID(d.Id(), "repository", "name")
	if err != nil {
		return diag.FromErr(err)
	}

	orgName := meta.(*Owner).name
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	githubLabel, resp, err := client.Issues.GetLabel(ctx,
		orgName, repoName, name)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing label from state because it no longer exists in GitHub", map[string]any{
					"name":      name,
					"org_name":  orgName,
					"repo_name": repoName,
				})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("name", name); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("color", githubLabel.GetColor()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("description", githubLabel.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("url", githubLabel.GetURL()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubIssueLabelDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	name := d.Get("name").(string)
	ctx = context.WithValue(ctx, ctxId, d.Id())

	_, err := client.Issues.DeleteLabel(ctx, orgName, repoName, name)
	return diag.FromErr(handleArchivedRepoDelete(err, "issue label", name, orgName, repoName))
}
