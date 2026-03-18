package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseSecurityConfiguration() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages a code security configuration for a GitHub Enterprise.",
		CreateContext: resourceGithubEnterpriseSecurityConfigurationCreate,
		ReadContext:   resourceGithubEnterpriseSecurityConfigurationRead,
		UpdateContext: resourceGithubEnterpriseSecurityConfigurationUpdate,
		DeleteContext: resourceGithubEnterpriseSecurityConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseSecurityConfigurationImport,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"configuration_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The numeric ID of the code security configuration.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the code security configuration.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "A description of the code security configuration.",
			},
			"advanced_security": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The advanced security configuration for the code security configuration. Can be one of 'enabled', 'disabled'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled",
				}, false)),
			},
			"dependency_graph": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The dependency graph configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"dependency_graph_autosubmit_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The dependency graph autosubmit action configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"dependency_graph_autosubmit_action_options": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The dependency graph autosubmit action options for the code security configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labeled_runners": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to use labeled runners for the dependency graph autosubmit action.",
						},
					},
				},
			},
			"dependabot_alerts": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The dependabot alerts configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"dependabot_security_updates": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The dependabot security updates configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"code_scanning_default_setup": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The code scanning default setup configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"code_scanning_default_setup_options": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The code scanning default setup options for the code security configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"runner_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "The type of runner to use for code scanning default setup. Can be one of 'standard', 'labeled'.",
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"standard", "labeled"}, false)),
						},
						"runner_label": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The label of the runner to use for code scanning default setup.",
						},
					},
				},
			},
			"code_scanning_delegated_alert_dismissal": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The code scanning delegated alert dismissal configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"code_scanning_options": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The code scanning options for the code security configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_advanced": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to allow advanced security for code scanning.",
						},
					},
				},
			},
			"code_security": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The code security configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The secret scanning configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_push_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The secret scanning push protection configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_validity_checks": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The secret scanning validity checks configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_non_provider_patterns": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The secret scanning non provider patterns configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_generic_secrets": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The secret scanning generic secrets configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_delegated_alert_dismissal": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The secret scanning delegated alert dismissal configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The secret protection configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"private_vulnerability_reporting": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The private vulnerability reporting configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"enforcement": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The enforcement configuration for the code security configuration. Can be one of 'enforced', 'unenforced'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enforced", "unenforced",
				}, false)),
			},
			"target_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target type of the code security configuration.",
			},
		},
	}
}

func resourceGithubEnterpriseSecurityConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterprise := d.Get("enterprise_slug").(string)
	name := d.Get("name").(string)

	tflog.Debug(ctx, "Creating enterprise code security configuration", map[string]any{
		"enterprise": enterprise,
		"name":       name,
	})

	config := expandCodeSecurityConfigurationCommon(d)

	configuration, _, err := client.Enterprise.CreateCodeSecurityConfiguration(ctx, enterprise, config)
	if err != nil {
		tflog.Error(ctx, "Failed to create enterprise code security configuration", map[string]any{
			"enterprise": enterprise,
			"name":       name,
			"error":      err.Error(),
		})
		return diag.FromErr(err)
	}

	id, err := buildID(enterprise, strconv.FormatInt(configuration.GetID(), 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if diags := setCodeSecurityConfigurationState(d, configuration); diags.HasError() {
		return diags
	}

	tflog.Info(ctx, "Created enterprise code security configuration", map[string]any{
		"enterprise": enterprise,
		"name":       name,
		"id":         configuration.GetID(),
	})

	return nil
}

func resourceGithubEnterpriseSecurityConfigurationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterprise := d.Get("enterprise_slug").(string)
	id := int64(d.Get("configuration_id").(int))

	tflog.Trace(ctx, "Reading enterprise code security configuration", map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	configuration, _, err := client.Enterprise.GetCodeSecurityConfiguration(ctx, enterprise, id)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing enterprise code security configuration from state because it no longer exists in GitHub", map[string]any{
					"enterprise": enterprise,
					"id":         id,
				})
				d.SetId("")
				return nil
			}
		}
		tflog.Error(ctx, "Failed to read enterprise code security configuration", map[string]any{
			"enterprise": enterprise,
			"id":         id,
			"error":      err.Error(),
		})
		return diag.FromErr(err)
	}

	if diags := setCodeSecurityConfigurationState(d, configuration); diags.HasError() {
		return diags
	}

	tflog.Trace(ctx, "Successfully read enterprise code security configuration", map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	return nil
}

func resourceGithubEnterpriseSecurityConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterprise := d.Get("enterprise_slug").(string)
	id := int64(d.Get("configuration_id").(int))

	tflog.Debug(ctx, "Updating enterprise code security configuration", map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	config := expandCodeSecurityConfigurationCommon(d)

	configuration, _, err := client.Enterprise.UpdateCodeSecurityConfiguration(ctx, enterprise, id, config)
	if err != nil {
		tflog.Error(ctx, "Failed to update enterprise code security configuration", map[string]any{
			"enterprise": enterprise,
			"id":         id,
			"error":      err.Error(),
		})
		return diag.FromErr(err)
	}

	if diags := setCodeSecurityConfigurationState(d, configuration); diags.HasError() {
		return diags
	}

	tflog.Info(ctx, "Updated enterprise code security configuration", map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	return nil
}

func resourceGithubEnterpriseSecurityConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterprise := d.Get("enterprise_slug").(string)
	id := int64(d.Get("configuration_id").(int))

	tflog.Debug(ctx, "Deleting enterprise code security configuration", map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	_, err := client.Enterprise.DeleteCodeSecurityConfiguration(ctx, enterprise, id)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Enterprise code security configuration already deleted", map[string]any{
				"enterprise": enterprise,
				"id":         id,
			})
			return nil
		}
		tflog.Error(ctx, "Failed to delete enterprise code security configuration", map[string]any{
			"enterprise": enterprise,
			"id":         id,
			"error":      err.Error(),
		})
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Deleted enterprise code security configuration", map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	return nil
}

func resourceGithubEnterpriseSecurityConfigurationImport(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	enterpriseSlug, configIDStr, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>:<configuration_id>. Parse error: %w", err)
	}

	configID, err := strconv.ParseInt(configIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration_id %q: %w", configIDStr, err)
	}

	id, err := buildID(enterpriseSlug, configIDStr)
	if err != nil {
		return nil, err
	}
	d.SetId(id)

	if err = d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}

	if err = d.Set("configuration_id", int(configID)); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

