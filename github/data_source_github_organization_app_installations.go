package github

import (
	"context"
	"encoding/json"

	"github.com/google/go-github/v92/github"
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
						"app_slug": {
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
						"single_file_paths": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The list of single file paths the GitHub App installation has access to.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date the GitHub App installation was created.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date the GitHub App installation was last updated.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationAppInstallationsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	owner := meta.name
	client := meta.v3client

	options := &github.ListOptions{
		PerPage: meta.maxPerPage,
	}

	results := make([]map[string]any, 0)
	for {
		appInstallations, resp, err := client.Organizations.ListInstallations(ctx, owner, options)
		if err != nil {
			return diag.FromErr(err)
		}

		installations, err := flattenGitHubAppInstallations(appInstallations.Installations)
		if err != nil {
			return diag.FromErr(err)
		}
		results = append(results, installations...)
		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	d.SetId(owner)
	if err := d.Set("installations", results); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func flattenGitHubAppInstallations(orgAppInstallations []*github.Installation) ([]map[string]any, error) {
	results := make([]map[string]any, 0)

	if orgAppInstallations == nil {
		return results, nil
	}

	for _, appInstallation := range orgAppInstallations {
		result := make(map[string]any)

		result["id"] = appInstallation.GetID()
		result["app_slug"] = appInstallation.GetAppSlug()
		result["app_id"] = appInstallation.GetAppID()
		result["repository_selection"] = appInstallation.GetRepositorySelection()
		result["client_id"] = appInstallation.GetClientID()
		result["target_id"] = appInstallation.GetTargetID()
		result["target_type"] = appInstallation.GetTargetType()
		result["suspended"] = !appInstallation.GetSuspendedAt().IsZero()
		if appInstallation.Events != nil {
			result["events"] = appInstallation.Events
		} else {
			result["events"] = []string{}
		}

		permissions, err := flattenInstallationPermissions(appInstallation.Permissions)
		if err != nil {
			return nil, err
		}
		result["permissions"] = permissions

		if appInstallation.SingleFilePaths != nil {
			result["single_file_paths"] = appInstallation.SingleFilePaths
		} else {
			result["single_file_paths"] = []string{}
		}

		result["created_at"] = appInstallation.GetCreatedAt().Format("2006-01-02T15:04:05Z")
		result["updated_at"] = appInstallation.GetUpdatedAt().Format("2006-01-02T15:04:05Z")

		results = append(results, result)
	}

	return results, nil
}

// Round-tripping through JSON keeps every permission go-github knows about instead of a hand-picked subset.
func flattenInstallationPermissions(perms *github.InstallationPermissions) (map[string]string, error) {
	permissions := make(map[string]string)
	if perms == nil {
		return permissions, nil
	}

	raw, err := json.Marshal(perms)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(raw, &permissions); err != nil {
		return nil, err
	}

	return permissions, nil
}
