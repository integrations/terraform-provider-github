package github

import (
	"context"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRepositoryPermissions() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsRepositoryPermissionsCreate,
		ReadContext:   resourceGithubActionsRepositoryPermissionsRead,
		UpdateContext: resourceGithubActionsRepositoryPermissionsUpdate,
		DeleteContext: resourceGithubActionsRepositoryPermissionsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"allowed_actions": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The permissions policy that controls the actions that are allowed to run. Can be one of: 'all', 'local_only', or 'selected'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "local_only", "selected"}, false)),
			},
			"allowed_actions_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Sets the actions that are allowed in an repository. Only available when 'allowed_actions' = 'selected'. ",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"github_owned_allowed": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether GitHub-owned actions are allowed in the repository.",
						},
						"patterns_allowed": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, 'monalisa/octocat@', 'monalisa/octocat@v2', 'monalisa/'.",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         schema.HashString,
						},
						"verified_allowed": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether actions in GitHub Marketplace from verified creators are allowed. Set to 'true' to allow all GitHub Marketplace actions by verified creators.",
						},
					},
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Should GitHub actions be enabled on this repository.",
			},
			"repository": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The GitHub repository.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
			},
			"sha_pinning_required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether pinning to a specific SHA is required for all actions and reusable workflows in a repository.",
			},
		},
	}
}

func resourceGithubActionsRepositoryAllowedObject(d *schema.ResourceData) *github.ActionsAllowed {
	allowed := &github.ActionsAllowed{}

	config := d.Get("allowed_actions_config").([]any)
	if len(config) > 0 {
		data := config[0].(map[string]any)
		switch x := data["github_owned_allowed"].(type) {
		case bool:
			allowed.GithubOwnedAllowed = &x
		}

		switch x := data["verified_allowed"].(type) {
		case bool:
			allowed.VerifiedAllowed = &x
		}

		patternsAllowed := []string{}

		switch t := data["patterns_allowed"].(type) {
		case *schema.Set:
			for _, value := range t.List() {
				patternsAllowed = append(patternsAllowed, value.(string))
			}
		}

		allowed.PatternsAllowed = patternsAllowed
	} else {
		return nil
	}

	return allowed
}

func resourceGithubActionsRepositoryPermissionsCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	allowedActions := d.Get("allowed_actions").(string)
	enabled := d.Get("enabled").(bool)
	tflog.Debug(ctx, "Create repository actions permissions.", map[string]any{"enabled": enabled})

	repoActionPermissions := github.ActionsPermissionsRepository{
		Enabled: &enabled,
	}

	// Only specify `allowed_actions` if actions are enabled
	if enabled {
		repoActionPermissions.AllowedActions = &allowedActions
	}

	if v, ok := d.GetOkExists("sha_pinning_required"); ok { //nolint:staticcheck,SA1019 // Use `GetOkExists` to detect explicit false for booleans.
		repoActionPermissions.SHAPinningRequired = github.Ptr(v.(bool))
	}

	_, _, err := client.Repositories.UpdateActionsPermissions(ctx,
		owner,
		repoName,
		repoActionPermissions,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	if allowedActions == "selected" {
		actionsAllowedData := resourceGithubActionsRepositoryAllowedObject(d)
		if actionsAllowedData != nil {
			tflog.Debug(ctx, "Set allowed actions configuration.")
			_, _, err = client.Repositories.EditActionsAllowed(ctx,
				owner,
				repoName,
				*actionsAllowedData)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			tflog.Debug(ctx, "Skip setting allowed actions configuration because none is specified.")
		}
	}

	d.SetId(repoName)
	return resourceGithubActionsRepositoryPermissionsRead(ctx, d, meta)
}

func resourceGithubActionsRepositoryPermissionsUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	allowedActions := d.Get("allowed_actions").(string)
	enabled := d.Get("enabled").(bool)
	tflog.Debug(ctx, "Update repository actions permissions.", map[string]any{"enabled": enabled})

	repoActionPermissions := github.ActionsPermissionsRepository{
		Enabled: &enabled,
	}

	// Specify `allowed_actions` only if actions are enabled.
	if enabled {
		repoActionPermissions.AllowedActions = &allowedActions
	}

	if d.HasChange("sha_pinning_required") {
		repoActionPermissions.SHAPinningRequired = github.Ptr(d.Get("sha_pinning_required").(bool))
	}

	_, _, err := client.Repositories.UpdateActionsPermissions(ctx,
		owner,
		repoName,
		repoActionPermissions,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	if allowedActions == "selected" {
		actionsAllowedData := resourceGithubActionsRepositoryAllowedObject(d)
		if actionsAllowedData != nil {
			tflog.Debug(ctx, "Update allowed actions configuration.")
			_, _, err = client.Repositories.EditActionsAllowed(ctx,
				owner,
				repoName,
				*actionsAllowedData)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			tflog.Debug(ctx, "Skip updating allowed actions configuration because none is specified.")
		}
	}

	if d.HasChange("sha_pinning_required") {
		if err := d.Set("sha_pinning_required", d.Get("sha_pinning_required").(bool)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubActionsRepositoryPermissionsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Id()

	actionsPermissions, _, err := client.Repositories.GetActionsPermissions(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	// only load and fill allowed_actions_config if allowed_actions_config is also set
	// in the TF code. (see #2105)
	// on initial import there might not be any value in the state, then we have to import the data
	// -> but we can only load an existing state if the current config is set to "selected" (see #2182)
	allowedActions := d.Get("allowed_actions").(string)
	allowedActionsConfig := d.Get("allowed_actions_config").([]any)

	serverHasAllowedActionsConfig := actionsPermissions.GetAllowedActions() == "selected" && actionsPermissions.GetEnabled()
	userWantsAllowedActionsConfig := (allowedActions == "selected" && len(allowedActionsConfig) > 0) || allowedActions == ""

	if serverHasAllowedActionsConfig && userWantsAllowedActionsConfig {
		actionsAllowed, _, err := client.Repositories.GetActionsAllowed(ctx, owner, repoName)
		if err != nil {
			return diag.FromErr(err)
		}

		// If actionsAllowed set to local/all by removing all actions config settings, the response will be empty
		if actionsAllowed != nil {
			if err = d.Set("allowed_actions_config", []any{
				map[string]any{
					"github_owned_allowed": actionsAllowed.GetGithubOwnedAllowed(),
					"patterns_allowed":     actionsAllowed.PatternsAllowed,
					"verified_allowed":     actionsAllowed.GetVerifiedAllowed(),
				},
			}); err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		if err = d.Set("allowed_actions_config", []any{}); err != nil {
			return diag.FromErr(err)
		}
	}

	if err = d.Set("allowed_actions", actionsPermissions.GetAllowedActions()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("enabled", actionsPermissions.GetEnabled()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("sha_pinning_required", actionsPermissions.GetSHAPinningRequired()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRepositoryPermissionsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()

	// Reset the repo to "default" settings
	repoActionPermissions := github.ActionsPermissionsRepository{
		AllowedActions: github.Ptr("all"),
		Enabled:        github.Ptr(true),
	}

	_, _, err := client.Repositories.UpdateActionsPermissions(ctx,
		owner,
		repoName,
		repoActionPermissions,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
