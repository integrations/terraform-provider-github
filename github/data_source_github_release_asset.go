package github

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubReleaseAsset() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieve information about a GitHub release asset.",
		ReadContext: dataSourceGithubReleaseAssetRead,

		Schema: map[string]*schema.Schema{
			"asset_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "ID of the release asset to retrieve",
			},
			"owner": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Owner of the repository",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository to retrieve the release asset from",
			},
			"body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The release asset body",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the asset",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Node ID of the asset",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "File name of the asset",
			},
			"label": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Label for the asset",
			},
			"content_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "MIME type of the asset",
			},
			"size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Asset size in bytes",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date the asset was created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date the asset was updated",
			},
			"browser_download_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Browser URL from which the release asset can be downloaded",
			},
		},
	}
}

func dataSourceGithubReleaseAssetRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	repository := d.Get("repository").(string)
	owner := d.Get("owner").(string)

	client := meta.(*Owner).v3client

	assetID := int64(d.Get("asset_id").(int))
	asset, _, err := client.Repositories.GetReleaseAsset(ctx, owner, repository, assetID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Use a client copy to avoid possible mutation of shared GitHub client state
	// by client.Repositories.DownloadReleaseAsset.
	clientCopy := client.Client()
	req, err := http.NewRequestWithContext(ctx, "GET", asset.GetBrowserDownloadURL(), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Accept", "application/octet-stream")
	resp, err := clientCopy.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return diag.Errorf("failed to get release asset (%s/%s %d): %s", owner, repository, assetID, resp.Status)
	}

	buf := new(strings.Builder)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildThreePartID(owner, repository, strconv.FormatInt(assetID, 10)))

	if err := d.Set("body", buf.String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", asset.URL); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("node_id", asset.NodeID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", asset.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("label", asset.Label); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("content_type", asset.ContentType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("size", asset.Size); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", asset.CreatedAt.String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", asset.UpdatedAt.String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("browser_download_url", asset.BrowserDownloadURL); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
