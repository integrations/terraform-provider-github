package github

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

type Actor struct {
	ID   githubv4.ID
	Name githubv4.String
	Slug githubv4.String
}

type ActorUser struct {
	ID    githubv4.ID
	Name  githubv4.String
	Login githubv4.String
}

type DismissalActorTypes struct {
	Actor struct {
		App  Actor     `graphql:"... on App"`
		Team Actor     `graphql:"... on Team"`
		User ActorUser `graphql:"... on User"`
	}
}

type BypassForcePushActorTypes struct {
	Actor struct {
		App  Actor     `graphql:"... on App"`
		Team Actor     `graphql:"... on Team"`
		User ActorUser `graphql:"... on User"`
	}
}

type BypassPullRequestActorTypes struct {
	Actor struct {
		App  Actor     `graphql:"... on App"`
		Team Actor     `graphql:"... on Team"`
		User ActorUser `graphql:"... on User"`
	}
}

type PushActorTypes struct {
	Actor struct {
		App  Actor     `graphql:"... on App"`
		Team Actor     `graphql:"... on Team"`
		User ActorUser `graphql:"... on User"`
	}
}

type BranchProtectionRule struct {
	Repository struct {
		ID   githubv4.String
		Name githubv4.String
	}
	PushAllowances struct {
		Nodes []PushActorTypes
	} `graphql:"pushAllowances(first: 100)"`
	ReviewDismissalAllowances struct {
		Nodes []DismissalActorTypes
	} `graphql:"reviewDismissalAllowances(first: 100)"`
	BypassForcePushAllowances struct {
		Nodes []BypassForcePushActorTypes
	} `graphql:"bypassForcePushAllowances(first: 100)"`
	BypassPullRequestAllowances struct {
		Nodes []BypassPullRequestActorTypes
	} `graphql:"bypassPullRequestAllowances(first: 100)"`
	AllowsDeletions                githubv4.Boolean
	AllowsForcePushes              githubv4.Boolean
	BlocksCreations                githubv4.Boolean
	DismissesStaleReviews          githubv4.Boolean
	ID                             githubv4.ID
	IsAdminEnforced                githubv4.Boolean
	Pattern                        githubv4.String
	RequiredApprovingReviewCount   githubv4.Int
	RequiredStatusCheckContexts    []githubv4.String
	RequiresApprovingReviews       githubv4.Boolean
	RequiresCodeOwnerReviews       githubv4.Boolean
	RequiresCommitSignatures       githubv4.Boolean
	RequiresLinearHistory          githubv4.Boolean
	RequiresConversationResolution githubv4.Boolean
	RequiresStatusChecks           githubv4.Boolean
	RequiresStrictStatusChecks     githubv4.Boolean
	RestrictsPushes                githubv4.Boolean
	RestrictsReviewDismissals      githubv4.Boolean
	RequireLastPushApproval        githubv4.Boolean
	LockBranch                     githubv4.Boolean
}

type BranchProtectionResourceData struct {
	AllowsDeletions                bool
	AllowsForcePushes              bool
	BlocksCreations                bool
	BranchProtectionRuleID         string
	BypassForcePushActorIDs        []string
	BypassPullRequestActorIDs      []string
	DismissesStaleReviews          bool
	IsAdminEnforced                bool
	Pattern                        string
	PushActorIDs                   []string
	RepositoryID                   string
	RequiredApprovingReviewCount   int
	RequiredStatusCheckContexts    []string
	RequiresApprovingReviews       bool
	RequiresCodeOwnerReviews       bool
	RequiresCommitSignatures       bool
	RequiresLinearHistory          bool
	RequiresConversationResolution bool
	RequiresStatusChecks           bool
	RequiresStrictStatusChecks     bool
	RestrictsPushes                bool
	RestrictsReviewDismissals      bool
	ReviewDismissalActorIDs        []string
	RequireLastPushApproval        bool
	LockBranch                     bool
}

func branchProtectionResourceData(d *schema.ResourceData, meta any) (BranchProtectionResourceData, error) {
	data := BranchProtectionResourceData{}

	if v, ok := d.GetOk(REPOSITORY_ID); ok {
		repoID, err := getRepositoryID(v.(string), meta)
		if err != nil {
			return data, err
		}
		data.RepositoryID = repoID.(string)
	}

	if v, ok := d.GetOk(PROTECTION_PATTERN); ok {
		data.Pattern = v.(string)
	}

	if v, ok := d.GetOk(PROTECTION_ALLOWS_DELETIONS); ok {
		data.AllowsDeletions = v.(bool)
	}

	if v, ok := d.GetOk(PROTECTION_ALLOWS_FORCE_PUSHES); ok {
		data.AllowsForcePushes = v.(bool)
	}

	if v, ok := d.GetOk(PROTECTION_IS_ADMIN_ENFORCED); ok {
		data.IsAdminEnforced = v.(bool)
	}

	if v, ok := d.GetOk(PROTECTION_REQUIRES_COMMIT_SIGNATURES); ok {
		data.RequiresCommitSignatures = v.(bool)
	}

	if v, ok := d.GetOk(PROTECTION_REQUIRES_LINEAR_HISTORY); ok {
		data.RequiresLinearHistory = v.(bool)
	}

	if v, ok := d.GetOk(PROTECTION_REQUIRES_CONVERSATION_RESOLUTION); ok {
		data.RequiresConversationResolution = v.(bool)
	}

	if v, ok := d.GetOk(PROTECTION_REQUIRES_APPROVING_REVIEWS); ok {
		vL := v.([]any)
		if len(vL) > 1 {
			return BranchProtectionResourceData{},
				fmt.Errorf("error multiple %s declarations", PROTECTION_REQUIRES_APPROVING_REVIEWS)
		}
		for _, v := range vL {
			if v == nil {
				break
			}

			data.RequiresApprovingReviews = true

			m := v.(map[string]any)
			if v, ok := m[PROTECTION_REQUIRED_APPROVING_REVIEW_COUNT]; ok {
				data.RequiredApprovingReviewCount = v.(int)
			}
			if v, ok := m[PROTECTION_DISMISSES_STALE_REVIEWS]; ok {
				data.DismissesStaleReviews = v.(bool)
			}
			if v, ok := m[PROTECTION_REQUIRES_CODE_OWNER_REVIEWS]; ok {
				data.RequiresCodeOwnerReviews = v.(bool)
			}
			if v, ok := m[PROTECTION_RESTRICTS_REVIEW_DISMISSALS]; ok {
				data.RestrictsReviewDismissals = v.(bool)
			}
			if v, ok := m[PROTECTION_REVIEW_DISMISSAL_ALLOWANCES]; ok {
				reviewDismissalActorIDs := make([]string, 0)
				vL := v.(*schema.Set).List()
				for _, v := range vL {
					reviewDismissalActorIDs = append(reviewDismissalActorIDs, v.(string))
				}
				if len(reviewDismissalActorIDs) > 0 {
					data.ReviewDismissalActorIDs = reviewDismissalActorIDs
					data.RestrictsReviewDismissals = true
				}
			}
			if v, ok := m[PROTECTION_PULL_REQUESTS_BYPASSERS]; ok {
				bypassPullRequestActorIDs := make([]string, 0)
				vL := v.(*schema.Set).List()
				for _, v := range vL {
					bypassPullRequestActorIDs = append(bypassPullRequestActorIDs, v.(string))
				}
				if len(bypassPullRequestActorIDs) > 0 {
					data.BypassPullRequestActorIDs = bypassPullRequestActorIDs
				}
			}
			if v, ok := m[PROTECTION_REQUIRE_LAST_PUSH_APPROVAL]; ok {
				data.RequireLastPushApproval = v.(bool)
			}
		}
	}

	if v, ok := d.GetOk(PROTECTION_REQUIRES_STATUS_CHECKS); ok {
		data.RequiresStatusChecks = true
		vL := v.([]any)
		if len(vL) > 1 {
			return BranchProtectionResourceData{},
				fmt.Errorf("error multiple %s declarations", PROTECTION_REQUIRES_STATUS_CHECKS)
		}
		for _, v := range vL {
			if v == nil {
				break
			}

			m := v.(map[string]any)
			if v, ok := m[PROTECTION_REQUIRES_STRICT_STATUS_CHECKS]; ok {
				data.RequiresStrictStatusChecks = v.(bool)
			}

			data.RequiredStatusCheckContexts = expandNestedSet(m, PROTECTION_REQUIRED_STATUS_CHECK_CONTEXTS)
		}
	}

	if v, ok := d.GetOk(PROTECTION_RESTRICTS_PUSHES); ok {
		vL := v.([]any)
		if len(vL) > 1 {
			return BranchProtectionResourceData{},
				fmt.Errorf("error multiple %s declarations", PROTECTION_RESTRICTS_PUSHES)
		}
		for _, v := range vL {
			if v == nil {
				break
			}

			data.RestrictsPushes = true

			m := v.(map[string]any)
			if v, ok := m[PROTECTION_BLOCKS_CREATIONS]; ok {
				data.BlocksCreations = v.(bool)
			}
			if v, ok := m[PROTECTION_PUSH_ALLOWANCES]; ok {
				pushActorIDs := make([]string, 0)
				vL := v.(*schema.Set).List()
				for _, v := range vL {
					pushActorIDs = append(pushActorIDs, v.(string))
				}
				if len(pushActorIDs) > 0 {
					data.PushActorIDs = pushActorIDs
				}
			}
		}
	}

	if v, ok := d.GetOk(PROTECTION_FORCE_PUSHES_BYPASSERS); ok {
		bypassForcePushActorIDs := make([]string, 0)
		vL := v.(*schema.Set).List()
		for _, v := range vL {
			bypassForcePushActorIDs = append(bypassForcePushActorIDs, v.(string))
		}
		if len(bypassForcePushActorIDs) > 0 {
			data.BypassForcePushActorIDs = bypassForcePushActorIDs
			data.AllowsForcePushes = false
		}
	}

	if v, ok := d.GetOk(PROTECTION_LOCK_BRANCH); ok {
		data.LockBranch = v.(bool)
	}

	return data, nil
}

func branchProtectionResourceDataActors(d *schema.ResourceData, meta any) (BranchProtectionResourceData, error) {
	data := BranchProtectionResourceData{}
	if v, ok := d.GetOk(PROTECTION_REQUIRES_APPROVING_REVIEWS); ok {
		vL := v.([]any)
		if len(vL) > 1 {
			return BranchProtectionResourceData{},
				fmt.Errorf("error multiple %s declarations", PROTECTION_REQUIRES_APPROVING_REVIEWS)
		}
		for _, v := range vL {
			if v == nil {
				break
			}

			m := v.(map[string]any)
			if v, ok := m[PROTECTION_REVIEW_DISMISSAL_ALLOWANCES]; ok {
				reviewDismissalActorIDs := make([]string, 0)
				vL := v.(*schema.Set).List()
				for _, v := range vL {
					reviewDismissalActorIDs = append(reviewDismissalActorIDs, v.(string))
				}
				if len(reviewDismissalActorIDs) > 0 {
					data.ReviewDismissalActorIDs = reviewDismissalActorIDs
					data.RestrictsReviewDismissals = true
				}
			}
			if v, ok := m[PROTECTION_PULL_REQUESTS_BYPASSERS]; ok {
				bypassPullRequestActorIDs := make([]string, 0)
				vL := v.(*schema.Set).List()
				for _, v := range vL {
					bypassPullRequestActorIDs = append(bypassPullRequestActorIDs, v.(string))
				}
				if len(bypassPullRequestActorIDs) > 0 {
					data.BypassPullRequestActorIDs = bypassPullRequestActorIDs
				}
			}
		}
	}

	if v, ok := d.GetOk(PROTECTION_RESTRICTS_PUSHES); ok {
		vL := v.([]any)
		if len(vL) > 1 {
			return BranchProtectionResourceData{},
				fmt.Errorf("error multiple %s declarations", PROTECTION_RESTRICTS_PUSHES)
		}
		for _, v := range vL {
			if v == nil {
				break
			}

			data.RestrictsPushes = true

			m := v.(map[string]any)
			if v, ok := m[PROTECTION_BLOCKS_CREATIONS]; ok {
				data.BlocksCreations = v.(bool)
			}
			if v, ok := m[PROTECTION_PUSH_ALLOWANCES]; ok {
				pushActorIDs := make([]string, 0)
				vL := v.(*schema.Set).List()
				for _, v := range vL {
					pushActorIDs = append(pushActorIDs, v.(string))
				}
				if len(pushActorIDs) > 0 {
					data.PushActorIDs = pushActorIDs
				}
			}
		}
	}

	if v, ok := d.GetOk(PROTECTION_FORCE_PUSHES_BYPASSERS); ok {
		bypassForcePushActorIDs := make([]string, 0)
		vL := v.(*schema.Set).List()
		for _, v := range vL {
			bypassForcePushActorIDs = append(bypassForcePushActorIDs, v.(string))
		}
		if len(bypassForcePushActorIDs) > 0 {
			data.BypassForcePushActorIDs = bypassForcePushActorIDs
			data.AllowsForcePushes = false
		}
	}

	return data, nil
}

func setDismissalActorIDs(actors []DismissalActorTypes, data BranchProtectionResourceData, meta any) []string {
	dismissalActors := make([]string, 0, len(actors))
	orgName := meta.(*Owner).name

	for _, a := range actors {
		IsID := false
		for _, v := range data.ReviewDismissalActorIDs {
			if (a.Actor.Team.ID != nil && a.Actor.Team.ID.(string) == v) || (a.Actor.User.ID != nil && a.Actor.User.ID.(string) == v) || (a.Actor.App.ID != nil && a.Actor.App.ID.(string) == v) {
				dismissalActors = append(dismissalActors, v)
				IsID = true
				break
			}
		}
		if !IsID {
			if a.Actor.Team.Slug != "" {
				dismissalActors = append(dismissalActors, orgName+"/"+string(a.Actor.Team.Slug))
				continue
			}
			if a.Actor.User.Login != "" {
				dismissalActors = append(dismissalActors, "/"+string(a.Actor.User.Login))
				continue
			}
			if a.Actor.App != (Actor{}) {
				dismissalActors = append(dismissalActors, a.Actor.App.ID.(string))
			}
		}
	}
	return dismissalActors
}

func setBypassForcePushActorIDs(actors []BypassForcePushActorTypes, data BranchProtectionResourceData, meta any) []string {
	bypassActors := make([]string, 0, len(actors))

	orgName := meta.(*Owner).name

	for _, a := range actors {
		IsID := false
		for _, v := range data.BypassForcePushActorIDs {
			if (a.Actor.Team.ID != nil && a.Actor.Team.ID.(string) == v) || (a.Actor.User.ID != nil && a.Actor.User.ID.(string) == v) || (a.Actor.App.ID != nil && a.Actor.App.ID.(string) == v) {
				bypassActors = append(bypassActors, v)
				IsID = true
				break
			}
		}
		if !IsID {
			if a.Actor.Team.Slug != "" {
				bypassActors = append(bypassActors, orgName+"/"+string(a.Actor.Team.Slug))
				continue
			}
			if a.Actor.User.Login != "" {
				bypassActors = append(bypassActors, "/"+string(a.Actor.User.Login))
				continue
			}
			if a.Actor.App != (Actor{}) {
				bypassActors = append(bypassActors, a.Actor.App.ID.(string))
			}
		}
	}
	return bypassActors
}

func setBypassPullRequestActorIDs(actors []BypassPullRequestActorTypes, data BranchProtectionResourceData, meta any) []string {
	bypassActors := make([]string, 0, len(actors))

	orgName := meta.(*Owner).name

	for _, a := range actors {
		IsID := false
		for _, v := range data.BypassPullRequestActorIDs {
			if (a.Actor.Team.ID != nil && a.Actor.Team.ID.(string) == v) || (a.Actor.User.ID != nil && a.Actor.User.ID.(string) == v) || (a.Actor.App.ID != nil && a.Actor.App.ID.(string) == v) {
				bypassActors = append(bypassActors, v)
				IsID = true
				break
			}
		}
		if !IsID {
			if a.Actor.Team.Slug != "" {
				bypassActors = append(bypassActors, orgName+"/"+string(a.Actor.Team.Slug))
				continue
			}
			if a.Actor.User.Login != "" {
				bypassActors = append(bypassActors, "/"+string(a.Actor.User.Login))
				continue
			}
			if a.Actor.App != (Actor{}) {
				bypassActors = append(bypassActors, a.Actor.App.ID.(string))
			}
		}
	}
	return bypassActors
}

func setPushActorIDs(actors []PushActorTypes, data BranchProtectionResourceData, meta any) []string {
	pushActors := make([]string, 0, len(actors))

	orgName := meta.(*Owner).name

	for _, a := range actors {
		IsID := false
		for _, v := range data.PushActorIDs {
			if (a.Actor.Team.ID != nil && a.Actor.Team.ID.(string) == v) || (a.Actor.User.ID != nil && a.Actor.User.ID.(string) == v) || (a.Actor.App.ID != nil && a.Actor.App.ID.(string) == v) {
				pushActors = append(pushActors, v)
				IsID = true
				break
			}
		}
		if !IsID {
			if a.Actor.Team.Slug != "" {
				pushActors = append(pushActors, orgName+"/"+string(a.Actor.Team.Slug))
				continue
			}
			if a.Actor.User.Login != "" {
				pushActors = append(pushActors, "/"+string(a.Actor.User.Login))
				continue
			}
			if a.Actor.App != (Actor{}) {
				pushActors = append(pushActors, a.Actor.App.ID.(string))
			}
		}
	}
	return pushActors
}

func setApprovingReviews(protection BranchProtectionRule, data BranchProtectionResourceData, meta any) any {
	if !protection.RequiresApprovingReviews {
		return nil
	}

	dismissalAllowances := protection.ReviewDismissalAllowances.Nodes
	dismissalActors := setDismissalActorIDs(dismissalAllowances, data, meta)

	bypassPullRequestAllowances := protection.BypassPullRequestAllowances.Nodes
	bypassPullRequestActors := setBypassPullRequestActorIDs(bypassPullRequestAllowances, data, meta)

	approvalReviews := []any{
		map[string]any{
			PROTECTION_REQUIRED_APPROVING_REVIEW_COUNT: protection.RequiredApprovingReviewCount,
			PROTECTION_REQUIRES_CODE_OWNER_REVIEWS:     protection.RequiresCodeOwnerReviews,
			PROTECTION_DISMISSES_STALE_REVIEWS:         protection.DismissesStaleReviews,
			PROTECTION_RESTRICTS_REVIEW_DISMISSALS:     protection.RestrictsReviewDismissals,
			PROTECTION_REVIEW_DISMISSAL_ALLOWANCES:     dismissalActors,
			PROTECTION_PULL_REQUESTS_BYPASSERS:         bypassPullRequestActors,
			PROTECTION_REQUIRE_LAST_PUSH_APPROVAL:      protection.RequireLastPushApproval,
		},
	}

	return approvalReviews
}

func setStatusChecks(protection BranchProtectionRule) any {
	if !protection.RequiresStatusChecks {
		return nil
	}

	statusChecks := []any{
		map[string]any{
			PROTECTION_REQUIRES_STRICT_STATUS_CHECKS:  protection.RequiresStrictStatusChecks,
			PROTECTION_REQUIRED_STATUS_CHECK_CONTEXTS: protection.RequiredStatusCheckContexts,
		},
	}

	return statusChecks
}

func setPushes(protection BranchProtectionRule, data BranchProtectionResourceData, meta any) any {
	if !protection.RestrictsPushes {
		return nil
	}

	pushAllowances := protection.PushAllowances.Nodes
	pushActors := setPushActorIDs(pushAllowances, data, meta)

	restrictsPushes := []any{
		map[string]any{
			PROTECTION_BLOCKS_CREATIONS: protection.BlocksCreations,
			PROTECTION_PUSH_ALLOWANCES:  pushActors,
		},
	}

	return restrictsPushes
}

func setForcePushBypassers(protection BranchProtectionRule, data BranchProtectionResourceData, meta any) []string {
	if protection.AllowsForcePushes {
		return nil
	}
	bypassForcePushAllowances := protection.BypassForcePushAllowances.Nodes
	bypassForcePushActors := setBypassForcePushActorIDs(bypassForcePushAllowances, data, meta)

	return bypassForcePushActors
}

func getBranchProtectionID(repoID githubv4.ID, pattern string, meta any) (githubv4.ID, error) {
	var query struct {
		Node struct {
			Repository struct {
				BranchProtectionRules struct {
					Nodes []struct {
						ID      string
						Pattern string
					}
					PageInfo PageInfo
				} `graphql:"branchProtectionRules(first: $first, after: $cursor)"`
				ID string
			} `graphql:"... on Repository"`
		} `graphql:"node(id: $id)"`
	}
	variables := map[string]any{
		"id":     repoID,
		"first":  githubv4.Int(100),
		"cursor": (*githubv4.String)(nil),
	}

	ctx := context.Background()
	client := meta.(*Owner).v4client

	var allRules []struct {
		ID      string
		Pattern string
	}
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}

		allRules = append(allRules, query.Node.Repository.BranchProtectionRules.Nodes...)

		if !query.Node.Repository.BranchProtectionRules.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Node.Repository.BranchProtectionRules.PageInfo.EndCursor)
	}

	for i := range allRules {
		if allRules[i].Pattern == pattern {
			return allRules[i].ID, nil
		}
	}

	return nil, fmt.Errorf("could not find a branch protection rule with the pattern '%s'", pattern)
}

func getActorIds(data []string, meta any) ([]string, error) {
	var actors []string
	for _, v := range data {
		id, err := getNodeIDv4(v, meta)
		if err != nil {
			return []string{}, err
		}
		log.Printf("[DEBUG] Retrieved node ID for user/team : %s - node ID : %s", v, id)
		actors = append(actors, id)
	}

	return actors, nil
}

// Given a string that is either a username or team slug, return the
// node id of the user or team it is referring to. Team slugs must be provided
// with the organization name as prefix (Ex.: exampleorg/exampleteam). Usernames
// must be provided with the "/" prefix otherwise getNodeIDv4 assumes that
// the provided string is a node ID.
func getNodeIDv4(userOrSlug string, meta any) (string, error) {
	orgName := meta.(*Owner).name
	ctx := context.Background()
	client := meta.(*Owner).v4client

	if strings.HasPrefix(userOrSlug, orgName+"/") {
		var queryTeam struct {
			Organization struct {
				Team struct {
					ID string
				} `graphql:"team(slug: $slug)"`
			} `graphql:"organization(login: $organization)"`
		}
		teamName := strings.TrimPrefix(userOrSlug, orgName+"/")
		variablesTeam := map[string]any{
			"slug":         githubv4.String(teamName),
			"organization": githubv4.String(orgName),
		}

		err := client.Query(ctx, &queryTeam, variablesTeam)
		if err != nil {
			return "", err
		}
		log.Printf("[DEBUG] Retrieved node ID for team %s. ID is %s", userOrSlug, queryTeam.Organization.Team.ID)
		return queryTeam.Organization.Team.ID, nil
	} else if strings.HasPrefix(userOrSlug, "/") {
		// The "/" prefix indicates a username
		var queryUser struct {
			User struct {
				ID string
			} `graphql:"user(login: $user)"`
		}
		userName := strings.TrimPrefix(userOrSlug, "/")
		variablesUser := map[string]any{
			"user": githubv4.String(userName),
		}

		err := client.Query(ctx, &queryUser, variablesUser)
		if err != nil {
			return "", err
		}
		log.Printf("[DEBUG] Retrieved node ID for user %s. ID is %s", userOrSlug, queryUser.User.ID)
		return queryUser.User.ID, nil
	} else {
		// If userOrSlug does not contain the team or username prefix, assume it is a node ID
		return userOrSlug, nil
	}
}
