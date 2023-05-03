package github

import (
	"context"
	"github.com/google/go-github/v53/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func dataSourceGithubActionsRepositoryRequiredWorkflows() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsRepositoryRequiredWorkflowsRead,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository.",
			},
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

func dataSourceGithubActionsRepositoryRequiredWorkflowsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)

	options := github.ListOptions{
		PerPage: 100,
	}

	var repoRequiredWorkflows []map[string]string
	for {
		workflows, resp, err := client.Actions.ListRepoRequiredWorkflows(ctx, owner, repo, &options)
		if err != nil {
			return err
		}
		for _, workflow := range workflows.RequiredWorkflows {
			newWorkflow := map[string]string{
				"required_workflow_id": strconv.FormatInt(workflow.GetID(), 10),
				"name":                 workflow.GetName(),
				"created_at":           workflow.GetCreatedAt().String(),
				"updated_at":           workflow.GetUpdatedAt().String(),
				"state":                workflow.GetState(),
				"badge_url":            workflow.GetBadgeURL(),
				"html_url":             workflow.GetHTMLURL(),
				"node_id":              workflow.GetNodeID(),
				"source_repository":    workflow.GetSourceRepository().GetName(),
				"path":                 workflow.GetPath(),
				"url":                  workflow.GetURL(),
			}
			repoRequiredWorkflows = append(repoRequiredWorkflows, newWorkflow)
		}
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	d.SetId(repo)

	d.Set("required_workflows", repoRequiredWorkflows)
	d.Set("total_count", len(repoRequiredWorkflows))

	return nil
}
