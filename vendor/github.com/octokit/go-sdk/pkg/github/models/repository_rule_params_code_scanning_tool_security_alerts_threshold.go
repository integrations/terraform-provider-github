package models
import (
    "errors"
)
// The severity level at which code scanning results that raise security alerts block a reference update. For more information on security severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
type RepositoryRuleParamsCodeScanningTool_security_alerts_threshold int

const (
    NONE_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD RepositoryRuleParamsCodeScanningTool_security_alerts_threshold = iota
    CRITICAL_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
    HIGH_OR_HIGHER_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
    MEDIUM_OR_HIGHER_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
    ALL_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
)

func (i RepositoryRuleParamsCodeScanningTool_security_alerts_threshold) String() string {
    return []string{"none", "critical", "high_or_higher", "medium_or_higher", "all"}[i]
}
func ParseRepositoryRuleParamsCodeScanningTool_security_alerts_threshold(v string) (any, error) {
    result := NONE_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
    switch v {
        case "none":
            result = NONE_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
        case "critical":
            result = CRITICAL_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
        case "high_or_higher":
            result = HIGH_OR_HIGHER_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
        case "medium_or_higher":
            result = MEDIUM_OR_HIGHER_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
        case "all":
            result = ALL_REPOSITORYRULEPARAMSCODESCANNINGTOOL_SECURITY_ALERTS_THRESHOLD
        default:
            return 0, errors.New("Unknown RepositoryRuleParamsCodeScanningTool_security_alerts_threshold value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleParamsCodeScanningTool_security_alerts_threshold(values []RepositoryRuleParamsCodeScanningTool_security_alerts_threshold) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleParamsCodeScanningTool_security_alerts_threshold) isMultiValue() bool {
    return false
}
