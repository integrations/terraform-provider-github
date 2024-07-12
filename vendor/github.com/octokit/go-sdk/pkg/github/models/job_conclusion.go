package models
import (
    "errors"
)
// The outcome of the job.
type Job_conclusion int

const (
    SUCCESS_JOB_CONCLUSION Job_conclusion = iota
    FAILURE_JOB_CONCLUSION
    NEUTRAL_JOB_CONCLUSION
    CANCELLED_JOB_CONCLUSION
    SKIPPED_JOB_CONCLUSION
    TIMED_OUT_JOB_CONCLUSION
    ACTION_REQUIRED_JOB_CONCLUSION
)

func (i Job_conclusion) String() string {
    return []string{"success", "failure", "neutral", "cancelled", "skipped", "timed_out", "action_required"}[i]
}
func ParseJob_conclusion(v string) (any, error) {
    result := SUCCESS_JOB_CONCLUSION
    switch v {
        case "success":
            result = SUCCESS_JOB_CONCLUSION
        case "failure":
            result = FAILURE_JOB_CONCLUSION
        case "neutral":
            result = NEUTRAL_JOB_CONCLUSION
        case "cancelled":
            result = CANCELLED_JOB_CONCLUSION
        case "skipped":
            result = SKIPPED_JOB_CONCLUSION
        case "timed_out":
            result = TIMED_OUT_JOB_CONCLUSION
        case "action_required":
            result = ACTION_REQUIRED_JOB_CONCLUSION
        default:
            return 0, errors.New("Unknown Job_conclusion value: " + v)
    }
    return &result, nil
}
func SerializeJob_conclusion(values []Job_conclusion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Job_conclusion) isMultiValue() bool {
    return false
}
