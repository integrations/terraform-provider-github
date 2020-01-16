package github

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	var p *schema.Provider
	// The actual provider
	p = &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_TOKEN", nil),
				Description: descriptions["token"],
			},
			"owner": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_OWNER", nil),
				Description: descriptions["owner"],
			},
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_BASE_URL", ""),
				Description: descriptions["base_url"],
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"github_team":                    resourceGithubTeam(),
			"github_team_membership":         resourceGithubTeamMembership(),
			"github_team_repository":         resourceGithubTeamRepository(),
			"github_membership":              resourceGithubMembership(),
			"github_repository":              resourceGithubRepository(),
			"github_repository_deploy_key":   resourceGithubRepositoryDeployKey(),
			"github_repository_webhook":      resourceGithubRepositoryWebhook(),
			"github_organization_webhook":    resourceGithubOrganizationWebhook(),
			"github_repository_collaborator": resourceGithubRepositoryCollaborator(),
			"github_issue_label":             resourceGithubIssueLabel(),
			"github_branch_protection":       resourceGithubBranchProtection(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"github_user":      dataSourceGithubUser(),
			"github_team":      dataSourceGithubTeam(),
			"github_ip_ranges": dataSourceGithubIpRanges(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"token": "The OAuth token used to connect to GitHub.",

		"owner": "The GitHub owner name to manage.",

		"base_url": "The GitHub Base API URL",
	}
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			Token:   d.Get("token").(string),
			Owner:   d.Get("owner").(string),
			BaseURL: d.Get("base_url").(string),
		}

		meta, err := config.Client()
		if err != nil {
			return nil, err
		}

		meta.(*Owner).StopContext = p.StopContext()

		return meta, nil
	}
}
