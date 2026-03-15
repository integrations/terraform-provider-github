package github

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var organizationNetworkConfigurationNamePattern = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

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
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(
					validation.StringLenBetween(1, 100),
					validation.StringMatch(
						organizationNetworkConfigurationNamePattern,
						"name may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'",
					),
				)),
				Description: "Name of the network configuration. Must be between 1 and 100 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.",
			},
			"compute_service": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "none",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"none", "actions"}, false)),
				Description:      "The hosted compute service to use for the network configuration. Can be one of: 'none', 'actions'. Defaults to 'none'.",
			},
			"network_settings_ids": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "An array containing exactly one network settings ID. A network settings resource can only be associated with one network configuration at a time.",
			},
			"created_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the network configuration was created.",
			},
		},
	}
}

func resourceGithubOrganizationNetworkConfigurationCreate(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	configuration, _, err := client.Organizations.CreateNetworkConfiguration(ctx, orgName, buildOrganizationNetworkConfigurationRequest(d))
	if err != nil {
		return err
	}

	d.SetId(configuration.GetID())

	return resourceGithubOrganizationNetworkConfigurationRead(d, meta)
}

func resourceGithubOrganizationNetworkConfigurationRead(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	networkConfigurationID := d.Id()
	ctx := context.WithValue(context.Background(), ctxId, networkConfigurationID)

	configuration, resp, err := client.Organizations.GetNetworkConfiguration(ctx, orgName, networkConfigurationID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			log.Printf("[WARN] Removing organization network configuration %s from state because it no longer exists in GitHub", networkConfigurationID)
			d.SetId("")
			return nil
		}

		return err
	}

	if resp != nil && resp.StatusCode == http.StatusNotModified {
		return nil
	}

	return setOrganizationNetworkConfigurationState(d, configuration)
}

func resourceGithubOrganizationNetworkConfigurationUpdate(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	networkConfigurationID := d.Id()
	ctx := context.WithValue(context.Background(), ctxId, networkConfigurationID)

	_, _, err := client.Organizations.UpdateNetworkConfiguration(ctx, orgName, networkConfigurationID, buildOrganizationNetworkConfigurationRequest(d))
	if err != nil {
		return err
	}

	return resourceGithubOrganizationNetworkConfigurationRead(d, meta)
}

func resourceGithubOrganizationNetworkConfigurationDelete(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	_, err := client.Organizations.DeleteNetworkConfigurations(ctx, orgName, d.Id())
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}

		return err
	}

	return nil
}

func expandOrganizationNetworkConfigurationComputeService(computeService string) *github.ComputeService {
	service := github.ComputeService(computeService)
	return &service
}

func expandOrganizationNetworkSettingsIDs(networkSettingsIDs []any) []string {
	ids := make([]string, 0, len(networkSettingsIDs))
	for _, networkSettingsID := range networkSettingsIDs {
		ids = append(ids, networkSettingsID.(string))
	}

	return ids
}

func buildOrganizationNetworkConfigurationRequest(d *schema.ResourceData) github.NetworkConfigurationRequest {
	return github.NetworkConfigurationRequest{
		Name:               github.Ptr(d.Get("name").(string)),
		ComputeService:     expandOrganizationNetworkConfigurationComputeService(d.Get("compute_service").(string)),
		NetworkSettingsIDs: expandOrganizationNetworkSettingsIDs(d.Get("network_settings_ids").([]any)),
	}
}

func setOrganizationNetworkConfigurationState(d *schema.ResourceData, configuration *github.NetworkConfiguration) error {
	if err := d.Set("name", configuration.GetName()); err != nil {
		return err
	}
	if configuration.ComputeService != nil {
		if err := d.Set("compute_service", string(*configuration.ComputeService)); err != nil {
			return err
		}
	}
	if err := d.Set("network_settings_ids", configuration.NetworkSettingsIDs); err != nil {
		return err
	}
	if configuration.CreatedOn != nil {
		if err := d.Set("created_on", configuration.CreatedOn.Format("2006-01-02T15:04:05Z07:00")); err != nil {
			return err
		}
	}

	return nil
}
