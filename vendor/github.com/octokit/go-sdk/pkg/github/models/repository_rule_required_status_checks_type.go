package models
import (
    "errors"
)
type RepositoryRuleRequiredStatusChecks_type int

const (
    REQUIRED_STATUS_CHECKS_REPOSITORYRULEREQUIREDSTATUSCHECKS_TYPE RepositoryRuleRequiredStatusChecks_type = iota
)

func (i RepositoryRuleRequiredStatusChecks_type) String() string {
    return []string{"required_status_checks"}[i]
}
func ParseRepositoryRuleRequiredStatusChecks_type(v string) (any, error) {
    result := REQUIRED_STATUS_CHECKS_REPOSITORYRULEREQUIREDSTATUSCHECKS_TYPE
    switch v {
        case "required_status_checks":
            result = REQUIRED_STATUS_CHECKS_REPOSITORYRULEREQUIREDSTATUSCHECKS_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleRequiredStatusChecks_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleRequiredStatusChecks_type(values []RepositoryRuleRequiredStatusChecks_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleRequiredStatusChecks_type) isMultiValue() bool {
    return false
}
