package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PorterLargeFile porter Large File
type PorterLargeFile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The oid property
    oid *string
    // The path property
    path *string
    // The ref_name property
    ref_name *string
    // The size property
    size *int32
}
// NewPorterLargeFile instantiates a new PorterLargeFile and sets the default values.
func NewPorterLargeFile()(*PorterLargeFile) {
    m := &PorterLargeFile{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePorterLargeFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePorterLargeFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPorterLargeFile(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PorterLargeFile) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PorterLargeFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["oid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOid(val)
        }
        return nil
    }
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
    res["ref_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRefName(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetOid gets the oid property value. The oid property
// returns a *string when successful
func (m *PorterLargeFile) GetOid()(*string) {
    return m.oid
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *PorterLargeFile) GetPath()(*string) {
    return m.path
}
// GetRefName gets the ref_name property value. The ref_name property
// returns a *string when successful
func (m *PorterLargeFile) GetRefName()(*string) {
    return m.ref_name
}
// GetSize gets the size property value. The size property
// returns a *int32 when successful
func (m *PorterLargeFile) GetSize()(*int32) {
    return m.size
}
// Serialize serializes information the current object
func (m *PorterLargeFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("oid", m.GetOid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref_name", m.GetRefName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size", m.GetSize())
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
func (m *PorterLargeFile) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOid sets the oid property value. The oid property
func (m *PorterLargeFile) SetOid(value *string)() {
    m.oid = value
}
// SetPath sets the path property value. The path property
func (m *PorterLargeFile) SetPath(value *string)() {
    m.path = value
}
// SetRefName sets the ref_name property value. The ref_name property
func (m *PorterLargeFile) SetRefName(value *string)() {
    m.ref_name = value
}
// SetSize sets the size property value. The size property
func (m *PorterLargeFile) SetSize(value *int32)() {
    m.size = value
}
type PorterLargeFileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOid()(*string)
    GetPath()(*string)
    GetRefName()(*string)
    GetSize()(*int32)
    SetOid(value *string)()
    SetPath(value *string)()
    SetRefName(value *string)()
    SetSize(value *int32)()
}
