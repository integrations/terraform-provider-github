package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemCodespacesNewGetResponse_defaults struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The devcontainer_path property
    devcontainer_path *string
    // The location property
    location *string
}
// NewItemItemCodespacesNewGetResponse_defaults instantiates a new ItemItemCodespacesNewGetResponse_defaults and sets the default values.
func NewItemItemCodespacesNewGetResponse_defaults()(*ItemItemCodespacesNewGetResponse_defaults) {
    m := &ItemItemCodespacesNewGetResponse_defaults{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCodespacesNewGetResponse_defaultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCodespacesNewGetResponse_defaultsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesNewGetResponse_defaults(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemCodespacesNewGetResponse_defaults) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDevcontainerPath gets the devcontainer_path property value. The devcontainer_path property
// returns a *string when successful
func (m *ItemItemCodespacesNewGetResponse_defaults) GetDevcontainerPath()(*string) {
    return m.devcontainer_path
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCodespacesNewGetResponse_defaults) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["devcontainer_path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevcontainerPath(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    return res
}
// GetLocation gets the location property value. The location property
// returns a *string when successful
func (m *ItemItemCodespacesNewGetResponse_defaults) GetLocation()(*string) {
    return m.location
}
// Serialize serializes information the current object
func (m *ItemItemCodespacesNewGetResponse_defaults) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("devcontainer_path", m.GetDevcontainerPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
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
func (m *ItemItemCodespacesNewGetResponse_defaults) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDevcontainerPath sets the devcontainer_path property value. The devcontainer_path property
func (m *ItemItemCodespacesNewGetResponse_defaults) SetDevcontainerPath(value *string)() {
    m.devcontainer_path = value
}
// SetLocation sets the location property value. The location property
func (m *ItemItemCodespacesNewGetResponse_defaults) SetLocation(value *string)() {
    m.location = value
}
type ItemItemCodespacesNewGetResponse_defaultsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevcontainerPath()(*string)
    GetLocation()(*string)
    SetDevcontainerPath(value *string)()
    SetLocation(value *string)()
}
