package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationProgrammaticAccessGrantRequest minimal representation of an organization programmatic access grant request for enumerations
type OrganizationProgrammaticAccessGrantRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Date and time when the request for access was created.
    created_at *string
    // Unique identifier of the request for access via fine-grained personal access token. The `pat_request_id` used to review PAT requests.
    id *int32
    // A GitHub user.
    owner SimpleUserable
    // Permissions requested, categorized by type of permission.
    permissions OrganizationProgrammaticAccessGrantRequest_permissionsable
    // Reason for requesting access.
    reason *string
    // URL to the list of repositories requested to be accessed via fine-grained personal access token. Should only be followed when `repository_selection` is `subset`.
    repositories_url *string
    // Type of repository selection requested.
    repository_selection *OrganizationProgrammaticAccessGrantRequest_repository_selection
    // Whether the associated fine-grained personal access token has expired.
    token_expired *bool
    // Date and time when the associated fine-grained personal access token expires.
    token_expires_at *string
    // Date and time when the associated fine-grained personal access token was last used for authentication.
    token_last_used_at *string
}
// NewOrganizationProgrammaticAccessGrantRequest instantiates a new OrganizationProgrammaticAccessGrantRequest and sets the default values.
func NewOrganizationProgrammaticAccessGrantRequest()(*OrganizationProgrammaticAccessGrantRequest) {
    m := &OrganizationProgrammaticAccessGrantRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationProgrammaticAccessGrantRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationProgrammaticAccessGrantRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationProgrammaticAccessGrantRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. Date and time when the request for access was created.
// returns a *string when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetCreatedAt()(*string) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
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
        val, err := n.GetObjectValue(CreateOrganizationProgrammaticAccessGrantRequest_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(OrganizationProgrammaticAccessGrantRequest_permissionsable))
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
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
        val, err := n.GetEnumValue(ParseOrganizationProgrammaticAccessGrantRequest_repository_selection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositorySelection(val.(*OrganizationProgrammaticAccessGrantRequest_repository_selection))
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
// GetId gets the id property value. Unique identifier of the request for access via fine-grained personal access token. The `pat_request_id` used to review PAT requests.
// returns a *int32 when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetId()(*int32) {
    return m.id
}
// GetOwner gets the owner property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetOwner()(SimpleUserable) {
    return m.owner
}
// GetPermissions gets the permissions property value. Permissions requested, categorized by type of permission.
// returns a OrganizationProgrammaticAccessGrantRequest_permissionsable when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetPermissions()(OrganizationProgrammaticAccessGrantRequest_permissionsable) {
    return m.permissions
}
// GetReason gets the reason property value. Reason for requesting access.
// returns a *string when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetReason()(*string) {
    return m.reason
}
// GetRepositoriesUrl gets the repositories_url property value. URL to the list of repositories requested to be accessed via fine-grained personal access token. Should only be followed when `repository_selection` is `subset`.
// returns a *string when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetRepositoriesUrl()(*string) {
    return m.repositories_url
}
// GetRepositorySelection gets the repository_selection property value. Type of repository selection requested.
// returns a *OrganizationProgrammaticAccessGrantRequest_repository_selection when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetRepositorySelection()(*OrganizationProgrammaticAccessGrantRequest_repository_selection) {
    return m.repository_selection
}
// GetTokenExpired gets the token_expired property value. Whether the associated fine-grained personal access token has expired.
// returns a *bool when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetTokenExpired()(*bool) {
    return m.token_expired
}
// GetTokenExpiresAt gets the token_expires_at property value. Date and time when the associated fine-grained personal access token expires.
// returns a *string when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetTokenExpiresAt()(*string) {
    return m.token_expires_at
}
// GetTokenLastUsedAt gets the token_last_used_at property value. Date and time when the associated fine-grained personal access token was last used for authentication.
// returns a *string when successful
func (m *OrganizationProgrammaticAccessGrantRequest) GetTokenLastUsedAt()(*string) {
    return m.token_last_used_at
}
// Serialize serializes information the current object
func (m *OrganizationProgrammaticAccessGrantRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteStringValue("reason", m.GetReason())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationProgrammaticAccessGrantRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. Date and time when the request for access was created.
func (m *OrganizationProgrammaticAccessGrantRequest) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetId sets the id property value. Unique identifier of the request for access via fine-grained personal access token. The `pat_request_id` used to review PAT requests.
func (m *OrganizationProgrammaticAccessGrantRequest) SetId(value *int32)() {
    m.id = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *OrganizationProgrammaticAccessGrantRequest) SetOwner(value SimpleUserable)() {
    m.owner = value
}
// SetPermissions sets the permissions property value. Permissions requested, categorized by type of permission.
func (m *OrganizationProgrammaticAccessGrantRequest) SetPermissions(value OrganizationProgrammaticAccessGrantRequest_permissionsable)() {
    m.permissions = value
}
// SetReason sets the reason property value. Reason for requesting access.
func (m *OrganizationProgrammaticAccessGrantRequest) SetReason(value *string)() {
    m.reason = value
}
// SetRepositoriesUrl sets the repositories_url property value. URL to the list of repositories requested to be accessed via fine-grained personal access token. Should only be followed when `repository_selection` is `subset`.
func (m *OrganizationProgrammaticAccessGrantRequest) SetRepositoriesUrl(value *string)() {
    m.repositories_url = value
}
// SetRepositorySelection sets the repository_selection property value. Type of repository selection requested.
func (m *OrganizationProgrammaticAccessGrantRequest) SetRepositorySelection(value *OrganizationProgrammaticAccessGrantRequest_repository_selection)() {
    m.repository_selection = value
}
// SetTokenExpired sets the token_expired property value. Whether the associated fine-grained personal access token has expired.
func (m *OrganizationProgrammaticAccessGrantRequest) SetTokenExpired(value *bool)() {
    m.token_expired = value
}
// SetTokenExpiresAt sets the token_expires_at property value. Date and time when the associated fine-grained personal access token expires.
func (m *OrganizationProgrammaticAccessGrantRequest) SetTokenExpiresAt(value *string)() {
    m.token_expires_at = value
}
// SetTokenLastUsedAt sets the token_last_used_at property value. Date and time when the associated fine-grained personal access token was last used for authentication.
func (m *OrganizationProgrammaticAccessGrantRequest) SetTokenLastUsedAt(value *string)() {
    m.token_last_used_at = value
}
type OrganizationProgrammaticAccessGrantRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetId()(*int32)
    GetOwner()(SimpleUserable)
    GetPermissions()(OrganizationProgrammaticAccessGrantRequest_permissionsable)
    GetReason()(*string)
    GetRepositoriesUrl()(*string)
    GetRepositorySelection()(*OrganizationProgrammaticAccessGrantRequest_repository_selection)
    GetTokenExpired()(*bool)
    GetTokenExpiresAt()(*string)
    GetTokenLastUsedAt()(*string)
    SetCreatedAt(value *string)()
    SetId(value *int32)()
    SetOwner(value SimpleUserable)()
    SetPermissions(value OrganizationProgrammaticAccessGrantRequest_permissionsable)()
    SetReason(value *string)()
    SetRepositoriesUrl(value *string)()
    SetRepositorySelection(value *OrganizationProgrammaticAccessGrantRequest_repository_selection)()
    SetTokenExpired(value *bool)()
    SetTokenExpiresAt(value *string)()
    SetTokenLastUsedAt(value *string)()
}
