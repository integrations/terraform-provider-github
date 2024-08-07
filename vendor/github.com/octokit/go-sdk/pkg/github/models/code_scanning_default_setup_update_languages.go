package models
import (
    "errors"
)
type CodeScanningDefaultSetupUpdate_languages int

const (
    CCPP_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES CodeScanningDefaultSetupUpdate_languages = iota
    CSHARP_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
    GO_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
    JAVAKOTLIN_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
    JAVASCRIPTTYPESCRIPT_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
    PYTHON_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
    RUBY_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
    SWIFT_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
)

func (i CodeScanningDefaultSetupUpdate_languages) String() string {
    return []string{"c-cpp", "csharp", "go", "java-kotlin", "javascript-typescript", "python", "ruby", "swift"}[i]
}
func ParseCodeScanningDefaultSetupUpdate_languages(v string) (any, error) {
    result := CCPP_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
    switch v {
        case "c-cpp":
            result = CCPP_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
        case "csharp":
            result = CSHARP_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
        case "go":
            result = GO_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
        case "java-kotlin":
            result = JAVAKOTLIN_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
        case "javascript-typescript":
            result = JAVASCRIPTTYPESCRIPT_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
        case "python":
            result = PYTHON_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
        case "ruby":
            result = RUBY_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
        case "swift":
            result = SWIFT_CODESCANNINGDEFAULTSETUPUPDATE_LANGUAGES
        default:
            return 0, errors.New("Unknown CodeScanningDefaultSetupUpdate_languages value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetupUpdate_languages(values []CodeScanningDefaultSetupUpdate_languages) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetupUpdate_languages) isMultiValue() bool {
    return false
}
