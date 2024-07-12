package models
import (
    "errors"
)
// The token status as of the latest validity check.
type SecretScanningAlert_validity int

const (
    ACTIVE_SECRETSCANNINGALERT_VALIDITY SecretScanningAlert_validity = iota
    INACTIVE_SECRETSCANNINGALERT_VALIDITY
    UNKNOWN_SECRETSCANNINGALERT_VALIDITY
)

func (i SecretScanningAlert_validity) String() string {
    return []string{"active", "inactive", "unknown"}[i]
}
func ParseSecretScanningAlert_validity(v string) (any, error) {
    result := ACTIVE_SECRETSCANNINGALERT_VALIDITY
    switch v {
        case "active":
            result = ACTIVE_SECRETSCANNINGALERT_VALIDITY
        case "inactive":
            result = INACTIVE_SECRETSCANNINGALERT_VALIDITY
        case "unknown":
            result = UNKNOWN_SECRETSCANNINGALERT_VALIDITY
        default:
            return 0, errors.New("Unknown SecretScanningAlert_validity value: " + v)
    }
    return &result, nil
}
func SerializeSecretScanningAlert_validity(values []SecretScanningAlert_validity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecretScanningAlert_validity) isMultiValue() bool {
    return false
}
