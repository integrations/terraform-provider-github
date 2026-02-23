package github

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseActionsHostedRunner() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseActionsHostedRunnerRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"runner_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The numeric ID of the hosted runner.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the hosted runner.",
			},
			"runner_group_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The runner group ID this runner belongs to.",
			},
			"platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The platform of the runner (e.g., 'linux-x64', 'win-x64').",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current status of the runner (e.g., 'Ready', 'Provisioning').",
			},
			"maximum_runners": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Maximum number of runners to scale up to.",
			},
			"public_ip_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether static public IP is enabled for this runner.",
			},
			"last_active_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RFC3339 timestamp indicating when the runner was last active.",
			},
			"image_details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details about the runner's image.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The image ID.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The image source (github, partner, or custom).",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The image version.",
						},
						"size_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the image in GB.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable display name for the image.",
						},
					},
				},
			},
			"machine_size_details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details about the runner's machine size.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Machine size identifier (e.g., '4-core', '8-core').",
						},
						"cpu_cores": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of CPU cores.",
						},
						"memory_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Amount of memory in GB.",
						},
						"storage_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Amount of SSD storage in GB.",
						},
					},
				},
			},
			"public_ips": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of public IP ranges assigned to this runner (only if public_ip_enabled is true).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this IP range is enabled.",
						},
						"prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address prefix.",
						},
						"length": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Subnet length (CIDR notation).",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseActionsHostedRunnerRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)
	runnerID := int64(d.Get("runner_id").(int))

	// Get the specific runner by ID
	runner, _, err := client.Enterprise.GetHostedRunner(ctx, enterpriseSlug, runnerID)
	if err != nil {
		return diag.Errorf("error reading enterprise hosted runner: %s", err.Error())
	}

	// Set the ID as enterprise_slug/runner_id
	id, err := buildID(enterpriseSlug, strconv.FormatInt(runner.GetID(), 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	runnerData := map[string]any{
		"name":                 runner.GetName(),
		"runner_group_id":      int(runner.GetRunnerGroupID()),
		"platform":             runner.GetPlatform(),
		"status":               runner.GetStatus(),
		"maximum_runners":      int(runner.GetMaximumRunners()),
		"public_ip_enabled":    runner.GetPublicIPEnabled(),
		"image_details":        flattenHostedRunnerImage(runner.ImageDetails),
		"machine_size_details": flattenHostedRunnerMachineSpec(runner.MachineSizeDetails),
		"public_ips":           flattenHostedRunnerPublicIPs(runner.PublicIPs),
	}

	if runner.LastActiveOn != nil {
		runnerData["last_active_on"] = runner.LastActiveOn.Format(time.RFC3339)
	}

	for k, v := range runnerData {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
