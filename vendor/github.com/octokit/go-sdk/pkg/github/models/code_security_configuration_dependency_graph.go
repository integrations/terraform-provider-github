package models
import (
    "errors"
)
// The enablement status of Dependency Graph
type CodeSecurityConfiguration_dependency_graph int

const (
    ENABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH CodeSecurityConfiguration_dependency_graph = iota
    DISABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH
    NOT_SET_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH
)

func (i CodeSecurityConfiguration_dependency_graph) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_dependency_graph(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH
        default:
            return 0, errors.New("Unknown CodeSecurityConfiguration_dependency_graph value: " + v)
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_dependency_graph(values []CodeSecurityConfiguration_dependency_graph) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_dependency_graph) isMultiValue() bool {
    return false
}
