package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportEscaped a repository import from an external source.
type ImportEscaped struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The authors_count property
    authors_count *int32
    // The authors_url property
    authors_url *string
    // The commit_count property
    commit_count *int32
    // The error_message property
    error_message *string
    // The failed_step property
    failed_step *string
    // The has_large_files property
    has_large_files *bool
    // The html_url property
    html_url *string
    // The import_percent property
    import_percent *int32
    // The large_files_count property
    large_files_count *int32
    // The large_files_size property
    large_files_size *int32
    // The message property
    message *string
    // The project_choices property
    project_choices []Import_project_choicesable
    // The push_percent property
    push_percent *int32
    // The repository_url property
    repository_url *string
    // The status property
    status *Import_status
    // The status_text property
    status_text *string
    // The svc_root property
    svc_root *string
    // The svn_root property
    svn_root *string
    // The tfvc_project property
    tfvc_project *string
    // The url property
    url *string
    // The use_lfs property
    use_lfs *bool
    // The vcs property
    vcs *string
    // The URL of the originating repository.
    vcs_url *string
}
// NewImportEscaped instantiates a new ImportEscaped and sets the default values.
func NewImportEscaped()(*ImportEscaped) {
    m := &ImportEscaped{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateImportEscapedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateImportEscapedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportEscaped(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ImportEscaped) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorsCount gets the authors_count property value. The authors_count property
// returns a *int32 when successful
func (m *ImportEscaped) GetAuthorsCount()(*int32) {
    return m.authors_count
}
// GetAuthorsUrl gets the authors_url property value. The authors_url property
// returns a *string when successful
func (m *ImportEscaped) GetAuthorsUrl()(*string) {
    return m.authors_url
}
// GetCommitCount gets the commit_count property value. The commit_count property
// returns a *int32 when successful
func (m *ImportEscaped) GetCommitCount()(*int32) {
    return m.commit_count
}
// GetErrorMessage gets the error_message property value. The error_message property
// returns a *string when successful
func (m *ImportEscaped) GetErrorMessage()(*string) {
    return m.error_message
}
// GetFailedStep gets the failed_step property value. The failed_step property
// returns a *string when successful
func (m *ImportEscaped) GetFailedStep()(*string) {
    return m.failed_step
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ImportEscaped) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authors_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorsCount(val)
        }
        return nil
    }
    res["authors_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorsUrl(val)
        }
        return nil
    }
    res["commit_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitCount(val)
        }
        return nil
    }
    res["error_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
        }
        return nil
    }
    res["failed_step"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedStep(val)
        }
        return nil
    }
    res["has_large_files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasLargeFiles(val)
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
    res["import_percent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportPercent(val)
        }
        return nil
    }
    res["large_files_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLargeFilesCount(val)
        }
        return nil
    }
    res["large_files_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLargeFilesSize(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["project_choices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImport_project_choicesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Import_project_choicesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Import_project_choicesable)
                }
            }
            m.SetProjectChoices(res)
        }
        return nil
    }
    res["push_percent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushPercent(val)
        }
        return nil
    }
    res["repository_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryUrl(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImport_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*Import_status))
        }
        return nil
    }
    res["status_text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusText(val)
        }
        return nil
    }
    res["svc_root"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSvcRoot(val)
        }
        return nil
    }
    res["svn_root"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSvnRoot(val)
        }
        return nil
    }
    res["tfvc_project"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTfvcProject(val)
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
    res["use_lfs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseLfs(val)
        }
        return nil
    }
    res["vcs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVcs(val)
        }
        return nil
    }
    res["vcs_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVcsUrl(val)
        }
        return nil
    }
    return res
}
// GetHasLargeFiles gets the has_large_files property value. The has_large_files property
// returns a *bool when successful
func (m *ImportEscaped) GetHasLargeFiles()(*bool) {
    return m.has_large_files
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *ImportEscaped) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetImportPercent gets the import_percent property value. The import_percent property
// returns a *int32 when successful
func (m *ImportEscaped) GetImportPercent()(*int32) {
    return m.import_percent
}
// GetLargeFilesCount gets the large_files_count property value. The large_files_count property
// returns a *int32 when successful
func (m *ImportEscaped) GetLargeFilesCount()(*int32) {
    return m.large_files_count
}
// GetLargeFilesSize gets the large_files_size property value. The large_files_size property
// returns a *int32 when successful
func (m *ImportEscaped) GetLargeFilesSize()(*int32) {
    return m.large_files_size
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *ImportEscaped) GetMessage()(*string) {
    return m.message
}
// GetProjectChoices gets the project_choices property value. The project_choices property
// returns a []Import_project_choicesable when successful
func (m *ImportEscaped) GetProjectChoices()([]Import_project_choicesable) {
    return m.project_choices
}
// GetPushPercent gets the push_percent property value. The push_percent property
// returns a *int32 when successful
func (m *ImportEscaped) GetPushPercent()(*int32) {
    return m.push_percent
}
// GetRepositoryUrl gets the repository_url property value. The repository_url property
// returns a *string when successful
func (m *ImportEscaped) GetRepositoryUrl()(*string) {
    return m.repository_url
}
// GetStatus gets the status property value. The status property
// returns a *Import_status when successful
func (m *ImportEscaped) GetStatus()(*Import_status) {
    return m.status
}
// GetStatusText gets the status_text property value. The status_text property
// returns a *string when successful
func (m *ImportEscaped) GetStatusText()(*string) {
    return m.status_text
}
// GetSvcRoot gets the svc_root property value. The svc_root property
// returns a *string when successful
func (m *ImportEscaped) GetSvcRoot()(*string) {
    return m.svc_root
}
// GetSvnRoot gets the svn_root property value. The svn_root property
// returns a *string when successful
func (m *ImportEscaped) GetSvnRoot()(*string) {
    return m.svn_root
}
// GetTfvcProject gets the tfvc_project property value. The tfvc_project property
// returns a *string when successful
func (m *ImportEscaped) GetTfvcProject()(*string) {
    return m.tfvc_project
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ImportEscaped) GetUrl()(*string) {
    return m.url
}
// GetUseLfs gets the use_lfs property value. The use_lfs property
// returns a *bool when successful
func (m *ImportEscaped) GetUseLfs()(*bool) {
    return m.use_lfs
}
// GetVcs gets the vcs property value. The vcs property
// returns a *string when successful
func (m *ImportEscaped) GetVcs()(*string) {
    return m.vcs
}
// GetVcsUrl gets the vcs_url property value. The URL of the originating repository.
// returns a *string when successful
func (m *ImportEscaped) GetVcsUrl()(*string) {
    return m.vcs_url
}
// Serialize serializes information the current object
func (m *ImportEscaped) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("authors_count", m.GetAuthorsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authors_url", m.GetAuthorsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("commit_count", m.GetCommitCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error_message", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failed_step", m.GetFailedStep())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_large_files", m.GetHasLargeFiles())
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
        err := writer.WriteInt32Value("import_percent", m.GetImportPercent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("large_files_count", m.GetLargeFilesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("large_files_size", m.GetLargeFilesSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    if m.GetProjectChoices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProjectChoices()))
        for i, v := range m.GetProjectChoices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("project_choices", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("push_percent", m.GetPushPercent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repository_url", m.GetRepositoryUrl())
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
        err := writer.WriteStringValue("status_text", m.GetStatusText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("svc_root", m.GetSvcRoot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("svn_root", m.GetSvnRoot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tfvc_project", m.GetTfvcProject())
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
        err := writer.WriteBoolValue("use_lfs", m.GetUseLfs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vcs", m.GetVcs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vcs_url", m.GetVcsUrl())
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
func (m *ImportEscaped) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorsCount sets the authors_count property value. The authors_count property
func (m *ImportEscaped) SetAuthorsCount(value *int32)() {
    m.authors_count = value
}
// SetAuthorsUrl sets the authors_url property value. The authors_url property
func (m *ImportEscaped) SetAuthorsUrl(value *string)() {
    m.authors_url = value
}
// SetCommitCount sets the commit_count property value. The commit_count property
func (m *ImportEscaped) SetCommitCount(value *int32)() {
    m.commit_count = value
}
// SetErrorMessage sets the error_message property value. The error_message property
func (m *ImportEscaped) SetErrorMessage(value *string)() {
    m.error_message = value
}
// SetFailedStep sets the failed_step property value. The failed_step property
func (m *ImportEscaped) SetFailedStep(value *string)() {
    m.failed_step = value
}
// SetHasLargeFiles sets the has_large_files property value. The has_large_files property
func (m *ImportEscaped) SetHasLargeFiles(value *bool)() {
    m.has_large_files = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *ImportEscaped) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetImportPercent sets the import_percent property value. The import_percent property
func (m *ImportEscaped) SetImportPercent(value *int32)() {
    m.import_percent = value
}
// SetLargeFilesCount sets the large_files_count property value. The large_files_count property
func (m *ImportEscaped) SetLargeFilesCount(value *int32)() {
    m.large_files_count = value
}
// SetLargeFilesSize sets the large_files_size property value. The large_files_size property
func (m *ImportEscaped) SetLargeFilesSize(value *int32)() {
    m.large_files_size = value
}
// SetMessage sets the message property value. The message property
func (m *ImportEscaped) SetMessage(value *string)() {
    m.message = value
}
// SetProjectChoices sets the project_choices property value. The project_choices property
func (m *ImportEscaped) SetProjectChoices(value []Import_project_choicesable)() {
    m.project_choices = value
}
// SetPushPercent sets the push_percent property value. The push_percent property
func (m *ImportEscaped) SetPushPercent(value *int32)() {
    m.push_percent = value
}
// SetRepositoryUrl sets the repository_url property value. The repository_url property
func (m *ImportEscaped) SetRepositoryUrl(value *string)() {
    m.repository_url = value
}
// SetStatus sets the status property value. The status property
func (m *ImportEscaped) SetStatus(value *Import_status)() {
    m.status = value
}
// SetStatusText sets the status_text property value. The status_text property
func (m *ImportEscaped) SetStatusText(value *string)() {
    m.status_text = value
}
// SetSvcRoot sets the svc_root property value. The svc_root property
func (m *ImportEscaped) SetSvcRoot(value *string)() {
    m.svc_root = value
}
// SetSvnRoot sets the svn_root property value. The svn_root property
func (m *ImportEscaped) SetSvnRoot(value *string)() {
    m.svn_root = value
}
// SetTfvcProject sets the tfvc_project property value. The tfvc_project property
func (m *ImportEscaped) SetTfvcProject(value *string)() {
    m.tfvc_project = value
}
// SetUrl sets the url property value. The url property
func (m *ImportEscaped) SetUrl(value *string)() {
    m.url = value
}
// SetUseLfs sets the use_lfs property value. The use_lfs property
func (m *ImportEscaped) SetUseLfs(value *bool)() {
    m.use_lfs = value
}
// SetVcs sets the vcs property value. The vcs property
func (m *ImportEscaped) SetVcs(value *string)() {
    m.vcs = value
}
// SetVcsUrl sets the vcs_url property value. The URL of the originating repository.
func (m *ImportEscaped) SetVcsUrl(value *string)() {
    m.vcs_url = value
}
type ImportEscapedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorsCount()(*int32)
    GetAuthorsUrl()(*string)
    GetCommitCount()(*int32)
    GetErrorMessage()(*string)
    GetFailedStep()(*string)
    GetHasLargeFiles()(*bool)
    GetHtmlUrl()(*string)
    GetImportPercent()(*int32)
    GetLargeFilesCount()(*int32)
    GetLargeFilesSize()(*int32)
    GetMessage()(*string)
    GetProjectChoices()([]Import_project_choicesable)
    GetPushPercent()(*int32)
    GetRepositoryUrl()(*string)
    GetStatus()(*Import_status)
    GetStatusText()(*string)
    GetSvcRoot()(*string)
    GetSvnRoot()(*string)
    GetTfvcProject()(*string)
    GetUrl()(*string)
    GetUseLfs()(*bool)
    GetVcs()(*string)
    GetVcsUrl()(*string)
    SetAuthorsCount(value *int32)()
    SetAuthorsUrl(value *string)()
    SetCommitCount(value *int32)()
    SetErrorMessage(value *string)()
    SetFailedStep(value *string)()
    SetHasLargeFiles(value *bool)()
    SetHtmlUrl(value *string)()
    SetImportPercent(value *int32)()
    SetLargeFilesCount(value *int32)()
    SetLargeFilesSize(value *int32)()
    SetMessage(value *string)()
    SetProjectChoices(value []Import_project_choicesable)()
    SetPushPercent(value *int32)()
    SetRepositoryUrl(value *string)()
    SetStatus(value *Import_status)()
    SetStatusText(value *string)()
    SetSvcRoot(value *string)()
    SetSvnRoot(value *string)()
    SetTfvcProject(value *string)()
    SetUrl(value *string)()
    SetUseLfs(value *bool)()
    SetVcs(value *string)()
    SetVcsUrl(value *string)()
}
