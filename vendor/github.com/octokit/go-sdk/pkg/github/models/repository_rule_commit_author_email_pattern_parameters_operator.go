package models
import (
    "errors"
)
// The operator to use for matching.
type RepositoryRuleCommitAuthorEmailPattern_parameters_operator int

const (
    STARTS_WITH_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_PARAMETERS_OPERATOR RepositoryRuleCommitAuthorEmailPattern_parameters_operator = iota
    ENDS_WITH_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_PARAMETERS_OPERATOR
    CONTAINS_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_PARAMETERS_OPERATOR
    REGEX_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_PARAMETERS_OPERATOR
)

func (i RepositoryRuleCommitAuthorEmailPattern_parameters_operator) String() string {
    return []string{"starts_with", "ends_with", "contains", "regex"}[i]
}
func ParseRepositoryRuleCommitAuthorEmailPattern_parameters_operator(v string) (any, error) {
    result := STARTS_WITH_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_PARAMETERS_OPERATOR
    switch v {
        case "starts_with":
            result = STARTS_WITH_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_PARAMETERS_OPERATOR
        case "ends_with":
            result = ENDS_WITH_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_PARAMETERS_OPERATOR
        case "contains":
            result = CONTAINS_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_PARAMETERS_OPERATOR
        case "regex":
            result = REGEX_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_PARAMETERS_OPERATOR
        default:
            return 0, errors.New("Unknown RepositoryRuleCommitAuthorEmailPattern_parameters_operator value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleCommitAuthorEmailPattern_parameters_operator(values []RepositoryRuleCommitAuthorEmailPattern_parameters_operator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleCommitAuthorEmailPattern_parameters_operator) isMultiValue() bool {
    return false
}
