package models
import (
    "errors"
)
// The side of the first line of the range for a multi-line comment.
type ReviewComment_side int

const (
    LEFT_REVIEWCOMMENT_SIDE ReviewComment_side = iota
    RIGHT_REVIEWCOMMENT_SIDE
)

func (i ReviewComment_side) String() string {
    return []string{"LEFT", "RIGHT"}[i]
}
func ParseReviewComment_side(v string) (any, error) {
    result := LEFT_REVIEWCOMMENT_SIDE
    switch v {
        case "LEFT":
            result = LEFT_REVIEWCOMMENT_SIDE
        case "RIGHT":
            result = RIGHT_REVIEWCOMMENT_SIDE
        default:
            return 0, errors.New("Unknown ReviewComment_side value: " + v)
    }
    return &result, nil
}
func SerializeReviewComment_side(values []ReviewComment_side) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReviewComment_side) isMultiValue() bool {
    return false
}
