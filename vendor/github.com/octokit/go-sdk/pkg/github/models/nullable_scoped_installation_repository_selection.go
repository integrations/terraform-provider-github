package models
import (
    "errors"
)
// Describe whether all repositories have been selected or there's a selection involved
type NullableScopedInstallation_repository_selection int

const (
    ALL_NULLABLESCOPEDINSTALLATION_REPOSITORY_SELECTION NullableScopedInstallation_repository_selection = iota
    SELECTED_NULLABLESCOPEDINSTALLATION_REPOSITORY_SELECTION
)

func (i NullableScopedInstallation_repository_selection) String() string {
    return []string{"all", "selected"}[i]
}
func ParseNullableScopedInstallation_repository_selection(v string) (any, error) {
    result := ALL_NULLABLESCOPEDINSTALLATION_REPOSITORY_SELECTION
    switch v {
        case "all":
            result = ALL_NULLABLESCOPEDINSTALLATION_REPOSITORY_SELECTION
        case "selected":
            result = SELECTED_NULLABLESCOPEDINSTALLATION_REPOSITORY_SELECTION
        default:
            return 0, errors.New("Unknown NullableScopedInstallation_repository_selection value: " + v)
    }
    return &result, nil
}
func SerializeNullableScopedInstallation_repository_selection(values []NullableScopedInstallation_repository_selection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NullableScopedInstallation_repository_selection) isMultiValue() bool {
    return false
}
