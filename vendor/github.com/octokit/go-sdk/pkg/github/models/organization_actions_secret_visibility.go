package models
import (
    "errors"
)
// Visibility of a secret
type OrganizationActionsSecret_visibility int

const (
    ALL_ORGANIZATIONACTIONSSECRET_VISIBILITY OrganizationActionsSecret_visibility = iota
    PRIVATE_ORGANIZATIONACTIONSSECRET_VISIBILITY
    SELECTED_ORGANIZATIONACTIONSSECRET_VISIBILITY
)

func (i OrganizationActionsSecret_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseOrganizationActionsSecret_visibility(v string) (any, error) {
    result := ALL_ORGANIZATIONACTIONSSECRET_VISIBILITY
    switch v {
        case "all":
            result = ALL_ORGANIZATIONACTIONSSECRET_VISIBILITY
        case "private":
            result = PRIVATE_ORGANIZATIONACTIONSSECRET_VISIBILITY
        case "selected":
            result = SELECTED_ORGANIZATIONACTIONSSECRET_VISIBILITY
        default:
            return 0, errors.New("Unknown OrganizationActionsSecret_visibility value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationActionsSecret_visibility(values []OrganizationActionsSecret_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationActionsSecret_visibility) isMultiValue() bool {
    return false
}
