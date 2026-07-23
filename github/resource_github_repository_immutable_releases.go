package github

import (
	"context"
	"net/http"
	"strconv"

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

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository name to configure immutable releases for.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the repository to configure immutable releases for.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether immutable releases are enabled for the repository.",
			},
			"enforced_by_owner": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether immutable releases are enforced by the repository owner (organization).",
			},
		},

		CustomizeDiff: diffRepository,
	}
}

func resourceGithubRepositoryImmutableReleasesCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	tflog.Info(ctx, "Creating repository immutable releases setting", map[string]any{"id": d.Id()})
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner := meta.name
	repoName := d.Get("repository").(string)

	immutableReleasesEnabled := d.Get("enabled").(bool)
	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	if repo.GetArchived() {
		return diag.Errorf("cannot configure immutable releases on archived repository %s/%s", owner, repoName)
	}
	if immutableReleasesEnabled {
		_, err = client.Repositories.EnableImmutableReleases(ctx, owner, repoName)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		_, err = client.Repositories.DisableImmutableReleases(ctx, owner, repoName)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(strconv.Itoa(int(repo.GetID())))

	if err = d.Set("repository_id", repo.GetID()); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryImmutableReleasesRead(ctx, d, m)
}

func resourceGithubRepositoryImmutableReleasesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	tflog.Info(ctx, "Reading repository immutable releases setting", map[string]any{"id": d.Id()})
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner := meta.name
	repoName := d.Get("repository").(string)
	status, resp, err := client.Repositories.AreImmutableReleasesEnabled(ctx, owner, repoName)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			// API may 404 when disabled; treat as enabled=false if the repo still exists.
			repo, _, repoErr := client.Repositories.Get(ctx, owner, repoName)
			if repoErr != nil {
				tflog.Warn(ctx, "Removing immutable releases from state; repository not found", map[string]any{"owner": owner, "repo": repoName})
				d.SetId("")
				return nil
			}
			if err = d.Set("enabled", false); err != nil {
				return diag.FromErr(err)
			}
			if err = d.Set("enforced_by_owner", false); err != nil {
				return diag.FromErr(err)
			}
			if err = d.Set("repository_id", repo.GetID()); err != nil {
				return diag.FromErr(err)
			}
			return nil
		}
		return diag.Errorf("error reading repository immutable releases: %s", err.Error())
	}

	if err = d.Set("enabled", status.GetEnabled()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("enforced_by_owner", status.GetEnforcedByOwner()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryImmutableReleasesUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	tflog.Info(ctx, "Updating repository immutable releases setting", map[string]any{"id": d.Id()})
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner := meta.name
	repoName := d.Get("repository").(string)

	immutableReleasesEnabled := d.Get("enabled").(bool)
	var err error
	if immutableReleasesEnabled {
		_, err = client.Repositories.EnableImmutableReleases(ctx, owner, repoName)
	} else {
		_, err = client.Repositories.DisableImmutableReleases(ctx, owner, repoName)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryImmutableReleasesRead(ctx, d, m)
}

func resourceGithubRepositoryImmutableReleasesDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	tflog.Info(ctx, "Deleting repository immutable releases setting", map[string]any{"id": d.Id()})
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner := meta.name
	repoName := d.Get("repository").(string)
	_, err := client.Repositories.DisableImmutableReleases(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(handleArchivedRepoDelete(err, "repository immutable releases", d.Id(), owner, repoName))
	}

	return nil
}

func resourceGithubRepositoryImmutableReleasesImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	tflog.Debug(ctx, "Importing repository immutable releases setting", map[string]any{"id": d.Id()})
	repoName := d.Id()
	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}

	meta, _ := m.(*Owner)
	owner := meta.name
	client := meta.v3client

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}

	d.SetId(strconv.Itoa(int(repo.GetID())))

	if err = d.Set("repository_id", repo.GetID()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
