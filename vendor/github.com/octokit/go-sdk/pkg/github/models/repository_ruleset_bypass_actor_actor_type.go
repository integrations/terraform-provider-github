package models
import (
    "errors"
)
// The type of actor that can bypass a ruleset.
type RepositoryRulesetBypassActor_actor_type int

const (
    INTEGRATION_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE RepositoryRulesetBypassActor_actor_type = iota
    ORGANIZATIONADMIN_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
    REPOSITORYROLE_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
    TEAM_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
    DEPLOYKEY_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
)

func (i RepositoryRulesetBypassActor_actor_type) String() string {
    return []string{"Integration", "OrganizationAdmin", "RepositoryRole", "Team", "DeployKey"}[i]
}
func ParseRepositoryRulesetBypassActor_actor_type(v string) (any, error) {
    result := INTEGRATION_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
    switch v {
        case "Integration":
            result = INTEGRATION_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
        case "OrganizationAdmin":
            result = ORGANIZATIONADMIN_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
        case "RepositoryRole":
            result = REPOSITORYROLE_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
        case "Team":
            result = TEAM_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
        case "DeployKey":
            result = DEPLOYKEY_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRulesetBypassActor_actor_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRulesetBypassActor_actor_type(values []RepositoryRulesetBypassActor_actor_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRulesetBypassActor_actor_type) isMultiValue() bool {
    return false
}
