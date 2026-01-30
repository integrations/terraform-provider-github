package github

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationSecurityConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationSecurityConfigurationCreate,
		Read:   resourceGithubOrganizationSecurityConfigurationRead,
		Update: resourceGithubOrganizationSecurityConfigurationUpdate,
		Delete: resourceGithubOrganizationSecurityConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
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
				Default:     "disabled",
				Description: "The advanced security configuration for the code security configuration. Can be one of 'enabled', 'disabled'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled",
				}, false),
			},
			"dependency_graph": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "enabled",
				Description: "The dependency graph configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"dependency_graph_autosubmit_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The dependency graph autosubmit action configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
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
				Default:     "disabled",
				Description: "The dependabot alerts configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"dependabot_security_updates": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The dependabot security updates configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"code_scanning_default_setup": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The code scanning default setup configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
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
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The type of runner to use for code scanning default setup. Can be one of 'standard', 'labeled'.",
							ValidateFunc: validation.StringInSlice([]string{"standard", "labeled"}, false),
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
				Default:     "disabled",
				Description: "The code scanning delegated alert dismissal configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
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
				Default:     "disabled",
				Description: "The code security configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"secret_scanning": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The secret scanning configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"secret_scanning_push_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The secret scanning push protection configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"secret_scanning_delegated_bypass": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The secret scanning delegated bypass configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
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
										Type:         schema.TypeString,
										Required:     true,
										Description:  "The type of the bypass reviewer. Can be one of 'Team', 'Role'.",
										ValidateFunc: validation.StringInSlice([]string{"Team", "Role"}, false),
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
				Default:     "disabled",
				Description: "The secret scanning validity checks configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"secret_scanning_non_provider_patterns": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The secret scanning non provider patterns configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"secret_scanning_generic_secrets": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The secret scanning generic secrets configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"secret_scanning_delegated_alert_dismissal": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The secret scanning delegated alert dismissal configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"secret_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The secret protection configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"private_vulnerability_reporting": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "The private vulnerability reporting configuration for the code security configuration. Can be one of 'enabled', 'disabled', 'not_set'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enabled", "disabled", "not_set",
				}, false),
			},
			"enforcement": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "enforced",
				Description: "The enforcement configuration for the code security configuration. Can be one of 'enforced', 'unenforced'.",
				ValidateFunc: validation.StringInSlice([]string{
					"enforced", "unenforced",
				}, false),
			},
			"target_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target type of the code security configuration.",
			},
		},
	}
}

func resourceGithubOrganizationSecurityConfigurationCreate(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	org := meta.(*Owner).name

	config := expandCodeSecurityConfiguration(d)

	configuration, _, err := client.Organizations.CreateCodeSecurityConfiguration(ctx, org, config)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(configuration.GetID(), 10))

	return resourceGithubOrganizationSecurityConfigurationRead(d, meta)
}

func resourceGithubOrganizationSecurityConfigurationRead(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	org := meta.(*Owner).name

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	configuration, _, err := client.Organizations.GetCodeSecurityConfiguration(ctx, org, id)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if err = d.Set("name", configuration.Name); err != nil {
		return err
	}
	if err = d.Set("description", configuration.Description); err != nil {
		return err
	}
	if err = d.Set("advanced_security", configuration.GetAdvancedSecurity()); err != nil {
		return err
	}
	if err = d.Set("dependency_graph", configuration.GetDependencyGraph()); err != nil {
		return err
	}
	if err = d.Set("dependency_graph_autosubmit_action", configuration.GetDependencyGraphAutosubmitAction()); err != nil {
		return err
	}
	if err := d.Set("dependency_graph_autosubmit_action_options", flattenDependencyGraphAutosubmitActionOptions(configuration.DependencyGraphAutosubmitActionOptions)); err != nil {
		return err
	}
	if err = d.Set("dependabot_alerts", configuration.GetDependabotAlerts()); err != nil {
		return err
	}
	if err = d.Set("dependabot_security_updates", configuration.GetDependabotSecurityUpdates()); err != nil {
		return err
	}
	if err = d.Set("code_scanning_default_setup", configuration.GetCodeScanningDefaultSetup()); err != nil {
		return err
	}
	if err := d.Set("code_scanning_default_setup_options", flattenCodeScanningDefaultSetupOptions(configuration.CodeScanningDefaultSetupOptions)); err != nil {
		return err
	}
	if err = d.Set("code_scanning_delegated_alert_dismissal", configuration.GetCodeScanningDelegatedAlertDismissal()); err != nil {
		return err
	}
	if err := d.Set("code_scanning_options", flattenCodeScanningOptions(configuration.CodeScanningOptions)); err != nil {
		return err
	}
	codeSec := configuration.GetCodeSecurity()
	if codeSec == "" {
		codeSec = "disabled"
	}
	if err = d.Set("code_security", codeSec); err != nil {
		return err
	}
	if err = d.Set("secret_scanning", configuration.GetSecretScanning()); err != nil {
		return err
	}
	if err = d.Set("secret_scanning_push_protection", configuration.GetSecretScanningPushProtection()); err != nil {
		return err
	}
	if err = d.Set("secret_scanning_delegated_bypass", configuration.GetSecretScanningDelegatedBypass()); err != nil {
		return err
	}
	if err := d.Set("secret_scanning_delegated_bypass_options", flattenSecretScanningDelegatedBypassOptions(configuration.SecretScanningDelegatedBypassOptions)); err != nil {
		return err
	}
	if err = d.Set("secret_scanning_validity_checks", configuration.GetSecretScanningValidityChecks()); err != nil {
		return err
	}
	if err = d.Set("secret_scanning_non_provider_patterns", configuration.GetSecretScanningNonProviderPatterns()); err != nil {
		return err
	}
	if err = d.Set("secret_scanning_generic_secrets", configuration.GetSecretScanningGenericSecrets()); err != nil {
		return err
	}
	if err = d.Set("secret_scanning_delegated_alert_dismissal", configuration.GetSecretScanningDelegatedAlertDismissal()); err != nil {
		return err
	}
	secretProt := configuration.GetSecretProtection()
	if secretProt == "" {
		secretProt = "disabled"
	}
	if err = d.Set("secret_protection", secretProt); err != nil {
		return err
	}
	if err = d.Set("private_vulnerability_reporting", configuration.GetPrivateVulnerabilityReporting()); err != nil {
		return err
	}
	if err = d.Set("enforcement", configuration.GetEnforcement()); err != nil {
		return err
	}
	if err = d.Set("target_type", configuration.GetTargetType()); err != nil {
		return err
	}

	return nil
}

func resourceGithubOrganizationSecurityConfigurationUpdate(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	org := meta.(*Owner).name

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	config := expandCodeSecurityConfiguration(d)

	_, _, err = client.Organizations.UpdateCodeSecurityConfiguration(ctx, org, id, config)
	if err != nil {
		return err
	}

	return resourceGithubOrganizationSecurityConfigurationRead(d, meta)
}

func resourceGithubOrganizationSecurityConfigurationDelete(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	org := meta.(*Owner).name

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting code security configuration %s", d.Id())

	_, err = client.Organizations.DeleteCodeSecurityConfiguration(ctx, org, id)
	return err
}

func expandCodeSecurityConfiguration(d *schema.ResourceData) github.CodeSecurityConfiguration {
	config := github.CodeSecurityConfiguration{
		Name:                                  d.Get("name").(string),
		Description:                           d.Get("description").(string),
		DependencyGraph:                       github.Ptr(d.Get("dependency_graph").(string)),
		DependencyGraphAutosubmitAction:       github.Ptr(d.Get("dependency_graph_autosubmit_action").(string)),
		DependabotAlerts:                      github.Ptr(d.Get("dependabot_alerts").(string)),
		DependabotSecurityUpdates:             github.Ptr(d.Get("dependabot_security_updates").(string)),
		CodeScanningDefaultSetup:              github.Ptr(d.Get("code_scanning_default_setup").(string)),
		CodeScanningDelegatedAlertDismissal:   github.Ptr(d.Get("code_scanning_delegated_alert_dismissal").(string)),
		SecretScanning:                        github.Ptr(d.Get("secret_scanning").(string)),
		SecretScanningPushProtection:          github.Ptr(d.Get("secret_scanning_push_protection").(string)),
		SecretScanningDelegatedBypass:         github.Ptr(d.Get("secret_scanning_delegated_bypass").(string)),
		SecretScanningValidityChecks:          github.Ptr(d.Get("secret_scanning_validity_checks").(string)),
		SecretScanningNonProviderPatterns:     github.Ptr(d.Get("secret_scanning_non_provider_patterns").(string)),
		SecretScanningGenericSecrets:          github.Ptr(d.Get("secret_scanning_generic_secrets").(string)),
		SecretScanningDelegatedAlertDismissal: github.Ptr(d.Get("secret_scanning_delegated_alert_dismissal").(string)),
		PrivateVulnerabilityReporting:         github.Ptr(d.Get("private_vulnerability_reporting").(string)),
		Enforcement:                           github.Ptr(d.Get("enforcement").(string)),
	}

	advSec := d.Get("advanced_security").(string)
	codeSec := d.Get("code_security").(string)
	secretProt := d.Get("secret_protection").(string)

	if advSec == "enabled" {
		config.AdvancedSecurity = github.Ptr(advSec)
	} else if codeSec == "enabled" || secretProt == "enabled" {
		config.CodeSecurity = github.Ptr(codeSec)
		config.SecretProtection = github.Ptr(secretProt)
	} else {
		config.AdvancedSecurity = github.Ptr(advSec)
	}

	if v, ok := d.GetOk("dependency_graph_autosubmit_action_options"); ok {
		optionsList := v.([]any)
		if len(optionsList) > 0 {
			m := optionsList[0].(map[string]any)
			config.DependencyGraphAutosubmitActionOptions = &github.DependencyGraphAutosubmitActionOptions{
				LabeledRunners: github.Ptr(m["labeled_runners"].(bool)),
			}
		}
	}

	if v, ok := d.GetOk("code_scanning_default_setup_options"); ok {
		optionsList := v.([]any)
		if len(optionsList) > 0 {
			m := optionsList[0].(map[string]any)
			config.CodeScanningDefaultSetupOptions = &github.CodeScanningDefaultSetupOptions{
				RunnerType: m["runner_type"].(string),
			}
			if runnerLabel, ok := m["runner_label"].(string); ok && runnerLabel != "" {
				config.CodeScanningDefaultSetupOptions.RunnerLabel = github.Ptr(runnerLabel)
			}
		}
	}

	if v, ok := d.GetOk("code_scanning_options"); ok {
		optionsList := v.([]any)
		if len(optionsList) > 0 {
			m := optionsList[0].(map[string]any)
			config.CodeScanningOptions = &github.CodeScanningOptions{
				AllowAdvanced: github.Ptr(m["allow_advanced"].(bool)),
			}
		}
	}

	if v, ok := d.GetOk("secret_scanning_delegated_bypass_options"); ok {
		optionsList := v.([]any)
		if len(optionsList) > 0 {
			m := optionsList[0].(map[string]any)
			options := &github.SecretScanningDelegatedBypassOptions{}
			if reviewersV, ok := m["reviewers"]; ok {
				reviewersList := reviewersV.([]any)
				reviewers := make([]*github.BypassReviewer, 0, len(reviewersList))
				for _, rV := range reviewersList {
					rM := rV.(map[string]any)
					reviewers = append(reviewers, &github.BypassReviewer{
						ReviewerID:   int64(rM["reviewer_id"].(int)),
						ReviewerType: rM["reviewer_type"].(string),
					})
				}
				options.Reviewers = reviewers
			}
			config.SecretScanningDelegatedBypassOptions = options
		}
	}

	return config
}

func flattenDependencyGraphAutosubmitActionOptions(options *github.DependencyGraphAutosubmitActionOptions) []any {
	if options == nil {
		return []any{}
	}
	m := make(map[string]any)
	if options.LabeledRunners != nil {
		m["labeled_runners"] = *options.LabeledRunners
	}
	return []any{m}
}

func flattenCodeScanningDefaultSetupOptions(options *github.CodeScanningDefaultSetupOptions) []any {
	if options == nil {
		return []any{}
	}
	m := make(map[string]any)
	m["runner_type"] = options.RunnerType
	if options.RunnerLabel != nil {
		m["runner_label"] = *options.RunnerLabel
	}
	return []any{m}
}

func flattenCodeScanningOptions(options *github.CodeScanningOptions) []any {
	if options == nil {
		return []any{}
	}
	m := make(map[string]any)
	if options.AllowAdvanced != nil {
		m["allow_advanced"] = *options.AllowAdvanced
	}
	return []any{m}
}

func flattenSecretScanningDelegatedBypassOptions(options *github.SecretScanningDelegatedBypassOptions) []any {
	if options == nil {
		return []any{}
	}
	m := make(map[string]any)
	if options.Reviewers != nil {
		reviewers := make([]any, 0, len(options.Reviewers))
		for _, r := range options.Reviewers {
			rM := make(map[string]any)
			rM["reviewer_id"] = r.ReviewerID
			rM["reviewer_type"] = r.ReviewerType
			reviewers = append(reviewers, rM)
		}
		m["reviewers"] = reviewers
	}
	return []any{m}
}
