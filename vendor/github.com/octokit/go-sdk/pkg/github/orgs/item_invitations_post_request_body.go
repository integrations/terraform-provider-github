package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemInvitationsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // **Required unless you provide `invitee_id`**. Email address of the person you are inviting, which can be an existing GitHub user.
    email *string
    // **Required unless you provide `email`**. GitHub user ID for the person you are inviting.
    invitee_id *int32
    // Specify IDs for the teams you want to invite new members to.
    team_ids []int32
}
// NewItemInvitationsPostRequestBody instantiates a new ItemInvitationsPostRequestBody and sets the default values.
func NewItemInvitationsPostRequestBody()(*ItemInvitationsPostRequestBody) {
    m := &ItemInvitationsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemInvitationsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemInvitationsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInvitationsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemInvitationsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. **Required unless you provide `invitee_id`**. Email address of the person you are inviting, which can be an existing GitHub user.
// returns a *string when successful
func (m *ItemInvitationsPostRequestBody) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemInvitationsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["invitee_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInviteeId(val)
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
// GetInviteeId gets the invitee_id property value. **Required unless you provide `email`**. GitHub user ID for the person you are inviting.
// returns a *int32 when successful
func (m *ItemInvitationsPostRequestBody) GetInviteeId()(*int32) {
    return m.invitee_id
}
// GetTeamIds gets the team_ids property value. Specify IDs for the teams you want to invite new members to.
// returns a []int32 when successful
func (m *ItemInvitationsPostRequestBody) GetTeamIds()([]int32) {
    return m.team_ids
}
// Serialize serializes information the current object
func (m *ItemInvitationsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("invitee_id", m.GetInviteeId())
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
func (m *ItemInvitationsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. **Required unless you provide `invitee_id`**. Email address of the person you are inviting, which can be an existing GitHub user.
func (m *ItemInvitationsPostRequestBody) SetEmail(value *string)() {
    m.email = value
}
// SetInviteeId sets the invitee_id property value. **Required unless you provide `email`**. GitHub user ID for the person you are inviting.
func (m *ItemInvitationsPostRequestBody) SetInviteeId(value *int32)() {
    m.invitee_id = value
}
// SetTeamIds sets the team_ids property value. Specify IDs for the teams you want to invite new members to.
func (m *ItemInvitationsPostRequestBody) SetTeamIds(value []int32)() {
    m.team_ids = value
}
type ItemInvitationsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetInviteeId()(*int32)
    GetTeamIds()([]int32)
    SetEmail(value *string)()
    SetInviteeId(value *int32)()
    SetTeamIds(value []int32)()
}
