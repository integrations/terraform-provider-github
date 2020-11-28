package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			PROVIDER_TOKEN: {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{PROVIDER_TOKEN, PROVIDER_APP},
				DefaultFunc:  schema.EnvDefaultFunc("GITHUB_TOKEN", nil),
				Description:  descriptions["token"],
			},
			PROVIDER_OWNER: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: OwnerOrOrgEnvDefaultFunc,
				Description: descriptions["owner"],
			},
			PROVIDER_ORGANIZATION: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: OwnerOrOrgEnvDefaultFunc,
				Description: descriptions["organization"],
			},
			PROVIDER_BASE_URL: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_BASE_URL", "https://api.github.com/"),
				Description: descriptions["base_url"],
			},
			PROVIDER_INSECURE: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["insecure"],
			},
			PROVIDER_APP: {
				Type:         schema.TypeList,
				Optional:     true,
				ExactlyOneOf: []string{PROVIDER_TOKEN, PROVIDER_APP},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PROVIDER_APP_ID: {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_ID", nil),
							Description: "The GitHub App ID.",
						},
						PROVIDER_APP_INSTALLATION_ID: {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_INSTALLATION_ID", nil),
							Description: "The GitHub App installation instance ID.",
						},
						PROVIDER_APP_PEM: {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("GITHUB_APP_PEM", nil),
							Description: "The GitHub App PEM string.",
							Sensitive:   true,
						},
					},
				},
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"github_actions_organization_secret": resourceGithubActionsOrganizationSecret(),
			"github_actions_secret":              resourceGithubActionsSecret(),
			"github_branch":                      resourceGithubBranch(),
			"github_branch_protection":           resourceGithubBranchProtection(),
			"github_issue_label":                 resourceGithubIssueLabel(),
			"github_membership":                  resourceGithubMembership(),
			"github_organization_block":          resourceOrganizationBlock(),
			"github_organization_project":        resourceGithubOrganizationProject(),
			"github_organization_webhook":        resourceGithubOrganizationWebhook(),
			"github_project_card":                resourceGithubProjectCard(),
			"github_project_column":              resourceGithubProjectColumn(),
			"github_repository_collaborator":     resourceGithubRepositoryCollaborator(),
			"github_repository_deploy_key":       resourceGithubRepositoryDeployKey(),
			"github_repository_file":             resourceGithubRepositoryFile(),
			"github_repository_milestone":        resourceGithubRepositoryMilestone(),
			"github_repository_project":          resourceGithubRepositoryProject(),
			"github_repository_webhook":          resourceGithubRepositoryWebhook(),
			"github_repository":                  resourceGithubRepository(),
			"github_team_membership":             resourceGithubTeamMembership(),
			"github_team_repository":             resourceGithubTeamRepository(),
			"github_team_sync_group_mapping":     resourceGithubTeamSyncGroupMapping(),
			"github_team":                        resourceGithubTeam(),
			"github_user_gpg_key":                resourceGithubUserGpgKey(),
			"github_user_invitation_accepter":    resourceGithubUserInvitationAccepter(),
			"github_user_ssh_key":                resourceGithubUserSshKey(),
			"github_branch_default":              resourceGithubBranchDefault(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"github_actions_public_key":            dataSourceGithubActionsPublicKey(),
			"github_branch":                        dataSourceGithubBranch(),
			"github_collaborators":                 dataSourceGithubCollaborators(),
			"github_ip_ranges":                     dataSourceGithubIpRanges(),
			"github_membership":                    dataSourceGithubMembership(),
			"github_organization":                  dataSourceGithubOrganization(),
			"github_organization_team_sync_groups": dataSourceGithubOrganizationTeamSyncGroups(),
			"github_release":                       dataSourceGithubRelease(),
			"github_repositories":                  dataSourceGithubRepositories(),
			"github_repository":                    dataSourceGithubRepository(),
			"github_repository_milestone":          dataSourceGithubRepositoryMilestone(),
			"github_team":                          dataSourceGithubTeam(),
			"github_user":                          dataSourceGithubUser(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"token": "The OAuth token used to connect to GitHub. " +
			"`anonymous` mode is enabled if `token` is not configured.",

		"base_url": "The GitHub Base API URL",

		"insecure": "Enable `insecure` mode for testing purposes",

		"owner": "The GitHub owner name to manage. " +
			"Use this field instead of `organization` when managing individual accounts.",

		"organization": "The GitHub organization name to manage. " +
			"Use this field instead of `owner` when managing organization accounts.",
	}
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		anonymous := true
		if d.Get("token").(string) != "" {
			anonymous = false
		}

		individual := true
		if d.Get("organization").(string) != "" {
			individual = false
		}

		owner := d.Get("owner").(string)
		if !individual {
			owner = d.Get("organization").(string)
		}

		config := Config{
			Token:        d.Get("token").(string),
			Organization: d.Get("organization").(string),
			BaseURL:      d.Get("base_url").(string),
			Insecure:     d.Get("insecure").(bool),
			Owner:        owner,
			Individual:   individual,
			Anonymous:    anonymous,
		}

		meta, err := config.Meta()
		if err != nil {
			return nil, err
		}

		meta.(*Owner).StopContext = p.StopContext()

		return meta, nil
	}
}
