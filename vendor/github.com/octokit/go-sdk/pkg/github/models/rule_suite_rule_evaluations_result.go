package models
import (
    "errors"
)
// The result of the evaluation of the individual rule.
type RuleSuite_rule_evaluations_result int

const (
    PASS_RULESUITE_RULE_EVALUATIONS_RESULT RuleSuite_rule_evaluations_result = iota
    FAIL_RULESUITE_RULE_EVALUATIONS_RESULT
)

func (i RuleSuite_rule_evaluations_result) String() string {
    return []string{"pass", "fail"}[i]
}
func ParseRuleSuite_rule_evaluations_result(v string) (any, error) {
    result := PASS_RULESUITE_RULE_EVALUATIONS_RESULT
    switch v {
        case "pass":
            result = PASS_RULESUITE_RULE_EVALUATIONS_RESULT
        case "fail":
            result = FAIL_RULESUITE_RULE_EVALUATIONS_RESULT
        default:
            return 0, errors.New("Unknown RuleSuite_rule_evaluations_result value: " + v)
    }
    return &result, nil
}
func SerializeRuleSuite_rule_evaluations_result(values []RuleSuite_rule_evaluations_result) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RuleSuite_rule_evaluations_result) isMultiValue() bool {
    return false
}
