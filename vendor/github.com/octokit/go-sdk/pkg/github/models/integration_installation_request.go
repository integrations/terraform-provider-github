package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntegrationInstallationRequest request to install an integration on a target
type IntegrationInstallationRequest struct {
    // The account property
    account IntegrationInstallationRequest_IntegrationInstallationRequest_accountable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Unique identifier of the request installation.
    id *int32
    // The node_id property
    node_id *string
    // A GitHub user.
    requester SimpleUserable
}
// IntegrationInstallationRequest_IntegrationInstallationRequest_account composed type wrapper for classes Enterpriseable, SimpleUserable
type IntegrationInstallationRequest_IntegrationInstallationRequest_account struct {
    // Composed type representation for type Enterpriseable
    enterprise Enterpriseable
    // Composed type representation for type SimpleUserable
    simpleUser SimpleUserable
}
// NewIntegrationInstallationRequest_IntegrationInstallationRequest_account instantiates a new IntegrationInstallationRequest_IntegrationInstallationRequest_account and sets the default values.
func NewIntegrationInstallationRequest_IntegrationInstallationRequest_account()(*IntegrationInstallationRequest_IntegrationInstallationRequest_account) {
    m := &IntegrationInstallationRequest_IntegrationInstallationRequest_account{
    }
    return m
}
// CreateIntegrationInstallationRequest_IntegrationInstallationRequest_accountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntegrationInstallationRequest_IntegrationInstallationRequest_accountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewIntegrationInstallationRequest_IntegrationInstallationRequest_account()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateEnterpriseFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(Enterpriseable); ok {
                result.SetEnterprise(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateSimpleUserFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(SimpleUserable); ok {
                result.SetSimpleUser(cast)
            }
        }
    }
    return result, nil
}
// GetEnterprise gets the enterprise property value. Composed type representation for type Enterpriseable
// returns a Enterpriseable when successful
func (m *IntegrationInstallationRequest_IntegrationInstallationRequest_account) GetEnterprise()(Enterpriseable) {
    return m.enterprise
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IntegrationInstallationRequest_IntegrationInstallationRequest_account) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *IntegrationInstallationRequest_IntegrationInstallationRequest_account) GetIsComposedType()(bool) {
    return true
}
// GetSimpleUser gets the simpleUser property value. Composed type representation for type SimpleUserable
// returns a SimpleUserable when successful
func (m *IntegrationInstallationRequest_IntegrationInstallationRequest_account) GetSimpleUser()(SimpleUserable) {
    return m.simpleUser
}
// Serialize serializes information the current object
func (m *IntegrationInstallationRequest_IntegrationInstallationRequest_account) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnterprise() != nil {
        err := writer.WriteObjectValue("", m.GetEnterprise())
        if err != nil {
            return err
        }
    } else if m.GetSimpleUser() != nil {
        err := writer.WriteObjectValue("", m.GetSimpleUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnterprise sets the enterprise property value. Composed type representation for type Enterpriseable
func (m *IntegrationInstallationRequest_IntegrationInstallationRequest_account) SetEnterprise(value Enterpriseable)() {
    m.enterprise = value
}
// SetSimpleUser sets the simpleUser property value. Composed type representation for type SimpleUserable
func (m *IntegrationInstallationRequest_IntegrationInstallationRequest_account) SetSimpleUser(value SimpleUserable)() {
    m.simpleUser = value
}
type IntegrationInstallationRequest_IntegrationInstallationRequest_accountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnterprise()(Enterpriseable)
    GetSimpleUser()(SimpleUserable)
    SetEnterprise(value Enterpriseable)()
    SetSimpleUser(value SimpleUserable)()
}
// NewIntegrationInstallationRequest instantiates a new IntegrationInstallationRequest and sets the default values.
func NewIntegrationInstallationRequest()(*IntegrationInstallationRequest) {
    m := &IntegrationInstallationRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntegrationInstallationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntegrationInstallationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntegrationInstallationRequest(), nil
}
// GetAccount gets the account property value. The account property
// returns a IntegrationInstallationRequest_IntegrationInstallationRequest_accountable when successful
func (m *IntegrationInstallationRequest) GetAccount()(IntegrationInstallationRequest_IntegrationInstallationRequest_accountable) {
    return m.account
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IntegrationInstallationRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *IntegrationInstallationRequest) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IntegrationInstallationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntegrationInstallationRequest_IntegrationInstallationRequest_accountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(IntegrationInstallationRequest_IntegrationInstallationRequest_accountable))
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
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
    res["requester"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequester(val.(SimpleUserable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique identifier of the request installation.
// returns a *int32 when successful
func (m *IntegrationInstallationRequest) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *IntegrationInstallationRequest) GetNodeId()(*string) {
    return m.node_id
}
// GetRequester gets the requester property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *IntegrationInstallationRequest) GetRequester()(SimpleUserable) {
    return m.requester
}
// Serialize serializes information the current object
func (m *IntegrationInstallationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requester", m.GetRequester())
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
// SetAccount sets the account property value. The account property
func (m *IntegrationInstallationRequest) SetAccount(value IntegrationInstallationRequest_IntegrationInstallationRequest_accountable)() {
    m.account = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntegrationInstallationRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *IntegrationInstallationRequest) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetId sets the id property value. Unique identifier of the request installation.
func (m *IntegrationInstallationRequest) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *IntegrationInstallationRequest) SetNodeId(value *string)() {
    m.node_id = value
}
// SetRequester sets the requester property value. A GitHub user.
func (m *IntegrationInstallationRequest) SetRequester(value SimpleUserable)() {
    m.requester = value
}
type IntegrationInstallationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(IntegrationInstallationRequest_IntegrationInstallationRequest_accountable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int32)
    GetNodeId()(*string)
    GetRequester()(SimpleUserable)
    SetAccount(value IntegrationInstallationRequest_IntegrationInstallationRequest_accountable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetRequester(value SimpleUserable)()
}
