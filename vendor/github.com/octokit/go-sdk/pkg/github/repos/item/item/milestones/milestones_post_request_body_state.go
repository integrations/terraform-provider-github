package milestones
import (
    "errors"
)
// The state of the milestone. Either `open` or `closed`.
type MilestonesPostRequestBody_state int

const (
    OPEN_MILESTONESPOSTREQUESTBODY_STATE MilestonesPostRequestBody_state = iota
    CLOSED_MILESTONESPOSTREQUESTBODY_STATE
)

func (i MilestonesPostRequestBody_state) String() string {
    return []string{"open", "closed"}[i]
}
func ParseMilestonesPostRequestBody_state(v string) (any, error) {
    result := OPEN_MILESTONESPOSTREQUESTBODY_STATE
    switch v {
        case "open":
            result = OPEN_MILESTONESPOSTREQUESTBODY_STATE
        case "closed":
            result = CLOSED_MILESTONESPOSTREQUESTBODY_STATE
        default:
            return 0, errors.New("Unknown MilestonesPostRequestBody_state value: " + v)
    }
    return &result, nil
}
func SerializeMilestonesPostRequestBody_state(values []MilestonesPostRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MilestonesPostRequestBody_state) isMultiValue() bool {
    return false
}
