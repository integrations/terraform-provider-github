package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsEnvironmentVariable() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubActionsEnvironmentVariableV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubActionsEnvironmentVariableStateUpgradeV0,
				Version: 0,
			},
		},

		CustomizeDiff: resourceGithubActionsEnvironmentVariableDiff,
		CreateContext: resourceGithubActionsEnvironmentVariableCreate,
		ReadContext:   resourceGithubActionsEnvironmentVariableRead,
		UpdateContext: resourceGithubActionsEnvironmentVariableUpdate,
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

func resourceGithubActionsEnvironmentVariableDiff(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	if len(diff.Id()) == 0 {
		return nil
	}

	if diff.HasChange("repository") {
		repoIDString, _, _, _, err := parseID4(diff.Id())
		if err != nil {
			return err
		}

		repoID, err := strconv.Atoi(repoIDString)
		if err != nil {
			return fmt.Errorf("failed to convert repository ID %s to integer: %w", repoIDString, err)
		}

		repoName := diff.Get("repository").(string)

		repo, _, err := client.Repositories.Get(ctx, owner, repoName)
		if err != nil {
			var ghErr *github.ErrorResponse
			if errors.As(err, &ghErr) {
				if ghErr.Response.StatusCode != http.StatusNotFound {
					return err
				}

				log.Printf("[INFO] Repository %s not found when checking repository change for actions environment variable %s", repoName, diff.Id())
			} else {
				return err
			}
		} else {
			log.Printf("[INFO] Repository %s found when checking repository change for actions environment variable %s", repoName, diff.Id())

			if repoID != int(repo.GetID()) {
				return diff.ForceNew("repository")
			}
		}
	}

	return nil
}

func resourceGithubActionsEnvironmentVariableCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	varName := d.Get("variable_name").(string)

	escapedEnvName := url.PathEscape(envName)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	repoID := int(repo.GetID())

	_, err = client.Actions.CreateEnvVariable(ctx, owner, repoName, escapedEnvName, &github.ActionsVariable{
		Name:  varName,
		Value: d.Get("value").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if id, err := buildID(strconv.Itoa(repoID), repoName, escapeIDPart(envName), varName); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

	// GitHub API does not return on create so we have to lookup the variable to get timestamps
	if variable, _, err := client.Actions.GetEnvVariable(ctx, owner, repoName, escapedEnvName, varName); err == nil {
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
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	varName := d.Get("variable_name").(string)

	variable, _, err := client.Actions.GetEnvVariable(ctx, owner, repoName, url.PathEscape(envName), varName)
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

func resourceGithubActionsEnvironmentVariableUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoIDString, _, _, _, err := parseID4(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	repoID, err := strconv.Atoi(repoIDString)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to convert repository ID %s to integer: %w", repoIDString, err))
	}

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	varName := d.Get("variable_name").(string)

	escapedEnvName := url.PathEscape(envName)

	_, err = client.Actions.UpdateEnvVariable(ctx, owner, repoName, escapedEnvName, &github.ActionsVariable{
		Name:  varName,
		Value: d.Get("value").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if id, err := buildID(strconv.Itoa(repoID), repoName, escapeIDPart(envName), varName); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

	// GitHub API does not return on create so we have to lookup the variable to get timestamps
	if variable, _, err := client.Actions.GetEnvVariable(ctx, owner, repoName, escapedEnvName, varName); err == nil {
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

func resourceGithubActionsEnvironmentVariableDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	_, repoName, envNamePart, varName, err := parseID4(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Actions.DeleteEnvVariable(ctx, owner, repoName, unescapeIDPart(envNamePart), varName)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsEnvironmentVariableImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
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

	if id, err := buildID(strconv.Itoa(repoID), repoName, envNamePart, varName); err != nil {
		return nil, err
	} else {
		d.SetId(id)
	}

	if err := d.Set("repository", repoName); err != nil {
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
