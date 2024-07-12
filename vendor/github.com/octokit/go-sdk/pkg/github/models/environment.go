package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Environment details of a deployment environment
type Environment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time that the environment was created, in ISO 8601 format.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
    deployment_branch_policy DeploymentBranchPolicySettingsable
    // The html_url property
    html_url *string
    // The id of the environment.
    id *int32
    // The name of the environment.
    name *string
    // The node_id property
    node_id *string
    // Built-in deployment protection rules for the environment.
    protection_rules []Environment_Environment_protection_rulesable
    // The time that the environment was last updated, in ISO 8601 format.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// Environment_Environment_protection_rules composed type wrapper for classes Environment_protection_rulesMember1able, Environment_protection_rulesMember2able, Environment_protection_rulesMember3able
type Environment_Environment_protection_rules struct {
    // Composed type representation for type Environment_protection_rulesMember1able
    environment_protection_rulesMember1 Environment_protection_rulesMember1able
    // Composed type representation for type Environment_protection_rulesMember2able
    environment_protection_rulesMember2 Environment_protection_rulesMember2able
    // Composed type representation for type Environment_protection_rulesMember3able
    environment_protection_rulesMember3 Environment_protection_rulesMember3able
}
// NewEnvironment_Environment_protection_rules instantiates a new Environment_Environment_protection_rules and sets the default values.
func NewEnvironment_Environment_protection_rules()(*Environment_Environment_protection_rules) {
    m := &Environment_Environment_protection_rules{
    }
    return m
}
// CreateEnvironment_Environment_protection_rulesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnvironment_Environment_protection_rulesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEnvironment_Environment_protection_rules()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateEnvironment_protection_rulesMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(Environment_protection_rulesMember1able); ok {
                result.SetEnvironmentProtectionRulesMember1(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateEnvironment_protection_rulesMember2FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(Environment_protection_rulesMember2able); ok {
                result.SetEnvironmentProtectionRulesMember2(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateEnvironment_protection_rulesMember3FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(Environment_protection_rulesMember3able); ok {
                result.SetEnvironmentProtectionRulesMember3(cast)
            }
        }
    }
    return result, nil
}
// GetEnvironmentProtectionRulesMember1 gets the environment_protection_rulesMember1 property value. Composed type representation for type Environment_protection_rulesMember1able
// returns a Environment_protection_rulesMember1able when successful
func (m *Environment_Environment_protection_rules) GetEnvironmentProtectionRulesMember1()(Environment_protection_rulesMember1able) {
    return m.environment_protection_rulesMember1
}
// GetEnvironmentProtectionRulesMember2 gets the environment_protection_rulesMember2 property value. Composed type representation for type Environment_protection_rulesMember2able
// returns a Environment_protection_rulesMember2able when successful
func (m *Environment_Environment_protection_rules) GetEnvironmentProtectionRulesMember2()(Environment_protection_rulesMember2able) {
    return m.environment_protection_rulesMember2
}
// GetEnvironmentProtectionRulesMember3 gets the environment_protection_rulesMember3 property value. Composed type representation for type Environment_protection_rulesMember3able
// returns a Environment_protection_rulesMember3able when successful
func (m *Environment_Environment_protection_rules) GetEnvironmentProtectionRulesMember3()(Environment_protection_rulesMember3able) {
    return m.environment_protection_rulesMember3
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Environment_Environment_protection_rules) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *Environment_Environment_protection_rules) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *Environment_Environment_protection_rules) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnvironmentProtectionRulesMember1() != nil {
        err := writer.WriteObjectValue("", m.GetEnvironmentProtectionRulesMember1())
        if err != nil {
            return err
        }
    } else if m.GetEnvironmentProtectionRulesMember2() != nil {
        err := writer.WriteObjectValue("", m.GetEnvironmentProtectionRulesMember2())
        if err != nil {
            return err
        }
    } else if m.GetEnvironmentProtectionRulesMember3() != nil {
        err := writer.WriteObjectValue("", m.GetEnvironmentProtectionRulesMember3())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnvironmentProtectionRulesMember1 sets the environment_protection_rulesMember1 property value. Composed type representation for type Environment_protection_rulesMember1able
func (m *Environment_Environment_protection_rules) SetEnvironmentProtectionRulesMember1(value Environment_protection_rulesMember1able)() {
    m.environment_protection_rulesMember1 = value
}
// SetEnvironmentProtectionRulesMember2 sets the environment_protection_rulesMember2 property value. Composed type representation for type Environment_protection_rulesMember2able
func (m *Environment_Environment_protection_rules) SetEnvironmentProtectionRulesMember2(value Environment_protection_rulesMember2able)() {
    m.environment_protection_rulesMember2 = value
}
// SetEnvironmentProtectionRulesMember3 sets the environment_protection_rulesMember3 property value. Composed type representation for type Environment_protection_rulesMember3able
func (m *Environment_Environment_protection_rules) SetEnvironmentProtectionRulesMember3(value Environment_protection_rulesMember3able)() {
    m.environment_protection_rulesMember3 = value
}
type Environment_Environment_protection_rulesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnvironmentProtectionRulesMember1()(Environment_protection_rulesMember1able)
    GetEnvironmentProtectionRulesMember2()(Environment_protection_rulesMember2able)
    GetEnvironmentProtectionRulesMember3()(Environment_protection_rulesMember3able)
    SetEnvironmentProtectionRulesMember1(value Environment_protection_rulesMember1able)()
    SetEnvironmentProtectionRulesMember2(value Environment_protection_rulesMember2able)()
    SetEnvironmentProtectionRulesMember3(value Environment_protection_rulesMember3able)()
}
// NewEnvironment instantiates a new Environment and sets the default values.
func NewEnvironment()(*Environment) {
    m := &Environment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnvironmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnvironment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Environment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The time that the environment was created, in ISO 8601 format.
// returns a *Time when successful
func (m *Environment) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDeploymentBranchPolicy gets the deployment_branch_policy property value. The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
// returns a DeploymentBranchPolicySettingsable when successful
func (m *Environment) GetDeploymentBranchPolicy()(DeploymentBranchPolicySettingsable) {
    return m.deployment_branch_policy
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Environment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["deployment_branch_policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentBranchPolicySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentBranchPolicy(val.(DeploymentBranchPolicySettingsable))
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
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
    res["protection_rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnvironment_Environment_protection_rulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Environment_Environment_protection_rulesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Environment_Environment_protection_rulesable)
                }
            }
            m.SetProtectionRules(res)
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
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Environment) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id of the environment.
// returns a *int32 when successful
func (m *Environment) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the environment.
// returns a *string when successful
func (m *Environment) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Environment) GetNodeId()(*string) {
    return m.node_id
}
// GetProtectionRules gets the protection_rules property value. Built-in deployment protection rules for the environment.
// returns a []Environment_Environment_protection_rulesable when successful
func (m *Environment) GetProtectionRules()([]Environment_Environment_protection_rulesable) {
    return m.protection_rules
}
// GetUpdatedAt gets the updated_at property value. The time that the environment was last updated, in ISO 8601 format.
// returns a *Time when successful
func (m *Environment) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Environment) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Environment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deployment_branch_policy", m.GetDeploymentBranchPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
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
    if m.GetProtectionRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProtectionRules()))
        for i, v := range m.GetProtectionRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("protection_rules", cast)
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
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *Environment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The time that the environment was created, in ISO 8601 format.
func (m *Environment) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDeploymentBranchPolicy sets the deployment_branch_policy property value. The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
func (m *Environment) SetDeploymentBranchPolicy(value DeploymentBranchPolicySettingsable)() {
    m.deployment_branch_policy = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Environment) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id of the environment.
func (m *Environment) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the environment.
func (m *Environment) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Environment) SetNodeId(value *string)() {
    m.node_id = value
}
// SetProtectionRules sets the protection_rules property value. Built-in deployment protection rules for the environment.
func (m *Environment) SetProtectionRules(value []Environment_Environment_protection_rulesable)() {
    m.protection_rules = value
}
// SetUpdatedAt sets the updated_at property value. The time that the environment was last updated, in ISO 8601 format.
func (m *Environment) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Environment) SetUrl(value *string)() {
    m.url = value
}
type Environmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentBranchPolicy()(DeploymentBranchPolicySettingsable)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetProtectionRules()([]Environment_Environment_protection_rulesable)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentBranchPolicy(value DeploymentBranchPolicySettingsable)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetProtectionRules(value []Environment_Environment_protection_rulesable)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
