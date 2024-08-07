package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemContentsItemWithPathPutRequestBody_author the author of the file. Default: The `committer` or the authenticated user if you omit `committer`.
type ItemItemContentsItemWithPathPutRequestBody_author struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date property
    date *string
    // The email of the author or committer of the commit. You'll receive a `422` status code if `email` is omitted.
    email *string
    // The name of the author or committer of the commit. You'll receive a `422` status code if `name` is omitted.
    name *string
}
// NewItemItemContentsItemWithPathPutRequestBody_author instantiates a new ItemItemContentsItemWithPathPutRequestBody_author and sets the default values.
func NewItemItemContentsItemWithPathPutRequestBody_author()(*ItemItemContentsItemWithPathPutRequestBody_author) {
    m := &ItemItemContentsItemWithPathPutRequestBody_author{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemContentsItemWithPathPutRequestBody_authorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemContentsItemWithPathPutRequestBody_authorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemContentsItemWithPathPutRequestBody_author(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemContentsItemWithPathPutRequestBody_author) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDate gets the date property value. The date property
// returns a *string when successful
func (m *ItemItemContentsItemWithPathPutRequestBody_author) GetDate()(*string) {
    return m.date
}
// GetEmail gets the email property value. The email of the author or committer of the commit. You'll receive a `422` status code if `email` is omitted.
// returns a *string when successful
func (m *ItemItemContentsItemWithPathPutRequestBody_author) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemContentsItemWithPathPutRequestBody_author) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
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
// GetName gets the name property value. The name of the author or committer of the commit. You'll receive a `422` status code if `name` is omitted.
// returns a *string when successful
func (m *ItemItemContentsItemWithPathPutRequestBody_author) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemContentsItemWithPathPutRequestBody_author) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
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
func (m *ItemItemContentsItemWithPathPutRequestBody_author) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDate sets the date property value. The date property
func (m *ItemItemContentsItemWithPathPutRequestBody_author) SetDate(value *string)() {
    m.date = value
}
// SetEmail sets the email property value. The email of the author or committer of the commit. You'll receive a `422` status code if `email` is omitted.
func (m *ItemItemContentsItemWithPathPutRequestBody_author) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. The name of the author or committer of the commit. You'll receive a `422` status code if `name` is omitted.
func (m *ItemItemContentsItemWithPathPutRequestBody_author) SetName(value *string)() {
    m.name = value
}
type ItemItemContentsItemWithPathPutRequestBody_authorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDate()(*string)
    GetEmail()(*string)
    GetName()(*string)
    SetDate(value *string)()
    SetEmail(value *string)()
    SetName(value *string)()
}
