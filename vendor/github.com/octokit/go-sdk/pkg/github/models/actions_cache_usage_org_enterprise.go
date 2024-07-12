package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsCacheUsageOrgEnterprise struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The count of active caches across all repositories of an enterprise or an organization.
    total_active_caches_count *int32
    // The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
    total_active_caches_size_in_bytes *int32
}
// NewActionsCacheUsageOrgEnterprise instantiates a new ActionsCacheUsageOrgEnterprise and sets the default values.
func NewActionsCacheUsageOrgEnterprise()(*ActionsCacheUsageOrgEnterprise) {
    m := &ActionsCacheUsageOrgEnterprise{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsCacheUsageOrgEnterpriseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsCacheUsageOrgEnterpriseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsCacheUsageOrgEnterprise(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsCacheUsageOrgEnterprise) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsCacheUsageOrgEnterprise) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["total_active_caches_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalActiveCachesCount(val)
        }
        return nil
    }
    res["total_active_caches_size_in_bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalActiveCachesSizeInBytes(val)
        }
        return nil
    }
    return res
}
// GetTotalActiveCachesCount gets the total_active_caches_count property value. The count of active caches across all repositories of an enterprise or an organization.
// returns a *int32 when successful
func (m *ActionsCacheUsageOrgEnterprise) GetTotalActiveCachesCount()(*int32) {
    return m.total_active_caches_count
}
// GetTotalActiveCachesSizeInBytes gets the total_active_caches_size_in_bytes property value. The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
// returns a *int32 when successful
func (m *ActionsCacheUsageOrgEnterprise) GetTotalActiveCachesSizeInBytes()(*int32) {
    return m.total_active_caches_size_in_bytes
}
// Serialize serializes information the current object
func (m *ActionsCacheUsageOrgEnterprise) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("total_active_caches_count", m.GetTotalActiveCachesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_active_caches_size_in_bytes", m.GetTotalActiveCachesSizeInBytes())
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
func (m *ActionsCacheUsageOrgEnterprise) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTotalActiveCachesCount sets the total_active_caches_count property value. The count of active caches across all repositories of an enterprise or an organization.
func (m *ActionsCacheUsageOrgEnterprise) SetTotalActiveCachesCount(value *int32)() {
    m.total_active_caches_count = value
}
// SetTotalActiveCachesSizeInBytes sets the total_active_caches_size_in_bytes property value. The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
func (m *ActionsCacheUsageOrgEnterprise) SetTotalActiveCachesSizeInBytes(value *int32)() {
    m.total_active_caches_size_in_bytes = value
}
type ActionsCacheUsageOrgEnterpriseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalActiveCachesCount()(*int32)
    GetTotalActiveCachesSizeInBytes()(*int32)
    SetTotalActiveCachesCount(value *int32)()
    SetTotalActiveCachesSizeInBytes(value *int32)()
}
