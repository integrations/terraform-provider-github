package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleMember2 note: max_file_path_length is in beta and subject to change.Prevent commits that include file paths that exceed a specified character limit from being pushed to the commit graph.
type RepositoryRuleMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters RepositoryRuleMember2_parametersable
    // The type property
    typeEscaped *RepositoryRuleMember2_type
}
// NewRepositoryRuleMember2 instantiates a new RepositoryRuleMember2 and sets the default values.
func NewRepositoryRuleMember2()(*RepositoryRuleMember2) {
    m := &RepositoryRuleMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleMember2_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(RepositoryRuleMember2_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleMember2_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleMember2_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a RepositoryRuleMember2_parametersable when successful
func (m *RepositoryRuleMember2) GetParameters()(RepositoryRuleMember2_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleMember2_type when successful
func (m *RepositoryRuleMember2) GetTypeEscaped()(*RepositoryRuleMember2_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RepositoryRuleMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *RepositoryRuleMember2) SetParameters(value RepositoryRuleMember2_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleMember2) SetTypeEscaped(value *RepositoryRuleMember2_type)() {
    m.typeEscaped = value
}
type RepositoryRuleMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(RepositoryRuleMember2_parametersable)
    GetTypeEscaped()(*RepositoryRuleMember2_type)
    SetParameters(value RepositoryRuleMember2_parametersable)()
    SetTypeEscaped(value *RepositoryRuleMember2_type)()
}
