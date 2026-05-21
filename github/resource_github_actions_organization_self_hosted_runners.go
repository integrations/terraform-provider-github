package github

import (
	"context"
	"errors"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsOrganizationSelfHostedRunners() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsOrganizationSelfHostedRunnersCreateOrUpdate,
		Read:   resourceGithubActionsOrganizationSelfHostedRunnersRead,
		Update: resourceGithubActionsOrganizationSelfHostedRunnersCreateOrUpdate,
		Delete: resourceGithubActionsOrganizationSelfHostedRunnersDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enabled_repositories": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The policy that controls which repositories in the organization can create self-hosted runners. Can be one of: 'all', 'selected', or 'none'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "selected", "none"}, false)),
			},
			"enabled_repositories_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Sets the list of selected repositories that are allowed to create self-hosted runners. Only available when 'enabled_repositories' = 'selected'.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository_ids": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "List of repository IDs allowed to create self-hosted runners.",
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
					},
				},
			},
		},
	}
}

func resourceGithubActionsOrganizationSelfHostedRunnersEnabledRepos(d *schema.ResourceData) ([]int64, error) {
	var repoIDs []int64

	config := d.Get("enabled_repositories_config").([]any)
	if len(config) > 0 {
		data := config[0].(map[string]any)
		switch x := data["repository_ids"].(type) {
		case *schema.Set:
			for _, value := range x.List() {
				repoIDs = append(repoIDs, int64(value.(int)))
			}
		}
	} else {
		return nil, errors.New("the enabled_repositories_config {} block must be specified if enabled_repositories == 'selected'")
	}
	return repoIDs, nil
}

func resourceGithubActionsOrganizationSelfHostedRunnersCreateOrUpdate(d *schema.ResourceData, meta any) error {
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

	enabledRepositories := d.Get("enabled_repositories").(string)

	_, err = client.Actions.UpdateSelfHostedRunnersSettingsInOrganization(ctx,
		orgName,
		github.SelfHostedRunnersSettingsOrganizationOpt{
			EnabledRepositories: &enabledRepositories,
		})
	if err != nil {
		return err
	}

	if enabledRepositories == "selected" {
		repoIDs, err := resourceGithubActionsOrganizationSelfHostedRunnersEnabledRepos(d)
		if err != nil {
			return err
		}
		_, err = client.Actions.SetRepositoriesSelfHostedRunnersAllowedInOrganization(ctx,
			orgName,
			repoIDs)
		if err != nil {
			return err
		}
	}

	d.SetId(orgName)
	return resourceGithubActionsOrganizationSelfHostedRunnersRead(d, meta)
}

func resourceGithubActionsOrganizationSelfHostedRunnersRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	settings, _, err := client.Actions.GetSelfHostedRunnersSettingsInOrganization(ctx, d.Id())
	if err != nil {
		return err
	}

	if err = d.Set("enabled_repositories", settings.GetEnabledRepositories()); err != nil {
		return err
	}

	if settings.GetEnabledRepositories() == "selected" {
		opts := github.ListOptions{PerPage: 10, Page: 1}
		var repoIDs []int64
		var allRepos []*github.Repository

		for {
			result, resp, err := client.Actions.ListRepositoriesSelfHostedRunnersAllowedInOrganization(ctx, d.Id(), &opts)
			if err != nil {
				return err
			}
			allRepos = append(allRepos, result.Repositories...)

			opts.Page = resp.NextPage

			if resp.NextPage == 0 {
				break
			}
		}
		for index := range allRepos {
			repoIDs = append(repoIDs, *allRepos[index].ID)
		}
		if allRepos != nil {
			if err = d.Set("enabled_repositories_config", []any{
				map[string]any{
					"repository_ids": repoIDs,
				},
			}); err != nil {
				return err
			}
		} else {
			if err = d.Set("enabled_repositories_config", []any{}); err != nil {
				return err
			}
		}
	} else {
		if err = d.Set("enabled_repositories_config", []any{}); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubActionsOrganizationSelfHostedRunnersDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	allPolicy := "all"
	_, err = client.Actions.UpdateSelfHostedRunnersSettingsInOrganization(ctx,
		orgName,
		github.SelfHostedRunnersSettingsOrganizationOpt{
			EnabledRepositories: &allPolicy,
		})
	if err != nil {
		return err
	}

	return nil
}
