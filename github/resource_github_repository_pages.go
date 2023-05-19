package github

import (
	"context"

	"github.com/google/go-github/v50/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubRepositoryPages() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryPagesCreate,
		Read:   resourceGithubRepositoryPagesRead,
		Update: resourceGithubRepositoryPagesUpdate,
		Delete: resourceGithubRepositoryPagesDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				d.Set("auto_init", false)
				return []*schema.ResourceData{d}, nil
			},
		},

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"cname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Pages site's custom domain",
			},
			"custom_404": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the rendered GitHub Pages site has a custom 404 page",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"html_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to the repository on the web.",
			},
			"public": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository for which Pages is being configured.",
			},
			"source": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "The source branch and directory for the rendered Pages site.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The repository branch used to publish the site's source files. (i.e. 'main' or 'gh-pages')",
						},
						"path": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "/",
							Description:  "The repository directory from which the site publishes (Default: '/')",
							ValidateFunc: validation.StringInSlice([]string{"/", "/docs"}, false),
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the most recent build of the Page.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The API address for accessing this Page resource.",
			},
		},
	}
}

func resourceGithubRepositoryPagesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	ctx := context.Background()

	pages := &github.Pages{}

	source := d.Get("source").([]interface{})[0].(map[string]interface{})
	sourceBranch := source["branch"].(string)
	sourcePath := ""
	if v, ok := source["path"].(string); ok {
		sourcePath = v
	}
	pages.Source = &github.PagesSource{Branch: &sourceBranch, Path: &sourcePath}

	pages, resp, err := client.Repositories.EnablePages(ctx, owner, repoName, pages)
	if err != nil {
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("url", pages.GetURL())
	d.Set("status", pages.GetStatus())
	d.Set("cname", pages.GetCNAME())
	d.Set("custom_404", pages.GetCustom404())
	d.Set("html_url", pages.GetHTMLURL())
	d.Set("source", flattenSource(pages))
	d.Set("public", pages.GetPublic())

	d.SetId(repoName)

	return resourceGithubRepositoryPagesUpdate(d, meta)
}

func resourceGithubRepositoryPagesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	pages, resp, err := client.Repositories.GetPagesInfo(ctx, owner, repoName)
	if err != nil {
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("url", pages.GetURL())
	d.Set("status", pages.GetStatus())
	d.Set("cname", pages.GetCNAME())
	d.Set("custom_404", pages.GetCustom404())
	d.Set("html_url", pages.GetHTMLURL())
	d.Set("source", flattenSource(pages))
	d.Set("public", pages.GetPublic())

	return nil
}

func resourceGithubRepositoryPagesUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	ctx := context.Background()

	update := &github.PagesUpdate{}

	// Only set the github.PagesUpdate CNAME field if the value is a non-empty string.
	// Leaving the CNAME field unset will remove the custom domain.
	if v, ok := d.Get("cname").(string); ok && v != "" {
		update.CNAME = github.String(v)
	}

	if v, ok := d.Get("https_enforced").(bool); ok {
		update.HTTPSEnforced = github.Bool(v)
	}

	if v, ok := d.Get("public").(bool); ok {
		update.Public = github.Bool(v)
	}

	// To update the GitHub Pages source, the github.PagesUpdate Source field
	// must include the branch name and optionally the subdirectory /docs.
	// e.g. "master" or "master /docs"
	source := d.Get("source").([]interface{})[0].(map[string]interface{})
	sourceBranch := source["branch"].(string)
	sourcePath := ""
	if v, ok := source["path"].(string); ok {
		sourcePath = v
	}
	update.Source = &github.PagesSource{Branch: &sourceBranch, Path: &sourcePath}

	_, err := client.Repositories.UpdatePages(ctx, owner, repoName, update)

	if err != nil {
		return err
	}

	return nil
}

func resourceGithubRepositoryPagesDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	repoName := d.Id()
	owner := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, err := client.Repositories.DisablePages(ctx, owner, repoName)

	return err
}

func flattenSource(pages *github.Pages) interface{} {
	if pages == nil {
		return []interface{}{}
	}

	sourceMap := make(map[string]interface{})
	sourceMap["branch"] = pages.GetSource().GetBranch()
	sourceMap["path"] = pages.GetSource().GetPath()

	return []interface{}{sourceMap}
}
