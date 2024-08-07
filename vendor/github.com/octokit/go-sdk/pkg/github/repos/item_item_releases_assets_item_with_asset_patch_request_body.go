package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemReleasesAssetsItemWithAsset_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An alternate short description of the asset. Used in place of the filename.
    label *string
    // The file name of the asset.
    name *string
    // The state property
    state *string
}
// NewItemItemReleasesAssetsItemWithAsset_PatchRequestBody instantiates a new ItemItemReleasesAssetsItemWithAsset_PatchRequestBody and sets the default values.
func NewItemItemReleasesAssetsItemWithAsset_PatchRequestBody()(*ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) {
    m := &ItemItemReleasesAssetsItemWithAsset_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemReleasesAssetsItemWithAsset_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemReleasesAssetsItemWithAsset_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemReleasesAssetsItemWithAsset_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
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
// GetLabel gets the label property value. An alternate short description of the asset. Used in place of the filename.
// returns a *string when successful
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) GetLabel()(*string) {
    return m.label
}
// GetName gets the name property value. The file name of the asset.
// returns a *string when successful
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) GetName()(*string) {
    return m.name
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("label", m.GetLabel())
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
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabel sets the label property value. An alternate short description of the asset. Used in place of the filename.
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) SetLabel(value *string)() {
    m.label = value
}
// SetName sets the name property value. The file name of the asset.
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetState sets the state property value. The state property
func (m *ItemItemReleasesAssetsItemWithAsset_PatchRequestBody) SetState(value *string)() {
    m.state = value
}
type ItemItemReleasesAssetsItemWithAsset_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabel()(*string)
    GetName()(*string)
    GetState()(*string)
    SetLabel(value *string)()
    SetName(value *string)()
    SetState(value *string)()
}
