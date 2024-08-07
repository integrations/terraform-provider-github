package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemActionsWorkflowsItemDispatchesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Input keys and values configured in the workflow file. The maximum number of properties is 10. Any default properties configured in the workflow file will be used when `inputs` are omitted.
    inputs ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsable
    // The git reference for the workflow. The reference can be a branch or tag name.
    ref *string
}
// NewItemItemActionsWorkflowsItemDispatchesPostRequestBody instantiates a new ItemItemActionsWorkflowsItemDispatchesPostRequestBody and sets the default values.
func NewItemItemActionsWorkflowsItemDispatchesPostRequestBody()(*ItemItemActionsWorkflowsItemDispatchesPostRequestBody) {
    m := &ItemItemActionsWorkflowsItemDispatchesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsWorkflowsItemDispatchesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemActionsWorkflowsItemDispatchesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsWorkflowsItemDispatchesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["inputs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInputs(val.(ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsable))
        }
        return nil
    }
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
        }
        return nil
    }
    return res
}
// GetInputs gets the inputs property value. Input keys and values configured in the workflow file. The maximum number of properties is 10. Any default properties configured in the workflow file will be used when `inputs` are omitted.
// returns a ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsable when successful
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody) GetInputs()(ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsable) {
    return m.inputs
}
// GetRef gets the ref property value. The git reference for the workflow. The reference can be a branch or tag name.
// returns a *string when successful
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody) GetRef()(*string) {
    return m.ref
}
// Serialize serializes information the current object
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("inputs", m.GetInputs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref", m.GetRef())
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
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInputs sets the inputs property value. Input keys and values configured in the workflow file. The maximum number of properties is 10. Any default properties configured in the workflow file will be used when `inputs` are omitted.
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody) SetInputs(value ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsable)() {
    m.inputs = value
}
// SetRef sets the ref property value. The git reference for the workflow. The reference can be a branch or tag name.
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody) SetRef(value *string)() {
    m.ref = value
}
type ItemItemActionsWorkflowsItemDispatchesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInputs()(ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsable)
    GetRef()(*string)
    SetInputs(value ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsable)()
    SetRef(value *string)()
}
