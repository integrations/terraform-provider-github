package models
import (
    "errors"
)
// The reason that the alert was dismissed.
type DependabotAlertWithRepository_dismissed_reason int

const (
    FIX_STARTED_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON DependabotAlertWithRepository_dismissed_reason = iota
    INACCURATE_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
    NO_BANDWIDTH_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
    NOT_USED_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
    TOLERABLE_RISK_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
)

func (i DependabotAlertWithRepository_dismissed_reason) String() string {
    return []string{"fix_started", "inaccurate", "no_bandwidth", "not_used", "tolerable_risk"}[i]
}
func ParseDependabotAlertWithRepository_dismissed_reason(v string) (any, error) {
    result := FIX_STARTED_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
    switch v {
        case "fix_started":
            result = FIX_STARTED_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
        case "inaccurate":
            result = INACCURATE_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
        case "no_bandwidth":
            result = NO_BANDWIDTH_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
        case "not_used":
            result = NOT_USED_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
        case "tolerable_risk":
            result = TOLERABLE_RISK_DEPENDABOTALERTWITHREPOSITORY_DISMISSED_REASON
        default:
            return 0, errors.New("Unknown DependabotAlertWithRepository_dismissed_reason value: " + v)
    }
    return &result, nil
}
func SerializeDependabotAlertWithRepository_dismissed_reason(values []DependabotAlertWithRepository_dismissed_reason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependabotAlertWithRepository_dismissed_reason) isMultiValue() bool {
    return false
}
