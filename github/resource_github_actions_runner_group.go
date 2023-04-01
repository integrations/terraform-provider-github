package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v50/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubActionsRunnerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsRunnerGroupCreate,
		Read:   resourceGithubActionsRunnerGroupRead,
		Update: resourceGithubActionsRunnerGroupUpdate,
		Delete: resourceGithubActionsRunnerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"allows_public_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
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
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The visibility of the runner group.",
				ValidateFunc: validation.StringInSlice([]string{"all", "selected", "private"}, false),
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

func resourceGithubActionsRunnerGroupCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	name := d.Get("name").(string)
	restrictedToWorkflows := d.Get("restricted_to_workflows").(bool)
	visibility := d.Get("visibility").(string)
	selectedRepositories, hasSelectedRepositories := d.GetOk("selected_repository_ids")

	selectedWorkflows := []string{}
	if workflows, ok := d.GetOk("selected_workflows"); ok {
		for _, workflow := range workflows.([]interface{}) {
			selectedWorkflows = append(selectedWorkflows, workflow.(string))
		}
	}

	if visibility != "selected" && hasSelectedRepositories {
		return fmt.Errorf("cannot use selected_repository_ids without visibility being set to selected")
	}

	selectedRepositoryIDs := []int64{}

	if hasSelectedRepositories {
		ids := selectedRepositories.(*schema.Set).List()

		for _, id := range ids {
			selectedRepositoryIDs = append(selectedRepositoryIDs, int64(id.(int)))
		}
	}

	ctx := context.Background()

	runnerGroup, resp, err := client.Actions.CreateOrganizationRunnerGroup(ctx,
		orgName,
		github.CreateRunnerGroupRequest{
			Name:                  &name,
			Visibility:            &visibility,
			RestrictedToWorkflows: &restrictedToWorkflows,
			SelectedRepositoryIDs: selectedRepositoryIDs,
			SelectedWorkflows:     selectedWorkflows,
		},
	)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(runnerGroup.GetID(), 10))
	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("allows_public_repositories", runnerGroup.GetAllowsPublicRepositories())
	d.Set("default", runnerGroup.GetDefault())
	d.Set("id", runnerGroup.GetID())
	d.Set("inherited", runnerGroup.GetInherited())
	d.Set("name", runnerGroup.GetName())
	d.Set("runners_url", runnerGroup.GetRunnersURL())
	d.Set("selected_repositories_url", runnerGroup.GetSelectedRepositoriesURL())
	d.Set("visibility", runnerGroup.GetVisibility())
	d.Set("selected_repository_ids", selectedRepositoryIDs) // Note: runnerGroup has no method to get selected repository IDs
	d.Set("restricted_to_workflows", runnerGroup.GetRestrictedToWorkflows())
	d.Set("selected_workflows", runnerGroup.SelectedWorkflows)

	return resourceGithubActionsRunnerGroupRead(d, meta)
}

func getOrganizationRunnerGroup(client *github.Client, ctx context.Context, org string, groupID int64) (*github.RunnerGroup, *github.Response, error) {
	runnerGroup, resp, err := client.Actions.GetOrganizationRunnerGroup(ctx, org, groupID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok && ghErr.Response.StatusCode == http.StatusNotModified {
			// ignore error StatusNotModified
			return runnerGroup, resp, nil
		}
	}
	return runnerGroup, resp, err
}

func resourceGithubActionsRunnerGroupRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	runnerGroup, resp, err := getOrganizationRunnerGroup(client, ctx, orgName, runnerGroupID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing organization runner group %s/%s from state because it no longer exists in GitHub",
					orgName, d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("allows_public_repositories", runnerGroup.GetAllowsPublicRepositories())
	d.Set("default", runnerGroup.GetDefault())
	d.Set("id", runnerGroup.GetID())
	d.Set("inherited", runnerGroup.GetInherited())
	d.Set("name", runnerGroup.GetName())
	d.Set("runners_url", runnerGroup.GetRunnersURL())
	d.Set("selected_repositories_url", runnerGroup.GetSelectedRepositoriesURL())
	d.Set("visibility", runnerGroup.GetVisibility())
	d.Set("restricted_to_workflows", runnerGroup.GetRestrictedToWorkflows())
	d.Set("selected_workflows", runnerGroup.SelectedWorkflows)

	selectedRepositoryIDs := []int64{}
	options := github.ListOptions{
		PerPage: maxPerPage,
	}

	for {
		runnerGroupRepositories, resp, err := client.Actions.ListRepositoryAccessRunnerGroup(ctx, orgName, runnerGroupID, &options)
		if err != nil {
			return err
		}

		for _, repo := range runnerGroupRepositories.Repositories {
			selectedRepositoryIDs = append(selectedRepositoryIDs, *repo.ID)
		}

		if resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	d.Set("selected_repository_ids", selectedRepositoryIDs)

	return nil
}

func resourceGithubActionsRunnerGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	name := d.Get("name").(string)
	visibility := d.Get("visibility").(string)
	restrictedToWorkflows := d.Get("restricted_to_workflows").(bool)
	selectedWorkflows := []string{}
	if workflows, ok := d.GetOk("selected_workflows"); ok {
		for _, workflow := range workflows.([]interface{}) {
			selectedWorkflows = append(selectedWorkflows, workflow.(string))
		}
	}

	options := github.UpdateRunnerGroupRequest{
		Name:                  &name,
		Visibility:            &visibility,
		RestrictedToWorkflows: &restrictedToWorkflows,
		SelectedWorkflows:     selectedWorkflows,
	}

	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	if _, _, err := client.Actions.UpdateOrganizationRunnerGroup(ctx, orgName, runnerGroupID, options); err != nil {
		return err
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
		return err
	}

	return resourceGithubActionsRunnerGroupRead(d, meta)
}

func resourceGithubActionsRunnerGroupDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[INFO] Deleting organization runner group: %s (%s)", d.Id(), orgName)
	_, err = client.Actions.DeleteOrganizationRunnerGroup(ctx, orgName, runnerGroupID)
	return err
}
