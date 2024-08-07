package invitations
import (
    "errors"
)
// The role for the new member.  * `admin` - Organization owners with full administrative rights to the organization and complete access to all repositories and teams.   * `direct_member` - Non-owner organization members with ability to see other members and join teams by invitation.   * `billing_manager` - Non-owner organization members with ability to manage the billing settings of your organization.  * `reinstate` - The previous role assigned to the invitee before they were removed from your organization. Can be one of the roles listed above. Only works if the invitee was previously part of your organization.
type InvitationsPostRequestBody_role int

const (
    ADMIN_INVITATIONSPOSTREQUESTBODY_ROLE InvitationsPostRequestBody_role = iota
    DIRECT_MEMBER_INVITATIONSPOSTREQUESTBODY_ROLE
    BILLING_MANAGER_INVITATIONSPOSTREQUESTBODY_ROLE
    REINSTATE_INVITATIONSPOSTREQUESTBODY_ROLE
)

func (i InvitationsPostRequestBody_role) String() string {
    return []string{"admin", "direct_member", "billing_manager", "reinstate"}[i]
}
func ParseInvitationsPostRequestBody_role(v string) (any, error) {
    result := ADMIN_INVITATIONSPOSTREQUESTBODY_ROLE
    switch v {
        case "admin":
            result = ADMIN_INVITATIONSPOSTREQUESTBODY_ROLE
        case "direct_member":
            result = DIRECT_MEMBER_INVITATIONSPOSTREQUESTBODY_ROLE
        case "billing_manager":
            result = BILLING_MANAGER_INVITATIONSPOSTREQUESTBODY_ROLE
        case "reinstate":
            result = REINSTATE_INVITATIONSPOSTREQUESTBODY_ROLE
        default:
            return 0, errors.New("Unknown InvitationsPostRequestBody_role value: " + v)
    }
    return &result, nil
}
func SerializeInvitationsPostRequestBody_role(values []InvitationsPostRequestBody_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i InvitationsPostRequestBody_role) isMultiValue() bool {
    return false
}
