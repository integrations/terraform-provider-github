package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCopilotBillingSelected_teamsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of team names within the organization to which to grant access to GitHub Copilot.
    selected_teams []string
}
// NewItemCopilotBillingSelected_teamsPostRequestBody instantiates a new ItemCopilotBillingSelected_teamsPostRequestBody and sets the default values.
func NewItemCopilotBillingSelected_teamsPostRequestBody()(*ItemCopilotBillingSelected_teamsPostRequestBody) {
    m := &ItemCopilotBillingSelected_teamsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCopilotBillingSelected_teamsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCopilotBillingSelected_teamsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSelected_teamsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCopilotBillingSelected_teamsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCopilotBillingSelected_teamsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["selected_teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSelectedTeams(res)
        }
        return nil
    }
    return res
}
// GetSelectedTeams gets the selected_teams property value. List of team names within the organization to which to grant access to GitHub Copilot.
// returns a []string when successful
func (m *ItemCopilotBillingSelected_teamsPostRequestBody) GetSelectedTeams()([]string) {
    return m.selected_teams
}
// Serialize serializes information the current object
func (m *ItemCopilotBillingSelected_teamsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSelectedTeams() != nil {
        err := writer.WriteCollectionOfStringValues("selected_teams", m.GetSelectedTeams())
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
func (m *ItemCopilotBillingSelected_teamsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSelectedTeams sets the selected_teams property value. List of team names within the organization to which to grant access to GitHub Copilot.
func (m *ItemCopilotBillingSelected_teamsPostRequestBody) SetSelectedTeams(value []string)() {
    m.selected_teams = value
}
type ItemCopilotBillingSelected_teamsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSelectedTeams()([]string)
    SetSelectedTeams(value []string)()
}
