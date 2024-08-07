package models
import (
    "errors"
)
// Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
type NullableCodespaceMachine_prebuild_availability int

const (
    NONE_NULLABLECODESPACEMACHINE_PREBUILD_AVAILABILITY NullableCodespaceMachine_prebuild_availability = iota
    READY_NULLABLECODESPACEMACHINE_PREBUILD_AVAILABILITY
    IN_PROGRESS_NULLABLECODESPACEMACHINE_PREBUILD_AVAILABILITY
)

func (i NullableCodespaceMachine_prebuild_availability) String() string {
    return []string{"none", "ready", "in_progress"}[i]
}
func ParseNullableCodespaceMachine_prebuild_availability(v string) (any, error) {
    result := NONE_NULLABLECODESPACEMACHINE_PREBUILD_AVAILABILITY
    switch v {
        case "none":
            result = NONE_NULLABLECODESPACEMACHINE_PREBUILD_AVAILABILITY
        case "ready":
            result = READY_NULLABLECODESPACEMACHINE_PREBUILD_AVAILABILITY
        case "in_progress":
            result = IN_PROGRESS_NULLABLECODESPACEMACHINE_PREBUILD_AVAILABILITY
        default:
            return 0, errors.New("Unknown NullableCodespaceMachine_prebuild_availability value: " + v)
    }
    return &result, nil
}
func SerializeNullableCodespaceMachine_prebuild_availability(values []NullableCodespaceMachine_prebuild_availability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NullableCodespaceMachine_prebuild_availability) isMultiValue() bool {
    return false
}
