package models
import (
    "errors"
)
// The side of the first line of the range for a multi-line comment.
type ReviewComment_start_side int

const (
    LEFT_REVIEWCOMMENT_START_SIDE ReviewComment_start_side = iota
    RIGHT_REVIEWCOMMENT_START_SIDE
)

func (i ReviewComment_start_side) String() string {
    return []string{"LEFT", "RIGHT"}[i]
}
func ParseReviewComment_start_side(v string) (any, error) {
    result := LEFT_REVIEWCOMMENT_START_SIDE
    switch v {
        case "LEFT":
            result = LEFT_REVIEWCOMMENT_START_SIDE
        case "RIGHT":
            result = RIGHT_REVIEWCOMMENT_START_SIDE
        default:
            return 0, errors.New("Unknown ReviewComment_start_side value: " + v)
    }
    return &result, nil
}
func SerializeReviewComment_start_side(values []ReviewComment_start_side) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReviewComment_start_side) isMultiValue() bool {
    return false
}
