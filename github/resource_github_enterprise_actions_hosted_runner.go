package github

import (
	"context"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseActionsHostedRunner() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubEnterpriseActionsHostedRunnerCreate,
		ReadContext:   resourceGithubEnterpriseActionsHostedRunnerRead,
		UpdateContext: resourceGithubEnterpriseActionsHostedRunnerUpdate,
		DeleteContext: resourceGithubEnterpriseActionsHostedRunnerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise. This is used to identify the enterprise in GitHub URLs and APIs.",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 64),
					validation.StringMatch(
						regexp.MustCompile(`^[a-zA-Z0-9._-]+$`),
						"name may only contain alphanumeric characters, '.', '-', and '_'",
					),
				),
				Description: "Name of the hosted runner. Must be between 1 and 64 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.",
			},
			"image": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "Image configuration for the hosted runner. This defines the operating system and software stack that will run on the runner. Cannot be changed after creation.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The image ID. For GitHub-owned images, use numeric IDs like '2306' for Ubuntu Latest 24.04. To get available images, use the GitHub API: GET /enterprises/{enterprise}/actions/hosted-runners/images/github-owned.",
						},
						"source": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "github",
							ValidateFunc: validation.StringInSlice([]string{"github", "partner", "custom"}, false),
							Description:  "The image source. Valid values are 'github' for GitHub-owned images, 'partner' for partner-provided images, or 'custom' for custom images. Defaults to 'github'.",
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "latest",
							Description: "The version of the runner image to deploy. For GitHub-owned images, this must be 'latest' (default). For custom images, you can specify a specific version.",
						},
						"size_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the image in GB. This is computed by the GitHub API and indicates the disk space required for the image.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable display name for this image. For example, '20.04' for Ubuntu 20.04.",
						},
					},
				},
			},
			"size": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Machine size for the hosted runner (e.g., '4-core', '8-core'). This determines the CPU, memory, and storage resources allocated to the runner. Can be updated to scale the runner. To list available sizes, use the GitHub API: GET /enterprises/{enterprise}/actions/hosted-runners/machine-sizes.",
			},
			"runner_group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the runner group to assign this runner to. Runner groups help organize runners and control which repositories or workflows can use them. You can get runner group IDs from the github_enterprise_actions_runner_group resource or data source.",
			},
			"maximum_runners": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "Maximum number of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit costs. If not specified, GitHub will use a default limit.",
			},
			"public_ip_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable static public IP for the runner. When enabled, the runner will be assigned a stable public IP address. Note that there are account-level limits for public IPs. To check limits, use the GitHub API: GET /enterprises/{enterprise}/actions/hosted-runners/limits. Defaults to false.",
			},
			"image_gen": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether this runner should be used to generate custom images. This is used for organizations that build their own custom runner images. Cannot be changed after creation. Defaults to false.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The hosted runner ID in the format {enterprise_slug}/{runner_id}. This is the unique identifier for the runner resource in Terraform.",
			},
			"runner_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The numeric ID of the hosted runner. This is the ID used in GitHub's API.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current status of the runner. Possible values include 'Ready', 'Provisioning', 'Deleting', etc. This indicates the operational state of the runner.",
			},
			"platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Platform of the runner. Examples: 'linux-x64', 'win-x64', 'macos-arm64'. This indicates the operating system and architecture of the runner.",
			},
			"machine_size_details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Detailed specifications of the machine size, including CPU cores, memory, and storage. This information is returned by the GitHub API and shows the actual resources allocated to the runner.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Machine size identifier. This matches the 'size' parameter used when creating the runner (e.g., '4-core', '8-core').",
						},
						"cpu_cores": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of CPU cores allocated to the runner. For example, a '4-core' runner has 4 CPU cores.",
						},
						"memory_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Amount of memory in gigabytes allocated to the runner. For example, a '4-core' runner typically has 16 GB of RAM.",
						},
						"storage_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Amount of SSD storage in gigabytes allocated to the runner. For example, a '4-core' runner typically has 150 GB of storage.",
						},
					},
				},
			},
			"public_ips": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of public IP ranges assigned to this runner. Only populated if 'public_ip_enabled' is true. These are the static IP addresses that will be used for outbound connections from the runner.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this IP range is enabled and active for the runner.",
						},
						"prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address prefix for the public IP range. Example: '20.80.208.150'. This is the base IP address.",
						},
						"length": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Subnet length for the IP range (CIDR notation length). Example: 28. This defines how many IP addresses are in the range.",
						},
					},
				},
			},
			"last_active_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RFC3339 timestamp indicating when the runner was last active. This helps track runner usage and can be used to identify idle runners.",
			},
		},
	}
}

func resourceGithubEnterpriseActionsHostedRunnerCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)

	// Build request using SDK struct
	request := &github.HostedRunnerRequest{
		Name:          d.Get("name").(string),
		Size:          d.Get("size").(string),
		RunnerGroupID: int64(d.Get("runner_group_id").(int)),
	}

	if image := expandHostedRunnerImage(d.Get("image").([]any)); image != nil {
		request.Image = *image
	}

	if v, ok := d.GetOk("maximum_runners"); ok {
		request.MaximumRunners = int64(v.(int))
	}

	if v, ok := d.GetOk("public_ip_enabled"); ok {
		request.EnableStaticIP = v.(bool)
	}

	runner, _, err := client.Enterprise.CreateHostedRunner(ctx, enterpriseSlug, request)
	if err != nil {
		return diag.Errorf("error creating enterprise hosted runner: %s", err.Error())
	}

	if runner == nil || runner.ID == nil {
		return diag.Errorf("no runner data returned from API")
	}

	// Set the ID in the format enterprise_slug/runner_id
	id, err := buildID(enterpriseSlug, strconv.FormatInt(runner.GetID(), 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	// Populate computed fields directly from API response
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}

	// runner.ID is guaranteed non-nil due to check above
	if err := d.Set("runner_id", int(runner.GetID())); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", runner.GetName()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("status", runner.GetStatus()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("platform", runner.GetPlatform()); err != nil {
		return diag.FromErr(err)
	}

	if runner.LastActiveOn != nil {
		if err := d.Set("last_active_on", runner.LastActiveOn.Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("public_ip_enabled", runner.GetPublicIPEnabled()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("image", flattenHostedRunnerImage(runner.ImageDetails)); err != nil {
		return diag.FromErr(err)
	}

	if runner.MachineSizeDetails != nil {
		if err := d.Set("size", runner.MachineSizeDetails.ID); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("machine_size_details", flattenHostedRunnerMachineSpec(runner.MachineSizeDetails)); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("runner_group_id", int(runner.GetRunnerGroupID())); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("maximum_runners", int(runner.GetMaximumRunners())); err != nil {
		return diag.FromErr(err)
	}

	if runner.PublicIPs != nil {
		if err := d.Set("public_ips", flattenHostedRunnerPublicIPs(runner.PublicIPs)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubEnterpriseActionsHostedRunnerRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug, runnerIDStr, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	runnerID, err := strconv.ParseInt(runnerIDStr, 10, 64)
	if err != nil {
		return diag.Errorf("invalid runner ID %q: %s", runnerIDStr, err.Error())
	}

	runner, resp, err := client.Enterprise.GetHostedRunner(ctx, enterpriseSlug, runnerID)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing enterprise hosted runner from state because it no longer exists in GitHub", map[string]any{
				"id":              d.Id(),
				"enterprise_slug": enterpriseSlug,
				"runner_id":       runnerID,
			})
			d.SetId("")
			return nil
		}
		return diag.Errorf("error reading enterprise hosted runner: %s", err.Error())
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("runner_id", int(runner.GetID())); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", runner.GetName()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("status", runner.GetStatus()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("platform", runner.GetPlatform()); err != nil {
		return diag.FromErr(err)
	}

	if runner.LastActiveOn != nil {
		if err := d.Set("last_active_on", runner.LastActiveOn.Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("public_ip_enabled", runner.GetPublicIPEnabled()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("image", flattenHostedRunnerImage(runner.ImageDetails)); err != nil {
		return diag.FromErr(err)
	}

	if runner.MachineSizeDetails != nil {
		if err := d.Set("size", runner.MachineSizeDetails.ID); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("machine_size_details", flattenHostedRunnerMachineSpec(runner.MachineSizeDetails)); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("runner_group_id", int(runner.GetRunnerGroupID())); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("maximum_runners", int(runner.GetMaximumRunners())); err != nil {
		return diag.FromErr(err)
	}

	if runner.PublicIPs != nil {
		if err := d.Set("public_ips", flattenHostedRunnerPublicIPs(runner.PublicIPs)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubEnterpriseActionsHostedRunnerUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug, runnerIDStr, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	runnerID, err := strconv.ParseInt(runnerIDStr, 10, 64)
	if err != nil {
		return diag.Errorf("invalid runner ID %q: %s", runnerIDStr, err.Error())
	}

	request := &github.HostedRunnerRequest{
		Name:           d.Get("name").(string),
		Size:           d.Get("size").(string),
		RunnerGroupID:  int64(d.Get("runner_group_id").(int)),
		MaximumRunners: int64(d.Get("maximum_runners").(int)),
		EnableStaticIP: d.Get("public_ip_enabled").(bool),
	}

	_, _, err = client.Enterprise.UpdateHostedRunner(ctx, enterpriseSlug, runnerID, *request)
	if err != nil {
		return diag.Errorf("error updating enterprise hosted runner: %s", err.Error())
	}

	return resourceGithubEnterpriseActionsHostedRunnerRead(ctx, d, meta)
}

func resourceGithubEnterpriseActionsHostedRunnerDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug, runnerIDStr, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	runnerID, err := strconv.ParseInt(runnerIDStr, 10, 64)
	if err != nil {
		return diag.Errorf("invalid runner ID %q: %s", runnerIDStr, err.Error())
	}

	_, resp, err := client.Enterprise.DeleteHostedRunner(ctx, enterpriseSlug, runnerID)

	// Check status code before error to handle expected HTTP responses.
	// Runner doesn't exist - treat as successful deletion since desired state is achieved.
	// This can happen if the runner was already deleted outside of Terraform.
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return nil
	}

	// Async deletion initiated (202 Accepted) - GitHub is processing the deletion.
	// Poll until the runner is actually removed from the API.
	if resp != nil && resp.StatusCode == http.StatusAccepted {
		if err := waitForEnterpriseRunnerDeletion(ctx, client, enterpriseSlug, runnerID, d.Timeout(schema.TimeoutDelete)); err != nil {
			return diag.Errorf("error waiting for enterprise hosted runner deletion: %s", err.Error())
		}
		return nil
	}

	// Any other error (API errors, network issues, etc.)
	if err != nil {
		return diag.Errorf("error deleting enterprise hosted runner: %s", err.Error())
	}

	// Successful synchronous deletion (200 OK)
	return nil
}

func waitForEnterpriseRunnerDeletion(ctx context.Context, client *github.Client, enterpriseSlug string, runnerID int64, timeout time.Duration) error {
	conf := &retry.StateChangeConf{
		Pending: []string{"deleting", "active"},
		Target:  []string{"deleted"},
		Refresh: func() (any, string, error) {
			_, resp, err := client.Enterprise.GetHostedRunner(ctx, enterpriseSlug, runnerID)
			if resp != nil && resp.StatusCode == http.StatusNotFound {
				return "deleted", "deleted", nil
			}

			if err != nil {
				return nil, "deleting", err
			}

			return "deleting", "deleting", nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	_, err := conf.WaitForStateContext(ctx)
	return err
}
