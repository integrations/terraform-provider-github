package github

import (
	"context"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationPermissions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsOrganizationPermissionsRead,

		Schema: map[string]*schema.Schema{
			"allowed_actions": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The permissions policy that controls the actions that are allowed to run. Can be one of: 'all', 'local_only', or 'selected'.",
			},
			"enabled_repositories": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: 'all', 'none', or 'selected'.",
			},
			"allowed_actions_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The actions that are allowed in the organization when 'allowed_actions' is 'selected'.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"github_owned_allowed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether GitHub-owned actions are allowed in the organization.",
						},
						"patterns_allowed": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"verified_allowed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether actions in GitHub Marketplace from verified creators are allowed.",
						},
					},
				},
			},
			"enabled_repositories_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of selected repositories that are enabled for GitHub Actions when 'enabled_repositories' is 'selected'.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of repository IDs enabled for GitHub Actions.",
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
					},
				},
			},
			"sha_pinning_required": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether pinning to a specific SHA is required for all actions and reusable workflows in an organization.",
			},
		},
	}
}

func dataSourceGithubActionsOrganizationPermissionsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	actionsPermissions, _, err := client.Actions.GetActionsPermissions(ctx, owner)
	if err != nil {
		return diag.FromErr(err)
	}

	if actionsPermissions.GetAllowedActions() == "selected" {
		actionsAllowed, _, err := client.Actions.GetActionsAllowed(ctx, owner)
		if err != nil {
			return diag.FromErr(err)
		}

		if actionsAllowed != nil {
			if err = d.Set("allowed_actions_config", []any{
				map[string]any{
					"github_owned_allowed": actionsAllowed.GetGithubOwnedAllowed(),
					"patterns_allowed":     actionsAllowed.PatternsAllowed,
					"verified_allowed":     actionsAllowed.GetVerifiedAllowed(),
				},
			}); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if actionsPermissions.GetEnabledRepositories() == "selected" {
		opts := github.ListOptions{PerPage: 10, Page: 1}
		var repoList []int64
		var allRepos []*github.Repository

		for {
			enabledRepos, resp, err := client.Actions.ListEnabledReposInOrg(ctx, owner, &opts)
			if err != nil {
				return diag.FromErr(err)
			}
			allRepos = append(allRepos, enabledRepos.Repositories...)
			opts.Page = resp.NextPage
			if resp.NextPage == 0 {
				break
			}
		}

		for _, repo := range allRepos {
			repoList = append(repoList, *repo.ID)
		}

		if allRepos != nil {
			if err = d.Set("enabled_repositories_config", []any{
				map[string]any{
					"repository_ids": repoList,
				},
			}); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if err = d.Set("allowed_actions", actionsPermissions.GetAllowedActions()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("enabled_repositories", actionsPermissions.GetEnabledRepositories()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("sha_pinning_required", actionsPermissions.GetSHAPinningRequired()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(owner)
	return nil
}
