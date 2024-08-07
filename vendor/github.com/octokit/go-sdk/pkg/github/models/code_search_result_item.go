package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeSearchResultItem code Search Result Item
type CodeSearchResultItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The file_size property
    file_size *int32
    // The git_url property
    git_url *string
    // The html_url property
    html_url *string
    // The language property
    language *string
    // The last_modified_at property
    last_modified_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The line_numbers property
    line_numbers []string
    // The name property
    name *string
    // The path property
    path *string
    // Minimal Repository
    repository MinimalRepositoryable
    // The score property
    score *float64
    // The sha property
    sha *string
    // The text_matches property
    text_matches []Codeable
    // The url property
    url *string
}
// NewCodeSearchResultItem instantiates a new CodeSearchResultItem and sets the default values.
func NewCodeSearchResultItem()(*CodeSearchResultItem) {
    m := &CodeSearchResultItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeSearchResultItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeSearchResultItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeSearchResultItem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeSearchResultItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeSearchResultItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["file_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSize(val)
        }
        return nil
    }
    res["git_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitUrl(val)
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
    res["last_modified_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedAt(val)
        }
        return nil
    }
    res["line_numbers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetLineNumbers(res)
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
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(MinimalRepositoryable))
        }
        return nil
    }
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
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
    res["text_matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Codeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Codeable)
                }
            }
            m.SetTextMatches(res)
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
// GetFileSize gets the file_size property value. The file_size property
// returns a *int32 when successful
func (m *CodeSearchResultItem) GetFileSize()(*int32) {
    return m.file_size
}
// GetGitUrl gets the git_url property value. The git_url property
// returns a *string when successful
func (m *CodeSearchResultItem) GetGitUrl()(*string) {
    return m.git_url
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *CodeSearchResultItem) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetLanguage gets the language property value. The language property
// returns a *string when successful
func (m *CodeSearchResultItem) GetLanguage()(*string) {
    return m.language
}
// GetLastModifiedAt gets the last_modified_at property value. The last_modified_at property
// returns a *Time when successful
func (m *CodeSearchResultItem) GetLastModifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_modified_at
}
// GetLineNumbers gets the line_numbers property value. The line_numbers property
// returns a []string when successful
func (m *CodeSearchResultItem) GetLineNumbers()([]string) {
    return m.line_numbers
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *CodeSearchResultItem) GetName()(*string) {
    return m.name
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *CodeSearchResultItem) GetPath()(*string) {
    return m.path
}
// GetRepository gets the repository property value. Minimal Repository
// returns a MinimalRepositoryable when successful
func (m *CodeSearchResultItem) GetRepository()(MinimalRepositoryable) {
    return m.repository
}
// GetScore gets the score property value. The score property
// returns a *float64 when successful
func (m *CodeSearchResultItem) GetScore()(*float64) {
    return m.score
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *CodeSearchResultItem) GetSha()(*string) {
    return m.sha
}
// GetTextMatches gets the text_matches property value. The text_matches property
// returns a []Codeable when successful
func (m *CodeSearchResultItem) GetTextMatches()([]Codeable) {
    return m.text_matches
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *CodeSearchResultItem) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CodeSearchResultItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("file_size", m.GetFileSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_url", m.GetGitUrl())
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
        err := writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("last_modified_at", m.GetLastModifiedAt())
        if err != nil {
            return err
        }
    }
    if m.GetLineNumbers() != nil {
        err := writer.WriteCollectionOfStringValues("line_numbers", m.GetLineNumbers())
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
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("score", m.GetScore())
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
    if m.GetTextMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTextMatches()))
        for i, v := range m.GetTextMatches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("text_matches", cast)
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
func (m *CodeSearchResultItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFileSize sets the file_size property value. The file_size property
func (m *CodeSearchResultItem) SetFileSize(value *int32)() {
    m.file_size = value
}
// SetGitUrl sets the git_url property value. The git_url property
func (m *CodeSearchResultItem) SetGitUrl(value *string)() {
    m.git_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *CodeSearchResultItem) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetLanguage sets the language property value. The language property
func (m *CodeSearchResultItem) SetLanguage(value *string)() {
    m.language = value
}
// SetLastModifiedAt sets the last_modified_at property value. The last_modified_at property
func (m *CodeSearchResultItem) SetLastModifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_modified_at = value
}
// SetLineNumbers sets the line_numbers property value. The line_numbers property
func (m *CodeSearchResultItem) SetLineNumbers(value []string)() {
    m.line_numbers = value
}
// SetName sets the name property value. The name property
func (m *CodeSearchResultItem) SetName(value *string)() {
    m.name = value
}
// SetPath sets the path property value. The path property
func (m *CodeSearchResultItem) SetPath(value *string)() {
    m.path = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *CodeSearchResultItem) SetRepository(value MinimalRepositoryable)() {
    m.repository = value
}
// SetScore sets the score property value. The score property
func (m *CodeSearchResultItem) SetScore(value *float64)() {
    m.score = value
}
// SetSha sets the sha property value. The sha property
func (m *CodeSearchResultItem) SetSha(value *string)() {
    m.sha = value
}
// SetTextMatches sets the text_matches property value. The text_matches property
func (m *CodeSearchResultItem) SetTextMatches(value []Codeable)() {
    m.text_matches = value
}
// SetUrl sets the url property value. The url property
func (m *CodeSearchResultItem) SetUrl(value *string)() {
    m.url = value
}
type CodeSearchResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileSize()(*int32)
    GetGitUrl()(*string)
    GetHtmlUrl()(*string)
    GetLanguage()(*string)
    GetLastModifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLineNumbers()([]string)
    GetName()(*string)
    GetPath()(*string)
    GetRepository()(MinimalRepositoryable)
    GetScore()(*float64)
    GetSha()(*string)
    GetTextMatches()([]Codeable)
    GetUrl()(*string)
    SetFileSize(value *int32)()
    SetGitUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetLanguage(value *string)()
    SetLastModifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLineNumbers(value []string)()
    SetName(value *string)()
    SetPath(value *string)()
    SetRepository(value MinimalRepositoryable)()
    SetScore(value *float64)()
    SetSha(value *string)()
    SetTextMatches(value []Codeable)()
    SetUrl(value *string)()
}
