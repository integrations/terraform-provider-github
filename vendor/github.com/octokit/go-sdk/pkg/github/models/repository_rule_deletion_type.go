package models
import (
    "errors"
)
type RepositoryRuleDeletion_type int

const (
    DELETION_REPOSITORYRULEDELETION_TYPE RepositoryRuleDeletion_type = iota
)

func (i RepositoryRuleDeletion_type) String() string {
    return []string{"deletion"}[i]
}
func ParseRepositoryRuleDeletion_type(v string) (any, error) {
    result := DELETION_REPOSITORYRULEDELETION_TYPE
    switch v {
        case "deletion":
            result = DELETION_REPOSITORYRULEDELETION_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleDeletion_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleDeletion_type(values []RepositoryRuleDeletion_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleDeletion_type) isMultiValue() bool {
    return false
}
