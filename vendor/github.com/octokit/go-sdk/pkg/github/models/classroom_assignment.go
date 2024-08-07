package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassroomAssignment a GitHub Classroom assignment
type ClassroomAssignment struct {
    // The number of students that have accepted the assignment.
    accepted *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub Classroom classroom
    classroom Classroomable
    // The time at which the assignment is due.
    deadline *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The selected editor for the assignment.
    editor *string
    // Whether feedback pull request will be created when a student accepts the assignment.
    feedback_pull_requests_enabled *bool
    // Unique identifier of the repository.
    id *int32
    // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
    invitations_enabled *bool
    // The link that a student can use to accept the assignment.
    invite_link *string
    // The programming language used in the assignment.
    language *string
    // The maximum allowable members per team.
    max_members *int32
    // The maximum allowable teams for the assignment.
    max_teams *int32
    // The number of students that have passed the assignment.
    passing *int32
    // Whether an accepted assignment creates a public repository.
    public_repo *bool
    // Sluggified name of the assignment.
    slug *string
    // A GitHub repository view for Classroom
    starter_code_repository SimpleClassroomRepositoryable
    // Whether students are admins on created repository when a student accepts the assignment.
    students_are_repo_admins *bool
    // The number of students that have submitted the assignment.
    submitted *int32
    // Assignment title.
    title *string
    // Whether it's a group assignment or individual assignment.
    typeEscaped *ClassroomAssignment_type
}
// NewClassroomAssignment instantiates a new ClassroomAssignment and sets the default values.
func NewClassroomAssignment()(*ClassroomAssignment) {
    m := &ClassroomAssignment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClassroomAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClassroomAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassroomAssignment(), nil
}
// GetAccepted gets the accepted property value. The number of students that have accepted the assignment.
// returns a *int32 when successful
func (m *ClassroomAssignment) GetAccepted()(*int32) {
    return m.accepted
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ClassroomAssignment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClassroom gets the classroom property value. A GitHub Classroom classroom
// returns a Classroomable when successful
func (m *ClassroomAssignment) GetClassroom()(Classroomable) {
    return m.classroom
}
// GetDeadline gets the deadline property value. The time at which the assignment is due.
// returns a *Time when successful
func (m *ClassroomAssignment) GetDeadline()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deadline
}
// GetEditor gets the editor property value. The selected editor for the assignment.
// returns a *string when successful
func (m *ClassroomAssignment) GetEditor()(*string) {
    return m.editor
}
// GetFeedbackPullRequestsEnabled gets the feedback_pull_requests_enabled property value. Whether feedback pull request will be created when a student accepts the assignment.
// returns a *bool when successful
func (m *ClassroomAssignment) GetFeedbackPullRequestsEnabled()(*bool) {
    return m.feedback_pull_requests_enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClassroomAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccepted(val)
        }
        return nil
    }
    res["classroom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClassroomFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroom(val.(Classroomable))
        }
        return nil
    }
    res["deadline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeadline(val)
        }
        return nil
    }
    res["editor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEditor(val)
        }
        return nil
    }
    res["feedback_pull_requests_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedbackPullRequestsEnabled(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["invitations_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitationsEnabled(val)
        }
        return nil
    }
    res["invite_link"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInviteLink(val)
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
    res["max_members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxMembers(val)
        }
        return nil
    }
    res["max_teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxTeams(val)
        }
        return nil
    }
    res["passing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassing(val)
        }
        return nil
    }
    res["public_repo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicRepo(val)
        }
        return nil
    }
    res["slug"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSlug(val)
        }
        return nil
    }
    res["starter_code_repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleClassroomRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStarterCodeRepository(val.(SimpleClassroomRepositoryable))
        }
        return nil
    }
    res["students_are_repo_admins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudentsAreRepoAdmins(val)
        }
        return nil
    }
    res["submitted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmitted(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClassroomAssignment_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*ClassroomAssignment_type))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique identifier of the repository.
// returns a *int32 when successful
func (m *ClassroomAssignment) GetId()(*int32) {
    return m.id
}
// GetInvitationsEnabled gets the invitations_enabled property value. Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
// returns a *bool when successful
func (m *ClassroomAssignment) GetInvitationsEnabled()(*bool) {
    return m.invitations_enabled
}
// GetInviteLink gets the invite_link property value. The link that a student can use to accept the assignment.
// returns a *string when successful
func (m *ClassroomAssignment) GetInviteLink()(*string) {
    return m.invite_link
}
// GetLanguage gets the language property value. The programming language used in the assignment.
// returns a *string when successful
func (m *ClassroomAssignment) GetLanguage()(*string) {
    return m.language
}
// GetMaxMembers gets the max_members property value. The maximum allowable members per team.
// returns a *int32 when successful
func (m *ClassroomAssignment) GetMaxMembers()(*int32) {
    return m.max_members
}
// GetMaxTeams gets the max_teams property value. The maximum allowable teams for the assignment.
// returns a *int32 when successful
func (m *ClassroomAssignment) GetMaxTeams()(*int32) {
    return m.max_teams
}
// GetPassing gets the passing property value. The number of students that have passed the assignment.
// returns a *int32 when successful
func (m *ClassroomAssignment) GetPassing()(*int32) {
    return m.passing
}
// GetPublicRepo gets the public_repo property value. Whether an accepted assignment creates a public repository.
// returns a *bool when successful
func (m *ClassroomAssignment) GetPublicRepo()(*bool) {
    return m.public_repo
}
// GetSlug gets the slug property value. Sluggified name of the assignment.
// returns a *string when successful
func (m *ClassroomAssignment) GetSlug()(*string) {
    return m.slug
}
// GetStarterCodeRepository gets the starter_code_repository property value. A GitHub repository view for Classroom
// returns a SimpleClassroomRepositoryable when successful
func (m *ClassroomAssignment) GetStarterCodeRepository()(SimpleClassroomRepositoryable) {
    return m.starter_code_repository
}
// GetStudentsAreRepoAdmins gets the students_are_repo_admins property value. Whether students are admins on created repository when a student accepts the assignment.
// returns a *bool when successful
func (m *ClassroomAssignment) GetStudentsAreRepoAdmins()(*bool) {
    return m.students_are_repo_admins
}
// GetSubmitted gets the submitted property value. The number of students that have submitted the assignment.
// returns a *int32 when successful
func (m *ClassroomAssignment) GetSubmitted()(*int32) {
    return m.submitted
}
// GetTitle gets the title property value. Assignment title.
// returns a *string when successful
func (m *ClassroomAssignment) GetTitle()(*string) {
    return m.title
}
// GetTypeEscaped gets the type property value. Whether it's a group assignment or individual assignment.
// returns a *ClassroomAssignment_type when successful
func (m *ClassroomAssignment) GetTypeEscaped()(*ClassroomAssignment_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ClassroomAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("accepted", m.GetAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("classroom", m.GetClassroom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("deadline", m.GetDeadline())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("editor", m.GetEditor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("feedback_pull_requests_enabled", m.GetFeedbackPullRequestsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("invitations_enabled", m.GetInvitationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("invite_link", m.GetInviteLink())
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
        err := writer.WriteInt32Value("max_members", m.GetMaxMembers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("max_teams", m.GetMaxTeams())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("passing", m.GetPassing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("public_repo", m.GetPublicRepo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("slug", m.GetSlug())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("starter_code_repository", m.GetStarterCodeRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("students_are_repo_admins", m.GetStudentsAreRepoAdmins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("submitted", m.GetSubmitted())
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
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
// SetAccepted sets the accepted property value. The number of students that have accepted the assignment.
func (m *ClassroomAssignment) SetAccepted(value *int32)() {
    m.accepted = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassroomAssignment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClassroom sets the classroom property value. A GitHub Classroom classroom
func (m *ClassroomAssignment) SetClassroom(value Classroomable)() {
    m.classroom = value
}
// SetDeadline sets the deadline property value. The time at which the assignment is due.
func (m *ClassroomAssignment) SetDeadline(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deadline = value
}
// SetEditor sets the editor property value. The selected editor for the assignment.
func (m *ClassroomAssignment) SetEditor(value *string)() {
    m.editor = value
}
// SetFeedbackPullRequestsEnabled sets the feedback_pull_requests_enabled property value. Whether feedback pull request will be created when a student accepts the assignment.
func (m *ClassroomAssignment) SetFeedbackPullRequestsEnabled(value *bool)() {
    m.feedback_pull_requests_enabled = value
}
// SetId sets the id property value. Unique identifier of the repository.
func (m *ClassroomAssignment) SetId(value *int32)() {
    m.id = value
}
// SetInvitationsEnabled sets the invitations_enabled property value. Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
func (m *ClassroomAssignment) SetInvitationsEnabled(value *bool)() {
    m.invitations_enabled = value
}
// SetInviteLink sets the invite_link property value. The link that a student can use to accept the assignment.
func (m *ClassroomAssignment) SetInviteLink(value *string)() {
    m.invite_link = value
}
// SetLanguage sets the language property value. The programming language used in the assignment.
func (m *ClassroomAssignment) SetLanguage(value *string)() {
    m.language = value
}
// SetMaxMembers sets the max_members property value. The maximum allowable members per team.
func (m *ClassroomAssignment) SetMaxMembers(value *int32)() {
    m.max_members = value
}
// SetMaxTeams sets the max_teams property value. The maximum allowable teams for the assignment.
func (m *ClassroomAssignment) SetMaxTeams(value *int32)() {
    m.max_teams = value
}
// SetPassing sets the passing property value. The number of students that have passed the assignment.
func (m *ClassroomAssignment) SetPassing(value *int32)() {
    m.passing = value
}
// SetPublicRepo sets the public_repo property value. Whether an accepted assignment creates a public repository.
func (m *ClassroomAssignment) SetPublicRepo(value *bool)() {
    m.public_repo = value
}
// SetSlug sets the slug property value. Sluggified name of the assignment.
func (m *ClassroomAssignment) SetSlug(value *string)() {
    m.slug = value
}
// SetStarterCodeRepository sets the starter_code_repository property value. A GitHub repository view for Classroom
func (m *ClassroomAssignment) SetStarterCodeRepository(value SimpleClassroomRepositoryable)() {
    m.starter_code_repository = value
}
// SetStudentsAreRepoAdmins sets the students_are_repo_admins property value. Whether students are admins on created repository when a student accepts the assignment.
func (m *ClassroomAssignment) SetStudentsAreRepoAdmins(value *bool)() {
    m.students_are_repo_admins = value
}
// SetSubmitted sets the submitted property value. The number of students that have submitted the assignment.
func (m *ClassroomAssignment) SetSubmitted(value *int32)() {
    m.submitted = value
}
// SetTitle sets the title property value. Assignment title.
func (m *ClassroomAssignment) SetTitle(value *string)() {
    m.title = value
}
// SetTypeEscaped sets the type property value. Whether it's a group assignment or individual assignment.
func (m *ClassroomAssignment) SetTypeEscaped(value *ClassroomAssignment_type)() {
    m.typeEscaped = value
}
type ClassroomAssignmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccepted()(*int32)
    GetClassroom()(Classroomable)
    GetDeadline()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEditor()(*string)
    GetFeedbackPullRequestsEnabled()(*bool)
    GetId()(*int32)
    GetInvitationsEnabled()(*bool)
    GetInviteLink()(*string)
    GetLanguage()(*string)
    GetMaxMembers()(*int32)
    GetMaxTeams()(*int32)
    GetPassing()(*int32)
    GetPublicRepo()(*bool)
    GetSlug()(*string)
    GetStarterCodeRepository()(SimpleClassroomRepositoryable)
    GetStudentsAreRepoAdmins()(*bool)
    GetSubmitted()(*int32)
    GetTitle()(*string)
    GetTypeEscaped()(*ClassroomAssignment_type)
    SetAccepted(value *int32)()
    SetClassroom(value Classroomable)()
    SetDeadline(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEditor(value *string)()
    SetFeedbackPullRequestsEnabled(value *bool)()
    SetId(value *int32)()
    SetInvitationsEnabled(value *bool)()
    SetInviteLink(value *string)()
    SetLanguage(value *string)()
    SetMaxMembers(value *int32)()
    SetMaxTeams(value *int32)()
    SetPassing(value *int32)()
    SetPublicRepo(value *bool)()
    SetSlug(value *string)()
    SetStarterCodeRepository(value SimpleClassroomRepositoryable)()
    SetStudentsAreRepoAdmins(value *bool)()
    SetSubmitted(value *int32)()
    SetTitle(value *string)()
    SetTypeEscaped(value *ClassroomAssignment_type)()
}
