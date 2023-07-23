package github

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v53/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubOrganizationRuleset() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationRulesetCreate,
		Read:   resourceGithubOrganizationRulesetRead,
		Update: resourceGithubOrganizationRulesetUpdate,
		Delete: resourceGithubOrganizationRulesetDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
				Description:  "The name of the ruleset.",
			},
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Possible values are `branch` and `tag`.",
			},
			"repository": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the repository to apply rulset to.",
			},
			"enforcement": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"disabled", "active", "evaluate"}, false),
				Description:  "Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.",
			},

			"bypass_actors": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The actors that can bypass the rules in this ruleset.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"actor_id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The ID of the actor that can bypass a ruleset",
						},
						"actor_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.",
						},
						"bypass_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"always", "pull_request"}, false),
							Description:  "When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.",
						},
					}},
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
									"inlcude": {
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
						"repository_name": {
							Type:          schema.TypeList,
							Optional:      true,
							MaxItems:      1,
							ConflictsWith: []string{"conditions.0.repository_id"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inlcude": {
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
										Description: "Whether renaming of target repositories is prevented.",
									},
								},
							},
						},
						"repository_id": {
							Type:          schema.TypeList,
							Optional:      true,
							MaxItems:      1,
							ConflictsWith: []string{"conditions.0.repository_name"},
							Description:   "The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
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
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
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
						"required_status_checks": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required_check": {
										Type:        schema.TypeList,
										MinItems:    1,
										Required:    true,
										Description: "Status checks that are required.",
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
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
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
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
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
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
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
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Parameters to be used for the branch_name_pattern rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
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
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Parameters to be used for the tag_name_pattern rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
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
					},
				},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		// CustomizeDiff: customDiffFunction,
	}
}

func resourceGithubOrganizationRulesetObject(d *schema.ResourceData) *github.Ruleset {
	return &github.Ruleset{
		Name:         d.Get("name").(string),
		Target:       github.String(d.Get("target").(string)),
		Source:       d.Get("repository").(string),
		Enforcement:  d.Get("enforcement").(string),
		BypassActors: expandOrganizationBypassActors(d.Get("bypass_actors").([]interface{})),
		Conditions:   expandOrganizationConditions(d.Get("conditions").([]interface{})),
		Rules:        expandOrganizationRules(d.Get("rules").([]interface{})),
	}
}

func resourceGithubOrganizationRulesetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	rulesetReq := resourceGithubOrganizationRulesetObject(d)

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	ctx := context.Background()

	var ruleset *github.Ruleset
	var err error

	ruleset, _, err = client.Repositories.CreateRuleset(ctx, owner, repoName, rulesetReq)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(*ruleset.ID, 10))

	return resourceGithubOrganizationRulesetRead(d, meta)
}

func resourceGithubOrganizationRulesetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	rulesetID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	// When the user has not authenticated the provider, AnonymousHTTPClient is used, therefore owner == "". In this
	// case lookup the owner in the data, and use that, if present.
	if explicitOwner, _, ok := resourceGithubParseFullName(d); ok && owner == "" {
		owner = explicitOwner
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	var ruleset *github.Ruleset
	var resp *github.Response

	ruleset, resp, err = client.Repositories.GetRuleset(ctx, owner, repoName, rulesetID, false)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing ruleset %s/%s: %d from state because it no longer exists in GitHub",
					owner, repoName, rulesetID)
				d.SetId("")
				return nil
			}
		}
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("name", ruleset.Name)
	d.Set("target", ruleset.GetTarget())
	d.Set("enforcement", ruleset.Enforcement)
	d.Set("bypass_actors", flattenBypassActors(ruleset.BypassActors))
	d.Set("conditions", flattenConditions(ruleset.GetConditions()))
	d.Set("rules", flattenRules(ruleset.Rules))
	d.Set("node_id", ruleset.GetNodeID())
	d.Set("ruleset_id", ruleset.ID)

	return nil
}

func resourceGithubOrganizationRulesetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	rulesetReq := resourceGithubRulesetObject(d)

	repoName := d.Get("repository").(string)
	owner := meta.(*Owner).name
	rulesetID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	ruleset, _, err := client.Repositories.UpdateRuleset(ctx, owner, repoName, rulesetID, rulesetReq)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(*ruleset.ID, 10))

	return resourceGithubOrganizationRulesetRead(d, meta)
}

func resourceGithubOrganizationRulesetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	repoName := d.Get("repository").(string)
	owner := meta.(*Owner).name
	rulesetID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting repository ruleset: %s/%s: %d", owner, repoName, rulesetID)
	_, err = client.Repositories.DeleteRuleset(ctx, owner, repoName, rulesetID)
	return err
}

func expandOrganizationBypassActors(input []interface{}) []*github.BypassActor {
	if len(input) == 0 {
		return nil
	}
	bypassActors := make([]*github.BypassActor, 0)

	for _, v := range input {
		inputMap := v.(map[string]interface{})
		actor := &github.BypassActor{}
		if v, ok := inputMap["actor_id"].(int); ok {
			actor.ActorID = github.Int64(int64(v))
		}

		if v, ok := inputMap["actor_type"].(string); ok {
			actor.ActorType = &v
		}

		if v, ok := inputMap["bypass_mode"].(string); ok {
			actor.BypassMode = &v
		}
		bypassActors = append(bypassActors, actor)
	}

	return bypassActors
}

func flattenOrganizationBypassActors(bypassActors []*github.BypassActor) []interface{} {
	if bypassActors == nil {
		return []interface{}{}
	}

	actorsSlice := make([]map[string]interface{}, 0)
	for _, v := range bypassActors {
		actorMap := make(map[string]interface{})

		actorMap["actor_id"] = v.GetActorID()
		actorMap["actor_type"] = v.GetActorType()
		actorMap["bypass_mode"] = v.GetBypassMode()

		actorsSlice = append(actorsSlice, actorMap)
	}

	return []interface{}{actorsSlice}
}

func expandOrganizationConditions(input []interface{}) *github.RulesetConditions {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	rulesetConditions := &github.RulesetConditions{}

	inputConditions := input[0].(map[string]interface{})

	// ref_name
	if v, ok := inputConditions["ref_name"].([]interface{}); ok && v != nil {
		inputRefName := v[0].(map[string]interface{})
		include := make([]string, 0)
		exclude := make([]string, 0)

		for _, v := range inputRefName["include"].([]interface{}) {
			if v != nil {
				include = append(include, v.(string))
			}
		}

		for _, v := range inputRefName["exclude"].([]interface{}) {
			if v != nil {
				exclude = append(exclude, v.(string))
			}
		}

		rulesetConditions.RefName = &github.RulesetRefConditionParameters{
			Include: include,
			Exclude: exclude,
		}
	}

	// repository_name
	if v, ok := inputConditions["repository_name"].([]interface{}); ok && v != nil {
		inputRepositoryName := v[0].(map[string]interface{})
		include := make([]string, 0)
		exclude := make([]string, 0)

		for _, v := range inputRepositoryName["include"].([]interface{}) {
			if v != nil {
				include = append(include, v.(string))
			}
		}

		for _, v := range inputRepositoryName["exclude"].([]interface{}) {
			if v != nil {
				exclude = append(exclude, v.(string))
			}
		}

		protected := inputRepositoryName["protected"].(bool)

		rulesetConditions.RepositoryName = &github.RulesetRepositoryNamesConditionParameters{
			Include:   include,
			Exclude:   exclude,
			Protected: &protected,
		}
	}

	// repository_id
	if v, ok := inputConditions["repository_id"].([]interface{}); ok && v != nil {
		repositoryIDs := make([]int64, 0)

		for _, v := range v {
			if v != nil {
				repositoryIDs = append(repositoryIDs, int64(v.(int)))
			}
		}

		rulesetConditions.RepositoryID = &github.RulesetRepositoryIDsConditionParameters{RepositoryIDs: repositoryIDs}
	}

	return rulesetConditions
}

func flattenOrganizationConditions(conditions *github.RulesetConditions) []interface{} {
	if conditions == nil || conditions.RefName == nil {
		return []interface{}{}
	}

	conditionsMap := make(map[string]interface{})
	refNameSlice := make([]map[string]interface{}, 0)
	repositoryNameSlice := make([]map[string]interface{}, 0)

	refNameSlice = append(refNameSlice, map[string]interface{}{
		"include": conditions.RefName.Include,
		"exclude": conditions.RefName.Exclude,
	})

	if conditions.RepositoryName != nil {
		repositoryNameSlice = append(refNameSlice, map[string]interface{}{
			"include":   conditions.RepositoryName.Include,
			"exclude":   conditions.RepositoryName.Exclude,
			"protected": *conditions.RepositoryName.Protected,
		})
		conditionsMap["repository_name"] = repositoryNameSlice
	}

	if conditions.RepositoryID != nil {
		conditionsMap["repository_id"] = conditions.RepositoryID.RepositoryIDs
	}

	conditionsMap["ref_name"] = refNameSlice

	return []interface{}{conditionsMap}
}

func expandOrganizationRules(input []interface{}) []*github.RepositoryRule {
	if len(input) == 0 || input[0] == nil {
		return nil
	}

	rulesMap := input[0].(map[string]interface{})
	rulesSlice := make([]*github.RepositoryRule, 0)

	// First we expand rules without parameters
	if v, ok := rulesMap["creation"].(bool); ok && v {
		rulesSlice = append(rulesSlice, github.NewCreationRule())
	}

	if v, ok := rulesMap["update"].(bool); ok && v {
		rulesSlice = append(rulesSlice, github.NewUpdateRule())
	}

	if v, ok := rulesMap["deletion"].(bool); ok && v {
		rulesSlice = append(rulesSlice, github.NewDeletionRule())
	}

	if v, ok := rulesMap["required_linear_history"].(bool); ok && v {
		rulesSlice = append(rulesSlice, github.NewRequiredLinearHistoryRule())
	}

	if v, ok := rulesMap["required_signatures"].(bool); ok && v {
		rulesSlice = append(rulesSlice, github.NewRequiredSignaturesRule())
	}

	if v, ok := rulesMap["non_fast_forward"].(bool); ok && v {
		rulesSlice = append(rulesSlice, github.NewNonFastForwardRule())
	}

	// Pattern parameter rules
	for _, k := range []string{"commit_message_pattern", "commit_author_email_pattern", "committer_email_pattern", "branch_name_pattern", "tag_name_pattern"} {
		if v, ok := rulesMap[k].([]interface{}); ok && len(v) != 0 {
			patternParametersMap := v[0].(map[string]interface{})
			if enabled, ok := patternParametersMap["enabled"].(bool); ok && enabled {

				name := patternParametersMap["name"].(string)
				negate := patternParametersMap["negate"].(bool)

				params := &github.RulePatternParameters{
					Name:     &name,
					Negate:   &negate,
					Operator: patternParametersMap["operator"].(string),
					Pattern:  patternParametersMap["pattern"].(string),
				}

				switch k {
				case "commit_message_pattern":
					rulesSlice = append(rulesSlice, github.NewCommitMessagePatternRule(params))
				case "commit_author_email_pattern":
					rulesSlice = append(rulesSlice, github.NewCommitAuthorEmailPatternRule(params))
				case "committer_email_pattern":
					rulesSlice = append(rulesSlice, github.NewCommitterEmailPatternRule(params))
				case "branch_name_pattern":
					rulesSlice = append(rulesSlice, github.NewBranchNamePatternRule(params))
				case "tag_name_pattern":
					rulesSlice = append(rulesSlice, github.NewTagNamePatternRule(params))
				}
			}
		}
	}

	// Pull request rule
	if v, ok := rulesMap["pull_request"].([]interface{}); ok && len(v) != 0 {
		pullRequestMap := v[0].(map[string]interface{})
		if enabled, ok := pullRequestMap["enabled"].(bool); ok && enabled {
			params := &github.PullRequestRuleParameters{
				DismissStaleReviewsOnPush:      pullRequestMap["dismiss_stale_reviews_on_push"].(bool),
				RequireCodeOwnerReview:         pullRequestMap["require_code_owner_review"].(bool),
				RequireLastPushApproval:        pullRequestMap["require_last_push_approval"].(bool),
				RequiredApprovingReviewCount:   pullRequestMap["required_approving_review_count"].(int),
				RequiredReviewThreadResolution: pullRequestMap["required_review_thread_resolution"].(bool),
			}

			rulesSlice = append(rulesSlice, github.NewPullRequestRule(params))
		}
	}

	// Required status checks rule
	if v, ok := rulesMap["required_status_checks"].([]interface{}); ok && len(v) != 0 {
		requiredStatusMap := v[0].(map[string]interface{})
		if enabled, ok := requiredStatusMap["enabled"].(bool); ok && enabled {
			requiredStatusChecks := make([]github.RuleRequiredStatusChecks, 0)
			if requiredStatusChecksInput, ok := requiredStatusMap["required_check"].([]map[string]interface{}); ok {
				for _, check := range requiredStatusChecksInput {
					requiredStatusChecks = append(requiredStatusChecks, github.RuleRequiredStatusChecks{
						Context:       check["context"].(string),
						IntegrationID: check["integration_id"].(*int64),
					})
				}
			}

			params := &github.RequiredStatusChecksRuleParameters{
				RequiredStatusChecks:             requiredStatusChecks,
				StrictRequiredStatusChecksPolicy: requiredStatusMap["strict_required_status_checks_policy"].(bool),
			}
			rulesSlice = append(rulesSlice, github.NewRequiredStatusChecksRule(params))
		}
	}

	return rulesSlice
}

func flattenOrganizationRules(rules []*github.RepositoryRule) []interface{} {
	if len(rules) == 0 || rules == nil {
		return []interface{}{}
	}

	rulesMap := make(map[string]interface{})
	for _, v := range rules {
		switch v.Type {
		case "creation", "update", "deletion", "required_linear_history", "required_signatures", "non_fast_forward":
			rulesMap[v.Type] = true

		case "commit_message_pattern", "commit_author_email_pattern", "committer_email_pattern", "branch_name_pattern", "tag_name_pattern":
			var params github.RulePatternParameters

			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}

			rule := make(map[string]interface{})
			rule["name"] = *params.Name
			rule["negate"] = *params.Negate
			rule["operator"] = params.Operator
			rule["pattern"] = params.Pattern
			rulesMap[v.Type] = rule

		case "pull_request":
			var params github.PullRequestRuleParameters

			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}

			rule := make(map[string]interface{})
			rule["dismiss_stale_reviews_on_push"] = params.DismissStaleReviewsOnPush
			rule["require_code_owner_review"] = params.RequireCodeOwnerReview
			rule["require_last_push_approval"] = params.RequireLastPushApproval
			rule["required_approving_review_count"] = params.RequiredApprovingReviewCount
			rule["required_review_thread_resolution"] = params.RequiredReviewThreadResolution
			rulesMap[v.Type] = rule

		case "required_status_checks":
			var params github.RequiredStatusChecksRuleParameters

			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}

			requiredStatusChecksSlice := make([]map[string]interface{}, 0)
			for _, check := range params.RequiredStatusChecks {
				requiredStatusChecksSlice = append(requiredStatusChecksSlice, map[string]interface{}{
					"context":        check.Context,
					"integration_id": check.IntegrationID,
				})
			}

			rule := make(map[string]interface{})
			rule["required_check"] = requiredStatusChecksSlice
			rule["strict_required_status_checks_policy"] = params.StrictRequiredStatusChecksPolicy
			rulesMap[v.Type] = rule
		}
	}

	return []interface{}{rulesMap}
}
