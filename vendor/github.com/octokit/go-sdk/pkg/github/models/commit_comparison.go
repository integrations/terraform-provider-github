package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommitComparison commit Comparison
type CommitComparison struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ahead_by property
    ahead_by *int32
    // Commit
    base_commit Commitable
    // The behind_by property
    behind_by *int32
    // The commits property
    commits []Commitable
    // The diff_url property
    diff_url *string
    // The files property
    files []DiffEntryable
    // The html_url property
    html_url *string
    // Commit
    merge_base_commit Commitable
    // The patch_url property
    patch_url *string
    // The permalink_url property
    permalink_url *string
    // The status property
    status *CommitComparison_status
    // The total_commits property
    total_commits *int32
    // The url property
    url *string
}
// NewCommitComparison instantiates a new CommitComparison and sets the default values.
func NewCommitComparison()(*CommitComparison) {
    m := &CommitComparison{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommitComparisonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommitComparisonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommitComparison(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommitComparison) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAheadBy gets the ahead_by property value. The ahead_by property
// returns a *int32 when successful
func (m *CommitComparison) GetAheadBy()(*int32) {
    return m.ahead_by
}
// GetBaseCommit gets the base_commit property value. Commit
// returns a Commitable when successful
func (m *CommitComparison) GetBaseCommit()(Commitable) {
    return m.base_commit
}
// GetBehindBy gets the behind_by property value. The behind_by property
// returns a *int32 when successful
func (m *CommitComparison) GetBehindBy()(*int32) {
    return m.behind_by
}
// GetCommits gets the commits property value. The commits property
// returns a []Commitable when successful
func (m *CommitComparison) GetCommits()([]Commitable) {
    return m.commits
}
// GetDiffUrl gets the diff_url property value. The diff_url property
// returns a *string when successful
func (m *CommitComparison) GetDiffUrl()(*string) {
    return m.diff_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommitComparison) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ahead_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAheadBy(val)
        }
        return nil
    }
    res["base_commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseCommit(val.(Commitable))
        }
        return nil
    }
    res["behind_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBehindBy(val)
        }
        return nil
    }
    res["commits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCommitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Commitable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Commitable)
                }
            }
            m.SetCommits(res)
        }
        return nil
    }
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
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDiffEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiffEntryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DiffEntryable)
                }
            }
            m.SetFiles(res)
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
    res["merge_base_commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeBaseCommit(val.(Commitable))
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
    res["permalink_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermalinkUrl(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCommitComparison_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CommitComparison_status))
        }
        return nil
    }
    res["total_commits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCommits(val)
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
// GetFiles gets the files property value. The files property
// returns a []DiffEntryable when successful
func (m *CommitComparison) GetFiles()([]DiffEntryable) {
    return m.files
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *CommitComparison) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetMergeBaseCommit gets the merge_base_commit property value. Commit
// returns a Commitable when successful
func (m *CommitComparison) GetMergeBaseCommit()(Commitable) {
    return m.merge_base_commit
}
// GetPatchUrl gets the patch_url property value. The patch_url property
// returns a *string when successful
func (m *CommitComparison) GetPatchUrl()(*string) {
    return m.patch_url
}
// GetPermalinkUrl gets the permalink_url property value. The permalink_url property
// returns a *string when successful
func (m *CommitComparison) GetPermalinkUrl()(*string) {
    return m.permalink_url
}
// GetStatus gets the status property value. The status property
// returns a *CommitComparison_status when successful
func (m *CommitComparison) GetStatus()(*CommitComparison_status) {
    return m.status
}
// GetTotalCommits gets the total_commits property value. The total_commits property
// returns a *int32 when successful
func (m *CommitComparison) GetTotalCommits()(*int32) {
    return m.total_commits
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *CommitComparison) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CommitComparison) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("ahead_by", m.GetAheadBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("base_commit", m.GetBaseCommit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("behind_by", m.GetBehindBy())
        if err != nil {
            return err
        }
    }
    if m.GetCommits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommits()))
        for i, v := range m.GetCommits() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("commits", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("diff_url", m.GetDiffUrl())
        if err != nil {
            return err
        }
    }
    if m.GetFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFiles()))
        for i, v := range m.GetFiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("files", cast)
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
        err := writer.WriteObjectValue("merge_base_commit", m.GetMergeBaseCommit())
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
        err := writer.WriteStringValue("permalink_url", m.GetPermalinkUrl())
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
        err := writer.WriteInt32Value("total_commits", m.GetTotalCommits())
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
func (m *CommitComparison) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAheadBy sets the ahead_by property value. The ahead_by property
func (m *CommitComparison) SetAheadBy(value *int32)() {
    m.ahead_by = value
}
// SetBaseCommit sets the base_commit property value. Commit
func (m *CommitComparison) SetBaseCommit(value Commitable)() {
    m.base_commit = value
}
// SetBehindBy sets the behind_by property value. The behind_by property
func (m *CommitComparison) SetBehindBy(value *int32)() {
    m.behind_by = value
}
// SetCommits sets the commits property value. The commits property
func (m *CommitComparison) SetCommits(value []Commitable)() {
    m.commits = value
}
// SetDiffUrl sets the diff_url property value. The diff_url property
func (m *CommitComparison) SetDiffUrl(value *string)() {
    m.diff_url = value
}
// SetFiles sets the files property value. The files property
func (m *CommitComparison) SetFiles(value []DiffEntryable)() {
    m.files = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *CommitComparison) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetMergeBaseCommit sets the merge_base_commit property value. Commit
func (m *CommitComparison) SetMergeBaseCommit(value Commitable)() {
    m.merge_base_commit = value
}
// SetPatchUrl sets the patch_url property value. The patch_url property
func (m *CommitComparison) SetPatchUrl(value *string)() {
    m.patch_url = value
}
// SetPermalinkUrl sets the permalink_url property value. The permalink_url property
func (m *CommitComparison) SetPermalinkUrl(value *string)() {
    m.permalink_url = value
}
// SetStatus sets the status property value. The status property
func (m *CommitComparison) SetStatus(value *CommitComparison_status)() {
    m.status = value
}
// SetTotalCommits sets the total_commits property value. The total_commits property
func (m *CommitComparison) SetTotalCommits(value *int32)() {
    m.total_commits = value
}
// SetUrl sets the url property value. The url property
func (m *CommitComparison) SetUrl(value *string)() {
    m.url = value
}
type CommitComparisonable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAheadBy()(*int32)
    GetBaseCommit()(Commitable)
    GetBehindBy()(*int32)
    GetCommits()([]Commitable)
    GetDiffUrl()(*string)
    GetFiles()([]DiffEntryable)
    GetHtmlUrl()(*string)
    GetMergeBaseCommit()(Commitable)
    GetPatchUrl()(*string)
    GetPermalinkUrl()(*string)
    GetStatus()(*CommitComparison_status)
    GetTotalCommits()(*int32)
    GetUrl()(*string)
    SetAheadBy(value *int32)()
    SetBaseCommit(value Commitable)()
    SetBehindBy(value *int32)()
    SetCommits(value []Commitable)()
    SetDiffUrl(value *string)()
    SetFiles(value []DiffEntryable)()
    SetHtmlUrl(value *string)()
    SetMergeBaseCommit(value Commitable)()
    SetPatchUrl(value *string)()
    SetPermalinkUrl(value *string)()
    SetStatus(value *CommitComparison_status)()
    SetTotalCommits(value *int32)()
    SetUrl(value *string)()
}
