package github

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func dataSourceGithubActionsRequiredWorkflow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsRequiredWorkflowRead,
		Schema: map[string]*schema.Schema{
			"required_workflow_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The unique identifier of the required workflow.",
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
	}
}

func dataSourceGithubActionsRequiredWorkflowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	owner := meta.(*Owner).name
	requiredWorkflowId := int64(d.Get("required_workflow_id").(int))

	requiredWorkflow, _, err := client.Actions.GetRequiredWorkflowByID(
		ctx,
		owner,
		requiredWorkflowId,
	)

	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(requiredWorkflow.GetID(), 10))
	d.Set("required_workflow_id", requiredWorkflow.GetID())
	d.Set("name", requiredWorkflow.GetName())
	d.Set("ref", requiredWorkflow.GetRef())
	d.Set("path", requiredWorkflow.GetPath())
	d.Set("created_at", requiredWorkflow.GetCreatedAt().String())
	d.Set("repository", requiredWorkflow.GetRepository().Name)
	d.Set("scope", requiredWorkflow.GetScope())
	d.Set("selected_repositories_url", requiredWorkflow.GetSelectedRepositoriesURL())
	d.Set("state", requiredWorkflow.GetState())
	d.Set("updated_at", requiredWorkflow.GetUpdatedAt().String())

	return nil
}
