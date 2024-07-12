package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemPagesPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specify a custom domain for the repository. Sending a `null` value will remove the custom domain. For more about custom domains, see "[Using a custom domain with GitHub Pages](https://docs.github.com/pages/configuring-a-custom-domain-for-your-github-pages-site)."
    cname *string
    // Specify whether HTTPS should be enforced for the repository.
    https_enforced *bool
    // The source property
    source ItemItemPagesPutRequestBody_PagesPutRequestBody_sourceable
}
// ItemItemPagesPutRequestBody_PagesPutRequestBody_source composed type wrapper for classes ItemItemPagesPutRequestBody_sourceMember1able, string
type ItemItemPagesPutRequestBody_PagesPutRequestBody_source struct {
    // Composed type representation for type ItemItemPagesPutRequestBody_sourceMember1able
    itemItemPagesPutRequestBody_sourceMember1 ItemItemPagesPutRequestBody_sourceMember1able
    // Composed type representation for type string
    string *string
}
// NewItemItemPagesPutRequestBody_PagesPutRequestBody_source instantiates a new ItemItemPagesPutRequestBody_PagesPutRequestBody_source and sets the default values.
func NewItemItemPagesPutRequestBody_PagesPutRequestBody_source()(*ItemItemPagesPutRequestBody_PagesPutRequestBody_source) {
    m := &ItemItemPagesPutRequestBody_PagesPutRequestBody_source{
    }
    return m
}
// CreateItemItemPagesPutRequestBody_PagesPutRequestBody_sourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPagesPutRequestBody_PagesPutRequestBody_sourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemItemPagesPutRequestBody_PagesPutRequestBody_source()
    if parseNode != nil {
        if val, err := parseNode.GetStringValue(); val != nil {
            if err != nil {
                return nil, err
            }
            result.SetString(val)
        } else if val, err := parseNode.GetObjectValue(CreateItemItemPagesPutRequestBody_sourceMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemPagesPutRequestBody_sourceMember1able); ok {
                result.SetItemItemPagesPutRequestBodySourceMember1(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPagesPutRequestBody_PagesPutRequestBody_source) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemItemPagesPutRequestBody_PagesPutRequestBody_source) GetIsComposedType()(bool) {
    return true
}
// GetItemItemPagesPutRequestBodySourceMember1 gets the ItemItemPagesPutRequestBody_sourceMember1 property value. Composed type representation for type ItemItemPagesPutRequestBody_sourceMember1able
// returns a ItemItemPagesPutRequestBody_sourceMember1able when successful
func (m *ItemItemPagesPutRequestBody_PagesPutRequestBody_source) GetItemItemPagesPutRequestBodySourceMember1()(ItemItemPagesPutRequestBody_sourceMember1able) {
    return m.itemItemPagesPutRequestBody_sourceMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ItemItemPagesPutRequestBody_PagesPutRequestBody_source) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ItemItemPagesPutRequestBody_PagesPutRequestBody_source) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetItemItemPagesPutRequestBodySourceMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemPagesPutRequestBodySourceMember1())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemPagesPutRequestBodySourceMember1 sets the ItemItemPagesPutRequestBody_sourceMember1 property value. Composed type representation for type ItemItemPagesPutRequestBody_sourceMember1able
func (m *ItemItemPagesPutRequestBody_PagesPutRequestBody_source) SetItemItemPagesPutRequestBodySourceMember1(value ItemItemPagesPutRequestBody_sourceMember1able)() {
    m.itemItemPagesPutRequestBody_sourceMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemItemPagesPutRequestBody_PagesPutRequestBody_source) SetString(value *string)() {
    m.string = value
}
type ItemItemPagesPutRequestBody_PagesPutRequestBody_sourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemPagesPutRequestBodySourceMember1()(ItemItemPagesPutRequestBody_sourceMember1able)
    GetString()(*string)
    SetItemItemPagesPutRequestBodySourceMember1(value ItemItemPagesPutRequestBody_sourceMember1able)()
    SetString(value *string)()
}
// NewItemItemPagesPutRequestBody instantiates a new ItemItemPagesPutRequestBody and sets the default values.
func NewItemItemPagesPutRequestBody()(*ItemItemPagesPutRequestBody) {
    m := &ItemItemPagesPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPagesPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPagesPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPagesPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPagesPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCname gets the cname property value. Specify a custom domain for the repository. Sending a `null` value will remove the custom domain. For more about custom domains, see "[Using a custom domain with GitHub Pages](https://docs.github.com/pages/configuring-a-custom-domain-for-your-github-pages-site)."
// returns a *string when successful
func (m *ItemItemPagesPutRequestBody) GetCname()(*string) {
    return m.cname
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPagesPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCname(val)
        }
        return nil
    }
    res["https_enforced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpsEnforced(val)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemPagesPutRequestBody_PagesPutRequestBody_sourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(ItemItemPagesPutRequestBody_PagesPutRequestBody_sourceable))
        }
        return nil
    }
    return res
}
// GetHttpsEnforced gets the https_enforced property value. Specify whether HTTPS should be enforced for the repository.
// returns a *bool when successful
func (m *ItemItemPagesPutRequestBody) GetHttpsEnforced()(*bool) {
    return m.https_enforced
}
// GetSource gets the source property value. The source property
// returns a ItemItemPagesPutRequestBody_PagesPutRequestBody_sourceable when successful
func (m *ItemItemPagesPutRequestBody) GetSource()(ItemItemPagesPutRequestBody_PagesPutRequestBody_sourceable) {
    return m.source
}
// Serialize serializes information the current object
func (m *ItemItemPagesPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cname", m.GetCname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("https_enforced", m.GetHttpsEnforced())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("source", m.GetSource())
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
func (m *ItemItemPagesPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCname sets the cname property value. Specify a custom domain for the repository. Sending a `null` value will remove the custom domain. For more about custom domains, see "[Using a custom domain with GitHub Pages](https://docs.github.com/pages/configuring-a-custom-domain-for-your-github-pages-site)."
func (m *ItemItemPagesPutRequestBody) SetCname(value *string)() {
    m.cname = value
}
// SetHttpsEnforced sets the https_enforced property value. Specify whether HTTPS should be enforced for the repository.
func (m *ItemItemPagesPutRequestBody) SetHttpsEnforced(value *bool)() {
    m.https_enforced = value
}
// SetSource sets the source property value. The source property
func (m *ItemItemPagesPutRequestBody) SetSource(value ItemItemPagesPutRequestBody_PagesPutRequestBody_sourceable)() {
    m.source = value
}
type ItemItemPagesPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCname()(*string)
    GetHttpsEnforced()(*bool)
    GetSource()(ItemItemPagesPutRequestBody_PagesPutRequestBody_sourceable)
    SetCname(value *string)()
    SetHttpsEnforced(value *bool)()
    SetSource(value ItemItemPagesPutRequestBody_PagesPutRequestBody_sourceable)()
}
