package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var supportedEnterpriseRulesetTargetTypes = []string{
	string(github.RulesetTargetBranch),
	string(github.RulesetTargetTag),
	string(github.RulesetTargetPush),
	string(github.RulesetTargetRepository),
}

func resourceGithubEnterpriseRuleset() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubEnterpriseRulesetCreate,
		ReadContext:   resourceGithubEnterpriseRulesetRead,
		UpdateContext: resourceGithubEnterpriseRulesetUpdate,
		DeleteContext: resourceGithubEnterpriseRulesetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseRulesetImport,
		},

		CustomizeDiff: resourceGithubEnterpriseRulesetDiff,

		Description: "Creates a GitHub enterprise ruleset. Enterprise rulesets apply to every repository in the organizations they target, and require an enterprise plan.",

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "The slug of the enterprise the ruleset belongs to.",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
				Description:      "The name of the ruleset.",
			},
			"target": {
				Type:     schema.TypeString,
				Required: true,
				// Updatable in place: the API accepts `target` in the body of
				// PUT /enterprises/{enterprise}/rulesets/{ruleset_id}, same as the org and
				// repository rulesets. CustomizeDiff validates that the conditions and rules
				// in the new configuration are legal for the new target.
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(supportedEnterpriseRulesetTargetTypes, false)),
				Description:      "The target of the ruleset. Possible values are " + strings.Join(supportedEnterpriseRulesetTargetTypes[:len(supportedEnterpriseRulesetTargetTypes)-1], ", ") + " and " + supportedEnterpriseRulesetTargetTypes[len(supportedEnterpriseRulesetTargetTypes)-1] + ". The `repository` target is only available for enterprise rulesets.",
			},
			"enforcement": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"disabled", "active", "evaluate"}, false)),
				Description:      "The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Possible values are `disabled`, `active`, and `evaluate`.",
			},
			"bypass_actors": {
				Type:             schema.TypeList, // The GitHub API returns these sorted by actor_id, so order is suppressed rather than enforced.
				Optional:         true,
				DiffSuppressFunc: suppressUnorderedListDiff("bypass_actors", bypassActorCompareIdentity),
				Description:      "The actors that can bypass the rules in this ruleset.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"actor_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     nil,
							Description: "The ID of the actor that can bypass a ruleset. Must be omitted for ID-less actor types: `OrganizationAdmin`, `EnterpriseOwner`, and `DeployKey` — the GitHub API does not use an ID for these types and will ignore any value set.",
						},
						"actor_type": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"Integration", "OrganizationAdmin", "RepositoryRole", "Team", "DeployKey", "EnterpriseOwner"}, false)),
							Description:      "The type of actor that can bypass a ruleset. Can be one of: `Integration`, `OrganizationAdmin`, `RepositoryRole`, `Team`, `DeployKey`, or `EnterpriseOwner`.",
						},
						"bypass_mode": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"always", "pull_request", "exempt"}, false)),
							Description:      "When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type. Also, `pull_request` is only applicable to branch rulesets. When `bypass_mode` is `exempt`, rules will not be run for that actor and a bypass audit entry will not be created. Can be one of: `always`, `pull_request`, `exempt`.",
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
				Description: "Parameters for an enterprise ruleset condition. Enterprise rulesets select the organizations they apply to with exactly one of `organization_name`, `organization_id` or `organization_property`, and the repositories within them with exactly one of `repository_name` or `repository_property`. The `branch` and `tag` targets additionally require `ref_name`; the `push` and `repository` targets must not set it.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organization_name": {
							Type:         schema.TypeList,
							Optional:     true,
							MaxItems:     1,
							ExactlyOneOf: enterpriseRulesetOrganizationSelectors,
							Description:  "Targets organizations that match the specified name patterns.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Array of organization names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~ALL` to include all organizations and `~EMUS` to include the Enterprise Managed Users organization.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"exclude": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Array of organization names or patterns to exclude. The condition will not pass if any of these patterns match.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"organization_id": {
							Type:         schema.TypeList,
							Optional:     true,
							ExactlyOneOf: enterpriseRulesetOrganizationSelectors,
							Description:  "The organization IDs that the ruleset applies to. One of these IDs must match for the condition to pass.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"organization_property": {
							Type:         schema.TypeList,
							Optional:     true,
							MaxItems:     1,
							ExactlyOneOf: enterpriseRulesetOrganizationSelectors,
							Description:  "Conditions to target organizations by custom properties.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Optional:    true,
										ConfigMode:  schema.SchemaConfigModeAttr,
										Description: "The organization properties and values to include. All of these properties must match for the condition to pass.",
										Elem:        enterpriseRulesetOrganizationPropertyElem(),
									},
									"exclude": {
										Type:        schema.TypeList,
										Optional:    true,
										ConfigMode:  schema.SchemaConfigModeAttr,
										Description: "The organization properties and values to exclude. The ruleset will not apply if any of these properties match.",
										Elem:        enterpriseRulesetOrganizationPropertyElem(),
									},
								},
							},
						},
						"ref_name": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Targets refs that match the specified patterns. Required for `branch` and `tag` targets.",
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
						"repository_property": {
							Type:         schema.TypeList,
							Optional:     true,
							MaxItems:     1,
							ExactlyOneOf: enterpriseRulesetRepositorySelectors,
							Description:  "Conditions to target repositories by custom or system properties.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": {
										Type:        schema.TypeList,
										Optional:    true,
										ConfigMode:  schema.SchemaConfigModeAttr,
										Description: "The repository properties and values to include. All of these properties must match for the condition to pass.",
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
													Type:             schema.TypeString,
													Optional:         true,
													Default:          "custom",
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"custom", "system"}, false)),
													Description:      "The source of the repository property. Defaults to 'custom' if not specified. Can be one of: custom, system",
												},
											},
										},
									},
									"exclude": {
										Type:        schema.TypeList,
										Optional:    true,
										ConfigMode:  schema.SchemaConfigModeAttr,
										Description: "The repository properties and values to exclude. The ruleset will not apply if any of these properties match.",
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
													Type:             schema.TypeString,
													Optional:         true,
													Default:          "custom",
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"custom", "system"}, false)),
													Description:      "The source of the repository property. Defaults to 'custom' if not specified. Can be one of: custom, system",
												},
											},
										},
									},
								},
							},
						},
						"repository_name": {
							Type:         schema.TypeList,
							Optional:     true,
							MaxItems:     1,
							ExactlyOneOf: enterpriseRulesetRepositorySelectors,
							Description:  "Targets repositories that match the specified name patterns.",
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
										Required:    true,
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
					},
				},
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Rules within the ruleset. Rules are target specific: `branch` and `tag` targets accept the ref lifecycle and merge requirement rules, `push` targets accept the file content rules, and `repository` targets accept the `repository_*` rules. Using a rule that the target does not support fails at plan time.",
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
													Type:             schema.TypeString,
													Required:         true,
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
													Description:      "The status check context name that must be present on the commit.",
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
										Default:     false,
										Description: "Allow repositories and branches to be created if a check would otherwise prohibit it.",
									},
								},
							},
						},
						"non_fast_forward": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Prevent users with push access from force pushing to refs.",
						},
						"commit_message_pattern": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Parameters to be used for the commit_message_pattern rule.",
							Elem:        rulesetPatternRuleElem(),
						},
						"commit_author_email_pattern": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Parameters to be used for the commit_author_email_pattern rule.",
							Elem:        rulesetPatternRuleElem(),
						},
						"committer_email_pattern": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Parameters to be used for the committer_email_pattern rule.",
							Elem:        rulesetPatternRuleElem(),
						},
						"branch_name_pattern": {
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"rules.0.tag_name_pattern"},
							Description:   "Parameters to be used for the branch_name_pattern rule. Conflicts with `tag_name_pattern` as it only applies to rulesets with target `branch`.",
							Elem:          rulesetPatternRuleElem(),
						},
						"tag_name_pattern": {
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"rules.0.branch_name_pattern"},
							Description:   "Parameters to be used for the tag_name_pattern rule. Conflicts with `branch_name_pattern` as it only applies to rulesets with target `tag`.",
							Elem:          rulesetPatternRuleElem(),
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
													Type:             schema.TypeString,
													Required:         true,
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringMatch(regexp.MustCompile(`^\.github\/workflows\/.*$`), "Path must be in the .github/workflows directory")),
													Description:      "The path to the workflow YAML definition file.",
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
													Type:             schema.TypeString,
													Required:         true,
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"none", "errors", "errors_and_warnings", "all"}, false)),
													Description:      "The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.",
												},
												"security_alerts_threshold": {
													Type:             schema.TypeString,
													Required:         true,
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"none", "critical", "high_or_higher", "medium_or_higher", "all"}, false)),
													Description:      "The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.",
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
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1, 100)),
										Description:      "The maximum allowed size of a file in megabytes (MB). Valid range is 1-100 MB.",
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
										ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1, 32767)),
										Description:      "The maximum allowed length of a file path.",
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
						"repository_create": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only allow users with bypass permission to create matching repositories. Only valid for the `repository` target.",
						},
						"repository_delete": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only allow users with bypass permission to delete matching repositories. Only valid for the `repository` target.",
						},
						"repository_transfer": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only allow users with bypass permission to transfer matching repositories. Only valid for the `repository` target.",
						},
						"repository_name": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Restrict repository names to the specified pattern. Only valid for the `repository` target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"negate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "If true, the rule will fail if the pattern matches.",
									},
									"pattern": {
										Type:             schema.TypeString,
										Required:         true,
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
										Description:      "The pattern to match repository names against.",
									},
								},
							},
						},
						"repository_visibility": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Restrict the visibilities a matching repository may have. Only valid for the `repository` target.",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the ruleset for caching purposes.",
			},
		},
	}
}

// rulesetPatternRuleElem returns the element schema shared by every
// `*_pattern` rule of an enterprise ruleset.
func rulesetPatternRuleElem() *schema.Resource {
	return &schema.Resource{
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
				Required:         true,
				ValidateDiagFunc: operatorValidation,
				Description:      "The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.",
			},
			"pattern": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The pattern to match with.",
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

	ruleset, resp, err := client.Enterprise.CreateRepositoryRuleset(ctx, enterpriseSlug, resourceGithubEnterpriseRulesetObject(d))
	if err != nil {
		tflog.Error(ctx, "Failed to create enterprise ruleset", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"name":            name,
			"error":           err.Error(),
		})
		return diag.FromErr(err)
	}

	id, err := buildID(escapeIDPart(enterpriseSlug), strconv.FormatInt(ruleset.GetID(), 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("ruleset_id", ruleset.GetID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("node_id", ruleset.GetNodeID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rules", flattenRules(ctx, ruleset.Rules, rulesetLevelEnterprise)); err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Created enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"name":            name,
		"ruleset_id":      ruleset.GetID(),
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

	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	ruleset, resp, err := client.Enterprise.GetRepositoryRuleset(ctx, enterpriseSlug, rulesetID)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				tflog.Debug(ctx, "API responded with StatusNotModified, not refreshing state", map[string]any{
					"enterprise_slug": enterpriseSlug,
					"ruleset_id":      rulesetID,
				})
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing enterprise ruleset from state because it no longer exists in GitHub", map[string]any{
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

func resourceGithubEnterpriseRulesetUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	rulesetID := int64(d.Get("ruleset_id").(int))
	name := d.Get("name").(string)

	tflog.Debug(ctx, "Updating enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
		"name":            name,
	})

	ruleset, resp, err := client.Enterprise.UpdateRepositoryRuleset(ctx, enterpriseSlug, rulesetID, resourceGithubEnterpriseRulesetObject(d))
	if err != nil {
		tflog.Error(ctx, "Failed to update enterprise ruleset", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"ruleset_id":      rulesetID,
			"error":           err.Error(),
		})
		return diag.FromErr(err)
	}

	if err := d.Set("rules", flattenRules(ctx, ruleset.Rules, rulesetLevelEnterprise)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ruleset_id", ruleset.GetID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("node_id", ruleset.GetNodeID()); err != nil {
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

	if _, err := client.Enterprise.DeleteRepositoryRuleset(ctx, enterpriseSlug, rulesetID); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Enterprise ruleset is already gone", map[string]any{
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

// resourceGithubEnterpriseRulesetImport seeds the two attributes that every other CRUD
// function reads from state — `enterprise_slug` and `ruleset_id` — from the composite
// import ID. Terraform calls Read straight afterwards to populate the rest.
func resourceGithubEnterpriseRulesetImport(ctx context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	enterpriseSlug, rulesetID, err := parseEnterpriseRulesetID(d.Id())
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Importing enterprise ruleset", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"ruleset_id":      rulesetID,
	})

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	if err := d.Set("ruleset_id", rulesetID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func resourceGithubEnterpriseRulesetDiff(ctx context.Context, d *schema.ResourceDiff, _ any) error {
	if err := validateRulesetConditions(ctx, d); err != nil {
		return err
	}

	return validateRulesetRules(ctx, d)
}

// resourceGithubEnterpriseRulesetObject builds the API payload for an enterprise ruleset.
func resourceGithubEnterpriseRulesetObject(d *schema.ResourceData) github.RepositoryRuleset {
	target := github.RulesetTarget(d.Get("target").(string))
	sourceType := github.RulesetSourceTypeEnterprise

	return github.RepositoryRuleset{
		Name:         d.Get("name").(string),
		Target:       &target,
		Source:       d.Get("enterprise_slug").(string),
		SourceType:   &sourceType,
		Enforcement:  github.RulesetEnforcement(d.Get("enforcement").(string)),
		BypassActors: expandBypassActors(d.Get("bypass_actors").([]any)),
		Conditions:   expandConditions(d.Get("conditions").([]any), rulesetLevelEnterprise),
		Rules:        expandRules(d.Get("rules").([]any), rulesetLevelEnterprise),
	}
}

// parseEnterpriseRulesetID splits the `<enterprise_slug>:<ruleset_id>` resource ID.
//
// The ID has to be composite: unlike organization rulesets, whose owner comes from the
// provider configuration, an enterprise slug has no provider-level source, so it can only
// be recovered from the ID itself when importing.
func parseEnterpriseRulesetID(id string) (string, int64, error) {
	enterpriseSlug, rulesetIDStr, err := parseID2(id)
	if err != nil {
		return "", 0, fmt.Errorf("invalid enterprise ruleset ID %q (expected format: <enterprise_slug>:<ruleset_id>): %w", id, err)
	}

	rulesetID, err := strconv.ParseInt(rulesetIDStr, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("invalid enterprise ruleset ID %q (expected format: <enterprise_slug>:<ruleset_id>): %w", id, unconvertibleIdErr(rulesetIDStr, err))
	}
	if rulesetID == 0 {
		return "", 0, fmt.Errorf("invalid enterprise ruleset ID %q: `ruleset_id` must be present and non-zero", id)
	}

	return unescapeIDPart(enterpriseSlug), rulesetID, nil
}

// enterpriseRulesetOrganizationPropertyElem returns the element schema for a single
// `organization_property` entry.
//
// It deliberately has no `source` field: the enterprise API accepts a source only for
// repository properties.
func enterpriseRulesetOrganizationPropertyElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the organization property to target.",
			},
			"property_values": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The values to match for the organization property.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// enterpriseRulesetOrganizationSelectors and enterpriseRulesetRepositorySelectors encode
// the condition combinations the enterprise ruleset API accepts: exactly one organization
// selector combined with exactly one repository selector, for six legal combinations.
// https://docs.github.com/en/rest/enterprise-admin/rules?apiVersion=2022-11-28
var (
	enterpriseRulesetOrganizationSelectors = []string{
		"conditions.0.organization_name",
		"conditions.0.organization_id",
		"conditions.0.organization_property",
	}
	enterpriseRulesetRepositorySelectors = []string{
		"conditions.0.repository_name",
		"conditions.0.repository_property",
	}
)
