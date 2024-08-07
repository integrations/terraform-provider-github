package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemTeamsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of the team.
    description *string
    // List GitHub IDs for organization members who will become team maintainers.
    maintainers []string
    // The name of the team.
    name *string
    // The ID of a team to set as the parent team.
    parent_team_id *int32
    // The full name (e.g., "organization-name/repository-name") of repositories to add the team to.
    repo_names []string
}
// NewItemTeamsPostRequestBody instantiates a new ItemTeamsPostRequestBody and sets the default values.
func NewItemTeamsPostRequestBody()(*ItemTeamsPostRequestBody) {
    m := &ItemTeamsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemTeamsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of the team.
// returns a *string when successful
func (m *ItemTeamsPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemTeamsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["maintainers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetMaintainers(res)
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
    res["repo_names"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRepoNames(res)
        }
        return nil
    }
    return res
}
// GetMaintainers gets the maintainers property value. List GitHub IDs for organization members who will become team maintainers.
// returns a []string when successful
func (m *ItemTeamsPostRequestBody) GetMaintainers()([]string) {
    return m.maintainers
}
// GetName gets the name property value. The name of the team.
// returns a *string when successful
func (m *ItemTeamsPostRequestBody) GetName()(*string) {
    return m.name
}
// GetParentTeamId gets the parent_team_id property value. The ID of a team to set as the parent team.
// returns a *int32 when successful
func (m *ItemTeamsPostRequestBody) GetParentTeamId()(*int32) {
    return m.parent_team_id
}
// GetRepoNames gets the repo_names property value. The full name (e.g., "organization-name/repository-name") of repositories to add the team to.
// returns a []string when successful
func (m *ItemTeamsPostRequestBody) GetRepoNames()([]string) {
    return m.repo_names
}
// Serialize serializes information the current object
func (m *ItemTeamsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetMaintainers() != nil {
        err := writer.WriteCollectionOfStringValues("maintainers", m.GetMaintainers())
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
    if m.GetRepoNames() != nil {
        err := writer.WriteCollectionOfStringValues("repo_names", m.GetRepoNames())
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
func (m *ItemTeamsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of the team.
func (m *ItemTeamsPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetMaintainers sets the maintainers property value. List GitHub IDs for organization members who will become team maintainers.
func (m *ItemTeamsPostRequestBody) SetMaintainers(value []string)() {
    m.maintainers = value
}
// SetName sets the name property value. The name of the team.
func (m *ItemTeamsPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetParentTeamId sets the parent_team_id property value. The ID of a team to set as the parent team.
func (m *ItemTeamsPostRequestBody) SetParentTeamId(value *int32)() {
    m.parent_team_id = value
}
// SetRepoNames sets the repo_names property value. The full name (e.g., "organization-name/repository-name") of repositories to add the team to.
func (m *ItemTeamsPostRequestBody) SetRepoNames(value []string)() {
    m.repo_names = value
}
type ItemTeamsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetMaintainers()([]string)
    GetName()(*string)
    GetParentTeamId()(*int32)
    GetRepoNames()([]string)
    SetDescription(value *string)()
    SetMaintainers(value []string)()
    SetName(value *string)()
    SetParentTeamId(value *int32)()
    SetRepoNames(value []string)()
}
