package models
import (
    "errors"
)
// The organization policy for allowing or disallowing organization members to use Copilot features within github.com.
type CopilotOrganizationDetails_platform_chat int

const (
    ENABLED_COPILOTORGANIZATIONDETAILS_PLATFORM_CHAT CopilotOrganizationDetails_platform_chat = iota
    DISABLED_COPILOTORGANIZATIONDETAILS_PLATFORM_CHAT
    UNCONFIGURED_COPILOTORGANIZATIONDETAILS_PLATFORM_CHAT
)

func (i CopilotOrganizationDetails_platform_chat) String() string {
    return []string{"enabled", "disabled", "unconfigured"}[i]
}
func ParseCopilotOrganizationDetails_platform_chat(v string) (any, error) {
    result := ENABLED_COPILOTORGANIZATIONDETAILS_PLATFORM_CHAT
    switch v {
        case "enabled":
            result = ENABLED_COPILOTORGANIZATIONDETAILS_PLATFORM_CHAT
        case "disabled":
            result = DISABLED_COPILOTORGANIZATIONDETAILS_PLATFORM_CHAT
        case "unconfigured":
            result = UNCONFIGURED_COPILOTORGANIZATIONDETAILS_PLATFORM_CHAT
        default:
            return 0, errors.New("Unknown CopilotOrganizationDetails_platform_chat value: " + v)
    }
    return &result, nil
}
func SerializeCopilotOrganizationDetails_platform_chat(values []CopilotOrganizationDetails_platform_chat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CopilotOrganizationDetails_platform_chat) isMultiValue() bool {
    return false
}
