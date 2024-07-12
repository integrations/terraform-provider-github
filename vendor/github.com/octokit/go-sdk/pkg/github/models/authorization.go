package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Authorization the authorization for an OAuth app, GitHub App, or a Personal Access Token.
type Authorization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The app property
    app Authorization_appable
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The expires_at property
    expires_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The fingerprint property
    fingerprint *string
    // The hashed_token property
    hashed_token *string
    // The id property
    id *int64
    // The installation property
    installation NullableScopedInstallationable
    // The note property
    note *string
    // The note_url property
    note_url *string
    // A list of scopes that this authorization is in.
    scopes []string
    // The token property
    token *string
    // The token_last_eight property
    token_last_eight *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewAuthorization instantiates a new Authorization and sets the default values.
func NewAuthorization()(*Authorization) {
    m := &Authorization{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthorizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthorizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthorization(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Authorization) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApp gets the app property value. The app property
// returns a Authorization_appable when successful
func (m *Authorization) GetApp()(Authorization_appable) {
    return m.app
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Authorization) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetExpiresAt gets the expires_at property value. The expires_at property
// returns a *Time when successful
func (m *Authorization) GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expires_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Authorization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorization_appFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApp(val.(Authorization_appable))
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
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
    res["fingerprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprint(val)
        }
        return nil
    }
    res["hashed_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashedToken(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["installation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableScopedInstallationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallation(val.(NullableScopedInstallationable))
        }
        return nil
    }
    res["note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNote(val)
        }
        return nil
    }
    res["note_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoteUrl(val)
        }
        return nil
    }
    res["scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetScopes(res)
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
    res["token_last_eight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenLastEight(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(NullableSimpleUserable))
        }
        return nil
    }
    return res
}
// GetFingerprint gets the fingerprint property value. The fingerprint property
// returns a *string when successful
func (m *Authorization) GetFingerprint()(*string) {
    return m.fingerprint
}
// GetHashedToken gets the hashed_token property value. The hashed_token property
// returns a *string when successful
func (m *Authorization) GetHashedToken()(*string) {
    return m.hashed_token
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *Authorization) GetId()(*int64) {
    return m.id
}
// GetInstallation gets the installation property value. The installation property
// returns a NullableScopedInstallationable when successful
func (m *Authorization) GetInstallation()(NullableScopedInstallationable) {
    return m.installation
}
// GetNote gets the note property value. The note property
// returns a *string when successful
func (m *Authorization) GetNote()(*string) {
    return m.note
}
// GetNoteUrl gets the note_url property value. The note_url property
// returns a *string when successful
func (m *Authorization) GetNoteUrl()(*string) {
    return m.note_url
}
// GetScopes gets the scopes property value. A list of scopes that this authorization is in.
// returns a []string when successful
func (m *Authorization) GetScopes()([]string) {
    return m.scopes
}
// GetToken gets the token property value. The token property
// returns a *string when successful
func (m *Authorization) GetToken()(*string) {
    return m.token
}
// GetTokenLastEight gets the token_last_eight property value. The token_last_eight property
// returns a *string when successful
func (m *Authorization) GetTokenLastEight()(*string) {
    return m.token_last_eight
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Authorization) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Authorization) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Authorization) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *Authorization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("app", m.GetApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expires_at", m.GetExpiresAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fingerprint", m.GetFingerprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hashed_token", m.GetHashedToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("installation", m.GetInstallation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("note", m.GetNote())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("note_url", m.GetNoteUrl())
        if err != nil {
            return err
        }
    }
    if m.GetScopes() != nil {
        err := writer.WriteCollectionOfStringValues("scopes", m.GetScopes())
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
        err := writer.WriteStringValue("token_last_eight", m.GetTokenLastEight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *Authorization) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApp sets the app property value. The app property
func (m *Authorization) SetApp(value Authorization_appable)() {
    m.app = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Authorization) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetExpiresAt sets the expires_at property value. The expires_at property
func (m *Authorization) SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expires_at = value
}
// SetFingerprint sets the fingerprint property value. The fingerprint property
func (m *Authorization) SetFingerprint(value *string)() {
    m.fingerprint = value
}
// SetHashedToken sets the hashed_token property value. The hashed_token property
func (m *Authorization) SetHashedToken(value *string)() {
    m.hashed_token = value
}
// SetId sets the id property value. The id property
func (m *Authorization) SetId(value *int64)() {
    m.id = value
}
// SetInstallation sets the installation property value. The installation property
func (m *Authorization) SetInstallation(value NullableScopedInstallationable)() {
    m.installation = value
}
// SetNote sets the note property value. The note property
func (m *Authorization) SetNote(value *string)() {
    m.note = value
}
// SetNoteUrl sets the note_url property value. The note_url property
func (m *Authorization) SetNoteUrl(value *string)() {
    m.note_url = value
}
// SetScopes sets the scopes property value. A list of scopes that this authorization is in.
func (m *Authorization) SetScopes(value []string)() {
    m.scopes = value
}
// SetToken sets the token property value. The token property
func (m *Authorization) SetToken(value *string)() {
    m.token = value
}
// SetTokenLastEight sets the token_last_eight property value. The token_last_eight property
func (m *Authorization) SetTokenLastEight(value *string)() {
    m.token_last_eight = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Authorization) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Authorization) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *Authorization) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type Authorizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApp()(Authorization_appable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFingerprint()(*string)
    GetHashedToken()(*string)
    GetId()(*int64)
    GetInstallation()(NullableScopedInstallationable)
    GetNote()(*string)
    GetNoteUrl()(*string)
    GetScopes()([]string)
    GetToken()(*string)
    GetTokenLastEight()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
    SetApp(value Authorization_appable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFingerprint(value *string)()
    SetHashedToken(value *string)()
    SetId(value *int64)()
    SetInstallation(value NullableScopedInstallationable)()
    SetNote(value *string)()
    SetNoteUrl(value *string)()
    SetScopes(value []string)()
    SetToken(value *string)()
    SetTokenLastEight(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
