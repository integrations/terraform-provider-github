package models
import (
    "errors"
)
// State of this codespace.
type Codespace_state int

const (
    UNKNOWN_CODESPACE_STATE Codespace_state = iota
    CREATED_CODESPACE_STATE
    QUEUED_CODESPACE_STATE
    PROVISIONING_CODESPACE_STATE
    AVAILABLE_CODESPACE_STATE
    AWAITING_CODESPACE_STATE
    UNAVAILABLE_CODESPACE_STATE
    DELETED_CODESPACE_STATE
    MOVED_CODESPACE_STATE
    SHUTDOWN_CODESPACE_STATE
    ARCHIVED_CODESPACE_STATE
    STARTING_CODESPACE_STATE
    SHUTTINGDOWN_CODESPACE_STATE
    FAILED_CODESPACE_STATE
    EXPORTING_CODESPACE_STATE
    UPDATING_CODESPACE_STATE
    REBUILDING_CODESPACE_STATE
)

func (i Codespace_state) String() string {
    return []string{"Unknown", "Created", "Queued", "Provisioning", "Available", "Awaiting", "Unavailable", "Deleted", "Moved", "Shutdown", "Archived", "Starting", "ShuttingDown", "Failed", "Exporting", "Updating", "Rebuilding"}[i]
}
func ParseCodespace_state(v string) (any, error) {
    result := UNKNOWN_CODESPACE_STATE
    switch v {
        case "Unknown":
            result = UNKNOWN_CODESPACE_STATE
        case "Created":
            result = CREATED_CODESPACE_STATE
        case "Queued":
            result = QUEUED_CODESPACE_STATE
        case "Provisioning":
            result = PROVISIONING_CODESPACE_STATE
        case "Available":
            result = AVAILABLE_CODESPACE_STATE
        case "Awaiting":
            result = AWAITING_CODESPACE_STATE
        case "Unavailable":
            result = UNAVAILABLE_CODESPACE_STATE
        case "Deleted":
            result = DELETED_CODESPACE_STATE
        case "Moved":
            result = MOVED_CODESPACE_STATE
        case "Shutdown":
            result = SHUTDOWN_CODESPACE_STATE
        case "Archived":
            result = ARCHIVED_CODESPACE_STATE
        case "Starting":
            result = STARTING_CODESPACE_STATE
        case "ShuttingDown":
            result = SHUTTINGDOWN_CODESPACE_STATE
        case "Failed":
            result = FAILED_CODESPACE_STATE
        case "Exporting":
            result = EXPORTING_CODESPACE_STATE
        case "Updating":
            result = UPDATING_CODESPACE_STATE
        case "Rebuilding":
            result = REBUILDING_CODESPACE_STATE
        default:
            return 0, errors.New("Unknown Codespace_state value: " + v)
    }
    return &result, nil
}
func SerializeCodespace_state(values []Codespace_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Codespace_state) isMultiValue() bool {
    return false
}
