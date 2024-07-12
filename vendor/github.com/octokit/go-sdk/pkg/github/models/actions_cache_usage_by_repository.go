package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsCacheUsageByRepository gitHub Actions Cache Usage by repository.
type ActionsCacheUsageByRepository struct {
    // The number of active caches in the repository.
    active_caches_count *int32
    // The sum of the size in bytes of all the active cache items in the repository.
    active_caches_size_in_bytes *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repository owner and name for the cache usage being shown.
    full_name *string
}
// NewActionsCacheUsageByRepository instantiates a new ActionsCacheUsageByRepository and sets the default values.
func NewActionsCacheUsageByRepository()(*ActionsCacheUsageByRepository) {
    m := &ActionsCacheUsageByRepository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsCacheUsageByRepositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsCacheUsageByRepositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsCacheUsageByRepository(), nil
}
// GetActiveCachesCount gets the active_caches_count property value. The number of active caches in the repository.
// returns a *int32 when successful
func (m *ActionsCacheUsageByRepository) GetActiveCachesCount()(*int32) {
    return m.active_caches_count
}
// GetActiveCachesSizeInBytes gets the active_caches_size_in_bytes property value. The sum of the size in bytes of all the active cache items in the repository.
// returns a *int32 when successful
func (m *ActionsCacheUsageByRepository) GetActiveCachesSizeInBytes()(*int32) {
    return m.active_caches_size_in_bytes
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsCacheUsageByRepository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsCacheUsageByRepository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active_caches_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveCachesCount(val)
        }
        return nil
    }
    res["active_caches_size_in_bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveCachesSizeInBytes(val)
        }
        return nil
    }
    res["full_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullName(val)
        }
        return nil
    }
    return res
}
// GetFullName gets the full_name property value. The repository owner and name for the cache usage being shown.
// returns a *string when successful
func (m *ActionsCacheUsageByRepository) GetFullName()(*string) {
    return m.full_name
}
// Serialize serializes information the current object
func (m *ActionsCacheUsageByRepository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("active_caches_count", m.GetActiveCachesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("active_caches_size_in_bytes", m.GetActiveCachesSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("full_name", m.GetFullName())
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
// SetActiveCachesCount sets the active_caches_count property value. The number of active caches in the repository.
func (m *ActionsCacheUsageByRepository) SetActiveCachesCount(value *int32)() {
    m.active_caches_count = value
}
// SetActiveCachesSizeInBytes sets the active_caches_size_in_bytes property value. The sum of the size in bytes of all the active cache items in the repository.
func (m *ActionsCacheUsageByRepository) SetActiveCachesSizeInBytes(value *int32)() {
    m.active_caches_size_in_bytes = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionsCacheUsageByRepository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFullName sets the full_name property value. The repository owner and name for the cache usage being shown.
func (m *ActionsCacheUsageByRepository) SetFullName(value *string)() {
    m.full_name = value
}
type ActionsCacheUsageByRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveCachesCount()(*int32)
    GetActiveCachesSizeInBytes()(*int32)
    GetFullName()(*string)
    SetActiveCachesCount(value *int32)()
    SetActiveCachesSizeInBytes(value *int32)()
    SetFullName(value *string)()
}
