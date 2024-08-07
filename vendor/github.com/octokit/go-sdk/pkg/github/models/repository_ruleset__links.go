package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleset__links struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The html property
    html RepositoryRuleset__links_htmlable
    // The self property
    self RepositoryRuleset__links_selfable
}
// NewRepositoryRuleset__links instantiates a new RepositoryRuleset__links and sets the default values.
func NewRepositoryRuleset__links()(*RepositoryRuleset__links) {
    m := &RepositoryRuleset__links{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleset__linksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleset__linksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleset__links(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleset__links) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleset__links) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleset__links_htmlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtml(val.(RepositoryRuleset__links_htmlable))
        }
        return nil
    }
    res["self"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleset__links_selfFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelf(val.(RepositoryRuleset__links_selfable))
        }
        return nil
    }
    return res
}
// GetHtml gets the html property value. The html property
// returns a RepositoryRuleset__links_htmlable when successful
func (m *RepositoryRuleset__links) GetHtml()(RepositoryRuleset__links_htmlable) {
    return m.html
}
// GetSelf gets the self property value. The self property
// returns a RepositoryRuleset__links_selfable when successful
func (m *RepositoryRuleset__links) GetSelf()(RepositoryRuleset__links_selfable) {
    return m.self
}
// Serialize serializes information the current object
func (m *RepositoryRuleset__links) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("html", m.GetHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("self", m.GetSelf())
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
func (m *RepositoryRuleset__links) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHtml sets the html property value. The html property
func (m *RepositoryRuleset__links) SetHtml(value RepositoryRuleset__links_htmlable)() {
    m.html = value
}
// SetSelf sets the self property value. The self property
func (m *RepositoryRuleset__links) SetSelf(value RepositoryRuleset__links_selfable)() {
    m.self = value
}
type RepositoryRuleset__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtml()(RepositoryRuleset__links_htmlable)
    GetSelf()(RepositoryRuleset__links_selfable)
    SetHtml(value RepositoryRuleset__links_htmlable)()
    SetSelf(value RepositoryRuleset__links_selfable)()
}
