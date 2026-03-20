package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseNetworkConfiguration() *schema.Resource {
	return &schema.Resource{
		Description:   "This resource allows you to create and manage hosted compute network configurations for a GitHub enterprise.",
		CreateContext: resourceGithubEnterpriseNetworkConfigurationCreate,
		ReadContext:   resourceGithubEnterpriseNetworkConfigurationRead,
		UpdateContext: resourceGithubEnterpriseNetworkConfigurationUpdate,
		DeleteContext: resourceGithubEnterpriseNetworkConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseNetworkConfigurationImport,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
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

func resourceGithubEnterpriseNetworkConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	enterpriseSlug := d.Get("enterprise_slug").(string)
	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)

	client := meta.(*Owner).v3client
	computeService := github.ComputeService(d.Get("compute_service").(string))
	networkSettingsIDs := []string{d.Get("network_settings_ids").([]any)[0].(string)}

	tflog.Debug(ctx, "Creating enterprise network configuration", map[string]any{
		"name":                 d.Get("name").(string),
		"compute_service":      d.Get("compute_service").(string),
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := client.Enterprise.CreateEnterpriseNetworkConfiguration(ctx, enterpriseSlug, github.NetworkConfigurationRequest{
		Name:               github.Ptr(d.Get("name").(string)),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	})
	if err != nil {
		return enterpriseNetworkConfigurationDiagnostics(err)
	}

	d.SetId(configuration.GetID())
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := setEnterpriseNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseNetworkConfigurationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	enterpriseSlug := d.Get("enterprise_slug").(string)
	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)
	ctx = tflog.SetField(ctx, "id", d.Id())

	client := meta.(*Owner).v3client
	networkConfigurationID := d.Id()
	ctx = context.WithValue(ctx, ctxId, networkConfigurationID)

	configuration, resp, err := client.Enterprise.GetEnterpriseNetworkConfiguration(ctx, enterpriseSlug, networkConfigurationID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Enterprise network configuration not found, removing from state", map[string]any{"id": networkConfigurationID})
			d.SetId("")
			return nil
		}

		return diag.FromErr(err)
	}

	if resp != nil && resp.StatusCode == http.StatusNotModified {
		return nil
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := setEnterpriseNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseNetworkConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	enterpriseSlug := d.Get("enterprise_slug").(string)
	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)
	ctx = tflog.SetField(ctx, "id", d.Id())

	client := meta.(*Owner).v3client
	networkConfigurationID := d.Id()
	ctx = context.WithValue(ctx, ctxId, networkConfigurationID)
	computeService := github.ComputeService(d.Get("compute_service").(string))
	networkSettingsIDs := []string{d.Get("network_settings_ids").([]any)[0].(string)}

	tflog.Debug(ctx, "Updating enterprise network configuration", map[string]any{
		"name":                 d.Get("name").(string),
		"compute_service":      d.Get("compute_service").(string),
		"network_settings_ids": networkSettingsIDs,
	})

	configuration, _, err := client.Enterprise.UpdateEnterpriseNetworkConfiguration(ctx, enterpriseSlug, networkConfigurationID, github.NetworkConfigurationRequest{
		Name:               github.Ptr(d.Get("name").(string)),
		ComputeService:     &computeService,
		NetworkSettingsIDs: networkSettingsIDs,
	})
	if err != nil {
		return enterpriseNetworkConfigurationDiagnostics(err)
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := setEnterpriseNetworkConfigurationState(d, configuration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseNetworkConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	enterpriseSlug := d.Get("enterprise_slug").(string)
	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)
	ctx = tflog.SetField(ctx, "id", d.Id())

	client := meta.(*Owner).v3client

	tflog.Debug(ctx, "Deleting enterprise network configuration")
	_, err := client.Enterprise.DeleteEnterpriseNetworkConfiguration(ctx, enterpriseSlug, d.Id())
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}

		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseNetworkConfigurationImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>/<network_configuration_id>")
	}

	enterpriseSlug, networkConfigurationID := parts[0], parts[1]
	d.SetId(networkConfigurationID)
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func setEnterpriseNetworkConfigurationState(d *schema.ResourceData, configuration *github.NetworkConfiguration) error {
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

func enterpriseNetworkConfigurationDiagnostics(err error) diag.Diagnostics {
	var ghErr *github.ErrorResponse
	if errors.As(err, &ghErr) && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusUnprocessableEntity {
		return diag.FromErr(fmt.Errorf("%w. if you are using Azure private networking, ensure the provided network settings GitHubId matches the enterprise scope; enterprise-level configurations may fail when the backing GitHub.Network/networkSettings resource was created with an organization databaseId", err))
	}

	return diag.FromErr(err)
}
