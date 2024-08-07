package models
import (
    "errors"
)
type PagesHttpsCertificate_state int

const (
    NEW_PAGESHTTPSCERTIFICATE_STATE PagesHttpsCertificate_state = iota
    AUTHORIZATION_CREATED_PAGESHTTPSCERTIFICATE_STATE
    AUTHORIZATION_PENDING_PAGESHTTPSCERTIFICATE_STATE
    AUTHORIZED_PAGESHTTPSCERTIFICATE_STATE
    AUTHORIZATION_REVOKED_PAGESHTTPSCERTIFICATE_STATE
    ISSUED_PAGESHTTPSCERTIFICATE_STATE
    UPLOADED_PAGESHTTPSCERTIFICATE_STATE
    APPROVED_PAGESHTTPSCERTIFICATE_STATE
    ERRORED_PAGESHTTPSCERTIFICATE_STATE
    BAD_AUTHZ_PAGESHTTPSCERTIFICATE_STATE
    DESTROY_PENDING_PAGESHTTPSCERTIFICATE_STATE
    DNS_CHANGED_PAGESHTTPSCERTIFICATE_STATE
)

func (i PagesHttpsCertificate_state) String() string {
    return []string{"new", "authorization_created", "authorization_pending", "authorized", "authorization_revoked", "issued", "uploaded", "approved", "errored", "bad_authz", "destroy_pending", "dns_changed"}[i]
}
func ParsePagesHttpsCertificate_state(v string) (any, error) {
    result := NEW_PAGESHTTPSCERTIFICATE_STATE
    switch v {
        case "new":
            result = NEW_PAGESHTTPSCERTIFICATE_STATE
        case "authorization_created":
            result = AUTHORIZATION_CREATED_PAGESHTTPSCERTIFICATE_STATE
        case "authorization_pending":
            result = AUTHORIZATION_PENDING_PAGESHTTPSCERTIFICATE_STATE
        case "authorized":
            result = AUTHORIZED_PAGESHTTPSCERTIFICATE_STATE
        case "authorization_revoked":
            result = AUTHORIZATION_REVOKED_PAGESHTTPSCERTIFICATE_STATE
        case "issued":
            result = ISSUED_PAGESHTTPSCERTIFICATE_STATE
        case "uploaded":
            result = UPLOADED_PAGESHTTPSCERTIFICATE_STATE
        case "approved":
            result = APPROVED_PAGESHTTPSCERTIFICATE_STATE
        case "errored":
            result = ERRORED_PAGESHTTPSCERTIFICATE_STATE
        case "bad_authz":
            result = BAD_AUTHZ_PAGESHTTPSCERTIFICATE_STATE
        case "destroy_pending":
            result = DESTROY_PENDING_PAGESHTTPSCERTIFICATE_STATE
        case "dns_changed":
            result = DNS_CHANGED_PAGESHTTPSCERTIFICATE_STATE
        default:
            return 0, errors.New("Unknown PagesHttpsCertificate_state value: " + v)
    }
    return &result, nil
}
func SerializePagesHttpsCertificate_state(values []PagesHttpsCertificate_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PagesHttpsCertificate_state) isMultiValue() bool {
    return false
}
