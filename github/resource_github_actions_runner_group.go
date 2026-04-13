package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRunnerGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsRunnerGroupCreate,
		ReadContext:   resourceGithubActionsRunnerGroupRead,
		UpdateContext: resourceGithubActionsRunnerGroupUpdate,
		DeleteContext: resourceGithubActionsRunnerGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the runner group.",
			},
			"allows_public_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether public repositories can be added to the runner group.",
			},
			"default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this is the default runner group.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the runner group object",
			},
			"inherited": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the runner group is inherited from the enterprise level",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the runner group.",
			},
			"runners_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GitHub API URL for the runner group's runners.",
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Optional:    true,
				Description: "List of repository IDs that can access the runner group.",
			},
			"selected_repositories_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitHub API URL for the runner group's repositories.",
			},
			"visibility": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The visibility of the runner group.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "selected", "private"}, false)),
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
			"network_configuration_id": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 255)),
				Description:      "The identifier of the hosted compute network configuration to associate with this runner group for GitHub-hosted private networking.",
			},
		},
	}
}

func getOrganizationRunnerGroup(client *github.Client, ctx context.Context, org string, groupID int64) (*github.RunnerGroup, *github.Response, error) {
	runnerGroup, resp, err := client.Actions.GetOrganizationRunnerGroup(ctx, org, groupID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotModified {
			// ignore error StatusNotModified
			return nil, resp, nil
		}
	}
	return runnerGroup, resp, err
}

func setGithubActionsRunnerGroupState(d *schema.ResourceData, runnerGroup *github.RunnerGroup, etag string, selectedRepositoryIDs []int64) error {
	if err := d.Set("etag", normalizeEtag(etag)); err != nil {
		return err
	}
	if err := d.Set("allows_public_repositories", runnerGroup.GetAllowsPublicRepositories()); err != nil {
		return err
	}
	if err := d.Set("default", runnerGroup.GetDefault()); err != nil {
		return err
	}
	if err := d.Set("id", strconv.FormatInt(runnerGroup.GetID(), 10)); err != nil {
		return err
	}
	if err := d.Set("inherited", runnerGroup.GetInherited()); err != nil {
		return err
	}
	if err := d.Set("name", runnerGroup.GetName()); err != nil {
		return err
	}
	if err := d.Set("runners_url", runnerGroup.GetRunnersURL()); err != nil {
		return err
	}
	if err := d.Set("selected_repositories_url", runnerGroup.GetSelectedRepositoriesURL()); err != nil {
		return err
	}
	if err := d.Set("visibility", runnerGroup.GetVisibility()); err != nil {
		return err
	}
	if err := d.Set("selected_repository_ids", selectedRepositoryIDs); err != nil {
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

func resourceGithubActionsRunnerGroupCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	name := d.Get("name").(string)
	restrictedToWorkflows := d.Get("restricted_to_workflows").(bool)
	visibility := d.Get("visibility").(string)
	selectedRepositories, hasSelectedRepositories := d.GetOk("selected_repository_ids")
	allowsPublicRepositories := d.Get("allows_public_repositories").(bool)

	selectedWorkflows := []string{}
	if workflows, ok := d.GetOk("selected_workflows"); ok {
		for _, workflow := range workflows.([]any) {
			selectedWorkflows = append(selectedWorkflows, workflow.(string))
		}
	}

	if visibility != "selected" && hasSelectedRepositories {
		return diag.FromErr(fmt.Errorf("cannot use selected_repository_ids without visibility being set to selected"))
	}

	selectedRepositoryIDs := []int64{}

	if hasSelectedRepositories {
		ids := selectedRepositories.(*schema.Set).List()

		for _, id := range ids {
			selectedRepositoryIDs = append(selectedRepositoryIDs, int64(id.(int)))
		}
	}

	var networkConfigurationIDPtr *string
	if networkConfigurationID, ok := d.GetOk("network_configuration_id"); ok {
		value := networkConfigurationID.(string)
		networkConfigurationIDPtr = &value
	}

	runnerGroup, resp, err := client.Actions.CreateOrganizationRunnerGroup(ctx,
		orgName,
		github.CreateRunnerGroupRequest{
			Name:                     &name,
			Visibility:               &visibility,
			RestrictedToWorkflows:    &restrictedToWorkflows,
			SelectedRepositoryIDs:    selectedRepositoryIDs,
			SelectedWorkflows:        selectedWorkflows,
			AllowsPublicRepositories: &allowsPublicRepositories,
			NetworkConfigurationID:   networkConfigurationIDPtr,
		},
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(runnerGroup.GetID(), 10))
	if err = setGithubActionsRunnerGroupState(d, runnerGroup, normalizeEtag(resp.Header.Get("ETag")), selectedRepositoryIDs); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("network_configuration_id", runnerGroup.NetworkConfigurationID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRunnerGroupRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	runnerGroup, resp, err := getOrganizationRunnerGroup(client, ctx, orgName, runnerGroupID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing organization runner group %s/%s from state because it no longer exists in GitHub",
					orgName, d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if runnerGroup == nil {
		return nil
	}

	selectedRepositoryIDs := []int64{}
	options := github.ListOptions{
		PerPage: maxPerPage,
	}

	for {
		runnerGroupRepositories, resp, err := client.Actions.ListRepositoryAccessRunnerGroup(ctx, orgName, runnerGroupID, &options)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, repo := range runnerGroupRepositories.Repositories {
			selectedRepositoryIDs = append(selectedRepositoryIDs, *repo.ID)
		}

		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	runnerGroupEtag := normalizeEtag(resp.Header.Get("ETag"))
	if err = setGithubActionsRunnerGroupState(d, runnerGroup, runnerGroupEtag, selectedRepositoryIDs); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("network_configuration_id", runnerGroup.NetworkConfigurationID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRunnerGroupUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	name := d.Get("name").(string)
	visibility := d.Get("visibility").(string)
	restrictedToWorkflows := d.Get("restricted_to_workflows").(bool)
	selectedWorkflows := []string{}
	allowsPublicRepositories := d.Get("allows_public_repositories").(bool)
	if workflows, ok := d.GetOk("selected_workflows"); ok {
		for _, workflow := range workflows.([]any) {
			selectedWorkflows = append(selectedWorkflows, workflow.(string))
		}
	}

	var networkConfigurationIDPtr *string
	if networkConfigurationID, ok := d.GetOk("network_configuration_id"); ok {
		value := networkConfigurationID.(string)
		networkConfigurationIDPtr = &value
	} else if d.HasChange("network_configuration_id") {
		// Field was removed — send empty string to clear it.
		// go-github's omitempty omits nil pointers, so empty string is used as a workaround.
		empty := ""
		networkConfigurationIDPtr = &empty
	}

	options := github.UpdateRunnerGroupRequest{
		Name:                     &name,
		Visibility:               &visibility,
		RestrictedToWorkflows:    &restrictedToWorkflows,
		SelectedWorkflows:        selectedWorkflows,
		AllowsPublicRepositories: &allowsPublicRepositories,
		NetworkConfigurationID:   networkConfigurationIDPtr,
	}

	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())

	runnerGroup, resp, err := client.Actions.UpdateOrganizationRunnerGroup(ctx, orgName, runnerGroupID, options)
	if err != nil {
		return diag.FromErr(err)
	}

	selectedRepositories, hasSelectedRepositories := d.GetOk("selected_repository_ids")
	selectedRepositoryIDs := []int64{}

	if hasSelectedRepositories {
		ids := selectedRepositories.(*schema.Set).List()

		for _, id := range ids {
			selectedRepositoryIDs = append(selectedRepositoryIDs, int64(id.(int)))
		}
	}

	reposOptions := github.SetRepoAccessRunnerGroupRequest{SelectedRepositoryIDs: selectedRepositoryIDs}

	if _, err := client.Actions.SetRepositoryAccessRunnerGroup(ctx, orgName, runnerGroupID, reposOptions); err != nil {
		return diag.FromErr(err)
	}

	runnerGroupEtag := normalizeEtag(resp.Header.Get("ETag"))

	if err := setGithubActionsRunnerGroupState(d, runnerGroup, runnerGroupEtag, selectedRepositoryIDs); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("network_configuration_id", runnerGroup.NetworkConfigurationID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRunnerGroupDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}
	ctx = context.WithValue(ctx, ctxId, d.Id())

	log.Printf("[INFO] Deleting organization runner group: %s (%s)", d.Id(), orgName)
	_, err = client.Actions.DeleteOrganizationRunnerGroup(ctx, orgName, runnerGroupID)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
