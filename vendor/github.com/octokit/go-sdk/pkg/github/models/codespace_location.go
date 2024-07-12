package models
import (
    "errors"
)
// The initally assigned location of a new codespace.
type Codespace_location int

const (
    EASTUS_CODESPACE_LOCATION Codespace_location = iota
    SOUTHEASTASIA_CODESPACE_LOCATION
    WESTEUROPE_CODESPACE_LOCATION
    WESTUS2_CODESPACE_LOCATION
)

func (i Codespace_location) String() string {
    return []string{"EastUs", "SouthEastAsia", "WestEurope", "WestUs2"}[i]
}
func ParseCodespace_location(v string) (any, error) {
    result := EASTUS_CODESPACE_LOCATION
    switch v {
        case "EastUs":
            result = EASTUS_CODESPACE_LOCATION
        case "SouthEastAsia":
            result = SOUTHEASTASIA_CODESPACE_LOCATION
        case "WestEurope":
            result = WESTEUROPE_CODESPACE_LOCATION
        case "WestUs2":
            result = WESTUS2_CODESPACE_LOCATION
        default:
            return 0, errors.New("Unknown Codespace_location value: " + v)
    }
    return &result, nil
}
func SerializeCodespace_location(values []Codespace_location) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Codespace_location) isMultiValue() bool {
    return false
}
