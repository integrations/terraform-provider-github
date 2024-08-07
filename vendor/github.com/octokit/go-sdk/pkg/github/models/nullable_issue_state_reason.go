package models
import (
    "errors"
)
// The reason for the current state
type NullableIssue_state_reason int

const (
    COMPLETED_NULLABLEISSUE_STATE_REASON NullableIssue_state_reason = iota
    REOPENED_NULLABLEISSUE_STATE_REASON
    NOT_PLANNED_NULLABLEISSUE_STATE_REASON
)

func (i NullableIssue_state_reason) String() string {
    return []string{"completed", "reopened", "not_planned"}[i]
}
func ParseNullableIssue_state_reason(v string) (any, error) {
    result := COMPLETED_NULLABLEISSUE_STATE_REASON
    switch v {
        case "completed":
            result = COMPLETED_NULLABLEISSUE_STATE_REASON
        case "reopened":
            result = REOPENED_NULLABLEISSUE_STATE_REASON
        case "not_planned":
            result = NOT_PLANNED_NULLABLEISSUE_STATE_REASON
        default:
            return 0, errors.New("Unknown NullableIssue_state_reason value: " + v)
    }
    return &result, nil
}
func SerializeNullableIssue_state_reason(values []NullableIssue_state_reason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NullableIssue_state_reason) isMultiValue() bool {
    return false
}
