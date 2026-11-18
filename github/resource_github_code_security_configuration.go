package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubCodeSecurityConfiguration() *schema.Resource {
	featureValueDiag := validation.ToDiagFunc(validation.StringInSlice([]string{"enabled", "disabled", "not_set"}, false))

	return &schema.Resource{
		Description:   "Manages a GitHub Code Security Configuration at the organization or enterprise level.",
		CreateContext: resourceGithubCodeSecurityConfigurationCreate,
		ReadContext:   resourceGithubCodeSecurityConfigurationRead,
		UpdateContext: resourceGithubCodeSecurityConfigurationUpdate,
		DeleteContext: resourceGithubCodeSecurityConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubCodeSecurityConfigurationImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the code security configuration. Must be unique within the organization or enterprise.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "A description of the code security configuration.",
			},
			"enterprise_slug": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise to create the configuration in. If omitted, the configuration is created at the organization level using the provider's configured owner.",
			},
			"advanced_security": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of GitHub Advanced Security. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"code_security": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The enablement status of GitHub Code Security. Can be 'enabled', 'disabled' or 'not_set'. Only sent when set.",
				ValidateDiagFunc: featureValueDiag,
			},
			"secret_protection": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The enablement status of GitHub Secret Protection. Can be 'enabled', 'disabled' or 'not_set'. Only sent when set.",
				ValidateDiagFunc: featureValueDiag,
			},
			"dependency_graph": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "enabled",
				Description:      "The enablement status of Dependency Graph. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"dependency_graph_autosubmit_action": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The enablement status of automatic dependency submission. Can be 'enabled', 'disabled' or 'not_set'. Only sent when set.",
				ValidateDiagFunc: featureValueDiag,
			},
			"dependency_graph_autosubmit_action_options": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Options for automatic dependency submission.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labeled_runners": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to use runners labeled with 'dependency-submission' or standard GitHub runners.",
						},
					},
				},
			},
			"dependabot_alerts": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of Dependabot alerts. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"dependabot_delegated_alert_dismissal": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The enablement status of Dependabot delegated alert dismissal. Can be 'enabled', 'disabled' or 'not_set'. Only sent when set.",
				ValidateDiagFunc: featureValueDiag,
			},
			"dependabot_security_updates": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of Dependabot security updates. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"code_scanning_default_setup": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of code scanning default setup. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"code_scanning_default_setup_options": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Options for code scanning default setup.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"runner_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "not_set",
							Description:      "The type of runner to use for code scanning default setup. Can be 'standard', 'labeled' or 'not_set'.",
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"standard", "labeled", "not_set"}, false)),
						},
						"runner_label": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The label of the runner to use for code scanning default setup when 'runner_type' is 'labeled'.",
						},
					},
				},
			},
			"code_scanning_delegated_alert_dismissal": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The enablement status of code scanning delegated alert dismissal. Can be 'enabled', 'disabled' or 'not_set'. Only sent when set.",
				ValidateDiagFunc: featureValueDiag,
			},
			"code_scanning_options": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Options for code scanning.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_advanced": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to allow repositories to use advanced (self-managed) code scanning setup.",
						},
					},
				},
			},
			"secret_scanning": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of secret scanning. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"secret_scanning_push_protection": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of secret scanning push protection. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"secret_scanning_delegated_bypass": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The enablement status of secret scanning delegated bypass. Can be 'enabled', 'disabled' or 'not_set'. Only sent when set.",
				ValidateDiagFunc: featureValueDiag,
			},
			"secret_scanning_delegated_bypass_options": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Options for secret scanning delegated bypass.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reviewers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Reviewers permitted to bypass secret scanning push protection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"reviewer_id": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The ID of the team or role selected as a bypass reviewer.",
									},
									"reviewer_type": {
										Type:             schema.TypeString,
										Required:         true,
										Description:      "The type of the bypass reviewer. Can be 'TEAM' or 'ROLE'.",
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"TEAM", "ROLE"}, false)),
									},
								},
							},
						},
					},
				},
			},
			"secret_scanning_validity_checks": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of secret scanning validity checks. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"secret_scanning_non_provider_patterns": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of secret scanning non-provider patterns. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"secret_scanning_generic_secrets": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The enablement status of Copilot secret scanning (generic secrets). Can be 'enabled', 'disabled' or 'not_set'. Only sent when set.",
				ValidateDiagFunc: featureValueDiag,
			},
			"secret_scanning_delegated_alert_dismissal": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The enablement status of secret scanning delegated alert dismissal. Can be 'enabled', 'disabled' or 'not_set'. Only sent when set.",
				ValidateDiagFunc: featureValueDiag,
			},
			"secret_scanning_extended_metadata": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The enablement status of secret scanning extended metadata. Can be 'enabled', 'disabled' or 'not_set'. Only sent when set.",
				ValidateDiagFunc: featureValueDiag,
			},
			"private_vulnerability_reporting": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of private vulnerability reporting. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"enforcement": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "enforced",
				Description:      "The enforcement status of the configuration. Can be 'enforced' or 'unenforced'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"enforced", "unenforced"}, false)),
			},
			"default_for_new_repos": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Which types of new repositories this configuration should be applied to by default. Can be 'all', 'none', 'private_and_internal' or 'public'. If omitted, the configuration is not set as a default.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "none", "private_and_internal", "public"}, false)),
			},
			"configuration_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The numeric ID of the code security configuration.",
			},
			"target_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target type of the configuration ('organization' or 'enterprise').",
			},
			"html_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the configuration in the GitHub UI.",
			},
		},
	}
}

// stringAttr returns the configured string at key, or "" when unset.
func stringAttr(d *schema.ResourceData, key string) string {
	s, _ := d.Get(key).(string)
	return s
}

// optionalString returns a pointer to the configured string at key, or nil when
// the attribute is unset. Attributes without a schema Default are only sent to
// the API when the user has set them, so GitHub keeps its own default.
func optionalString(d *schema.ResourceData, key string) *string {
	v, ok := d.GetOk(key)
	if !ok {
		return nil
	}
	s, _ := v.(string)
	return new(s)
}

// firstBlock returns the single nested block at key, or nil when the block is
// not configured.
func firstBlock(d *schema.ResourceData, key string) map[string]any {
	v, ok := d.GetOk(key)
	if !ok {
		return nil
	}
	list, _ := v.([]any)
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	block, _ := list[0].(map[string]any)
	return block
}

func buildCodeSecurityConfiguration(d *schema.ResourceData) github.CodeSecurityConfiguration {
	config := github.CodeSecurityConfiguration{
		Name:        stringAttr(d, "name"),
		Description: stringAttr(d, "description"),

		// Attributes with a schema Default are always sent.
		AdvancedSecurity:                  new(stringAttr(d, "advanced_security")),
		DependencyGraph:                   new(stringAttr(d, "dependency_graph")),
		DependabotAlerts:                  new(stringAttr(d, "dependabot_alerts")),
		DependabotSecurityUpdates:         new(stringAttr(d, "dependabot_security_updates")),
		CodeScanningDefaultSetup:          new(stringAttr(d, "code_scanning_default_setup")),
		SecretScanning:                    new(stringAttr(d, "secret_scanning")),
		SecretScanningPushProtection:      new(stringAttr(d, "secret_scanning_push_protection")),
		SecretScanningValidityChecks:      new(stringAttr(d, "secret_scanning_validity_checks")),
		SecretScanningNonProviderPatterns: new(stringAttr(d, "secret_scanning_non_provider_patterns")),
		PrivateVulnerabilityReporting:     new(stringAttr(d, "private_vulnerability_reporting")),
		Enforcement:                       new(stringAttr(d, "enforcement")),

		// Attributes without a schema Default are only sent when configured.
		CodeSecurity:                          optionalString(d, "code_security"),
		SecretProtection:                      optionalString(d, "secret_protection"),
		DependencyGraphAutosubmitAction:       optionalString(d, "dependency_graph_autosubmit_action"),
		DependabotDelegatedAlertDismissal:     optionalString(d, "dependabot_delegated_alert_dismissal"),
		CodeScanningDelegatedAlertDismissal:   optionalString(d, "code_scanning_delegated_alert_dismissal"),
		SecretScanningDelegatedBypass:         optionalString(d, "secret_scanning_delegated_bypass"),
		SecretScanningGenericSecrets:          optionalString(d, "secret_scanning_generic_secrets"),
		SecretScanningDelegatedAlertDismissal: optionalString(d, "secret_scanning_delegated_alert_dismissal"),
		SecretScanningExtendedMetadata:        optionalString(d, "secret_scanning_extended_metadata"),
	}

	if block := firstBlock(d, "dependency_graph_autosubmit_action_options"); block != nil {
		labeledRunners, _ := block["labeled_runners"].(bool)
		config.DependencyGraphAutosubmitActionOptions = &github.DependencyGraphAutosubmitActionOptions{
			LabeledRunners: new(labeledRunners),
		}
	}

	if block := firstBlock(d, "code_scanning_default_setup_options"); block != nil {
		runnerType, _ := block["runner_type"].(string)
		if runnerType == "" {
			runnerType = "not_set"
		}
		options := &github.CodeScanningDefaultSetupOptions{RunnerType: runnerType}
		if runnerLabel, _ := block["runner_label"].(string); runnerLabel != "" {
			options.RunnerLabel = new(runnerLabel)
		}
		config.CodeScanningDefaultSetupOptions = options
	}

	if block := firstBlock(d, "code_scanning_options"); block != nil {
		allowAdvanced, _ := block["allow_advanced"].(bool)
		config.CodeScanningOptions = &github.CodeScanningOptions{
			AllowAdvanced: new(allowAdvanced),
		}
	}

	if block := firstBlock(d, "secret_scanning_delegated_bypass_options"); block != nil {
		reviewersList, _ := block["reviewers"].([]any)
		reviewers := make([]*github.BypassReviewer, 0, len(reviewersList))
		for _, r := range reviewersList {
			reviewer, _ := r.(map[string]any)
			id, _ := reviewer["reviewer_id"].(int)
			reviewerType, _ := reviewer["reviewer_type"].(string)
			reviewers = append(reviewers, &github.BypassReviewer{
				ReviewerID:   int64(id),
				ReviewerType: reviewerType,
			})
		}
		config.SecretScanningDelegatedBypassOptions = &github.SecretScanningDelegatedBypassOptions{
			Reviewers: reviewers,
		}
	}

	return config
}

func flattenDependencyGraphAutosubmitActionOptions(options *github.DependencyGraphAutosubmitActionOptions) []any {
	if options == nil {
		return []any{}
	}
	return []any{map[string]any{
		"labeled_runners": options.GetLabeledRunners(),
	}}
}

func flattenCodeScanningDefaultSetupOptions(options *github.CodeScanningDefaultSetupOptions) []any {
	if options == nil {
		return []any{}
	}
	block := map[string]any{
		"runner_type": options.RunnerType,
	}
	if options.RunnerLabel != nil {
		block["runner_label"] = options.GetRunnerLabel()
	}
	return []any{block}
}

func flattenCodeScanningOptions(options *github.CodeScanningOptions) []any {
	if options == nil {
		return []any{}
	}
	return []any{map[string]any{
		"allow_advanced": options.GetAllowAdvanced(),
	}}
}

func flattenSecretScanningDelegatedBypassOptions(options *github.SecretScanningDelegatedBypassOptions) []any {
	if options == nil {
		return []any{}
	}
	reviewers := make([]any, 0, len(options.Reviewers))
	for _, reviewer := range options.Reviewers {
		if reviewer == nil {
			continue
		}
		reviewers = append(reviewers, map[string]any{
			"reviewer_id":   int(reviewer.ReviewerID),
			"reviewer_type": reviewer.ReviewerType,
		})
	}
	return []any{map[string]any{
		"reviewers": reviewers,
	}}
}

// setCodeSecurityConfigurationState writes the API representation into state.
//
// Attributes with a schema Default are always written. Attributes without a
// Default are only written when they are present in the configuration, so an
// unmanaged attribute never produces a diff against whatever GitHub holds for
// it. Nested option blocks follow the same rule. This pattern relies on
// d.GetOk and therefore does not support top-level bool attributes; all
// optional-only attributes here are strings or blocks.
func setCodeSecurityConfigurationState(d *schema.ResourceData, config *github.CodeSecurityConfiguration) error {
	always := []struct {
		key   string
		value any
	}{
		{"name", config.Name},
		{"description", config.Description},
		{"advanced_security", config.GetAdvancedSecurity()},
		{"dependency_graph", config.GetDependencyGraph()},
		{"dependabot_alerts", config.GetDependabotAlerts()},
		{"dependabot_security_updates", config.GetDependabotSecurityUpdates()},
		{"code_scanning_default_setup", config.GetCodeScanningDefaultSetup()},
		{"secret_scanning", config.GetSecretScanning()},
		{"secret_scanning_push_protection", config.GetSecretScanningPushProtection()},
		{"secret_scanning_validity_checks", config.GetSecretScanningValidityChecks()},
		{"secret_scanning_non_provider_patterns", config.GetSecretScanningNonProviderPatterns()},
		{"private_vulnerability_reporting", config.GetPrivateVulnerabilityReporting()},
		{"enforcement", config.GetEnforcement()},
		{"configuration_id", config.GetID()},
		{"target_type", config.GetTargetType()},
		{"html_url", config.GetHTMLURL()},
	}
	for _, attr := range always {
		if err := d.Set(attr.key, attr.value); err != nil {
			return err
		}
	}

	whenConfigured := []struct {
		key   string
		value any
	}{
		{"code_security", config.GetCodeSecurity()},
		{"secret_protection", config.GetSecretProtection()},
		{"dependency_graph_autosubmit_action", config.GetDependencyGraphAutosubmitAction()},
		{"dependency_graph_autosubmit_action_options", flattenDependencyGraphAutosubmitActionOptions(config.DependencyGraphAutosubmitActionOptions)},
		{"dependabot_delegated_alert_dismissal", config.GetDependabotDelegatedAlertDismissal()},
		{"code_scanning_default_setup_options", flattenCodeScanningDefaultSetupOptions(config.CodeScanningDefaultSetupOptions)},
		{"code_scanning_delegated_alert_dismissal", config.GetCodeScanningDelegatedAlertDismissal()},
		{"code_scanning_options", flattenCodeScanningOptions(config.CodeScanningOptions)},
		{"secret_scanning_delegated_bypass", config.GetSecretScanningDelegatedBypass()},
		{"secret_scanning_delegated_bypass_options", flattenSecretScanningDelegatedBypassOptions(config.SecretScanningDelegatedBypassOptions)},
		{"secret_scanning_generic_secrets", config.GetSecretScanningGenericSecrets()},
		{"secret_scanning_delegated_alert_dismissal", config.GetSecretScanningDelegatedAlertDismissal()},
		{"secret_scanning_extended_metadata", config.GetSecretScanningExtendedMetadata()},
	}
	for _, attr := range whenConfigured {
		if _, ok := d.GetOk(attr.key); !ok {
			continue
		}
		if err := d.Set(attr.key, attr.value); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubCodeSecurityConfigurationCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	enterpriseSlug := stringAttr(d, "enterprise_slug")
	body := buildCodeSecurityConfiguration(d)

	// The API rejects delegated bypass reviewers on the initial create request
	// but accepts them on update, so they are applied in a follow-up call.
	bypassOptions := body.SecretScanningDelegatedBypassOptions
	body.SecretScanningDelegatedBypassOptions = nil

	var config *github.CodeSecurityConfiguration
	var err error
	if enterpriseSlug != "" {
		tflog.Debug(ctx, "Creating code security configuration for enterprise", map[string]any{"name": body.Name, "enterprise_slug": enterpriseSlug})
		config, _, err = client.Enterprise.CreateCodeSecurityConfiguration(ctx, enterpriseSlug, body)
	} else {
		if err := checkOrganization(m); err != nil {
			return diag.FromErr(err)
		}
		tflog.Debug(ctx, "Creating code security configuration for organization", map[string]any{"name": body.Name, "owner": owner})
		config, _, err = client.Organizations.CreateCodeSecurityConfiguration(ctx, owner, body)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(config.GetID(), 10))

	if bypassOptions != nil {
		body.SecretScanningDelegatedBypassOptions = bypassOptions
		tflog.Debug(ctx, "Applying delegated bypass reviewers to code security configuration", map[string]any{"configuration_id": config.GetID()})
		if enterpriseSlug != "" {
			config, _, err = client.Enterprise.UpdateCodeSecurityConfiguration(ctx, enterpriseSlug, config.GetID(), body)
		} else {
			config, _, err = client.Organizations.UpdateCodeSecurityConfiguration(ctx, owner, config.GetID(), body)
		}
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if err := setCodeSecurityConfigurationState(d, config); err != nil {
		return diag.FromErr(err)
	}

	if defaultForNewRepos := stringAttr(d, "default_for_new_repos"); defaultForNewRepos != "" {
		if err := setCodeSecurityConfigurationDefault(ctx, client, owner, enterpriseSlug, config.GetID(), defaultForNewRepos); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubCodeSecurityConfigurationRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	enterpriseSlug := stringAttr(d, "enterprise_slug")

	var config *github.CodeSecurityConfiguration
	if enterpriseSlug != "" {
		config, _, err = client.Enterprise.GetCodeSecurityConfiguration(ctx, enterpriseSlug, configID)
	} else {
		config, _, err = client.Organizations.GetCodeSecurityConfiguration(ctx, owner, configID)
	}
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing code security configuration from state because it no longer exists in GitHub.", map[string]any{"resource_id": d.Id(), "owner": owner, "enterprise_slug": enterpriseSlug})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := setCodeSecurityConfigurationState(d, config); err != nil {
		return diag.FromErr(err)
	}

	// default_for_new_repos is only surfaced via the /defaults listing.
	defaultForNewRepos, err := readCodeSecurityConfigurationDefault(ctx, client, owner, enterpriseSlug, configID)
	if err != nil {
		return diag.FromErr(err)
	}
	// A configuration that is not a default (or is explicitly "none") does not
	// appear in the defaults listing; preserve an explicit "none" in state.
	if defaultForNewRepos == "" && stringAttr(d, "default_for_new_repos") == "none" {
		defaultForNewRepos = "none"
	}
	if err := d.Set("default_for_new_repos", defaultForNewRepos); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubCodeSecurityConfigurationUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	enterpriseSlug := stringAttr(d, "enterprise_slug")
	body := buildCodeSecurityConfiguration(d)

	var config *github.CodeSecurityConfiguration
	if enterpriseSlug != "" {
		tflog.Debug(ctx, "Updating code security configuration for enterprise", map[string]any{"configuration_id": configID, "enterprise_slug": enterpriseSlug})
		config, _, err = client.Enterprise.UpdateCodeSecurityConfiguration(ctx, enterpriseSlug, configID, body)
	} else {
		tflog.Debug(ctx, "Updating code security configuration for organization", map[string]any{"configuration_id": configID, "owner": owner})
		config, _, err = client.Organizations.UpdateCodeSecurityConfiguration(ctx, owner, configID, body)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	if config != nil {
		if err := setCodeSecurityConfigurationState(d, config); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("default_for_new_repos") {
		newReposParam := stringAttr(d, "default_for_new_repos")
		if newReposParam == "" {
			// Removing the attribute reverts the default to "none".
			newReposParam = "none"
		}
		if err := setCodeSecurityConfigurationDefault(ctx, client, owner, enterpriseSlug, configID, newReposParam); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubCodeSecurityConfigurationDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	enterpriseSlug := stringAttr(d, "enterprise_slug")

	if enterpriseSlug != "" {
		tflog.Debug(ctx, "Deleting code security configuration for enterprise", map[string]any{"configuration_id": configID, "enterprise_slug": enterpriseSlug})
		_, err = client.Enterprise.DeleteCodeSecurityConfiguration(ctx, enterpriseSlug, configID)
	} else {
		tflog.Debug(ctx, "Deleting code security configuration for organization", map[string]any{"configuration_id": configID, "owner": owner})
		_, err = client.Organizations.DeleteCodeSecurityConfiguration(ctx, owner, configID)
	}
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Code security configuration no longer exists in GitHub; treating delete as successful.", map[string]any{"resource_id": d.Id(), "owner": owner, "enterprise_slug": enterpriseSlug})
			return nil
		}
		return diag.FromErr(err)
	}
	return nil
}

// resourceGithubCodeSecurityConfigurationImport supports two import ID formats:
//   - "<configuration_id>" for organization-level configurations
//   - "<enterprise_slug>:<configuration_id>" for enterprise-level configurations
func resourceGithubCodeSecurityConfigurationImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")
	switch len(parts) {
	case 1:
		// organization-level: ID is the configuration ID as-is
	case 2:
		if err := d.Set("enterprise_slug", parts[0]); err != nil {
			return nil, err
		}
		d.SetId(parts[1])
	default:
		return nil, fmt.Errorf("invalid import ID %q: expected '<configuration_id>' or '<enterprise_slug>:<configuration_id>'", d.Id())
	}

	if _, err := strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return nil, fmt.Errorf("invalid configuration ID %q: %w", d.Id(), err)
	}

	return []*schema.ResourceData{d}, nil
}

func setCodeSecurityConfigurationDefault(ctx context.Context, client *github.Client, org, enterpriseSlug string, configID int64, newReposParam string) error {
	var err error
	if enterpriseSlug != "" {
		tflog.Debug(ctx, "Setting code security configuration as default for enterprise", map[string]any{"configuration_id": configID, "default_for_new_repos": newReposParam, "enterprise_slug": enterpriseSlug})
		_, _, err = client.Enterprise.SetDefaultCodeSecurityConfiguration(ctx, enterpriseSlug, configID, newReposParam)
	} else {
		tflog.Debug(ctx, "Setting code security configuration as default for organization", map[string]any{"configuration_id": configID, "default_for_new_repos": newReposParam, "owner": org})
		_, _, err = client.Organizations.SetDefaultCodeSecurityConfiguration(ctx, org, configID, newReposParam)
	}
	return err
}

func readCodeSecurityConfigurationDefault(ctx context.Context, client *github.Client, org, enterpriseSlug string, configID int64) (string, error) {
	var defaults []*github.CodeSecurityConfigurationWithDefaultForNewRepos
	var err error
	if enterpriseSlug != "" {
		defaults, _, err = client.Enterprise.ListDefaultCodeSecurityConfigurations(ctx, enterpriseSlug)
	} else {
		defaults, _, err = client.Organizations.ListDefaultCodeSecurityConfigurations(ctx, org)
	}
	if err != nil {
		return "", err
	}
	for _, def := range defaults {
		if def.GetConfiguration().GetID() == configID {
			value := def.GetDefaultForNewRepos()
			if value == "none" {
				return "", nil
			}
			return value, nil
		}
	}
	return "", nil
}
