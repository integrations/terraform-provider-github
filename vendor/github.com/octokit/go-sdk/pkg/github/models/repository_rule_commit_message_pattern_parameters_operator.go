package models
import (
    "errors"
)
// The operator to use for matching.
type RepositoryRuleCommitMessagePattern_parameters_operator int

const (
    STARTS_WITH_REPOSITORYRULECOMMITMESSAGEPATTERN_PARAMETERS_OPERATOR RepositoryRuleCommitMessagePattern_parameters_operator = iota
    ENDS_WITH_REPOSITORYRULECOMMITMESSAGEPATTERN_PARAMETERS_OPERATOR
    CONTAINS_REPOSITORYRULECOMMITMESSAGEPATTERN_PARAMETERS_OPERATOR
    REGEX_REPOSITORYRULECOMMITMESSAGEPATTERN_PARAMETERS_OPERATOR
)

func (i RepositoryRuleCommitMessagePattern_parameters_operator) String() string {
    return []string{"starts_with", "ends_with", "contains", "regex"}[i]
}
func ParseRepositoryRuleCommitMessagePattern_parameters_operator(v string) (any, error) {
    result := STARTS_WITH_REPOSITORYRULECOMMITMESSAGEPATTERN_PARAMETERS_OPERATOR
    switch v {
        case "starts_with":
            result = STARTS_WITH_REPOSITORYRULECOMMITMESSAGEPATTERN_PARAMETERS_OPERATOR
        case "ends_with":
            result = ENDS_WITH_REPOSITORYRULECOMMITMESSAGEPATTERN_PARAMETERS_OPERATOR
        case "contains":
            result = CONTAINS_REPOSITORYRULECOMMITMESSAGEPATTERN_PARAMETERS_OPERATOR
        case "regex":
            result = REGEX_REPOSITORYRULECOMMITMESSAGEPATTERN_PARAMETERS_OPERATOR
        default:
            return 0, errors.New("Unknown RepositoryRuleCommitMessagePattern_parameters_operator value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleCommitMessagePattern_parameters_operator(values []RepositoryRuleCommitMessagePattern_parameters_operator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleCommitMessagePattern_parameters_operator) isMultiValue() bool {
    return false
}
