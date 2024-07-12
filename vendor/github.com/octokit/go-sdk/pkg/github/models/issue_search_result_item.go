package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssueSearchResultItem issue Search Result Item
type IssueSearchResultItem struct {
    // The active_lock_reason property
    active_lock_reason *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    assignee NullableSimpleUserable
    // The assignees property
    assignees []SimpleUserable
    // How the author is associated with the repository.
    author_association *AuthorAssociation
    // The body property
    body *string
    // The body_html property
    body_html *string
    // The body_text property
    body_text *string
    // The closed_at property
    closed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The comments property
    comments *int32
    // The comments_url property
    comments_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The draft property
    draft *bool
    // The events_url property
    events_url *string
    // The html_url property
    html_url *string
    // The id property
    id *int64
    // The labels property
    labels []IssueSearchResultItem_labelsable
    // The labels_url property
    labels_url *string
    // The locked property
    locked *bool
    // A collection of related issues and pull requests.
    milestone NullableMilestoneable
    // The node_id property
    node_id *string
    // The number property
    number *int32
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    performed_via_github_app NullableIntegrationable
    // The pull_request property
    pull_request IssueSearchResultItem_pull_requestable
    // The reactions property
    reactions ReactionRollupable
    // A repository on GitHub.
    repository Repositoryable
    // The repository_url property
    repository_url *string
    // The score property
    score *float64
    // The state property
    state *string
    // The state_reason property
    state_reason *string
    // The text_matches property
    text_matches []Issuesable
    // The timeline_url property
    timeline_url *string
    // The title property
    title *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewIssueSearchResultItem instantiates a new IssueSearchResultItem and sets the default values.
func NewIssueSearchResultItem()(*IssueSearchResultItem) {
    m := &IssueSearchResultItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIssueSearchResultItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssueSearchResultItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIssueSearchResultItem(), nil
}
// GetActiveLockReason gets the active_lock_reason property value. The active_lock_reason property
// returns a *string when successful
func (m *IssueSearchResultItem) GetActiveLockReason()(*string) {
    return m.active_lock_reason
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IssueSearchResultItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignee gets the assignee property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *IssueSearchResultItem) GetAssignee()(NullableSimpleUserable) {
    return m.assignee
}
// GetAssignees gets the assignees property value. The assignees property
// returns a []SimpleUserable when successful
func (m *IssueSearchResultItem) GetAssignees()([]SimpleUserable) {
    return m.assignees
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *IssueSearchResultItem) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *IssueSearchResultItem) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *IssueSearchResultItem) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyText gets the body_text property value. The body_text property
// returns a *string when successful
func (m *IssueSearchResultItem) GetBodyText()(*string) {
    return m.body_text
}
// GetClosedAt gets the closed_at property value. The closed_at property
// returns a *Time when successful
func (m *IssueSearchResultItem) GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closed_at
}
// GetComments gets the comments property value. The comments property
// returns a *int32 when successful
func (m *IssueSearchResultItem) GetComments()(*int32) {
    return m.comments
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *IssueSearchResultItem) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *IssueSearchResultItem) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDraft gets the draft property value. The draft property
// returns a *bool when successful
func (m *IssueSearchResultItem) GetDraft()(*bool) {
    return m.draft
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *IssueSearchResultItem) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IssueSearchResultItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active_lock_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveLockReason(val)
        }
        return nil
    }
    res["assignee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignee(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["assignees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimpleUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SimpleUserable)
                }
            }
            m.SetAssignees(res)
        }
        return nil
    }
    res["author_association"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthorAssociation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorAssociation(val.(*AuthorAssociation))
        }
        return nil
    }
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["body_html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyHtml(val)
        }
        return nil
    }
    res["body_text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyText(val)
        }
        return nil
    }
    res["closed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedAt(val)
        }
        return nil
    }
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComments(val)
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
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["draft"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraft(val)
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
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIssueSearchResultItem_labelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IssueSearchResultItem_labelsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IssueSearchResultItem_labelsable)
                }
            }
            m.SetLabels(res)
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
    res["locked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocked(val)
        }
        return nil
    }
    res["milestone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableMilestoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestone(val.(NullableMilestoneable))
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
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["performed_via_github_app"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformedViaGithubApp(val.(NullableIntegrationable))
        }
        return nil
    }
    res["pull_request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIssueSearchResultItem_pull_requestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequest(val.(IssueSearchResultItem_pull_requestable))
        }
        return nil
    }
    res["reactions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReactionRollupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReactions(val.(ReactionRollupable))
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(Repositoryable))
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["state_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateReason(val)
        }
        return nil
    }
    res["text_matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIssuesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Issuesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Issuesable)
                }
            }
            m.SetTextMatches(res)
        }
        return nil
    }
    res["timeline_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimelineUrl(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
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
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(NullableSimpleUserable))
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *IssueSearchResultItem) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *IssueSearchResultItem) GetId()(*int64) {
    return m.id
}
// GetLabels gets the labels property value. The labels property
// returns a []IssueSearchResultItem_labelsable when successful
func (m *IssueSearchResultItem) GetLabels()([]IssueSearchResultItem_labelsable) {
    return m.labels
}
// GetLabelsUrl gets the labels_url property value. The labels_url property
// returns a *string when successful
func (m *IssueSearchResultItem) GetLabelsUrl()(*string) {
    return m.labels_url
}
// GetLocked gets the locked property value. The locked property
// returns a *bool when successful
func (m *IssueSearchResultItem) GetLocked()(*bool) {
    return m.locked
}
// GetMilestone gets the milestone property value. A collection of related issues and pull requests.
// returns a NullableMilestoneable when successful
func (m *IssueSearchResultItem) GetMilestone()(NullableMilestoneable) {
    return m.milestone
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *IssueSearchResultItem) GetNodeId()(*string) {
    return m.node_id
}
// GetNumber gets the number property value. The number property
// returns a *int32 when successful
func (m *IssueSearchResultItem) GetNumber()(*int32) {
    return m.number
}
// GetPerformedViaGithubApp gets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
// returns a NullableIntegrationable when successful
func (m *IssueSearchResultItem) GetPerformedViaGithubApp()(NullableIntegrationable) {
    return m.performed_via_github_app
}
// GetPullRequest gets the pull_request property value. The pull_request property
// returns a IssueSearchResultItem_pull_requestable when successful
func (m *IssueSearchResultItem) GetPullRequest()(IssueSearchResultItem_pull_requestable) {
    return m.pull_request
}
// GetReactions gets the reactions property value. The reactions property
// returns a ReactionRollupable when successful
func (m *IssueSearchResultItem) GetReactions()(ReactionRollupable) {
    return m.reactions
}
// GetRepository gets the repository property value. A repository on GitHub.
// returns a Repositoryable when successful
func (m *IssueSearchResultItem) GetRepository()(Repositoryable) {
    return m.repository
}
// GetRepositoryUrl gets the repository_url property value. The repository_url property
// returns a *string when successful
func (m *IssueSearchResultItem) GetRepositoryUrl()(*string) {
    return m.repository_url
}
// GetScore gets the score property value. The score property
// returns a *float64 when successful
func (m *IssueSearchResultItem) GetScore()(*float64) {
    return m.score
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *IssueSearchResultItem) GetState()(*string) {
    return m.state
}
// GetStateReason gets the state_reason property value. The state_reason property
// returns a *string when successful
func (m *IssueSearchResultItem) GetStateReason()(*string) {
    return m.state_reason
}
// GetTextMatches gets the text_matches property value. The text_matches property
// returns a []Issuesable when successful
func (m *IssueSearchResultItem) GetTextMatches()([]Issuesable) {
    return m.text_matches
}
// GetTimelineUrl gets the timeline_url property value. The timeline_url property
// returns a *string when successful
func (m *IssueSearchResultItem) GetTimelineUrl()(*string) {
    return m.timeline_url
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *IssueSearchResultItem) GetTitle()(*string) {
    return m.title
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *IssueSearchResultItem) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *IssueSearchResultItem) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *IssueSearchResultItem) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *IssueSearchResultItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("active_lock_reason", m.GetActiveLockReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("assignee", m.GetAssignee())
        if err != nil {
            return err
        }
    }
    if m.GetAssignees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignees()))
        for i, v := range m.GetAssignees() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("assignees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthorAssociation() != nil {
        cast := (*m.GetAuthorAssociation()).String()
        err := writer.WriteStringValue("author_association", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body_html", m.GetBodyHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body_text", m.GetBodyText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("closed_at", m.GetClosedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("comments", m.GetComments())
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
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("draft", m.GetDraft())
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
    if m.GetLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("labels", cast)
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
        err := writer.WriteBoolValue("locked", m.GetLocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("milestone", m.GetMilestone())
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
        err := writer.WriteInt32Value("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("performed_via_github_app", m.GetPerformedViaGithubApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pull_request", m.GetPullRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reactions", m.GetReactions())
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
        err := writer.WriteStringValue("repository_url", m.GetRepositoryUrl())
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
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state_reason", m.GetStateReason())
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
        err := writer.WriteStringValue("timeline_url", m.GetTimelineUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
        err := writer.WriteObjectValue("user", m.GetUser())
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
// SetActiveLockReason sets the active_lock_reason property value. The active_lock_reason property
func (m *IssueSearchResultItem) SetActiveLockReason(value *string)() {
    m.active_lock_reason = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IssueSearchResultItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignee sets the assignee property value. A GitHub user.
func (m *IssueSearchResultItem) SetAssignee(value NullableSimpleUserable)() {
    m.assignee = value
}
// SetAssignees sets the assignees property value. The assignees property
func (m *IssueSearchResultItem) SetAssignees(value []SimpleUserable)() {
    m.assignees = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *IssueSearchResultItem) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetBody sets the body property value. The body property
func (m *IssueSearchResultItem) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *IssueSearchResultItem) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyText sets the body_text property value. The body_text property
func (m *IssueSearchResultItem) SetBodyText(value *string)() {
    m.body_text = value
}
// SetClosedAt sets the closed_at property value. The closed_at property
func (m *IssueSearchResultItem) SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closed_at = value
}
// SetComments sets the comments property value. The comments property
func (m *IssueSearchResultItem) SetComments(value *int32)() {
    m.comments = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *IssueSearchResultItem) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *IssueSearchResultItem) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDraft sets the draft property value. The draft property
func (m *IssueSearchResultItem) SetDraft(value *bool)() {
    m.draft = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *IssueSearchResultItem) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *IssueSearchResultItem) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *IssueSearchResultItem) SetId(value *int64)() {
    m.id = value
}
// SetLabels sets the labels property value. The labels property
func (m *IssueSearchResultItem) SetLabels(value []IssueSearchResultItem_labelsable)() {
    m.labels = value
}
// SetLabelsUrl sets the labels_url property value. The labels_url property
func (m *IssueSearchResultItem) SetLabelsUrl(value *string)() {
    m.labels_url = value
}
// SetLocked sets the locked property value. The locked property
func (m *IssueSearchResultItem) SetLocked(value *bool)() {
    m.locked = value
}
// SetMilestone sets the milestone property value. A collection of related issues and pull requests.
func (m *IssueSearchResultItem) SetMilestone(value NullableMilestoneable)() {
    m.milestone = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *IssueSearchResultItem) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNumber sets the number property value. The number property
func (m *IssueSearchResultItem) SetNumber(value *int32)() {
    m.number = value
}
// SetPerformedViaGithubApp sets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *IssueSearchResultItem) SetPerformedViaGithubApp(value NullableIntegrationable)() {
    m.performed_via_github_app = value
}
// SetPullRequest sets the pull_request property value. The pull_request property
func (m *IssueSearchResultItem) SetPullRequest(value IssueSearchResultItem_pull_requestable)() {
    m.pull_request = value
}
// SetReactions sets the reactions property value. The reactions property
func (m *IssueSearchResultItem) SetReactions(value ReactionRollupable)() {
    m.reactions = value
}
// SetRepository sets the repository property value. A repository on GitHub.
func (m *IssueSearchResultItem) SetRepository(value Repositoryable)() {
    m.repository = value
}
// SetRepositoryUrl sets the repository_url property value. The repository_url property
func (m *IssueSearchResultItem) SetRepositoryUrl(value *string)() {
    m.repository_url = value
}
// SetScore sets the score property value. The score property
func (m *IssueSearchResultItem) SetScore(value *float64)() {
    m.score = value
}
// SetState sets the state property value. The state property
func (m *IssueSearchResultItem) SetState(value *string)() {
    m.state = value
}
// SetStateReason sets the state_reason property value. The state_reason property
func (m *IssueSearchResultItem) SetStateReason(value *string)() {
    m.state_reason = value
}
// SetTextMatches sets the text_matches property value. The text_matches property
func (m *IssueSearchResultItem) SetTextMatches(value []Issuesable)() {
    m.text_matches = value
}
// SetTimelineUrl sets the timeline_url property value. The timeline_url property
func (m *IssueSearchResultItem) SetTimelineUrl(value *string)() {
    m.timeline_url = value
}
// SetTitle sets the title property value. The title property
func (m *IssueSearchResultItem) SetTitle(value *string)() {
    m.title = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *IssueSearchResultItem) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *IssueSearchResultItem) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *IssueSearchResultItem) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type IssueSearchResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveLockReason()(*string)
    GetAssignee()(NullableSimpleUserable)
    GetAssignees()([]SimpleUserable)
    GetAuthorAssociation()(*AuthorAssociation)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyText()(*string)
    GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetComments()(*int32)
    GetCommentsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDraft()(*bool)
    GetEventsUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetLabels()([]IssueSearchResultItem_labelsable)
    GetLabelsUrl()(*string)
    GetLocked()(*bool)
    GetMilestone()(NullableMilestoneable)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetPullRequest()(IssueSearchResultItem_pull_requestable)
    GetReactions()(ReactionRollupable)
    GetRepository()(Repositoryable)
    GetRepositoryUrl()(*string)
    GetScore()(*float64)
    GetState()(*string)
    GetStateReason()(*string)
    GetTextMatches()([]Issuesable)
    GetTimelineUrl()(*string)
    GetTitle()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
    SetActiveLockReason(value *string)()
    SetAssignee(value NullableSimpleUserable)()
    SetAssignees(value []SimpleUserable)()
    SetAuthorAssociation(value *AuthorAssociation)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyText(value *string)()
    SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetComments(value *int32)()
    SetCommentsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDraft(value *bool)()
    SetEventsUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetLabels(value []IssueSearchResultItem_labelsable)()
    SetLabelsUrl(value *string)()
    SetLocked(value *bool)()
    SetMilestone(value NullableMilestoneable)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetPullRequest(value IssueSearchResultItem_pull_requestable)()
    SetReactions(value ReactionRollupable)()
    SetRepository(value Repositoryable)()
    SetRepositoryUrl(value *string)()
    SetScore(value *float64)()
    SetState(value *string)()
    SetStateReason(value *string)()
    SetTextMatches(value []Issuesable)()
    SetTimelineUrl(value *string)()
    SetTitle(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
