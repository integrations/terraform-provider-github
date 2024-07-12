package models
import (
    "errors"
)
// The organization policy for allowing or disallowing organization members to use Copilot Chat within their editor.
type CopilotOrganizationDetails_copilot_chat int

const (
    ENABLED_COPILOTORGANIZATIONDETAILS_COPILOT_CHAT CopilotOrganizationDetails_copilot_chat = iota
    DISABLED_COPILOTORGANIZATIONDETAILS_COPILOT_CHAT
    UNCONFIGURED_COPILOTORGANIZATIONDETAILS_COPILOT_CHAT
)

func (i CopilotOrganizationDetails_copilot_chat) String() string {
    return []string{"enabled", "disabled", "unconfigured"}[i]
}
func ParseCopilotOrganizationDetails_copilot_chat(v string) (any, error) {
    result := ENABLED_COPILOTORGANIZATIONDETAILS_COPILOT_CHAT
    switch v {
        case "enabled":
            result = ENABLED_COPILOTORGANIZATIONDETAILS_COPILOT_CHAT
        case "disabled":
            result = DISABLED_COPILOTORGANIZATIONDETAILS_COPILOT_CHAT
        case "unconfigured":
            result = UNCONFIGURED_COPILOTORGANIZATIONDETAILS_COPILOT_CHAT
        default:
            return 0, errors.New("Unknown CopilotOrganizationDetails_copilot_chat value: " + v)
    }
    return &result, nil
}
func SerializeCopilotOrganizationDetails_copilot_chat(values []CopilotOrganizationDetails_copilot_chat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CopilotOrganizationDetails_copilot_chat) isMultiValue() bool {
    return false
}
