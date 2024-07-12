package models
import (
    "errors"
)
type RepositoryRuleNonFastForward_type int

const (
    NON_FAST_FORWARD_REPOSITORYRULENONFASTFORWARD_TYPE RepositoryRuleNonFastForward_type = iota
)

func (i RepositoryRuleNonFastForward_type) String() string {
    return []string{"non_fast_forward"}[i]
}
func ParseRepositoryRuleNonFastForward_type(v string) (any, error) {
    result := NON_FAST_FORWARD_REPOSITORYRULENONFASTFORWARD_TYPE
    switch v {
        case "non_fast_forward":
            result = NON_FAST_FORWARD_REPOSITORYRULENONFASTFORWARD_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleNonFastForward_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleNonFastForward_type(values []RepositoryRuleNonFastForward_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleNonFastForward_type) isMultiValue() bool {
    return false
}
