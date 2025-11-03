package github

import (
	"context"
	"log"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseActionsWorkflowPermissions() *schema.Resource {
	return &schema.Resource{
		Description: "GitHub Enterprise Actions Workflow Permissions management.",
		Create:      resourceGithubEnterpriseActionsWorkflowPermissionsCreateOrUpdate,
		Read:        resourceGithubEnterpriseActionsWorkflowPermissionsRead,
		Update:      resourceGithubEnterpriseActionsWorkflowPermissionsCreateOrUpdate,
		Delete:      resourceGithubEnterpriseActionsWorkflowPermissionsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"default_workflow_permissions": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "read",
				Description:  "The default workflow permissions granted to the GITHUB_TOKEN when running workflows. Can be 'read' or 'write'.",
				ValidateFunc: validation.StringInSlice([]string{"read", "write"}, false),
			},
			"can_approve_pull_request_reviews": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether GitHub Actions can approve pull request reviews.",
			},
		},
	}
}

func resourceGithubEnterpriseActionsWorkflowPermissionsCreateOrUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Get("enterprise_slug").(string)
	d.SetId(enterpriseSlug)

	workflowPerms := github.DefaultWorkflowPermissionEnterprise{}

	if v, ok := d.GetOk("default_workflow_permissions"); ok {
		workflowPerms.DefaultWorkflowPermissions = github.String(v.(string))
	}

	if v, ok := d.GetOk("can_approve_pull_request_reviews"); ok {
		workflowPerms.CanApprovePullRequestReviews = github.Bool(v.(bool))
	}

	log.Printf("[DEBUG] Updating workflow permissions for enterprise: %s", enterpriseSlug)
	_, _, err := client.Actions.EditDefaultWorkflowPermissionsInEnterprise(ctx, enterpriseSlug, workflowPerms)
	if err != nil {
		return err
	}

	return resourceGithubEnterpriseActionsWorkflowPermissionsRead(d, meta)
}

func resourceGithubEnterpriseActionsWorkflowPermissionsRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Id()
	log.Printf("[DEBUG] Reading workflow permissions for enterprise: %s", enterpriseSlug)

	workflowPerms, _, err := client.Actions.GetDefaultWorkflowPermissionsInEnterprise(ctx, enterpriseSlug)
	if err != nil {
		return err
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return err
	}
	if err := d.Set("default_workflow_permissions", workflowPerms.DefaultWorkflowPermissions); err != nil {
		return err
	}
	if err := d.Set("can_approve_pull_request_reviews", workflowPerms.CanApprovePullRequestReviews); err != nil {
		return err
	}

	return nil
}

func resourceGithubEnterpriseActionsWorkflowPermissionsDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Id()
	log.Printf("[DEBUG] Resetting workflow permissions to defaults for enterprise: %s", enterpriseSlug)

	// Reset to safe defaults
	workflowPerms := github.DefaultWorkflowPermissionEnterprise{
		DefaultWorkflowPermissions:   github.String("read"),
		CanApprovePullRequestReviews: github.Bool(false),
	}

	_, _, err := client.Actions.EditDefaultWorkflowPermissionsInEnterprise(ctx, enterpriseSlug, workflowPerms)
	if err != nil {
		return err
	}

	return nil
}
