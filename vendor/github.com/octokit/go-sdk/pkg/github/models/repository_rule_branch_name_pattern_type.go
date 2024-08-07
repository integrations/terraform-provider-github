package models
import (
    "errors"
)
type RepositoryRuleBranchNamePattern_type int

const (
    BRANCH_NAME_PATTERN_REPOSITORYRULEBRANCHNAMEPATTERN_TYPE RepositoryRuleBranchNamePattern_type = iota
)

func (i RepositoryRuleBranchNamePattern_type) String() string {
    return []string{"branch_name_pattern"}[i]
}
func ParseRepositoryRuleBranchNamePattern_type(v string) (any, error) {
    result := BRANCH_NAME_PATTERN_REPOSITORYRULEBRANCHNAMEPATTERN_TYPE
    switch v {
        case "branch_name_pattern":
            result = BRANCH_NAME_PATTERN_REPOSITORYRULEBRANCHNAMEPATTERN_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleBranchNamePattern_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleBranchNamePattern_type(values []RepositoryRuleBranchNamePattern_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleBranchNamePattern_type) isMultiValue() bool {
    return false
}
