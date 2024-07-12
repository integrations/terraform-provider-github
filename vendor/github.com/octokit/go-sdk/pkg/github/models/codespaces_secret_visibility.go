package models
import (
    "errors"
)
// The type of repositories in the organization that the secret is visible to
type CodespacesSecret_visibility int

const (
    ALL_CODESPACESSECRET_VISIBILITY CodespacesSecret_visibility = iota
    PRIVATE_CODESPACESSECRET_VISIBILITY
    SELECTED_CODESPACESSECRET_VISIBILITY
)

func (i CodespacesSecret_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseCodespacesSecret_visibility(v string) (any, error) {
    result := ALL_CODESPACESSECRET_VISIBILITY
    switch v {
        case "all":
            result = ALL_CODESPACESSECRET_VISIBILITY
        case "private":
            result = PRIVATE_CODESPACESSECRET_VISIBILITY
        case "selected":
            result = SELECTED_CODESPACESSECRET_VISIBILITY
        default:
            return 0, errors.New("Unknown CodespacesSecret_visibility value: " + v)
    }
    return &result, nil
}
func SerializeCodespacesSecret_visibility(values []CodespacesSecret_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodespacesSecret_visibility) isMultiValue() bool {
    return false
}
