package models
import (
    "errors"
)
// The level of privacy this team should have
type TeamFull_privacy int

const (
    CLOSED_TEAMFULL_PRIVACY TeamFull_privacy = iota
    SECRET_TEAMFULL_PRIVACY
)

func (i TeamFull_privacy) String() string {
    return []string{"closed", "secret"}[i]
}
func ParseTeamFull_privacy(v string) (any, error) {
    result := CLOSED_TEAMFULL_PRIVACY
    switch v {
        case "closed":
            result = CLOSED_TEAMFULL_PRIVACY
        case "secret":
            result = SECRET_TEAMFULL_PRIVACY
        default:
            return 0, errors.New("Unknown TeamFull_privacy value: " + v)
    }
    return &result, nil
}
func SerializeTeamFull_privacy(values []TeamFull_privacy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamFull_privacy) isMultiValue() bool {
    return false
}
