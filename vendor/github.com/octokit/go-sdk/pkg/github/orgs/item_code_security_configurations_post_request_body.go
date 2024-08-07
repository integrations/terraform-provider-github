package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCodeSecurityConfigurationsPostRequestBody struct {
    // A description of the code security configuration
    description *string
    // The name of the code security configuration. Must be unique within the organization.
    name *string
}
// NewItemCodeSecurityConfigurationsPostRequestBody instantiates a new ItemCodeSecurityConfigurationsPostRequestBody and sets the default values.
func NewItemCodeSecurityConfigurationsPostRequestBody()(*ItemCodeSecurityConfigurationsPostRequestBody) {
    m := &ItemCodeSecurityConfigurationsPostRequestBody{
    }
    return m
}
// CreateItemCodeSecurityConfigurationsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCodeSecurityConfigurationsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodeSecurityConfigurationsPostRequestBody(), nil
}
// GetDescription gets the description property value. A description of the code security configuration
// returns a *string when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    return res
}
// GetName gets the name property value. The name of the code security configuration. Must be unique within the organization.
// returns a *string when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemCodeSecurityConfigurationsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
    return nil
}
// SetDescription sets the description property value. A description of the code security configuration
func (m *ItemCodeSecurityConfigurationsPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name of the code security configuration. Must be unique within the organization.
func (m *ItemCodeSecurityConfigurationsPostRequestBody) SetName(value *string)() {
    m.name = value
}
type ItemCodeSecurityConfigurationsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetName()(*string)
    SetDescription(value *string)()
    SetName(value *string)()
}
