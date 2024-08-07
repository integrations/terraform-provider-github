package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id of the user or team who can review the deployment
    id *int32
    // The type of reviewer.
    typeEscaped *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentReviewerType
}
// NewItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers instantiates a new ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers and sets the default values.
func NewItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers()(*ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers) {
    m := &ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ParseDeploymentReviewerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentReviewerType))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id of the user or team who can review the deployment
// returns a *int32 when successful
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers) GetId()(*int32) {
    return m.id
}
// GetTypeEscaped gets the type property value. The type of reviewer.
// returns a *DeploymentReviewerType when successful
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers) GetTypeEscaped()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentReviewerType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id of the user or team who can review the deployment
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers) SetId(value *int32)() {
    m.id = value
}
// SetTypeEscaped sets the type property value. The type of reviewer.
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewers) SetTypeEscaped(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentReviewerType)() {
    m.typeEscaped = value
}
type ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetTypeEscaped()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentReviewerType)
    SetId(value *int32)()
    SetTypeEscaped(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentReviewerType)()
}
