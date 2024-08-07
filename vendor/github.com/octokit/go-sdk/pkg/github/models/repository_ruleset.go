package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleset a set of rules to apply when specified conditions are met.
type RepositoryRuleset struct {
    // The _links property
    _links RepositoryRuleset__linksable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The actors that can bypass the rules in this ruleset
    bypass_actors []RepositoryRulesetBypassActorable
    // The conditions property
    conditions RepositoryRuleset_RepositoryRuleset_conditionsable
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The bypass type of the user making the API request for this ruleset. This field is only returned whenquerying the repository-level endpoint.
    current_user_can_bypass *RepositoryRuleset_current_user_can_bypass
    // The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
    enforcement *RepositoryRuleEnforcement
    // The ID of the ruleset
    id *int32
    // The name of the ruleset
    name *string
    // The node_id property
    node_id *string
    // The rules property
    rules []RepositoryRuleable
    // The name of the source
    source *string
    // The type of the source of the ruleset
    source_type *RepositoryRuleset_source_type
    // The target of the ruleset**Note**: The `push` target is in beta and is subject to change.
    target *RepositoryRuleset_target
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// RepositoryRuleset_RepositoryRuleset_conditions composed type wrapper for classes OrgRulesetConditionsable, RepositoryRulesetConditionsable
type RepositoryRuleset_RepositoryRuleset_conditions struct {
    // Composed type representation for type OrgRulesetConditionsable
    orgRulesetConditions OrgRulesetConditionsable
    // Composed type representation for type RepositoryRulesetConditionsable
    repositoryRulesetConditions RepositoryRulesetConditionsable
}
// NewRepositoryRuleset_RepositoryRuleset_conditions instantiates a new RepositoryRuleset_RepositoryRuleset_conditions and sets the default values.
func NewRepositoryRuleset_RepositoryRuleset_conditions()(*RepositoryRuleset_RepositoryRuleset_conditions) {
    m := &RepositoryRuleset_RepositoryRuleset_conditions{
    }
    return m
}
// CreateRepositoryRuleset_RepositoryRuleset_conditionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleset_RepositoryRuleset_conditionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewRepositoryRuleset_RepositoryRuleset_conditions()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateOrgRulesetConditionsFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(OrgRulesetConditionsable); ok {
                result.SetOrgRulesetConditions(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateRepositoryRulesetConditionsFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(RepositoryRulesetConditionsable); ok {
                result.SetRepositoryRulesetConditions(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleset_RepositoryRuleset_conditions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *RepositoryRuleset_RepositoryRuleset_conditions) GetIsComposedType()(bool) {
    return true
}
// GetOrgRulesetConditions gets the orgRulesetConditions property value. Composed type representation for type OrgRulesetConditionsable
// returns a OrgRulesetConditionsable when successful
func (m *RepositoryRuleset_RepositoryRuleset_conditions) GetOrgRulesetConditions()(OrgRulesetConditionsable) {
    return m.orgRulesetConditions
}
// GetRepositoryRulesetConditions gets the repositoryRulesetConditions property value. Composed type representation for type RepositoryRulesetConditionsable
// returns a RepositoryRulesetConditionsable when successful
func (m *RepositoryRuleset_RepositoryRuleset_conditions) GetRepositoryRulesetConditions()(RepositoryRulesetConditionsable) {
    return m.repositoryRulesetConditions
}
// Serialize serializes information the current object
func (m *RepositoryRuleset_RepositoryRuleset_conditions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOrgRulesetConditions() != nil {
        err := writer.WriteObjectValue("", m.GetOrgRulesetConditions())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRulesetConditions() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRulesetConditions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrgRulesetConditions sets the orgRulesetConditions property value. Composed type representation for type OrgRulesetConditionsable
func (m *RepositoryRuleset_RepositoryRuleset_conditions) SetOrgRulesetConditions(value OrgRulesetConditionsable)() {
    m.orgRulesetConditions = value
}
// SetRepositoryRulesetConditions sets the repositoryRulesetConditions property value. Composed type representation for type RepositoryRulesetConditionsable
func (m *RepositoryRuleset_RepositoryRuleset_conditions) SetRepositoryRulesetConditions(value RepositoryRulesetConditionsable)() {
    m.repositoryRulesetConditions = value
}
type RepositoryRuleset_RepositoryRuleset_conditionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOrgRulesetConditions()(OrgRulesetConditionsable)
    GetRepositoryRulesetConditions()(RepositoryRulesetConditionsable)
    SetOrgRulesetConditions(value OrgRulesetConditionsable)()
    SetRepositoryRulesetConditions(value RepositoryRulesetConditionsable)()
}
// NewRepositoryRuleset instantiates a new RepositoryRuleset and sets the default values.
func NewRepositoryRuleset()(*RepositoryRuleset) {
    m := &RepositoryRuleset{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRulesetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRulesetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleset(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleset) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBypassActors gets the bypass_actors property value. The actors that can bypass the rules in this ruleset
// returns a []RepositoryRulesetBypassActorable when successful
func (m *RepositoryRuleset) GetBypassActors()([]RepositoryRulesetBypassActorable) {
    return m.bypass_actors
}
// GetConditions gets the conditions property value. The conditions property
// returns a RepositoryRuleset_RepositoryRuleset_conditionsable when successful
func (m *RepositoryRuleset) GetConditions()(RepositoryRuleset_RepositoryRuleset_conditionsable) {
    return m.conditions
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *RepositoryRuleset) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCurrentUserCanBypass gets the current_user_can_bypass property value. The bypass type of the user making the API request for this ruleset. This field is only returned whenquerying the repository-level endpoint.
// returns a *RepositoryRuleset_current_user_can_bypass when successful
func (m *RepositoryRuleset) GetCurrentUserCanBypass()(*RepositoryRuleset_current_user_can_bypass) {
    return m.current_user_can_bypass
}
// GetEnforcement gets the enforcement property value. The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
// returns a *RepositoryRuleEnforcement when successful
func (m *RepositoryRuleset) GetEnforcement()(*RepositoryRuleEnforcement) {
    return m.enforcement
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleset) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleset__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(RepositoryRuleset__linksable))
        }
        return nil
    }
    res["bypass_actors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryRulesetBypassActorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryRulesetBypassActorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryRulesetBypassActorable)
                }
            }
            m.SetBypassActors(res)
        }
        return nil
    }
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleset_RepositoryRuleset_conditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(RepositoryRuleset_RepositoryRuleset_conditionsable))
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["current_user_can_bypass"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleset_current_user_can_bypass)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserCanBypass(val.(*RepositoryRuleset_current_user_can_bypass))
        }
        return nil
    }
    res["enforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleEnforcement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcement(val.(*RepositoryRuleEnforcement))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryRuleable)
                }
            }
            m.SetRules(res)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    res["source_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleset_source_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceType(val.(*RepositoryRuleset_source_type))
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleset_target)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(*RepositoryRuleset_target))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the ruleset
// returns a *int32 when successful
func (m *RepositoryRuleset) GetId()(*int32) {
    return m.id
}
// GetLinks gets the _links property value. The _links property
// returns a RepositoryRuleset__linksable when successful
func (m *RepositoryRuleset) GetLinks()(RepositoryRuleset__linksable) {
    return m._links
}
// GetName gets the name property value. The name of the ruleset
// returns a *string when successful
func (m *RepositoryRuleset) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *RepositoryRuleset) GetNodeId()(*string) {
    return m.node_id
}
// GetRules gets the rules property value. The rules property
// returns a []RepositoryRuleable when successful
func (m *RepositoryRuleset) GetRules()([]RepositoryRuleable) {
    return m.rules
}
// GetSource gets the source property value. The name of the source
// returns a *string when successful
func (m *RepositoryRuleset) GetSource()(*string) {
    return m.source
}
// GetSourceType gets the source_type property value. The type of the source of the ruleset
// returns a *RepositoryRuleset_source_type when successful
func (m *RepositoryRuleset) GetSourceType()(*RepositoryRuleset_source_type) {
    return m.source_type
}
// GetTarget gets the target property value. The target of the ruleset**Note**: The `push` target is in beta and is subject to change.
// returns a *RepositoryRuleset_target when successful
func (m *RepositoryRuleset) GetTarget()(*RepositoryRuleset_target) {
    return m.target
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *RepositoryRuleset) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *RepositoryRuleset) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBypassActors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBypassActors()))
        for i, v := range m.GetBypassActors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("bypass_actors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    if m.GetCurrentUserCanBypass() != nil {
        cast := (*m.GetCurrentUserCanBypass()).String()
        err := writer.WriteStringValue("current_user_can_bypass", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnforcement() != nil {
        cast := (*m.GetEnforcement()).String()
        err := writer.WriteStringValue("enforcement", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    if m.GetSourceType() != nil {
        cast := (*m.GetSourceType()).String()
        err := writer.WriteStringValue("source_type", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTarget() != nil {
        cast := (*m.GetTarget()).String()
        err := writer.WriteStringValue("target", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("_links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoryRuleset) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBypassActors sets the bypass_actors property value. The actors that can bypass the rules in this ruleset
func (m *RepositoryRuleset) SetBypassActors(value []RepositoryRulesetBypassActorable)() {
    m.bypass_actors = value
}
// SetConditions sets the conditions property value. The conditions property
func (m *RepositoryRuleset) SetConditions(value RepositoryRuleset_RepositoryRuleset_conditionsable)() {
    m.conditions = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *RepositoryRuleset) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCurrentUserCanBypass sets the current_user_can_bypass property value. The bypass type of the user making the API request for this ruleset. This field is only returned whenquerying the repository-level endpoint.
func (m *RepositoryRuleset) SetCurrentUserCanBypass(value *RepositoryRuleset_current_user_can_bypass)() {
    m.current_user_can_bypass = value
}
// SetEnforcement sets the enforcement property value. The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
func (m *RepositoryRuleset) SetEnforcement(value *RepositoryRuleEnforcement)() {
    m.enforcement = value
}
// SetId sets the id property value. The ID of the ruleset
func (m *RepositoryRuleset) SetId(value *int32)() {
    m.id = value
}
// SetLinks sets the _links property value. The _links property
func (m *RepositoryRuleset) SetLinks(value RepositoryRuleset__linksable)() {
    m._links = value
}
// SetName sets the name property value. The name of the ruleset
func (m *RepositoryRuleset) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *RepositoryRuleset) SetNodeId(value *string)() {
    m.node_id = value
}
// SetRules sets the rules property value. The rules property
func (m *RepositoryRuleset) SetRules(value []RepositoryRuleable)() {
    m.rules = value
}
// SetSource sets the source property value. The name of the source
func (m *RepositoryRuleset) SetSource(value *string)() {
    m.source = value
}
// SetSourceType sets the source_type property value. The type of the source of the ruleset
func (m *RepositoryRuleset) SetSourceType(value *RepositoryRuleset_source_type)() {
    m.source_type = value
}
// SetTarget sets the target property value. The target of the ruleset**Note**: The `push` target is in beta and is subject to change.
func (m *RepositoryRuleset) SetTarget(value *RepositoryRuleset_target)() {
    m.target = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *RepositoryRuleset) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type RepositoryRulesetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBypassActors()([]RepositoryRulesetBypassActorable)
    GetConditions()(RepositoryRuleset_RepositoryRuleset_conditionsable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentUserCanBypass()(*RepositoryRuleset_current_user_can_bypass)
    GetEnforcement()(*RepositoryRuleEnforcement)
    GetId()(*int32)
    GetLinks()(RepositoryRuleset__linksable)
    GetName()(*string)
    GetNodeId()(*string)
    GetRules()([]RepositoryRuleable)
    GetSource()(*string)
    GetSourceType()(*RepositoryRuleset_source_type)
    GetTarget()(*RepositoryRuleset_target)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetBypassActors(value []RepositoryRulesetBypassActorable)()
    SetConditions(value RepositoryRuleset_RepositoryRuleset_conditionsable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentUserCanBypass(value *RepositoryRuleset_current_user_can_bypass)()
    SetEnforcement(value *RepositoryRuleEnforcement)()
    SetId(value *int32)()
    SetLinks(value RepositoryRuleset__linksable)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetRules(value []RepositoryRuleable)()
    SetSource(value *string)()
    SetSourceType(value *RepositoryRuleset_source_type)()
    SetTarget(value *RepositoryRuleset_target)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
