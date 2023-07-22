package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

		// TODO: Implement this
		// Importer: &schema.ResourceImporter{
		// 	State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
		// 		_, baseRepository, _, err := parsePullRequestID(d)
		// 		if err != nil {
		// 			return nil, err
		// 		}
		// 		d.Set("base_repository", baseRepository)

		// 		return []*schema.ResourceData{d}, nil
		// 	},
		// },

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
			"rules": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The rules that the ruleset will enforce",
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
						"non_fast_forward": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Prevent users with push access from force pushing to branches.",
						},
						"required_deployments": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"pull_request": {
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
						"required_status_checks": {
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
										// TODO: The API spec says that the context is the only required field. Maybe change this implementation to only use the context and ignore the integration_id?
										Type:        schema.TypeSet,
										Required:    true,
										Description: "The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the 'context' and 'integration_id' like so 'context:integration_id'",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						// TODO: There seems to be some rules available in https://github.com/github/rest-api-description/blob/main/descriptions/api.github.com/api.github.com.2022-11-28.json which are not available through the UI. I'll skip these for now

					},
				},
			},
			"bypass_actors": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of actors that can bypass rules in a ruleset.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// TODO: Is there a better API for this upcoming? Currently you have to set the bypass list by ID, which is not really user friendly
						"actor_id": {
							Type:        schema.TypeInt,
							Description: "The ID of the actor that can bypass a ruleset",
							Required:    true,
						},
						"actor_type": {
							Type:        schema.TypeString,
							Description: "The type of actor that can bypass a ruleset. One of RepositoryRole, Team, Integration or OrganizationAdmin",
							Required:    true,
							ValidateFunc: validation.StringInSlice([]string{
								"RepositoryRole",
								"Team",
								"Integration",
								"OrganizationAdmin",
							}, false),
						},
						"bypass_mode": {
							Type:     schema.TypeString,
							Required: true,
							Description: "When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests.",
							ValidateFunc: validation.StringInSlice([]string{
								"always",
								"pull_request",	  
							}, false),
						},
					},
				},
			},
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
			// TODO: Look into BypassMode
		},
	}
}

func resourceGithubRepositoryRulesetCreate(d *schema.ResourceData, meta interface{}) error {
	ctx := context.TODO()
	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	var ruleset *github.Ruleset

	client.Repositories.CreateRuleset(ctx, orgName, repoName, ruleset)

	// For convenience, by default we expect that the base repository and head
	// repository owners are the same, and both belong to the caller, indicating
	// a "PR within the same repo" scenario. The head will *always* belong to
	// the current caller, the base - not necessarily. The base will belong to
	// another namespace in case of forks, and this resource supports them.
	headOwner := meta.(*Owner).name

	baseOwner := headOwner
	if explicitBaseOwner, ok := d.GetOk("owner"); ok {
		baseOwner = explicitBaseOwner.(string)
	}

	baseRepository := d.Get("base_repository").(string)

	head := d.Get("head_ref").(string)
	if headOwner != baseOwner {
		head = strings.Join([]string{headOwner, head}, ":")
	}

	pullRequest, _, err := client.PullRequests.Create(ctx, baseOwner, baseRepository, &github.NewPullRequest{
		Title:               github.String(d.Get("title").(string)),
		Head:                github.String(head),
		Base:                github.String(d.Get("base_ref").(string)),
		Body:                github.String(d.Get("body").(string)),
		MaintainerCanModify: github.Bool(d.Get("maintainer_can_modify").(bool)),
	})

	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(baseOwner, baseRepository, strconv.Itoa(pullRequest.GetNumber())))

	return resourceGithubRepositoryPullRequestRead(d, meta)
}

func resourceGithubRepositoryRulesetRead(d *schema.ResourceData, meta interface{}) error {
	// TODO: Implement
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

	conditions := []interface{}{
		map[string]interface{}{
			"include": ruleset.GetConditions().RefName.Include,
			"exclude": ruleset.GetConditions().RefName.Exclude,
		},
	}

	rules_toggleable := map[string]bool{
		"creation":                false,
		"update":                  false,
		"deletion":                false,
		"required_linear_history": false,
		"required_signatures":     false,
		"non_fast_forward":        false,
	}

	for _, rule := range ruleset.Rules {
		switch rule_type := rule.Type; rule_type {
		case "required_deployments":

			rule.GetParameters()

			fmt.Println("TODO: Implement this")
		case "pull_request":
			fmt.Println("TODO: Implement this")
		case "required_status_checks":
			fmt.Println("TODO: Implement this")

		default:
			// TODO: Is there a better way of doing this?
			if _, ok := rules_toggleable[rule_type]; !ok {
				return fmt.Errorf("Unexpected rule %q.", rule_type)
			}

			rules_toggleable[rule_type] = true
		}
	}

	d.Set("ruleset_id", ruleset.ID)
	d.Set("name", ruleset.Name)
	d.Set("target", ruleset.GetTarget())
	d.Set("enforcement", ruleset.Enforcement)
	d.Set("conditions", conditions)
	// d.Set("rules", rules)
	// d.Set("bypass_actors", ruleset.BypassActors)
	d.Set("source_type", ruleset.GetSourceType())
	d.Set("source", ruleset.Source)

	return nil
}

func resourceGithubRepositoryRulesetUpdate(d *schema.ResourceData, meta interface{}) error {
	// TODO: Implement
	ctx := context.TODO()
	client := meta.(*Owner).v3client

	owner, repository, number, err := parsePullRequestID(d)
	if err != nil {
		return err
	}

	update := &github.PullRequest{
		Title:               github.String(d.Get("title").(string)),
		Body:                github.String(d.Get("body").(string)),
		MaintainerCanModify: github.Bool(d.Get("maintainer_can_modify").(bool)),
	}

	if d.HasChange("base_ref") {
		update.Base = &github.PullRequestBranch{
			Ref: github.String(d.Get("base_ref").(string)),
		}
	}

	_, _, err = client.PullRequests.Edit(ctx, owner, repository, number, update)
	if err == nil {
		return resourceGithubRepositoryPullRequestRead(d, meta)
	}

	errors := []string{fmt.Sprintf("could not update the Pull Request: %v", err)}

	if err := resourceGithubRepositoryPullRequestRead(d, meta); err != nil {
		errors = append(errors, fmt.Sprintf("could not read the Pull Request after the failed update: %v", err))
	}

	return fmt.Errorf(strings.Join(errors, ", "))
}

func resourceGithubRepositoryRulesetDelete(d *schema.ResourceData, meta interface{}) error {
	// TODO: Implement this

	// It's not entirely clear how to treat PR deletion according to Terraform's
	// CRUD semantics. The approach we're taking here is to close the PR unless
	// it's already closed or merged. Merging it feels intuitively wrong in what
	// effectively is a destructor.
	if d.Get("state").(string) != "open" {
		d.SetId("")
		return nil
	}

	ctx := context.TODO()
	client := meta.(*Owner).v3client

	owner, repository, number, err := parsePullRequestID(d)
	if err != nil {
		return err
	}

	update := &github.PullRequest{State: github.String("closed")}
	if _, _, err = client.PullRequests.Edit(ctx, owner, repository, number, update); err != nil {
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
		err = fmt.Errorf("invalid PR number %s: %w", strNumber, err)
	}

	return
}
