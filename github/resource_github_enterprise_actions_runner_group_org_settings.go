package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseActionsRunnerGroupOrgSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubEnterpriseActionsRunnerGroupOrgSettingsCreate,
		Read:   resourceGithubEnterpriseActionsRunnerGroupOrgSettingsRead,
		Update: resourceGithubEnterpriseActionsRunnerGroupOrgSettingsUpdate,
		Delete: resourceGithubEnterpriseActionsRunnerGroupOrgSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubEnterpriseActionsRunnerGroupOrgSettingsImport,
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
				ForceNew:    true,
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
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{"all", "selected", "private"}, false), "visibility"),
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

func findInheritedEnterpriseRunnerGroupByName(client *github.Client, ctx context.Context, org string, name string) (*github.RunnerGroup, error) {
	opts := &github.ListOrgRunnerGroupOptions{
		ListOptions: github.ListOptions{
			PerPage: maxPerPage,
		},
	}

	for {
		groups, resp, err := client.Actions.ListOrganizationRunnerGroups(ctx, org, opts)
		if err != nil {
			return nil, err
		}

		for _, group := range groups.RunnerGroups {
			// Only match runner groups that are inherited from the enterprise
			if group.GetName() == name && group.GetInherited() {
				return group, nil
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil, fmt.Errorf("inherited enterprise runner group '%s' not found in organization '%s'. Ensure the enterprise runner group is shared with this organization", name, org)
}

func resourceGithubEnterpriseActionsRunnerGroupOrgSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

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
		return err
	}

	// Verify it's actually inherited
	if !runnerGroup.GetInherited() {
		return fmt.Errorf("runner group '%s' exists but is not inherited from the enterprise. This resource only manages inherited enterprise runner groups", enterpriseRunnerGroupName)
	}

	runnerGroupID := runnerGroup.GetID()
	d.SetId(fmt.Sprintf("%s:%d", org, runnerGroupID))

	// Set the runner group ID and inherited flag
	if err := d.Set("runner_group_id", int(runnerGroupID)); err != nil {
		return err
	}
	if err := d.Set("inherited", runnerGroup.GetInherited()); err != nil {
		return err
	}

	// Update runner group settings
	updateReq := github.UpdateRunnerGroupRequest{
		Visibility:               github.String(visibility),
		AllowsPublicRepositories: &allowsPublicRepositories,
		RestrictedToWorkflows:    &restrictedToWorkflows,
		SelectedWorkflows:        selectedWorkflows,
	}

	_, _, err = client.Actions.UpdateOrganizationRunnerGroup(ctx, org, runnerGroupID, updateReq)
	if err != nil {
		return fmt.Errorf("failed to update runner group: %w", err)
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
			return fmt.Errorf("failed to set repository access: %w", err)
		}
	}

	return resourceGithubEnterpriseActionsRunnerGroupOrgSettingsRead(d, meta)
}

func resourceGithubEnterpriseActionsRunnerGroupOrgSettingsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	org := d.Get("organization").(string)
	runnerGroupID := int64(d.Get("runner_group_id").(int))

	// Get the runner group details
	runnerGroup, _, err := client.Actions.GetOrganizationRunnerGroup(ctx, org, runnerGroupID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing actions organization runner group %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if runnerGroup == nil {
		return nil
	}

	// Verify it's still inherited from the enterprise
	if !runnerGroup.GetInherited() {
		log.Printf("[WARN] Runner group %s is no longer inherited from the enterprise", d.Id())
	}

	// Set inherited flag
	if err := d.Set("inherited", runnerGroup.GetInherited()); err != nil {
		return err
	}

	// Set visibility
	if err := d.Set("visibility", runnerGroup.GetVisibility()); err != nil {
		return err
	}

	// Get repository access list only if visibility is "selected"
	if runnerGroup.GetVisibility() == "selected" {
		selectedRepositoryIDs := []int64{}
		opts := &github.ListOptions{
			PerPage: maxPerPage,
		}

		for {
			repos, resp, err := client.Actions.ListRepositoryAccessRunnerGroup(ctx, org, runnerGroupID, opts)
			if err != nil {
				return fmt.Errorf("failed to list repository access: %w", err)
			}

			for _, repo := range repos.Repositories {
				selectedRepositoryIDs = append(selectedRepositoryIDs, repo.GetID())
			}

			if resp.NextPage == 0 {
				break
			}
			opts.Page = resp.NextPage
		}

		if err := d.Set("selected_repository_ids", selectedRepositoryIDs); err != nil {
			return err
		}
	} else {
		// Clear selected_repository_ids if visibility is not "selected"
		if err := d.Set("selected_repository_ids", []int64{}); err != nil {
			return err
		}
	}

	if err := d.Set("allows_public_repositories", runnerGroup.GetAllowsPublicRepositories()); err != nil {
		return err
	}

	if err := d.Set("restricted_to_workflows", runnerGroup.GetRestrictedToWorkflows()); err != nil {
		return err
	}

	if err := d.Set("selected_workflows", runnerGroup.SelectedWorkflows); err != nil {
		return err
	}

	return nil
}

func resourceGithubEnterpriseActionsRunnerGroupOrgSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

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
			Visibility:               github.String(visibility),
			AllowsPublicRepositories: &allowsPublicRepositories,
			RestrictedToWorkflows:    &restrictedToWorkflows,
			SelectedWorkflows:        selectedWorkflows,
		}

		_, _, err := client.Actions.UpdateOrganizationRunnerGroup(ctx, org, runnerGroupID, updateReq)
		if err != nil {
			return fmt.Errorf("failed to update runner group: %w", err)
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
			return fmt.Errorf("failed to set repository access: %w", err)
		}
	}

	return resourceGithubEnterpriseActionsRunnerGroupOrgSettingsRead(d, meta)
}

func resourceGithubEnterpriseActionsRunnerGroupOrgSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	org := d.Get("organization").(string)
	runnerGroupID := int64(d.Get("runner_group_id").(int))

	log.Printf("[INFO] Removing repository access for runner group: %s", d.Id())

	// Reset to "all" visibility and clear repository access
	updateReq := github.UpdateRunnerGroupRequest{
		Visibility: github.String("all"),
	}

	_, _, err := client.Actions.UpdateOrganizationRunnerGroup(ctx, org, runnerGroupID, updateReq)
	if err != nil {
		// If the runner group doesn't exist, that's fine
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				return nil
			}
		}
		return fmt.Errorf("failed to reset runner group visibility: %w", err)
	}

	return nil
}

func resourceGithubEnterpriseActionsRunnerGroupOrgSettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import ID format, expected 'organization:enterprise_runner_group_name' or 'organization:runner_group_id'")
	}

	org := parts[0]
	identifier := parts[1]

	client := meta.(*Owner).v3client
	ctx := context.Background()

	var runnerGroup *github.RunnerGroup
	var err error

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

	d.SetId(fmt.Sprintf("%s:%d", org, runnerGroup.GetID()))
	d.Set("organization", org)
	d.Set("enterprise_runner_group_name", runnerGroup.GetName())
	d.Set("runner_group_id", int(runnerGroup.GetID()))
	d.Set("inherited", runnerGroup.GetInherited())

	return []*schema.ResourceData{d}, nil
}
