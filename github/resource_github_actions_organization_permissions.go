package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsOrganizationPermissions() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsOrganizationPermissionsCreate,
		ReadContext:   resourceGithubActionsOrganizationPermissionsRead,
		UpdateContext: resourceGithubActionsOrganizationPermissionsUpdate,
		DeleteContext: resourceGithubActionsOrganizationPermissionsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsOrganizationPermissionsImport,
		},

		CustomizeDiff: diffGithubActionsOrganizationPermissions,

		Description: "Resource to manage the GitHub Actions organization permissions.",

		Schema: map[string]*schema.Schema{
			"enabled_repositories": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Policy that controls the repositories in the organization that are allowed to run GitHub Actions; this can be one of `all`, `none`, or `selected`.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "none", "selected"}, false)),
			},
			"enabled_repositories_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Configure the repositories that are enabled for GitHub Actions in an organization; this is only supported when `enabled_repositories` is set to `selected`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository_ids": {
							Type:        schema.TypeSet,
							Description: "List of repository IDs to enable for GitHub Actions.",
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Required:    true,
						},
					},
				},
			},
			"allowed_actions": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Policy that controls the actions that are allowed to run; this can be one of `all`, `local_only`, or `selected`.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "local_only", "selected"}, false)),
			},
			"allowed_actions_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Configure the actions that are allowed in an organization; this is only supported when `allowed_actions` is set to `selected`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"github_owned_allowed": {
							Type:        schema.TypeBool,
							Required:    true, // TODO: Make optional in v7
							Description: "If GitHub-owned actions are allowed.",
						},
						"verified_allowed": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If actions in GitHub Marketplace from verified creators are allowed.",
						},
						"patterns_allowed": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of string-matching patterns to allow specific actions; wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`.",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         schema.HashString,
						},
					},
				},
			},
			"sha_pinning_required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true, // TODO: Remove computed in v7
				Description: "If all actions and reusable workflows in an organization require pinning to a specific SHA.",
			},
		},
	}
}

func diffGithubActionsOrganizationPermissions(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	if !diff.NewValueKnown("enabled_repositories") || !diff.NewValueKnown("enabled_repositories_config") || !diff.NewValueKnown("allowed_actions") || !diff.NewValueKnown("allowed_actions_config") {
		return nil
	}

	enabledRepos, _ := diff.Get("enabled_repositories").(string)
	enabledReposConfig := false

	if v, ok := diff.GetOk("enabled_repositories_config"); ok {
		if p, ok := v.([]any); ok && len(p) > 0 {
			enabledReposConfig = true
		}
	}

	if enabledRepos != "selected" && enabledReposConfig {
		return fmt.Errorf("enabled_repositories_config is only valid when enabled_repositories is set to selected")
	}

	if enabledRepos == "selected" && !enabledReposConfig {
		return fmt.Errorf("enabled_repositories_config is required when enabled_repositories is set to selected")
	}

	allowedActions, _ := diff.Get("allowed_actions").(string)
	allowedActionsConfig := false

	if v, ok := diff.GetOk("allowed_actions_config"); ok {
		if p, ok := v.([]any); ok && len(p) > 0 {
			allowedActionsConfig = true
		}
	}

	if allowedActions != "" && enabledRepos == "none" {
		return fmt.Errorf("allowed_actions is only valid when enabled_repositories is set to all or selected")
	}

	if allowedActions != "selected" && allowedActionsConfig {
		return fmt.Errorf("allowed_actions_config is only valid when allowed_actions is set to selected")
	}

	if allowedActions == "selected" && !allowedActionsConfig {
		return fmt.Errorf("allowed_actions_config is required when allowed_actions is set to selected")
	}

	return nil
}

func resourceGithubActionsOrganizationPermissionsCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	enabledRepositories, _ := d.Get("enabled_repositories").(string)

	req := github.ActionsPermissions{
		EnabledRepositories: &enabledRepositories,
	}

	if v, ok := d.GetOk("allowed_actions"); ok {
		s, _ := v.(string)
		req.AllowedActions = &s
	}

	if v, ok := d.GetOkExists("sha_pinning_required"); ok { //nolint:staticcheck // SA1019: d.GetOkExists is deprecated but necessary for bool fields
		b, _ := v.(bool)
		req.SHAPinningRequired = &b
	}

	if _, _, err := meta.v3client.Actions.UpdateActionsPermissions(ctx, meta.name, req); err != nil {
		return diag.FromErr(err)
	}

	if req.SHAPinningRequired == nil {
		perm, _, err := meta.v3client.Actions.GetActionsPermissions(ctx, meta.name)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("sha_pinning_required", perm.GetSHAPinningRequired()); err != nil {
			return diag.FromErr(err)
		}
	}

	if *req.EnabledRepositories == "selected" {
		if v, ok := d.GetOk("enabled_repositories_config"); ok {
			if p, ok := v.([]any); ok && len(p) > 0 && p[0] != nil {
				block, _ := p[0].(map[string]any)
				col, _ := block["repository_ids"].(*schema.Set)

				repoIDs := make([]int64, col.Len())
				for i, a := range col.List() {
					id, _ := a.(int)
					repoIDs[i] = int64(id)
				}

				if _, err := meta.v3client.Actions.SetEnabledReposInOrg(ctx, meta.name, repoIDs); err != nil {
					return diag.FromErr(err)
				}
			}
		}
	}

	if req.AllowedActions != nil && *req.AllowedActions == "selected" {
		if v, ok := d.GetOk("allowed_actions_config"); ok {
			if p, ok := v.([]any); ok && len(p) > 0 && p[0] != nil {
				block, _ := p[0].(map[string]any)

				ownerAllowed, _ := block["github_owned_allowed"].(bool)

				req := github.ActionsAllowed{
					GithubOwnedAllowed: &ownerAllowed,
				}

				if v, ok := block["verified_allowed"]; ok {
					if b, ok := v.(bool); ok {
						req.VerifiedAllowed = &b
					}
				}

				if v, ok := block["patterns_allowed"]; ok {
					if p, ok := v.(*schema.Set); ok {
						patterns := make([]string, p.Len())
						for i, a := range p.List() {
							patterns[i], _ = a.(string)
						}
						req.PatternsAllowed = patterns
					}
				}

				if _, _, err := meta.v3client.Actions.UpdateActionsAllowed(ctx, meta.name, req); err != nil {
					return diag.FromErr(err)
				}
			}
		}
	}

	d.SetId(meta.name)

	return nil
}

func resourceGithubActionsOrganizationPermissionsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	perm, _, err := meta.v3client.Actions.GetActionsPermissions(ctx, meta.name)
	if err != nil {
		return diag.FromErr(err)
	}

	enabledRepositories := perm.GetEnabledRepositories()
	if err := d.Set("enabled_repositories", enabledRepositories); err != nil {
		return diag.FromErr(err)
	}

	var reposCfg []any
	if enabledRepositories == "selected" {
		repoIDs := []int{}

		for repo, err := range meta.v3client.Actions.ListEnabledReposInOrgIter(ctx, meta.name, &github.ListOptions{PerPage: meta.maxPerPage}) {
			if err != nil {
				return diag.FromErr(err)
			}

			repoIDs = append(repoIDs, int(repo.GetID()))
		}

		reposCfg = []any{map[string]any{
			"repository_ids": repoIDs,
		}}
	}

	if err := d.Set("enabled_repositories_config", reposCfg); err != nil {
		return diag.FromErr(err)
	}

	if _, ok := d.GetOk("allowed_actions"); ok {
		allowedActions := perm.GetAllowedActions()
		if err := d.Set("allowed_actions", allowedActions); err != nil {
			return diag.FromErr(err)
		}

		var actionsCfg []any
		if allowedActions == "selected" {
			act, _, err := meta.v3client.Actions.GetActionsAllowed(ctx, meta.name)
			if err != nil {
				return diag.FromErr(err)
			}

			actionsCfg = []any{map[string]any{
				"github_owned_allowed": act.GetGithubOwnedAllowed(),
				"patterns_allowed":     act.GetPatternsAllowed(),
				"verified_allowed":     act.GetVerifiedAllowed(),
			}}
		}

		if err = d.Set("allowed_actions_config", actionsCfg); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("sha_pinning_required", perm.GetSHAPinningRequired()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationPermissionsUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	enabledRepositories, _ := d.Get("enabled_repositories").(string)
	shaPinningRequired, _ := d.Get("sha_pinning_required").(bool) // TODO: Update for v7 when no longer computed

	req := github.ActionsPermissions{
		EnabledRepositories: &enabledRepositories,
		SHAPinningRequired:  &shaPinningRequired,
	}

	if v, ok := d.GetOk("allowed_actions"); ok {
		s, _ := v.(string)
		req.AllowedActions = &s
	}

	if _, _, err := meta.v3client.Actions.UpdateActionsPermissions(ctx, meta.name, req); err != nil {
		return diag.FromErr(err)
	}

	if *req.EnabledRepositories == "selected" {
		if v, ok := d.GetOk("enabled_repositories_config"); ok {
			if p, ok := v.([]any); ok && len(p) > 0 && p[0] != nil {
				block, _ := p[0].(map[string]any)
				col, _ := block["repository_ids"].(*schema.Set)

				repoIDs := make([]int64, col.Len())
				for i, a := range col.List() {
					id, _ := a.(int)
					repoIDs[i] = int64(id)
				}

				if _, err := meta.v3client.Actions.SetEnabledReposInOrg(ctx, meta.name, repoIDs); err != nil {
					return diag.FromErr(err)
				}
			}
		}
	}

	if req.AllowedActions != nil && *req.AllowedActions == "selected" {
		if v, ok := d.GetOk("allowed_actions_config"); ok {
			if p, ok := v.([]any); ok && len(p) > 0 && p[0] != nil {
				block, _ := p[0].(map[string]any)

				ownerAllowed, _ := block["github_owned_allowed"].(bool)

				req := github.ActionsAllowed{
					GithubOwnedAllowed: &ownerAllowed,
				}

				if v, ok := block["verified_allowed"]; ok {
					if b, ok := v.(bool); ok {
						req.VerifiedAllowed = &b
					}
				}

				if v, ok := block["patterns_allowed"]; ok {
					if p, ok := v.(*schema.Set); ok {
						patterns := make([]string, p.Len())
						for i, a := range p.List() {
							patterns[i], _ = a.(string)
						}
						req.PatternsAllowed = patterns
					}
				}

				if _, _, err := meta.v3client.Actions.UpdateActionsAllowed(ctx, meta.name, req); err != nil {
					return diag.FromErr(err)
				}
			}
		}
	}

	return nil
}

func resourceGithubActionsOrganizationPermissionsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	req := github.ActionsPermissions{
		AllowedActions:      new("all"),
		EnabledRepositories: new("all"),
	}

	if _, _, err := meta.v3client.Actions.UpdateActionsPermissions(ctx, meta.name, req); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == 409 {
				tflog.Info(ctx, "Organization actions permissions couldn't be unset as they're already configured at enterprise level.")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationPermissionsImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)

	if d.Id() != meta.name {
		return nil, fmt.Errorf("invalid id %q, expected %q", d.Id(), meta.name)
	}

	perm, _, err := meta.v3client.Actions.GetActionsPermissions(ctx, meta.name)
	if err != nil {
		return nil, err
	}

	enabledRepositories := perm.GetEnabledRepositories()
	if err := d.Set("enabled_repositories", enabledRepositories); err != nil {
		return nil, err
	}

	var reposCfg []any
	if enabledRepositories == "selected" {
		repoIDs := []int{}

		for repo, err := range meta.v3client.Actions.ListEnabledReposInOrgIter(ctx, meta.name, &github.ListOptions{PerPage: meta.maxPerPage}) {
			if err != nil {
				return nil, err
			}

			repoIDs = append(repoIDs, int(repo.GetID()))
		}

		reposCfg = []any{map[string]any{
			"repository_ids": repoIDs,
		}}
	}

	if err := d.Set("enabled_repositories_config", reposCfg); err != nil {
		return nil, err
	}

	allowedActions := perm.GetAllowedActions()
	if err := d.Set("allowed_actions", allowedActions); err != nil {
		return nil, err
	}

	var actionsCfg []any
	if allowedActions == "selected" {
		act, _, err := meta.v3client.Actions.GetActionsAllowed(ctx, meta.name)
		if err != nil {
			return nil, err
		}

		actionsCfg = []any{map[string]any{
			"github_owned_allowed": act.GetGithubOwnedAllowed(),
			"patterns_allowed":     act.GetPatternsAllowed(),
			"verified_allowed":     act.GetVerifiedAllowed(),
		}}
	}

	if err = d.Set("allowed_actions_config", actionsCfg); err != nil {
		return nil, err
	}

	if err := d.Set("sha_pinning_required", perm.GetSHAPinningRequired()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
