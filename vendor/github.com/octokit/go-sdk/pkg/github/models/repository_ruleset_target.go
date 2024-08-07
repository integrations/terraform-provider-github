package models
import (
    "errors"
)
// The target of the ruleset**Note**: The `push` target is in beta and is subject to change.
type RepositoryRuleset_target int

const (
    BRANCH_REPOSITORYRULESET_TARGET RepositoryRuleset_target = iota
    TAG_REPOSITORYRULESET_TARGET
    PUSH_REPOSITORYRULESET_TARGET
)

func (i RepositoryRuleset_target) String() string {
    return []string{"branch", "tag", "push"}[i]
}
func ParseRepositoryRuleset_target(v string) (any, error) {
    result := BRANCH_REPOSITORYRULESET_TARGET
    switch v {
        case "branch":
            result = BRANCH_REPOSITORYRULESET_TARGET
        case "tag":
            result = TAG_REPOSITORYRULESET_TARGET
        case "push":
            result = PUSH_REPOSITORYRULESET_TARGET
        default:
            return 0, errors.New("Unknown RepositoryRuleset_target value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleset_target(values []RepositoryRuleset_target) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleset_target) isMultiValue() bool {
    return false
}
