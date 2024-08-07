package migrations
import (
    "errors"
)
type MigrationsPostRequestBody_exclude int

const (
    REPOSITORIES_MIGRATIONSPOSTREQUESTBODY_EXCLUDE MigrationsPostRequestBody_exclude = iota
)

func (i MigrationsPostRequestBody_exclude) String() string {
    return []string{"repositories"}[i]
}
func ParseMigrationsPostRequestBody_exclude(v string) (any, error) {
    result := REPOSITORIES_MIGRATIONSPOSTREQUESTBODY_EXCLUDE
    switch v {
        case "repositories":
            result = REPOSITORIES_MIGRATIONSPOSTREQUESTBODY_EXCLUDE
        default:
            return 0, errors.New("Unknown MigrationsPostRequestBody_exclude value: " + v)
    }
    return &result, nil
}
func SerializeMigrationsPostRequestBody_exclude(values []MigrationsPostRequestBody_exclude) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MigrationsPostRequestBody_exclude) isMultiValue() bool {
    return false
}
