package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/shurcooL/githubv4"
)

type Actor struct {
	ID   githubv4.ID
	Name githubv4.String
}

type ActorTypes struct {
	Actor struct {
		Team Actor `graphql:"... on Team"`
		User Actor `graphql:"... on User"`
	}
}

type BranchProtectionRule struct {
	Repository struct {
		ID   githubv4.String
		Name githubv4.String
	}
	PushAllowances struct {
		Nodes []ActorTypes
	} `graphql:"pushAllowances(first: 100)"`
	ReviewDismissalAllowances struct {
		Nodes []ActorTypes
	} `graphql:"reviewDismissalAllowances(first: 100)"`
	DismissesStaleReviews        githubv4.Boolean
	ID                           githubv4.ID
	IsAdminEnforced              githubv4.Boolean
	Pattern                      githubv4.String
	RequiredApprovingReviewCount githubv4.Int
	RequiredStatusCheckContexts  []githubv4.String
	RequiresApprovingReviews     githubv4.Boolean
	RequiresCodeOwnerReviews     githubv4.Boolean
	RequiresCommitSignatures     githubv4.Boolean
	RequiresStatusChecks         githubv4.Boolean
	RequiresStrictStatusChecks   githubv4.Boolean
	RestrictsPushes              githubv4.Boolean
	RestrictsReviewDismissals    githubv4.Boolean
}

type BranchProtectionResourceData struct {
	BranchProtectionRuleID       string
	DismissesStaleReviews        bool
	IsAdminEnforced              bool
	Pattern                      string
	PushActorIDs                 []string
	RepositoryID                 string
	RequiredApprovingReviewCount int
	RequiredStatusCheckContexts  []string
	RequiresApprovingReviews     bool
	RequiresCodeOwnerReviews     bool
	RequiresCommitSignatures     bool
	RequiresStatusChecks         bool
	RequiresStrictStatusChecks   bool
	RestrictsPushes              bool
	RestrictsReviewDismissals    bool
	ReviewDismissalActorIDs      []string
}

func branchProtectionResourceData(d *schema.ResourceData, meta interface{}) (BranchProtectionResourceData, error) {
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

	if v, ok := d.GetOk(PROTECTION_IS_ADMIN_ENFORCED); ok {
		data.IsAdminEnforced = v.(bool)
	}

	if v, ok := d.GetOk(PROTECTION_REQUIRES_COMMIT_SIGNATURES); ok {
		data.RequiresCommitSignatures = v.(bool)
	}

	if v, ok := d.GetOk(PROTECTION_REQUIRES_APPROVING_REVIEWS); ok {
		vL := v.([]interface{})
		if len(vL) > 1 {
			return BranchProtectionResourceData{},
				fmt.Errorf("error multiple %s declarations", PROTECTION_REQUIRES_APPROVING_REVIEWS)
		}
		for _, v := range vL {
			if v == nil {
				break
			}

			data.RequiresApprovingReviews = true

			m := v.(map[string]interface{})
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
		}
	}

	if v, ok := d.GetOk(PROTECTION_REQUIRES_STATUS_CHECKS); ok {
		vL := v.([]interface{})
		if len(vL) > 1 {
			return BranchProtectionResourceData{},
				fmt.Errorf("error multiple %s declarations", PROTECTION_REQUIRES_STATUS_CHECKS)
		}
		for _, v := range vL {
			if v == nil {
				break
			}

			m := v.(map[string]interface{})
			if v, ok := m[PROTECTION_REQUIRES_STRICT_STATUS_CHECKS]; ok {
				data.RequiresStrictStatusChecks = v.(bool)
			}

			data.RequiredStatusCheckContexts = expandNestedSet(m, PROTECTION_REQUIRED_STATUS_CHECK_CONTEXTS)
			if len(data.RequiredStatusCheckContexts) > 0 {
				data.RequiresStatusChecks = true
			}
		}
	}

	if v, ok := d.GetOk(PROTECTION_RESTRICTS_PUSHES); ok {
		pushActorIDs := make([]string, 0)
		vL := v.(*schema.Set).List()
		for _, v := range vL {
			pushActorIDs = append(pushActorIDs, v.(string))
		}
		if len(pushActorIDs) > 0 {
			data.PushActorIDs = pushActorIDs
			data.RestrictsPushes = true
		}
	}

	return data, nil
}

func setActorIDs(actors []ActorTypes) []string {
	pushActors := make([]string, 0, len(actors))
	for _, a := range actors {
		if a.Actor.Team != (Actor{}) {
			pushActors = append(pushActors, a.Actor.Team.ID.(string))
		}
		if a.Actor.User != (Actor{}) {
			pushActors = append(pushActors, a.Actor.Team.ID.(string))
		}
	}

	return pushActors
}

func setApprovingReviews(protection BranchProtectionRule) interface{} {
	if !protection.RequiresApprovingReviews {
		return nil
	}

	dismissalAllowances := protection.ReviewDismissalAllowances.Nodes
	dismissalActors := setActorIDs(dismissalAllowances)
	approvalReviews := []interface{}{
		map[string]interface{}{
			PROTECTION_REQUIRED_APPROVING_REVIEW_COUNT: protection.RequiredApprovingReviewCount,
			PROTECTION_REQUIRES_CODE_OWNER_REVIEWS:     protection.RequiresCodeOwnerReviews,
			PROTECTION_DISMISSES_STALE_REVIEWS:         protection.DismissesStaleReviews,
			PROTECTION_RESTRICTS_REVIEW_DISMISSALS:     dismissalActors,
		},
	}

	return approvalReviews
}

func setStatusChecks(protection BranchProtectionRule) interface{} {
	if !protection.RequiresStatusChecks {
		return nil
	}

	statusChecks := []interface{}{
		map[string]interface{}{
			PROTECTION_REQUIRES_STRICT_STATUS_CHECKS:  protection.RequiresStrictStatusChecks,
			PROTECTION_REQUIRED_STATUS_CHECK_CONTEXTS: protection.RequiredStatusCheckContexts,
		},
	}

	return statusChecks
}

func setPushes(protection BranchProtectionRule) []string {
	if !protection.RestrictsPushes {
		return nil
	}
	pushAllowances := protection.PushAllowances.Nodes
	pushActors := setActorIDs(pushAllowances)

	return pushActors
}

func getBranchProtectionID(name string, pattern string, meta interface{}) (githubv4.ID, error) {
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
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	variables := map[string]interface{}{
		"owner":  githubv4.String(meta.(*Owner).name),
		"name":   githubv4.String(name),
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

	var id string
	for i := range allRules {
		if allRules[i].Pattern == pattern {
			id = allRules[i].ID
			break
		}
	}

	return id, nil
}
