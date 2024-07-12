package models
import (
    "errors"
)
type DependencyGraphDiff_change_type int

const (
    ADDED_DEPENDENCYGRAPHDIFF_CHANGE_TYPE DependencyGraphDiff_change_type = iota
    REMOVED_DEPENDENCYGRAPHDIFF_CHANGE_TYPE
)

func (i DependencyGraphDiff_change_type) String() string {
    return []string{"added", "removed"}[i]
}
func ParseDependencyGraphDiff_change_type(v string) (any, error) {
    result := ADDED_DEPENDENCYGRAPHDIFF_CHANGE_TYPE
    switch v {
        case "added":
            result = ADDED_DEPENDENCYGRAPHDIFF_CHANGE_TYPE
        case "removed":
            result = REMOVED_DEPENDENCYGRAPHDIFF_CHANGE_TYPE
        default:
            return 0, errors.New("Unknown DependencyGraphDiff_change_type value: " + v)
    }
    return &result, nil
}
func SerializeDependencyGraphDiff_change_type(values []DependencyGraphDiff_change_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependencyGraphDiff_change_type) isMultiValue() bool {
    return false
}
