package personalaccesstokenrequests
import (
    "errors"
)
// Action to apply to the requests.
type PersonalAccessTokenRequestsPostRequestBody_action int

const (
    APPROVE_PERSONALACCESSTOKENREQUESTSPOSTREQUESTBODY_ACTION PersonalAccessTokenRequestsPostRequestBody_action = iota
    DENY_PERSONALACCESSTOKENREQUESTSPOSTREQUESTBODY_ACTION
)

func (i PersonalAccessTokenRequestsPostRequestBody_action) String() string {
    return []string{"approve", "deny"}[i]
}
func ParsePersonalAccessTokenRequestsPostRequestBody_action(v string) (any, error) {
    result := APPROVE_PERSONALACCESSTOKENREQUESTSPOSTREQUESTBODY_ACTION
    switch v {
        case "approve":
            result = APPROVE_PERSONALACCESSTOKENREQUESTSPOSTREQUESTBODY_ACTION
        case "deny":
            result = DENY_PERSONALACCESSTOKENREQUESTSPOSTREQUESTBODY_ACTION
        default:
            return 0, errors.New("Unknown PersonalAccessTokenRequestsPostRequestBody_action value: " + v)
    }
    return &result, nil
}
func SerializePersonalAccessTokenRequestsPostRequestBody_action(values []PersonalAccessTokenRequestsPostRequestBody_action) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PersonalAccessTokenRequestsPostRequestBody_action) isMultiValue() bool {
    return false
}
