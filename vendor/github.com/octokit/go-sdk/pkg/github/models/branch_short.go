package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchShort branch Short
type BranchShort struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commit property
    commit BranchShort_commitable
    // The name property
    name *string
    // The protected property
    protected *bool
}
// NewBranchShort instantiates a new BranchShort and sets the default values.
func NewBranchShort()(*BranchShort) {
    m := &BranchShort{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBranchShortFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBranchShortFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBranchShort(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BranchShort) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommit gets the commit property value. The commit property
// returns a BranchShort_commitable when successful
func (m *BranchShort) GetCommit()(BranchShort_commitable) {
    return m.commit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BranchShort) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchShort_commitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommit(val.(BranchShort_commitable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["protected"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtected(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *BranchShort) GetName()(*string) {
    return m.name
}
// GetProtected gets the protected property value. The protected property
// returns a *bool when successful
func (m *BranchShort) GetProtected()(*bool) {
    return m.protected
}
// Serialize serializes information the current object
func (m *BranchShort) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("commit", m.GetCommit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("protected", m.GetProtected())
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
func (m *BranchShort) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommit sets the commit property value. The commit property
func (m *BranchShort) SetCommit(value BranchShort_commitable)() {
    m.commit = value
}
// SetName sets the name property value. The name property
func (m *BranchShort) SetName(value *string)() {
    m.name = value
}
// SetProtected sets the protected property value. The protected property
func (m *BranchShort) SetProtected(value *bool)() {
    m.protected = value
}
type BranchShortable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommit()(BranchShort_commitable)
    GetName()(*string)
    GetProtected()(*bool)
    SetCommit(value BranchShort_commitable)()
    SetName(value *string)()
    SetProtected(value *bool)()
}
