package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationNetworkConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubOrganizationNetworkConfigurationCreate,
		ReadContext:   resourceGithubOrganizationNetworkConfigurationRead,
		UpdateContext: resourceGithubOrganizationNetworkConfigurationUpdate,
		DeleteContext: resourceGithubOrganizationNetworkConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Description: "Resource to manage a hosted compute network configuration for a GitHub organization. " +
			"The organization is determined by the provider's `owner` setting.\n\n" +
			"A network configuration associates an Azure virtual network with GitHub-hosted runners. " +
			"Assign it to a runner group with [`github_actions_runner_group.network_configuration_id`](actions_runner_group).\n\n" +
			"First create an Azure `GitHub.Network/networkSettings` resource registered against the same organization. " +
			"Pass its `GitHubId`, not its Azure resource ID, in `network_settings_ids`.\n\n" +
			"See the [GitHub REST API documentation](https://docs.github.com/en/rest/orgs/network-configurations#create-a-hosted-compute-network-configuration-for-an-organization) for required permissions.",

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
					Type:             schema.TypeString,
					ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotWhiteSpace),
				},
				Description: "A list containing exactly one nonempty network settings GitHub ID registered against this organization. A network settings resource can only be associated with one network configuration at a time.",
			},
			"created_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the network configuration was created, in RFC3339 format. Empty when GitHub does not return a creation timestamp.",
			},
		},
	}
}

func resourceGithubOrganizationNetworkConfigurationCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	orgName := meta.name
	name, _ := d.Get("name").(string)
	computeServiceName, _ := d.Get("compute_service").(string)
	computeService := github.ComputeService(computeServiceName)
	ids, _ := d.Get("network_settings_ids").([]any)
	networkSettingsIDs := expandStringList(ids)

	ctx = tflog.SetField(ctx, "organization", orgName)
	tflog.Debug(ctx, "Creating organization network configuration", map[string]any{
		"name":                 name,
		"compute_service":      computeService,
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := meta.v3client.Organizations.CreateNetworkConfiguration(ctx, orgName, github.NetworkConfigurationRequest{
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

func resourceGithubOrganizationNetworkConfigurationRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	orgName := meta.name

	ctx = tflog.SetField(ctx, "organization", orgName)

	configuration, _, err := meta.v3client.Organizations.GetNetworkConfiguration(ctx, orgName, d.Id())
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

func resourceGithubOrganizationNetworkConfigurationUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	orgName := meta.name
	name, _ := d.Get("name").(string)
	computeServiceName, _ := d.Get("compute_service").(string)
	computeService := github.ComputeService(computeServiceName)
	ids, _ := d.Get("network_settings_ids").([]any)
	networkSettingsIDs := expandStringList(ids)

	ctx = tflog.SetField(ctx, "organization", orgName)
	tflog.Debug(ctx, "Updating organization network configuration", map[string]any{
		"name":                 name,
		"compute_service":      computeService,
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := meta.v3client.Organizations.UpdateNetworkConfiguration(ctx, orgName, d.Id(), github.NetworkConfigurationRequest{
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

func resourceGithubOrganizationNetworkConfigurationDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	orgName := meta.name

	ctx = tflog.SetField(ctx, "organization", orgName)
	tflog.Debug(ctx, "Deleting organization network configuration")

	if _, err := meta.v3client.Organizations.DeleteNetworkConfigurations(ctx, orgName, d.Id()); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}

		return diag.FromErr(err)
	}

	return nil
}
