package cards
import (
    "errors"
)
type GetArchived_stateQueryParameterType int

const (
    ALL_GETARCHIVED_STATEQUERYPARAMETERTYPE GetArchived_stateQueryParameterType = iota
    ARCHIVED_GETARCHIVED_STATEQUERYPARAMETERTYPE
    NOT_ARCHIVED_GETARCHIVED_STATEQUERYPARAMETERTYPE
)

func (i GetArchived_stateQueryParameterType) String() string {
    return []string{"all", "archived", "not_archived"}[i]
}
func ParseGetArchived_stateQueryParameterType(v string) (any, error) {
    result := ALL_GETARCHIVED_STATEQUERYPARAMETERTYPE
    switch v {
        case "all":
            result = ALL_GETARCHIVED_STATEQUERYPARAMETERTYPE
        case "archived":
            result = ARCHIVED_GETARCHIVED_STATEQUERYPARAMETERTYPE
        case "not_archived":
            result = NOT_ARCHIVED_GETARCHIVED_STATEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetArchived_stateQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetArchived_stateQueryParameterType(values []GetArchived_stateQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetArchived_stateQueryParameterType) isMultiValue() bool {
    return false
}
