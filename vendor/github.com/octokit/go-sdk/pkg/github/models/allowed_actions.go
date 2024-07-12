package models
import (
    "errors"
)
// The permissions policy that controls the actions and reusable workflows that are allowed to run.
type AllowedActions int

const (
    ALL_ALLOWEDACTIONS AllowedActions = iota
    LOCAL_ONLY_ALLOWEDACTIONS
    SELECTED_ALLOWEDACTIONS
)

func (i AllowedActions) String() string {
    return []string{"all", "local_only", "selected"}[i]
}
func ParseAllowedActions(v string) (any, error) {
    result := ALL_ALLOWEDACTIONS
    switch v {
        case "all":
            result = ALL_ALLOWEDACTIONS
        case "local_only":
            result = LOCAL_ONLY_ALLOWEDACTIONS
        case "selected":
            result = SELECTED_ALLOWEDACTIONS
        default:
            return 0, errors.New("Unknown AllowedActions value: " + v)
    }
    return &result, nil
}
func SerializeAllowedActions(values []AllowedActions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AllowedActions) isMultiValue() bool {
    return false
}
