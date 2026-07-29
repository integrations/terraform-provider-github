package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseRuleset() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseRulesetRead,

		Description: "Retrieves an existing GitHub enterprise ruleset by its ID.",

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "The slug of the enterprise the ruleset belongs to.",
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
				Description: "The target of the ruleset. One of `branch`, `tag`, `push` or `repository`.",
			},
			"enforcement": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The enforcement level of the ruleset. One of `disabled`, `active` or `evaluate`.",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GraphQL global node id for use with v4 API.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the ruleset for caching purposes.",
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
				Description: "The organizations, repositories and refs the ruleset applies to.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organization_name": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Organization names or patterns the ruleset targets.",
							Elem:        dataSourceRulesetIncludeExcludeElem("organization"),
						},
						"organization_id": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Organization IDs the ruleset targets.",
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
						"organization_property": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Organization properties the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The organization properties and values to include.",
										Elem:        dataSourceRulesetOrganizationPropertyElem(),
									},
									"exclude": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The organization properties and values to exclude.",
										Elem:        dataSourceRulesetOrganizationPropertyElem(),
									},
								},
							},
						},
						"ref_name": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Ref names or patterns the ruleset targets.",
							Elem:        dataSourceRulesetIncludeExcludeElem("ref"),
						},
						"repository_name": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Repository names or patterns the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of repository names or patterns to include.",
									},
									"exclude": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of repository names or patterns to exclude.",
									},
									"protected": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether renaming of target repositories is prevented.",
									},
								},
							},
						},
						"repository_property": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Repository properties the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The repository properties and values to include.",
										Elem:        dataSourceRulesetRepositoryPropertyElem(),
									},
									"exclude": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The repository properties and values to exclude.",
										Elem:        dataSourceRulesetRepositoryPropertyElem(),
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
							Description: "Prevent users with push access from force pushing to refs.",
						},
						"pull_request": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Pull request requirements for matching refs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_merge_methods": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Array of allowed merge methods.",
									},
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
									"required_reviewers": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Reviewers that must approve pull requests touching the matching files.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"reviewer": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The reviewer that must review matching files.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The ID of the reviewer that must review.",
															},
															"type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The type of reviewer.",
															},
														},
													},
												},
												"file_patterns": {
													Type:        schema.TypeList,
													Computed:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "File patterns (fnmatch syntax) that this reviewer must approve.",
												},
												"minimum_approvals": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Minimum number of approvals required from this reviewer.",
												},
											},
										},
									},
								},
							},
						},
						"copilot_code_review": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Copilot code review configuration for matching pull requests.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"review_on_push": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Copilot automatically reviews each new push to the pull request.",
									},
									"review_draft_pull_requests": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Copilot automatically reviews draft pull requests before they are marked as ready for review.",
									},
								},
							},
						},
						"required_status_checks": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Status checks that must pass before matching refs can be updated.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_check": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Status checks that are required.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"context": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The status check context name that must be present on the commit.",
												},
												"integration_id": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The integration ID that this status check must originate from.",
												},
											},
										},
									},
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
						"commit_message_pattern": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Parameters used for the commit_message_pattern rule.",
							Elem:        dataSourceRulesetPatternRuleElem(),
						},
						"commit_author_email_pattern": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Parameters used for the commit_author_email_pattern rule.",
							Elem:        dataSourceRulesetPatternRuleElem(),
						},
						"committer_email_pattern": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Parameters used for the committer_email_pattern rule.",
							Elem:        dataSourceRulesetPatternRuleElem(),
						},
						"branch_name_pattern": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Parameters used for the branch_name_pattern rule.",
							Elem:        dataSourceRulesetPatternRuleElem(),
						},
						"tag_name_pattern": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Parameters used for the tag_name_pattern rule.",
							Elem:        dataSourceRulesetPatternRuleElem(),
						},
						"required_workflows": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Actions workflows that must pass before matching refs can be updated.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"do_not_enforce_on_create": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allow repositories and branches to be created if a check would otherwise prohibit it.",
									},
									"required_workflow": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Actions workflows that are required.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"repository_id": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The repository in which the workflow is defined.",
												},
												"path": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The path to the workflow YAML definition file.",
												},
												"ref": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The ref (branch or tag) of the workflow file to use.",
												},
											},
										},
									},
								},
							},
						},
						"required_code_scanning": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Code scanning tools that must provide results before matching refs can be updated.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_code_scanning_tool": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Tools that must provide code scanning results for this rule to pass.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alerts_threshold": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The severity level at which code scanning results that raise alerts block a reference update.",
												},
												"security_alerts_threshold": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The severity level at which code scanning results that raise security alerts block a reference update.",
												},
												"tool": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of a code scanning tool.",
												},
											},
										},
									},
								},
							},
						},
						"file_path_restriction": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "File paths that are restricted from being pushed to the commit graph.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"restricted_file_paths": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The restricted file paths.",
									},
								},
							},
						},
						"max_file_size": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Maximum allowed file size.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_file_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum allowed size of a file in megabytes (MB).",
									},
								},
							},
						},
						"max_file_path_length": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Maximum allowed file path length.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_file_path_length": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum allowed length of a file path.",
									},
								},
							},
						},
						"file_extension_restriction": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "File extensions that are restricted from being pushed to the commit graph.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"restricted_file_extensions": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The restricted file extensions.",
									},
								},
							},
						},
						"repository_create": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Only allow users with bypass permission to create matching repositories.",
						},
						"repository_delete": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Only allow users with bypass permission to delete matching repositories.",
						},
						"repository_transfer": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Only allow users with bypass permission to transfer matching repositories.",
						},
						"repository_name": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Pattern that matching repository names must satisfy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If true, the rule fails if the pattern matches.",
									},
									"pattern": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The pattern repository names are matched against.",
									},
								},
							},
						},
						"repository_visibility": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Visibilities a matching repository may have.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"internal": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether internal visibility is allowed.",
									},
									"private": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether private visibility is allowed.",
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

// dataSourceRulesetIncludeExcludeElem returns the read-only element schema for the
// include/exclude pattern conditions, whose shape is identical for every subject.
func dataSourceRulesetIncludeExcludeElem(subject string) *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"include": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: fmt.Sprintf("Array of %s names or patterns to include.", subject),
			},
			"exclude": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: fmt.Sprintf("Array of %s names or patterns to exclude.", subject),
			},
		},
	}
}

// dataSourceRulesetRepositoryPropertyElem returns the read-only element schema for a
// single repository property condition.
func dataSourceRulesetRepositoryPropertyElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the repository property.",
			},
			"property_values": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The values matched for the repository property.",
			},
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source of the repository property. One of `custom` or `system`.",
			},
		},
	}
}

// dataSourceRulesetOrganizationPropertyElem returns the read-only element schema for a
// single `organization_property` entry. Organization properties have no `source`.
func dataSourceRulesetOrganizationPropertyElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the organization property.",
			},
			"property_values": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The values matched for the organization property.",
			},
		},
	}
}

// dataSourceRulesetPatternRuleElem returns the read-only element schema shared by
// every `*_pattern` rule.
func dataSourceRulesetPatternRuleElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "How this rule appears to users.",
			},
			"negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, the rule fails if the pattern matches.",
			},
			"operator": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The operator used for matching.",
			},
			"pattern": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The pattern matched against.",
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
		tflog.Error(ctx, "Failed to read enterprise ruleset", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"ruleset_id":      rulesetID,
			"error":           err.Error(),
		})
		return diag.FromErr(err)
	}

	id, err := buildID(escapeIDPart(enterpriseSlug), strconv.FormatInt(ruleset.GetID(), 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("name", ruleset.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target", ruleset.GetTarget()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enforcement", ruleset.Enforcement); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("bypass_actors", flattenBypassActors(ctx, ruleset.BypassActors)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("conditions", flattenConditions(ctx, ruleset.GetConditions(), rulesetLevelEnterprise)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rules", flattenRules(ctx, ruleset.Rules, rulesetLevelEnterprise)); err != nil {
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
