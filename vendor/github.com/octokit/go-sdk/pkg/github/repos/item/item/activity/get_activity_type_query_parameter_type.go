package activity
import (
    "errors"
)
type GetActivity_typeQueryParameterType int

const (
    PUSH_GETACTIVITY_TYPEQUERYPARAMETERTYPE GetActivity_typeQueryParameterType = iota
    FORCE_PUSH_GETACTIVITY_TYPEQUERYPARAMETERTYPE
    BRANCH_CREATION_GETACTIVITY_TYPEQUERYPARAMETERTYPE
    BRANCH_DELETION_GETACTIVITY_TYPEQUERYPARAMETERTYPE
    PR_MERGE_GETACTIVITY_TYPEQUERYPARAMETERTYPE
    MERGE_QUEUE_MERGE_GETACTIVITY_TYPEQUERYPARAMETERTYPE
)

func (i GetActivity_typeQueryParameterType) String() string {
    return []string{"push", "force_push", "branch_creation", "branch_deletion", "pr_merge", "merge_queue_merge"}[i]
}
func ParseGetActivity_typeQueryParameterType(v string) (any, error) {
    result := PUSH_GETACTIVITY_TYPEQUERYPARAMETERTYPE
    switch v {
        case "push":
            result = PUSH_GETACTIVITY_TYPEQUERYPARAMETERTYPE
        case "force_push":
            result = FORCE_PUSH_GETACTIVITY_TYPEQUERYPARAMETERTYPE
        case "branch_creation":
            result = BRANCH_CREATION_GETACTIVITY_TYPEQUERYPARAMETERTYPE
        case "branch_deletion":
            result = BRANCH_DELETION_GETACTIVITY_TYPEQUERYPARAMETERTYPE
        case "pr_merge":
            result = PR_MERGE_GETACTIVITY_TYPEQUERYPARAMETERTYPE
        case "merge_queue_merge":
            result = MERGE_QUEUE_MERGE_GETACTIVITY_TYPEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetActivity_typeQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetActivity_typeQueryParameterType(values []GetActivity_typeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetActivity_typeQueryParameterType) isMultiValue() bool {
    return false
}
