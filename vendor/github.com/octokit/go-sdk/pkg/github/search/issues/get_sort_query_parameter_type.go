package issues
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    COMMENTS_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    REACTIONS_GETSORTQUERYPARAMETERTYPE
    REACTIONS_PLUS_1_GETSORTQUERYPARAMETERTYPE
    REACTIONS1_GETSORTQUERYPARAMETERTYPE
    REACTIONSSMILE_GETSORTQUERYPARAMETERTYPE
    REACTIONSTHINKING_FACE_GETSORTQUERYPARAMETERTYPE
    REACTIONSHEART_GETSORTQUERYPARAMETERTYPE
    REACTIONSTADA_GETSORTQUERYPARAMETERTYPE
    INTERACTIONS_GETSORTQUERYPARAMETERTYPE
    CREATED_GETSORTQUERYPARAMETERTYPE
    UPDATED_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"comments", "reactions", "reactions-+1", "reactions--1", "reactions-smile", "reactions-thinking_face", "reactions-heart", "reactions-tada", "interactions", "created", "updated"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := COMMENTS_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "comments":
            result = COMMENTS_GETSORTQUERYPARAMETERTYPE
        case "reactions":
            result = REACTIONS_GETSORTQUERYPARAMETERTYPE
        case "reactions-+1":
            result = REACTIONS_PLUS_1_GETSORTQUERYPARAMETERTYPE
        case "reactions--1":
            result = REACTIONS1_GETSORTQUERYPARAMETERTYPE
        case "reactions-smile":
            result = REACTIONSSMILE_GETSORTQUERYPARAMETERTYPE
        case "reactions-thinking_face":
            result = REACTIONSTHINKING_FACE_GETSORTQUERYPARAMETERTYPE
        case "reactions-heart":
            result = REACTIONSHEART_GETSORTQUERYPARAMETERTYPE
        case "reactions-tada":
            result = REACTIONSTADA_GETSORTQUERYPARAMETERTYPE
        case "interactions":
            result = INTERACTIONS_GETSORTQUERYPARAMETERTYPE
        case "created":
            result = CREATED_GETSORTQUERYPARAMETERTYPE
        case "updated":
            result = UPDATED_GETSORTQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetSortQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetSortQueryParameterType(values []GetSortQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSortQueryParameterType) isMultiValue() bool {
    return false
}
