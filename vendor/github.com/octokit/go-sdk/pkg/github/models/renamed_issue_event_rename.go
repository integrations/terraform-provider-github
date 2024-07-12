package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RenamedIssueEvent_rename struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The from property
    from *string
    // The to property
    to *string
}
// NewRenamedIssueEvent_rename instantiates a new RenamedIssueEvent_rename and sets the default values.
func NewRenamedIssueEvent_rename()(*RenamedIssueEvent_rename) {
    m := &RenamedIssueEvent_rename{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRenamedIssueEvent_renameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRenamedIssueEvent_renameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRenamedIssueEvent_rename(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RenamedIssueEvent_rename) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RenamedIssueEvent_rename) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["from"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val)
        }
        return nil
    }
    res["to"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTo(val)
        }
        return nil
    }
    return res
}
// GetFrom gets the from property value. The from property
// returns a *string when successful
func (m *RenamedIssueEvent_rename) GetFrom()(*string) {
    return m.from
}
// GetTo gets the to property value. The to property
// returns a *string when successful
func (m *RenamedIssueEvent_rename) GetTo()(*string) {
    return m.to
}
// Serialize serializes information the current object
func (m *RenamedIssueEvent_rename) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("to", m.GetTo())
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
func (m *RenamedIssueEvent_rename) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFrom sets the from property value. The from property
func (m *RenamedIssueEvent_rename) SetFrom(value *string)() {
    m.from = value
}
// SetTo sets the to property value. The to property
func (m *RenamedIssueEvent_rename) SetTo(value *string)() {
    m.to = value
}
type RenamedIssueEvent_renameable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFrom()(*string)
    GetTo()(*string)
    SetFrom(value *string)()
    SetTo(value *string)()
}
