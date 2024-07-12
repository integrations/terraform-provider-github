package models
import (
    "errors"
)
// State of the release asset.
type ReleaseAsset_state int

const (
    UPLOADED_RELEASEASSET_STATE ReleaseAsset_state = iota
    OPEN_RELEASEASSET_STATE
)

func (i ReleaseAsset_state) String() string {
    return []string{"uploaded", "open"}[i]
}
func ParseReleaseAsset_state(v string) (any, error) {
    result := UPLOADED_RELEASEASSET_STATE
    switch v {
        case "uploaded":
            result = UPLOADED_RELEASEASSET_STATE
        case "open":
            result = OPEN_RELEASEASSET_STATE
        default:
            return 0, errors.New("Unknown ReleaseAsset_state value: " + v)
    }
    return &result, nil
}
func SerializeReleaseAsset_state(values []ReleaseAsset_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReleaseAsset_state) isMultiValue() bool {
    return false
}
