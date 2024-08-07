package models
import (
    "errors"
)
// The permission associated with the invitation.
type RepositoryInvitation_permissions int

const (
    READ_REPOSITORYINVITATION_PERMISSIONS RepositoryInvitation_permissions = iota
    WRITE_REPOSITORYINVITATION_PERMISSIONS
    ADMIN_REPOSITORYINVITATION_PERMISSIONS
    TRIAGE_REPOSITORYINVITATION_PERMISSIONS
    MAINTAIN_REPOSITORYINVITATION_PERMISSIONS
)

func (i RepositoryInvitation_permissions) String() string {
    return []string{"read", "write", "admin", "triage", "maintain"}[i]
}
func ParseRepositoryInvitation_permissions(v string) (any, error) {
    result := READ_REPOSITORYINVITATION_PERMISSIONS
    switch v {
        case "read":
            result = READ_REPOSITORYINVITATION_PERMISSIONS
        case "write":
            result = WRITE_REPOSITORYINVITATION_PERMISSIONS
        case "admin":
            result = ADMIN_REPOSITORYINVITATION_PERMISSIONS
        case "triage":
            result = TRIAGE_REPOSITORYINVITATION_PERMISSIONS
        case "maintain":
            result = MAINTAIN_REPOSITORYINVITATION_PERMISSIONS
        default:
            return 0, errors.New("Unknown RepositoryInvitation_permissions value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryInvitation_permissions(values []RepositoryInvitation_permissions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryInvitation_permissions) isMultiValue() bool {
    return false
}
