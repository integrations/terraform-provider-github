package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Artifact_workflow_run struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The head_branch property
    head_branch *string
    // The head_repository_id property
    head_repository_id *int32
    // The head_sha property
    head_sha *string
    // The id property
    id *int32
    // The repository_id property
    repository_id *int32
}
// NewArtifact_workflow_run instantiates a new Artifact_workflow_run and sets the default values.
func NewArtifact_workflow_run()(*Artifact_workflow_run) {
    m := &Artifact_workflow_run{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateArtifact_workflow_runFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateArtifact_workflow_runFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewArtifact_workflow_run(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Artifact_workflow_run) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Artifact_workflow_run) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["head_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadBranch(val)
        }
        return nil
    }
    res["head_repository_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadRepositoryId(val)
        }
        return nil
    }
    res["head_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadSha(val)
        }
        return nil
    }
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
    return res
}
// GetHeadBranch gets the head_branch property value. The head_branch property
// returns a *string when successful
func (m *Artifact_workflow_run) GetHeadBranch()(*string) {
    return m.head_branch
}
// GetHeadRepositoryId gets the head_repository_id property value. The head_repository_id property
// returns a *int32 when successful
func (m *Artifact_workflow_run) GetHeadRepositoryId()(*int32) {
    return m.head_repository_id
}
// GetHeadSha gets the head_sha property value. The head_sha property
// returns a *string when successful
func (m *Artifact_workflow_run) GetHeadSha()(*string) {
    return m.head_sha
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Artifact_workflow_run) GetId()(*int32) {
    return m.id
}
// GetRepositoryId gets the repository_id property value. The repository_id property
// returns a *int32 when successful
func (m *Artifact_workflow_run) GetRepositoryId()(*int32) {
    return m.repository_id
}
// Serialize serializes information the current object
func (m *Artifact_workflow_run) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("head_branch", m.GetHeadBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("head_repository_id", m.GetHeadRepositoryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("head_sha", m.GetHeadSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Artifact_workflow_run) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHeadBranch sets the head_branch property value. The head_branch property
func (m *Artifact_workflow_run) SetHeadBranch(value *string)() {
    m.head_branch = value
}
// SetHeadRepositoryId sets the head_repository_id property value. The head_repository_id property
func (m *Artifact_workflow_run) SetHeadRepositoryId(value *int32)() {
    m.head_repository_id = value
}
// SetHeadSha sets the head_sha property value. The head_sha property
func (m *Artifact_workflow_run) SetHeadSha(value *string)() {
    m.head_sha = value
}
// SetId sets the id property value. The id property
func (m *Artifact_workflow_run) SetId(value *int32)() {
    m.id = value
}
// SetRepositoryId sets the repository_id property value. The repository_id property
func (m *Artifact_workflow_run) SetRepositoryId(value *int32)() {
    m.repository_id = value
}
type Artifact_workflow_runable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeadBranch()(*string)
    GetHeadRepositoryId()(*int32)
    GetHeadSha()(*string)
    GetId()(*int32)
    GetRepositoryId()(*int32)
    SetHeadBranch(value *string)()
    SetHeadRepositoryId(value *int32)()
    SetHeadSha(value *string)()
    SetId(value *int32)()
    SetRepositoryId(value *int32)()
}
