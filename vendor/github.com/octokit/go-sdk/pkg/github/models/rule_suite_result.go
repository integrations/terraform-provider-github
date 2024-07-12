package models
import (
    "errors"
)
// The result of the rule evaluations for rules with the `active` enforcement status.
type RuleSuite_result int

const (
    PASS_RULESUITE_RESULT RuleSuite_result = iota
    FAIL_RULESUITE_RESULT
    BYPASS_RULESUITE_RESULT
)

func (i RuleSuite_result) String() string {
    return []string{"pass", "fail", "bypass"}[i]
}
func ParseRuleSuite_result(v string) (any, error) {
    result := PASS_RULESUITE_RESULT
    switch v {
        case "pass":
            result = PASS_RULESUITE_RESULT
        case "fail":
            result = FAIL_RULESUITE_RESULT
        case "bypass":
            result = BYPASS_RULESUITE_RESULT
        default:
            return 0, errors.New("Unknown RuleSuite_result value: " + v)
    }
    return &result, nil
}
func SerializeRuleSuite_result(values []RuleSuite_result) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RuleSuite_result) isMultiValue() bool {
    return false
}
