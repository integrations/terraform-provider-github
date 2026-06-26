package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v88/github"
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
				Description: "A description of the code security configuration.",
			},
			"advanced_security": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The advanced security configuration for the code security configuration. Can be one of 'enabled', 'disabled'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled",
				}, false)),
			},
			"dependency_graph": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The dependency graph configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"dependency_graph_autosubmit_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The dependency graph autosubmit action configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"dependency_graph_autosubmit_action_options": {
				Type:        schema.TypeList,
				Optional:    true,
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
				Description: "The dependabot alerts configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"dependabot_security_updates": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The dependabot security updates configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"code_scanning_default_setup": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The code scanning default setup configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"code_scanning_default_setup_options": {
				Type:        schema.TypeList,
				Optional:    true,
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
				Description: "The code scanning delegated alert dismissal configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"code_scanning_options": {
				Type:        schema.TypeList,
				Optional:    true,
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
				Description: "The code security setting. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The secret scanning configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_push_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The secret scanning push protection configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_validity_checks": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The secret scanning validity checks configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_non_provider_patterns": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The secret scanning non provider patterns configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_generic_secrets": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The secret scanning generic secrets configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_delegated_alert_dismissal": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The secret scanning delegated alert dismissal configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The secret protection configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"private_vulnerability_reporting": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The private vulnerability reporting configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"enforcement": {
				Type:        schema.TypeString,
				Optional:    true,
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

	if diags := resourceGithubEnterpriseSecurityConfigurationSetState(d, configuration); diags.HasError() {
		return diags
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

	if diags := resourceGithubEnterpriseSecurityConfigurationSetState(d, configuration); diags.HasError() {
		return diags
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

	configuration, _, err := client.Enterprise.UpdateCodeSecurityConfiguration(ctx, enterprise, id, config)
	if err != nil {
		tflog.Error(ctx, "Failed to update enterprise code security configuration", map[string]any{"enterprise": enterprise, "id": id, "error": err.Error()})
		return diag.FromErr(err)
	}

	if diags := resourceGithubEnterpriseSecurityConfigurationSetState(d, configuration); diags.HasError() {
		return diags
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

// resourceGithubEnterpriseSecurityConfigurationSetState writes the configuration returned by the API to Terraform state.
func resourceGithubEnterpriseSecurityConfigurationSetState(d *schema.ResourceData, configuration *github.CodeSecurityConfiguration) diag.Diagnostics {
	// Only persist attributes the user configured. The optional config attributes are not
	// Computed, and GitHub server-assigns a value to each one even when it is omitted, so
	// writing those server values to state would produce a perpetual diff. Unmanaged values
	// can be read via the data source instead. During import there is no config, so every
	// attribute is populated.
	// Only persist attributes the user manages (a non-empty value in config on create, or in
	// prior state on read). d.Get reflects config during create/update and prior state during
	// read, so unmanaged attributes stay out of state and GitHub's server-assigned defaults
	// don't surface as perpetual diffs.
	managed := func(key string) bool {
		switch v := d.Get(key).(type) {
		case string:
			return v != ""
		case []any:
			return len(v) > 0
		default:
			return v != nil
		}
	}

	if err := d.Set("configuration_id", int(configuration.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", configuration.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target_type", configuration.GetTargetType()); err != nil {
		return diag.FromErr(err)
	}

	optional := []struct {
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
	for _, attr := range optional {
		if !managed(attr.key) {
			continue
		}
		if err := d.Set(attr.key, attr.value); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
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
