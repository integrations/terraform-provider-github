package models
import (
    "errors"
)
// The organization policy for allowing or disallowing organization members to use Copilot Chat within their editor.
type CopilotOrganizationDetails_ide_chat int

const (
    ENABLED_COPILOTORGANIZATIONDETAILS_IDE_CHAT CopilotOrganizationDetails_ide_chat = iota
    DISABLED_COPILOTORGANIZATIONDETAILS_IDE_CHAT
    UNCONFIGURED_COPILOTORGANIZATIONDETAILS_IDE_CHAT
)

func (i CopilotOrganizationDetails_ide_chat) String() string {
    return []string{"enabled", "disabled", "unconfigured"}[i]
}
func ParseCopilotOrganizationDetails_ide_chat(v string) (any, error) {
    result := ENABLED_COPILOTORGANIZATIONDETAILS_IDE_CHAT
    switch v {
        case "enabled":
            result = ENABLED_COPILOTORGANIZATIONDETAILS_IDE_CHAT
        case "disabled":
            result = DISABLED_COPILOTORGANIZATIONDETAILS_IDE_CHAT
        case "unconfigured":
            result = UNCONFIGURED_COPILOTORGANIZATIONDETAILS_IDE_CHAT
        default:
            return 0, errors.New("Unknown CopilotOrganizationDetails_ide_chat value: " + v)
    }
    return &result, nil
}
func SerializeCopilotOrganizationDetails_ide_chat(values []CopilotOrganizationDetails_ide_chat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CopilotOrganizationDetails_ide_chat) isMultiValue() bool {
    return false
}
