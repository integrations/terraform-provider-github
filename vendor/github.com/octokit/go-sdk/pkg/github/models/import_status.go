package models
import (
    "errors"
)
type Import_status int

const (
    AUTH_IMPORT_STATUS Import_status = iota
    ERROR_IMPORT_STATUS
    NONE_IMPORT_STATUS
    DETECTING_IMPORT_STATUS
    CHOOSE_IMPORT_STATUS
    AUTH_FAILED_IMPORT_STATUS
    IMPORTING_IMPORT_STATUS
    MAPPING_IMPORT_STATUS
    WAITING_TO_PUSH_IMPORT_STATUS
    PUSHING_IMPORT_STATUS
    COMPLETE_IMPORT_STATUS
    SETUP_IMPORT_STATUS
    UNKNOWN_IMPORT_STATUS
    DETECTION_FOUND_MULTIPLE_IMPORT_STATUS
    DETECTION_FOUND_NOTHING_IMPORT_STATUS
    DETECTION_NEEDS_AUTH_IMPORT_STATUS
)

func (i Import_status) String() string {
    return []string{"auth", "error", "none", "detecting", "choose", "auth_failed", "importing", "mapping", "waiting_to_push", "pushing", "complete", "setup", "unknown", "detection_found_multiple", "detection_found_nothing", "detection_needs_auth"}[i]
}
func ParseImport_status(v string) (any, error) {
    result := AUTH_IMPORT_STATUS
    switch v {
        case "auth":
            result = AUTH_IMPORT_STATUS
        case "error":
            result = ERROR_IMPORT_STATUS
        case "none":
            result = NONE_IMPORT_STATUS
        case "detecting":
            result = DETECTING_IMPORT_STATUS
        case "choose":
            result = CHOOSE_IMPORT_STATUS
        case "auth_failed":
            result = AUTH_FAILED_IMPORT_STATUS
        case "importing":
            result = IMPORTING_IMPORT_STATUS
        case "mapping":
            result = MAPPING_IMPORT_STATUS
        case "waiting_to_push":
            result = WAITING_TO_PUSH_IMPORT_STATUS
        case "pushing":
            result = PUSHING_IMPORT_STATUS
        case "complete":
            result = COMPLETE_IMPORT_STATUS
        case "setup":
            result = SETUP_IMPORT_STATUS
        case "unknown":
            result = UNKNOWN_IMPORT_STATUS
        case "detection_found_multiple":
            result = DETECTION_FOUND_MULTIPLE_IMPORT_STATUS
        case "detection_found_nothing":
            result = DETECTION_FOUND_NOTHING_IMPORT_STATUS
        case "detection_needs_auth":
            result = DETECTION_NEEDS_AUTH_IMPORT_STATUS
        default:
            return 0, errors.New("Unknown Import_status value: " + v)
    }
    return &result, nil
}
func SerializeImport_status(values []Import_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Import_status) isMultiValue() bool {
    return false
}
