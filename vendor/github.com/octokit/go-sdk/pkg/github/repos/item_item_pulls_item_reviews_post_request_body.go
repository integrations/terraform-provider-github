package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsItemReviewsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // **Required** when using `REQUEST_CHANGES` or `COMMENT` for the `event` parameter. The body text of the pull request review.
    body *string
    // Use the following table to specify the location, destination, and contents of the draft review comment.
    comments []ItemItemPullsItemReviewsPostRequestBody_commentsable
    // The SHA of the commit that needs a review. Not using the latest commit SHA may render your review comment outdated if a subsequent commit modifies the line you specify as the `position`. Defaults to the most recent commit in the pull request when you do not specify a value.
    commit_id *string
}
// NewItemItemPullsItemReviewsPostRequestBody instantiates a new ItemItemPullsItemReviewsPostRequestBody and sets the default values.
func NewItemItemPullsItemReviewsPostRequestBody()(*ItemItemPullsItemReviewsPostRequestBody) {
    m := &ItemItemPullsItemReviewsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemReviewsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsItemReviewsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemReviewsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsItemReviewsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. **Required** when using `REQUEST_CHANGES` or `COMMENT` for the `event` parameter. The body text of the pull request review.
// returns a *string when successful
func (m *ItemItemPullsItemReviewsPostRequestBody) GetBody()(*string) {
    return m.body
}
// GetComments gets the comments property value. Use the following table to specify the location, destination, and contents of the draft review comment.
// returns a []ItemItemPullsItemReviewsPostRequestBody_commentsable when successful
func (m *ItemItemPullsItemReviewsPostRequestBody) GetComments()([]ItemItemPullsItemReviewsPostRequestBody_commentsable) {
    return m.comments
}
// GetCommitId gets the commit_id property value. The SHA of the commit that needs a review. Not using the latest commit SHA may render your review comment outdated if a subsequent commit modifies the line you specify as the `position`. Defaults to the most recent commit in the pull request when you do not specify a value.
// returns a *string when successful
func (m *ItemItemPullsItemReviewsPostRequestBody) GetCommitId()(*string) {
    return m.commit_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsItemReviewsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemItemPullsItemReviewsPostRequestBody_commentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemItemPullsItemReviewsPostRequestBody_commentsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemItemPullsItemReviewsPostRequestBody_commentsable)
                }
            }
            m.SetComments(res)
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
    return res
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemReviewsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    if m.GetComments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComments()))
        for i, v := range m.GetComments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("comments", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemPullsItemReviewsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. **Required** when using `REQUEST_CHANGES` or `COMMENT` for the `event` parameter. The body text of the pull request review.
func (m *ItemItemPullsItemReviewsPostRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetComments sets the comments property value. Use the following table to specify the location, destination, and contents of the draft review comment.
func (m *ItemItemPullsItemReviewsPostRequestBody) SetComments(value []ItemItemPullsItemReviewsPostRequestBody_commentsable)() {
    m.comments = value
}
// SetCommitId sets the commit_id property value. The SHA of the commit that needs a review. Not using the latest commit SHA may render your review comment outdated if a subsequent commit modifies the line you specify as the `position`. Defaults to the most recent commit in the pull request when you do not specify a value.
func (m *ItemItemPullsItemReviewsPostRequestBody) SetCommitId(value *string)() {
    m.commit_id = value
}
type ItemItemPullsItemReviewsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetComments()([]ItemItemPullsItemReviewsPostRequestBody_commentsable)
    GetCommitId()(*string)
    SetBody(value *string)()
    SetComments(value []ItemItemPullsItemReviewsPostRequestBody_commentsable)()
    SetCommitId(value *string)()
}
