package models
import (
    "errors"
)
// The merge method to use.
type AutoMerge_merge_method int

const (
    MERGE_AUTOMERGE_MERGE_METHOD AutoMerge_merge_method = iota
    SQUASH_AUTOMERGE_MERGE_METHOD
    REBASE_AUTOMERGE_MERGE_METHOD
)

func (i AutoMerge_merge_method) String() string {
    return []string{"merge", "squash", "rebase"}[i]
}
func ParseAutoMerge_merge_method(v string) (any, error) {
    result := MERGE_AUTOMERGE_MERGE_METHOD
    switch v {
        case "merge":
            result = MERGE_AUTOMERGE_MERGE_METHOD
        case "squash":
            result = SQUASH_AUTOMERGE_MERGE_METHOD
        case "rebase":
            result = REBASE_AUTOMERGE_MERGE_METHOD
        default:
            return 0, errors.New("Unknown AutoMerge_merge_method value: " + v)
    }
    return &result, nil
}
func SerializeAutoMerge_merge_method(values []AutoMerge_merge_method) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AutoMerge_merge_method) isMultiValue() bool {
    return false
}
