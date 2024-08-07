package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunityProfile community Profile
type CommunityProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content_reports_enabled property
    content_reports_enabled *bool
    // The description property
    description *string
    // The documentation property
    documentation *string
    // The files property
    files CommunityProfile_filesable
    // The health_percentage property
    health_percentage *int32
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCommunityProfile instantiates a new CommunityProfile and sets the default values.
func NewCommunityProfile()(*CommunityProfile) {
    m := &CommunityProfile{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommunityProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommunityProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunityProfile(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommunityProfile) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentReportsEnabled gets the content_reports_enabled property value. The content_reports_enabled property
// returns a *bool when successful
func (m *CommunityProfile) GetContentReportsEnabled()(*bool) {
    return m.content_reports_enabled
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *CommunityProfile) GetDescription()(*string) {
    return m.description
}
// GetDocumentation gets the documentation property value. The documentation property
// returns a *string when successful
func (m *CommunityProfile) GetDocumentation()(*string) {
    return m.documentation
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommunityProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content_reports_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentReportsEnabled(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["documentation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentation(val)
        }
        return nil
    }
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommunityProfile_filesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFiles(val.(CommunityProfile_filesable))
        }
        return nil
    }
    res["health_percentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthPercentage(val)
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
    return res
}
// GetFiles gets the files property value. The files property
// returns a CommunityProfile_filesable when successful
func (m *CommunityProfile) GetFiles()(CommunityProfile_filesable) {
    return m.files
}
// GetHealthPercentage gets the health_percentage property value. The health_percentage property
// returns a *int32 when successful
func (m *CommunityProfile) GetHealthPercentage()(*int32) {
    return m.health_percentage
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *CommunityProfile) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *CommunityProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("content_reports_enabled", m.GetContentReportsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("documentation", m.GetDocumentation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("files", m.GetFiles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("health_percentage", m.GetHealthPercentage())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunityProfile) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentReportsEnabled sets the content_reports_enabled property value. The content_reports_enabled property
func (m *CommunityProfile) SetContentReportsEnabled(value *bool)() {
    m.content_reports_enabled = value
}
// SetDescription sets the description property value. The description property
func (m *CommunityProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDocumentation sets the documentation property value. The documentation property
func (m *CommunityProfile) SetDocumentation(value *string)() {
    m.documentation = value
}
// SetFiles sets the files property value. The files property
func (m *CommunityProfile) SetFiles(value CommunityProfile_filesable)() {
    m.files = value
}
// SetHealthPercentage sets the health_percentage property value. The health_percentage property
func (m *CommunityProfile) SetHealthPercentage(value *int32)() {
    m.health_percentage = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *CommunityProfile) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type CommunityProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentReportsEnabled()(*bool)
    GetDescription()(*string)
    GetDocumentation()(*string)
    GetFiles()(CommunityProfile_filesable)
    GetHealthPercentage()(*int32)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetContentReportsEnabled(value *bool)()
    SetDescription(value *string)()
    SetDocumentation(value *string)()
    SetFiles(value CommunityProfile_filesable)()
    SetHealthPercentage(value *int32)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
