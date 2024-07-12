package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationToken authentication Token
type AuthenticationToken struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time this token expires
    expires_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The permissions property
    permissions AuthenticationToken_permissionsable
    // The repositories this token has access to
    repositories []Repositoryable
    // Describe whether all repositories have been selected or there's a selection involved
    repository_selection *AuthenticationToken_repository_selection
    // The single_file property
    single_file *string
    // The token used for authentication
    token *string
}
// NewAuthenticationToken instantiates a new AuthenticationToken and sets the default values.
func NewAuthenticationToken()(*AuthenticationToken) {
    m := &AuthenticationToken{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationToken(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthenticationToken) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiresAt gets the expires_at property value. The time this token expires
// returns a *Time when successful
func (m *AuthenticationToken) GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expires_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAt(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationToken_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(AuthenticationToken_permissionsable))
        }
        return nil
    }
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Repositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Repositoryable)
                }
            }
            m.SetRepositories(res)
        }
        return nil
    }
    res["repository_selection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationToken_repository_selection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositorySelection(val.(*AuthenticationToken_repository_selection))
        }
        return nil
    }
    res["single_file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleFile(val)
        }
        return nil
    }
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    return res
}
// GetPermissions gets the permissions property value. The permissions property
// returns a AuthenticationToken_permissionsable when successful
func (m *AuthenticationToken) GetPermissions()(AuthenticationToken_permissionsable) {
    return m.permissions
}
// GetRepositories gets the repositories property value. The repositories this token has access to
// returns a []Repositoryable when successful
func (m *AuthenticationToken) GetRepositories()([]Repositoryable) {
    return m.repositories
}
// GetRepositorySelection gets the repository_selection property value. Describe whether all repositories have been selected or there's a selection involved
// returns a *AuthenticationToken_repository_selection when successful
func (m *AuthenticationToken) GetRepositorySelection()(*AuthenticationToken_repository_selection) {
    return m.repository_selection
}
// GetSingleFile gets the single_file property value. The single_file property
// returns a *string when successful
func (m *AuthenticationToken) GetSingleFile()(*string) {
    return m.single_file
}
// GetToken gets the token property value. The token used for authentication
// returns a *string when successful
func (m *AuthenticationToken) GetToken()(*string) {
    return m.token
}
// Serialize serializes information the current object
func (m *AuthenticationToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expires_at", m.GetExpiresAt())
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
    if m.GetRepositories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositories()))
        for i, v := range m.GetRepositories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("repositories", cast)
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
        err := writer.WriteStringValue("single_file", m.GetSingleFile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token", m.GetToken())
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
func (m *AuthenticationToken) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiresAt sets the expires_at property value. The time this token expires
func (m *AuthenticationToken) SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expires_at = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *AuthenticationToken) SetPermissions(value AuthenticationToken_permissionsable)() {
    m.permissions = value
}
// SetRepositories sets the repositories property value. The repositories this token has access to
func (m *AuthenticationToken) SetRepositories(value []Repositoryable)() {
    m.repositories = value
}
// SetRepositorySelection sets the repository_selection property value. Describe whether all repositories have been selected or there's a selection involved
func (m *AuthenticationToken) SetRepositorySelection(value *AuthenticationToken_repository_selection)() {
    m.repository_selection = value
}
// SetSingleFile sets the single_file property value. The single_file property
func (m *AuthenticationToken) SetSingleFile(value *string)() {
    m.single_file = value
}
// SetToken sets the token property value. The token used for authentication
func (m *AuthenticationToken) SetToken(value *string)() {
    m.token = value
}
type AuthenticationTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPermissions()(AuthenticationToken_permissionsable)
    GetRepositories()([]Repositoryable)
    GetRepositorySelection()(*AuthenticationToken_repository_selection)
    GetSingleFile()(*string)
    GetToken()(*string)
    SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPermissions(value AuthenticationToken_permissionsable)()
    SetRepositories(value []Repositoryable)()
    SetRepositorySelection(value *AuthenticationToken_repository_selection)()
    SetSingleFile(value *string)()
    SetToken(value *string)()
}
