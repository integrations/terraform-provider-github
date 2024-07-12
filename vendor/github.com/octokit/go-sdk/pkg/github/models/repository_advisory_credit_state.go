package models
import (
    "errors"
)
// The state of the user's acceptance of the credit.
type RepositoryAdvisoryCredit_state int

const (
    ACCEPTED_REPOSITORYADVISORYCREDIT_STATE RepositoryAdvisoryCredit_state = iota
    DECLINED_REPOSITORYADVISORYCREDIT_STATE
    PENDING_REPOSITORYADVISORYCREDIT_STATE
)

func (i RepositoryAdvisoryCredit_state) String() string {
    return []string{"accepted", "declined", "pending"}[i]
}
func ParseRepositoryAdvisoryCredit_state(v string) (any, error) {
    result := ACCEPTED_REPOSITORYADVISORYCREDIT_STATE
    switch v {
        case "accepted":
            result = ACCEPTED_REPOSITORYADVISORYCREDIT_STATE
        case "declined":
            result = DECLINED_REPOSITORYADVISORYCREDIT_STATE
        case "pending":
            result = PENDING_REPOSITORYADVISORYCREDIT_STATE
        default:
            return 0, errors.New("Unknown RepositoryAdvisoryCredit_state value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryAdvisoryCredit_state(values []RepositoryAdvisoryCredit_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryAdvisoryCredit_state) isMultiValue() bool {
    return false
}
