package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemMergeUpstreamPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the branch which should be updated to match upstream.
    branch *string
}
// NewItemItemMergeUpstreamPostRequestBody instantiates a new ItemItemMergeUpstreamPostRequestBody and sets the default values.
func NewItemItemMergeUpstreamPostRequestBody()(*ItemItemMergeUpstreamPostRequestBody) {
    m := &ItemItemMergeUpstreamPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemMergeUpstreamPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemMergeUpstreamPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemMergeUpstreamPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemMergeUpstreamPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBranch gets the branch property value. The name of the branch which should be updated to match upstream.
// returns a *string when successful
func (m *ItemItemMergeUpstreamPostRequestBody) GetBranch()(*string) {
    return m.branch
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemMergeUpstreamPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranch(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemMergeUpstreamPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("branch", m.GetBranch())
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
func (m *ItemItemMergeUpstreamPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBranch sets the branch property value. The name of the branch which should be updated to match upstream.
func (m *ItemItemMergeUpstreamPostRequestBody) SetBranch(value *string)() {
    m.branch = value
}
type ItemItemMergeUpstreamPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranch()(*string)
    SetBranch(value *string)()
}
