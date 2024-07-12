package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleCommitAuthorEmailPattern parameters to be used for the commit_author_email_pattern rule
type RepositoryRuleCommitAuthorEmailPattern struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters RepositoryRuleCommitAuthorEmailPattern_parametersable
    // The type property
    typeEscaped *RepositoryRuleCommitAuthorEmailPattern_type
}
// NewRepositoryRuleCommitAuthorEmailPattern instantiates a new RepositoryRuleCommitAuthorEmailPattern and sets the default values.
func NewRepositoryRuleCommitAuthorEmailPattern()(*RepositoryRuleCommitAuthorEmailPattern) {
    m := &RepositoryRuleCommitAuthorEmailPattern{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleCommitAuthorEmailPatternFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleCommitAuthorEmailPatternFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleCommitAuthorEmailPattern(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleCommitAuthorEmailPattern) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleCommitAuthorEmailPattern) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleCommitAuthorEmailPattern_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(RepositoryRuleCommitAuthorEmailPattern_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleCommitAuthorEmailPattern_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleCommitAuthorEmailPattern_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a RepositoryRuleCommitAuthorEmailPattern_parametersable when successful
func (m *RepositoryRuleCommitAuthorEmailPattern) GetParameters()(RepositoryRuleCommitAuthorEmailPattern_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleCommitAuthorEmailPattern_type when successful
func (m *RepositoryRuleCommitAuthorEmailPattern) GetTypeEscaped()(*RepositoryRuleCommitAuthorEmailPattern_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleCommitAuthorEmailPattern) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RepositoryRuleCommitAuthorEmailPattern) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *RepositoryRuleCommitAuthorEmailPattern) SetParameters(value RepositoryRuleCommitAuthorEmailPattern_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleCommitAuthorEmailPattern) SetTypeEscaped(value *RepositoryRuleCommitAuthorEmailPattern_type)() {
    m.typeEscaped = value
}
type RepositoryRuleCommitAuthorEmailPatternable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(RepositoryRuleCommitAuthorEmailPattern_parametersable)
    GetTypeEscaped()(*RepositoryRuleCommitAuthorEmailPattern_type)
    SetParameters(value RepositoryRuleCommitAuthorEmailPattern_parametersable)()
    SetTypeEscaped(value *RepositoryRuleCommitAuthorEmailPattern_type)()
}
