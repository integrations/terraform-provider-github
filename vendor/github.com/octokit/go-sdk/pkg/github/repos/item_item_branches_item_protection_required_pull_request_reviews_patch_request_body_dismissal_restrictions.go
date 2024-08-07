package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions specify which users, teams, and apps can dismiss pull request reviews. Pass an empty `dismissal_restrictions` object to disable. User and team `dismissal_restrictions` are only available for organization-owned repositories. Omit this parameter for personal repositories.
type ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The list of app `slug`s with dismissal access
    apps []string
    // The list of team `slug`s with dismissal access
    teams []string
    // The list of user `login`s with dismissal access
    users []string
}
// NewItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions instantiates a new ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions and sets the default values.
func NewItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions()(*ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) {
    m := &ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApps gets the apps property value. The list of app `slug`s with dismissal access
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) GetApps()([]string) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetTeams gets the teams property value. The list of team `slug`s with dismissal access
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) GetTeams()([]string) {
    return m.teams
}
// GetUsers gets the users property value. The list of user `login`s with dismissal access
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) GetUsers()([]string) {
    return m.users
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApps sets the apps property value. The list of app `slug`s with dismissal access
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) SetApps(value []string)() {
    m.apps = value
}
// SetTeams sets the teams property value. The list of team `slug`s with dismissal access
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) SetTeams(value []string)() {
    m.teams = value
}
// SetUsers sets the users property value. The list of user `login`s with dismissal access
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictions) SetUsers(value []string)() {
    m.users = value
}
type ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApps()([]string)
    GetTeams()([]string)
    GetUsers()([]string)
    SetApps(value []string)()
    SetTeams(value []string)()
    SetUsers(value []string)()
}
