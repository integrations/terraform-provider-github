package models
import (
    "errors"
)
// The execution scope of the vulnerable dependency.
type DependabotAlert_dependency_scope int

const (
    DEVELOPMENT_DEPENDABOTALERT_DEPENDENCY_SCOPE DependabotAlert_dependency_scope = iota
    RUNTIME_DEPENDABOTALERT_DEPENDENCY_SCOPE
)

func (i DependabotAlert_dependency_scope) String() string {
    return []string{"development", "runtime"}[i]
}
func ParseDependabotAlert_dependency_scope(v string) (any, error) {
    result := DEVELOPMENT_DEPENDABOTALERT_DEPENDENCY_SCOPE
    switch v {
        case "development":
            result = DEVELOPMENT_DEPENDABOTALERT_DEPENDENCY_SCOPE
        case "runtime":
            result = RUNTIME_DEPENDABOTALERT_DEPENDENCY_SCOPE
        default:
            return 0, errors.New("Unknown DependabotAlert_dependency_scope value: " + v)
    }
    return &result, nil
}
func SerializeDependabotAlert_dependency_scope(values []DependabotAlert_dependency_scope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependabotAlert_dependency_scope) isMultiValue() bool {
    return false
}
