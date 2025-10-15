package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/google/go-github/v66/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsEnvironmentVariables() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsEnvironmentVariablesCreate,
		Read:   resourceGithubActionsEnvironmentVariablesRead,
		Update: resourceGithubActionsEnvironmentVariablesUpdate,
		Delete: resourceGithubActionsEnvironmentVariablesDelete,
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
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the environment.",
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

type environmentVariable struct {
	name      string
	value     string
	createdAt string
	updatedAt string
}

func (v environmentVariable) Empty() bool {
	return v == environmentVariable{}
}

func flattenEnvironmentVariable(variable environmentVariable) map[string]interface{} {
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

func flattenEnvironmentVariables(variables []environmentVariable) []interface{} {
	if variables == nil {
		return nil
	}

	// Sort variables by name for consistent ordering
	sort.SliceStable(variables, func(i, j int) bool {
		return variables[i].name < variables[j].name
	})

	result := make([]interface{}, len(variables))
	for i, variable := range variables {
		result[i] = flattenEnvironmentVariable(variable)
	}

	return result
}

// List all environment variables for a repository environment
func listEnvironmentVariables(client *github.Client, ctx context.Context, owner, repo, envName string) ([]environmentVariable, error) {
	escapedEnvName := url.PathEscape(envName)
	options := github.ListOptions{
		PerPage: 100,
	}

	var allVariables []environmentVariable
	for {
		variables, resp, err := client.Actions.ListEnvVariables(ctx, owner, repo, escapedEnvName, &options)
		if err != nil {
			return nil, err
		}

		for _, variable := range variables.Variables {
			allVariables = append(allVariables, environmentVariable{
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
func syncEnvironmentVariables(ctx context.Context, client *github.Client, owner, repo, envName string,
	wantVariables []interface{}, existingVariables []environmentVariable) error {

	escapedEnvName := url.PathEscape(envName)

	// Map of existing variables by name for easy lookup
	existingMap := make(map[string]environmentVariable)
	for _, v := range existingVariables {
		existingMap[v.name] = v
	}

	// Track variables to create, update, or delete
	for _, v := range wantVariables {
		varConfig := v.(map[string]interface{})
		name := strings.ToUpper(varConfig["name"].(string))
		value := varConfig["value"].(string)

		if existing, exists := existingMap[name]; exists {
			// Variable exists, check if value has changed
			if existing.value != value {
				// Update variable
				variable := &github.ActionsVariable{
					Name:  name,
					Value: value,
				}

				_, err := client.Actions.UpdateEnvVariable(ctx, owner, repo, escapedEnvName, variable)
				if err != nil {
					return fmt.Errorf("error updating environment variable %s: %v", name, err)
				}
				log.Printf("[DEBUG] Updated environment variable: %s", name)
			}

			// Remove from map to track what variables to keep
			delete(existingMap, name)
		} else {
			// Create new variable
			variable := &github.ActionsVariable{
				Name:  name,
				Value: value,
			}

			_, err := client.Actions.CreateEnvVariable(ctx, owner, repo, escapedEnvName, variable)
			if err != nil {
				return fmt.Errorf("error creating environment variable %s: %v", name, err)
			}
			log.Printf("[DEBUG] Created environment variable: %s", name)
		}
	}

	// Delete variables that are no longer in config
	for name := range existingMap {
		_, err := client.Actions.DeleteEnvVariable(ctx, owner, repo, escapedEnvName, name)
		if err != nil {
			return fmt.Errorf("error deleting environment variable %s: %v", name, err)
		}
		log.Printf("[DEBUG] Deleted environment variable: %s", name)
	}

	return nil
}

func resourceGithubActionsEnvironmentVariablesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	variables := d.Get("variable").(*schema.Set).List()

	// Check for 100 item environment variable limit
	if len(variables) > 100 {
		return fmt.Errorf("environment variable set cannot contain more than 100 items")
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
	existingVariables, err := listEnvironmentVariables(client, ctx, owner, repo, envName)
	if err != nil {
		return err
	}

	// Sync variables (create, update, delete as needed)
	err = syncEnvironmentVariables(ctx, client, owner, repo, envName, variables, existingVariables)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repo, envName))
	return resourceGithubActionsEnvironmentVariablesRead(d, meta)
}

func resourceGithubActionsEnvironmentVariablesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo, envName, err := parseTwoPartID(d.Id(), "repository", "environment")
	if err != nil {
		return err
	}

	variables, err := listEnvironmentVariables(client, ctx, owner, repo, envName)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing environment variables %s from state because the environment or repository no longer exists in GitHub", d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("repository", repo)
	d.Set("environment", envName)
	d.Set("variable", flattenEnvironmentVariables(variables))

	return nil
}

func resourceGithubActionsEnvironmentVariablesUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	variables := d.Get("variable").(*schema.Set).List()

	// Check for 100 item environment variable limit
	if len(variables) > 100 {
		return fmt.Errorf("environment variable set cannot contain more than 100 items")
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
	existingVariables, err := listEnvironmentVariables(client, ctx, owner, repo, envName)
	if err != nil {
		return err
	}

	// Sync variables (create, update, delete as needed)
	err = syncEnvironmentVariables(ctx, client, owner, repo, envName, variables, existingVariables)
	if err != nil {
		return err
	}

	return resourceGithubActionsEnvironmentVariablesRead(d, meta)
}

func resourceGithubActionsEnvironmentVariablesDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo, envName, err := parseTwoPartID(d.Id(), "repository", "environment")
	if err != nil {
		return err
	}

	escapedEnvName := url.PathEscape(envName)

	// List all variables
	variables, err := listEnvironmentVariables(client, ctx, owner, repo, envName)
	if err != nil {
		return deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "environment variables (%s/%s)", repo, envName)
	}

	// Delete each variable
	for _, variable := range variables {
		_, err = client.Actions.DeleteEnvVariable(ctx, owner, repo, escapedEnvName, variable.name)
		if err != nil {
			return fmt.Errorf("error deleting environment variable %s: %v", variable.name, err)
		}
		log.Printf("[DEBUG] Deleted environment variable: %s", variable.name)
	}

	return nil
}
