package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssueEvent issue Event
type IssueEvent struct {
    // A GitHub user.
    actor NullableSimpleUserable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    assignee NullableSimpleUserable
    // A GitHub user.
    assigner NullableSimpleUserable
    // How the author is associated with the repository.
    author_association *AuthorAssociation
    // The commit_id property
    commit_id *string
    // The commit_url property
    commit_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The dismissed_review property
    dismissed_review IssueEventDismissedReviewable
    // The event property
    event *string
    // The id property
    id *int64
    // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
    issue NullableIssueable
    // Issue Event Label
    label IssueEventLabelable
    // The lock_reason property
    lock_reason *string
    // Issue Event Milestone
    milestone IssueEventMilestoneable
    // The node_id property
    node_id *string
    // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
    performed_via_github_app NullableIntegrationable
    // Issue Event Project Card
    project_card IssueEventProjectCardable
    // Issue Event Rename
    rename IssueEventRenameable
    // A GitHub user.
    requested_reviewer NullableSimpleUserable
    // Groups of organization members that gives permissions on specified repositories.
    requested_team Teamable
    // A GitHub user.
    review_requester NullableSimpleUserable
    // The url property
    url *string
}
// NewIssueEvent instantiates a new IssueEvent and sets the default values.
func NewIssueEvent()(*IssueEvent) {
    m := &IssueEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIssueEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssueEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIssueEvent(), nil
}
// GetActor gets the actor property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *IssueEvent) GetActor()(NullableSimpleUserable) {
    return m.actor
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IssueEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignee gets the assignee property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *IssueEvent) GetAssignee()(NullableSimpleUserable) {
    return m.assignee
}
// GetAssigner gets the assigner property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *IssueEvent) GetAssigner()(NullableSimpleUserable) {
    return m.assigner
}
// GetAuthorAssociation gets the author_association property value. How the author is associated with the repository.
// returns a *AuthorAssociation when successful
func (m *IssueEvent) GetAuthorAssociation()(*AuthorAssociation) {
    return m.author_association
}
// GetCommitId gets the commit_id property value. The commit_id property
// returns a *string when successful
func (m *IssueEvent) GetCommitId()(*string) {
    return m.commit_id
}
// GetCommitUrl gets the commit_url property value. The commit_url property
// returns a *string when successful
func (m *IssueEvent) GetCommitUrl()(*string) {
    return m.commit_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *IssueEvent) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDismissedReview gets the dismissed_review property value. The dismissed_review property
// returns a IssueEventDismissedReviewable when successful
func (m *IssueEvent) GetDismissedReview()(IssueEventDismissedReviewable) {
    return m.dismissed_review
}
// GetEvent gets the event property value. The event property
// returns a *string when successful
func (m *IssueEvent) GetEvent()(*string) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IssueEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(NullableSimpleUserable))
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
    res["assigner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigner(val.(NullableSimpleUserable))
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
    res["commit_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitId(val)
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
    res["dismissed_review"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIssueEventDismissedReviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedReview(val.(IssueEventDismissedReviewable))
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val)
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
    res["issue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssue(val.(NullableIssueable))
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIssueEventLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val.(IssueEventLabelable))
        }
        return nil
    }
    res["lock_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockReason(val)
        }
        return nil
    }
    res["milestone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIssueEventMilestoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestone(val.(IssueEventMilestoneable))
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
    res["project_card"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIssueEventProjectCardFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProjectCard(val.(IssueEventProjectCardable))
        }
        return nil
    }
    res["rename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIssueEventRenameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRename(val.(IssueEventRenameable))
        }
        return nil
    }
    res["requested_reviewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedReviewer(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["requested_team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedTeam(val.(Teamable))
        }
        return nil
    }
    res["review_requester"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewRequester(val.(NullableSimpleUserable))
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
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *IssueEvent) GetId()(*int64) {
    return m.id
}
// GetIssue gets the issue property value. Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
// returns a NullableIssueable when successful
func (m *IssueEvent) GetIssue()(NullableIssueable) {
    return m.issue
}
// GetLabel gets the label property value. Issue Event Label
// returns a IssueEventLabelable when successful
func (m *IssueEvent) GetLabel()(IssueEventLabelable) {
    return m.label
}
// GetLockReason gets the lock_reason property value. The lock_reason property
// returns a *string when successful
func (m *IssueEvent) GetLockReason()(*string) {
    return m.lock_reason
}
// GetMilestone gets the milestone property value. Issue Event Milestone
// returns a IssueEventMilestoneable when successful
func (m *IssueEvent) GetMilestone()(IssueEventMilestoneable) {
    return m.milestone
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *IssueEvent) GetNodeId()(*string) {
    return m.node_id
}
// GetPerformedViaGithubApp gets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
// returns a NullableIntegrationable when successful
func (m *IssueEvent) GetPerformedViaGithubApp()(NullableIntegrationable) {
    return m.performed_via_github_app
}
// GetProjectCard gets the project_card property value. Issue Event Project Card
// returns a IssueEventProjectCardable when successful
func (m *IssueEvent) GetProjectCard()(IssueEventProjectCardable) {
    return m.project_card
}
// GetRename gets the rename property value. Issue Event Rename
// returns a IssueEventRenameable when successful
func (m *IssueEvent) GetRename()(IssueEventRenameable) {
    return m.rename
}
// GetRequestedReviewer gets the requested_reviewer property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *IssueEvent) GetRequestedReviewer()(NullableSimpleUserable) {
    return m.requested_reviewer
}
// GetRequestedTeam gets the requested_team property value. Groups of organization members that gives permissions on specified repositories.
// returns a Teamable when successful
func (m *IssueEvent) GetRequestedTeam()(Teamable) {
    return m.requested_team
}
// GetReviewRequester gets the review_requester property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *IssueEvent) GetReviewRequester()(NullableSimpleUserable) {
    return m.review_requester
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *IssueEvent) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *IssueEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("actor", m.GetActor())
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
    {
        err := writer.WriteObjectValue("assigner", m.GetAssigner())
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
        err := writer.WriteStringValue("commit_id", m.GetCommitId())
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
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dismissed_review", m.GetDismissedReview())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event", m.GetEvent())
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
        err := writer.WriteObjectValue("issue", m.GetIssue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lock_reason", m.GetLockReason())
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
        err := writer.WriteObjectValue("performed_via_github_app", m.GetPerformedViaGithubApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("project_card", m.GetProjectCard())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rename", m.GetRename())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requested_reviewer", m.GetRequestedReviewer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requested_team", m.GetRequestedTeam())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("review_requester", m.GetReviewRequester())
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
// SetActor sets the actor property value. A GitHub user.
func (m *IssueEvent) SetActor(value NullableSimpleUserable)() {
    m.actor = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IssueEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignee sets the assignee property value. A GitHub user.
func (m *IssueEvent) SetAssignee(value NullableSimpleUserable)() {
    m.assignee = value
}
// SetAssigner sets the assigner property value. A GitHub user.
func (m *IssueEvent) SetAssigner(value NullableSimpleUserable)() {
    m.assigner = value
}
// SetAuthorAssociation sets the author_association property value. How the author is associated with the repository.
func (m *IssueEvent) SetAuthorAssociation(value *AuthorAssociation)() {
    m.author_association = value
}
// SetCommitId sets the commit_id property value. The commit_id property
func (m *IssueEvent) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetCommitUrl sets the commit_url property value. The commit_url property
func (m *IssueEvent) SetCommitUrl(value *string)() {
    m.commit_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *IssueEvent) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDismissedReview sets the dismissed_review property value. The dismissed_review property
func (m *IssueEvent) SetDismissedReview(value IssueEventDismissedReviewable)() {
    m.dismissed_review = value
}
// SetEvent sets the event property value. The event property
func (m *IssueEvent) SetEvent(value *string)() {
    m.event = value
}
// SetId sets the id property value. The id property
func (m *IssueEvent) SetId(value *int64)() {
    m.id = value
}
// SetIssue sets the issue property value. Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
func (m *IssueEvent) SetIssue(value NullableIssueable)() {
    m.issue = value
}
// SetLabel sets the label property value. Issue Event Label
func (m *IssueEvent) SetLabel(value IssueEventLabelable)() {
    m.label = value
}
// SetLockReason sets the lock_reason property value. The lock_reason property
func (m *IssueEvent) SetLockReason(value *string)() {
    m.lock_reason = value
}
// SetMilestone sets the milestone property value. Issue Event Milestone
func (m *IssueEvent) SetMilestone(value IssueEventMilestoneable)() {
    m.milestone = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *IssueEvent) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPerformedViaGithubApp sets the performed_via_github_app property value. GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
func (m *IssueEvent) SetPerformedViaGithubApp(value NullableIntegrationable)() {
    m.performed_via_github_app = value
}
// SetProjectCard sets the project_card property value. Issue Event Project Card
func (m *IssueEvent) SetProjectCard(value IssueEventProjectCardable)() {
    m.project_card = value
}
// SetRename sets the rename property value. Issue Event Rename
func (m *IssueEvent) SetRename(value IssueEventRenameable)() {
    m.rename = value
}
// SetRequestedReviewer sets the requested_reviewer property value. A GitHub user.
func (m *IssueEvent) SetRequestedReviewer(value NullableSimpleUserable)() {
    m.requested_reviewer = value
}
// SetRequestedTeam sets the requested_team property value. Groups of organization members that gives permissions on specified repositories.
func (m *IssueEvent) SetRequestedTeam(value Teamable)() {
    m.requested_team = value
}
// SetReviewRequester sets the review_requester property value. A GitHub user.
func (m *IssueEvent) SetReviewRequester(value NullableSimpleUserable)() {
    m.review_requester = value
}
// SetUrl sets the url property value. The url property
func (m *IssueEvent) SetUrl(value *string)() {
    m.url = value
}
type IssueEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActor()(NullableSimpleUserable)
    GetAssignee()(NullableSimpleUserable)
    GetAssigner()(NullableSimpleUserable)
    GetAuthorAssociation()(*AuthorAssociation)
    GetCommitId()(*string)
    GetCommitUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedReview()(IssueEventDismissedReviewable)
    GetEvent()(*string)
    GetId()(*int64)
    GetIssue()(NullableIssueable)
    GetLabel()(IssueEventLabelable)
    GetLockReason()(*string)
    GetMilestone()(IssueEventMilestoneable)
    GetNodeId()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetProjectCard()(IssueEventProjectCardable)
    GetRename()(IssueEventRenameable)
    GetRequestedReviewer()(NullableSimpleUserable)
    GetRequestedTeam()(Teamable)
    GetReviewRequester()(NullableSimpleUserable)
    GetUrl()(*string)
    SetActor(value NullableSimpleUserable)()
    SetAssignee(value NullableSimpleUserable)()
    SetAssigner(value NullableSimpleUserable)()
    SetAuthorAssociation(value *AuthorAssociation)()
    SetCommitId(value *string)()
    SetCommitUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedReview(value IssueEventDismissedReviewable)()
    SetEvent(value *string)()
    SetId(value *int64)()
    SetIssue(value NullableIssueable)()
    SetLabel(value IssueEventLabelable)()
    SetLockReason(value *string)()
    SetMilestone(value IssueEventMilestoneable)()
    SetNodeId(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetProjectCard(value IssueEventProjectCardable)()
    SetRename(value IssueEventRenameable)()
    SetRequestedReviewer(value NullableSimpleUserable)()
    SetRequestedTeam(value Teamable)()
    SetReviewRequester(value NullableSimpleUserable)()
    SetUrl(value *string)()
}
