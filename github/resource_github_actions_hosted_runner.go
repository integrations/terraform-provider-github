package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsHostedRunner() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsHostedRunnerCreate,
		Read:   resourceGithubActionsHostedRunnerRead,
		Update: resourceGithubActionsHostedRunnerUpdate,
		Delete: resourceGithubActionsHostedRunnerDelete,
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
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "github",
							ValidateFunc: validation.StringInSlice([]string{"github", "partner", "custom"}, false),
							Description:  "The image source (github, partner, or custom).",
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
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "Maximum number of runners to scale up to.",
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

func expandImage(imageList []interface{}) map[string]interface{} {
	if len(imageList) == 0 {
		return nil
	}

	imageMap := imageList[0].(map[string]interface{})
	result := make(map[string]interface{})

	if id, ok := imageMap["id"].(string); ok {
		result["id"] = id
	}
	if source, ok := imageMap["source"].(string); ok {
		result["source"] = source
	}

	return result
}

func flattenImage(image map[string]interface{}) []interface{} {
	if image == nil {
		return []interface{}{}
	}

	result := make(map[string]interface{})

	// Handle id as either string or number
	if id, ok := image["id"].(string); ok {
		result["id"] = id
	} else if id, ok := image["id"].(float64); ok {
		result["id"] = fmt.Sprintf("%.0f", id)
	}

	if source, ok := image["source"].(string); ok {
		result["source"] = source
	}
	if size, ok := image["size"].(float64); ok {
		result["size_gb"] = int(size)
	}

	return []interface{}{result}
}

func flattenMachineSizeDetails(details map[string]interface{}) []interface{} {
	if details == nil {
		return []interface{}{}
	}

	result := make(map[string]interface{})
	if id, ok := details["id"].(string); ok {
		result["id"] = id
	}
	if cpuCores, ok := details["cpu_cores"].(float64); ok {
		result["cpu_cores"] = int(cpuCores)
	}
	if memoryGB, ok := details["memory_gb"].(float64); ok {
		result["memory_gb"] = int(memoryGB)
	}
	if storageGB, ok := details["storage_gb"].(float64); ok {
		result["storage_gb"] = int(storageGB)
	}

	return []interface{}{result}
}

func flattenPublicIPs(ips []interface{}) []interface{} {
	if ips == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0, len(ips))
	for _, ip := range ips {
		ipMap, ok := ip.(map[string]interface{})
		if !ok {
			continue
		}

		ipResult := make(map[string]interface{})
		if enabled, ok := ipMap["enabled"].(bool); ok {
			ipResult["enabled"] = enabled
		}
		if prefix, ok := ipMap["prefix"].(string); ok {
			ipResult["prefix"] = prefix
		}
		if length, ok := ipMap["length"].(float64); ok {
			ipResult["length"] = int(length)
		}
		result = append(result, ipResult)
	}

	return result
}

func resourceGithubActionsHostedRunnerCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	// Build request payload
	payload := map[string]interface{}{
		"name":            d.Get("name").(string),
		"image":           expandImage(d.Get("image").([]interface{})),
		"size":            d.Get("size").(string),
		"runner_group_id": d.Get("runner_group_id").(int),
	}

	if v, ok := d.GetOk("maximum_runners"); ok {
		payload["maximum_runners"] = v.(int)
	}

	if v, ok := d.GetOk("public_ip_enabled"); ok {
		payload["enable_static_ip"] = v.(bool)
	}

	if v, ok := d.GetOk("image_version"); ok {
		payload["image_version"] = v.(string)
	}

	if v, ok := d.GetOk("image_gen"); ok {
		payload["image_gen"] = v.(bool)
	}

	// Create HTTP request
	req, err := client.NewRequest("POST", fmt.Sprintf("orgs/%s/actions/hosted-runners", orgName), payload)
	if err != nil {
		return err
	}

	var runner map[string]interface{}
	_, err = client.Do(ctx, req, &runner)
	if err != nil {
		if _, ok := err.(*github.AcceptedError); !ok {
			return err
		}
	}

	if runner == nil {
		return fmt.Errorf("no runner data returned from API")
	}

	// Set the ID
	if id, ok := runner["id"].(float64); ok {
		d.SetId(strconv.Itoa(int(id)))
	} else {
		return fmt.Errorf("failed to get runner ID from response: %+v", runner)
	}

	return resourceGithubActionsHostedRunnerRead(d, meta)
}

func resourceGithubActionsHostedRunnerRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	runnerID := d.Id()
	ctx := context.WithValue(context.Background(), ctxId, runnerID)

	// Create GET request
	req, err := client.NewRequest("GET", fmt.Sprintf("orgs/%s/actions/hosted-runners/%s", orgName, runnerID), nil)
	if err != nil {
		return err
	}

	var runner map[string]interface{}
	_, err = client.Do(ctx, req, &runner)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing hosted runner %s from state because it no longer exists in GitHub", runnerID)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if runner == nil {
		return fmt.Errorf("no runner data returned from API")
	}

	if name, ok := runner["name"].(string); ok {
		d.Set("name", name)
	}
	if status, ok := runner["status"].(string); ok {
		d.Set("status", status)
	}
	if platform, ok := runner["platform"].(string); ok {
		d.Set("platform", platform)
	}
	if lastActiveOn, ok := runner["last_active_on"].(string); ok {
		d.Set("last_active_on", lastActiveOn)
	}
	if publicIPEnabled, ok := runner["public_ip_enabled"].(bool); ok {
		d.Set("public_ip_enabled", publicIPEnabled)
	}

	if image, ok := runner["image"].(map[string]interface{}); ok {
		d.Set("image", flattenImage(image))
	}

	if machineSizeDetails, ok := runner["machine_size_details"].(map[string]interface{}); ok {
		d.Set("size", machineSizeDetails["id"])
		d.Set("machine_size_details", flattenMachineSizeDetails(machineSizeDetails))
	}

	if runnerGroupID, ok := runner["runner_group_id"].(float64); ok {
		d.Set("runner_group_id", int(runnerGroupID))
	}

	if maxRunners, ok := runner["maximum_runners"].(float64); ok {
		d.Set("maximum_runners", int(maxRunners))
	}

	if publicIPs, ok := runner["public_ips"].([]interface{}); ok {
		d.Set("public_ips", flattenPublicIPs(publicIPs))
	}

	if imageGen, ok := runner["image_gen"].(bool); ok {
		d.Set("image_gen", imageGen)
	}

	return nil
}

func resourceGithubActionsHostedRunnerUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	runnerID := d.Id()
	ctx := context.WithValue(context.Background(), ctxId, runnerID)

	payload := make(map[string]interface{})

	if d.HasChange("name") {
		payload["name"] = d.Get("name").(string)
	}
	if d.HasChange("size") {
		payload["size"] = d.Get("size").(string)
	}
	if d.HasChange("runner_group_id") {
		payload["runner_group_id"] = d.Get("runner_group_id").(int)
	}
	if d.HasChange("maximum_runners") {
		payload["maximum_runners"] = d.Get("maximum_runners").(int)
	}
	if d.HasChange("public_ip_enabled") {
		payload["enable_static_ip"] = d.Get("public_ip_enabled").(bool)
	}
	if d.HasChange("image_version") {
		payload["image_version"] = d.Get("image_version").(string)
	}

	if len(payload) == 0 {
		return resourceGithubActionsHostedRunnerRead(d, meta)
	}

	// Create PATCH request
	req, err := client.NewRequest("PATCH", fmt.Sprintf("orgs/%s/actions/hosted-runners/%s", orgName, runnerID), payload)
	if err != nil {
		return err
	}

	var runner map[string]interface{}
	_, err = client.Do(ctx, req, &runner)
	if err != nil {
		if _, ok := err.(*github.AcceptedError); !ok {
			return err
		}
	}

	return resourceGithubActionsHostedRunnerRead(d, meta)
}

func resourceGithubActionsHostedRunnerDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	runnerID := d.Id()

	// Send DELETE request
	req, err := client.NewRequest("DELETE", fmt.Sprintf("orgs/%s/actions/hosted-runners/%s", orgName, runnerID), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(ctx, req, nil)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil
		}
		if _, ok := err.(*github.AcceptedError); ok {
			return waitForRunnerDeletion(ctx, client, orgName, runnerID, d.Timeout(schema.TimeoutDelete))
		}
		return err
	}

	if resp != nil && resp.StatusCode == http.StatusAccepted {
		return waitForRunnerDeletion(ctx, client, orgName, runnerID, d.Timeout(schema.TimeoutDelete))
	}

	return nil
}

func waitForRunnerDeletion(ctx context.Context, client *github.Client, orgName, runnerID string, timeout time.Duration) error {
	conf := &retry.StateChangeConf{
		Pending: []string{"deleting", "active"},
		Target:  []string{"deleted"},
		Refresh: func() (interface{}, string, error) {
			req, err := client.NewRequest("GET", fmt.Sprintf("orgs/%s/actions/hosted-runners/%s", orgName, runnerID), nil)
			if err != nil {
				return nil, "", err
			}

			resp, err := client.Do(ctx, req, nil)
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
