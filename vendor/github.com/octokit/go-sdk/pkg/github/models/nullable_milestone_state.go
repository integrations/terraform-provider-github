package models
import (
    "errors"
)
// The state of the milestone.
type NullableMilestone_state int

const (
    OPEN_NULLABLEMILESTONE_STATE NullableMilestone_state = iota
    CLOSED_NULLABLEMILESTONE_STATE
)

func (i NullableMilestone_state) String() string {
    return []string{"open", "closed"}[i]
}
func ParseNullableMilestone_state(v string) (any, error) {
    result := OPEN_NULLABLEMILESTONE_STATE
    switch v {
        case "open":
            result = OPEN_NULLABLEMILESTONE_STATE
        case "closed":
            result = CLOSED_NULLABLEMILESTONE_STATE
        default:
            return 0, errors.New("Unknown NullableMilestone_state value: " + v)
    }
    return &result, nil
}
func SerializeNullableMilestone_state(values []NullableMilestone_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NullableMilestone_state) isMultiValue() bool {
    return false
}
