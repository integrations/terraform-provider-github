package github

import (
	"strconv"

	"github.com/google/go-github/v50/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubOrganization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"orgname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repositories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"members": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Deprecated: "Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version.",
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	client4 := meta.(*Owner).v4client
	client3 := meta.(*Owner).v3client
	ctx := meta.(*Owner).StopContext

	organization, _, err := client3.Organizations.Get(ctx, name)
	if err != nil {
		return err
	}

	var planName string

	if plan := organization.GetPlan(); plan != nil {
		planName = plan.GetName()
	}

	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100, Page: 1},
	}

	var repoList []string
	var allRepos []*github.Repository

	for {
		repos, resp, err := client3.Repositories.ListByOrg(ctx, name, opts)
		if err != nil {
			return err
		}
		allRepos = append(allRepos, repos...)

		opts.Page = resp.NextPage

		if resp.NextPage == 0 {
			break
		}
	}
	for index := range allRepos {
		repoList = append(repoList, allRepos[index].GetFullName())
	}

	var query struct {
		Organization struct {
			MembersWithRole struct {
				Edges []struct {
					Role githubv4.String
					Node struct {
						Id    githubv4.String
						Login githubv4.String
						Email githubv4.String
					}
				}
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"membersWithRole(first: 100, after: $after)"`
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]interface{}{
		"login": githubv4.String(name),
		"after": (*githubv4.String)(nil),
	}
	var members []string
	var users []map[string]string
	for {
		err := client4.Query(ctx, &query, variables)
		if err != nil {
			return err
		}
		for _, edge := range query.Organization.MembersWithRole.Edges {
			members = append(members, string(edge.Node.Login))
			users = append(users, map[string]string{
				"id":    string(edge.Node.Id),
				"login": string(edge.Node.Login),
				"email": string(edge.Node.Email),
				"role":  string(edge.Role),
			})
		}
		if !query.Organization.MembersWithRole.PageInfo.HasNextPage {
			break
		}
		variables["after"] = githubv4.NewString(query.Organization.MembersWithRole.PageInfo.EndCursor)
	}

	d.SetId(strconv.FormatInt(organization.GetID(), 10))
	d.Set("login", organization.GetLogin())
	d.Set("name", organization.GetName())
	d.Set("orgname", name)
	d.Set("node_id", organization.GetNodeID())
	d.Set("description", organization.GetDescription())
	d.Set("plan", planName)
	d.Set("repositories", repoList)
	d.Set("members", members)
	d.Set("users", users)

	return nil
}
