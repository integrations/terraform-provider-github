package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRelease() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubReleaseCreate,
		ReadContext:   resourceGithubReleaseRead,
		UpdateContext: resourceGithubReleaseUpdate,
		DeleteContext: resourceGithubReleaseDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubReleaseImport,
		},

		CustomizeDiff: diffRepository,

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubReleaseV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubReleaseStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "Resource to manage a GitHub release.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the repository.",
			},
			"tag_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the tag.",
			},
			"target_commitish": {
				Type:        schema.TypeString,
				Default:     "main",
				Optional:    true,
				ForceNew:    true,
				Description: "The branch name or commit SHA the tag is created from; this defaults to `main`.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the release.",
			},
			"body": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Text describing the contents of the tag.",
			},
			"draft": {
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				Description: "Set to `false` to create a published release.",
			},
			"prerelease": {
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				Description: "Set to `false` to identify the release as a full release.",
			},
			"generate_release_notes": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Set to `true` to automatically generate the name and body for this release when it is created. If `name` is specified, the specified name will be used; otherwise, a name will be automatically generated. If `body` is specified, the body will be pre-pended to the automatically generated notes.",
			},
			"discussion_category_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. If there is already a discussion linked to the release, this parameter is ignored.",
			},
			"release_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the release.",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The node ID of the release.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the release was created.",
			},
			"published_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the release was published.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the release.",
			},
			"html_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTML URL for the release.",
			},
			"assets_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the release assets.",
			},
			"upload_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the uploaded assets of release.",
			},
			"zipball_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the zipball of the release.",
			},
			"tarball_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the tarball of the release.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubReleaseCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	tagName, _ := d.Get("tag_name").(string)
	targetCommitish, _ := d.Get("target_commitish").(string)
	draft, _ := d.Get("draft").(bool)
	prerelease, _ := d.Get("prerelease").(bool)
	generateReleaseNotes, _ := d.Get("generate_release_notes").(bool)

	req := github.CreateReleaseRequest{
		TagName:              tagName,
		TargetCommitish:      new(targetCommitish),
		Draft:                new(draft),
		Prerelease:           new(prerelease),
		GenerateReleaseNotes: new(generateReleaseNotes),
	}

	if v, ok := d.GetOk("body"); ok {
		s, _ := v.(string)
		req.Body = new(s)
	}

	if v, ok := d.GetOk("name"); ok {
		s, _ := v.(string)
		req.Name = new(s)
	}

	if v, ok := d.GetOk("discussion_category_name"); ok {
		s, _ := v.(string)
		req.DiscussionCategoryName = new(s)
	}

	tflog.Debug(ctx, "Creating release.", map[string]any{"target_commitish": targetCommitish, "release_tag": tagName, "repository": repoName, "owner": owner})

	release, _, err := client.Repositories.CreateRelease(ctx, owner, repoName, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(release.GetID(), 10))
	if err := d.Set("release_id", release.GetID()); err != nil {
		return diag.FromErr(err)
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	repoID := int(repo.GetID())

	if err := d.Set("repository_id", repoID); err != nil {
		return diag.FromErr(err)
	}

	if err := repositoryReleaseToComputedResourceData(d, release); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubReleaseRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repository := d.Get("repository").(string)
	releaseID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to convert release id %s to int64: %w", d.Id(), err))
	}

	release, _, err := client.Repositories.GetRelease(ctx, owner, repository, releaseID)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing release from state because it no longer exists on GitHub.", map[string]any{"release_id": releaseID, "repository": repository})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := repositoryReleaseToComputedResourceData(d, release); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubReleaseUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)

	releaseID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to convert release id %s to int64: %w", d.Id(), err))
	}

	draft, _ := d.Get("draft").(bool)
	prerelease, _ := d.Get("prerelease").(bool)

	req := github.UpdateReleaseRequest{
		Draft:      new(draft),
		Prerelease: new(prerelease),
	}

	if v, ok := d.GetOk("body"); ok {
		req.Body = new(v.(string))
	}

	if v, ok := d.GetOk("name"); ok {
		req.Name = new(v.(string))
	}

	if v, ok := d.GetOk("discussion_category_name"); ok {
		req.DiscussionCategoryName = new(v.(string))
	}

	tflog.Debug(ctx, "Updating release.", map[string]any{"release_id": releaseID, "repository": repoName, "owner": owner})

	release, _, err := client.Repositories.UpdateRelease(ctx, owner, repoName, releaseID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := repositoryReleaseToComputedResourceData(d, release); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubReleaseDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repository, _ := d.Get("repository").(string)

	releaseID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to convert release id %s to int64: %w", d.Id(), err))
	}

	if _, err = client.Repositories.DeleteRelease(ctx, owner, repository, releaseID); err != nil {
		return diag.Errorf("error deleting release %s/%s/%d: %v", owner, repository, releaseID, err)
	}
	return nil
}

func resourceGithubReleaseImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, releaseIDStr, err := parseID2(d.Id())
	if err != nil {
		return nil, err
	}

	releaseID, err := strconv.ParseInt(releaseIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("unable to convert release id %s to int64: %w", releaseIDStr, err)
	}
	if releaseID == 0 {
		return nil, fmt.Errorf("release_id must be present")
	}

	tflog.Debug(ctx, "Importing release.", map[string]any{"release_id": releaseID, "repository": repoName})

	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	if err = d.Set("repository", repository.GetName()); err != nil {
		return nil, err
	}
	if err = d.Set("repository_id", repository.GetID()); err != nil {
		return nil, err
	}

	release, _, err := client.Repositories.GetRelease(ctx, owner, repository.GetName(), releaseID)
	if err != nil {
		return nil, err
	}
	d.SetId(strconv.FormatInt(release.GetID(), 10))
	if err := d.Set("release_id", release.GetID()); err != nil {
		return nil, err
	}
	if err := d.Set("tag_name", release.GetTagName()); err != nil {
		return nil, err
	}
	if err := d.Set("target_commitish", release.GetTargetCommitish()); err != nil {
		return nil, err
	}
	if release.Name != nil {
		if err := d.Set("name", release.GetName()); err != nil {
			return nil, err
		}
	}
	if release.Body != nil {
		if err := d.Set("body", release.GetBody()); err != nil {
			return nil, err
		}
	}
	if err := d.Set("draft", release.GetDraft()); err != nil {
		return nil, err
	}
	if err := d.Set("prerelease", release.GetPrerelease()); err != nil {
		return nil, err
	}
	if err := d.Set("generate_release_notes", false); err != nil {
		return nil, err
	}
	if err := repositoryReleaseToComputedResourceData(d, release); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func repositoryReleaseToComputedResourceData(d *schema.ResourceData, release *github.RepositoryRelease) error {
	if err := d.Set("node_id", release.GetNodeID()); err != nil {
		return err
	}
	if err := d.Set("created_at", release.GetCreatedAt().String()); err != nil {
		return err
	}
	if err := d.Set("published_at", release.GetPublishedAt().String()); err != nil {
		return err
	}
	if err := d.Set("url", release.GetURL()); err != nil {
		return err
	}
	if err := d.Set("html_url", release.GetHTMLURL()); err != nil {
		return err
	}
	if err := d.Set("assets_url", release.GetAssetsURL()); err != nil {
		return err
	}
	if err := d.Set("upload_url", release.GetUploadURL()); err != nil {
		return err
	}
	if err := d.Set("zipball_url", release.GetZipballURL()); err != nil {
		return err
	}
	if err := d.Set("tarball_url", release.GetTarballURL()); err != nil {
		return err
	}

	return nil
}
