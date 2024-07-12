package repos

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemGitTagsPostRequestBody_tagger an object with information about the individual creating the tag.
type ItemItemGitTagsPostRequestBody_tagger struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // When this object was tagged. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The email of the author of the tag
    email *string
    // The name of the author of the tag
    name *string
}
// NewItemItemGitTagsPostRequestBody_tagger instantiates a new ItemItemGitTagsPostRequestBody_tagger and sets the default values.
func NewItemItemGitTagsPostRequestBody_tagger()(*ItemItemGitTagsPostRequestBody_tagger) {
    m := &ItemItemGitTagsPostRequestBody_tagger{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemGitTagsPostRequestBody_taggerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemGitTagsPostRequestBody_taggerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemGitTagsPostRequestBody_tagger(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemGitTagsPostRequestBody_tagger) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDate gets the date property value. When this object was tagged. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *ItemItemGitTagsPostRequestBody_tagger) GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date
}
// GetEmail gets the email property value. The email of the author of the tag
// returns a *string when successful
func (m *ItemItemGitTagsPostRequestBody_tagger) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemGitTagsPostRequestBody_tagger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
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
// GetName gets the name property value. The name of the author of the tag
// returns a *string when successful
func (m *ItemItemGitTagsPostRequestBody_tagger) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemGitTagsPostRequestBody_tagger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("date", m.GetDate())
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
func (m *ItemItemGitTagsPostRequestBody_tagger) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDate sets the date property value. When this object was tagged. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *ItemItemGitTagsPostRequestBody_tagger) SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date = value
}
// SetEmail sets the email property value. The email of the author of the tag
func (m *ItemItemGitTagsPostRequestBody_tagger) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. The name of the author of the tag
func (m *ItemItemGitTagsPostRequestBody_tagger) SetName(value *string)() {
    m.name = value
}
type ItemItemGitTagsPostRequestBody_taggerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmail()(*string)
    GetName()(*string)
    SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmail(value *string)()
    SetName(value *string)()
}
