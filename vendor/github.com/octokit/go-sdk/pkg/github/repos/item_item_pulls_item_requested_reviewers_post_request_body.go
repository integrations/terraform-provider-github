package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsItemRequested_reviewersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An array of user `login`s that will be requested.
    reviewers []string
    // An array of team `slug`s that will be requested.
    team_reviewers []string
}
// NewItemItemPullsItemRequested_reviewersPostRequestBody instantiates a new ItemItemPullsItemRequested_reviewersPostRequestBody and sets the default values.
func NewItemItemPullsItemRequested_reviewersPostRequestBody()(*ItemItemPullsItemRequested_reviewersPostRequestBody) {
    m := &ItemItemPullsItemRequested_reviewersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemRequested_reviewersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsItemRequested_reviewersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemRequested_reviewersPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsItemRequested_reviewersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsItemRequested_reviewersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetReviewers(res)
        }
        return nil
    }
    res["team_reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTeamReviewers(res)
        }
        return nil
    }
    return res
}
// GetReviewers gets the reviewers property value. An array of user `login`s that will be requested.
// returns a []string when successful
func (m *ItemItemPullsItemRequested_reviewersPostRequestBody) GetReviewers()([]string) {
    return m.reviewers
}
// GetTeamReviewers gets the team_reviewers property value. An array of team `slug`s that will be requested.
// returns a []string when successful
func (m *ItemItemPullsItemRequested_reviewersPostRequestBody) GetTeamReviewers()([]string) {
    return m.team_reviewers
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemRequested_reviewersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReviewers() != nil {
        err := writer.WriteCollectionOfStringValues("reviewers", m.GetReviewers())
        if err != nil {
            return err
        }
    }
    if m.GetTeamReviewers() != nil {
        err := writer.WriteCollectionOfStringValues("team_reviewers", m.GetTeamReviewers())
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
func (m *ItemItemPullsItemRequested_reviewersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReviewers sets the reviewers property value. An array of user `login`s that will be requested.
func (m *ItemItemPullsItemRequested_reviewersPostRequestBody) SetReviewers(value []string)() {
    m.reviewers = value
}
// SetTeamReviewers sets the team_reviewers property value. An array of team `slug`s that will be requested.
func (m *ItemItemPullsItemRequested_reviewersPostRequestBody) SetTeamReviewers(value []string)() {
    m.team_reviewers = value
}
type ItemItemPullsItemRequested_reviewersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewers()([]string)
    GetTeamReviewers()([]string)
    SetReviewers(value []string)()
    SetTeamReviewers(value []string)()
}
