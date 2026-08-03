package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// networkConfigurationNamePattern mirrors the name constraint the REST API enforces on hosted
// compute network configurations so that invalid names are rejected at plan time.
var networkConfigurationNamePattern = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

func resourceGithubOrganizationNetworkConfiguration() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages a hosted compute network configuration for a GitHub organization.",
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
						networkConfigurationNamePattern,
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
				Description:      "The hosted compute service the network configuration supports. Can be one of: 'none', 'actions'. Defaults to 'none'.",
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
				Description: "Timestamp of when the network configuration was created, in RFC3339 format.",
			},
		},
	}
}

func resourceGithubOrganizationNetworkConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	name := d.Get("name").(string)
	computeService := github.ComputeService(d.Get("compute_service").(string))
	networkSettingsIDs := expandNetworkSettingsIDs(d)

	ctx = tflog.SetField(ctx, "organization", orgName)
	tflog.Debug(ctx, "Creating organization network configuration", map[string]any{
		"name":                 name,
		"compute_service":      computeService,
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := client.Organizations.CreateNetworkConfiguration(ctx, orgName, github.NetworkConfigurationRequest{
		Name:               new(name),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	})
	if err != nil {
		return diag.FromErr(networkSettingsScopeError(err, "organization"))
	}

	d.SetId(configuration.GetID())
	if err := setNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationNetworkConfigurationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	ctx = tflog.SetField(ctx, "organization", orgName)

	configuration, _, err := client.Organizations.GetNetworkConfiguration(ctx, orgName, d.Id())
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing organization network configuration from state because it no longer exists in GitHub")
				d.SetId("")
				return nil
			}
		}

		return diag.FromErr(err)
	}

	if err := setNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationNetworkConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	name := d.Get("name").(string)
	computeService := github.ComputeService(d.Get("compute_service").(string))
	networkSettingsIDs := expandNetworkSettingsIDs(d)

	ctx = tflog.SetField(ctx, "organization", orgName)
	tflog.Debug(ctx, "Updating organization network configuration", map[string]any{
		"name":                 name,
		"compute_service":      computeService,
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := client.Organizations.UpdateNetworkConfiguration(ctx, orgName, d.Id(), github.NetworkConfigurationRequest{
		Name:               new(name),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	})
	if err != nil {
		return diag.FromErr(networkSettingsScopeError(err, "organization"))
	}

	if err := setNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationNetworkConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	ctx = tflog.SetField(ctx, "organization", orgName)
	tflog.Debug(ctx, "Deleting organization network configuration")

	if _, err := client.Organizations.DeleteNetworkConfigurations(ctx, orgName, d.Id()); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}

		return diag.FromErr(err)
	}

	return nil
}

// expandNetworkSettingsIDs reads network_settings_ids, which the schema constrains to exactly
// one element.
func expandNetworkSettingsIDs(d *schema.ResourceData) []string {
	ids := d.Get("network_settings_ids").([]any)
	networkSettingsIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		networkSettingsIDs = append(networkSettingsIDs, id.(string))
	}

	return networkSettingsIDs
}

// setNetworkConfigurationState writes the attributes shared by the organization and enterprise
// network configuration resources. The REST API returns an identical payload for both scopes.
func setNetworkConfigurationState(d *schema.ResourceData, configuration *github.NetworkConfiguration) error {
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

// networkSettingsScopeError annotates the 422 the API returns when the referenced network
// settings resource belongs to a different scope than the configuration being written. Azure
// issues distinct GitHub IDs for organization and enterprise network settings, and mixing them
// up is the most common cause of this error.
func networkSettingsScopeError(err error, scope string) error {
	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusUnprocessableEntity {
		return fmt.Errorf("%w. verify the network settings ID belongs to the same %s: Azure GitHub.Network/networkSettings resources are registered against a single organization or enterprise and cannot be shared across scopes", err, scope)
	}

	return err
}
