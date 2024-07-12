package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Collaborator collaborator
type Collaborator struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The avatar_url property
    avatar_url *string
    // The email property
    email *string
    // The events_url property
    events_url *string
    // The followers_url property
    followers_url *string
    // The following_url property
    following_url *string
    // The gists_url property
    gists_url *string
    // The gravatar_id property
    gravatar_id *string
    // The html_url property
    html_url *string
    // The id property
    id *int64
    // The login property
    login *string
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The organizations_url property
    organizations_url *string
    // The permissions property
    permissions Collaborator_permissionsable
    // The received_events_url property
    received_events_url *string
    // The repos_url property
    repos_url *string
    // The role_name property
    role_name *string
    // The site_admin property
    site_admin *bool
    // The starred_url property
    starred_url *string
    // The subscriptions_url property
    subscriptions_url *string
    // The type property
    typeEscaped *string
    // The url property
    url *string
}
// NewCollaborator instantiates a new Collaborator and sets the default values.
func NewCollaborator()(*Collaborator) {
    m := &Collaborator{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCollaboratorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCollaboratorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCollaborator(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Collaborator) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvatarUrl gets the avatar_url property value. The avatar_url property
// returns a *string when successful
func (m *Collaborator) GetAvatarUrl()(*string) {
    return m.avatar_url
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *Collaborator) GetEmail()(*string) {
    return m.email
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *Collaborator) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Collaborator) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["avatar_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvatarUrl(val)
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
    res["events_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventsUrl(val)
        }
        return nil
    }
    res["followers_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFollowersUrl(val)
        }
        return nil
    }
    res["following_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFollowingUrl(val)
        }
        return nil
    }
    res["gists_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGistsUrl(val)
        }
        return nil
    }
    res["gravatar_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGravatarId(val)
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
        val, err := n.GetInt64Value()
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
    res["organizations_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationsUrl(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCollaborator_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(Collaborator_permissionsable))
        }
        return nil
    }
    res["received_events_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedEventsUrl(val)
        }
        return nil
    }
    res["repos_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReposUrl(val)
        }
        return nil
    }
    res["role_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleName(val)
        }
        return nil
    }
    res["site_admin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteAdmin(val)
        }
        return nil
    }
    res["starred_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStarredUrl(val)
        }
        return nil
    }
    res["subscriptions_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionsUrl(val)
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
// GetFollowersUrl gets the followers_url property value. The followers_url property
// returns a *string when successful
func (m *Collaborator) GetFollowersUrl()(*string) {
    return m.followers_url
}
// GetFollowingUrl gets the following_url property value. The following_url property
// returns a *string when successful
func (m *Collaborator) GetFollowingUrl()(*string) {
    return m.following_url
}
// GetGistsUrl gets the gists_url property value. The gists_url property
// returns a *string when successful
func (m *Collaborator) GetGistsUrl()(*string) {
    return m.gists_url
}
// GetGravatarId gets the gravatar_id property value. The gravatar_id property
// returns a *string when successful
func (m *Collaborator) GetGravatarId()(*string) {
    return m.gravatar_id
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Collaborator) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *Collaborator) GetId()(*int64) {
    return m.id
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *Collaborator) GetLogin()(*string) {
    return m.login
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Collaborator) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Collaborator) GetNodeId()(*string) {
    return m.node_id
}
// GetOrganizationsUrl gets the organizations_url property value. The organizations_url property
// returns a *string when successful
func (m *Collaborator) GetOrganizationsUrl()(*string) {
    return m.organizations_url
}
// GetPermissions gets the permissions property value. The permissions property
// returns a Collaborator_permissionsable when successful
func (m *Collaborator) GetPermissions()(Collaborator_permissionsable) {
    return m.permissions
}
// GetReceivedEventsUrl gets the received_events_url property value. The received_events_url property
// returns a *string when successful
func (m *Collaborator) GetReceivedEventsUrl()(*string) {
    return m.received_events_url
}
// GetReposUrl gets the repos_url property value. The repos_url property
// returns a *string when successful
func (m *Collaborator) GetReposUrl()(*string) {
    return m.repos_url
}
// GetRoleName gets the role_name property value. The role_name property
// returns a *string when successful
func (m *Collaborator) GetRoleName()(*string) {
    return m.role_name
}
// GetSiteAdmin gets the site_admin property value. The site_admin property
// returns a *bool when successful
func (m *Collaborator) GetSiteAdmin()(*bool) {
    return m.site_admin
}
// GetStarredUrl gets the starred_url property value. The starred_url property
// returns a *string when successful
func (m *Collaborator) GetStarredUrl()(*string) {
    return m.starred_url
}
// GetSubscriptionsUrl gets the subscriptions_url property value. The subscriptions_url property
// returns a *string when successful
func (m *Collaborator) GetSubscriptionsUrl()(*string) {
    return m.subscriptions_url
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *Collaborator) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *Collaborator) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Collaborator) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("avatar_url", m.GetAvatarUrl())
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
        err := writer.WriteStringValue("events_url", m.GetEventsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("followers_url", m.GetFollowersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("following_url", m.GetFollowingUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("gists_url", m.GetGistsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("gravatar_id", m.GetGravatarId())
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
        err := writer.WriteInt64Value("id", m.GetId())
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
        err := writer.WriteStringValue("organizations_url", m.GetOrganizationsUrl())
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
        err := writer.WriteStringValue("received_events_url", m.GetReceivedEventsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repos_url", m.GetReposUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("role_name", m.GetRoleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("site_admin", m.GetSiteAdmin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("starred_url", m.GetStarredUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptions_url", m.GetSubscriptionsUrl())
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
func (m *Collaborator) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvatarUrl sets the avatar_url property value. The avatar_url property
func (m *Collaborator) SetAvatarUrl(value *string)() {
    m.avatar_url = value
}
// SetEmail sets the email property value. The email property
func (m *Collaborator) SetEmail(value *string)() {
    m.email = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *Collaborator) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetFollowersUrl sets the followers_url property value. The followers_url property
func (m *Collaborator) SetFollowersUrl(value *string)() {
    m.followers_url = value
}
// SetFollowingUrl sets the following_url property value. The following_url property
func (m *Collaborator) SetFollowingUrl(value *string)() {
    m.following_url = value
}
// SetGistsUrl sets the gists_url property value. The gists_url property
func (m *Collaborator) SetGistsUrl(value *string)() {
    m.gists_url = value
}
// SetGravatarId sets the gravatar_id property value. The gravatar_id property
func (m *Collaborator) SetGravatarId(value *string)() {
    m.gravatar_id = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Collaborator) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *Collaborator) SetId(value *int64)() {
    m.id = value
}
// SetLogin sets the login property value. The login property
func (m *Collaborator) SetLogin(value *string)() {
    m.login = value
}
// SetName sets the name property value. The name property
func (m *Collaborator) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Collaborator) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOrganizationsUrl sets the organizations_url property value. The organizations_url property
func (m *Collaborator) SetOrganizationsUrl(value *string)() {
    m.organizations_url = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *Collaborator) SetPermissions(value Collaborator_permissionsable)() {
    m.permissions = value
}
// SetReceivedEventsUrl sets the received_events_url property value. The received_events_url property
func (m *Collaborator) SetReceivedEventsUrl(value *string)() {
    m.received_events_url = value
}
// SetReposUrl sets the repos_url property value. The repos_url property
func (m *Collaborator) SetReposUrl(value *string)() {
    m.repos_url = value
}
// SetRoleName sets the role_name property value. The role_name property
func (m *Collaborator) SetRoleName(value *string)() {
    m.role_name = value
}
// SetSiteAdmin sets the site_admin property value. The site_admin property
func (m *Collaborator) SetSiteAdmin(value *bool)() {
    m.site_admin = value
}
// SetStarredUrl sets the starred_url property value. The starred_url property
func (m *Collaborator) SetStarredUrl(value *string)() {
    m.starred_url = value
}
// SetSubscriptionsUrl sets the subscriptions_url property value. The subscriptions_url property
func (m *Collaborator) SetSubscriptionsUrl(value *string)() {
    m.subscriptions_url = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Collaborator) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The url property
func (m *Collaborator) SetUrl(value *string)() {
    m.url = value
}
type Collaboratorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetEmail()(*string)
    GetEventsUrl()(*string)
    GetFollowersUrl()(*string)
    GetFollowingUrl()(*string)
    GetGistsUrl()(*string)
    GetGravatarId()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetLogin()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetOrganizationsUrl()(*string)
    GetPermissions()(Collaborator_permissionsable)
    GetReceivedEventsUrl()(*string)
    GetReposUrl()(*string)
    GetRoleName()(*string)
    GetSiteAdmin()(*bool)
    GetStarredUrl()(*string)
    GetSubscriptionsUrl()(*string)
    GetTypeEscaped()(*string)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetEmail(value *string)()
    SetEventsUrl(value *string)()
    SetFollowersUrl(value *string)()
    SetFollowingUrl(value *string)()
    SetGistsUrl(value *string)()
    SetGravatarId(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetLogin(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOrganizationsUrl(value *string)()
    SetPermissions(value Collaborator_permissionsable)()
    SetReceivedEventsUrl(value *string)()
    SetReposUrl(value *string)()
    SetRoleName(value *string)()
    SetSiteAdmin(value *bool)()
    SetStarredUrl(value *string)()
    SetSubscriptionsUrl(value *string)()
    SetTypeEscaped(value *string)()
    SetUrl(value *string)()
}
