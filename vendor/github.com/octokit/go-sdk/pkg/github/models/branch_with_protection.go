package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchWithProtection branch With Protection
type BranchWithProtection struct {
    // The _links property
    _links BranchWithProtection__linksable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Commit
    commit Commitable
    // The name property
    name *string
    // The pattern property
    pattern *string
    // The protected property
    protected *bool
    // Branch Protection
    protection BranchProtectionable
    // The protection_url property
    protection_url *string
    // The required_approving_review_count property
    required_approving_review_count *int32
}
// NewBranchWithProtection instantiates a new BranchWithProtection and sets the default values.
func NewBranchWithProtection()(*BranchWithProtection) {
    m := &BranchWithProtection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBranchWithProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBranchWithProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBranchWithProtection(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BranchWithProtection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommit gets the commit property value. Commit
// returns a Commitable when successful
func (m *BranchWithProtection) GetCommit()(Commitable) {
    return m.commit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BranchWithProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchWithProtection__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(BranchWithProtection__linksable))
        }
        return nil
    }
    res["commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommit(val.(Commitable))
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
    res["pattern"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPattern(val)
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
    res["required_approving_review_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredApprovingReviewCount(val)
        }
        return nil
    }
    return res
}
// GetLinks gets the _links property value. The _links property
// returns a BranchWithProtection__linksable when successful
func (m *BranchWithProtection) GetLinks()(BranchWithProtection__linksable) {
    return m._links
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *BranchWithProtection) GetName()(*string) {
    return m.name
}
// GetPattern gets the pattern property value. The pattern property
// returns a *string when successful
func (m *BranchWithProtection) GetPattern()(*string) {
    return m.pattern
}
// GetProtected gets the protected property value. The protected property
// returns a *bool when successful
func (m *BranchWithProtection) GetProtected()(*bool) {
    return m.protected
}
// GetProtection gets the protection property value. Branch Protection
// returns a BranchProtectionable when successful
func (m *BranchWithProtection) GetProtection()(BranchProtectionable) {
    return m.protection
}
// GetProtectionUrl gets the protection_url property value. The protection_url property
// returns a *string when successful
func (m *BranchWithProtection) GetProtectionUrl()(*string) {
    return m.protection_url
}
// GetRequiredApprovingReviewCount gets the required_approving_review_count property value. The required_approving_review_count property
// returns a *int32 when successful
func (m *BranchWithProtection) GetRequiredApprovingReviewCount()(*int32) {
    return m.required_approving_review_count
}
// Serialize serializes information the current object
func (m *BranchWithProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("pattern", m.GetPattern())
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
        err := writer.WriteInt32Value("required_approving_review_count", m.GetRequiredApprovingReviewCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("_links", m.GetLinks())
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
func (m *BranchWithProtection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommit sets the commit property value. Commit
func (m *BranchWithProtection) SetCommit(value Commitable)() {
    m.commit = value
}
// SetLinks sets the _links property value. The _links property
func (m *BranchWithProtection) SetLinks(value BranchWithProtection__linksable)() {
    m._links = value
}
// SetName sets the name property value. The name property
func (m *BranchWithProtection) SetName(value *string)() {
    m.name = value
}
// SetPattern sets the pattern property value. The pattern property
func (m *BranchWithProtection) SetPattern(value *string)() {
    m.pattern = value
}
// SetProtected sets the protected property value. The protected property
func (m *BranchWithProtection) SetProtected(value *bool)() {
    m.protected = value
}
// SetProtection sets the protection property value. Branch Protection
func (m *BranchWithProtection) SetProtection(value BranchProtectionable)() {
    m.protection = value
}
// SetProtectionUrl sets the protection_url property value. The protection_url property
func (m *BranchWithProtection) SetProtectionUrl(value *string)() {
    m.protection_url = value
}
// SetRequiredApprovingReviewCount sets the required_approving_review_count property value. The required_approving_review_count property
func (m *BranchWithProtection) SetRequiredApprovingReviewCount(value *int32)() {
    m.required_approving_review_count = value
}
type BranchWithProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommit()(Commitable)
    GetLinks()(BranchWithProtection__linksable)
    GetName()(*string)
    GetPattern()(*string)
    GetProtected()(*bool)
    GetProtection()(BranchProtectionable)
    GetProtectionUrl()(*string)
    GetRequiredApprovingReviewCount()(*int32)
    SetCommit(value Commitable)()
    SetLinks(value BranchWithProtection__linksable)()
    SetName(value *string)()
    SetPattern(value *string)()
    SetProtected(value *bool)()
    SetProtection(value BranchProtectionable)()
    SetProtectionUrl(value *string)()
    SetRequiredApprovingReviewCount(value *int32)()
}
