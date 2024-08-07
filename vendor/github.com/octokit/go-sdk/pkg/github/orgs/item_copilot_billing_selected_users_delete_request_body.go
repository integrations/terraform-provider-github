package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCopilotBillingSelected_usersDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The usernames of the organization members for which to revoke access to GitHub Copilot.
    selected_usernames []string
}
// NewItemCopilotBillingSelected_usersDeleteRequestBody instantiates a new ItemCopilotBillingSelected_usersDeleteRequestBody and sets the default values.
func NewItemCopilotBillingSelected_usersDeleteRequestBody()(*ItemCopilotBillingSelected_usersDeleteRequestBody) {
    m := &ItemCopilotBillingSelected_usersDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCopilotBillingSelected_usersDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCopilotBillingSelected_usersDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSelected_usersDeleteRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCopilotBillingSelected_usersDeleteRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCopilotBillingSelected_usersDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["selected_usernames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSelectedUsernames(res)
        }
        return nil
    }
    return res
}
// GetSelectedUsernames gets the selected_usernames property value. The usernames of the organization members for which to revoke access to GitHub Copilot.
// returns a []string when successful
func (m *ItemCopilotBillingSelected_usersDeleteRequestBody) GetSelectedUsernames()([]string) {
    return m.selected_usernames
}
// Serialize serializes information the current object
func (m *ItemCopilotBillingSelected_usersDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSelectedUsernames() != nil {
        err := writer.WriteCollectionOfStringValues("selected_usernames", m.GetSelectedUsernames())
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
func (m *ItemCopilotBillingSelected_usersDeleteRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSelectedUsernames sets the selected_usernames property value. The usernames of the organization members for which to revoke access to GitHub Copilot.
func (m *ItemCopilotBillingSelected_usersDeleteRequestBody) SetSelectedUsernames(value []string)() {
    m.selected_usernames = value
}
type ItemCopilotBillingSelected_usersDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSelectedUsernames()([]string)
    SetSelectedUsernames(value []string)()
}
