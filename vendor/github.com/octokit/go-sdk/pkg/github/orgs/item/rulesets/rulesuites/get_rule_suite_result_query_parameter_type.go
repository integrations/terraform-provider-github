package rulesuites
import (
    "errors"
)
type GetRule_suite_resultQueryParameterType int

const (
    PASS_GETRULE_SUITE_RESULTQUERYPARAMETERTYPE GetRule_suite_resultQueryParameterType = iota
    FAIL_GETRULE_SUITE_RESULTQUERYPARAMETERTYPE
    BYPASS_GETRULE_SUITE_RESULTQUERYPARAMETERTYPE
    ALL_GETRULE_SUITE_RESULTQUERYPARAMETERTYPE
)

func (i GetRule_suite_resultQueryParameterType) String() string {
    return []string{"pass", "fail", "bypass", "all"}[i]
}
func ParseGetRule_suite_resultQueryParameterType(v string) (any, error) {
    result := PASS_GETRULE_SUITE_RESULTQUERYPARAMETERTYPE
    switch v {
        case "pass":
            result = PASS_GETRULE_SUITE_RESULTQUERYPARAMETERTYPE
        case "fail":
            result = FAIL_GETRULE_SUITE_RESULTQUERYPARAMETERTYPE
        case "bypass":
            result = BYPASS_GETRULE_SUITE_RESULTQUERYPARAMETERTYPE
        case "all":
            result = ALL_GETRULE_SUITE_RESULTQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetRule_suite_resultQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetRule_suite_resultQueryParameterType(values []GetRule_suite_resultQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetRule_suite_resultQueryParameterType) isMultiValue() bool {
    return false
}
