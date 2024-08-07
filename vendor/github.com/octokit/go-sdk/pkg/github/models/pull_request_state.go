package models
import (
    "errors"
)
// State of this Pull Request. Either `open` or `closed`.
type PullRequest_state int

const (
    OPEN_PULLREQUEST_STATE PullRequest_state = iota
    CLOSED_PULLREQUEST_STATE
)

func (i PullRequest_state) String() string {
    return []string{"open", "closed"}[i]
}
func ParsePullRequest_state(v string) (any, error) {
    result := OPEN_PULLREQUEST_STATE
    switch v {
        case "open":
            result = OPEN_PULLREQUEST_STATE
        case "closed":
            result = CLOSED_PULLREQUEST_STATE
        default:
            return 0, errors.New("Unknown PullRequest_state value: " + v)
    }
    return &result, nil
}
func SerializePullRequest_state(values []PullRequest_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PullRequest_state) isMultiValue() bool {
    return false
}
