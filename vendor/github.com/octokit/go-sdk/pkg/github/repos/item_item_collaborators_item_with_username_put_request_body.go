package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemCollaboratorsItemWithUsernamePutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The permission to grant the collaborator. **Only valid on organization-owned repositories.** We accept the following permissions to be set: `pull`, `triage`, `push`, `maintain`, `admin` and you can also specify a custom repository role name, if the owning organization has defined any.
    permission *string
}
// NewItemItemCollaboratorsItemWithUsernamePutRequestBody instantiates a new ItemItemCollaboratorsItemWithUsernamePutRequestBody and sets the default values.
func NewItemItemCollaboratorsItemWithUsernamePutRequestBody()(*ItemItemCollaboratorsItemWithUsernamePutRequestBody) {
    m := &ItemItemCollaboratorsItemWithUsernamePutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    permissionValue := "push"
    m.SetPermission(&permissionValue)
    return m
}
// CreateItemItemCollaboratorsItemWithUsernamePutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCollaboratorsItemWithUsernamePutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCollaboratorsItemWithUsernamePutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemCollaboratorsItemWithUsernamePutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCollaboratorsItemWithUsernamePutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val)
        }
        return nil
    }
    return res
}
// GetPermission gets the permission property value. The permission to grant the collaborator. **Only valid on organization-owned repositories.** We accept the following permissions to be set: `pull`, `triage`, `push`, `maintain`, `admin` and you can also specify a custom repository role name, if the owning organization has defined any.
// returns a *string when successful
func (m *ItemItemCollaboratorsItemWithUsernamePutRequestBody) GetPermission()(*string) {
    return m.permission
}
// Serialize serializes information the current object
func (m *ItemItemCollaboratorsItemWithUsernamePutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("permission", m.GetPermission())
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
func (m *ItemItemCollaboratorsItemWithUsernamePutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPermission sets the permission property value. The permission to grant the collaborator. **Only valid on organization-owned repositories.** We accept the following permissions to be set: `pull`, `triage`, `push`, `maintain`, `admin` and you can also specify a custom repository role name, if the owning organization has defined any.
func (m *ItemItemCollaboratorsItemWithUsernamePutRequestBody) SetPermission(value *string)() {
    m.permission = value
}
type ItemItemCollaboratorsItemWithUsernamePutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPermission()(*string)
    SetPermission(value *string)()
}
