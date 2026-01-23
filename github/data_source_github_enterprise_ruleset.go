package github

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseRuleset() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseRulesetRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"ruleset_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the ruleset to retrieve.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the ruleset.",
			},
			"target": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target of the ruleset (branch, tag, or push).",
			},
			"enforcement": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The enforcement level of the ruleset (disabled, active, or evaluate).",
			},
			"bypass_actors": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The actors that can bypass the rules in this ruleset.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"actor_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the actor that can bypass a ruleset.",
						},
						"actor_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of actor that can bypass a ruleset.",
						},
						"bypass_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When the specified actor can bypass the ruleset.",
						},
					},
				},
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GraphQL global node id for use with v4 API.",
			},
			"conditions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Parameters for an enterprise ruleset condition.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organization_name": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Conditions for organization names that the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of organization name patterns to include.",
									},
									"exclude": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of organization name patterns to exclude.",
									},
								},
							},
						},
						"organization_id": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Conditions for organization IDs that the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"organization_ids": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeInt},
										Description: "Array of organization IDs to target.",
									},
								},
							},
						},
						"repository_name": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Conditions for repository names that the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of repository name patterns to include.",
									},
									"exclude": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of repository name patterns to exclude.",
									},
									"protected": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to target only protected repositories.",
									},
								},
							},
						},
						"repository_id": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Conditions for repository IDs that the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"repository_ids": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeInt},
										Description: "Array of repository IDs to target.",
									},
								},
							},
						},
						"ref_name": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Conditions for ref names (branches or tags) that the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of ref name patterns to include.",
									},
									"exclude": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of ref name patterns to exclude.",
									},
								},
							},
						},
					},
				},
			},
			"rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Rules for the ruleset.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creation": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Only allow users with bypass permission to create matching refs.",
						},
						"update": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Only allow users with bypass permission to update matching refs.",
						},
						"deletion": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Only allow users with bypass permissions to delete matching refs.",
						},
						"required_linear_history": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Prevent merge commits from being pushed to matching branches.",
						},
						"required_signatures": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Commits pushed to matching branches must have verified signatures.",
						},
						"merge_queue": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Merges must be performed via a merge queue.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"check_response_timeout_minutes": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Maximum time for a required status check to report a conclusion.",
									},
									"grouping_strategy": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "When set to ALLGREEN, the merge commit created by merge queue for each PR in the group must pass all required checks to merge. When set to HEADGREEN, only the commit at the head of the merge group must pass its required checks to merge.",
									},
									"max_entries_to_build": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Limit the number of queued pull requests requesting checks and workflow runs at the same time.",
									},
									"max_entries_to_merge": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum number of PRs that will be merged together in a group.",
									},
									"merge_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Method to use when merging changes from queued pull requests.",
									},
									"min_entries_to_merge": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum number of PRs that will be merged together in a group.",
									},
									"min_entries_to_merge_wait_minutes": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The time merge queue should wait after the first PR is added to the queue for the minimum group size to be met.",
									},
								},
							},
						},
						"required_deployments": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_deployment_environments": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The environments that must be successfully deployed to before branches can be merged.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"non_fast_forward": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Prevent users with push access from force pushing to branches.",
						},
						"pull_request": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Require all commits be made to a non-target branch and submitted via a pull request before they can be merged.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dismiss_stale_reviews_on_push": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "New, reviewable commits pushed will dismiss previous pull request review approvals.",
									},
									"require_code_owner_review": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Require an approving review in pull requests that modify files that have a designated code owner.",
									},
									"require_last_push_approval": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the most recent reviewable push must be approved by someone other than the person who pushed it.",
									},
									"required_approving_review_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of approving reviews that are required before a pull request can be merged.",
									},
									"required_review_thread_resolution": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "All conversations on code must be resolved before a pull request can be merged.",
									},
								},
							},
						},
						"required_status_checks": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Choose which status checks must pass before branches can be merged into a branch that matches this rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"strict_required_status_checks_policy": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether pull requests targeting a matching branch must be tested with the latest code.",
									},
									"do_not_enforce_on_create": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allow repositories and branches to be created if a check would otherwise prohibit it.",
									},
								},
							},
						},
						"required_workflows": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Choose which Actions workflows must pass before branches can be merged into a branch that matches this rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"do_not_enforce_on_create": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allow repositories and branches to be created if a check would otherwise prohibit it.",
									},
								},
							},
						},
						// Repository target rules (only valid when target = "repository")
						"repository_creation": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Only allow users with bypass permission to create repositories. Only valid for `repository` target.",
						},
						"repository_deletion": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Only allow users with bypass permission to delete repositories. Only valid for `repository` target.",
						},
						"repository_transfer": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Only allow users with bypass permission to transfer repositories. Only valid for `repository` target.",
						},
						"repository_name": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Restrict repository names to match specified patterns. Only valid for `repository` target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If true, the rule will fail if the pattern matches.",
									},
									"pattern": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The pattern to match repository names against.",
									},
								},
							},
						},
						"repository_visibility": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Restrict repository visibility changes. Only valid for `repository` target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"internal": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allow internal visibility for repositories.",
									},
									"private": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allow private visibility for repositories.",
									},
								},
							},
						},
					},
				},
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ETag of the ruleset for conditional updates.",
			},
		},
	}
}

func dataSourceGithubEnterpriseRulesetRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	rulesetID := int64(d.Get("ruleset_id").(int))

	tflog.Trace(ctx, "Reading enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
	})

	ruleset, resp, err := client.Enterprise.GetRepositoryRuleset(ctx, enterpriseSlug, rulesetID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Error(ctx, "Enterprise ruleset not found", map[string]any{
					"enterprise_slug": enterpriseSlug,
					"ruleset_id":      rulesetID,
				})
				return diag.Errorf("enterprise ruleset %d not found in enterprise %s", rulesetID, enterpriseSlug)
			}
		}
		tflog.Error(ctx, "Failed to read enterprise ruleset", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"ruleset_id":      rulesetID,
			"error":           err.Error(),
		})
		return diag.FromErr(err)
	}

	// Set the ID to the ruleset ID
	d.SetId(strconv.FormatInt(ruleset.GetID(), 10))

	// Set all computed attributes
	if err := d.Set("ruleset_id", ruleset.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", ruleset.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target", ruleset.GetTarget()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enforcement", ruleset.Enforcement); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("bypass_actors", flattenBypassActors(ruleset.BypassActors)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("conditions", flattenConditions(ruleset.GetConditions(), true)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rules", flattenRules(ruleset.Rules, true)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("node_id", ruleset.GetNodeID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Successfully read enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
		"name":            ruleset.Name,
	})

	return nil
}
