package models
import (
    "errors"
)
// The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
type EnabledRepositories int

const (
    ALL_ENABLEDREPOSITORIES EnabledRepositories = iota
    NONE_ENABLEDREPOSITORIES
    SELECTED_ENABLEDREPOSITORIES
)

func (i EnabledRepositories) String() string {
    return []string{"all", "none", "selected"}[i]
}
func ParseEnabledRepositories(v string) (any, error) {
    result := ALL_ENABLEDREPOSITORIES
    switch v {
        case "all":
            result = ALL_ENABLEDREPOSITORIES
        case "none":
            result = NONE_ENABLEDREPOSITORIES
        case "selected":
            result = SELECTED_ENABLEDREPOSITORIES
        default:
            return 0, errors.New("Unknown EnabledRepositories value: " + v)
    }
    return &result, nil
}
func SerializeEnabledRepositories(values []EnabledRepositories) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EnabledRepositories) isMultiValue() bool {
    return false
}
