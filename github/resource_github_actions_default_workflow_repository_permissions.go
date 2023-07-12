package github

import (
	"context"

	"github.com/google/go-github/v53/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubActionsDefaultWorkflowRepositoryPermissions() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsDefaultWorkflowRepositoryPermissionsCreateOrUpdate,
		Read:   resourceGithubActionsDefaultWorkflowRepositoryPermissionsRead,
		Update: resourceGithubActionsDefaultWorkflowRepositoryPermissionsCreateOrUpdate,
		Delete: resourceGithubActionsDefaultWorkflowRepositoryPermissionsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"default_workflow_permissions": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "read",
				Description:  "The default workflow permissions granted to the GITHUB_TOKEN when running workflows. Can be one of: 'read' or 'write'.",
				ValidateFunc: validation.StringInSlice([]string{"read", "write"}, false),
			},
			"can_approve_pull_request_reviews": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.",
			},
			"repository": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The GitHub repository.",
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
		},
	}
}

func resourceGithubActionsDefaultWorkflowRepositoryPermissionsCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	ctx := context.Background()
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	defaultWorkflowPermissions := d.Get("default_workflow_permissions").(string)
	canApprovePullRequestReviews := d.Get("can_approve_pull_request_reviews").(bool)

	repoDefaultWorkflowPermissions := github.DefaultWorkflowPermissionsRepository{
		DefaultWorkflowPermissions:   &defaultWorkflowPermissions,
		CanApprovePullRequestReviews: &canApprovePullRequestReviews,
	}

	_, _, err := client.Repositories.EditDefaultWorkflowPermissions(ctx,
		owner,
		repoName,
		repoDefaultWorkflowPermissions,
	)
	if err != nil {
		return err
	}

	d.SetId(repoName)
	return resourceGithubActionsDefaultWorkflowRepositoryPermissionsRead(d, meta)
}

func resourceGithubActionsDefaultWorkflowRepositoryPermissionsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Id()
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	defaultWorkflowPermissions, _, err := client.Repositories.GetDefaultWorkflowPermissions(ctx, owner, repoName)
	if err != nil {
		return err
	}

	d.Set("default_workflow_permissions", defaultWorkflowPermissions.GetDefaultWorkflowPermissions())
	d.Set("can_approve_pull_request_reviews", defaultWorkflowPermissions.GetCanApprovePullRequestReviews())
	d.Set("repository", repoName)

	return nil
}

func resourceGithubActionsDefaultWorkflowRepositoryPermissionsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	// Reset the repo to "default" settings
	repoActionPermissions := github.DefaultWorkflowPermissionsRepository{
		DefaultWorkflowPermissions:   github.String("read"),
		CanApprovePullRequestReviews: github.Bool(false),
	}

	_, _, err := client.Repositories.EditDefaultWorkflowPermissions(ctx,
		owner,
		repoName,
		repoActionPermissions,
	)
	if err != nil {
		return err
	}

	return nil
}
