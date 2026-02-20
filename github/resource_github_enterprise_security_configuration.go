package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v82/github"
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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the code security configuration.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
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

	tflog.Debug(ctx, fmt.Sprintf("Creating enterprise code security configuration: %s/%s", enterprise, name), map[string]any{
		"enterprise": enterprise,
		"name":       name,
	})

	config := expandEnterpriseCodeSecurityConfiguration(d)

	configuration, _, err := client.Enterprise.CreateCodeSecurityConfiguration(ctx, enterprise, config)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create enterprise code security configuration: %s/%s", enterprise, name), map[string]any{
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

	tflog.Info(ctx, fmt.Sprintf("Created enterprise code security configuration: %s/%s (ID: %d)", enterprise, name, configuration.GetID()), map[string]any{
		"enterprise": enterprise,
		"name":       name,
		"id":         configuration.GetID(),
	})

	return resourceGithubEnterpriseSecurityConfigurationRead(ctx, d, meta)
}

func resourceGithubEnterpriseSecurityConfigurationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterprise, idStr, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, fmt.Sprintf("Reading enterprise code security configuration: %s/%d", enterprise, id), map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	configuration, _, err := client.Enterprise.GetCodeSecurityConfiguration(ctx, enterprise, id)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, fmt.Sprintf("Removing enterprise code security configuration %s/%d from state because it no longer exists in GitHub", enterprise, id), map[string]any{
					"enterprise": enterprise,
					"id":         id,
				})
				d.SetId("")
				return nil
			}
		}
		tflog.Error(ctx, fmt.Sprintf("Failed to read enterprise code security configuration: %s/%d", enterprise, id), map[string]any{
			"enterprise": enterprise,
			"id":         id,
			"error":      err.Error(),
		})
		return diag.FromErr(err)
	}

	if err = d.Set("enterprise_slug", enterprise); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("name", configuration.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("description", configuration.Description); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("advanced_security", configuration.GetAdvancedSecurity()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("dependency_graph", configuration.GetDependencyGraph()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("dependency_graph_autosubmit_action", configuration.GetDependencyGraphAutosubmitAction()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dependency_graph_autosubmit_action_options", flattenDependencyGraphAutosubmitActionOptions(configuration.DependencyGraphAutosubmitActionOptions)); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("dependabot_alerts", configuration.GetDependabotAlerts()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("dependabot_security_updates", configuration.GetDependabotSecurityUpdates()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("code_scanning_default_setup", configuration.GetCodeScanningDefaultSetup()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("code_scanning_default_setup_options", flattenCodeScanningDefaultSetupOptions(configuration.CodeScanningDefaultSetupOptions)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("code_scanning_options", flattenCodeScanningOptions(configuration.CodeScanningOptions)); err != nil {
		return diag.FromErr(err)
	}
	codeSec := configuration.GetCodeSecurity()
	if codeSec == "" {
		codeSec = "disabled"
	}
	if err = d.Set("code_security", codeSec); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("secret_scanning", configuration.GetSecretScanning()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("secret_scanning_push_protection", configuration.GetSecretScanningPushProtection()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("secret_scanning_validity_checks", configuration.GetSecretScanningValidityChecks()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("secret_scanning_non_provider_patterns", configuration.GetSecretScanningNonProviderPatterns()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("secret_scanning_generic_secrets", configuration.GetSecretScanningGenericSecrets()); err != nil {
		return diag.FromErr(err)
	}
	secretProt := configuration.GetSecretProtection()
	if secretProt == "" {
		secretProt = "disabled"
	}
	if err = d.Set("secret_protection", secretProt); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("private_vulnerability_reporting", configuration.GetPrivateVulnerabilityReporting()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("enforcement", configuration.GetEnforcement()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("target_type", configuration.GetTargetType()); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, fmt.Sprintf("Successfully read enterprise code security configuration: %s/%d", enterprise, id), map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	return nil
}

func resourceGithubEnterpriseSecurityConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterprise, idStr, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating enterprise code security configuration: %s/%d", enterprise, id), map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	config := expandEnterpriseCodeSecurityConfiguration(d)

	_, _, err = client.Enterprise.UpdateCodeSecurityConfiguration(ctx, enterprise, id, config)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to update enterprise code security configuration: %s/%d", enterprise, id), map[string]any{
			"enterprise": enterprise,
			"id":         id,
			"error":      err.Error(),
		})
		return diag.FromErr(err)
	}

	tflog.Info(ctx, fmt.Sprintf("Updated enterprise code security configuration: %s/%d", enterprise, id), map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	return resourceGithubEnterpriseSecurityConfigurationRead(ctx, d, meta)
}

func resourceGithubEnterpriseSecurityConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterprise, idStr, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting enterprise code security configuration: %s/%d", enterprise, id), map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	_, err = client.Enterprise.DeleteCodeSecurityConfiguration(ctx, enterprise, id)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, fmt.Sprintf("Enterprise code security configuration %s/%d already deleted", enterprise, id), map[string]any{
				"enterprise": enterprise,
				"id":         id,
			})
			return nil
		}
		tflog.Error(ctx, fmt.Sprintf("Failed to delete enterprise code security configuration: %s/%d", enterprise, id), map[string]any{
			"enterprise": enterprise,
			"id":         id,
			"error":      err.Error(),
		})
		return diag.FromErr(err)
	}

	tflog.Info(ctx, fmt.Sprintf("Deleted enterprise code security configuration: %s/%d", enterprise, id), map[string]any{
		"enterprise": enterprise,
		"id":         id,
	})

	return nil
}

func resourceGithubEnterpriseSecurityConfigurationImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	enterpriseSlug, configID, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>:<configuration_id>. Parse error: %w", err)
	}

	id, err := buildID(enterpriseSlug, configID)
	if err != nil {
		return nil, err
	}
	d.SetId(id)

	if err = d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func expandEnterpriseCodeSecurityConfiguration(d *schema.ResourceData) github.CodeSecurityConfiguration {
	config := github.CodeSecurityConfiguration{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	// Only set optional fields if they are explicitly configured
	if val, ok := d.GetOk("advanced_security"); ok {
		config.AdvancedSecurity = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("dependency_graph"); ok {
		config.DependencyGraph = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("dependency_graph_autosubmit_action"); ok {
		config.DependencyGraphAutosubmitAction = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("dependabot_alerts"); ok {
		config.DependabotAlerts = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("dependabot_security_updates"); ok {
		config.DependabotSecurityUpdates = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("code_scanning_default_setup"); ok {
		config.CodeScanningDefaultSetup = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("code_security"); ok {
		config.CodeSecurity = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("secret_scanning"); ok {
		config.SecretScanning = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("secret_scanning_push_protection"); ok {
		config.SecretScanningPushProtection = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("secret_scanning_validity_checks"); ok {
		config.SecretScanningValidityChecks = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("secret_scanning_non_provider_patterns"); ok {
		config.SecretScanningNonProviderPatterns = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("secret_scanning_generic_secrets"); ok {
		config.SecretScanningGenericSecrets = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("secret_protection"); ok {
		config.SecretProtection = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("private_vulnerability_reporting"); ok {
		config.PrivateVulnerabilityReporting = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("enforcement"); ok {
		config.Enforcement = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("dependency_graph_autosubmit_action_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			autosubmitOpts := optionsList[0].(map[string]any)
			config.DependencyGraphAutosubmitActionOptions = &github.DependencyGraphAutosubmitActionOptions{
				LabeledRunners: github.Ptr(autosubmitOpts["labeled_runners"].(bool)),
			}
		}
	}

	if val, ok := d.GetOk("code_scanning_default_setup_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			setupOpts := optionsList[0].(map[string]any)
			config.CodeScanningDefaultSetupOptions = &github.CodeScanningDefaultSetupOptions{
				RunnerType: setupOpts["runner_type"].(string),
			}
			if runnerLabel, ok := setupOpts["runner_label"].(string); ok && runnerLabel != "" {
				config.CodeScanningDefaultSetupOptions.RunnerLabel = github.Ptr(runnerLabel)
			}
		}
	}

	if val, ok := d.GetOk("code_scanning_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			scanOpts := optionsList[0].(map[string]any)
			config.CodeScanningOptions = &github.CodeScanningOptions{
				AllowAdvanced: github.Ptr(scanOpts["allow_advanced"].(bool)),
			}
		}
	}

	return config
}
