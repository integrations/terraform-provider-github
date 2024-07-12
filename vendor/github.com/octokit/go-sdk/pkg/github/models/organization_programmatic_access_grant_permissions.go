package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationProgrammaticAccessGrant_permissions permissions requested, categorized by type of permission.
type OrganizationProgrammaticAccessGrant_permissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The organization property
    organization OrganizationProgrammaticAccessGrant_permissions_organizationable
    // The other property
    other OrganizationProgrammaticAccessGrant_permissions_otherable
    // The repository property
    repository OrganizationProgrammaticAccessGrant_permissions_repositoryable
}
// NewOrganizationProgrammaticAccessGrant_permissions instantiates a new OrganizationProgrammaticAccessGrant_permissions and sets the default values.
func NewOrganizationProgrammaticAccessGrant_permissions()(*OrganizationProgrammaticAccessGrant_permissions) {
    m := &OrganizationProgrammaticAccessGrant_permissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationProgrammaticAccessGrant_permissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationProgrammaticAccessGrant_permissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationProgrammaticAccessGrant_permissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationProgrammaticAccessGrant_permissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationProgrammaticAccessGrant_permissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationProgrammaticAccessGrant_permissions_organizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val.(OrganizationProgrammaticAccessGrant_permissions_organizationable))
        }
        return nil
    }
    res["other"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationProgrammaticAccessGrant_permissions_otherFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOther(val.(OrganizationProgrammaticAccessGrant_permissions_otherable))
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationProgrammaticAccessGrant_permissions_repositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(OrganizationProgrammaticAccessGrant_permissions_repositoryable))
        }
        return nil
    }
    return res
}
// GetOrganization gets the organization property value. The organization property
// returns a OrganizationProgrammaticAccessGrant_permissions_organizationable when successful
func (m *OrganizationProgrammaticAccessGrant_permissions) GetOrganization()(OrganizationProgrammaticAccessGrant_permissions_organizationable) {
    return m.organization
}
// GetOther gets the other property value. The other property
// returns a OrganizationProgrammaticAccessGrant_permissions_otherable when successful
func (m *OrganizationProgrammaticAccessGrant_permissions) GetOther()(OrganizationProgrammaticAccessGrant_permissions_otherable) {
    return m.other
}
// GetRepository gets the repository property value. The repository property
// returns a OrganizationProgrammaticAccessGrant_permissions_repositoryable when successful
func (m *OrganizationProgrammaticAccessGrant_permissions) GetRepository()(OrganizationProgrammaticAccessGrant_permissions_repositoryable) {
    return m.repository
}
// Serialize serializes information the current object
func (m *OrganizationProgrammaticAccessGrant_permissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("organization", m.GetOrganization())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("other", m.GetOther())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
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
func (m *OrganizationProgrammaticAccessGrant_permissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOrganization sets the organization property value. The organization property
func (m *OrganizationProgrammaticAccessGrant_permissions) SetOrganization(value OrganizationProgrammaticAccessGrant_permissions_organizationable)() {
    m.organization = value
}
// SetOther sets the other property value. The other property
func (m *OrganizationProgrammaticAccessGrant_permissions) SetOther(value OrganizationProgrammaticAccessGrant_permissions_otherable)() {
    m.other = value
}
// SetRepository sets the repository property value. The repository property
func (m *OrganizationProgrammaticAccessGrant_permissions) SetRepository(value OrganizationProgrammaticAccessGrant_permissions_repositoryable)() {
    m.repository = value
}
type OrganizationProgrammaticAccessGrant_permissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOrganization()(OrganizationProgrammaticAccessGrant_permissions_organizationable)
    GetOther()(OrganizationProgrammaticAccessGrant_permissions_otherable)
    GetRepository()(OrganizationProgrammaticAccessGrant_permissions_repositoryable)
    SetOrganization(value OrganizationProgrammaticAccessGrant_permissions_organizationable)()
    SetOther(value OrganizationProgrammaticAccessGrant_permissions_otherable)()
    SetRepository(value OrganizationProgrammaticAccessGrant_permissions_repositoryable)()
}
