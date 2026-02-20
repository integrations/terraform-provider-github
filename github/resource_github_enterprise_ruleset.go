package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseRuleset() *schema.Resource {
	return &schema.Resource{
		Description: "Manages GitHub enterprise rulesets",

		CreateContext: resourceGithubEnterpriseRulesetCreate,
		ReadContext:   resourceGithubEnterpriseRulesetRead,
		UpdateContext: resourceGithubEnterpriseRulesetUpdate,
		DeleteContext: resourceGithubEnterpriseRulesetDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseRulesetImport,
		},

		CustomizeDiff: resourceGithubEnterpriseRulesetCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
				Description:  "The name of the ruleset.",
			},
			"target": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(supportedEnterpriseRulesetTargetTypes, false)),
				Description:      "Possible values are `branch`, `tag`, `push` and `repository`. Note: The `repository` target is in preview and is subject to change.",
			},
			"enforcement": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"disabled", "active", "evaluate"}, false)),
				Description:  "Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.",
			},
			"bypass_actors": {
				Type:             schema.TypeList,
				Optional:         true,
				DiffSuppressFunc: bypassActorsDiffSuppressFunc,
				Description:      "The actors that can bypass the rules in this ruleset.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"actor_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     nil,
							Description: "The ID of the actor that can bypass a ruleset. When `actor_type` is `OrganizationAdmin`, this should be set to `1`. Some resources such as DeployKey do not have an ID and this should be omitted.",
						},
						"actor_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"Integration", "OrganizationAdmin", "RepositoryRole", "Team", "DeployKey", "EnterpriseOwner"}, false)),
							Description:  "The type of actor that can bypass a ruleset. See https://docs.github.com/en/rest/enterprise-admin/rules for more information",
						},
						"bypass_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"always", "pull_request", "exempt"}, false)),
							Description:  "When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`, `exempt`.",
						},
					},
				},
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GraphQL global node id for use with v4 API.",
			},
			"ruleset_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "GitHub ID for the ruleset.",
			},
			"conditions": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Parameters for an enterprise ruleset condition. Enterprise rulesets must include organization targeting (organization_name or organization_id) and repository targeting (repository_name or repository_id). For branch and tag targets, ref_name is also required.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organization_name": {
							Type:          schema.TypeList,
							Optional:      true,
							MaxItems:      1,
							ConflictsWith: []string{"conditions.0.organization_id"},
							Description:   "Conditions for organization names that the ruleset targets. Conflicts with `organization_id`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Array of organization names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~ALL` to include all organizations.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"exclude": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Array of organization names or patterns to exclude. The condition will not pass if any of these patterns match.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"organization_id": {
							Type:          schema.TypeList,
							Optional:      true,
							ConflictsWith: []string{"conditions.0.organization_name"},
							Description:   "Organization IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `organization_name`.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"ref_name": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Conditions for ref names (branches or tags) that the ruleset targets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"exclude": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"repository_name": {
							Type:         schema.TypeList,
							Optional:     true,
							MaxItems:     1,
							Description:  "Conditions for repository names that the ruleset targets. Exactly one of `repository_name`, `repository_id`, or `repository_property` must be set.",
							ExactlyOneOf: []string{"conditions.0.repository_name", "conditions.0.repository_id", "conditions.0.repository_property"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Array of repository names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~ALL` to include all repositories.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"exclude": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Array of repository names or patterns to exclude. The condition will not pass if any of these patterns match.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"protected": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Whether renaming of target repositories is prevented.",
									},
								},
							},
						},
						"repository_id": {
							Type:         schema.TypeList,
							Optional:     true,
							ExactlyOneOf: []string{"conditions.0.repository_name", "conditions.0.repository_id", "conditions.0.repository_property"},
							Description:  "The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Exactly one of `repository_name`, `repository_id`, or `repository_property` must be set.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"repository_property": {
							Type:         schema.TypeList,
							Optional:     true,
							MaxItems:     1,
							Description:  "Conditions based on repository properties. Exactly one of `repository_name`, `repository_id`, or `repository_property` must be set.",
							ExactlyOneOf: []string{"conditions.0.repository_name", "conditions.0.repository_id", "conditions.0.repository_property"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Array of repository property conditions to include.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of the repository property to target.",
												},
												"property_values": {
													Type:        schema.TypeList,
													Required:    true,
													Description: "The values to match for the repository property.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The source of the repository property.",
												},
											},
										},
									},
									"exclude": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Array of repository property conditions to exclude.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of the repository property to target.",
												},
												"property_values": {
													Type:        schema.TypeList,
													Required:    true,
													Description: "The values to match for the repository property.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The source of the repository property.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Rules within the ruleset.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only allow users with bypass permission to create matching refs.",
						},
						"update": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only allow users with bypass permission to update matching refs.",
						},
						"deletion": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only allow users with bypass permissions to delete matching refs.",
						},
						"required_linear_history": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Prevent merge commits from being pushed to matching branches.",
						},
						"required_signatures": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Commits pushed to matching branches must have verified signatures.",
						},
						"pull_request": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Require all commits be made to a non-target branch and submitted via a pull request before they can be merged.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_merge_methods": {
										Type:        schema.TypeList,
										Optional:    true,
										MinItems:    1,
										Description: "Array of allowed merge methods. Allowed values include `merge`, `squash`, and `rebase`. At least one option must be enabled.",
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"merge", "squash", "rebase"}, false)),
										},
									},
									"dismiss_stale_reviews_on_push": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "New, reviewable commits pushed will dismiss previous pull request review approvals. Defaults to `false`.",
									},
									"require_code_owner_review": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Require an approving review in pull requests that modify files that have a designated code owner. Defaults to `false`.",
									},
									"require_last_push_approval": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Whether the most recent reviewable push must be approved by someone other than the person who pushed it. Defaults to `false`.",
									},
									"required_approving_review_count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "The number of approving reviews that are required before a pull request can be merged. Defaults to `0`.",
									},
									"required_review_thread_resolution": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "All conversations on code must be resolved before a pull request can be merged. Defaults to `false`.",
									},
								},
							},
						},
						"copilot_code_review": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Automatically request Copilot code review for new pull requests if the author has access to Copilot code review and their premium requests quota has not reached the limit.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"review_on_push": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Copilot automatically reviews each new push to the pull request. Defaults to `false`.",
									},
									"review_draft_pull_requests": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Copilot automatically reviews draft pull requests before they are marked as ready for review. Defaults to `false`.",
									},
								},
							},
						},
						"required_status_checks": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_check": {
										Type:        schema.TypeSet,
										MinItems:    1,
										Required:    true,
										Description: "Status checks that are required. Several can be defined.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"context": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The status check context name that must be present on the commit.",
												},
												"integration_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "The optional integration ID that this status check must originate from.",
												},
											},
										},
									},
									"strict_required_status_checks_policy": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.",
									},
									"do_not_enforce_on_create": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Allow repositories and branches to be created if a check would otherwise prohibit it.",
										Default:     false,
									},
								},
							},
						},
						"non_fast_forward": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Prevent users with push access from force pushing to branches.",
						},
						"commit_message_pattern": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Parameters to be used for the commit_message_pattern rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "How this rule will appear to users.",
									},
									"negate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the rule will fail if the pattern matches.",
									},
									"operator": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
									},
									"pattern": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The pattern to match with.",
									},
								},
							},
						},
						"commit_author_email_pattern": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Parameters to be used for the commit_author_email_pattern rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "How this rule will appear to users.",
									},
									"negate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the rule will fail if the pattern matches.",
									},
									"operator": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
									},
									"pattern": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The pattern to match with.",
									},
								},
							},
						},
						"committer_email_pattern": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Parameters to be used for the committer_email_pattern rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "How this rule will appear to users.",
									},
									"negate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the rule will fail if the pattern matches.",
									},
									"operator": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
									},
									"pattern": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The pattern to match with.",
									},
								},
							},
						},
						"branch_name_pattern": {
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"rules.0.tag_name_pattern"},
							Description:   "Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applies to rulesets with target `branch`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "How this rule will appear to users.",
									},
									"negate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the rule will fail if the pattern matches.",
									},
									"operator": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
									},
									"pattern": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The pattern to match with.",
									},
								},
							},
						},
						"tag_name_pattern": {
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"rules.0.branch_name_pattern"},
							Description:   "Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applies to rulesets with target `tag`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "How this rule will appear to users.",
									},
									"negate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the rule will fail if the pattern matches.",
									},
									"operator": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
									},
									"pattern": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The pattern to match with.",
									},
								},
							},
						},
						"required_workflows": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Choose which Actions workflows must pass before branches can be merged into a branch that matches this rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"do_not_enforce_on_create": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Allow repositories and branches to be created if a check would otherwise prohibit it.",
									},
									"required_workflow": {
										Type:        schema.TypeSet,
										MinItems:    1,
										Required:    true,
										Description: "Actions workflows that are required. Several can be defined.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"repository_id": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "The repository in which the workflow is defined.",
												},
												"path": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The path to the workflow YAML definition file.",
												},
												"ref": {
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "master",
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
							MaxItems:    1,
							Optional:    true,
							Description: "Choose which tools must provide code scanning results before the reference is updated. When configured, code scanning must be enabled and have results for both the commit and the reference being updated.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_code_scanning_tool": {
										Type:        schema.TypeSet,
										MinItems:    1,
										Required:    true,
										Description: "Tools that must provide code scanning results for this rule to pass.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alerts_threshold": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.",
												},
												"security_alerts_threshold": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.",
												},
												"tool": {
													Type:        schema.TypeString,
													Required:    true,
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
							Optional:    true,
							MaxItems:    1,
							Description: "Prevent commits that include changes in specified file paths from being pushed to the commit graph.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"restricted_file_paths": {
										Type:        schema.TypeList,
										MinItems:    1,
										Required:    true,
										Description: "The file paths that are restricted from being pushed to the commit graph.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"max_file_size": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Prevent pushes based on file size.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_file_size": {
										Type:             schema.TypeInt,
										Required:         true,
										Description:      "The maximum allowed size of a file in megabytes (MB). Valid range is 1-100 MB.",
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1, 100)),
									},
								},
							},
						},
						"max_file_path_length": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Prevent pushes based on file path length.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_file_path_length": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The maximum allowed length of a file path.",
									},
								},
							},
						},
						"file_extension_restriction": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Prevent pushes based on file extensions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"restricted_file_extensions": {
										Type:        schema.TypeSet,
										MinItems:    1,
										Required:    true,
										Description: "The file extensions that are restricted from being pushed to the commit graph.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						// Repository target rules (only valid when target = "repository")
						"repository_creation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only allow users with bypass permission to create repositories. Only valid for `repository` target.",
						},
						"repository_deletion": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only allow users with bypass permission to delete repositories. Only valid for `repository` target.",
						},
						"repository_transfer": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only allow users with bypass permission to transfer repositories. Only valid for `repository` target.",
						},
						"repository_name": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Restrict repository names to match specified patterns. Only valid for `repository` target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"negate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "If true, the rule will fail if the pattern matches.",
									},
									"pattern": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The pattern to match repository names against.",
									},
								},
							},
						},
						"repository_visibility": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Restrict repository visibility changes. Only valid for `repository` target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"internal": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Allow internal visibility for repositories.",
									},
									"private": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Allow private visibility for repositories.",
									},
								},
							},
						},
					},
				},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubEnterpriseRulesetCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	name := d.Get("name").(string)

	tflog.Debug(ctx, "Creating enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"name":            name,
	})

	rulesetReq := resourceGithubEnterpriseRulesetObject(d)

	ruleset, resp, err := client.Enterprise.CreateRepositoryRuleset(ctx, enterpriseSlug, rulesetReq)
	if err != nil {
		tflog.Error(ctx, "Failed to create enterprise ruleset", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"name":            name,
			"error":           err.Error(),
		})
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(*ruleset.ID, 10))
	if err := d.Set("ruleset_id", ruleset.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("node_id", ruleset.GetNodeID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Created enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"name":            name,
		"ruleset_id":      *ruleset.ID,
	})

	return nil
}

func resourceGithubEnterpriseRulesetRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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
			if ghErr.Response.StatusCode == http.StatusNotModified {
				tflog.Debug(ctx, "API responded with StatusNotModified, not refreshing state", map[string]any{
					"enterprise_slug": enterpriseSlug,
					"ruleset_id":      rulesetID,
				})
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing ruleset from state because it no longer exists in GitHub", map[string]any{
					"enterprise_slug": enterpriseSlug,
					"ruleset_id":      rulesetID,
				})
				d.SetId("")
				return nil
			}
		}
		tflog.Error(ctx, "Failed to read enterprise ruleset", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"ruleset_id":      rulesetID,
			"error":           err.Error(),
		})
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
	if err := d.Set("conditions", flattenConditions(ctx, ruleset.GetConditions(), true)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rules", flattenRules(ctx, ruleset.Rules, true)); err != nil {
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

func resourceGithubEnterpriseRulesetUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	name := d.Get("name").(string)
	rulesetID := int64(d.Get("ruleset_id").(int))

	tflog.Debug(ctx, "Updating enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
		"name":            name,
	})

	rulesetReq := resourceGithubEnterpriseRulesetObject(d)

	_, resp, err := client.Enterprise.UpdateRepositoryRuleset(ctx, enterpriseSlug, rulesetID, rulesetReq)
	if err != nil {
		tflog.Error(ctx, "Failed to update enterprise ruleset", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"ruleset_id":      rulesetID,
			"error":           err.Error(),
		})
		return diag.FromErr(err)
	}

	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Updated enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
		"name":            name,
	})

	return nil
}

func resourceGithubEnterpriseRulesetDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	rulesetID := int64(d.Get("ruleset_id").(int))

	tflog.Debug(ctx, "Deleting enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
	})

	_, err := client.Enterprise.DeleteRepositoryRuleset(ctx, enterpriseSlug, rulesetID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Enterprise ruleset already deleted", map[string]any{
				"enterprise_slug": enterpriseSlug,
				"ruleset_id":      rulesetID,
			})
			return nil
		}
		tflog.Error(ctx, "Failed to delete enterprise ruleset", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"ruleset_id":      rulesetID,
			"error":           err.Error(),
		})
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Deleted enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
	})

	return nil
}

func resourceGithubEnterpriseRulesetImport(ctx context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	enterpriseSlug, rulesetIDStr, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("error importing enterprise ruleset (expected format: <enterprise_slug>:<ruleset_id>): %w", err)
	}

	rulesetID, err := strconv.ParseInt(rulesetIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error importing enterprise ruleset (expected format: <enterprise_slug>:<ruleset_id>): %w", unconvertibleIdErr(rulesetIDStr, err))
	}
	if rulesetID == 0 {
		return nil, fmt.Errorf("error importing enterprise ruleset (expected format: <enterprise_slug>:<ruleset_id>): ruleset_id must be present")
	}

	tflog.Debug(ctx, "Importing enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
	})

	d.SetId(rulesetIDStr)
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	if err := d.Set("ruleset_id", rulesetID); err != nil {
		return nil, err
	}

	tflog.Info(ctx, "Imported enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
	})

	return []*schema.ResourceData{d}, nil
}

var supportedEnterpriseRulesetTargetTypes = []string{
	string(github.RulesetTargetBranch),
	string(github.RulesetTargetTag),
	string(github.RulesetTargetPush),
	string(github.RulesetTargetRepository),
}

// resourceGithubEnterpriseRulesetObject creates a GitHub RepositoryRuleset object for enterprise-level rulesets
func resourceGithubEnterpriseRulesetObject(d *schema.ResourceData) github.RepositoryRuleset {
	return github.RepositoryRuleset{
		Name:         d.Get("name").(string),
		Target:       github.Ptr(github.RulesetTarget(d.Get("target").(string))),
		Source:       d.Get("enterprise_slug").(string),
		SourceType:   github.Ptr(github.RulesetSourceType("Enterprise")),
		Enforcement:  github.RulesetEnforcement(d.Get("enforcement").(string)),
		BypassActors: expandBypassActors(d.Get("bypass_actors").([]any)),
		Conditions:   expandConditions(d.Get("conditions").([]any), true),
		Rules:        expandRules(d.Get("rules").([]any), true),
	}
}
