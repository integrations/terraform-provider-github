package models
import (
    "errors"
)
// The state of the milestone.
type Milestone_state int

const (
    OPEN_MILESTONE_STATE Milestone_state = iota
    CLOSED_MILESTONE_STATE
)

func (i Milestone_state) String() string {
    return []string{"open", "closed"}[i]
}
func ParseMilestone_state(v string) (any, error) {
    result := OPEN_MILESTONE_STATE
    switch v {
        case "open":
            result = OPEN_MILESTONE_STATE
        case "closed":
            result = CLOSED_MILESTONE_STATE
        default:
            return 0, errors.New("Unknown Milestone_state value: " + v)
    }
    return &result, nil
}
func SerializeMilestone_state(values []Milestone_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Milestone_state) isMultiValue() bool {
    return false
}
