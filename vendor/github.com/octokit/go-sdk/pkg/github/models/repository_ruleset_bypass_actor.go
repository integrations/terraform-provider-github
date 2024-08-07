package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRulesetBypassActor an actor that can bypass rules in a ruleset
type RepositoryRulesetBypassActor struct {
    // The ID of the actor that can bypass a ruleset. If `actor_type` is `OrganizationAdmin`, this should be `1`. If `actor_type` is `DeployKey`, this should be null. `OrganizationAdmin` is not applicable for personal repositories.
    actor_id *int32
    // The type of actor that can bypass a ruleset.
    actor_type *RepositoryRulesetBypassActor_actor_type
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type.
    bypass_mode *RepositoryRulesetBypassActor_bypass_mode
}
// NewRepositoryRulesetBypassActor instantiates a new RepositoryRulesetBypassActor and sets the default values.
func NewRepositoryRulesetBypassActor()(*RepositoryRulesetBypassActor) {
    m := &RepositoryRulesetBypassActor{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRulesetBypassActorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRulesetBypassActorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRulesetBypassActor(), nil
}
// GetActorId gets the actor_id property value. The ID of the actor that can bypass a ruleset. If `actor_type` is `OrganizationAdmin`, this should be `1`. If `actor_type` is `DeployKey`, this should be null. `OrganizationAdmin` is not applicable for personal repositories.
// returns a *int32 when successful
func (m *RepositoryRulesetBypassActor) GetActorId()(*int32) {
    return m.actor_id
}
// GetActorType gets the actor_type property value. The type of actor that can bypass a ruleset.
// returns a *RepositoryRulesetBypassActor_actor_type when successful
func (m *RepositoryRulesetBypassActor) GetActorType()(*RepositoryRulesetBypassActor_actor_type) {
    return m.actor_type
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRulesetBypassActor) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBypassMode gets the bypass_mode property value. When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type.
// returns a *RepositoryRulesetBypassActor_bypass_mode when successful
func (m *RepositoryRulesetBypassActor) GetBypassMode()(*RepositoryRulesetBypassActor_bypass_mode) {
    return m.bypass_mode
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRulesetBypassActor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actor_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorId(val)
        }
        return nil
    }
    res["actor_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRulesetBypassActor_actor_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorType(val.(*RepositoryRulesetBypassActor_actor_type))
        }
        return nil
    }
    res["bypass_mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRulesetBypassActor_bypass_mode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassMode(val.(*RepositoryRulesetBypassActor_bypass_mode))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RepositoryRulesetBypassActor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("actor_id", m.GetActorId())
        if err != nil {
            return err
        }
    }
    if m.GetActorType() != nil {
        cast := (*m.GetActorType()).String()
        err := writer.WriteStringValue("actor_type", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBypassMode() != nil {
        cast := (*m.GetBypassMode()).String()
        err := writer.WriteStringValue("bypass_mode", &cast)
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
// SetActorId sets the actor_id property value. The ID of the actor that can bypass a ruleset. If `actor_type` is `OrganizationAdmin`, this should be `1`. If `actor_type` is `DeployKey`, this should be null. `OrganizationAdmin` is not applicable for personal repositories.
func (m *RepositoryRulesetBypassActor) SetActorId(value *int32)() {
    m.actor_id = value
}
// SetActorType sets the actor_type property value. The type of actor that can bypass a ruleset.
func (m *RepositoryRulesetBypassActor) SetActorType(value *RepositoryRulesetBypassActor_actor_type)() {
    m.actor_type = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoryRulesetBypassActor) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBypassMode sets the bypass_mode property value. When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type.
func (m *RepositoryRulesetBypassActor) SetBypassMode(value *RepositoryRulesetBypassActor_bypass_mode)() {
    m.bypass_mode = value
}
type RepositoryRulesetBypassActorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActorId()(*int32)
    GetActorType()(*RepositoryRulesetBypassActor_actor_type)
    GetBypassMode()(*RepositoryRulesetBypassActor_bypass_mode)
    SetActorId(value *int32)()
    SetActorType(value *RepositoryRulesetBypassActor_actor_type)()
    SetBypassMode(value *RepositoryRulesetBypassActor_bypass_mode)()
}
