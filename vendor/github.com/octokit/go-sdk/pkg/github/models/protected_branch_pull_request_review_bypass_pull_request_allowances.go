package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectedBranchPullRequestReview_bypass_pull_request_allowances allow specific users, teams, or apps to bypass pull request requirements.
type ProtectedBranchPullRequestReview_bypass_pull_request_allowances struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The list of apps allowed to bypass pull request requirements.
    apps []Integrationable
    // The list of teams allowed to bypass pull request requirements.
    teams []Teamable
    // The list of users allowed to bypass pull request requirements.
    users []SimpleUserable
}
// NewProtectedBranchPullRequestReview_bypass_pull_request_allowances instantiates a new ProtectedBranchPullRequestReview_bypass_pull_request_allowances and sets the default values.
func NewProtectedBranchPullRequestReview_bypass_pull_request_allowances()(*ProtectedBranchPullRequestReview_bypass_pull_request_allowances) {
    m := &ProtectedBranchPullRequestReview_bypass_pull_request_allowances{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProtectedBranchPullRequestReview_bypass_pull_request_allowancesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectedBranchPullRequestReview_bypass_pull_request_allowancesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectedBranchPullRequestReview_bypass_pull_request_allowances(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApps gets the apps property value. The list of apps allowed to bypass pull request requirements.
// returns a []Integrationable when successful
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) GetApps()([]Integrationable) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Integrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Integrationable)
                }
            }
            m.SetApps(res)
        }
        return nil
    }
    res["teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Teamable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Teamable)
                }
            }
            m.SetTeams(res)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimpleUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SimpleUserable)
                }
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetTeams gets the teams property value. The list of teams allowed to bypass pull request requirements.
// returns a []Teamable when successful
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) GetTeams()([]Teamable) {
    return m.teams
}
// GetUsers gets the users property value. The list of users allowed to bypass pull request requirements.
// returns a []SimpleUserable when successful
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) GetUsers()([]SimpleUserable) {
    return m.users
}
// Serialize serializes information the current object
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApps sets the apps property value. The list of apps allowed to bypass pull request requirements.
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) SetApps(value []Integrationable)() {
    m.apps = value
}
// SetTeams sets the teams property value. The list of teams allowed to bypass pull request requirements.
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) SetTeams(value []Teamable)() {
    m.teams = value
}
// SetUsers sets the users property value. The list of users allowed to bypass pull request requirements.
func (m *ProtectedBranchPullRequestReview_bypass_pull_request_allowances) SetUsers(value []SimpleUserable)() {
    m.users = value
}
type ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApps()([]Integrationable)
    GetTeams()([]Teamable)
    GetUsers()([]SimpleUserable)
    SetApps(value []Integrationable)()
    SetTeams(value []Teamable)()
    SetUsers(value []SimpleUserable)()
}
