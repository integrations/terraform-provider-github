package github

import (
	"context"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRelease() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubReleaseRead,
		Description: "Use this data source to retrieve information about a GitHub release in a specific repository.",
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository to retrieve the release from.",
			},
			"owner": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Owner of the repository.",
			},
			"retrieve_by": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Describes how to fetch the release. Valid values are `id`, `tag`, `latest`.",
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{
					"latest",
					"id",
					"tag",
				}, false), "retrieve_by"),
			},
			"release_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the release to retrieve. Must be specified when `retrieve_by` = `tag`.",
			},
			"release_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of the release to retrieve. Must be specified when `retrieve_by` = `id`.",
			},
			"target_commitish": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Commitish value that determines where the Git release is created from.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the release.",
			},
			"body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Contents of the description (body) of a release.",
			},
			"draft": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the release is a draft.",
			},
			"prerelease": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the release is a prerelease.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of release creation.",
			},
			"published_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of release publishing.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Base URL of the release.",
			},
			"html_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL directing to detailed information on the release.",
			},
			"assets_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of any associated assets with the release.",
			},
			"asserts_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Deprecated:  "use assets_url instead",
				Description: "URL of any associated assets with the release.",
			},
			"upload_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL that can be used to upload assets to the release.",
			},
			"zipball_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Download URL of the release in zip format.",
			},
			"tarball_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Download URL of the release in tar.gz format.",
			},
			"assets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of assets for the release.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of the asset.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the asset.",
						},
						"node_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Node ID of the asset.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The file name of the asset.",
						},
						"label": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Label for the asset.",
						},
						"content_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "MIME type of the asset.",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Size of the asset in bytes.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date the asset was created.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date the asset was last updated.",
						},
						"browser_download_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Browser download URL.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubReleaseRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	repository := d.Get("repository").(string)
	owner := d.Get("owner").(string)

	client := meta.(*Owner).v3client

	var err error
	var release *github.RepositoryRelease

	switch retrieveBy := strings.ToLower(d.Get("retrieve_by").(string)); retrieveBy {
	case "latest":
		release, _, err = client.Repositories.GetLatestRelease(ctx, owner, repository)
	case "id":
		releaseID := int64(d.Get("release_id").(int))
		if releaseID == 0 {
			return diag.Errorf("`release_id` must be set when `retrieve_by` = `id`")
		}

		release, _, err = client.Repositories.GetRelease(ctx, owner, repository, releaseID)
	case "tag":
		tag := d.Get("release_tag").(string)
		if tag == "" {
			return diag.Errorf("`release_tag` must be set when `retrieve_by` = `tag`")
		}

		release, _, err = client.Repositories.GetReleaseByTag(ctx, owner, repository, tag)
	default:
		return diag.Errorf("one of: `latest`, `id`, `tag` must be set for `retrieve_by`")
	}

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(release.GetID(), 10))
	if err = d.Set("release_tag", release.GetTagName()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("target_commitish", release.GetTargetCommitish()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("name", release.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("body", release.GetBody()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("draft", release.GetDraft()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("prerelease", release.GetPrerelease()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("created_at", release.GetCreatedAt().String()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("published_at", release.GetPublishedAt().String()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("url", release.GetURL()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("html_url", release.GetHTMLURL()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("assets_url", release.GetAssetsURL()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("asserts_url", release.GetAssetsURL()); err != nil { // Deprecated, original version of assets_url
		return diag.FromErr(err)
	}
	if err = d.Set("upload_url", release.GetUploadURL()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("zipball_url", release.GetZipballURL()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("tarball_url", release.GetTarballURL()); err != nil {
		return diag.FromErr(err)
	}

	assets := make([]any, 0, len(release.Assets))
	for _, releaseAsset := range release.Assets {
		if releaseAsset == nil {
			continue
		}

		assets = append(assets, map[string]any{
			"id":                   releaseAsset.GetID(),
			"url":                  releaseAsset.GetURL(),
			"node_id":              releaseAsset.GetNodeID(),
			"name":                 releaseAsset.GetName(),
			"label":                releaseAsset.GetLabel(),
			"content_type":         releaseAsset.GetContentType(),
			"size":                 releaseAsset.GetSize(),
			"created_at":           releaseAsset.GetCreatedAt().String(),
			"updated_at":           releaseAsset.GetUpdatedAt().String(),
			"browser_download_url": releaseAsset.GetBrowserDownloadURL(),
		})
	}

	if err = d.Set("assets", assets); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
