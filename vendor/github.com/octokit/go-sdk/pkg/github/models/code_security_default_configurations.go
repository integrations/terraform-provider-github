package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeSecurityDefaultConfigurations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A code security configuration
    configuration CodeSecurityConfigurationable
    // The visibility of newly created repositories for which the code security configuration will be applied to by default
    default_for_new_repos *CodeSecurityDefaultConfigurations_default_for_new_repos
}
// NewCodeSecurityDefaultConfigurations instantiates a new CodeSecurityDefaultConfigurations and sets the default values.
func NewCodeSecurityDefaultConfigurations()(*CodeSecurityDefaultConfigurations) {
    m := &CodeSecurityDefaultConfigurations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeSecurityDefaultConfigurationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeSecurityDefaultConfigurationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeSecurityDefaultConfigurations(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeSecurityDefaultConfigurations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfiguration gets the configuration property value. A code security configuration
// returns a CodeSecurityConfigurationable when successful
func (m *CodeSecurityDefaultConfigurations) GetConfiguration()(CodeSecurityConfigurationable) {
    return m.configuration
}
// GetDefaultForNewRepos gets the default_for_new_repos property value. The visibility of newly created repositories for which the code security configuration will be applied to by default
// returns a *CodeSecurityDefaultConfigurations_default_for_new_repos when successful
func (m *CodeSecurityDefaultConfigurations) GetDefaultForNewRepos()(*CodeSecurityDefaultConfigurations_default_for_new_repos) {
    return m.default_for_new_repos
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeSecurityDefaultConfigurations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeSecurityConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(CodeSecurityConfigurationable))
        }
        return nil
    }
    res["default_for_new_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityDefaultConfigurations_default_for_new_repos)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultForNewRepos(val.(*CodeSecurityDefaultConfigurations_default_for_new_repos))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CodeSecurityDefaultConfigurations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultForNewRepos() != nil {
        cast := (*m.GetDefaultForNewRepos()).String()
        err := writer.WriteStringValue("default_for_new_repos", &cast)
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
func (m *CodeSecurityDefaultConfigurations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfiguration sets the configuration property value. A code security configuration
func (m *CodeSecurityDefaultConfigurations) SetConfiguration(value CodeSecurityConfigurationable)() {
    m.configuration = value
}
// SetDefaultForNewRepos sets the default_for_new_repos property value. The visibility of newly created repositories for which the code security configuration will be applied to by default
func (m *CodeSecurityDefaultConfigurations) SetDefaultForNewRepos(value *CodeSecurityDefaultConfigurations_default_for_new_repos)() {
    m.default_for_new_repos = value
}
type CodeSecurityDefaultConfigurationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguration()(CodeSecurityConfigurationable)
    GetDefaultForNewRepos()(*CodeSecurityDefaultConfigurations_default_for_new_repos)
    SetConfiguration(value CodeSecurityConfigurationable)()
    SetDefaultForNewRepos(value *CodeSecurityDefaultConfigurations_default_for_new_repos)()
}
