package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Snapshot_detector a description of the detector used.
type Snapshot_detector struct {
    // The name of the detector used.
    name *string
    // The url of the detector used.
    url *string
    // The version of the detector used.
    version *string
}
// NewSnapshot_detector instantiates a new Snapshot_detector and sets the default values.
func NewSnapshot_detector()(*Snapshot_detector) {
    m := &Snapshot_detector{
    }
    return m
}
// CreateSnapshot_detectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSnapshot_detectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSnapshot_detector(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Snapshot_detector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetName gets the name property value. The name of the detector used.
// returns a *string when successful
func (m *Snapshot_detector) GetName()(*string) {
    return m.name
}
// GetUrl gets the url property value. The url of the detector used.
// returns a *string when successful
func (m *Snapshot_detector) GetUrl()(*string) {
    return m.url
}
// GetVersion gets the version property value. The version of the detector used.
// returns a *string when successful
func (m *Snapshot_detector) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *Snapshot_detector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name of the detector used.
func (m *Snapshot_detector) SetName(value *string)() {
    m.name = value
}
// SetUrl sets the url property value. The url of the detector used.
func (m *Snapshot_detector) SetUrl(value *string)() {
    m.url = value
}
// SetVersion sets the version property value. The version of the detector used.
func (m *Snapshot_detector) SetVersion(value *string)() {
    m.version = value
}
type Snapshot_detectorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetUrl()(*string)
    GetVersion()(*string)
    SetName(value *string)()
    SetUrl(value *string)()
    SetVersion(value *string)()
}
