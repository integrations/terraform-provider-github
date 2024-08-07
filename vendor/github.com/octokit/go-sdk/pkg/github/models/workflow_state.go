package models
import (
    "errors"
)
type Workflow_state int

const (
    ACTIVE_WORKFLOW_STATE Workflow_state = iota
    DELETED_WORKFLOW_STATE
    DISABLED_FORK_WORKFLOW_STATE
    DISABLED_INACTIVITY_WORKFLOW_STATE
    DISABLED_MANUALLY_WORKFLOW_STATE
)

func (i Workflow_state) String() string {
    return []string{"active", "deleted", "disabled_fork", "disabled_inactivity", "disabled_manually"}[i]
}
func ParseWorkflow_state(v string) (any, error) {
    result := ACTIVE_WORKFLOW_STATE
    switch v {
        case "active":
            result = ACTIVE_WORKFLOW_STATE
        case "deleted":
            result = DELETED_WORKFLOW_STATE
        case "disabled_fork":
            result = DISABLED_FORK_WORKFLOW_STATE
        case "disabled_inactivity":
            result = DISABLED_INACTIVITY_WORKFLOW_STATE
        case "disabled_manually":
            result = DISABLED_MANUALLY_WORKFLOW_STATE
        default:
            return 0, errors.New("Unknown Workflow_state value: " + v)
    }
    return &result, nil
}
func SerializeWorkflow_state(values []Workflow_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Workflow_state) isMultiValue() bool {
    return false
}
