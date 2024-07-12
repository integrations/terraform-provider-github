package models
import (
    "errors"
)
type InstallationToken_repository_selection int

const (
    ALL_INSTALLATIONTOKEN_REPOSITORY_SELECTION InstallationToken_repository_selection = iota
    SELECTED_INSTALLATIONTOKEN_REPOSITORY_SELECTION
)

func (i InstallationToken_repository_selection) String() string {
    return []string{"all", "selected"}[i]
}
func ParseInstallationToken_repository_selection(v string) (any, error) {
    result := ALL_INSTALLATIONTOKEN_REPOSITORY_SELECTION
    switch v {
        case "all":
            result = ALL_INSTALLATIONTOKEN_REPOSITORY_SELECTION
        case "selected":
            result = SELECTED_INSTALLATIONTOKEN_REPOSITORY_SELECTION
        default:
            return 0, errors.New("Unknown InstallationToken_repository_selection value: " + v)
    }
    return &result, nil
}
func SerializeInstallationToken_repository_selection(values []InstallationToken_repository_selection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i InstallationToken_repository_selection) isMultiValue() bool {
    return false
}
