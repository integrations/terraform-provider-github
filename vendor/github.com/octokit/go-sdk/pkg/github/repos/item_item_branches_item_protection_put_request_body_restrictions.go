package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionPutRequestBody_restrictions restrict who can push to the protected branch. User, app, and team `restrictions` are only available for organization-owned repositories. Set to `null` to disable.
type ItemItemBranchesItemProtectionPutRequestBody_restrictions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The list of app `slug`s with push access
    apps []string
    // The list of team `slug`s with push access
    teams []string
    // The list of user `login`s with push access
    users []string
}
// NewItemItemBranchesItemProtectionPutRequestBody_restrictions instantiates a new ItemItemBranchesItemProtectionPutRequestBody_restrictions and sets the default values.
func NewItemItemBranchesItemProtectionPutRequestBody_restrictions()(*ItemItemBranchesItemProtectionPutRequestBody_restrictions) {
    m := &ItemItemBranchesItemProtectionPutRequestBody_restrictions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionPutRequestBody_restrictionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBranchesItemProtectionPutRequestBody_restrictionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionPutRequestBody_restrictions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApps gets the apps property value. The list of app `slug`s with push access
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) GetApps()([]string) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetApps(res)
        }
        return nil
    }
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
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetTeams gets the teams property value. The list of team `slug`s with push access
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) GetTeams()([]string) {
    return m.teams
}
// GetUsers gets the users property value. The list of user `login`s with push access
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) GetUsers()([]string) {
    return m.users
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApps() != nil {
        err := writer.WriteCollectionOfStringValues("apps", m.GetApps())
        if err != nil {
            return err
        }
    }
    if m.GetTeams() != nil {
        err := writer.WriteCollectionOfStringValues("teams", m.GetTeams())
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        err := writer.WriteCollectionOfStringValues("users", m.GetUsers())
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
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApps sets the apps property value. The list of app `slug`s with push access
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) SetApps(value []string)() {
    m.apps = value
}
// SetTeams sets the teams property value. The list of team `slug`s with push access
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) SetTeams(value []string)() {
    m.teams = value
}
// SetUsers sets the users property value. The list of user `login`s with push access
func (m *ItemItemBranchesItemProtectionPutRequestBody_restrictions) SetUsers(value []string)() {
    m.users = value
}
type ItemItemBranchesItemProtectionPutRequestBody_restrictionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApps()([]string)
    GetTeams()([]string)
    GetUsers()([]string)
    SetApps(value []string)()
    SetTeams(value []string)()
    SetUsers(value []string)()
}
