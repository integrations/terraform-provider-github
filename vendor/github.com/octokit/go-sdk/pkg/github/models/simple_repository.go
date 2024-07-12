package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimpleRepository a GitHub repository.
type SimpleRepository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A template for the API URL to download the repository as an archive.
    archive_url *string
    // A template for the API URL to list the available assignees for issues in the repository.
    assignees_url *string
    // A template for the API URL to create or retrieve a raw Git blob in the repository.
    blobs_url *string
    // A template for the API URL to get information about branches in the repository.
    branches_url *string
    // A template for the API URL to get information about collaborators of the repository.
    collaborators_url *string
    // A template for the API URL to get information about comments on the repository.
    comments_url *string
    // A template for the API URL to get information about commits on the repository.
    commits_url *string
    // A template for the API URL to compare two commits or refs.
    compare_url *string
    // A template for the API URL to get the contents of the repository.
    contents_url *string
    // A template for the API URL to list the contributors to the repository.
    contributors_url *string
    // The API URL to list the deployments of the repository.
    deployments_url *string
    // The repository description.
    description *string
    // The API URL to list the downloads on the repository.
    downloads_url *string
    // The API URL to list the events of the repository.
    events_url *string
    // Whether the repository is a fork.
    fork *bool
    // The API URL to list the forks of the repository.
    forks_url *string
    // The full, globally unique, name of the repository.
    full_name *string
    // A template for the API URL to get information about Git commits of the repository.
    git_commits_url *string
    // A template for the API URL to get information about Git refs of the repository.
    git_refs_url *string
    // A template for the API URL to get information about Git tags of the repository.
    git_tags_url *string
    // The API URL to list the hooks on the repository.
    hooks_url *string
    // The URL to view the repository on GitHub.com.
    html_url *string
    // A unique identifier of the repository.
    id *int64
    // A template for the API URL to get information about issue comments on the repository.
    issue_comment_url *string
    // A template for the API URL to get information about issue events on the repository.
    issue_events_url *string
    // A template for the API URL to get information about issues on the repository.
    issues_url *string
    // A template for the API URL to get information about deploy keys on the repository.
    keys_url *string
    // A template for the API URL to get information about labels of the repository.
    labels_url *string
    // The API URL to get information about the languages of the repository.
    languages_url *string
    // The API URL to merge branches in the repository.
    merges_url *string
    // A template for the API URL to get information about milestones of the repository.
    milestones_url *string
    // The name of the repository.
    name *string
    // The GraphQL identifier of the repository.
    node_id *string
    // A template for the API URL to get information about notifications on the repository.
    notifications_url *string
    // A GitHub user.
    owner SimpleUserable
    // Whether the repository is private.
    private *bool
    // A template for the API URL to get information about pull requests on the repository.
    pulls_url *string
    // A template for the API URL to get information about releases on the repository.
    releases_url *string
    // The API URL to list the stargazers on the repository.
    stargazers_url *string
    // A template for the API URL to get information about statuses of a commit.
    statuses_url *string
    // The API URL to list the subscribers on the repository.
    subscribers_url *string
    // The API URL to subscribe to notifications for this repository.
    subscription_url *string
    // The API URL to get information about tags on the repository.
    tags_url *string
    // The API URL to list the teams on the repository.
    teams_url *string
    // A template for the API URL to create or retrieve a raw Git tree of the repository.
    trees_url *string
    // The URL to get more information about the repository from the GitHub API.
    url *string
}
// NewSimpleRepository instantiates a new SimpleRepository and sets the default values.
func NewSimpleRepository()(*SimpleRepository) {
    m := &SimpleRepository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSimpleRepositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSimpleRepositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimpleRepository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SimpleRepository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchiveUrl gets the archive_url property value. A template for the API URL to download the repository as an archive.
// returns a *string when successful
func (m *SimpleRepository) GetArchiveUrl()(*string) {
    return m.archive_url
}
// GetAssigneesUrl gets the assignees_url property value. A template for the API URL to list the available assignees for issues in the repository.
// returns a *string when successful
func (m *SimpleRepository) GetAssigneesUrl()(*string) {
    return m.assignees_url
}
// GetBlobsUrl gets the blobs_url property value. A template for the API URL to create or retrieve a raw Git blob in the repository.
// returns a *string when successful
func (m *SimpleRepository) GetBlobsUrl()(*string) {
    return m.blobs_url
}
// GetBranchesUrl gets the branches_url property value. A template for the API URL to get information about branches in the repository.
// returns a *string when successful
func (m *SimpleRepository) GetBranchesUrl()(*string) {
    return m.branches_url
}
// GetCollaboratorsUrl gets the collaborators_url property value. A template for the API URL to get information about collaborators of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetCollaboratorsUrl()(*string) {
    return m.collaborators_url
}
// GetCommentsUrl gets the comments_url property value. A template for the API URL to get information about comments on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCommitsUrl gets the commits_url property value. A template for the API URL to get information about commits on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetCommitsUrl()(*string) {
    return m.commits_url
}
// GetCompareUrl gets the compare_url property value. A template for the API URL to compare two commits or refs.
// returns a *string when successful
func (m *SimpleRepository) GetCompareUrl()(*string) {
    return m.compare_url
}
// GetContentsUrl gets the contents_url property value. A template for the API URL to get the contents of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetContentsUrl()(*string) {
    return m.contents_url
}
// GetContributorsUrl gets the contributors_url property value. A template for the API URL to list the contributors to the repository.
// returns a *string when successful
func (m *SimpleRepository) GetContributorsUrl()(*string) {
    return m.contributors_url
}
// GetDeploymentsUrl gets the deployments_url property value. The API URL to list the deployments of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetDeploymentsUrl()(*string) {
    return m.deployments_url
}
// GetDescription gets the description property value. The repository description.
// returns a *string when successful
func (m *SimpleRepository) GetDescription()(*string) {
    return m.description
}
// GetDownloadsUrl gets the downloads_url property value. The API URL to list the downloads on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetDownloadsUrl()(*string) {
    return m.downloads_url
}
// GetEventsUrl gets the events_url property value. The API URL to list the events of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SimpleRepository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archive_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchiveUrl(val)
        }
        return nil
    }
    res["assignees_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigneesUrl(val)
        }
        return nil
    }
    res["blobs_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlobsUrl(val)
        }
        return nil
    }
    res["branches_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranchesUrl(val)
        }
        return nil
    }
    res["collaborators_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollaboratorsUrl(val)
        }
        return nil
    }
    res["comments_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentsUrl(val)
        }
        return nil
    }
    res["commits_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitsUrl(val)
        }
        return nil
    }
    res["compare_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompareUrl(val)
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
    res["contributors_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributorsUrl(val)
        }
        return nil
    }
    res["deployments_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentsUrl(val)
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
    res["downloads_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadsUrl(val)
        }
        return nil
    }
    res["events_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventsUrl(val)
        }
        return nil
    }
    res["fork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFork(val)
        }
        return nil
    }
    res["forks_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForksUrl(val)
        }
        return nil
    }
    res["full_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullName(val)
        }
        return nil
    }
    res["git_commits_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitCommitsUrl(val)
        }
        return nil
    }
    res["git_refs_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitRefsUrl(val)
        }
        return nil
    }
    res["git_tags_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitTagsUrl(val)
        }
        return nil
    }
    res["hooks_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHooksUrl(val)
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["issue_comment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueCommentUrl(val)
        }
        return nil
    }
    res["issue_events_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueEventsUrl(val)
        }
        return nil
    }
    res["issues_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuesUrl(val)
        }
        return nil
    }
    res["keys_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeysUrl(val)
        }
        return nil
    }
    res["labels_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelsUrl(val)
        }
        return nil
    }
    res["languages_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguagesUrl(val)
        }
        return nil
    }
    res["merges_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergesUrl(val)
        }
        return nil
    }
    res["milestones_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestonesUrl(val)
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
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["notifications_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationsUrl(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(SimpleUserable))
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
    res["pulls_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullsUrl(val)
        }
        return nil
    }
    res["releases_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleasesUrl(val)
        }
        return nil
    }
    res["stargazers_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStargazersUrl(val)
        }
        return nil
    }
    res["statuses_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusesUrl(val)
        }
        return nil
    }
    res["subscribers_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscribersUrl(val)
        }
        return nil
    }
    res["subscription_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionUrl(val)
        }
        return nil
    }
    res["tags_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTagsUrl(val)
        }
        return nil
    }
    res["teams_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsUrl(val)
        }
        return nil
    }
    res["trees_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTreesUrl(val)
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
// GetFork gets the fork property value. Whether the repository is a fork.
// returns a *bool when successful
func (m *SimpleRepository) GetFork()(*bool) {
    return m.fork
}
// GetForksUrl gets the forks_url property value. The API URL to list the forks of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetForksUrl()(*string) {
    return m.forks_url
}
// GetFullName gets the full_name property value. The full, globally unique, name of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetFullName()(*string) {
    return m.full_name
}
// GetGitCommitsUrl gets the git_commits_url property value. A template for the API URL to get information about Git commits of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetGitCommitsUrl()(*string) {
    return m.git_commits_url
}
// GetGitRefsUrl gets the git_refs_url property value. A template for the API URL to get information about Git refs of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetGitRefsUrl()(*string) {
    return m.git_refs_url
}
// GetGitTagsUrl gets the git_tags_url property value. A template for the API URL to get information about Git tags of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetGitTagsUrl()(*string) {
    return m.git_tags_url
}
// GetHooksUrl gets the hooks_url property value. The API URL to list the hooks on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetHooksUrl()(*string) {
    return m.hooks_url
}
// GetHtmlUrl gets the html_url property value. The URL to view the repository on GitHub.com.
// returns a *string when successful
func (m *SimpleRepository) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. A unique identifier of the repository.
// returns a *int64 when successful
func (m *SimpleRepository) GetId()(*int64) {
    return m.id
}
// GetIssueCommentUrl gets the issue_comment_url property value. A template for the API URL to get information about issue comments on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetIssueCommentUrl()(*string) {
    return m.issue_comment_url
}
// GetIssueEventsUrl gets the issue_events_url property value. A template for the API URL to get information about issue events on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetIssueEventsUrl()(*string) {
    return m.issue_events_url
}
// GetIssuesUrl gets the issues_url property value. A template for the API URL to get information about issues on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetIssuesUrl()(*string) {
    return m.issues_url
}
// GetKeysUrl gets the keys_url property value. A template for the API URL to get information about deploy keys on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetKeysUrl()(*string) {
    return m.keys_url
}
// GetLabelsUrl gets the labels_url property value. A template for the API URL to get information about labels of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetLabelsUrl()(*string) {
    return m.labels_url
}
// GetLanguagesUrl gets the languages_url property value. The API URL to get information about the languages of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetLanguagesUrl()(*string) {
    return m.languages_url
}
// GetMergesUrl gets the merges_url property value. The API URL to merge branches in the repository.
// returns a *string when successful
func (m *SimpleRepository) GetMergesUrl()(*string) {
    return m.merges_url
}
// GetMilestonesUrl gets the milestones_url property value. A template for the API URL to get information about milestones of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetMilestonesUrl()(*string) {
    return m.milestones_url
}
// GetName gets the name property value. The name of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The GraphQL identifier of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetNodeId()(*string) {
    return m.node_id
}
// GetNotificationsUrl gets the notifications_url property value. A template for the API URL to get information about notifications on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetNotificationsUrl()(*string) {
    return m.notifications_url
}
// GetOwner gets the owner property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *SimpleRepository) GetOwner()(SimpleUserable) {
    return m.owner
}
// GetPrivate gets the private property value. Whether the repository is private.
// returns a *bool when successful
func (m *SimpleRepository) GetPrivate()(*bool) {
    return m.private
}
// GetPullsUrl gets the pulls_url property value. A template for the API URL to get information about pull requests on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetPullsUrl()(*string) {
    return m.pulls_url
}
// GetReleasesUrl gets the releases_url property value. A template for the API URL to get information about releases on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetReleasesUrl()(*string) {
    return m.releases_url
}
// GetStargazersUrl gets the stargazers_url property value. The API URL to list the stargazers on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetStargazersUrl()(*string) {
    return m.stargazers_url
}
// GetStatusesUrl gets the statuses_url property value. A template for the API URL to get information about statuses of a commit.
// returns a *string when successful
func (m *SimpleRepository) GetStatusesUrl()(*string) {
    return m.statuses_url
}
// GetSubscribersUrl gets the subscribers_url property value. The API URL to list the subscribers on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetSubscribersUrl()(*string) {
    return m.subscribers_url
}
// GetSubscriptionUrl gets the subscription_url property value. The API URL to subscribe to notifications for this repository.
// returns a *string when successful
func (m *SimpleRepository) GetSubscriptionUrl()(*string) {
    return m.subscription_url
}
// GetTagsUrl gets the tags_url property value. The API URL to get information about tags on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetTagsUrl()(*string) {
    return m.tags_url
}
// GetTeamsUrl gets the teams_url property value. The API URL to list the teams on the repository.
// returns a *string when successful
func (m *SimpleRepository) GetTeamsUrl()(*string) {
    return m.teams_url
}
// GetTreesUrl gets the trees_url property value. A template for the API URL to create or retrieve a raw Git tree of the repository.
// returns a *string when successful
func (m *SimpleRepository) GetTreesUrl()(*string) {
    return m.trees_url
}
// GetUrl gets the url property value. The URL to get more information about the repository from the GitHub API.
// returns a *string when successful
func (m *SimpleRepository) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *SimpleRepository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("archive_url", m.GetArchiveUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assignees_url", m.GetAssigneesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blobs_url", m.GetBlobsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("branches_url", m.GetBranchesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("collaborators_url", m.GetCollaboratorsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comments_url", m.GetCommentsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commits_url", m.GetCommitsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("compare_url", m.GetCompareUrl())
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
        err := writer.WriteStringValue("contributors_url", m.GetContributorsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deployments_url", m.GetDeploymentsUrl())
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
        err := writer.WriteStringValue("downloads_url", m.GetDownloadsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("events_url", m.GetEventsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("fork", m.GetFork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("forks_url", m.GetForksUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("full_name", m.GetFullName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_commits_url", m.GetGitCommitsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_refs_url", m.GetGitRefsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_tags_url", m.GetGitTagsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hooks_url", m.GetHooksUrl())
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
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issues_url", m.GetIssuesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issue_comment_url", m.GetIssueCommentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issue_events_url", m.GetIssueEventsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("keys_url", m.GetKeysUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labels_url", m.GetLabelsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languages_url", m.GetLanguagesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("merges_url", m.GetMergesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("milestones_url", m.GetMilestonesUrl())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notifications_url", m.GetNotificationsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("owner", m.GetOwner())
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
        err := writer.WriteStringValue("pulls_url", m.GetPullsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("releases_url", m.GetReleasesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stargazers_url", m.GetStargazersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("statuses_url", m.GetStatusesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscribers_url", m.GetSubscribersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscription_url", m.GetSubscriptionUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tags_url", m.GetTagsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teams_url", m.GetTeamsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trees_url", m.GetTreesUrl())
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
func (m *SimpleRepository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchiveUrl sets the archive_url property value. A template for the API URL to download the repository as an archive.
func (m *SimpleRepository) SetArchiveUrl(value *string)() {
    m.archive_url = value
}
// SetAssigneesUrl sets the assignees_url property value. A template for the API URL to list the available assignees for issues in the repository.
func (m *SimpleRepository) SetAssigneesUrl(value *string)() {
    m.assignees_url = value
}
// SetBlobsUrl sets the blobs_url property value. A template for the API URL to create or retrieve a raw Git blob in the repository.
func (m *SimpleRepository) SetBlobsUrl(value *string)() {
    m.blobs_url = value
}
// SetBranchesUrl sets the branches_url property value. A template for the API URL to get information about branches in the repository.
func (m *SimpleRepository) SetBranchesUrl(value *string)() {
    m.branches_url = value
}
// SetCollaboratorsUrl sets the collaborators_url property value. A template for the API URL to get information about collaborators of the repository.
func (m *SimpleRepository) SetCollaboratorsUrl(value *string)() {
    m.collaborators_url = value
}
// SetCommentsUrl sets the comments_url property value. A template for the API URL to get information about comments on the repository.
func (m *SimpleRepository) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCommitsUrl sets the commits_url property value. A template for the API URL to get information about commits on the repository.
func (m *SimpleRepository) SetCommitsUrl(value *string)() {
    m.commits_url = value
}
// SetCompareUrl sets the compare_url property value. A template for the API URL to compare two commits or refs.
func (m *SimpleRepository) SetCompareUrl(value *string)() {
    m.compare_url = value
}
// SetContentsUrl sets the contents_url property value. A template for the API URL to get the contents of the repository.
func (m *SimpleRepository) SetContentsUrl(value *string)() {
    m.contents_url = value
}
// SetContributorsUrl sets the contributors_url property value. A template for the API URL to list the contributors to the repository.
func (m *SimpleRepository) SetContributorsUrl(value *string)() {
    m.contributors_url = value
}
// SetDeploymentsUrl sets the deployments_url property value. The API URL to list the deployments of the repository.
func (m *SimpleRepository) SetDeploymentsUrl(value *string)() {
    m.deployments_url = value
}
// SetDescription sets the description property value. The repository description.
func (m *SimpleRepository) SetDescription(value *string)() {
    m.description = value
}
// SetDownloadsUrl sets the downloads_url property value. The API URL to list the downloads on the repository.
func (m *SimpleRepository) SetDownloadsUrl(value *string)() {
    m.downloads_url = value
}
// SetEventsUrl sets the events_url property value. The API URL to list the events of the repository.
func (m *SimpleRepository) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetFork sets the fork property value. Whether the repository is a fork.
func (m *SimpleRepository) SetFork(value *bool)() {
    m.fork = value
}
// SetForksUrl sets the forks_url property value. The API URL to list the forks of the repository.
func (m *SimpleRepository) SetForksUrl(value *string)() {
    m.forks_url = value
}
// SetFullName sets the full_name property value. The full, globally unique, name of the repository.
func (m *SimpleRepository) SetFullName(value *string)() {
    m.full_name = value
}
// SetGitCommitsUrl sets the git_commits_url property value. A template for the API URL to get information about Git commits of the repository.
func (m *SimpleRepository) SetGitCommitsUrl(value *string)() {
    m.git_commits_url = value
}
// SetGitRefsUrl sets the git_refs_url property value. A template for the API URL to get information about Git refs of the repository.
func (m *SimpleRepository) SetGitRefsUrl(value *string)() {
    m.git_refs_url = value
}
// SetGitTagsUrl sets the git_tags_url property value. A template for the API URL to get information about Git tags of the repository.
func (m *SimpleRepository) SetGitTagsUrl(value *string)() {
    m.git_tags_url = value
}
// SetHooksUrl sets the hooks_url property value. The API URL to list the hooks on the repository.
func (m *SimpleRepository) SetHooksUrl(value *string)() {
    m.hooks_url = value
}
// SetHtmlUrl sets the html_url property value. The URL to view the repository on GitHub.com.
func (m *SimpleRepository) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. A unique identifier of the repository.
func (m *SimpleRepository) SetId(value *int64)() {
    m.id = value
}
// SetIssueCommentUrl sets the issue_comment_url property value. A template for the API URL to get information about issue comments on the repository.
func (m *SimpleRepository) SetIssueCommentUrl(value *string)() {
    m.issue_comment_url = value
}
// SetIssueEventsUrl sets the issue_events_url property value. A template for the API URL to get information about issue events on the repository.
func (m *SimpleRepository) SetIssueEventsUrl(value *string)() {
    m.issue_events_url = value
}
// SetIssuesUrl sets the issues_url property value. A template for the API URL to get information about issues on the repository.
func (m *SimpleRepository) SetIssuesUrl(value *string)() {
    m.issues_url = value
}
// SetKeysUrl sets the keys_url property value. A template for the API URL to get information about deploy keys on the repository.
func (m *SimpleRepository) SetKeysUrl(value *string)() {
    m.keys_url = value
}
// SetLabelsUrl sets the labels_url property value. A template for the API URL to get information about labels of the repository.
func (m *SimpleRepository) SetLabelsUrl(value *string)() {
    m.labels_url = value
}
// SetLanguagesUrl sets the languages_url property value. The API URL to get information about the languages of the repository.
func (m *SimpleRepository) SetLanguagesUrl(value *string)() {
    m.languages_url = value
}
// SetMergesUrl sets the merges_url property value. The API URL to merge branches in the repository.
func (m *SimpleRepository) SetMergesUrl(value *string)() {
    m.merges_url = value
}
// SetMilestonesUrl sets the milestones_url property value. A template for the API URL to get information about milestones of the repository.
func (m *SimpleRepository) SetMilestonesUrl(value *string)() {
    m.milestones_url = value
}
// SetName sets the name property value. The name of the repository.
func (m *SimpleRepository) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The GraphQL identifier of the repository.
func (m *SimpleRepository) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNotificationsUrl sets the notifications_url property value. A template for the API URL to get information about notifications on the repository.
func (m *SimpleRepository) SetNotificationsUrl(value *string)() {
    m.notifications_url = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *SimpleRepository) SetOwner(value SimpleUserable)() {
    m.owner = value
}
// SetPrivate sets the private property value. Whether the repository is private.
func (m *SimpleRepository) SetPrivate(value *bool)() {
    m.private = value
}
// SetPullsUrl sets the pulls_url property value. A template for the API URL to get information about pull requests on the repository.
func (m *SimpleRepository) SetPullsUrl(value *string)() {
    m.pulls_url = value
}
// SetReleasesUrl sets the releases_url property value. A template for the API URL to get information about releases on the repository.
func (m *SimpleRepository) SetReleasesUrl(value *string)() {
    m.releases_url = value
}
// SetStargazersUrl sets the stargazers_url property value. The API URL to list the stargazers on the repository.
func (m *SimpleRepository) SetStargazersUrl(value *string)() {
    m.stargazers_url = value
}
// SetStatusesUrl sets the statuses_url property value. A template for the API URL to get information about statuses of a commit.
func (m *SimpleRepository) SetStatusesUrl(value *string)() {
    m.statuses_url = value
}
// SetSubscribersUrl sets the subscribers_url property value. The API URL to list the subscribers on the repository.
func (m *SimpleRepository) SetSubscribersUrl(value *string)() {
    m.subscribers_url = value
}
// SetSubscriptionUrl sets the subscription_url property value. The API URL to subscribe to notifications for this repository.
func (m *SimpleRepository) SetSubscriptionUrl(value *string)() {
    m.subscription_url = value
}
// SetTagsUrl sets the tags_url property value. The API URL to get information about tags on the repository.
func (m *SimpleRepository) SetTagsUrl(value *string)() {
    m.tags_url = value
}
// SetTeamsUrl sets the teams_url property value. The API URL to list the teams on the repository.
func (m *SimpleRepository) SetTeamsUrl(value *string)() {
    m.teams_url = value
}
// SetTreesUrl sets the trees_url property value. A template for the API URL to create or retrieve a raw Git tree of the repository.
func (m *SimpleRepository) SetTreesUrl(value *string)() {
    m.trees_url = value
}
// SetUrl sets the url property value. The URL to get more information about the repository from the GitHub API.
func (m *SimpleRepository) SetUrl(value *string)() {
    m.url = value
}
type SimpleRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchiveUrl()(*string)
    GetAssigneesUrl()(*string)
    GetBlobsUrl()(*string)
    GetBranchesUrl()(*string)
    GetCollaboratorsUrl()(*string)
    GetCommentsUrl()(*string)
    GetCommitsUrl()(*string)
    GetCompareUrl()(*string)
    GetContentsUrl()(*string)
    GetContributorsUrl()(*string)
    GetDeploymentsUrl()(*string)
    GetDescription()(*string)
    GetDownloadsUrl()(*string)
    GetEventsUrl()(*string)
    GetFork()(*bool)
    GetForksUrl()(*string)
    GetFullName()(*string)
    GetGitCommitsUrl()(*string)
    GetGitRefsUrl()(*string)
    GetGitTagsUrl()(*string)
    GetHooksUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetIssueCommentUrl()(*string)
    GetIssueEventsUrl()(*string)
    GetIssuesUrl()(*string)
    GetKeysUrl()(*string)
    GetLabelsUrl()(*string)
    GetLanguagesUrl()(*string)
    GetMergesUrl()(*string)
    GetMilestonesUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetNotificationsUrl()(*string)
    GetOwner()(SimpleUserable)
    GetPrivate()(*bool)
    GetPullsUrl()(*string)
    GetReleasesUrl()(*string)
    GetStargazersUrl()(*string)
    GetStatusesUrl()(*string)
    GetSubscribersUrl()(*string)
    GetSubscriptionUrl()(*string)
    GetTagsUrl()(*string)
    GetTeamsUrl()(*string)
    GetTreesUrl()(*string)
    GetUrl()(*string)
    SetArchiveUrl(value *string)()
    SetAssigneesUrl(value *string)()
    SetBlobsUrl(value *string)()
    SetBranchesUrl(value *string)()
    SetCollaboratorsUrl(value *string)()
    SetCommentsUrl(value *string)()
    SetCommitsUrl(value *string)()
    SetCompareUrl(value *string)()
    SetContentsUrl(value *string)()
    SetContributorsUrl(value *string)()
    SetDeploymentsUrl(value *string)()
    SetDescription(value *string)()
    SetDownloadsUrl(value *string)()
    SetEventsUrl(value *string)()
    SetFork(value *bool)()
    SetForksUrl(value *string)()
    SetFullName(value *string)()
    SetGitCommitsUrl(value *string)()
    SetGitRefsUrl(value *string)()
    SetGitTagsUrl(value *string)()
    SetHooksUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetIssueCommentUrl(value *string)()
    SetIssueEventsUrl(value *string)()
    SetIssuesUrl(value *string)()
    SetKeysUrl(value *string)()
    SetLabelsUrl(value *string)()
    SetLanguagesUrl(value *string)()
    SetMergesUrl(value *string)()
    SetMilestonesUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNotificationsUrl(value *string)()
    SetOwner(value SimpleUserable)()
    SetPrivate(value *bool)()
    SetPullsUrl(value *string)()
    SetReleasesUrl(value *string)()
    SetStargazersUrl(value *string)()
    SetStatusesUrl(value *string)()
    SetSubscribersUrl(value *string)()
    SetSubscriptionUrl(value *string)()
    SetTagsUrl(value *string)()
    SetTeamsUrl(value *string)()
    SetTreesUrl(value *string)()
    SetUrl(value *string)()
}
