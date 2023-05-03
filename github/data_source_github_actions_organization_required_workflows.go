package github

import (
	"context"
	"github.com/google/go-github/v53/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func dataSourceGithubActionsOrganizationRequiredWorkflows() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsOrganizationRequiredWorkflowsRead,
		Schema: map[string]*schema.Schema{
			"required_workflows": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"required_workflow_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"repository": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scope": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"selected_repositories_url": {
							Type:     schema.TypeString,
							Computed: true,
							Required: false,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"total_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubActionsOrganizationRequiredWorkflowsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	owner := meta.(*Owner).name

	options := github.ListOptions{
		PerPage: 100,
	}

	var orgRequiredWorkflows []map[string]string
	for {
		workflows, resp, err := client.Actions.ListOrgRequiredWorkflows(ctx, owner, &options)
		if err != nil {
			return err
		}
		for _, workflow := range workflows.RequiredWorkflows {
			newWorkflow := map[string]string{
				"required_workflow_id":      strconv.FormatInt(workflow.GetID(), 10),
				"name":                      workflow.GetName(),
				"created_at":                workflow.GetCreatedAt().String(),
				"updated_at":                workflow.GetUpdatedAt().String(),
				"state":                     workflow.GetState(),
				"ref":                       workflow.GetRef(),
				"repository":                workflow.GetRepository().GetName(),
				"scope":                     workflow.GetScope(),
				"selected_repositories_url": workflow.GetSelectedRepositoriesURL(),
				"path":                      workflow.GetPath(),
			}
			orgRequiredWorkflows = append(orgRequiredWorkflows, newWorkflow)
		}
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	d.SetId(owner)

	d.Set("required_workflows", orgRequiredWorkflows)
	d.Set("total_count", len(orgRequiredWorkflows))

	return nil
}
