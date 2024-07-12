package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemImportAuthorsItemWithAuthor_PatchRequestBody struct {
    // The new Git author email.
    email *string
    // The new Git author name.
    name *string
}
// NewItemItemImportAuthorsItemWithAuthor_PatchRequestBody instantiates a new ItemItemImportAuthorsItemWithAuthor_PatchRequestBody and sets the default values.
func NewItemItemImportAuthorsItemWithAuthor_PatchRequestBody()(*ItemItemImportAuthorsItemWithAuthor_PatchRequestBody) {
    m := &ItemItemImportAuthorsItemWithAuthor_PatchRequestBody{
    }
    return m
}
// CreateItemItemImportAuthorsItemWithAuthor_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemImportAuthorsItemWithAuthor_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemImportAuthorsItemWithAuthor_PatchRequestBody(), nil
}
// GetEmail gets the email property value. The new Git author email.
// returns a *string when successful
func (m *ItemItemImportAuthorsItemWithAuthor_PatchRequestBody) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemImportAuthorsItemWithAuthor_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetName gets the name property value. The new Git author name.
// returns a *string when successful
func (m *ItemItemImportAuthorsItemWithAuthor_PatchRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemImportAuthorsItemWithAuthor_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetEmail sets the email property value. The new Git author email.
func (m *ItemItemImportAuthorsItemWithAuthor_PatchRequestBody) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. The new Git author name.
func (m *ItemItemImportAuthorsItemWithAuthor_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
type ItemItemImportAuthorsItemWithAuthor_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetName()(*string)
    SetEmail(value *string)()
    SetName(value *string)()
}
