package models
import (
    "errors"
)
// The phase of the lifecycle that the job is currently in.
type Job_steps_status int

const (
    QUEUED_JOB_STEPS_STATUS Job_steps_status = iota
    IN_PROGRESS_JOB_STEPS_STATUS
    COMPLETED_JOB_STEPS_STATUS
)

func (i Job_steps_status) String() string {
    return []string{"queued", "in_progress", "completed"}[i]
}
func ParseJob_steps_status(v string) (any, error) {
    result := QUEUED_JOB_STEPS_STATUS
    switch v {
        case "queued":
            result = QUEUED_JOB_STEPS_STATUS
        case "in_progress":
            result = IN_PROGRESS_JOB_STEPS_STATUS
        case "completed":
            result = COMPLETED_JOB_STEPS_STATUS
        default:
            return 0, errors.New("Unknown Job_steps_status value: " + v)
    }
    return &result, nil
}
func SerializeJob_steps_status(values []Job_steps_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Job_steps_status) isMultiValue() bool {
    return false
}
