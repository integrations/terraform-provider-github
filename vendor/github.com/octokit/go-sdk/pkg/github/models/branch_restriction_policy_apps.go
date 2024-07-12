package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BranchRestrictionPolicy_apps struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *string
    // The description property
    description *string
    // The events property
    events []string
    // The external_url property
    external_url *string
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The owner property
    owner BranchRestrictionPolicy_apps_ownerable
    // The permissions property
    permissions BranchRestrictionPolicy_apps_permissionsable
    // The slug property
    slug *string
    // The updated_at property
    updated_at *string
}
// NewBranchRestrictionPolicy_apps instantiates a new BranchRestrictionPolicy_apps and sets the default values.
func NewBranchRestrictionPolicy_apps()(*BranchRestrictionPolicy_apps) {
    m := &BranchRestrictionPolicy_apps{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBranchRestrictionPolicy_appsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBranchRestrictionPolicy_appsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBranchRestrictionPolicy_apps(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BranchRestrictionPolicy_apps) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps) GetCreatedAt()(*string) {
    return m.created_at
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps) GetDescription()(*string) {
    return m.description
}
// GetEvents gets the events property value. The events property
// returns a []string when successful
func (m *BranchRestrictionPolicy_apps) GetEvents()([]string) {
    return m.events
}
// GetExternalUrl gets the external_url property value. The external_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps) GetExternalUrl()(*string) {
    return m.external_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BranchRestrictionPolicy_apps) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
        val, err := n.GetObjectValue(CreateBranchRestrictionPolicy_apps_ownerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(BranchRestrictionPolicy_apps_ownerable))
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBranchRestrictionPolicy_apps_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(BranchRestrictionPolicy_apps_permissionsable))
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *BranchRestrictionPolicy_apps) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps) GetNodeId()(*string) {
    return m.node_id
}
// GetOwner gets the owner property value. The owner property
// returns a BranchRestrictionPolicy_apps_ownerable when successful
func (m *BranchRestrictionPolicy_apps) GetOwner()(BranchRestrictionPolicy_apps_ownerable) {
    return m.owner
}
// GetPermissions gets the permissions property value. The permissions property
// returns a BranchRestrictionPolicy_apps_permissionsable when successful
func (m *BranchRestrictionPolicy_apps) GetPermissions()(BranchRestrictionPolicy_apps_permissionsable) {
    return m.permissions
}
// GetSlug gets the slug property value. The slug property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps) GetSlug()(*string) {
    return m.slug
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps) GetUpdatedAt()(*string) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *BranchRestrictionPolicy_apps) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
func (m *BranchRestrictionPolicy_apps) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *BranchRestrictionPolicy_apps) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetDescription sets the description property value. The description property
func (m *BranchRestrictionPolicy_apps) SetDescription(value *string)() {
    m.description = value
}
// SetEvents sets the events property value. The events property
func (m *BranchRestrictionPolicy_apps) SetEvents(value []string)() {
    m.events = value
}
// SetExternalUrl sets the external_url property value. The external_url property
func (m *BranchRestrictionPolicy_apps) SetExternalUrl(value *string)() {
    m.external_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *BranchRestrictionPolicy_apps) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *BranchRestrictionPolicy_apps) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *BranchRestrictionPolicy_apps) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *BranchRestrictionPolicy_apps) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOwner sets the owner property value. The owner property
func (m *BranchRestrictionPolicy_apps) SetOwner(value BranchRestrictionPolicy_apps_ownerable)() {
    m.owner = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *BranchRestrictionPolicy_apps) SetPermissions(value BranchRestrictionPolicy_apps_permissionsable)() {
    m.permissions = value
}
// SetSlug sets the slug property value. The slug property
func (m *BranchRestrictionPolicy_apps) SetSlug(value *string)() {
    m.slug = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *BranchRestrictionPolicy_apps) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
type BranchRestrictionPolicy_appsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetDescription()(*string)
    GetEvents()([]string)
    GetExternalUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetOwner()(BranchRestrictionPolicy_apps_ownerable)
    GetPermissions()(BranchRestrictionPolicy_apps_permissionsable)
    GetSlug()(*string)
    GetUpdatedAt()(*string)
    SetCreatedAt(value *string)()
    SetDescription(value *string)()
    SetEvents(value []string)()
    SetExternalUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOwner(value BranchRestrictionPolicy_apps_ownerable)()
    SetPermissions(value BranchRestrictionPolicy_apps_permissionsable)()
    SetSlug(value *string)()
    SetUpdatedAt(value *string)()
}
