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

func resourceGithubCodeSecurityConfigurationCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	enterpriseSlug := d.Get("enterprise_slug").(string)
	body := buildCodeSecurityConfiguration(d)

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

	if err := d.Set("configuration_id", config.GetID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target_type", config.GetTargetType()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("html_url", config.GetHTMLURL()); err != nil {
		return diag.FromErr(err)
	}

	if v, ok := d.GetOk("default_for_new_repos"); ok {
		if err := setCodeSecurityConfigurationDefault(ctx, client, owner, enterpriseSlug, config.GetID(), v.(string)); err != nil {
			return diag.FromErr(err)
		}
	}

	if v, ok := d.GetOk("attach_scope"); ok {
		if err := attachCodeSecurityConfiguration(ctx, client, owner, enterpriseSlug, config.GetID(), v.(string)); err != nil {
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
	enterpriseSlug := d.Get("enterprise_slug").(string)

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

	if err := d.Set("name", config.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", config.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("advanced_security", config.GetAdvancedSecurity()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dependency_graph", config.GetDependencyGraph()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dependabot_alerts", config.GetDependabotAlerts()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dependabot_security_updates", config.GetDependabotSecurityUpdates()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("code_scanning_default_setup", config.GetCodeScanningDefaultSetup()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret_scanning", config.GetSecretScanning()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret_scanning_push_protection", config.GetSecretScanningPushProtection()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret_scanning_validity_checks", config.GetSecretScanningValidityChecks()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret_scanning_non_provider_patterns", config.GetSecretScanningNonProviderPatterns()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("private_vulnerability_reporting", config.GetPrivateVulnerabilityReporting()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enforcement", config.GetEnforcement()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("configuration_id", config.GetID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target_type", config.GetTargetType()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("html_url", config.GetHTMLURL()); err != nil {
		return diag.FromErr(err)
	}

	// default_for_new_repos is only surfaced via the /defaults listing.
	defaultForNewRepos, err := readCodeSecurityConfigurationDefault(ctx, client, owner, enterpriseSlug, configID)
	if err != nil {
		return diag.FromErr(err)
	}
	// A configuration that is not a default (or is explicitly "none") does not
	// appear in the defaults listing; preserve an explicit "none" in state.
	if defaultForNewRepos == "" && d.Get("default_for_new_repos").(string) == "none" {
		defaultForNewRepos = "none"
	}
	if err := d.Set("default_for_new_repos", defaultForNewRepos); err != nil {
		return diag.FromErr(err)
	}

	// attach_scope is write-only: the API does not expose the scope a
	// configuration was attached with, so it is intentionally not read back.

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
	enterpriseSlug := d.Get("enterprise_slug").(string)
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
		if err := d.Set("configuration_id", config.GetID()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("target_type", config.GetTargetType()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("html_url", config.GetHTMLURL()); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("default_for_new_repos") {
		newReposParam := d.Get("default_for_new_repos").(string)
		if newReposParam == "" {
			// Removing the attribute reverts the default to "none".
			newReposParam = "none"
		}
		if err := setCodeSecurityConfigurationDefault(ctx, client, owner, enterpriseSlug, configID, newReposParam); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("attach_scope") {
		if v, ok := d.GetOk("attach_scope"); ok {
			if err := attachCodeSecurityConfiguration(ctx, client, owner, enterpriseSlug, configID, v.(string)); err != nil {
				return diag.FromErr(err)
			}
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
	enterpriseSlug := d.Get("enterprise_slug").(string)

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

func attachCodeSecurityConfiguration(ctx context.Context, client *github.Client, org, enterpriseSlug string, configID int64, scope string) error {
	// The go-github v89 Attach helpers dereference the response status code even
	// when the request fails before a response exists, which panics on transport
	// errors. Issue the request directly so a nil response is handled safely.
	var u string
	if enterpriseSlug != "" {
		tflog.Debug(ctx, "Attaching code security configuration for enterprise", map[string]any{"configuration_id": configID, "scope": scope, "enterprise_slug": enterpriseSlug})
		u = fmt.Sprintf("enterprises/%s/code-security/configurations/%d/attach", enterpriseSlug, configID)
	} else {
		tflog.Debug(ctx, "Attaching code security configuration for organization", map[string]any{"configuration_id": configID, "scope": scope, "owner": org})
		u = fmt.Sprintf("orgs/%s/code-security/configurations/%d/attach", org, configID)
	}

	req, err := client.NewRequest(ctx, http.MethodPost, u, map[string]any{"scope": scope})
	if err != nil {
		return err
	}

	resp, err := client.Do(req, nil)
	if err != nil {
		// The attach endpoint responds 202 Accepted; go-github surfaces that as
		// an AcceptedError, which is a success for our purposes.
		if _, ok := errors.AsType[*github.AcceptedError](err); ok {
			return nil
		}
		if resp != nil && resp.StatusCode == http.StatusAccepted {
			return nil
		}
		return err
	}
	return nil
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
