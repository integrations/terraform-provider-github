package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsHostedRunnerCustomImageVersions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsHostedRunnerCustomImageVersionsRead,

		Schema: map[string]*schema.Schema{
			"image_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The custom image definition ID.",
			},
			"versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version string (e.g., 1.0.0).",
						},
						"size_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Size of the image version in GB.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State of the version (e.g., Ready).",
						},
						"created_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when the version was created.",
						},
					},
				},
				Description: "List of versions for this custom image.",
			},
		},
	}
}

func dataSourceGithubActionsHostedRunnerCustomImageVersionsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	imageID := d.Get("image_id").(int)

	req, err := client.NewRequest("GET", fmt.Sprintf("orgs/%s/actions/hosted-runners/images/custom/%d/versions", orgName, imageID), nil)
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

	var allVersions []map[string]any
	for _, v := range result.ImageVersions {
		m := map[string]any{}
		if ver, ok := v["version"].(string); ok {
			m["version"] = ver
		}
		if size, ok := v["size_gb"].(float64); ok {
			m["size_gb"] = int(size)
		}
		if state, ok := v["state"].(string); ok {
			m["state"] = state
		}
		if created, ok := v["created_on"].(string); ok {
			m["created_on"] = created
		}
		allVersions = append(allVersions, m)
	}

	d.SetId(fmt.Sprintf("%s/%s", orgName, strconv.Itoa(imageID)))
	if err := d.Set("versions", allVersions); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
