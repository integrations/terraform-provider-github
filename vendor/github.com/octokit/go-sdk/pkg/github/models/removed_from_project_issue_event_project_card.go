package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RemovedFromProjectIssueEvent_project_card struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The column_name property
    column_name *string
    // The id property
    id *int32
    // The previous_column_name property
    previous_column_name *string
    // The project_id property
    project_id *int32
    // The project_url property
    project_url *string
    // The url property
    url *string
}
// NewRemovedFromProjectIssueEvent_project_card instantiates a new RemovedFromProjectIssueEvent_project_card and sets the default values.
func NewRemovedFromProjectIssueEvent_project_card()(*RemovedFromProjectIssueEvent_project_card) {
    m := &RemovedFromProjectIssueEvent_project_card{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRemovedFromProjectIssueEvent_project_cardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRemovedFromProjectIssueEvent_project_cardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemovedFromProjectIssueEvent_project_card(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RemovedFromProjectIssueEvent_project_card) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColumnName gets the column_name property value. The column_name property
// returns a *string when successful
func (m *RemovedFromProjectIssueEvent_project_card) GetColumnName()(*string) {
    return m.column_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RemovedFromProjectIssueEvent_project_card) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["previous_column_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousColumnName(val)
        }
        return nil
    }
    res["project_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *RemovedFromProjectIssueEvent_project_card) GetId()(*int32) {
    return m.id
}
// GetPreviousColumnName gets the previous_column_name property value. The previous_column_name property
// returns a *string when successful
func (m *RemovedFromProjectIssueEvent_project_card) GetPreviousColumnName()(*string) {
    return m.previous_column_name
}
// GetProjectId gets the project_id property value. The project_id property
// returns a *int32 when successful
func (m *RemovedFromProjectIssueEvent_project_card) GetProjectId()(*int32) {
    return m.project_id
}
// GetProjectUrl gets the project_url property value. The project_url property
// returns a *string when successful
func (m *RemovedFromProjectIssueEvent_project_card) GetProjectUrl()(*string) {
    return m.project_url
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *RemovedFromProjectIssueEvent_project_card) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *RemovedFromProjectIssueEvent_project_card) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("column_name", m.GetColumnName())
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
        err := writer.WriteStringValue("previous_column_name", m.GetPreviousColumnName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("project_id", m.GetProjectId())
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
func (m *RemovedFromProjectIssueEvent_project_card) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColumnName sets the column_name property value. The column_name property
func (m *RemovedFromProjectIssueEvent_project_card) SetColumnName(value *string)() {
    m.column_name = value
}
// SetId sets the id property value. The id property
func (m *RemovedFromProjectIssueEvent_project_card) SetId(value *int32)() {
    m.id = value
}
// SetPreviousColumnName sets the previous_column_name property value. The previous_column_name property
func (m *RemovedFromProjectIssueEvent_project_card) SetPreviousColumnName(value *string)() {
    m.previous_column_name = value
}
// SetProjectId sets the project_id property value. The project_id property
func (m *RemovedFromProjectIssueEvent_project_card) SetProjectId(value *int32)() {
    m.project_id = value
}
// SetProjectUrl sets the project_url property value. The project_url property
func (m *RemovedFromProjectIssueEvent_project_card) SetProjectUrl(value *string)() {
    m.project_url = value
}
// SetUrl sets the url property value. The url property
func (m *RemovedFromProjectIssueEvent_project_card) SetUrl(value *string)() {
    m.url = value
}
type RemovedFromProjectIssueEvent_project_cardable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColumnName()(*string)
    GetId()(*int32)
    GetPreviousColumnName()(*string)
    GetProjectId()(*int32)
    GetProjectUrl()(*string)
    GetUrl()(*string)
    SetColumnName(value *string)()
    SetId(value *int32)()
    SetPreviousColumnName(value *string)()
    SetProjectId(value *int32)()
    SetProjectUrl(value *string)()
    SetUrl(value *string)()
}
