package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SelectedActions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
    github_owned_allowed *bool
    // Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`.**Note**: The `patterns_allowed` setting only applies to public repositories.
    patterns_allowed []string
    // Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
    verified_allowed *bool
}
// NewSelectedActions instantiates a new SelectedActions and sets the default values.
func NewSelectedActions()(*SelectedActions) {
    m := &SelectedActions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSelectedActionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSelectedActionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSelectedActions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SelectedActions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SelectedActions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["github_owned_allowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubOwnedAllowed(val)
        }
        return nil
    }
    res["patterns_allowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPatternsAllowed(res)
        }
        return nil
    }
    res["verified_allowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedAllowed(val)
        }
        return nil
    }
    return res
}
// GetGithubOwnedAllowed gets the github_owned_allowed property value. Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
// returns a *bool when successful
func (m *SelectedActions) GetGithubOwnedAllowed()(*bool) {
    return m.github_owned_allowed
}
// GetPatternsAllowed gets the patterns_allowed property value. Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`.**Note**: The `patterns_allowed` setting only applies to public repositories.
// returns a []string when successful
func (m *SelectedActions) GetPatternsAllowed()([]string) {
    return m.patterns_allowed
}
// GetVerifiedAllowed gets the verified_allowed property value. Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
// returns a *bool when successful
func (m *SelectedActions) GetVerifiedAllowed()(*bool) {
    return m.verified_allowed
}
// Serialize serializes information the current object
func (m *SelectedActions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("github_owned_allowed", m.GetGithubOwnedAllowed())
        if err != nil {
            return err
        }
    }
    if m.GetPatternsAllowed() != nil {
        err := writer.WriteCollectionOfStringValues("patterns_allowed", m.GetPatternsAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("verified_allowed", m.GetVerifiedAllowed())
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
func (m *SelectedActions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGithubOwnedAllowed sets the github_owned_allowed property value. Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
func (m *SelectedActions) SetGithubOwnedAllowed(value *bool)() {
    m.github_owned_allowed = value
}
// SetPatternsAllowed sets the patterns_allowed property value. Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`.**Note**: The `patterns_allowed` setting only applies to public repositories.
func (m *SelectedActions) SetPatternsAllowed(value []string)() {
    m.patterns_allowed = value
}
// SetVerifiedAllowed sets the verified_allowed property value. Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
func (m *SelectedActions) SetVerifiedAllowed(value *bool)() {
    m.verified_allowed = value
}
type SelectedActionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGithubOwnedAllowed()(*bool)
    GetPatternsAllowed()([]string)
    GetVerifiedAllowed()(*bool)
    SetGithubOwnedAllowed(value *bool)()
    SetPatternsAllowed(value []string)()
    SetVerifiedAllowed(value *bool)()
}
