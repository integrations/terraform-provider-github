package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubBranchProtectionRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubBranchProtectionRulesRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pattern": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubBranchProtectionRulesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v4client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	var query struct {
		Repository struct {
			ID                    githubv4.String
			BranchProtectionRules struct {
				Nodes []struct {
					Pattern githubv4.String
				}
				PageInfo PageInfo
			} `graphql:"branchProtectionRules(first:$first, after:$cursor)"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}
	variables := map[string]any{
		"first":  githubv4.Int(100),
		"name":   githubv4.String(repoName),
		"owner":  githubv4.String(orgName),
		"cursor": (*githubv4.String)(nil),
	}

	var rules []any
	for {
		err := client.Query(meta.(*Owner).StopContext, &query, variables)
		if err != nil {
			return err
		}

		additionalRules := make([]any, len(query.Repository.BranchProtectionRules.Nodes))
		for i, rule := range query.Repository.BranchProtectionRules.Nodes {
			r := make(map[string]any)
			r["pattern"] = rule.Pattern
			additionalRules[i] = r
		}
		rules = append(rules, additionalRules...)

		if !query.Repository.BranchProtectionRules.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Repository.BranchProtectionRules.PageInfo.EndCursor)
	}

	d.SetId(string(query.Repository.ID))
	err := d.Set("rules", rules)
	if err != nil {
		return err
	}

	return nil
}
