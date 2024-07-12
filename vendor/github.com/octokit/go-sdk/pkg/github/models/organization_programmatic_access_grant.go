package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationProgrammaticAccessGrant minimal representation of an organization programmatic access grant for enumerations
type OrganizationProgrammaticAccessGrant struct {
    // Date and time when the fine-grained personal access token was approved to access the organization.
    access_granted_at *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique identifier of the fine-grained personal access token. The `pat_id` used to get details about an approved fine-grained personal access token.
    id *int32
    // A GitHub user.
    owner SimpleUserable
    // Permissions requested, categorized by type of permission.
    permissions OrganizationProgrammaticAccessGrant_permissionsable
    // URL to the list of repositories the fine-grained personal access token can access. Only follow when `repository_selection` is `subset`.
    repositories_url *string
    // Type of repository selection requested.
    repository_selection *OrganizationProgrammaticAccessGrant_repository_selection
    // Whether the associated fine-grained personal access token has expired.
    token_expired *bool
    // Date and time when the associated fine-grained personal access token expires.
    token_expires_at *string
    // Date and time when the associated fine-grained personal access token was last used for authentication.
    token_last_used_at *string
}
// NewOrganizationProgrammaticAccessGrant instantiates a new OrganizationProgrammaticAccessGrant and sets the default values.
func NewOrganizationProgrammaticAccessGrant()(*OrganizationProgrammaticAccessGrant) {
    m := &OrganizationProgrammaticAccessGrant{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationProgrammaticAccessGrantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationProgrammaticAccessGrantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationProgrammaticAccessGrant(), nil
}
// GetAccessGrantedAt gets the access_granted_at property value. Date and time when the fine-grained personal access token was approved to access the organization.
// returns a *string when successful
func (m *OrganizationProgrammaticAccessGrant) GetAccessGrantedAt()(*string) {
    return m.access_granted_at
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationProgrammaticAccessGrant) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationProgrammaticAccessGrant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access_granted_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessGrantedAt(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(SimpleUserable))
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationProgrammaticAccessGrant_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(OrganizationProgrammaticAccessGrant_permissionsable))
        }
        return nil
    }
    res["repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoriesUrl(val)
        }
        return nil
    }
    res["repository_selection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationProgrammaticAccessGrant_repository_selection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositorySelection(val.(*OrganizationProgrammaticAccessGrant_repository_selection))
        }
        return nil
    }
    res["token_expired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenExpired(val)
        }
        return nil
    }
    res["token_expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenExpiresAt(val)
        }
        return nil
    }
    res["token_last_used_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenLastUsedAt(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique identifier of the fine-grained personal access token. The `pat_id` used to get details about an approved fine-grained personal access token.
// returns a *int32 when successful
func (m *OrganizationProgrammaticAccessGrant) GetId()(*int32) {
    return m.id
}
// GetOwner gets the owner property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *OrganizationProgrammaticAccessGrant) GetOwner()(SimpleUserable) {
    return m.owner
}
// GetPermissions gets the permissions property value. Permissions requested, categorized by type of permission.
// returns a OrganizationProgrammaticAccessGrant_permissionsable when successful
func (m *OrganizationProgrammaticAccessGrant) GetPermissions()(OrganizationProgrammaticAccessGrant_permissionsable) {
    return m.permissions
}
// GetRepositoriesUrl gets the repositories_url property value. URL to the list of repositories the fine-grained personal access token can access. Only follow when `repository_selection` is `subset`.
// returns a *string when successful
func (m *OrganizationProgrammaticAccessGrant) GetRepositoriesUrl()(*string) {
    return m.repositories_url
}
// GetRepositorySelection gets the repository_selection property value. Type of repository selection requested.
// returns a *OrganizationProgrammaticAccessGrant_repository_selection when successful
func (m *OrganizationProgrammaticAccessGrant) GetRepositorySelection()(*OrganizationProgrammaticAccessGrant_repository_selection) {
    return m.repository_selection
}
// GetTokenExpired gets the token_expired property value. Whether the associated fine-grained personal access token has expired.
// returns a *bool when successful
func (m *OrganizationProgrammaticAccessGrant) GetTokenExpired()(*bool) {
    return m.token_expired
}
// GetTokenExpiresAt gets the token_expires_at property value. Date and time when the associated fine-grained personal access token expires.
// returns a *string when successful
func (m *OrganizationProgrammaticAccessGrant) GetTokenExpiresAt()(*string) {
    return m.token_expires_at
}
// GetTokenLastUsedAt gets the token_last_used_at property value. Date and time when the associated fine-grained personal access token was last used for authentication.
// returns a *string when successful
func (m *OrganizationProgrammaticAccessGrant) GetTokenLastUsedAt()(*string) {
    return m.token_last_used_at
}
// Serialize serializes information the current object
func (m *OrganizationProgrammaticAccessGrant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("access_granted_at", m.GetAccessGrantedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repositories_url", m.GetRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    if m.GetRepositorySelection() != nil {
        cast := (*m.GetRepositorySelection()).String()
        err := writer.WriteStringValue("repository_selection", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("token_expired", m.GetTokenExpired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token_expires_at", m.GetTokenExpiresAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token_last_used_at", m.GetTokenLastUsedAt())
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
// SetAccessGrantedAt sets the access_granted_at property value. Date and time when the fine-grained personal access token was approved to access the organization.
func (m *OrganizationProgrammaticAccessGrant) SetAccessGrantedAt(value *string)() {
    m.access_granted_at = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationProgrammaticAccessGrant) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Unique identifier of the fine-grained personal access token. The `pat_id` used to get details about an approved fine-grained personal access token.
func (m *OrganizationProgrammaticAccessGrant) SetId(value *int32)() {
    m.id = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *OrganizationProgrammaticAccessGrant) SetOwner(value SimpleUserable)() {
    m.owner = value
}
// SetPermissions sets the permissions property value. Permissions requested, categorized by type of permission.
func (m *OrganizationProgrammaticAccessGrant) SetPermissions(value OrganizationProgrammaticAccessGrant_permissionsable)() {
    m.permissions = value
}
// SetRepositoriesUrl sets the repositories_url property value. URL to the list of repositories the fine-grained personal access token can access. Only follow when `repository_selection` is `subset`.
func (m *OrganizationProgrammaticAccessGrant) SetRepositoriesUrl(value *string)() {
    m.repositories_url = value
}
// SetRepositorySelection sets the repository_selection property value. Type of repository selection requested.
func (m *OrganizationProgrammaticAccessGrant) SetRepositorySelection(value *OrganizationProgrammaticAccessGrant_repository_selection)() {
    m.repository_selection = value
}
// SetTokenExpired sets the token_expired property value. Whether the associated fine-grained personal access token has expired.
func (m *OrganizationProgrammaticAccessGrant) SetTokenExpired(value *bool)() {
    m.token_expired = value
}
// SetTokenExpiresAt sets the token_expires_at property value. Date and time when the associated fine-grained personal access token expires.
func (m *OrganizationProgrammaticAccessGrant) SetTokenExpiresAt(value *string)() {
    m.token_expires_at = value
}
// SetTokenLastUsedAt sets the token_last_used_at property value. Date and time when the associated fine-grained personal access token was last used for authentication.
func (m *OrganizationProgrammaticAccessGrant) SetTokenLastUsedAt(value *string)() {
    m.token_last_used_at = value
}
type OrganizationProgrammaticAccessGrantable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessGrantedAt()(*string)
    GetId()(*int32)
    GetOwner()(SimpleUserable)
    GetPermissions()(OrganizationProgrammaticAccessGrant_permissionsable)
    GetRepositoriesUrl()(*string)
    GetRepositorySelection()(*OrganizationProgrammaticAccessGrant_repository_selection)
    GetTokenExpired()(*bool)
    GetTokenExpiresAt()(*string)
    GetTokenLastUsedAt()(*string)
    SetAccessGrantedAt(value *string)()
    SetId(value *int32)()
    SetOwner(value SimpleUserable)()
    SetPermissions(value OrganizationProgrammaticAccessGrant_permissionsable)()
    SetRepositoriesUrl(value *string)()
    SetRepositorySelection(value *OrganizationProgrammaticAccessGrant_repository_selection)()
    SetTokenExpired(value *bool)()
    SetTokenExpiresAt(value *string)()
    SetTokenLastUsedAt(value *string)()
}
