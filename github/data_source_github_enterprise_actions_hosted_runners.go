package github

import (
	"context"
	"time"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseActionsHostedRunners() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubEnterpriseActionsHostedRunnersRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"runners": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of hosted runners for the enterprise.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the hosted runner.",
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
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseActionsHostedRunnersRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Get("enterprise_slug").(string)

	// List all hosted runners with pagination
	opts := &github.ListOptions{PerPage: 100}
	var allRunners []*github.HostedRunner

	for {
		runners, resp, err := client.Enterprise.ListHostedRunners(ctx, enterpriseSlug, opts)
		if err != nil {
			return err
		}

		allRunners = append(allRunners, runners.Runners...)

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	// Set the ID as the enterprise slug
	d.SetId(enterpriseSlug)

	// Flatten runners data
	if err := d.Set("runners", flattenHostedRunners(allRunners)); err != nil {
		return err
	}

	return nil
}

func flattenHostedRunners(runners []*github.HostedRunner) []any {
	if runners == nil {
		return []any{}
	}

	result := make([]any, 0, len(runners))
	for _, runner := range runners {
		if runner == nil {
			continue
		}

		runnerMap := make(map[string]any)

		if runner.ID != nil {
			runnerMap["id"] = int(*runner.ID)
		}
		if runner.Name != nil {
			runnerMap["name"] = *runner.Name
		}
		if runner.RunnerGroupID != nil {
			runnerMap["runner_group_id"] = int(*runner.RunnerGroupID)
		}
		if runner.Platform != nil {
			runnerMap["platform"] = *runner.Platform
		}
		if runner.Status != nil {
			runnerMap["status"] = *runner.Status
		}
		if runner.MaximumRunners != nil {
			runnerMap["maximum_runners"] = int(*runner.MaximumRunners)
		}
		if runner.PublicIPEnabled != nil {
			runnerMap["public_ip_enabled"] = *runner.PublicIPEnabled
		}
		if runner.LastActiveOn != nil {
			runnerMap["last_active_on"] = runner.LastActiveOn.Format(time.RFC3339)
		}

		if runner.ImageDetails != nil {
			runnerMap["image_details"] = flattenHostedRunnerImage(runner.ImageDetails)
		}

		if runner.MachineSizeDetails != nil {
			runnerMap["machine_size_details"] = flattenHostedRunnerMachineSpec(runner.MachineSizeDetails)
		}

		result = append(result, runnerMap)
	}

	return result
}
