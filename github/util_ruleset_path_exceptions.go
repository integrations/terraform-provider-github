package github

import (
	"context"
	"encoding/json"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// These wrappers preserve ignored_file_paths until go-github supports it.
// All other ruleset fields continue to use the SDK's encoding.
type rulesetWithPathExceptions struct {
	github.RepositoryRuleset
	Rules *rulesWithPathExceptions `json:"rules,omitempty"`
}

type rulesWithPathExceptions struct {
	rules            *github.RepositoryRulesetRules
	ignoredFilePaths map[string][]string
}

func (r rulesWithPathExceptions) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(r.rules)
	if err != nil {
		return nil, err
	}
	var rules []struct {
		Type       string                     `json:"type"`
		Parameters map[string]json.RawMessage `json:"parameters,omitempty"`
	}
	if err := json.Unmarshal(data, &rules); err != nil {
		return nil, err
	}
	for i := range rules {
		paths, ok := r.ignoredFilePaths[rules[i].Type]
		if !ok {
			continue
		}
		if paths == nil {
			paths = []string{}
		}
		data, err := json.Marshal(paths)
		if err != nil {
			return nil, err
		}
		rules[i].Parameters["ignored_file_paths"] = data
	}
	return json.Marshal(rules)
}

func (r *rulesWithPathExceptions) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &r.rules); err != nil {
		return err
	}
	var rules []struct {
		Type       string `json:"type"`
		Parameters struct {
			IgnoredFilePaths []string `json:"ignored_file_paths"`
		} `json:"parameters"`
	}
	if err := json.Unmarshal(data, &rules); err != nil {
		return err
	}
	r.ignoredFilePaths = make(map[string][]string)
	for _, rule := range rules {
		switch rule.Type {
		case "file_path_restriction", "max_file_size":
			r.ignoredFilePaths[rule.Type] = rule.Parameters.IgnoredFilePaths
		}
	}
	return nil
}

func expandRulesetWithPathExceptions(d *schema.ResourceData, org string) rulesetWithPathExceptions {
	ruleset := resourceGithubRulesetObject(d, org)
	rules := &rulesWithPathExceptions{
		rules:            ruleset.Rules,
		ignoredFilePaths: make(map[string][]string),
	}
	for _, ruleType := range []string{"file_path_restriction", "max_file_size"} {
		if paths, ok := d.Get("rules.0." + ruleType + ".0.ignored_file_paths").([]any); ok {
			rules.ignoredFilePaths[ruleType] = expandStringList(paths)
		}
	}
	return rulesetWithPathExceptions{RepositoryRuleset: ruleset, Rules: rules}
}

func flattenRulesWithPathExceptions(ctx context.Context, rules *rulesWithPathExceptions, org bool) []any {
	if rules == nil {
		return []any{}
	}
	result := flattenRules(ctx, rules.rules, org)
	if len(result) == 0 {
		return result
	}
	rulesMap, _ := result[0].(map[string]any)
	for ruleType, paths := range rules.ignoredFilePaths {
		if rule, ok := rulesMap[ruleType].([]map[string]any); ok && len(rule) > 0 {
			rule[0]["ignored_file_paths"] = paths
		}
	}
	return result
}

func doRulesetRequest(ctx context.Context, client *github.Client, method, path string, body any) (*rulesetWithPathExceptions, *github.Response, error) {
	req, err := client.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, nil, err
	}
	var ruleset *rulesetWithPathExceptions
	resp, err := client.Do(req, &ruleset)
	if err != nil {
		return nil, resp, err
	}
	return ruleset, resp, nil
}
