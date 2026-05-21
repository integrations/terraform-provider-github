package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsHostedRunnerCustomImages() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsHostedRunnerCustomImagesRead,

		Schema: map[string]*schema.Schema{
			"images": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
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
				},
				Description: "List of custom images for GitHub-hosted runners.",
			},
		},
	}
}

func dataSourceGithubActionsHostedRunnerCustomImagesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	req, err := client.NewRequest("GET", fmt.Sprintf("orgs/%s/actions/hosted-runners/images/custom", orgName), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	var result struct {
		TotalCount    int              `json:"total_count"`
		ImageVersions []map[string]any `json:"image_versions"`
	}
	_, err = client.Do(ctx, req, &result)
	if err != nil {
		return diag.FromErr(err)
	}

	var allImages []map[string]any
	for _, img := range result.ImageVersions {
		m := map[string]any{}
		if id, ok := img["id"].(float64); ok {
			m["id"] = strconv.FormatInt(int64(id), 10)
		}
		if v, ok := img["platform"].(string); ok {
			m["platform"] = v
		}
		if v, ok := img["name"].(string); ok {
			m["name"] = v
		}
		if v, ok := img["source"].(string); ok {
			m["source"] = v
		}
		if v, ok := img["versions_count"].(float64); ok {
			m["versions_count"] = int(v)
		}
		if v, ok := img["total_versions_size"].(float64); ok {
			m["total_versions_size"] = int(v)
		}
		if v, ok := img["latest_version"].(string); ok {
			m["latest_version"] = v
		}
		if v, ok := img["state"].(string); ok {
			m["state"] = v
		}
		allImages = append(allImages, m)
	}

	d.SetId(orgName)
	if err := d.Set("images", allImages); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
