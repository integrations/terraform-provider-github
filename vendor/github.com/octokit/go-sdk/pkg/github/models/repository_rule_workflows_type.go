package models
import (
    "errors"
)
type RepositoryRuleWorkflows_type int

const (
    WORKFLOWS_REPOSITORYRULEWORKFLOWS_TYPE RepositoryRuleWorkflows_type = iota
)

func (i RepositoryRuleWorkflows_type) String() string {
    return []string{"workflows"}[i]
}
func ParseRepositoryRuleWorkflows_type(v string) (any, error) {
    result := WORKFLOWS_REPOSITORYRULEWORKFLOWS_TYPE
    switch v {
        case "workflows":
            result = WORKFLOWS_REPOSITORYRULEWORKFLOWS_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleWorkflows_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleWorkflows_type(values []RepositoryRuleWorkflows_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleWorkflows_type) isMultiValue() bool {
    return false
}
