package models
import (
    "errors"
)
type RepositoryRuleCreation_type int

const (
    CREATION_REPOSITORYRULECREATION_TYPE RepositoryRuleCreation_type = iota
)

func (i RepositoryRuleCreation_type) String() string {
    return []string{"creation"}[i]
}
func ParseRepositoryRuleCreation_type(v string) (any, error) {
    result := CREATION_REPOSITORYRULECREATION_TYPE
    switch v {
        case "creation":
            result = CREATION_REPOSITORYRULECREATION_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleCreation_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleCreation_type(values []RepositoryRuleCreation_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleCreation_type) isMultiValue() bool {
    return false
}
