package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type githubActionsOrganizationWorkflowPermissionsErrorResponse struct {
	Message          string `json:"message"`
	Errors           string `json:"errors"`
	DocumentationURL string `json:"documentation_url"`
	Status           string `json:"status"`
}

func resourceGithubActionsOrganizationWorkflowPermissions() *schema.Resource {
	return &schema.Resource{
		Description:   "This resource allows you to manage GitHub Actions workflow permissions for a GitHub Organization account. This controls the default permissions granted to the GITHUB_TOKEN when running workflows and whether GitHub Actions can approve pull request reviews.\n\nYou must have organization admin access to use this resource.",
		CreateContext: resourceGithubActionsOrganizationWorkflowPermissionsCreate,
		ReadContext:   resourceGithubActionsOrganizationWorkflowPermissionsRead,
		UpdateContext: resourceGithubActionsOrganizationWorkflowPermissionsUpdate,
		DeleteContext: resourceGithubActionsOrganizationWorkflowPermissionsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"organization_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the Organization.",
			},
			"default_workflow_permissions": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "read",
				Description:      "The default workflow permissions granted to the GITHUB_TOKEN when running workflows in any repository in the organization. Can be 'read' or 'write'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"read", "write"}, false)),
			},
			"can_approve_pull_request_reviews": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether GitHub Actions can approve pull request reviews in any repository in the organization.",
			},
		},
	}
}

func handleEditWorkflowPermissionsError(ctx context.Context, err error, resp *github.Response) diag.Diagnostics {
	var ghErr *github.ErrorResponse
	if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusConflict {
		tflog.Info(ctx, "Detected conflict with workflow permissions", map[string]any{
			"status_code": ghErr.Response.StatusCode,
		})

		errorResponse := &githubActionsOrganizationWorkflowPermissionsErrorResponse{}
		if resp != nil && resp.Body != nil {
			data, readError := io.ReadAll(resp.Body)
			if readError != nil {
				tflog.Error(ctx, "Failed to read workflow permissions conflict response", map[string]any{
					"error": readError.Error(),
				})
				return diag.FromErr(readError)
			}

			if len(data) > 0 {
				if unmarshalError := json.Unmarshal(data, errorResponse); unmarshalError != nil {
					tflog.Error(ctx, "Failed to unmarshal workflow permissions conflict response", map[string]any{
						"error": unmarshalError.Error(),
					})
					return diag.FromErr(unmarshalError)
				}
			}
		}

		return diag.FromErr(fmt.Errorf("you are trying to modify a value restricted by the Enterprise's settings.\n Message: %s\n Errors: %s\n Documentation URL: %s\n Status: %s\nerr: %w", errorResponse.Message, errorResponse.Errors, errorResponse.DocumentationURL, errorResponse.Status, err))
	}

	return diag.FromErr(err)
}

func resourceGithubActionsOrganizationWorkflowPermissionsCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	organizationSlug := d.Get("organization_slug").(string)
	defaultPermissions := d.Get("default_workflow_permissions").(string)
	canApprovePRReviews := d.Get("can_approve_pull_request_reviews").(bool)

	ctx = tflog.SetField(ctx, "organization_slug", organizationSlug)
	tflog.Info(ctx, "Creating workflow permissions")

	workflowPerms := github.DefaultWorkflowPermissionOrganization{
		DefaultWorkflowPermissions:   github.Ptr(defaultPermissions),
		CanApprovePullRequestReviews: github.Ptr(canApprovePRReviews),
	}

	tflog.Debug(ctx, "Calling GitHub API to create workflow permissions", map[string]any{
		"default_workflow_permissions":     defaultPermissions,
		"can_approve_pull_request_reviews": canApprovePRReviews,
	})
	_, resp, err := client.Actions.UpdateDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug, workflowPerms)
	if err != nil {
		return handleEditWorkflowPermissionsError(ctx, err, resp)
	}

	d.SetId(organizationSlug)

	tflog.Trace(ctx, "Created workflow permissions successfully")

	return resourceGithubActionsOrganizationWorkflowPermissionsRead(ctx, d, m)
}

func resourceGithubActionsOrganizationWorkflowPermissionsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	organizationSlug := d.Id()

	ctx = tflog.SetField(ctx, "id", d.Id())
	ctx = tflog.SetField(ctx, "organization_slug", organizationSlug)
	tflog.Info(ctx, "Reading workflow permissions")

	workflowPerms, _, err := client.Actions.GetDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Retrieved workflow permissions from API", map[string]any{
		"default_workflow_permissions":     workflowPerms.DefaultWorkflowPermissions,
		"can_approve_pull_request_reviews": workflowPerms.CanApprovePullRequestReviews,
	})

	if err := d.Set("organization_slug", organizationSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_workflow_permissions", workflowPerms.DefaultWorkflowPermissions); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("can_approve_pull_request_reviews", workflowPerms.CanApprovePullRequestReviews); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Read workflow permissions successfully")

	return nil
}

func resourceGithubActionsOrganizationWorkflowPermissionsUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	organizationSlug := d.Id()
	defaultPermissions := d.Get("default_workflow_permissions").(string)
	canApprovePRReviews := d.Get("can_approve_pull_request_reviews").(bool)

	ctx = tflog.SetField(ctx, "id", d.Id())
	ctx = tflog.SetField(ctx, "organization_slug", organizationSlug)
	tflog.Info(ctx, "Updating workflow permissions")

	workflowPerms := github.DefaultWorkflowPermissionOrganization{
		DefaultWorkflowPermissions:   github.Ptr(defaultPermissions),
		CanApprovePullRequestReviews: github.Ptr(canApprovePRReviews),
	}

	tflog.Debug(ctx, "Calling GitHub API to update workflow permissions", map[string]any{
		"default_workflow_permissions":     defaultPermissions,
		"can_approve_pull_request_reviews": canApprovePRReviews,
	})
	_, resp, err := client.Actions.UpdateDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug, workflowPerms)
	if err != nil {
		return handleEditWorkflowPermissionsError(ctx, err, resp)
	}

	d.SetId(organizationSlug)

	tflog.Trace(ctx, "Updated workflow permissions successfully")

	return resourceGithubActionsOrganizationWorkflowPermissionsRead(ctx, d, m)
}

func resourceGithubActionsOrganizationWorkflowPermissionsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	organizationSlug := d.Id()

	ctx = tflog.SetField(ctx, "id", d.Id())
	ctx = tflog.SetField(ctx, "organization_slug", organizationSlug)
	tflog.Info(ctx, "Deleting organization workflow permissions (resetting to defaults)")

	// Reset to safe defaults
	workflowPerms := github.DefaultWorkflowPermissionOrganization{
		DefaultWorkflowPermissions:   github.Ptr("read"),
		CanApprovePullRequestReviews: github.Ptr(false),
	}

	tflog.Debug(ctx, "Calling GitHub API to reset workflow permissions", map[string]any{
		"workflow_permissions": workflowPerms,
	})

	_, resp, err := client.Actions.UpdateDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug, workflowPerms)
	if err != nil {
		return handleEditWorkflowPermissionsError(ctx, err, resp)
	}

	tflog.Trace(ctx, "Deleted workflow permissions successfully")

	return nil
}
