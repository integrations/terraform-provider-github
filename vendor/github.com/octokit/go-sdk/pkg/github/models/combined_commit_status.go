package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CombinedCommitStatus combined Commit Status
type CombinedCommitStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commit_url property
    commit_url *string
    // Minimal Repository
    repository MinimalRepositoryable
    // The sha property
    sha *string
    // The state property
    state *string
    // The statuses property
    statuses []SimpleCommitStatusable
    // The total_count property
    total_count *int32
    // The url property
    url *string
}
// NewCombinedCommitStatus instantiates a new CombinedCommitStatus and sets the default values.
func NewCombinedCommitStatus()(*CombinedCommitStatus) {
    m := &CombinedCommitStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCombinedCommitStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCombinedCommitStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCombinedCommitStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CombinedCommitStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommitUrl gets the commit_url property value. The commit_url property
// returns a *string when successful
func (m *CombinedCommitStatus) GetCommitUrl()(*string) {
    return m.commit_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CombinedCommitStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commit_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitUrl(val)
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(MinimalRepositoryable))
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["statuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimpleCommitStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimpleCommitStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SimpleCommitStatusable)
                }
            }
            m.SetStatuses(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetRepository gets the repository property value. Minimal Repository
// returns a MinimalRepositoryable when successful
func (m *CombinedCommitStatus) GetRepository()(MinimalRepositoryable) {
    return m.repository
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *CombinedCommitStatus) GetSha()(*string) {
    return m.sha
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *CombinedCommitStatus) GetState()(*string) {
    return m.state
}
// GetStatuses gets the statuses property value. The statuses property
// returns a []SimpleCommitStatusable when successful
func (m *CombinedCommitStatus) GetStatuses()([]SimpleCommitStatusable) {
    return m.statuses
}
// GetTotalCount gets the total_count property value. The total_count property
// returns a *int32 when successful
func (m *CombinedCommitStatus) GetTotalCount()(*int32) {
    return m.total_count
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *CombinedCommitStatus) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CombinedCommitStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("commit_url", m.GetCommitUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
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
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    if m.GetStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStatuses()))
        for i, v := range m.GetStatuses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("statuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *CombinedCommitStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommitUrl sets the commit_url property value. The commit_url property
func (m *CombinedCommitStatus) SetCommitUrl(value *string)() {
    m.commit_url = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *CombinedCommitStatus) SetRepository(value MinimalRepositoryable)() {
    m.repository = value
}
// SetSha sets the sha property value. The sha property
func (m *CombinedCommitStatus) SetSha(value *string)() {
    m.sha = value
}
// SetState sets the state property value. The state property
func (m *CombinedCommitStatus) SetState(value *string)() {
    m.state = value
}
// SetStatuses sets the statuses property value. The statuses property
func (m *CombinedCommitStatus) SetStatuses(value []SimpleCommitStatusable)() {
    m.statuses = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *CombinedCommitStatus) SetTotalCount(value *int32)() {
    m.total_count = value
}
// SetUrl sets the url property value. The url property
func (m *CombinedCommitStatus) SetUrl(value *string)() {
    m.url = value
}
type CombinedCommitStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommitUrl()(*string)
    GetRepository()(MinimalRepositoryable)
    GetSha()(*string)
    GetState()(*string)
    GetStatuses()([]SimpleCommitStatusable)
    GetTotalCount()(*int32)
    GetUrl()(*string)
    SetCommitUrl(value *string)()
    SetRepository(value MinimalRepositoryable)()
    SetSha(value *string)()
    SetState(value *string)()
    SetStatuses(value []SimpleCommitStatusable)()
    SetTotalCount(value *int32)()
    SetUrl(value *string)()
}
