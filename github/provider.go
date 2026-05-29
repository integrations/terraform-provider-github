package github

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/integrations/terraform-provider-github/v6/internal/ghclient"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

// NewProvider returns a function that returns the schema.Provider for this provider.
func NewProvider() func() *schema.Provider {
	return func() *schema.Provider {
		return &schema.Provider{
			Schema: map[string]*schema.Schema{
				"base_url": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("GITHUB_BASE_URL", DotComAPIURL),
					Description: "The base URL for the GitHub API; this defaults to the GitHub API URL. If you are using GitHub Enterprise Server (GHES) or GitHub Enterprise Cloud with Data Residency (GHEC-DR), this is required. This can also be set by the `GITHUB_BASE_URL` environment variable.",
				},
				"owner": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("GITHUB_OWNER", nil),
					Description: "GitHub organization or user account to manage; this is required when authenticating using a GitHub App. If the owner is not provided and a token is provided, the provider will attempt to auto-detect the owner associated with the token. This can also be set by the `GITHUB_OWNER` environment variable.",
				},
				"organization": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "GitHub organization to manage. This can also be set by the `GITHUB_ORGANIZATION` environment variable.",
					Deprecated:  "This argument is deprecated and will be removed in a future major release; use `owner` instead.",
				},
				"token": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("GITHUB_TOKEN", nil),
					Description: "GitHub OAuth or Personal Access Token (PAT) to use for authentication. This can also be set by the `GITHUB_TOKEN` environment variable.",
					// ConflictsWith: []string{"app_auth"}, // TODO: Enable as part of v7.
				},
				"app_auth": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Authenticate using a GitHub App.",
					// ConflictsWith: []string{"token"}, // TODO: Enable as part of v7.
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"id": {
								Type:        schema.TypeString,
								Required:    true,
								DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_ID", nil),
								Description: "The GitHub App's identifier. This can also be set by the `GITHUB_APP_ID` environment variable.",
							},
							"installation_id": {
								Type:        schema.TypeString,
								Required:    true,
								DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_INSTALLATION_ID", nil),
								Description: "The GitHub App's installation identifier. This can also be set by the `GITHUB_APP_INSTALLATION_ID` environment variable.",
							},
							"pem_file": {
								Type:        schema.TypeString,
								Required:    true,
								Sensitive:   true,
								DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_PEM_FILE", nil),
								Description: "The GitHub App's PEM file content; `\\n` can be used for newlines. This can also be set by the `GITHUB_APP_PEM_FILE` environment variable.",
							},
						},
					},
				},
				"read_delay_ms": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          0,
					ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
					Description:      "The delay in milliseconds between read operations; this defaults to `0`. This can be used to mitigate rate limiting issues when performing a large number of read operations. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation is GitHub rate limit aware.",
				},
				"write_delay_ms": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          1000,
					ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
					Description:      "The delay in milliseconds between write operations; this defaults to `1000`. This is used to mitigate the GitHub API's abuse rate limits when writing. Note that **ALL** requests to the GraphQL API are implemented as `POST` requests under the hood, so this setting affects those calls as well. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation is GitHub rate limit aware.",
				},
				"retry_delay_ms": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          1000,
					ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
					Description:      "The delay in milliseconds between retry attempts; this defaults to `1000`. This setting only applies when `max_retries` is greater than `0`.",
				},
				"retryable_errors": {
					Type:        schema.TypeList,
					Elem:        &schema.Schema{Type: schema.TypeInt},
					Optional:    true,
					Description: "List of HTTP status codes that should be retried; if not set this uses the provider defaults. This setting only applies when `max_retries` is greater than `0`. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation handles the retry logic.",
				},
				"max_retries": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          3,
					ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
					Description:      "The maximum number of retries for failed requests; this defaults to `3`.",
				},
				"max_per_page": {
					Type:             schema.TypeInt,
					Optional:         true,
					DefaultFunc:      schema.EnvDefaultFunc("GITHUB_MAX_PER_PAGE", 100),
					ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
					Description:      "The maximum number of results per page for paginated API requests; this defaults to `100`. This can also be set by the `GITHUB_MAX_PER_PAGE` environment variable.",
				},
				"parallel_requests": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "Allow the provider to make parallel API calls; this is experimental and may cause concurrency and rate limiting issues. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation is designed to safely handle parallel requests.",
				},
				"insecure": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "Allow insecure server connections when using SSL.",
					Deprecated:  "This argument is deprecated as it's currently not used and will be removed in the next major version.",
				},
				"legacy_client": {
					Type:        schema.TypeBool,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("GITHUB_LEGACY_CLIENT", true),
					Description: "Use the legacy GitHub client implementation; if set to `false`, the new client implementation is used. This can also be set by the `GITHUB_LEGACY_CLIENT` environment variable.",
				},
				"cache_path": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("GITHUB_CACHE_PATH", ""),
					Description: "The path to the cache directory for persisting GitHub API requests between runs; if not set there will be no caching between runs. This can also be set by the `GITHUB_CACHE_PATH` environment variable.",
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

			ConfigureContextFunc: configureProvider(),
		}
	}
}

// configureProvider initializes the provider meta parameter with the necessary clients and owner information based on the provided configuration. It returns the initialized meta parameter or an error if the configuration is invalid or if there are issues initializing the clients.
func configureProvider() func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		config := &Config{
			GraphQLAPIPath: "graphql",
		}

		if v, ok := d.GetOk("legacy_client"); ok {
			if b, ok := v.(bool); ok {
				config.LegacyClient = b
			}
		}

		if v, ok := d.GetOk("base_url"); ok {
			if s, ok := v.(string); ok && s != "" {
				baseURL, isGHES, err := getBaseURL(s)
				if err != nil {
					return nil, diag.FromErr(err)
				}

				tflog.Debug(ctx, "Using base URL from provider configuration.", map[string]any{"base_url": baseURL.String()})
				config.BaseURL = baseURL

				if isGHES {
					tflog.Debug(ctx, "Base URL indicates GitHub Enterprise Server (GHES) usage; enabling GHES mode.", map[string]any{"base_url": baseURL.String()})
					config.RESTAPIPath = GHESRESTAPIPath
					config.GraphQLAPIPath = GHESGraphQLAPIPath
				}
			}
		}

		// TODO: In v7 remove organization and the associated backwards compatibility code, and require owner to be set either via the provider configuration or GITHUB_OWNER environment variable with the provider configuration taking precedence.
		if v, ok := d.GetOk("organization"); ok {
			if s, ok := v.(string); ok && s != "" {
				tflog.Debug(ctx, "Using organization environment variable or attribute as owner.", map[string]any{"owner": s})
				config.Owner = s
			}
		}

		if config.Owner == "" {
			if s, ok := os.LookupEnv("GITHUB_OWNER"); ok && s != "" {
				tflog.Debug(ctx, "Using GITHUB_OWNER environment variable as owner.", map[string]any{"owner": s})
				config.Owner = s
			}
		}

		if config.Owner == "" {
			if v, ok := d.GetOk("owner"); ok {
				if s, ok := v.(string); ok && s != "" {
					tflog.Debug(ctx, "Using owner attribute as owner.", map[string]any{"owner": s})
					config.Owner = s
				}
			}
		}

		if appID, appInstallationID, appPEM, ok := getAppAuth(d); ok {
			tflog.Debug(ctx, "Using GitHub App authentication.", map[string]any{"app_id": appID, "app_installation_id": appInstallationID})
			config.AppID = appID
			config.AppInstallationID = appInstallationID
			config.AppPEM = appPEM
		}

		if config.AppID == nil {
			if _, ok := d.GetOk("app_auth"); ok {
				return nil, diag.Errorf("app_auth block is set but required fields are missing or contains empty values")
			}

			if v, ok := d.GetOk("token"); ok {
				if s, ok := v.(string); ok && s != "" {
					tflog.Debug(ctx, "Using token from provider configuration.")
					config.Token = s
				}
			}
		}

		if config.Owner == "" && config.AppID != nil {
			return nil, diag.Errorf("owner must be set for github app authentication")
		}

		if config.LegacyClient {
			if config.AppID != nil {
				appToken, err := GenerateOAuthTokenFromApp(config.BaseURL.JoinPath(config.RESTAPIPath), *config.AppID, *config.AppInstallationID, string(config.AppPEM))
				if err != nil {
					return nil, diag.FromErr(err)
				}
				config.Token = appToken
			}

			if config.Token == "" {
				tflog.Debug(ctx, "No token found, using GitHub CLI to get token from base URL.", map[string]any{"base_url": config.BaseURL.String()})
				config.Token = tokenFromGHCLI(ctx, config.BaseURL)
			}
		}

		if v, ok := d.GetOk("read_delay_ms"); ok {
			if i, ok := v.(int); ok {
				tflog.Debug(ctx, "Using read delay from provider configuration.", map[string]any{"read_delay_ms": i})
				config.ReadDelay = time.Duration(i) * time.Millisecond
			}
		}

		if v, ok := d.GetOk("write_delay_ms"); ok {
			if i, ok := v.(int); ok {
				tflog.Debug(ctx, "Using write delay from provider configuration.", map[string]any{"write_delay_ms": i})
				config.WriteDelay = time.Duration(i) * time.Millisecond
			}
		}

		if v, ok := d.GetOk("retry_delay_ms"); ok {
			if i, ok := v.(int); ok {
				tflog.Debug(ctx, "Using retry delay from provider configuration.", map[string]any{"retry_delay_ms": i})
				config.RetryDelay = time.Duration(i) * time.Millisecond
			}
		}

		if v, ok := d.GetOk("retryable_errors"); ok {
			if c, ok := v.([]any); ok && len(c) > 0 {
				retryableErrors := make(map[int]bool)
				for _, status := range c {
					i, ok := status.(int)
					if !ok {
						return nil, diag.Errorf("retryable_errors must be a list of integers")
					}
					retryableErrors[i] = true
				}

				tflog.Debug(ctx, "Using retryable errors from provider configuration.", map[string]any{"retryable_errors": retryableErrors})
				config.RetryableErrors = retryableErrors
			}
		}

		if config.RetryableErrors == nil {
			config.RetryableErrors = getDefaultRetryableErrors()
		}

		if v, ok := d.GetOk("max_retries"); ok {
			if i, ok := v.(int); ok {
				tflog.Debug(ctx, "Using max retries from provider configuration.", map[string]any{"max_retries": i})
				config.MaxRetries = i
			}
		}

		if v, ok := d.GetOk("max_per_page"); ok {
			if i, ok := v.(int); ok {
				tflog.Debug(ctx, "Using max per page from provider configuration.", map[string]any{"max_per_page": i})
				// TODO: Move max per page to the provider metadata and remove the global variable.
				maxPerPage = i
			}
		}

		if v, ok := d.GetOk("parallel_requests"); ok {
			if b, ok := v.(bool); ok && b {
				tflog.Warn(ctx, "Parallel requests are enabled; this may cause concurrency and rate limiting issues.")
				config.ParallelRequests = true
			}
		}

		if v, ok := d.GetOk("insecure"); ok {
			if b, ok := v.(bool); ok && b {
				tflog.Warn(ctx, "Insecure mode enabled; SSL certificate verification is disabled. This is not recommended for production environments.")
				config.Insecure = true
			}
		}

		if v, ok := d.GetOk("cache_path"); ok {
			if s, ok := v.(string); ok && s != "" {
				tflog.Debug(ctx, "Using cache path from provider configuration.", map[string]any{"cache_path": s})
				config.CachePath = &s
			}
		}

		meta, err := configureProviderMeta(ctx, config)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return meta, nil
	}
}

// configureProviderMeta initializes the provider metadata, including setting up the GitHub API clients based on the provided configuration. It returns the initialized metadata or an error if the configuration is invalid or if there are issues initializing the clients.
func configureProviderMeta(ctx context.Context, c *Config) (*Owner, error) {
	owner := &Owner{
		name: c.Owner,
	}

	if c.LegacyClient {
		var client *http.Client
		if c.Anonymous() {
			client = c.AnonymousHTTPClient()
		} else {
			client = c.AuthenticatedHTTPClient()
		}

		v3client, err := c.NewRESTClient(client)
		if err != nil {
			return nil, err
		}
		owner.v3client = v3client

		v4client, err := c.NewGraphQLClient(client)
		if err != nil {
			return nil, err
		}
		owner.v4client = v4client
	} else {
		options := ghclient.Options{
			RESTAPIURL:   new(c.BaseURL.JoinPath(c.RESTAPIPath).String()),
			GraphQLURL:   new(c.BaseURL.JoinPath(c.GraphQLAPIPath).String()),
			CachePath:    c.CachePath,
			RetryMax:     c.MaxRetries,
			RetryWaitMin: c.RetryDelay,
			RetryWaitMax: c.RetryDelay,
		}

		var source ghclient.Source
		if c.AppID != nil {
			appSource, err := ghclient.NewAppSource(*c.AppID, c.AppPEM, options)
			if err != nil {
				return nil, fmt.Errorf("failed to create app source: %w", err)
			}
			source = appSource
		} else if c.Token != "" {
			tokenSource, err := ghclient.NewTokenSource(c.Token, options)
			if err != nil {
				return nil, fmt.Errorf("failed to create token source: %w", err)
			}
			source = tokenSource
		} else {
			anonymousSource, err := ghclient.NewAnonymousSource(options)
			if err != nil {
				return nil, fmt.Errorf("failed to create anonymous source: %w", err)
			}
			source = anonymousSource
		}

		v3client, err := source.OwnerRESTClient(ctx, owner.name)
		if err != nil {
			return nil, fmt.Errorf("failed to create rest client for owner %q: %w", owner.name, err)
		}

		v4client, err := source.OwnerGraphQLClient(ctx, owner.name)
		if err != nil {
			return nil, fmt.Errorf("failed to create graphql client for owner %q: %w", owner.name, err)
		}

		owner.v3client = v3client
		owner.v4client = v4client
	}

	if owner.name == "" && c.Token != "" {
		user, _, err := owner.v3client.Users.Get(ctx, "")
		if err != nil {
			return nil, err
		}
		owner.name = user.GetLogin()
	}

	if org, _, err := owner.v3client.Organizations.Get(ctx, owner.name); err == nil && org != nil {
		owner.id = org.GetID()
		owner.IsOrganization = true
	}

	return owner, nil
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
func tokenFromGHCLI(ctx context.Context, u *url.URL) string {
	ghCliPath, ok := os.LookupEnv("GH_PATH")
	if !ok {
		ghCliPath = "gh"
	}

	host := ghCLIHostFromAPIHost(u.Host)

	out, err := exec.CommandContext(ctx, ghCliPath, "auth", "token", "--hostname", host).Output()
	if err != nil {
		tflog.Debug(ctx, "Error getting token from GitHub CLI; ensure GitHub CLI is installed and authenticated if you intend to use it for authentication.", map[string]any{"error": err.Error()})
		// GH CLI is either not installed or there was no `gh auth login` command issued,
		// which is fine. don't return the error to keep the flow going
		return ""
	}

	tflog.Info(ctx, "Using the token from GitHub CLI.")
	return strings.TrimSpace(string(out))
}

// getAppAuth retrieves GitHub App authentication parameters from the provider configuration, environment variables, or defaults, and validates them. It returns the app ID, installation ID, PEM file content, and a boolean indicating whether valid app authentication parameters were found.
func getAppAuth(d *schema.ResourceData) (*string, *string, []byte, bool) {
	appID := os.Getenv("GITHUB_APP_ID")
	appInstallationID := os.Getenv("GITHUB_APP_INSTALLATION_ID")
	appPEM := os.Getenv("GITHUB_APP_PEM_FILE")

	v, ok := d.GetOk("app_auth")
	if !ok {
		return validateAppAuth(appID, appInstallationID, appPEM)
	}

	c, ok := v.([]any)
	if !ok || len(c) == 0 || c[0] == nil {
		return validateAppAuth(appID, appInstallationID, appPEM)
	}

	appAuthAttr, ok := c[0].(map[string]any)
	if !ok {
		return validateAppAuth(appID, appInstallationID, appPEM)
	}

	if o, ok := appAuthAttr["id"]; ok {
		if s, ok := o.(string); ok && s != "" {
			appID = s
		}
	}

	if o, ok := appAuthAttr["installation_id"]; ok {
		if s, ok := o.(string); ok && s != "" {
			appInstallationID = s
		}
	}

	if o, ok := appAuthAttr["pem_file"]; ok {
		if s, ok := o.(string); ok && s != "" {
			appPEM = s
		}
	}

	return validateAppAuth(appID, appInstallationID, appPEM)
}

// validateAppAuth checks if the provided app authentication parameters are valid (non-empty) and returns them along with a boolean indicating validity.
func validateAppAuth(appID, appInstallationID, appPEM string) (*string, *string, []byte, bool) {
	if appID == "" || appInstallationID == "" || appPEM == "" {
		return nil, nil, nil, false
	}

	return &appID, &appInstallationID, []byte(strings.ReplaceAll(appPEM, `\n`, "\n")), true
}

// getDefaultRetryableErrors returns the default set of retryable errors.
func getDefaultRetryableErrors() map[int]bool {
	return map[int]bool{
		500: true,
		502: true,
		503: true,
		504: true,
	}
}
