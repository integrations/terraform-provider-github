package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var supportedRulesetTargetTypes = []string{string(github.RulesetTargetBranch), string(github.RulesetTargetPush), string(github.RulesetTargetTag)}

func resourceGithubRepositoryRuleset() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryRulesetCreate,
		ReadContext:   resourceGithubRepositoryRulesetRead,
		UpdateContext: resourceGithubRepositoryRulesetUpdate,
		DeleteContext: resourceGithubRepositoryRulesetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryRulesetImport,
		},

		SchemaVersion: 1,

		CustomizeDiff: resourceGithubRepositoryRulesetDiff,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
				Description:      "The name of the ruleset.",
			},
			"target": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(supportedRulesetTargetTypes, false)),
				Description:      "Possible values are " + strings.Join(supportedRulesetTargetTypes[:len(supportedRulesetTargetTypes)-1], ", ") + " and " + supportedRulesetTargetTypes[len(supportedRulesetTargetTypes)-1],
			},
			"repository": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringMatch(regexp.MustCompile(`^[-a-zA-Z0-9_.]{1,100}$`), "must include only alphanumeric characters, underscores or hyphens and consist of 100 characters or less")),
				Description:      "Name of the repository to apply ruleset to.",
			},
			"enforcement": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"disabled", "active", "evaluate"}, false)),
				Description:      "Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.",
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
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"RepositoryRole", "Team", "Integration", "OrganizationAdmin", "DeployKey"}, false)),
							Description:      "The type of actor that can bypass a ruleset. See https://docs.github.com/en/rest/repos/rules for more information.",
						},
						"bypass_mode": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"always", "pull_request", "exempt"}, false)),
							Description:      "When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`, `exempt`.",
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
				Description: "Parameters for a repository ruleset ref name condition.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ref_name": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
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
										Required:    true,
										Description: "Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
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
						"update_allows_fetch_and_merge": {
							Type:         schema.TypeBool,
							Optional:     true,
							Default:      false,
							RequiredWith: []string{"rules.0.update"},
							Description:  "Branch can pull changes from its upstream repository. This is only applicable to forked repositories. Requires `update` to be set to `true`.",
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
						"required_deployments": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_deployment_environments": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "The environments that must be successfully deployed to before branches can be merged.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
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
										Computed:    true,
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
										Type:             schema.TypeInt,
										Optional:         true,
										Default:          0,
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 10)),
										Description:      "The number of approving reviews that are required before a pull request can be merged. Defaults to `0`.",
									},
									"required_review_thread_resolution": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "All conversations on code must be resolved before a pull request can be merged. Defaults to `false`.",
									},
									"required_reviewers": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Require specific reviewers to approve pull requests targeting matching branches. Note: This feature is in beta and subject to change.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"reviewer": {
													Type:        schema.TypeList,
													Required:    true,
													MaxItems:    1,
													Description: "The reviewer that must review matching files.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Required:    true,
																Description: "The ID of the reviewer that must review.",
															},
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"Team"}, false)),
																Description:      "The type of reviewer. Currently only `Team` is supported.",
															},
														},
													},
												},
												"file_patterns": {
													Type:        schema.TypeList,
													Required:    true,
													MinItems:    1,
													Description: "File patterns (fnmatch syntax) that this reviewer must approve.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"minimum_approvals": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Minimum number of approvals required from this reviewer. Set to 0 to make approval optional.",
												},
											},
										},
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
						"merge_queue": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Merges must be performed via a merge queue.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"check_response_timeout_minutes": {
										Type:             schema.TypeInt,
										Optional:         true,
										Default:          60,
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 360)),
										Description:      "Maximum time for a required status check to report a conclusion. After this much time has elapsed, checks that have not reported a conclusion will be assumed to have failed. Defaults to `60`.",
									},
									"grouping_strategy": {
										Type:             schema.TypeString,
										Optional:         true,
										Default:          "ALLGREEN",
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"ALLGREEN", "HEADGREEN"}, false)),
										Description:      "When set to ALLGREEN, the merge commit created by merge queue for each PR in the group must pass all required checks to merge. When set to HEADGREEN, only the commit at the head of the merge group, i.e. the commit containing changes from all of the PRs in the group, must pass its required checks to merge. Can be one of: ALLGREEN, HEADGREEN. Defaults to `ALLGREEN`.",
									},
									"max_entries_to_build": {
										Type:             schema.TypeInt,
										Optional:         true,
										Default:          5,
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 100)),
										Description:      "Limit the number of queued pull requests requesting checks and workflow runs at the same time. Defaults to `5`.",
									},
									"max_entries_to_merge": {
										Type:             schema.TypeInt,
										Optional:         true,
										Default:          5,
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 100)),
										Description:      "The maximum number of PRs that will be merged together in a group. Defaults to `5`.",
									},
									"merge_method": {
										Type:             schema.TypeString,
										Optional:         true,
										Default:          "MERGE",
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"MERGE", "SQUASH", "REBASE"}, false)),
										Description:      "Method to use when merging changes from queued pull requests. Can be one of: MERGE, SQUASH, REBASE. Defaults to `MERGE`.",
									},
									"min_entries_to_merge": {
										Type:             schema.TypeInt,
										Optional:         true,
										Default:          1,
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 100)),
										Description:      "The minimum number of PRs that will be merged together in a group. Defaults to `1`.",
									},
									"min_entries_to_merge_wait_minutes": {
										Type:             schema.TypeInt,
										Optional:         true,
										Default:          5,
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 360)),
										Description:      "The time merge queue should wait after the first PR is added to the queue for the minimum group size to be met. After this time has elapsed, the minimum group size will be ignored and a smaller group will be merged. Defaults to `5`.",
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
							Description: "Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations.",
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
										Type:             schema.TypeString,
										ValidateDiagFunc: operatorValidation,
										Required:         true,
										Description:      "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
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
							Description: "Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations.",
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
										Type:             schema.TypeString,
										ValidateDiagFunc: operatorValidation,
										Required:         true,
										Description:      "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
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
							Description: "Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations.",
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
										Type:             schema.TypeString,
										ValidateDiagFunc: operatorValidation,
										Required:         true,
										Description:      "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
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
										Type:             schema.TypeString,
										ValidateDiagFunc: operatorValidation,
										Required:         true,
										Description:      "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
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
										Type:             schema.TypeString,
										ValidateDiagFunc: operatorValidation,
										Required:         true,
										Description:      "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
									},
									"pattern": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The pattern to match with.",
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
													Description:      "The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.",
													Required:         true,
													Type:             schema.TypeString,
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"none", "errors", "errors_and_warnings", "all"}, false)),
												},
												"security_alerts_threshold": {
													Description:      "The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.",
													Required:         true,
													Type:             schema.TypeString,
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"none", "critical", "high_or_higher", "medium_or_higher", "all"}, false)),
												},
												"tool": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of a code scanning tool",
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
										Type:             schema.TypeInt,
										Required:         true,
										Description:      "The maximum allowed length of a file path.",
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1, 32767)),
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
										Description: "A list of file extensions.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

func resourceGithubRepositoryRulesetCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	rulesetReq := resourceGithubRulesetObject(d, "")

	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)

	// Check if repository is archived - cannot create rulesets on archived repos (attempts PUT on read-only resource)
	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	if repo.GetArchived() {
		return diag.Errorf("cannot create ruleset on archived repository %s/%s", owner, repoName)
	}

	ruleset, resp, err := client.Repositories.CreateRuleset(ctx, owner, repoName, rulesetReq)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(ruleset.GetID(), 10))
	if err := d.Set("ruleset_id", ruleset.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("node_id", ruleset.GetNodeID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rules", flattenRules(ctx, ruleset.Rules, false)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryRulesetRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	rulesetID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}

	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	ruleset, resp, err := client.Repositories.GetRuleset(ctx, owner, repoName, rulesetID, false)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing ruleset from state because it no longer exists in GitHub", map[string]any{"owner": owner, "repo_name": repoName, "ruleset_id": rulesetID})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if ruleset == nil {
		tflog.Info(ctx, "Removing ruleset from state because it no longer exists in GitHub (empty response)", map[string]any{"owner": owner, "repo_name": repoName, "ruleset_id": rulesetID})
		d.SetId("")
		return nil
	}

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
	if err := d.Set("conditions", flattenConditions(ctx, ruleset.GetConditions(), false)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rules", flattenRules(ctx, ruleset.GetRules(), false)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("node_id", ruleset.GetNodeID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryRulesetUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	rulesetReq := resourceGithubRulesetObject(d, "")

	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	rulesetID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}

	// Check if repository is archived - skip update if it is
	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	if repo.GetArchived() {
		tflog.Info(ctx, "Repository is archived, skipping ruleset update", map[string]any{"owner": owner, "repo_name": repoName})
		return nil
	}

	ruleset, resp, err := client.Repositories.UpdateRuleset(ctx, owner, repoName, rulesetID, rulesetReq)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(ruleset.GetID(), 10))
	if err := d.Set("ruleset_id", ruleset.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("node_id", ruleset.GetNodeID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryRulesetDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	rulesetID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}

	tflog.Debug(ctx, "Deleting repository ruleset", map[string]any{"owner": owner, "repo_name": repoName, "ruleset_id": rulesetID})
	_, err = client.Repositories.DeleteRuleset(ctx, owner, repoName, rulesetID)
	return diag.FromErr(handleArchivedRepoDelete(err, "repository ruleset", strconv.FormatInt(rulesetID, 10), owner, repoName))
}

func resourceGithubRepositoryRulesetImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	repoName, rulesetIDStr, err := parseTwoPartID(d.Id(), "repository", "ruleset")
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	rulesetID, err := strconv.ParseInt(rulesetIDStr, 10, 64)
	if err != nil {
		return []*schema.ResourceData{d}, unconvertibleIdErr(rulesetIDStr, err)
	}
	if rulesetID == 0 {
		return []*schema.ResourceData{d}, fmt.Errorf("`ruleset_id` must be present")
	}
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	tflog.Debug(ctx, "Importing repository ruleset", map[string]any{"owner": owner, "repo_name": repoName, "ruleset_id": rulesetID})

	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if repository == nil || err != nil {
		return []*schema.ResourceData{d}, err
	}
	if err := d.Set("repository", repository.GetName()); err != nil {
		return []*schema.ResourceData{d}, err
	}

	ruleset, _, err := client.Repositories.GetRuleset(ctx, owner, repository.GetName(), rulesetID, false)
	if ruleset == nil || err != nil {
		return []*schema.ResourceData{d}, err
	}
	d.SetId(strconv.FormatInt(ruleset.GetID(), 10))

	return []*schema.ResourceData{d}, nil
}

func resourceGithubRepositoryRulesetDiff(ctx context.Context, d *schema.ResourceDiff, meta any) error {
	err := validateRulesetConditions(ctx, d, false)
	if err != nil {
		return err
	}

	err = validateRulesetRules(ctx, d)
	if err != nil {
		return err
	}

	return nil
}
