package github

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsHostedRunner() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsHostedRunnerCreate,
		ReadContext:   resourceGithubActionsHostedRunnerRead,
		UpdateContext: resourceGithubActionsHostedRunnerUpdate,
		DeleteContext: resourceGithubActionsHostedRunnerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(
					validation.StringLenBetween(1, 64),
					validation.StringMatch(
						regexp.MustCompile(`^[a-zA-Z0-9._-]+$`),
						"name may only contain alphanumeric characters, '.', '-', and '_'",
					),
				)),
				Description: "Name of the hosted runner. Must be between 1 and 64 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.",
			},
			"image": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The image ID.",
						},
						"source": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "github",
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"github", "partner", "custom"}, false)),
							Description:      "The image source (github, partner, or custom).",
						},
						"size_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the image in GB.",
						},
					},
				},
				Description: "Image configuration for the hosted runner. Cannot be changed after creation.",
			},
			"size": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Machine size (e.g., '4-core', '8-core'). Can be updated to scale the runner.",
			},
			"runner_group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The runner group ID.",
			},
			"maximum_runners": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Description:      "Maximum number of runners to scale up to.",
			},
			"public_ip_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable static public IP.",
			},
			"image_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of the runner image to deploy. This is relevant only for runners using custom images.",
			},
			"image_gen": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether this runner should be used to generate custom images. Cannot be changed after creation.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The hosted runner ID.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current status of the runner.",
			},
			"platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Platform of the runner.",
			},
			"machine_size_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Machine size ID.",
						},
						"cpu_cores": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of CPU cores.",
						},
						"memory_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory in GB.",
						},
						"storage_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Storage in GB.",
						},
					},
				},
				Description: "Detailed machine size specifications.",
			},
			"public_ips": {
				Type:     schema.TypeList,
				Computed: true,
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
							Description: "Subnet length.",
						},
					},
				},
				Description: "List of public IP ranges assigned to this runner.",
			},
			"last_active_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the runner was last active.",
			},
		},
	}
}

func expandImage(imageList []any) github.HostedRunnerImage {
	if len(imageList) == 0 {
		return github.HostedRunnerImage{}
	}
	runnerImage := github.HostedRunnerImage{}

	imageMap := imageList[0].(map[string]any)

	if id, ok := imageMap["id"].(string); ok {
		runnerImage.ID = id
	}
	if source, ok := imageMap["source"].(string); ok {
		runnerImage.Source = source
	}

	return runnerImage
}

func flattenImage(image *github.HostedRunnerImageDetail) []any {
	if image == nil {
		return []any{}
	}

	result := make(map[string]any)

	result["id"] = image.GetID()
	result["source"] = image.GetSource()
	result["size_gb"] = image.GetSizeGB()

	return []any{result}
}

func flattenMachineSizeDetails(details *github.HostedRunnerMachineSpec) []any {
	if details == nil {
		return []any{}
	}

	result := make(map[string]any)
	result["id"] = details.GetID()
	result["cpu_cores"] = details.GetCPUCores()
	result["memory_gb"] = details.GetMemoryGB()
	result["storage_gb"] = details.GetStorageGB()

	return []any{result}
}

func flattenPublicIPs(ips []*github.HostedRunnerPublicIP) []any {
	if ips == nil {
		return []any{}
	}

	result := make([]any, 0, len(ips))
	for _, ip := range ips {
		ipResult := make(map[string]any)
		ipResult["enabled"] = ip.GetEnabled()
		ipResult["prefix"] = ip.GetPrefix()
		ipResult["length"] = ip.GetLength()
		result = append(result, ipResult)
	}

	return result
}

func resourceGithubActionsHostedRunnerCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	runnerName, _ := d.Get("name").(string)
	runnerImageDetails, _ := d.Get("image").([]any)
	runnerSize, _ := d.Get("size").(string)
	runnerGroupID, _ := d.Get("runner_group_id").(int)

	req := github.CreateHostedRunnerRequest{
		Name:          runnerName,
		Image:         expandImage(runnerImageDetails),
		Size:          runnerSize,
		RunnerGroupID: int64(runnerGroupID),
	}

	if v, ok := d.GetOk("maximum_runners"); ok {
		maxRunners, _ := v.(int)
		req.MaximumRunners = new(int64(maxRunners))
	}

	if v, okExists := d.GetOkExists("public_ip_enabled"); okExists { //nolint:staticcheck // SA1019: d.GetOkExists is deprecated but necessary for bool fields
		publicIPEnabled, _ := v.(bool)
		req.EnableStaticIP = new(publicIPEnabled)
	}

	if v, ok := d.GetOk("image_version"); ok {
		runnerImageVersion, _ := v.(string)
		req.Image.Version = new(runnerImageVersion)
	}

	if v, ok := d.GetOk("image_gen"); ok {
		useRunnerForImageGen, _ := v.(bool)
		req.ImageGen = new(useRunnerForImageGen)
	}

	runner, _, err := client.Actions.CreateHostedRunner(ctx, orgName, req)
	if err != nil {
		if _, ok := errors.AsType[*github.AcceptedError](err); !ok {
			return diag.FromErr(err)
		}
	}

	d.SetId(strconv.Itoa(int(runner.GetID())))

	if err := d.Set("status", runner.GetStatus()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("platform", runner.GetPlatform()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_active_on", runner.GetLastActiveOn().GoString()); err != nil {
		return diag.FromErr(err)
	}

	if machineSizeDetails := runner.GetMachineSizeDetails(); machineSizeDetails != nil {
		if err := d.Set("machine_size_details", flattenMachineSizeDetails(machineSizeDetails)); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("maximum_runners", runner.GetMaximumRunners()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("public_ips", flattenPublicIPs(runner.GetPublicIPs())); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsHostedRunnerRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	runnerIDStr := d.Id()
	runnerID, err := strconv.ParseInt(runnerIDStr, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	runner, _, err := client.Actions.GetHostedRunner(ctx, orgName, runnerID)
	if err != nil {
		if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == http.StatusNotFound {
			tflog.Warn(ctx, "Removing hosted runner from state because it no longer exists in GitHub", map[string]any{"runner_id": runnerIDStr})
			d.SetId("")
			return nil
		}
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
	if err := d.Set("last_active_on", runner.GetLastActiveOn().GoString()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("public_ip_enabled", runner.GetPublicIPEnabled()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("image", flattenImage(runner.GetImageDetails())); err != nil {
		return diag.FromErr(err)
	}

	if machineSizeDetails := runner.GetMachineSizeDetails(); machineSizeDetails != nil {
		if err := d.Set("size", machineSizeDetails.GetID()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("machine_size_details", flattenMachineSizeDetails(machineSizeDetails)); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("runner_group_id", runner.GetRunnerGroupID()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("maximum_runners", runner.GetMaximumRunners()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("public_ips", flattenPublicIPs(runner.GetPublicIPs())); err != nil {
		return diag.FromErr(err)
	}

	// TODO: Uncomment when go-github supports image_gen field in the HostedRunner struct
	// if err := d.Set("image_gen", runner.GetImageGen()); err != nil {
	// 	return diag.FromErr(err)
	// }

	return nil
}

func resourceGithubActionsHostedRunnerUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	runnerID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	req := github.UpdateHostedRunnerRequest{}

	if d.HasChange("name") {
		name, _ := d.Get("name").(string)
		req.Name = new(name)
	}
	if d.HasChange("size") {
		size, _ := d.Get("size").(string)
		req.Size = new(size)
	}
	if d.HasChange("runner_group_id") {
		runnerGroupID, _ := d.Get("runner_group_id").(int)
		req.RunnerGroupID = new(int64(runnerGroupID))
	}
	if d.HasChange("maximum_runners") {
		maximumRunners, _ := d.Get("maximum_runners").(int)
		req.MaximumRunners = new(int64(maximumRunners))
	}
	if d.HasChange("public_ip_enabled") {
		publicIPEnabled, _ := d.Get("public_ip_enabled").(bool)
		req.EnableStaticIP = new(publicIPEnabled)
	}
	if d.HasChange("image_version") {
		imageVersion, _ := d.Get("image_version").(string)
		req.ImageVersion = new(imageVersion)
	}

	// Only update the runner if any of the configured fields have changed
	if d.HasChanges("name", "size", "runner_group_id", "maximum_runners", "public_ip_enabled", "image_version") {

		runner, _, err := client.Actions.UpdateHostedRunner(ctx, orgName, runnerID, req)
		if err != nil {
			if _, ok := errors.AsType[*github.AcceptedError](err); !ok {
				return diag.FromErr(err)
			}
		}

		if err := d.Set("status", runner.GetStatus()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("platform", runner.GetPlatform()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("last_active_on", runner.GetLastActiveOn().GoString()); err != nil {
			return diag.FromErr(err)
		}

		if machineSizeDetails := runner.GetMachineSizeDetails(); machineSizeDetails != nil {
			if err := d.Set("machine_size_details", flattenMachineSizeDetails(machineSizeDetails)); err != nil {
				return diag.FromErr(err)
			}
		}

		if err := d.Set("maximum_runners", runner.GetMaximumRunners()); err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("public_ips", flattenPublicIPs(runner.GetPublicIPs())); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubActionsHostedRunnerDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	runnerID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	runner, resp, err := client.Actions.DeleteHostedRunner(ctx, orgName, runnerID)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil
		}
		if _, ok := errors.AsType[*github.AcceptedError](err); ok {
			return diag.FromErr(waitForRunnerDeletion(ctx, client, orgName, runner.GetID(), d.Timeout(schema.TimeoutDelete)))
		}
		return diag.FromErr(err)
	}

	if resp != nil && resp.StatusCode == http.StatusAccepted {
		return diag.FromErr(waitForRunnerDeletion(ctx, client, orgName, runner.GetID(), d.Timeout(schema.TimeoutDelete)))
	}

	return nil
}

func waitForRunnerDeletion(ctx context.Context, client *github.Client, orgName string, runnerID int64, timeout time.Duration) error {
	conf := &retry.StateChangeConf{
		Pending: []string{"deleting", "active", "ready", "provisioning"},
		Target:  []string{"deleted"},
		Refresh: func() (any, string, error) {
			_, resp, err := client.Actions.GetHostedRunner(ctx, orgName, runnerID)

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
