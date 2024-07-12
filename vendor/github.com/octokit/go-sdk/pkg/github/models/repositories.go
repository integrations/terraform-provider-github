package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Repositories struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The fragment property
    fragment *string
    // The matches property
    matches []Repositories_matchesable
    // The object_type property
    object_type *string
    // The object_url property
    object_url *string
    // The property property
    property *string
}
// NewRepositories instantiates a new Repositories and sets the default values.
func NewRepositories()(*Repositories) {
    m := &Repositories{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoriesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoriesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositories(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Repositories) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Repositories) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fragment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFragment(val)
        }
        return nil
    }
    res["matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositories_matchesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Repositories_matchesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Repositories_matchesable)
                }
            }
            m.SetMatches(res)
        }
        return nil
    }
    res["object_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectType(val)
        }
        return nil
    }
    res["object_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectUrl(val)
        }
        return nil
    }
    res["property"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperty(val)
        }
        return nil
    }
    return res
}
// GetFragment gets the fragment property value. The fragment property
// returns a *string when successful
func (m *Repositories) GetFragment()(*string) {
    return m.fragment
}
// GetMatches gets the matches property value. The matches property
// returns a []Repositories_matchesable when successful
func (m *Repositories) GetMatches()([]Repositories_matchesable) {
    return m.matches
}
// GetObjectType gets the object_type property value. The object_type property
// returns a *string when successful
func (m *Repositories) GetObjectType()(*string) {
    return m.object_type
}
// GetObjectUrl gets the object_url property value. The object_url property
// returns a *string when successful
func (m *Repositories) GetObjectUrl()(*string) {
    return m.object_url
}
// GetProperty gets the property property value. The property property
// returns a *string when successful
func (m *Repositories) GetProperty()(*string) {
    return m.property
}
// Serialize serializes information the current object
func (m *Repositories) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fragment", m.GetFragment())
        if err != nil {
            return err
        }
    }
    if m.GetMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatches()))
        for i, v := range m.GetMatches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("matches", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("object_type", m.GetObjectType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("object_url", m.GetObjectUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("property", m.GetProperty())
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
func (m *Repositories) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFragment sets the fragment property value. The fragment property
func (m *Repositories) SetFragment(value *string)() {
    m.fragment = value
}
// SetMatches sets the matches property value. The matches property
func (m *Repositories) SetMatches(value []Repositories_matchesable)() {
    m.matches = value
}
// SetObjectType sets the object_type property value. The object_type property
func (m *Repositories) SetObjectType(value *string)() {
    m.object_type = value
}
// SetObjectUrl sets the object_url property value. The object_url property
func (m *Repositories) SetObjectUrl(value *string)() {
    m.object_url = value
}
// SetProperty sets the property property value. The property property
func (m *Repositories) SetProperty(value *string)() {
    m.property = value
}
type Repositoriesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFragment()(*string)
    GetMatches()([]Repositories_matchesable)
    GetObjectType()(*string)
    GetObjectUrl()(*string)
    GetProperty()(*string)
    SetFragment(value *string)()
    SetMatches(value []Repositories_matchesable)()
    SetObjectType(value *string)()
    SetObjectUrl(value *string)()
    SetProperty(value *string)()
}
