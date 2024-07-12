package models
import (
    "errors"
)
// When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type.
type RepositoryRulesetBypassActor_bypass_mode int

const (
    ALWAYS_REPOSITORYRULESETBYPASSACTOR_BYPASS_MODE RepositoryRulesetBypassActor_bypass_mode = iota
    PULL_REQUEST_REPOSITORYRULESETBYPASSACTOR_BYPASS_MODE
)

func (i RepositoryRulesetBypassActor_bypass_mode) String() string {
    return []string{"always", "pull_request"}[i]
}
func ParseRepositoryRulesetBypassActor_bypass_mode(v string) (any, error) {
    result := ALWAYS_REPOSITORYRULESETBYPASSACTOR_BYPASS_MODE
    switch v {
        case "always":
            result = ALWAYS_REPOSITORYRULESETBYPASSACTOR_BYPASS_MODE
        case "pull_request":
            result = PULL_REQUEST_REPOSITORYRULESETBYPASSACTOR_BYPASS_MODE
        default:
            return 0, errors.New("Unknown RepositoryRulesetBypassActor_bypass_mode value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRulesetBypassActor_bypass_mode(values []RepositoryRulesetBypassActor_bypass_mode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRulesetBypassActor_bypass_mode) isMultiValue() bool {
    return false
}
