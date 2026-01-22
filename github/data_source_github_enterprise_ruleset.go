package github

import (
	"context"
	"errors"
	"fmt"
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
						"repository_property": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Conditions for repository properties that the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of repository property patterns to include.",
									},
									"exclude": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of repository property patterns to exclude.",
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

	tflog.Trace(ctx, fmt.Sprintf("Reading enterprise ruleset: %s/%d", enterpriseSlug, rulesetID), map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
	})

	ruleset, resp, err := client.Enterprise.GetRepositoryRuleset(ctx, enterpriseSlug, rulesetID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Error(ctx, fmt.Sprintf("Enterprise ruleset not found: %s/%d", enterpriseSlug, rulesetID), map[string]any{
					"enterprise_slug": enterpriseSlug,
					"ruleset_id":      rulesetID,
				})
				return diag.Errorf("enterprise ruleset %d not found in enterprise %s", rulesetID, enterpriseSlug)
			}
		}
		tflog.Error(ctx, fmt.Sprintf("Failed to read enterprise ruleset: %s/%d", enterpriseSlug, rulesetID), map[string]any{
			"enterprise_slug": enterpriseSlug,
			"ruleset_id":      rulesetID,
			"error":           err.Error(),
		})
		return diag.FromErr(err)
	}

	// Set the ID to the ruleset ID
	d.SetId(strconv.FormatInt(ruleset.GetID(), 10))

	// Set all computed attributes
	_ = d.Set("ruleset_id", ruleset.ID)
	_ = d.Set("name", ruleset.Name)
	_ = d.Set("target", ruleset.GetTarget())
	_ = d.Set("enforcement", ruleset.Enforcement)
	_ = d.Set("bypass_actors", flattenBypassActors(ruleset.BypassActors))
	_ = d.Set("conditions", flattenConditions(ruleset.GetConditions(), true))
	_ = d.Set("rules", flattenRules(ruleset.Rules, true))
	_ = d.Set("node_id", ruleset.GetNodeID())
	_ = d.Set("etag", resp.Header.Get("ETag"))

	tflog.Trace(ctx, fmt.Sprintf("Successfully read enterprise ruleset: %s/%d", enterpriseSlug, rulesetID), map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
		"name":            ruleset.Name,
	})

	return nil
}
