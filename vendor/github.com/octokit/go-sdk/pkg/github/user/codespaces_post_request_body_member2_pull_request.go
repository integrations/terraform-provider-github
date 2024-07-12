package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesPostRequestBodyMember2_pull_request pull request number for this codespace
type CodespacesPostRequestBodyMember2_pull_request struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Pull request number
    pull_request_number *int32
    // Repository id for this codespace
    repository_id *int32
}
// NewCodespacesPostRequestBodyMember2_pull_request instantiates a new CodespacesPostRequestBodyMember2_pull_request and sets the default values.
func NewCodespacesPostRequestBodyMember2_pull_request()(*CodespacesPostRequestBodyMember2_pull_request) {
    m := &CodespacesPostRequestBodyMember2_pull_request{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesPostRequestBodyMember2_pull_requestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodespacesPostRequestBodyMember2_pull_requestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesPostRequestBodyMember2_pull_request(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodespacesPostRequestBodyMember2_pull_request) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodespacesPostRequestBodyMember2_pull_request) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pull_request_number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestNumber(val)
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
// GetPullRequestNumber gets the pull_request_number property value. Pull request number
// returns a *int32 when successful
func (m *CodespacesPostRequestBodyMember2_pull_request) GetPullRequestNumber()(*int32) {
    return m.pull_request_number
}
// GetRepositoryId gets the repository_id property value. Repository id for this codespace
// returns a *int32 when successful
func (m *CodespacesPostRequestBodyMember2_pull_request) GetRepositoryId()(*int32) {
    return m.repository_id
}
// Serialize serializes information the current object
func (m *CodespacesPostRequestBodyMember2_pull_request) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("pull_request_number", m.GetPullRequestNumber())
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
func (m *CodespacesPostRequestBodyMember2_pull_request) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPullRequestNumber sets the pull_request_number property value. Pull request number
func (m *CodespacesPostRequestBodyMember2_pull_request) SetPullRequestNumber(value *int32)() {
    m.pull_request_number = value
}
// SetRepositoryId sets the repository_id property value. Repository id for this codespace
func (m *CodespacesPostRequestBodyMember2_pull_request) SetRepositoryId(value *int32)() {
    m.repository_id = value
}
type CodespacesPostRequestBodyMember2_pull_requestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPullRequestNumber()(*int32)
    GetRepositoryId()(*int32)
    SetPullRequestNumber(value *int32)()
    SetRepositoryId(value *int32)()
}
