package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-github/v53/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func buildRulesetRequest(d *schema.ResourceData, sourceType *string) (*github.Ruleset, error) {
	target := d.Get("target").(string)
	req := &github.Ruleset{
		// ID:           0,
		Name:        d.Get("name").(string),
		Target:      &target,
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

	conditions, err := expandConditions(d)
	if err != nil {
		return nil, err
	}
	rulesetConditions := github.RulesetConditions{
		RefName:        conditions,
		// RepositoryName: &github.RulesetRepositoryConditionParameters{}, // TODO: Implement for org stuff
	}
	req.Conditions = &rulesetConditions

	rules, err := expandRules(d)
	if err != nil {
		return nil, err
	}
	req.Rules = rules

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
			// actorBypasMode := m["bypass_mode"].(string) // Pending a bump of the underlying sdk (needs https://github.com/google/go-github/blob/c030d43bc8e3003715a3de91972b1a594039d262/github/repos_rules.go#L15-L21)
			bypassActor := &github.BypassActor{
				ActorID:   actorID,
				ActorType: actorType,
				// BypassMode: actorBypasMode,
			}
			bypassActors = append(bypassActors, bypassActor)
		}

		return bypassActors, nil
	}
	return nil, nil	
}

func expandConditions(d *schema.ResourceData) (*github.RulesetRefConditionParameters, error) {
	if v, ok := d.GetOk("conditions"); ok {
		vL := v.([]interface{})
		if len(vL) > 1 {
			return nil, errors.New("cannot specify conditions more than one time")
		}
		rulesetConditions := new(github.RulesetRefConditionParameters)

		for _, v := range vL {
			// List can only have one item, safe to early return here
			if v == nil {
				return nil, nil
			}
			m := v.(map[string]interface{})
			rulesetConditions.Include = expandNestedSet(m, "include")
			rulesetConditions.Exclude = expandNestedSet(m, "exclude")
		}

		return rulesetConditions, nil
	}

	return nil, nil
}

func expandRules(d *schema.ResourceData) ([]*github.RepositoryRule, error) {
	rulesetRules := make([]*github.RepositoryRule, 0)

	rules_toggleable := []string{
		"rule_creation",               
		"rule_update",                 
		"rule_deletion",               
		"rule_required_linear_history",
		"rule_required_signatures",    
		"rule_non_fast_forward",       
	}
	for _, ruleName := range rules_toggleable {
		rulesOn := d.Get(ruleName).(bool)
		if rulesOn {
			toggleableRule := &github.RepositoryRule{
				Type:       ruleName,
			}
			rulesetRules = append(rulesetRules, toggleableRule)
		}
	}

	if v, ok := d.GetOk("rule_required_deployments"); ok {
		vL := v.([]interface{})
		if len(vL) > 0 {
			requiredDeploymentEnvs := make([]string, len(vL))
			for i, env := range vL {
				requiredDeploymentEnvs[i] = env.(string)
			}

			mRequiredDeploymentParams := github.RequiredDeploymentEnvironmentsRuleParameters{
				RequiredDeploymentEnvironments: requiredDeploymentEnvs,
			}
			bytes, _ := json.Marshal(mRequiredDeploymentParams)
			rawParams := json.RawMessage(bytes)
			requiredDeploymentRule := &github.RepositoryRule{
				Type:       "required_deployments", // Drop the "rule_", that is only to make the provider implementation easier
				Parameters: &rawParams,
			}
			rulesetRules = append(rulesetRules, requiredDeploymentRule)		
		}
	}

	if v, ok := d.GetOk("rule_pull_request"); ok {
		vL := v.([]interface{})
		if len(vL) > 1 {
			return nil, errors.New("cannot specify rule_pull_request more than one time")
		}
		for _, v := range vL {
			// List can only have one item, safe to early return here
			if v == nil {
				return nil, nil
			}
			m := v.(map[string]interface{})
			
			pullRequestRuleParams := github.PullRequestRuleParameters{
				DismissStaleReviewsOnPush:      m["dismiss_stale_reviews_on_push"].(bool),
				RequireCodeOwnerReview:         m["require_code_owner_review"].(bool),
				RequireLastPushApproval:        m["require_last_push_approval"].(bool),
				RequiredApprovingReviewCount:   m["required_approving_review_count"].(int),
				RequiredReviewThreadResolution: m["required_review_thread_resolution"].(bool),
			}
			bytes, _ := json.Marshal(pullRequestRuleParams)
			rawParams := json.RawMessage(bytes)
			requiredDeploymentRule := &github.RepositoryRule{
				Type:       "pull_request",
				Parameters: &rawParams,
			}
			rulesetRules = append(rulesetRules, requiredDeploymentRule)
		}

		return rulesetRules, nil
	}

	if v, ok := d.GetOk("rule_required_status_checks"); ok {
		vL := v.([]interface{})
		if len(vL) > 1 {
			return nil, errors.New("cannot specify rule_required_status_checks more than one time")
		}
		for _, v := range vL {
			// List can only have one item, safe to early return here
			if v == nil {
				return nil, nil
			}
			m := v.(map[string]interface{})

			requiredStatusChecks := m["strict_required_status_checks_policy"].([]string)
			requiredStatusChecksList := make([]github.RuleRequiredStatusChecks, 0)
			for _, statusCheck := range requiredStatusChecks {

				// Expect a string of "context:integration_id", allowing for the absence of "integration_id"
				parts := strings.SplitN(statusCheck, ":", 2)
				var cContext, cIntegrationId string
				switch len(parts) {
					case 1:
						cContext, cIntegrationId = parts[0], ""
					case 2:
						cContext, cIntegrationId = parts[0], parts[1]
					default:
						// TODO: What is the prefered way of throwing errors? fmt.Errorf() or errors.New()?
						return nil, fmt.Errorf("Could not parse check '%s'. Expected `context:integration_id` or `context`", statusCheck)
				}

				var rrscCheck *github.RuleRequiredStatusChecks
				if cIntegrationId != "" {
					// If we have a valid app_id, include it in the RSC
					rrscIntegrationId, err := strconv.Atoi(cIntegrationId)
					if err != nil {
						return nil, fmt.Errorf("Could not parse %v as valid integration_id", cIntegrationId)
					}
					rrscIntegrationId64 := int64(rrscIntegrationId)
					rrscCheck = &github.RuleRequiredStatusChecks{Context: cContext, IntegrationID: &rrscIntegrationId64}
				} else {
					// Else simply provide the context
					rrscCheck = &github.RuleRequiredStatusChecks{Context: cContext}
				}

				requiredStatusChecksList = append(requiredStatusChecksList, *rrscCheck)

			}
			
			requiredStatusChecksRuleParams := github.RequiredStatusChecksRuleParameters{
				RequiredStatusChecks:             requiredStatusChecksList,
				StrictRequiredStatusChecksPolicy: m["strict_required_status_checks_policy"].(bool),
			}
			bytes, _ := json.Marshal(requiredStatusChecksRuleParams)
			rawParams := json.RawMessage(bytes)
			requiredDeploymentRule := &github.RepositoryRule{
				Type:       "required_status_checks",
				Parameters: &rawParams,
			}
			rulesetRules = append(rulesetRules, requiredDeploymentRule)
		}

		return rulesetRules, nil
	}


	return nil, nil
}
