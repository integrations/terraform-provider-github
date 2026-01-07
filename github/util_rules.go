package github

import (
	"reflect"
	"sort"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Helper function to safely convert interface{} to int, handling both int and float64.
func toInt(v any) int {
	switch val := v.(type) {
	case int:
		return val
	case float64:
		return int(val)
	case int64:
		return int(val)
	default:
		return 0
	}
}

// Helper function to safely convert interface{} to int64, handling both int and float64.
func toInt64(v any) int64 {
	switch val := v.(type) {
	case int:
		return int64(val)
	case int64:
		return val
	case float64:
		return int64(val)
	default:
		return 0
	}
}

func resourceGithubRulesetObject(d *schema.ResourceData, org string) github.RepositoryRuleset {
	isOrgLevel := len(org) > 0

	var source, sourceType string
	if isOrgLevel {
		source = org
		sourceType = "Organization"
	} else {
		source = d.Get("repository").(string)
		sourceType = "Repository"
	}

	target := github.RulesetTarget(d.Get("target").(string))
	enforcement := github.RulesetEnforcement(d.Get("enforcement").(string))
	sourceTypeEnum := github.RulesetSourceType(sourceType)

	return github.RepositoryRuleset{
		Name:         d.Get("name").(string),
		Target:       &target,
		Source:       source,
		SourceType:   &sourceTypeEnum,
		Enforcement:  enforcement,
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
		if actorIDVal, ok := inputMap["actor_id"]; ok {
			actorID := toInt(actorIDVal)
			if actorID == 0 {
				actor.ActorID = nil
			} else {
				actor.ActorID = github.Ptr(int64(actorID))
			}
		}

		if v, ok := inputMap["actor_type"].(string); ok {
			actorType := github.BypassActorType(v)
			actor.ActorType = &actorType
		}

		if v, ok := inputMap["bypass_mode"].(string); ok {
			bypassMode := github.BypassMode(v)
			actor.BypassMode = &bypassMode
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

func expandConditions(input []any, org bool) *github.RepositoryRulesetConditions {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	rulesetConditions := &github.RepositoryRulesetConditions{}
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

		rulesetConditions.RefName = &github.RepositoryRulesetRefConditionParameters{
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

			rulesetConditions.RepositoryName = &github.RepositoryRulesetRepositoryNamesConditionParameters{
				Include:   include,
				Exclude:   exclude,
				Protected: &protected,
			}
		} else if v, ok := inputConditions["repository_id"].([]any); ok && v != nil && len(v) != 0 {
			repositoryIDs := make([]int64, 0)

			for _, v := range v {
				if v != nil {
					repositoryIDs = append(repositoryIDs, toInt64(v))
				}
			}

			rulesetConditions.RepositoryID = &github.RepositoryRulesetRepositoryIDsConditionParameters{RepositoryIDs: repositoryIDs}
		}
	}

	return rulesetConditions
}

func flattenConditions(conditions *github.RepositoryRulesetConditions, org bool) []any {
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

func expandRules(input []any, org bool) *github.RepositoryRulesetRules {
	if len(input) == 0 || input[0] == nil {
		return &github.RepositoryRulesetRules{}
	}

	rulesMap := input[0].(map[string]any)
	rulesetRules := &github.RepositoryRulesetRules{}

	// Simple rules without parameters
	if v, ok := rulesMap["creation"].(bool); ok && v {
		rulesetRules.Creation = &github.EmptyRuleParameters{}
	}

	if v, ok := rulesMap["deletion"].(bool); ok && v {
		rulesetRules.Deletion = &github.EmptyRuleParameters{}
	}

	if v, ok := rulesMap["required_linear_history"].(bool); ok && v {
		rulesetRules.RequiredLinearHistory = &github.EmptyRuleParameters{}
	}

	if v, ok := rulesMap["required_signatures"].(bool); ok && v {
		rulesetRules.RequiredSignatures = &github.EmptyRuleParameters{}
	}

	if v, ok := rulesMap["non_fast_forward"].(bool); ok && v {
		rulesetRules.NonFastForward = &github.EmptyRuleParameters{}
	}

	// Update rule with parameters
	if v, ok := rulesMap["update"].(bool); ok && v {
		updateParams := &github.UpdateRuleParameters{}
		if updateMerge, ok := rulesMap["update_allows_fetch_and_merge"].(bool); ok {
			updateParams.UpdateAllowsFetchAndMerge = updateMerge
		}
		rulesetRules.Update = updateParams
	}

	// Required deployments rule
	if v, ok := rulesMap["required_deployments"].([]any); ok && len(v) != 0 {
		requiredDeploymentsMap := v[0].(map[string]any)
		envs := make([]string, 0)
		for _, env := range requiredDeploymentsMap["required_deployment_environments"].([]any) {
			envs = append(envs, env.(string))
		}
		rulesetRules.RequiredDeployments = &github.RequiredDeploymentsRuleParameters{
			RequiredDeploymentEnvironments: envs,
		}
	}

	// Pull request rule
	if v, ok := rulesMap["pull_request"].([]any); ok && len(v) != 0 {
		pullRequestMap := v[0].(map[string]any)
		params := &github.PullRequestRuleParameters{
			AllowedMergeMethods:            []github.PullRequestMergeMethod{github.PullRequestMergeMethodMerge, github.PullRequestMergeMethodSquash, github.PullRequestMergeMethodRebase},
			DismissStaleReviewsOnPush:      pullRequestMap["dismiss_stale_reviews_on_push"].(bool),
			RequireCodeOwnerReview:         pullRequestMap["require_code_owner_review"].(bool),
			RequireLastPushApproval:        pullRequestMap["require_last_push_approval"].(bool),
			RequiredApprovingReviewCount:   toInt(pullRequestMap["required_approving_review_count"]),
			RequiredReviewThreadResolution: pullRequestMap["required_review_thread_resolution"].(bool),
		}
		rulesetRules.PullRequest = params
	}

	// Merge queue rule
	if v, ok := rulesMap["merge_queue"].([]any); ok && len(v) != 0 {
		mergeQueueMap := v[0].(map[string]any)
		params := &github.MergeQueueRuleParameters{
			CheckResponseTimeoutMinutes:  toInt(mergeQueueMap["check_response_timeout_minutes"]),
			GroupingStrategy:             github.MergeGroupingStrategy(mergeQueueMap["grouping_strategy"].(string)),
			MaxEntriesToBuild:            toInt(mergeQueueMap["max_entries_to_build"]),
			MaxEntriesToMerge:            toInt(mergeQueueMap["max_entries_to_merge"]),
			MergeMethod:                  github.MergeQueueMergeMethod(mergeQueueMap["merge_method"].(string)),
			MinEntriesToMerge:            toInt(mergeQueueMap["min_entries_to_merge"]),
			MinEntriesToMergeWaitMinutes: toInt(mergeQueueMap["min_entries_to_merge_wait_minutes"]),
		}
		rulesetRules.MergeQueue = params
	}

	// Required status checks rule
	if v, ok := rulesMap["required_status_checks"].([]any); ok && len(v) != 0 {
		requiredStatusMap := v[0].(map[string]any)
		requiredStatusChecks := make([]*github.RuleStatusCheck, 0)

		if requiredStatusChecksInput, ok := requiredStatusMap["required_check"]; ok {
			requiredStatusChecksSet := requiredStatusChecksInput.(*schema.Set)
			for _, checkMap := range requiredStatusChecksSet.List() {
				check := checkMap.(map[string]any)
				integrationID := toInt64(check["integration_id"])

				params := &github.RuleStatusCheck{
					Context: check["context"].(string),
				}

				if integrationID != 0 {
					params.IntegrationID = github.Ptr(integrationID)
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
		rulesetRules.RequiredStatusChecks = params
	}

	// Pattern parameter rules
	patternRules := map[string]*github.PatternRuleParameters{
		"commit_message_pattern":      nil,
		"commit_author_email_pattern": nil,
		"committer_email_pattern":     nil,
		"branch_name_pattern":         nil,
		"tag_name_pattern":            nil,
	}

	for k := range patternRules {
		if v, ok := rulesMap[k].([]any); ok && len(v) != 0 {
			patternParametersMap := v[0].(map[string]any)

			name := patternParametersMap["name"].(string)
			negate := patternParametersMap["negate"].(bool)

			params := &github.PatternRuleParameters{
				Name:     &name,
				Negate:   &negate,
				Operator: github.PatternRuleOperator(patternParametersMap["operator"].(string)),
				Pattern:  patternParametersMap["pattern"].(string),
			}

			switch k {
			case "commit_message_pattern":
				rulesetRules.CommitMessagePattern = params
			case "commit_author_email_pattern":
				rulesetRules.CommitAuthorEmailPattern = params
			case "committer_email_pattern":
				rulesetRules.CommitterEmailPattern = params
			case "branch_name_pattern":
				rulesetRules.BranchNamePattern = params
			case "tag_name_pattern":
				rulesetRules.TagNamePattern = params
			}
		}
	}

	// Required workflows rule (org-only)
	if org {
		if v, ok := rulesMap["required_workflows"].([]any); ok && len(v) != 0 {
			requiredWorkflowsMap := v[0].(map[string]any)
			requiredWorkflows := make([]*github.RuleWorkflow, 0)

			if requiredWorkflowsInput, ok := requiredWorkflowsMap["required_workflow"]; ok {
				requiredWorkflowsSet := requiredWorkflowsInput.(*schema.Set)
				for _, workflowMap := range requiredWorkflowsSet.List() {
					workflow := workflowMap.(map[string]any)

					params := &github.RuleWorkflow{
						RepositoryID: github.Ptr(toInt64(workflow["repository_id"])),
						Path:         workflow["path"].(string),
						Ref:          github.Ptr(workflow["ref"].(string)),
					}

					requiredWorkflows = append(requiredWorkflows, params)
				}
			}

			doNotEnforceOnCreate := requiredWorkflowsMap["do_not_enforce_on_create"].(bool)
			params := &github.WorkflowsRuleParameters{
				DoNotEnforceOnCreate: &doNotEnforceOnCreate,
				Workflows:            requiredWorkflows,
			}
			rulesetRules.Workflows = params
		}
	}

	// Required code scanning rule
	if v, ok := rulesMap["required_code_scanning"].([]any); ok && len(v) != 0 {
		requiredCodeScanningMap := v[0].(map[string]any)
		requiredCodeScanningTools := make([]*github.RuleCodeScanningTool, 0)

		if requiredCodeScanningInput, ok := requiredCodeScanningMap["required_code_scanning_tool"]; ok {
			requiredCodeScanningSet := requiredCodeScanningInput.(*schema.Set)
			for _, codeScanningMap := range requiredCodeScanningSet.List() {
				codeScanningTool := codeScanningMap.(map[string]any)

				params := &github.RuleCodeScanningTool{
					AlertsThreshold:         github.CodeScanningAlertsThreshold(codeScanningTool["alerts_threshold"].(string)),
					SecurityAlertsThreshold: github.CodeScanningSecurityAlertsThreshold(codeScanningTool["security_alerts_threshold"].(string)),
					Tool:                    codeScanningTool["tool"].(string),
				}

				requiredCodeScanningTools = append(requiredCodeScanningTools, params)
			}
		}

		params := &github.CodeScanningRuleParameters{
			CodeScanningTools: requiredCodeScanningTools,
		}
		rulesetRules.CodeScanning = params
	}

	// File path restriction rule
	if v, ok := rulesMap["file_path_restriction"].([]any); ok && len(v) != 0 {
		filePathRestrictionMap := v[0].(map[string]any)
		restrictedFilePaths := make([]string, 0)
		for _, path := range filePathRestrictionMap["restricted_file_paths"].([]any) {
			restrictedFilePaths = append(restrictedFilePaths, path.(string))
		}
		params := &github.FilePathRestrictionRuleParameters{
			RestrictedFilePaths: restrictedFilePaths,
		}
		rulesetRules.FilePathRestriction = params
	}

	// Max file size rule
	if v, ok := rulesMap["max_file_size"].([]any); ok && len(v) != 0 {
		maxFileSizeMap := v[0].(map[string]any)
		maxFileSize := toInt64(maxFileSizeMap["max_file_size"])
		params := &github.MaxFileSizeRuleParameters{
			MaxFileSize: maxFileSize,
		}
		rulesetRules.MaxFileSize = params
	}

	// Max file path length rule
	if v, ok := rulesMap["max_file_path_length"].([]any); ok && len(v) != 0 {
		maxFilePathLengthMap := v[0].(map[string]any)
		maxFilePathLength := toInt(maxFilePathLengthMap["max_file_path_length"])
		params := &github.MaxFilePathLengthRuleParameters{
			MaxFilePathLength: maxFilePathLength,
		}
		rulesetRules.MaxFilePathLength = params
	}

	// File extension restriction rule
	if v, ok := rulesMap["file_extension_restriction"].([]any); ok && len(v) != 0 {
		fileExtensionRestrictionMap := v[0].(map[string]any)
		restrictedFileExtensions := make([]string, 0)

		// Handle schema.TypeSet
		extensionSet := fileExtensionRestrictionMap["restricted_file_extensions"].(*schema.Set)
		for _, extension := range extensionSet.List() {
			restrictedFileExtensions = append(restrictedFileExtensions, extension.(string))
		}
		params := &github.FileExtensionRestrictionRuleParameters{
			RestrictedFileExtensions: restrictedFileExtensions,
		}
		rulesetRules.FileExtensionRestriction = params
	}

	return rulesetRules
}

func flattenRules(rules *github.RepositoryRulesetRules, org bool) []any {
	if rules == nil {
		return []any{}
	}

	rulesMap := make(map[string]any)

	// Simple boolean rules - explicitly set all to false first, then override with true if present
	rulesMap["creation"] = rules.Creation != nil
	rulesMap["deletion"] = rules.Deletion != nil
	rulesMap["required_linear_history"] = rules.RequiredLinearHistory != nil
	rulesMap["required_signatures"] = rules.RequiredSignatures != nil
	rulesMap["non_fast_forward"] = rules.NonFastForward != nil

	// Update rule with parameters
	if rules.Update != nil {
		rulesMap["update"] = true
		rulesMap["update_allows_fetch_and_merge"] = rules.Update.UpdateAllowsFetchAndMerge
	} else {
		rulesMap["update"] = false
		rulesMap["update_allows_fetch_and_merge"] = false
	} // Required deployments rule
	if rules.RequiredDeployments != nil {
		requiredDeploymentsSlice := make([]map[string]any, 0)
		requiredDeploymentsSlice = append(requiredDeploymentsSlice, map[string]any{
			"required_deployment_environments": rules.RequiredDeployments.RequiredDeploymentEnvironments,
		})
		rulesMap["required_deployments"] = requiredDeploymentsSlice
	}

	// Pull request rule
	if rules.PullRequest != nil {
		pullRequestSlice := make([]map[string]any, 0)
		pullRequestSlice = append(pullRequestSlice, map[string]any{
			"dismiss_stale_reviews_on_push":     rules.PullRequest.DismissStaleReviewsOnPush,
			"require_code_owner_review":         rules.PullRequest.RequireCodeOwnerReview,
			"require_last_push_approval":        rules.PullRequest.RequireLastPushApproval,
			"required_approving_review_count":   rules.PullRequest.RequiredApprovingReviewCount,
			"required_review_thread_resolution": rules.PullRequest.RequiredReviewThreadResolution,
		})
		rulesMap["pull_request"] = pullRequestSlice
	}

	// Merge queue rule
	if rules.MergeQueue != nil {
		mergeQueueSlice := make([]map[string]any, 0)
		mergeQueueSlice = append(mergeQueueSlice, map[string]any{
			"check_response_timeout_minutes":    rules.MergeQueue.CheckResponseTimeoutMinutes,
			"grouping_strategy":                 string(rules.MergeQueue.GroupingStrategy),
			"max_entries_to_build":              rules.MergeQueue.MaxEntriesToBuild,
			"max_entries_to_merge":              rules.MergeQueue.MaxEntriesToMerge,
			"merge_method":                      string(rules.MergeQueue.MergeMethod),
			"min_entries_to_merge":              rules.MergeQueue.MinEntriesToMerge,
			"min_entries_to_merge_wait_minutes": rules.MergeQueue.MinEntriesToMergeWaitMinutes,
		})
		rulesMap["merge_queue"] = mergeQueueSlice
	}

	// Required status checks rule
	if rules.RequiredStatusChecks != nil {
		requiredStatusSlice := make([]map[string]any, 0)
		requiredChecks := make([]map[string]any, 0)

		for _, check := range rules.RequiredStatusChecks.RequiredStatusChecks {
			checkMap := map[string]any{
				"context": check.Context,
			}
			if check.IntegrationID != nil {
				checkMap["integration_id"] = int(*check.IntegrationID)
			} else {
				checkMap["integration_id"] = 0
			}
			requiredChecks = append(requiredChecks, checkMap)
		}

		statusChecksMap := map[string]any{
			"required_check":                       requiredChecks,
			"strict_required_status_checks_policy": rules.RequiredStatusChecks.StrictRequiredStatusChecksPolicy,
		}

		if rules.RequiredStatusChecks.DoNotEnforceOnCreate != nil {
			statusChecksMap["do_not_enforce_on_create"] = *rules.RequiredStatusChecks.DoNotEnforceOnCreate
		} else {
			statusChecksMap["do_not_enforce_on_create"] = false
		}

		requiredStatusSlice = append(requiredStatusSlice, statusChecksMap)
		rulesMap["required_status_checks"] = requiredStatusSlice
	}

	// Pattern parameter rules
	patternRules := map[string]*github.PatternRuleParameters{
		"commit_message_pattern":      rules.CommitMessagePattern,
		"commit_author_email_pattern": rules.CommitAuthorEmailPattern,
		"committer_email_pattern":     rules.CommitterEmailPattern,
		"branch_name_pattern":         rules.BranchNamePattern,
		"tag_name_pattern":            rules.TagNamePattern,
	}

	for k, v := range patternRules {
		if v != nil {
			patternSlice := make([]map[string]any, 0)
			patternMap := map[string]any{
				"operator": string(v.Operator),
				"pattern":  v.Pattern,
			}
			if v.Name != nil {
				patternMap["name"] = *v.Name
			}
			if v.Negate != nil {
				patternMap["negate"] = *v.Negate
			}
			patternSlice = append(patternSlice, patternMap)
			rulesMap[k] = patternSlice
		}
	}

	// Required workflows rule (org-only)
	if org && rules.Workflows != nil {
		requiredWorkflowsSlice := make([]map[string]any, 0)
		requiredWorkflows := make([]map[string]any, 0)

		for _, workflow := range rules.Workflows.Workflows {
			workflowMap := map[string]any{
				"path": workflow.Path,
			}
			if workflow.RepositoryID != nil {
				workflowMap["repository_id"] = int(*workflow.RepositoryID)
			}
			if workflow.Ref != nil {
				workflowMap["ref"] = *workflow.Ref
			}
			requiredWorkflows = append(requiredWorkflows, workflowMap)
		}

		workflowsMap := map[string]any{
			"required_workflow": requiredWorkflows,
		}

		if rules.Workflows.DoNotEnforceOnCreate != nil {
			workflowsMap["do_not_enforce_on_create"] = *rules.Workflows.DoNotEnforceOnCreate
		} else {
			workflowsMap["do_not_enforce_on_create"] = false
		}

		requiredWorkflowsSlice = append(requiredWorkflowsSlice, workflowsMap)
		rulesMap["required_workflows"] = requiredWorkflowsSlice
	}

	// Required code scanning rule
	if rules.CodeScanning != nil {
		requiredCodeScanningSlice := make([]map[string]any, 0)
		requiredCodeScanningTools := make([]map[string]any, 0)

		for _, tool := range rules.CodeScanning.CodeScanningTools {
			toolMap := map[string]any{
				"alerts_threshold":          string(tool.AlertsThreshold),
				"security_alerts_threshold": string(tool.SecurityAlertsThreshold),
				"tool":                      tool.Tool,
			}
			requiredCodeScanningTools = append(requiredCodeScanningTools, toolMap)
		}

		codeScanningMap := map[string]any{
			"required_code_scanning_tool": requiredCodeScanningTools,
		}

		requiredCodeScanningSlice = append(requiredCodeScanningSlice, codeScanningMap)
		rulesMap["required_code_scanning"] = requiredCodeScanningSlice
	}

	// File path restriction rule
	if rules.FilePathRestriction != nil {
		filePathRestrictionSlice := make([]map[string]any, 0)
		filePathRestrictionSlice = append(filePathRestrictionSlice, map[string]any{
			"restricted_file_paths": rules.FilePathRestriction.RestrictedFilePaths,
		})
		rulesMap["file_path_restriction"] = filePathRestrictionSlice
	}

	// Max file size rule
	if rules.MaxFileSize != nil {
		maxFileSizeSlice := make([]map[string]any, 0)
		maxFileSizeSlice = append(maxFileSizeSlice, map[string]any{
			"max_file_size": rules.MaxFileSize.MaxFileSize,
		})
		rulesMap["max_file_size"] = maxFileSizeSlice
	}

	// Max file path length rule
	if rules.MaxFilePathLength != nil {
		maxFilePathLengthSlice := make([]map[string]any, 0)
		maxFilePathLengthSlice = append(maxFilePathLengthSlice, map[string]any{
			"max_file_path_length": rules.MaxFilePathLength.MaxFilePathLength,
		})
		rulesMap["max_file_path_length"] = maxFilePathLengthSlice
	}

	// File extension restriction rule
	if rules.FileExtensionRestriction != nil {
		fileExtensionRestrictionSlice := make([]map[string]any, 0)
		fileExtensionRestrictionSlice = append(fileExtensionRestrictionSlice, map[string]any{
			"restricted_file_extensions": rules.FileExtensionRestriction.RestrictedFileExtensions,
		})
		rulesMap["file_extension_restriction"] = fileExtensionRestrictionSlice
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
