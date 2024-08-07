package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Issue_pull_request struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The diff_url property
    diff_url *string
    // The html_url property
    html_url *string
    // The merged_at property
    merged_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The patch_url property
    patch_url *string
    // The url property
    url *string
}
// NewIssue_pull_request instantiates a new Issue_pull_request and sets the default values.
func NewIssue_pull_request()(*Issue_pull_request) {
    m := &Issue_pull_request{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIssue_pull_requestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssue_pull_requestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIssue_pull_request(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Issue_pull_request) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDiffUrl gets the diff_url property value. The diff_url property
// returns a *string when successful
func (m *Issue_pull_request) GetDiffUrl()(*string) {
    return m.diff_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Issue_pull_request) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["diff_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiffUrl(val)
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
    res["merged_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergedAt(val)
        }
        return nil
    }
    res["patch_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatchUrl(val)
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
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Issue_pull_request) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetMergedAt gets the merged_at property value. The merged_at property
// returns a *Time when successful
func (m *Issue_pull_request) GetMergedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.merged_at
}
// GetPatchUrl gets the patch_url property value. The patch_url property
// returns a *string when successful
func (m *Issue_pull_request) GetPatchUrl()(*string) {
    return m.patch_url
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Issue_pull_request) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Issue_pull_request) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("diff_url", m.GetDiffUrl())
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
        err := writer.WriteTimeValue("merged_at", m.GetMergedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("patch_url", m.GetPatchUrl())
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
func (m *Issue_pull_request) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDiffUrl sets the diff_url property value. The diff_url property
func (m *Issue_pull_request) SetDiffUrl(value *string)() {
    m.diff_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Issue_pull_request) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetMergedAt sets the merged_at property value. The merged_at property
func (m *Issue_pull_request) SetMergedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.merged_at = value
}
// SetPatchUrl sets the patch_url property value. The patch_url property
func (m *Issue_pull_request) SetPatchUrl(value *string)() {
    m.patch_url = value
}
// SetUrl sets the url property value. The url property
func (m *Issue_pull_request) SetUrl(value *string)() {
    m.url = value
}
type Issue_pull_requestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDiffUrl()(*string)
    GetHtmlUrl()(*string)
    GetMergedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPatchUrl()(*string)
    GetUrl()(*string)
    SetDiffUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetMergedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPatchUrl(value *string)()
    SetUrl(value *string)()
}
