package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchRestrictionPolicy branch Restriction Policy
type BranchRestrictionPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The apps property
    apps []BranchRestrictionPolicy_appsable
    // The apps_url property
    apps_url *string
    // The teams property
    teams []BranchRestrictionPolicy_teamsable
    // The teams_url property
    teams_url *string
    // The url property
    url *string
    // The users property
    users []BranchRestrictionPolicy_usersable
    // The users_url property
    users_url *string
}
// NewBranchRestrictionPolicy instantiates a new BranchRestrictionPolicy and sets the default values.
func NewBranchRestrictionPolicy()(*BranchRestrictionPolicy) {
    m := &BranchRestrictionPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBranchRestrictionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBranchRestrictionPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBranchRestrictionPolicy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BranchRestrictionPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApps gets the apps property value. The apps property
// returns a []BranchRestrictionPolicy_appsable when successful
func (m *BranchRestrictionPolicy) GetApps()([]BranchRestrictionPolicy_appsable) {
    return m.apps
}
// GetAppsUrl gets the apps_url property value. The apps_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy) GetAppsUrl()(*string) {
    return m.apps_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BranchRestrictionPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBranchRestrictionPolicy_appsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BranchRestrictionPolicy_appsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BranchRestrictionPolicy_appsable)
                }
            }
            m.SetApps(res)
        }
        return nil
    }
    res["apps_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsUrl(val)
        }
        return nil
    }
    res["teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBranchRestrictionPolicy_teamsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BranchRestrictionPolicy_teamsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BranchRestrictionPolicy_teamsable)
                }
            }
            m.SetTeams(res)
        }
        return nil
    }
    res["teams_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsUrl(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBranchRestrictionPolicy_usersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BranchRestrictionPolicy_usersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BranchRestrictionPolicy_usersable)
                }
            }
            m.SetUsers(res)
        }
        return nil
    }
    res["users_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersUrl(val)
        }
        return nil
    }
    return res
}
// GetTeams gets the teams property value. The teams property
// returns a []BranchRestrictionPolicy_teamsable when successful
func (m *BranchRestrictionPolicy) GetTeams()([]BranchRestrictionPolicy_teamsable) {
    return m.teams
}
// GetTeamsUrl gets the teams_url property value. The teams_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy) GetTeamsUrl()(*string) {
    return m.teams_url
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *BranchRestrictionPolicy) GetUrl()(*string) {
    return m.url
}
// GetUsers gets the users property value. The users property
// returns a []BranchRestrictionPolicy_usersable when successful
func (m *BranchRestrictionPolicy) GetUsers()([]BranchRestrictionPolicy_usersable) {
    return m.users
}
// GetUsersUrl gets the users_url property value. The users_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy) GetUsersUrl()(*string) {
    return m.users_url
}
// Serialize serializes information the current object
func (m *BranchRestrictionPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("apps_url", m.GetAppsUrl())
        if err != nil {
            return err
        }
    }
    if m.GetTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTeams()))
        for i, v := range m.GetTeams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("teams", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teams_url", m.GetTeamsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("users_url", m.GetUsersUrl())
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
func (m *BranchRestrictionPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApps sets the apps property value. The apps property
func (m *BranchRestrictionPolicy) SetApps(value []BranchRestrictionPolicy_appsable)() {
    m.apps = value
}
// SetAppsUrl sets the apps_url property value. The apps_url property
func (m *BranchRestrictionPolicy) SetAppsUrl(value *string)() {
    m.apps_url = value
}
// SetTeams sets the teams property value. The teams property
func (m *BranchRestrictionPolicy) SetTeams(value []BranchRestrictionPolicy_teamsable)() {
    m.teams = value
}
// SetTeamsUrl sets the teams_url property value. The teams_url property
func (m *BranchRestrictionPolicy) SetTeamsUrl(value *string)() {
    m.teams_url = value
}
// SetUrl sets the url property value. The url property
func (m *BranchRestrictionPolicy) SetUrl(value *string)() {
    m.url = value
}
// SetUsers sets the users property value. The users property
func (m *BranchRestrictionPolicy) SetUsers(value []BranchRestrictionPolicy_usersable)() {
    m.users = value
}
// SetUsersUrl sets the users_url property value. The users_url property
func (m *BranchRestrictionPolicy) SetUsersUrl(value *string)() {
    m.users_url = value
}
type BranchRestrictionPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApps()([]BranchRestrictionPolicy_appsable)
    GetAppsUrl()(*string)
    GetTeams()([]BranchRestrictionPolicy_teamsable)
    GetTeamsUrl()(*string)
    GetUrl()(*string)
    GetUsers()([]BranchRestrictionPolicy_usersable)
    GetUsersUrl()(*string)
    SetApps(value []BranchRestrictionPolicy_appsable)()
    SetAppsUrl(value *string)()
    SetTeams(value []BranchRestrictionPolicy_teamsable)()
    SetTeamsUrl(value *string)()
    SetUrl(value *string)()
    SetUsers(value []BranchRestrictionPolicy_usersable)()
    SetUsersUrl(value *string)()
}
