package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Organization gitHub account for managing multiple users, teams, and repositories
type Organization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The avatar_url property
    avatar_url *string
    // Display blog url for the organization
    blog *string
    // Display company name for the organization
    company *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // Display email for the organization
    email *string
    // The events_url property
    events_url *string
    // The followers property
    followers *int32
    // The following property
    following *int32
    // Specifies if organization projects are enabled for this org
    has_organization_projects *bool
    // Specifies if repository projects are enabled for repositories that belong to this org
    has_repository_projects *bool
    // The hooks_url property
    hooks_url *string
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The is_verified property
    is_verified *bool
    // The issues_url property
    issues_url *string
    // Display location for the organization
    location *string
    // Unique login name of the organization
    login *string
    // The members_url property
    members_url *string
    // Display name for the organization
    name *string
    // The node_id property
    node_id *string
    // The plan property
    plan Organization_planable
    // The public_gists property
    public_gists *int32
    // The public_members_url property
    public_members_url *string
    // The public_repos property
    public_repos *int32
    // The repos_url property
    repos_url *string
    // The type property
    typeEscaped *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // URL for the organization
    url *string
}
// NewOrganization instantiates a new Organization and sets the default values.
func NewOrganization()(*Organization) {
    m := &Organization{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganization(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Organization) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvatarUrl gets the avatar_url property value. The avatar_url property
// returns a *string when successful
func (m *Organization) GetAvatarUrl()(*string) {
    return m.avatar_url
}
// GetBlog gets the blog property value. Display blog url for the organization
// returns a *string when successful
func (m *Organization) GetBlog()(*string) {
    return m.blog
}
// GetCompany gets the company property value. Display company name for the organization
// returns a *string when successful
func (m *Organization) GetCompany()(*string) {
    return m.company
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Organization) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Organization) GetDescription()(*string) {
    return m.description
}
// GetEmail gets the email property value. Display email for the organization
// returns a *string when successful
func (m *Organization) GetEmail()(*string) {
    return m.email
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *Organization) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Organization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["has_organization_projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasOrganizationProjects(val)
        }
        return nil
    }
    res["has_repository_projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasRepositoryProjects(val)
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
    res["is_verified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVerified(val)
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
    res["plan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganization_planFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlan(val.(Organization_planable))
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
func (m *Organization) GetFollowers()(*int32) {
    return m.followers
}
// GetFollowing gets the following property value. The following property
// returns a *int32 when successful
func (m *Organization) GetFollowing()(*int32) {
    return m.following
}
// GetHasOrganizationProjects gets the has_organization_projects property value. Specifies if organization projects are enabled for this org
// returns a *bool when successful
func (m *Organization) GetHasOrganizationProjects()(*bool) {
    return m.has_organization_projects
}
// GetHasRepositoryProjects gets the has_repository_projects property value. Specifies if repository projects are enabled for repositories that belong to this org
// returns a *bool when successful
func (m *Organization) GetHasRepositoryProjects()(*bool) {
    return m.has_repository_projects
}
// GetHooksUrl gets the hooks_url property value. The hooks_url property
// returns a *string when successful
func (m *Organization) GetHooksUrl()(*string) {
    return m.hooks_url
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Organization) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Organization) GetId()(*int32) {
    return m.id
}
// GetIssuesUrl gets the issues_url property value. The issues_url property
// returns a *string when successful
func (m *Organization) GetIssuesUrl()(*string) {
    return m.issues_url
}
// GetIsVerified gets the is_verified property value. The is_verified property
// returns a *bool when successful
func (m *Organization) GetIsVerified()(*bool) {
    return m.is_verified
}
// GetLocation gets the location property value. Display location for the organization
// returns a *string when successful
func (m *Organization) GetLocation()(*string) {
    return m.location
}
// GetLogin gets the login property value. Unique login name of the organization
// returns a *string when successful
func (m *Organization) GetLogin()(*string) {
    return m.login
}
// GetMembersUrl gets the members_url property value. The members_url property
// returns a *string when successful
func (m *Organization) GetMembersUrl()(*string) {
    return m.members_url
}
// GetName gets the name property value. Display name for the organization
// returns a *string when successful
func (m *Organization) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Organization) GetNodeId()(*string) {
    return m.node_id
}
// GetPlan gets the plan property value. The plan property
// returns a Organization_planable when successful
func (m *Organization) GetPlan()(Organization_planable) {
    return m.plan
}
// GetPublicGists gets the public_gists property value. The public_gists property
// returns a *int32 when successful
func (m *Organization) GetPublicGists()(*int32) {
    return m.public_gists
}
// GetPublicMembersUrl gets the public_members_url property value. The public_members_url property
// returns a *string when successful
func (m *Organization) GetPublicMembersUrl()(*string) {
    return m.public_members_url
}
// GetPublicRepos gets the public_repos property value. The public_repos property
// returns a *int32 when successful
func (m *Organization) GetPublicRepos()(*int32) {
    return m.public_repos
}
// GetReposUrl gets the repos_url property value. The repos_url property
// returns a *string when successful
func (m *Organization) GetReposUrl()(*string) {
    return m.repos_url
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *Organization) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Organization) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. URL for the organization
// returns a *string when successful
func (m *Organization) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Organization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("avatar_url", m.GetAvatarUrl())
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
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteInt32Value("following", m.GetFollowing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_organization_projects", m.GetHasOrganizationProjects())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_repository_projects", m.GetHasRepositoryProjects())
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
        err := writer.WriteBoolValue("is_verified", m.GetIsVerified())
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
    {
        err := writer.WriteObjectValue("plan", m.GetPlan())
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
        err := writer.WriteStringValue("public_members_url", m.GetPublicMembersUrl())
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
        err := writer.WriteStringValue("repos_url", m.GetReposUrl())
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
func (m *Organization) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvatarUrl sets the avatar_url property value. The avatar_url property
func (m *Organization) SetAvatarUrl(value *string)() {
    m.avatar_url = value
}
// SetBlog sets the blog property value. Display blog url for the organization
func (m *Organization) SetBlog(value *string)() {
    m.blog = value
}
// SetCompany sets the company property value. Display company name for the organization
func (m *Organization) SetCompany(value *string)() {
    m.company = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Organization) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDescription sets the description property value. The description property
func (m *Organization) SetDescription(value *string)() {
    m.description = value
}
// SetEmail sets the email property value. Display email for the organization
func (m *Organization) SetEmail(value *string)() {
    m.email = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *Organization) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetFollowers sets the followers property value. The followers property
func (m *Organization) SetFollowers(value *int32)() {
    m.followers = value
}
// SetFollowing sets the following property value. The following property
func (m *Organization) SetFollowing(value *int32)() {
    m.following = value
}
// SetHasOrganizationProjects sets the has_organization_projects property value. Specifies if organization projects are enabled for this org
func (m *Organization) SetHasOrganizationProjects(value *bool)() {
    m.has_organization_projects = value
}
// SetHasRepositoryProjects sets the has_repository_projects property value. Specifies if repository projects are enabled for repositories that belong to this org
func (m *Organization) SetHasRepositoryProjects(value *bool)() {
    m.has_repository_projects = value
}
// SetHooksUrl sets the hooks_url property value. The hooks_url property
func (m *Organization) SetHooksUrl(value *string)() {
    m.hooks_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Organization) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *Organization) SetId(value *int32)() {
    m.id = value
}
// SetIssuesUrl sets the issues_url property value. The issues_url property
func (m *Organization) SetIssuesUrl(value *string)() {
    m.issues_url = value
}
// SetIsVerified sets the is_verified property value. The is_verified property
func (m *Organization) SetIsVerified(value *bool)() {
    m.is_verified = value
}
// SetLocation sets the location property value. Display location for the organization
func (m *Organization) SetLocation(value *string)() {
    m.location = value
}
// SetLogin sets the login property value. Unique login name of the organization
func (m *Organization) SetLogin(value *string)() {
    m.login = value
}
// SetMembersUrl sets the members_url property value. The members_url property
func (m *Organization) SetMembersUrl(value *string)() {
    m.members_url = value
}
// SetName sets the name property value. Display name for the organization
func (m *Organization) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Organization) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPlan sets the plan property value. The plan property
func (m *Organization) SetPlan(value Organization_planable)() {
    m.plan = value
}
// SetPublicGists sets the public_gists property value. The public_gists property
func (m *Organization) SetPublicGists(value *int32)() {
    m.public_gists = value
}
// SetPublicMembersUrl sets the public_members_url property value. The public_members_url property
func (m *Organization) SetPublicMembersUrl(value *string)() {
    m.public_members_url = value
}
// SetPublicRepos sets the public_repos property value. The public_repos property
func (m *Organization) SetPublicRepos(value *int32)() {
    m.public_repos = value
}
// SetReposUrl sets the repos_url property value. The repos_url property
func (m *Organization) SetReposUrl(value *string)() {
    m.repos_url = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Organization) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Organization) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. URL for the organization
func (m *Organization) SetUrl(value *string)() {
    m.url = value
}
type Organizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetBlog()(*string)
    GetCompany()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetEmail()(*string)
    GetEventsUrl()(*string)
    GetFollowers()(*int32)
    GetFollowing()(*int32)
    GetHasOrganizationProjects()(*bool)
    GetHasRepositoryProjects()(*bool)
    GetHooksUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetIssuesUrl()(*string)
    GetIsVerified()(*bool)
    GetLocation()(*string)
    GetLogin()(*string)
    GetMembersUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetPlan()(Organization_planable)
    GetPublicGists()(*int32)
    GetPublicMembersUrl()(*string)
    GetPublicRepos()(*int32)
    GetReposUrl()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetBlog(value *string)()
    SetCompany(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetEmail(value *string)()
    SetEventsUrl(value *string)()
    SetFollowers(value *int32)()
    SetFollowing(value *int32)()
    SetHasOrganizationProjects(value *bool)()
    SetHasRepositoryProjects(value *bool)()
    SetHooksUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetIssuesUrl(value *string)()
    SetIsVerified(value *bool)()
    SetLocation(value *string)()
    SetLogin(value *string)()
    SetMembersUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetPlan(value Organization_planable)()
    SetPublicGists(value *int32)()
    SetPublicMembersUrl(value *string)()
    SetPublicRepos(value *int32)()
    SetReposUrl(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
