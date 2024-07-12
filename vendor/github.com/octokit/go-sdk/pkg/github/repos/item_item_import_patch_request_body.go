package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemImportPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // For a tfvc import, the name of the project that is being imported.
    tfvc_project *string
    // The password to provide to the originating repository.
    vcs_password *string
    // The username to provide to the originating repository.
    vcs_username *string
}
// NewItemItemImportPatchRequestBody instantiates a new ItemItemImportPatchRequestBody and sets the default values.
func NewItemItemImportPatchRequestBody()(*ItemItemImportPatchRequestBody) {
    m := &ItemItemImportPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemImportPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemImportPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemImportPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemImportPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemImportPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tfvc_project"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTfvcProject(val)
        }
        return nil
    }
    res["vcs_password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVcsPassword(val)
        }
        return nil
    }
    res["vcs_username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVcsUsername(val)
        }
        return nil
    }
    return res
}
// GetTfvcProject gets the tfvc_project property value. For a tfvc import, the name of the project that is being imported.
// returns a *string when successful
func (m *ItemItemImportPatchRequestBody) GetTfvcProject()(*string) {
    return m.tfvc_project
}
// GetVcsPassword gets the vcs_password property value. The password to provide to the originating repository.
// returns a *string when successful
func (m *ItemItemImportPatchRequestBody) GetVcsPassword()(*string) {
    return m.vcs_password
}
// GetVcsUsername gets the vcs_username property value. The username to provide to the originating repository.
// returns a *string when successful
func (m *ItemItemImportPatchRequestBody) GetVcsUsername()(*string) {
    return m.vcs_username
}
// Serialize serializes information the current object
func (m *ItemItemImportPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("tfvc_project", m.GetTfvcProject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vcs_password", m.GetVcsPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vcs_username", m.GetVcsUsername())
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
func (m *ItemItemImportPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTfvcProject sets the tfvc_project property value. For a tfvc import, the name of the project that is being imported.
func (m *ItemItemImportPatchRequestBody) SetTfvcProject(value *string)() {
    m.tfvc_project = value
}
// SetVcsPassword sets the vcs_password property value. The password to provide to the originating repository.
func (m *ItemItemImportPatchRequestBody) SetVcsPassword(value *string)() {
    m.vcs_password = value
}
// SetVcsUsername sets the vcs_username property value. The username to provide to the originating repository.
func (m *ItemItemImportPatchRequestBody) SetVcsUsername(value *string)() {
    m.vcs_username = value
}
type ItemItemImportPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTfvcProject()(*string)
    GetVcsPassword()(*string)
    GetVcsUsername()(*string)
    SetTfvcProject(value *string)()
    SetVcsPassword(value *string)()
    SetVcsUsername(value *string)()
}
