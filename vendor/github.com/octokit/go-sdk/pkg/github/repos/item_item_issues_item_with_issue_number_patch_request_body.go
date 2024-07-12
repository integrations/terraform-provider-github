package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemIssuesItemWithIssue_numberPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Username to assign to this issue. **This field is deprecated.**
    assignee *string
    // Usernames to assign to this issue. Pass one or more user logins to _replace_ the set of assignees on this issue. Send an empty array (`[]`) to clear all assignees from the issue. Only users with push access can set assignees for new issues. Without push access to the repository, assignee changes are silently dropped.
    assignees []string
    // The contents of the issue.
    body *string
    // Labels to associate with this issue. Pass one or more labels to _replace_ the set of labels on this issue. Send an empty array (`[]`) to clear all labels from the issue. Only users with push access can set labels for issues. Without push access to the repository, label changes are silently dropped.
    labels []string
    // The milestone property
    milestone ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneable
    // The title of the issue.
    title ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleable
}
// ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone composed type wrapper for classes int32, string
type ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone struct {
    // Composed type representation for type int32
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone instantiates a new ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone and sets the default values.
func NewItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone()(*ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone) {
    m := &ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone{
    }
    return m
}
// CreateItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone()
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
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type int32
// returns a *int32 when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestone) SetString(value *string)() {
    m.string = value
}
// ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title composed type wrapper for classes int32, string
type ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title struct {
    // Composed type representation for type int32
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title instantiates a new ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title and sets the default values.
func NewItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title()(*ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title) {
    m := &ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title{
    }
    return m
}
// CreateItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title()
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
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type int32
// returns a *int32 when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_title) SetString(value *string)() {
    m.string = value
}
type ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
type ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewItemItemIssuesItemWithIssue_numberPatchRequestBody instantiates a new ItemItemIssuesItemWithIssue_numberPatchRequestBody and sets the default values.
func NewItemItemIssuesItemWithIssue_numberPatchRequestBody()(*ItemItemIssuesItemWithIssue_numberPatchRequestBody) {
    m := &ItemItemIssuesItemWithIssue_numberPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesItemWithIssue_numberPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesItemWithIssue_numberPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemIssuesItemWithIssue_numberPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignee gets the assignee property value. Username to assign to this issue. **This field is deprecated.**
// returns a *string when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) GetAssignee()(*string) {
    return m.assignee
}
// GetAssignees gets the assignees property value. Usernames to assign to this issue. Pass one or more user logins to _replace_ the set of assignees on this issue. Send an empty array (`[]`) to clear all assignees from the issue. Only users with push access can set assignees for new issues. Without push access to the repository, assignee changes are silently dropped.
// returns a []string when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) GetAssignees()([]string) {
    return m.assignees
}
// GetBody gets the body property value. The contents of the issue.
// returns a *string when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestone(val.(ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleable))
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. Labels to associate with this issue. Pass one or more labels to _replace_ the set of labels on this issue. Send an empty array (`[]`) to clear all labels from the issue. Only users with push access can set labels for issues. Without push access to the repository, label changes are silently dropped.
// returns a []string when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) GetLabels()([]string) {
    return m.labels
}
// GetMilestone gets the milestone property value. The milestone property
// returns a ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneable when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) GetMilestone()(ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneable) {
    return m.milestone
}
// GetTitle gets the title property value. The title of the issue.
// returns a ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleable when successful
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) GetTitle()(ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleable) {
    return m.title
}
// Serialize serializes information the current object
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignee sets the assignee property value. Username to assign to this issue. **This field is deprecated.**
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) SetAssignee(value *string)() {
    m.assignee = value
}
// SetAssignees sets the assignees property value. Usernames to assign to this issue. Pass one or more user logins to _replace_ the set of assignees on this issue. Send an empty array (`[]`) to clear all assignees from the issue. Only users with push access can set assignees for new issues. Without push access to the repository, assignee changes are silently dropped.
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) SetAssignees(value []string)() {
    m.assignees = value
}
// SetBody sets the body property value. The contents of the issue.
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetLabels sets the labels property value. Labels to associate with this issue. Pass one or more labels to _replace_ the set of labels on this issue. Send an empty array (`[]`) to clear all labels from the issue. Only users with push access can set labels for issues. Without push access to the repository, label changes are silently dropped.
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) SetLabels(value []string)() {
    m.labels = value
}
// SetMilestone sets the milestone property value. The milestone property
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) SetMilestone(value ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneable)() {
    m.milestone = value
}
// SetTitle sets the title property value. The title of the issue.
func (m *ItemItemIssuesItemWithIssue_numberPatchRequestBody) SetTitle(value ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleable)() {
    m.title = value
}
type ItemItemIssuesItemWithIssue_numberPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignee()(*string)
    GetAssignees()([]string)
    GetBody()(*string)
    GetLabels()([]string)
    GetMilestone()(ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneable)
    GetTitle()(ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleable)
    SetAssignee(value *string)()
    SetAssignees(value []string)()
    SetBody(value *string)()
    SetLabels(value []string)()
    SetMilestone(value ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_milestoneable)()
    SetTitle(value ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberPatchRequestBody_titleable)()
}
