package models
import (
    "errors"
)
// **Required when the `state` is `resolved`.** The reason for resolving the alert.
type SecretScanningAlertResolution int

const (
    FALSE_POSITIVE_SECRETSCANNINGALERTRESOLUTION SecretScanningAlertResolution = iota
    WONT_FIX_SECRETSCANNINGALERTRESOLUTION
    REVOKED_SECRETSCANNINGALERTRESOLUTION
    USED_IN_TESTS_SECRETSCANNINGALERTRESOLUTION
)

func (i SecretScanningAlertResolution) String() string {
    return []string{"false_positive", "wont_fix", "revoked", "used_in_tests"}[i]
}
func ParseSecretScanningAlertResolution(v string) (any, error) {
    result := FALSE_POSITIVE_SECRETSCANNINGALERTRESOLUTION
    switch v {
        case "false_positive":
            result = FALSE_POSITIVE_SECRETSCANNINGALERTRESOLUTION
        case "wont_fix":
            result = WONT_FIX_SECRETSCANNINGALERTRESOLUTION
        case "revoked":
            result = REVOKED_SECRETSCANNINGALERTRESOLUTION
        case "used_in_tests":
            result = USED_IN_TESTS_SECRETSCANNINGALERTRESOLUTION
        default:
            return 0, errors.New("Unknown SecretScanningAlertResolution value: " + v)
    }
    return &result, nil
}
func SerializeSecretScanningAlertResolution(values []SecretScanningAlertResolution) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecretScanningAlertResolution) isMultiValue() bool {
    return false
}
