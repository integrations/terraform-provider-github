package models
import (
    "errors"
)
// The reason that the alert was dismissed.
type DependabotAlert_dismissed_reason int

const (
    FIX_STARTED_DEPENDABOTALERT_DISMISSED_REASON DependabotAlert_dismissed_reason = iota
    INACCURATE_DEPENDABOTALERT_DISMISSED_REASON
    NO_BANDWIDTH_DEPENDABOTALERT_DISMISSED_REASON
    NOT_USED_DEPENDABOTALERT_DISMISSED_REASON
    TOLERABLE_RISK_DEPENDABOTALERT_DISMISSED_REASON
)

func (i DependabotAlert_dismissed_reason) String() string {
    return []string{"fix_started", "inaccurate", "no_bandwidth", "not_used", "tolerable_risk"}[i]
}
func ParseDependabotAlert_dismissed_reason(v string) (any, error) {
    result := FIX_STARTED_DEPENDABOTALERT_DISMISSED_REASON
    switch v {
        case "fix_started":
            result = FIX_STARTED_DEPENDABOTALERT_DISMISSED_REASON
        case "inaccurate":
            result = INACCURATE_DEPENDABOTALERT_DISMISSED_REASON
        case "no_bandwidth":
            result = NO_BANDWIDTH_DEPENDABOTALERT_DISMISSED_REASON
        case "not_used":
            result = NOT_USED_DEPENDABOTALERT_DISMISSED_REASON
        case "tolerable_risk":
            result = TOLERABLE_RISK_DEPENDABOTALERT_DISMISSED_REASON
        default:
            return 0, errors.New("Unknown DependabotAlert_dismissed_reason value: " + v)
    }
    return &result, nil
}
func SerializeDependabotAlert_dismissed_reason(values []DependabotAlert_dismissed_reason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependabotAlert_dismissed_reason) isMultiValue() bool {
    return false
}
