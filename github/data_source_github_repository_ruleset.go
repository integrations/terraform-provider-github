package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryRuleset() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoryRulesetRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository to get the ruleset from.",
			},
			"ruleset_id": {
				Type:        schema.TypeInt,
				Required:    true,
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
			"conditions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Parameters for a repository ruleset ref name condition.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ref_name": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Array of ref names or patterns to include.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"exclude": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Array of ref names or patterns to exclude.",
										Elem:        &schema.Schema{Type: schema.TypeString},
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
				Description: "Rules within the ruleset.",
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
						"update_allows_fetch_and_merge": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Branch can pull changes from its upstream repository.",
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
						"required_deployments": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Environments that must be successfully deployed to before branches can be merged.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_deployment_environments": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"pull_request": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Require all commits be made via a pull request before they can be merged.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_merge_methods": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"dismiss_stale_reviews_on_push": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"require_code_owner_review": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"require_last_push_approval": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"required_approving_review_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"required_review_thread_resolution": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"required_reviewers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"reviewer": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"file_patterns": {
													Type:     schema.TypeList,
													Computed: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"minimum_approvals": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"required_status_checks": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Status checks that must pass before branches can be merged.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_check": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"context": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"integration_id": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"strict_required_status_checks_policy": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"do_not_enforce_on_create": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"merge_queue": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Merges must be performed via a merge queue.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"check_response_timeout_minutes": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"grouping_strategy": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_entries_to_build": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"max_entries_to_merge": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"merge_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"min_entries_to_merge": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"min_entries_to_merge_wait_minutes": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"commit_message_pattern": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name":     {Type: schema.TypeString, Computed: true},
									"negate":   {Type: schema.TypeBool, Computed: true},
									"operator": {Type: schema.TypeString, Computed: true},
									"pattern":  {Type: schema.TypeString, Computed: true},
								},
							},
						},
						"commit_author_email_pattern": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name":     {Type: schema.TypeString, Computed: true},
									"negate":   {Type: schema.TypeBool, Computed: true},
									"operator": {Type: schema.TypeString, Computed: true},
									"pattern":  {Type: schema.TypeString, Computed: true},
								},
							},
						},
						"committer_email_pattern": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name":     {Type: schema.TypeString, Computed: true},
									"negate":   {Type: schema.TypeBool, Computed: true},
									"operator": {Type: schema.TypeString, Computed: true},
									"pattern":  {Type: schema.TypeString, Computed: true},
								},
							},
						},
						"branch_name_pattern": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name":     {Type: schema.TypeString, Computed: true},
									"negate":   {Type: schema.TypeBool, Computed: true},
									"operator": {Type: schema.TypeString, Computed: true},
									"pattern":  {Type: schema.TypeString, Computed: true},
								},
							},
						},
						"tag_name_pattern": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name":     {Type: schema.TypeString, Computed: true},
									"negate":   {Type: schema.TypeBool, Computed: true},
									"operator": {Type: schema.TypeString, Computed: true},
									"pattern":  {Type: schema.TypeString, Computed: true},
								},
							},
						},
						"required_code_scanning": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_code_scanning_tool": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alerts_threshold":          {Type: schema.TypeString, Computed: true},
												"security_alerts_threshold": {Type: schema.TypeString, Computed: true},
												"tool":                      {Type: schema.TypeString, Computed: true},
											},
										},
									},
								},
							},
						},
						"file_path_restriction": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"restricted_file_paths": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"max_file_size": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_file_size": {Type: schema.TypeInt, Computed: true},
								},
							},
						},
						"max_file_path_length": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_file_path_length": {Type: schema.TypeInt, Computed: true},
								},
							},
						},
						"file_extension_restriction": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"restricted_file_extensions": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"copilot_code_review": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"review_on_push": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"review_draft_pull_requests": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryRulesetRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	owner, _ := m.(*Owner)
	client := owner.v3client
	ownerName := owner.name

	repoName, ok := d.Get("repository").(string)
	if !ok {
		return diag.Errorf(`expected "repository" to be string`)
	}
	rulesetIDInt, ok := d.Get("ruleset_id").(int)
	if !ok {
		return diag.Errorf(`expected "ruleset_id" to be int`)
	}
	rulesetId := int64(rulesetIDInt)

	ruleset, _, err := client.Repositories.GetRuleset(ctx, ownerName, repoName, rulesetId, false)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(ruleset.GetID(), 10))

	if err := d.Set("name", ruleset.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target", ruleset.GetTarget()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enforcement", ruleset.GetEnforcement()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("node_id", ruleset.GetNodeID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("bypass_actors", flattenBypassActors(ruleset.BypassActors)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("conditions", flattenConditions(ctx, ruleset.GetConditions(), false)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rules", flattenRules(ctx, ruleset.GetRules(), false)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
