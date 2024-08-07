package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GistCommit_change_status struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The additions property
    additions *int32
    // The deletions property
    deletions *int32
    // The total property
    total *int32
}
// NewGistCommit_change_status instantiates a new GistCommit_change_status and sets the default values.
func NewGistCommit_change_status()(*GistCommit_change_status) {
    m := &GistCommit_change_status{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistCommit_change_statusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGistCommit_change_statusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGistCommit_change_status(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GistCommit_change_status) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdditions gets the additions property value. The additions property
// returns a *int32 when successful
func (m *GistCommit_change_status) GetAdditions()(*int32) {
    return m.additions
}
// GetDeletions gets the deletions property value. The deletions property
// returns a *int32 when successful
func (m *GistCommit_change_status) GetDeletions()(*int32) {
    return m.deletions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GistCommit_change_status) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditions(val)
        }
        return nil
    }
    res["deletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletions(val)
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
// GetTotal gets the total property value. The total property
// returns a *int32 when successful
func (m *GistCommit_change_status) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *GistCommit_change_status) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("additions", m.GetAdditions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deletions", m.GetDeletions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
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
func (m *GistCommit_change_status) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdditions sets the additions property value. The additions property
func (m *GistCommit_change_status) SetAdditions(value *int32)() {
    m.additions = value
}
// SetDeletions sets the deletions property value. The deletions property
func (m *GistCommit_change_status) SetDeletions(value *int32)() {
    m.deletions = value
}
// SetTotal sets the total property value. The total property
func (m *GistCommit_change_status) SetTotal(value *int32)() {
    m.total = value
}
type GistCommit_change_statusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditions()(*int32)
    GetDeletions()(*int32)
    GetTotal()(*int32)
    SetAdditions(value *int32)()
    SetDeletions(value *int32)()
    SetTotal(value *int32)()
}
