package github

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v50/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubActionsEnvironmentVariable() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsEnvironmentVariableCreate,
		Read:   resourceGithubActionsEnvironmentVariableRead,
		Update: resourceGithubActionsEnvironmentVariableUpdate,
		Delete: resourceGithubActionsEnvironmentVariableDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Name of the variable.",
				ValidateFunc: validateSecretNameFunc,
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

func resourceGithubActionsEnvironmentVariableCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repoName := d.Get("repository").(string)
	env := d.Get("environment").(string)
	name := d.Get("variable_name").(string)

	variable := &github.ActionsVariable{
		Name:  name,
		Value: d.Get("value").(string),
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return err
	}

	_, err = client.Actions.CreateEnvVariable(ctx, int(repo.GetID()), env, variable)
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(repoName, env, name))
	return resourceGithubActionsEnvironmentVariableRead(d, meta)
}

func resourceGithubActionsEnvironmentVariableUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repoName := d.Get("repository").(string)
	env := d.Get("environment").(string)
	name := d.Get("variable_name").(string)

	variable := &github.ActionsVariable{
		Name:  name,
		Value: d.Get("value").(string),
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return err
	}
	_, err = client.Actions.UpdateEnvVariable(ctx, int(repo.GetID()), env, variable)
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(repoName, env, name))
	return resourceGithubActionsEnvironmentVariableRead(d, meta)
}

func resourceGithubActionsEnvironmentVariableRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repoName, env, name, err := parseThreePartID(d.Id(), "repository", "environment", "variable_name")
	if err != nil {
		return err
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return err
	}

	variable, _, err := client.Actions.GetEnvVariable(ctx, int(repo.GetID()), env, name)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing actions variable %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("repository", repoName)
	d.Set("environment", env)
	d.Set("variable_name", name)
	d.Set("value", variable.Value)
	d.Set("created_at", variable.CreatedAt.String())
	d.Set("updated_at", variable.UpdatedAt.String())

	return nil
}

func resourceGithubActionsEnvironmentVariableDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	repoName, env, name, err := parseThreePartID(d.Id(), "repository", "environment", "variable_name")
	if err != nil {
		return err
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return err
	}

	_, err = client.Actions.DeleteEnvVariable(ctx, int(repo.GetID()), env, name)

	return err
}
