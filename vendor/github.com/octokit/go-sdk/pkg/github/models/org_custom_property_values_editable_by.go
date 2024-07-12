package models
import (
    "errors"
)
// Who can edit the values of the property
type OrgCustomProperty_values_editable_by int

const (
    ORG_ACTORS_ORGCUSTOMPROPERTY_VALUES_EDITABLE_BY OrgCustomProperty_values_editable_by = iota
    ORG_AND_REPO_ACTORS_ORGCUSTOMPROPERTY_VALUES_EDITABLE_BY
)

func (i OrgCustomProperty_values_editable_by) String() string {
    return []string{"org_actors", "org_and_repo_actors"}[i]
}
func ParseOrgCustomProperty_values_editable_by(v string) (any, error) {
    result := ORG_ACTORS_ORGCUSTOMPROPERTY_VALUES_EDITABLE_BY
    switch v {
        case "org_actors":
            result = ORG_ACTORS_ORGCUSTOMPROPERTY_VALUES_EDITABLE_BY
        case "org_and_repo_actors":
            result = ORG_AND_REPO_ACTORS_ORGCUSTOMPROPERTY_VALUES_EDITABLE_BY
        default:
            return 0, errors.New("Unknown OrgCustomProperty_values_editable_by value: " + v)
    }
    return &result, nil
}
func SerializeOrgCustomProperty_values_editable_by(values []OrgCustomProperty_values_editable_by) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrgCustomProperty_values_editable_by) isMultiValue() bool {
    return false
}
