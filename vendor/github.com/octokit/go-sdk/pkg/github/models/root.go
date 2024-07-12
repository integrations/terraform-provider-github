package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Root struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The authorizations_url property
    authorizations_url *string
    // The code_search_url property
    code_search_url *string
    // The commit_search_url property
    commit_search_url *string
    // The current_user_authorizations_html_url property
    current_user_authorizations_html_url *string
    // The current_user_repositories_url property
    current_user_repositories_url *string
    // The current_user_url property
    current_user_url *string
    // The emails_url property
    emails_url *string
    // The emojis_url property
    emojis_url *string
    // The events_url property
    events_url *string
    // The feeds_url property
    feeds_url *string
    // The followers_url property
    followers_url *string
    // The following_url property
    following_url *string
    // The gists_url property
    gists_url *string
    // The hub_url property
    // Deprecated: 
    hub_url *string
    // The issue_search_url property
    issue_search_url *string
    // The issues_url property
    issues_url *string
    // The keys_url property
    keys_url *string
    // The label_search_url property
    label_search_url *string
    // The notifications_url property
    notifications_url *string
    // The organization_repositories_url property
    organization_repositories_url *string
    // The organization_teams_url property
    organization_teams_url *string
    // The organization_url property
    organization_url *string
    // The public_gists_url property
    public_gists_url *string
    // The rate_limit_url property
    rate_limit_url *string
    // The repository_search_url property
    repository_search_url *string
    // The repository_url property
    repository_url *string
    // The starred_gists_url property
    starred_gists_url *string
    // The starred_url property
    starred_url *string
    // The topic_search_url property
    topic_search_url *string
    // The user_organizations_url property
    user_organizations_url *string
    // The user_repositories_url property
    user_repositories_url *string
    // The user_search_url property
    user_search_url *string
    // The user_url property
    user_url *string
}
// NewRoot instantiates a new Root and sets the default values.
func NewRoot()(*Root) {
    m := &Root{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoot(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Root) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorizationsUrl gets the authorizations_url property value. The authorizations_url property
// returns a *string when successful
func (m *Root) GetAuthorizationsUrl()(*string) {
    return m.authorizations_url
}
// GetCodeSearchUrl gets the code_search_url property value. The code_search_url property
// returns a *string when successful
func (m *Root) GetCodeSearchUrl()(*string) {
    return m.code_search_url
}
// GetCommitSearchUrl gets the commit_search_url property value. The commit_search_url property
// returns a *string when successful
func (m *Root) GetCommitSearchUrl()(*string) {
    return m.commit_search_url
}
// GetCurrentUserAuthorizationsHtmlUrl gets the current_user_authorizations_html_url property value. The current_user_authorizations_html_url property
// returns a *string when successful
func (m *Root) GetCurrentUserAuthorizationsHtmlUrl()(*string) {
    return m.current_user_authorizations_html_url
}
// GetCurrentUserRepositoriesUrl gets the current_user_repositories_url property value. The current_user_repositories_url property
// returns a *string when successful
func (m *Root) GetCurrentUserRepositoriesUrl()(*string) {
    return m.current_user_repositories_url
}
// GetCurrentUserUrl gets the current_user_url property value. The current_user_url property
// returns a *string when successful
func (m *Root) GetCurrentUserUrl()(*string) {
    return m.current_user_url
}
// GetEmailsUrl gets the emails_url property value. The emails_url property
// returns a *string when successful
func (m *Root) GetEmailsUrl()(*string) {
    return m.emails_url
}
// GetEmojisUrl gets the emojis_url property value. The emojis_url property
// returns a *string when successful
func (m *Root) GetEmojisUrl()(*string) {
    return m.emojis_url
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *Root) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFeedsUrl gets the feeds_url property value. The feeds_url property
// returns a *string when successful
func (m *Root) GetFeedsUrl()(*string) {
    return m.feeds_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Root) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authorizations_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationsUrl(val)
        }
        return nil
    }
    res["code_search_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeSearchUrl(val)
        }
        return nil
    }
    res["commit_search_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitSearchUrl(val)
        }
        return nil
    }
    res["current_user_authorizations_html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserAuthorizationsHtmlUrl(val)
        }
        return nil
    }
    res["current_user_repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserRepositoriesUrl(val)
        }
        return nil
    }
    res["current_user_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserUrl(val)
        }
        return nil
    }
    res["emails_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailsUrl(val)
        }
        return nil
    }
    res["emojis_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmojisUrl(val)
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
    res["feeds_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedsUrl(val)
        }
        return nil
    }
    res["followers_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFollowersUrl(val)
        }
        return nil
    }
    res["following_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFollowingUrl(val)
        }
        return nil
    }
    res["gists_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGistsUrl(val)
        }
        return nil
    }
    res["hub_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHubUrl(val)
        }
        return nil
    }
    res["issue_search_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueSearchUrl(val)
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
    res["label_search_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelSearchUrl(val)
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
    res["organization_repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationRepositoriesUrl(val)
        }
        return nil
    }
    res["organization_teams_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationTeamsUrl(val)
        }
        return nil
    }
    res["organization_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationUrl(val)
        }
        return nil
    }
    res["public_gists_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicGistsUrl(val)
        }
        return nil
    }
    res["rate_limit_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRateLimitUrl(val)
        }
        return nil
    }
    res["repository_search_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositorySearchUrl(val)
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
    res["starred_gists_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStarredGistsUrl(val)
        }
        return nil
    }
    res["starred_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStarredUrl(val)
        }
        return nil
    }
    res["topic_search_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicSearchUrl(val)
        }
        return nil
    }
    res["user_organizations_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserOrganizationsUrl(val)
        }
        return nil
    }
    res["user_repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRepositoriesUrl(val)
        }
        return nil
    }
    res["user_search_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSearchUrl(val)
        }
        return nil
    }
    res["user_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserUrl(val)
        }
        return nil
    }
    return res
}
// GetFollowersUrl gets the followers_url property value. The followers_url property
// returns a *string when successful
func (m *Root) GetFollowersUrl()(*string) {
    return m.followers_url
}
// GetFollowingUrl gets the following_url property value. The following_url property
// returns a *string when successful
func (m *Root) GetFollowingUrl()(*string) {
    return m.following_url
}
// GetGistsUrl gets the gists_url property value. The gists_url property
// returns a *string when successful
func (m *Root) GetGistsUrl()(*string) {
    return m.gists_url
}
// GetHubUrl gets the hub_url property value. The hub_url property
// Deprecated: 
// returns a *string when successful
func (m *Root) GetHubUrl()(*string) {
    return m.hub_url
}
// GetIssueSearchUrl gets the issue_search_url property value. The issue_search_url property
// returns a *string when successful
func (m *Root) GetIssueSearchUrl()(*string) {
    return m.issue_search_url
}
// GetIssuesUrl gets the issues_url property value. The issues_url property
// returns a *string when successful
func (m *Root) GetIssuesUrl()(*string) {
    return m.issues_url
}
// GetKeysUrl gets the keys_url property value. The keys_url property
// returns a *string when successful
func (m *Root) GetKeysUrl()(*string) {
    return m.keys_url
}
// GetLabelSearchUrl gets the label_search_url property value. The label_search_url property
// returns a *string when successful
func (m *Root) GetLabelSearchUrl()(*string) {
    return m.label_search_url
}
// GetNotificationsUrl gets the notifications_url property value. The notifications_url property
// returns a *string when successful
func (m *Root) GetNotificationsUrl()(*string) {
    return m.notifications_url
}
// GetOrganizationRepositoriesUrl gets the organization_repositories_url property value. The organization_repositories_url property
// returns a *string when successful
func (m *Root) GetOrganizationRepositoriesUrl()(*string) {
    return m.organization_repositories_url
}
// GetOrganizationTeamsUrl gets the organization_teams_url property value. The organization_teams_url property
// returns a *string when successful
func (m *Root) GetOrganizationTeamsUrl()(*string) {
    return m.organization_teams_url
}
// GetOrganizationUrl gets the organization_url property value. The organization_url property
// returns a *string when successful
func (m *Root) GetOrganizationUrl()(*string) {
    return m.organization_url
}
// GetPublicGistsUrl gets the public_gists_url property value. The public_gists_url property
// returns a *string when successful
func (m *Root) GetPublicGistsUrl()(*string) {
    return m.public_gists_url
}
// GetRateLimitUrl gets the rate_limit_url property value. The rate_limit_url property
// returns a *string when successful
func (m *Root) GetRateLimitUrl()(*string) {
    return m.rate_limit_url
}
// GetRepositorySearchUrl gets the repository_search_url property value. The repository_search_url property
// returns a *string when successful
func (m *Root) GetRepositorySearchUrl()(*string) {
    return m.repository_search_url
}
// GetRepositoryUrl gets the repository_url property value. The repository_url property
// returns a *string when successful
func (m *Root) GetRepositoryUrl()(*string) {
    return m.repository_url
}
// GetStarredGistsUrl gets the starred_gists_url property value. The starred_gists_url property
// returns a *string when successful
func (m *Root) GetStarredGistsUrl()(*string) {
    return m.starred_gists_url
}
// GetStarredUrl gets the starred_url property value. The starred_url property
// returns a *string when successful
func (m *Root) GetStarredUrl()(*string) {
    return m.starred_url
}
// GetTopicSearchUrl gets the topic_search_url property value. The topic_search_url property
// returns a *string when successful
func (m *Root) GetTopicSearchUrl()(*string) {
    return m.topic_search_url
}
// GetUserOrganizationsUrl gets the user_organizations_url property value. The user_organizations_url property
// returns a *string when successful
func (m *Root) GetUserOrganizationsUrl()(*string) {
    return m.user_organizations_url
}
// GetUserRepositoriesUrl gets the user_repositories_url property value. The user_repositories_url property
// returns a *string when successful
func (m *Root) GetUserRepositoriesUrl()(*string) {
    return m.user_repositories_url
}
// GetUserSearchUrl gets the user_search_url property value. The user_search_url property
// returns a *string when successful
func (m *Root) GetUserSearchUrl()(*string) {
    return m.user_search_url
}
// GetUserUrl gets the user_url property value. The user_url property
// returns a *string when successful
func (m *Root) GetUserUrl()(*string) {
    return m.user_url
}
// Serialize serializes information the current object
func (m *Root) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authorizations_url", m.GetAuthorizationsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("code_search_url", m.GetCodeSearchUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_search_url", m.GetCommitSearchUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("current_user_authorizations_html_url", m.GetCurrentUserAuthorizationsHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("current_user_repositories_url", m.GetCurrentUserRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("current_user_url", m.GetCurrentUserUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("emails_url", m.GetEmailsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("emojis_url", m.GetEmojisUrl())
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
        err := writer.WriteStringValue("feeds_url", m.GetFeedsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("followers_url", m.GetFollowersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("following_url", m.GetFollowingUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("gists_url", m.GetGistsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hub_url", m.GetHubUrl())
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
        err := writer.WriteStringValue("issue_search_url", m.GetIssueSearchUrl())
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
        err := writer.WriteStringValue("label_search_url", m.GetLabelSearchUrl())
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
        err := writer.WriteStringValue("organization_repositories_url", m.GetOrganizationRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organization_teams_url", m.GetOrganizationTeamsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organization_url", m.GetOrganizationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("public_gists_url", m.GetPublicGistsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rate_limit_url", m.GetRateLimitUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repository_search_url", m.GetRepositorySearchUrl())
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
    {
        err := writer.WriteStringValue("starred_gists_url", m.GetStarredGistsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("starred_url", m.GetStarredUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("topic_search_url", m.GetTopicSearchUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user_organizations_url", m.GetUserOrganizationsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user_repositories_url", m.GetUserRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user_search_url", m.GetUserSearchUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user_url", m.GetUserUrl())
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
func (m *Root) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorizationsUrl sets the authorizations_url property value. The authorizations_url property
func (m *Root) SetAuthorizationsUrl(value *string)() {
    m.authorizations_url = value
}
// SetCodeSearchUrl sets the code_search_url property value. The code_search_url property
func (m *Root) SetCodeSearchUrl(value *string)() {
    m.code_search_url = value
}
// SetCommitSearchUrl sets the commit_search_url property value. The commit_search_url property
func (m *Root) SetCommitSearchUrl(value *string)() {
    m.commit_search_url = value
}
// SetCurrentUserAuthorizationsHtmlUrl sets the current_user_authorizations_html_url property value. The current_user_authorizations_html_url property
func (m *Root) SetCurrentUserAuthorizationsHtmlUrl(value *string)() {
    m.current_user_authorizations_html_url = value
}
// SetCurrentUserRepositoriesUrl sets the current_user_repositories_url property value. The current_user_repositories_url property
func (m *Root) SetCurrentUserRepositoriesUrl(value *string)() {
    m.current_user_repositories_url = value
}
// SetCurrentUserUrl sets the current_user_url property value. The current_user_url property
func (m *Root) SetCurrentUserUrl(value *string)() {
    m.current_user_url = value
}
// SetEmailsUrl sets the emails_url property value. The emails_url property
func (m *Root) SetEmailsUrl(value *string)() {
    m.emails_url = value
}
// SetEmojisUrl sets the emojis_url property value. The emojis_url property
func (m *Root) SetEmojisUrl(value *string)() {
    m.emojis_url = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *Root) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetFeedsUrl sets the feeds_url property value. The feeds_url property
func (m *Root) SetFeedsUrl(value *string)() {
    m.feeds_url = value
}
// SetFollowersUrl sets the followers_url property value. The followers_url property
func (m *Root) SetFollowersUrl(value *string)() {
    m.followers_url = value
}
// SetFollowingUrl sets the following_url property value. The following_url property
func (m *Root) SetFollowingUrl(value *string)() {
    m.following_url = value
}
// SetGistsUrl sets the gists_url property value. The gists_url property
func (m *Root) SetGistsUrl(value *string)() {
    m.gists_url = value
}
// SetHubUrl sets the hub_url property value. The hub_url property
// Deprecated: 
func (m *Root) SetHubUrl(value *string)() {
    m.hub_url = value
}
// SetIssueSearchUrl sets the issue_search_url property value. The issue_search_url property
func (m *Root) SetIssueSearchUrl(value *string)() {
    m.issue_search_url = value
}
// SetIssuesUrl sets the issues_url property value. The issues_url property
func (m *Root) SetIssuesUrl(value *string)() {
    m.issues_url = value
}
// SetKeysUrl sets the keys_url property value. The keys_url property
func (m *Root) SetKeysUrl(value *string)() {
    m.keys_url = value
}
// SetLabelSearchUrl sets the label_search_url property value. The label_search_url property
func (m *Root) SetLabelSearchUrl(value *string)() {
    m.label_search_url = value
}
// SetNotificationsUrl sets the notifications_url property value. The notifications_url property
func (m *Root) SetNotificationsUrl(value *string)() {
    m.notifications_url = value
}
// SetOrganizationRepositoriesUrl sets the organization_repositories_url property value. The organization_repositories_url property
func (m *Root) SetOrganizationRepositoriesUrl(value *string)() {
    m.organization_repositories_url = value
}
// SetOrganizationTeamsUrl sets the organization_teams_url property value. The organization_teams_url property
func (m *Root) SetOrganizationTeamsUrl(value *string)() {
    m.organization_teams_url = value
}
// SetOrganizationUrl sets the organization_url property value. The organization_url property
func (m *Root) SetOrganizationUrl(value *string)() {
    m.organization_url = value
}
// SetPublicGistsUrl sets the public_gists_url property value. The public_gists_url property
func (m *Root) SetPublicGistsUrl(value *string)() {
    m.public_gists_url = value
}
// SetRateLimitUrl sets the rate_limit_url property value. The rate_limit_url property
func (m *Root) SetRateLimitUrl(value *string)() {
    m.rate_limit_url = value
}
// SetRepositorySearchUrl sets the repository_search_url property value. The repository_search_url property
func (m *Root) SetRepositorySearchUrl(value *string)() {
    m.repository_search_url = value
}
// SetRepositoryUrl sets the repository_url property value. The repository_url property
func (m *Root) SetRepositoryUrl(value *string)() {
    m.repository_url = value
}
// SetStarredGistsUrl sets the starred_gists_url property value. The starred_gists_url property
func (m *Root) SetStarredGistsUrl(value *string)() {
    m.starred_gists_url = value
}
// SetStarredUrl sets the starred_url property value. The starred_url property
func (m *Root) SetStarredUrl(value *string)() {
    m.starred_url = value
}
// SetTopicSearchUrl sets the topic_search_url property value. The topic_search_url property
func (m *Root) SetTopicSearchUrl(value *string)() {
    m.topic_search_url = value
}
// SetUserOrganizationsUrl sets the user_organizations_url property value. The user_organizations_url property
func (m *Root) SetUserOrganizationsUrl(value *string)() {
    m.user_organizations_url = value
}
// SetUserRepositoriesUrl sets the user_repositories_url property value. The user_repositories_url property
func (m *Root) SetUserRepositoriesUrl(value *string)() {
    m.user_repositories_url = value
}
// SetUserSearchUrl sets the user_search_url property value. The user_search_url property
func (m *Root) SetUserSearchUrl(value *string)() {
    m.user_search_url = value
}
// SetUserUrl sets the user_url property value. The user_url property
func (m *Root) SetUserUrl(value *string)() {
    m.user_url = value
}
type Rootable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizationsUrl()(*string)
    GetCodeSearchUrl()(*string)
    GetCommitSearchUrl()(*string)
    GetCurrentUserAuthorizationsHtmlUrl()(*string)
    GetCurrentUserRepositoriesUrl()(*string)
    GetCurrentUserUrl()(*string)
    GetEmailsUrl()(*string)
    GetEmojisUrl()(*string)
    GetEventsUrl()(*string)
    GetFeedsUrl()(*string)
    GetFollowersUrl()(*string)
    GetFollowingUrl()(*string)
    GetGistsUrl()(*string)
    GetHubUrl()(*string)
    GetIssueSearchUrl()(*string)
    GetIssuesUrl()(*string)
    GetKeysUrl()(*string)
    GetLabelSearchUrl()(*string)
    GetNotificationsUrl()(*string)
    GetOrganizationRepositoriesUrl()(*string)
    GetOrganizationTeamsUrl()(*string)
    GetOrganizationUrl()(*string)
    GetPublicGistsUrl()(*string)
    GetRateLimitUrl()(*string)
    GetRepositorySearchUrl()(*string)
    GetRepositoryUrl()(*string)
    GetStarredGistsUrl()(*string)
    GetStarredUrl()(*string)
    GetTopicSearchUrl()(*string)
    GetUserOrganizationsUrl()(*string)
    GetUserRepositoriesUrl()(*string)
    GetUserSearchUrl()(*string)
    GetUserUrl()(*string)
    SetAuthorizationsUrl(value *string)()
    SetCodeSearchUrl(value *string)()
    SetCommitSearchUrl(value *string)()
    SetCurrentUserAuthorizationsHtmlUrl(value *string)()
    SetCurrentUserRepositoriesUrl(value *string)()
    SetCurrentUserUrl(value *string)()
    SetEmailsUrl(value *string)()
    SetEmojisUrl(value *string)()
    SetEventsUrl(value *string)()
    SetFeedsUrl(value *string)()
    SetFollowersUrl(value *string)()
    SetFollowingUrl(value *string)()
    SetGistsUrl(value *string)()
    SetHubUrl(value *string)()
    SetIssueSearchUrl(value *string)()
    SetIssuesUrl(value *string)()
    SetKeysUrl(value *string)()
    SetLabelSearchUrl(value *string)()
    SetNotificationsUrl(value *string)()
    SetOrganizationRepositoriesUrl(value *string)()
    SetOrganizationTeamsUrl(value *string)()
    SetOrganizationUrl(value *string)()
    SetPublicGistsUrl(value *string)()
    SetRateLimitUrl(value *string)()
    SetRepositorySearchUrl(value *string)()
    SetRepositoryUrl(value *string)()
    SetStarredGistsUrl(value *string)()
    SetStarredUrl(value *string)()
    SetTopicSearchUrl(value *string)()
    SetUserOrganizationsUrl(value *string)()
    SetUserRepositoriesUrl(value *string)()
    SetUserSearchUrl(value *string)()
    SetUserUrl(value *string)()
}
