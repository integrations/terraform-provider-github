package models
import (
    "errors"
)
// The reason for the current state
type Issue_state_reason int

const (
    COMPLETED_ISSUE_STATE_REASON Issue_state_reason = iota
    REOPENED_ISSUE_STATE_REASON
    NOT_PLANNED_ISSUE_STATE_REASON
)

func (i Issue_state_reason) String() string {
    return []string{"completed", "reopened", "not_planned"}[i]
}
func ParseIssue_state_reason(v string) (any, error) {
    result := COMPLETED_ISSUE_STATE_REASON
    switch v {
        case "completed":
            result = COMPLETED_ISSUE_STATE_REASON
        case "reopened":
            result = REOPENED_ISSUE_STATE_REASON
        case "not_planned":
            result = NOT_PLANNED_ISSUE_STATE_REASON
        default:
            return 0, errors.New("Unknown Issue_state_reason value: " + v)
    }
    return &result, nil
}
func SerializeIssue_state_reason(values []Issue_state_reason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Issue_state_reason) isMultiValue() bool {
    return false
}
