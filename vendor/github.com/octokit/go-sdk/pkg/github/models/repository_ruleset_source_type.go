package models
import (
    "errors"
)
// The type of the source of the ruleset
type RepositoryRuleset_source_type int

const (
    REPOSITORY_REPOSITORYRULESET_SOURCE_TYPE RepositoryRuleset_source_type = iota
    ORGANIZATION_REPOSITORYRULESET_SOURCE_TYPE
)

func (i RepositoryRuleset_source_type) String() string {
    return []string{"Repository", "Organization"}[i]
}
func ParseRepositoryRuleset_source_type(v string) (any, error) {
    result := REPOSITORY_REPOSITORYRULESET_SOURCE_TYPE
    switch v {
        case "Repository":
            result = REPOSITORY_REPOSITORYRULESET_SOURCE_TYPE
        case "Organization":
            result = ORGANIZATION_REPOSITORYRULESET_SOURCE_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleset_source_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleset_source_type(values []RepositoryRuleset_source_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleset_source_type) isMultiValue() bool {
    return false
}
