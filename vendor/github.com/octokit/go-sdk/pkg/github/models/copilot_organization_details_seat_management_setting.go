package models
import (
    "errors"
)
// The mode of assigning new seats.
type CopilotOrganizationDetails_seat_management_setting int

const (
    ASSIGN_ALL_COPILOTORGANIZATIONDETAILS_SEAT_MANAGEMENT_SETTING CopilotOrganizationDetails_seat_management_setting = iota
    ASSIGN_SELECTED_COPILOTORGANIZATIONDETAILS_SEAT_MANAGEMENT_SETTING
    DISABLED_COPILOTORGANIZATIONDETAILS_SEAT_MANAGEMENT_SETTING
    UNCONFIGURED_COPILOTORGANIZATIONDETAILS_SEAT_MANAGEMENT_SETTING
)

func (i CopilotOrganizationDetails_seat_management_setting) String() string {
    return []string{"assign_all", "assign_selected", "disabled", "unconfigured"}[i]
}
func ParseCopilotOrganizationDetails_seat_management_setting(v string) (any, error) {
    result := ASSIGN_ALL_COPILOTORGANIZATIONDETAILS_SEAT_MANAGEMENT_SETTING
    switch v {
        case "assign_all":
            result = ASSIGN_ALL_COPILOTORGANIZATIONDETAILS_SEAT_MANAGEMENT_SETTING
        case "assign_selected":
            result = ASSIGN_SELECTED_COPILOTORGANIZATIONDETAILS_SEAT_MANAGEMENT_SETTING
        case "disabled":
            result = DISABLED_COPILOTORGANIZATIONDETAILS_SEAT_MANAGEMENT_SETTING
        case "unconfigured":
            result = UNCONFIGURED_COPILOTORGANIZATIONDETAILS_SEAT_MANAGEMENT_SETTING
        default:
            return 0, errors.New("Unknown CopilotOrganizationDetails_seat_management_setting value: " + v)
    }
    return &result, nil
}
func SerializeCopilotOrganizationDetails_seat_management_setting(values []CopilotOrganizationDetails_seat_management_setting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CopilotOrganizationDetails_seat_management_setting) isMultiValue() bool {
    return false
}
