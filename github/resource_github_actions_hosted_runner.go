package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
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

		Description: "This resource allows you to create and manage GitHub-hosted runners within your GitHub organization. You must have admin access to an organization to use this resource.",

		CustomizeDiff: customdiff.All(resourceGithubActionsHostedRunnerValidation),

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
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "Image configuration for the hosted runner. Cannot be changed after creation.",
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
			},
			"size": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Machine size for the hosted runner (e.g., `4-core`, `8-core`). Can be updated to scale the runner. To list available sizes, use the GitHub API: `GET /orgs/{org}/actions/hosted-runners/machine-sizes`.",
			},
			"runner_group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the runner group to assign this runner to.",
			},
			"maximum_runners": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Description:      "Maximum number of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit costs.",
			},
			"public_ip_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable static public IP for the runner. Note there are account limits. To list limits, use the GitHub API: `GET /orgs/{org}/actions/hosted-runners/limits`. Defaults to false.",
			},
			"image_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of the runner image to deploy. This is only relevant for runners using custom images.",
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
			"machine_size_details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Detailed machine size specifications.",
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
			},
			"platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Platform of the runner.",
			},
			"public_ips": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of public IP ranges assigned to this runner.",
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
			},
			"last_active_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the runner was last active.",
			},
		},
	}
}

func resourceGithubActionsHostedRunnerCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	err := checkOrganization(m)
	if err != nil {
		return diag.FromErr(err)
	}

	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	runnerName, _ := d.Get("name").(string)
	runnerImageDetails, _ := d.Get("image").([]any)
	runnerSize, _ := d.Get("size").(string)
	runnerGroupID, _ := d.Get("runner_group_id").(int)
	expandedImageDetails, err := expandImage(runnerImageDetails)
	if err != nil {
		return diag.FromErr(err)
	}

	req := github.CreateHostedRunnerRequest{
		Name:          runnerName,
		Image:         *expandedImageDetails,
		Size:          runnerSize,
		RunnerGroupID: int64(runnerGroupID),
	}

	if v, ok := d.GetOk("maximum_runners"); ok {
		maxRunners, _ := v.(int)
		req.MaximumRunners = new(int64(maxRunners))
	}

	publicIPEnabled, _ := d.Get("public_ip_enabled").(bool)
	req.EnableStaticIP = new(publicIPEnabled)

	if v, ok := d.GetOk("image_version"); ok {
		runnerImageVersion, _ := v.(string)
		req.Image.Version = new(runnerImageVersion)
	}

	useRunnerForImageGen, _ := d.Get("image_gen").(bool)
	req.ImageGen = new(useRunnerForImageGen)

	runner, _, err := client.Actions.CreateHostedRunner(ctx, orgName, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(int(runner.GetID())))

	if err := d.Set("status", runner.GetStatus()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("platform", runner.GetPlatform()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_active_on", runner.GetLastActiveOn().Format(time.RFC3339)); err != nil {
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

func resourceGithubActionsHostedRunnerRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	err := checkOrganization(m)
	if err != nil {
		return diag.FromErr(err)
	}

	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

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
	if err := d.Set("last_active_on", runner.GetLastActiveOn().Format(time.RFC3339)); err != nil {
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

func resourceGithubActionsHostedRunnerUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	err := checkOrganization(m)
	if err != nil {
		return diag.FromErr(err)
	}

	// Only update the runner if any of the configured fields have changed
	if !d.HasChanges("name", "size", "runner_group_id", "maximum_runners", "public_ip_enabled", "image_version") {
		return nil
	}

	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	runnerID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	name, _ := d.Get("name").(string)
	size, _ := d.Get("size").(string)
	runnerGroupID, _ := d.Get("runner_group_id").(int)
	maximumRunners, _ := d.Get("maximum_runners").(int)
	publicIPEnabled, _ := d.Get("public_ip_enabled").(bool)

	req := github.UpdateHostedRunnerRequest{
		Name:           new(name),
		Size:           new(size),
		RunnerGroupID:  new(int64(runnerGroupID)),
		MaximumRunners: new(int64(maximumRunners)),
		EnableStaticIP: new(publicIPEnabled),
	}

	// image_version is only settable for runners with a custom image, so we only include it in the request if it has changed
	if d.HasChange("image_version") {
		imageVersion, _ := d.Get("image_version").(string)
		req.ImageVersion = new(imageVersion)
	}

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
	if err := d.Set("last_active_on", runner.GetLastActiveOn().Format(time.RFC3339)); err != nil {
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

func resourceGithubActionsHostedRunnerDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	err := checkOrganization(m)
	if err != nil {
		return diag.FromErr(err)
	}

	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	runnerID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	_, resp, err := client.Actions.DeleteHostedRunner(ctx, orgName, runnerID)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil
		}
		if _, ok := errors.AsType[*github.AcceptedError](err); ok {
			return diag.FromErr(waitForRunnerDeletion(ctx, client, orgName, runnerID, d.Timeout(schema.TimeoutDelete)))
		}
		return diag.FromErr(err)
	}

	if resp != nil && resp.StatusCode == http.StatusAccepted {
		return diag.FromErr(waitForRunnerDeletion(ctx, client, orgName, runnerID, d.Timeout(schema.TimeoutDelete)))
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

func expandImage(imageList []any) (*github.HostedRunnerImage, error) {
	if len(imageList) == 0 {
		return &github.HostedRunnerImage{}, nil
	}
	runnerImage := &github.HostedRunnerImage{}

	imageMap, ok := imageList[0].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("expected map[string]any, got %T", imageList[0])
	}

	if id, ok := imageMap["id"].(string); ok {
		runnerImage.ID = id
	}
	if source, ok := imageMap["source"].(string); ok {
		runnerImage.Source = source
	}

	return runnerImage, nil
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

func resourceGithubActionsHostedRunnerValidation(ctx context.Context, d *schema.ResourceDiff, meta any) error {
	// validates that image_version is only set for images with source "custom"
	if d.HasChange("image_version") {
		imageList, ok := d.GetOk("image")
		imageSeq, okSeq := imageList.([]any)
		if !ok || !okSeq || len(imageSeq) == 0 {
			return fmt.Errorf("`image_version` can only be set when `image` is configured")
		}

		imageMap, ok := imageSeq[0].(map[string]any)
		if !ok {
			return fmt.Errorf("unexpected type for image: %T", imageSeq[0])
		}

		source, ok := imageMap["source"].(string)
		if !ok || source != "custom" {
			return fmt.Errorf("`image_version` can only be set when `image[0].source` is 'custom'")
		}
	}
	return nil
}
