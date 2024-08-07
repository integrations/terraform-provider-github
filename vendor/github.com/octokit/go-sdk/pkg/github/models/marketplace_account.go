package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MarketplaceAccount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email property
    email *string
    // The id property
    id *int32
    // The login property
    login *string
    // The node_id property
    node_id *string
    // The organization_billing_email property
    organization_billing_email *string
    // The type property
    typeEscaped *string
    // The url property
    url *string
}
// NewMarketplaceAccount instantiates a new MarketplaceAccount and sets the default values.
func NewMarketplaceAccount()(*MarketplaceAccount) {
    m := &MarketplaceAccount{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMarketplaceAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMarketplaceAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMarketplaceAccount(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MarketplaceAccount) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *MarketplaceAccount) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MarketplaceAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["organization_billing_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationBillingEmail(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *MarketplaceAccount) GetId()(*int32) {
    return m.id
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *MarketplaceAccount) GetLogin()(*string) {
    return m.login
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *MarketplaceAccount) GetNodeId()(*string) {
    return m.node_id
}
// GetOrganizationBillingEmail gets the organization_billing_email property value. The organization_billing_email property
// returns a *string when successful
func (m *MarketplaceAccount) GetOrganizationBillingEmail()(*string) {
    return m.organization_billing_email
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *MarketplaceAccount) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *MarketplaceAccount) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *MarketplaceAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("login", m.GetLogin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organization_billing_email", m.GetOrganizationBillingEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *MarketplaceAccount) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. The email property
func (m *MarketplaceAccount) SetEmail(value *string)() {
    m.email = value
}
// SetId sets the id property value. The id property
func (m *MarketplaceAccount) SetId(value *int32)() {
    m.id = value
}
// SetLogin sets the login property value. The login property
func (m *MarketplaceAccount) SetLogin(value *string)() {
    m.login = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *MarketplaceAccount) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOrganizationBillingEmail sets the organization_billing_email property value. The organization_billing_email property
func (m *MarketplaceAccount) SetOrganizationBillingEmail(value *string)() {
    m.organization_billing_email = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *MarketplaceAccount) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The url property
func (m *MarketplaceAccount) SetUrl(value *string)() {
    m.url = value
}
type MarketplaceAccountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetId()(*int32)
    GetLogin()(*string)
    GetNodeId()(*string)
    GetOrganizationBillingEmail()(*string)
    GetTypeEscaped()(*string)
    GetUrl()(*string)
    SetEmail(value *string)()
    SetId(value *int32)()
    SetLogin(value *string)()
    SetNodeId(value *string)()
    SetOrganizationBillingEmail(value *string)()
    SetTypeEscaped(value *string)()
    SetUrl(value *string)()
}
