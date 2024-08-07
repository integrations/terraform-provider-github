package models
import (
    "errors"
)
// Whether it's a group assignment or individual assignment.
type ClassroomAssignment_type int

const (
    INDIVIDUAL_CLASSROOMASSIGNMENT_TYPE ClassroomAssignment_type = iota
    GROUP_CLASSROOMASSIGNMENT_TYPE
)

func (i ClassroomAssignment_type) String() string {
    return []string{"individual", "group"}[i]
}
func ParseClassroomAssignment_type(v string) (any, error) {
    result := INDIVIDUAL_CLASSROOMASSIGNMENT_TYPE
    switch v {
        case "individual":
            result = INDIVIDUAL_CLASSROOMASSIGNMENT_TYPE
        case "group":
            result = GROUP_CLASSROOMASSIGNMENT_TYPE
        default:
            return 0, errors.New("Unknown ClassroomAssignment_type value: " + v)
    }
    return &result, nil
}
func SerializeClassroomAssignment_type(values []ClassroomAssignment_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClassroomAssignment_type) isMultiValue() bool {
    return false
}
