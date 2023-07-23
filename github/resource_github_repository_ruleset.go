package github

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/go-github/v53/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubRepositoryRuleset() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryRulesetCreate,
		Read:   resourceGithubRepositoryRulesetRead,
		Update: resourceGithubRepositoryRulesetUpdate, 
		Delete: resourceGithubRepositoryRulesetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the repository to add the ruleset to.",
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Owner of the repository. If not provided, the provider's default owner is used.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Ruleset within the repository.",
			},
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The target of the ruleset. Either branch or tag.",
				ValidateFunc: validation.StringInSlice([]string{
					"branch",
					"tag",
				}, false),
			},
			"enforcement": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The enforcement level of the ruleset. One of active, disabled or evaluate. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).",
				ValidateFunc: validation.StringInSlice([]string{
					"active",
					"disabled",
					"evaluate",
				}, false),
			},
			"conditions": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Target branches/tags. Both an inclusion and exclusion list, supporting regexes as well as ALL branches/tags and the default branch",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// TODO: Should the default branch + ALL branches/tags have it's own field?
						"include": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"exclude": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"rule_creation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Only allow users with bypass permission to create matching refs.",
			},
			// TODO: Borken currently. When underlying sdk gets a version bump, its fixed
			// "rule_update": {
			// 	Type:        schema.TypeBool,
			// 	Optional:    true,
			// 	ForceNew:    true, // TODO: Remove this when updating is implemented
			// 	Description: "Only allow users with bypass permission to update matching refs.",
			// },
			"rule_deletion": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Only allow users with bypass permissions to delete matching refs.",
			},
			"rule_required_linear_history": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Prevent merge commits from being pushed to matching branches.",
			},
			"rule_required_signatures": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Commits pushed to matching branches must have verified signatures.",
			},
			"rule_non_fast_forward": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Prevent users with push access from force pushing to branches.",
			},
			"rule_required_deployments": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"rule_pull_request": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dismiss_stale_reviews_on_push": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "New, reviewable commits pushed will dismiss previous pull request review approvals.",
						},
						"require_code_owner_review": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Require an approving review in pull requests that modify files that have a designated code owner.",
						},
						"require_last_push_approval": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether the most recent reviewable push must be approved by someone other than the person who pushed it.",
						},
						"required_approving_review_count": {
							Type:         schema.TypeInt,
							Required:     true,
							Description:  "The number of approving reviews that are required before a pull request can be merged.",
							ValidateFunc: validation.IntBetween(0, 10),
						},
						"required_review_thread_resolution": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "All conversations on code must be resolved before a pull request can be merged.",
						},
					},
				},
			},
			"rule_required_status_checks": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"strict_required_status_checks_policy": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled.",
						},
						"required_status_checks": {
							// This field is based on "checks" in github_branch_protection_v3, which combines the context and integration_id (or app_id which it's called in branch_protection rules) into one string
							Type:        schema.TypeSet,
							Required:    true,
							Description: "The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the 'context' and 'integration_id' like so 'context:integration_id'. Also supports only the 'context'",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			// TODO: Broken until a bump in the underlying sdk
			// "": {
			// 	Type:        schema.TypeSet,
			// 	Optional:    true,
			// 	ForceNew:    true, // TODO: Remove this when updating is implemented
			// 	Description: "A list of actors that can bypass rules in a ruleset.",
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			// TODO: Is there a better API for this upcoming? Currently you have to set the bypass list by ID, which is not really user friendly
			// 			"actor_id": {
			// 				Type:        schema.TypeInt,
			// 				Description: "The ID of the actor that can bypass a ruleset",
			// 				Required:    true,
			// 			},
			// 			"actor_type": {
			// 				Type:        schema.TypeString,
			// 				Description: "The type of actor that can bypass a ruleset. One of RepositoryRole, Team, Integration or OrganizationAdmin",
			// 				Required:    true,
			// 				ValidateFunc: validation.StringInSlice([]string{
			// 					"RepositoryRole",
			// 					"Team",
			// 					"Integration",
			// 					"OrganizationAdmin",
			// 				}, false),
			// 			},
			// 			// TODO: Needs a bump in the underlying sdk
			// 			// "bypass_mode": {
			// 			// 	Type:     schema.TypeString,
			// 			// 	Required: true,
			// 			// 	Description: "When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests.",
			// 			// 	ValidateFunc: validation.StringInSlice([]string{
			// 			// 		"always",
			// 			// 		"pull_request",	  
			// 			// 	}, false),
			// 			// },
			// 		},
			// 	},
			// },
			"ruleset_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The id of the Ruleset within the repository.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the source of the ruleset. Either Repository or Organization (for this resource it will always be Repository).",
			},
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Source of the ruleset (OWNER/REPO).",
			},
		},
	}
}

func resourceGithubRepositoryRulesetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	if explicitOwner, ok := d.GetOk("owner"); ok {
		owner = explicitOwner.(string)
	}

	repository := d.Get("repository").(string)

	sourceType := "Repository"
	rulesetRequest, err := buildRulesetRequest(d, &sourceType)
	if err != nil {
		return err
	}
	ctx := context.Background()

	ruleset, _, err := client.Repositories.CreateRuleset(ctx, owner, repository, rulesetRequest)
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(owner, repository, strconv.FormatInt(ruleset.ID, 10)))

	return resourceGithubRepositoryRulesetRead(d, meta)
}

func resourceGithubRepositoryRulesetRead(d *schema.ResourceData, meta interface{}) error {
	ctx := context.TODO()
	client := meta.(*Owner).v3client

	owner, repository, id, err := parseRulesetID(d)
	if err != nil {
		return err
	}

	ruleset, _, err := client.Repositories.GetRuleset(ctx, owner, repository, id, false)
	if err != nil {
		return err
	}

	if err := flattenAndSetRulesetConditions(d, ruleset); err != nil {
		return fmt.Errorf("error setting conditions: %v", err)
	}

	rules_toggleable := map[string]bool{
		"creation":                false,
		// "update":                  false, // TODO: Currently broken
		"deletion":                false,
		"required_linear_history": false,
		"required_signatures":     false,
		"non_fast_forward":        false,
	}

	for _, rule := range ruleset.Rules {
		switch rule_type := rule.Type; rule_type {
		case "required_deployments":
			var rderParams *github.RequiredDeploymentEnvironmentsRuleParameters
			err := json.Unmarshal(rule.GetParameters(), &rderParams)
			if err != nil {
				return err
			}
			d.Set("rule_required_deployments", rderParams.RequiredDeploymentEnvironments)

		case "pull_request":
			var prrParams *github.PullRequestRuleParameters
			err := json.Unmarshal(rule.GetParameters(), &prrParams)
			if err != nil {
				return err
			}

			d.Set("rule_pull_request", []interface{}{
				map[string]interface{}{
					"dismiss_stale_reviews_on_push": prrParams.DismissStaleReviewsOnPush,
					"require_code_owner_review": prrParams.RequireCodeOwnerReview,
					"require_last_push_approval": prrParams.RequireLastPushApproval,
					"required_approving_review_count": prrParams.RequiredApprovingReviewCount,
					"required_review_thread_resolution": prrParams.RequiredReviewThreadResolution,
				},
			})

		case "required_status_checks":
			var rscrParams *github.RequiredStatusChecksRuleParameters
			err := json.Unmarshal(rule.GetParameters(), &rscrParams)
			if err != nil {
				return err
			}

			var checks []interface{}
			for _, chk := range rscrParams.RequiredStatusChecks {
				if chk.GetIntegrationID() != 0 {
					checks = append(checks, fmt.Sprintf("%s:%d", chk.Context, chk.GetIntegrationID()))
				} else {
					checks = append(checks, chk.Context)
				}
			}

			d.Set("rule_required_status_checks", []interface{}{
				map[string]interface{}{
					"strict_required_status_checks_policy": rscrParams.StrictRequiredStatusChecksPolicy,
					"required_status_checks": schema.NewSet(schema.HashString, checks),
				},
			})

		default:
			// TODO: Is there a better way of doing this?
			if _, ok := rules_toggleable[rule_type]; !ok {
				return fmt.Errorf("unexpected rule %q", rule_type)
			}

			d.Set(fmt.Sprintf("rule_%s", rule_type), true)
		}
	}

	// TODO: Waiting until go-github has been bumped 
	// d.Set("", ruleset.BypassActors)

	d.Set("ruleset_id", ruleset.ID)
	d.Set("repository", repository)
	d.Set("owner", owner)
	d.Set("name", ruleset.Name)
	d.Set("target", ruleset.GetTarget())
	d.Set("enforcement", ruleset.Enforcement)
	d.Set("source_type", ruleset.GetSourceType())
	d.Set("source", ruleset.Source)

	return nil
}

func resourceGithubRepositoryRulesetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner, repository, id, err := parseRulesetID(d)
	if err != nil {
		return err
	}

	sourceType := "Repository"
	rulesetRequest, err := buildRulesetRequest(d, &sourceType)
	if err != nil {
		return err
	}

	// TODO: wtf even is this? I have no idea what golang contexts are
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	ruleset, _, err := client.Repositories.UpdateRuleset(ctx, owner, repository, id, rulesetRequest)
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(owner, repository, strconv.FormatInt(ruleset.ID, 10)))

	return resourceGithubRepositoryRulesetRead(d, meta)
}

func resourceGithubRepositoryRulesetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner, repository, id, err := parseRulesetID(d)
	if err != nil {
		return err
	}

	_, err = client.Repositories.DeleteRuleset(context.TODO(), owner, repository, id)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func parseRulesetID(d *schema.ResourceData) (owner, repository string, id int64, err error) {
	var strNumber string

	if owner, repository, strNumber, err = parseThreePartID(d.Id(), "owner", "repository", "ruleset_id"); err != nil {
		return
	}

	if id, err = strconv.ParseInt(strNumber, 10, 64); err != nil {
		err = fmt.Errorf("invalid Ruleset number %s: %w", strNumber, err)
	}

	return
}
