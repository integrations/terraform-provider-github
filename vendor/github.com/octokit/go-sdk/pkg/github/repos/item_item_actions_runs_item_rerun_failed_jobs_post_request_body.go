package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemActionsRunsItemRerunFailedJobsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether to enable debug logging for the re-run.
    enable_debug_logging *bool
}
// NewItemItemActionsRunsItemRerunFailedJobsPostRequestBody instantiates a new ItemItemActionsRunsItemRerunFailedJobsPostRequestBody and sets the default values.
func NewItemItemActionsRunsItemRerunFailedJobsPostRequestBody()(*ItemItemActionsRunsItemRerunFailedJobsPostRequestBody) {
    m := &ItemItemActionsRunsItemRerunFailedJobsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsRunsItemRerunFailedJobsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemActionsRunsItemRerunFailedJobsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunsItemRerunFailedJobsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemActionsRunsItemRerunFailedJobsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnableDebugLogging gets the enable_debug_logging property value. Whether to enable debug logging for the re-run.
// returns a *bool when successful
func (m *ItemItemActionsRunsItemRerunFailedJobsPostRequestBody) GetEnableDebugLogging()(*bool) {
    return m.enable_debug_logging
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemActionsRunsItemRerunFailedJobsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enable_debug_logging"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableDebugLogging(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemActionsRunsItemRerunFailedJobsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enable_debug_logging", m.GetEnableDebugLogging())
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
func (m *ItemItemActionsRunsItemRerunFailedJobsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnableDebugLogging sets the enable_debug_logging property value. Whether to enable debug logging for the re-run.
func (m *ItemItemActionsRunsItemRerunFailedJobsPostRequestBody) SetEnableDebugLogging(value *bool)() {
    m.enable_debug_logging = value
}
type ItemItemActionsRunsItemRerunFailedJobsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnableDebugLogging()(*bool)
    SetEnableDebugLogging(value *bool)()
}
