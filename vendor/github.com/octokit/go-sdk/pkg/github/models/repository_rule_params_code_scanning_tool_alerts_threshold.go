package models
import (
    "errors"
)
// The severity level at which code scanning results that raise alerts block a reference update. For more information on alert severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
type RepositoryRuleParamsCodeScanningTool_alerts_threshold int

const (
    NONE_REPOSITORYRULEPARAMSCODESCANNINGTOOL_ALERTS_THRESHOLD RepositoryRuleParamsCodeScanningTool_alerts_threshold = iota
    ERRORS_REPOSITORYRULEPARAMSCODESCANNINGTOOL_ALERTS_THRESHOLD
    ERRORS_AND_WARNINGS_REPOSITORYRULEPARAMSCODESCANNINGTOOL_ALERTS_THRESHOLD
    ALL_REPOSITORYRULEPARAMSCODESCANNINGTOOL_ALERTS_THRESHOLD
)

func (i RepositoryRuleParamsCodeScanningTool_alerts_threshold) String() string {
    return []string{"none", "errors", "errors_and_warnings", "all"}[i]
}
func ParseRepositoryRuleParamsCodeScanningTool_alerts_threshold(v string) (any, error) {
    result := NONE_REPOSITORYRULEPARAMSCODESCANNINGTOOL_ALERTS_THRESHOLD
    switch v {
        case "none":
            result = NONE_REPOSITORYRULEPARAMSCODESCANNINGTOOL_ALERTS_THRESHOLD
        case "errors":
            result = ERRORS_REPOSITORYRULEPARAMSCODESCANNINGTOOL_ALERTS_THRESHOLD
        case "errors_and_warnings":
            result = ERRORS_AND_WARNINGS_REPOSITORYRULEPARAMSCODESCANNINGTOOL_ALERTS_THRESHOLD
        case "all":
            result = ALL_REPOSITORYRULEPARAMSCODESCANNINGTOOL_ALERTS_THRESHOLD
        default:
            return 0, errors.New("Unknown RepositoryRuleParamsCodeScanningTool_alerts_threshold value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleParamsCodeScanningTool_alerts_threshold(values []RepositoryRuleParamsCodeScanningTool_alerts_threshold) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleParamsCodeScanningTool_alerts_threshold) isMultiValue() bool {
    return false
}
