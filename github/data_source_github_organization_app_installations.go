package github

import (
	"context"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationAppInstallations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationAppInstallationsRead,
		Description: "Use this data source to retrieve all GitHub App installations of the organization.",

		Schema: map[string]*schema.Schema{
			"installations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of GitHub App installations in the organization.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the GitHub App installation.",
						},
						"slug": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL-friendly name of the GitHub App.",
						},
						"app_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the GitHub App.",
						},
						"repository_selection": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether the installation has access to all repositories or only selected ones. Possible values are 'all' or 'selected'.",
						},
						"permissions": {
							Type:        schema.TypeMap,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The permissions granted to the GitHub App installation.",
						},
						"events": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The list of events the GitHub App installation subscribes to.",
						},
						"html_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to the GitHub App installation's settings page.",
						},
						"client_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The OAuth client ID of the GitHub App.",
						},
						"target_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the account the GitHub App is installed on.",
						},
						"target_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of account the GitHub App is installed on. Possible values are 'Organization' or 'User'.",
						},
						"suspended": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the GitHub App installation is currently suspended.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationAppInstallationsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	owner := meta.(*Owner).name

	client := meta.(*Owner).v3client

	options := &github.ListOptions{
		PerPage: maxPerPage,
	}

	results := make([]map[string]interface{}, 0)
	for {
		appInstallations, resp, err := client.Organizations.ListInstallations(ctx, owner, options)
		if err != nil {
			return diag.FromErr(err)
		}

		results = append(results, flattenGitHubAppInstallations(appInstallations.Installations)...)
		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	d.SetId(owner)
	err := d.Set("installations", results)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func flattenGitHubAppInstallations(orgAppInstallations []*github.Installation) []map[string]interface{} {
	results := make([]map[string]interface{}, 0)

	if orgAppInstallations == nil {
		return results
	}

	for _, appInstallation := range orgAppInstallations {
		result := make(map[string]interface{})

		result["id"] = appInstallation.GetID()
		result["slug"] = appInstallation.GetAppSlug()
		result["app_id"] = appInstallation.GetAppID()
		result["repository_selection"] = appInstallation.GetRepositorySelection()
		result["html_url"] = appInstallation.GetHTMLURL()
		result["client_id"] = appInstallation.GetClientID()
		result["target_id"] = appInstallation.GetTargetID()
		result["target_type"] = appInstallation.GetTargetType()
		result["suspended"] = !appInstallation.GetSuspendedAt().IsZero()
		if appInstallation.Events != nil {
			result["events"] = appInstallation.Events
		} else {
			result["events"] = []string{}
		}

		result["permissions"] = flattenInstallationPermissions(appInstallation.Permissions)

		results = append(results, result)
	}

	return results
}

func flattenInstallationPermissions(perms *github.InstallationPermissions) map[string]string {
	permissions := make(map[string]string)
	if perms == nil {
		return permissions
	}

	if v := perms.GetActions(); v != "" {
		permissions["actions"] = v
	}
	if v := perms.GetAdministration(); v != "" {
		permissions["administration"] = v
	}
	if v := perms.GetChecks(); v != "" {
		permissions["checks"] = v
	}
	if v := perms.GetContents(); v != "" {
		permissions["contents"] = v
	}
	if v := perms.GetDeployments(); v != "" {
		permissions["deployments"] = v
	}
	if v := perms.GetEnvironments(); v != "" {
		permissions["environments"] = v
	}
	if v := perms.GetIssues(); v != "" {
		permissions["issues"] = v
	}
	if v := perms.GetMetadata(); v != "" {
		permissions["metadata"] = v
	}
	if v := perms.GetMembers(); v != "" {
		permissions["members"] = v
	}
	if v := perms.GetOrganizationAdministration(); v != "" {
		permissions["organization_administration"] = v
	}
	if v := perms.GetOrganizationHooks(); v != "" {
		permissions["organization_hooks"] = v
	}
	if v := perms.GetOrganizationPlan(); v != "" {
		permissions["organization_plan"] = v
	}
	if v := perms.GetOrganizationProjects(); v != "" {
		permissions["organization_projects"] = v
	}
	if v := perms.GetOrganizationSecrets(); v != "" {
		permissions["organization_secrets"] = v
	}
	if v := perms.GetOrganizationSelfHostedRunners(); v != "" {
		permissions["organization_self_hosted_runners"] = v
	}
	if v := perms.GetOrganizationUserBlocking(); v != "" {
		permissions["organization_user_blocking"] = v
	}
	if v := perms.GetPackages(); v != "" {
		permissions["packages"] = v
	}
	if v := perms.GetPages(); v != "" {
		permissions["pages"] = v
	}
	if v := perms.GetPullRequests(); v != "" {
		permissions["pull_requests"] = v
	}
	if v := perms.GetRepositoryHooks(); v != "" {
		permissions["repository_hooks"] = v
	}
	if v := perms.GetRepositoryProjects(); v != "" {
		permissions["repository_projects"] = v
	}
	if v := perms.GetRepositoryPreReceiveHooks(); v != "" {
		permissions["repository_pre_receive_hooks"] = v
	}
	if v := perms.GetSecrets(); v != "" {
		permissions["secrets"] = v
	}
	if v := perms.GetSecretScanningAlerts(); v != "" {
		permissions["secret_scanning_alerts"] = v
	}
	if v := perms.GetSecurityEvents(); v != "" {
		permissions["security_events"] = v
	}
	if v := perms.GetSingleFile(); v != "" {
		permissions["single_file"] = v
	}
	if v := perms.GetStatuses(); v != "" {
		permissions["statuses"] = v
	}
	if v := perms.GetTeamDiscussions(); v != "" {
		permissions["team_discussions"] = v
	}
	if v := perms.GetVulnerabilityAlerts(); v != "" {
		permissions["vulnerability_alerts"] = v
	}
	if v := perms.GetWorkflows(); v != "" {
		permissions["workflows"] = v
	}

	return permissions
}
