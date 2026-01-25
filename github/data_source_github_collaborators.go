package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubCollaborators() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the collaborators for a repository.",
		Read:        dataSourceGithubCollaboratorsRead,

		Schema: map[string]*schema.Schema{
			"owner": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The organization that owns the repository.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository.",
			},
			"affiliation": {
				Type:        schema.TypeString,
				Description: "Filter collaborators returned by their affiliation. Can be one of: outside, direct, all. Defaults to all.",
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{
					"all",
					"direct",
					"outside",
				}, false), "affiliation"),
				Optional: true,
				Default:  "all",
			},
			"permission": {
				Type:        schema.TypeString,
				Description: "Filter collaborators returned by their permission. Can be one of: pull, triage, push, maintain, admin.",
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{
					"pull",
					"triage",
					"push",
					"maintain",
					"admin",
				}, false), "permission"),
				Optional: true,
				Default:  "",
			},
			"collaborator": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of collaborators for the repository.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"login": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The collaborator's login.",
						},
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the collaborator.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for the collaborator.",
						},
						"html_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub HTML URL for the collaborator.",
						},
						"followers_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for the collaborator's followers.",
						},
						"following_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for those following the collaborator.",
						},
						"gists_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for the collaborator's gists.",
						},
						"starred_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for the collaborator's starred repositories.",
						},
						"subscriptions_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for the collaborator's subscribed repositories.",
						},
						"organizations_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for the collaborator's organizations.",
						},
						"repos_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for the collaborator's repositories.",
						},
						"events_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for the collaborator's events.",
						},
						"received_events_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GitHub API URL for the collaborator's received events.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the collaborator (e.g., user).",
						},
						"site_admin": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the user is a GitHub admin.",
						},
						"permission": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The permission of the collaborator.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubCollaboratorsRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner := d.Get("owner").(string)
	repo := d.Get("repository").(string)
	affiliation := d.Get("affiliation").(string)
	permission := d.Get("permission").(string)

	options := &github.ListCollaboratorsOptions{
		Affiliation: affiliation,
		Permission:  permission,
		ListOptions: github.ListOptions{
			PerPage: maxPerPage,
		},
	}

	if len(permission) == 0 {
		d.SetId(fmt.Sprintf("%s/%s/%s", owner, repo, affiliation))
	} else {
		d.SetId(fmt.Sprintf("%s/%s/%s/%s", owner, repo, affiliation, permission))
	}
	err := d.Set("owner", owner)
	if err != nil {
		return err
	}
	err = d.Set("repository", repo)
	if err != nil {
		return err
	}
	err = d.Set("affiliation", affiliation)
	if err != nil {
		return err
	}
	err = d.Set("permission", permission)
	if err != nil {
		return err
	}

	totalCollaborators := make([]any, 0)
	for {
		collaborators, resp, err := client.Repositories.ListCollaborators(ctx, owner, repo, options)
		if err != nil {
			return err
		}

		result := flattenGitHubCollaborators(collaborators)

		totalCollaborators = append(totalCollaborators, result...)

		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	err = d.Set("collaborator", totalCollaborators)
	if err != nil {
		return err
	}

	return nil
}

func flattenGitHubCollaborators(collaborators []*github.User) []any {
	if collaborators == nil {
		return make([]any, 0)
	}

	results := make([]any, 0)

	for _, c := range collaborators {
		result := make(map[string]any)

		result["login"] = c.GetLogin()
		result["id"] = c.GetID()
		result["url"] = c.GetURL()
		result["html_url"] = c.GetHTMLURL()
		result["following_url"] = c.GetFollowingURL()
		result["followers_url"] = c.GetFollowersURL()
		result["gists_url"] = c.GetGistsURL()
		result["starred_url"] = c.GetStarredURL()
		result["subscriptions_url"] = c.GetSubscriptionsURL()
		result["organizations_url"] = c.GetOrganizationsURL()
		result["repos_url"] = c.GetReposURL()
		result["events_url"] = c.GetEventsURL()
		result["received_events_url"] = c.GetReceivedEventsURL()
		result["type"] = c.GetType()
		result["site_admin"] = c.GetSiteAdmin()
		result["permission"] = getPermission(c.GetRoleName())

		results = append(results, result)
	}

	return results
}
