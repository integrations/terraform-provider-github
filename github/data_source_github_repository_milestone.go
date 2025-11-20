package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryMilestone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryMilestoneRead,

		Schema: map[string]*schema.Schema{
			"owner": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"number": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"due_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubRepositoryMilestoneRead(d *schema.ResourceData, meta any) error {
	conn := meta.(*Owner).v3client
	ctx := context.Background()

	owner := d.Get("owner").(string)
	repoName := d.Get("repository").(string)

	number := d.Get("number").(int)
	milestone, _, err := conn.Issues.GetMilestone(ctx, owner, repoName, number)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(milestone.GetID(), 10))
	if err = d.Set("description", milestone.GetDescription()); err != nil {
		return err
	}
	if err = d.Set("due_date", milestone.GetDueOn().Format(layoutISO)); err != nil {
		return err
	}
	if err = d.Set("state", milestone.GetState()); err != nil {
		return err
	}
	if err = d.Set("title", milestone.GetTitle()); err != nil {
		return err
	}

	return nil
}
