package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemCodespacesDevcontainersGetResponse_devcontainers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display_name property
    display_name *string
    // The name property
    name *string
    // The path property
    path *string
}
// NewItemItemCodespacesDevcontainersGetResponse_devcontainers instantiates a new ItemItemCodespacesDevcontainersGetResponse_devcontainers and sets the default values.
func NewItemItemCodespacesDevcontainersGetResponse_devcontainers()(*ItemItemCodespacesDevcontainersGetResponse_devcontainers) {
    m := &ItemItemCodespacesDevcontainersGetResponse_devcontainers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCodespacesDevcontainersGetResponse_devcontainersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCodespacesDevcontainersGetResponse_devcontainersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesDevcontainersGetResponse_devcontainers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the display_name property value. The display_name property
// returns a *string when successful
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) GetDisplayName()(*string) {
    return m.display_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) GetName()(*string) {
    return m.name
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) GetPath()(*string) {
    return m.path
}
// Serialize serializes information the current object
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
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
        err := writer.WriteStringValue("path", m.GetPath())
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
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the display_name property value. The display_name property
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetName sets the name property value. The name property
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) SetName(value *string)() {
    m.name = value
}
// SetPath sets the path property value. The path property
func (m *ItemItemCodespacesDevcontainersGetResponse_devcontainers) SetPath(value *string)() {
    m.path = value
}
type ItemItemCodespacesDevcontainersGetResponse_devcontainersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetName()(*string)
    GetPath()(*string)
    SetDisplayName(value *string)()
    SetName(value *string)()
    SetPath(value *string)()
}
