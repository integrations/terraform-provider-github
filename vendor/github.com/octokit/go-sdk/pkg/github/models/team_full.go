package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamFull groups of organization members that gives permissions on specified repositories.
type TeamFull struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The html_url property
    html_url *string
    // Unique identifier of the team
    id *int32
    // Distinguished Name (DN) that team maps to within LDAP environment
    ldap_dn *string
    // The members_count property
    members_count *int32
    // The members_url property
    members_url *string
    // Name of the team
    name *string
    // The node_id property
    node_id *string
    // The notification setting the team has set
    notification_setting *TeamFull_notification_setting
    // Team Organization
    organization TeamOrganizationable
    // Groups of organization members that gives permissions on specified repositories.
    parent NullableTeamSimpleable
    // Permission that the team will have for its repositories
    permission *string
    // The level of privacy this team should have
    privacy *TeamFull_privacy
    // The repos_count property
    repos_count *int32
    // The repositories_url property
    repositories_url *string
    // The slug property
    slug *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // URL for the team
    url *string
}
// NewTeamFull instantiates a new TeamFull and sets the default values.
func NewTeamFull()(*TeamFull) {
    m := &TeamFull{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamFullFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamFullFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamFull(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamFull) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *TeamFull) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TeamFull) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamFull) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["ldap_dn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLdapDn(val)
        }
        return nil
    }
    res["members_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCount(val)
        }
        return nil
    }
    res["members_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersUrl(val)
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
    res["notification_setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamFull_notification_setting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationSetting(val.(*TeamFull_notification_setting))
        }
        return nil
    }
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamOrganizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val.(TeamOrganizationable))
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableTeamSimpleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(NullableTeamSimpleable))
        }
        return nil
    }
    res["permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val)
        }
        return nil
    }
    res["privacy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamFull_privacy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacy(val.(*TeamFull_privacy))
        }
        return nil
    }
    res["repos_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReposCount(val)
        }
        return nil
    }
    res["repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoriesUrl(val)
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
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *TeamFull) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. Unique identifier of the team
// returns a *int32 when successful
func (m *TeamFull) GetId()(*int32) {
    return m.id
}
// GetLdapDn gets the ldap_dn property value. Distinguished Name (DN) that team maps to within LDAP environment
// returns a *string when successful
func (m *TeamFull) GetLdapDn()(*string) {
    return m.ldap_dn
}
// GetMembersCount gets the members_count property value. The members_count property
// returns a *int32 when successful
func (m *TeamFull) GetMembersCount()(*int32) {
    return m.members_count
}
// GetMembersUrl gets the members_url property value. The members_url property
// returns a *string when successful
func (m *TeamFull) GetMembersUrl()(*string) {
    return m.members_url
}
// GetName gets the name property value. Name of the team
// returns a *string when successful
func (m *TeamFull) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *TeamFull) GetNodeId()(*string) {
    return m.node_id
}
// GetNotificationSetting gets the notification_setting property value. The notification setting the team has set
// returns a *TeamFull_notification_setting when successful
func (m *TeamFull) GetNotificationSetting()(*TeamFull_notification_setting) {
    return m.notification_setting
}
// GetOrganization gets the organization property value. Team Organization
// returns a TeamOrganizationable when successful
func (m *TeamFull) GetOrganization()(TeamOrganizationable) {
    return m.organization
}
// GetParent gets the parent property value. Groups of organization members that gives permissions on specified repositories.
// returns a NullableTeamSimpleable when successful
func (m *TeamFull) GetParent()(NullableTeamSimpleable) {
    return m.parent
}
// GetPermission gets the permission property value. Permission that the team will have for its repositories
// returns a *string when successful
func (m *TeamFull) GetPermission()(*string) {
    return m.permission
}
// GetPrivacy gets the privacy property value. The level of privacy this team should have
// returns a *TeamFull_privacy when successful
func (m *TeamFull) GetPrivacy()(*TeamFull_privacy) {
    return m.privacy
}
// GetReposCount gets the repos_count property value. The repos_count property
// returns a *int32 when successful
func (m *TeamFull) GetReposCount()(*int32) {
    return m.repos_count
}
// GetRepositoriesUrl gets the repositories_url property value. The repositories_url property
// returns a *string when successful
func (m *TeamFull) GetRepositoriesUrl()(*string) {
    return m.repositories_url
}
// GetSlug gets the slug property value. The slug property
// returns a *string when successful
func (m *TeamFull) GetSlug()(*string) {
    return m.slug
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *TeamFull) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. URL for the team
// returns a *string when successful
func (m *TeamFull) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *TeamFull) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("ldap_dn", m.GetLdapDn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("members_count", m.GetMembersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("members_url", m.GetMembersUrl())
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
    if m.GetNotificationSetting() != nil {
        cast := (*m.GetNotificationSetting()).String()
        err := writer.WriteStringValue("notification_setting", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("organization", m.GetOrganization())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("permission", m.GetPermission())
        if err != nil {
            return err
        }
    }
    if m.GetPrivacy() != nil {
        cast := (*m.GetPrivacy()).String()
        err := writer.WriteStringValue("privacy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repositories_url", m.GetRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("repos_count", m.GetReposCount())
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
func (m *TeamFull) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *TeamFull) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDescription sets the description property value. The description property
func (m *TeamFull) SetDescription(value *string)() {
    m.description = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *TeamFull) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. Unique identifier of the team
func (m *TeamFull) SetId(value *int32)() {
    m.id = value
}
// SetLdapDn sets the ldap_dn property value. Distinguished Name (DN) that team maps to within LDAP environment
func (m *TeamFull) SetLdapDn(value *string)() {
    m.ldap_dn = value
}
// SetMembersCount sets the members_count property value. The members_count property
func (m *TeamFull) SetMembersCount(value *int32)() {
    m.members_count = value
}
// SetMembersUrl sets the members_url property value. The members_url property
func (m *TeamFull) SetMembersUrl(value *string)() {
    m.members_url = value
}
// SetName sets the name property value. Name of the team
func (m *TeamFull) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *TeamFull) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNotificationSetting sets the notification_setting property value. The notification setting the team has set
func (m *TeamFull) SetNotificationSetting(value *TeamFull_notification_setting)() {
    m.notification_setting = value
}
// SetOrganization sets the organization property value. Team Organization
func (m *TeamFull) SetOrganization(value TeamOrganizationable)() {
    m.organization = value
}
// SetParent sets the parent property value. Groups of organization members that gives permissions on specified repositories.
func (m *TeamFull) SetParent(value NullableTeamSimpleable)() {
    m.parent = value
}
// SetPermission sets the permission property value. Permission that the team will have for its repositories
func (m *TeamFull) SetPermission(value *string)() {
    m.permission = value
}
// SetPrivacy sets the privacy property value. The level of privacy this team should have
func (m *TeamFull) SetPrivacy(value *TeamFull_privacy)() {
    m.privacy = value
}
// SetReposCount sets the repos_count property value. The repos_count property
func (m *TeamFull) SetReposCount(value *int32)() {
    m.repos_count = value
}
// SetRepositoriesUrl sets the repositories_url property value. The repositories_url property
func (m *TeamFull) SetRepositoriesUrl(value *string)() {
    m.repositories_url = value
}
// SetSlug sets the slug property value. The slug property
func (m *TeamFull) SetSlug(value *string)() {
    m.slug = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *TeamFull) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. URL for the team
func (m *TeamFull) SetUrl(value *string)() {
    m.url = value
}
type TeamFullable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLdapDn()(*string)
    GetMembersCount()(*int32)
    GetMembersUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetNotificationSetting()(*TeamFull_notification_setting)
    GetOrganization()(TeamOrganizationable)
    GetParent()(NullableTeamSimpleable)
    GetPermission()(*string)
    GetPrivacy()(*TeamFull_privacy)
    GetReposCount()(*int32)
    GetRepositoriesUrl()(*string)
    GetSlug()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLdapDn(value *string)()
    SetMembersCount(value *int32)()
    SetMembersUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNotificationSetting(value *TeamFull_notification_setting)()
    SetOrganization(value TeamOrganizationable)()
    SetParent(value NullableTeamSimpleable)()
    SetPermission(value *string)()
    SetPrivacy(value *TeamFull_privacy)()
    SetReposCount(value *int32)()
    SetRepositoriesUrl(value *string)()
    SetSlug(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
