package models
import (
    "errors"
)
// The enforcement level of this rule source.
type RuleSuite_rule_evaluations_enforcement int

const (
    ACTIVE_RULESUITE_RULE_EVALUATIONS_ENFORCEMENT RuleSuite_rule_evaluations_enforcement = iota
    EVALUATE_RULESUITE_RULE_EVALUATIONS_ENFORCEMENT
    DELETEDRULESET_RULESUITE_RULE_EVALUATIONS_ENFORCEMENT
)

func (i RuleSuite_rule_evaluations_enforcement) String() string {
    return []string{"active", "evaluate", "deleted ruleset"}[i]
}
func ParseRuleSuite_rule_evaluations_enforcement(v string) (any, error) {
    result := ACTIVE_RULESUITE_RULE_EVALUATIONS_ENFORCEMENT
    switch v {
        case "active":
            result = ACTIVE_RULESUITE_RULE_EVALUATIONS_ENFORCEMENT
        case "evaluate":
            result = EVALUATE_RULESUITE_RULE_EVALUATIONS_ENFORCEMENT
        case "deleted ruleset":
            result = DELETEDRULESET_RULESUITE_RULE_EVALUATIONS_ENFORCEMENT
        default:
            return 0, errors.New("Unknown RuleSuite_rule_evaluations_enforcement value: " + v)
    }
    return &result, nil
}
func SerializeRuleSuite_rule_evaluations_enforcement(values []RuleSuite_rule_evaluations_enforcement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RuleSuite_rule_evaluations_enforcement) isMultiValue() bool {
    return false
}
