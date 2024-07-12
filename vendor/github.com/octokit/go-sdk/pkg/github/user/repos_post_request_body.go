package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ReposPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether to allow Auto-merge to be used on pull requests.
    allow_auto_merge *bool
    // Whether to allow merge commits for pull requests.
    allow_merge_commit *bool
    // Whether to allow rebase merges for pull requests.
    allow_rebase_merge *bool
    // Whether to allow squash merges for pull requests.
    allow_squash_merge *bool
    // Whether the repository is initialized with a minimal README.
    auto_init *bool
    // Whether to delete head branches when pull requests are merged
    delete_branch_on_merge *bool
    // A short description of the repository.
    description *string
    // The desired language or platform to apply to the .gitignore.
    gitignore_template *string
    // Whether discussions are enabled.
    has_discussions *bool
    // Whether downloads are enabled.
    has_downloads *bool
    // Whether issues are enabled.
    has_issues *bool
    // Whether projects are enabled.
    has_projects *bool
    // Whether the wiki is enabled.
    has_wiki *bool
    // A URL with more information about the repository.
    homepage *string
    // Whether this repository acts as a template that can be used to generate new repositories.
    is_template *bool
    // The license keyword of the open source license for this repository.
    license_template *string
    // The name of the repository.
    name *string
    // Whether the repository is private.
    private *bool
    // The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization.
    team_id *int32
}
// NewReposPostRequestBody instantiates a new ReposPostRequestBody and sets the default values.
func NewReposPostRequestBody()(*ReposPostRequestBody) {
    m := &ReposPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReposPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReposPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReposPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReposPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowAutoMerge gets the allow_auto_merge property value. Whether to allow Auto-merge to be used on pull requests.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetAllowAutoMerge()(*bool) {
    return m.allow_auto_merge
}
// GetAllowMergeCommit gets the allow_merge_commit property value. Whether to allow merge commits for pull requests.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetAllowMergeCommit()(*bool) {
    return m.allow_merge_commit
}
// GetAllowRebaseMerge gets the allow_rebase_merge property value. Whether to allow rebase merges for pull requests.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetAllowRebaseMerge()(*bool) {
    return m.allow_rebase_merge
}
// GetAllowSquashMerge gets the allow_squash_merge property value. Whether to allow squash merges for pull requests.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetAllowSquashMerge()(*bool) {
    return m.allow_squash_merge
}
// GetAutoInit gets the auto_init property value. Whether the repository is initialized with a minimal README.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetAutoInit()(*bool) {
    return m.auto_init
}
// GetDeleteBranchOnMerge gets the delete_branch_on_merge property value. Whether to delete head branches when pull requests are merged
// returns a *bool when successful
func (m *ReposPostRequestBody) GetDeleteBranchOnMerge()(*bool) {
    return m.delete_branch_on_merge
}
// GetDescription gets the description property value. A short description of the repository.
// returns a *string when successful
func (m *ReposPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReposPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_auto_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAutoMerge(val)
        }
        return nil
    }
    res["allow_merge_commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMergeCommit(val)
        }
        return nil
    }
    res["allow_rebase_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowRebaseMerge(val)
        }
        return nil
    }
    res["allow_squash_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSquashMerge(val)
        }
        return nil
    }
    res["auto_init"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoInit(val)
        }
        return nil
    }
    res["delete_branch_on_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleteBranchOnMerge(val)
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
    res["gitignore_template"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitignoreTemplate(val)
        }
        return nil
    }
    res["has_discussions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasDiscussions(val)
        }
        return nil
    }
    res["has_downloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasDownloads(val)
        }
        return nil
    }
    res["has_issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasIssues(val)
        }
        return nil
    }
    res["has_projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasProjects(val)
        }
        return nil
    }
    res["has_wiki"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasWiki(val)
        }
        return nil
    }
    res["homepage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomepage(val)
        }
        return nil
    }
    res["is_template"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTemplate(val)
        }
        return nil
    }
    res["license_template"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseTemplate(val)
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
    res["private"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivate(val)
        }
        return nil
    }
    res["team_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    return res
}
// GetGitignoreTemplate gets the gitignore_template property value. The desired language or platform to apply to the .gitignore.
// returns a *string when successful
func (m *ReposPostRequestBody) GetGitignoreTemplate()(*string) {
    return m.gitignore_template
}
// GetHasDiscussions gets the has_discussions property value. Whether discussions are enabled.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetHasDiscussions()(*bool) {
    return m.has_discussions
}
// GetHasDownloads gets the has_downloads property value. Whether downloads are enabled.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetHasDownloads()(*bool) {
    return m.has_downloads
}
// GetHasIssues gets the has_issues property value. Whether issues are enabled.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetHasIssues()(*bool) {
    return m.has_issues
}
// GetHasProjects gets the has_projects property value. Whether projects are enabled.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetHasProjects()(*bool) {
    return m.has_projects
}
// GetHasWiki gets the has_wiki property value. Whether the wiki is enabled.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetHasWiki()(*bool) {
    return m.has_wiki
}
// GetHomepage gets the homepage property value. A URL with more information about the repository.
// returns a *string when successful
func (m *ReposPostRequestBody) GetHomepage()(*string) {
    return m.homepage
}
// GetIsTemplate gets the is_template property value. Whether this repository acts as a template that can be used to generate new repositories.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetIsTemplate()(*bool) {
    return m.is_template
}
// GetLicenseTemplate gets the license_template property value. The license keyword of the open source license for this repository.
// returns a *string when successful
func (m *ReposPostRequestBody) GetLicenseTemplate()(*string) {
    return m.license_template
}
// GetName gets the name property value. The name of the repository.
// returns a *string when successful
func (m *ReposPostRequestBody) GetName()(*string) {
    return m.name
}
// GetPrivate gets the private property value. Whether the repository is private.
// returns a *bool when successful
func (m *ReposPostRequestBody) GetPrivate()(*bool) {
    return m.private
}
// GetTeamId gets the team_id property value. The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization.
// returns a *int32 when successful
func (m *ReposPostRequestBody) GetTeamId()(*int32) {
    return m.team_id
}
// Serialize serializes information the current object
func (m *ReposPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allow_auto_merge", m.GetAllowAutoMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_merge_commit", m.GetAllowMergeCommit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_rebase_merge", m.GetAllowRebaseMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_squash_merge", m.GetAllowSquashMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("auto_init", m.GetAutoInit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("delete_branch_on_merge", m.GetDeleteBranchOnMerge())
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
        err := writer.WriteStringValue("gitignore_template", m.GetGitignoreTemplate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_discussions", m.GetHasDiscussions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_downloads", m.GetHasDownloads())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_issues", m.GetHasIssues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_projects", m.GetHasProjects())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_wiki", m.GetHasWiki())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("homepage", m.GetHomepage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_template", m.GetIsTemplate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("license_template", m.GetLicenseTemplate())
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
        err := writer.WriteBoolValue("private", m.GetPrivate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("team_id", m.GetTeamId())
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
func (m *ReposPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowAutoMerge sets the allow_auto_merge property value. Whether to allow Auto-merge to be used on pull requests.
func (m *ReposPostRequestBody) SetAllowAutoMerge(value *bool)() {
    m.allow_auto_merge = value
}
// SetAllowMergeCommit sets the allow_merge_commit property value. Whether to allow merge commits for pull requests.
func (m *ReposPostRequestBody) SetAllowMergeCommit(value *bool)() {
    m.allow_merge_commit = value
}
// SetAllowRebaseMerge sets the allow_rebase_merge property value. Whether to allow rebase merges for pull requests.
func (m *ReposPostRequestBody) SetAllowRebaseMerge(value *bool)() {
    m.allow_rebase_merge = value
}
// SetAllowSquashMerge sets the allow_squash_merge property value. Whether to allow squash merges for pull requests.
func (m *ReposPostRequestBody) SetAllowSquashMerge(value *bool)() {
    m.allow_squash_merge = value
}
// SetAutoInit sets the auto_init property value. Whether the repository is initialized with a minimal README.
func (m *ReposPostRequestBody) SetAutoInit(value *bool)() {
    m.auto_init = value
}
// SetDeleteBranchOnMerge sets the delete_branch_on_merge property value. Whether to delete head branches when pull requests are merged
func (m *ReposPostRequestBody) SetDeleteBranchOnMerge(value *bool)() {
    m.delete_branch_on_merge = value
}
// SetDescription sets the description property value. A short description of the repository.
func (m *ReposPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetGitignoreTemplate sets the gitignore_template property value. The desired language or platform to apply to the .gitignore.
func (m *ReposPostRequestBody) SetGitignoreTemplate(value *string)() {
    m.gitignore_template = value
}
// SetHasDiscussions sets the has_discussions property value. Whether discussions are enabled.
func (m *ReposPostRequestBody) SetHasDiscussions(value *bool)() {
    m.has_discussions = value
}
// SetHasDownloads sets the has_downloads property value. Whether downloads are enabled.
func (m *ReposPostRequestBody) SetHasDownloads(value *bool)() {
    m.has_downloads = value
}
// SetHasIssues sets the has_issues property value. Whether issues are enabled.
func (m *ReposPostRequestBody) SetHasIssues(value *bool)() {
    m.has_issues = value
}
// SetHasProjects sets the has_projects property value. Whether projects are enabled.
func (m *ReposPostRequestBody) SetHasProjects(value *bool)() {
    m.has_projects = value
}
// SetHasWiki sets the has_wiki property value. Whether the wiki is enabled.
func (m *ReposPostRequestBody) SetHasWiki(value *bool)() {
    m.has_wiki = value
}
// SetHomepage sets the homepage property value. A URL with more information about the repository.
func (m *ReposPostRequestBody) SetHomepage(value *string)() {
    m.homepage = value
}
// SetIsTemplate sets the is_template property value. Whether this repository acts as a template that can be used to generate new repositories.
func (m *ReposPostRequestBody) SetIsTemplate(value *bool)() {
    m.is_template = value
}
// SetLicenseTemplate sets the license_template property value. The license keyword of the open source license for this repository.
func (m *ReposPostRequestBody) SetLicenseTemplate(value *string)() {
    m.license_template = value
}
// SetName sets the name property value. The name of the repository.
func (m *ReposPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetPrivate sets the private property value. Whether the repository is private.
func (m *ReposPostRequestBody) SetPrivate(value *bool)() {
    m.private = value
}
// SetTeamId sets the team_id property value. The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization.
func (m *ReposPostRequestBody) SetTeamId(value *int32)() {
    m.team_id = value
}
type ReposPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAutoMerge()(*bool)
    GetAllowMergeCommit()(*bool)
    GetAllowRebaseMerge()(*bool)
    GetAllowSquashMerge()(*bool)
    GetAutoInit()(*bool)
    GetDeleteBranchOnMerge()(*bool)
    GetDescription()(*string)
    GetGitignoreTemplate()(*string)
    GetHasDiscussions()(*bool)
    GetHasDownloads()(*bool)
    GetHasIssues()(*bool)
    GetHasProjects()(*bool)
    GetHasWiki()(*bool)
    GetHomepage()(*string)
    GetIsTemplate()(*bool)
    GetLicenseTemplate()(*string)
    GetName()(*string)
    GetPrivate()(*bool)
    GetTeamId()(*int32)
    SetAllowAutoMerge(value *bool)()
    SetAllowMergeCommit(value *bool)()
    SetAllowRebaseMerge(value *bool)()
    SetAllowSquashMerge(value *bool)()
    SetAutoInit(value *bool)()
    SetDeleteBranchOnMerge(value *bool)()
    SetDescription(value *string)()
    SetGitignoreTemplate(value *string)()
    SetHasDiscussions(value *bool)()
    SetHasDownloads(value *bool)()
    SetHasIssues(value *bool)()
    SetHasProjects(value *bool)()
    SetHasWiki(value *bool)()
    SetHomepage(value *string)()
    SetIsTemplate(value *bool)()
    SetLicenseTemplate(value *string)()
    SetName(value *string)()
    SetPrivate(value *bool)()
    SetTeamId(value *int32)()
}
