package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type GithubActionsOrganizationWorkflowPermissionsErrorResponse struct {
	Message          string `json:"message"`
	Errors           string `json:"errors"`
	DocumentationURL string `json:"documentation_url"`
	Status           string `json:"status"`
}

func resourceGithubActionsOrganizationWorkflowPermissions() *schema.Resource {
	return &schema.Resource{
		Description:   "This resource allows you to manage GitHub Actions workflow permissions for a GitHub Organization account. This controls the default permissions granted to the GITHUB_TOKEN when running workflows and whether GitHub Actions can approve pull request reviews.\n\nYou must have organization admin access to use this resource.",
		CreateContext: resourceGithubActionsOrganizationWorkflowPermissionsCreateOrUpdate,
		ReadContext:   resourceGithubActionsOrganizationWorkflowPermissionsRead,
		UpdateContext: resourceGithubActionsOrganizationWorkflowPermissionsCreateOrUpdate,
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "read",
				Description:  "The default workflow permissions granted to the GITHUB_TOKEN when running workflows in any repository in the organization. Can be 'read' or 'write'.",
				ValidateFunc: validation.StringInSlice([]string{"read", "write"}, false),
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
	if errors.As(err, &ghErr) {
		if ghErr.Response.StatusCode == http.StatusConflict {
			tflog.Info(ctx, "Detected conflict with workflow permissions", map[string]any{
				"status_code": ghErr.Response.StatusCode,
			})

			errorResponse := &GithubActionsOrganizationWorkflowPermissionsErrorResponse{}
			data, readError := io.ReadAll(resp.Body)
			if readError == nil && data != nil {
				unmarshalError := json.Unmarshal(data, errorResponse)
				if unmarshalError != nil {
					tflog.Error(ctx, "Failed to unmarshal error response", map[string]any{
						"error": unmarshalError.Error(),
					})
					return diag.FromErr(unmarshalError)
				}

				tflog.Debug(ctx, "Parsed workflow permissions conflict error", map[string]any{
					"message":           errorResponse.Message,
					"errors":            errorResponse.Errors,
					"documentation_url": errorResponse.DocumentationURL,
					"status":            errorResponse.Status,
				})
			}
			return diag.FromErr(fmt.Errorf("you are trying to modify a value restricted by the Enterprise's settings.\n Message: %s\n Errors: %s\n Documentation URL: %s\n Status: %s\nerr: %w", errorResponse.Message, errorResponse.Errors, errorResponse.DocumentationURL, errorResponse.Status, err))
		}
	}

	tflog.Trace(ctx, "Returning generic error", map[string]any{
		"error": err.Error(),
	})

	return diag.FromErr(err)
}

func resourceGithubActionsOrganizationWorkflowPermissionsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Trace(ctx, "Entering Create/Update workflow permissions", map[string]any{
		"organization_slug": d.Get("organization_slug").(string),
	})

	client := meta.(*Owner).v3client

	organizationSlug := d.Get("organization_slug").(string)
	d.SetId(organizationSlug)

	if d.IsNewResource() {
		tflog.Info(ctx, "Creating organization workflow permissions", map[string]any{
			"organization_slug": organizationSlug,
		})
	} else {
		tflog.Info(ctx, "Updating organization workflow permissions", map[string]any{
			"organization_slug": organizationSlug,
		})
	}

	workflowPerms := github.DefaultWorkflowPermissionOrganization{}

	if v, ok := d.GetOk("default_workflow_permissions"); ok {
		workflowPerms.DefaultWorkflowPermissions = github.Ptr(v.(string))
	}

	if v, ok := d.GetOk("can_approve_pull_request_reviews"); ok {
		workflowPerms.CanApprovePullRequestReviews = github.Ptr(v.(bool))
	}

	tflog.Debug(ctx, "Calling GitHub API to update workflow permissions", map[string]any{
		"organization_slug":                organizationSlug,
		"default_workflow_permissions":     workflowPerms.DefaultWorkflowPermissions,
		"can_approve_pull_request_reviews": workflowPerms.CanApprovePullRequestReviews,
	})
	_, resp, err := client.Actions.UpdateDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug, workflowPerms)
	if err != nil {
		return handleEditWorkflowPermissionsError(ctx, err, resp)
	}

	tflog.Trace(ctx, "Exiting Create/Update workflow permissions successfully", map[string]any{
		"organization_slug": organizationSlug,
	})
	return nil
}

func resourceGithubActionsOrganizationWorkflowPermissionsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Trace(ctx, "Entering Read workflow permissions", map[string]any{
		"organization_slug": d.Id(),
	})

	client := meta.(*Owner).v3client

	organizationSlug := d.Id()
	tflog.Debug(ctx, "Calling GitHub API to read workflow permissions", map[string]any{
		"organization_slug": organizationSlug,
	})

	workflowPerms, _, err := client.Actions.GetDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Retrieved workflow permissions from API", map[string]any{
		"organization_slug":                organizationSlug,
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

	tflog.Trace(ctx, "Exiting Read workflow permissions successfully", map[string]any{
		"organization_slug": organizationSlug,
	})

	return nil
}

func resourceGithubActionsOrganizationWorkflowPermissionsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Trace(ctx, "Entering Delete workflow permissions", map[string]any{
		"organization_slug": d.Id(),
	})

	client := meta.(*Owner).v3client

	organizationSlug := d.Id()
	tflog.Info(ctx, "Deleting organization workflow permissions (resetting to defaults)", map[string]any{
		"organization_slug": organizationSlug,
	})

	// Reset to safe defaults
	workflowPerms := github.DefaultWorkflowPermissionOrganization{
		DefaultWorkflowPermissions:   github.Ptr("read"),
		CanApprovePullRequestReviews: github.Ptr(false),
	}

	tflog.Debug(ctx, "Using safe default values", map[string]any{
		"default_workflow_permissions":     "read",
		"can_approve_pull_request_reviews": false,
	})

	tflog.Debug(ctx, "Calling GitHub API to reset workflow permissions", map[string]any{
		"organization_slug":    organizationSlug,
		"workflow_permissions": workflowPerms,
	})

	_, resp, err := client.Actions.UpdateDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug, workflowPerms)
	if err != nil {
		return handleEditWorkflowPermissionsError(ctx, err, resp)
	}

	tflog.Trace(ctx, "Exiting Delete workflow permissions successfully", map[string]any{
		"organization_slug": organizationSlug,
	})

	return nil
}
