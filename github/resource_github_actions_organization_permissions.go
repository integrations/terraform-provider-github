package github

import (
	"context"
	"errors"
	"log"

	"github.com/google/go-github/v42/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubActionsOrganizationPermissions() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsOrganizationPermissionsCreateOrUpdate,
		Read:   resourceGithubActionsOrganizationPermissionsRead,
		Update: resourceGithubActionsOrganizationPermissionsCreateOrUpdate,
		Delete: resourceGithubActionsOrganizationPermissionsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"allowed_actions": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"all", "local_only", "selected"}, false),
			},
			"enabled_repositories": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"all", "none", "selected"}, false),
			},
			"allowed_actions_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"github_owned_allowed": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"patterns_allowed": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"verified_allowed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"enabled_repositories_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository_ids": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceGithubActionsOrganizationAllowedObject(d *schema.ResourceData) (*github.ActionsAllowed, error) {
	allowed := &github.ActionsAllowed{}

	config := d.Get("allowed_actions_config").([]interface{})
	if len(config) > 0 {
		data := config[0].(map[string]interface{})
		switch x := data["github_owned_allowed"].(type) {
		case bool:
			allowed.GithubOwnedAllowed = &x
		}

		switch x := data["verified_allowed"].(type) {
		case bool:
			allowed.VerifiedAllowed = &x
		}

		patternsAllowed := []string{}

		switch t := data["patterns_allowed"].(type) {
		case *schema.Set:
			for _, value := range t.List() {
				patternsAllowed = append(patternsAllowed, value.(string))
			}
		}

		allowed.PatternsAllowed = patternsAllowed
	} else {
		return &github.ActionsAllowed{},
			errors.New("The allowed_actions_config {} block must be specified if allowed_actions == 'selected'.")
	}

	return allowed, nil
}

func resourceGithubActionsEnabledRepositoriesObject(d *schema.ResourceData) ([]int64, error) {
	var enabled []int64

	config := d.Get("enabled_repositories_config").([]interface{})
	log.Printf("[help] length of config in actopms enabled is %v", len(config))
	if len(config) > 0 {
		data := config[0].(map[string]interface{})
		switch x := data["repository_ids"].(type) {
		case *schema.Set:
			for _, value := range x.List() {
				enabled = append(enabled, int64(value.(int)))
			}
		}
	} else {
		return nil, errors.New("The enabled_repositories_config {} block must be specified if enabled_repositories == 'selected'.")
	}
	return enabled, nil
}

func resourceGithubActionsOrganizationPermissionsCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	allowedActions := d.Get("allowed_actions").(string)
	enabledRepositories := d.Get("enabled_repositories").(string)

	_, _, err = client.Organizations.EditActionsPermissions(ctx,
		orgName,
		github.ActionsPermissions{
			AllowedActions:      &allowedActions,
			EnabledRepositories: &enabledRepositories,
		})
	if err != nil {
		return err
	}

	if allowedActions == "selected" {
		actionsAllowedData, err := resourceGithubActionsOrganizationAllowedObject(d)
		if err != nil {
			return err
		}
		_, _, err = client.Organizations.EditActionsAllowed(ctx,
			orgName,
			*actionsAllowedData)
		if err != nil {
			return err
		}
	}

	if enabledRepositories == "selected" {
		enabledReposData, err := resourceGithubActionsEnabledRepositoriesObject(d)
		if err != nil {
			return err
		}
		_, err = client.Actions.SetEnabledReposInOrg(ctx,
			orgName,
			enabledReposData)
		if err != nil {
			return err
		}
	}

	d.SetId(orgName)
	return resourceGithubActionsOrganizationPermissionsRead(d, meta)
}

func resourceGithubActionsOrganizationPermissionsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	actionsPermissions, _, err := client.Organizations.GetActionsPermissions(ctx, d.Id())
	if err != nil {
		return err
	}

	if actionsPermissions.GetAllowedActions() == "selected" {
		actionsAllowed, _, err := client.Organizations.GetActionsAllowed(ctx, d.Id())
		if err != nil {
			return err
		}

		// If actionsAllowed set to local/all by removing all actions config settings, the response will be empty
		if actionsAllowed != nil {
			d.Set("allowed_actions_config", []interface{}{
				map[string]interface{}{
					"github_owned_allowed": actionsAllowed.GetGithubOwnedAllowed(),
					"patterns_allowed":     actionsAllowed.PatternsAllowed,
					"verified_allowed":     actionsAllowed.GetVerifiedAllowed(),
				},
			})
		}
	} else {
		d.Set("allowed_actions_config", []interface{}{})
	}

	if actionsPermissions.GetEnabledRepositories() == "selected" {
		opts := github.ListOptions{PerPage: 10, Page: 1}
		var repoList []int64
		var allRepos []*github.Repository

		for {
			enabledRepos, resp, err := client.Actions.ListEnabledReposInOrg(ctx, d.Id(), &opts)
			if err != nil {
				return err
			}
			allRepos = append(allRepos, enabledRepos.Repositories...)

			opts.Page = resp.NextPage

			if resp.NextPage == 0 {
				break
			}
		}
		for index := range allRepos {
			repoList = append(repoList, *allRepos[index].ID)
		}
		if allRepos != nil {
			d.Set("enabled_repositories_config", []interface{}{
				map[string]interface{}{
					"repository_ids": repoList,
				},
			})
		} else {
			d.Set("enabled_repositories_config", []interface{}{})
		}
	}

	d.Set("allowed_actions", actionsPermissions.GetAllowedActions())
	d.Set("enabled_repositories", actionsPermissions.GetEnabledRepositories())

	return nil
}

func resourceGithubActionsOrganizationPermissionsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	// This will nullify any allowedActions elements
	_, _, err = client.Organizations.EditActionsPermissions(ctx,
		orgName,
		github.ActionsPermissions{
			AllowedActions:      github.String("all"),
			EnabledRepositories: github.String("all"),
		})
	if err != nil {
		return err
	}

	return nil
}
