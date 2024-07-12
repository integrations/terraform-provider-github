package runs
import (
    "errors"
)
type GetStatusQueryParameterType int

const (
    COMPLETED_GETSTATUSQUERYPARAMETERTYPE GetStatusQueryParameterType = iota
    ACTION_REQUIRED_GETSTATUSQUERYPARAMETERTYPE
    CANCELLED_GETSTATUSQUERYPARAMETERTYPE
    FAILURE_GETSTATUSQUERYPARAMETERTYPE
    NEUTRAL_GETSTATUSQUERYPARAMETERTYPE
    SKIPPED_GETSTATUSQUERYPARAMETERTYPE
    STALE_GETSTATUSQUERYPARAMETERTYPE
    SUCCESS_GETSTATUSQUERYPARAMETERTYPE
    TIMED_OUT_GETSTATUSQUERYPARAMETERTYPE
    IN_PROGRESS_GETSTATUSQUERYPARAMETERTYPE
    QUEUED_GETSTATUSQUERYPARAMETERTYPE
    REQUESTED_GETSTATUSQUERYPARAMETERTYPE
    WAITING_GETSTATUSQUERYPARAMETERTYPE
    PENDING_GETSTATUSQUERYPARAMETERTYPE
)

func (i GetStatusQueryParameterType) String() string {
    return []string{"completed", "action_required", "cancelled", "failure", "neutral", "skipped", "stale", "success", "timed_out", "in_progress", "queued", "requested", "waiting", "pending"}[i]
}
func ParseGetStatusQueryParameterType(v string) (any, error) {
    result := COMPLETED_GETSTATUSQUERYPARAMETERTYPE
    switch v {
        case "completed":
            result = COMPLETED_GETSTATUSQUERYPARAMETERTYPE
        case "action_required":
            result = ACTION_REQUIRED_GETSTATUSQUERYPARAMETERTYPE
        case "cancelled":
            result = CANCELLED_GETSTATUSQUERYPARAMETERTYPE
        case "failure":
            result = FAILURE_GETSTATUSQUERYPARAMETERTYPE
        case "neutral":
            result = NEUTRAL_GETSTATUSQUERYPARAMETERTYPE
        case "skipped":
            result = SKIPPED_GETSTATUSQUERYPARAMETERTYPE
        case "stale":
            result = STALE_GETSTATUSQUERYPARAMETERTYPE
        case "success":
            result = SUCCESS_GETSTATUSQUERYPARAMETERTYPE
        case "timed_out":
            result = TIMED_OUT_GETSTATUSQUERYPARAMETERTYPE
        case "in_progress":
            result = IN_PROGRESS_GETSTATUSQUERYPARAMETERTYPE
        case "queued":
            result = QUEUED_GETSTATUSQUERYPARAMETERTYPE
        case "requested":
            result = REQUESTED_GETSTATUSQUERYPARAMETERTYPE
        case "waiting":
            result = WAITING_GETSTATUSQUERYPARAMETERTYPE
        case "pending":
            result = PENDING_GETSTATUSQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetStatusQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetStatusQueryParameterType(values []GetStatusQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetStatusQueryParameterType) isMultiValue() bool {
    return false
}
