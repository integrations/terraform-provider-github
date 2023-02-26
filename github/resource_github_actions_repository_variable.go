package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubActionsRepositoryVariable() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsRepositoryVariableCreate,
		Read:   resourceGithubActionsRepositoryVariableRead,
		Update: resourceGithubActionsRepositoryVariableUpdate,
		Delete: resourceGithubActionsRepositoryVariableDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the repository in which to create the variable.",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Name of the repository variable.",
				ValidateFunc: validateVariableName,
				StateFunc: func(in interface{}) string {
					// Names are always stored as upper case.
					return strings.ToUpper(in.(string))
				},
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Value of the repository variable.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Timestamp of when the repository variable was created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Timestamp of when the repository variable was last updated.",
			},
		},
	}
}

func resourceGithubActionsRepositoryVariableCreate(d *schema.ResourceData, meta interface{}) error {
	ctx := context.Background()

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repo := d.Get("repository").(string)
	name := d.Get("name").(string)
	value := d.Get("value").(string)

	if _, err := client.Actions.CreateRepoVariable(ctx, owner, repo, &github.ActionsVariable{
		Name:  name,
		Value: value,
	}); err != nil {
		return fmt.Errorf("failed to create repository variable %s: %w", name, err)
	}

	d.SetId(buildTwoPartID(repo, name))
	return resourceGithubActionsRepositoryVariableRead(d, meta)
}

func resourceGithubActionsRepositoryVariableRead(d *schema.ResourceData, meta interface{}) error {
	ctx := context.Background()

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repo, name, err := parseTwoPartID(d.Id(), "repository", "name")
	if err != nil {
		return err
	}

	variable, _, err := client.Actions.GetRepoVariable(ctx, owner, repo, name)

	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing actions variable %q from state because it no longer exists on GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return fmt.Errorf("failed to lookup repo variable %s: %w", name, err)
	}

	d.Set("repository", repo)
	d.Set("name", variable.Name)
	d.Set("value", variable.Value)
	d.Set("created_at", variable.CreatedAt.String())
	d.Set("updated_at", variable.UpdatedAt.String())

	return nil
}

func resourceGithubActionsRepositoryVariableUpdate(d *schema.ResourceData, meta interface{}) error {
	ctx := context.Background()

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repo := d.Get("repository").(string)
	name := d.Get("name").(string)
	value := d.Get("value").(string)

	if d.HasChange("repository") || d.HasChange("name") || d.HasChange("value") {
		if _, err := client.Actions.UpdateRepoVariable(ctx, owner, repo, &github.ActionsVariable{
			Name:  name,
			Value: value,
		}); err != nil {
			return fmt.Errorf("failed to update repository variable %s: %w", name, err)
		}

		d.SetId(buildTwoPartID(repo, name))
	}

	return resourceGithubActionsRepositoryVariableRead(d, meta)
}

func resourceGithubActionsRepositoryVariableDelete(d *schema.ResourceData, meta interface{}) error {
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repo, name, err := parseTwoPartID(d.Id(), "repository", "name")
	if err != nil {
		return err
	}

	if _, err := client.Actions.DeleteRepoVariable(ctx, owner, repo, name); err != nil {
		return fmt.Errorf("failed to delete repo variable %s: %w", name, err)
	}
	return nil
}
