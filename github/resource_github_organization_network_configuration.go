package github

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/google/go-github/v77/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationNetworkConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationNetworkConfigurationCreate,
		Read:   resourceGithubOrganizationNetworkConfigurationRead,
		Update: resourceGithubOrganizationNetworkConfigurationUpdate,
		Delete: resourceGithubOrganizationNetworkConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 100),
					validation.StringMatch(
						regexp.MustCompile(`^[a-zA-Z0-9._-]+$`),
						"name may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'",
					),
				),
				Description: "Name of the network configuration. Must be between 1 and 100 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.",
			},
			"compute_service": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: validation.StringInSlice([]string{"none", "actions"}, false),
				Description:  "The hosted compute service to use for the network configuration. Can be one of: 'none', 'actions'. Defaults to 'none'.",
			},
			"network_settings_ids": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The identifier of the network settings to use for the network configuration. Exactly one network settings must be specified.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network configuration ID.",
			},
			"created_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the network configuration was created.",
			},
		},
	}
}

func resourceGithubOrganizationNetworkConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	// Build request
	computeService := github.ComputeService(d.Get("compute_service").(string))

	networkSettingsIDs := []string{}
	for _, id := range d.Get("network_settings_ids").([]interface{}) {
		networkSettingsIDs = append(networkSettingsIDs, id.(string))
	}

	req := github.NetworkConfigurationRequest{
		Name:               github.String(d.Get("name").(string)),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	}

	log.Printf("[DEBUG] Creating network configuration: %s", d.Get("name").(string))
	configuration, _, err := client.Organizations.CreateNetworkConfiguration(ctx, orgName, req)
	if err != nil {
		return err
	}

	d.SetId(configuration.GetID())
	return resourceGithubOrganizationNetworkConfigurationRead(d, meta)
}

func resourceGithubOrganizationNetworkConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	networkID := d.Id()
	ctx := context.WithValue(context.Background(), ctxId, networkID)

	log.Printf("[DEBUG] Reading network configuration: %s", networkID)
	configuration, resp, err := client.Organizations.GetNetworkConfiguration(ctx, orgName, networkID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing network configuration %s from state because it no longer exists in GitHub", networkID)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if resp.StatusCode == http.StatusNotModified {
		return nil
	}

	d.Set("name", configuration.GetName())
	if configuration.ComputeService != nil {
		d.Set("compute_service", string(*configuration.ComputeService))
	}
	d.Set("network_settings_ids", configuration.NetworkSettingsIDs)
	if configuration.CreatedOn != nil {
		d.Set("created_on", configuration.CreatedOn.Format("2006-01-02T15:04:05Z07:00"))
	}

	return nil
}

func resourceGithubOrganizationNetworkConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	networkID := d.Id()
	ctx := context.WithValue(context.Background(), ctxId, networkID)

	// Build request with changed fields
	computeService := github.ComputeService(d.Get("compute_service").(string))

	networkSettingsIDs := []string{}
	for _, id := range d.Get("network_settings_ids").([]interface{}) {
		networkSettingsIDs = append(networkSettingsIDs, id.(string))
	}

	req := github.NetworkConfigurationRequest{
		Name:               github.String(d.Get("name").(string)),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	}

	log.Printf("[DEBUG] Updating network configuration: %s", networkID)
	_, _, err = client.Organizations.UpdateNetworkConfiguration(ctx, orgName, networkID, req)
	if err != nil {
		return err
	}

	return resourceGithubOrganizationNetworkConfigurationRead(d, meta)
}

func resourceGithubOrganizationNetworkConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	networkID := d.Id()
	ctx := context.Background()

	log.Printf("[DEBUG] Deleting network configuration: %s", networkID)
	_, err = client.Organizations.DeleteNetworkConfigurations(ctx, orgName, networkID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				return nil
			}
		}
		return err
	}

	return nil
}
