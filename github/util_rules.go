package github

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func toPullRequestMergeMethods(input any) []github.PullRequestMergeMethod {
	value, ok := input.([]any)
	if !ok || len(value) == 0 {
		return []github.PullRequestMergeMethod{}
	}

	mergeMethods := make([]github.PullRequestMergeMethod, 0, len(value))
	for _, item := range value {
		if method, ok := item.(string); ok {
			mergeMethods = append(mergeMethods, github.PullRequestMergeMethod(method))
		}
	}
	return mergeMethods
}

// expandRequiredReviewers converts Terraform schema data to go-github RequiredReviewers.
func expandRequiredReviewers(input []any) []*github.RulesetRequiredReviewer {
	if len(input) == 0 {
		return nil
	}

	reviewers := make([]*github.RulesetRequiredReviewer, 0, len(input))
	for _, item := range input {
		reviewerMap, ok := item.(map[string]any)
		if !ok {
			continue
		}

		var reviewer *github.RulesetReviewer
		if rv, ok := reviewerMap["reviewer"].([]any); ok && len(rv) != 0 {
			if reviewerData, ok := rv[0].(map[string]any); ok {
				reviewerTypeStr, _ := reviewerData["type"].(string)
				reviewerType := github.RulesetReviewerType(reviewerTypeStr)
				reviewerID, _ := reviewerData["id"].(int)
				reviewer = &github.RulesetReviewer{
					ID:   new(int64(reviewerID)),
					Type: &reviewerType,
				}
			}
		}

		filePatterns := make([]string, 0)
		if fp, ok := reviewerMap["file_patterns"].([]any); ok {
			for _, p := range fp {
				if s, ok := p.(string); ok {
					filePatterns = append(filePatterns, s)
				}
			}
		}

		minApprovals, _ := reviewerMap["minimum_approvals"].(int)
		reviewers = append(reviewers, &github.RulesetRequiredReviewer{
			MinimumApprovals: new(minApprovals),
			FilePatterns:     filePatterns,
			Reviewer:         reviewer,
		})
	}
	return reviewers
}

// flattenRequiredReviewers converts go-github RequiredReviewers to Terraform schema data.
func flattenRequiredReviewers(reviewers []*github.RulesetRequiredReviewer) []map[string]any {
	if len(reviewers) == 0 {
		return nil
	}

	reviewersList := make([]map[string]any, 0, len(reviewers))
	for _, rr := range reviewers {
		reviewerMap := map[string]any{
			"file_patterns":     rr.FilePatterns,
			"minimum_approvals": 0,
		}
		if rr.MinimumApprovals != nil {
			reviewerMap["minimum_approvals"] = *rr.MinimumApprovals
		}
		if rr.Reviewer != nil {
			reviewerData := map[string]any{}
			if rr.Reviewer.ID != nil {
				reviewerData["id"] = int(*rr.Reviewer.ID)
			}
			if rr.Reviewer.Type != nil {
				reviewerData["type"] = string(*rr.Reviewer.Type)
			}
			reviewerMap["reviewer"] = []map[string]any{reviewerData}
		}
		reviewersList = append(reviewersList, reviewerMap)
	}
	return reviewersList
}

func resourceGithubRulesetObject(d *schema.ResourceData, org string) github.RepositoryRuleset {
	isOrgLevel := len(org) > 0

	var source, sourceType string
	if isOrgLevel {
		source = org
		sourceType = "Organization"
	} else {
		source, _ = d.Get("repository").(string)
		sourceType = "Repository"
	}

	targetStr, _ := d.Get("target").(string)
	enforcementStr, _ := d.Get("enforcement").(string)
	name, _ := d.Get("name").(string)
	bypassActorsRaw, _ := d.Get("bypass_actors").([]any)
	conditionsRaw, _ := d.Get("conditions").([]any)
	rulesRaw, _ := d.Get("rules").([]any)

	target := github.RulesetTarget(targetStr)
	enforcement := github.RulesetEnforcement(enforcementStr)
	sourceTypeEnum := github.RulesetSourceType(sourceType)

	return github.RepositoryRuleset{
		Name:         name,
		Target:       &target,
		Source:       source,
		SourceType:   &sourceTypeEnum,
		Enforcement:  enforcement,
		BypassActors: expandBypassActors(bypassActorsRaw),
		Conditions:   expandConditions(conditionsRaw, isOrgLevel),
		Rules:        expandRules(rulesRaw, isOrgLevel),
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
		inputMap, ok := v.(map[string]any)
		if !ok {
			continue
		}
		actor := &github.BypassActor{}
		if actorIDVal, ok := inputMap["actor_id"]; ok {
			actorID := toInt(actorIDVal)
			if actorID == 0 {
				actor.ActorID = nil
			} else {
				actor.ActorID = new(int64(actorID))
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

func flattenBypassActors(ctx context.Context, bypassActors []*github.BypassActor) []any {
	tflog.Debug(ctx, "Flattening bypass actors", map[string]any{"count_bypass_actors": len(bypassActors)})

	if bypassActors == nil {
		return []any{}
	}

	actorsSlice := make([]any, 0)
	for _, v := range bypassActors {
		actorMap := make(map[string]any)
		actorID := v.GetActorID()
		actorType := v.GetActorType()
		actorMap["actor_id"] = actorID
		actorMap["actor_type"] = actorType
		actorMap["bypass_mode"] = v.GetBypassMode()

		actorsSlice = append(actorsSlice, actorMap)
	}

	tflog.Debug(ctx, "Flattened bypass actors", map[string]any{"count_bypass_actors": len(bypassActors), "flattened_bypass_actors": actorsSlice})

	return actorsSlice
}

func expandConditions(input []any, org bool) *github.RepositoryRulesetConditions {
	if len(input) == 0 {
		return nil
	}
	inputConditions, ok := input[0].(map[string]any)
	if !ok {
		return nil
	}
	rulesetConditions := &github.RepositoryRulesetConditions{}

	// ref_name is available for both repo and org rulesets
	if v, ok := inputConditions["ref_name"].([]any); ok && len(v) != 0 {
		if inputRefName, ok := v[0].(map[string]any); ok {
			include := make([]string, 0)
			exclude := make([]string, 0)

			if inc, ok := inputRefName["include"].([]any); ok {
				for _, v := range inc {
					if s, ok := v.(string); ok {
						include = append(include, s)
					}
				}
			}

			if exc, ok := inputRefName["exclude"].([]any); ok {
				for _, v := range exc {
					if s, ok := v.(string); ok {
						exclude = append(exclude, s)
					}
				}
			}

			rulesetConditions.RefName = &github.RepositoryRulesetRefConditionParameters{
				Include: include,
				Exclude: exclude,
			}
		}
	}

	// org-only fields
	if org {
		// repository_name and repository_id
		if v, ok := inputConditions["repository_name"].([]any); ok && len(v) != 0 {
			if inputRepositoryName, ok := v[0].(map[string]any); ok {
				include := make([]string, 0)
				exclude := make([]string, 0)

				if inc, ok := inputRepositoryName["include"].([]any); ok {
					for _, v := range inc {
						if s, ok := v.(string); ok {
							include = append(include, s)
						}
					}
				}

				if exc, ok := inputRepositoryName["exclude"].([]any); ok {
					for _, v := range exc {
						if s, ok := v.(string); ok {
							exclude = append(exclude, s)
						}
					}
				}

				protected, _ := inputRepositoryName["protected"].(bool)

				rulesetConditions.RepositoryName = &github.RepositoryRulesetRepositoryNamesConditionParameters{
					Include:   include,
					Exclude:   exclude,
					Protected: &protected,
				}
			}
		} else if v, ok := inputConditions["repository_id"].([]any); ok && len(v) != 0 {
			repositoryIDs := make([]int64, 0)

			for _, v := range v {
				if v != nil {
					repositoryIDs = append(repositoryIDs, toInt64(v))
				}
			}

			rulesetConditions.RepositoryID = &github.RepositoryRulesetRepositoryIDsConditionParameters{RepositoryIDs: repositoryIDs}
		} else if v, ok := inputConditions["repository_property"].([]any); ok && len(v) != 0 {
			rulesetConditions.RepositoryProperty = expandRepositoryPropertyConditions(v)
		}
	}

	return rulesetConditions
}

func expandRepositoryPropertyConditions(v []any) *github.RepositoryRulesetRepositoryPropertyConditionParameters {
	if len(v) == 0 {
		return nil
	}
	repositoryProperties, ok := v[0].(map[string]any)
	if !ok {
		return nil
	}
	include := make([]*github.RepositoryRulesetRepositoryPropertyTargetParameters, 0)
	exclude := make([]*github.RepositoryRulesetRepositoryPropertyTargetParameters, 0)

	if inc, ok := repositoryProperties["include"].([]any); ok {
		for _, v := range inc {
			propertyMap, ok := v.(map[string]any)
			if !ok {
				continue
			}
			propertyValues := make([]string, 0)
			if pvs, ok := propertyMap["property_values"].([]any); ok {
				for _, pv := range pvs {
					if s, ok := pv.(string); ok {
						propertyValues = append(propertyValues, s)
					}
				}
			}
			name, _ := propertyMap["name"].(string)
			source, _ := propertyMap["source"].(string)
			property := &github.RepositoryRulesetRepositoryPropertyTargetParameters{
				Name:           name,
				Source:         new(source),
				PropertyValues: propertyValues,
			}
			include = append(include, property)
		}
	}

	if exc, ok := repositoryProperties["exclude"].([]any); ok {
		for _, v := range exc {
			propertyMap, ok := v.(map[string]any)
			if !ok {
				continue
			}
			propertyValues := make([]string, 0)
			if pvs, ok := propertyMap["property_values"].([]any); ok {
				for _, pv := range pvs {
					if s, ok := pv.(string); ok {
						propertyValues = append(propertyValues, s)
					}
				}
			}
			name, _ := propertyMap["name"].(string)
			source, _ := propertyMap["source"].(string)
			property := &github.RepositoryRulesetRepositoryPropertyTargetParameters{
				Name:           name,
				Source:         new(source),
				PropertyValues: propertyValues,
			}
			exclude = append(exclude, property)
		}
	}

	return &github.RepositoryRulesetRepositoryPropertyConditionParameters{
		Include: include,
		Exclude: exclude,
	}
}

func flattenConditions(ctx context.Context, conditions *github.RepositoryRulesetConditions, org bool) []any {
	if conditions == nil || reflect.DeepEqual(conditions, &github.RepositoryRulesetConditions{}) {
		tflog.Debug(ctx, "Conditions are empty, returning empty list")
		return []any{}
	}

	conditionsMap := make(map[string]any)
	refNameSlice := make([]map[string]any, 0)

	if conditions.RefName != nil {
		refNameSlice = append(refNameSlice, map[string]any{
			"include": conditions.RefName.Include,
			"exclude": conditions.RefName.Exclude,
		})

		conditionsMap["ref_name"] = refNameSlice
	}

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

		if conditions.RepositoryProperty != nil {
			repositoryPropertySlice := make([]map[string]any, 0)

			repositoryPropertySlice = append(repositoryPropertySlice, map[string]any{
				"include": flattenRulesetRepositoryPropertyTargetParameters(conditions.RepositoryProperty.Include),
				"exclude": flattenRulesetRepositoryPropertyTargetParameters(conditions.RepositoryProperty.Exclude),
			})
			conditionsMap["repository_property"] = repositoryPropertySlice
		}
	}

	return []any{conditionsMap}
}

func flattenRulesetRepositoryPropertyTargetParameters(input []*github.RepositoryRulesetRepositoryPropertyTargetParameters) []map[string]any {
	result := make([]map[string]any, 0)

	for _, v := range input {
		propertyMap := make(map[string]any)
		propertyMap["name"] = v.Name
		source := v.GetSource()
		if source == "" {
			source = "custom"
		}
		propertyMap["source"] = source
		propertyMap["property_values"] = v.PropertyValues
		result = append(result, propertyMap)
	}

	return result
}

func expandRules(input []any, org bool) *github.RepositoryRulesetRules {
	if len(input) == 0 {
		return &github.RepositoryRulesetRules{}
	}

	rulesMap, ok := input[0].(map[string]any)
	if !ok {
		return &github.RepositoryRulesetRules{}
	}
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
		if requiredDeploymentsMap, ok := v[0].(map[string]any); ok {
			envs := make([]string, 0)
			if envList, ok := requiredDeploymentsMap["required_deployment_environments"].([]any); ok {
				for _, env := range envList {
					if envStr, ok := env.(string); ok {
						envs = append(envs, envStr)
					}
				}
			}
			rulesetRules.RequiredDeployments = &github.RequiredDeploymentsRuleParameters{
				RequiredDeploymentEnvironments: envs,
			}
		}
	}

	// Pull request rule
	if v, ok := rulesMap["pull_request"].([]any); ok && len(v) != 0 {
		if pullRequestMap, ok := v[0].(map[string]any); ok {
			allowedMergeMethods := pullRequestMap["allowed_merge_methods"]
			dismissStale, _ := pullRequestMap["dismiss_stale_reviews_on_push"].(bool)
			requireCodeOwner, _ := pullRequestMap["require_code_owner_review"].(bool)
			requireLastPush, _ := pullRequestMap["require_last_push_approval"].(bool)
			requireThreadResolution, _ := pullRequestMap["required_review_thread_resolution"].(bool)

			params := &github.PullRequestRuleParameters{
				DismissStaleReviewsOnPush:      dismissStale,
				RequireCodeOwnerReview:         requireCodeOwner,
				RequireLastPushApproval:        requireLastPush,
				RequiredApprovingReviewCount:   toInt(pullRequestMap["required_approving_review_count"]),
				RequiredReviewThreadResolution: requireThreadResolution,
				AllowedMergeMethods:            toPullRequestMergeMethods(allowedMergeMethods),
			}

			// Add required reviewers if provided
			if reqReviewers, ok := pullRequestMap["required_reviewers"].([]any); ok && len(reqReviewers) != 0 {
				params.RequiredReviewers = expandRequiredReviewers(reqReviewers)
			}

			rulesetRules.PullRequest = params
		}
	}

	// Merge queue rule
	if v, ok := rulesMap["merge_queue"].([]any); ok && len(v) != 0 {
		if mergeQueueMap, ok := v[0].(map[string]any); ok {
			groupingStrategy, _ := mergeQueueMap["grouping_strategy"].(string)
			mergeMethod, _ := mergeQueueMap["merge_method"].(string)
			params := &github.MergeQueueRuleParameters{
				CheckResponseTimeoutMinutes:  toInt(mergeQueueMap["check_response_timeout_minutes"]),
				GroupingStrategy:             github.MergeGroupingStrategy(groupingStrategy),
				MaxEntriesToBuild:            toInt(mergeQueueMap["max_entries_to_build"]),
				MaxEntriesToMerge:            toInt(mergeQueueMap["max_entries_to_merge"]),
				MergeMethod:                  github.MergeQueueMergeMethod(mergeMethod),
				MinEntriesToMerge:            toInt(mergeQueueMap["min_entries_to_merge"]),
				MinEntriesToMergeWaitMinutes: toInt(mergeQueueMap["min_entries_to_merge_wait_minutes"]),
			}
			rulesetRules.MergeQueue = params
		}
	}

	// Required status checks rule
	if v, ok := rulesMap["required_status_checks"].([]any); ok && len(v) != 0 {
		if requiredStatusMap, ok := v[0].(map[string]any); ok {
			requiredStatusChecks := make([]*github.RuleStatusCheck, 0)

			if requiredStatusChecksSet, ok := requiredStatusMap["required_check"].(*schema.Set); ok {
				for _, checkMap := range requiredStatusChecksSet.List() {
					check, ok := checkMap.(map[string]any)
					if !ok {
						continue
					}
					integrationID := toInt64(check["integration_id"])
					contextStr, _ := check["context"].(string)

					params := &github.RuleStatusCheck{
						Context: contextStr,
					}

					if integrationID != 0 {
						params.IntegrationID = new(integrationID)
					}

					requiredStatusChecks = append(requiredStatusChecks, params)
				}
			}

			doNotEnforceOnCreate, _ := requiredStatusMap["do_not_enforce_on_create"].(bool)
			strictPolicy, _ := requiredStatusMap["strict_required_status_checks_policy"].(bool)
			params := &github.RequiredStatusChecksRuleParameters{
				RequiredStatusChecks:             requiredStatusChecks,
				StrictRequiredStatusChecksPolicy: strictPolicy,
				DoNotEnforceOnCreate:             &doNotEnforceOnCreate,
			}
			rulesetRules.RequiredStatusChecks = params
		}
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
			patternParametersMap, ok := v[0].(map[string]any)
			if !ok {
				continue
			}

			name, _ := patternParametersMap["name"].(string)
			negate, _ := patternParametersMap["negate"].(bool)
			operator, _ := patternParametersMap["operator"].(string)
			pattern, _ := patternParametersMap["pattern"].(string)

			params := &github.PatternRuleParameters{
				Name:     &name,
				Negate:   &negate,
				Operator: github.PatternRuleOperator(operator),
				Pattern:  pattern,
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
			if requiredWorkflowsMap, ok := v[0].(map[string]any); ok {
				requiredWorkflows := make([]*github.RuleWorkflow, 0)

				if requiredWorkflowsSet, ok := requiredWorkflowsMap["required_workflow"].(*schema.Set); ok {
					for _, workflowMap := range requiredWorkflowsSet.List() {
						workflow, ok := workflowMap.(map[string]any)
						if !ok {
							continue
						}
						path, _ := workflow["path"].(string)
						ref, _ := workflow["ref"].(string)

						params := &github.RuleWorkflow{
							RepositoryID: new(toInt64(workflow["repository_id"])),
							Path:         path,
							Ref:          new(ref),
						}

						requiredWorkflows = append(requiredWorkflows, params)
					}
				}

				doNotEnforceOnCreate, _ := requiredWorkflowsMap["do_not_enforce_on_create"].(bool)
				params := &github.WorkflowsRuleParameters{
					DoNotEnforceOnCreate: &doNotEnforceOnCreate,
					Workflows:            requiredWorkflows,
				}
				rulesetRules.Workflows = params
			}
		}
	}

	// Required code scanning rule
	if v, ok := rulesMap["required_code_scanning"].([]any); ok && len(v) != 0 {
		if requiredCodeScanningMap, ok := v[0].(map[string]any); ok {
			requiredCodeScanningTools := make([]*github.RuleCodeScanningTool, 0)

			if requiredCodeScanningSet, ok := requiredCodeScanningMap["required_code_scanning_tool"].(*schema.Set); ok {
				for _, codeScanningMap := range requiredCodeScanningSet.List() {
					codeScanningTool, ok := codeScanningMap.(map[string]any)
					if !ok {
						continue
					}
					alertsThreshold, _ := codeScanningTool["alerts_threshold"].(string)
					securityAlertsThreshold, _ := codeScanningTool["security_alerts_threshold"].(string)
					tool, _ := codeScanningTool["tool"].(string)

					params := &github.RuleCodeScanningTool{
						AlertsThreshold:         github.CodeScanningAlertsThreshold(alertsThreshold),
						SecurityAlertsThreshold: github.CodeScanningSecurityAlertsThreshold(securityAlertsThreshold),
						Tool:                    tool,
					}

					requiredCodeScanningTools = append(requiredCodeScanningTools, params)
				}
			}

			params := &github.CodeScanningRuleParameters{
				CodeScanningTools: requiredCodeScanningTools,
			}
			rulesetRules.CodeScanning = params
		}
	}

	// File path restriction rule
	if v, ok := rulesMap["file_path_restriction"].([]any); ok && len(v) != 0 {
		if filePathRestrictionMap, ok := v[0].(map[string]any); ok {
			restrictedFilePaths := make([]string, 0)
			if paths, ok := filePathRestrictionMap["restricted_file_paths"].([]any); ok {
				for _, path := range paths {
					if pathStr, ok := path.(string); ok {
						restrictedFilePaths = append(restrictedFilePaths, pathStr)
					}
				}
			}
			params := &github.FilePathRestrictionRuleParameters{
				RestrictedFilePaths: restrictedFilePaths,
			}
			rulesetRules.FilePathRestriction = params
		}
	}

	// Max file size rule
	if v, ok := rulesMap["max_file_size"].([]any); ok && len(v) != 0 {
		if maxFileSizeMap, ok := v[0].(map[string]any); ok {
			maxFileSize := toInt64(maxFileSizeMap["max_file_size"])
			params := &github.MaxFileSizeRuleParameters{
				MaxFileSize: maxFileSize,
			}
			rulesetRules.MaxFileSize = params
		}
	}

	// Max file path length rule
	if v, ok := rulesMap["max_file_path_length"].([]any); ok && len(v) != 0 {
		if maxFilePathLengthMap, ok := v[0].(map[string]any); ok {
			maxFilePathLength := toInt(maxFilePathLengthMap["max_file_path_length"])
			params := &github.MaxFilePathLengthRuleParameters{
				MaxFilePathLength: maxFilePathLength,
			}
			rulesetRules.MaxFilePathLength = params
		}
	}

	// File extension restriction rule
	if v, ok := rulesMap["file_extension_restriction"].([]any); ok && len(v) != 0 {
		if fileExtensionRestrictionMap, ok := v[0].(map[string]any); ok {
			restrictedFileExtensions := make([]string, 0)

			// Handle schema.TypeSet
			if extensionSet, ok := fileExtensionRestrictionMap["restricted_file_extensions"].(*schema.Set); ok {
				for _, extension := range extensionSet.List() {
					if extStr, ok := extension.(string); ok {
						restrictedFileExtensions = append(restrictedFileExtensions, extStr)
					}
				}
			}
			params := &github.FileExtensionRestrictionRuleParameters{
				RestrictedFileExtensions: restrictedFileExtensions,
			}
			rulesetRules.FileExtensionRestriction = params
		}
	}

	// Copilot code review rule
	if v, ok := rulesMap["copilot_code_review"].([]any); ok && len(v) != 0 {
		if copilotCodeReviewMap, ok := v[0].(map[string]any); ok {
			reviewOnPush, _ := copilotCodeReviewMap["review_on_push"].(bool)
			reviewDraftPullRequests, _ := copilotCodeReviewMap["review_draft_pull_requests"].(bool)
			params := &github.CopilotCodeReviewRuleParameters{
				ReviewOnPush:            reviewOnPush,
				ReviewDraftPullRequests: reviewDraftPullRequests,
			}
			rulesetRules.CopilotCodeReview = params
		}
	}

	return rulesetRules
}

func flattenRules(ctx context.Context, rules *github.RepositoryRulesetRules, org bool) []any {
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
		if !org {
			rulesMap["update_allows_fetch_and_merge"] = rules.Update.UpdateAllowsFetchAndMerge
		}
	} else {
		rulesMap["update"] = false
		if !org {
			rulesMap["update_allows_fetch_and_merge"] = false
		}
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
			"allowed_merge_methods":             rules.PullRequest.AllowedMergeMethods,
			"required_reviewers":                flattenRequiredReviewers(rules.PullRequest.RequiredReviewers),
		})
		tflog.Debug(ctx, "Flattened Pull Request rules slice", map[string]any{"pull_request": pullRequestSlice})
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

	// Copilot code review rule
	if rules.CopilotCodeReview != nil {
		copilotCodeReviewSlice := make([]map[string]any, 0)
		copilotCodeReviewSlice = append(copilotCodeReviewSlice, map[string]any{
			"review_on_push":             rules.CopilotCodeReview.ReviewOnPush,
			"review_draft_pull_requests": rules.CopilotCodeReview.ReviewDraftPullRequests,
		})
		rulesMap["copilot_code_review"] = copilotCodeReviewSlice
	}

	return []any{rulesMap}
}

// bypassActorCompareIdentity compares two bypass actors based on their actor_type and actor_id.
func bypassActorCompareIdentity(a, b any) int {
	mapA, ok := a.(map[string]any)
	if !ok {
		return 0
	}
	mapB, ok := b.(map[string]any)
	if !ok {
		return 0
	}

	typeA, ok := mapA["actor_type"].(string)
	if !ok {
		return 0
	}
	typeB, ok := mapB["actor_type"].(string)
	if !ok {
		return 0
	}

	idA, ok := mapA["actor_id"].(int)
	if !ok {
		return 0
	}
	idB, ok := mapB["actor_id"].(int)
	if !ok {
		return 0
	}

	return strings.Compare(fmt.Sprintf("%s:%d", typeA, idA), fmt.Sprintf("%s:%d", typeB, idB))
}
