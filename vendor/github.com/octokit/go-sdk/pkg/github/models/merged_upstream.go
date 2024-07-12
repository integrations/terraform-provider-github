package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MergedUpstream results of a successful merge upstream request
type MergedUpstream struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The base_branch property
    base_branch *string
    // The merge_type property
    merge_type *MergedUpstream_merge_type
    // The message property
    message *string
}
// NewMergedUpstream instantiates a new MergedUpstream and sets the default values.
func NewMergedUpstream()(*MergedUpstream) {
    m := &MergedUpstream{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMergedUpstreamFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMergedUpstreamFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMergedUpstream(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MergedUpstream) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBaseBranch gets the base_branch property value. The base_branch property
// returns a *string when successful
func (m *MergedUpstream) GetBaseBranch()(*string) {
    return m.base_branch
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MergedUpstream) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["base_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseBranch(val)
        }
        return nil
    }
    res["merge_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMergedUpstream_merge_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeType(val.(*MergedUpstream_merge_type))
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetMergeType gets the merge_type property value. The merge_type property
// returns a *MergedUpstream_merge_type when successful
func (m *MergedUpstream) GetMergeType()(*MergedUpstream_merge_type) {
    return m.merge_type
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *MergedUpstream) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *MergedUpstream) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("base_branch", m.GetBaseBranch())
        if err != nil {
            return err
        }
    }
    if m.GetMergeType() != nil {
        cast := (*m.GetMergeType()).String()
        err := writer.WriteStringValue("merge_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *MergedUpstream) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBaseBranch sets the base_branch property value. The base_branch property
func (m *MergedUpstream) SetBaseBranch(value *string)() {
    m.base_branch = value
}
// SetMergeType sets the merge_type property value. The merge_type property
func (m *MergedUpstream) SetMergeType(value *MergedUpstream_merge_type)() {
    m.merge_type = value
}
// SetMessage sets the message property value. The message property
func (m *MergedUpstream) SetMessage(value *string)() {
    m.message = value
}
type MergedUpstreamable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseBranch()(*string)
    GetMergeType()(*MergedUpstream_merge_type)
    GetMessage()(*string)
    SetBaseBranch(value *string)()
    SetMergeType(value *MergedUpstream_merge_type)()
    SetMessage(value *string)()
}
