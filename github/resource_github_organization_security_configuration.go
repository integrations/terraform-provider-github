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

func resourceGithubOrganizationSecurityConfiguration() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages a code security configuration for a GitHub Organization.",
		CreateContext: resourceGithubOrganizationSecurityConfigurationCreate,
		ReadContext:   resourceGithubOrganizationSecurityConfigurationRead,
		UpdateContext: resourceGithubOrganizationSecurityConfigurationUpdate,
		DeleteContext: resourceGithubOrganizationSecurityConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubOrganizationSecurityConfigurationImport,
		},

		Schema: map[string]*schema.Schema{
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
							Required:    true,
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
							Required:    true,
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
			"secret_scanning_delegated_bypass": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The secret scanning delegated bypass configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false)),
			},
			"secret_scanning_delegated_bypass_options": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The secret scanning delegated bypass options for the code security configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reviewers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The bypass reviewers for the secret scanning delegated bypass.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"reviewer_id": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The ID of the bypass reviewer.",
									},
									"reviewer_type": {
										Type:             schema.TypeString,
										Required:         true,
										Description:      "The type of the bypass reviewer. Can be one of 'TEAM', 'ROLE'.",
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"TEAM", "ROLE"}, false)),
									},
								},
							},
						},
					},
				},
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

func resourceGithubOrganizationSecurityConfigurationCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	client := meta.v3client
	org := meta.name
	name, _ := d.Get("name").(string)

	tflog.Debug(ctx, "Creating organization code security configuration", map[string]any{"organization": org, "name": name})

	config := resourceGithubOrganizationSecurityConfigurationExpand(d)

	// The GitHub API returns HTTP 500 when secret_scanning_delegated_bypass_options reviewers
	// are sent in the initial create request, even though the configuration is created. Setting
	// the reviewers via a follow-up update succeeds, so we defer them: create without the bypass
	// options, then apply them with an update if any were specified.
	bypassOptions := config.SecretScanningDelegatedBypassOptions
	config.SecretScanningDelegatedBypassOptions = nil

	configuration, _, err := client.Organizations.CreateCodeSecurityConfiguration(ctx, org, config)
	if err != nil {
		tflog.Error(ctx, "Failed to create organization code security configuration", map[string]any{"organization": org, "name": name, "error": err.Error()})
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(configuration.GetID(), 10))

	if bypassOptions != nil {
		config.SecretScanningDelegatedBypassOptions = bypassOptions
		configuration, _, err = client.Organizations.UpdateCodeSecurityConfiguration(ctx, org, configuration.GetID(), config)
		if err != nil {
			tflog.Error(ctx, "Failed to set secret scanning delegated bypass options on organization code security configuration", map[string]any{"organization": org, "name": name, "id": configuration.GetID(), "error": err.Error()})
			return diag.FromErr(err)
		}
	}

	// Optional (non-computed) attributes are persisted from config by the SDK; only the
	// computed values need to be written here.
	if err := d.Set("configuration_id", int(configuration.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target_type", configuration.GetTargetType()); err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Created organization code security configuration", map[string]any{"organization": org, "name": name, "id": configuration.GetID()})

	return nil
}

func resourceGithubOrganizationSecurityConfigurationRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	client := meta.v3client
	org := meta.name
	configID, _ := d.Get("configuration_id").(int)
	id := int64(configID)

	tflog.Trace(ctx, "Reading organization code security configuration", map[string]any{"organization": org, "id": id})

	configuration, _, err := client.Organizations.GetCodeSecurityConfiguration(ctx, org, id)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing organization code security configuration from state because it no longer exists in GitHub", map[string]any{"organization": org, "id": id})
			d.SetId("")
			return nil
		}
		tflog.Error(ctx, "Failed to read organization code security configuration", map[string]any{"organization": org, "id": id, "error": err.Error()})
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
		{"secret_scanning_delegated_bypass", configuration.GetSecretScanningDelegatedBypass()},
		{"secret_scanning_delegated_bypass_options", flattenSecretScanningDelegatedBypassOptions(configuration.SecretScanningDelegatedBypassOptions)},
		{"secret_scanning_validity_checks", configuration.GetSecretScanningValidityChecks()},
		{"secret_scanning_non_provider_patterns", configuration.GetSecretScanningNonProviderPatterns()},
		{"secret_scanning_generic_secrets", configuration.GetSecretScanningGenericSecrets()},
		{"secret_scanning_delegated_alert_dismissal", configuration.GetSecretScanningDelegatedAlertDismissal()},
		{"secret_protection", configuration.GetSecretProtection()},
		{"private_vulnerability_reporting", configuration.GetPrivateVulnerabilityReporting()},
		{"enforcement", configuration.GetEnforcement()},
	}
	for _, attr := range managed {
		if _, ok := d.GetOk(attr.key); !ok {
			continue
		}
		if err := d.Set(attr.key, attr.value); err != nil {
			return diag.FromErr(err)
		}
	}

	tflog.Trace(ctx, "Successfully read organization code security configuration", map[string]any{"organization": org, "id": id})

	return nil
}

func resourceGithubOrganizationSecurityConfigurationUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	client := meta.v3client
	org := meta.name
	configID, _ := d.Get("configuration_id").(int)
	id := int64(configID)

	tflog.Debug(ctx, "Updating organization code security configuration", map[string]any{"organization": org, "id": id})

	config := resourceGithubOrganizationSecurityConfigurationExpand(d)

	// Optional (non-computed) attributes are persisted from config by the SDK and the computed
	// values don't change on update, so there is nothing to write back to state here.
	if _, _, err := client.Organizations.UpdateCodeSecurityConfiguration(ctx, org, id, config); err != nil {
		tflog.Error(ctx, "Failed to update organization code security configuration", map[string]any{"organization": org, "id": id, "error": err.Error()})
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Updated organization code security configuration", map[string]any{"organization": org, "id": id})

	return nil
}

func resourceGithubOrganizationSecurityConfigurationDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	client := meta.v3client
	org := meta.name
	configID, _ := d.Get("configuration_id").(int)
	id := int64(configID)

	tflog.Debug(ctx, "Deleting organization code security configuration", map[string]any{"organization": org, "id": id})

	_, err := client.Organizations.DeleteCodeSecurityConfiguration(ctx, org, id)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Organization code security configuration already deleted", map[string]any{"organization": org, "id": id})
			return nil
		}
		tflog.Error(ctx, "Failed to delete organization code security configuration", map[string]any{"organization": org, "id": id, "error": err.Error()})
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Deleted organization code security configuration", map[string]any{"organization": org, "id": id})

	return nil
}

func resourceGithubOrganizationSecurityConfigurationImport(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration_id %q: %w", d.Id(), err)
	}

	if err = d.Set("configuration_id", int(configID)); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

// resourceGithubOrganizationSecurityConfigurationExpand builds a CodeSecurityConfiguration from Terraform resource data,
// including the organization-only secret scanning delegated bypass fields.
func resourceGithubOrganizationSecurityConfigurationExpand(d *schema.ResourceData) github.CodeSecurityConfiguration {
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

	if val, ok := d.GetOk("secret_scanning_delegated_bypass"); ok {
		config.SecretScanningDelegatedBypass = new(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_delegated_bypass_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			bypassOpts := optionsList[0].(map[string]any)
			options := &github.SecretScanningDelegatedBypassOptions{}
			if reviewersVal, ok := bypassOpts["reviewers"]; ok {
				reviewersList := reviewersVal.([]any)
				reviewers := make([]*github.BypassReviewer, 0, len(reviewersList))
				for _, reviewerRaw := range reviewersList {
					reviewerMap := reviewerRaw.(map[string]any)
					reviewers = append(reviewers, &github.BypassReviewer{
						ReviewerID:   int64(reviewerMap["reviewer_id"].(int)),
						ReviewerType: reviewerMap["reviewer_type"].(string),
					})
				}
				options.Reviewers = reviewers
			}
			config.SecretScanningDelegatedBypassOptions = options
		}
	}

	return config
}
