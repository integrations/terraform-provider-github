package models
import (
    "errors"
)
// The state of the status.
type DeploymentStatus_state int

const (
    ERROR_DEPLOYMENTSTATUS_STATE DeploymentStatus_state = iota
    FAILURE_DEPLOYMENTSTATUS_STATE
    INACTIVE_DEPLOYMENTSTATUS_STATE
    PENDING_DEPLOYMENTSTATUS_STATE
    SUCCESS_DEPLOYMENTSTATUS_STATE
    QUEUED_DEPLOYMENTSTATUS_STATE
    IN_PROGRESS_DEPLOYMENTSTATUS_STATE
)

func (i DeploymentStatus_state) String() string {
    return []string{"error", "failure", "inactive", "pending", "success", "queued", "in_progress"}[i]
}
func ParseDeploymentStatus_state(v string) (any, error) {
    result := ERROR_DEPLOYMENTSTATUS_STATE
    switch v {
        case "error":
            result = ERROR_DEPLOYMENTSTATUS_STATE
        case "failure":
            result = FAILURE_DEPLOYMENTSTATUS_STATE
        case "inactive":
            result = INACTIVE_DEPLOYMENTSTATUS_STATE
        case "pending":
            result = PENDING_DEPLOYMENTSTATUS_STATE
        case "success":
            result = SUCCESS_DEPLOYMENTSTATUS_STATE
        case "queued":
            result = QUEUED_DEPLOYMENTSTATUS_STATE
        case "in_progress":
            result = IN_PROGRESS_DEPLOYMENTSTATUS_STATE
        default:
            return 0, errors.New("Unknown DeploymentStatus_state value: " + v)
    }
    return &result, nil
}
func SerializeDeploymentStatus_state(values []DeploymentStatus_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeploymentStatus_state) isMultiValue() bool {
    return false
}
