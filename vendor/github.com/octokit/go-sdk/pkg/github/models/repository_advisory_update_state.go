package models
import (
    "errors"
)
// The state of the advisory.
type RepositoryAdvisoryUpdate_state int

const (
    PUBLISHED_REPOSITORYADVISORYUPDATE_STATE RepositoryAdvisoryUpdate_state = iota
    CLOSED_REPOSITORYADVISORYUPDATE_STATE
    DRAFT_REPOSITORYADVISORYUPDATE_STATE
)

func (i RepositoryAdvisoryUpdate_state) String() string {
    return []string{"published", "closed", "draft"}[i]
}
func ParseRepositoryAdvisoryUpdate_state(v string) (any, error) {
    result := PUBLISHED_REPOSITORYADVISORYUPDATE_STATE
    switch v {
        case "published":
            result = PUBLISHED_REPOSITORYADVISORYUPDATE_STATE
        case "closed":
            result = CLOSED_REPOSITORYADVISORYUPDATE_STATE
        case "draft":
            result = DRAFT_REPOSITORYADVISORYUPDATE_STATE
        default:
            return 0, errors.New("Unknown RepositoryAdvisoryUpdate_state value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryAdvisoryUpdate_state(values []RepositoryAdvisoryUpdate_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryAdvisoryUpdate_state) isMultiValue() bool {
    return false
}
