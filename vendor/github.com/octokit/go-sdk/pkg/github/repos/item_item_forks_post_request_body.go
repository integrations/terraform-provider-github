package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemForksPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // When forking from an existing repository, fork with only the default branch.
    default_branch_only *bool
    // When forking from an existing repository, a new name for the fork.
    name *string
    // Optional parameter to specify the organization name if forking into an organization.
    organization *string
}
// NewItemItemForksPostRequestBody instantiates a new ItemItemForksPostRequestBody and sets the default values.
func NewItemItemForksPostRequestBody()(*ItemItemForksPostRequestBody) {
    m := &ItemItemForksPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemForksPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemForksPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemForksPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemForksPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultBranchOnly gets the default_branch_only property value. When forking from an existing repository, fork with only the default branch.
// returns a *bool when successful
func (m *ItemItemForksPostRequestBody) GetDefaultBranchOnly()(*bool) {
    return m.default_branch_only
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemForksPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["default_branch_only"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultBranchOnly(val)
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
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. When forking from an existing repository, a new name for the fork.
// returns a *string when successful
func (m *ItemItemForksPostRequestBody) GetName()(*string) {
    return m.name
}
// GetOrganization gets the organization property value. Optional parameter to specify the organization name if forking into an organization.
// returns a *string when successful
func (m *ItemItemForksPostRequestBody) GetOrganization()(*string) {
    return m.organization
}
// Serialize serializes information the current object
func (m *ItemItemForksPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("default_branch_only", m.GetDefaultBranchOnly())
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
        err := writer.WriteStringValue("organization", m.GetOrganization())
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
func (m *ItemItemForksPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultBranchOnly sets the default_branch_only property value. When forking from an existing repository, fork with only the default branch.
func (m *ItemItemForksPostRequestBody) SetDefaultBranchOnly(value *bool)() {
    m.default_branch_only = value
}
// SetName sets the name property value. When forking from an existing repository, a new name for the fork.
func (m *ItemItemForksPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetOrganization sets the organization property value. Optional parameter to specify the organization name if forking into an organization.
func (m *ItemItemForksPostRequestBody) SetOrganization(value *string)() {
    m.organization = value
}
type ItemItemForksPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultBranchOnly()(*bool)
    GetName()(*string)
    GetOrganization()(*string)
    SetDefaultBranchOnly(value *bool)()
    SetName(value *string)()
    SetOrganization(value *string)()
}
