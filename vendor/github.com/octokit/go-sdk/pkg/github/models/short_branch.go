package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ShortBranch short Branch
type ShortBranch struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commit property
    commit ShortBranch_commitable
    // The name property
    name *string
    // The protected property
    protected *bool
    // Branch Protection
    protection BranchProtectionable
    // The protection_url property
    protection_url *string
}
// NewShortBranch instantiates a new ShortBranch and sets the default values.
func NewShortBranch()(*ShortBranch) {
    m := &ShortBranch{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateShortBranchFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateShortBranchFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewShortBranch(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ShortBranch) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommit gets the commit property value. The commit property
// returns a ShortBranch_commitable when successful
func (m *ShortBranch) GetCommit()(ShortBranch_commitable) {
    return m.commit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ShortBranch) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShortBranch_commitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommit(val.(ShortBranch_commitable))
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
    res["protection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtection(val.(BranchProtectionable))
        }
        return nil
    }
    res["protection_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionUrl(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ShortBranch) GetName()(*string) {
    return m.name
}
// GetProtected gets the protected property value. The protected property
// returns a *bool when successful
func (m *ShortBranch) GetProtected()(*bool) {
    return m.protected
}
// GetProtection gets the protection property value. Branch Protection
// returns a BranchProtectionable when successful
func (m *ShortBranch) GetProtection()(BranchProtectionable) {
    return m.protection
}
// GetProtectionUrl gets the protection_url property value. The protection_url property
// returns a *string when successful
func (m *ShortBranch) GetProtectionUrl()(*string) {
    return m.protection_url
}
// Serialize serializes information the current object
func (m *ShortBranch) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("protection", m.GetProtection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("protection_url", m.GetProtectionUrl())
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
func (m *ShortBranch) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommit sets the commit property value. The commit property
func (m *ShortBranch) SetCommit(value ShortBranch_commitable)() {
    m.commit = value
}
// SetName sets the name property value. The name property
func (m *ShortBranch) SetName(value *string)() {
    m.name = value
}
// SetProtected sets the protected property value. The protected property
func (m *ShortBranch) SetProtected(value *bool)() {
    m.protected = value
}
// SetProtection sets the protection property value. Branch Protection
func (m *ShortBranch) SetProtection(value BranchProtectionable)() {
    m.protection = value
}
// SetProtectionUrl sets the protection_url property value. The protection_url property
func (m *ShortBranch) SetProtectionUrl(value *string)() {
    m.protection_url = value
}
type ShortBranchable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommit()(ShortBranch_commitable)
    GetName()(*string)
    GetProtected()(*bool)
    GetProtection()(BranchProtectionable)
    GetProtectionUrl()(*string)
    SetCommit(value ShortBranch_commitable)()
    SetName(value *string)()
    SetProtected(value *bool)()
    SetProtection(value BranchProtectionable)()
    SetProtectionUrl(value *string)()
}
