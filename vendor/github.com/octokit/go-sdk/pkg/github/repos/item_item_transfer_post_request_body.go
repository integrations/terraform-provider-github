package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemTransferPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The new name to be given to the repository.
    new_name *string
    // The username or organization name the repository will be transferred to.
    new_owner *string
    // ID of the team or teams to add to the repository. Teams can only be added to organization-owned repositories.
    team_ids []int32
}
// NewItemItemTransferPostRequestBody instantiates a new ItemItemTransferPostRequestBody and sets the default values.
func NewItemItemTransferPostRequestBody()(*ItemItemTransferPostRequestBody) {
    m := &ItemItemTransferPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemTransferPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemTransferPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemTransferPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemTransferPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemTransferPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["new_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewName(val)
        }
        return nil
    }
    res["new_owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewOwner(val)
        }
        return nil
    }
    res["team_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetTeamIds(res)
        }
        return nil
    }
    return res
}
// GetNewName gets the new_name property value. The new name to be given to the repository.
// returns a *string when successful
func (m *ItemItemTransferPostRequestBody) GetNewName()(*string) {
    return m.new_name
}
// GetNewOwner gets the new_owner property value. The username or organization name the repository will be transferred to.
// returns a *string when successful
func (m *ItemItemTransferPostRequestBody) GetNewOwner()(*string) {
    return m.new_owner
}
// GetTeamIds gets the team_ids property value. ID of the team or teams to add to the repository. Teams can only be added to organization-owned repositories.
// returns a []int32 when successful
func (m *ItemItemTransferPostRequestBody) GetTeamIds()([]int32) {
    return m.team_ids
}
// Serialize serializes information the current object
func (m *ItemItemTransferPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("new_name", m.GetNewName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("new_owner", m.GetNewOwner())
        if err != nil {
            return err
        }
    }
    if m.GetTeamIds() != nil {
        err := writer.WriteCollectionOfInt32Values("team_ids", m.GetTeamIds())
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
func (m *ItemItemTransferPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewName sets the new_name property value. The new name to be given to the repository.
func (m *ItemItemTransferPostRequestBody) SetNewName(value *string)() {
    m.new_name = value
}
// SetNewOwner sets the new_owner property value. The username or organization name the repository will be transferred to.
func (m *ItemItemTransferPostRequestBody) SetNewOwner(value *string)() {
    m.new_owner = value
}
// SetTeamIds sets the team_ids property value. ID of the team or teams to add to the repository. Teams can only be added to organization-owned repositories.
func (m *ItemItemTransferPostRequestBody) SetTeamIds(value []int32)() {
    m.team_ids = value
}
type ItemItemTransferPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNewName()(*string)
    GetNewOwner()(*string)
    GetTeamIds()([]int32)
    SetNewName(value *string)()
    SetNewOwner(value *string)()
    SetTeamIds(value []int32)()
}
