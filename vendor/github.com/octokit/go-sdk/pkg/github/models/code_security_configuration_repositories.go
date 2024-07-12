package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeSecurityConfigurationRepositories repositories associated with a code security configuration and attachment status
type CodeSecurityConfigurationRepositories struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub repository.
    repository SimpleRepositoryable
    // The attachment status of the code security configuration on the repository.
    status *CodeSecurityConfigurationRepositories_status
}
// NewCodeSecurityConfigurationRepositories instantiates a new CodeSecurityConfigurationRepositories and sets the default values.
func NewCodeSecurityConfigurationRepositories()(*CodeSecurityConfigurationRepositories) {
    m := &CodeSecurityConfigurationRepositories{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeSecurityConfigurationRepositoriesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeSecurityConfigurationRepositoriesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeSecurityConfigurationRepositories(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeSecurityConfigurationRepositories) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeSecurityConfigurationRepositories) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(SimpleRepositoryable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfigurationRepositories_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CodeSecurityConfigurationRepositories_status))
        }
        return nil
    }
    return res
}
// GetRepository gets the repository property value. A GitHub repository.
// returns a SimpleRepositoryable when successful
func (m *CodeSecurityConfigurationRepositories) GetRepository()(SimpleRepositoryable) {
    return m.repository
}
// GetStatus gets the status property value. The attachment status of the code security configuration on the repository.
// returns a *CodeSecurityConfigurationRepositories_status when successful
func (m *CodeSecurityConfigurationRepositories) GetStatus()(*CodeSecurityConfigurationRepositories_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *CodeSecurityConfigurationRepositories) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *CodeSecurityConfigurationRepositories) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepository sets the repository property value. A GitHub repository.
func (m *CodeSecurityConfigurationRepositories) SetRepository(value SimpleRepositoryable)() {
    m.repository = value
}
// SetStatus sets the status property value. The attachment status of the code security configuration on the repository.
func (m *CodeSecurityConfigurationRepositories) SetStatus(value *CodeSecurityConfigurationRepositories_status)() {
    m.status = value
}
type CodeSecurityConfigurationRepositoriesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepository()(SimpleRepositoryable)
    GetStatus()(*CodeSecurityConfigurationRepositories_status)
    SetRepository(value SimpleRepositoryable)()
    SetStatus(value *CodeSecurityConfigurationRepositories_status)()
}
