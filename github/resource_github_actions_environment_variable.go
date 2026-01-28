package github

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsEnvironmentVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsEnvironmentVariableCreateOrUpdate,
		ReadContext:   resourceGithubActionsEnvironmentVariableRead,
		UpdateContext: resourceGithubActionsEnvironmentVariableCreateOrUpdate,
		DeleteContext: resourceGithubActionsEnvironmentVariableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsEnvironmentVariableImport,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
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
				Description: "Date of 'actions_variable' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_variable' update.",
			},
		},
	}
}

func resourceGithubActionsEnvironmentVariableCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	name := d.Get("variable_name").(string)

	variable := &github.ActionsVariable{
		Name:  name,
		Value: d.Get("value").(string),
	}

	// Try to create the variable first
	_, err := client.Actions.CreateEnvVariable(ctx, owner, repoName, url.PathEscape(envName), variable)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusConflict {
			// Variable already exists, try to update instead
			// If it fails here, we want to return the error otherwise continue
			_, err = client.Actions.UpdateEnvVariable(ctx, owner, repoName, url.PathEscape(envName), variable)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			return diag.FromErr(err)
		}
	}

	if id, err := buildID(repoName, escapeIDPart(envName), name); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

	return resourceGithubActionsEnvironmentVariableRead(ctx, d, meta)
}

func resourceGithubActionsEnvironmentVariableRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName, envNamePart, name, err := parseID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	envName := unescapeIDPart(envNamePart)

	variable, _, err := client.Actions.GetEnvVariable(ctx, owner, repoName, url.PathEscape(envName), name)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing actions variable %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	_ = d.Set("repository", repoName)
	_ = d.Set("environment", envName)
	_ = d.Set("variable_name", name)
	_ = d.Set("value", variable.Value)
	_ = d.Set("created_at", variable.CreatedAt.String())
	_ = d.Set("updated_at", variable.UpdatedAt.String())

	return nil
}

func resourceGithubActionsEnvironmentVariableDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName, envNamePart, name, err := parseID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	envName := unescapeIDPart(envNamePart)

	_, err = client.Actions.DeleteEnvVariable(ctx, owner, repoName, url.PathEscape(envName), name)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsEnvironmentVariableImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	repoName, envNamePart, name, err := parseID3(d.Id())
	if err != nil {
		return nil, err
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("environment", unescapeIDPart(envNamePart)); err != nil {
		return nil, err
	}
	if err := d.Set("variable_name", name); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
