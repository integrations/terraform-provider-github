package models
import (
    "errors"
)
// Describe whether all repositories have been selected or there's a selection involved
type Installation_repository_selection int

const (
    ALL_INSTALLATION_REPOSITORY_SELECTION Installation_repository_selection = iota
    SELECTED_INSTALLATION_REPOSITORY_SELECTION
)

func (i Installation_repository_selection) String() string {
    return []string{"all", "selected"}[i]
}
func ParseInstallation_repository_selection(v string) (any, error) {
    result := ALL_INSTALLATION_REPOSITORY_SELECTION
    switch v {
        case "all":
            result = ALL_INSTALLATION_REPOSITORY_SELECTION
        case "selected":
            result = SELECTED_INSTALLATION_REPOSITORY_SELECTION
        default:
            return 0, errors.New("Unknown Installation_repository_selection value: " + v)
    }
    return &result, nil
}
func SerializeInstallation_repository_selection(values []Installation_repository_selection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Installation_repository_selection) isMultiValue() bool {
    return false
}
