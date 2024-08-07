package models
import (
    "errors"
)
// Visibility of a variable
type OrganizationActionsVariable_visibility int

const (
    ALL_ORGANIZATIONACTIONSVARIABLE_VISIBILITY OrganizationActionsVariable_visibility = iota
    PRIVATE_ORGANIZATIONACTIONSVARIABLE_VISIBILITY
    SELECTED_ORGANIZATIONACTIONSVARIABLE_VISIBILITY
)

func (i OrganizationActionsVariable_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseOrganizationActionsVariable_visibility(v string) (any, error) {
    result := ALL_ORGANIZATIONACTIONSVARIABLE_VISIBILITY
    switch v {
        case "all":
            result = ALL_ORGANIZATIONACTIONSVARIABLE_VISIBILITY
        case "private":
            result = PRIVATE_ORGANIZATIONACTIONSVARIABLE_VISIBILITY
        case "selected":
            result = SELECTED_ORGANIZATIONACTIONSVARIABLE_VISIBILITY
        default:
            return 0, errors.New("Unknown OrganizationActionsVariable_visibility value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationActionsVariable_visibility(values []OrganizationActionsVariable_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationActionsVariable_visibility) isMultiValue() bool {
    return false
}
