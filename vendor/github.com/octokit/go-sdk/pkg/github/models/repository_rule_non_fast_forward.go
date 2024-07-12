package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleNonFastForward prevent users with push access from force pushing to refs.
type RepositoryRuleNonFastForward struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type property
    typeEscaped *RepositoryRuleNonFastForward_type
}
// NewRepositoryRuleNonFastForward instantiates a new RepositoryRuleNonFastForward and sets the default values.
func NewRepositoryRuleNonFastForward()(*RepositoryRuleNonFastForward) {
    m := &RepositoryRuleNonFastForward{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleNonFastForwardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleNonFastForwardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleNonFastForward(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleNonFastForward) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleNonFastForward) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleNonFastForward_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleNonFastForward_type))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleNonFastForward_type when successful
func (m *RepositoryRuleNonFastForward) GetTypeEscaped()(*RepositoryRuleNonFastForward_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleNonFastForward) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *RepositoryRuleNonFastForward) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleNonFastForward) SetTypeEscaped(value *RepositoryRuleNonFastForward_type)() {
    m.typeEscaped = value
}
type RepositoryRuleNonFastForwardable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*RepositoryRuleNonFastForward_type)
    SetTypeEscaped(value *RepositoryRuleNonFastForward_type)()
}
