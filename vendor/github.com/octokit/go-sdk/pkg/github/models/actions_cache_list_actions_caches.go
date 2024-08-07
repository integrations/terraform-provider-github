package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsCacheList_actions_caches struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The id property
    id *int32
    // The key property
    key *string
    // The last_accessed_at property
    last_accessed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ref property
    ref *string
    // The size_in_bytes property
    size_in_bytes *int32
    // The version property
    version *string
}
// NewActionsCacheList_actions_caches instantiates a new ActionsCacheList_actions_caches and sets the default values.
func NewActionsCacheList_actions_caches()(*ActionsCacheList_actions_caches) {
    m := &ActionsCacheList_actions_caches{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsCacheList_actions_cachesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsCacheList_actions_cachesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsCacheList_actions_caches(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsCacheList_actions_caches) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *ActionsCacheList_actions_caches) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsCacheList_actions_caches) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
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
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["last_accessed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAccessedAt(val)
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
    res["size_in_bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *ActionsCacheList_actions_caches) GetId()(*int32) {
    return m.id
}
// GetKey gets the key property value. The key property
// returns a *string when successful
func (m *ActionsCacheList_actions_caches) GetKey()(*string) {
    return m.key
}
// GetLastAccessedAt gets the last_accessed_at property value. The last_accessed_at property
// returns a *Time when successful
func (m *ActionsCacheList_actions_caches) GetLastAccessedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_accessed_at
}
// GetRef gets the ref property value. The ref property
// returns a *string when successful
func (m *ActionsCacheList_actions_caches) GetRef()(*string) {
    return m.ref
}
// GetSizeInBytes gets the size_in_bytes property value. The size_in_bytes property
// returns a *int32 when successful
func (m *ActionsCacheList_actions_caches) GetSizeInBytes()(*int32) {
    return m.size_in_bytes
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *ActionsCacheList_actions_caches) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *ActionsCacheList_actions_caches) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("last_accessed_at", m.GetLastAccessedAt())
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
        err := writer.WriteInt32Value("size_in_bytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *ActionsCacheList_actions_caches) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *ActionsCacheList_actions_caches) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetId sets the id property value. The id property
func (m *ActionsCacheList_actions_caches) SetId(value *int32)() {
    m.id = value
}
// SetKey sets the key property value. The key property
func (m *ActionsCacheList_actions_caches) SetKey(value *string)() {
    m.key = value
}
// SetLastAccessedAt sets the last_accessed_at property value. The last_accessed_at property
func (m *ActionsCacheList_actions_caches) SetLastAccessedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_accessed_at = value
}
// SetRef sets the ref property value. The ref property
func (m *ActionsCacheList_actions_caches) SetRef(value *string)() {
    m.ref = value
}
// SetSizeInBytes sets the size_in_bytes property value. The size_in_bytes property
func (m *ActionsCacheList_actions_caches) SetSizeInBytes(value *int32)() {
    m.size_in_bytes = value
}
// SetVersion sets the version property value. The version property
func (m *ActionsCacheList_actions_caches) SetVersion(value *string)() {
    m.version = value
}
type ActionsCacheList_actions_cachesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int32)
    GetKey()(*string)
    GetLastAccessedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRef()(*string)
    GetSizeInBytes()(*int32)
    GetVersion()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int32)()
    SetKey(value *string)()
    SetLastAccessedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRef(value *string)()
    SetSizeInBytes(value *int32)()
    SetVersion(value *string)()
}
