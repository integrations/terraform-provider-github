package models
import (
    "errors"
)
// The state of the advisory.
type RepositoryAdvisory_state int

const (
    PUBLISHED_REPOSITORYADVISORY_STATE RepositoryAdvisory_state = iota
    CLOSED_REPOSITORYADVISORY_STATE
    WITHDRAWN_REPOSITORYADVISORY_STATE
    DRAFT_REPOSITORYADVISORY_STATE
    TRIAGE_REPOSITORYADVISORY_STATE
)

func (i RepositoryAdvisory_state) String() string {
    return []string{"published", "closed", "withdrawn", "draft", "triage"}[i]
}
func ParseRepositoryAdvisory_state(v string) (any, error) {
    result := PUBLISHED_REPOSITORYADVISORY_STATE
    switch v {
        case "published":
            result = PUBLISHED_REPOSITORYADVISORY_STATE
        case "closed":
            result = CLOSED_REPOSITORYADVISORY_STATE
        case "withdrawn":
            result = WITHDRAWN_REPOSITORYADVISORY_STATE
        case "draft":
            result = DRAFT_REPOSITORYADVISORY_STATE
        case "triage":
            result = TRIAGE_REPOSITORYADVISORY_STATE
        default:
            return 0, errors.New("Unknown RepositoryAdvisory_state value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryAdvisory_state(values []RepositoryAdvisory_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryAdvisory_state) isMultiValue() bool {
    return false
}
