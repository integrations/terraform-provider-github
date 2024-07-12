package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemReposPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Either `true` to allow auto-merge on pull requests, or `false` to disallow auto-merge.
    allow_auto_merge *bool
    // Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits.
    allow_merge_commit *bool
    // Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging.
    allow_rebase_merge *bool
    // Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging.
    allow_squash_merge *bool
    // Pass `true` to create an initial commit with empty README.
    auto_init *bool
    // The custom properties for the new repository. The keys are the custom property names, and the values are the corresponding custom property values.
    custom_properties ItemReposPostRequestBody_custom_propertiesable
    // Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion. **The authenticated user must be an organization owner to set this property to `true`.**
    delete_branch_on_merge *bool
    // A short description of the repository.
    description *string
    // Desired language or platform [.gitignore template](https://github.com/github/gitignore) to apply. Use the name of the template without the extension. For example, "Haskell".
    gitignore_template *string
    // Whether downloads are enabled.
    has_downloads *bool
    // Either `true` to enable issues for this repository or `false` to disable them.
    has_issues *bool
    // Either `true` to enable projects for this repository or `false` to disable them. **Note:** If you're creating a repository in an organization that has disabled repository projects, the default is `false`, and if you pass `true`, the API returns an error.
    has_projects *bool
    // Either `true` to enable the wiki for this repository or `false` to disable it.
    has_wiki *bool
    // A URL with more information about the repository.
    homepage *string
    // Either `true` to make this repo available as a template repository or `false` to prevent it.
    is_template *bool
    // Choose an [open source license template](https://choosealicense.com/) that best suits your needs, and then use the [license keyword](https://docs.github.com/articles/licensing-a-repository/#searching-github-by-license-type) as the `license_template` string. For example, "mit" or "mpl-2.0".
    license_template *string
    // The name of the repository.
    name *string
    // Whether the repository is private.
    private *bool
    // The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization.
    team_id *int32
    // Either `true` to allow squash-merge commits to use pull request title, or `false` to use commit message. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
    // Deprecated: 
    use_squash_pr_title_as_default *bool
}
// NewItemReposPostRequestBody instantiates a new ItemReposPostRequestBody and sets the default values.
func NewItemReposPostRequestBody()(*ItemReposPostRequestBody) {
    m := &ItemReposPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemReposPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemReposPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemReposPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemReposPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowAutoMerge gets the allow_auto_merge property value. Either `true` to allow auto-merge on pull requests, or `false` to disallow auto-merge.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetAllowAutoMerge()(*bool) {
    return m.allow_auto_merge
}
// GetAllowMergeCommit gets the allow_merge_commit property value. Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetAllowMergeCommit()(*bool) {
    return m.allow_merge_commit
}
// GetAllowRebaseMerge gets the allow_rebase_merge property value. Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetAllowRebaseMerge()(*bool) {
    return m.allow_rebase_merge
}
// GetAllowSquashMerge gets the allow_squash_merge property value. Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetAllowSquashMerge()(*bool) {
    return m.allow_squash_merge
}
// GetAutoInit gets the auto_init property value. Pass `true` to create an initial commit with empty README.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetAutoInit()(*bool) {
    return m.auto_init
}
// GetCustomProperties gets the custom_properties property value. The custom properties for the new repository. The keys are the custom property names, and the values are the corresponding custom property values.
// returns a ItemReposPostRequestBody_custom_propertiesable when successful
func (m *ItemReposPostRequestBody) GetCustomProperties()(ItemReposPostRequestBody_custom_propertiesable) {
    return m.custom_properties
}
// GetDeleteBranchOnMerge gets the delete_branch_on_merge property value. Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion. **The authenticated user must be an organization owner to set this property to `true`.**
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetDeleteBranchOnMerge()(*bool) {
    return m.delete_branch_on_merge
}
// GetDescription gets the description property value. A short description of the repository.
// returns a *string when successful
func (m *ItemReposPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemReposPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["custom_properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemReposPostRequestBody_custom_propertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomProperties(val.(ItemReposPostRequestBody_custom_propertiesable))
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
    res["use_squash_pr_title_as_default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSquashPrTitleAsDefault(val)
        }
        return nil
    }
    return res
}
// GetGitignoreTemplate gets the gitignore_template property value. Desired language or platform [.gitignore template](https://github.com/github/gitignore) to apply. Use the name of the template without the extension. For example, "Haskell".
// returns a *string when successful
func (m *ItemReposPostRequestBody) GetGitignoreTemplate()(*string) {
    return m.gitignore_template
}
// GetHasDownloads gets the has_downloads property value. Whether downloads are enabled.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetHasDownloads()(*bool) {
    return m.has_downloads
}
// GetHasIssues gets the has_issues property value. Either `true` to enable issues for this repository or `false` to disable them.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetHasIssues()(*bool) {
    return m.has_issues
}
// GetHasProjects gets the has_projects property value. Either `true` to enable projects for this repository or `false` to disable them. **Note:** If you're creating a repository in an organization that has disabled repository projects, the default is `false`, and if you pass `true`, the API returns an error.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetHasProjects()(*bool) {
    return m.has_projects
}
// GetHasWiki gets the has_wiki property value. Either `true` to enable the wiki for this repository or `false` to disable it.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetHasWiki()(*bool) {
    return m.has_wiki
}
// GetHomepage gets the homepage property value. A URL with more information about the repository.
// returns a *string when successful
func (m *ItemReposPostRequestBody) GetHomepage()(*string) {
    return m.homepage
}
// GetIsTemplate gets the is_template property value. Either `true` to make this repo available as a template repository or `false` to prevent it.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetIsTemplate()(*bool) {
    return m.is_template
}
// GetLicenseTemplate gets the license_template property value. Choose an [open source license template](https://choosealicense.com/) that best suits your needs, and then use the [license keyword](https://docs.github.com/articles/licensing-a-repository/#searching-github-by-license-type) as the `license_template` string. For example, "mit" or "mpl-2.0".
// returns a *string when successful
func (m *ItemReposPostRequestBody) GetLicenseTemplate()(*string) {
    return m.license_template
}
// GetName gets the name property value. The name of the repository.
// returns a *string when successful
func (m *ItemReposPostRequestBody) GetName()(*string) {
    return m.name
}
// GetPrivate gets the private property value. Whether the repository is private.
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetPrivate()(*bool) {
    return m.private
}
// GetTeamId gets the team_id property value. The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization.
// returns a *int32 when successful
func (m *ItemReposPostRequestBody) GetTeamId()(*int32) {
    return m.team_id
}
// GetUseSquashPrTitleAsDefault gets the use_squash_pr_title_as_default property value. Either `true` to allow squash-merge commits to use pull request title, or `false` to use commit message. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
// Deprecated: 
// returns a *bool when successful
func (m *ItemReposPostRequestBody) GetUseSquashPrTitleAsDefault()(*bool) {
    return m.use_squash_pr_title_as_default
}
// Serialize serializes information the current object
func (m *ItemReposPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("custom_properties", m.GetCustomProperties())
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
        err := writer.WriteBoolValue("use_squash_pr_title_as_default", m.GetUseSquashPrTitleAsDefault())
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
func (m *ItemReposPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowAutoMerge sets the allow_auto_merge property value. Either `true` to allow auto-merge on pull requests, or `false` to disallow auto-merge.
func (m *ItemReposPostRequestBody) SetAllowAutoMerge(value *bool)() {
    m.allow_auto_merge = value
}
// SetAllowMergeCommit sets the allow_merge_commit property value. Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits.
func (m *ItemReposPostRequestBody) SetAllowMergeCommit(value *bool)() {
    m.allow_merge_commit = value
}
// SetAllowRebaseMerge sets the allow_rebase_merge property value. Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging.
func (m *ItemReposPostRequestBody) SetAllowRebaseMerge(value *bool)() {
    m.allow_rebase_merge = value
}
// SetAllowSquashMerge sets the allow_squash_merge property value. Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging.
func (m *ItemReposPostRequestBody) SetAllowSquashMerge(value *bool)() {
    m.allow_squash_merge = value
}
// SetAutoInit sets the auto_init property value. Pass `true` to create an initial commit with empty README.
func (m *ItemReposPostRequestBody) SetAutoInit(value *bool)() {
    m.auto_init = value
}
// SetCustomProperties sets the custom_properties property value. The custom properties for the new repository. The keys are the custom property names, and the values are the corresponding custom property values.
func (m *ItemReposPostRequestBody) SetCustomProperties(value ItemReposPostRequestBody_custom_propertiesable)() {
    m.custom_properties = value
}
// SetDeleteBranchOnMerge sets the delete_branch_on_merge property value. Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion. **The authenticated user must be an organization owner to set this property to `true`.**
func (m *ItemReposPostRequestBody) SetDeleteBranchOnMerge(value *bool)() {
    m.delete_branch_on_merge = value
}
// SetDescription sets the description property value. A short description of the repository.
func (m *ItemReposPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetGitignoreTemplate sets the gitignore_template property value. Desired language or platform [.gitignore template](https://github.com/github/gitignore) to apply. Use the name of the template without the extension. For example, "Haskell".
func (m *ItemReposPostRequestBody) SetGitignoreTemplate(value *string)() {
    m.gitignore_template = value
}
// SetHasDownloads sets the has_downloads property value. Whether downloads are enabled.
func (m *ItemReposPostRequestBody) SetHasDownloads(value *bool)() {
    m.has_downloads = value
}
// SetHasIssues sets the has_issues property value. Either `true` to enable issues for this repository or `false` to disable them.
func (m *ItemReposPostRequestBody) SetHasIssues(value *bool)() {
    m.has_issues = value
}
// SetHasProjects sets the has_projects property value. Either `true` to enable projects for this repository or `false` to disable them. **Note:** If you're creating a repository in an organization that has disabled repository projects, the default is `false`, and if you pass `true`, the API returns an error.
func (m *ItemReposPostRequestBody) SetHasProjects(value *bool)() {
    m.has_projects = value
}
// SetHasWiki sets the has_wiki property value. Either `true` to enable the wiki for this repository or `false` to disable it.
func (m *ItemReposPostRequestBody) SetHasWiki(value *bool)() {
    m.has_wiki = value
}
// SetHomepage sets the homepage property value. A URL with more information about the repository.
func (m *ItemReposPostRequestBody) SetHomepage(value *string)() {
    m.homepage = value
}
// SetIsTemplate sets the is_template property value. Either `true` to make this repo available as a template repository or `false` to prevent it.
func (m *ItemReposPostRequestBody) SetIsTemplate(value *bool)() {
    m.is_template = value
}
// SetLicenseTemplate sets the license_template property value. Choose an [open source license template](https://choosealicense.com/) that best suits your needs, and then use the [license keyword](https://docs.github.com/articles/licensing-a-repository/#searching-github-by-license-type) as the `license_template` string. For example, "mit" or "mpl-2.0".
func (m *ItemReposPostRequestBody) SetLicenseTemplate(value *string)() {
    m.license_template = value
}
// SetName sets the name property value. The name of the repository.
func (m *ItemReposPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetPrivate sets the private property value. Whether the repository is private.
func (m *ItemReposPostRequestBody) SetPrivate(value *bool)() {
    m.private = value
}
// SetTeamId sets the team_id property value. The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization.
func (m *ItemReposPostRequestBody) SetTeamId(value *int32)() {
    m.team_id = value
}
// SetUseSquashPrTitleAsDefault sets the use_squash_pr_title_as_default property value. Either `true` to allow squash-merge commits to use pull request title, or `false` to use commit message. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
// Deprecated: 
func (m *ItemReposPostRequestBody) SetUseSquashPrTitleAsDefault(value *bool)() {
    m.use_squash_pr_title_as_default = value
}
type ItemReposPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAutoMerge()(*bool)
    GetAllowMergeCommit()(*bool)
    GetAllowRebaseMerge()(*bool)
    GetAllowSquashMerge()(*bool)
    GetAutoInit()(*bool)
    GetCustomProperties()(ItemReposPostRequestBody_custom_propertiesable)
    GetDeleteBranchOnMerge()(*bool)
    GetDescription()(*string)
    GetGitignoreTemplate()(*string)
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
    GetUseSquashPrTitleAsDefault()(*bool)
    SetAllowAutoMerge(value *bool)()
    SetAllowMergeCommit(value *bool)()
    SetAllowRebaseMerge(value *bool)()
    SetAllowSquashMerge(value *bool)()
    SetAutoInit(value *bool)()
    SetCustomProperties(value ItemReposPostRequestBody_custom_propertiesable)()
    SetDeleteBranchOnMerge(value *bool)()
    SetDescription(value *string)()
    SetGitignoreTemplate(value *string)()
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
    SetUseSquashPrTitleAsDefault(value *bool)()
}
