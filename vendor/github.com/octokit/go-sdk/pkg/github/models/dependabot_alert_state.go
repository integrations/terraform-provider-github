package models
import (
    "errors"
)
// The state of the Dependabot alert.
type DependabotAlert_state int

const (
    AUTO_DISMISSED_DEPENDABOTALERT_STATE DependabotAlert_state = iota
    DISMISSED_DEPENDABOTALERT_STATE
    FIXED_DEPENDABOTALERT_STATE
    OPEN_DEPENDABOTALERT_STATE
)

func (i DependabotAlert_state) String() string {
    return []string{"auto_dismissed", "dismissed", "fixed", "open"}[i]
}
func ParseDependabotAlert_state(v string) (any, error) {
    result := AUTO_DISMISSED_DEPENDABOTALERT_STATE
    switch v {
        case "auto_dismissed":
            result = AUTO_DISMISSED_DEPENDABOTALERT_STATE
        case "dismissed":
            result = DISMISSED_DEPENDABOTALERT_STATE
        case "fixed":
            result = FIXED_DEPENDABOTALERT_STATE
        case "open":
            result = OPEN_DEPENDABOTALERT_STATE
        default:
            return 0, errors.New("Unknown DependabotAlert_state value: " + v)
    }
    return &result, nil
}
func SerializeDependabotAlert_state(values []DependabotAlert_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependabotAlert_state) isMultiValue() bool {
    return false
}
