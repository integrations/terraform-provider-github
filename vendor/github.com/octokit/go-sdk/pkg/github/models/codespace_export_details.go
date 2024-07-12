package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespaceExportDetails an export of a codespace. Also, latest export details for a codespace can be fetched with id = latest
type CodespaceExportDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the exported branch
    branch *string
    // Completion time of the last export operation
    completed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Url for fetching export details
    export_url *string
    // Web url for the exported branch
    html_url *string
    // Id for the export details
    id *string
    // Git commit SHA of the exported branch
    sha *string
    // State of the latest export
    state *string
}
// NewCodespaceExportDetails instantiates a new CodespaceExportDetails and sets the default values.
func NewCodespaceExportDetails()(*CodespaceExportDetails) {
    m := &CodespaceExportDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespaceExportDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodespaceExportDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespaceExportDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodespaceExportDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBranch gets the branch property value. Name of the exported branch
// returns a *string when successful
func (m *CodespaceExportDetails) GetBranch()(*string) {
    return m.branch
}
// GetCompletedAt gets the completed_at property value. Completion time of the last export operation
// returns a *Time when successful
func (m *CodespaceExportDetails) GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completed_at
}
// GetExportUrl gets the export_url property value. Url for fetching export details
// returns a *string when successful
func (m *CodespaceExportDetails) GetExportUrl()(*string) {
    return m.export_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodespaceExportDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranch(val)
        }
        return nil
    }
    res["completed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedAt(val)
        }
        return nil
    }
    res["export_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportUrl(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. Web url for the exported branch
// returns a *string when successful
func (m *CodespaceExportDetails) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. Id for the export details
// returns a *string when successful
func (m *CodespaceExportDetails) GetId()(*string) {
    return m.id
}
// GetSha gets the sha property value. Git commit SHA of the exported branch
// returns a *string when successful
func (m *CodespaceExportDetails) GetSha()(*string) {
    return m.sha
}
// GetState gets the state property value. State of the latest export
// returns a *string when successful
func (m *CodespaceExportDetails) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *CodespaceExportDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("branch", m.GetBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("completed_at", m.GetCompletedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("export_url", m.GetExportUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *CodespaceExportDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBranch sets the branch property value. Name of the exported branch
func (m *CodespaceExportDetails) SetBranch(value *string)() {
    m.branch = value
}
// SetCompletedAt sets the completed_at property value. Completion time of the last export operation
func (m *CodespaceExportDetails) SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completed_at = value
}
// SetExportUrl sets the export_url property value. Url for fetching export details
func (m *CodespaceExportDetails) SetExportUrl(value *string)() {
    m.export_url = value
}
// SetHtmlUrl sets the html_url property value. Web url for the exported branch
func (m *CodespaceExportDetails) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. Id for the export details
func (m *CodespaceExportDetails) SetId(value *string)() {
    m.id = value
}
// SetSha sets the sha property value. Git commit SHA of the exported branch
func (m *CodespaceExportDetails) SetSha(value *string)() {
    m.sha = value
}
// SetState sets the state property value. State of the latest export
func (m *CodespaceExportDetails) SetState(value *string)() {
    m.state = value
}
type CodespaceExportDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranch()(*string)
    GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExportUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*string)
    GetSha()(*string)
    GetState()(*string)
    SetBranch(value *string)()
    SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExportUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *string)()
    SetSha(value *string)()
    SetState(value *string)()
}
