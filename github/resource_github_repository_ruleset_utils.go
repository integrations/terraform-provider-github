package github

import (
	"errors"

	"github.com/google/go-github/v53/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func buildRulesetRequest(d *schema.ResourceData, sourceType *string) (*github.Ruleset, error) {
	req := &github.Ruleset{
		// ID:           0,
		Name:        d.Get("name").(string),
		Target:      d.Get("target").(*string),
		SourceType:  sourceType,
		Enforcement: d.Get("enforcement").(string),
		// BypassMode:   new(string), TODO: what is this?
		// Links:        &github.RulesetLinks{}, TODO: Ignore this?
	}

	bypassActors, err := expandBypassActors(d)
	if err != nil {
		return nil, err
	}
	req.BypassActors = bypassActors

	// TOOD: expandBypassActors()
	// TOOD: expandConditions()
	// TODO. expandRules()

	// rsc, err := expandRequiredStatusChecks(d)
	// if err != nil {
	// 	return nil, err
	// }
	// req.RequiredStatusChecks = rsc

	// rprr, err := expandRequiredPullRequestReviews(d)
	// if err != nil {
	// 	return nil, err
	// }
	// req.RequiredPullRequestReviews = rprr

	// res, err := expandRestrictions(d)
	// if err != nil {
	// 	return nil, err
	// }
	// req.Restrictions = res

	return req, nil
}

func expandBypassActors(d *schema.ResourceData) ([]*github.BypassActor, error) {
	bypassActors := make([]*github.BypassActor, 0)

	if v, ok := d.GetOk("bypass_actors"); ok {
		vL := v.(*schema.Set).List()
		for _, v := range vL {
			m := v.(map[string]interface{})
			actorID := m["actor_id"].(*int64)
			actorType := m["actor_type"].(*string)
			// actorBypasMode := m["bypass_mode"].(string) // Pending a bump of the underlying sdk
			bypassActor := &github.BypassActor{
				ActorID:   actorID,
				ActorType: actorType,
			}
			groups = append(groups, bypassActor)
		}
	}
	return &github.IDPGroupList{Groups: groups}



	if v, ok := d.GetOk("bypass_actors"); ok {
		bypassActorsSchema := v.([]interface{})
		if len(bypassActorsSchema) > 1 {
			return nil, errors.New("cannot specify bypass_actors more than one time")
		}
		bypassActors := new(github.BypassActor)

		for _, v := range bypassActorsSchema {
			// List can only have one item, safe to early return here
			// TODO: This might not be empty 
			if v == nil {
				return nil, nil
			}
			m := v.(map[string]interface{})

			users := expandNestedSet(m, "dismissal_users")
			if len(users) > 0 {
				drr.Users = &users
			}
			teams := expandNestedSet(m, "dismissal_teams")
			if len(teams) > 0 {
				drr.Teams = &teams
			}

			bpra, err := expandBypassPullRequestAllowances(m)
			if err != nil {
				return nil, err
			}

			rprr.DismissalRestrictionsRequest = drr
			rprr.DismissStaleReviews = m["dismiss_stale_reviews"].(bool)
			rprr.RequireCodeOwnerReviews = m["require_code_owner_reviews"].(bool)
			rprr.RequiredApprovingReviewCount = m["required_approving_review_count"].(int)
			rprr.BypassPullRequestAllowancesRequest = bpra
		}

		return rprr, nil
	}

	return nil, nil
}
