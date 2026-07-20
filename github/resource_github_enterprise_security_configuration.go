package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseSecurityConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubEnterpriseSecurityConfigurationCreate,
		ReadContext:   resourceGithubEnterpriseSecurityConfigurationRead,
		UpdateContext: resourceGithubEnterpriseSecurityConfigurationUpdate,
		DeleteContext: resourceGithubEnterpriseSecurityConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseSecurityConfigurationImport,
		},

		Description: "Resource to manage a GitHub code security configuration for an enterprise.",

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Enterprise slug.",
			},
			"configuration_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Numeric ID of the code security configuration.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the code security configuration.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the code security configuration.",
			},
			"advanced_security": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Advanced security setting. Can be one of 'enabled', 'disabled'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled",
				}, false)),
			},
			"dependency_graph": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Dependency graph setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"dependency_graph_autosubmit_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Dependency graph autosubmit action setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"dependency_graph_autosubmit_action_options": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Dependency graph autosubmit action options.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labeled_runners": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to use labeled runners for the dependency graph autosubmit action.",
						},
					},
				},
			},
			"dependabot_alerts": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Dependabot alerts setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"dependabot_security_updates": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Dependabot security updates setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"code_scanning_default_setup": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Code scanning default setup. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"code_scanning_default_setup_options": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Code scanning default setup options.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"runner_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Type of runner to use for code scanning default setup. Can be one of 'standard', 'labeled'.",
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"standard", "labeled"}, false)),
						},
						"runner_label": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Label of the runner to use for code scanning default setup.",
						},
					},
				},
			},
			"code_scanning_delegated_alert_dismissal": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Code scanning delegated alert dismissal setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"code_scanning_options": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Code scanning options.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_advanced": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to allow advanced security for code scanning.",
						},
					},
				},
			},
			"code_security": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Code security setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret scanning setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_push_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret scanning push protection setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_validity_checks": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret scanning validity checks setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_non_provider_patterns": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret scanning non-provider patterns setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_generic_secrets": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret scanning generic secrets setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_delegated_alert_dismissal": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret scanning delegated alert dismissal setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret protection setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"private_vulnerability_reporting": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Private vulnerability reporting setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"enforcement": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enforcement setting. Can be one of 'enforced', 'unenforced'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enforced", "unenforced",
				}, false)),
			},
			"target_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target type of the code security configuration.",
			},
		},
	}
}

func resourceGithubEnterpriseSecurityConfigurationCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	enterprise, _ := d.Get("enterprise_slug").(string)
	name, _ := d.Get("name").(string)

	tflog.Debug(ctx, "Creating enterprise code security configuration", map[string]any{"enterprise": enterprise, "name": name})

	config := resourceGithubEnterpriseSecurityConfigurationExpand(d)

	configuration, _, err := client.Enterprise.CreateCodeSecurityConfiguration(ctx, enterprise, config)
	if err != nil {
		tflog.Error(ctx, "Failed to create enterprise code security configuration", map[string]any{"enterprise": enterprise, "name": name, "error": err.Error()})
		return diag.FromErr(err)
	}

	id, err := buildID(enterprise, strconv.FormatInt(configuration.GetID(), 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	// Optional (non-computed) attributes are persisted from config by the SDK; only the
	// computed values need to be written here.
	if err := d.Set("configuration_id", int(configuration.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target_type", configuration.GetTargetType()); err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Created enterprise code security configuration", map[string]any{"enterprise": enterprise, "name": name, "id": configuration.GetID()})

	return nil
}

func resourceGithubEnterpriseSecurityConfigurationRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	enterprise, _ := d.Get("enterprise_slug").(string)
	configID, _ := d.Get("configuration_id").(int)
	id := int64(configID)

	tflog.Trace(ctx, "Reading enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id})

	configuration, _, err := client.Enterprise.GetCodeSecurityConfiguration(ctx, enterprise, id)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing enterprise code security configuration from state because it no longer exists in GitHub", map[string]any{"enterprise": enterprise, "id": id})
			d.SetId("")
			return nil
		}
		tflog.Error(ctx, "Failed to read enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id, "error": err.Error()})
		return diag.FromErr(err)
	}

	if err := d.Set("name", configuration.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("configuration_id", int(configuration.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target_type", configuration.GetTargetType()); err != nil {
		return diag.FromErr(err)
	}
	// Reconcile only the attributes under management (already present in state). GitHub
	// server-assigns a value to every omitted attribute, so refreshing those would produce a
	// perpetual diff; unmanaged values can be read via the data source instead.
	managed := []struct {
		key   string
		value any
	}{
		{"description", configuration.Description},
		{"advanced_security", configuration.GetAdvancedSecurity()},
		{"dependency_graph", configuration.GetDependencyGraph()},
		{"dependency_graph_autosubmit_action", configuration.GetDependencyGraphAutosubmitAction()},
		{"dependency_graph_autosubmit_action_options", flattenDependencyGraphAutosubmitActionOptions(configuration.DependencyGraphAutosubmitActionOptions)},
		{"dependabot_alerts", configuration.GetDependabotAlerts()},
		{"dependabot_security_updates", configuration.GetDependabotSecurityUpdates()},
		{"code_scanning_default_setup", configuration.GetCodeScanningDefaultSetup()},
		{"code_scanning_default_setup_options", flattenCodeScanningDefaultSetupOptions(configuration.CodeScanningDefaultSetupOptions)},
		{"code_scanning_delegated_alert_dismissal", configuration.GetCodeScanningDelegatedAlertDismissal()},
		{"code_scanning_options", flattenCodeScanningOptions(configuration.CodeScanningOptions)},
		{"code_security", configuration.GetCodeSecurity()},
		{"secret_scanning", configuration.GetSecretScanning()},
		{"secret_scanning_push_protection", configuration.GetSecretScanningPushProtection()},
		{"secret_scanning_validity_checks", configuration.GetSecretScanningValidityChecks()},
		{"secret_scanning_non_provider_patterns", configuration.GetSecretScanningNonProviderPatterns()},
		{"secret_scanning_generic_secrets", configuration.GetSecretScanningGenericSecrets()},
		{"secret_scanning_delegated_alert_dismissal", configuration.GetSecretScanningDelegatedAlertDismissal()},
		{"secret_protection", configuration.GetSecretProtection()},
		{"private_vulnerability_reporting", configuration.GetPrivateVulnerabilityReporting()},
		{"enforcement", configuration.GetEnforcement()},
	}
	// Note: this pattern does not support bool attributes, as d.GetOk for a bool field always
	// returns true. All attributes reconciled here are strings or nested blocks, so that is fine.
	for _, attr := range managed {
		if _, ok := d.GetOk(attr.key); !ok {
			continue
		}
		if err := d.Set(attr.key, attr.value); err != nil {
			return diag.FromErr(err)
		}
	}

	tflog.Trace(ctx, "Successfully read enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id})

	return nil
}

func resourceGithubEnterpriseSecurityConfigurationUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	enterprise, _ := d.Get("enterprise_slug").(string)
	configID, _ := d.Get("configuration_id").(int)
	id := int64(configID)

	tflog.Debug(ctx, "Updating enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id})

	config := resourceGithubEnterpriseSecurityConfigurationExpand(d)

	// Optional (non-computed) attributes are persisted from config by the SDK and the computed
	// values don't change on update, so there is nothing to write back to state here.
	if _, _, err := client.Enterprise.UpdateCodeSecurityConfiguration(ctx, enterprise, id, config); err != nil {
		tflog.Error(ctx, "Failed to update enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id, "error": err.Error()})
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Updated enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id})

	return nil
}

func resourceGithubEnterpriseSecurityConfigurationDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	enterprise, _ := d.Get("enterprise_slug").(string)
	configID, _ := d.Get("configuration_id").(int)
	id := int64(configID)

	tflog.Debug(ctx, "Deleting enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id})

	_, err := client.Enterprise.DeleteCodeSecurityConfiguration(ctx, enterprise, id)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Enterprise code security configuration already deleted", map[string]any{"enterprise": enterprise, "id": id})
			return nil
		}
		tflog.Error(ctx, "Failed to delete enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id, "error": err.Error()})
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Deleted enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id})

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

	if err = d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}

	if err = d.Set("configuration_id", int(configID)); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

// resourceGithubEnterpriseSecurityConfigurationExpand builds a CodeSecurityConfiguration from Terraform resource data.
func resourceGithubEnterpriseSecurityConfigurationExpand(d *schema.ResourceData) github.CodeSecurityConfiguration {
	config := github.CodeSecurityConfiguration{
		Name: d.Get("name").(string),
	}
	if val, ok := d.GetOk("description"); ok {
		config.Description = val.(string)
	}

	if val, ok := d.GetOk("advanced_security"); ok {
		config.AdvancedSecurity = new(val.(string))
	}
	if val, ok := d.GetOk("dependency_graph"); ok {
		config.DependencyGraph = new(val.(string))
	}
	if val, ok := d.GetOk("dependency_graph_autosubmit_action"); ok {
		config.DependencyGraphAutosubmitAction = new(val.(string))
	}
	if val, ok := d.GetOk("dependabot_alerts"); ok {
		config.DependabotAlerts = new(val.(string))
	}
	if val, ok := d.GetOk("dependabot_security_updates"); ok {
		config.DependabotSecurityUpdates = new(val.(string))
	}
	if val, ok := d.GetOk("code_scanning_default_setup"); ok {
		config.CodeScanningDefaultSetup = new(val.(string))
	}
	if val, ok := d.GetOk("code_scanning_delegated_alert_dismissal"); ok {
		config.CodeScanningDelegatedAlertDismissal = new(val.(string))
	}
	if val, ok := d.GetOk("code_security"); ok {
		config.CodeSecurity = new(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning"); ok {
		config.SecretScanning = new(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_push_protection"); ok {
		config.SecretScanningPushProtection = new(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_validity_checks"); ok {
		config.SecretScanningValidityChecks = new(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_non_provider_patterns"); ok {
		config.SecretScanningNonProviderPatterns = new(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_generic_secrets"); ok {
		config.SecretScanningGenericSecrets = new(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_delegated_alert_dismissal"); ok {
		config.SecretScanningDelegatedAlertDismissal = new(val.(string))
	}
	if val, ok := d.GetOk("secret_protection"); ok {
		config.SecretProtection = new(val.(string))
	}
	if val, ok := d.GetOk("private_vulnerability_reporting"); ok {
		config.PrivateVulnerabilityReporting = new(val.(string))
	}
	if val, ok := d.GetOk("enforcement"); ok {
		config.Enforcement = new(val.(string))
	}

	if val, ok := d.GetOk("dependency_graph_autosubmit_action_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			autosubmitOpts := optionsList[0].(map[string]any)
			config.DependencyGraphAutosubmitActionOptions = &github.DependencyGraphAutosubmitActionOptions{
				LabeledRunners: new(autosubmitOpts["labeled_runners"].(bool)),
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
				config.CodeScanningDefaultSetupOptions.RunnerLabel = new(runnerLabel)
			}
		}
	}

	if val, ok := d.GetOk("code_scanning_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			scanOpts := optionsList[0].(map[string]any)
			config.CodeScanningOptions = &github.CodeScanningOptions{
				AllowAdvanced: new(scanOpts["allow_advanced"].(bool)),
			}
		}
	}

	return config
}
