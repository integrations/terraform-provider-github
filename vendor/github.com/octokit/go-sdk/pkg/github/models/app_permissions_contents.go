package models
import (
    "errors"
)
// The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
type AppPermissions_contents int

const (
    READ_APPPERMISSIONS_CONTENTS AppPermissions_contents = iota
    WRITE_APPPERMISSIONS_CONTENTS
)

func (i AppPermissions_contents) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_contents(v string) (any, error) {
    result := READ_APPPERMISSIONS_CONTENTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_CONTENTS
        case "write":
            result = WRITE_APPPERMISSIONS_CONTENTS
        default:
            return 0, errors.New("Unknown AppPermissions_contents value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_contents(values []AppPermissions_contents) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_contents) isMultiValue() bool {
    return false
}
