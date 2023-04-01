package github

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v50/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubActionsVariable() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsVariableCreate,
		Read:   resourceGithubActionsVariableRead,
		Update: resourceGithubActionsVariableUpdate,
		Delete: resourceGithubActionsVariableDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
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

func resourceGithubActionsVariableCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	variable := &github.ActionsVariable{
		Name:  d.Get("variable_name").(string),
		Value: d.Get("value").(string),
	}

	_, err := client.Actions.CreateRepoVariable(ctx, owner, repo, variable)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repo, d.Get("variable_name").(string)))
	return resourceGithubActionsVariableRead(d, meta)
}

func resourceGithubActionsVariableUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	variable := &github.ActionsVariable{
		Name:  d.Get("variable_name").(string),
		Value: d.Get("value").(string),
	}

	_, err := client.Actions.UpdateRepoVariable(ctx, owner, repo, variable)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repo, d.Get("variable_name").(string)))
	return resourceGithubActionsVariableRead(d, meta)
}

func resourceGithubActionsVariableRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repoName, variableName, err := parseTwoPartID(d.Id(), "repository", "variable_name")
	if err != nil {
		return err
	}

	variable, _, err := client.Actions.GetRepoVariable(ctx, owner, repoName, variableName)
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
	d.Set("variable_name", variableName)
	d.Set("value", variable.Value)
	d.Set("created_at", variable.CreatedAt.String())
	d.Set("updated_at", variable.UpdatedAt.String())

	return nil
}

func resourceGithubActionsVariableDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	repoName, variableName, err := parseTwoPartID(d.Id(), "repository", "variable_name")
	if err != nil {
		return err
	}

	_, err = client.Actions.DeleteRepoVariable(ctx, orgName, repoName, variableName)

	return err
}
