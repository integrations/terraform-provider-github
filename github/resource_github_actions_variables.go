package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsVariables() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsVariablesCreate,
		Read:   resourceGithubActionsVariablesRead,
		Update: resourceGithubActionsVariablesUpdate,
		Delete: resourceGithubActionsVariablesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the repository.",
			},
			"variable": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of variables to manage.",
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					return schema.HashString(strings.ToUpper(m["name"].(string)))
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:             schema.TypeString,
							Required:         true,
							Description:      "Name of the variable.",
							ValidateDiagFunc: validateSecretNameFunc,
							DiffSuppressFunc: caseInsensitive(),
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Value of the variable.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date of variable creation.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date of variable update.",
						},
					},
				},
			},
		},
	}
}

type repoVariable struct {
	name      string
	value     string
	createdAt string
	updatedAt string
}

func (v repoVariable) Empty() bool {
	return v == repoVariable{}
}

func flattenRepoVariable(variable repoVariable) map[string]interface{} {
	if variable.Empty() {
		return nil
	}

	return map[string]interface{}{
		"name":       variable.name,
		"value":      variable.value,
		"created_at": variable.createdAt,
		"updated_at": variable.updatedAt,
	}
}

func flattenRepoVariables(variables []repoVariable) []interface{} {
	if variables == nil {
		return nil
	}

	// Sort variables by name for consistent ordering
	sort.SliceStable(variables, func(i, j int) bool {
		return variables[i].name < variables[j].name
	})

	result := make([]interface{}, len(variables))
	for i, variable := range variables {
		result[i] = flattenRepoVariable(variable)
	}

	return result
}

// List all repository variables
func listRepoVariables(client *github.Client, ctx context.Context, owner, repo string) ([]repoVariable, error) {
	options := github.ListOptions{
		PerPage: 100,
	}

	var allVariables []repoVariable
	for {
		variables, resp, err := client.Actions.ListRepoVariables(ctx, owner, repo, &options)
		if err != nil {
			return nil, err
		}

		for _, variable := range variables.Variables {
			allVariables = append(allVariables, repoVariable{
				name:      variable.Name,
				value:     variable.Value,
				createdAt: variable.CreatedAt.String(),
				updatedAt: variable.UpdatedAt.String(),
			})
		}

		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	return allVariables, nil
}

// Create or update variables to match desired state
func syncRepoVariables(ctx context.Context, client *github.Client, owner, repo string,
	wantVariables []interface{}, existingVariables []repoVariable) error {

	// Map of existing variables by name for easy lookup
	existingMap := make(map[string]repoVariable)
	for _, v := range existingVariables {
		existingMap[v.name] = v
	}

	// Track variables to create, update, or delete
	for _, v := range wantVariables {
		varConfig := v.(map[string]interface{})
		name := varConfig["name"].(string)
		value := varConfig["value"].(string)

		if existing, exists := existingMap[name]; exists {
			// Variable exists, check if value has changed
			if existing.value != value {
				// Update variable
				variable := &github.ActionsVariable{
					Name:  name,
					Value: value,
				}

				_, err := client.Actions.UpdateRepoVariable(ctx, owner, repo, variable)
				if err != nil {
					return fmt.Errorf("error updating repository variable %s: %v", name, err)
				}
				log.Printf("[DEBUG] Updated repository variable: %s", name)
			}

			// Remove from map to track what variables to keep
			delete(existingMap, name)
		} else {
			// Create new variable
			variable := &github.ActionsVariable{
				Name:  name,
				Value: value,
			}

			_, err := client.Actions.CreateRepoVariable(ctx, owner, repo, variable)
			if err != nil {
				return fmt.Errorf("error creating repository variable %s: %v", name, err)
			}
			log.Printf("[DEBUG] Created repository variable: %s", name)
		}
	}

	// Delete variables that are no longer in config
	for name := range existingMap {
		_, err := client.Actions.DeleteRepoVariable(ctx, owner, repo, name)
		if err != nil {
			return fmt.Errorf("error deleting repository variable %s: %v", name, err)
		}
		log.Printf("[DEBUG] Deleted repository variable: %s", name)
	}

	return nil
}

func resourceGithubActionsVariablesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	variables := d.Get("variable").(*schema.Set).List()

	// Check for 500 item variable limit
	if len(variables) > 500 {
		return fmt.Errorf("variable set cannot contain more than 500 items")
	}

	// Check for any duplicate variable names
	namesMap := make(map[string]struct{})
	for _, v := range variables {
		variableConfig := v.(map[string]interface{})
		name := strings.ToUpper(variableConfig["name"].(string))
		if _, exists := namesMap[name]; exists {
			return fmt.Errorf("duplicate variable name detected: %s", name)
		}
		namesMap[name] = struct{}{}
	}

	// List existing variables
	existingVariables, err := listRepoVariables(client, ctx, owner, repo)
	if err != nil {
		return err
	}

	// Sync variables (create, update, delete as needed)
	err = syncRepoVariables(ctx, client, owner, repo, variables, existingVariables)
	if err != nil {
		return err
	}

	d.SetId(repo)
	return resourceGithubActionsVariablesRead(d, meta)
}

func resourceGithubActionsVariablesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Id()

	variables, err := listRepoVariables(client, ctx, owner, repo)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing repository variables %s from state because the repository no longer exists in GitHub", d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("repository", repo)
	d.Set("variable", flattenRepoVariables(variables))

	return nil
}

func resourceGithubActionsVariablesUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	variables := d.Get("variable").(*schema.Set).List()

	// Check for 500 item variable limit
	if len(variables) > 500 {
		return fmt.Errorf("variable set cannot contain more than 500 items")
	}

	// Check for any duplicate variable names
	namesMap := make(map[string]struct{})
	for _, v := range variables {
		variableConfig := v.(map[string]interface{})
		name := strings.ToUpper(variableConfig["name"].(string))
		if _, exists := namesMap[name]; exists {
			return fmt.Errorf("duplicate variable name detected: %s", name)
		}
		namesMap[name] = struct{}{}
	}

	// List existing variables
	existingVariables, err := listRepoVariables(client, ctx, owner, repo)
	if err != nil {
		return err
	}

	// Sync variables (create, update, delete as needed)
	err = syncRepoVariables(ctx, client, owner, repo, variables, existingVariables)
	if err != nil {
		return err
	}

	return resourceGithubActionsVariablesRead(d, meta)
}

func resourceGithubActionsVariablesDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)

	// List all variables
	variables, err := listRepoVariables(client, ctx, owner, repo)
	if err != nil {
		return deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "repository variables (%s)", repo)
	}

	// Delete each variable
	for _, variable := range variables {
		_, err = client.Actions.DeleteRepoVariable(ctx, owner, repo, variable.name)
		if err != nil {
			return fmt.Errorf("error deleting repository variable %s: %v", variable.name, err)
		}
		log.Printf("[DEBUG] Deleted repository variable: %s", variable.name)
	}

	return nil
}
