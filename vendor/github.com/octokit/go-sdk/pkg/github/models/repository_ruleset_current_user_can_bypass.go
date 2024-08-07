package models
import (
    "errors"
)
// The bypass type of the user making the API request for this ruleset. This field is only returned whenquerying the repository-level endpoint.
type RepositoryRuleset_current_user_can_bypass int

const (
    ALWAYS_REPOSITORYRULESET_CURRENT_USER_CAN_BYPASS RepositoryRuleset_current_user_can_bypass = iota
    PULL_REQUESTS_ONLY_REPOSITORYRULESET_CURRENT_USER_CAN_BYPASS
    NEVER_REPOSITORYRULESET_CURRENT_USER_CAN_BYPASS
)

func (i RepositoryRuleset_current_user_can_bypass) String() string {
    return []string{"always", "pull_requests_only", "never"}[i]
}
func ParseRepositoryRuleset_current_user_can_bypass(v string) (any, error) {
    result := ALWAYS_REPOSITORYRULESET_CURRENT_USER_CAN_BYPASS
    switch v {
        case "always":
            result = ALWAYS_REPOSITORYRULESET_CURRENT_USER_CAN_BYPASS
        case "pull_requests_only":
            result = PULL_REQUESTS_ONLY_REPOSITORYRULESET_CURRENT_USER_CAN_BYPASS
        case "never":
            result = NEVER_REPOSITORYRULESET_CURRENT_USER_CAN_BYPASS
        default:
            return 0, errors.New("Unknown RepositoryRuleset_current_user_can_bypass value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleset_current_user_can_bypass(values []RepositoryRuleset_current_user_can_bypass) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleset_current_user_can_bypass) isMultiValue() bool {
    return false
}
