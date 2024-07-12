package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemAutolinksPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether this autolink reference matches alphanumeric characters. If true, the `<num>` parameter of the `url_template` matches alphanumeric characters `A-Z` (case insensitive), `0-9`, and `-`. If false, this autolink reference only matches numeric characters.
    is_alphanumeric *bool
    // This prefix appended by certain characters will generate a link any time it is found in an issue, pull request, or commit.
    key_prefix *string
    // The URL must contain `<num>` for the reference number. `<num>` matches different characters depending on the value of `is_alphanumeric`.
    url_template *string
}
// NewItemItemAutolinksPostRequestBody instantiates a new ItemItemAutolinksPostRequestBody and sets the default values.
func NewItemItemAutolinksPostRequestBody()(*ItemItemAutolinksPostRequestBody) {
    m := &ItemItemAutolinksPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemAutolinksPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemAutolinksPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemAutolinksPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemAutolinksPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemAutolinksPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["is_alphanumeric"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAlphanumeric(val)
        }
        return nil
    }
    res["key_prefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyPrefix(val)
        }
        return nil
    }
    res["url_template"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrlTemplate(val)
        }
        return nil
    }
    return res
}
// GetIsAlphanumeric gets the is_alphanumeric property value. Whether this autolink reference matches alphanumeric characters. If true, the `<num>` parameter of the `url_template` matches alphanumeric characters `A-Z` (case insensitive), `0-9`, and `-`. If false, this autolink reference only matches numeric characters.
// returns a *bool when successful
func (m *ItemItemAutolinksPostRequestBody) GetIsAlphanumeric()(*bool) {
    return m.is_alphanumeric
}
// GetKeyPrefix gets the key_prefix property value. This prefix appended by certain characters will generate a link any time it is found in an issue, pull request, or commit.
// returns a *string when successful
func (m *ItemItemAutolinksPostRequestBody) GetKeyPrefix()(*string) {
    return m.key_prefix
}
// GetUrlTemplate gets the url_template property value. The URL must contain `<num>` for the reference number. `<num>` matches different characters depending on the value of `is_alphanumeric`.
// returns a *string when successful
func (m *ItemItemAutolinksPostRequestBody) GetUrlTemplate()(*string) {
    return m.url_template
}
// Serialize serializes information the current object
func (m *ItemItemAutolinksPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("is_alphanumeric", m.GetIsAlphanumeric())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key_prefix", m.GetKeyPrefix())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url_template", m.GetUrlTemplate())
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
func (m *ItemItemAutolinksPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsAlphanumeric sets the is_alphanumeric property value. Whether this autolink reference matches alphanumeric characters. If true, the `<num>` parameter of the `url_template` matches alphanumeric characters `A-Z` (case insensitive), `0-9`, and `-`. If false, this autolink reference only matches numeric characters.
func (m *ItemItemAutolinksPostRequestBody) SetIsAlphanumeric(value *bool)() {
    m.is_alphanumeric = value
}
// SetKeyPrefix sets the key_prefix property value. This prefix appended by certain characters will generate a link any time it is found in an issue, pull request, or commit.
func (m *ItemItemAutolinksPostRequestBody) SetKeyPrefix(value *string)() {
    m.key_prefix = value
}
// SetUrlTemplate sets the url_template property value. The URL must contain `<num>` for the reference number. `<num>` matches different characters depending on the value of `is_alphanumeric`.
func (m *ItemItemAutolinksPostRequestBody) SetUrlTemplate(value *string)() {
    m.url_template = value
}
type ItemItemAutolinksPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsAlphanumeric()(*bool)
    GetKeyPrefix()(*string)
    GetUrlTemplate()(*string)
    SetIsAlphanumeric(value *bool)()
    SetKeyPrefix(value *string)()
    SetUrlTemplate(value *string)()
}
