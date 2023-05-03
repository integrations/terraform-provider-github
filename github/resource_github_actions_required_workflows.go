package github

import (
	"context"
	"github.com/google/go-github/v53/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"strconv"
)

func resourceGithubActionsRequiredWorkflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsRequiredWorkflowCreate,
		Read:   resourceGithubActionsRequiredWorkflowRead,
		Update: resourceGithubActionsRequiredWorkflowUpdate,
		Delete: resourceGithubActionsRequiredWorkflowDelete,

		Schema: map[string]*schema.Schema{
			"repository_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the repository that contains the workflow file.",
			},
			"required_workflow_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The unique identifier of the required workflow.",
			},
			"scope": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Enable the required workflow for all repositories or selected repositories in the organization. Can be one of: 'selected', 'all'",
				ValidateFunc: validation.StringInSlice([]string{"all", "selected"}, false),
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Optional:    true,
				Description: "A list of repository IDs where you want to enable the required workflow. You can only provide a list of repository ids when the scope is set to selected.",
			},
			"workflow_file_path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The path of the workflow file to be configured as a required workflow.",
			},
		},
	}
}

func resourceGithubActionsRequiredWorkflowCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoId := int64(d.Get("repository_id").(int))
	scope := d.Get("scope").(string)
	workflowFilePath := d.Get("workflow_file_path").(string)

	ctx := context.Background()
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	requiredWorkflowOptions := github.CreateUpdateRequiredWorkflowOptions{
		WorkflowFilePath: &workflowFilePath,
		RepositoryID:     &repoId,
		Scope:            &scope,
	}

	// Only set SelectedRepositoryIDs if the scope is 'selected'
	if scope == "selected" {
		selectedRepoIdsValue := d.Get("selected_repository_ids")
		var selectedRepoIds github.SelectedRepoIDs
		ids := selectedRepoIdsValue.(*schema.Set).List()
		for _, id := range ids {
			selectedRepoIds = append(selectedRepoIds, int64(id.(int)))
		}

		requiredWorkflowOptions.SelectedRepositoryIDs = &selectedRepoIds
	}

	create, _, err := client.Actions.CreateRequiredWorkflow(
		ctx,
		owner,
		&requiredWorkflowOptions,
	)

	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(create.GetID(), 10))
	d.Set("required_workflow_id", create.GetID())

	return resourceGithubActionsRequiredWorkflowRead(d, meta)
}

func resourceGithubActionsRequiredWorkflowUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoId := int64(d.Get("repository_id").(int))
	requiredWorkflowId := int64(d.Get("required_workflow_id").(int))

	scope := d.Get("scope").(string)
	workflowFilePath := d.Get("workflow_file_path").(string)

	ctx := context.Background()
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	requiredWorkflowOptions := github.CreateUpdateRequiredWorkflowOptions{
		WorkflowFilePath: &workflowFilePath,
		RepositoryID:     &repoId,
		Scope:            &scope,
	}

	// Only set SelectedRepositoryIDs if the scope is 'selected'
	if scope == "selected" {
		selectedRepoIdsValue := d.Get("selected_repository_ids")
		var selectedRepoIds github.SelectedRepoIDs
		ids := selectedRepoIdsValue.(*schema.Set).List()
		for _, id := range ids {
			selectedRepoIds = append(selectedRepoIds, int64(id.(int)))
		}

		requiredWorkflowOptions.SelectedRepositoryIDs = &selectedRepoIds
	}

	_, _, err := client.Actions.UpdateRequiredWorkflow(
		ctx,
		owner,
		requiredWorkflowId,
		&requiredWorkflowOptions,
	)

	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(requiredWorkflowId, 10))
	return resourceGithubActionsRequiredWorkflowRead(d, meta)
}

func resourceGithubActionsRequiredWorkflowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	requiredWorkflowId := int64(d.Get("required_workflow_id").(int))

	_, _, err := client.Actions.GetRequiredWorkflowByID(ctx, owner, requiredWorkflowId)
	if err != nil {
		return err
	}

	return nil
}

func resourceGithubActionsRequiredWorkflowDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	requiredWorkflowId := int64(d.Get("required_workflow_id").(int))

	_, err := client.Actions.DeleteRequiredWorkflow(ctx,
		owner,
		requiredWorkflowId,
	)
	log.Printf("[DEBUG] Deleting workflow: %d", requiredWorkflowId)
	if err != nil {
		return err
	}

	return nil
}
