package models
import (
    "errors"
)
type CheckRun_conclusion int

const (
    SUCCESS_CHECKRUN_CONCLUSION CheckRun_conclusion = iota
    FAILURE_CHECKRUN_CONCLUSION
    NEUTRAL_CHECKRUN_CONCLUSION
    CANCELLED_CHECKRUN_CONCLUSION
    SKIPPED_CHECKRUN_CONCLUSION
    TIMED_OUT_CHECKRUN_CONCLUSION
    ACTION_REQUIRED_CHECKRUN_CONCLUSION
)

func (i CheckRun_conclusion) String() string {
    return []string{"success", "failure", "neutral", "cancelled", "skipped", "timed_out", "action_required"}[i]
}
func ParseCheckRun_conclusion(v string) (any, error) {
    result := SUCCESS_CHECKRUN_CONCLUSION
    switch v {
        case "success":
            result = SUCCESS_CHECKRUN_CONCLUSION
        case "failure":
            result = FAILURE_CHECKRUN_CONCLUSION
        case "neutral":
            result = NEUTRAL_CHECKRUN_CONCLUSION
        case "cancelled":
            result = CANCELLED_CHECKRUN_CONCLUSION
        case "skipped":
            result = SKIPPED_CHECKRUN_CONCLUSION
        case "timed_out":
            result = TIMED_OUT_CHECKRUN_CONCLUSION
        case "action_required":
            result = ACTION_REQUIRED_CHECKRUN_CONCLUSION
        default:
            return 0, errors.New("Unknown CheckRun_conclusion value: " + v)
    }
    return &result, nil
}
func SerializeCheckRun_conclusion(values []CheckRun_conclusion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CheckRun_conclusion) isMultiValue() bool {
    return false
}
