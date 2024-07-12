package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassroomAssignmentGrade grade for a student or groups GitHub Classroom assignment
type ClassroomAssignmentGrade struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the assignment
    assignment_name *string
    // URL of the assignment
    assignment_url *string
    // GitHub username of the student
    github_username *string
    // If a group assignment, name of the group the student is in
    group_name *string
    // Number of points available for the assignment
    points_available *int32
    // Number of points awarded to the student
    points_awarded *int32
    // Roster identifier of the student
    roster_identifier *string
    // URL of the starter code for the assignment
    starter_code_url *string
    // Name of the student's assignment repository
    student_repository_name *string
    // URL of the student's assignment repository
    student_repository_url *string
    // Timestamp of the student's assignment submission
    submission_timestamp *string
}
// NewClassroomAssignmentGrade instantiates a new ClassroomAssignmentGrade and sets the default values.
func NewClassroomAssignmentGrade()(*ClassroomAssignmentGrade) {
    m := &ClassroomAssignmentGrade{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClassroomAssignmentGradeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClassroomAssignmentGradeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassroomAssignmentGrade(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ClassroomAssignmentGrade) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignmentName gets the assignment_name property value. Name of the assignment
// returns a *string when successful
func (m *ClassroomAssignmentGrade) GetAssignmentName()(*string) {
    return m.assignment_name
}
// GetAssignmentUrl gets the assignment_url property value. URL of the assignment
// returns a *string when successful
func (m *ClassroomAssignmentGrade) GetAssignmentUrl()(*string) {
    return m.assignment_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClassroomAssignmentGrade) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignment_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentName(val)
        }
        return nil
    }
    res["assignment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentUrl(val)
        }
        return nil
    }
    res["github_username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubUsername(val)
        }
        return nil
    }
    res["group_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    res["points_available"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPointsAvailable(val)
        }
        return nil
    }
    res["points_awarded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPointsAwarded(val)
        }
        return nil
    }
    res["roster_identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRosterIdentifier(val)
        }
        return nil
    }
    res["starter_code_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStarterCodeUrl(val)
        }
        return nil
    }
    res["student_repository_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudentRepositoryName(val)
        }
        return nil
    }
    res["student_repository_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudentRepositoryUrl(val)
        }
        return nil
    }
    res["submission_timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmissionTimestamp(val)
        }
        return nil
    }
    return res
}
// GetGithubUsername gets the github_username property value. GitHub username of the student
// returns a *string when successful
func (m *ClassroomAssignmentGrade) GetGithubUsername()(*string) {
    return m.github_username
}
// GetGroupName gets the group_name property value. If a group assignment, name of the group the student is in
// returns a *string when successful
func (m *ClassroomAssignmentGrade) GetGroupName()(*string) {
    return m.group_name
}
// GetPointsAvailable gets the points_available property value. Number of points available for the assignment
// returns a *int32 when successful
func (m *ClassroomAssignmentGrade) GetPointsAvailable()(*int32) {
    return m.points_available
}
// GetPointsAwarded gets the points_awarded property value. Number of points awarded to the student
// returns a *int32 when successful
func (m *ClassroomAssignmentGrade) GetPointsAwarded()(*int32) {
    return m.points_awarded
}
// GetRosterIdentifier gets the roster_identifier property value. Roster identifier of the student
// returns a *string when successful
func (m *ClassroomAssignmentGrade) GetRosterIdentifier()(*string) {
    return m.roster_identifier
}
// GetStarterCodeUrl gets the starter_code_url property value. URL of the starter code for the assignment
// returns a *string when successful
func (m *ClassroomAssignmentGrade) GetStarterCodeUrl()(*string) {
    return m.starter_code_url
}
// GetStudentRepositoryName gets the student_repository_name property value. Name of the student's assignment repository
// returns a *string when successful
func (m *ClassroomAssignmentGrade) GetStudentRepositoryName()(*string) {
    return m.student_repository_name
}
// GetStudentRepositoryUrl gets the student_repository_url property value. URL of the student's assignment repository
// returns a *string when successful
func (m *ClassroomAssignmentGrade) GetStudentRepositoryUrl()(*string) {
    return m.student_repository_url
}
// GetSubmissionTimestamp gets the submission_timestamp property value. Timestamp of the student's assignment submission
// returns a *string when successful
func (m *ClassroomAssignmentGrade) GetSubmissionTimestamp()(*string) {
    return m.submission_timestamp
}
// Serialize serializes information the current object
func (m *ClassroomAssignmentGrade) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignment_name", m.GetAssignmentName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assignment_url", m.GetAssignmentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("github_username", m.GetGithubUsername())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("group_name", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("points_available", m.GetPointsAvailable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("points_awarded", m.GetPointsAwarded())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roster_identifier", m.GetRosterIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("starter_code_url", m.GetStarterCodeUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("student_repository_name", m.GetStudentRepositoryName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("student_repository_url", m.GetStudentRepositoryUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("submission_timestamp", m.GetSubmissionTimestamp())
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
func (m *ClassroomAssignmentGrade) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignmentName sets the assignment_name property value. Name of the assignment
func (m *ClassroomAssignmentGrade) SetAssignmentName(value *string)() {
    m.assignment_name = value
}
// SetAssignmentUrl sets the assignment_url property value. URL of the assignment
func (m *ClassroomAssignmentGrade) SetAssignmentUrl(value *string)() {
    m.assignment_url = value
}
// SetGithubUsername sets the github_username property value. GitHub username of the student
func (m *ClassroomAssignmentGrade) SetGithubUsername(value *string)() {
    m.github_username = value
}
// SetGroupName sets the group_name property value. If a group assignment, name of the group the student is in
func (m *ClassroomAssignmentGrade) SetGroupName(value *string)() {
    m.group_name = value
}
// SetPointsAvailable sets the points_available property value. Number of points available for the assignment
func (m *ClassroomAssignmentGrade) SetPointsAvailable(value *int32)() {
    m.points_available = value
}
// SetPointsAwarded sets the points_awarded property value. Number of points awarded to the student
func (m *ClassroomAssignmentGrade) SetPointsAwarded(value *int32)() {
    m.points_awarded = value
}
// SetRosterIdentifier sets the roster_identifier property value. Roster identifier of the student
func (m *ClassroomAssignmentGrade) SetRosterIdentifier(value *string)() {
    m.roster_identifier = value
}
// SetStarterCodeUrl sets the starter_code_url property value. URL of the starter code for the assignment
func (m *ClassroomAssignmentGrade) SetStarterCodeUrl(value *string)() {
    m.starter_code_url = value
}
// SetStudentRepositoryName sets the student_repository_name property value. Name of the student's assignment repository
func (m *ClassroomAssignmentGrade) SetStudentRepositoryName(value *string)() {
    m.student_repository_name = value
}
// SetStudentRepositoryUrl sets the student_repository_url property value. URL of the student's assignment repository
func (m *ClassroomAssignmentGrade) SetStudentRepositoryUrl(value *string)() {
    m.student_repository_url = value
}
// SetSubmissionTimestamp sets the submission_timestamp property value. Timestamp of the student's assignment submission
func (m *ClassroomAssignmentGrade) SetSubmissionTimestamp(value *string)() {
    m.submission_timestamp = value
}
type ClassroomAssignmentGradeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentName()(*string)
    GetAssignmentUrl()(*string)
    GetGithubUsername()(*string)
    GetGroupName()(*string)
    GetPointsAvailable()(*int32)
    GetPointsAwarded()(*int32)
    GetRosterIdentifier()(*string)
    GetStarterCodeUrl()(*string)
    GetStudentRepositoryName()(*string)
    GetStudentRepositoryUrl()(*string)
    GetSubmissionTimestamp()(*string)
    SetAssignmentName(value *string)()
    SetAssignmentUrl(value *string)()
    SetGithubUsername(value *string)()
    SetGroupName(value *string)()
    SetPointsAvailable(value *int32)()
    SetPointsAwarded(value *int32)()
    SetRosterIdentifier(value *string)()
    SetStarterCodeUrl(value *string)()
    SetStudentRepositoryName(value *string)()
    SetStudentRepositoryUrl(value *string)()
    SetSubmissionTimestamp(value *string)()
}
