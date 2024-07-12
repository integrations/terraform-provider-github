package models
import (
    "errors"
)
// Whether it's a Group Assignment or Individual Assignment.
type SimpleClassroomAssignment_type int

const (
    INDIVIDUAL_SIMPLECLASSROOMASSIGNMENT_TYPE SimpleClassroomAssignment_type = iota
    GROUP_SIMPLECLASSROOMASSIGNMENT_TYPE
)

func (i SimpleClassroomAssignment_type) String() string {
    return []string{"individual", "group"}[i]
}
func ParseSimpleClassroomAssignment_type(v string) (any, error) {
    result := INDIVIDUAL_SIMPLECLASSROOMASSIGNMENT_TYPE
    switch v {
        case "individual":
            result = INDIVIDUAL_SIMPLECLASSROOMASSIGNMENT_TYPE
        case "group":
            result = GROUP_SIMPLECLASSROOMASSIGNMENT_TYPE
        default:
            return 0, errors.New("Unknown SimpleClassroomAssignment_type value: " + v)
    }
    return &result, nil
}
func SerializeSimpleClassroomAssignment_type(values []SimpleClassroomAssignment_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SimpleClassroomAssignment_type) isMultiValue() bool {
    return false
}
