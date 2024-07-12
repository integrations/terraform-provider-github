package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemTeamsItemWithTeam_slugPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of the team.
    description *string
    // The name of the team.
    name *string
    // The ID of a team to set as the parent team.
    parent_team_id *int32
}
// NewItemTeamsItemWithTeam_slugPatchRequestBody instantiates a new ItemTeamsItemWithTeam_slugPatchRequestBody and sets the default values.
func NewItemTeamsItemWithTeam_slugPatchRequestBody()(*ItemTeamsItemWithTeam_slugPatchRequestBody) {
    m := &ItemTeamsItemWithTeam_slugPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamsItemWithTeam_slugPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamsItemWithTeam_slugPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamsItemWithTeam_slugPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of the team.
// returns a *string when successful
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["parent_team_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentTeamId(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the team.
// returns a *string when successful
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) GetName()(*string) {
    return m.name
}
// GetParentTeamId gets the parent_team_id property value. The ID of a team to set as the parent team.
// returns a *int32 when successful
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) GetParentTeamId()(*int32) {
    return m.parent_team_id
}
// Serialize serializes information the current object
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteInt32Value("parent_team_id", m.GetParentTeamId())
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
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of the team.
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name of the team.
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetParentTeamId sets the parent_team_id property value. The ID of a team to set as the parent team.
func (m *ItemTeamsItemWithTeam_slugPatchRequestBody) SetParentTeamId(value *int32)() {
    m.parent_team_id = value
}
type ItemTeamsItemWithTeam_slugPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetName()(*string)
    GetParentTeamId()(*int32)
    SetDescription(value *string)()
    SetName(value *string)()
    SetParentTeamId(value *int32)()
}
