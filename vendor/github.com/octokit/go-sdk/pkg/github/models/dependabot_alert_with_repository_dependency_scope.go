package models
import (
    "errors"
)
// The execution scope of the vulnerable dependency.
type DependabotAlertWithRepository_dependency_scope int

const (
    DEVELOPMENT_DEPENDABOTALERTWITHREPOSITORY_DEPENDENCY_SCOPE DependabotAlertWithRepository_dependency_scope = iota
    RUNTIME_DEPENDABOTALERTWITHREPOSITORY_DEPENDENCY_SCOPE
)

func (i DependabotAlertWithRepository_dependency_scope) String() string {
    return []string{"development", "runtime"}[i]
}
func ParseDependabotAlertWithRepository_dependency_scope(v string) (any, error) {
    result := DEVELOPMENT_DEPENDABOTALERTWITHREPOSITORY_DEPENDENCY_SCOPE
    switch v {
        case "development":
            result = DEVELOPMENT_DEPENDABOTALERTWITHREPOSITORY_DEPENDENCY_SCOPE
        case "runtime":
            result = RUNTIME_DEPENDABOTALERTWITHREPOSITORY_DEPENDENCY_SCOPE
        default:
            return 0, errors.New("Unknown DependabotAlertWithRepository_dependency_scope value: " + v)
    }
    return &result, nil
}
func SerializeDependabotAlertWithRepository_dependency_scope(values []DependabotAlertWithRepository_dependency_scope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependabotAlertWithRepository_dependency_scope) isMultiValue() bool {
    return false
}
