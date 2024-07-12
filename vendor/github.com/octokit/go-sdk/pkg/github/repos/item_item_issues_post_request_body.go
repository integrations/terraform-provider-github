package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemIssuesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Login for the user that this issue should be assigned to. _NOTE: Only users with push access can set the assignee for new issues. The assignee is silently dropped otherwise. **This field is deprecated.**_
    assignee *string
    // Logins for Users to assign to this issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._
    assignees []string
    // The contents of the issue.
    body *string
    // Labels to associate with this issue. _NOTE: Only users with push access can set labels for new issues. Labels are silently dropped otherwise._
    labels []string
    // The milestone property
    milestone ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneable
    // The title of the issue.
    title ItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleable
}
// ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone composed type wrapper for classes int32, string
type ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone struct {
    // Composed type representation for type int32
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone instantiates a new ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone and sets the default values.
func NewItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone()(*ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone) {
    m := &ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone{
    }
    return m
}
// CreateItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetInt32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInteger(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type int32
// returns a *int32 when successful
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteger() != nil {
        err := writer.WriteInt32Value("", m.GetInteger())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteger sets the integer property value. Composed type representation for type int32
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestone) SetString(value *string)() {
    m.string = value
}
// ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title composed type wrapper for classes int32, string
type ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title struct {
    // Composed type representation for type int32
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewItemItemIssuesPostRequestBody_IssuesPostRequestBody_title instantiates a new ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title and sets the default values.
func NewItemItemIssuesPostRequestBody_IssuesPostRequestBody_title()(*ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title) {
    m := &ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title{
    }
    return m
}
// CreateItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemItemIssuesPostRequestBody_IssuesPostRequestBody_title()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetInt32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInteger(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type int32
// returns a *int32 when successful
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteger() != nil {
        err := writer.WriteInt32Value("", m.GetInteger())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteger sets the integer property value. Composed type representation for type int32
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemItemIssuesPostRequestBody_IssuesPostRequestBody_title) SetString(value *string)() {
    m.string = value
}
type ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
type ItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewItemItemIssuesPostRequestBody instantiates a new ItemItemIssuesPostRequestBody and sets the default values.
func NewItemItemIssuesPostRequestBody()(*ItemItemIssuesPostRequestBody) {
    m := &ItemItemIssuesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemIssuesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemIssuesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignee gets the assignee property value. Login for the user that this issue should be assigned to. _NOTE: Only users with push access can set the assignee for new issues. The assignee is silently dropped otherwise. **This field is deprecated.**_
// returns a *string when successful
func (m *ItemItemIssuesPostRequestBody) GetAssignee()(*string) {
    return m.assignee
}
// GetAssignees gets the assignees property value. Logins for Users to assign to this issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._
// returns a []string when successful
func (m *ItemItemIssuesPostRequestBody) GetAssignees()([]string) {
    return m.assignees
}
// GetBody gets the body property value. The contents of the issue.
// returns a *string when successful
func (m *ItemItemIssuesPostRequestBody) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemIssuesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignee(val)
        }
        return nil
    }
    res["assignees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAssignees(res)
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
    res["milestone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestone(val.(ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(ItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleable))
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. Labels to associate with this issue. _NOTE: Only users with push access can set labels for new issues. Labels are silently dropped otherwise._
// returns a []string when successful
func (m *ItemItemIssuesPostRequestBody) GetLabels()([]string) {
    return m.labels
}
// GetMilestone gets the milestone property value. The milestone property
// returns a ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneable when successful
func (m *ItemItemIssuesPostRequestBody) GetMilestone()(ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneable) {
    return m.milestone
}
// GetTitle gets the title property value. The title of the issue.
// returns a ItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleable when successful
func (m *ItemItemIssuesPostRequestBody) GetTitle()(ItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleable) {
    return m.title
}
// Serialize serializes information the current object
func (m *ItemItemIssuesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignee", m.GetAssignee())
        if err != nil {
            return err
        }
    }
    if m.GetAssignees() != nil {
        err := writer.WriteCollectionOfStringValues("assignees", m.GetAssignees())
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
    if m.GetLabels() != nil {
        err := writer.WriteCollectionOfStringValues("labels", m.GetLabels())
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
        err := writer.WriteObjectValue("title", m.GetTitle())
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
func (m *ItemItemIssuesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignee sets the assignee property value. Login for the user that this issue should be assigned to. _NOTE: Only users with push access can set the assignee for new issues. The assignee is silently dropped otherwise. **This field is deprecated.**_
func (m *ItemItemIssuesPostRequestBody) SetAssignee(value *string)() {
    m.assignee = value
}
// SetAssignees sets the assignees property value. Logins for Users to assign to this issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._
func (m *ItemItemIssuesPostRequestBody) SetAssignees(value []string)() {
    m.assignees = value
}
// SetBody sets the body property value. The contents of the issue.
func (m *ItemItemIssuesPostRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetLabels sets the labels property value. Labels to associate with this issue. _NOTE: Only users with push access can set labels for new issues. Labels are silently dropped otherwise._
func (m *ItemItemIssuesPostRequestBody) SetLabels(value []string)() {
    m.labels = value
}
// SetMilestone sets the milestone property value. The milestone property
func (m *ItemItemIssuesPostRequestBody) SetMilestone(value ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneable)() {
    m.milestone = value
}
// SetTitle sets the title property value. The title of the issue.
func (m *ItemItemIssuesPostRequestBody) SetTitle(value ItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleable)() {
    m.title = value
}
type ItemItemIssuesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignee()(*string)
    GetAssignees()([]string)
    GetBody()(*string)
    GetLabels()([]string)
    GetMilestone()(ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneable)
    GetTitle()(ItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleable)
    SetAssignee(value *string)()
    SetAssignees(value []string)()
    SetBody(value *string)()
    SetLabels(value []string)()
    SetMilestone(value ItemItemIssuesPostRequestBody_IssuesPostRequestBody_milestoneable)()
    SetTitle(value ItemItemIssuesPostRequestBody_IssuesPostRequestBody_titleable)()
}
