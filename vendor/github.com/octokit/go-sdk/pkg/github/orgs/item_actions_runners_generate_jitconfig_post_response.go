package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemActionsRunnersGenerateJitconfigPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The base64 encoded runner configuration.
    encoded_jit_config *string
    // A self hosted runner
    runner i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable
}
// NewItemActionsRunnersGenerateJitconfigPostResponse instantiates a new ItemActionsRunnersGenerateJitconfigPostResponse and sets the default values.
func NewItemActionsRunnersGenerateJitconfigPostResponse()(*ItemActionsRunnersGenerateJitconfigPostResponse) {
    m := &ItemActionsRunnersGenerateJitconfigPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsRunnersGenerateJitconfigPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsRunnersGenerateJitconfigPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnersGenerateJitconfigPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsRunnersGenerateJitconfigPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEncodedJitConfig gets the encoded_jit_config property value. The base64 encoded runner configuration.
// returns a *string when successful
func (m *ItemActionsRunnersGenerateJitconfigPostResponse) GetEncodedJitConfig()(*string) {
    return m.encoded_jit_config
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsRunnersGenerateJitconfigPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encoded_jit_config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncodedJitConfig(val)
        }
        return nil
    }
    res["runner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRunnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunner(val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable))
        }
        return nil
    }
    return res
}
// GetRunner gets the runner property value. A self hosted runner
// returns a Runnerable when successful
func (m *ItemActionsRunnersGenerateJitconfigPostResponse) GetRunner()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable) {
    return m.runner
}
// Serialize serializes information the current object
func (m *ItemActionsRunnersGenerateJitconfigPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("encoded_jit_config", m.GetEncodedJitConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("runner", m.GetRunner())
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
func (m *ItemActionsRunnersGenerateJitconfigPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEncodedJitConfig sets the encoded_jit_config property value. The base64 encoded runner configuration.
func (m *ItemActionsRunnersGenerateJitconfigPostResponse) SetEncodedJitConfig(value *string)() {
    m.encoded_jit_config = value
}
// SetRunner sets the runner property value. A self hosted runner
func (m *ItemActionsRunnersGenerateJitconfigPostResponse) SetRunner(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable)() {
    m.runner = value
}
type ItemActionsRunnersGenerateJitconfigPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncodedJitConfig()(*string)
    GetRunner()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable)
    SetEncodedJitConfig(value *string)()
    SetRunner(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Runnerable)()
}
