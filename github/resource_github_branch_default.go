package github

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// defaultBranchRenameTimeout bounds how long to wait for GitHub to converge on a
// renamed default branch. RenameBranch is subject to read-after-write eventual
// consistency: a GET can briefly return the previous default branch (and a new
// ETag) before the rename propagates.
const defaultBranchRenameTimeout = 2 * time.Minute

// waitForDefaultBranch polls the repository until GitHub reports the expected
// default branch, then returns the ETag of that converged response. The polling
// GET is unconditional so a stale 304 cannot mask an unpropagated rename.
// This is necessary because the GitHub API is eventually consistent for default branch renames,
// and a read immediately after a rename may return the old default branch with a new ETag.
func waitForDefaultBranch(ctx context.Context, client *github.Client, owner, repoName, expected string, timeout time.Duration) error {
	conf := &retry.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"converged"},
		Refresh: func() (any, string, error) {
			repository, _, err := client.Repositories.Get(ctx, owner, repoName)
			if err != nil {
				return nil, "", err
			}
			if repository.GetDefaultBranch() != expected {
				return repository, "pending", nil
			}
			return repository, "converged", nil
		},
		Timeout:    timeout,
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	if _, err := conf.WaitForStateContext(ctx); err != nil {
		return err
	}

	return nil
}

func resourceGithubBranchDefault() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubBranchDefaultCreate,
		ReadContext:   resourceGithubBranchDefaultRead,
		UpdateContext: resourceGithubBranchDefaultUpdate,
		DeleteContext: resourceGithubBranchDefaultDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubBranchDefaultImport,
		},

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubBranchDefaultV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubBranchDefaultStateUpgradeV0,
				Version: 0,
			},
		},

		CustomizeDiff: diffRepository,

		Description: "Configures the default branch for a GitHub repository.",

		Schema: map[string]*schema.Schema{
			"branch": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the branch to set as the default (e.g. 'main').",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the GitHub repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the GitHub repository.",
			},
			"rename": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicate if the current default branch should be renamed rather than switching to an existing branch. Defaults to 'false'.",
			},
			"etag": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ETag header for the repository API response.",
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					return true
				},
				DiffSuppressOnRefresh: true,
			},
		},
	}
}

func resourceGithubBranchDefaultCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	defaultBranch, _ := d.Get("branch").(string)
	rename, _ := d.Get("rename").(bool)

	tflog.Trace(ctx, "Creating default branch resource", map[string]any{"owner": owner, "repository": repoName, "branch": defaultBranch, "rename": rename})

	repository, resp, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	etag := resp.Header.Get("ETag")

	tflog.Debug(ctx, "Fetched repository", map[string]any{"current_default_branch": repository.GetDefaultBranch()})

	if repository.GetDefaultBranch() != defaultBranch {
		if rename {
			tflog.Debug(ctx, "Renaming branch to new default")
			if _, _, err := client.Repositories.RenameBranch(ctx, owner, repoName, repository.GetDefaultBranch(), defaultBranch); err != nil {
				return diag.FromErr(err)
			}
			err := waitForDefaultBranch(ctx, client, owner, repoName, defaultBranch, defaultBranchRenameTimeout)
			if err != nil {
				return diag.FromErr(err)
			}
			etag = ""
		} else {
			tflog.Debug(ctx, "Setting new default branch")
			repository := &github.Repository{
				DefaultBranch: new(defaultBranch),
			}

			if _, resp, err := client.Repositories.Edit(ctx, owner, repoName, repository); err != nil {
				return diag.FromErr(err)
			} else {
				etag = resp.Header.Get("ETag")
			}

		}
	} else {
		tflog.Debug(ctx, "Default branch already set to desired branch, skipping update")
	}

	d.SetId(repoName)

	if err := d.Set("repository_id", int(repository.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("etag", etag); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Finished creating default branch resource", map[string]any{"resource_id": d.Id()})

	return nil
}

func resourceGithubBranchDefaultRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "resource_id", d.Id())
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)

	tflog.Trace(ctx, "Reading default branch resource", map[string]any{"owner": owner, "repository": repoName})

	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	repository, resp, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				tflog.Debug(ctx, "Repository not modified, skipping read")
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing repository from state because it no longer exists in GitHub", map[string]any{"owner": owner, "repository": repoName})
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

	if err := d.Set("repository_id", int(repository.GetID())); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("branch", repository.GetDefaultBranch()); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Finished reading default branch resource")
	return nil
}

func resourceGithubBranchDefaultUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	defaultBranch, _ := d.Get("branch").(string)
	rename, _ := d.Get("rename").(bool)

	tflog.Trace(ctx, "Updating default branch resource", map[string]any{"owner": owner, "repository": repoName, "branch": defaultBranch, "rename": rename})

	var etag string

	if rename {
		tflog.Debug(ctx, "Rename enabled, checking if branch rename is needed")
		repository, resp, err := client.Repositories.Get(ctx, owner, repoName)
		if err != nil {
			return diag.FromErr(err)
		}
		etag = resp.Header.Get("ETag")
		if repository.GetDefaultBranch() != defaultBranch {
			tflog.Debug(ctx, "Renaming branch to new default")
			if _, _, err := client.Repositories.RenameBranch(ctx, owner, repoName, repository.GetDefaultBranch(), defaultBranch); err != nil {
				return diag.FromErr(err)
			}
			err := waitForDefaultBranch(ctx, client, owner, repoName, defaultBranch, defaultBranchRenameTimeout)
			if err != nil {
				return diag.FromErr(err)
			}
			etag = ""
		}
	} else {
		tflog.Debug(ctx, "Setting new default branch")
		repository := &github.Repository{
			DefaultBranch: new(defaultBranch),
		}

		if _, resp, err := client.Repositories.Edit(ctx, owner, repoName, repository); err != nil {
			return diag.FromErr(err)
		} else {
			etag = resp.Header.Get("ETag")
		}
	}

	if d.HasChange("repository") {
		d.SetId(repoName)
	}

	if err := d.Set("etag", etag); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Finished updating default branch resource")
	return nil
}

func resourceGithubBranchDefaultDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "resource_id", d.Id())
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repoName, _ := d.Get("repository").(string)

	tflog.Trace(ctx, "Deleting default branch resource", map[string]any{"owner": owner, "repository": repoName})

	repository := &github.Repository{
		DefaultBranch: nil,
	}

	_, _, err := client.Repositories.Edit(ctx, owner, repoName, repository)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing resource from state because repository no longer exists upstream", map[string]any{"owner": owner, "repository": repoName})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Finished deleting default branch resource")
	return nil
}

func resourceGithubBranchDefaultImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	repoName := d.Id()
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repository, resp, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("branch", repository.GetDefaultBranch()); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", int(repository.GetID())); err != nil {
		return nil, err
	}
	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
