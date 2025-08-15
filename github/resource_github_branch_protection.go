package github

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"
)

func resourceGithubBranchProtection() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 2,

		Schema: map[string]*schema.Schema{
			// Input
			REPOSITORY_ID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name or node ID of the repository associated with this branch protection rule.",
			},
			PROTECTION_PATTERN: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identifies the protection rule pattern.",
			},
			PROTECTION_ALLOWS_DELETIONS: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' to allow the branch to be deleted.",
			},
			PROTECTION_ALLOWS_FORCE_PUSHES: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' to allow force pushes on the branch.",
			},
			PROTECTION_IS_ADMIN_ENFORCED: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' enforces status checks for repository administrators.",
			},
			PROTECTION_REQUIRES_COMMIT_SIGNATURES: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' requires all commits to be signed with GPG.",
			},
			PROTECTION_REQUIRES_LINEAR_HISTORY: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch.",
			},
			PROTECTION_REQUIRES_CONVERSATION_RESOLUTION: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' requires all conversations on code must be resolved before a pull request can be merged.",
			},
			PROTECTION_LOCK_BRANCH: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to 'true' will make the branch read-only and preventing any pushes to it.",
			},
			PROTECTION_REQUIRES_APPROVING_REVIEWS: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Enforce restrictions for pull request reviews.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PROTECTION_REQUIRED_APPROVING_REVIEW_COUNT: {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          1,
							Description:      "Require 'x' number of approvals to satisfy branch protection requirements. If this is specified it must be a number between 0-6.",
							ValidateDiagFunc: toDiagFunc(validation.IntBetween(0, 6), PROTECTION_REQUIRED_APPROVING_REVIEW_COUNT),
						},
						PROTECTION_REQUIRES_CODE_OWNER_REVIEWS: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Require an approved review in pull requests including files with a designated code owner.",
						},
						PROTECTION_DISMISSES_STALE_REVIEWS: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Dismiss approved reviews automatically when a new commit is pushed.",
						},
						PROTECTION_RESTRICTS_REVIEW_DISMISSALS: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Restrict pull request review dismissals.",
						},
						PROTECTION_REVIEW_DISMISSAL_ALLOWANCES: {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of actor Names/IDs with dismissal access. If not empty, 'restrict_dismissals' is ignored. Actor names must either begin with a '/' for users or the organization name followed by a '/' for teams.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						PROTECTION_PULL_REQUESTS_BYPASSERS: {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of actor Names/IDs that are allowed to bypass pull request requirements. Actor names must either begin with a '/' for users or the organization name followed by a '/' for teams.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						PROTECTION_REQUIRE_LAST_PUSH_APPROVAL: {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Require that The most recent push must be approved by someone other than the last pusher.",
						},
					},
				},
			},
			PROTECTION_REQUIRES_STATUS_CHECKS: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Enforce restrictions for required status checks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PROTECTION_REQUIRES_STRICT_STATUS_CHECKS: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Require branches to be up to date before merging.",
						},
						PROTECTION_REQUIRED_STATUS_CHECK_CONTEXTS: {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of status checks to require in order to merge into this branch. No status checks are required by default.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			PROTECTION_RESTRICTS_PUSHES: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Restrict who can push to matching branches.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PROTECTION_BLOCKS_CREATIONS: {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Restrict pushes that create matching branches.",
						},
						PROTECTION_PUSH_ALLOWANCES: {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of actor Names/IDs that may push to the branch. Actor names must either begin with a '/' for users or the organization name followed by a '/' for teams.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			PROTECTION_FORCE_PUSHES_BYPASSERS: {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a '/' for users or the organization name followed by a '/' for teams.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},

		Create: resourceGithubBranchProtectionCreate,
		Read:   resourceGithubBranchProtectionRead,
		Update: resourceGithubBranchProtectionUpdate,
		Delete: resourceGithubBranchProtectionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGithubBranchProtectionImport,
		},

		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubBranchProtectionV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubBranchProtectionUpgradeV0,
				Version: 0,
			},
			{
				Type:    resourceGithubBranchProtectionV1().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubBranchProtectionUpgradeV1,
				Version: 1,
			},
		},
	}
}

func resourceGithubBranchProtectionCreate(d *schema.ResourceData, meta any) error {
	var mutate struct {
		CreateBranchProtectionRule struct {
			BranchProtectionRule struct {
				ID githubv4.ID
			}
		} `graphql:"createBranchProtectionRule(input: $input)"`
	}
	data, err := branchProtectionResourceData(d, meta)
	if err != nil {
		return err
	}

	var reviewIds, pushIds, bypassForcePushIds, bypassPullRequestIds []string
	reviewIds, err = getActorIds(data.ReviewDismissalActorIDs, meta)
	if err != nil {
		return err
	}

	pushIds, err = getActorIds(data.PushActorIDs, meta)
	if err != nil {
		return err
	}

	bypassForcePushIds, err = getActorIds(data.BypassForcePushActorIDs, meta)
	if err != nil {
		return err
	}

	bypassPullRequestIds, err = getActorIds(data.BypassPullRequestActorIDs, meta)
	if err != nil {
		return err
	}

	data.PushActorIDs = pushIds
	data.ReviewDismissalActorIDs = reviewIds
	data.BypassForcePushActorIDs = bypassForcePushIds
	data.BypassPullRequestActorIDs = bypassPullRequestIds

	input := githubv4.CreateBranchProtectionRuleInput{
		AllowsDeletions:                githubv4.NewBoolean(githubv4.Boolean(data.AllowsDeletions)),
		AllowsForcePushes:              githubv4.NewBoolean(githubv4.Boolean(data.AllowsForcePushes)),
		BlocksCreations:                githubv4.NewBoolean(githubv4.Boolean(data.BlocksCreations)),
		BypassForcePushActorIDs:        githubv4NewIDSlice(githubv4IDSliceEmpty(data.BypassForcePushActorIDs)),
		BypassPullRequestActorIDs:      githubv4NewIDSlice(githubv4IDSliceEmpty(data.BypassPullRequestActorIDs)),
		DismissesStaleReviews:          githubv4.NewBoolean(githubv4.Boolean(data.DismissesStaleReviews)),
		IsAdminEnforced:                githubv4.NewBoolean(githubv4.Boolean(data.IsAdminEnforced)),
		Pattern:                        githubv4.String(data.Pattern),
		PushActorIDs:                   githubv4NewIDSlice(githubv4IDSlice(data.PushActorIDs)),
		RepositoryID:                   githubv4.NewID(githubv4.ID(data.RepositoryID)),
		RequiredApprovingReviewCount:   githubv4.NewInt(githubv4.Int(data.RequiredApprovingReviewCount)),
		RequiredStatusCheckContexts:    githubv4NewStringSlice(githubv4StringSliceEmpty(data.RequiredStatusCheckContexts)),
		RequiresApprovingReviews:       githubv4.NewBoolean(githubv4.Boolean(data.RequiresApprovingReviews)),
		RequiresCodeOwnerReviews:       githubv4.NewBoolean(githubv4.Boolean(data.RequiresCodeOwnerReviews)),
		RequiresCommitSignatures:       githubv4.NewBoolean(githubv4.Boolean(data.RequiresCommitSignatures)),
		RequiresConversationResolution: githubv4.NewBoolean(githubv4.Boolean(data.RequiresConversationResolution)),
		RequiresLinearHistory:          githubv4.NewBoolean(githubv4.Boolean(data.RequiresLinearHistory)),
		RequiresStatusChecks:           githubv4.NewBoolean(githubv4.Boolean(data.RequiresStatusChecks)),
		RequiresStrictStatusChecks:     githubv4.NewBoolean(githubv4.Boolean(data.RequiresStrictStatusChecks)),
		RestrictsPushes:                githubv4.NewBoolean(githubv4.Boolean(data.RestrictsPushes)),
		RestrictsReviewDismissals:      githubv4.NewBoolean(githubv4.Boolean(data.RestrictsReviewDismissals)),
		ReviewDismissalActorIDs:        githubv4NewIDSlice(githubv4IDSlice(data.ReviewDismissalActorIDs)),
		LockBranch:                     githubv4.NewBoolean(githubv4.Boolean(data.LockBranch)),
		RequireLastPushApproval:        githubv4.NewBoolean(githubv4.Boolean(data.RequireLastPushApproval)),
	}

	ctx := context.Background()
	client := meta.(*Owner).v4client
	err = client.Mutate(ctx, &mutate, input, nil)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s", mutate.CreateBranchProtectionRule.BranchProtectionRule.ID))

	return resourceGithubBranchProtectionRead(d, meta)
}

func resourceGithubBranchProtectionRead(d *schema.ResourceData, meta any) error {
	var query struct {
		Node struct {
			Node BranchProtectionRule `graphql:"... on BranchProtectionRule"`
		} `graphql:"node(id: $id)"`
	}
	variables := map[string]any{
		"id": d.Id(),
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	client := meta.(*Owner).v4client
	err := client.Query(ctx, &query, variables)
	if err != nil {
		if strings.Contains(err.Error(), "Could not resolve to a node with the global id") {
			log.Printf("[INFO] Removing branch protection (%s) from state because it no longer exists in GitHub", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}
	protection := query.Node.Node

	err = d.Set(PROTECTION_PATTERN, protection.Pattern)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_PATTERN, protection.Repository.Name, protection.Pattern, d.Id())
	}

	err = d.Set(PROTECTION_ALLOWS_DELETIONS, protection.AllowsDeletions)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_ALLOWS_DELETIONS, protection.Repository.Name, protection.Pattern, d.Id())
	}

	err = d.Set(PROTECTION_ALLOWS_FORCE_PUSHES, protection.AllowsForcePushes)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_ALLOWS_FORCE_PUSHES, protection.Repository.Name, protection.Pattern, d.Id())
	}

	err = d.Set(PROTECTION_IS_ADMIN_ENFORCED, protection.IsAdminEnforced)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_IS_ADMIN_ENFORCED, protection.Repository.Name, protection.Pattern, d.Id())
	}

	err = d.Set(PROTECTION_REQUIRES_COMMIT_SIGNATURES, protection.RequiresCommitSignatures)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_REQUIRES_COMMIT_SIGNATURES, protection.Repository.Name, protection.Pattern, d.Id())
	}

	err = d.Set(PROTECTION_REQUIRES_LINEAR_HISTORY, protection.RequiresLinearHistory)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_REQUIRES_LINEAR_HISTORY, protection.Repository.Name, protection.Pattern, d.Id())
	}

	err = d.Set(PROTECTION_REQUIRES_CONVERSATION_RESOLUTION, protection.RequiresConversationResolution)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_REQUIRES_CONVERSATION_RESOLUTION, protection.Repository.Name, protection.Pattern, d.Id())
	}

	data, err := branchProtectionResourceDataActors(d, meta)
	if err != nil {
		return err
	}

	approvingReviews := setApprovingReviews(protection, data, meta)
	err = d.Set(PROTECTION_REQUIRES_APPROVING_REVIEWS, approvingReviews)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_REQUIRES_APPROVING_REVIEWS, protection.Repository.Name, protection.Pattern, d.Id())
	}

	statusChecks := setStatusChecks(protection)
	err = d.Set(PROTECTION_REQUIRES_STATUS_CHECKS, statusChecks)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_REQUIRES_STATUS_CHECKS, protection.Repository.Name, protection.Pattern, d.Id())
	}

	restrictsPushes := setPushes(protection, data, meta)
	err = d.Set(PROTECTION_RESTRICTS_PUSHES, restrictsPushes)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_RESTRICTS_PUSHES, protection.Repository.Name, protection.Pattern, d.Id())
	}

	forcePushBypassers := setForcePushBypassers(protection, data, meta)
	err = d.Set(PROTECTION_FORCE_PUSHES_BYPASSERS, forcePushBypassers)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_FORCE_PUSHES_BYPASSERS, protection.Repository.Name, protection.Pattern, d.Id())
	}

	err = d.Set(PROTECTION_LOCK_BRANCH, protection.LockBranch)
	if err != nil {
		log.Printf("[DEBUG] Problem setting '%s' in %s %s branch protection (%s)", PROTECTION_LOCK_BRANCH, protection.Repository.Name, protection.Pattern, d.Id())
	}

	return nil
}

func resourceGithubBranchProtectionUpdate(d *schema.ResourceData, meta any) error {
	var mutate struct {
		UpdateBranchProtectionRule struct {
			BranchProtectionRule struct {
				ID githubv4.ID
			}
		} `graphql:"updateBranchProtectionRule(input: $input)"`
	}
	data, err := branchProtectionResourceData(d, meta)
	if err != nil {
		return err
	}

	var reviewIds, pushIds, bypassForcePushIds, bypassPullRequestIds []string
	reviewIds, err = getActorIds(data.ReviewDismissalActorIDs, meta)
	if err != nil {
		return err
	}

	pushIds, err = getActorIds(data.PushActorIDs, meta)
	if err != nil {
		return err
	}

	bypassForcePushIds, err = getActorIds(data.BypassForcePushActorIDs, meta)
	if err != nil {
		return err
	}

	bypassPullRequestIds, err = getActorIds(data.BypassPullRequestActorIDs, meta)
	if err != nil {
		return err
	}

	data.PushActorIDs = pushIds
	data.ReviewDismissalActorIDs = reviewIds
	data.BypassForcePushActorIDs = bypassForcePushIds
	data.BypassPullRequestActorIDs = bypassPullRequestIds

	input := githubv4.UpdateBranchProtectionRuleInput{
		BranchProtectionRuleID:         d.Id(),
		AllowsDeletions:                githubv4.NewBoolean(githubv4.Boolean(data.AllowsDeletions)),
		AllowsForcePushes:              githubv4.NewBoolean(githubv4.Boolean(data.AllowsForcePushes)),
		BlocksCreations:                githubv4.NewBoolean(githubv4.Boolean(data.BlocksCreations)),
		BypassForcePushActorIDs:        githubv4NewIDSlice(githubv4IDSliceEmpty(data.BypassForcePushActorIDs)),
		BypassPullRequestActorIDs:      githubv4NewIDSlice(githubv4IDSliceEmpty(data.BypassPullRequestActorIDs)),
		DismissesStaleReviews:          githubv4.NewBoolean(githubv4.Boolean(data.DismissesStaleReviews)),
		IsAdminEnforced:                githubv4.NewBoolean(githubv4.Boolean(data.IsAdminEnforced)),
		Pattern:                        githubv4.NewString(githubv4.String(data.Pattern)),
		PushActorIDs:                   githubv4NewIDSlice(githubv4IDSlice(data.PushActorIDs)),
		RequiredApprovingReviewCount:   githubv4.NewInt(githubv4.Int(data.RequiredApprovingReviewCount)),
		RequiredStatusCheckContexts:    githubv4NewStringSlice(githubv4StringSliceEmpty(data.RequiredStatusCheckContexts)),
		RequiresApprovingReviews:       githubv4.NewBoolean(githubv4.Boolean(data.RequiresApprovingReviews)),
		RequiresCodeOwnerReviews:       githubv4.NewBoolean(githubv4.Boolean(data.RequiresCodeOwnerReviews)),
		RequiresCommitSignatures:       githubv4.NewBoolean(githubv4.Boolean(data.RequiresCommitSignatures)),
		RequiresConversationResolution: githubv4.NewBoolean(githubv4.Boolean(data.RequiresConversationResolution)),
		RequiresLinearHistory:          githubv4.NewBoolean(githubv4.Boolean(data.RequiresLinearHistory)),
		RequiresStatusChecks:           githubv4.NewBoolean(githubv4.Boolean(data.RequiresStatusChecks)),
		RequiresStrictStatusChecks:     githubv4.NewBoolean(githubv4.Boolean(data.RequiresStrictStatusChecks)),
		RestrictsPushes:                githubv4.NewBoolean(githubv4.Boolean(data.RestrictsPushes)),
		RestrictsReviewDismissals:      githubv4.NewBoolean(githubv4.Boolean(data.RestrictsReviewDismissals)),
		ReviewDismissalActorIDs:        githubv4NewIDSlice(githubv4IDSlice(data.ReviewDismissalActorIDs)),
		LockBranch:                     githubv4.NewBoolean(githubv4.Boolean(data.LockBranch)),
		RequireLastPushApproval:        githubv4.NewBoolean(githubv4.Boolean(data.RequireLastPushApproval)),
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	client := meta.(*Owner).v4client
	err = client.Mutate(ctx, &mutate, input, nil)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s", mutate.UpdateBranchProtectionRule.BranchProtectionRule.ID))

	return resourceGithubBranchProtectionRead(d, meta)
}

func resourceGithubBranchProtectionDelete(d *schema.ResourceData, meta any) error {
	var mutate struct {
		DeleteBranchProtectionRule struct { // Empty struct does not work
			ClientMutationId githubv4.ID
		} `graphql:"deleteBranchProtectionRule(input: $input)"`
	}
	input := githubv4.DeleteBranchProtectionRuleInput{
		BranchProtectionRuleID: d.Id(),
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	client := meta.(*Owner).v4client
	err := client.Mutate(ctx, &mutate, input, nil)

	return err
}

func resourceGithubBranchProtectionImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	repoName, pattern, err := parseTwoPartID(d.Id(), "repository", "pattern")
	if err != nil {
		return nil, err
	}

	repoID, err := getRepositoryID(repoName, meta)
	if err != nil {
		return nil, err
	}
	if err = d.Set("repository_id", repoID); err != nil {
		return nil, err
	}

	id, err := getBranchProtectionID(repoID, pattern, meta)
	if err != nil {
		return nil, err
	}
	d.SetId(fmt.Sprintf("%s", id))

	return []*schema.ResourceData{d}, resourceGithubBranchProtectionRead(d, meta)
}
