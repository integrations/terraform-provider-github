package models
import (
    "errors"
)
type CodeScanningDefaultSetup_languages int

const (
    CCPP_CODESCANNINGDEFAULTSETUP_LANGUAGES CodeScanningDefaultSetup_languages = iota
    CSHARP_CODESCANNINGDEFAULTSETUP_LANGUAGES
    GO_CODESCANNINGDEFAULTSETUP_LANGUAGES
    JAVAKOTLIN_CODESCANNINGDEFAULTSETUP_LANGUAGES
    JAVASCRIPTTYPESCRIPT_CODESCANNINGDEFAULTSETUP_LANGUAGES
    JAVASCRIPT_CODESCANNINGDEFAULTSETUP_LANGUAGES
    PYTHON_CODESCANNINGDEFAULTSETUP_LANGUAGES
    RUBY_CODESCANNINGDEFAULTSETUP_LANGUAGES
    TYPESCRIPT_CODESCANNINGDEFAULTSETUP_LANGUAGES
    SWIFT_CODESCANNINGDEFAULTSETUP_LANGUAGES
)

func (i CodeScanningDefaultSetup_languages) String() string {
    return []string{"c-cpp", "csharp", "go", "java-kotlin", "javascript-typescript", "javascript", "python", "ruby", "typescript", "swift"}[i]
}
func ParseCodeScanningDefaultSetup_languages(v string) (any, error) {
    result := CCPP_CODESCANNINGDEFAULTSETUP_LANGUAGES
    switch v {
        case "c-cpp":
            result = CCPP_CODESCANNINGDEFAULTSETUP_LANGUAGES
        case "csharp":
            result = CSHARP_CODESCANNINGDEFAULTSETUP_LANGUAGES
        case "go":
            result = GO_CODESCANNINGDEFAULTSETUP_LANGUAGES
        case "java-kotlin":
            result = JAVAKOTLIN_CODESCANNINGDEFAULTSETUP_LANGUAGES
        case "javascript-typescript":
            result = JAVASCRIPTTYPESCRIPT_CODESCANNINGDEFAULTSETUP_LANGUAGES
        case "javascript":
            result = JAVASCRIPT_CODESCANNINGDEFAULTSETUP_LANGUAGES
        case "python":
            result = PYTHON_CODESCANNINGDEFAULTSETUP_LANGUAGES
        case "ruby":
            result = RUBY_CODESCANNINGDEFAULTSETUP_LANGUAGES
        case "typescript":
            result = TYPESCRIPT_CODESCANNINGDEFAULTSETUP_LANGUAGES
        case "swift":
            result = SWIFT_CODESCANNINGDEFAULTSETUP_LANGUAGES
        default:
            return 0, errors.New("Unknown CodeScanningDefaultSetup_languages value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetup_languages(values []CodeScanningDefaultSetup_languages) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetup_languages) isMultiValue() bool {
    return false
}
