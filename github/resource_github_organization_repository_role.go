package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationRepositoryRole() *schema.Resource {
	return &schema.Resource{
		Description: "Manage a custom organization repository role.",

		CreateContext: resourceGithubOrganizationRepositoryRoleCreate,
		ReadContext:   resourceGithubOrganizationRepositoryRoleRead,
		UpdateContext: resourceGithubOrganizationRepositoryRoleUpdate,
		DeleteContext: resourceGithubOrganizationRepositoryRoleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description: "The ID of the organization repository role.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"name": {
				Description: "The name of the organization repository role.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the organization repository role.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"base_role": {
				Description:      "The base role for the organization repository role.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validateValueFunc([]string{"read", "triage", "write", "maintain"}),
			},
			"permissions": {
				Description: "The permissions for the organization repository role.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				MinItems:    1,
			},
		},

		CustomizeDiff: resourceGithubOrganizationRepositoryRoleCustomizeDiff,
	}
}

func resourceGithubOrganizationRepositoryRoleCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	permissions := d.Get("permissions").(*schema.Set).List()
	permissionsStr := make([]string, len(permissions))
	for i, v := range permissions {
		permissionsStr[i] = v.(string)
	}

	role, _, err := client.Organizations.CreateCustomRepoRole(ctx, orgName, &github.CreateOrUpdateCustomRepoRoleOptions{
		Name:        github.Ptr(d.Get("name").(string)),
		Description: github.Ptr(d.Get("description").(string)),
		BaseRole:    github.Ptr(d.Get("base_role").(string)),
		Permissions: permissionsStr,
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating GitHub organization repository role (%s/%s): %w", orgName, d.Get("name").(string), err))
	}

	if err = d.Set("role_id", role.GetID()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprint(role.GetID()))
	return nil
}

func resourceGithubOrganizationRepositoryRoleRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	role, _, err := client.Organizations.GetCustomRepoRole(ctx, orgName, roleId)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Warn(ctx, "GitHub organization repository role not found, removing from state", map[string]any{
					"orgName": orgName,
					"roleId":  roleId,
				})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("role_id", role.GetID()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("name", role.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("description", role.Description); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("base_role", role.BaseRole); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("permissions", role.Permissions); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRepositoryRoleUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	permissions := d.Get("permissions").(*schema.Set).List()
	permissionsStr := make([]string, len(permissions))
	for i, v := range permissions {
		permissionsStr[i] = v.(string)
	}

	update := &github.CreateOrUpdateCustomRepoRoleOptions{
		Name:        github.Ptr(d.Get("name").(string)),
		Description: github.Ptr(d.Get("description").(string)),
		BaseRole:    github.Ptr(d.Get("base_role").(string)),
		Permissions: permissionsStr,
	}

	_, _, err = client.Organizations.UpdateCustomRepoRole(ctx, orgName, roleId, update)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating GitHub organization repository role (%s/%s): %w", orgName, d.Get("name").(string), err))
	}

	return nil
}

func resourceGithubOrganizationRepositoryRoleDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Organizations.DeleteCustomRepoRole(ctx, orgName, roleId)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error deleting organization repository role %d: %w", roleId, err))
	}

	return nil
}

// Snapshot of the response to https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/custom-roles?apiVersion=2022-11-28#list-repository-fine-grained-permissions-for-an-organization
// The endpoint isn't covered in the SDK yet.
var validRolePermissions = []string{
	"add_assignee",
	"add_label",
	"bypass_branch_protection",
	"close_discussion",
	"close_issue",
	"close_pull_request",
	"convert_issues_to_discussions",
	"create_discussion_category",
	"create_solo_merge_queue_entry",
	"create_tag",
	"delete_alerts_code_scanning",
	"delete_discussion",
	"delete_discussion_comment",
	"delete_issue",
	"delete_tag",
	"edit_category_on_discussion",
	"edit_discussion_category",
	"edit_discussion_comment",
	"edit_repo_custom_properties_values",
	"edit_repo_metadata",
	"edit_repo_protections",
	"jump_merge_queue",
	"manage_deploy_keys",
	"manage_settings_merge_types",
	"manage_settings_pages",
	"manage_settings_projects",
	"manage_settings_wiki",
	"manage_webhooks",
	"mark_as_duplicate",
	"push_protected_branch",
	"read_code_quality",
	"read_code_scanning",
	"reopen_discussion",
	"reopen_issue",
	"reopen_pull_request",
	"request_pr_review",
	"resolve_dependabot_alerts",
	"resolve_secret_scanning_alerts",
	"set_interaction_limits",
	"set_issue_type",
	"set_milestone",
	"set_social_preview",
	"toggle_discussion_answer",
	"toggle_discussion_comment_minimize",
	"view_dependabot_alerts",
	"view_secret_scanning_alerts",
	"write_code_quality",
	"write_code_scanning",
	"write_repository_actions_environments",
	"write_repository_actions_runners",
	"write_repository_actions_secrets",
	"write_repository_actions_settings",
	"write_repository_actions_variables",
}

func resourceGithubOrganizationRepositoryRoleCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, m any) error {
	tflog.Debug(ctx, "Customizing diff for GitHub organization repository role", map[string]any{"permissionsChanged": d.HasChange("permissions")})
	if d.HasChange("permissions") {
		newPermissions := d.Get("permissions").(*schema.Set).List()
		tflog.Debug(ctx, "Validating permissions values", map[string]any{"newPermissions": newPermissions})
		for _, permission := range newPermissions {
			if !slices.Contains(validRolePermissions, permission.(string)) {
				return fmt.Errorf("invalid permission: %+v", permission)
			}
		}
	}
	return nil
}
