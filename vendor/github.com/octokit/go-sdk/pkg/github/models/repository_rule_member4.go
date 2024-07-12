package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleMember4 note: max_file_size is in beta and subject to change.Prevent commits that exceed a specified file size limit from being pushed to the commit.
type RepositoryRuleMember4 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters RepositoryRuleMember4_parametersable
    // The type property
    typeEscaped *RepositoryRuleMember4_type
}
// NewRepositoryRuleMember4 instantiates a new RepositoryRuleMember4 and sets the default values.
func NewRepositoryRuleMember4()(*RepositoryRuleMember4) {
    m := &RepositoryRuleMember4{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleMember4FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleMember4FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleMember4(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleMember4) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleMember4) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleMember4_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(RepositoryRuleMember4_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleMember4_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleMember4_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a RepositoryRuleMember4_parametersable when successful
func (m *RepositoryRuleMember4) GetParameters()(RepositoryRuleMember4_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleMember4_type when successful
func (m *RepositoryRuleMember4) GetTypeEscaped()(*RepositoryRuleMember4_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleMember4) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("parameters", m.GetParameters())
        if err != nil {
            return err
        }
    }
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
func (m *RepositoryRuleMember4) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *RepositoryRuleMember4) SetParameters(value RepositoryRuleMember4_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleMember4) SetTypeEscaped(value *RepositoryRuleMember4_type)() {
    m.typeEscaped = value
}
type RepositoryRuleMember4able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(RepositoryRuleMember4_parametersable)
    GetTypeEscaped()(*RepositoryRuleMember4_type)
    SetParameters(value RepositoryRuleMember4_parametersable)()
    SetTypeEscaped(value *RepositoryRuleMember4_type)()
}
