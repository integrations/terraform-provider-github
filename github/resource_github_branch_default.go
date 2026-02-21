package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
			// StateContext: resourceGithubBranchDefaultImport,
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
			// TODO add repository_id and diffRepository to handle repository renames
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

func resourceGithubBranchDefaultCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	defaultBranch := d.Get("branch").(string)
	rename := d.Get("rename").(bool)

	tflog.Trace(ctx, "Creating default branch resource", map[string]any{
		"owner":      owner,
		"repository": repoName,
		"branch":     defaultBranch,
		"rename":     rename,
	})

	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Fetched repository", map[string]any{
		"current_default_branch": repository.GetDefaultBranch(),
	})

	if repository.GetDefaultBranch() != defaultBranch {
		if rename {
			tflog.Debug(ctx, "Renaming branch to new default")
			if _, _, err := client.Repositories.RenameBranch(ctx, owner, repoName, repository.GetDefaultBranch(), defaultBranch); err != nil {
				return diag.FromErr(err)
			}
		} else {
			tflog.Debug(ctx, "Setting new default branch")
			repository := &github.Repository{
				DefaultBranch: github.Ptr(defaultBranch),
			}

			if _, _, err := client.Repositories.Edit(ctx, owner, repoName, repository); err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		tflog.Debug(ctx, "Default branch already set to desired branch, skipping update")
	}

	d.SetId(repoName)

	tflog.Trace(ctx, "Finished creating default branch resource", map[string]any{
		"resource_id": d.Id(),
	})
	return resourceGithubBranchDefaultRead(ctx, d, meta)
}

func resourceGithubBranchDefaultRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "resource_id", d.Id())
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	// repoName := d.Get("repository").(string)
	repoName := d.Id()

	tflog.Trace(ctx, "Reading default branch resource", map[string]any{
		"owner":      owner,
		"repository": repoName,
	})

	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	repository, resp, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				tflog.Debug(ctx, "Repository not modified, skipping read")
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing repository from state because it no longer exists in GitHub", map[string]any{
					"owner":      owner,
					"repository": repoName,
				})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if repository.DefaultBranch == nil {
		tflog.Warn(ctx, "Default branch is nil, removing resource from state")
		d.SetId("")
		return nil
	}

	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("branch", repository.GetDefaultBranch()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository", repository.GetName()); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Finished reading default branch resource")
	return nil
}

func resourceGithubBranchDefaultUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "resource_id", d.Id())
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	// repoName := d.Get("repository").(string)
	repoName := d.Id()
	defaultBranch := d.Get("branch").(string)
	rename := d.Get("rename").(bool)

	tflog.Trace(ctx, "Updating default branch resource", map[string]any{
		"owner":      owner,
		"repository": repoName,
		"branch":     defaultBranch,
		"rename":     rename,
	})

	if rename {
		tflog.Debug(ctx, "Renaming branch to new default")
		repository, _, err := client.Repositories.Get(ctx, owner, repoName)
		if err != nil {
			return diag.FromErr(err)
		}
		if _, _, err := client.Repositories.RenameBranch(ctx, owner, repoName, repository.GetDefaultBranch(), defaultBranch); err != nil {
			return diag.FromErr(err)
		}
	} else {
		tflog.Debug(ctx, "Setting new default branch")
		repository := &github.Repository{
			DefaultBranch: github.Ptr(defaultBranch),
		}

		if _, _, err := client.Repositories.Edit(ctx, owner, repoName, repository); err != nil {
			return diag.FromErr(err)
		}
	}

	tflog.Trace(ctx, "Finished updating default branch resource")
	return resourceGithubBranchDefaultRead(ctx, d, meta)
}

func resourceGithubBranchDefaultDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "resource_id", d.Id())
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	// repoName := d.Get("repository").(string)
	repoName := d.Id()

	tflog.Trace(ctx, "Deleting default branch resource", map[string]any{
		"owner":      owner,
		"repository": repoName,
	})

	repository := &github.Repository{
		DefaultBranch: nil,
	}

	_, _, err := client.Repositories.Edit(ctx, owner, repoName, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Finished deleting default branch resource")
	return nil
}

// func resourceGithubBranchDefaultImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
// 	repoName, defaultBranch, err := parseID2(d.Id())
// 	if err != nil {
// 		return nil, err
// 	}

// 	d.SetId(repoName)
// 	if err := d.Set("branch", defaultBranch); err != nil {
// 		return nil, err
// 	}
// 	if err := d.Set("repository", repoName); err != nil {
// 		return nil, err
// 	}

// 	return []*schema.ResourceData{d}, nil
// }
