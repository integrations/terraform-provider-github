package github

import (
	"context"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
					DefaultFunc: schema.EnvDefaultFunc("GITHUB_ORGANIZATION", nil),
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
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     0,
					Description: "The delay in milliseconds between read operations; this defaults to `0`. This can be used to mitigate rate limiting issues when performing a large number of read operations.",
				},
				"write_delay_ms": {
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     1000,
					Description: "The delay in milliseconds between write operations; this defaults to `1000`. This is used to mitigate the GitHub API's abuse rate limits when writing. Note that **ALL** requests to the GraphQL API are implemented as `POST` requests under the hood, so this setting affects those calls as well.",
				},
				"retry_delay_ms": {
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     1000,
					Description: "The delay in milliseconds between retry attempts; this defaults to `1000`. This setting only applies when `max_retries` is greater than `0`.",
				},
				"retryable_errors": {
					Type:        schema.TypeList,
					Elem:        &schema.Schema{Type: schema.TypeInt},
					Optional:    true,
					Description: "List of HTTP status codes that should be retried; if not set this uses the provider defaults. This setting only applies when `max_retries` is greater than `0`.",
				},
				"max_retries": {
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     3,
					Description: "The maximum number of retries for failed requests; this defaults to `3`.",
				},
				"max_per_page": {
					Type:        schema.TypeInt,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("GITHUB_MAX_PER_PAGE", 100),
					Description: "The maximum number of results per page for paginated API requests; this defaults to `100`. This can also be set by the `GITHUB_MAX_PER_PAGE` environment variable.",
				},
				"parallel_requests": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "Allow the provider to make parallel API calls; this is experimental and may cause concurrency and rate limiting issues.",
				},
				"insecure": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "Allow insecure server connections when using SSL.",
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
				"github_copilot_organization_seat_assignment":                           resourceGithubCopilotOrganizationSeatAssignment(),
				"github_copilot_team_seat_assignment":                                   resourceGithubCopilotTeamSeatAssignment(),
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

func configureProvider() func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		owner := d.Get("owner").(string)
		token := d.Get("token").(string)
		insecure := d.Get("insecure").(bool)

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
			log.Printf("[INFO] Selecting organization attribute as owner: %s", org)
			owner = org
		}

		if appAuth, ok := d.Get("app_auth").([]any); ok && len(appAuth) > 0 && appAuth[0] != nil {
			appAuthAttr := appAuth[0].(map[string]any)

			var appID, appInstallationID, appPemFile string

			if v, ok := appAuthAttr["id"].(string); ok && v != "" {
				appID = v
			} else {
				return nil, diag.Errorf("app_auth.id must be set and contain a non-empty value")
			}

			if v, ok := appAuthAttr["installation_id"].(string); ok && v != "" {
				appInstallationID = v
			} else {
				return nil, diag.Errorf("app_auth.installation_id must be set and contain a non-empty value")
			}

			if v, ok := appAuthAttr["pem_file"].(string); ok && v != "" {
				// The Go encoding/pem package only decodes PEM formatted blocks
				// that contain new lines. Some platforms, like Terraform Cloud,
				// do not support new lines within Environment Variables.
				// Any occurrence of \n in the `pem_file` argument's value
				// (explicit value, or default value taken from
				// GITHUB_APP_PEM_FILE Environment Variable) is replaced with an
				// actual new line character before decoding.
				appPemFile = strings.ReplaceAll(v, `\n`, "\n")
			} else {
				return nil, diag.Errorf("app_auth.pem_file must be set and contain a non-empty value")
			}

			apiPath := ""
			if isGHES {
				apiPath = GHESRESTAPIPath
			}

			appToken, err := GenerateOAuthTokenFromApp(baseURL.JoinPath(apiPath), appID, appInstallationID, appPemFile)
			if err != nil {
				return nil, diag.FromErr(err)
			}

			token = appToken
		}

		if token == "" {
			log.Printf("[INFO] No token found, using GitHub CLI to get token from hostname %s", baseURL.Host)
			token = tokenFromGHCLI(baseURL)
		}

		writeDelay := d.Get("write_delay_ms").(int)
		if writeDelay <= 0 {
			return nil, diag.Errorf("write_delay_ms must be greater than 0ms")
		}
		log.Printf("[INFO] Setting write_delay_ms to %d", writeDelay)

		readDelay := d.Get("read_delay_ms").(int)
		if readDelay < 0 {
			return nil, diag.Errorf("read_delay_ms must be greater than or equal to 0ms")
		}
		log.Printf("[DEBUG] Setting read_delay_ms to %d", readDelay)

		retryDelay, _ := d.Get("retry_delay_ms").(int)
		if retryDelay < 0 {
			return nil, diag.Errorf("retry_delay_ms must be greater than or equal to 0ms")
		}
		log.Printf("[DEBUG] Setting retry_delay_ms to %d", retryDelay)

		maxRetries := d.Get("max_retries").(int)
		if maxRetries < 0 {
			return nil, diag.Errorf("max_retries must be greater than or equal to 0")
		}
		log.Printf("[DEBUG] Setting max_retries to %d", maxRetries)
		retryableErrors := make(map[int]bool)
		if maxRetries > 0 {
			reParam := d.Get("retryable_errors").([]any)
			if len(reParam) == 0 {
				retryableErrors = getDefaultRetryableErrors()
			} else {
				for _, status := range reParam {
					retryableErrors[status.(int)] = true
				}
			}

			log.Printf("[DEBUG] Setting retryableErrors to %v", retryableErrors)
		}

		_maxPerPage := d.Get("max_per_page").(int)
		if _maxPerPage <= 0 {
			return nil, diag.Errorf("max_per_page must be greater than 0")
		}
		log.Printf("[DEBUG] Setting max_per_page to %d", _maxPerPage)
		maxPerPage = _maxPerPage

		parallelRequests := d.Get("parallel_requests").(bool)

		log.Printf("[DEBUG] Setting parallel_requests to %t", parallelRequests)

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
			return nil, diag.FromErr(err)
		}

		return meta, nil
	}
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

// getDefaultRetryableErrors returns the default set of retryable errors.
func getDefaultRetryableErrors() map[int]bool {
	return map[int]bool{
		500: true,
		502: true,
		503: true,
		504: true,
	}
}
