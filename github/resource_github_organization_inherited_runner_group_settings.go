package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v85/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationInheritedRunnerGroupSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Manages organization-level settings for an enterprise Actions runner group inherited by the organization.",

		CreateContext: resourceGithubOrganizationInheritedRunnerGroupSettingsCreate,
		ReadContext:   resourceGithubOrganizationInheritedRunnerGroupSettingsRead,
		UpdateContext: resourceGithubOrganizationInheritedRunnerGroupSettingsUpdate,
		DeleteContext: resourceGithubOrganizationInheritedRunnerGroupSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubOrganizationInheritedRunnerGroupSettingsImport,
		},

		Schema: map[string]*schema.Schema{
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub organization name.",
			},
			"enterprise_runner_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the enterprise runner group inherited by the organization.",
			},
			"runner_group_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the inherited enterprise runner group in the organization.",
			},
			"inherited": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this runner group is inherited from the enterprise.",
			},
			"visibility": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "selected",
				Description:      "The visibility of the runner group. Can be 'all', 'selected', or 'private'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "selected", "private"}, false)),
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Optional:    true,
				Description: "List of repository IDs that can access the runner group. Only applicable when visibility is set to 'selected'.",
			},
			"allows_public_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether public repositories can be added to the runner group.",
			},
			"restricted_to_workflows": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If 'true', the runner group will be restricted to running only the workflows specified in the 'selected_workflows' array. Defaults to 'false'.",
			},
			"selected_workflows": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to 'true'.",
			},
		},
	}
}

func findInheritedEnterpriseRunnerGroupByName(client *github.Client, ctx context.Context, org, name string) (*github.RunnerGroup, error) {
	for group, err := range client.Actions.ListOrganizationRunnerGroupsIter(ctx, org, nil) {
		if err != nil {
			return nil, err
		}
		if group.GetInherited() && group.GetName() == name {
			return group, nil
		}
	}

	return nil, fmt.Errorf("inherited enterprise runner group '%s' not found in organization '%s'. Ensure the enterprise runner group is shared with this organization", name, org)
}

func resourceGithubOrganizationInheritedRunnerGroupSettingsCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	org := d.Get("organization").(string)
	enterpriseRunnerGroupName := d.Get("enterprise_runner_group_name").(string)
	visibility := d.Get("visibility").(string)
	allowsPublicRepositories := d.Get("allows_public_repositories").(bool)
	restrictedToWorkflows := d.Get("restricted_to_workflows").(bool)
	selectedRepositories, hasSelectedRepositories := d.GetOk("selected_repository_ids")

	selectedWorkflows := []string{}
	if workflows, ok := d.GetOk("selected_workflows"); ok {
		for _, workflow := range workflows.([]any) {
			selectedWorkflows = append(selectedWorkflows, workflow.(string))
		}
	}

	// Find the inherited enterprise runner group by name
	runnerGroup, err := findInheritedEnterpriseRunnerGroupByName(client, ctx, org, enterpriseRunnerGroupName)
	if err != nil {
		return diag.FromErr(err)
	}

	runnerGroupID := runnerGroup.GetID()
	id, err := buildID(org, strconv.FormatInt(runnerGroupID, 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("runner_group_id", int(runnerGroupID)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("inherited", runnerGroup.GetInherited()); err != nil {
		return diag.FromErr(err)
	}

	// Update runner group settings
	updateReq := github.UpdateRunnerGroupRequest{
		Visibility:               new(visibility),
		AllowsPublicRepositories: new(allowsPublicRepositories),
		RestrictedToWorkflows:    new(restrictedToWorkflows),
		SelectedWorkflows:        selectedWorkflows,
	}

	_, _, err = client.Actions.UpdateOrganizationRunnerGroup(ctx, org, runnerGroupID, updateReq)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set repository access if visibility is "selected"
	if visibility == "selected" && hasSelectedRepositories {
		selectedRepositoryIDs := []int64{}
		for _, id := range selectedRepositories.(*schema.Set).List() {
			selectedRepositoryIDs = append(selectedRepositoryIDs, int64(id.(int)))
		}

		repoAccessReq := github.SetRepoAccessRunnerGroupRequest{
			SelectedRepositoryIDs: selectedRepositoryIDs,
		}

		_, err = client.Actions.SetRepositoryAccessRunnerGroup(ctx, org, runnerGroupID, repoAccessReq)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubOrganizationInheritedRunnerGroupSettingsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	org := d.Get("organization").(string)
	runnerGroupID := int64(d.Get("runner_group_id").(int))

	// Get the runner group details
	runnerGroup, _, err := client.Actions.GetOrganizationRunnerGroup(ctx, org, runnerGroupID)
	if err != nil {
		ghErr := &github.ErrorResponse{}
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing actions organization runner group from state because it no longer exists in GitHub", map[string]any{
					"runner_group_id": d.Id(),
				})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err := d.Set("inherited", runnerGroup.GetInherited()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("visibility", runnerGroup.GetVisibility()); err != nil {
		return diag.FromErr(err)
	}

	// Get repository access list only if visibility is "selected"
	if runnerGroup.GetVisibility() == "selected" {
		selectedRepositoryIDs := []int64{}

		for repo, err := range client.Actions.ListRepositoryAccessRunnerGroupIter(ctx, org, runnerGroupID, nil) {
			if err != nil {
				return diag.FromErr(err)
			}
			selectedRepositoryIDs = append(selectedRepositoryIDs, repo.GetID())
		}

		if err := d.Set("selected_repository_ids", selectedRepositoryIDs); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("selected_repository_ids", []int64{}); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("allows_public_repositories", runnerGroup.GetAllowsPublicRepositories()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("restricted_to_workflows", runnerGroup.GetRestrictedToWorkflows()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("selected_workflows", runnerGroup.SelectedWorkflows); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationInheritedRunnerGroupSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	org := d.Get("organization").(string)
	runnerGroupID := int64(d.Get("runner_group_id").(int))
	visibility := d.Get("visibility").(string)
	selectedRepositories, hasSelectedRepositories := d.GetOk("selected_repository_ids")

	// Update runner group settings if any relevant fields changed
	if d.HasChange("visibility") || d.HasChange("allows_public_repositories") || d.HasChange("restricted_to_workflows") || d.HasChange("selected_workflows") {
		allowsPublicRepositories := d.Get("allows_public_repositories").(bool)
		restrictedToWorkflows := d.Get("restricted_to_workflows").(bool)

		selectedWorkflows := []string{}
		if workflows, ok := d.GetOk("selected_workflows"); ok {
			for _, workflow := range workflows.([]any) {
				selectedWorkflows = append(selectedWorkflows, workflow.(string))
			}
		}

		updateReq := github.UpdateRunnerGroupRequest{
			Visibility:               new(visibility),
			AllowsPublicRepositories: new(allowsPublicRepositories),
			RestrictedToWorkflows:    new(restrictedToWorkflows),
			SelectedWorkflows:        selectedWorkflows,
		}

		_, _, err := client.Actions.UpdateOrganizationRunnerGroup(ctx, org, runnerGroupID, updateReq)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Update repository access if changed and visibility is "selected"
	if d.HasChange("selected_repository_ids") && visibility == "selected" && hasSelectedRepositories {
		selectedRepositoryIDs := []int64{}

		for _, id := range selectedRepositories.(*schema.Set).List() {
			selectedRepositoryIDs = append(selectedRepositoryIDs, int64(id.(int)))
		}

		repoAccessReq := github.SetRepoAccessRunnerGroupRequest{
			SelectedRepositoryIDs: selectedRepositoryIDs,
		}

		_, err := client.Actions.SetRepositoryAccessRunnerGroup(ctx, org, runnerGroupID, repoAccessReq)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubOrganizationInheritedRunnerGroupSettingsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	org := d.Get("organization").(string)
	runnerGroupID := int64(d.Get("runner_group_id").(int))

	tflog.Info(ctx, "Removing repository access for runner group", map[string]any{
		"runner_group_id": d.Id(),
	})

	// Reset to "all" visibility and clear repository access
	updateReq := github.UpdateRunnerGroupRequest{
		Visibility: new("all"),
	}

	_, _, err := client.Actions.UpdateOrganizationRunnerGroup(ctx, org, runnerGroupID, updateReq)
	if err != nil {
		// If the runner group doesn't exist, that's fine
		ghErr := &github.ErrorResponse{}
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				return nil
			}
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationInheritedRunnerGroupSettingsImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	org, identifier, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid import ID format, expected 'organization:enterprise_runner_group_name' or 'organization:organization_runner_group_id'")
	}

	client := meta.(*Owner).v3client

	var runnerGroup *github.RunnerGroup

	// Try to parse as ID first
	if id, parseErr := strconv.ParseInt(identifier, 10, 64); parseErr == nil {
		// It's an ID - get the runner group and verify it's inherited
		runnerGroup, _, err = client.Actions.GetOrganizationRunnerGroup(ctx, org, id)
		if err != nil {
			return nil, fmt.Errorf("failed to get runner group: %w", err)
		}
	} else {
		// It's a name - find the inherited enterprise runner group
		runnerGroup, err = findInheritedEnterpriseRunnerGroupByName(client, ctx, org, identifier)
		if err != nil {
			return nil, err
		}
	}

	// Verify the runner group is inherited from the enterprise
	if !runnerGroup.GetInherited() {
		return nil, fmt.Errorf("runner group '%s' is not inherited from the enterprise. This resource only manages inherited enterprise runner groups", runnerGroup.GetName())
	}

	id, err := buildID(org, strconv.FormatInt(runnerGroup.GetID(), 10))
	if err != nil {
		return nil, err
	}
	d.SetId(id)
	if err = d.Set("organization", org); err != nil {
		return nil, err
	}
	if err = d.Set("enterprise_runner_group_name", runnerGroup.GetName()); err != nil {
		return nil, err
	}
	if err = d.Set("runner_group_id", int(runnerGroup.GetID())); err != nil {
		return nil, err
	}
	if err = d.Set("inherited", runnerGroup.GetInherited()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
