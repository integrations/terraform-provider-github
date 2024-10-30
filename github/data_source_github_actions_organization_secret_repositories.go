package github

import (
	"context"

	"github.com/google/go-github/v57/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationSecretRepositories() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsOrganizationSecretRepositoriesRead,

		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repositories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"full_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubActionsOrganizationSecretRepositoriesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	name := d.Get("secret_name").(string)

	options := github.ListOptions{
		PerPage: 100,
	}

	var all_repositories []map[string]string
	for {
		repositories, resp, err := client.Actions.ListSelectedReposForOrgSecret(context.TODO(), owner, name, &options)
		if err != nil {
			return err
		}
		for _, repository := range repositories.Repositories {
			newRepository := map[string]string{
				"name":      *repository.Name,
				"full_name": *repository.FullName,
			}
			all_repositories = append(all_repositories, newRepository)

		}
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	d.SetId(owner)
	err := d.Set("repositories", all_repositories)
	if err != nil {
		return err
	}

	return nil
}
