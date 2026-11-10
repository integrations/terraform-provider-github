package github

import (
	"github.com/google/go-github/v93/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// actionsPolicyAllWorkflows is the workflow_path pattern that matches all workflows.
const actionsPolicyAllWorkflows = "~ALL"

var actionsPolicyEvents = []string{
	"branch_protection_rule", "check_run", "check_suite", "create", "delete", "deployment",
	"deployment_status", "discussion", "discussion_comment", "fork", "gollum", "image_version",
	"issue_comment", "issues", "label", "merge_group", "milestone", "page_build", "project",
	"project_card", "project_column", "public", "pull_request", "pull_request_review",
	"pull_request_review_comment", "pull_request_target", "push", "registry_package", "release",
	"repository_dispatch", "schedule", "status", "watch", "workflow_call", "workflow_dispatch",
	"workflow_run",
}

var actionsPolicyActorTypes = []string{
	string(github.ActionsPolicyActorTypeUser),
	string(github.ActionsPolicyActorTypeBot),
	string(github.ActionsPolicyActorTypeTeam),
	string(github.ActionsPolicyActorTypeBusinessTeam),
	string(github.ActionsPolicyActorTypeEnterpriseTeam),
	string(github.ActionsPolicyActorTypeIntegrationInstallation),
	string(github.ActionsPolicyActorTypeApp),
	string(github.ActionsPolicyActorTypeRepositoryRole),
}

// actionsPolicyCommonSchema returns the schema fields shared by all Actions policy resources.
func actionsPolicyCommonSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the policy.",
		},
		"enforcement": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The enforcement level of the policy. Can be one of `disabled`, `active`, or `evaluate`. `evaluate` is only available with GitHub Enterprise.",
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
				string(github.RulesetEnforcementDisabled),
				string(github.RulesetEnforcementActive),
				string(github.RulesetEnforcementEvaluate),
			}, false)),
		},
		"workflow_path": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Restrict the policy to matching workflow files. If omitted, the policy applies to all workflows.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"include": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Workflow file paths or glob patterns to include. An empty list includes all workflows not matched by `exclude`.",
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"exclude": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Workflow file paths or glob patterns to exclude.",
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		"allowed_events": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The events that are allowed to trigger matching workflows.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(actionsPolicyEvents, false)),
			},
		},
		"allowed_actor": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The actors that are allowed to trigger matching workflows.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeInt,
						Required:    true,
						Description: "The ID of the actor.",
					},
					"type": {
						Type:             schema.TypeString,
						Required:         true,
						Description:      "The type of the actor. Can be one of `User`, `Bot`, `Team`, `BusinessTeam`, `EnterpriseTeam`, `IntegrationInstallation`, `App`, or `RepositoryRole`.",
						ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(actionsPolicyActorTypes, false)),
					},
				},
			},
		},
		"policy_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The ID of the policy.",
		},
	}
}

// expandActionsPolicyRules builds the policy rules from the allowed_actor and allowed_events fields.
// It always returns a non-nil slice so that removing all rules on update clears them.
func expandActionsPolicyRules(d *schema.ResourceData) []*github.ActionsPolicyRule {
	rules := []*github.ActionsPolicyRule{}

	if v, ok := d.GetOk("allowed_actor"); ok {
		var actors []*github.ActionsPolicyActor
		for _, a := range v.(*schema.Set).List() {
			actor := a.(map[string]any)
			actors = append(actors, &github.ActionsPolicyActor{
				ID:   int64(actor["id"].(int)),
				Type: github.ActionsPolicyActorType(actor["type"].(string)),
			})
		}
		rules = append(rules, &github.ActionsPolicyRule{
			Type:       github.ActionsPolicyRuleTypeRestrictActionsActors,
			Parameters: &github.ActionsPolicyRuleParameters{AllowedActors: actors},
		})
	}

	if v, ok := d.GetOk("allowed_events"); ok {
		rules = append(rules, &github.ActionsPolicyRule{
			Type:       github.ActionsPolicyRuleTypeRestrictActionEvents,
			Parameters: &github.ActionsPolicyRuleParameters{AllowedEvents: expandStringList(v.(*schema.Set).List())},
		})
	}

	return rules
}

// expandActionsPolicyWorkflowPath builds the workflow_path condition. When
// updating, an omitted workflow_path is sent as ~ALL because the API preserves
// the existing workflow targeting if the condition is omitted.
func expandActionsPolicyWorkflowPath(d *schema.ResourceData, isUpdate bool) *github.ActionsPolicyWorkflowPathConditionParameters {
	v := d.Get("workflow_path").([]any)
	if len(v) == 0 || v[0] == nil {
		if isUpdate {
			return &github.ActionsPolicyWorkflowPathConditionParameters{
				Include: []string{actionsPolicyAllWorkflows},
				Exclude: []string{},
			}
		}
		return nil
	}

	wp := v[0].(map[string]any)
	return &github.ActionsPolicyWorkflowPathConditionParameters{
		Include: expandStringList(wp["include"].([]any)),
		Exclude: expandStringList(wp["exclude"].([]any)),
	}
}

// flattenActionsPolicyWorkflowPath returns nil when the condition targets all workflows.
func flattenActionsPolicyWorkflowPath(conditions *github.ActionsPolicyConditions) []any {
	if conditions == nil || conditions.WorkflowPath == nil {
		return nil
	}
	wp := conditions.WorkflowPath
	if len(wp.Exclude) == 0 && (len(wp.Include) == 0 || (len(wp.Include) == 1 && wp.Include[0] == actionsPolicyAllWorkflows)) {
		return nil
	}
	return []any{map[string]any{
		"include": flattenStringList(wp.Include),
		"exclude": flattenStringList(wp.Exclude),
	}}
}

// flattenActionsPolicyCommon sets the schema fields shared by all Actions policy resources.
func flattenActionsPolicyCommon(d *schema.ResourceData, policy *github.ActionsPolicy) error {
	var events []any
	var actors []any
	for _, rule := range policy.Rules {
		if rule == nil || rule.Parameters == nil {
			continue
		}
		switch rule.Type {
		case github.ActionsPolicyRuleTypeRestrictActionEvents:
			events = append(events, flattenStringList(rule.Parameters.AllowedEvents)...)
		case github.ActionsPolicyRuleTypeRestrictActionsActors:
			for _, a := range rule.Parameters.AllowedActors {
				actors = append(actors, map[string]any{"id": int(a.ID), "type": string(a.Type)})
			}
		}
	}

	values := map[string]any{
		"policy_id":      int(policy.ID),
		"name":           policy.Name,
		"enforcement":    string(policy.Enforcement),
		"workflow_path":  flattenActionsPolicyWorkflowPath(policy.Conditions),
		"allowed_events": events,
		"allowed_actor":  actors,
	}
	for k, v := range values {
		if err := d.Set(k, v); err != nil {
			return err
		}
	}
	return nil
}
