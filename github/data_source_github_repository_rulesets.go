package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryRulesets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoryRulesetsRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository to get rulesets from.",
			},
			"rulesets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of rulesets for the repository.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ruleset_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "GitHub ID for the ruleset.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the ruleset.",
						},
						"target": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Possible values are `branch`, `push`, and `tag`.",
						},
						"enforcement": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Possible values for Enforcement are `disabled`, `active`, `evaluate`.",
						},
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GraphQL global node id for use with v4 API.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryRulesetsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner, _ := meta.(*Owner)
	client := owner.v3client
	ownerName := owner.name

	repoName, ok := d.Get("repository").(string)
	if !ok {
		return diag.Errorf(`expected "repository" to be string`)
	}

	opts := &github.RepositoryListRulesetsOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var allRulesets []*github.RepositoryRuleset
	for {
		rulesets, resp, err := client.Repositories.GetAllRulesets(ctx, ownerName, repoName, opts)
		if err != nil {
			return diag.FromErr(err)
		}
		allRulesets = append(allRulesets, rulesets...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	result := make([]map[string]any, len(allRulesets))
	for i, rs := range allRulesets {
		result[i] = map[string]any{
			"ruleset_id":  rs.GetID(),
			"name":        rs.GetName(),
			"target":      rs.GetTarget(),
			"enforcement": rs.GetEnforcement(),
			"node_id":     rs.GetNodeID(),
		}
	}

	d.SetId(repoName + ":" + strconv.Itoa(len(allRulesets)))

	if err := d.Set("rulesets", result); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
