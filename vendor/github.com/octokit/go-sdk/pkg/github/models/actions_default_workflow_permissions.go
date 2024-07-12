package models
import (
    "errors"
)
// The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
type ActionsDefaultWorkflowPermissions int

const (
    READ_ACTIONSDEFAULTWORKFLOWPERMISSIONS ActionsDefaultWorkflowPermissions = iota
    WRITE_ACTIONSDEFAULTWORKFLOWPERMISSIONS
)

func (i ActionsDefaultWorkflowPermissions) String() string {
    return []string{"read", "write"}[i]
}
func ParseActionsDefaultWorkflowPermissions(v string) (any, error) {
    result := READ_ACTIONSDEFAULTWORKFLOWPERMISSIONS
    switch v {
        case "read":
            result = READ_ACTIONSDEFAULTWORKFLOWPERMISSIONS
        case "write":
            result = WRITE_ACTIONSDEFAULTWORKFLOWPERMISSIONS
        default:
            return 0, errors.New("Unknown ActionsDefaultWorkflowPermissions value: " + v)
    }
    return &result, nil
}
func SerializeActionsDefaultWorkflowPermissions(values []ActionsDefaultWorkflowPermissions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ActionsDefaultWorkflowPermissions) isMultiValue() bool {
    return false
}
