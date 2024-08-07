package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemContentsItemWithPathDeleteRequestBody_committer object containing information about the committer.
type ItemItemContentsItemWithPathDeleteRequestBody_committer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email of the author (or committer) of the commit
    email *string
    // The name of the author (or committer) of the commit
    name *string
}
// NewItemItemContentsItemWithPathDeleteRequestBody_committer instantiates a new ItemItemContentsItemWithPathDeleteRequestBody_committer and sets the default values.
func NewItemItemContentsItemWithPathDeleteRequestBody_committer()(*ItemItemContentsItemWithPathDeleteRequestBody_committer) {
    m := &ItemItemContentsItemWithPathDeleteRequestBody_committer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemContentsItemWithPathDeleteRequestBody_committerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemContentsItemWithPathDeleteRequestBody_committerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemContentsItemWithPathDeleteRequestBody_committer(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody_committer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. The email of the author (or committer) of the commit
// returns a *string when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody_committer) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody_committer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
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
    return res
}
// GetName gets the name property value. The name of the author (or committer) of the commit
// returns a *string when successful
func (m *ItemItemContentsItemWithPathDeleteRequestBody_committer) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemContentsItemWithPathDeleteRequestBody_committer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *ItemItemContentsItemWithPathDeleteRequestBody_committer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. The email of the author (or committer) of the commit
func (m *ItemItemContentsItemWithPathDeleteRequestBody_committer) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. The name of the author (or committer) of the commit
func (m *ItemItemContentsItemWithPathDeleteRequestBody_committer) SetName(value *string)() {
    m.name = value
}
type ItemItemContentsItemWithPathDeleteRequestBody_committerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetName()(*string)
    SetEmail(value *string)()
    SetName(value *string)()
}
