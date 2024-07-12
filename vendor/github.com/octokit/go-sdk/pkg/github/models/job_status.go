package models
import (
    "errors"
)
// The phase of the lifecycle that the job is currently in.
type Job_status int

const (
    QUEUED_JOB_STATUS Job_status = iota
    IN_PROGRESS_JOB_STATUS
    COMPLETED_JOB_STATUS
    WAITING_JOB_STATUS
    REQUESTED_JOB_STATUS
    PENDING_JOB_STATUS
)

func (i Job_status) String() string {
    return []string{"queued", "in_progress", "completed", "waiting", "requested", "pending"}[i]
}
func ParseJob_status(v string) (any, error) {
    result := QUEUED_JOB_STATUS
    switch v {
        case "queued":
            result = QUEUED_JOB_STATUS
        case "in_progress":
            result = IN_PROGRESS_JOB_STATUS
        case "completed":
            result = COMPLETED_JOB_STATUS
        case "waiting":
            result = WAITING_JOB_STATUS
        case "requested":
            result = REQUESTED_JOB_STATUS
        case "pending":
            result = PENDING_JOB_STATUS
        default:
            return 0, errors.New("Unknown Job_status value: " + v)
    }
    return &result, nil
}
func SerializeJob_status(values []Job_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Job_status) isMultiValue() bool {
    return false
}
