package models
import (
    "errors"
)
type RepositoryRuleRequiredLinearHistory_type int

const (
    REQUIRED_LINEAR_HISTORY_REPOSITORYRULEREQUIREDLINEARHISTORY_TYPE RepositoryRuleRequiredLinearHistory_type = iota
)

func (i RepositoryRuleRequiredLinearHistory_type) String() string {
    return []string{"required_linear_history"}[i]
}
func ParseRepositoryRuleRequiredLinearHistory_type(v string) (any, error) {
    result := REQUIRED_LINEAR_HISTORY_REPOSITORYRULEREQUIREDLINEARHISTORY_TYPE
    switch v {
        case "required_linear_history":
            result = REQUIRED_LINEAR_HISTORY_REPOSITORYRULEREQUIREDLINEARHISTORY_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleRequiredLinearHistory_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleRequiredLinearHistory_type(values []RepositoryRuleRequiredLinearHistory_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleRequiredLinearHistory_type) isMultiValue() bool {
    return false
}
