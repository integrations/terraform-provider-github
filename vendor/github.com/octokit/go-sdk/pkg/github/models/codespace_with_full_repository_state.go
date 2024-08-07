package models
import (
    "errors"
)
// State of this codespace.
type CodespaceWithFullRepository_state int

const (
    UNKNOWN_CODESPACEWITHFULLREPOSITORY_STATE CodespaceWithFullRepository_state = iota
    CREATED_CODESPACEWITHFULLREPOSITORY_STATE
    QUEUED_CODESPACEWITHFULLREPOSITORY_STATE
    PROVISIONING_CODESPACEWITHFULLREPOSITORY_STATE
    AVAILABLE_CODESPACEWITHFULLREPOSITORY_STATE
    AWAITING_CODESPACEWITHFULLREPOSITORY_STATE
    UNAVAILABLE_CODESPACEWITHFULLREPOSITORY_STATE
    DELETED_CODESPACEWITHFULLREPOSITORY_STATE
    MOVED_CODESPACEWITHFULLREPOSITORY_STATE
    SHUTDOWN_CODESPACEWITHFULLREPOSITORY_STATE
    ARCHIVED_CODESPACEWITHFULLREPOSITORY_STATE
    STARTING_CODESPACEWITHFULLREPOSITORY_STATE
    SHUTTINGDOWN_CODESPACEWITHFULLREPOSITORY_STATE
    FAILED_CODESPACEWITHFULLREPOSITORY_STATE
    EXPORTING_CODESPACEWITHFULLREPOSITORY_STATE
    UPDATING_CODESPACEWITHFULLREPOSITORY_STATE
    REBUILDING_CODESPACEWITHFULLREPOSITORY_STATE
)

func (i CodespaceWithFullRepository_state) String() string {
    return []string{"Unknown", "Created", "Queued", "Provisioning", "Available", "Awaiting", "Unavailable", "Deleted", "Moved", "Shutdown", "Archived", "Starting", "ShuttingDown", "Failed", "Exporting", "Updating", "Rebuilding"}[i]
}
func ParseCodespaceWithFullRepository_state(v string) (any, error) {
    result := UNKNOWN_CODESPACEWITHFULLREPOSITORY_STATE
    switch v {
        case "Unknown":
            result = UNKNOWN_CODESPACEWITHFULLREPOSITORY_STATE
        case "Created":
            result = CREATED_CODESPACEWITHFULLREPOSITORY_STATE
        case "Queued":
            result = QUEUED_CODESPACEWITHFULLREPOSITORY_STATE
        case "Provisioning":
            result = PROVISIONING_CODESPACEWITHFULLREPOSITORY_STATE
        case "Available":
            result = AVAILABLE_CODESPACEWITHFULLREPOSITORY_STATE
        case "Awaiting":
            result = AWAITING_CODESPACEWITHFULLREPOSITORY_STATE
        case "Unavailable":
            result = UNAVAILABLE_CODESPACEWITHFULLREPOSITORY_STATE
        case "Deleted":
            result = DELETED_CODESPACEWITHFULLREPOSITORY_STATE
        case "Moved":
            result = MOVED_CODESPACEWITHFULLREPOSITORY_STATE
        case "Shutdown":
            result = SHUTDOWN_CODESPACEWITHFULLREPOSITORY_STATE
        case "Archived":
            result = ARCHIVED_CODESPACEWITHFULLREPOSITORY_STATE
        case "Starting":
            result = STARTING_CODESPACEWITHFULLREPOSITORY_STATE
        case "ShuttingDown":
            result = SHUTTINGDOWN_CODESPACEWITHFULLREPOSITORY_STATE
        case "Failed":
            result = FAILED_CODESPACEWITHFULLREPOSITORY_STATE
        case "Exporting":
            result = EXPORTING_CODESPACEWITHFULLREPOSITORY_STATE
        case "Updating":
            result = UPDATING_CODESPACEWITHFULLREPOSITORY_STATE
        case "Rebuilding":
            result = REBUILDING_CODESPACEWITHFULLREPOSITORY_STATE
        default:
            return 0, errors.New("Unknown CodespaceWithFullRepository_state value: " + v)
    }
    return &result, nil
}
func SerializeCodespaceWithFullRepository_state(values []CodespaceWithFullRepository_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodespaceWithFullRepository_state) isMultiValue() bool {
    return false
}
