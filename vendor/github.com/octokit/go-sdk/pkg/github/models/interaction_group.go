package models
import (
    "errors"
)
// The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
type InteractionGroup int

const (
    EXISTING_USERS_INTERACTIONGROUP InteractionGroup = iota
    CONTRIBUTORS_ONLY_INTERACTIONGROUP
    COLLABORATORS_ONLY_INTERACTIONGROUP
)

func (i InteractionGroup) String() string {
    return []string{"existing_users", "contributors_only", "collaborators_only"}[i]
}
func ParseInteractionGroup(v string) (any, error) {
    result := EXISTING_USERS_INTERACTIONGROUP
    switch v {
        case "existing_users":
            result = EXISTING_USERS_INTERACTIONGROUP
        case "contributors_only":
            result = CONTRIBUTORS_ONLY_INTERACTIONGROUP
        case "collaborators_only":
            result = COLLABORATORS_ONLY_INTERACTIONGROUP
        default:
            return 0, errors.New("Unknown InteractionGroup value: " + v)
    }
    return &result, nil
}
func SerializeInteractionGroup(values []InteractionGroup) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i InteractionGroup) isMultiValue() bool {
    return false
}
