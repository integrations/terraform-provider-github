package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Issue issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
type Issue struct {
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
    // Contents of the issue
    body *string
    // The body_html property
    body_html *string
    // The body_text property
    body_text *string
    // The closed_at property
    closed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    closed_by NullableSimpleUserable
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
    // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
    labels []string
    // The labels_url property
    labels_url *string
    // The locked property
    locked *bool
    // A collection of related issues and pull requests.
    milestone NullableMilestoneable
    // The node_id property
    node_id *string
    // Number uniquely identifying the issue within its repository
    number *int32
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    performed_via_github_app NullableIntegrationable
    // The pull_request property
    pull_request Issue_pull_requestable
    // The reactions property
    reactions ReactionRollupable
    // A repository on GitHub.
    repository Repositoryable
    // The repository_url property
    repository_url *string
    // State of the issue; either 'open' or 'closed'
    state *string
    // The reason for the current state
    state_reason *Issue_state_reason
    // The timeline_url property
    timeline_url *string
    // Title of the issue
    title *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // URL for the issue
    url *string
    // A GitHub user.
    user NullableSimpleUserable
}
// NewIssue instantiates a new Issue and sets the default values.
func NewIssue()(*Issue) {
    m := &Issue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIssueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIssue(), nil
}
// GetActiveLockReason gets the active_lock_reason property value. The active_lock_reason property
// returns a *string when successful
func (m *Issue) GetActiveLockReason()(*string) {
    return m.active_lock_reason
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Issue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignee gets the assignee property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Issue) GetAssignee()(NullableSimpleUserable) {
    return m.assignee
}
// GetAssignees gets the assignees property value. The assignees property
// returns a []SimpleUserable when successful
func (m *Issue) GetAssignees()([]SimpleUserable) {
    return m.assignees
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *Issue) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetBody gets the body property value. Contents of the issue
// returns a *string when successful
func (m *Issue) GetBody()(*string) {
    return m.body
}
// GetBodyHtml gets the body_html property value. The body_html property
// returns a *string when successful
func (m *Issue) GetBodyHtml()(*string) {
    return m.body_html
}
// GetBodyText gets the body_text property value. The body_text property
// returns a *string when successful
func (m *Issue) GetBodyText()(*string) {
    return m.body_text
}
// GetClosedAt gets the closed_at property value. The closed_at property
// returns a *Time when successful
func (m *Issue) GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closed_at
}
// GetClosedBy gets the closed_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Issue) GetClosedBy()(NullableSimpleUserable) {
    return m.closed_by
}
// GetComments gets the comments property value. The comments property
// returns a *int32 when successful
func (m *Issue) GetComments()(*int32) {
    return m.comments
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *Issue) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Issue) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDraft gets the draft property value. The draft property
// returns a *bool when successful
func (m *Issue) GetDraft()(*bool) {
    return m.draft
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *Issue) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Issue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["closed_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedBy(val.(NullableSimpleUserable))
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
        val, err := n.GetObjectValue(CreateIssue_pull_requestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequest(val.(Issue_pull_requestable))
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
        val, err := n.GetEnumValue(ParseIssue_state_reason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateReason(val.(*Issue_state_reason))
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
func (m *Issue) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *Issue) GetId()(*int64) {
    return m.id
}
// GetLabels gets the labels property value. Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
// returns a []string when successful
func (m *Issue) GetLabels()([]string) {
    return m.labels
}
// GetLabelsUrl gets the labels_url property value. The labels_url property
// returns a *string when successful
func (m *Issue) GetLabelsUrl()(*string) {
    return m.labels_url
}
// GetLocked gets the locked property value. The locked property
// returns a *bool when successful
func (m *Issue) GetLocked()(*bool) {
    return m.locked
}
// GetMilestone gets the milestone property value. A collection of related issues and pull requests.
// returns a NullableMilestoneable when successful
func (m *Issue) GetMilestone()(NullableMilestoneable) {
    return m.milestone
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Issue) GetNodeId()(*string) {
    return m.node_id
}
// GetNumber gets the number property value. Number uniquely identifying the issue within its repository
// returns a *int32 when successful
func (m *Issue) GetNumber()(*int32) {
    return m.number
}
// GetPerformedViaGithubApp gets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
// returns a NullableIntegrationable when successful
func (m *Issue) GetPerformedViaGithubApp()(NullableIntegrationable) {
    return m.performed_via_github_app
}
// GetPullRequest gets the pull_request property value. The pull_request property
// returns a Issue_pull_requestable when successful
func (m *Issue) GetPullRequest()(Issue_pull_requestable) {
    return m.pull_request
}
// GetReactions gets the reactions property value. The reactions property
// returns a ReactionRollupable when successful
func (m *Issue) GetReactions()(ReactionRollupable) {
    return m.reactions
}
// GetRepository gets the repository property value. A repository on GitHub.
// returns a Repositoryable when successful
func (m *Issue) GetRepository()(Repositoryable) {
    return m.repository
}
// GetRepositoryUrl gets the repository_url property value. The repository_url property
// returns a *string when successful
func (m *Issue) GetRepositoryUrl()(*string) {
    return m.repository_url
}
// GetState gets the state property value. State of the issue; either 'open' or 'closed'
// returns a *string when successful
func (m *Issue) GetState()(*string) {
    return m.state
}
// GetStateReason gets the state_reason property value. The reason for the current state
// returns a *Issue_state_reason when successful
func (m *Issue) GetStateReason()(*Issue_state_reason) {
    return m.state_reason
}
// GetTimelineUrl gets the timeline_url property value. The timeline_url property
// returns a *string when successful
func (m *Issue) GetTimelineUrl()(*string) {
    return m.timeline_url
}
// GetTitle gets the title property value. Title of the issue
// returns a *string when successful
func (m *Issue) GetTitle()(*string) {
    return m.title
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Issue) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. URL for the issue
// returns a *string when successful
func (m *Issue) GetUrl()(*string) {
    return m.url
}
// GetUser gets the user property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Issue) GetUser()(NullableSimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *Issue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("closed_by", m.GetClosedBy())
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
        err := writer.WriteCollectionOfStringValues("labels", m.GetLabels())
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
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    if m.GetStateReason() != nil {
        cast := (*m.GetStateReason()).String()
        err := writer.WriteStringValue("state_reason", &cast)
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
func (m *Issue) SetActiveLockReason(value *string)() {
    m.active_lock_reason = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Issue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignee sets the assignee property value. A GitHub user.
func (m *Issue) SetAssignee(value NullableSimpleUserable)() {
    m.assignee = value
}
// SetAssignees sets the assignees property value. The assignees property
func (m *Issue) SetAssignees(value []SimpleUserable)() {
    m.assignees = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *Issue) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetBody sets the body property value. Contents of the issue
func (m *Issue) SetBody(value *string)() {
    m.body = value
}
// SetBodyHtml sets the body_html property value. The body_html property
func (m *Issue) SetBodyHtml(value *string)() {
    m.body_html = value
}
// SetBodyText sets the body_text property value. The body_text property
func (m *Issue) SetBodyText(value *string)() {
    m.body_text = value
}
// SetClosedAt sets the closed_at property value. The closed_at property
func (m *Issue) SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closed_at = value
}
// SetClosedBy sets the closed_by property value. A GitHub user.
func (m *Issue) SetClosedBy(value NullableSimpleUserable)() {
    m.closed_by = value
}
// SetComments sets the comments property value. The comments property
func (m *Issue) SetComments(value *int32)() {
    m.comments = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *Issue) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Issue) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDraft sets the draft property value. The draft property
func (m *Issue) SetDraft(value *bool)() {
    m.draft = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *Issue) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Issue) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *Issue) SetId(value *int64)() {
    m.id = value
}
// SetLabels sets the labels property value. Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
func (m *Issue) SetLabels(value []string)() {
    m.labels = value
}
// SetLabelsUrl sets the labels_url property value. The labels_url property
func (m *Issue) SetLabelsUrl(value *string)() {
    m.labels_url = value
}
// SetLocked sets the locked property value. The locked property
func (m *Issue) SetLocked(value *bool)() {
    m.locked = value
}
// SetMilestone sets the milestone property value. A collection of related issues and pull requests.
func (m *Issue) SetMilestone(value NullableMilestoneable)() {
    m.milestone = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Issue) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNumber sets the number property value. Number uniquely identifying the issue within its repository
func (m *Issue) SetNumber(value *int32)() {
    m.number = value
}
// SetPerformedViaGithubApp sets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *Issue) SetPerformedViaGithubApp(value NullableIntegrationable)() {
    m.performed_via_github_app = value
}
// SetPullRequest sets the pull_request property value. The pull_request property
func (m *Issue) SetPullRequest(value Issue_pull_requestable)() {
    m.pull_request = value
}
// SetReactions sets the reactions property value. The reactions property
func (m *Issue) SetReactions(value ReactionRollupable)() {
    m.reactions = value
}
// SetRepository sets the repository property value. A repository on GitHub.
func (m *Issue) SetRepository(value Repositoryable)() {
    m.repository = value
}
// SetRepositoryUrl sets the repository_url property value. The repository_url property
func (m *Issue) SetRepositoryUrl(value *string)() {
    m.repository_url = value
}
// SetState sets the state property value. State of the issue; either 'open' or 'closed'
func (m *Issue) SetState(value *string)() {
    m.state = value
}
// SetStateReason sets the state_reason property value. The reason for the current state
func (m *Issue) SetStateReason(value *Issue_state_reason)() {
    m.state_reason = value
}
// SetTimelineUrl sets the timeline_url property value. The timeline_url property
func (m *Issue) SetTimelineUrl(value *string)() {
    m.timeline_url = value
}
// SetTitle sets the title property value. Title of the issue
func (m *Issue) SetTitle(value *string)() {
    m.title = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Issue) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. URL for the issue
func (m *Issue) SetUrl(value *string)() {
    m.url = value
}
// SetUser sets the user property value. A GitHub user.
func (m *Issue) SetUser(value NullableSimpleUserable)() {
    m.user = value
}
type Issueable interface {
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
    GetClosedBy()(NullableSimpleUserable)
    GetComments()(*int32)
    GetCommentsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDraft()(*bool)
    GetEventsUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetLabels()([]string)
    GetLabelsUrl()(*string)
    GetLocked()(*bool)
    GetMilestone()(NullableMilestoneable)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetPullRequest()(Issue_pull_requestable)
    GetReactions()(ReactionRollupable)
    GetRepository()(Repositoryable)
    GetRepositoryUrl()(*string)
    GetState()(*string)
    GetStateReason()(*Issue_state_reason)
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
    SetClosedBy(value NullableSimpleUserable)()
    SetComments(value *int32)()
    SetCommentsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDraft(value *bool)()
    SetEventsUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetLabels(value []string)()
    SetLabelsUrl(value *string)()
    SetLocked(value *bool)()
    SetMilestone(value NullableMilestoneable)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetPullRequest(value Issue_pull_requestable)()
    SetReactions(value ReactionRollupable)()
    SetRepository(value Repositoryable)()
    SetRepositoryUrl(value *string)()
    SetState(value *string)()
    SetStateReason(value *Issue_state_reason)()
    SetTimelineUrl(value *string)()
    SetTitle(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
