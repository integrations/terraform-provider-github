package github

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseCostCenterUsers() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages user assignments for a GitHub enterprise cost center (authoritative).",
		CreateContext: resourceGithubEnterpriseCostCenterUsersCreate,
		ReadContext:   resourceGithubEnterpriseCostCenterUsersRead,
		UpdateContext: resourceGithubEnterpriseCostCenterUsersUpdate,
		DeleteContext: resourceGithubEnterpriseCostCenterUsersDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseCostCenterUsersImport,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "The slug of the enterprise.",
			},
			"cost_center_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "The ID of the cost center.",
			},
			"usernames": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Set:      caseInsensitiveStringHash,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
					StateFunc:        caseInsensitiveStringState,
				},
				Description: "Usernames to assign to the cost center. This is authoritative - users not in this set will be removed.",
			},
		},
	}
}

func resourceGithubEnterpriseCostCenterUsersCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("unexpected provider metadata type %T", meta)
	}
	enterpriseSlug, ok := resourceKeysGetOk[string](d, "enterprise_slug")
	if !ok {
		return diag.Errorf("expected enterprise_slug to be a non-empty string")
	}
	costCenterID, ok := resourceKeysGetOk[string](d, "cost_center_id")
	if !ok {
		return diag.Errorf("expected cost_center_id to be a non-empty string")
	}

	cc, err := getEnterpriseCostCenter(ctx, owner.v3client, enterpriseSlug, costCenterID, owner.maxPerPage)
	if err != nil {
		return diag.FromErr(err)
	}
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeUser {
			return diag.Errorf("cost center %q already has users assigned; import the existing assignments first or remove them manually", costCenterID)
		}
	}

	usernames, ok := resourceKeysGetOk[*schema.Set](d, "usernames")
	if !ok {
		return diag.Errorf("expected usernames to be a non-empty set")
	}
	toAdd := expandStringList(usernames.List())

	tflog.Info(ctx, "Adding users to cost center", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"cost_center_id":  costCenterID,
		"count":           len(toAdd),
	})

	d.SetId(costCenterID)
	for batch := range slices.Chunk(toAdd, maxCostCenterResourcesPerRequest) {
		if _, _, err := owner.v3client.Enterprise.AddResourcesToCostCenter(ctx, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterUsersUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("unexpected provider metadata type %T", meta)
	}
	enterpriseSlug, ok := resourceKeysGetOk[string](d, "enterprise_slug")
	if !ok {
		return diag.Errorf("expected enterprise_slug to be a non-empty string")
	}
	costCenterID, ok := resourceKeysGetOk[string](d, "cost_center_id")
	if !ok {
		return diag.Errorf("expected cost_center_id to be a non-empty string")
	}

	cc, err := getEnterpriseCostCenter(ctx, owner.v3client, enterpriseSlug, costCenterID, owner.maxPerPage)
	if err != nil {
		return diag.FromErr(err)
	}

	var currentUsers []string
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeUser {
			currentUsers = append(currentUsers, ccResource.Name)
		}
	}

	usernames, ok := resourceKeysGetOk[*schema.Set](d, "usernames")
	if !ok {
		return diag.Errorf("expected usernames to be a non-empty set")
	}
	desiredUsers := expandStringList(usernames.List())
	toAdd, toRemove := caseInsensitiveStringDifference(currentUsers, desiredUsers)

	if len(toRemove) > 0 {
		tflog.Info(ctx, "Removing users from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toRemove),
		})

		for batch := range slices.Chunk(toRemove, maxCostCenterResourcesPerRequest) {
			if _, _, err := owner.v3client.Enterprise.RemoveResourcesFromCostCenter(ctx, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if len(toAdd) > 0 {
		tflog.Info(ctx, "Adding users to cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toAdd),
		})

		for batch := range slices.Chunk(toAdd, maxCostCenterResourcesPerRequest) {
			if _, _, err := owner.v3client.Enterprise.AddResourcesToCostCenter(ctx, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterUsersRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("unexpected provider metadata type %T", meta)
	}
	enterpriseSlug, ok := resourceKeysGetOk[string](d, "enterprise_slug")
	if !ok {
		return diag.Errorf("expected enterprise_slug to be a non-empty string")
	}
	costCenterID, ok := resourceKeysGetOk[string](d, "cost_center_id")
	if !ok {
		return diag.Errorf("expected cost_center_id to be a non-empty string")
	}

	cc, err := getEnterpriseCostCenter(ctx, owner.v3client, enterpriseSlug, costCenterID, owner.maxPerPage)
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
	if cc.GetState() == "deleted" {
		tflog.Warn(ctx, "Cost center is archived, removing user assignments from state", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
		})
		d.SetId("")
		return nil
	}

	var users []string
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeUser {
			users = append(users, caseInsensitiveStringState(ccResource.Name))
		}
	}

	if err := d.Set("usernames", flattenStringList(users)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterUsersDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("unexpected provider metadata type %T", meta)
	}
	enterpriseSlug, ok := resourceKeysGetOk[string](d, "enterprise_slug")
	if !ok {
		return diag.Errorf("expected enterprise_slug to be a non-empty string")
	}
	costCenterID, ok := resourceKeysGetOk[string](d, "cost_center_id")
	if !ok {
		return diag.Errorf("expected cost_center_id to be a non-empty string")
	}

	cc, err := getEnterpriseCostCenter(ctx, owner.v3client, enterpriseSlug, costCenterID, owner.maxPerPage)
	if err != nil {
		if errIs404(err) {
			return nil
		}
		return diag.FromErr(err)
	}
	if cc.GetState() == "deleted" {
		return nil
	}

	var usernames []string
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeUser {
			usernames = append(usernames, ccResource.Name)
		}
	}

	if len(usernames) > 0 {
		tflog.Info(ctx, "Removing all users from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(usernames),
		})

		for batch := range slices.Chunk(usernames, maxCostCenterResourcesPerRequest) {
			if _, _, err := owner.v3client.Enterprise.RemoveResourcesFromCostCenter(ctx, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterUsersImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	enterpriseSlug, costCenterID, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid import ID %q: expected format <enterprise_slug>:<cost_center_id>", d.Id())
	}

	d.SetId(costCenterID)
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	if err := d.Set("cost_center_id", costCenterID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
