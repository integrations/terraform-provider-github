package models
import (
    "errors"
)
type CheckSuite_conclusion int

const (
    SUCCESS_CHECKSUITE_CONCLUSION CheckSuite_conclusion = iota
    FAILURE_CHECKSUITE_CONCLUSION
    NEUTRAL_CHECKSUITE_CONCLUSION
    CANCELLED_CHECKSUITE_CONCLUSION
    SKIPPED_CHECKSUITE_CONCLUSION
    TIMED_OUT_CHECKSUITE_CONCLUSION
    ACTION_REQUIRED_CHECKSUITE_CONCLUSION
    STARTUP_FAILURE_CHECKSUITE_CONCLUSION
    STALE_CHECKSUITE_CONCLUSION
)

func (i CheckSuite_conclusion) String() string {
    return []string{"success", "failure", "neutral", "cancelled", "skipped", "timed_out", "action_required", "startup_failure", "stale"}[i]
}
func ParseCheckSuite_conclusion(v string) (any, error) {
    result := SUCCESS_CHECKSUITE_CONCLUSION
    switch v {
        case "success":
            result = SUCCESS_CHECKSUITE_CONCLUSION
        case "failure":
            result = FAILURE_CHECKSUITE_CONCLUSION
        case "neutral":
            result = NEUTRAL_CHECKSUITE_CONCLUSION
        case "cancelled":
            result = CANCELLED_CHECKSUITE_CONCLUSION
        case "skipped":
            result = SKIPPED_CHECKSUITE_CONCLUSION
        case "timed_out":
            result = TIMED_OUT_CHECKSUITE_CONCLUSION
        case "action_required":
            result = ACTION_REQUIRED_CHECKSUITE_CONCLUSION
        case "startup_failure":
            result = STARTUP_FAILURE_CHECKSUITE_CONCLUSION
        case "stale":
            result = STALE_CHECKSUITE_CONCLUSION
        default:
            return 0, errors.New("Unknown CheckSuite_conclusion value: " + v)
    }
    return &result, nil
}
func SerializeCheckSuite_conclusion(values []CheckSuite_conclusion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CheckSuite_conclusion) isMultiValue() bool {
    return false
}
