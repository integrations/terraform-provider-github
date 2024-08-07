package models
import (
    "errors"
)
// The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
type RepositoryRuleEnforcement int

const (
    DISABLED_REPOSITORYRULEENFORCEMENT RepositoryRuleEnforcement = iota
    ACTIVE_REPOSITORYRULEENFORCEMENT
    EVALUATE_REPOSITORYRULEENFORCEMENT
)

func (i RepositoryRuleEnforcement) String() string {
    return []string{"disabled", "active", "evaluate"}[i]
}
func ParseRepositoryRuleEnforcement(v string) (any, error) {
    result := DISABLED_REPOSITORYRULEENFORCEMENT
    switch v {
        case "disabled":
            result = DISABLED_REPOSITORYRULEENFORCEMENT
        case "active":
            result = ACTIVE_REPOSITORYRULEENFORCEMENT
        case "evaluate":
            result = EVALUATE_REPOSITORYRULEENFORCEMENT
        default:
            return 0, errors.New("Unknown RepositoryRuleEnforcement value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleEnforcement(values []RepositoryRuleEnforcement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleEnforcement) isMultiValue() bool {
    return false
}
