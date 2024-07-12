package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPullsItemRequested_reviewersDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An array of user `login`s that will be removed.
    reviewers []string
    // An array of team `slug`s that will be removed.
    team_reviewers []string
}
// NewItemItemPullsItemRequested_reviewersDeleteRequestBody instantiates a new ItemItemPullsItemRequested_reviewersDeleteRequestBody and sets the default values.
func NewItemItemPullsItemRequested_reviewersDeleteRequestBody()(*ItemItemPullsItemRequested_reviewersDeleteRequestBody) {
    m := &ItemItemPullsItemRequested_reviewersDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemRequested_reviewersDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPullsItemRequested_reviewersDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemRequested_reviewersDeleteRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPullsItemRequested_reviewersDeleteRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPullsItemRequested_reviewersDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetReviewers gets the reviewers property value. An array of user `login`s that will be removed.
// returns a []string when successful
func (m *ItemItemPullsItemRequested_reviewersDeleteRequestBody) GetReviewers()([]string) {
    return m.reviewers
}
// GetTeamReviewers gets the team_reviewers property value. An array of team `slug`s that will be removed.
// returns a []string when successful
func (m *ItemItemPullsItemRequested_reviewersDeleteRequestBody) GetTeamReviewers()([]string) {
    return m.team_reviewers
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemRequested_reviewersDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemPullsItemRequested_reviewersDeleteRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReviewers sets the reviewers property value. An array of user `login`s that will be removed.
func (m *ItemItemPullsItemRequested_reviewersDeleteRequestBody) SetReviewers(value []string)() {
    m.reviewers = value
}
// SetTeamReviewers sets the team_reviewers property value. An array of team `slug`s that will be removed.
func (m *ItemItemPullsItemRequested_reviewersDeleteRequestBody) SetTeamReviewers(value []string)() {
    m.team_reviewers = value
}
type ItemItemPullsItemRequested_reviewersDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewers()([]string)
    GetTeamReviewers()([]string)
    SetReviewers(value []string)()
    SetTeamReviewers(value []string)()
}
