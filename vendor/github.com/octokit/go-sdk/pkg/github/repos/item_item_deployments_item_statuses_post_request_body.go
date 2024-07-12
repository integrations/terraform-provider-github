package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemDeploymentsItemStatusesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Adds a new `inactive` status to all prior non-transient, non-production environment deployments with the same repository and `environment` name as the created status's deployment. An `inactive` status is only added to deployments that had a `success` state. Default: `true`
    auto_inactive *bool
    // A short description of the status. The maximum description length is 140 characters.
    description *string
    // Name for the target deployment environment, which can be changed when setting a deploy status. For example, `production`, `staging`, or `qa`. If not defined, the environment of the previous status on the deployment will be used, if it exists. Otherwise, the environment of the deployment will be used.
    environment *string
    // Sets the URL for accessing your environment. Default: `""`
    environment_url *string
    // The full URL of the deployment's output. This parameter replaces `target_url`. We will continue to accept `target_url` to support legacy uses, but we recommend replacing `target_url` with `log_url`. Setting `log_url` will automatically set `target_url` to the same value. Default: `""`
    log_url *string
    // The target URL to associate with this status. This URL should contain output to keep the user updated while the task is running or serve as historical information for what happened in the deployment. **Note:** It's recommended to use the `log_url` parameter, which replaces `target_url`.
    target_url *string
}
// NewItemItemDeploymentsItemStatusesPostRequestBody instantiates a new ItemItemDeploymentsItemStatusesPostRequestBody and sets the default values.
func NewItemItemDeploymentsItemStatusesPostRequestBody()(*ItemItemDeploymentsItemStatusesPostRequestBody) {
    m := &ItemItemDeploymentsItemStatusesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemDeploymentsItemStatusesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemDeploymentsItemStatusesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemDeploymentsItemStatusesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAutoInactive gets the auto_inactive property value. Adds a new `inactive` status to all prior non-transient, non-production environment deployments with the same repository and `environment` name as the created status's deployment. An `inactive` status is only added to deployments that had a `success` state. Default: `true`
// returns a *bool when successful
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) GetAutoInactive()(*bool) {
    return m.auto_inactive
}
// GetDescription gets the description property value. A short description of the status. The maximum description length is 140 characters.
// returns a *string when successful
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetEnvironment gets the environment property value. Name for the target deployment environment, which can be changed when setting a deploy status. For example, `production`, `staging`, or `qa`. If not defined, the environment of the previous status on the deployment will be used, if it exists. Otherwise, the environment of the deployment will be used.
// returns a *string when successful
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) GetEnvironment()(*string) {
    return m.environment
}
// GetEnvironmentUrl gets the environment_url property value. Sets the URL for accessing your environment. Default: `""`
// returns a *string when successful
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) GetEnvironmentUrl()(*string) {
    return m.environment_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["auto_inactive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoInactive(val)
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
    res["environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironment(val)
        }
        return nil
    }
    res["environment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironmentUrl(val)
        }
        return nil
    }
    res["log_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogUrl(val)
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
// GetLogUrl gets the log_url property value. The full URL of the deployment's output. This parameter replaces `target_url`. We will continue to accept `target_url` to support legacy uses, but we recommend replacing `target_url` with `log_url`. Setting `log_url` will automatically set `target_url` to the same value. Default: `""`
// returns a *string when successful
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) GetLogUrl()(*string) {
    return m.log_url
}
// GetTargetUrl gets the target_url property value. The target URL to associate with this status. This URL should contain output to keep the user updated while the task is running or serve as historical information for what happened in the deployment. **Note:** It's recommended to use the `log_url` parameter, which replaces `target_url`.
// returns a *string when successful
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) GetTargetUrl()(*string) {
    return m.target_url
}
// Serialize serializes information the current object
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("auto_inactive", m.GetAutoInactive())
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
        err := writer.WriteStringValue("environment", m.GetEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("environment_url", m.GetEnvironmentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("log_url", m.GetLogUrl())
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
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAutoInactive sets the auto_inactive property value. Adds a new `inactive` status to all prior non-transient, non-production environment deployments with the same repository and `environment` name as the created status's deployment. An `inactive` status is only added to deployments that had a `success` state. Default: `true`
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) SetAutoInactive(value *bool)() {
    m.auto_inactive = value
}
// SetDescription sets the description property value. A short description of the status. The maximum description length is 140 characters.
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetEnvironment sets the environment property value. Name for the target deployment environment, which can be changed when setting a deploy status. For example, `production`, `staging`, or `qa`. If not defined, the environment of the previous status on the deployment will be used, if it exists. Otherwise, the environment of the deployment will be used.
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) SetEnvironment(value *string)() {
    m.environment = value
}
// SetEnvironmentUrl sets the environment_url property value. Sets the URL for accessing your environment. Default: `""`
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) SetEnvironmentUrl(value *string)() {
    m.environment_url = value
}
// SetLogUrl sets the log_url property value. The full URL of the deployment's output. This parameter replaces `target_url`. We will continue to accept `target_url` to support legacy uses, but we recommend replacing `target_url` with `log_url`. Setting `log_url` will automatically set `target_url` to the same value. Default: `""`
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) SetLogUrl(value *string)() {
    m.log_url = value
}
// SetTargetUrl sets the target_url property value. The target URL to associate with this status. This URL should contain output to keep the user updated while the task is running or serve as historical information for what happened in the deployment. **Note:** It's recommended to use the `log_url` parameter, which replaces `target_url`.
func (m *ItemItemDeploymentsItemStatusesPostRequestBody) SetTargetUrl(value *string)() {
    m.target_url = value
}
type ItemItemDeploymentsItemStatusesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoInactive()(*bool)
    GetDescription()(*string)
    GetEnvironment()(*string)
    GetEnvironmentUrl()(*string)
    GetLogUrl()(*string)
    GetTargetUrl()(*string)
    SetAutoInactive(value *bool)()
    SetDescription(value *string)()
    SetEnvironment(value *string)()
    SetEnvironmentUrl(value *string)()
    SetLogUrl(value *string)()
    SetTargetUrl(value *string)()
}
