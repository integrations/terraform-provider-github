package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesPermissionsCheckForDevcontainer permission check result for a given devcontainer config.
type CodespacesPermissionsCheckForDevcontainer struct {
    // Whether the user has accepted the permissions defined by the devcontainer config
    accepted *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewCodespacesPermissionsCheckForDevcontainer instantiates a new CodespacesPermissionsCheckForDevcontainer and sets the default values.
func NewCodespacesPermissionsCheckForDevcontainer()(*CodespacesPermissionsCheckForDevcontainer) {
    m := &CodespacesPermissionsCheckForDevcontainer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesPermissionsCheckForDevcontainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodespacesPermissionsCheckForDevcontainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesPermissionsCheckForDevcontainer(), nil
}
// GetAccepted gets the accepted property value. Whether the user has accepted the permissions defined by the devcontainer config
// returns a *bool when successful
func (m *CodespacesPermissionsCheckForDevcontainer) GetAccepted()(*bool) {
    return m.accepted
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodespacesPermissionsCheckForDevcontainer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodespacesPermissionsCheckForDevcontainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccepted(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CodespacesPermissionsCheckForDevcontainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("accepted", m.GetAccepted())
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
// SetAccepted sets the accepted property value. Whether the user has accepted the permissions defined by the devcontainer config
func (m *CodespacesPermissionsCheckForDevcontainer) SetAccepted(value *bool)() {
    m.accepted = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodespacesPermissionsCheckForDevcontainer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type CodespacesPermissionsCheckForDevcontainerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccepted()(*bool)
    SetAccepted(value *bool)()
}
