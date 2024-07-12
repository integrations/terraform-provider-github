package models
import (
    "errors"
)
// Type of repository selection requested.
type OrganizationProgrammaticAccessGrantRequest_repository_selection int

const (
    NONE_ORGANIZATIONPROGRAMMATICACCESSGRANTREQUEST_REPOSITORY_SELECTION OrganizationProgrammaticAccessGrantRequest_repository_selection = iota
    ALL_ORGANIZATIONPROGRAMMATICACCESSGRANTREQUEST_REPOSITORY_SELECTION
    SUBSET_ORGANIZATIONPROGRAMMATICACCESSGRANTREQUEST_REPOSITORY_SELECTION
)

func (i OrganizationProgrammaticAccessGrantRequest_repository_selection) String() string {
    return []string{"none", "all", "subset"}[i]
}
func ParseOrganizationProgrammaticAccessGrantRequest_repository_selection(v string) (any, error) {
    result := NONE_ORGANIZATIONPROGRAMMATICACCESSGRANTREQUEST_REPOSITORY_SELECTION
    switch v {
        case "none":
            result = NONE_ORGANIZATIONPROGRAMMATICACCESSGRANTREQUEST_REPOSITORY_SELECTION
        case "all":
            result = ALL_ORGANIZATIONPROGRAMMATICACCESSGRANTREQUEST_REPOSITORY_SELECTION
        case "subset":
            result = SUBSET_ORGANIZATIONPROGRAMMATICACCESSGRANTREQUEST_REPOSITORY_SELECTION
        default:
            return 0, errors.New("Unknown OrganizationProgrammaticAccessGrantRequest_repository_selection value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationProgrammaticAccessGrantRequest_repository_selection(values []OrganizationProgrammaticAccessGrantRequest_repository_selection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationProgrammaticAccessGrantRequest_repository_selection) isMultiValue() bool {
    return false
}
