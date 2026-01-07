package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/go-github/v81/github"
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

func handleEditWorkflowPermissionsError(err error, resp *github.Response) diag.Diagnostics {
	var ghErr *github.ErrorResponse
	if errors.As(err, &ghErr) {
		if ghErr.Response.StatusCode == http.StatusConflict {
			errorResponse := &GithubActionsOrganizationWorkflowPermissionsErrorResponse{}
			data, readError := io.ReadAll(resp.Body)
			if readError == nil && data != nil {
				unmarshalError := json.Unmarshal(data, errorResponse)
				if unmarshalError != nil {
					return diag.FromErr(unmarshalError)
				}
			}
			return diag.FromErr(fmt.Errorf("you are trying to modify a value restricted by the Enterprise's settings.\n Message: %s\n Errors: %s\n Documentation URL: %s\n Status: %s\nerr: %w", errorResponse.Message, errorResponse.Errors, errorResponse.DocumentationURL, errorResponse.Status, err))
		}
	}
	return diag.FromErr(err)
}

func resourceGithubActionsOrganizationWorkflowPermissionsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	organizationSlug := d.Get("organization_slug").(string)
	d.SetId(organizationSlug)

	workflowPerms := github.DefaultWorkflowPermissionOrganization{}

	if v, ok := d.GetOk("default_workflow_permissions"); ok {
		workflowPerms.DefaultWorkflowPermissions = github.String(v.(string))
	}

	if v, ok := d.GetOk("can_approve_pull_request_reviews"); ok {
		workflowPerms.CanApprovePullRequestReviews = github.Bool(v.(bool))
	}

	log.Printf("[DEBUG] Updating workflow permissions for Organization: %s", organizationSlug)
	_, resp, err := client.Actions.UpdateDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug, workflowPerms)
	if err != nil {
		return handleEditWorkflowPermissionsError(err, resp)
	}

	// Calling read is necessary as the Update API returns 204 with Empty Body on success
	return resourceGithubActionsOrganizationWorkflowPermissionsRead(ctx, d, meta)
}

func resourceGithubActionsOrganizationWorkflowPermissionsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	organizationSlug := d.Id()
	log.Printf("[DEBUG] Reading workflow permissions for Organization: %s", organizationSlug)

	workflowPerms, _, err := client.Actions.GetDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("organization_slug", organizationSlug); err != nil {
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

func resourceGithubActionsOrganizationWorkflowPermissionsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	organizationSlug := d.Id()
	log.Printf("[DEBUG] Resetting workflow permissions to defaults for Organization: %s", organizationSlug)

	// Reset to safe defaults
	workflowPerms := github.DefaultWorkflowPermissionOrganization{
		DefaultWorkflowPermissions:   github.Ptr("read"),
		CanApprovePullRequestReviews: github.Ptr(false),
	}

	_, resp, err := client.Actions.UpdateDefaultWorkflowPermissionsInOrganization(ctx, organizationSlug, workflowPerms)
	if err != nil {
		return handleEditWorkflowPermissionsError(err, resp)
	}

	return nil
}
