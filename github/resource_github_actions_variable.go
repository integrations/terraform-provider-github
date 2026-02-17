package github

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsVariable() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages a GitHub Actions variable within a repository.",
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubActionsVariableV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubActionsVariableStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the repository.",
			},
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

		CustomizeDiff: diffRepository,

		CreateContext: resourceGithubActionsVariableCreate,
		ReadContext:   resourceGithubActionsVariableRead,
		UpdateContext: resourceGithubActionsVariableUpdate,
		DeleteContext: resourceGithubActionsVariableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsVariableImport,
		},
	}
}

func resourceGithubActionsVariableCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	varName := d.Get("variable_name").(string)

	variable := github.ActionsVariable{
		Name:  varName,
		Value: d.Get("value").(string),
	}

	_, err := client.Actions.CreateRepoVariable(ctx, owner, repoName, &variable)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, varName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	repoID := int(repo.GetID())

	if err := d.Set("repository_id", repoID); err != nil {
		return diag.FromErr(err)
	}

	// GitHub API does not return on create so we have to lookup the variable to get timestamps
	if variable, _, err := client.Actions.GetRepoVariable(ctx, owner, repoName, varName); err == nil {
		if err := d.Set("created_at", variable.CreatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("updated_at", variable.UpdatedAt.String()); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubActionsVariableRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	varName := d.Get("variable_name").(string)

	variable, _, err := client.Actions.GetRepoVariable(ctx, owner, repoName, varName)
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

	if err = d.Set("value", variable.Value); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("created_at", variable.CreatedAt.String()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("updated_at", variable.UpdatedAt.String()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsVariableUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	varName := d.Get("variable_name").(string)

	variable := github.ActionsVariable{
		Name:  varName,
		Value: d.Get("value").(string),
	}

	_, err := client.Actions.UpdateRepoVariable(ctx, owner, repoName, &variable)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, varName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	// GitHub API does not return on create so we have to lookup the variable to get timestamps
	if variable, _, err := client.Actions.GetRepoVariable(ctx, owner, repoName, varName); err == nil {
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

func resourceGithubActionsVariableDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	varName := d.Get("variable_name").(string)

	_, err := client.Actions.DeleteRepoVariable(ctx, owner, repoName, varName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsVariableImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, varName, err := parseID2(d.Id())
	if err != nil {
		return nil, err
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	repoID := int(repo.GetID())

	variable, _, err := client.Actions.GetRepoVariable(ctx, owner, repoName, varName)
	if err != nil {
		return nil, err
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", repoID); err != nil {
		return nil, err
	}
	if err := d.Set("variable_name", varName); err != nil {
		return nil, err
	}
	if err := d.Set("value", variable.Value); err != nil {
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
