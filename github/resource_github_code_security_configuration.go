package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubCodeSecurityConfiguration() *schema.Resource {
	featureValueDiag := validation.ToDiagFunc(validation.StringInSlice([]string{"enabled", "disabled", "not_set"}, false))

	return &schema.Resource{
		Description: "Manages a GitHub Code Security Configuration at the organization or enterprise level.",
		Create:      resourceGithubCodeSecurityConfigurationCreate,
		Read:        resourceGithubCodeSecurityConfigurationRead,
		Update:      resourceGithubCodeSecurityConfigurationUpdate,
		Delete:      resourceGithubCodeSecurityConfigurationDelete,
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
			"dependency_graph": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "enabled",
				Description:      "The enablement status of Dependency Graph. Can be 'enabled', 'disabled' or 'not_set'.",
				ValidateDiagFunc: featureValueDiag,
			},
			"dependabot_alerts": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "The enablement status of Dependabot alerts. Can be 'enabled', 'disabled' or 'not_set'.",
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
			"attach_scope": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The scope of repositories to attach the configuration to. Can be 'all' or 'all_without_configurations'. Attachment is applied on create and whenever this value changes; it cannot be read back from the API, so removing this attribute does not detach repositories.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "all_without_configurations"}, false)),
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

func buildCodeSecurityConfiguration(d *schema.ResourceData) github.CodeSecurityConfiguration {
	return github.CodeSecurityConfiguration{
		Name:                              d.Get("name").(string),
		Description:                       d.Get("description").(string),
		AdvancedSecurity:                  new(d.Get("advanced_security").(string)),
		DependencyGraph:                   new(d.Get("dependency_graph").(string)),
		DependabotAlerts:                  new(d.Get("dependabot_alerts").(string)),
		DependabotSecurityUpdates:         new(d.Get("dependabot_security_updates").(string)),
		CodeScanningDefaultSetup:          new(d.Get("code_scanning_default_setup").(string)),
		SecretScanning:                    new(d.Get("secret_scanning").(string)),
		SecretScanningPushProtection:      new(d.Get("secret_scanning_push_protection").(string)),
		SecretScanningValidityChecks:      new(d.Get("secret_scanning_validity_checks").(string)),
		SecretScanningNonProviderPatterns: new(d.Get("secret_scanning_non_provider_patterns").(string)),
		PrivateVulnerabilityReporting:     new(d.Get("private_vulnerability_reporting").(string)),
		Enforcement:                       new(d.Get("enforcement").(string)),
	}
}

func resourceGithubCodeSecurityConfigurationCreate(d *schema.ResourceData, meta any) error {
	owner := meta.(*Owner)
	client := owner.v3client
	ctx := context.Background()

	enterpriseSlug := d.Get("enterprise_slug").(string)
	body := buildCodeSecurityConfiguration(d)

	var config *github.CodeSecurityConfiguration
	var err error
	if enterpriseSlug != "" {
		log.Printf("[DEBUG] Creating code security configuration %q for enterprise: %s", body.Name, enterpriseSlug)
		config, _, err = client.Enterprise.CreateCodeSecurityConfiguration(ctx, enterpriseSlug, body)
	} else {
		if err := checkOrganization(meta); err != nil {
			return err
		}
		log.Printf("[DEBUG] Creating code security configuration %q for organization: %s", body.Name, owner.name)
		config, _, err = client.Organizations.CreateCodeSecurityConfiguration(ctx, owner.name, body)
	}
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(config.GetID(), 10))

	if v, ok := d.GetOk("default_for_new_repos"); ok {
		if err := setCodeSecurityConfigurationDefault(ctx, client, owner.name, enterpriseSlug, config.GetID(), v.(string)); err != nil {
			return err
		}
	}

	if v, ok := d.GetOk("attach_scope"); ok {
		if err := attachCodeSecurityConfiguration(ctx, client, owner.name, enterpriseSlug, config.GetID(), v.(string)); err != nil {
			return err
		}
	}

	return resourceGithubCodeSecurityConfigurationRead(d, meta)
}

func resourceGithubCodeSecurityConfigurationRead(d *schema.ResourceData, meta any) error {
	owner := meta.(*Owner)
	client := owner.v3client
	ctx := context.Background()

	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	enterpriseSlug := d.Get("enterprise_slug").(string)

	var config *github.CodeSecurityConfiguration
	if enterpriseSlug != "" {
		config, _, err = client.Enterprise.GetCodeSecurityConfiguration(ctx, enterpriseSlug, configID)
	} else {
		config, _, err = client.Organizations.GetCodeSecurityConfiguration(ctx, owner.name, configID)
	}
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			log.Printf("[INFO] Removing code security configuration %s from state because it no longer exists in GitHub", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	if err := d.Set("name", config.Name); err != nil {
		return err
	}
	if err := d.Set("description", config.Description); err != nil {
		return err
	}
	if err := d.Set("advanced_security", config.GetAdvancedSecurity()); err != nil {
		return err
	}
	if err := d.Set("dependency_graph", config.GetDependencyGraph()); err != nil {
		return err
	}
	if err := d.Set("dependabot_alerts", config.GetDependabotAlerts()); err != nil {
		return err
	}
	if err := d.Set("dependabot_security_updates", config.GetDependabotSecurityUpdates()); err != nil {
		return err
	}
	if err := d.Set("code_scanning_default_setup", config.GetCodeScanningDefaultSetup()); err != nil {
		return err
	}
	if err := d.Set("secret_scanning", config.GetSecretScanning()); err != nil {
		return err
	}
	if err := d.Set("secret_scanning_push_protection", config.GetSecretScanningPushProtection()); err != nil {
		return err
	}
	if err := d.Set("secret_scanning_validity_checks", config.GetSecretScanningValidityChecks()); err != nil {
		return err
	}
	if err := d.Set("secret_scanning_non_provider_patterns", config.GetSecretScanningNonProviderPatterns()); err != nil {
		return err
	}
	if err := d.Set("private_vulnerability_reporting", config.GetPrivateVulnerabilityReporting()); err != nil {
		return err
	}
	if err := d.Set("enforcement", config.GetEnforcement()); err != nil {
		return err
	}
	if err := d.Set("configuration_id", config.GetID()); err != nil {
		return err
	}
	if err := d.Set("target_type", config.GetTargetType()); err != nil {
		return err
	}
	if err := d.Set("html_url", config.GetHTMLURL()); err != nil {
		return err
	}

	// default_for_new_repos is only surfaced via the /defaults listing.
	defaultForNewRepos, err := readCodeSecurityConfigurationDefault(ctx, client, owner.name, enterpriseSlug, configID)
	if err != nil {
		return err
	}
	// A configuration that is not a default (or is explicitly "none") does not
	// appear in the defaults listing; preserve an explicit "none" in state.
	if defaultForNewRepos == "" && d.Get("default_for_new_repos").(string) == "none" {
		defaultForNewRepos = "none"
	}
	if err := d.Set("default_for_new_repos", defaultForNewRepos); err != nil {
		return err
	}

	// attach_scope is write-only: the API does not expose the scope a
	// configuration was attached with, so it is intentionally not read back.

	return nil
}

func resourceGithubCodeSecurityConfigurationUpdate(d *schema.ResourceData, meta any) error {
	owner := meta.(*Owner)
	client := owner.v3client
	ctx := context.Background()

	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	enterpriseSlug := d.Get("enterprise_slug").(string)
	body := buildCodeSecurityConfiguration(d)

	if enterpriseSlug != "" {
		log.Printf("[DEBUG] Updating code security configuration %d for enterprise: %s", configID, enterpriseSlug)
		_, _, err = client.Enterprise.UpdateCodeSecurityConfiguration(ctx, enterpriseSlug, configID, body)
	} else {
		log.Printf("[DEBUG] Updating code security configuration %d for organization: %s", configID, owner.name)
		_, _, err = client.Organizations.UpdateCodeSecurityConfiguration(ctx, owner.name, configID, body)
	}
	if err != nil {
		return err
	}

	if d.HasChange("default_for_new_repos") {
		newReposParam := d.Get("default_for_new_repos").(string)
		if newReposParam == "" {
			// Removing the attribute reverts the default to "none".
			newReposParam = "none"
		}
		if err := setCodeSecurityConfigurationDefault(ctx, client, owner.name, enterpriseSlug, configID, newReposParam); err != nil {
			return err
		}
	}

	if d.HasChange("attach_scope") {
		if v, ok := d.GetOk("attach_scope"); ok {
			if err := attachCodeSecurityConfiguration(ctx, client, owner.name, enterpriseSlug, configID, v.(string)); err != nil {
				return err
			}
		}
	}

	return resourceGithubCodeSecurityConfigurationRead(d, meta)
}

func resourceGithubCodeSecurityConfigurationDelete(d *schema.ResourceData, meta any) error {
	owner := meta.(*Owner)
	client := owner.v3client
	ctx := context.Background()

	configID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	enterpriseSlug := d.Get("enterprise_slug").(string)

	if enterpriseSlug != "" {
		log.Printf("[DEBUG] Deleting code security configuration %d for enterprise: %s", configID, enterpriseSlug)
		_, err = client.Enterprise.DeleteCodeSecurityConfiguration(ctx, enterpriseSlug, configID)
	} else {
		log.Printf("[DEBUG] Deleting code security configuration %d for organization: %s", configID, owner.name)
		_, err = client.Organizations.DeleteCodeSecurityConfiguration(ctx, owner.name, configID)
	}
	return err
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
		log.Printf("[DEBUG] Setting code security configuration %d as default (%s) for enterprise: %s", configID, newReposParam, enterpriseSlug)
		_, _, err = client.Enterprise.SetDefaultCodeSecurityConfiguration(ctx, enterpriseSlug, configID, newReposParam)
	} else {
		log.Printf("[DEBUG] Setting code security configuration %d as default (%s) for organization: %s", configID, newReposParam, org)
		_, _, err = client.Organizations.SetDefaultCodeSecurityConfiguration(ctx, org, configID, newReposParam)
	}
	return err
}

func attachCodeSecurityConfiguration(ctx context.Context, client *github.Client, org, enterpriseSlug string, configID int64, scope string) error {
	var err error
	if enterpriseSlug != "" {
		log.Printf("[DEBUG] Attaching code security configuration %d with scope %s for enterprise: %s", configID, scope, enterpriseSlug)
		_, err = client.Enterprise.AttachCodeSecurityConfigurationToRepositories(ctx, enterpriseSlug, configID, scope)
	} else {
		log.Printf("[DEBUG] Attaching code security configuration %d with scope %s for organization: %s", configID, scope, org)
		_, err = client.Organizations.AttachCodeSecurityConfigurationToRepositories(ctx, org, configID, scope, nil)
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
