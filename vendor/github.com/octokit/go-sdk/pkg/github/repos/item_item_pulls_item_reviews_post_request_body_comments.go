package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsItemReviewsPostRequestBody_comments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Text of the review comment.
    body *string
    // The line property
    line *int32
    // The relative path to the file that necessitates a review comment.
    path *string
    // The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. The `position` value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.
    position *int32
    // The side property
    side *string
    // The start_line property
    start_line *int32
    // The start_side property
    start_side *string
}
// NewItemItemPullsItemReviewsPostRequestBody_comments instantiates a new ItemItemPullsItemReviewsPostRequestBody_comments and sets the default values.
func NewItemItemPullsItemReviewsPostRequestBody_comments()(*ItemItemPullsItemReviewsPostRequestBody_comments) {
    m := &ItemItemPullsItemReviewsPostRequestBody_comments{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemReviewsPostRequestBody_commentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsItemReviewsPostRequestBody_commentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemReviewsPostRequestBody_comments(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. Text of the review comment.
// returns a *string when successful
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["side"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSide(val)
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
    res["start_side"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartSide(val)
        }
        return nil
    }
    return res
}
// GetLine gets the line property value. The line property
// returns a *int32 when successful
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) GetLine()(*int32) {
    return m.line
}
// GetPath gets the path property value. The relative path to the file that necessitates a review comment.
// returns a *string when successful
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) GetPath()(*string) {
    return m.path
}
// GetPosition gets the position property value. The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. The `position` value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.
// returns a *int32 when successful
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) GetPosition()(*int32) {
    return m.position
}
// GetSide gets the side property value. The side property
// returns a *string when successful
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) GetSide()(*string) {
    return m.side
}
// GetStartLine gets the start_line property value. The start_line property
// returns a *int32 when successful
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) GetStartLine()(*int32) {
    return m.start_line
}
// GetStartSide gets the start_side property value. The start_side property
// returns a *string when successful
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) GetStartSide()(*string) {
    return m.start_side
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
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
        err := writer.WriteStringValue("side", m.GetSide())
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
        err := writer.WriteStringValue("start_side", m.GetStartSide())
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
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. Text of the review comment.
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) SetBody(value *string)() {
    m.body = value
}
// SetLine sets the line property value. The line property
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) SetLine(value *int32)() {
    m.line = value
}
// SetPath sets the path property value. The relative path to the file that necessitates a review comment.
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) SetPath(value *string)() {
    m.path = value
}
// SetPosition sets the position property value. The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. The `position` value equals the number of lines down from the first "@@" hunk header in the file you want to add a comment. The line just below the "@@" line is position 1, the next line is position 2, and so on. The position in the diff continues to increase through lines of whitespace and additional hunks until the beginning of a new file.
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) SetPosition(value *int32)() {
    m.position = value
}
// SetSide sets the side property value. The side property
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) SetSide(value *string)() {
    m.side = value
}
// SetStartLine sets the start_line property value. The start_line property
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) SetStartLine(value *int32)() {
    m.start_line = value
}
// SetStartSide sets the start_side property value. The start_side property
func (m *ItemItemPullsItemReviewsPostRequestBody_comments) SetStartSide(value *string)() {
    m.start_side = value
}
type ItemItemPullsItemReviewsPostRequestBody_commentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetLine()(*int32)
    GetPath()(*string)
    GetPosition()(*int32)
    GetSide()(*string)
    GetStartLine()(*int32)
    GetStartSide()(*string)
    SetBody(value *string)()
    SetLine(value *int32)()
    SetPath(value *string)()
    SetPosition(value *int32)()
    SetSide(value *string)()
    SetStartLine(value *int32)()
    SetStartSide(value *string)()
}
