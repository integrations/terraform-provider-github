package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleset__links_self struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URL of the ruleset
    href *string
}
// NewRepositoryRuleset__links_self instantiates a new RepositoryRuleset__links_self and sets the default values.
func NewRepositoryRuleset__links_self()(*RepositoryRuleset__links_self) {
    m := &RepositoryRuleset__links_self{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleset__links_selfFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleset__links_selfFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleset__links_self(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleset__links_self) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleset__links_self) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["href"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHref(val)
        }
        return nil
    }
    return res
}
// GetHref gets the href property value. The URL of the ruleset
// returns a *string when successful
func (m *RepositoryRuleset__links_self) GetHref()(*string) {
    return m.href
}
// Serialize serializes information the current object
func (m *RepositoryRuleset__links_self) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("href", m.GetHref())
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
func (m *RepositoryRuleset__links_self) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHref sets the href property value. The URL of the ruleset
func (m *RepositoryRuleset__links_self) SetHref(value *string)() {
    m.href = value
}
type RepositoryRuleset__links_selfable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHref()(*string)
    SetHref(value *string)()
}
