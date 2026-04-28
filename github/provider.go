package github

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("GITHUB_AUTH_MODE", nil),
				Description:  descriptions["auth_mode"],
				ValidateFunc: validation.StringInSlice([]string{"anonymous", "token", "app"}, false),
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_TOKEN", nil),
				Description: descriptions["token"],
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_OWNER", nil),
				Description: descriptions["owner"],
			},
			"retryable_errors": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				DefaultFunc: func() (any, error) {
					defaultErrors := []int{500, 502, 503, 504}
					errorInterfaces := make([]any, len(defaultErrors))
					for i, v := range defaultErrors {
						errorInterfaces[i] = v
					}
					return errorInterfaces, nil
				},
				Description: descriptions["retryable_errors"],
			},
			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     3,
				Description: descriptions["max_retries"],
			},
			"organization": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_ORGANIZATION", nil),
				Description: descriptions["organization"],
				Deprecated:  "Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION)",
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_BASE_URL", "https://api.github.com/"),
				Description: descriptions["base_url"],
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["insecure"],
			},
			"write_delay_ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1000,
				Description: descriptions["write_delay_ms"],
			},
			"read_delay_ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: descriptions["read_delay_ms"],
			},
			"retry_delay_ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1000,
				Description: descriptions["retry_delay_ms"],
			},
			"parallel_requests": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["parallel_requests"],
			},
			"app_auth": {
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				Description:   descriptions["app_auth"],
				Deprecated:    "Use top-level app_id, app_installation_id, and app_private_key instead.",
				ConflictsWith: []string{"app_id", "app_installation_id", "app_private_key"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_ID", nil),
							Description: descriptions["app_auth.id"],
						},
						"installation_id": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_INSTALLATION_ID", nil),
							Description: descriptions["app_auth.installation_id"],
						},
						"pem_file": {
							Type:        schema.TypeString,
							Required:    true,
							Sensitive:   true,
							DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_PEM_FILE", nil),
							Description: descriptions["app_auth.pem_file"],
						},
					},
				},
			},
			"app_id": {
				Type:          schema.TypeString,
				Optional:      true,
				DefaultFunc:   schema.EnvDefaultFunc("GITHUB_APP_ID", nil),
				Description:   descriptions["app_id"],
				ConflictsWith: []string{"app_auth"},
			},
			"app_installation_id": {
				Type:          schema.TypeString,
				Optional:      true,
				DefaultFunc:   schema.EnvDefaultFunc("GITHUB_APP_INSTALLATION_ID", nil),
				Description:   descriptions["app_installation_id"],
				ConflictsWith: []string{"app_auth"},
			},
			"app_private_key": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				DefaultFunc:   schema.EnvDefaultFunc("GITHUB_APP_PRIVATE_KEY", nil),
				Description:   descriptions["app_private_key"],
				ConflictsWith: []string{"app_auth"},
			},
			// https://developer.github.com/guides/traversing-with-pagination/#basics-of-pagination
			"max_per_page": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_MAX_PER_PAGE", "100"),
				Description: descriptions["max_per_page"],
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"github_enterprise_actions_permissions":                                 resourceGithubActionsEnterprisePermissions(),
			"github_actions_environment_secret":                                     resourceGithubActionsEnvironmentSecret(),
			"github_actions_environment_variable":                                   resourceGithubActionsEnvironmentVariable(),
			"github_actions_organization_oidc_subject_claim_customization_template": resourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplate(),
			"github_actions_organization_permissions":                               resourceGithubActionsOrganizationPermissions(),
			"github_actions_organization_secret":                                    resourceGithubActionsOrganizationSecret(),
			"github_actions_organization_secret_repositories":                       resourceGithubActionsOrganizationSecretRepositories(),
			"github_actions_organization_secret_repository":                         resourceGithubActionsOrganizationSecretRepository(),
			"github_actions_organization_variable":                                  resourceGithubActionsOrganizationVariable(),
			"github_actions_organization_variable_repositories":                     resourceGithubActionsOrganizationVariableRepositories(),
			"github_actions_organization_variable_repository":                       resourceGithubActionsOrganizationVariableRepository(),
			"github_actions_repository_access_level":                                resourceGithubActionsRepositoryAccessLevel(),
			"github_actions_repository_oidc_subject_claim_customization_template":   resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate(),
			"github_actions_repository_permissions":                                 resourceGithubActionsRepositoryPermissions(),
			"github_actions_runner_group":                                           resourceGithubActionsRunnerGroup(),
			"github_actions_hosted_runner":                                          resourceGithubActionsHostedRunner(),
			"github_actions_secret":                                                 resourceGithubActionsSecret(),
			"github_actions_variable":                                               resourceGithubActionsVariable(),
			"github_app_installation_repositories":                                  resourceGithubAppInstallationRepositories(),
			"github_app_installation_repository":                                    resourceGithubAppInstallationRepository(),
			"github_branch":                                                         resourceGithubBranch(),
			"github_branch_default":                                                 resourceGithubBranchDefault(),
			"github_branch_protection":                                              resourceGithubBranchProtection(),
			"github_branch_protection_v3":                                           resourceGithubBranchProtectionV3(),
			"github_codespaces_organization_secret":                                 resourceGithubCodespacesOrganizationSecret(),
			"github_codespaces_organization_secret_repositories":                    resourceGithubCodespacesOrganizationSecretRepositories(),
			"github_codespaces_secret":                                              resourceGithubCodespacesSecret(),
			"github_codespaces_user_secret":                                         resourceGithubCodespacesUserSecret(),
			"github_dependabot_organization_secret":                                 resourceGithubDependabotOrganizationSecret(),
			"github_dependabot_organization_secret_repositories":                    resourceGithubDependabotOrganizationSecretRepositories(),
			"github_dependabot_organization_secret_repository":                      resourceGithubDependabotOrganizationSecretRepository(),
			"github_dependabot_secret":                                              resourceGithubDependabotSecret(),
			"github_emu_group_mapping":                                              resourceGithubEMUGroupMapping(),
			"github_issue":                                                          resourceGithubIssue(),
			"github_issue_label":                                                    resourceGithubIssueLabel(),
			"github_issue_labels":                                                   resourceGithubIssueLabels(),
			"github_membership":                                                     resourceGithubMembership(),
			"github_organization_block":                                             resourceOrganizationBlock(),
			"github_organization_custom_role":                                       resourceGithubOrganizationCustomRole(),
			"github_organization_custom_properties":                                 resourceGithubOrganizationCustomProperties(),
			"github_organization_project":                                           resourceGithubOrganizationProject(),
			"github_organization_repository_role":                                   resourceGithubOrganizationRepositoryRole(),
			"github_organization_role":                                              resourceGithubOrganizationRole(),
			"github_organization_role_team":                                         resourceGithubOrganizationRoleTeam(),
			"github_organization_role_user":                                         resourceGithubOrganizationRoleUser(),
			"github_organization_role_team_assignment":                              resourceGithubOrganizationRoleTeamAssignment(),
			"github_organization_ruleset":                                           resourceGithubOrganizationRuleset(),
			"github_organization_security_manager":                                  resourceGithubOrganizationSecurityManager(),
			"github_organization_settings":                                          resourceGithubOrganizationSettings(),
			"github_organization_webhook":                                           resourceGithubOrganizationWebhook(),
			"github_project_card":                                                   resourceGithubProjectCard(),
			"github_project_column":                                                 resourceGithubProjectColumn(),
			"github_release":                                                        resourceGithubRelease(),
			"github_repository":                                                     resourceGithubRepository(),
			"github_repository_autolink_reference":                                  resourceGithubRepositoryAutolinkReference(),
			"github_repository_dependabot_security_updates":                         resourceGithubRepositoryDependabotSecurityUpdates(),
			"github_repository_collaborator":                                        resourceGithubRepositoryCollaborator(),
			"github_repository_collaborators":                                       resourceGithubRepositoryCollaborators(),
			"github_repository_custom_property":                                     resourceGithubRepositoryCustomProperty(),
			"github_repository_deploy_key":                                          resourceGithubRepositoryDeployKey(),
			"github_repository_deployment_branch_policy":                            resourceGithubRepositoryDeploymentBranchPolicy(),
			"github_repository_environment":                                         resourceGithubRepositoryEnvironment(),
			"github_repository_environment_deployment_policy":                       resourceGithubRepositoryEnvironmentDeploymentPolicy(),
			"github_repository_file":                                                resourceGithubRepositoryFile(),
			"github_repository_milestone":                                           resourceGithubRepositoryMilestone(),
			"github_repository_pages":                                               resourceGithubRepositoryPages(),
			"github_repository_project":                                             resourceGithubRepositoryProject(),
			"github_repository_pull_request":                                        resourceGithubRepositoryPullRequest(),
			"github_repository_ruleset":                                             resourceGithubRepositoryRuleset(),
			"github_repository_topics":                                              resourceGithubRepositoryTopics(),
			"github_repository_webhook":                                             resourceGithubRepositoryWebhook(),
			"github_repository_vulnerability_alerts":                                resourceGithubRepositoryVulnerabilityAlerts(),
			"github_team":                                                           resourceGithubTeam(),
			"github_team_members":                                                   resourceGithubTeamMembers(),
			"github_team_membership":                                                resourceGithubTeamMembership(),
			"github_team_repository":                                                resourceGithubTeamRepository(),
			"github_team_settings":                                                  resourceGithubTeamSettings(),
			"github_team_sync_group_mapping":                                        resourceGithubTeamSyncGroupMapping(),
			"github_user_gpg_key":                                                   resourceGithubUserGpgKey(),
			"github_user_invitation_accepter":                                       resourceGithubUserInvitationAccepter(),
			"github_user_ssh_key":                                                   resourceGithubUserSshKey(),
			"github_enterprise_organization":                                        resourceGithubEnterpriseOrganization(),
			"github_enterprise_actions_runner_group":                                resourceGithubActionsEnterpriseRunnerGroup(),
			"github_enterprise_ip_allow_list_entry":                                 resourceGithubEnterpriseIpAllowListEntry(),
			"github_enterprise_actions_workflow_permissions":                        resourceGithubEnterpriseActionsWorkflowPermissions(),
			"github_actions_organization_workflow_permissions":                      resourceGithubActionsOrganizationWorkflowPermissions(),
			"github_enterprise_security_analysis_settings":                          resourceGithubEnterpriseSecurityAnalysisSettings(),
			"github_workflow_repository_permissions":                                resourceGithubWorkflowRepositoryPermissions(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"github_actions_environment_public_key":                                 dataSourceGithubActionsEnvironmentPublicKey(),
			"github_actions_environment_secrets":                                    dataSourceGithubActionsEnvironmentSecrets(),
			"github_actions_environment_variables":                                  dataSourceGithubActionsEnvironmentVariables(),
			"github_actions_organization_oidc_subject_claim_customization_template": dataSourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplate(),
			"github_actions_organization_public_key":                                dataSourceGithubActionsOrganizationPublicKey(),
			"github_actions_organization_registration_token":                        dataSourceGithubActionsOrganizationRegistrationToken(),
			"github_actions_organization_secrets":                                   dataSourceGithubActionsOrganizationSecrets(),
			"github_actions_organization_variables":                                 dataSourceGithubActionsOrganizationVariables(),
			"github_actions_public_key":                                             dataSourceGithubActionsPublicKey(),
			"github_actions_registration_token":                                     dataSourceGithubActionsRegistrationToken(),
			"github_actions_repository_oidc_subject_claim_customization_template":   dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate(),
			"github_actions_secrets":                                                dataSourceGithubActionsSecrets(),
			"github_actions_variables":                                              dataSourceGithubActionsVariables(),
			"github_app":                                                            dataSourceGithubApp(),
			"github_app_token":                                                      dataSourceGithubAppToken(),
			"github_branch":                                                         dataSourceGithubBranch(),
			"github_branch_protection_rules":                                        dataSourceGithubBranchProtectionRules(),
			"github_collaborators":                                                  dataSourceGithubCollaborators(),
			"github_codespaces_organization_public_key":                             dataSourceGithubCodespacesOrganizationPublicKey(),
			"github_codespaces_organization_secrets":                                dataSourceGithubCodespacesOrganizationSecrets(),
			"github_codespaces_public_key":                                          dataSourceGithubCodespacesPublicKey(),
			"github_codespaces_secrets":                                             dataSourceGithubCodespacesSecrets(),
			"github_codespaces_user_public_key":                                     dataSourceGithubCodespacesUserPublicKey(),
			"github_codespaces_user_secrets":                                        dataSourceGithubCodespacesUserSecrets(),
			"github_dependabot_organization_public_key":                             dataSourceGithubDependabotOrganizationPublicKey(),
			"github_dependabot_organization_secrets":                                dataSourceGithubDependabotOrganizationSecrets(),
			"github_dependabot_public_key":                                          dataSourceGithubDependabotPublicKey(),
			"github_dependabot_secrets":                                             dataSourceGithubDependabotSecrets(),
			"github_external_groups":                                                dataSourceGithubExternalGroups(),
			"github_ip_ranges":                                                      dataSourceGithubIpRanges(),
			"github_issue_labels":                                                   dataSourceGithubIssueLabels(),
			"github_membership":                                                     dataSourceGithubMembership(),
			"github_organization":                                                   dataSourceGithubOrganization(),
			"github_organization_custom_role":                                       dataSourceGithubOrganizationCustomRole(),
			"github_organization_custom_properties":                                 dataSourceGithubOrganizationCustomProperties(),
			"github_organization_external_identities":                               dataSourceGithubOrganizationExternalIdentities(),
			"github_organization_ip_allow_list":                                     dataSourceGithubOrganizationIpAllowList(),
			"github_organization_repository_role":                                   dataSourceGithubOrganizationRepositoryRole(),
			"github_organization_repository_roles":                                  dataSourceGithubOrganizationRepositoryRoles(),
			"github_organization_role":                                              dataSourceGithubOrganizationRole(),
			"github_organization_role_teams":                                        dataSourceGithubOrganizationRoleTeams(),
			"github_organization_role_users":                                        dataSourceGithubOrganizationRoleUsers(),
			"github_organization_roles":                                             dataSourceGithubOrganizationRoles(),
			"github_organization_security_managers":                                 dataSourceGithubOrganizationSecurityManagers(),
			"github_organization_team_sync_groups":                                  dataSourceGithubOrganizationTeamSyncGroups(),
			"github_organization_teams":                                             dataSourceGithubOrganizationTeams(),
			"github_organization_webhooks":                                          dataSourceGithubOrganizationWebhooks(),
			"github_organization_app_installations":                                 dataSourceGithubOrganizationAppInstallations(),
			"github_ref":                                                            dataSourceGithubRef(),
			"github_release":                                                        dataSourceGithubRelease(),
			"github_release_asset":                                                  dataSourceGithubReleaseAsset(),
			"github_repositories":                                                   dataSourceGithubRepositories(),
			"github_repository":                                                     dataSourceGithubRepository(),
			"github_repository_autolink_references":                                 dataSourceGithubRepositoryAutolinkReferences(),
			"github_repository_branches":                                            dataSourceGithubRepositoryBranches(),
			"github_repository_custom_properties":                                   dataSourceGithubRepositoryCustomProperties(),
			"github_repository_environments":                                        dataSourceGithubRepositoryEnvironments(),
			"github_repository_deploy_keys":                                         dataSourceGithubRepositoryDeployKeys(),
			"github_repository_deployment_branch_policies":                          dataSourceGithubRepositoryDeploymentBranchPolicies(),
			"github_repository_file":                                                dataSourceGithubRepositoryFile(),
			"github_repository_milestone":                                           dataSourceGithubRepositoryMilestone(),
			"github_repository_pages":                                               dataSourceGithubRepositoryPages(),
			"github_repository_pull_request":                                        dataSourceGithubRepositoryPullRequest(),
			"github_repository_pull_requests":                                       dataSourceGithubRepositoryPullRequests(),
			"github_repository_teams":                                               dataSourceGithubRepositoryTeams(),
			"github_repository_webhooks":                                            dataSourceGithubRepositoryWebhooks(),
			"github_rest_api":                                                       dataSourceGithubRestApi(),
			"github_ssh_keys":                                                       dataSourceGithubSshKeys(),
			"github_team":                                                           dataSourceGithubTeam(),
			"github_tree":                                                           dataSourceGithubTree(),
			"github_user":                                                           dataSourceGithubUser(),
			"github_user_external_identity":                                         dataSourceGithubUserExternalIdentity(),
			"github_users":                                                          dataSourceGithubUsers(),
			"github_enterprise":                                                     dataSourceGithubEnterprise(),
			"github_repository_environment_deployment_policies":                     dataSourceGithubRepositoryEnvironmentDeploymentPolicies(),
		},
	}

	p.ConfigureContextFunc = providerConfigure(p)

	return p
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_mode": "Explicit authentication mode. Valid values are `anonymous`, `token`, and `app`. " +
			"When not set, the provider auto-detects the mode based on provided credentials for backward compatibility.",

		"token": "The OAuth token used to connect to GitHub. " +
			"When `auth_mode` is not set, anonymous mode is enabled if no credentials are provided.",

		"base_url": "The GitHub Base API URL",

		"insecure": "Enable `insecure` mode for testing purposes",

		"owner": "The GitHub owner name to manage. " +
			"Use this field instead of `organization` when managing individual accounts.",

		"organization": "The GitHub organization name to manage. " +
			"Use this field instead of `owner` when managing organization accounts.",

		"app_auth": "Deprecated: use top-level `app_id`, `app_installation_id`, and `app_private_key` instead. " +
			"The GitHub App credentials used to connect to GitHub.",
		"app_auth.id":              "The GitHub App ID.",
		"app_auth.installation_id": "The GitHub App installation instance ID.",
		"app_auth.pem_file":        "The GitHub App PEM file contents.",
		"app_id":                   "The GitHub App ID.",
		"app_installation_id":      "The GitHub App installation instance ID.",
		"app_private_key":          "The GitHub App private key in PEM format.",
		"write_delay_ms": "Amount of time in milliseconds to sleep in between writes to GitHub API. " +
			"Defaults to 1000ms or 1s if not set.",
		"read_delay_ms": "Amount of time in milliseconds to sleep in between non-write requests to GitHub API. " +
			"Defaults to 0ms if not set.",
		"retry_delay_ms": "Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. " +
			"Defaults to 1000ms or 1s if not set, the max_retries must be set to greater than zero.",
		"parallel_requests": "Allow the provider to make parallel API calls to GitHub. " +
			"You may want to set it to true when you have a private Github Enterprise without strict rate limits. " +
			"While it is possible to enable this setting on github.com, " +
			"github.com's best practices recommend using serialization to avoid hitting abuse rate limits" +
			"Defaults to false if not set",
		"retryable_errors": "Allow the provider to retry after receiving an error status code, the max_retries should be set for this to work" +
			"Defaults to [500, 502, 503, 504]",
		"max_retries": "Number of times to retry a request after receiving an error status code" +
			"Defaults to 3",
		"max_per_page": "Number of items per page for pagination" +
			"Defaults to 100",
	}
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		var diags diag.Diagnostics

		owner := d.Get("owner").(string)
		insecure := d.Get("insecure").(bool)
		authMode := d.Get("auth_mode").(string)

		// BEGIN backwards compatibility
		// OwnerOrOrgEnvDefaultFunc used to be the default value for both
		// 'owner' and 'organization'. This meant that if 'owner' and
		// 'GITHUB_OWNER' were set, 'GITHUB_OWNER' would be used as the default
		// value of 'organization' and therefore override 'owner'.
		//
		// This seems undesirable (an environment variable should not override
		// an explicitly set value in a provider block), but is necessary
		// for backwards compatibility. We could remove this backwards compatibility
		// code in a future major release.
		env := ownerOrOrgEnvDefaultFunc()
		if env.(string) != "" {
			owner = env.(string)
		}
		// END backwards compatibility

		baseURL, isGHES, err := getBaseURL(d.Get("base_url").(string))
		if err != nil {
			return nil, diag.FromErr(err)
		}

		org := d.Get("organization").(string)
		if org != "" {
			tflog.Info(ctx, "Selecting organization attribute as owner", map[string]any{"owner": org})
			owner = org
		}

		var token string

		switch authMode {
		case "anonymous":
			tflog.Info(ctx, "Auth mode: anonymous")

		case "token":
			token = d.Get("token").(string)
			if token == "" {
				return nil, diag.Errorf(
					"auth_mode is set to \"token\" but no token was provided; " +
						"set the `token` argument or `GITHUB_TOKEN` environment variable")
			}
			tflog.Info(ctx, "Auth mode: token")

		case "app":
			appID, appInstallationID, appPemFile := getAppCredentials(d)
			var missingFields []string
			if appID == "" {
				missingFields = append(missingFields, "app_id (GITHUB_APP_ID)")
			}
			if appInstallationID == "" {
				missingFields = append(missingFields, "app_installation_id (GITHUB_APP_INSTALLATION_ID)")
			}
			if appPemFile == "" {
				missingFields = append(missingFields, "app_private_key (GITHUB_APP_PRIVATE_KEY)")
			}
			if len(missingFields) > 0 {
				return nil, diag.Errorf(
					"auth_mode is set to \"app\" but the following app credentials are missing: %s",
					strings.Join(missingFields, ", "))
			}

			apiPath := ""
			if isGHES {
				apiPath = GHESRESTAPIPath
			}

			appToken, err := GenerateOAuthTokenFromApp(baseURL.JoinPath(apiPath), appID, appInstallationID, appPemFile)
			if err != nil {
				return nil, wrapErrors([]error{err})
			}

			token = appToken
			tflog.Info(ctx, "Auth mode: app", map[string]any{"app_id": appID, "installation_id": appInstallationID})

		default: // auto-detect (backward compatibility)
			token = d.Get("token").(string)

			if token == "" {
				// Top-level app fields require an explicit auth_mode.
				// Skip this check when app_auth is configured for backward compatibility.
				appAuth, _ := d.Get("app_auth").([]any)
				hasAppAuth := len(appAuth) > 0 && appAuth[0] != nil
				topLevelAppSet := d.Get("app_id").(string) != "" ||
					d.Get("app_installation_id").(string) != "" ||
					d.Get("app_private_key").(string) != ""
				if topLevelAppSet && !hasAppAuth {
					return nil, diag.Errorf(
						"top-level app credentials (app_id, app_installation_id, app_private_key) " +
							"require auth_mode = \"app\" to be set explicitly; use the `auth_mode` " +
							"provider argument or the GITHUB_AUTH_MODE environment variable")
				}

				appID, appInstallationID, appPemFile := getAppCredentials(d)
				if appID != "" && appInstallationID != "" && appPemFile != "" {
					apiPath := ""
					if isGHES {
						apiPath = GHESRESTAPIPath
					}

					appToken, err := GenerateOAuthTokenFromApp(baseURL.JoinPath(apiPath), appID, appInstallationID, appPemFile)
					if err != nil {
						return nil, wrapErrors([]error{err})
					}
					token = appToken
					tflog.Info(ctx, "Auth mode: app", map[string]any{"app_id": appID, "installation_id": appInstallationID})
				}
			}

			if token == "" {
				tflog.Info(ctx, "No token found, trying GitHub CLI to get token", map[string]any{"hostname": baseURL.Host})
				ghToken := tokenFromGHCLI(baseURL)
				if ghToken != "" {
					token = ghToken
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  "GitHub CLI token fallback is deprecated",
						Detail: "Automatic token detection from `gh auth token` is deprecated and will be removed in a future major release. " +
							"Please set the `token` provider argument or `GITHUB_TOKEN` environment variable explicitly. " +
							"You can use `export GITHUB_TOKEN=$(gh auth token)` as a replacement.",
					})
				}
			}
		}

		writeDelay := d.Get("write_delay_ms").(int)
		if writeDelay <= 0 {
			return nil, wrapErrors([]error{fmt.Errorf("write_delay_ms must be greater than 0ms")})
		}
		tflog.Info(ctx, "Setting write_delay_ms", map[string]any{"write_delay_ms": writeDelay})

		readDelay := d.Get("read_delay_ms").(int)
		if readDelay < 0 {
			return nil, wrapErrors([]error{fmt.Errorf("read_delay_ms must be greater than or equal to 0ms")})
		}
		tflog.Debug(ctx, "Setting read_delay_ms", map[string]any{"read_delay_ms": readDelay})

		retryDelay := d.Get("read_delay_ms").(int)
		if retryDelay < 0 {
			return nil, diag.Errorf("retry_delay_ms must be greater than or equal to 0ms")
		}
		tflog.Debug(ctx, "Setting retry_delay_ms", map[string]any{"retry_delay_ms": retryDelay})

		maxRetries := d.Get("max_retries").(int)
		if maxRetries < 0 {
			return nil, diag.Errorf("max_retries must be greater than or equal to 0")
		}
		tflog.Debug(ctx, "Setting max_retries", map[string]any{"max_retries": maxRetries})
		retryableErrors := make(map[int]bool)
		if maxRetries > 0 {
			reParam := d.Get("retryable_errors").([]any)
			if len(reParam) == 0 {
				retryableErrors = getDefaultRetriableErrors()
			} else {
				for _, status := range reParam {
					retryableErrors[status.(int)] = true
				}
			}

			tflog.Debug(ctx, "Setting retryable_errors", map[string]any{"retryable_errors": retryableErrors})
		}

		_maxPerPage := d.Get("max_per_page").(int)
		if _maxPerPage <= 0 {
			return nil, diag.Errorf("max_per_page must be greater than than 0")
		}
		tflog.Debug(ctx, "Setting max_per_page", map[string]any{"max_per_page": _maxPerPage})
		maxPerPage = _maxPerPage

		parallelRequests := d.Get("parallel_requests").(bool)

		tflog.Debug(ctx, "Setting parallel_requests", map[string]any{"parallel_requests": parallelRequests})

		config := Config{
			Token:            token,
			BaseURL:          baseURL,
			Insecure:         insecure,
			Owner:            owner,
			WriteDelay:       time.Duration(writeDelay) * time.Millisecond,
			ReadDelay:        time.Duration(readDelay) * time.Millisecond,
			RetryDelay:       time.Duration(retryDelay) * time.Millisecond,
			RetryableErrors:  retryableErrors,
			MaxRetries:       maxRetries,
			ParallelRequests: parallelRequests,
			IsGHES:           isGHES,
		}

		meta, err := config.Meta()
		if err != nil {
			return nil, wrapErrors([]error{err})
		}

		return meta, diags
	}
}

func getAppCredentials(d *schema.ResourceData) (appID, appInstallationID, appPemFile string) {
	// Try top-level fields first
	if v, ok := d.Get("app_id").(string); ok && v != "" {
		appID = v
	}
	if v, ok := d.Get("app_installation_id").(string); ok && v != "" {
		appInstallationID = v
	}
	if v, ok := d.Get("app_private_key").(string); ok && v != "" {
		// The Go encoding/pem package only decodes PEM formatted blocks
		// that contain new lines. Some platforms, like Terraform Cloud,
		// do not support new lines within Environment Variables.
		// Any occurrence of \n in the `app_private_key` argument's value
		// (explicit value, or default value taken from
		// GITHUB_APP_PRIVATE_KEY Environment Variable) is replaced with an
		// actual new line character before decoding.
		appPemFile = strings.ReplaceAll(v, `\n`, "\n")
	}

	// Fall back to app_auth block for any missing values
	if appID == "" || appInstallationID == "" || appPemFile == "" {
		if appAuth, ok := d.Get("app_auth").([]any); ok && len(appAuth) > 0 && appAuth[0] != nil {
			appAuthAttr := appAuth[0].(map[string]any)
			if appID == "" {
				if v, ok := appAuthAttr["id"].(string); ok && v != "" {
					appID = v
				}
			}
			if appInstallationID == "" {
				if v, ok := appAuthAttr["installation_id"].(string); ok && v != "" {
					appInstallationID = v
				}
			}
			if appPemFile == "" {
				if v, ok := appAuthAttr["pem_file"].(string); ok && v != "" {
					appPemFile = strings.ReplaceAll(v, `\n`, "\n")
				}
			}
		}
	}

	return
}

// ghCLIHostFromAPIHost maps an API hostname to the corresponding
// gh-CLI --hostname value.  For example api.github.com -> github.com
// and api.<slug>.ghe.com -> <slug>.ghe.com.
// for unrecognized hostnames, input is returned unmodified.
func ghCLIHostFromAPIHost(host string) string {
	if host == DotComAPIHost {
		return DotComHost
	} else if GHECAPIHostMatch.MatchString(host) {
		return strings.TrimPrefix(host, "api.")
	}
	return host
}

// See https://github.com/integrations/terraform-provider-github/issues/1822
func tokenFromGHCLI(u *url.URL) string {
	ghCliPath := os.Getenv("GH_PATH")
	if ghCliPath == "" {
		ghCliPath = "gh"
	}

	host := ghCLIHostFromAPIHost(u.Host)

	out, err := exec.Command(ghCliPath, "auth", "token", "--hostname", host).Output()
	if err != nil {
		log.Printf("[DEBUG] Error getting token from GitHub CLI: %s", err.Error())
		// GH CLI is either not installed or there was no `gh auth login` command issued,
		// which is fine. don't return the error to keep the flow going
		return ""
	}

	log.Printf("[INFO] Using the token from GitHub CLI")
	return strings.TrimSpace(string(out))
}

func ownerOrOrgEnvDefaultFunc() any {
	if organization := os.Getenv("GITHUB_ORGANIZATION"); organization != "" {
		log.Printf("[INFO] Selecting owner %s from GITHUB_ORGANIZATION environment variable", organization)
		return organization
	}
	owner := os.Getenv("GITHUB_OWNER")
	log.Printf("[INFO] Selecting owner %s from GITHUB_OWNER environment variable", owner)
	return owner
}
