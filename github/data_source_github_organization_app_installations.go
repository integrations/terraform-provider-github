package github

import (
	"context"

	"github.com/google/go-github/v77/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationAppInstallations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationAppInstallationsRead,

		Schema: map[string]*schema.Schema{
			"installations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"app_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"repository_selection": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"permissions": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"events": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"html_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"target_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"suspended": {
							Type:     schema.TypeBool,
							Computed: true,
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
		PerPage: 100,
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
		result["suspended"] = appInstallation.GetSuspendedAt() != nil
		result["events"] = appInstallation.Events

		permissions := make(map[string]string)
		if appInstallation.Permissions != nil {
			if v := appInstallation.Permissions.Actions; v != nil {
				permissions["actions"] = *v
			}
			if v := appInstallation.Permissions.Administration; v != nil {
				permissions["administration"] = *v
			}
			if v := appInstallation.Permissions.Checks; v != nil {
				permissions["checks"] = *v
			}
			if v := appInstallation.Permissions.Contents; v != nil {
				permissions["contents"] = *v
			}
			if v := appInstallation.Permissions.Deployments; v != nil {
				permissions["deployments"] = *v
			}
			if v := appInstallation.Permissions.Environments; v != nil {
				permissions["environments"] = *v
			}
			if v := appInstallation.Permissions.Issues; v != nil {
				permissions["issues"] = *v
			}
			if v := appInstallation.Permissions.Metadata; v != nil {
				permissions["metadata"] = *v
			}
			if v := appInstallation.Permissions.Members; v != nil {
				permissions["members"] = *v
			}
			if v := appInstallation.Permissions.OrganizationAdministration; v != nil {
				permissions["organization_administration"] = *v
			}
			if v := appInstallation.Permissions.OrganizationHooks; v != nil {
				permissions["organization_hooks"] = *v
			}
			if v := appInstallation.Permissions.OrganizationPlan; v != nil {
				permissions["organization_plan"] = *v
			}
			if v := appInstallation.Permissions.OrganizationProjects; v != nil {
				permissions["organization_projects"] = *v
			}
			if v := appInstallation.Permissions.OrganizationSecrets; v != nil {
				permissions["organization_secrets"] = *v
			}
			if v := appInstallation.Permissions.OrganizationSelfHostedRunners; v != nil {
				permissions["organization_self_hosted_runners"] = *v
			}
			if v := appInstallation.Permissions.OrganizationUserBlocking; v != nil {
				permissions["organization_user_blocking"] = *v
			}
			if v := appInstallation.Permissions.Packages; v != nil {
				permissions["packages"] = *v
			}
			if v := appInstallation.Permissions.Pages; v != nil {
				permissions["pages"] = *v
			}
			if v := appInstallation.Permissions.PullRequests; v != nil {
				permissions["pull_requests"] = *v
			}
			if v := appInstallation.Permissions.RepositoryHooks; v != nil {
				permissions["repository_hooks"] = *v
			}
			if v := appInstallation.Permissions.RepositoryProjects; v != nil {
				permissions["repository_projects"] = *v
			}
			if v := appInstallation.Permissions.RepositoryPreReceiveHooks; v != nil {
				permissions["repository_pre_receive_hooks"] = *v
			}
			if v := appInstallation.Permissions.Secrets; v != nil {
				permissions["secrets"] = *v
			}
			if v := appInstallation.Permissions.SecretScanningAlerts; v != nil {
				permissions["secret_scanning_alerts"] = *v
			}
			if v := appInstallation.Permissions.SecurityEvents; v != nil {
				permissions["security_events"] = *v
			}
			if v := appInstallation.Permissions.SingleFile; v != nil {
				permissions["single_file"] = *v
			}
			if v := appInstallation.Permissions.Statuses; v != nil {
				permissions["statuses"] = *v
			}
			if v := appInstallation.Permissions.TeamDiscussions; v != nil {
				permissions["team_discussions"] = *v
			}
			if v := appInstallation.Permissions.VulnerabilityAlerts; v != nil {
				permissions["vulnerability_alerts"] = *v
			}
			if v := appInstallation.Permissions.Workflows; v != nil {
				permissions["workflows"] = *v
			}
		}
		result["permissions"] = permissions

		results = append(results, result)
	}

	return results
}
