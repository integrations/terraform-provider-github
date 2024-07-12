package models
import (
    "errors"
)
// The level of permission to grant the access token to update GitHub Actions workflow files.
type AppPermissions_workflows int

const (
    WRITE_APPPERMISSIONS_WORKFLOWS AppPermissions_workflows = iota
)

func (i AppPermissions_workflows) String() string {
    return []string{"write"}[i]
}
func ParseAppPermissions_workflows(v string) (any, error) {
    result := WRITE_APPPERMISSIONS_WORKFLOWS
    switch v {
        case "write":
            result = WRITE_APPPERMISSIONS_WORKFLOWS
        default:
            return 0, errors.New("Unknown AppPermissions_workflows value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_workflows(values []AppPermissions_workflows) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_workflows) isMultiValue() bool {
    return false
}
