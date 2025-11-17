package github

import (
	"encoding/json"
	"log"
	"reflect"
	"sort"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRulesetObject(d *schema.ResourceData, org string) *github.Ruleset {
	isOrgLevel := len(org) > 0

	var source, sourceType string
	if isOrgLevel {
		source = org
		sourceType = "Organization"
	} else {
		source = d.Get("repository").(string)
		sourceType = "Repository"
	}

	return &github.Ruleset{
		Name:         d.Get("name").(string),
		Target:       github.String(d.Get("target").(string)),
		Source:       source,
		SourceType:   &sourceType,
		Enforcement:  d.Get("enforcement").(string),
		BypassActors: expandBypassActors(d.Get("bypass_actors").([]any)),
		Conditions:   expandConditions(d.Get("conditions").([]any), isOrgLevel),
		Rules:        expandRules(d.Get("rules").([]any), isOrgLevel),
	}
}

func expandBypassActors(input []any) []*github.BypassActor {
	if len(input) == 0 {
		// IMPORTANT:
		// Always return an empty slice ([]), not nil.
		// If this function returns nil, go-github serializes the field as `"bypass_actors": null`,
		// which causes GitHub API to reject the request with:
		//   422 "Invalid property /bypass_actors: data cannot be null."
		//
		// According to the GitHub REST API specification:
		//   - The "bypass_actors" field must be an array (even if empty).
		//   - Sending `null` is invalid; sending `[]` explicitly clears the list.
		// Reference:
		//   https://docs.github.com/en/rest/repos/rules#get-a-repository-ruleset
		return []*github.BypassActor{}
	}
	bypassActors := make([]*github.BypassActor, 0)

	for _, v := range input {
		inputMap := v.(map[string]any)
		actor := &github.BypassActor{}
		if v, ok := inputMap["actor_id"].(int); ok {
			if v == 0 {
				actor.ActorID = nil
			} else {
				actor.ActorID = github.Int64(int64(v))
			}
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

func flattenBypassActors(bypassActors []*github.BypassActor) []any {
	if bypassActors == nil {
		return []any{}
	}

	actorsSlice := make([]any, 0)
	for _, v := range bypassActors {
		actorMap := make(map[string]any)

		actorMap["actor_id"] = v.GetActorID()
		actorMap["actor_type"] = v.GetActorType()
		actorMap["bypass_mode"] = v.GetBypassMode()

		actorsSlice = append(actorsSlice, actorMap)
	}

	return actorsSlice
}

func expandConditions(input []any, org bool) *github.RulesetConditions {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	rulesetConditions := &github.RulesetConditions{}
	inputConditions := input[0].(map[string]any)

	// ref_name is available for both repo and org rulesets
	if v, ok := inputConditions["ref_name"].([]any); ok && v != nil && len(v) != 0 {
		inputRefName := v[0].(map[string]any)
		include := make([]string, 0)
		exclude := make([]string, 0)

		for _, v := range inputRefName["include"].([]any) {
			if v != nil {
				include = append(include, v.(string))
			}
		}

		for _, v := range inputRefName["exclude"].([]any) {
			if v != nil {
				exclude = append(exclude, v.(string))
			}
		}

		rulesetConditions.RefName = &github.RulesetRefConditionParameters{
			Include: include,
			Exclude: exclude,
		}
	}

	// org-only fields
	if org {
		// repository_name and repository_id
		if v, ok := inputConditions["repository_name"].([]any); ok && v != nil && len(v) != 0 {
			inputRepositoryName := v[0].(map[string]any)
			include := make([]string, 0)
			exclude := make([]string, 0)

			for _, v := range inputRepositoryName["include"].([]any) {
				if v != nil {
					include = append(include, v.(string))
				}
			}

			for _, v := range inputRepositoryName["exclude"].([]any) {
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
		} else if v, ok := inputConditions["repository_id"].([]any); ok && v != nil && len(v) != 0 {
			repositoryIDs := make([]int64, 0)

			for _, v := range v {
				if v != nil {
					repositoryIDs = append(repositoryIDs, int64(v.(int)))
				}
			}

			rulesetConditions.RepositoryID = &github.RulesetRepositoryIDsConditionParameters{RepositoryIDs: repositoryIDs}
		}
	}

	return rulesetConditions
}

func flattenConditions(conditions *github.RulesetConditions, org bool) []any {
	if conditions == nil || conditions.RefName == nil {
		return []any{}
	}

	conditionsMap := make(map[string]any)
	refNameSlice := make([]map[string]any, 0)

	refNameSlice = append(refNameSlice, map[string]any{
		"include": conditions.RefName.Include,
		"exclude": conditions.RefName.Exclude,
	})

	conditionsMap["ref_name"] = refNameSlice

	// org-only fields
	if org {
		repositoryNameSlice := make([]map[string]any, 0)

		if conditions.RepositoryName != nil {
			var protected bool

			if conditions.RepositoryName.Protected != nil {
				protected = *conditions.RepositoryName.Protected
			}

			repositoryNameSlice = append(repositoryNameSlice, map[string]any{
				"include":   conditions.RepositoryName.Include,
				"exclude":   conditions.RepositoryName.Exclude,
				"protected": protected,
			})
			conditionsMap["repository_name"] = repositoryNameSlice
		}

		if conditions.RepositoryID != nil {
			conditionsMap["repository_id"] = conditions.RepositoryID.RepositoryIDs
		}
	}

	return []any{conditionsMap}
}

func expandRules(input []any, org bool) []*github.RepositoryRule {
	if len(input) == 0 || input[0] == nil {
		return nil
	}

	rulesMap := input[0].(map[string]any)
	rulesSlice := make([]*github.RepositoryRule, 0)

	// First we expand rules without parameters
	if v, ok := rulesMap["creation"].(bool); ok && v {
		rulesSlice = append(rulesSlice, github.NewCreationRule())
	}

	if v, ok := rulesMap["update"].(bool); ok && v {
		params := github.UpdateAllowsFetchAndMergeRuleParameters{}
		if fetchAndMerge, ok := rulesMap["update"].(bool); ok && fetchAndMerge {
			params.UpdateAllowsFetchAndMerge = true
		} else {
			params.UpdateAllowsFetchAndMerge = false
		}
		rulesSlice = append(rulesSlice, github.NewUpdateRule(&params))
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

	// Required deployments rule
	if !org {
		if v, ok := rulesMap["required_deployments"].([]any); ok && len(v) != 0 {
			requiredDeploymentsMap := make(map[string]any)
			// If the rule's block is present but has an empty environments list
			if v[0] == nil {
				requiredDeploymentsMap["required_deployment_environments"] = make([]any, 0)
			} else {
				requiredDeploymentsMap = v[0].(map[string]any)
			}
			envs := make([]string, 0)
			for _, v := range requiredDeploymentsMap["required_deployment_environments"].([]any) {
				envs = append(envs, v.(string))
			}

			params := &github.RequiredDeploymentEnvironmentsRuleParameters{
				RequiredDeploymentEnvironments: envs,
			}

			rulesSlice = append(rulesSlice, github.NewRequiredDeploymentsRule(params))
		}
	}

	// Pattern parameter rules
	for _, k := range []string{"commit_message_pattern", "commit_author_email_pattern", "committer_email_pattern", "branch_name_pattern", "tag_name_pattern"} {
		if v, ok := rulesMap[k].([]any); ok && len(v) != 0 {
			patternParametersMap := v[0].(map[string]any)

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

	// Pull request rule
	if v, ok := rulesMap["pull_request"].([]any); ok && len(v) != 0 {
		pullRequestMap := v[0].(map[string]any)
		params := &github.PullRequestRuleParameters{
			DismissStaleReviewsOnPush:      pullRequestMap["dismiss_stale_reviews_on_push"].(bool),
			RequireCodeOwnerReview:         pullRequestMap["require_code_owner_review"].(bool),
			RequireLastPushApproval:        pullRequestMap["require_last_push_approval"].(bool),
			RequiredApprovingReviewCount:   pullRequestMap["required_approving_review_count"].(int),
			RequiredReviewThreadResolution: pullRequestMap["required_review_thread_resolution"].(bool),
		}

		rulesSlice = append(rulesSlice, github.NewPullRequestRule(params))
	}

	// Merge queue rule
	if v, ok := rulesMap["merge_queue"].([]any); ok && len(v) != 0 {
		mergeQueueMap := v[0].(map[string]any)
		params := &github.MergeQueueRuleParameters{
			CheckResponseTimeoutMinutes:  mergeQueueMap["check_response_timeout_minutes"].(int),
			GroupingStrategy:             mergeQueueMap["grouping_strategy"].(string),
			MaxEntriesToBuild:            mergeQueueMap["max_entries_to_build"].(int),
			MaxEntriesToMerge:            mergeQueueMap["max_entries_to_merge"].(int),
			MergeMethod:                  mergeQueueMap["merge_method"].(string),
			MinEntriesToMerge:            mergeQueueMap["min_entries_to_merge"].(int),
			MinEntriesToMergeWaitMinutes: mergeQueueMap["min_entries_to_merge_wait_minutes"].(int),
		}

		rulesSlice = append(rulesSlice, github.NewMergeQueueRule(params))
	}

	// Required status checks rule
	if v, ok := rulesMap["required_status_checks"].([]any); ok && len(v) != 0 {
		requiredStatusMap := v[0].(map[string]any)
		requiredStatusChecks := make([]github.RuleRequiredStatusChecks, 0)

		if requiredStatusChecksInput, ok := requiredStatusMap["required_check"]; ok {

			requiredStatusChecksSet := requiredStatusChecksInput.(*schema.Set)
			for _, checkMap := range requiredStatusChecksSet.List() {
				check := checkMap.(map[string]any)
				integrationID := github.Int64(int64(check["integration_id"].(int)))

				params := github.RuleRequiredStatusChecks{
					Context: check["context"].(string),
				}

				if *integrationID != 0 {
					params.IntegrationID = integrationID
				}

				requiredStatusChecks = append(requiredStatusChecks, params)
			}
		}

		doNotEnforceOnCreate := requiredStatusMap["do_not_enforce_on_create"].(bool)
		params := &github.RequiredStatusChecksRuleParameters{
			RequiredStatusChecks:             requiredStatusChecks,
			StrictRequiredStatusChecksPolicy: requiredStatusMap["strict_required_status_checks_policy"].(bool),
			DoNotEnforceOnCreate:             &doNotEnforceOnCreate,
		}
		rulesSlice = append(rulesSlice, github.NewRequiredStatusChecksRule(params))
	}

	// Required workflows to pass before merging rule
	if v, ok := rulesMap["required_workflows"].([]any); ok && len(v) != 0 {
		requiredWorkflowsMap := v[0].(map[string]any)
		requiredWorkflows := make([]*github.RuleRequiredWorkflow, 0)

		if requiredWorkflowsInput, ok := requiredWorkflowsMap["required_workflow"]; ok {

			requiredWorkflowsSet := requiredWorkflowsInput.(*schema.Set)
			for _, workflowMap := range requiredWorkflowsSet.List() {
				workflow := workflowMap.(map[string]any)

				// Get all parameters
				repositoryID := github.Int64(int64(workflow["repository_id"].(int)))
				ref := github.String(workflow["ref"].(string))

				params := &github.RuleRequiredWorkflow{
					RepositoryID: repositoryID,
					Path:         workflow["path"].(string),
					Ref:          ref,
				}

				requiredWorkflows = append(requiredWorkflows, params)
			}
		}

		params := &github.RequiredWorkflowsRuleParameters{
			DoNotEnforceOnCreate: requiredWorkflowsMap["do_not_enforce_on_create"].(bool),
			RequiredWorkflows:    requiredWorkflows,
		}
		rulesSlice = append(rulesSlice, github.NewRequiredWorkflowsRule(params))
	}

	// Required code scanning to pass before merging rule
	if v, ok := rulesMap["required_code_scanning"].([]any); ok && len(v) != 0 {
		requiredCodeScanningMap := v[0].(map[string]any)
		requiredCodeScanningTools := make([]*github.RuleRequiredCodeScanningTool, 0)

		if requiredCodeScanningInput, ok := requiredCodeScanningMap["required_code_scanning_tool"]; ok {

			requiredCodeScanningSet := requiredCodeScanningInput.(*schema.Set)
			for _, codeScanningMap := range requiredCodeScanningSet.List() {
				codeScanningTool := codeScanningMap.(map[string]any)

				// Get all parameters
				alertsThreshold := github.String(codeScanningTool["alerts_threshold"].(string))
				securityAlertsThreshold := github.String(codeScanningTool["security_alerts_threshold"].(string))
				tool := github.String(codeScanningTool["tool"].(string))

				params := &github.RuleRequiredCodeScanningTool{
					AlertsThreshold:         *alertsThreshold,
					SecurityAlertsThreshold: *securityAlertsThreshold,
					Tool:                    *tool,
				}

				requiredCodeScanningTools = append(requiredCodeScanningTools, params)
			}
		}

		params := &github.RequiredCodeScanningRuleParameters{
			RequiredCodeScanningTools: requiredCodeScanningTools,
		}
		rulesSlice = append(rulesSlice, github.NewRequiredCodeScanningRule(params))
	}

	// file_path_restriction rule
	if v, ok := rulesMap["file_path_restriction"].([]any); ok && len(v) != 0 {
		filePathRestrictionMap := v[0].(map[string]any)
		restrictedFilePaths := make([]string, 0)
		for _, path := range filePathRestrictionMap["restricted_file_paths"].([]any) {
			restrictedFilePaths = append(restrictedFilePaths, path.(string))
		}
		params := &github.RuleFileParameters{
			RestrictedFilePaths: &restrictedFilePaths,
		}
		rulesSlice = append(rulesSlice, github.NewFilePathRestrictionRule(params))
	}

	// max_file_size rule
	if v, ok := rulesMap["max_file_size"].([]any); ok && len(v) != 0 {
		maxFileSizeMap := v[0].(map[string]any)
		maxFileSize := int64(maxFileSizeMap["max_file_size"].(float64))
		params := &github.RuleMaxFileSizeParameters{
			MaxFileSize: maxFileSize,
		}
		rulesSlice = append(rulesSlice, github.NewMaxFileSizeRule(params))

	}

	// max_file_path_length rule
	if v, ok := rulesMap["max_file_path_length"].([]any); ok && len(v) != 0 {
		maxFilePathLengthMap := v[0].(map[string]any)
		maxFilePathLength := maxFilePathLengthMap["max_file_path_length"].(int)
		params := &github.RuleMaxFilePathLengthParameters{
			MaxFilePathLength: maxFilePathLength,
		}
		rulesSlice = append(rulesSlice, github.NewMaxFilePathLengthRule(params))

	}

	// file_extension_restriction rule
	if v, ok := rulesMap["file_extension_restriction"].([]any); ok && len(v) != 0 {
		fileExtensionRestrictionMap := v[0].(map[string]any)
		restrictedFileExtensions := make([]string, 0)
		for _, extension := range fileExtensionRestrictionMap["restricted_file_extensions"].([]any) {
			restrictedFileExtensions = append(restrictedFileExtensions, extension.(string))
		}
		params := &github.RuleFileExtensionRestrictionParameters{
			RestrictedFileExtensions: restrictedFileExtensions,
		}
		rulesSlice = append(rulesSlice, github.NewFileExtensionRestrictionRule(params))
	}

	return rulesSlice
}

func flattenRules(rules []*github.RepositoryRule, org bool) []any {
	if len(rules) == 0 || rules == nil {
		return []any{}
	}

	rulesMap := make(map[string]any)
	for _, v := range rules {
		switch v.Type {
		case "creation", "deletion", "required_linear_history", "required_signatures", "non_fast_forward":
			rulesMap[v.Type] = true

		case "update":
			var params github.UpdateAllowsFetchAndMergeRuleParameters
			if v.Parameters != nil {
				err := json.Unmarshal(*v.Parameters, &params)
				if err != nil {
					log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
						v.Type, v.Parameters)
				}
				rulesMap["update_allows_fetch_and_merge"] = params.UpdateAllowsFetchAndMerge
			} else {
				rulesMap["update_allows_fetch_and_merge"] = false
			}
			rulesMap[v.Type] = true

		case "commit_message_pattern", "commit_author_email_pattern", "committer_email_pattern", "branch_name_pattern", "tag_name_pattern":
			var params github.RulePatternParameters
			var name string
			var negate bool

			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}

			if params.Name != nil {
				name = *params.Name
			}
			if params.Negate != nil {
				negate = *params.Negate
			}

			rule := make(map[string]any)
			rule["name"] = name
			rule["negate"] = negate
			rule["operator"] = params.Operator
			rule["pattern"] = params.Pattern
			rulesMap[v.Type] = []map[string]any{rule}

		case "required_deployments":
			if !org {
				var params github.RequiredDeploymentEnvironmentsRuleParameters

				err := json.Unmarshal(*v.Parameters, &params)
				if err != nil {
					log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
						v.Type, v.Parameters)
				}

				rule := make(map[string]any)
				rule["required_deployment_environments"] = params.RequiredDeploymentEnvironments
				rulesMap[v.Type] = []map[string]any{rule}
			}

		case "pull_request":
			var params github.PullRequestRuleParameters

			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}

			rule := make(map[string]any)
			rule["dismiss_stale_reviews_on_push"] = params.DismissStaleReviewsOnPush
			rule["require_code_owner_review"] = params.RequireCodeOwnerReview
			rule["require_last_push_approval"] = params.RequireLastPushApproval
			rule["required_approving_review_count"] = params.RequiredApprovingReviewCount
			rule["required_review_thread_resolution"] = params.RequiredReviewThreadResolution
			rulesMap[v.Type] = []map[string]any{rule}

		case "required_status_checks":
			var params github.RequiredStatusChecksRuleParameters

			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}

			requiredStatusChecksSlice := make([]map[string]any, 0)
			for _, check := range params.RequiredStatusChecks {
				integrationID := int64(0)
				if check.IntegrationID != nil {
					integrationID = *check.IntegrationID
				}
				requiredStatusChecksSlice = append(requiredStatusChecksSlice, map[string]any{
					"context":        check.Context,
					"integration_id": integrationID,
				})
			}

			rule := make(map[string]any)
			rule["required_check"] = requiredStatusChecksSlice
			rule["strict_required_status_checks_policy"] = params.StrictRequiredStatusChecksPolicy
			rule["do_not_enforce_on_create"] = params.DoNotEnforceOnCreate
			rulesMap[v.Type] = []map[string]any{rule}

		case "workflows":
			var params github.RequiredWorkflowsRuleParameters

			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}

			requiredWorkflowsSlice := make([]map[string]any, 0)
			for _, check := range params.RequiredWorkflows {
				requiredWorkflowsSlice = append(requiredWorkflowsSlice, map[string]any{
					"repository_id": check.RepositoryID,
					"path":          check.Path,
					"ref":           check.Ref,
				})
			}

			rule := make(map[string]any)
			rule["do_not_enforce_on_create"] = params.DoNotEnforceOnCreate
			rule["required_workflow"] = requiredWorkflowsSlice
			rulesMap["required_workflows"] = []map[string]any{rule}

		case "code_scanning":
			var params github.RequiredCodeScanningRuleParameters

			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}

			requiredCodeScanningSlice := make([]map[string]any, 0)
			for _, check := range params.RequiredCodeScanningTools {
				requiredCodeScanningSlice = append(requiredCodeScanningSlice, map[string]any{
					"alerts_threshold":          check.AlertsThreshold,
					"security_alerts_threshold": check.SecurityAlertsThreshold,
					"tool":                      check.Tool,
				})
			}

			rule := make(map[string]any)
			rule["required_code_scanning_tool"] = requiredCodeScanningSlice
			rulesMap["required_code_scanning"] = []map[string]any{rule}

		case "merge_queue":
			var params github.MergeQueueRuleParameters

			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}

			rule := make(map[string]any)
			rule["check_response_timeout_minutes"] = params.CheckResponseTimeoutMinutes
			rule["grouping_strategy"] = params.GroupingStrategy
			rule["max_entries_to_build"] = params.MaxEntriesToBuild
			rule["max_entries_to_merge"] = params.MaxEntriesToMerge
			rule["merge_method"] = params.MergeMethod
			rule["min_entries_to_merge"] = params.MinEntriesToMerge
			rule["min_entries_to_merge_wait_minutes"] = params.MinEntriesToMergeWaitMinutes
			rulesMap[v.Type] = []map[string]any{rule}

		case "file_path_restriction":
			var params github.RuleFileParameters
			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}
			rule := make(map[string]any)
			rule["restricted_file_paths"] = params.GetRestrictedFilePaths()
			rulesMap[v.Type] = []map[string]any{rule}

		case "max_file_size":
			var params github.RuleMaxFileSizeParameters
			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}
			rule := make(map[string]any)
			rule["max_file_size"] = params.MaxFileSize
			rulesMap[v.Type] = []map[string]any{rule}

		case "max_file_path_length":
			var params github.RuleMaxFilePathLengthParameters
			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}
			rule := make(map[string]any)
			rule["max_file_path_length"] = params.MaxFilePathLength
			rulesMap[v.Type] = []map[string]any{rule}

		case "file_extension_restriction":
			var params github.RuleFileExtensionRestrictionParameters
			err := json.Unmarshal(*v.Parameters, &params)
			if err != nil {
				log.Printf("[INFO] Unexpected error unmarshalling rule %s with parameters: %v",
					v.Type, v.Parameters)
			}
			rule := make(map[string]any)
			rule["restricted_file_extensions"] = params.RestrictedFileExtensions
			rulesMap[v.Type] = []map[string]any{rule}

		default:
			// Handle unknown rule types (like Copilot code review, etc.) gracefully
			// Log the unknown rule type but don't cause Terraform to fail or see a diff
			log.Printf("[DEBUG] Ignoring unknown repository rule type: %s. This rule was likely added outside of Terraform (e.g., via GitHub UI) and is not yet supported by the provider.", v.Type)
			// Note: We intentionally don't add this to rulesMap to avoid causing diffs for rules that aren't managed by Terraform
		}
	}

	return []any{rulesMap}
}

func bypassActorsDiffSuppressFunc(k, o, n string, d *schema.ResourceData) bool {
	// If the length has changed, no need to suppress
	if k == "bypass_actors.#" {
		return o == n
	}

	// Get change to bypass actors
	oba, nba := d.GetChange("bypass_actors")
	oldBypassActors := oba.([]any)
	newBypassActors := nba.([]any)

	sort.SliceStable(oldBypassActors, func(i, j int) bool {
		return oldBypassActors[i].(map[string]any)["actor_id"].(int) > oldBypassActors[j].(map[string]any)["actor_id"].(int)
	})
	sort.SliceStable(newBypassActors, func(i, j int) bool {
		return newBypassActors[i].(map[string]any)["actor_id"].(int) > newBypassActors[j].(map[string]any)["actor_id"].(int)
	})

	return reflect.DeepEqual(oldBypassActors, newBypassActors)
}
