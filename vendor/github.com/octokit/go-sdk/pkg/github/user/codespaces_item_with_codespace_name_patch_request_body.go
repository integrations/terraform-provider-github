package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodespacesItemWithCodespace_namePatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Display name for this codespace
    display_name *string
    // A valid machine to transition this codespace to.
    machine *string
    // Recently opened folders inside the codespace. It is currently used by the clients to determine the folder path to load the codespace in.
    recent_folders []string
}
// NewCodespacesItemWithCodespace_namePatchRequestBody instantiates a new CodespacesItemWithCodespace_namePatchRequestBody and sets the default values.
func NewCodespacesItemWithCodespace_namePatchRequestBody()(*CodespacesItemWithCodespace_namePatchRequestBody) {
    m := &CodespacesItemWithCodespace_namePatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesItemWithCodespace_namePatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodespacesItemWithCodespace_namePatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesItemWithCodespace_namePatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodespacesItemWithCodespace_namePatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the display_name property value. Display name for this codespace
// returns a *string when successful
func (m *CodespacesItemWithCodespace_namePatchRequestBody) GetDisplayName()(*string) {
    return m.display_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodespacesItemWithCodespace_namePatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["machine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachine(val)
        }
        return nil
    }
    res["recent_folders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRecentFolders(res)
        }
        return nil
    }
    return res
}
// GetMachine gets the machine property value. A valid machine to transition this codespace to.
// returns a *string when successful
func (m *CodespacesItemWithCodespace_namePatchRequestBody) GetMachine()(*string) {
    return m.machine
}
// GetRecentFolders gets the recent_folders property value. Recently opened folders inside the codespace. It is currently used by the clients to determine the folder path to load the codespace in.
// returns a []string when successful
func (m *CodespacesItemWithCodespace_namePatchRequestBody) GetRecentFolders()([]string) {
    return m.recent_folders
}
// Serialize serializes information the current object
func (m *CodespacesItemWithCodespace_namePatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("machine", m.GetMachine())
        if err != nil {
            return err
        }
    }
    if m.GetRecentFolders() != nil {
        err := writer.WriteCollectionOfStringValues("recent_folders", m.GetRecentFolders())
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
func (m *CodespacesItemWithCodespace_namePatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the display_name property value. Display name for this codespace
func (m *CodespacesItemWithCodespace_namePatchRequestBody) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetMachine sets the machine property value. A valid machine to transition this codespace to.
func (m *CodespacesItemWithCodespace_namePatchRequestBody) SetMachine(value *string)() {
    m.machine = value
}
// SetRecentFolders sets the recent_folders property value. Recently opened folders inside the codespace. It is currently used by the clients to determine the folder path to load the codespace in.
func (m *CodespacesItemWithCodespace_namePatchRequestBody) SetRecentFolders(value []string)() {
    m.recent_folders = value
}
type CodespacesItemWithCodespace_namePatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetMachine()(*string)
    GetRecentFolders()([]string)
    SetDisplayName(value *string)()
    SetMachine(value *string)()
    SetRecentFolders(value []string)()
}
