package models
import (
    "errors"
)
// A classification of the file. For example to identify it as generated.
type CodeScanningAlertClassification int

const (
    SOURCE_CODESCANNINGALERTCLASSIFICATION CodeScanningAlertClassification = iota
    GENERATED_CODESCANNINGALERTCLASSIFICATION
    TEST_CODESCANNINGALERTCLASSIFICATION
    LIBRARY_CODESCANNINGALERTCLASSIFICATION
)

func (i CodeScanningAlertClassification) String() string {
    return []string{"source", "generated", "test", "library"}[i]
}
func ParseCodeScanningAlertClassification(v string) (any, error) {
    result := SOURCE_CODESCANNINGALERTCLASSIFICATION
    switch v {
        case "source":
            result = SOURCE_CODESCANNINGALERTCLASSIFICATION
        case "generated":
            result = GENERATED_CODESCANNINGALERTCLASSIFICATION
        case "test":
            result = TEST_CODESCANNINGALERTCLASSIFICATION
        case "library":
            result = LIBRARY_CODESCANNINGALERTCLASSIFICATION
        default:
            return 0, errors.New("Unknown CodeScanningAlertClassification value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertClassification(values []CodeScanningAlertClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAlertClassification) isMultiValue() bool {
    return false
}
