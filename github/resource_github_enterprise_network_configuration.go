package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseNetworkConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubEnterpriseNetworkConfigurationCreate,
		ReadContext:   resourceGithubEnterpriseNetworkConfigurationRead,
		UpdateContext: resourceGithubEnterpriseNetworkConfigurationUpdate,
		DeleteContext: resourceGithubEnterpriseNetworkConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseNetworkConfigurationImport,
		},

		Description: "Resource to manage a hosted compute network configuration for a GitHub enterprise.\n\n" +
			"A network configuration associates an Azure virtual network with GitHub-hosted runners. " +
			"Assign it to a runner group with [`github_enterprise_actions_runner_group.network_configuration_id`](enterprise_actions_runner_group).\n\n" +
			"First create an Azure `GitHub.Network/networkSettings` resource registered against the same enterprise. " +
			"Pass its `GitHubId`, not its Azure resource ID, in `network_settings_ids`.\n\n" +
			"This API does not support GitHub App tokens or fine-grained personal access tokens. " +
			"See the [GitHub REST API documentation](https://docs.github.com/en/enterprise-cloud@latest/rest/enterprise-admin/network-configurations#create-a-hosted-compute-network-configuration-for-an-enterprise) for authentication requirements.",

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotWhiteSpace),
				Description:      "The slug of the enterprise. Changing this forces a new resource to be created.",
			},
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
				Description: "A list containing exactly one nonempty network settings GitHub ID registered against this enterprise. A network settings resource can only be associated with one network configuration at a time.",
			},
			"created_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the network configuration was created, in RFC3339 format. Empty when GitHub does not return a creation timestamp.",
			},
		},
	}
}

func resourceGithubEnterpriseNetworkConfigurationCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	enterpriseSlug, _ := d.Get("enterprise_slug").(string)
	name, _ := d.Get("name").(string)
	computeServiceName, _ := d.Get("compute_service").(string)
	computeService := github.ComputeService(computeServiceName)
	ids, _ := d.Get("network_settings_ids").([]any)
	networkSettingsIDs := expandStringList(ids)

	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)
	tflog.Debug(ctx, "Creating enterprise network configuration", map[string]any{
		"name":                 name,
		"compute_service":      computeService,
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := meta.v3client.Enterprise.CreateEnterpriseNetworkConfiguration(ctx, enterpriseSlug, github.NetworkConfigurationRequest{
		Name:               new(name),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	})
	if err != nil {
		return diag.FromErr(networkSettingsScopeError(err, "enterprise"))
	}

	d.SetId(configuration.GetID())
	if err := setNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseNetworkConfigurationRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	enterpriseSlug, _ := d.Get("enterprise_slug").(string)

	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)

	configuration, _, err := meta.v3client.Enterprise.GetEnterpriseNetworkConfiguration(ctx, enterpriseSlug, d.Id())
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing enterprise network configuration from state because it no longer exists in GitHub")
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

func resourceGithubEnterpriseNetworkConfigurationUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	enterpriseSlug, _ := d.Get("enterprise_slug").(string)
	name, _ := d.Get("name").(string)
	computeServiceName, _ := d.Get("compute_service").(string)
	computeService := github.ComputeService(computeServiceName)
	ids, _ := d.Get("network_settings_ids").([]any)
	networkSettingsIDs := expandStringList(ids)

	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)
	tflog.Debug(ctx, "Updating enterprise network configuration", map[string]any{
		"name":                 name,
		"compute_service":      computeService,
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := meta.v3client.Enterprise.UpdateEnterpriseNetworkConfiguration(ctx, enterpriseSlug, d.Id(), github.NetworkConfigurationRequest{
		Name:               new(name),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	})
	if err != nil {
		return diag.FromErr(networkSettingsScopeError(err, "enterprise"))
	}

	if err := setNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseNetworkConfigurationDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	enterpriseSlug, _ := d.Get("enterprise_slug").(string)

	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)
	tflog.Debug(ctx, "Deleting enterprise network configuration")

	if _, err := meta.v3client.Enterprise.DeleteEnterpriseNetworkConfiguration(ctx, enterpriseSlug, d.Id()); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}

		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseNetworkConfigurationImport(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	enterpriseSlug, networkConfigurationID, ok := strings.Cut(d.Id(), "/")
	if !ok || strings.TrimSpace(enterpriseSlug) == "" || strings.TrimSpace(networkConfigurationID) == "" || strings.Contains(networkConfigurationID, "/") {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>/<network_configuration_id>")
	}

	d.SetId(networkConfigurationID)
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
