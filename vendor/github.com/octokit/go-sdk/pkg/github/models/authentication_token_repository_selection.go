package models
import (
    "errors"
)
// Describe whether all repositories have been selected or there's a selection involved
type AuthenticationToken_repository_selection int

const (
    ALL_AUTHENTICATIONTOKEN_REPOSITORY_SELECTION AuthenticationToken_repository_selection = iota
    SELECTED_AUTHENTICATIONTOKEN_REPOSITORY_SELECTION
)

func (i AuthenticationToken_repository_selection) String() string {
    return []string{"all", "selected"}[i]
}
func ParseAuthenticationToken_repository_selection(v string) (any, error) {
    result := ALL_AUTHENTICATIONTOKEN_REPOSITORY_SELECTION
    switch v {
        case "all":
            result = ALL_AUTHENTICATIONTOKEN_REPOSITORY_SELECTION
        case "selected":
            result = SELECTED_AUTHENTICATIONTOKEN_REPOSITORY_SELECTION
        default:
            return 0, errors.New("Unknown AuthenticationToken_repository_selection value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationToken_repository_selection(values []AuthenticationToken_repository_selection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthenticationToken_repository_selection) isMultiValue() bool {
    return false
}
