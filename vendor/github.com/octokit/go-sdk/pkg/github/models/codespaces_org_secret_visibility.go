package models
import (
    "errors"
)
// The type of repositories in the organization that the secret is visible to
type CodespacesOrgSecret_visibility int

const (
    ALL_CODESPACESORGSECRET_VISIBILITY CodespacesOrgSecret_visibility = iota
    PRIVATE_CODESPACESORGSECRET_VISIBILITY
    SELECTED_CODESPACESORGSECRET_VISIBILITY
)

func (i CodespacesOrgSecret_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseCodespacesOrgSecret_visibility(v string) (any, error) {
    result := ALL_CODESPACESORGSECRET_VISIBILITY
    switch v {
        case "all":
            result = ALL_CODESPACESORGSECRET_VISIBILITY
        case "private":
            result = PRIVATE_CODESPACESORGSECRET_VISIBILITY
        case "selected":
            result = SELECTED_CODESPACESORGSECRET_VISIBILITY
        default:
            return 0, errors.New("Unknown CodespacesOrgSecret_visibility value: " + v)
    }
    return &result, nil
}
func SerializeCodespacesOrgSecret_visibility(values []CodespacesOrgSecret_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodespacesOrgSecret_visibility) isMultiValue() bool {
    return false
}
