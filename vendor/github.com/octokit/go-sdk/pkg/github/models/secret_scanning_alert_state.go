package models
import (
    "errors"
)
// Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
type SecretScanningAlertState int

const (
    OPEN_SECRETSCANNINGALERTSTATE SecretScanningAlertState = iota
    RESOLVED_SECRETSCANNINGALERTSTATE
)

func (i SecretScanningAlertState) String() string {
    return []string{"open", "resolved"}[i]
}
func ParseSecretScanningAlertState(v string) (any, error) {
    result := OPEN_SECRETSCANNINGALERTSTATE
    switch v {
        case "open":
            result = OPEN_SECRETSCANNINGALERTSTATE
        case "resolved":
            result = RESOLVED_SECRETSCANNINGALERTSTATE
        default:
            return 0, errors.New("Unknown SecretScanningAlertState value: " + v)
    }
    return &result, nil
}
func SerializeSecretScanningAlertState(values []SecretScanningAlertState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecretScanningAlertState) isMultiValue() bool {
    return false
}
