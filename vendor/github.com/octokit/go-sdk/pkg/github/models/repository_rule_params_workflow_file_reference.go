package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleParamsWorkflowFileReference a workflow that must run for this rule to pass
type RepositoryRuleParamsWorkflowFileReference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The path to the workflow file
    path *string
    // The ref (branch or tag) of the workflow file to use
    ref *string
    // The ID of the repository where the workflow is defined
    repository_id *int32
    // The commit SHA of the workflow file to use
    sha *string
}
// NewRepositoryRuleParamsWorkflowFileReference instantiates a new RepositoryRuleParamsWorkflowFileReference and sets the default values.
func NewRepositoryRuleParamsWorkflowFileReference()(*RepositoryRuleParamsWorkflowFileReference) {
    m := &RepositoryRuleParamsWorkflowFileReference{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleParamsWorkflowFileReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleParamsWorkflowFileReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleParamsWorkflowFileReference(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleParamsWorkflowFileReference) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleParamsWorkflowFileReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
        }
        return nil
    }
    res["repository_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryId(val)
        }
        return nil
    }
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    return res
}
// GetPath gets the path property value. The path to the workflow file
// returns a *string when successful
func (m *RepositoryRuleParamsWorkflowFileReference) GetPath()(*string) {
    return m.path
}
// GetRef gets the ref property value. The ref (branch or tag) of the workflow file to use
// returns a *string when successful
func (m *RepositoryRuleParamsWorkflowFileReference) GetRef()(*string) {
    return m.ref
}
// GetRepositoryId gets the repository_id property value. The ID of the repository where the workflow is defined
// returns a *int32 when successful
func (m *RepositoryRuleParamsWorkflowFileReference) GetRepositoryId()(*int32) {
    return m.repository_id
}
// GetSha gets the sha property value. The commit SHA of the workflow file to use
// returns a *string when successful
func (m *RepositoryRuleParamsWorkflowFileReference) GetSha()(*string) {
    return m.sha
}
// Serialize serializes information the current object
func (m *RepositoryRuleParamsWorkflowFileReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref", m.GetRef())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("repository_id", m.GetRepositoryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
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
func (m *RepositoryRuleParamsWorkflowFileReference) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPath sets the path property value. The path to the workflow file
func (m *RepositoryRuleParamsWorkflowFileReference) SetPath(value *string)() {
    m.path = value
}
// SetRef sets the ref property value. The ref (branch or tag) of the workflow file to use
func (m *RepositoryRuleParamsWorkflowFileReference) SetRef(value *string)() {
    m.ref = value
}
// SetRepositoryId sets the repository_id property value. The ID of the repository where the workflow is defined
func (m *RepositoryRuleParamsWorkflowFileReference) SetRepositoryId(value *int32)() {
    m.repository_id = value
}
// SetSha sets the sha property value. The commit SHA of the workflow file to use
func (m *RepositoryRuleParamsWorkflowFileReference) SetSha(value *string)() {
    m.sha = value
}
type RepositoryRuleParamsWorkflowFileReferenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPath()(*string)
    GetRef()(*string)
    GetRepositoryId()(*int32)
    GetSha()(*string)
    SetPath(value *string)()
    SetRef(value *string)()
    SetRepositoryId(value *int32)()
    SetSha(value *string)()
}
