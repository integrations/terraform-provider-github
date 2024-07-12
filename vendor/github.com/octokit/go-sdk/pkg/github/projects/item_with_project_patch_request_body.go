package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemWithProject_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Body of the project
    body *string
    // Name of the project
    name *string
    // Whether or not this project can be seen by everyone.
    private *bool
    // State of the project; either 'open' or 'closed'
    state *string
}
// NewItemWithProject_PatchRequestBody instantiates a new ItemWithProject_PatchRequestBody and sets the default values.
func NewItemWithProject_PatchRequestBody()(*ItemWithProject_PatchRequestBody) {
    m := &ItemWithProject_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWithProject_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemWithProject_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWithProject_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemWithProject_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. Body of the project
// returns a *string when successful
func (m *ItemWithProject_PatchRequestBody) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemWithProject_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the project
// returns a *string when successful
func (m *ItemWithProject_PatchRequestBody) GetName()(*string) {
    return m.name
}
// GetPrivate gets the private property value. Whether or not this project can be seen by everyone.
// returns a *bool when successful
func (m *ItemWithProject_PatchRequestBody) GetPrivate()(*bool) {
    return m.private
}
// GetState gets the state property value. State of the project; either 'open' or 'closed'
// returns a *string when successful
func (m *ItemWithProject_PatchRequestBody) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *ItemWithProject_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
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
        err := writer.WriteBoolValue("private", m.GetPrivate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *ItemWithProject_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. Body of the project
func (m *ItemWithProject_PatchRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetName sets the name property value. Name of the project
func (m *ItemWithProject_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetPrivate sets the private property value. Whether or not this project can be seen by everyone.
func (m *ItemWithProject_PatchRequestBody) SetPrivate(value *bool)() {
    m.private = value
}
// SetState sets the state property value. State of the project; either 'open' or 'closed'
func (m *ItemWithProject_PatchRequestBody) SetState(value *string)() {
    m.state = value
}
type ItemWithProject_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetName()(*string)
    GetPrivate()(*bool)
    GetState()(*string)
    SetBody(value *string)()
    SetName(value *string)()
    SetPrivate(value *bool)()
    SetState(value *string)()
}
