package github

import (
	"context"
	"log"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseSettings() *schema.Resource {
	return &schema.Resource{
		Description: `
GitHub Enterprise Settings management.

Provides a resource to manage various settings for a GitHub Enterprise account.

~> **Note:** The managing account must have enterprise admin permissions.
~> **Note:** This resource requires a GitHub Enterprise account.

`,
		Create: resourceGithubEnterpriseSettingsCreateOrUpdate,
		Read:   resourceGithubEnterpriseSettingsRead,
		Update: resourceGithubEnterpriseSettingsCreateOrUpdate,
		Delete: resourceGithubEnterpriseSettingsDelete,
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

			// Actions Permissions - Available in current go-github
			"actions_enabled_organizations": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be 'all', 'none', or 'selected'.",
			},
			"actions_allowed_actions": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The permissions policy that controls the actions and reusable workflows that are allowed to run. Can be 'all', 'local_only', or 'selected'.",
			},
			"actions_github_owned_allowed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether GitHub-owned actions are allowed.",
			},
			"actions_verified_allowed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether verified Marketplace actions are allowed.",
			},
			"actions_patterns_allowed": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s).",
			},
			"default_workflow_permissions": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The default workflow permissions granted to the GITHUB_TOKEN when running workflows. Can be 'read' or 'write'.",
			},
			"can_approve_pull_request_reviews": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether GitHub Actions can approve pull request reviews.",
			},
		},
	}
}

func resourceGithubEnterpriseSettingsCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Get("enterprise_slug").(string)
	d.SetId(enterpriseSlug)

	hasActionsChange := d.HasChange("actions_enabled_organizations") ||
		d.HasChange("actions_allowed_actions")

	if d.IsNewResource() || hasActionsChange {
		actionsPerms := github.ActionsPermissionsEnterprise{}

		if v, ok := d.GetOk("actions_enabled_organizations"); ok {
			actionsPerms.EnabledOrganizations = github.String(v.(string))
		}

		if v, ok := d.GetOk("actions_allowed_actions"); ok {
			actionsPerms.AllowedActions = github.String(v.(string))
		}

		_, _, err := client.Actions.EditActionsPermissionsInEnterprise(ctx, enterpriseSlug, actionsPerms)
		if err != nil {
			return err
		}
	}

	// Handle allowed actions (only when actions_allowed_actions is "selected")
	hasAllowedActionsChange := d.HasChange("actions_github_owned_allowed") ||
		d.HasChange("actions_verified_allowed") ||
		d.HasChange("actions_patterns_allowed")

	if (d.IsNewResource() || hasAllowedActionsChange) && d.Get("actions_allowed_actions").(string) == "selected" {
		allowedActions := github.ActionsAllowed{}

		if v, ok := d.GetOk("actions_github_owned_allowed"); ok {
			allowedActions.GithubOwnedAllowed = github.Bool(v.(bool))
		}

		if v, ok := d.GetOk("actions_verified_allowed"); ok {
			allowedActions.VerifiedAllowed = github.Bool(v.(bool))
		}

		if v, ok := d.GetOk("actions_patterns_allowed"); ok {
			patterns := expandStringSet(v.(*schema.Set))
			if len(patterns) > 0 {
				allowedActions.PatternsAllowed = patterns
			}
		}

		_, _, err := client.Actions.EditActionsAllowedInEnterprise(ctx, enterpriseSlug, allowedActions)
		if err != nil {
			return err
		}
	}

	hasWorkflowChange := d.HasChange("default_workflow_permissions") ||
		d.HasChange("can_approve_pull_request_reviews")

	if d.IsNewResource() || hasWorkflowChange {
		workflowPerms := github.DefaultWorkflowPermissionEnterprise{}

		if v, ok := d.GetOk("default_workflow_permissions"); ok {
			workflowPerms.DefaultWorkflowPermissions = github.String(v.(string))
		}

		if v, ok := d.GetOk("can_approve_pull_request_reviews"); ok {
			workflowPerms.CanApprovePullRequestReviews = github.Bool(v.(bool))
		}

		_, _, err := client.Actions.EditDefaultWorkflowPermissionsInEnterprise(ctx, enterpriseSlug, workflowPerms)
		if err != nil {
			return err
		}
	}

	return resourceGithubEnterpriseSettingsRead(d, meta)
}

func resourceGithubEnterpriseSettingsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Id()
	log.Printf("[DEBUG] Reading enterprise settings: %s", enterpriseSlug)

	actionsPerms, _, err := client.Actions.GetActionsPermissionsInEnterprise(ctx, enterpriseSlug)
	if err != nil {
		return err
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return err
	}
	if err := d.Set("actions_enabled_organizations", actionsPerms.EnabledOrganizations); err != nil {
		return err
	}
	if err := d.Set("actions_allowed_actions", actionsPerms.AllowedActions); err != nil {
		return err
	}

	// Read allowed actions if policy is "selected"
	if actionsPerms.AllowedActions != nil && *actionsPerms.AllowedActions == "selected" {
		allowedActions, _, err := client.Actions.GetActionsAllowedInEnterprise(ctx, enterpriseSlug)
		if err != nil {
			return err
		}

		if err := d.Set("actions_github_owned_allowed", allowedActions.GithubOwnedAllowed); err != nil {
			return err
		}
		if err := d.Set("actions_verified_allowed", allowedActions.VerifiedAllowed); err != nil {
			return err
		}
		if len(allowedActions.PatternsAllowed) > 0 {
			if err := d.Set("actions_patterns_allowed", allowedActions.PatternsAllowed); err != nil {
				return err
			}
		}
	}

	// Read workflow permissions
	workflowPerms, _, err := client.Actions.GetDefaultWorkflowPermissionsInEnterprise(ctx, enterpriseSlug)
	if err != nil {
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

func resourceGithubEnterpriseSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	// Enterprise settings don't get "deleted", they revert to defaults - just remove from state
	log.Printf("[DEBUG] Removing enterprise settings from Terraform state: %s", d.Id())
	return nil
}

// expandStringSet converts a schema.Set to a slice of strings
// TODO: might be useful in other places, consider moving to utility
func expandStringSet(set *schema.Set) []string {
	result := make([]string, set.Len())
	for i, v := range set.List() {
		result[i] = v.(string)
	}
	return result
}
