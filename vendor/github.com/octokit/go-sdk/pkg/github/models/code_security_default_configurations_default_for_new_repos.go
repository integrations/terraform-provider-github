package models
import (
    "errors"
)
// The visibility of newly created repositories for which the code security configuration will be applied to by default
type CodeSecurityDefaultConfigurations_default_for_new_repos int

const (
    PUBLIC_CODESECURITYDEFAULTCONFIGURATIONS_DEFAULT_FOR_NEW_REPOS CodeSecurityDefaultConfigurations_default_for_new_repos = iota
    PRIVATE_AND_INTERNAL_CODESECURITYDEFAULTCONFIGURATIONS_DEFAULT_FOR_NEW_REPOS
    ALL_CODESECURITYDEFAULTCONFIGURATIONS_DEFAULT_FOR_NEW_REPOS
)

func (i CodeSecurityDefaultConfigurations_default_for_new_repos) String() string {
    return []string{"public", "private_and_internal", "all"}[i]
}
func ParseCodeSecurityDefaultConfigurations_default_for_new_repos(v string) (any, error) {
    result := PUBLIC_CODESECURITYDEFAULTCONFIGURATIONS_DEFAULT_FOR_NEW_REPOS
    switch v {
        case "public":
            result = PUBLIC_CODESECURITYDEFAULTCONFIGURATIONS_DEFAULT_FOR_NEW_REPOS
        case "private_and_internal":
            result = PRIVATE_AND_INTERNAL_CODESECURITYDEFAULTCONFIGURATIONS_DEFAULT_FOR_NEW_REPOS
        case "all":
            result = ALL_CODESECURITYDEFAULTCONFIGURATIONS_DEFAULT_FOR_NEW_REPOS
        default:
            return 0, errors.New("Unknown CodeSecurityDefaultConfigurations_default_for_new_repos value: " + v)
    }
    return &result, nil
}
func SerializeCodeSecurityDefaultConfigurations_default_for_new_repos(values []CodeSecurityDefaultConfigurations_default_for_new_repos) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityDefaultConfigurations_default_for_new_repos) isMultiValue() bool {
    return false
}
