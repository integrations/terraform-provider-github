package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseCostCenter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubEnterpriseCostCenterCreate,
		ReadContext:   resourceGithubEnterpriseCostCenterRead,
		UpdateContext: resourceGithubEnterpriseCostCenterUpdate,
		DeleteContext: resourceGithubEnterpriseCostCenterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseCostCenterImport,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the cost center.",
			},
			"users": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The usernames assigned to this cost center.",
			},
			"organizations": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The organization logins assigned to this cost center.",
			},
			"repositories": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The repositories (full name) assigned to this cost center.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the cost center.",
			},
			"azure_subscription": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Azure subscription associated with the cost center.",
			},
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource identifier (username, organization name, or repo full name).",
						},
					},
				},
			},
		},
	}
}

func resourceGithubEnterpriseCostCenterCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	name := d.Get("name").(string)

	ctx = context.WithValue(ctx, ctxId, fmt.Sprintf("%s/%s", enterpriseSlug, name))
	log.Printf("[INFO] Creating enterprise cost center: %s (%s)", name, enterpriseSlug)

	cc, err := enterpriseCostCenterCreate(ctx, client, enterpriseSlug, name)
	if err != nil {
		return diag.FromErr(err)
	}

	if cc == nil || cc.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to create cost center: missing id in response"))
	}

	d.SetId(cc.ID)

	if hasCostCenterAssignmentsConfigured(d) {
		// Ensure we operate on fresh API state before mutations.
		current, err := enterpriseCostCenterGet(ctx, client, enterpriseSlug, cc.ID)
		if err != nil {
			return diag.FromErr(err)
		}
		if diags := syncEnterpriseCostCenterAssignments(ctx, d, client, enterpriseSlug, cc.ID, current.Resources); diags.HasError() {
			return diags
		}
	}

	return resourceGithubEnterpriseCostCenterRead(ctx, d, meta)
}

func resourceGithubEnterpriseCostCenterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Id()

	ctx = context.WithValue(ctx, ctxId, fmt.Sprintf("%s/%s", enterpriseSlug, costCenterID))

	cc, err := enterpriseCostCenterGet(ctx, client, enterpriseSlug, costCenterID)
	if err != nil {
		if is404(err) {
			// If the API starts returning 404 for archived cost centers, we remove it from state.
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	_ = d.Set("name", cc.Name)

	state := strings.ToLower(cc.State)
	if state == "" {
		state = "active"
	}
	_ = d.Set("state", state)
	_ = d.Set("azure_subscription", cc.AzureSubscription)

	resources := make([]map[string]any, 0)
	for _, r := range cc.Resources {
		resources = append(resources, map[string]any{
			"type": r.Type,
			"name": r.Name,
		})
	}
	_ = d.Set("resources", resources)

	users, organizations, repositories := enterpriseCostCenterSplitResources(cc.Resources)
	sort.Strings(users)
	sort.Strings(organizations)
	sort.Strings(repositories)
	_ = d.Set("users", stringSliceToAnySlice(users))
	_ = d.Set("organizations", stringSliceToAnySlice(organizations))
	_ = d.Set("repositories", stringSliceToAnySlice(repositories))

	return nil
}

func resourceGithubEnterpriseCostCenterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Id()

	ctx = context.WithValue(ctx, ctxId, fmt.Sprintf("%s/%s", enterpriseSlug, costCenterID))

	cc, err := enterpriseCostCenterGet(ctx, client, enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}
	if strings.EqualFold(cc.State, "deleted") {
		return diag.FromErr(fmt.Errorf("cannot update cost center %q because it is archived", costCenterID))
	}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		log.Printf("[INFO] Updating enterprise cost center: %s/%s", enterpriseSlug, costCenterID)
		_, err := enterpriseCostCenterUpdate(ctx, client, enterpriseSlug, costCenterID, name)
		if err != nil {
			return diag.FromErr(err)
		}

		cc, err = enterpriseCostCenterGet(ctx, client, enterpriseSlug, costCenterID)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("users") || d.HasChange("organizations") || d.HasChange("repositories") {
		if diags := syncEnterpriseCostCenterAssignments(ctx, d, client, enterpriseSlug, costCenterID, cc.Resources); diags.HasError() {
			return diags
		}
	}

	return resourceGithubEnterpriseCostCenterRead(ctx, d, meta)
}

func resourceGithubEnterpriseCostCenterDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Id()

	ctx = context.WithValue(ctx, ctxId, fmt.Sprintf("%s/%s", enterpriseSlug, costCenterID))
	log.Printf("[INFO] Archiving enterprise cost center: %s/%s", enterpriseSlug, costCenterID)

	_, err := enterpriseCostCenterArchive(ctx, client, enterpriseSlug, costCenterID)
	if err != nil {
		if is404(err) {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>/<cost_center_id>")
	}

	enterpriseSlug, costCenterID := parts[0], parts[1]
	d.SetId(costCenterID)
	_ = d.Set("enterprise_slug", enterpriseSlug)

	return []*schema.ResourceData{d}, nil
}

func syncEnterpriseCostCenterAssignments(ctx context.Context, d *schema.ResourceData, client *github.Client, enterpriseSlug, costCenterID string, currentResources []enterpriseCostCenterResource) diag.Diagnostics {
	desiredUsers := expandStringSet(getStringSetOrEmpty(d, "users"))
	desiredOrgs := expandStringSet(getStringSetOrEmpty(d, "organizations"))
	desiredRepos := expandStringSet(getStringSetOrEmpty(d, "repositories"))

	currentUsers, currentOrgs, currentRepos := enterpriseCostCenterSplitResources(currentResources)

	toAddUsers, toRemoveUsers := diffStringSlices(currentUsers, desiredUsers)
	toAddOrgs, toRemoveOrgs := diffStringSlices(currentOrgs, desiredOrgs)
	toAddRepos, toRemoveRepos := diffStringSlices(currentRepos, desiredRepos)

	const maxResourcesPerRequest = 50
	const costCenterResourcesRetryTimeout = 5 * time.Minute

	retryRemove := func(req enterpriseCostCenterResourcesRequest) diag.Diagnostics {
		//nolint:staticcheck
		err := resource.RetryContext(ctx, costCenterResourcesRetryTimeout, func() *resource.RetryError {
			_, err := enterpriseCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, req)
			if err == nil {
				return nil
			}
			if isRetryableGithubResponseError(err) {
				//nolint:staticcheck
				return resource.RetryableError(err)
			}
			//nolint:staticcheck
			return resource.NonRetryableError(err)
		})
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}

	retryAssign := func(req enterpriseCostCenterResourcesRequest) diag.Diagnostics {
		//nolint:staticcheck
		err := resource.RetryContext(ctx, costCenterResourcesRetryTimeout, func() *resource.RetryError {
			_, err := enterpriseCostCenterAssignResources(ctx, client, enterpriseSlug, costCenterID, req)
			if err == nil {
				return nil
			}
			if isRetryableGithubResponseError(err) {
				//nolint:staticcheck
				return resource.RetryableError(err)
			}
			//nolint:staticcheck
			return resource.NonRetryableError(err)
		})
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}

	chunk := func(items []string) [][]string {
		if len(items) == 0 {
			return nil
		}
		const size = maxResourcesPerRequest
		chunks := make([][]string, 0, (len(items)+size-1)/size)
		for start := 0; start < len(items); start += size {
			end := min(start+size, len(items))
			chunks = append(chunks, items[start:end])
		}
		return chunks
	}

	if len(toRemoveUsers)+len(toRemoveOrgs)+len(toRemoveRepos) > 0 {
		log.Printf("[INFO] Removing enterprise cost center resources: %s/%s", enterpriseSlug, costCenterID)

		for _, batch := range chunk(toRemoveUsers) {
			if diags := retryRemove(enterpriseCostCenterResourcesRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunk(toRemoveOrgs) {
			if diags := retryRemove(enterpriseCostCenterResourcesRequest{Organizations: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunk(toRemoveRepos) {
			if diags := retryRemove(enterpriseCostCenterResourcesRequest{Repositories: batch}); diags.HasError() {
				return diags
			}
		}
	}

	if len(toAddUsers)+len(toAddOrgs)+len(toAddRepos) > 0 {
		log.Printf("[INFO] Assigning enterprise cost center resources: %s/%s", enterpriseSlug, costCenterID)

		for _, batch := range chunk(toAddUsers) {
			if diags := retryAssign(enterpriseCostCenterResourcesRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunk(toAddOrgs) {
			if diags := retryAssign(enterpriseCostCenterResourcesRequest{Organizations: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunk(toAddRepos) {
			if diags := retryAssign(enterpriseCostCenterResourcesRequest{Repositories: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
}

func hasCostCenterAssignmentsConfigured(d *schema.ResourceData) bool {
	assignmentKeys := []string{"users", "organizations", "repositories"}
	for _, key := range assignmentKeys {
		if v, ok := d.GetOkExists(key); ok {
			if set, ok := v.(*schema.Set); ok && set != nil && set.Len() > 0 {
				return true
			}
			if !ok {
				// Non-set values still indicate explicit configuration.
				return true
			}
		}
	}
	return false
}

func expandStringSet(set *schema.Set) []string {
	if set == nil {
		return nil
	}

	list := set.List()
	out := make([]string, 0, len(list))
	for _, v := range list {
		out = append(out, v.(string))
	}
	sort.Strings(out)
	return out
}

func getStringSetOrEmpty(d *schema.ResourceData, key string) *schema.Set {
	v, ok := d.GetOk(key)
	if !ok || v == nil {
		return schema.NewSet(schema.HashString, []any{})
	}

	set, ok := v.(*schema.Set)
	if !ok || set == nil {
		return schema.NewSet(schema.HashString, []any{})
	}

	return set
}

func diffStringSlices(current, desired []string) (toAdd, toRemove []string) {
	cur := schema.NewSet(schema.HashString, stringSliceToAnySlice(current))
	des := schema.NewSet(schema.HashString, stringSliceToAnySlice(desired))

	for _, v := range des.Difference(cur).List() {
		toAdd = append(toAdd, v.(string))
	}
	for _, v := range cur.Difference(des).List() {
		toRemove = append(toRemove, v.(string))
	}

	sort.Strings(toAdd)
	sort.Strings(toRemove)
	return toAdd, toRemove
}

func isRetryableGithubResponseError(err error) bool {
	var ghErr *github.ErrorResponse
	if errors.As(err, &ghErr) && ghErr.Response != nil {
		switch ghErr.Response.StatusCode {
		case 404, 409, 500, 502, 503, 504:
			return true
		default:
			return false
		}
	}
	return false
}
