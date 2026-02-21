package github

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubBranchDefault() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubBranchDefaultCreate,
		ReadContext:   resourceGithubBranchDefaultRead,
		UpdateContext: resourceGithubBranchDefaultUpdate,
		DeleteContext: resourceGithubBranchDefaultDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"branch": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The branch (e.g. 'main').",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub repository.",
			},
			"rename": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicate if it should rename the branch rather than use an existing branch. Defaults to 'false'.",
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

func resourceGithubBranchDefaultCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	defaultBranch := d.Get("branch").(string)
	rename := d.Get("rename").(bool)

	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	if *repository.DefaultBranch != defaultBranch {
		if rename {
			if _, _, err := client.Repositories.RenameBranch(ctx, owner, repoName, *repository.DefaultBranch, defaultBranch); err != nil {
				return diag.FromErr(err)
			}
		} else {
			repository := &github.Repository{
				DefaultBranch: &defaultBranch,
			}

			if _, _, err := client.Repositories.Edit(ctx, owner, repoName, repository); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId(repoName)

	return resourceGithubBranchDefaultRead(ctx, d, meta)
}

func resourceGithubBranchDefaultRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()

	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	repository, resp, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing repository %s/%s from state because it no longer exists in GitHub",
					owner, repoName)
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if repository.DefaultBranch == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("etag", resp.Header.Get("ETag"))
	_ = d.Set("branch", *repository.DefaultBranch)
	_ = d.Set("repository", *repository.Name)
	return nil
}

func resourceGithubBranchDefaultUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()
	defaultBranch := d.Get("branch").(string)
	rename := d.Get("rename").(bool)

	if rename {
		repository, _, err := client.Repositories.Get(ctx, owner, repoName)
		if err != nil {
			return diag.FromErr(err)
		}
		if _, _, err := client.Repositories.RenameBranch(ctx, owner, repoName, *repository.DefaultBranch, defaultBranch); err != nil {
			return diag.FromErr(err)
		}
	} else {
		repository := &github.Repository{
			DefaultBranch: &defaultBranch,
		}

		if _, _, err := client.Repositories.Edit(ctx, owner, repoName, repository); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceGithubBranchDefaultRead(ctx, d, meta)
}

func resourceGithubBranchDefaultDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()

	repository := &github.Repository{
		DefaultBranch: nil,
	}

	_, _, err := client.Repositories.Edit(ctx, owner, repoName, repository)
	return diag.FromErr(err)
}
