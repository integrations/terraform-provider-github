package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The slug values for teams
    teams []string
}
// NewItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 instantiates a new ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1()(*ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1) {
    m := &ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTeams(res)
        }
        return nil
    }
    return res
}
// GetTeams gets the teams property value. The slug values for teams
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1) GetTeams()([]string) {
    return m.teams
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTeams() != nil {
        err := writer.WriteCollectionOfStringValues("teams", m.GetTeams())
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
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTeams sets the teams property value. The slug values for teams
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1) SetTeams(value []string)() {
    m.teams = value
}
type ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTeams()([]string)
    SetTeams(value []string)()
}
