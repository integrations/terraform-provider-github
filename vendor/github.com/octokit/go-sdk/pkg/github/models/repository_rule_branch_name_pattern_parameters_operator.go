package models
import (
    "errors"
)
// The operator to use for matching.
type RepositoryRuleBranchNamePattern_parameters_operator int

const (
    STARTS_WITH_REPOSITORYRULEBRANCHNAMEPATTERN_PARAMETERS_OPERATOR RepositoryRuleBranchNamePattern_parameters_operator = iota
    ENDS_WITH_REPOSITORYRULEBRANCHNAMEPATTERN_PARAMETERS_OPERATOR
    CONTAINS_REPOSITORYRULEBRANCHNAMEPATTERN_PARAMETERS_OPERATOR
    REGEX_REPOSITORYRULEBRANCHNAMEPATTERN_PARAMETERS_OPERATOR
)

func (i RepositoryRuleBranchNamePattern_parameters_operator) String() string {
    return []string{"starts_with", "ends_with", "contains", "regex"}[i]
}
func ParseRepositoryRuleBranchNamePattern_parameters_operator(v string) (any, error) {
    result := STARTS_WITH_REPOSITORYRULEBRANCHNAMEPATTERN_PARAMETERS_OPERATOR
    switch v {
        case "starts_with":
            result = STARTS_WITH_REPOSITORYRULEBRANCHNAMEPATTERN_PARAMETERS_OPERATOR
        case "ends_with":
            result = ENDS_WITH_REPOSITORYRULEBRANCHNAMEPATTERN_PARAMETERS_OPERATOR
        case "contains":
            result = CONTAINS_REPOSITORYRULEBRANCHNAMEPATTERN_PARAMETERS_OPERATOR
        case "regex":
            result = REGEX_REPOSITORYRULEBRANCHNAMEPATTERN_PARAMETERS_OPERATOR
        default:
            return 0, errors.New("Unknown RepositoryRuleBranchNamePattern_parameters_operator value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleBranchNamePattern_parameters_operator(values []RepositoryRuleBranchNamePattern_parameters_operator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleBranchNamePattern_parameters_operator) isMultiValue() bool {
    return false
}
