package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodespacesItemPublishPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A name for the new repository.
    name *string
    // Whether the new repository should be private.
    private *bool
}
// NewCodespacesItemPublishPostRequestBody instantiates a new CodespacesItemPublishPostRequestBody and sets the default values.
func NewCodespacesItemPublishPostRequestBody()(*CodespacesItemPublishPostRequestBody) {
    m := &CodespacesItemPublishPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesItemPublishPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodespacesItemPublishPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesItemPublishPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodespacesItemPublishPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodespacesItemPublishPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["private"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivate(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. A name for the new repository.
// returns a *string when successful
func (m *CodespacesItemPublishPostRequestBody) GetName()(*string) {
    return m.name
}
// GetPrivate gets the private property value. Whether the new repository should be private.
// returns a *bool when successful
func (m *CodespacesItemPublishPostRequestBody) GetPrivate()(*bool) {
    return m.private
}
// Serialize serializes information the current object
func (m *CodespacesItemPublishPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("private", m.GetPrivate())
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
func (m *CodespacesItemPublishPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. A name for the new repository.
func (m *CodespacesItemPublishPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetPrivate sets the private property value. Whether the new repository should be private.
func (m *CodespacesItemPublishPostRequestBody) SetPrivate(value *bool)() {
    m.private = value
}
type CodespacesItemPublishPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetPrivate()(*bool)
    SetName(value *string)()
    SetPrivate(value *bool)()
}
