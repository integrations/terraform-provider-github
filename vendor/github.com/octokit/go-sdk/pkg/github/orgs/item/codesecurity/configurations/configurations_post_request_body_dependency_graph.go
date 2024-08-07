package configurations
import (
    "errors"
)
// The enablement status of Dependency Graph
type ConfigurationsPostRequestBody_dependency_graph int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH ConfigurationsPostRequestBody_dependency_graph = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH
)

func (i ConfigurationsPostRequestBody_dependency_graph) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_dependency_graph(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH
        default:
            return 0, errors.New("Unknown ConfigurationsPostRequestBody_dependency_graph value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_dependency_graph(values []ConfigurationsPostRequestBody_dependency_graph) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_dependency_graph) isMultiValue() bool {
    return false
}
