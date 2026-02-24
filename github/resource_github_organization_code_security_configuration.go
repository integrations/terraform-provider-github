package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v66/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationCodeSecurityConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationCodeSecurityConfigurationCreate,
		Read:   resourceGithubOrganizationCodeSecurityConfigurationRead,
		Update: resourceGithubOrganizationCodeSecurityConfigurationUpdate,
		Delete: resourceGithubOrganizationCodeSecurityConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the code security configuration. Must be unique within the organization.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "A description of the code security configuration.",
			},
			"advanced_security": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of GitHub Advanced Security features. enabled will enable both Code Security or Secret Protection features. Can be one of: `enabled`, `disabled`, `code_security` and `secret_protection`.",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled", "code_security", "secret_protection"}),
			},
			"dependency_graph": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "enabled",
				Description:      "The enablement status of Dependency Graph. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"dependency_graph_autosubmit_action": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of Automatic dependency submission. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"dependency_graph_autosubmit_action_options": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The options for the Automatic dependency submission action.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labeled_runners": {
							Type:        schema.TypeBool,
							Default:     false,
							Optional:    true,
							Description: "Whether to use runners labeled with 'dependency-submission' or standard GitHub runners.",
						},
					},
				},
			},
			"dependabot_alerts": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of Dependabot alerts. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"dependabot_security_updates": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of Dependabot security updates. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"code_scanning_default_setup": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of code scanning default setup. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"secret_scanning": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of secret scanning. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"secret_scanning_push_protection": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of secret scanning push protection. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"secret_scanning_validity_checks": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of secret scanning validity checks. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"secret_scanning_non_provider_patterns": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of secret scanning non-provider patterns. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"private_vulnerability_reporting": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of private vulnerability reporting. Can be one of: `enabled` or `disabled`",
				ValidateDiagFunc: validateValueFunc([]string{"enabled", "disabled"}),
			},
			"enforcement": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "enforced",
				Description:      "The enforcement status for a security configuration. Can be one of: `enforced` or `unenforced`",
				ValidateDiagFunc: validateValueFunc([]string{"enforced", "unenforced"}),
			},
		},
	}
}

func resourceGithubOrganizationCodeSecurityConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	if err := checkOrganization(meta); err != nil {
		return err
	}

	options := &github.CodeSecurityConfiguration{
		Name:                              github.String(d.Get("name").(string)),
		Description:                       github.String(d.Get("description").(string)),
		AdvancedSecurity:                  github.String(d.Get("advanced_security").(string)),
		DependencyGraph:                   github.String(d.Get("dependency_graph").(string)),
		DependencyGraphAutosubmitAction:   github.String(d.Get("dependency_graph_autosubmit_action").(string)),
		DependabotAlerts:                  github.String(d.Get("dependabot_alerts").(string)),
		DependabotSecurityUpdates:         github.String(d.Get("dependabot_security_updates").(string)),
		CodeScanningDefaultSetup:          github.String(d.Get("code_scanning_default_setup").(string)),
		SecretScanningPushProtection:      github.String(d.Get("secret_scanning_push_protection").(string)),
		SecretScanningValidityChecks:      github.String(d.Get("secret_scanning_validity_checks").(string)),
		SecretScanningNonProviderPatterns: github.String(d.Get("secret_scanning_non_provider_patterns").(string)),
		PrivateVulnerabilityReporting:     github.String(d.Get("private_vulnerability_reporting").(string)),
		SecretScanning:                    github.String(d.Get("secret_scanning").(string)),
		Enforcement:                       github.String(d.Get("enforcement").(string)),
	}

	config, _, err := client.Organizations.CreateCodeSecurityConfiguration(ctx, orgName, options)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprint(config.GetID()))
	return resourceGithubOrganizationCodeSecurityConfigurationRead(d, meta)
}

func resourceGithubOrganizationCodeSecurityConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()
	if err := checkOrganization(meta); err != nil {
		return err
	}
	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	config, _, err := client.Organizations.GetCodeSecurityConfiguration(ctx, orgName, configID)
	if err != nil {
		return err
	}

	d.Set("name", config.GetName())
	d.Set("description", config.GetDescription())
	d.Set("advanced_security", config.GetAdvancedSecurity())
	d.Set("dependency_graph", config.GetDependencyGraph())
	d.Set("dependency_graph_autosubmit_action", config.GetDependencyGraphAutosubmitAction())
	d.Set("dependabot_alerts", config.GetDependabotAlerts())
	d.Set("dependabot_security_updates", config.GetDependabotSecurityUpdates())
	d.Set("code_scanning_default_setup", config.GetCodeScanningDefaultSetup())
	d.Set("secret_scanning_push_protection", config.GetSecretScanningPushProtection())
	d.Set("secret_scanning_validity_checks", config.GetSecretScanningValidityChecks())
	d.Set("secret_scanning_non_provider_patterns", config.GetSecretScanningNonProviderPatterns())
	d.Set("private_vulnerability_reporting", config.GetPrivateVulnerabilityReporting())
	d.Set("secret_scanning", config.GetSecretScanning())
	d.Set("enforcement", config.GetEnforcement())

	return nil
}
func resourceGithubOrganizationCodeSecurityConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()
	if err := checkOrganization(meta); err != nil {
		return err
	}
	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	options := resourceGithubOrganizationCodeSecurityConfigurationObject(d)

	config, _, err := client.Organizations.UpdateCodeSecurityConfiguration(ctx, orgName, configID, options)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprint(config.GetID()))
	return resourceGithubOrganizationCodeSecurityConfigurationRead(d, meta)
}
func resourceGithubOrganizationCodeSecurityConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()
	if err := checkOrganization(meta); err != nil {
		return err
	}
	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	_, err = client.Organizations.DeleteCodeSecurityConfiguration(ctx, orgName, configID)
	return err
}

func resourceGithubOrganizationCodeSecurityConfigurationObject(d *schema.ResourceData) *github.CodeSecurityConfiguration {
	config := &github.CodeSecurityConfiguration{
		Name:                              github.String(d.Get("name").(string)),
		Description:                       github.String(d.Get("description").(string)),
		AdvancedSecurity:                  github.String(d.Get("advanced_security").(string)),
		DependencyGraph:                   github.String(d.Get("dependency_graph").(string)),
		DependencyGraphAutosubmitAction:   github.String(d.Get("dependency_graph_autosubmit_action").(string)),
		DependabotAlerts:                  github.String(d.Get("dependabot_alerts").(string)),
		DependabotSecurityUpdates:         github.String(d.Get("dependabot_security_updates").(string)),
		CodeScanningDefaultSetup:          github.String(d.Get("code_scanning_default_setup").(string)),
		SecretScanningPushProtection:      github.String(d.Get("secret_scanning_push_protection").(string)),
		SecretScanningValidityChecks:      github.String(d.Get("secret_scanning_validity_checks").(string)),
		SecretScanningNonProviderPatterns: github.String(d.Get("secret_scanning_non_provider_patterns").(string)),
		PrivateVulnerabilityReporting:     github.String(d.Get("private_vulnerability_reporting").(string)),
		SecretScanning:                    github.String(d.Get("secret_scanning").(string)),
		Enforcement:                       github.String(d.Get("enforcement").(string)),
	}

	if options := d.Get("dependency_graph_autosubmit_action_options"); options != nil && len(options.([]interface{})) > 0 {
		graphOptions := options.([]interface{})[0].(map[string]interface{})
		config.DependencyGraphAutosubmitActionOptions = &github.DependencyGraphAutosubmitActionOptions{
			LabeledRunners: github.Bool(graphOptions["labeled_runners"].(bool)),
		}
	}

	return config
}
