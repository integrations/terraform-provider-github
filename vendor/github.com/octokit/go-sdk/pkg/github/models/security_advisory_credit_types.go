package models
import (
    "errors"
)
// The type of credit the user is receiving.
type SecurityAdvisoryCreditTypes int

const (
    ANALYST_SECURITYADVISORYCREDITTYPES SecurityAdvisoryCreditTypes = iota
    FINDER_SECURITYADVISORYCREDITTYPES
    REPORTER_SECURITYADVISORYCREDITTYPES
    COORDINATOR_SECURITYADVISORYCREDITTYPES
    REMEDIATION_DEVELOPER_SECURITYADVISORYCREDITTYPES
    REMEDIATION_REVIEWER_SECURITYADVISORYCREDITTYPES
    REMEDIATION_VERIFIER_SECURITYADVISORYCREDITTYPES
    TOOL_SECURITYADVISORYCREDITTYPES
    SPONSOR_SECURITYADVISORYCREDITTYPES
    OTHER_SECURITYADVISORYCREDITTYPES
)

func (i SecurityAdvisoryCreditTypes) String() string {
    return []string{"analyst", "finder", "reporter", "coordinator", "remediation_developer", "remediation_reviewer", "remediation_verifier", "tool", "sponsor", "other"}[i]
}
func ParseSecurityAdvisoryCreditTypes(v string) (any, error) {
    result := ANALYST_SECURITYADVISORYCREDITTYPES
    switch v {
        case "analyst":
            result = ANALYST_SECURITYADVISORYCREDITTYPES
        case "finder":
            result = FINDER_SECURITYADVISORYCREDITTYPES
        case "reporter":
            result = REPORTER_SECURITYADVISORYCREDITTYPES
        case "coordinator":
            result = COORDINATOR_SECURITYADVISORYCREDITTYPES
        case "remediation_developer":
            result = REMEDIATION_DEVELOPER_SECURITYADVISORYCREDITTYPES
        case "remediation_reviewer":
            result = REMEDIATION_REVIEWER_SECURITYADVISORYCREDITTYPES
        case "remediation_verifier":
            result = REMEDIATION_VERIFIER_SECURITYADVISORYCREDITTYPES
        case "tool":
            result = TOOL_SECURITYADVISORYCREDITTYPES
        case "sponsor":
            result = SPONSOR_SECURITYADVISORYCREDITTYPES
        case "other":
            result = OTHER_SECURITYADVISORYCREDITTYPES
        default:
            return 0, errors.New("Unknown SecurityAdvisoryCreditTypes value: " + v)
    }
    return &result, nil
}
func SerializeSecurityAdvisoryCreditTypes(values []SecurityAdvisoryCreditTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAdvisoryCreditTypes) isMultiValue() bool {
    return false
}
