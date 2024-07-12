package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationWikiCommit represents a 'wiki_commit' secret scanning location type. This location type shows that a secret was detected inside a commit to a repository wiki.
type SecretScanningLocationWikiCommit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // SHA-1 hash ID of the associated blob
    blob_sha *string
    // SHA-1 hash ID of the associated commit
    commit_sha *string
    // The GitHub URL to get the associated wiki commit
    commit_url *string
    // The column at which the secret ends within the end line when the file is interpreted as 8-bit ASCII.
    end_column *float64
    // Line number at which the secret ends in the file
    end_line *float64
    // The GitHub URL to get the associated wiki page
    page_url *string
    // The file path of the wiki page
    path *string
    // The column at which the secret starts within the start line when the file is interpreted as 8-bit ASCII.
    start_column *float64
    // Line number at which the secret starts in the file
    start_line *float64
}
// NewSecretScanningLocationWikiCommit instantiates a new SecretScanningLocationWikiCommit and sets the default values.
func NewSecretScanningLocationWikiCommit()(*SecretScanningLocationWikiCommit) {
    m := &SecretScanningLocationWikiCommit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationWikiCommitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationWikiCommitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocationWikiCommit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocationWikiCommit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlobSha gets the blob_sha property value. SHA-1 hash ID of the associated blob
// returns a *string when successful
func (m *SecretScanningLocationWikiCommit) GetBlobSha()(*string) {
    return m.blob_sha
}
// GetCommitSha gets the commit_sha property value. SHA-1 hash ID of the associated commit
// returns a *string when successful
func (m *SecretScanningLocationWikiCommit) GetCommitSha()(*string) {
    return m.commit_sha
}
// GetCommitUrl gets the commit_url property value. The GitHub URL to get the associated wiki commit
// returns a *string when successful
func (m *SecretScanningLocationWikiCommit) GetCommitUrl()(*string) {
    return m.commit_url
}
// GetEndColumn gets the end_column property value. The column at which the secret ends within the end line when the file is interpreted as 8-bit ASCII.
// returns a *float64 when successful
func (m *SecretScanningLocationWikiCommit) GetEndColumn()(*float64) {
    return m.end_column
}
// GetEndLine gets the end_line property value. Line number at which the secret ends in the file
// returns a *float64 when successful
func (m *SecretScanningLocationWikiCommit) GetEndLine()(*float64) {
    return m.end_line
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocationWikiCommit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blob_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlobSha(val)
        }
        return nil
    }
    res["commit_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitSha(val)
        }
        return nil
    }
    res["commit_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitUrl(val)
        }
        return nil
    }
    res["end_column"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndColumn(val)
        }
        return nil
    }
    res["end_line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndLine(val)
        }
        return nil
    }
    res["page_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageUrl(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["start_column"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartColumn(val)
        }
        return nil
    }
    res["start_line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartLine(val)
        }
        return nil
    }
    return res
}
// GetPageUrl gets the page_url property value. The GitHub URL to get the associated wiki page
// returns a *string when successful
func (m *SecretScanningLocationWikiCommit) GetPageUrl()(*string) {
    return m.page_url
}
// GetPath gets the path property value. The file path of the wiki page
// returns a *string when successful
func (m *SecretScanningLocationWikiCommit) GetPath()(*string) {
    return m.path
}
// GetStartColumn gets the start_column property value. The column at which the secret starts within the start line when the file is interpreted as 8-bit ASCII.
// returns a *float64 when successful
func (m *SecretScanningLocationWikiCommit) GetStartColumn()(*float64) {
    return m.start_column
}
// GetStartLine gets the start_line property value. Line number at which the secret starts in the file
// returns a *float64 when successful
func (m *SecretScanningLocationWikiCommit) GetStartLine()(*float64) {
    return m.start_line
}
// Serialize serializes information the current object
func (m *SecretScanningLocationWikiCommit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("blob_sha", m.GetBlobSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_sha", m.GetCommitSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_url", m.GetCommitUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("end_column", m.GetEndColumn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("end_line", m.GetEndLine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("page_url", m.GetPageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("start_column", m.GetStartColumn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("start_line", m.GetStartLine())
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
func (m *SecretScanningLocationWikiCommit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlobSha sets the blob_sha property value. SHA-1 hash ID of the associated blob
func (m *SecretScanningLocationWikiCommit) SetBlobSha(value *string)() {
    m.blob_sha = value
}
// SetCommitSha sets the commit_sha property value. SHA-1 hash ID of the associated commit
func (m *SecretScanningLocationWikiCommit) SetCommitSha(value *string)() {
    m.commit_sha = value
}
// SetCommitUrl sets the commit_url property value. The GitHub URL to get the associated wiki commit
func (m *SecretScanningLocationWikiCommit) SetCommitUrl(value *string)() {
    m.commit_url = value
}
// SetEndColumn sets the end_column property value. The column at which the secret ends within the end line when the file is interpreted as 8-bit ASCII.
func (m *SecretScanningLocationWikiCommit) SetEndColumn(value *float64)() {
    m.end_column = value
}
// SetEndLine sets the end_line property value. Line number at which the secret ends in the file
func (m *SecretScanningLocationWikiCommit) SetEndLine(value *float64)() {
    m.end_line = value
}
// SetPageUrl sets the page_url property value. The GitHub URL to get the associated wiki page
func (m *SecretScanningLocationWikiCommit) SetPageUrl(value *string)() {
    m.page_url = value
}
// SetPath sets the path property value. The file path of the wiki page
func (m *SecretScanningLocationWikiCommit) SetPath(value *string)() {
    m.path = value
}
// SetStartColumn sets the start_column property value. The column at which the secret starts within the start line when the file is interpreted as 8-bit ASCII.
func (m *SecretScanningLocationWikiCommit) SetStartColumn(value *float64)() {
    m.start_column = value
}
// SetStartLine sets the start_line property value. Line number at which the secret starts in the file
func (m *SecretScanningLocationWikiCommit) SetStartLine(value *float64)() {
    m.start_line = value
}
type SecretScanningLocationWikiCommitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlobSha()(*string)
    GetCommitSha()(*string)
    GetCommitUrl()(*string)
    GetEndColumn()(*float64)
    GetEndLine()(*float64)
    GetPageUrl()(*string)
    GetPath()(*string)
    GetStartColumn()(*float64)
    GetStartLine()(*float64)
    SetBlobSha(value *string)()
    SetCommitSha(value *string)()
    SetCommitUrl(value *string)()
    SetEndColumn(value *float64)()
    SetEndLine(value *float64)()
    SetPageUrl(value *string)()
    SetPath(value *string)()
    SetStartColumn(value *float64)()
    SetStartLine(value *float64)()
}
