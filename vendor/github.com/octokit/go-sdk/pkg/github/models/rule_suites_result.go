package models
import (
    "errors"
)
// The result of the rule evaluations for rules with the `active` enforcement status.
type RuleSuites_result int

const (
    PASS_RULESUITES_RESULT RuleSuites_result = iota
    FAIL_RULESUITES_RESULT
    BYPASS_RULESUITES_RESULT
)

func (i RuleSuites_result) String() string {
    return []string{"pass", "fail", "bypass"}[i]
}
func ParseRuleSuites_result(v string) (any, error) {
    result := PASS_RULESUITES_RESULT
    switch v {
        case "pass":
            result = PASS_RULESUITES_RESULT
        case "fail":
            result = FAIL_RULESUITES_RESULT
        case "bypass":
            result = BYPASS_RULESUITES_RESULT
        default:
            return 0, errors.New("Unknown RuleSuites_result value: " + v)
    }
    return &result, nil
}
func SerializeRuleSuites_result(values []RuleSuites_result) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RuleSuites_result) isMultiValue() bool {
    return false
}
