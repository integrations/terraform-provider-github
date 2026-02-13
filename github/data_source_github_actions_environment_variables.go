package github

import (
	"context"
	"net/url"

	"github.com/google/go-github/v82/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsEnvironmentVariables() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsEnvironmentVariablesRead,

		Schema: map[string]*schema.Schema{
			"full_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"name"},
				Description:   "Full name of the repository (in org/name format).",
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"full_name"},
				Description:   "The name of the repository.",
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the environment.",
			},
			"variables": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of variables for the environment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the variable.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value of the variable.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of the variable creation.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of the variable last update.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubActionsEnvironmentVariablesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	var repoName string

	envName := d.Get("environment").(string)

	if fullName, ok := d.GetOk("full_name"); ok {
		var err error
		owner, repoName, err = splitRepoFullName(fullName.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if name, ok := d.GetOk("name"); ok {
		repoName = name.(string)
	}

	if repoName == "" {
		return diag.Errorf("one of %q or %q has to be provided", "full_name", "name")
	}

	options := github.ListOptions{
		PerPage: maxPerPage,
	}

	var all_variables []map[string]string
	for {
		variables, resp, err := client.Actions.ListEnvVariables(ctx, owner, repoName, url.PathEscape(envName), &options)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, variable := range variables.Variables {
			new_variable := map[string]string{
				"name":       variable.Name,
				"value":      variable.Value,
				"created_at": variable.CreatedAt.String(),
				"updated_at": variable.UpdatedAt.String(),
			}
			all_variables = append(all_variables, new_variable)
		}
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	id, err := buildID(repoName, escapeIDPart(envName))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("variables", all_variables); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
