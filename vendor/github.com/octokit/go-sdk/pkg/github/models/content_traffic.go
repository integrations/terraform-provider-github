package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentTraffic content Traffic
type ContentTraffic struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The count property
    count *int32
    // The path property
    path *string
    // The title property
    title *string
    // The uniques property
    uniques *int32
}
// NewContentTraffic instantiates a new ContentTraffic and sets the default values.
func NewContentTraffic()(*ContentTraffic) {
    m := &ContentTraffic{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContentTrafficFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContentTrafficFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentTraffic(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContentTraffic) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *ContentTraffic) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContentTraffic) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
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
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["uniques"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniques(val)
        }
        return nil
    }
    return res
}
// GetPath gets the path property value. The path property
// returns a *string when successful
func (m *ContentTraffic) GetPath()(*string) {
    return m.path
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *ContentTraffic) GetTitle()(*string) {
    return m.title
}
// GetUniques gets the uniques property value. The uniques property
// returns a *int32 when successful
func (m *ContentTraffic) GetUniques()(*int32) {
    return m.uniques
}
// Serialize serializes information the current object
func (m *ContentTraffic) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
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
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("uniques", m.GetUniques())
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
func (m *ContentTraffic) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. The count property
func (m *ContentTraffic) SetCount(value *int32)() {
    m.count = value
}
// SetPath sets the path property value. The path property
func (m *ContentTraffic) SetPath(value *string)() {
    m.path = value
}
// SetTitle sets the title property value. The title property
func (m *ContentTraffic) SetTitle(value *string)() {
    m.title = value
}
// SetUniques sets the uniques property value. The uniques property
func (m *ContentTraffic) SetUniques(value *int32)() {
    m.uniques = value
}
type ContentTrafficable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetPath()(*string)
    GetTitle()(*string)
    GetUniques()(*int32)
    SetCount(value *int32)()
    SetPath(value *string)()
    SetTitle(value *string)()
    SetUniques(value *int32)()
}
