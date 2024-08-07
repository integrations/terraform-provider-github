package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PullRequestMinimal_base struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ref property
    ref *string
    // The repo property
    repo PullRequestMinimal_base_repoable
    // The sha property
    sha *string
}
// NewPullRequestMinimal_base instantiates a new PullRequestMinimal_base and sets the default values.
func NewPullRequestMinimal_base()(*PullRequestMinimal_base) {
    m := &PullRequestMinimal_base{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestMinimal_baseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestMinimal_baseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestMinimal_base(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequestMinimal_base) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestMinimal_base) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetObjectValue(CreatePullRequestMinimal_base_repoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepo(val.(PullRequestMinimal_base_repoable))
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
    return res
}
// GetRef gets the ref property value. The ref property
// returns a *string when successful
func (m *PullRequestMinimal_base) GetRef()(*string) {
    return m.ref
}
// GetRepo gets the repo property value. The repo property
// returns a PullRequestMinimal_base_repoable when successful
func (m *PullRequestMinimal_base) GetRepo()(PullRequestMinimal_base_repoable) {
    return m.repo
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *PullRequestMinimal_base) GetSha()(*string) {
    return m.sha
}
// Serialize serializes information the current object
func (m *PullRequestMinimal_base) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PullRequestMinimal_base) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRef sets the ref property value. The ref property
func (m *PullRequestMinimal_base) SetRef(value *string)() {
    m.ref = value
}
// SetRepo sets the repo property value. The repo property
func (m *PullRequestMinimal_base) SetRepo(value PullRequestMinimal_base_repoable)() {
    m.repo = value
}
// SetSha sets the sha property value. The sha property
func (m *PullRequestMinimal_base) SetSha(value *string)() {
    m.sha = value
}
type PullRequestMinimal_baseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRef()(*string)
    GetRepo()(PullRequestMinimal_base_repoable)
    GetSha()(*string)
    SetRef(value *string)()
    SetRepo(value PullRequestMinimal_base_repoable)()
    SetSha(value *string)()
}
