package models
import (
    "errors"
)
// The type of the activity that was performed.
type Activity_activity_type int

const (
    PUSH_ACTIVITY_ACTIVITY_TYPE Activity_activity_type = iota
    FORCE_PUSH_ACTIVITY_ACTIVITY_TYPE
    BRANCH_DELETION_ACTIVITY_ACTIVITY_TYPE
    BRANCH_CREATION_ACTIVITY_ACTIVITY_TYPE
    PR_MERGE_ACTIVITY_ACTIVITY_TYPE
    MERGE_QUEUE_MERGE_ACTIVITY_ACTIVITY_TYPE
)

func (i Activity_activity_type) String() string {
    return []string{"push", "force_push", "branch_deletion", "branch_creation", "pr_merge", "merge_queue_merge"}[i]
}
func ParseActivity_activity_type(v string) (any, error) {
    result := PUSH_ACTIVITY_ACTIVITY_TYPE
    switch v {
        case "push":
            result = PUSH_ACTIVITY_ACTIVITY_TYPE
        case "force_push":
            result = FORCE_PUSH_ACTIVITY_ACTIVITY_TYPE
        case "branch_deletion":
            result = BRANCH_DELETION_ACTIVITY_ACTIVITY_TYPE
        case "branch_creation":
            result = BRANCH_CREATION_ACTIVITY_ACTIVITY_TYPE
        case "pr_merge":
            result = PR_MERGE_ACTIVITY_ACTIVITY_TYPE
        case "merge_queue_merge":
            result = MERGE_QUEUE_MERGE_ACTIVITY_ACTIVITY_TYPE
        default:
            return 0, errors.New("Unknown Activity_activity_type value: " + v)
    }
    return &result, nil
}
func SerializeActivity_activity_type(values []Activity_activity_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Activity_activity_type) isMultiValue() bool {
    return false
}
