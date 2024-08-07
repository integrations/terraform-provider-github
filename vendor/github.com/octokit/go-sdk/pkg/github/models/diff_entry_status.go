package models
import (
    "errors"
)
type DiffEntry_status int

const (
    ADDED_DIFFENTRY_STATUS DiffEntry_status = iota
    REMOVED_DIFFENTRY_STATUS
    MODIFIED_DIFFENTRY_STATUS
    RENAMED_DIFFENTRY_STATUS
    COPIED_DIFFENTRY_STATUS
    CHANGED_DIFFENTRY_STATUS
    UNCHANGED_DIFFENTRY_STATUS
)

func (i DiffEntry_status) String() string {
    return []string{"added", "removed", "modified", "renamed", "copied", "changed", "unchanged"}[i]
}
func ParseDiffEntry_status(v string) (any, error) {
    result := ADDED_DIFFENTRY_STATUS
    switch v {
        case "added":
            result = ADDED_DIFFENTRY_STATUS
        case "removed":
            result = REMOVED_DIFFENTRY_STATUS
        case "modified":
            result = MODIFIED_DIFFENTRY_STATUS
        case "renamed":
            result = RENAMED_DIFFENTRY_STATUS
        case "copied":
            result = COPIED_DIFFENTRY_STATUS
        case "changed":
            result = CHANGED_DIFFENTRY_STATUS
        case "unchanged":
            result = UNCHANGED_DIFFENTRY_STATUS
        default:
            return 0, errors.New("Unknown DiffEntry_status value: " + v)
    }
    return &result, nil
}
func SerializeDiffEntry_status(values []DiffEntry_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DiffEntry_status) isMultiValue() bool {
    return false
}
