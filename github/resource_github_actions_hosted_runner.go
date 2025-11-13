package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the hosted runner.",
			},
			"image": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The image ID.",
						},
						"source": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "github",
							Description: "The image source.",
						},
					},
				},
				Description: "Image configuration for the hosted runner.",
			},
			"size": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Machine size (e.g., '4-core', '8-core').",
			},
			"runner_group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The runner group ID.",
			},
			"maximum_runners": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum number of runners to scale up to.",
			},
			"enable_static_ip": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable static IP.",
			},
			// Computed fields
			"id": {
				Type:        schema.TypeInt,
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
		},
	}
}

// expandImage converts the Terraform image configuration to a map for API calls
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

// flattenImage converts the API image response to Terraform format
func flattenImage(image map[string]interface{}) []interface{} {
	if image == nil {
		return []interface{}{}
	}

	result := make(map[string]interface{})
	if id, ok := image["id"].(string); ok {
		result["id"] = id
	}
	if source, ok := image["source"].(string); ok {
		result["source"] = source
	}

	return []interface{}{result}
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

	if v, ok := d.GetOk("enable_static_ip"); ok {
		payload["enable_static_ip"] = v.(bool)
	}

	// Create HTTP request
	req, err := client.NewRequest("POST", fmt.Sprintf("orgs/%s/actions/hosted-runners", orgName), payload)
	if err != nil {
		return err
	}

	var runner map[string]interface{}
	resp, err := client.Do(ctx, req, &runner)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Set the ID
	if id, ok := runner["id"].(float64); ok {
		d.SetId(strconv.Itoa(int(id)))
	} else {
		return fmt.Errorf("failed to get runner ID from response")
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
	ctx := context.Background()

	runnerID := d.Id()

	// Create GET request
	req, err := client.NewRequest("GET", fmt.Sprintf("orgs/%s/actions/hosted-runners/%s", orgName, runnerID), nil)
	if err != nil {
		return err
	}

	var runner map[string]interface{}
	resp, err := client.Do(ctx, req, &runner)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			log.Printf("[WARN] Removing hosted runner %s from state because it no longer exists in GitHub", runnerID)
			d.SetId("")
			return nil
		}
		return err
	}

	// Set computed attributes
	if name, ok := runner["name"].(string); ok {
		d.Set("name", name)
	}
	if status, ok := runner["status"].(string); ok {
		d.Set("status", status)
	}
	if platform, ok := runner["platform"].(string); ok {
		d.Set("platform", platform)
	}
	if image, ok := runner["image"].(map[string]interface{}); ok {
		d.Set("image", flattenImage(image))
	}
	if size, ok := runner["size"].(string); ok {
		d.Set("size", size)
	}
	if runnerGroupID, ok := runner["runner_group_id"].(float64); ok {
		d.Set("runner_group_id", int(runnerGroupID))
	}
	if maxRunners, ok := runner["maximum_runners"].(float64); ok {
		d.Set("maximum_runners", int(maxRunners))
	}
	if staticIP, ok := runner["enable_static_ip"].(bool); ok {
		d.Set("enable_static_ip", staticIP)
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
	ctx := context.Background()

	runnerID := d.Id()

	// Build update payload
	payload := make(map[string]interface{})

	if d.HasChange("name") {
		payload["name"] = d.Get("name").(string)
	}
	if d.HasChange("image") {
		payload["image"] = expandImage(d.Get("image").([]interface{}))
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
	if d.HasChange("enable_static_ip") {
		payload["enable_static_ip"] = d.Get("enable_static_ip").(bool)
	}

	// Create PATCH request
	req, err := client.NewRequest("PATCH", fmt.Sprintf("orgs/%s/actions/hosted-runners/%s", orgName, runnerID), payload)
	if err != nil {
		return err
	}

	var runner map[string]interface{}
	resp, err := client.Do(ctx, req, &runner)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
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
			// Already deleted
			return nil
		}
		return err
	}

	// Handle async deletion (202 Accepted)
	if resp.StatusCode == http.StatusAccepted {
		log.Printf("[DEBUG] Hosted runner %s deletion accepted, polling for completion", runnerID)
		return waitForRunnerDeletion(client, orgName, runnerID, ctx)
	}

	return nil
}

// waitForRunnerDeletion polls the API until the runner is deleted or times out
func waitForRunnerDeletion(client *http.Client, orgName, runnerID string, ctx context.Context) error {
	timeout := time.After(10 * time.Minute)
	interval := 30 * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	attempts := 0
	maxInterval := 2 * time.Minute

	for {
		select {
		case <-timeout:
			return fmt.Errorf("timeout waiting for hosted runner %s to be deleted after 10 minutes", runnerID)
		case <-ticker.C:
			attempts++

			// Check if runner still exists
			req, err := client.NewRequest("GET", fmt.Sprintf("orgs/%s/actions/hosted-runners/%s", orgName, runnerID), nil)
			if err != nil {
				return err
			}

			resp, err := client.Do(ctx, req, nil)
			if err != nil {
				// If 404, runner is deleted successfully
				if resp != nil && resp.StatusCode == http.StatusNotFound {
					log.Printf("[DEBUG] Hosted runner %s successfully deleted after %d attempts", runnerID, attempts)
					return nil
				}
				return err
			}

			// Runner still exists, continue polling with exponential backoff
			log.Printf("[DEBUG] Hosted runner %s still exists, continuing to poll (attempt %d)", runnerID, attempts)

			// Increase interval with exponential backoff, capped at maxInterval
			newInterval := time.Duration(float64(interval) * 1.5)
			if newInterval > maxInterval {
				newInterval = maxInterval
			}
			interval = newInterval
			ticker.Reset(interval)
		}
	}
}
