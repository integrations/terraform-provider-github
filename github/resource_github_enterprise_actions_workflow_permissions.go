package github

import (
	"context"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseActionsWorkflowPermissions() *schema.Resource {
	return &schema.Resource{
		Description:   "GitHub Enterprise Actions Workflow Permissions management.",
		CreateContext: resourceGithubEnterpriseActionsWorkflowPermissionsCreateOrUpdate,
		ReadContext:   resourceGithubEnterpriseActionsWorkflowPermissionsRead,
		UpdateContext: resourceGithubEnterpriseActionsWorkflowPermissionsCreateOrUpdate,
		DeleteContext: resourceGithubEnterpriseActionsWorkflowPermissionsDelete,
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
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "read",
				Description:      "The default workflow permissions granted to the GITHUB_TOKEN when running workflows. Can be 'read' or 'write'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"read", "write"}, false)),
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

func resourceGithubEnterpriseActionsWorkflowPermissionsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)
	d.SetId(enterpriseSlug)

	workflowPerms := github.DefaultWorkflowPermissionEnterprise{}

	if v, ok := d.GetOk("default_workflow_permissions"); ok {
		workflowPerms.DefaultWorkflowPermissions = new(v.(string))
	}

	if v, ok := d.GetOk("can_approve_pull_request_reviews"); ok {
		workflowPerms.CanApprovePullRequestReviews = new(v.(bool))
	}
	tflog.Debug(ctx, "Updating workflow permissions for enterprise", map[string]any{"enterprise": enterpriseSlug})
	_, _, err := client.Actions.UpdateDefaultWorkflowPermissionsInEnterprise(ctx, enterpriseSlug, workflowPerms)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubEnterpriseActionsWorkflowPermissionsRead(ctx, d, meta)
}

func resourceGithubEnterpriseActionsWorkflowPermissionsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Id()
	tflog.Debug(ctx, "Reading workflow permissions for enterprise", map[string]any{"enterprise": enterpriseSlug})

	workflowPerms, _, err := client.Actions.GetDefaultWorkflowPermissionsInEnterprise(ctx, enterpriseSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_workflow_permissions", workflowPerms.DefaultWorkflowPermissions); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("can_approve_pull_request_reviews", workflowPerms.CanApprovePullRequestReviews); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseActionsWorkflowPermissionsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Id()
	tflog.Debug(ctx, "Resetting workflow permissions to defaults for enterprise", map[string]any{"enterprise": enterpriseSlug})

	// Reset to safe defaults
	workflowPerms := github.DefaultWorkflowPermissionEnterprise{
		DefaultWorkflowPermissions:   new("read"),
		CanApprovePullRequestReviews: new(false),
	}

	_, _, err := client.Actions.UpdateDefaultWorkflowPermissionsInEnterprise(ctx, enterpriseSlug, workflowPerms)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
