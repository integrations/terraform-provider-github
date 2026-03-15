package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type enterpriseRunnerGroup struct {
	NetworkConfigurationID *string `json:"network_configuration_id,omitempty"`
}

func resourceGithubActionsEnterpriseRunnerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsEnterpriseRunnerGroupCreate,
		Read:   resourceGithubActionsEnterpriseRunnerGroupRead,
		Update: resourceGithubActionsEnterpriseRunnerGroupUpdate,
		Delete: resourceGithubActionsEnterpriseRunnerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubActionsEnterpriseRunnerGroupImport,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
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
			"visibility": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The visibility of the runner group.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "selected"}, false)),
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
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "The identifier of the hosted compute network configuration to associate with this runner group for GitHub-hosted private networking.",
			},
			"selected_organization_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Optional:    true,
				Description: "List of organization IDs that can access the runner group.",
			},
			"selected_organizations_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitHub API URL for the runner group's organizations.",
			},
		},
	}
}

func getEnterpriseRunnerGroupNetworking(client *github.Client, ctx context.Context, enterprise string, groupID int64) (*enterpriseRunnerGroup, *github.Response, error) {
	req, err := client.NewRequest("GET", fmt.Sprintf("enterprises/%s/actions/runner-groups/%d", enterprise, groupID), nil)
	if err != nil {
		return nil, nil, err
	}

	var runnerGroup enterpriseRunnerGroup
	resp, err := client.Do(ctx, req, &runnerGroup)
	if err != nil {
		return nil, resp, err
	}

	return &runnerGroup, resp, nil
}

func updateEnterpriseRunnerGroupNetworking(client *github.Client, ctx context.Context, enterprise string, groupID int64, networkConfigurationID *string) (*github.Response, error) {
	payload := map[string]any{
		"network_configuration_id": networkConfigurationID,
	}

	req, err := client.NewRequest("PATCH", fmt.Sprintf("enterprises/%s/actions/runner-groups/%d", enterprise, groupID), payload)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func setGithubActionsEnterpriseRunnerGroupNetworkingState(d *schema.ResourceData, runnerGroup *enterpriseRunnerGroup) error {
	if runnerGroup != nil && runnerGroup.NetworkConfigurationID != nil && *runnerGroup.NetworkConfigurationID != "" {
		if err := d.Set("network_configuration_id", *runnerGroup.NetworkConfigurationID); err != nil {
			return err
		}
	} else {
		if err := d.Set("network_configuration_id", nil); err != nil {
			return err
		}
	}

	return nil
}

func setGithubActionsEnterpriseRunnerGroupState(d *schema.ResourceData, runnerGroup *github.EnterpriseRunnerGroup, etag string, enterpriseSlug string, selectedOrganizationIDs []int64) error {
	if err := d.Set("etag", etag); err != nil {
		return err
	}
	if err := d.Set("allows_public_repositories", runnerGroup.GetAllowsPublicRepositories()); err != nil {
		return err
	}
	if err := d.Set("default", runnerGroup.GetDefault()); err != nil {
		return err
	}
	if err := d.Set("name", runnerGroup.GetName()); err != nil {
		return err
	}
	if err := d.Set("runners_url", runnerGroup.GetRunnersURL()); err != nil {
		return err
	}
	if err := d.Set("selected_organizations_url", runnerGroup.GetSelectedOrganizationsURL()); err != nil {
		return err
	}
	if err := d.Set("visibility", runnerGroup.GetVisibility()); err != nil {
		return err
	}
	if err := d.Set("selected_organization_ids", selectedOrganizationIDs); err != nil {
		return err
	}
	if err := d.Set("restricted_to_workflows", runnerGroup.GetRestrictedToWorkflows()); err != nil {
		return err
	}
	if err := d.Set("selected_workflows", runnerGroup.SelectedWorkflows); err != nil {
		return err
	}
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return err
	}

	return nil
}

func resourceGithubActionsEnterpriseRunnerGroupCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client

	name := d.Get("name").(string)
	enterpriseSlug := d.Get("enterprise_slug").(string)
	restrictedToWorkflows := d.Get("restricted_to_workflows").(bool)
	visibility := d.Get("visibility").(string)
	selectedOrganizations, hasSelectedOrganizations := d.GetOk("selected_organization_ids")
	allowsPublicRepositories := d.Get("allows_public_repositories").(bool)

	selectedWorkflows := []string{}
	if workflows, ok := d.GetOk("selected_workflows"); ok {
		for _, workflow := range workflows.([]any) {
			selectedWorkflows = append(selectedWorkflows, workflow.(string))
		}
	}

	if visibility != "selected" && hasSelectedOrganizations {
		return fmt.Errorf("cannot use selected_organization_ids without visibility being set to selected")
	}

	selectedOrganizationIDs := []int64{}

	if hasSelectedOrganizations {
		ids := selectedOrganizations.(*schema.Set).List()

		for _, id := range ids {
			selectedOrganizationIDs = append(selectedOrganizationIDs, int64(id.(int)))
		}
	}

	ctx := context.Background()

	runnerGroup, resp, err := client.Enterprise.CreateEnterpriseRunnerGroup(ctx,
		enterpriseSlug,
		github.CreateEnterpriseRunnerGroupRequest{
			Name:                     &name,
			Visibility:               &visibility,
			SelectedOrganizationIDs:  selectedOrganizationIDs,
			AllowsPublicRepositories: &allowsPublicRepositories,
			RestrictedToWorkflows:    &restrictedToWorkflows,
			SelectedWorkflows:        selectedWorkflows,
		},
	)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(runnerGroup.GetID(), 10))
	if err = setGithubActionsEnterpriseRunnerGroupState(d, runnerGroup, resp.Header.Get("ETag"), enterpriseSlug, selectedOrganizationIDs); err != nil {
		return err
	}

	if networkConfigurationID, ok := d.GetOk("network_configuration_id"); ok {
		networkConfigurationIDValue := networkConfigurationID.(string)
		// The create endpoint does not accept network_configuration_id, so private networking
		// must be attached with a follow-up PATCH after the runner group has been created.
		if _, err = updateEnterpriseRunnerGroupNetworking(client, ctx, enterpriseSlug, runnerGroup.GetID(), &networkConfigurationIDValue); err != nil {
			return err
		}

		if err = setGithubActionsEnterpriseRunnerGroupNetworkingState(d, &enterpriseRunnerGroup{NetworkConfigurationID: &networkConfigurationIDValue}); err != nil {
			return err
		}
	} else {
		if err = setGithubActionsEnterpriseRunnerGroupNetworkingState(d, nil); err != nil {
			return err
		}
	}

	return nil
}

func getEnterpriseRunnerGroup(client *github.Client, ctx context.Context, ent string, groupID int64) (*github.EnterpriseRunnerGroup, *github.Response, error) {
	enterpriseRunnerGroup, resp, err := client.Enterprise.GetEnterpriseRunnerGroup(ctx, ent, groupID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotModified {
			// ignore error StatusNotModified
			return enterpriseRunnerGroup, resp, nil
		}
	}
	return enterpriseRunnerGroup, resp, err
}

func resourceGithubActionsEnterpriseRunnerGroupRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)
	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)
	ctx = tflog.SetField(ctx, "id", d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	enterpriseRunnerGroup, resp, err := getEnterpriseRunnerGroup(client, ctx, enterpriseSlug, runnerGroupID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing enterprise runner group from state because it no longer exists in GitHub")
				d.SetId("")
				return nil
			}
		}
		return err
	}

	// if runner group is nil (typically not modified) we can return early
	if enterpriseRunnerGroup == nil {
		return nil
	}
	runnerGroupEtag := resp.Header.Get("ETag")

	runnerGroupNetworking, _, err := getEnterpriseRunnerGroupNetworking(client, context.WithValue(context.Background(), ctxId, d.Id()), enterpriseSlug, runnerGroupID)
	if err != nil {
		return err
	}

	selectedOrganizationIDs := []int64{}
	optionsOrgs := github.ListOptions{
		PerPage: maxPerPage,
	}

	for {
		enterpriseRunnerGroupOrganizations, resp, err := client.Enterprise.ListOrganizationAccessRunnerGroup(ctx, enterpriseSlug, runnerGroupID, &optionsOrgs)
		if err != nil {
			return err
		}

		for _, org := range enterpriseRunnerGroupOrganizations.Organizations {
			selectedOrganizationIDs = append(selectedOrganizationIDs, *org.ID)
		}

		if resp.NextPage == 0 {
			break
		}

		optionsOrgs.Page = resp.NextPage
	}

	if err = setGithubActionsEnterpriseRunnerGroupState(d, enterpriseRunnerGroup, runnerGroupEtag, enterpriseSlug, selectedOrganizationIDs); err != nil {
		return err
	}
	if err = setGithubActionsEnterpriseRunnerGroupNetworkingState(d, runnerGroupNetworking); err != nil {
		return err
	}

	return nil
}

func resourceGithubActionsEnterpriseRunnerGroupUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client

	name := d.Get("name").(string)
	enterpriseSlug := d.Get("enterprise_slug").(string)
	visibility := d.Get("visibility").(string)
	restrictedToWorkflows := d.Get("restricted_to_workflows").(bool)
	selectedWorkflows := []string{}
	allowsPublicRepositories := d.Get("allows_public_repositories").(bool)
	if workflows, ok := d.GetOk("selected_workflows"); ok {
		for _, workflow := range workflows.([]any) {
			selectedWorkflows = append(selectedWorkflows, workflow.(string))
		}
	}

	options := github.UpdateEnterpriseRunnerGroupRequest{
		Name:                     &name,
		Visibility:               &visibility,
		RestrictedToWorkflows:    &restrictedToWorkflows,
		SelectedWorkflows:        selectedWorkflows,
		AllowsPublicRepositories: &allowsPublicRepositories,
	}

	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	if _, _, err := client.Enterprise.UpdateEnterpriseRunnerGroup(ctx, enterpriseSlug, runnerGroupID, options); err != nil {
		return err
	}

	if d.HasChange("network_configuration_id") {
		var networkConfigurationIDValue *string
		if networkConfigurationID, ok := d.GetOk("network_configuration_id"); ok {
			value := networkConfigurationID.(string)
			networkConfigurationIDValue = &value
		}

		if _, err := updateEnterpriseRunnerGroupNetworking(client, ctx, enterpriseSlug, runnerGroupID, networkConfigurationIDValue); err != nil {
			return err
		}
	}

	selectedOrganizations, hasSelectedOrganizations := d.GetOk("selected_organization_ids")
	selectedOrganizationIDs := []int64{}

	if hasSelectedOrganizations {
		ids := selectedOrganizations.(*schema.Set).List()

		for _, id := range ids {
			selectedOrganizationIDs = append(selectedOrganizationIDs, int64(id.(int)))
		}
	}

	orgOptions := github.SetOrgAccessRunnerGroupRequest{SelectedOrganizationIDs: selectedOrganizationIDs}

	if _, err := client.Enterprise.SetOrganizationAccessRunnerGroup(ctx, enterpriseSlug, runnerGroupID, orgOptions); err != nil {
		return err
	}

	return resourceGithubActionsEnterpriseRunnerGroupRead(d, meta)
}

func resourceGithubActionsEnterpriseRunnerGroupDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	enterpriseRunnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	ctx = tflog.SetField(ctx, "enterprise_slug", enterpriseSlug)
	ctx = tflog.SetField(ctx, "id", d.Id())
	ctx = tflog.SetField(ctx, "name", d.Get("name"))

	tflog.Debug(ctx, "Deleting enterprise runner group")
	_, err = client.Enterprise.DeleteEnterpriseRunnerGroup(ctx, enterpriseSlug, enterpriseRunnerGroupID)
	return err
}

func resourceGithubActionsEnterpriseRunnerGroupImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>/<runner_group_id>")
	}

	enterpriseId, runnerGroupID := parts[0], parts[1]

	d.SetId(runnerGroupID)
	_ = d.Set("enterprise_slug", enterpriseId)

	return []*schema.ResourceData{d}, nil
}
