package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemStatusesItemWithShaPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A string label to differentiate this status from the status of other systems. This field is case-insensitive.
    context *string
    // A short description of the status.
    description *string
    // The target URL to associate with this status. This URL will be linked from the GitHub UI to allow users to easily see the source of the status.  For example, if your continuous integration system is posting build status, you would want to provide the deep link for the build output for this specific SHA:  `http://ci.example.com/user/repo/build/sha`
    target_url *string
}
// NewItemItemStatusesItemWithShaPostRequestBody instantiates a new ItemItemStatusesItemWithShaPostRequestBody and sets the default values.
func NewItemItemStatusesItemWithShaPostRequestBody()(*ItemItemStatusesItemWithShaPostRequestBody) {
    m := &ItemItemStatusesItemWithShaPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    contextValue := "default"
    m.SetContext(&contextValue)
    return m
}
// CreateItemItemStatusesItemWithShaPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemStatusesItemWithShaPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemStatusesItemWithShaPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemStatusesItemWithShaPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContext gets the context property value. A string label to differentiate this status from the status of other systems. This field is case-insensitive.
// returns a *string when successful
func (m *ItemItemStatusesItemWithShaPostRequestBody) GetContext()(*string) {
    return m.context
}
// GetDescription gets the description property value. A short description of the status.
// returns a *string when successful
func (m *ItemItemStatusesItemWithShaPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemStatusesItemWithShaPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["context"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["target_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetUrl(val)
        }
        return nil
    }
    return res
}
// GetTargetUrl gets the target_url property value. The target URL to associate with this status. This URL will be linked from the GitHub UI to allow users to easily see the source of the status.  For example, if your continuous integration system is posting build status, you would want to provide the deep link for the build output for this specific SHA:  `http://ci.example.com/user/repo/build/sha`
// returns a *string when successful
func (m *ItemItemStatusesItemWithShaPostRequestBody) GetTargetUrl()(*string) {
    return m.target_url
}
// Serialize serializes information the current object
func (m *ItemItemStatusesItemWithShaPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("context", m.GetContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target_url", m.GetTargetUrl())
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
func (m *ItemItemStatusesItemWithShaPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContext sets the context property value. A string label to differentiate this status from the status of other systems. This field is case-insensitive.
func (m *ItemItemStatusesItemWithShaPostRequestBody) SetContext(value *string)() {
    m.context = value
}
// SetDescription sets the description property value. A short description of the status.
func (m *ItemItemStatusesItemWithShaPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetTargetUrl sets the target_url property value. The target URL to associate with this status. This URL will be linked from the GitHub UI to allow users to easily see the source of the status.  For example, if your continuous integration system is posting build status, you would want to provide the deep link for the build output for this specific SHA:  `http://ci.example.com/user/repo/build/sha`
func (m *ItemItemStatusesItemWithShaPostRequestBody) SetTargetUrl(value *string)() {
    m.target_url = value
}
type ItemItemStatusesItemWithShaPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContext()(*string)
    GetDescription()(*string)
    GetTargetUrl()(*string)
    SetContext(value *string)()
    SetDescription(value *string)()
    SetTargetUrl(value *string)()
}
