package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemActionsRunsItemPending_deploymentsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A comment to accompany the deployment review
    comment *string
    // The list of environment ids to approve or reject
    environment_ids []int32
}
// NewItemItemActionsRunsItemPending_deploymentsPostRequestBody instantiates a new ItemItemActionsRunsItemPending_deploymentsPostRequestBody and sets the default values.
func NewItemItemActionsRunsItemPending_deploymentsPostRequestBody()(*ItemItemActionsRunsItemPending_deploymentsPostRequestBody) {
    m := &ItemItemActionsRunsItemPending_deploymentsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsRunsItemPending_deploymentsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemActionsRunsItemPending_deploymentsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunsItemPending_deploymentsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemActionsRunsItemPending_deploymentsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. A comment to accompany the deployment review
// returns a *string when successful
func (m *ItemItemActionsRunsItemPending_deploymentsPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetEnvironmentIds gets the environment_ids property value. The list of environment ids to approve or reject
// returns a []int32 when successful
func (m *ItemItemActionsRunsItemPending_deploymentsPostRequestBody) GetEnvironmentIds()([]int32) {
    return m.environment_ids
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemActionsRunsItemPending_deploymentsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["environment_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetEnvironmentIds(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemActionsRunsItemPending_deploymentsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    if m.GetEnvironmentIds() != nil {
        err := writer.WriteCollectionOfInt32Values("environment_ids", m.GetEnvironmentIds())
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
func (m *ItemItemActionsRunsItemPending_deploymentsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. A comment to accompany the deployment review
func (m *ItemItemActionsRunsItemPending_deploymentsPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetEnvironmentIds sets the environment_ids property value. The list of environment ids to approve or reject
func (m *ItemItemActionsRunsItemPending_deploymentsPostRequestBody) SetEnvironmentIds(value []int32)() {
    m.environment_ids = value
}
type ItemItemActionsRunsItemPending_deploymentsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetEnvironmentIds()([]int32)
    SetComment(value *string)()
    SetEnvironmentIds(value []int32)()
}
