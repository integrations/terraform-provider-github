package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceGithubCollaborators() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubCollaboratorsRead,

		Schema: map[string]*schema.Schema{
			"owner": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"affiliation": {
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{
					"all",
					"direct",
					"outside",
				}, false),
				Optional: true,
				Default:  "all",
			},
			"collaborator": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"login": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"html_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"followers_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"following_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gists_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"starred_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscriptions_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizations_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"repos_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"events_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"received_events_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_admin": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"permission": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubCollaboratorsRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Organization).v3client
	ctx := context.Background()

	owner := d.Get("owner").(string)
	repo := d.Get("repository").(string)
	affiliation := d.Get("affiliation").(string)

	options := &github.ListCollaboratorsOptions{
		Affiliation: affiliation,
		ListOptions: github.ListOptions{
			PerPage: maxPerPage,
		},
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", owner, repo, affiliation))
	d.Set("owner", owner)
	d.Set("repository", repo)
	d.Set("affiliation", affiliation)

	totalCollaborators := make([]interface{}, 0)
	for {
		collaborators, resp, err := client.Repositories.ListCollaborators(ctx, owner, repo, options)
		if err != nil {
			return err
		}

		result, err := flattenGitHubCollaborators(collaborators)
		if err != nil {
			return fmt.Errorf("unable to flatten GitHub Collaborators (Owner: %q/Repository: %q) : %+v", owner, repo, err)
		}

		totalCollaborators = append(totalCollaborators, result...)

		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	d.Set("collaborator", totalCollaborators)

	return nil
}

func flattenGitHubCollaborators(collaborators []*github.User) ([]interface{}, error) {
	if collaborators == nil {
		return make([]interface{}, 0), nil
	}

	results := make([]interface{}, 0)

	for _, c := range collaborators {
		result := make(map[string]interface{})

		result["login"] = c.Login
		result["id"] = c.ID
		result["url"] = c.URL
		result["html_url"] = c.HTMLURL
		result["following_url"] = c.FollowingURL
		result["followers_url"] = c.FollowersURL
		result["gists_url"] = c.GistsURL
		result["starred_url"] = c.StarredURL
		result["subscriptions_url"] = c.SubscriptionsURL
		result["organizations_url"] = c.OrganizationsURL
		result["repos_url"] = c.ReposURL
		result["events_url"] = c.EventsURL
		result["received_events_url"] = c.ReceivedEventsURL
		result["type"] = c.Type
		result["site_admin"] = c.SiteAdmin

		permissionName, err := getRepoPermission(c.Permissions)
		if err != nil {
			return nil, err
		}

		result["permission"] = permissionName
		results = append(results, result)
	}

	return results, nil
}
