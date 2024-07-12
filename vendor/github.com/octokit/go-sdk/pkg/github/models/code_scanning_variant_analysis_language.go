package models
import (
    "errors"
)
// The language targeted by the CodeQL query
type CodeScanningVariantAnalysisLanguage int

const (
    CPP_CODESCANNINGVARIANTANALYSISLANGUAGE CodeScanningVariantAnalysisLanguage = iota
    CSHARP_CODESCANNINGVARIANTANALYSISLANGUAGE
    GO_CODESCANNINGVARIANTANALYSISLANGUAGE
    JAVA_CODESCANNINGVARIANTANALYSISLANGUAGE
    JAVASCRIPT_CODESCANNINGVARIANTANALYSISLANGUAGE
    PYTHON_CODESCANNINGVARIANTANALYSISLANGUAGE
    RUBY_CODESCANNINGVARIANTANALYSISLANGUAGE
    SWIFT_CODESCANNINGVARIANTANALYSISLANGUAGE
)

func (i CodeScanningVariantAnalysisLanguage) String() string {
    return []string{"cpp", "csharp", "go", "java", "javascript", "python", "ruby", "swift"}[i]
}
func ParseCodeScanningVariantAnalysisLanguage(v string) (any, error) {
    result := CPP_CODESCANNINGVARIANTANALYSISLANGUAGE
    switch v {
        case "cpp":
            result = CPP_CODESCANNINGVARIANTANALYSISLANGUAGE
        case "csharp":
            result = CSHARP_CODESCANNINGVARIANTANALYSISLANGUAGE
        case "go":
            result = GO_CODESCANNINGVARIANTANALYSISLANGUAGE
        case "java":
            result = JAVA_CODESCANNINGVARIANTANALYSISLANGUAGE
        case "javascript":
            result = JAVASCRIPT_CODESCANNINGVARIANTANALYSISLANGUAGE
        case "python":
            result = PYTHON_CODESCANNINGVARIANTANALYSISLANGUAGE
        case "ruby":
            result = RUBY_CODESCANNINGVARIANTANALYSISLANGUAGE
        case "swift":
            result = SWIFT_CODESCANNINGVARIANTANALYSISLANGUAGE
        default:
            return 0, errors.New("Unknown CodeScanningVariantAnalysisLanguage value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningVariantAnalysisLanguage(values []CodeScanningVariantAnalysisLanguage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningVariantAnalysisLanguage) isMultiValue() bool {
    return false
}
