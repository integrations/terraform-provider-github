package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemActionsVariablesItemWithNamePatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the variable.
    name *string
    // An array of repository ids that can access the organization variable. You can only provide a list of repository ids when the `visibility` is set to `selected`.
    selected_repository_ids []int32
    // The value of the variable.
    value *string
}
// NewItemActionsVariablesItemWithNamePatchRequestBody instantiates a new ItemActionsVariablesItemWithNamePatchRequestBody and sets the default values.
func NewItemActionsVariablesItemWithNamePatchRequestBody()(*ItemActionsVariablesItemWithNamePatchRequestBody) {
    m := &ItemActionsVariablesItemWithNamePatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsVariablesItemWithNamePatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsVariablesItemWithNamePatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsVariablesItemWithNamePatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["selected_repository_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetSelectedRepositoryIds(res)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the variable.
// returns a *string when successful
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) GetName()(*string) {
    return m.name
}
// GetSelectedRepositoryIds gets the selected_repository_ids property value. An array of repository ids that can access the organization variable. You can only provide a list of repository ids when the `visibility` is set to `selected`.
// returns a []int32 when successful
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) GetSelectedRepositoryIds()([]int32) {
    return m.selected_repository_ids
}
// GetValue gets the value property value. The value of the variable.
// returns a *string when successful
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetSelectedRepositoryIds() != nil {
        err := writer.WriteCollectionOfInt32Values("selected_repository_ids", m.GetSelectedRepositoryIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name of the variable.
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetSelectedRepositoryIds sets the selected_repository_ids property value. An array of repository ids that can access the organization variable. You can only provide a list of repository ids when the `visibility` is set to `selected`.
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) SetSelectedRepositoryIds(value []int32)() {
    m.selected_repository_ids = value
}
// SetValue sets the value property value. The value of the variable.
func (m *ItemActionsVariablesItemWithNamePatchRequestBody) SetValue(value *string)() {
    m.value = value
}
type ItemActionsVariablesItemWithNamePatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetSelectedRepositoryIds()([]int32)
    GetValue()(*string)
    SetName(value *string)()
    SetSelectedRepositoryIds(value []int32)()
    SetValue(value *string)()
}
