package models
import (
    "errors"
)
// How the author is associated with the repository.
type AuthorAssociation int

const (
    COLLABORATOR_AUTHORASSOCIATION AuthorAssociation = iota
    CONTRIBUTOR_AUTHORASSOCIATION
    FIRST_TIMER_AUTHORASSOCIATION
    FIRST_TIME_CONTRIBUTOR_AUTHORASSOCIATION
    MANNEQUIN_AUTHORASSOCIATION
    MEMBER_AUTHORASSOCIATION
    NONE_AUTHORASSOCIATION
    OWNER_AUTHORASSOCIATION
)

func (i AuthorAssociation) String() string {
    return []string{"COLLABORATOR", "CONTRIBUTOR", "FIRST_TIMER", "FIRST_TIME_CONTRIBUTOR", "MANNEQUIN", "MEMBER", "NONE", "OWNER"}[i]
}
func ParseAuthorAssociation(v string) (any, error) {
    result := COLLABORATOR_AUTHORASSOCIATION
    switch v {
        case "COLLABORATOR":
            result = COLLABORATOR_AUTHORASSOCIATION
        case "CONTRIBUTOR":
            result = CONTRIBUTOR_AUTHORASSOCIATION
        case "FIRST_TIMER":
            result = FIRST_TIMER_AUTHORASSOCIATION
        case "FIRST_TIME_CONTRIBUTOR":
            result = FIRST_TIME_CONTRIBUTOR_AUTHORASSOCIATION
        case "MANNEQUIN":
            result = MANNEQUIN_AUTHORASSOCIATION
        case "MEMBER":
            result = MEMBER_AUTHORASSOCIATION
        case "NONE":
            result = NONE_AUTHORASSOCIATION
        case "OWNER":
            result = OWNER_AUTHORASSOCIATION
        default:
            return 0, errors.New("Unknown AuthorAssociation value: " + v)
    }
    return &result, nil
}
func SerializeAuthorAssociation(values []AuthorAssociation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthorAssociation) isMultiValue() bool {
    return false
}
