package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningCodeqlDatabase a CodeQL database.
type CodeScanningCodeqlDatabase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commit SHA of the repository at the time the CodeQL database was created.
    commit_oid *string
    // The MIME type of the CodeQL database file.
    content_type *string
    // The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ID of the CodeQL database.
    id *int32
    // The language of the CodeQL database.
    language *string
    // The name of the CodeQL database.
    name *string
    // The size of the CodeQL database file in bytes.
    size *int32
    // The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    uploader SimpleUserable
    // The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
    url *string
}
// NewCodeScanningCodeqlDatabase instantiates a new CodeScanningCodeqlDatabase and sets the default values.
func NewCodeScanningCodeqlDatabase()(*CodeScanningCodeqlDatabase) {
    m := &CodeScanningCodeqlDatabase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningCodeqlDatabaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningCodeqlDatabaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningCodeqlDatabase(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningCodeqlDatabase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommitOid gets the commit_oid property value. The commit SHA of the repository at the time the CodeQL database was created.
// returns a *string when successful
func (m *CodeScanningCodeqlDatabase) GetCommitOid()(*string) {
    return m.commit_oid
}
// GetContentType gets the content_type property value. The MIME type of the CodeQL database file.
// returns a *string when successful
func (m *CodeScanningCodeqlDatabase) GetContentType()(*string) {
    return m.content_type
}
// GetCreatedAt gets the created_at property value. The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
// returns a *Time when successful
func (m *CodeScanningCodeqlDatabase) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningCodeqlDatabase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commit_oid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitOid(val)
        }
        return nil
    }
    res["content_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
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
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
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
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
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
    res["uploader"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploader(val.(SimpleUserable))
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
    return res
}
// GetId gets the id property value. The ID of the CodeQL database.
// returns a *int32 when successful
func (m *CodeScanningCodeqlDatabase) GetId()(*int32) {
    return m.id
}
// GetLanguage gets the language property value. The language of the CodeQL database.
// returns a *string when successful
func (m *CodeScanningCodeqlDatabase) GetLanguage()(*string) {
    return m.language
}
// GetName gets the name property value. The name of the CodeQL database.
// returns a *string when successful
func (m *CodeScanningCodeqlDatabase) GetName()(*string) {
    return m.name
}
// GetSize gets the size property value. The size of the CodeQL database file in bytes.
// returns a *int32 when successful
func (m *CodeScanningCodeqlDatabase) GetSize()(*int32) {
    return m.size
}
// GetUpdatedAt gets the updated_at property value. The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
// returns a *Time when successful
func (m *CodeScanningCodeqlDatabase) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUploader gets the uploader property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *CodeScanningCodeqlDatabase) GetUploader()(SimpleUserable) {
    return m.uploader
}
// GetUrl gets the url property value. The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
// returns a *string when successful
func (m *CodeScanningCodeqlDatabase) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CodeScanningCodeqlDatabase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("commit_oid", m.GetCommitOid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content_type", m.GetContentType())
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
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
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
        err := writer.WriteInt32Value("size", m.GetSize())
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
        err := writer.WriteObjectValue("uploader", m.GetUploader())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodeScanningCodeqlDatabase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommitOid sets the commit_oid property value. The commit SHA of the repository at the time the CodeQL database was created.
func (m *CodeScanningCodeqlDatabase) SetCommitOid(value *string)() {
    m.commit_oid = value
}
// SetContentType sets the content_type property value. The MIME type of the CodeQL database file.
func (m *CodeScanningCodeqlDatabase) SetContentType(value *string)() {
    m.content_type = value
}
// SetCreatedAt sets the created_at property value. The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
func (m *CodeScanningCodeqlDatabase) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetId sets the id property value. The ID of the CodeQL database.
func (m *CodeScanningCodeqlDatabase) SetId(value *int32)() {
    m.id = value
}
// SetLanguage sets the language property value. The language of the CodeQL database.
func (m *CodeScanningCodeqlDatabase) SetLanguage(value *string)() {
    m.language = value
}
// SetName sets the name property value. The name of the CodeQL database.
func (m *CodeScanningCodeqlDatabase) SetName(value *string)() {
    m.name = value
}
// SetSize sets the size property value. The size of the CodeQL database file in bytes.
func (m *CodeScanningCodeqlDatabase) SetSize(value *int32)() {
    m.size = value
}
// SetUpdatedAt sets the updated_at property value. The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
func (m *CodeScanningCodeqlDatabase) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUploader sets the uploader property value. A GitHub user.
func (m *CodeScanningCodeqlDatabase) SetUploader(value SimpleUserable)() {
    m.uploader = value
}
// SetUrl sets the url property value. The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
func (m *CodeScanningCodeqlDatabase) SetUrl(value *string)() {
    m.url = value
}
type CodeScanningCodeqlDatabaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommitOid()(*string)
    GetContentType()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int32)
    GetLanguage()(*string)
    GetName()(*string)
    GetSize()(*int32)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUploader()(SimpleUserable)
    GetUrl()(*string)
    SetCommitOid(value *string)()
    SetContentType(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int32)()
    SetLanguage(value *string)()
    SetName(value *string)()
    SetSize(value *int32)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUploader(value SimpleUserable)()
    SetUrl(value *string)()
}
