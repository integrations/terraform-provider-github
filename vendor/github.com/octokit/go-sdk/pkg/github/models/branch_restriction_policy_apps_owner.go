package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BranchRestrictionPolicy_apps_owner struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The avatar_url property
    avatar_url *string
    // The description property
    description *string
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
    // The hooks_url property
    hooks_url *string
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The issues_url property
    issues_url *string
    // The login property
    login *string
    // The members_url property
    members_url *string
    // The node_id property
    node_id *string
    // The organizations_url property
    organizations_url *string
    // The public_members_url property
    public_members_url *string
    // The received_events_url property
    received_events_url *string
    // The repos_url property
    repos_url *string
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
// NewBranchRestrictionPolicy_apps_owner instantiates a new BranchRestrictionPolicy_apps_owner and sets the default values.
func NewBranchRestrictionPolicy_apps_owner()(*BranchRestrictionPolicy_apps_owner) {
    m := &BranchRestrictionPolicy_apps_owner{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBranchRestrictionPolicy_apps_ownerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBranchRestrictionPolicy_apps_ownerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBranchRestrictionPolicy_apps_owner(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BranchRestrictionPolicy_apps_owner) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvatarUrl gets the avatar_url property value. The avatar_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetAvatarUrl()(*string) {
    return m.avatar_url
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetDescription()(*string) {
    return m.description
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BranchRestrictionPolicy_apps_owner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["hooks_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHooksUrl(val)
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
    res["issues_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuesUrl(val)
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
    res["public_members_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicMembersUrl(val)
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
func (m *BranchRestrictionPolicy_apps_owner) GetFollowersUrl()(*string) {
    return m.followers_url
}
// GetFollowingUrl gets the following_url property value. The following_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetFollowingUrl()(*string) {
    return m.following_url
}
// GetGistsUrl gets the gists_url property value. The gists_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetGistsUrl()(*string) {
    return m.gists_url
}
// GetGravatarId gets the gravatar_id property value. The gravatar_id property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetGravatarId()(*string) {
    return m.gravatar_id
}
// GetHooksUrl gets the hooks_url property value. The hooks_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetHooksUrl()(*string) {
    return m.hooks_url
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *BranchRestrictionPolicy_apps_owner) GetId()(*int32) {
    return m.id
}
// GetIssuesUrl gets the issues_url property value. The issues_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetIssuesUrl()(*string) {
    return m.issues_url
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetLogin()(*string) {
    return m.login
}
// GetMembersUrl gets the members_url property value. The members_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetMembersUrl()(*string) {
    return m.members_url
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetNodeId()(*string) {
    return m.node_id
}
// GetOrganizationsUrl gets the organizations_url property value. The organizations_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetOrganizationsUrl()(*string) {
    return m.organizations_url
}
// GetPublicMembersUrl gets the public_members_url property value. The public_members_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetPublicMembersUrl()(*string) {
    return m.public_members_url
}
// GetReceivedEventsUrl gets the received_events_url property value. The received_events_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetReceivedEventsUrl()(*string) {
    return m.received_events_url
}
// GetReposUrl gets the repos_url property value. The repos_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetReposUrl()(*string) {
    return m.repos_url
}
// GetSiteAdmin gets the site_admin property value. The site_admin property
// returns a *bool when successful
func (m *BranchRestrictionPolicy_apps_owner) GetSiteAdmin()(*bool) {
    return m.site_admin
}
// GetStarredUrl gets the starred_url property value. The starred_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetStarredUrl()(*string) {
    return m.starred_url
}
// GetSubscriptionsUrl gets the subscriptions_url property value. The subscriptions_url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetSubscriptionsUrl()(*string) {
    return m.subscriptions_url
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *BranchRestrictionPolicy_apps_owner) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *BranchRestrictionPolicy_apps_owner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("avatar_url", m.GetAvatarUrl())
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
        err := writer.WriteStringValue("hooks_url", m.GetHooksUrl())
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
        err := writer.WriteStringValue("issues_url", m.GetIssuesUrl())
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
        err := writer.WriteStringValue("members_url", m.GetMembersUrl())
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
        err := writer.WriteStringValue("public_members_url", m.GetPublicMembersUrl())
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
func (m *BranchRestrictionPolicy_apps_owner) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvatarUrl sets the avatar_url property value. The avatar_url property
func (m *BranchRestrictionPolicy_apps_owner) SetAvatarUrl(value *string)() {
    m.avatar_url = value
}
// SetDescription sets the description property value. The description property
func (m *BranchRestrictionPolicy_apps_owner) SetDescription(value *string)() {
    m.description = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *BranchRestrictionPolicy_apps_owner) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetFollowersUrl sets the followers_url property value. The followers_url property
func (m *BranchRestrictionPolicy_apps_owner) SetFollowersUrl(value *string)() {
    m.followers_url = value
}
// SetFollowingUrl sets the following_url property value. The following_url property
func (m *BranchRestrictionPolicy_apps_owner) SetFollowingUrl(value *string)() {
    m.following_url = value
}
// SetGistsUrl sets the gists_url property value. The gists_url property
func (m *BranchRestrictionPolicy_apps_owner) SetGistsUrl(value *string)() {
    m.gists_url = value
}
// SetGravatarId sets the gravatar_id property value. The gravatar_id property
func (m *BranchRestrictionPolicy_apps_owner) SetGravatarId(value *string)() {
    m.gravatar_id = value
}
// SetHooksUrl sets the hooks_url property value. The hooks_url property
func (m *BranchRestrictionPolicy_apps_owner) SetHooksUrl(value *string)() {
    m.hooks_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *BranchRestrictionPolicy_apps_owner) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *BranchRestrictionPolicy_apps_owner) SetId(value *int32)() {
    m.id = value
}
// SetIssuesUrl sets the issues_url property value. The issues_url property
func (m *BranchRestrictionPolicy_apps_owner) SetIssuesUrl(value *string)() {
    m.issues_url = value
}
// SetLogin sets the login property value. The login property
func (m *BranchRestrictionPolicy_apps_owner) SetLogin(value *string)() {
    m.login = value
}
// SetMembersUrl sets the members_url property value. The members_url property
func (m *BranchRestrictionPolicy_apps_owner) SetMembersUrl(value *string)() {
    m.members_url = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *BranchRestrictionPolicy_apps_owner) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOrganizationsUrl sets the organizations_url property value. The organizations_url property
func (m *BranchRestrictionPolicy_apps_owner) SetOrganizationsUrl(value *string)() {
    m.organizations_url = value
}
// SetPublicMembersUrl sets the public_members_url property value. The public_members_url property
func (m *BranchRestrictionPolicy_apps_owner) SetPublicMembersUrl(value *string)() {
    m.public_members_url = value
}
// SetReceivedEventsUrl sets the received_events_url property value. The received_events_url property
func (m *BranchRestrictionPolicy_apps_owner) SetReceivedEventsUrl(value *string)() {
    m.received_events_url = value
}
// SetReposUrl sets the repos_url property value. The repos_url property
func (m *BranchRestrictionPolicy_apps_owner) SetReposUrl(value *string)() {
    m.repos_url = value
}
// SetSiteAdmin sets the site_admin property value. The site_admin property
func (m *BranchRestrictionPolicy_apps_owner) SetSiteAdmin(value *bool)() {
    m.site_admin = value
}
// SetStarredUrl sets the starred_url property value. The starred_url property
func (m *BranchRestrictionPolicy_apps_owner) SetStarredUrl(value *string)() {
    m.starred_url = value
}
// SetSubscriptionsUrl sets the subscriptions_url property value. The subscriptions_url property
func (m *BranchRestrictionPolicy_apps_owner) SetSubscriptionsUrl(value *string)() {
    m.subscriptions_url = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *BranchRestrictionPolicy_apps_owner) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The url property
func (m *BranchRestrictionPolicy_apps_owner) SetUrl(value *string)() {
    m.url = value
}
type BranchRestrictionPolicy_apps_ownerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetDescription()(*string)
    GetEventsUrl()(*string)
    GetFollowersUrl()(*string)
    GetFollowingUrl()(*string)
    GetGistsUrl()(*string)
    GetGravatarId()(*string)
    GetHooksUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetIssuesUrl()(*string)
    GetLogin()(*string)
    GetMembersUrl()(*string)
    GetNodeId()(*string)
    GetOrganizationsUrl()(*string)
    GetPublicMembersUrl()(*string)
    GetReceivedEventsUrl()(*string)
    GetReposUrl()(*string)
    GetSiteAdmin()(*bool)
    GetStarredUrl()(*string)
    GetSubscriptionsUrl()(*string)
    GetTypeEscaped()(*string)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetDescription(value *string)()
    SetEventsUrl(value *string)()
    SetFollowersUrl(value *string)()
    SetFollowingUrl(value *string)()
    SetGistsUrl(value *string)()
    SetGravatarId(value *string)()
    SetHooksUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetIssuesUrl(value *string)()
    SetLogin(value *string)()
    SetMembersUrl(value *string)()
    SetNodeId(value *string)()
    SetOrganizationsUrl(value *string)()
    SetPublicMembersUrl(value *string)()
    SetReceivedEventsUrl(value *string)()
    SetReposUrl(value *string)()
    SetSiteAdmin(value *bool)()
    SetStarredUrl(value *string)()
    SetSubscriptionsUrl(value *string)()
    SetTypeEscaped(value *string)()
    SetUrl(value *string)()
}
