package models
import (
    "errors"
)
// Whether deployment to the environment(s) was approved or rejected or pending (with comments)
type EnvironmentApprovals_state int

const (
    APPROVED_ENVIRONMENTAPPROVALS_STATE EnvironmentApprovals_state = iota
    REJECTED_ENVIRONMENTAPPROVALS_STATE
    PENDING_ENVIRONMENTAPPROVALS_STATE
)

func (i EnvironmentApprovals_state) String() string {
    return []string{"approved", "rejected", "pending"}[i]
}
func ParseEnvironmentApprovals_state(v string) (any, error) {
    result := APPROVED_ENVIRONMENTAPPROVALS_STATE
    switch v {
        case "approved":
            result = APPROVED_ENVIRONMENTAPPROVALS_STATE
        case "rejected":
            result = REJECTED_ENVIRONMENTAPPROVALS_STATE
        case "pending":
            result = PENDING_ENVIRONMENTAPPROVALS_STATE
        default:
            return 0, errors.New("Unknown EnvironmentApprovals_state value: " + v)
    }
    return &result, nil
}
func SerializeEnvironmentApprovals_state(values []EnvironmentApprovals_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EnvironmentApprovals_state) isMultiValue() bool {
    return false
}
