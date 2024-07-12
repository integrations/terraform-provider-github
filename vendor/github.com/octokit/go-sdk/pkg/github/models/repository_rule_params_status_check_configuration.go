package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleParamsStatusCheckConfiguration required status check
type RepositoryRuleParamsStatusCheckConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The status check context name that must be present on the commit.
    context *string
    // The optional integration ID that this status check must originate from.
    integration_id *int32
}
// NewRepositoryRuleParamsStatusCheckConfiguration instantiates a new RepositoryRuleParamsStatusCheckConfiguration and sets the default values.
func NewRepositoryRuleParamsStatusCheckConfiguration()(*RepositoryRuleParamsStatusCheckConfiguration) {
    m := &RepositoryRuleParamsStatusCheckConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleParamsStatusCheckConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleParamsStatusCheckConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleParamsStatusCheckConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleParamsStatusCheckConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContext gets the context property value. The status check context name that must be present on the commit.
// returns a *string when successful
func (m *RepositoryRuleParamsStatusCheckConfiguration) GetContext()(*string) {
    return m.context
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleParamsStatusCheckConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["context"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val)
        }
        return nil
    }
    res["integration_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrationId(val)
        }
        return nil
    }
    return res
}
// GetIntegrationId gets the integration_id property value. The optional integration ID that this status check must originate from.
// returns a *int32 when successful
func (m *RepositoryRuleParamsStatusCheckConfiguration) GetIntegrationId()(*int32) {
    return m.integration_id
}
// Serialize serializes information the current object
func (m *RepositoryRuleParamsStatusCheckConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("context", m.GetContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("integration_id", m.GetIntegrationId())
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
func (m *RepositoryRuleParamsStatusCheckConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContext sets the context property value. The status check context name that must be present on the commit.
func (m *RepositoryRuleParamsStatusCheckConfiguration) SetContext(value *string)() {
    m.context = value
}
// SetIntegrationId sets the integration_id property value. The optional integration ID that this status check must originate from.
func (m *RepositoryRuleParamsStatusCheckConfiguration) SetIntegrationId(value *int32)() {
    m.integration_id = value
}
type RepositoryRuleParamsStatusCheckConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContext()(*string)
    GetIntegrationId()(*int32)
    SetContext(value *string)()
    SetIntegrationId(value *int32)()
}
