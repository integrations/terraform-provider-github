package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProjectCard project cards represent a scope of work.
type ProjectCard struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether or not the card is archived
    archived *bool
    // The column_name property
    column_name *string
    // The column_url property
    column_url *string
    // The content_url property
    content_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    creator NullableSimpleUserable
    // The project card's ID
    id *int64
    // The node_id property
    node_id *string
    // The note property
    note *string
    // The project_id property
    project_id *string
    // The project_url property
    project_url *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewProjectCard instantiates a new ProjectCard and sets the default values.
func NewProjectCard()(*ProjectCard) {
    m := &ProjectCard{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProjectCardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProjectCardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProjectCard(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProjectCard) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchived gets the archived property value. Whether or not the card is archived
// returns a *bool when successful
func (m *ProjectCard) GetArchived()(*bool) {
    return m.archived
}
// GetColumnName gets the column_name property value. The column_name property
// returns a *string when successful
func (m *ProjectCard) GetColumnName()(*string) {
    return m.column_name
}
// GetColumnUrl gets the column_url property value. The column_url property
// returns a *string when successful
func (m *ProjectCard) GetColumnUrl()(*string) {
    return m.column_url
}
// GetContentUrl gets the content_url property value. The content_url property
// returns a *string when successful
func (m *ProjectCard) GetContentUrl()(*string) {
    return m.content_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *ProjectCard) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCreator gets the creator property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *ProjectCard) GetCreator()(NullableSimpleUserable) {
    return m.creator
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProjectCard) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchived(val)
        }
        return nil
    }
    res["column_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnName(val)
        }
        return nil
    }
    res["column_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnUrl(val)
        }
        return nil
    }
    res["content_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentUrl(val)
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
    res["creator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreator(val.(NullableSimpleUserable))
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
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
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
    res["project_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProjectId(val)
        }
        return nil
    }
    res["project_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProjectUrl(val)
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
    return res
}
// GetId gets the id property value. The project card's ID
// returns a *int64 when successful
func (m *ProjectCard) GetId()(*int64) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *ProjectCard) GetNodeId()(*string) {
    return m.node_id
}
// GetNote gets the note property value. The note property
// returns a *string when successful
func (m *ProjectCard) GetNote()(*string) {
    return m.note
}
// GetProjectId gets the project_id property value. The project_id property
// returns a *string when successful
func (m *ProjectCard) GetProjectId()(*string) {
    return m.project_id
}
// GetProjectUrl gets the project_url property value. The project_url property
// returns a *string when successful
func (m *ProjectCard) GetProjectUrl()(*string) {
    return m.project_url
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *ProjectCard) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ProjectCard) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ProjectCard) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("archived", m.GetArchived())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("column_name", m.GetColumnName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("column_url", m.GetColumnUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content_url", m.GetContentUrl())
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
        err := writer.WriteObjectValue("creator", m.GetCreator())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
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
        err := writer.WriteStringValue("project_id", m.GetProjectId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("project_url", m.GetProjectUrl())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProjectCard) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchived sets the archived property value. Whether or not the card is archived
func (m *ProjectCard) SetArchived(value *bool)() {
    m.archived = value
}
// SetColumnName sets the column_name property value. The column_name property
func (m *ProjectCard) SetColumnName(value *string)() {
    m.column_name = value
}
// SetColumnUrl sets the column_url property value. The column_url property
func (m *ProjectCard) SetColumnUrl(value *string)() {
    m.column_url = value
}
// SetContentUrl sets the content_url property value. The content_url property
func (m *ProjectCard) SetContentUrl(value *string)() {
    m.content_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *ProjectCard) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCreator sets the creator property value. A GitHub user.
func (m *ProjectCard) SetCreator(value NullableSimpleUserable)() {
    m.creator = value
}
// SetId sets the id property value. The project card's ID
func (m *ProjectCard) SetId(value *int64)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *ProjectCard) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNote sets the note property value. The note property
func (m *ProjectCard) SetNote(value *string)() {
    m.note = value
}
// SetProjectId sets the project_id property value. The project_id property
func (m *ProjectCard) SetProjectId(value *string)() {
    m.project_id = value
}
// SetProjectUrl sets the project_url property value. The project_url property
func (m *ProjectCard) SetProjectUrl(value *string)() {
    m.project_url = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *ProjectCard) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *ProjectCard) SetUrl(value *string)() {
    m.url = value
}
type ProjectCardable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchived()(*bool)
    GetColumnName()(*string)
    GetColumnUrl()(*string)
    GetContentUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(NullableSimpleUserable)
    GetId()(*int64)
    GetNodeId()(*string)
    GetNote()(*string)
    GetProjectId()(*string)
    GetProjectUrl()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetArchived(value *bool)()
    SetColumnName(value *string)()
    SetColumnUrl(value *string)()
    SetContentUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value NullableSimpleUserable)()
    SetId(value *int64)()
    SetNodeId(value *string)()
    SetNote(value *string)()
    SetProjectId(value *string)()
    SetProjectUrl(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
