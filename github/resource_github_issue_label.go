package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubIssueLabel() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubIssueLabelCreate,
		ReadContext:   resourceGithubIssueLabelRead,
		UpdateContext: resourceGithubIssueLabelUpdate,
		DeleteContext: resourceGithubIssueLabelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubIssueLabelImport,
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

func resourceGithubIssueLabelCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name
	repoName, ok := d.Get("repository").(string)
	if !ok {
		return diag.Errorf(`expected "repository" to be string`)
	}
	name, ok := d.Get("name").(string)
	if !ok {
		return diag.Errorf(`expected "name" to be string`)
	}
	color, ok := d.Get("color").(string)
	if !ok {
		return diag.Errorf(`expected "color" to be string`)
	}

	label := &github.Label{
		Name:  new(name),
		Color: new(color),
	}

	description, ok := d.Get("description").(string)
	if !ok {
		return diag.Errorf(`expected "description" to be string`)
	}
	label.Description = &description
	githubLabel, resp, err := client.Issues.CreateLabel(ctx, orgName, repoName, label)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	if err := d.Set("url", githubLabel.GetURL()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubIssueLabelRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	repoName, ok := d.Get("repository").(string)
	if !ok {
		return diag.Errorf(`expected "repository" to be string`)
	}
	name, ok := d.Get("name").(string)
	if !ok {
		return diag.Errorf(`expected "name" to be string`)
	}

	orgName := meta.name
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	githubLabel, resp, err := client.Issues.GetLabel(ctx, orgName, repoName, name)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing label from state because it no longer exists in GitHub", map[string]any{"name": name, "org_name": orgName, "repo_name": repoName})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("color", githubLabel.GetColor()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("description", githubLabel.GetDescription()); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("url", githubLabel.GetURL()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubIssueLabelUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name
	repoName, ok := d.Get("repository").(string)
	if !ok {
		return diag.Errorf(`expected "repository" to be string`)
	}
	name, ok := d.Get("name").(string)
	if !ok {
		return diag.Errorf(`expected "name" to be string`)
	}
	color, ok := d.Get("color").(string)
	if !ok {
		return diag.Errorf(`expected "color" to be string`)
	}

	originalName := name
	if d.HasChange("name") {
		oldName, _ := d.GetChange("name")
		oldNameString, ok := oldName.(string)
		if !ok {
			return diag.Errorf(`expected old "name" to be string`)
		}
		originalName = oldNameString
	}
	label := &github.Label{
		Name:  new(name),
		Color: new(color),
	}
	description, ok := d.Get("description").(string)
	if !ok {
		return diag.Errorf(`expected "description" to be string`)
	}
	label.Description = &description
	githubLabel, resp, err := client.Issues.EditLabel(ctx, orgName, repoName, originalName, label)
	if err != nil {
		return diag.FromErr(err)
	}
	id, err := buildID(repoName, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("url", githubLabel.GetURL()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubIssueLabelDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	repoName, ok := d.Get("repository").(string)
	if !ok {
		return diag.Errorf(`expected "repository" to be string`)
	}
	name, ok := d.Get("name").(string)
	if !ok {
		return diag.Errorf(`expected "name" to be string`)
	}
	_, err := client.Issues.DeleteLabel(ctx, orgName, repoName, name)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusForbidden {
				tflog.Info(ctx, "Ignoring delete of issue label in archived repository", map[string]any{"name": name, "org_name": orgName, "repo_name": repoName})
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubIssueLabelImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	repoName, name, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid import ID %q; expected format %q: %w", d.Id(), "<repository name>:<label name>", err)
	}
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	label, _, err := client.Issues.GetLabel(ctx, orgName, repoName, name)
	if err != nil {
		return nil, err
	}
	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("name", label.GetName()); err != nil {
		return nil, err
	}
	if err := d.Set("color", label.GetColor()); err != nil {
		return nil, err
	}
	if err := d.Set("description", label.GetDescription()); err != nil {
		return nil, err
	}
	if err := d.Set("url", label.GetURL()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
