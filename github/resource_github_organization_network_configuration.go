package github

import (
	"context"
	"net/http"
	"regexp"
	"time"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var organizationNetworkConfigurationNamePattern = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

func resourceGithubOrganizationNetworkConfiguration() *schema.Resource {
	return &schema.Resource{
		Description:   "This resource allows you to create and manage hosted compute network configurations for a GitHub organization.",
		CreateContext: resourceGithubOrganizationNetworkConfigurationCreate,
		ReadContext:   resourceGithubOrganizationNetworkConfigurationRead,
		UpdateContext: resourceGithubOrganizationNetworkConfigurationUpdate,
		DeleteContext: resourceGithubOrganizationNetworkConfigurationDelete,
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

func resourceGithubOrganizationNetworkConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "organization", meta.(*Owner).name)

	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	computeService := github.ComputeService(d.Get("compute_service").(string))
	networkSettingsIDs := []string{d.Get("network_settings_ids").([]any)[0].(string)}

	tflog.Debug(ctx, "Creating organization network configuration", map[string]any{
		"name":                 d.Get("name").(string),
		"compute_service":      d.Get("compute_service").(string),
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := client.Organizations.CreateNetworkConfiguration(ctx, orgName, github.NetworkConfigurationRequest{
		Name:               github.Ptr(d.Get("name").(string)),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(configuration.GetID())
	if err := setOrganizationNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationNetworkConfigurationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())
	ctx = tflog.SetField(ctx, "organization", meta.(*Owner).name)

	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	networkConfigurationID := d.Id()
	ctx = context.WithValue(ctx, ctxId, networkConfigurationID)

	configuration, resp, err := client.Organizations.GetNetworkConfiguration(ctx, orgName, networkConfigurationID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Organization network configuration not found, removing from state", map[string]any{"id": networkConfigurationID})
			d.SetId("")
			return nil
		}

		return diag.FromErr(err)
	}

	if resp != nil && resp.StatusCode == http.StatusNotModified {
		return nil
	}

	if err := setOrganizationNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationNetworkConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())
	ctx = tflog.SetField(ctx, "organization", meta.(*Owner).name)

	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	networkConfigurationID := d.Id()
	ctx = context.WithValue(ctx, ctxId, networkConfigurationID)
	computeService := github.ComputeService(d.Get("compute_service").(string))
	networkSettingsIDs := []string{d.Get("network_settings_ids").([]any)[0].(string)}

	tflog.Debug(ctx, "Updating organization network configuration", map[string]any{
		"name":                 d.Get("name").(string),
		"compute_service":      d.Get("compute_service").(string),
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := client.Organizations.UpdateNetworkConfiguration(ctx, orgName, networkConfigurationID, github.NetworkConfigurationRequest{
		Name:               github.Ptr(d.Get("name").(string)),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setOrganizationNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationNetworkConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())
	ctx = tflog.SetField(ctx, "organization", meta.(*Owner).name)

	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	tflog.Debug(ctx, "Deleting organization network configuration")
	_, err := client.Organizations.DeleteNetworkConfigurations(ctx, orgName, d.Id())
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}

		return diag.FromErr(err)
	}

	return nil
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
		if err := d.Set("created_on", configuration.CreatedOn.Format(time.RFC3339)); err != nil {
			return err
		}
	}

	return nil
}
