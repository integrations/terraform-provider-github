package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseActionsHostedRunner() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubEnterpriseActionsHostedRunnerCreate,
		Read:   resourceGithubEnterpriseActionsHostedRunnerRead,
		Update: resourceGithubEnterpriseActionsHostedRunnerUpdate,
		Delete: resourceGithubEnterpriseActionsHostedRunnerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
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

func resourceGithubEnterpriseActionsHostedRunnerCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	ctx := context.Background()

	// Build request payload
	payload := map[string]any{
		"name":            d.Get("name").(string),
		"image":           expandImage(d.Get("image").([]any)),
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
	req, err := client.NewRequest("POST", fmt.Sprintf("enterprises/%s/actions/hosted-runners", enterpriseSlug), payload)
	if err != nil {
		return err
	}

	var runner map[string]any
	_, err = client.Do(ctx, req, &runner)
	if err != nil {
		var acceptedErr *github.AcceptedError
		if !errors.As(err, &acceptedErr) {
			return err
		}
	}

	if runner == nil {
		return fmt.Errorf("no runner data returned from API")
	}

	// Set the ID
	if id, ok := runner["id"].(float64); ok {
		d.SetId(fmt.Sprintf("%s/%d", enterpriseSlug, int(id)))
	} else {
		return fmt.Errorf("failed to get runner ID from response: %+v", runner)
	}

	return resourceGithubEnterpriseActionsHostedRunnerRead(d, meta)
}

func resourceGithubEnterpriseActionsHostedRunnerRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	enterpriseSlug, runnerID, err := parseEnterpriseRunnerID(d.Id())
	if err != nil {
		return err
	}

	// Create GET request
	req, err := client.NewRequest("GET", fmt.Sprintf("enterprises/%s/actions/hosted-runners/%s", enterpriseSlug, runnerID), nil)
	if err != nil {
		return err
	}

	var runner map[string]any
	_, err = client.Do(ctx, req, &runner)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing enterprise hosted runner %s from state because it no longer exists in GitHub", runnerID)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if runner == nil {
		return fmt.Errorf("no runner data returned from API")
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return err
	}

	if name, ok := runner["name"].(string); ok {
		if err := d.Set("name", name); err != nil {
			return err
		}
	}
	if status, ok := runner["status"].(string); ok {
		if err := d.Set("status", status); err != nil {
			return err
		}
	}
	if platform, ok := runner["platform"].(string); ok {
		if err := d.Set("platform", platform); err != nil {
			return err
		}
	}
	if lastActiveOn, ok := runner["last_active_on"].(string); ok {
		if err := d.Set("last_active_on", lastActiveOn); err != nil {
			return err
		}
	}
	if publicIPEnabled, ok := runner["public_ip_enabled"].(bool); ok {
		if err := d.Set("public_ip_enabled", publicIPEnabled); err != nil {
			return err
		}
	}

	if image, ok := runner["image"].(map[string]any); ok {
		if err := d.Set("image", flattenImage(image)); err != nil {
			return err
		}
	}

	if machineSizeDetails, ok := runner["machine_size_details"].(map[string]any); ok {
		if err := d.Set("size", machineSizeDetails["id"]); err != nil {
			return err
		}
		if err := d.Set("machine_size_details", flattenMachineSizeDetails(machineSizeDetails)); err != nil {
			return err
		}
	}

	if runnerGroupID, ok := runner["runner_group_id"].(float64); ok {
		if err := d.Set("runner_group_id", int(runnerGroupID)); err != nil {
			return err
		}
	}

	if maxRunners, ok := runner["maximum_runners"].(float64); ok {
		if err := d.Set("maximum_runners", int(maxRunners)); err != nil {
			return err
		}
	}

	if publicIPs, ok := runner["public_ips"].([]any); ok {
		if err := d.Set("public_ips", flattenPublicIPs(publicIPs)); err != nil {
			return err
		}
	}

	if imageGen, ok := runner["image_gen"].(bool); ok {
		if err := d.Set("image_gen", imageGen); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubEnterpriseActionsHostedRunnerUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	enterpriseSlug, runnerID, err := parseEnterpriseRunnerID(d.Id())
	if err != nil {
		return err
	}

	payload := make(map[string]any)

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
		return resourceGithubEnterpriseActionsHostedRunnerRead(d, meta)
	}

	// Create PATCH request
	req, err := client.NewRequest("PATCH", fmt.Sprintf("enterprises/%s/actions/hosted-runners/%s", enterpriseSlug, runnerID), payload)
	if err != nil {
		return err
	}

	var runner map[string]any
	_, err = client.Do(ctx, req, &runner)
	if err != nil {
		var acceptedErr *github.AcceptedError
		if !errors.As(err, &acceptedErr) {
			return err
		}
	}

	return resourceGithubEnterpriseActionsHostedRunnerRead(d, meta)
}

func resourceGithubEnterpriseActionsHostedRunnerDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug, runnerID, err := parseEnterpriseRunnerID(d.Id())
	if err != nil {
		return err
	}

	// Send DELETE request
	req, err := client.NewRequest("DELETE", fmt.Sprintf("enterprises/%s/actions/hosted-runners/%s", enterpriseSlug, runnerID), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(ctx, req, nil)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil
		}
		var acceptedErr *github.AcceptedError
		if errors.As(err, &acceptedErr) {
			return waitForEnterpriseRunnerDeletion(ctx, client, enterpriseSlug, runnerID, d.Timeout(schema.TimeoutDelete))
		}
		return err
	}

	if resp != nil && resp.StatusCode == http.StatusAccepted {
		return waitForEnterpriseRunnerDeletion(ctx, client, enterpriseSlug, runnerID, d.Timeout(schema.TimeoutDelete))
	}

	return nil
}

func waitForEnterpriseRunnerDeletion(ctx context.Context, client *github.Client, enterpriseSlug, runnerID string, timeout time.Duration) error {
	conf := &retry.StateChangeConf{
		Pending: []string{"deleting", "active"},
		Target:  []string{"deleted"},
		Refresh: func() (any, string, error) {
			req, err := client.NewRequest("GET", fmt.Sprintf("enterprises/%s/actions/hosted-runners/%s", enterpriseSlug, runnerID), nil)
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

func parseEnterpriseRunnerID(id string) (string, string, error) {
	parts := make([]string, 0)
	for i, part := range regexp.MustCompile(`/`).Split(id, -1) {
		if i == 0 {
			parts = append(parts, part)
		} else {
			parts = append(parts, part)
			break
		}
	}

	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid ID format: %s (expected: enterprise_slug/runner_id)", id)
	}

	return parts[0], parts[1], nil
}
