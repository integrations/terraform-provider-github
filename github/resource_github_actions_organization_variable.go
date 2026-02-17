package github

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsOrganizationVariable() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a GitHub Actions variable within an organization.",
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
				Description: "Date of 'actions_variable' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_variable' update.",
			},
		},

		CustomizeDiff: diffSecretVariableVisibility,

		CreateContext: resourceGithubActionsOrganizationVariableCreate,
		ReadContext:   resourceGithubActionsOrganizationVariableRead,
		UpdateContext: resourceGithubActionsOrganizationVariableUpdate,
		DeleteContext: resourceGithubActionsOrganizationVariableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsOrganizationVariableImport,
		},
	}
}

func resourceGithubActionsOrganizationVariableCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	varName := d.Get("variable_name").(string)
	visibility := d.Get("visibility").(string)
	repoIDs := github.SelectedRepoIDs{}

	if v, ok := d.GetOk("selected_repository_ids"); ok {
		ids := v.(*schema.Set).List()

		for _, id := range ids {
			repoIDs = append(repoIDs, int64(id.(int)))
		}
	}

	variable := &github.ActionsVariable{
		Name:                  varName,
		Value:                 d.Get("value").(string),
		Visibility:            github.Ptr(visibility),
		SelectedRepositoryIDs: github.Ptr(repoIDs),
	}
	_, err := client.Actions.CreateOrgVariable(ctx, owner, variable)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(varName)

	// GitHub API does not return on create so we have to lookup the variable to get timestamps
	if variable, _, err := client.Actions.GetOrgVariable(ctx, owner, varName); err == nil {
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
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	varName := d.Get("variable_name").(string)

	variable, _, err := client.Actions.GetOrgVariable(ctx, owner, varName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing actions variable %s from state because it no longer exists in GitHub", d.Id())
				d.SetId("")
				return nil
			}
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
			repoIDs := []int64{}
			opt := &github.ListOptions{
				PerPage: maxPerPage,
			}
			for {
				results, resp, err := client.Actions.ListSelectedReposForOrgVariable(ctx, owner, varName, opt)
				if err != nil {
					return diag.FromErr(err)
				}

				for _, repo := range results.Repositories {
					repoIDs = append(repoIDs, repo.GetID())
				}

				if resp.NextPage == 0 {
					break
				}
				opt.Page = resp.NextPage
			}

			if err := d.Set("selected_repository_ids", repoIDs); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}

func resourceGithubActionsOrganizationVariableUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	varName := d.Get("variable_name").(string)
	visibility := d.Get("visibility").(string)
	repoIDs := github.SelectedRepoIDs{}

	if v, ok := d.GetOk("selected_repository_ids"); ok {
		ids := v.(*schema.Set).List()

		for _, id := range ids {
			repoIDs = append(repoIDs, int64(id.(int)))
		}
	}

	variable := &github.ActionsVariable{
		Name:                  varName,
		Value:                 d.Get("value").(string),
		Visibility:            github.Ptr(visibility),
		SelectedRepositoryIDs: github.Ptr(repoIDs),
	}

	_, err := client.Actions.UpdateOrgVariable(ctx, owner, variable)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(varName)

	// GitHub API does not return on create so we have to lookup the variable to get timestamps
	if variable, _, err := client.Actions.GetOrgVariable(ctx, owner, varName); err == nil {
		if err := d.Set("created_at", variable.CreatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("updated_at", variable.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("updated_at", nil); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubActionsOrganizationVariableDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	varName := d.Get("variable_name").(string)

	_, err := client.Actions.DeleteOrgVariable(ctx, owner, varName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationVariableImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
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

	selectedRepositoryIDs := []int64{}
	if variable.GetVisibility() == "selected" {
		opt := &github.ListOptions{
			PerPage: maxPerPage,
		}
		for {
			results, resp, err := client.Actions.ListSelectedReposForOrgVariable(ctx, owner, varName, opt)
			if err != nil {
				return nil, err
			}

			for _, repo := range results.Repositories {
				selectedRepositoryIDs = append(selectedRepositoryIDs, repo.GetID())
			}

			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}

	if err := d.Set("selected_repository_ids", selectedRepositoryIDs); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
