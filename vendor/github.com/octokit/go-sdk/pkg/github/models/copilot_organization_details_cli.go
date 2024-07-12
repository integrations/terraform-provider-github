package models
import (
    "errors"
)
// The organization policy for allowing or disallowing organization members to use Copilot within their CLI.
type CopilotOrganizationDetails_cli int

const (
    ENABLED_COPILOTORGANIZATIONDETAILS_CLI CopilotOrganizationDetails_cli = iota
    DISABLED_COPILOTORGANIZATIONDETAILS_CLI
    UNCONFIGURED_COPILOTORGANIZATIONDETAILS_CLI
)

func (i CopilotOrganizationDetails_cli) String() string {
    return []string{"enabled", "disabled", "unconfigured"}[i]
}
func ParseCopilotOrganizationDetails_cli(v string) (any, error) {
    result := ENABLED_COPILOTORGANIZATIONDETAILS_CLI
    switch v {
        case "enabled":
            result = ENABLED_COPILOTORGANIZATIONDETAILS_CLI
        case "disabled":
            result = DISABLED_COPILOTORGANIZATIONDETAILS_CLI
        case "unconfigured":
            result = UNCONFIGURED_COPILOTORGANIZATIONDETAILS_CLI
        default:
            return 0, errors.New("Unknown CopilotOrganizationDetails_cli value: " + v)
    }
    return &result, nil
}
func SerializeCopilotOrganizationDetails_cli(values []CopilotOrganizationDetails_cli) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CopilotOrganizationDetails_cli) isMultiValue() bool {
    return false
}
