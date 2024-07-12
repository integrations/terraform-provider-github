package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemGitRefsItemWithRefPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether to force the update or to make sure the update is a fast-forward update. Leaving this out or setting it to `false` will make sure you're not overwriting work.
    force *bool
    // The SHA1 value to set this reference to
    sha *string
}
// NewItemItemGitRefsItemWithRefPatchRequestBody instantiates a new ItemItemGitRefsItemWithRefPatchRequestBody and sets the default values.
func NewItemItemGitRefsItemWithRefPatchRequestBody()(*ItemItemGitRefsItemWithRefPatchRequestBody) {
    m := &ItemItemGitRefsItemWithRefPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemGitRefsItemWithRefPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemGitRefsItemWithRefPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemGitRefsItemWithRefPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemGitRefsItemWithRefPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemGitRefsItemWithRefPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["force"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForce(val)
        }
        return nil
    }
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    return res
}
// GetForce gets the force property value. Indicates whether to force the update or to make sure the update is a fast-forward update. Leaving this out or setting it to `false` will make sure you're not overwriting work.
// returns a *bool when successful
func (m *ItemItemGitRefsItemWithRefPatchRequestBody) GetForce()(*bool) {
    return m.force
}
// GetSha gets the sha property value. The SHA1 value to set this reference to
// returns a *string when successful
func (m *ItemItemGitRefsItemWithRefPatchRequestBody) GetSha()(*string) {
    return m.sha
}
// Serialize serializes information the current object
func (m *ItemItemGitRefsItemWithRefPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("force", m.GetForce())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
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
func (m *ItemItemGitRefsItemWithRefPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetForce sets the force property value. Indicates whether to force the update or to make sure the update is a fast-forward update. Leaving this out or setting it to `false` will make sure you're not overwriting work.
func (m *ItemItemGitRefsItemWithRefPatchRequestBody) SetForce(value *bool)() {
    m.force = value
}
// SetSha sets the sha property value. The SHA1 value to set this reference to
func (m *ItemItemGitRefsItemWithRefPatchRequestBody) SetSha(value *string)() {
    m.sha = value
}
type ItemItemGitRefsItemWithRefPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetForce()(*bool)
    GetSha()(*string)
    SetForce(value *bool)()
    SetSha(value *string)()
}
