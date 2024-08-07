package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PullRequestSimple_base struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The label property
    label *string
    // The ref property
    ref *string
    // A repository on GitHub.
    repo Repositoryable
    // The sha property
    sha *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewPullRequestSimple_base instantiates a new PullRequestSimple_base and sets the default values.
func NewPullRequestSimple_base()(*PullRequestSimple_base) {
    m := &PullRequestSimple_base{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestSimple_baseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestSimple_baseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestSimple_base(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequestSimple_base) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestSimple_base) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
        }
        return nil
    }
    res["repo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepo(val.(Repositoryable))
        }
        return nil
    }
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(NullableSimpleUserable))
        }
        return nil
    }
    return res
}
// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *PullRequestSimple_base) GetLabel()(*string) {
    return m.label
}
// GetRef gets the ref property value. The ref property
// returns a *string when successful
func (m *PullRequestSimple_base) GetRef()(*string) {
    return m.ref
}
// GetRepo gets the repo property value. A repository on GitHub.
// returns a Repositoryable when successful
func (m *PullRequestSimple_base) GetRepo()(Repositoryable) {
    return m.repo
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *PullRequestSimple_base) GetSha()(*string) {
    return m.sha
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *PullRequestSimple_base) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *PullRequestSimple_base) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref", m.GetRef())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repo", m.GetRepo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *PullRequestSimple_base) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabel sets the label property value. The label property
func (m *PullRequestSimple_base) SetLabel(value *string)() {
    m.label = value
}
// SetRef sets the ref property value. The ref property
func (m *PullRequestSimple_base) SetRef(value *string)() {
    m.ref = value
}
// SetRepo sets the repo property value. A repository on GitHub.
func (m *PullRequestSimple_base) SetRepo(value Repositoryable)() {
    m.repo = value
}
// SetSha sets the sha property value. The sha property
func (m *PullRequestSimple_base) SetSha(value *string)() {
    m.sha = value
}
// SetUser sets the user property value. A GitHub user.
func (m *PullRequestSimple_base) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type PullRequestSimple_baseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabel()(*string)
    GetRef()(*string)
    GetRepo()(Repositoryable)
    GetSha()(*string)
    GetUser()(NullableSimpleUserable)
    SetLabel(value *string)()
    SetRef(value *string)()
    SetRepo(value Repositoryable)()
    SetSha(value *string)()
    SetUser(value NullableSimpleUserable)()
}
