package personalaccesstokens
import (
    "errors"
)
// Action to apply to the fine-grained personal access token.
type PersonalAccessTokensPostRequestBody_action int

const (
    REVOKE_PERSONALACCESSTOKENSPOSTREQUESTBODY_ACTION PersonalAccessTokensPostRequestBody_action = iota
)

func (i PersonalAccessTokensPostRequestBody_action) String() string {
    return []string{"revoke"}[i]
}
func ParsePersonalAccessTokensPostRequestBody_action(v string) (any, error) {
    result := REVOKE_PERSONALACCESSTOKENSPOSTREQUESTBODY_ACTION
    switch v {
        case "revoke":
            result = REVOKE_PERSONALACCESSTOKENSPOSTREQUESTBODY_ACTION
        default:
            return 0, errors.New("Unknown PersonalAccessTokensPostRequestBody_action value: " + v)
    }
    return &result, nil
}
func SerializePersonalAccessTokensPostRequestBody_action(values []PersonalAccessTokensPostRequestBody_action) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PersonalAccessTokensPostRequestBody_action) isMultiValue() bool {
    return false
}
