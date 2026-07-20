package github

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsEnvironmentVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsEnvironmentVariableCreate,
		ReadContext:   resourceGithubActionsEnvironmentVariableRead,
		UpdateContext: resourceGithubActionsEnvironmentVariableUpdate,
		DeleteContext: resourceGithubActionsEnvironmentVariableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsEnvironmentVariableImport,
		},

		CustomizeDiff: diffRepository,

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubActionsEnvironmentVariableV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubActionsEnvironmentVariableStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "Resource to manage a GitHub Actions environment variable for a repository environment.",

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
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the environment.",
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
				Description: "Timestamp for when the variable was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp for when the variable was last updated.",
			},
		},
	}
}

func resourceGithubActionsEnvironmentVariableCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	varName, _ := d.Get("variable_name").(string)
	varValue, _ := d.Get("value").(string)

	escapedEnvName := url.PathEscape(envName)

	varReq := github.ActionsVariableCreateRequest{
		Name:  varName,
		Value: varValue,
	}

	_, err := client.Actions.CreateEnvVariable(ctx, owner, repoName, escapedEnvName, varReq)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), varName)
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

	// GitHub API does not return on create so we have to lookup the variable to get timestamps.
	if variable, err := retryUntilResourceFound(ctx, func() (*github.ActionsVariable, error) {
		val, _, err := client.Actions.GetEnvVariable(ctx, owner, repoName, escapedEnvName, varName)
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

func resourceGithubActionsEnvironmentVariableRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	varName, _ := d.Get("variable_name").(string)

	variable, _, err := client.Actions.GetEnvVariable(ctx, owner, repoName, url.PathEscape(envName), varName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing actions variable from state because it no longer exists in GitHub.", map[string]any{"variable_name": varName, "repository": repoName, "environment": envName})
			d.SetId("")
			return nil
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

func resourceGithubActionsEnvironmentVariableUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	varName, _ := d.Get("variable_name").(string)
	varValue, _ := d.Get("value").(string)

	escapedEnvName := url.PathEscape(envName)

	varReq := github.ActionsVariableUpdateRequest{
		Name:  new(varName),
		Value: new(varValue),
	}

	_, err := client.Actions.UpdateEnvVariable(ctx, owner, repoName, escapedEnvName, varName, varReq)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), varName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	// GitHub API does not return on update so we have to lookup the variable to get timestamps.
	if variable, err := retryUntilResourceFound(ctx, func() (*github.ActionsVariable, error) {
		val, _, err := client.Actions.GetEnvVariable(ctx, owner, repoName, escapedEnvName, varName)
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

func resourceGithubActionsEnvironmentVariableDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	varName, _ := d.Get("variable_name").(string)

	_, err := client.Actions.DeleteEnvVariable(ctx, owner, repoName, url.PathEscape(envName), varName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsEnvironmentVariableImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, envNamePart, varName, err := parseID3(d.Id())
	if err != nil {
		return nil, err
	}

	envName := unescapeIDPart(envNamePart)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	repoID := int(repo.GetID())

	variable, _, err := client.Actions.GetEnvVariable(ctx, owner, repoName, url.PathEscape(envName), varName)
	if err != nil {
		return nil, err
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", repoID); err != nil {
		return nil, err
	}
	if err := d.Set("environment", envName); err != nil {
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
