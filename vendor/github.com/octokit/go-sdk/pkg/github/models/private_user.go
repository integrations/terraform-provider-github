package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivateUser private User
type PrivateUser struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The avatar_url property
    avatar_url *string
    // The bio property
    bio *string
    // The blog property
    blog *string
    // The business_plus property
    business_plus *bool
    // The collaborators property
    collaborators *int32
    // The company property
    company *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The disk_usage property
    disk_usage *int32
    // The email property
    email *string
    // The events_url property
    events_url *string
    // The followers property
    followers *int32
    // The followers_url property
    followers_url *string
    // The following property
    following *int32
    // The following_url property
    following_url *string
    // The gists_url property
    gists_url *string
    // The gravatar_id property
    gravatar_id *string
    // The hireable property
    hireable *bool
    // The html_url property
    html_url *string
    // The id property
    id *int64
    // The ldap_dn property
    ldap_dn *string
    // The location property
    location *string
    // The login property
    login *string
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The notification_email property
    notification_email *string
    // The organizations_url property
    organizations_url *string
    // The owned_private_repos property
    owned_private_repos *int32
    // The plan property
    plan PrivateUser_planable
    // The private_gists property
    private_gists *int32
    // The public_gists property
    public_gists *int32
    // The public_repos property
    public_repos *int32
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
    // The suspended_at property
    suspended_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The total_private_repos property
    total_private_repos *int32
    // The twitter_username property
    twitter_username *string
    // The two_factor_authentication property
    two_factor_authentication *bool
    // The type property
    typeEscaped *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewPrivateUser instantiates a new PrivateUser and sets the default values.
func NewPrivateUser()(*PrivateUser) {
    m := &PrivateUser{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrivateUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivateUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivateUser(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PrivateUser) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvatarUrl gets the avatar_url property value. The avatar_url property
// returns a *string when successful
func (m *PrivateUser) GetAvatarUrl()(*string) {
    return m.avatar_url
}
// GetBio gets the bio property value. The bio property
// returns a *string when successful
func (m *PrivateUser) GetBio()(*string) {
    return m.bio
}
// GetBlog gets the blog property value. The blog property
// returns a *string when successful
func (m *PrivateUser) GetBlog()(*string) {
    return m.blog
}
// GetBusinessPlus gets the business_plus property value. The business_plus property
// returns a *bool when successful
func (m *PrivateUser) GetBusinessPlus()(*bool) {
    return m.business_plus
}
// GetCollaborators gets the collaborators property value. The collaborators property
// returns a *int32 when successful
func (m *PrivateUser) GetCollaborators()(*int32) {
    return m.collaborators
}
// GetCompany gets the company property value. The company property
// returns a *string when successful
func (m *PrivateUser) GetCompany()(*string) {
    return m.company
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *PrivateUser) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDiskUsage gets the disk_usage property value. The disk_usage property
// returns a *int32 when successful
func (m *PrivateUser) GetDiskUsage()(*int32) {
    return m.disk_usage
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *PrivateUser) GetEmail()(*string) {
    return m.email
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *PrivateUser) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivateUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["bio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBio(val)
        }
        return nil
    }
    res["blog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlog(val)
        }
        return nil
    }
    res["business_plus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessPlus(val)
        }
        return nil
    }
    res["collaborators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollaborators(val)
        }
        return nil
    }
    res["company"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompany(val)
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
    res["disk_usage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskUsage(val)
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
    res["followers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFollowers(val)
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
    res["following"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFollowing(val)
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
    res["hireable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHireable(val)
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
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
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
    res["notification_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationEmail(val)
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
    res["owned_private_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnedPrivateRepos(val)
        }
        return nil
    }
    res["plan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivateUser_planFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlan(val.(PrivateUser_planable))
        }
        return nil
    }
    res["private_gists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateGists(val)
        }
        return nil
    }
    res["public_gists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicGists(val)
        }
        return nil
    }
    res["public_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicRepos(val)
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
    res["suspended_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuspendedAt(val)
        }
        return nil
    }
    res["total_private_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPrivateRepos(val)
        }
        return nil
    }
    res["twitter_username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwitterUsername(val)
        }
        return nil
    }
    res["two_factor_authentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwoFactorAuthentication(val)
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
// GetFollowers gets the followers property value. The followers property
// returns a *int32 when successful
func (m *PrivateUser) GetFollowers()(*int32) {
    return m.followers
}
// GetFollowersUrl gets the followers_url property value. The followers_url property
// returns a *string when successful
func (m *PrivateUser) GetFollowersUrl()(*string) {
    return m.followers_url
}
// GetFollowing gets the following property value. The following property
// returns a *int32 when successful
func (m *PrivateUser) GetFollowing()(*int32) {
    return m.following
}
// GetFollowingUrl gets the following_url property value. The following_url property
// returns a *string when successful
func (m *PrivateUser) GetFollowingUrl()(*string) {
    return m.following_url
}
// GetGistsUrl gets the gists_url property value. The gists_url property
// returns a *string when successful
func (m *PrivateUser) GetGistsUrl()(*string) {
    return m.gists_url
}
// GetGravatarId gets the gravatar_id property value. The gravatar_id property
// returns a *string when successful
func (m *PrivateUser) GetGravatarId()(*string) {
    return m.gravatar_id
}
// GetHireable gets the hireable property value. The hireable property
// returns a *bool when successful
func (m *PrivateUser) GetHireable()(*bool) {
    return m.hireable
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *PrivateUser) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *PrivateUser) GetId()(*int64) {
    return m.id
}
// GetLdapDn gets the ldap_dn property value. The ldap_dn property
// returns a *string when successful
func (m *PrivateUser) GetLdapDn()(*string) {
    return m.ldap_dn
}
// GetLocation gets the location property value. The location property
// returns a *string when successful
func (m *PrivateUser) GetLocation()(*string) {
    return m.location
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *PrivateUser) GetLogin()(*string) {
    return m.login
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *PrivateUser) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *PrivateUser) GetNodeId()(*string) {
    return m.node_id
}
// GetNotificationEmail gets the notification_email property value. The notification_email property
// returns a *string when successful
func (m *PrivateUser) GetNotificationEmail()(*string) {
    return m.notification_email
}
// GetOrganizationsUrl gets the organizations_url property value. The organizations_url property
// returns a *string when successful
func (m *PrivateUser) GetOrganizationsUrl()(*string) {
    return m.organizations_url
}
// GetOwnedPrivateRepos gets the owned_private_repos property value. The owned_private_repos property
// returns a *int32 when successful
func (m *PrivateUser) GetOwnedPrivateRepos()(*int32) {
    return m.owned_private_repos
}
// GetPlan gets the plan property value. The plan property
// returns a PrivateUser_planable when successful
func (m *PrivateUser) GetPlan()(PrivateUser_planable) {
    return m.plan
}
// GetPrivateGists gets the private_gists property value. The private_gists property
// returns a *int32 when successful
func (m *PrivateUser) GetPrivateGists()(*int32) {
    return m.private_gists
}
// GetPublicGists gets the public_gists property value. The public_gists property
// returns a *int32 when successful
func (m *PrivateUser) GetPublicGists()(*int32) {
    return m.public_gists
}
// GetPublicRepos gets the public_repos property value. The public_repos property
// returns a *int32 when successful
func (m *PrivateUser) GetPublicRepos()(*int32) {
    return m.public_repos
}
// GetReceivedEventsUrl gets the received_events_url property value. The received_events_url property
// returns a *string when successful
func (m *PrivateUser) GetReceivedEventsUrl()(*string) {
    return m.received_events_url
}
// GetReposUrl gets the repos_url property value. The repos_url property
// returns a *string when successful
func (m *PrivateUser) GetReposUrl()(*string) {
    return m.repos_url
}
// GetSiteAdmin gets the site_admin property value. The site_admin property
// returns a *bool when successful
func (m *PrivateUser) GetSiteAdmin()(*bool) {
    return m.site_admin
}
// GetStarredUrl gets the starred_url property value. The starred_url property
// returns a *string when successful
func (m *PrivateUser) GetStarredUrl()(*string) {
    return m.starred_url
}
// GetSubscriptionsUrl gets the subscriptions_url property value. The subscriptions_url property
// returns a *string when successful
func (m *PrivateUser) GetSubscriptionsUrl()(*string) {
    return m.subscriptions_url
}
// GetSuspendedAt gets the suspended_at property value. The suspended_at property
// returns a *Time when successful
func (m *PrivateUser) GetSuspendedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.suspended_at
}
// GetTotalPrivateRepos gets the total_private_repos property value. The total_private_repos property
// returns a *int32 when successful
func (m *PrivateUser) GetTotalPrivateRepos()(*int32) {
    return m.total_private_repos
}
// GetTwitterUsername gets the twitter_username property value. The twitter_username property
// returns a *string when successful
func (m *PrivateUser) GetTwitterUsername()(*string) {
    return m.twitter_username
}
// GetTwoFactorAuthentication gets the two_factor_authentication property value. The two_factor_authentication property
// returns a *bool when successful
func (m *PrivateUser) GetTwoFactorAuthentication()(*bool) {
    return m.two_factor_authentication
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *PrivateUser) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *PrivateUser) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PrivateUser) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PrivateUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("avatar_url", m.GetAvatarUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bio", m.GetBio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blog", m.GetBlog())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("business_plus", m.GetBusinessPlus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("collaborators", m.GetCollaborators())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("company", m.GetCompany())
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
        err := writer.WriteInt32Value("disk_usage", m.GetDiskUsage())
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
        err := writer.WriteInt32Value("followers", m.GetFollowers())
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
        err := writer.WriteInt32Value("following", m.GetFollowing())
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
        err := writer.WriteBoolValue("hireable", m.GetHireable())
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
        err := writer.WriteStringValue("ldap_dn", m.GetLdapDn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
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
        err := writer.WriteStringValue("notification_email", m.GetNotificationEmail())
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
        err := writer.WriteInt32Value("owned_private_repos", m.GetOwnedPrivateRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("plan", m.GetPlan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("private_gists", m.GetPrivateGists())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("public_gists", m.GetPublicGists())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("public_repos", m.GetPublicRepos())
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
        err := writer.WriteTimeValue("suspended_at", m.GetSuspendedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_private_repos", m.GetTotalPrivateRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("twitter_username", m.GetTwitterUsername())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("two_factor_authentication", m.GetTwoFactorAuthentication())
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
func (m *PrivateUser) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvatarUrl sets the avatar_url property value. The avatar_url property
func (m *PrivateUser) SetAvatarUrl(value *string)() {
    m.avatar_url = value
}
// SetBio sets the bio property value. The bio property
func (m *PrivateUser) SetBio(value *string)() {
    m.bio = value
}
// SetBlog sets the blog property value. The blog property
func (m *PrivateUser) SetBlog(value *string)() {
    m.blog = value
}
// SetBusinessPlus sets the business_plus property value. The business_plus property
func (m *PrivateUser) SetBusinessPlus(value *bool)() {
    m.business_plus = value
}
// SetCollaborators sets the collaborators property value. The collaborators property
func (m *PrivateUser) SetCollaborators(value *int32)() {
    m.collaborators = value
}
// SetCompany sets the company property value. The company property
func (m *PrivateUser) SetCompany(value *string)() {
    m.company = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *PrivateUser) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDiskUsage sets the disk_usage property value. The disk_usage property
func (m *PrivateUser) SetDiskUsage(value *int32)() {
    m.disk_usage = value
}
// SetEmail sets the email property value. The email property
func (m *PrivateUser) SetEmail(value *string)() {
    m.email = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *PrivateUser) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetFollowers sets the followers property value. The followers property
func (m *PrivateUser) SetFollowers(value *int32)() {
    m.followers = value
}
// SetFollowersUrl sets the followers_url property value. The followers_url property
func (m *PrivateUser) SetFollowersUrl(value *string)() {
    m.followers_url = value
}
// SetFollowing sets the following property value. The following property
func (m *PrivateUser) SetFollowing(value *int32)() {
    m.following = value
}
// SetFollowingUrl sets the following_url property value. The following_url property
func (m *PrivateUser) SetFollowingUrl(value *string)() {
    m.following_url = value
}
// SetGistsUrl sets the gists_url property value. The gists_url property
func (m *PrivateUser) SetGistsUrl(value *string)() {
    m.gists_url = value
}
// SetGravatarId sets the gravatar_id property value. The gravatar_id property
func (m *PrivateUser) SetGravatarId(value *string)() {
    m.gravatar_id = value
}
// SetHireable sets the hireable property value. The hireable property
func (m *PrivateUser) SetHireable(value *bool)() {
    m.hireable = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *PrivateUser) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *PrivateUser) SetId(value *int64)() {
    m.id = value
}
// SetLdapDn sets the ldap_dn property value. The ldap_dn property
func (m *PrivateUser) SetLdapDn(value *string)() {
    m.ldap_dn = value
}
// SetLocation sets the location property value. The location property
func (m *PrivateUser) SetLocation(value *string)() {
    m.location = value
}
// SetLogin sets the login property value. The login property
func (m *PrivateUser) SetLogin(value *string)() {
    m.login = value
}
// SetName sets the name property value. The name property
func (m *PrivateUser) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *PrivateUser) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNotificationEmail sets the notification_email property value. The notification_email property
func (m *PrivateUser) SetNotificationEmail(value *string)() {
    m.notification_email = value
}
// SetOrganizationsUrl sets the organizations_url property value. The organizations_url property
func (m *PrivateUser) SetOrganizationsUrl(value *string)() {
    m.organizations_url = value
}
// SetOwnedPrivateRepos sets the owned_private_repos property value. The owned_private_repos property
func (m *PrivateUser) SetOwnedPrivateRepos(value *int32)() {
    m.owned_private_repos = value
}
// SetPlan sets the plan property value. The plan property
func (m *PrivateUser) SetPlan(value PrivateUser_planable)() {
    m.plan = value
}
// SetPrivateGists sets the private_gists property value. The private_gists property
func (m *PrivateUser) SetPrivateGists(value *int32)() {
    m.private_gists = value
}
// SetPublicGists sets the public_gists property value. The public_gists property
func (m *PrivateUser) SetPublicGists(value *int32)() {
    m.public_gists = value
}
// SetPublicRepos sets the public_repos property value. The public_repos property
func (m *PrivateUser) SetPublicRepos(value *int32)() {
    m.public_repos = value
}
// SetReceivedEventsUrl sets the received_events_url property value. The received_events_url property
func (m *PrivateUser) SetReceivedEventsUrl(value *string)() {
    m.received_events_url = value
}
// SetReposUrl sets the repos_url property value. The repos_url property
func (m *PrivateUser) SetReposUrl(value *string)() {
    m.repos_url = value
}
// SetSiteAdmin sets the site_admin property value. The site_admin property
func (m *PrivateUser) SetSiteAdmin(value *bool)() {
    m.site_admin = value
}
// SetStarredUrl sets the starred_url property value. The starred_url property
func (m *PrivateUser) SetStarredUrl(value *string)() {
    m.starred_url = value
}
// SetSubscriptionsUrl sets the subscriptions_url property value. The subscriptions_url property
func (m *PrivateUser) SetSubscriptionsUrl(value *string)() {
    m.subscriptions_url = value
}
// SetSuspendedAt sets the suspended_at property value. The suspended_at property
func (m *PrivateUser) SetSuspendedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.suspended_at = value
}
// SetTotalPrivateRepos sets the total_private_repos property value. The total_private_repos property
func (m *PrivateUser) SetTotalPrivateRepos(value *int32)() {
    m.total_private_repos = value
}
// SetTwitterUsername sets the twitter_username property value. The twitter_username property
func (m *PrivateUser) SetTwitterUsername(value *string)() {
    m.twitter_username = value
}
// SetTwoFactorAuthentication sets the two_factor_authentication property value. The two_factor_authentication property
func (m *PrivateUser) SetTwoFactorAuthentication(value *bool)() {
    m.two_factor_authentication = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *PrivateUser) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *PrivateUser) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *PrivateUser) SetUrl(value *string)() {
    m.url = value
}
type PrivateUserable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetBio()(*string)
    GetBlog()(*string)
    GetBusinessPlus()(*bool)
    GetCollaborators()(*int32)
    GetCompany()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDiskUsage()(*int32)
    GetEmail()(*string)
    GetEventsUrl()(*string)
    GetFollowers()(*int32)
    GetFollowersUrl()(*string)
    GetFollowing()(*int32)
    GetFollowingUrl()(*string)
    GetGistsUrl()(*string)
    GetGravatarId()(*string)
    GetHireable()(*bool)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetLdapDn()(*string)
    GetLocation()(*string)
    GetLogin()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetNotificationEmail()(*string)
    GetOrganizationsUrl()(*string)
    GetOwnedPrivateRepos()(*int32)
    GetPlan()(PrivateUser_planable)
    GetPrivateGists()(*int32)
    GetPublicGists()(*int32)
    GetPublicRepos()(*int32)
    GetReceivedEventsUrl()(*string)
    GetReposUrl()(*string)
    GetSiteAdmin()(*bool)
    GetStarredUrl()(*string)
    GetSubscriptionsUrl()(*string)
    GetSuspendedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTotalPrivateRepos()(*int32)
    GetTwitterUsername()(*string)
    GetTwoFactorAuthentication()(*bool)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetBio(value *string)()
    SetBlog(value *string)()
    SetBusinessPlus(value *bool)()
    SetCollaborators(value *int32)()
    SetCompany(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDiskUsage(value *int32)()
    SetEmail(value *string)()
    SetEventsUrl(value *string)()
    SetFollowers(value *int32)()
    SetFollowersUrl(value *string)()
    SetFollowing(value *int32)()
    SetFollowingUrl(value *string)()
    SetGistsUrl(value *string)()
    SetGravatarId(value *string)()
    SetHireable(value *bool)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetLdapDn(value *string)()
    SetLocation(value *string)()
    SetLogin(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNotificationEmail(value *string)()
    SetOrganizationsUrl(value *string)()
    SetOwnedPrivateRepos(value *int32)()
    SetPlan(value PrivateUser_planable)()
    SetPrivateGists(value *int32)()
    SetPublicGists(value *int32)()
    SetPublicRepos(value *int32)()
    SetReceivedEventsUrl(value *string)()
    SetReposUrl(value *string)()
    SetSiteAdmin(value *bool)()
    SetStarredUrl(value *string)()
    SetSubscriptionsUrl(value *string)()
    SetSuspendedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTotalPrivateRepos(value *int32)()
    SetTwitterUsername(value *string)()
    SetTwoFactorAuthentication(value *bool)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
