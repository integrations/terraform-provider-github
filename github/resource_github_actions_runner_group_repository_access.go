package github

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsRunnerGroupRepositoryAccess() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages the access of a repository to an organization runner group.",
		CreateContext: resourceGithubActionsRunnerGroupRepositoryAccessCreate,
		ReadContext:   resourceGithubActionsRunnerGroupRepositoryAccessRead,
		// Omitting update function since this resource expresses a simple membership in the set of repositories with access to the runner group, it either exists or doesn't
		// UpdateContext: resourceGithubActionsRunnerGroupRepositoryAccessUpdate,
		DeleteContext: resourceGithubActionsRunnerGroupRepositoryAccessDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsRunnerGroupRepositoryAccessImport,
		},

		Schema: map[string]*schema.Schema{
			"runner_group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the runner group to grant the repository access on",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the repository to grant access to the runner group",
			},
		},
	}
}

func resourceGithubActionsRunnerGroupRepositoryAccessCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	runnerGroupId := d.Get("runner_group_id").(int)
	repositoryId := d.Get("repository_id").(int)

	_, err := client.Actions.AddRepositoryAccessRunnerGroup(ctx, orgName, int64(runnerGroupId), int64(repositoryId))

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d/%d", runnerGroupId, repositoryId))

	return resourceGithubActionsRunnerGroupRepositoryAccessRead(ctx, d, meta)
}

func resourceGithubActionsRunnerGroupRepositoryAccessRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	runnerGroupId := d.Get("runner_group_id").(int)
	repositoryId := d.Get("repository_id").(int)

	for repo, err := range client.Actions.ListRepositoryAccessRunnerGroupIter(ctx, orgName, int64(runnerGroupId), nil) {
		if err != nil {
			return diag.FromErr(err)
		}

		if *repo.ID == int64(repositoryId) {
			// Resource matches the state in github exactly (repo has access), no need for further modifications
			return nil
		}
	}
	// We reached the end of the list without a match for our desired repository access
	log.Printf("[INFO] Removing access of repository with id %d to runner group %s/%d from state because access no longer exists in GitHub", repositoryId, orgName, runnerGroupId)
	d.SetId("")
	return nil

}

func resourceGithubActionsRunnerGroupRepositoryAccessDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	runnerGroupId := d.Get("runner_group_id").(int)
	repositoryId := d.Get("repository_id").(int)

	_, err := client.Actions.RemoveRepositoryAccessRunnerGroup(ctx, orgName, int64(runnerGroupId), int64(repositoryId))
	return diag.FromErr(err)

}

func resourceGithubActionsRunnerGroupRepositoryAccessImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	id := d.Id()
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <runner_group_id>/<repository_id>")
	}

	runnerGroupID, err := strconv.Atoi(parts[0])

	if err != nil {
		return nil, fmt.Errorf("runner_group_id in id must be convertible to an int: %w", err)
	}

	err = d.Set("runner_group_id", runnerGroupID)
	if err != nil {
		return nil, err
	}

	repositoryId, err := strconv.Atoi(parts[1])

	if err != nil {
		return nil, fmt.Errorf("runner_id in id must be convertible to an int: %w", err)
	}

	err = d.Set("repository_id", repositoryId)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
