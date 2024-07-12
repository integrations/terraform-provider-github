package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemOwnerPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Either `true` to allow auto-merge on pull requests, or `false` to disallow auto-merge.
    allow_auto_merge *bool
    // Either `true` to allow private forks, or `false` to prevent private forks.
    allow_forking *bool
    // Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits.
    allow_merge_commit *bool
    // Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging.
    allow_rebase_merge *bool
    // Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging.
    allow_squash_merge *bool
    // Either `true` to always allow a pull request head branch that is behind its base branch to be updated even if it is not required to be up to date before merging, or false otherwise.
    allow_update_branch *bool
    // Whether to archive this repository. `false` will unarchive a previously archived repository.
    archived *bool
    // Updates the default branch for this repository.
    default_branch *string
    // Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion.
    delete_branch_on_merge *bool
    // A short description of the repository.
    description *string
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
    // The name of the repository.
    name *string
    // Either `true` to make the repository private or `false` to make it public. Default: `false`.  **Note**: You will get a `422` error if the organization restricts [changing repository visibility](https://docs.github.com/articles/repository-permission-levels-for-an-organization#changing-the-visibility-of-repositories) to organization owners and a non-owner tries to change the value of private.
    private *bool
    // Specify which security and analysis features to enable or disable for the repository.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."For example, to enable GitHub Advanced Security, use this data in the body of the `PATCH` request:`{ "security_and_analysis": {"advanced_security": { "status": "enabled" } } }`.You can check which security and analysis features are currently enabled by using a `GET /repos/{owner}/{repo}` request.
    security_and_analysis ItemItemOwnerPatchRequestBody_security_and_analysisable
    // Either `true` to allow squash-merge commits to use pull request title, or `false` to use commit message. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
    // Deprecated: 
    use_squash_pr_title_as_default *bool
    // Either `true` to require contributors to sign off on web-based commits, or `false` to not require contributors to sign off on web-based commits.
    web_commit_signoff_required *bool
}
// NewItemItemOwnerPatchRequestBody instantiates a new ItemItemOwnerPatchRequestBody and sets the default values.
func NewItemItemOwnerPatchRequestBody()(*ItemItemOwnerPatchRequestBody) {
    m := &ItemItemOwnerPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemOwnerPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemOwnerPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemOwnerPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemOwnerPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowAutoMerge gets the allow_auto_merge property value. Either `true` to allow auto-merge on pull requests, or `false` to disallow auto-merge.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetAllowAutoMerge()(*bool) {
    return m.allow_auto_merge
}
// GetAllowForking gets the allow_forking property value. Either `true` to allow private forks, or `false` to prevent private forks.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetAllowForking()(*bool) {
    return m.allow_forking
}
// GetAllowMergeCommit gets the allow_merge_commit property value. Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetAllowMergeCommit()(*bool) {
    return m.allow_merge_commit
}
// GetAllowRebaseMerge gets the allow_rebase_merge property value. Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetAllowRebaseMerge()(*bool) {
    return m.allow_rebase_merge
}
// GetAllowSquashMerge gets the allow_squash_merge property value. Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetAllowSquashMerge()(*bool) {
    return m.allow_squash_merge
}
// GetAllowUpdateBranch gets the allow_update_branch property value. Either `true` to always allow a pull request head branch that is behind its base branch to be updated even if it is not required to be up to date before merging, or false otherwise.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetAllowUpdateBranch()(*bool) {
    return m.allow_update_branch
}
// GetArchived gets the archived property value. Whether to archive this repository. `false` will unarchive a previously archived repository.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetArchived()(*bool) {
    return m.archived
}
// GetDefaultBranch gets the default_branch property value. Updates the default branch for this repository.
// returns a *string when successful
func (m *ItemItemOwnerPatchRequestBody) GetDefaultBranch()(*string) {
    return m.default_branch
}
// GetDeleteBranchOnMerge gets the delete_branch_on_merge property value. Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetDeleteBranchOnMerge()(*bool) {
    return m.delete_branch_on_merge
}
// GetDescription gets the description property value. A short description of the repository.
// returns a *string when successful
func (m *ItemItemOwnerPatchRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemOwnerPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["allow_forking"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowForking(val)
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
    res["allow_update_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUpdateBranch(val)
        }
        return nil
    }
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
    res["default_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultBranch(val)
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
    res["security_and_analysis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemOwnerPatchRequestBody_security_and_analysisFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityAndAnalysis(val.(ItemItemOwnerPatchRequestBody_security_and_analysisable))
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
    res["web_commit_signoff_required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebCommitSignoffRequired(val)
        }
        return nil
    }
    return res
}
// GetHasIssues gets the has_issues property value. Either `true` to enable issues for this repository or `false` to disable them.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetHasIssues()(*bool) {
    return m.has_issues
}
// GetHasProjects gets the has_projects property value. Either `true` to enable projects for this repository or `false` to disable them. **Note:** If you're creating a repository in an organization that has disabled repository projects, the default is `false`, and if you pass `true`, the API returns an error.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetHasProjects()(*bool) {
    return m.has_projects
}
// GetHasWiki gets the has_wiki property value. Either `true` to enable the wiki for this repository or `false` to disable it.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetHasWiki()(*bool) {
    return m.has_wiki
}
// GetHomepage gets the homepage property value. A URL with more information about the repository.
// returns a *string when successful
func (m *ItemItemOwnerPatchRequestBody) GetHomepage()(*string) {
    return m.homepage
}
// GetIsTemplate gets the is_template property value. Either `true` to make this repo available as a template repository or `false` to prevent it.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetIsTemplate()(*bool) {
    return m.is_template
}
// GetName gets the name property value. The name of the repository.
// returns a *string when successful
func (m *ItemItemOwnerPatchRequestBody) GetName()(*string) {
    return m.name
}
// GetPrivate gets the private property value. Either `true` to make the repository private or `false` to make it public. Default: `false`.  **Note**: You will get a `422` error if the organization restricts [changing repository visibility](https://docs.github.com/articles/repository-permission-levels-for-an-organization#changing-the-visibility-of-repositories) to organization owners and a non-owner tries to change the value of private.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetPrivate()(*bool) {
    return m.private
}
// GetSecurityAndAnalysis gets the security_and_analysis property value. Specify which security and analysis features to enable or disable for the repository.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."For example, to enable GitHub Advanced Security, use this data in the body of the `PATCH` request:`{ "security_and_analysis": {"advanced_security": { "status": "enabled" } } }`.You can check which security and analysis features are currently enabled by using a `GET /repos/{owner}/{repo}` request.
// returns a ItemItemOwnerPatchRequestBody_security_and_analysisable when successful
func (m *ItemItemOwnerPatchRequestBody) GetSecurityAndAnalysis()(ItemItemOwnerPatchRequestBody_security_and_analysisable) {
    return m.security_and_analysis
}
// GetUseSquashPrTitleAsDefault gets the use_squash_pr_title_as_default property value. Either `true` to allow squash-merge commits to use pull request title, or `false` to use commit message. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
// Deprecated: 
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetUseSquashPrTitleAsDefault()(*bool) {
    return m.use_squash_pr_title_as_default
}
// GetWebCommitSignoffRequired gets the web_commit_signoff_required property value. Either `true` to require contributors to sign off on web-based commits, or `false` to not require contributors to sign off on web-based commits.
// returns a *bool when successful
func (m *ItemItemOwnerPatchRequestBody) GetWebCommitSignoffRequired()(*bool) {
    return m.web_commit_signoff_required
}
// Serialize serializes information the current object
func (m *ItemItemOwnerPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allow_auto_merge", m.GetAllowAutoMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_forking", m.GetAllowForking())
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
        err := writer.WriteBoolValue("allow_update_branch", m.GetAllowUpdateBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("archived", m.GetArchived())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("default_branch", m.GetDefaultBranch())
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
        err := writer.WriteObjectValue("security_and_analysis", m.GetSecurityAndAnalysis())
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
        err := writer.WriteBoolValue("web_commit_signoff_required", m.GetWebCommitSignoffRequired())
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
func (m *ItemItemOwnerPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowAutoMerge sets the allow_auto_merge property value. Either `true` to allow auto-merge on pull requests, or `false` to disallow auto-merge.
func (m *ItemItemOwnerPatchRequestBody) SetAllowAutoMerge(value *bool)() {
    m.allow_auto_merge = value
}
// SetAllowForking sets the allow_forking property value. Either `true` to allow private forks, or `false` to prevent private forks.
func (m *ItemItemOwnerPatchRequestBody) SetAllowForking(value *bool)() {
    m.allow_forking = value
}
// SetAllowMergeCommit sets the allow_merge_commit property value. Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits.
func (m *ItemItemOwnerPatchRequestBody) SetAllowMergeCommit(value *bool)() {
    m.allow_merge_commit = value
}
// SetAllowRebaseMerge sets the allow_rebase_merge property value. Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging.
func (m *ItemItemOwnerPatchRequestBody) SetAllowRebaseMerge(value *bool)() {
    m.allow_rebase_merge = value
}
// SetAllowSquashMerge sets the allow_squash_merge property value. Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging.
func (m *ItemItemOwnerPatchRequestBody) SetAllowSquashMerge(value *bool)() {
    m.allow_squash_merge = value
}
// SetAllowUpdateBranch sets the allow_update_branch property value. Either `true` to always allow a pull request head branch that is behind its base branch to be updated even if it is not required to be up to date before merging, or false otherwise.
func (m *ItemItemOwnerPatchRequestBody) SetAllowUpdateBranch(value *bool)() {
    m.allow_update_branch = value
}
// SetArchived sets the archived property value. Whether to archive this repository. `false` will unarchive a previously archived repository.
func (m *ItemItemOwnerPatchRequestBody) SetArchived(value *bool)() {
    m.archived = value
}
// SetDefaultBranch sets the default_branch property value. Updates the default branch for this repository.
func (m *ItemItemOwnerPatchRequestBody) SetDefaultBranch(value *string)() {
    m.default_branch = value
}
// SetDeleteBranchOnMerge sets the delete_branch_on_merge property value. Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion.
func (m *ItemItemOwnerPatchRequestBody) SetDeleteBranchOnMerge(value *bool)() {
    m.delete_branch_on_merge = value
}
// SetDescription sets the description property value. A short description of the repository.
func (m *ItemItemOwnerPatchRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetHasIssues sets the has_issues property value. Either `true` to enable issues for this repository or `false` to disable them.
func (m *ItemItemOwnerPatchRequestBody) SetHasIssues(value *bool)() {
    m.has_issues = value
}
// SetHasProjects sets the has_projects property value. Either `true` to enable projects for this repository or `false` to disable them. **Note:** If you're creating a repository in an organization that has disabled repository projects, the default is `false`, and if you pass `true`, the API returns an error.
func (m *ItemItemOwnerPatchRequestBody) SetHasProjects(value *bool)() {
    m.has_projects = value
}
// SetHasWiki sets the has_wiki property value. Either `true` to enable the wiki for this repository or `false` to disable it.
func (m *ItemItemOwnerPatchRequestBody) SetHasWiki(value *bool)() {
    m.has_wiki = value
}
// SetHomepage sets the homepage property value. A URL with more information about the repository.
func (m *ItemItemOwnerPatchRequestBody) SetHomepage(value *string)() {
    m.homepage = value
}
// SetIsTemplate sets the is_template property value. Either `true` to make this repo available as a template repository or `false` to prevent it.
func (m *ItemItemOwnerPatchRequestBody) SetIsTemplate(value *bool)() {
    m.is_template = value
}
// SetName sets the name property value. The name of the repository.
func (m *ItemItemOwnerPatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetPrivate sets the private property value. Either `true` to make the repository private or `false` to make it public. Default: `false`.  **Note**: You will get a `422` error if the organization restricts [changing repository visibility](https://docs.github.com/articles/repository-permission-levels-for-an-organization#changing-the-visibility-of-repositories) to organization owners and a non-owner tries to change the value of private.
func (m *ItemItemOwnerPatchRequestBody) SetPrivate(value *bool)() {
    m.private = value
}
// SetSecurityAndAnalysis sets the security_and_analysis property value. Specify which security and analysis features to enable or disable for the repository.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."For example, to enable GitHub Advanced Security, use this data in the body of the `PATCH` request:`{ "security_and_analysis": {"advanced_security": { "status": "enabled" } } }`.You can check which security and analysis features are currently enabled by using a `GET /repos/{owner}/{repo}` request.
func (m *ItemItemOwnerPatchRequestBody) SetSecurityAndAnalysis(value ItemItemOwnerPatchRequestBody_security_and_analysisable)() {
    m.security_and_analysis = value
}
// SetUseSquashPrTitleAsDefault sets the use_squash_pr_title_as_default property value. Either `true` to allow squash-merge commits to use pull request title, or `false` to use commit message. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
// Deprecated: 
func (m *ItemItemOwnerPatchRequestBody) SetUseSquashPrTitleAsDefault(value *bool)() {
    m.use_squash_pr_title_as_default = value
}
// SetWebCommitSignoffRequired sets the web_commit_signoff_required property value. Either `true` to require contributors to sign off on web-based commits, or `false` to not require contributors to sign off on web-based commits.
func (m *ItemItemOwnerPatchRequestBody) SetWebCommitSignoffRequired(value *bool)() {
    m.web_commit_signoff_required = value
}
type ItemItemOwnerPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAutoMerge()(*bool)
    GetAllowForking()(*bool)
    GetAllowMergeCommit()(*bool)
    GetAllowRebaseMerge()(*bool)
    GetAllowSquashMerge()(*bool)
    GetAllowUpdateBranch()(*bool)
    GetArchived()(*bool)
    GetDefaultBranch()(*string)
    GetDeleteBranchOnMerge()(*bool)
    GetDescription()(*string)
    GetHasIssues()(*bool)
    GetHasProjects()(*bool)
    GetHasWiki()(*bool)
    GetHomepage()(*string)
    GetIsTemplate()(*bool)
    GetName()(*string)
    GetPrivate()(*bool)
    GetSecurityAndAnalysis()(ItemItemOwnerPatchRequestBody_security_and_analysisable)
    GetUseSquashPrTitleAsDefault()(*bool)
    GetWebCommitSignoffRequired()(*bool)
    SetAllowAutoMerge(value *bool)()
    SetAllowForking(value *bool)()
    SetAllowMergeCommit(value *bool)()
    SetAllowRebaseMerge(value *bool)()
    SetAllowSquashMerge(value *bool)()
    SetAllowUpdateBranch(value *bool)()
    SetArchived(value *bool)()
    SetDefaultBranch(value *string)()
    SetDeleteBranchOnMerge(value *bool)()
    SetDescription(value *string)()
    SetHasIssues(value *bool)()
    SetHasProjects(value *bool)()
    SetHasWiki(value *bool)()
    SetHomepage(value *string)()
    SetIsTemplate(value *bool)()
    SetName(value *string)()
    SetPrivate(value *bool)()
    SetSecurityAndAnalysis(value ItemItemOwnerPatchRequestBody_security_and_analysisable)()
    SetUseSquashPrTitleAsDefault(value *bool)()
    SetWebCommitSignoffRequired(value *bool)()
}
