package github

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryImmutableReleases() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryImmutableReleasesCreate,
		ReadContext:   resourceGithubRepositoryImmutableReleasesRead,
		UpdateContext: resourceGithubRepositoryImmutableReleasesUpdate,
		DeleteContext: resourceGithubRepositoryImmutableReleasesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryImmutableReleasesImport,
		},

		CustomizeDiff: diffRepository,

		Description: "Manages immutable releases for a GitHub repository.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository to configure immutable releases for.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the repository.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether immutable releases are enabled for the repository. Defaults to `true`.",
			},
		},
	}
}

func resourceGithubRepositoryImmutableReleasesCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	enabled, _ := d.Get("enabled").(bool)

	tflog.Debug(ctx, "Creating repository immutable releases", map[string]any{"owner": owner, "repository": repoName, "enabled": enabled})

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setRepositoryImmutableReleases(ctx, client, owner, repoName, enabled); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(repo.GetID(), 10))

	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryImmutableReleasesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)

	var enabled bool
	status, _, err := client.Repositories.AreImmutableReleasesEnabled(ctx, owner, repoName)
	if err != nil {
		ghErr, ok := errors.AsType[*github.ErrorResponse](err)
		if !ok || ghErr.Response.StatusCode != http.StatusNotFound {
			return diag.FromErr(err)
		}

		// The API documents a 404 for disabled immutable releases, although it
		// currently responds with enabled=false instead. Look the repository up
		// so a 404 only removes the resource when the repository is gone.
		if _, _, err := client.Repositories.Get(ctx, owner, repoName); err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing repository immutable releases from state because the repository no longer exists", map[string]any{"owner": owner, "repository": repoName})
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
	} else {
		enabled = status.GetEnabled()
	}

	if err := d.Set("enabled", enabled); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryImmutableReleasesUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if !d.HasChange("enabled") {
		return nil
	}

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	enabled, _ := d.Get("enabled").(bool)

	tflog.Debug(ctx, "Updating repository immutable releases", map[string]any{"owner": owner, "repository": repoName, "enabled": enabled})

	if err := setRepositoryImmutableReleases(ctx, client, owner, repoName, enabled); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryImmutableReleasesDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)

	tflog.Debug(ctx, "Deleting repository immutable releases", map[string]any{"owner": owner, "repository": repoName})

	if _, err := client.Repositories.DisableImmutableReleases(ctx, owner, repoName); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(handleArchivedRepoDelete(err, "repository immutable releases", d.Id(), owner, repoName))
	}

	return nil
}

func resourceGithubRepositoryImmutableReleasesImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Id()

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}

	d.SetId(strconv.FormatInt(repo.GetID(), 10))

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func setRepositoryImmutableReleases(ctx context.Context, client *github.Client, owner, repoName string, enabled bool) error {
	var err error
	if enabled {
		_, err = client.Repositories.EnableImmutableReleases(ctx, owner, repoName)
	} else {
		_, err = client.Repositories.DisableImmutableReleases(ctx, owner, repoName)
	}
	return err
}
