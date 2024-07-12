package models
import (
    "errors"
)
// Type of repository selection requested.
type OrganizationProgrammaticAccessGrant_repository_selection int

const (
    NONE_ORGANIZATIONPROGRAMMATICACCESSGRANT_REPOSITORY_SELECTION OrganizationProgrammaticAccessGrant_repository_selection = iota
    ALL_ORGANIZATIONPROGRAMMATICACCESSGRANT_REPOSITORY_SELECTION
    SUBSET_ORGANIZATIONPROGRAMMATICACCESSGRANT_REPOSITORY_SELECTION
)

func (i OrganizationProgrammaticAccessGrant_repository_selection) String() string {
    return []string{"none", "all", "subset"}[i]
}
func ParseOrganizationProgrammaticAccessGrant_repository_selection(v string) (any, error) {
    result := NONE_ORGANIZATIONPROGRAMMATICACCESSGRANT_REPOSITORY_SELECTION
    switch v {
        case "none":
            result = NONE_ORGANIZATIONPROGRAMMATICACCESSGRANT_REPOSITORY_SELECTION
        case "all":
            result = ALL_ORGANIZATIONPROGRAMMATICACCESSGRANT_REPOSITORY_SELECTION
        case "subset":
            result = SUBSET_ORGANIZATIONPROGRAMMATICACCESSGRANT_REPOSITORY_SELECTION
        default:
            return 0, errors.New("Unknown OrganizationProgrammaticAccessGrant_repository_selection value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationProgrammaticAccessGrant_repository_selection(values []OrganizationProgrammaticAccessGrant_repository_selection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationProgrammaticAccessGrant_repository_selection) isMultiValue() bool {
    return false
}
