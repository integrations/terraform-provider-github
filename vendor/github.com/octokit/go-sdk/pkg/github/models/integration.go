package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Integration gitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
type Integration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The client_id property
    client_id *string
    // The client_secret property
    client_secret *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The list of events for the GitHub app
    events []string
    // The external_url property
    external_url *string
    // The html_url property
    html_url *string
    // Unique identifier of the GitHub app
    id *int32
    // The number of installations associated with the GitHub app
    installations_count *int32
    // The name of the GitHub app
    name *string
    // The node_id property
    node_id *string
    // A GitHub user.
    owner NullableSimpleUserable
    // The pem property
    pem *string
    // The set of permissions for the GitHub app
    permissions Integration_permissionsable
    // The slug name of the GitHub app
    slug *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The webhook_secret property
    webhook_secret *string
}
// NewIntegration instantiates a new Integration and sets the default values.
func NewIntegration()(*Integration) {
    m := &Integration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntegrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntegrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntegration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Integration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the client_id property value. The client_id property
// returns a *string when successful
func (m *Integration) GetClientId()(*string) {
    return m.client_id
}
// GetClientSecret gets the client_secret property value. The client_secret property
// returns a *string when successful
func (m *Integration) GetClientSecret()(*string) {
    return m.client_secret
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Integration) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Integration) GetDescription()(*string) {
    return m.description
}
// GetEvents gets the events property value. The list of events for the GitHub app
// returns a []string when successful
func (m *Integration) GetEvents()([]string) {
    return m.events
}
// GetExternalUrl gets the external_url property value. The external_url property
// returns a *string when successful
func (m *Integration) GetExternalUrl()(*string) {
    return m.external_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Integration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["client_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["client_secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["external_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUrl(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
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
    res["installations_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallationsCount(val)
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["pem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPem(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntegration_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(Integration_permissionsable))
        }
        return nil
    }
    res["slug"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSlug(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["webhook_secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebhookSecret(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Integration) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. Unique identifier of the GitHub app
// returns a *int32 when successful
func (m *Integration) GetId()(*int32) {
    return m.id
}
// GetInstallationsCount gets the installations_count property value. The number of installations associated with the GitHub app
// returns a *int32 when successful
func (m *Integration) GetInstallationsCount()(*int32) {
    return m.installations_count
}
// GetName gets the name property value. The name of the GitHub app
// returns a *string when successful
func (m *Integration) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Integration) GetNodeId()(*string) {
    return m.node_id
}
// GetOwner gets the owner property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Integration) GetOwner()(NullableSimpleUserable) {
    return m.owner
}
// GetPem gets the pem property value. The pem property
// returns a *string when successful
func (m *Integration) GetPem()(*string) {
    return m.pem
}
// GetPermissions gets the permissions property value. The set of permissions for the GitHub app
// returns a Integration_permissionsable when successful
func (m *Integration) GetPermissions()(Integration_permissionsable) {
    return m.permissions
}
// GetSlug gets the slug property value. The slug name of the GitHub app
// returns a *string when successful
func (m *Integration) GetSlug()(*string) {
    return m.slug
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Integration) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetWebhookSecret gets the webhook_secret property value. The webhook_secret property
// returns a *string when successful
func (m *Integration) GetWebhookSecret()(*string) {
    return m.webhook_secret
}
// Serialize serializes information the current object
func (m *Integration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client_id", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("client_secret", m.GetClientSecret())
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
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        err := writer.WriteCollectionOfStringValues("events", m.GetEvents())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("external_url", m.GetExternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
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
        err := writer.WriteInt32Value("installations_count", m.GetInstallationsCount())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pem", m.GetPem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("slug", m.GetSlug())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webhook_secret", m.GetWebhookSecret())
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
func (m *Integration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the client_id property value. The client_id property
func (m *Integration) SetClientId(value *string)() {
    m.client_id = value
}
// SetClientSecret sets the client_secret property value. The client_secret property
func (m *Integration) SetClientSecret(value *string)() {
    m.client_secret = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Integration) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDescription sets the description property value. The description property
func (m *Integration) SetDescription(value *string)() {
    m.description = value
}
// SetEvents sets the events property value. The list of events for the GitHub app
func (m *Integration) SetEvents(value []string)() {
    m.events = value
}
// SetExternalUrl sets the external_url property value. The external_url property
func (m *Integration) SetExternalUrl(value *string)() {
    m.external_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Integration) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. Unique identifier of the GitHub app
func (m *Integration) SetId(value *int32)() {
    m.id = value
}
// SetInstallationsCount sets the installations_count property value. The number of installations associated with the GitHub app
func (m *Integration) SetInstallationsCount(value *int32)() {
    m.installations_count = value
}
// SetName sets the name property value. The name of the GitHub app
func (m *Integration) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Integration) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *Integration) SetOwner(value NullableSimpleUserable)() {
    m.owner = value
}
// SetPem sets the pem property value. The pem property
func (m *Integration) SetPem(value *string)() {
    m.pem = value
}
// SetPermissions sets the permissions property value. The set of permissions for the GitHub app
func (m *Integration) SetPermissions(value Integration_permissionsable)() {
    m.permissions = value
}
// SetSlug sets the slug property value. The slug name of the GitHub app
func (m *Integration) SetSlug(value *string)() {
    m.slug = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Integration) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetWebhookSecret sets the webhook_secret property value. The webhook_secret property
func (m *Integration) SetWebhookSecret(value *string)() {
    m.webhook_secret = value
}
type Integrationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetEvents()([]string)
    GetExternalUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetInstallationsCount()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetOwner()(NullableSimpleUserable)
    GetPem()(*string)
    GetPermissions()(Integration_permissionsable)
    GetSlug()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetWebhookSecret()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetEvents(value []string)()
    SetExternalUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetInstallationsCount(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOwner(value NullableSimpleUserable)()
    SetPem(value *string)()
    SetPermissions(value Integration_permissionsable)()
    SetSlug(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetWebhookSecret(value *string)()
}
