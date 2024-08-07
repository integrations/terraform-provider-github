package models
import (
    "errors"
)
// Whether to approve or reject deployment to the specified environments.
type ReviewCustomGatesStateRequired_state int

const (
    APPROVED_REVIEWCUSTOMGATESSTATEREQUIRED_STATE ReviewCustomGatesStateRequired_state = iota
    REJECTED_REVIEWCUSTOMGATESSTATEREQUIRED_STATE
)

func (i ReviewCustomGatesStateRequired_state) String() string {
    return []string{"approved", "rejected"}[i]
}
func ParseReviewCustomGatesStateRequired_state(v string) (any, error) {
    result := APPROVED_REVIEWCUSTOMGATESSTATEREQUIRED_STATE
    switch v {
        case "approved":
            result = APPROVED_REVIEWCUSTOMGATESSTATEREQUIRED_STATE
        case "rejected":
            result = REJECTED_REVIEWCUSTOMGATESSTATEREQUIRED_STATE
        default:
            return 0, errors.New("Unknown ReviewCustomGatesStateRequired_state value: " + v)
    }
    return &result, nil
}
func SerializeReviewCustomGatesStateRequired_state(values []ReviewCustomGatesStateRequired_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReviewCustomGatesStateRequired_state) isMultiValue() bool {
    return false
}
