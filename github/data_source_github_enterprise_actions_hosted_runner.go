package github

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseActionsHostedRunner() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubEnterpriseActionsHostedRunnerRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the hosted runner to lookup.",
			},
			"runner_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The numeric ID of the hosted runner.",
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

func dataSourceGithubEnterpriseActionsHostedRunnerRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Get("enterprise_slug").(string)
	runnerName := d.Get("name").(string)

	// List all runners and find the one matching the name
	opts := &github.ListOptions{PerPage: 100}
	var foundRunner *github.HostedRunner

	for {
		runners, resp, err := client.Enterprise.ListHostedRunners(ctx, enterpriseSlug, opts)
		if err != nil {
			return err
		}

		for _, runner := range runners.Runners {
			if runner.Name != nil && *runner.Name == runnerName {
				foundRunner = runner
				break
			}
		}

		if foundRunner != nil || resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	if foundRunner == nil {
		return fmt.Errorf("no hosted runner found with name %q in enterprise %q", runnerName, enterpriseSlug)
	}

	// Set the ID as enterprise_slug/runner_id
	d.SetId(fmt.Sprintf("%s/%d", enterpriseSlug, *foundRunner.ID))

	if foundRunner.ID != nil {
		if err := d.Set("runner_id", int(*foundRunner.ID)); err != nil {
			return err
		}
	}

	if foundRunner.RunnerGroupID != nil {
		if err := d.Set("runner_group_id", int(*foundRunner.RunnerGroupID)); err != nil {
			return err
		}
	}

	if foundRunner.Platform != nil {
		if err := d.Set("platform", *foundRunner.Platform); err != nil {
			return err
		}
	}

	if foundRunner.Status != nil {
		if err := d.Set("status", *foundRunner.Status); err != nil {
			return err
		}
	}

	if foundRunner.MaximumRunners != nil {
		if err := d.Set("maximum_runners", int(*foundRunner.MaximumRunners)); err != nil {
			return err
		}
	}

	if foundRunner.PublicIPEnabled != nil {
		if err := d.Set("public_ip_enabled", *foundRunner.PublicIPEnabled); err != nil {
			return err
		}
	}

	if foundRunner.LastActiveOn != nil {
		if err := d.Set("last_active_on", foundRunner.LastActiveOn.Format(time.RFC3339)); err != nil {
			return err
		}
	}

	if foundRunner.ImageDetails != nil {
		if err := d.Set("image_details", flattenHostedRunnerImage(foundRunner.ImageDetails)); err != nil {
			return err
		}
	}

	if foundRunner.MachineSizeDetails != nil {
		if err := d.Set("machine_size_details", flattenHostedRunnerMachineSpec(foundRunner.MachineSizeDetails)); err != nil {
			return err
		}
	}

	if foundRunner.PublicIPs != nil {
		if err := d.Set("public_ips", flattenHostedRunnerPublicIPs(foundRunner.PublicIPs)); err != nil {
			return err
		}
	}

	return nil
}
