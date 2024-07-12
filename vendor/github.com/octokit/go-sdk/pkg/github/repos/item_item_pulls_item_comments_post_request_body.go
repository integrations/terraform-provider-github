package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsItemCommentsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The text of the review comment.
    body *string
    // The SHA of the commit needing a comment. Not using the latest commit SHA may render your comment outdated if a subsequent commit modifies the line you specify as the `position`.
    commit_id *string
    // The ID of the review comment to reply to. To find the ID of a review comment with ["List review comments on a pull request"](#list-review-comments-on-a-pull-request). When specified, all parameters other than `body` in the request body are ignored.
    in_reply_to *int32
    // **Required unless using `subject_type:file`**. The line of the blob in the pull request diff that the comment applies to. For a multi-line comment, the last line of the range that your comment applies to.
    line *int32
    // The relative path to the file that necessitates a comment.
    path *string
    // **This parameter is deprecated. Use `line` instead**. The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. The position value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.
    // Deprecated: 
    position *int32
    // **Required when using multi-line comments unless using `in_reply_to`**. The `start_line` is the first line in the pull request diff that your multi-line comment applies to. To learn more about multi-line comments, see "[Commenting on a pull request](https://docs.github.com/articles/commenting-on-a-pull-request#adding-line-comments-to-a-pull-request)" in the GitHub Help documentation.
    start_line *int32
}
// NewItemItemPullsItemCommentsPostRequestBody instantiates a new ItemItemPullsItemCommentsPostRequestBody and sets the default values.
func NewItemItemPullsItemCommentsPostRequestBody()(*ItemItemPullsItemCommentsPostRequestBody) {
    m := &ItemItemPullsItemCommentsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemCommentsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsItemCommentsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemCommentsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsItemCommentsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The text of the review comment.
// returns a *string when successful
func (m *ItemItemPullsItemCommentsPostRequestBody) GetBody()(*string) {
    return m.body
}
// GetCommitId gets the commit_id property value. The SHA of the commit needing a comment. Not using the latest commit SHA may render your comment outdated if a subsequent commit modifies the line you specify as the `position`.
// returns a *string when successful
func (m *ItemItemPullsItemCommentsPostRequestBody) GetCommitId()(*string) {
    return m.commit_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsItemCommentsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["in_reply_to"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInReplyTo(val)
        }
        return nil
    }
    res["line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLine(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    res["start_line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartLine(val)
        }
        return nil
    }
    return res
}
// GetInReplyTo gets the in_reply_to property value. The ID of the review comment to reply to. To find the ID of a review comment with ["List review comments on a pull request"](#list-review-comments-on-a-pull-request). When specified, all parameters other than `body` in the request body are ignored.
// returns a *int32 when successful
func (m *ItemItemPullsItemCommentsPostRequestBody) GetInReplyTo()(*int32) {
    return m.in_reply_to
}
// GetLine gets the line property value. **Required unless using `subject_type:file`**. The line of the blob in the pull request diff that the comment applies to. For a multi-line comment, the last line of the range that your comment applies to.
// returns a *int32 when successful
func (m *ItemItemPullsItemCommentsPostRequestBody) GetLine()(*int32) {
    return m.line
}
// GetPath gets the path property value. The relative path to the file that necessitates a comment.
// returns a *string when successful
func (m *ItemItemPullsItemCommentsPostRequestBody) GetPath()(*string) {
    return m.path
}
// GetPosition gets the position property value. **This parameter is deprecated. Use `line` instead**. The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. The position value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.
// Deprecated: 
// returns a *int32 when successful
func (m *ItemItemPullsItemCommentsPostRequestBody) GetPosition()(*int32) {
    return m.position
}
// GetStartLine gets the start_line property value. **Required when using multi-line comments unless using `in_reply_to`**. The `start_line` is the first line in the pull request diff that your multi-line comment applies to. To learn more about multi-line comments, see "[Commenting on a pull request](https://docs.github.com/articles/commenting-on-a-pull-request#adding-line-comments-to-a-pull-request)" in the GitHub Help documentation.
// returns a *int32 when successful
func (m *ItemItemPullsItemCommentsPostRequestBody) GetStartLine()(*int32) {
    return m.start_line
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemCommentsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
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
        err := writer.WriteInt32Value("in_reply_to", m.GetInReplyTo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("line", m.GetLine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("start_line", m.GetStartLine())
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
func (m *ItemItemPullsItemCommentsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The text of the review comment.
func (m *ItemItemPullsItemCommentsPostRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetCommitId sets the commit_id property value. The SHA of the commit needing a comment. Not using the latest commit SHA may render your comment outdated if a subsequent commit modifies the line you specify as the `position`.
func (m *ItemItemPullsItemCommentsPostRequestBody) SetCommitId(value *string)() {
    m.commit_id = value
}
// SetInReplyTo sets the in_reply_to property value. The ID of the review comment to reply to. To find the ID of a review comment with ["List review comments on a pull request"](#list-review-comments-on-a-pull-request). When specified, all parameters other than `body` in the request body are ignored.
func (m *ItemItemPullsItemCommentsPostRequestBody) SetInReplyTo(value *int32)() {
    m.in_reply_to = value
}
// SetLine sets the line property value. **Required unless using `subject_type:file`**. The line of the blob in the pull request diff that the comment applies to. For a multi-line comment, the last line of the range that your comment applies to.
func (m *ItemItemPullsItemCommentsPostRequestBody) SetLine(value *int32)() {
    m.line = value
}
// SetPath sets the path property value. The relative path to the file that necessitates a comment.
func (m *ItemItemPullsItemCommentsPostRequestBody) SetPath(value *string)() {
    m.path = value
}
// SetPosition sets the position property value. **This parameter is deprecated. Use `line` instead**. The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. The position value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.
// Deprecated: 
func (m *ItemItemPullsItemCommentsPostRequestBody) SetPosition(value *int32)() {
    m.position = value
}
// SetStartLine sets the start_line property value. **Required when using multi-line comments unless using `in_reply_to`**. The `start_line` is the first line in the pull request diff that your multi-line comment applies to. To learn more about multi-line comments, see "[Commenting on a pull request](https://docs.github.com/articles/commenting-on-a-pull-request#adding-line-comments-to-a-pull-request)" in the GitHub Help documentation.
func (m *ItemItemPullsItemCommentsPostRequestBody) SetStartLine(value *int32)() {
    m.start_line = value
}
type ItemItemPullsItemCommentsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetCommitId()(*string)
    GetInReplyTo()(*int32)
    GetLine()(*int32)
    GetPath()(*string)
    GetPosition()(*int32)
    GetStartLine()(*int32)
    SetBody(value *string)()
    SetCommitId(value *string)()
    SetInReplyTo(value *int32)()
    SetLine(value *int32)()
    SetPath(value *string)()
    SetPosition(value *int32)()
    SetStartLine(value *int32)()
}
