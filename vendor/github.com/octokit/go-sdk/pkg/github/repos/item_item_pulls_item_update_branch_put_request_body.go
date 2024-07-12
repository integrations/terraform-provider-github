package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsItemUpdateBranchPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The expected SHA of the pull request's HEAD ref. This is the most recent commit on the pull request's branch. If the expected SHA does not match the pull request's HEAD, you will receive a `422 Unprocessable Entity` status. You can use the "[List commits](https://docs.github.com/rest/commits/commits#list-commits)" endpoint to find the most recent commit SHA. Default: SHA of the pull request's current HEAD ref.
    expected_head_sha *string
}
// NewItemItemPullsItemUpdateBranchPutRequestBody instantiates a new ItemItemPullsItemUpdateBranchPutRequestBody and sets the default values.
func NewItemItemPullsItemUpdateBranchPutRequestBody()(*ItemItemPullsItemUpdateBranchPutRequestBody) {
    m := &ItemItemPullsItemUpdateBranchPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemUpdateBranchPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsItemUpdateBranchPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemUpdateBranchPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsItemUpdateBranchPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpectedHeadSha gets the expected_head_sha property value. The expected SHA of the pull request's HEAD ref. This is the most recent commit on the pull request's branch. If the expected SHA does not match the pull request's HEAD, you will receive a `422 Unprocessable Entity` status. You can use the "[List commits](https://docs.github.com/rest/commits/commits#list-commits)" endpoint to find the most recent commit SHA. Default: SHA of the pull request's current HEAD ref.
// returns a *string when successful
func (m *ItemItemPullsItemUpdateBranchPutRequestBody) GetExpectedHeadSha()(*string) {
    return m.expected_head_sha
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsItemUpdateBranchPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expected_head_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpectedHeadSha(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemUpdateBranchPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expected_head_sha", m.GetExpectedHeadSha())
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
func (m *ItemItemPullsItemUpdateBranchPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpectedHeadSha sets the expected_head_sha property value. The expected SHA of the pull request's HEAD ref. This is the most recent commit on the pull request's branch. If the expected SHA does not match the pull request's HEAD, you will receive a `422 Unprocessable Entity` status. You can use the "[List commits](https://docs.github.com/rest/commits/commits#list-commits)" endpoint to find the most recent commit SHA. Default: SHA of the pull request's current HEAD ref.
func (m *ItemItemPullsItemUpdateBranchPutRequestBody) SetExpectedHeadSha(value *string)() {
    m.expected_head_sha = value
}
type ItemItemPullsItemUpdateBranchPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpectedHeadSha()(*string)
    SetExpectedHeadSha(value *string)()
}
