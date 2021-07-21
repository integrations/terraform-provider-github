package github

import (
	"context"
	"log"
	"strconv"

	"github.com/google/go-github/v36/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubOrganization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
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
			"admins": {
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
			},
		},
	}
}

func dataSourceGithubOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	log.Printf("[INFO] Refreshing GitHub Organization: %s", name)

	client := meta.(*Owner).v3client
	ctx := meta.(*Owner).StopContext

	organization, _, err := client.Organizations.Get(ctx, name)
	if err != nil {
		return err
	}

	plan := organization.GetPlan()

	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10, Page: 1},
	}

	var repoList []string
	var allRepos []*github.Repository

	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, name, opts)
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

	memberList, err := listMembersByRole(ctx, client, name, "member")
	if err != nil {
		return err
	}

	adminList, err := listMembersByRole(ctx, client, name, "admin")
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(organization.GetID(), 10))
	d.Set("login", organization.GetLogin())
	d.Set("name", organization.GetName())
	d.Set("description", organization.GetDescription())
	d.Set("plan", plan.Name)
	d.Set("repositories", repoList)
	d.Set("members", memberList)
	d.Set("admins", adminList)

	return nil
}

func listMembersByRole(ctx context.Context, client *github.Client, name string, role string) ([]string, error) {
	membershipOpts := &github.ListMembersOptions{
		Role:        role,
		ListOptions: github.ListOptions{PerPage: 10, Page: 1},
	}

	var memberList []string
	var allMembers []*github.User

	for {
		members, resp, err := client.Organizations.ListMembers(ctx, name, membershipOpts)
		if err != nil {
			return nil, err
		}
		allMembers = append(allMembers, members...)

		membershipOpts.Page = resp.NextPage

		if resp.NextPage == 0 {
			break
		}
	}
	for index := range allMembers {
		memberList = append(memberList, *allMembers[index].Login)
	}

	return memberList, nil
}
