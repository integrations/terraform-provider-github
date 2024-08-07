package models
import (
    "errors"
)
// The result of the rule evaluations for rules with the `active` and `evaluate` enforcement statuses, demonstrating whether rules would pass or fail if all rules in the rule suite were `active`.
type RuleSuites_evaluation_result int

const (
    PASS_RULESUITES_EVALUATION_RESULT RuleSuites_evaluation_result = iota
    FAIL_RULESUITES_EVALUATION_RESULT
)

func (i RuleSuites_evaluation_result) String() string {
    return []string{"pass", "fail"}[i]
}
func ParseRuleSuites_evaluation_result(v string) (any, error) {
    result := PASS_RULESUITES_EVALUATION_RESULT
    switch v {
        case "pass":
            result = PASS_RULESUITES_EVALUATION_RESULT
        case "fail":
            result = FAIL_RULESUITES_EVALUATION_RESULT
        default:
            return 0, errors.New("Unknown RuleSuites_evaluation_result value: " + v)
    }
    return &result, nil
}
func SerializeRuleSuites_evaluation_result(values []RuleSuites_evaluation_result) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RuleSuites_evaluation_result) isMultiValue() bool {
    return false
}
