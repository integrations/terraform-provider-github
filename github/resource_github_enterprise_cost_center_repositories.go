package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseCostCenterRepositories() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages repository assignments for a GitHub enterprise cost center (authoritative).",
		CreateContext: resourceGithubEnterpriseCostCenterRepositoriesCreate,
		ReadContext:   resourceGithubEnterpriseCostCenterRepositoriesRead,
		UpdateContext: resourceGithubEnterpriseCostCenterRepositoriesUpdate,
		DeleteContext: resourceGithubEnterpriseCostCenterRepositoriesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseCostCenterRepositoriesImport,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"cost_center_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the cost center.",
			},
			"repository_names": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Repository names (full name, e.g. org/repo) to assign to the cost center. This is authoritative - repositories not in this set will be removed.",
			},
		},
	}
}

func resourceGithubEnterpriseCostCenterRepositoriesCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	id, err := buildID(enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	// Get desired repositories from config
	desiredReposSet := d.Get("repository_names").(*schema.Set)
	toAdd := expandStringList(desiredReposSet.List())

	// Add repositories
	if len(toAdd) > 0 {
		tflog.Info(ctx, "Adding repositories to cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toAdd),
		})

		for _, batch := range chunkStringSlice(toAdd, maxResourcesPerRequest) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Repositories: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterRepositoriesUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}

	currentRepos := make(map[string]bool)
	for _, r := range cc.Resources {
		if r != nil && r.Type == "Repo" {
			currentRepos[r.Name] = true
		}
	}

	desiredReposSet := d.Get("repository_names").(*schema.Set)
	desiredRepos := make(map[string]bool)
	for _, repo := range desiredReposSet.List() {
		desiredRepos[repo.(string)] = true
	}

	var toAdd, toRemove []string
	for repo := range desiredRepos {
		if !currentRepos[repo] {
			toAdd = append(toAdd, repo)
		}
	}
	for repo := range currentRepos {
		if !desiredRepos[repo] {
			toRemove = append(toRemove, repo)
		}
	}

	if len(toRemove) > 0 {
		tflog.Info(ctx, "Removing repositories from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toRemove),
		})

		for _, batch := range chunkStringSlice(toRemove, maxResourcesPerRequest) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Repositories: batch}); diags.HasError() {
				return diags
			}
		}
	}

	if len(toAdd) > 0 {
		tflog.Info(ctx, "Adding repositories to cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toAdd),
		})

		for _, batch := range chunkStringSlice(toAdd, maxResourcesPerRequest) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Repositories: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterRepositoriesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		if errIs404(err) {
			tflog.Warn(ctx, "Cost center not found, removing from state", map[string]any{
				"enterprise_slug": enterpriseSlug,
				"cost_center_id":  costCenterID,
			})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	var repositories []string
	for _, r := range cc.Resources {
		if r != nil && r.Type == "Repo" {
			repositories = append(repositories, r.Name)
		}
	}

	if err := d.Set("repository_names", flattenStringList(repositories)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterRepositoriesDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	repositoriesSet := d.Get("repository_names").(*schema.Set)
	repositories := expandStringList(repositoriesSet.List())

	if len(repositories) > 0 {
		tflog.Info(ctx, "Removing all repositories from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(repositories),
		})

		for _, batch := range chunkStringSlice(repositories, maxResourcesPerRequest) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Repositories: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterRepositoriesImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	enterpriseSlug, costCenterID, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid import ID %q: expected format <enterprise_slug>:<cost_center_id>", d.Id())
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	if err := d.Set("cost_center_id", costCenterID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
