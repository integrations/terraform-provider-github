package models
import (
    "errors"
)
// The initally assigned location of a new codespace.
type CodespaceWithFullRepository_location int

const (
    EASTUS_CODESPACEWITHFULLREPOSITORY_LOCATION CodespaceWithFullRepository_location = iota
    SOUTHEASTASIA_CODESPACEWITHFULLREPOSITORY_LOCATION
    WESTEUROPE_CODESPACEWITHFULLREPOSITORY_LOCATION
    WESTUS2_CODESPACEWITHFULLREPOSITORY_LOCATION
)

func (i CodespaceWithFullRepository_location) String() string {
    return []string{"EastUs", "SouthEastAsia", "WestEurope", "WestUs2"}[i]
}
func ParseCodespaceWithFullRepository_location(v string) (any, error) {
    result := EASTUS_CODESPACEWITHFULLREPOSITORY_LOCATION
    switch v {
        case "EastUs":
            result = EASTUS_CODESPACEWITHFULLREPOSITORY_LOCATION
        case "SouthEastAsia":
            result = SOUTHEASTASIA_CODESPACEWITHFULLREPOSITORY_LOCATION
        case "WestEurope":
            result = WESTEUROPE_CODESPACEWITHFULLREPOSITORY_LOCATION
        case "WestUs2":
            result = WESTUS2_CODESPACEWITHFULLREPOSITORY_LOCATION
        default:
            return 0, errors.New("Unknown CodespaceWithFullRepository_location value: " + v)
    }
    return &result, nil
}
func SerializeCodespaceWithFullRepository_location(values []CodespaceWithFullRepository_location) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodespaceWithFullRepository_location) isMultiValue() bool {
    return false
}
