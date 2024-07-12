package models
import (
    "errors"
)
// Defines the level of access that workflows outside of the repository have to actions and reusable workflows within therepository.`none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repositories only. `organization` level access allows sharing across the organization.
type ActionsWorkflowAccessToRepository_access_level int

const (
    NONE_ACTIONSWORKFLOWACCESSTOREPOSITORY_ACCESS_LEVEL ActionsWorkflowAccessToRepository_access_level = iota
    USER_ACTIONSWORKFLOWACCESSTOREPOSITORY_ACCESS_LEVEL
    ORGANIZATION_ACTIONSWORKFLOWACCESSTOREPOSITORY_ACCESS_LEVEL
)

func (i ActionsWorkflowAccessToRepository_access_level) String() string {
    return []string{"none", "user", "organization"}[i]
}
func ParseActionsWorkflowAccessToRepository_access_level(v string) (any, error) {
    result := NONE_ACTIONSWORKFLOWACCESSTOREPOSITORY_ACCESS_LEVEL
    switch v {
        case "none":
            result = NONE_ACTIONSWORKFLOWACCESSTOREPOSITORY_ACCESS_LEVEL
        case "user":
            result = USER_ACTIONSWORKFLOWACCESSTOREPOSITORY_ACCESS_LEVEL
        case "organization":
            result = ORGANIZATION_ACTIONSWORKFLOWACCESSTOREPOSITORY_ACCESS_LEVEL
        default:
            return 0, errors.New("Unknown ActionsWorkflowAccessToRepository_access_level value: " + v)
    }
    return &result, nil
}
func SerializeActionsWorkflowAccessToRepository_access_level(values []ActionsWorkflowAccessToRepository_access_level) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ActionsWorkflowAccessToRepository_access_level) isMultiValue() bool {
    return false
}
