package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsOrganizationVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsOrganizationVariableCreate,
		ReadContext:   resourceGithubActionsOrganizationVariableRead,
		UpdateContext: resourceGithubActionsOrganizationVariableUpdate,
		DeleteContext: resourceGithubActionsOrganizationVariableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsOrganizationVariableImport,
		},

		CustomizeDiff: diffSecretVariableVisibility,

		Description: "Resource to manage a GitHub Actions variable for an organization.",

		Schema: map[string]*schema.Schema{
			"variable_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the variable.",
				ValidateDiagFunc: validateSecretNameFunc,
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Value of the variable.",
			},
			"visibility": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "private", "selected"}, false)),
				Description:      "Configures the access that repositories have to the organization variable. Must be one of 'all', 'private', or 'selected'. 'selected_repository_ids' is required if set to 'selected'.",
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Set:  schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional:    true,
				Description: "An array of repository ids that can access the organization variable.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the variable was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the variable was last updated.",
			},
		},
	}
}

func resourceGithubActionsOrganizationVariableCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	varName, _ := d.Get("variable_name").(string)
	visibility, _ := d.Get("visibility").(string)
	value, _ := d.Get("value").(string)

	var repoIDs []int64

	if v, ok := d.GetOk("selected_repository_ids"); ok {
		ids := v.(*schema.Set).List()

		for _, id := range ids {
			repoIDs = append(repoIDs, int64(id.(int)))
		}
	}

	varReq := github.OrgActionsVariableCreateRequest{
		Name:                  varName,
		Value:                 value,
		Visibility:            visibility,
		SelectedRepositoryIDs: repoIDs,
	}

	if _, err := client.Actions.CreateOrgVariable(ctx, owner, varReq); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(varName)

	// GitHub API does not return on create so we have to lookup the variable to get timestamps.
	if variable, err := retryUntilResourceFound(ctx, func() (*github.ActionsVariable, error) {
		val, _, err := client.Actions.GetOrgVariable(ctx, owner, varName)
		return val, err
	}, nil); err == nil {
		if err := d.Set("created_at", variable.CreatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("updated_at", variable.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubActionsOrganizationVariableRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	varName, _ := d.Get("variable_name").(string)

	variable, _, err := client.Actions.GetOrgVariable(ctx, owner, varName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing organization actions variable from state because it no longer exists.", map[string]any{"variable_name": varName})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("value", variable.Value); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("visibility", variable.Visibility); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", variable.CreatedAt.String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", variable.UpdatedAt.String()); err != nil {
		return diag.FromErr(err)
	}

	if variable.GetVisibility() == "selected" {
		if _, ok := d.GetOk("selected_repository_ids"); ok {
			var repoIDs []int64
			for repo, err := range client.Actions.ListSelectedReposForOrgVariableIter(ctx, owner, varName, &github.ListOptions{PerPage: maxPerPage}) {
				if err != nil {
					return diag.FromErr(err)
				}

				repoIDs = append(repoIDs, repo.GetID())
			}

			if err := d.Set("selected_repository_ids", repoIDs); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}

func resourceGithubActionsOrganizationVariableUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	varName, _ := d.Get("variable_name").(string)
	varValue, _ := d.Get("value").(string)
	visibility, _ := d.Get("visibility").(string)

	var repoIDs []int64
	if v, ok := d.GetOk("selected_repository_ids"); ok {
		ids := v.(*schema.Set).List()

		for _, id := range ids {
			repoIDs = append(repoIDs, int64(id.(int)))
		}
	}

	varReq := github.OrgActionsVariableUpdateRequest{
		Name:                  new(varName),
		Value:                 new(varValue),
		Visibility:            new(visibility),
		SelectedRepositoryIDs: repoIDs,
	}

	if _, err := client.Actions.UpdateOrgVariable(ctx, owner, varName, varReq); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(varName)

	// GitHub API does not return on update so we have to lookup the variable to get timestamps.
	if variable, err := retryUntilResourceFound(ctx, func() (*github.ActionsVariable, error) {
		val, _, err := client.Actions.GetOrgVariable(ctx, owner, varName)
		return val, err
	}, nil); err == nil {
		if err := d.Set("created_at", variable.CreatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("updated_at", variable.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubActionsOrganizationVariableDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	varName, _ := d.Get("variable_name").(string)

	if _, err := client.Actions.DeleteOrgVariable(ctx, owner, varName); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationVariableImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	varName := d.Id()

	variable, _, err := client.Actions.GetOrgVariable(ctx, owner, varName)
	if err != nil {
		return nil, err
	}

	if err := d.Set("variable_name", varName); err != nil {
		return nil, err
	}
	if err := d.Set("value", variable.Value); err != nil {
		return nil, err
	}
	if err := d.Set("visibility", variable.Visibility); err != nil {
		return nil, err
	}
	if err := d.Set("created_at", variable.CreatedAt.String()); err != nil {
		return nil, err
	}
	if err := d.Set("updated_at", variable.UpdatedAt.String()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
