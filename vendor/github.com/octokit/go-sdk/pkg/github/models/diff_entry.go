package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DiffEntry diff Entry
type DiffEntry struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The additions property
    additions *int32
    // The blob_url property
    blob_url *string
    // The changes property
    changes *int32
    // The contents_url property
    contents_url *string
    // The deletions property
    deletions *int32
    // The filename property
    filename *string
    // The patch property
    patch *string
    // The previous_filename property
    previous_filename *string
    // The raw_url property
    raw_url *string
    // The sha property
    sha *string
    // The status property
    status *DiffEntry_status
}
// NewDiffEntry instantiates a new DiffEntry and sets the default values.
func NewDiffEntry()(*DiffEntry) {
    m := &DiffEntry{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDiffEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDiffEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDiffEntry(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DiffEntry) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdditions gets the additions property value. The additions property
// returns a *int32 when successful
func (m *DiffEntry) GetAdditions()(*int32) {
    return m.additions
}
// GetBlobUrl gets the blob_url property value. The blob_url property
// returns a *string when successful
func (m *DiffEntry) GetBlobUrl()(*string) {
    return m.blob_url
}
// GetChanges gets the changes property value. The changes property
// returns a *int32 when successful
func (m *DiffEntry) GetChanges()(*int32) {
    return m.changes
}
// GetContentsUrl gets the contents_url property value. The contents_url property
// returns a *string when successful
func (m *DiffEntry) GetContentsUrl()(*string) {
    return m.contents_url
}
// GetDeletions gets the deletions property value. The deletions property
// returns a *int32 when successful
func (m *DiffEntry) GetDeletions()(*int32) {
    return m.deletions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DiffEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditions(val)
        }
        return nil
    }
    res["blob_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlobUrl(val)
        }
        return nil
    }
    res["changes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChanges(val)
        }
        return nil
    }
    res["contents_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentsUrl(val)
        }
        return nil
    }
    res["deletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletions(val)
        }
        return nil
    }
    res["filename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilename(val)
        }
        return nil
    }
    res["patch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatch(val)
        }
        return nil
    }
    res["previous_filename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousFilename(val)
        }
        return nil
    }
    res["raw_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRawUrl(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiffEntry_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DiffEntry_status))
        }
        return nil
    }
    return res
}
// GetFilename gets the filename property value. The filename property
// returns a *string when successful
func (m *DiffEntry) GetFilename()(*string) {
    return m.filename
}
// GetPatch gets the patch property value. The patch property
// returns a *string when successful
func (m *DiffEntry) GetPatch()(*string) {
    return m.patch
}
// GetPreviousFilename gets the previous_filename property value. The previous_filename property
// returns a *string when successful
func (m *DiffEntry) GetPreviousFilename()(*string) {
    return m.previous_filename
}
// GetRawUrl gets the raw_url property value. The raw_url property
// returns a *string when successful
func (m *DiffEntry) GetRawUrl()(*string) {
    return m.raw_url
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *DiffEntry) GetSha()(*string) {
    return m.sha
}
// GetStatus gets the status property value. The status property
// returns a *DiffEntry_status when successful
func (m *DiffEntry) GetStatus()(*DiffEntry_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *DiffEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("additions", m.GetAdditions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blob_url", m.GetBlobUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("changes", m.GetChanges())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contents_url", m.GetContentsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deletions", m.GetDeletions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filename", m.GetFilename())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("patch", m.GetPatch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previous_filename", m.GetPreviousFilename())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("raw_url", m.GetRawUrl())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *DiffEntry) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdditions sets the additions property value. The additions property
func (m *DiffEntry) SetAdditions(value *int32)() {
    m.additions = value
}
// SetBlobUrl sets the blob_url property value. The blob_url property
func (m *DiffEntry) SetBlobUrl(value *string)() {
    m.blob_url = value
}
// SetChanges sets the changes property value. The changes property
func (m *DiffEntry) SetChanges(value *int32)() {
    m.changes = value
}
// SetContentsUrl sets the contents_url property value. The contents_url property
func (m *DiffEntry) SetContentsUrl(value *string)() {
    m.contents_url = value
}
// SetDeletions sets the deletions property value. The deletions property
func (m *DiffEntry) SetDeletions(value *int32)() {
    m.deletions = value
}
// SetFilename sets the filename property value. The filename property
func (m *DiffEntry) SetFilename(value *string)() {
    m.filename = value
}
// SetPatch sets the patch property value. The patch property
func (m *DiffEntry) SetPatch(value *string)() {
    m.patch = value
}
// SetPreviousFilename sets the previous_filename property value. The previous_filename property
func (m *DiffEntry) SetPreviousFilename(value *string)() {
    m.previous_filename = value
}
// SetRawUrl sets the raw_url property value. The raw_url property
func (m *DiffEntry) SetRawUrl(value *string)() {
    m.raw_url = value
}
// SetSha sets the sha property value. The sha property
func (m *DiffEntry) SetSha(value *string)() {
    m.sha = value
}
// SetStatus sets the status property value. The status property
func (m *DiffEntry) SetStatus(value *DiffEntry_status)() {
    m.status = value
}
type DiffEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditions()(*int32)
    GetBlobUrl()(*string)
    GetChanges()(*int32)
    GetContentsUrl()(*string)
    GetDeletions()(*int32)
    GetFilename()(*string)
    GetPatch()(*string)
    GetPreviousFilename()(*string)
    GetRawUrl()(*string)
    GetSha()(*string)
    GetStatus()(*DiffEntry_status)
    SetAdditions(value *int32)()
    SetBlobUrl(value *string)()
    SetChanges(value *int32)()
    SetContentsUrl(value *string)()
    SetDeletions(value *int32)()
    SetFilename(value *string)()
    SetPatch(value *string)()
    SetPreviousFilename(value *string)()
    SetRawUrl(value *string)()
    SetSha(value *string)()
    SetStatus(value *DiffEntry_status)()
}
