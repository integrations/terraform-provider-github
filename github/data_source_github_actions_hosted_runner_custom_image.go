package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsHostedRunnerCustomImage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsHostedRunnerCustomImageRead,

		Schema: map[string]*schema.Schema{
			"image_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The custom image definition ID.",
			},
			"platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Platform of the image (e.g., linux-x64).",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the custom image.",
			},
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Source of the image.",
			},
			"versions_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of versions of this image.",
			},
			"total_versions_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total size of all versions in GB.",
			},
			"latest_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Latest version string.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the image (e.g., Ready).",
			},
		},
	}
}

func dataSourceGithubActionsHostedRunnerCustomImageRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	imageID := d.Get("image_id").(int)

	req, err := client.NewRequest("GET", fmt.Sprintf("orgs/%s/actions/hosted-runners/images/custom/%d", orgName, imageID), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	var image map[string]any
	_, err = client.Do(ctx, req, &image)
	if err != nil {
		return diag.FromErr(err)
	}

	if id, ok := image["id"].(float64); ok {
		d.SetId(strconv.FormatInt(int64(id), 10))
	}
	if v, ok := image["platform"].(string); ok {
		if err := d.Set("platform", v); err != nil {
			return diag.FromErr(err)
		}
	}
	if v, ok := image["name"].(string); ok {
		if err := d.Set("name", v); err != nil {
			return diag.FromErr(err)
		}
	}
	if v, ok := image["source"].(string); ok {
		if err := d.Set("source", v); err != nil {
			return diag.FromErr(err)
		}
	}
	if v, ok := image["versions_count"].(float64); ok {
		if err := d.Set("versions_count", int(v)); err != nil {
			return diag.FromErr(err)
		}
	}
	if v, ok := image["total_versions_size"].(float64); ok {
		if err := d.Set("total_versions_size", int(v)); err != nil {
			return diag.FromErr(err)
		}
	}
	if v, ok := image["latest_version"].(string); ok {
		if err := d.Set("latest_version", v); err != nil {
			return diag.FromErr(err)
		}
	}
	if v, ok := image["state"].(string); ok {
		if err := d.Set("state", v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
