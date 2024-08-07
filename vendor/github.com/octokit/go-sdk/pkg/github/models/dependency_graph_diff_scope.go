package models
import (
    "errors"
)
// Where the dependency is utilized. `development` means that the dependency is only utilized in the development environment. `runtime` means that the dependency is utilized at runtime and in the development environment.
type DependencyGraphDiff_scope int

const (
    UNKNOWN_DEPENDENCYGRAPHDIFF_SCOPE DependencyGraphDiff_scope = iota
    RUNTIME_DEPENDENCYGRAPHDIFF_SCOPE
    DEVELOPMENT_DEPENDENCYGRAPHDIFF_SCOPE
)

func (i DependencyGraphDiff_scope) String() string {
    return []string{"unknown", "runtime", "development"}[i]
}
func ParseDependencyGraphDiff_scope(v string) (any, error) {
    result := UNKNOWN_DEPENDENCYGRAPHDIFF_SCOPE
    switch v {
        case "unknown":
            result = UNKNOWN_DEPENDENCYGRAPHDIFF_SCOPE
        case "runtime":
            result = RUNTIME_DEPENDENCYGRAPHDIFF_SCOPE
        case "development":
            result = DEVELOPMENT_DEPENDENCYGRAPHDIFF_SCOPE
        default:
            return 0, errors.New("Unknown DependencyGraphDiff_scope value: " + v)
    }
    return &result, nil
}
func SerializeDependencyGraphDiff_scope(values []DependencyGraphDiff_scope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependencyGraphDiff_scope) isMultiValue() bool {
    return false
}
