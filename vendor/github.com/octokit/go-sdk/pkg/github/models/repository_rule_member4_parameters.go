package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleMember4_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The maximum file size allowed in megabytes. This limit does not apply to Git Large File Storage (Git LFS).
    max_file_size *int32
}
// NewRepositoryRuleMember4_parameters instantiates a new RepositoryRuleMember4_parameters and sets the default values.
func NewRepositoryRuleMember4_parameters()(*RepositoryRuleMember4_parameters) {
    m := &RepositoryRuleMember4_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleMember4_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleMember4_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleMember4_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleMember4_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleMember4_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["max_file_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxFileSize(val)
        }
        return nil
    }
    return res
}
// GetMaxFileSize gets the max_file_size property value. The maximum file size allowed in megabytes. This limit does not apply to Git Large File Storage (Git LFS).
// returns a *int32 when successful
func (m *RepositoryRuleMember4_parameters) GetMaxFileSize()(*int32) {
    return m.max_file_size
}
// Serialize serializes information the current object
func (m *RepositoryRuleMember4_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("max_file_size", m.GetMaxFileSize())
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
func (m *RepositoryRuleMember4_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaxFileSize sets the max_file_size property value. The maximum file size allowed in megabytes. This limit does not apply to Git Large File Storage (Git LFS).
func (m *RepositoryRuleMember4_parameters) SetMaxFileSize(value *int32)() {
    m.max_file_size = value
}
type RepositoryRuleMember4_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxFileSize()(*int32)
    SetMaxFileSize(value *int32)()
}
