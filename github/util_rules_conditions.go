package github

import (
	"context"
	"reflect"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func expandConditions(input []any, org bool) *github.RepositoryRulesetConditions {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	rulesetConditions := &github.RepositoryRulesetConditions{}
	inputConditions := input[0].(map[string]any)

	// ref_name is available for both repo and org rulesets
	if v, ok := inputConditions["ref_name"].([]any); ok && v != nil && len(v) != 0 && v[0] != nil {
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
		if v, ok := inputConditions["repository_name"].([]any); ok && v != nil && len(v) != 0 && v[0] != nil {
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
		} else if v, ok := inputConditions["repository_property"].([]any); ok && v != nil && len(v) != 0 {
			rulesetConditions.RepositoryProperty = expandRepositoryPropertyConditions(v)
		}
	}

	return rulesetConditions
}

func expandRepositoryPropertyConditions(v []any) *github.RepositoryRulesetRepositoryPropertyConditionParameters {
	repositoryProperties := v[0].(map[string]any)
	include := make([]*github.RepositoryRulesetRepositoryPropertyTargetParameters, 0)
	exclude := make([]*github.RepositoryRulesetRepositoryPropertyTargetParameters, 0)

	for _, v := range repositoryProperties["include"].([]any) {
		if v != nil {
			propertyMap := v.(map[string]any)
			propertyValues := make([]string, 0)
			for _, pv := range propertyMap["property_values"].([]any) {
				if pv != nil {
					propertyValues = append(propertyValues, pv.(string))
				}
			}
			property := &github.RepositoryRulesetRepositoryPropertyTargetParameters{
				Name:           propertyMap["name"].(string),
				Source:         new(propertyMap["source"].(string)),
				PropertyValues: propertyValues,
			}
			include = append(include, property)
		}
	}

	for _, v := range repositoryProperties["exclude"].([]any) {
		if v != nil {
			propertyMap := v.(map[string]any)
			propertyValues := make([]string, 0)
			for _, pv := range propertyMap["property_values"].([]any) {
				if pv != nil {
					propertyValues = append(propertyValues, pv.(string))
				}
			}
			property := &github.RepositoryRulesetRepositoryPropertyTargetParameters{
				Name:           propertyMap["name"].(string),
				Source:         new(propertyMap["source"].(string)),
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
