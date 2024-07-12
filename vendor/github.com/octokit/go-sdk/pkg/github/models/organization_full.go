package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationFull organization Full
type OrganizationFull struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization.This field is only visible to organization owners or members of a team with the security manager role.
    advanced_security_enabled_for_new_repositories *bool
    // The archived_at property
    archived_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The avatar_url property
    avatar_url *string
    // The billing_email property
    billing_email *string
    // The blog property
    blog *string
    // The collaborators property
    collaborators *int32
    // The company property
    company *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The default_repository_permission property
    default_repository_permission *string
    // Whether GitHub Advanced Security is automatically enabled for new repositories and repositories transferred tothis organization.This field is only visible to organization owners or members of a team with the security manager role.
    dependabot_alerts_enabled_for_new_repositories *bool
    // Whether dependabot security updates are automatically enabled for new repositories and repositories transferredto this organization.This field is only visible to organization owners or members of a team with the security manager role.
    dependabot_security_updates_enabled_for_new_repositories *bool
    // Whether dependency graph is automatically enabled for new repositories and repositories transferred to thisorganization.This field is only visible to organization owners or members of a team with the security manager role.
    dependency_graph_enabled_for_new_repositories *bool
    // The description property
    description *string
    // The disk_usage property
    disk_usage *int32
    // The email property
    email *string
    // The events_url property
    events_url *string
    // The followers property
    followers *int32
    // The following property
    following *int32
    // The has_organization_projects property
    has_organization_projects *bool
    // The has_repository_projects property
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
    // The location property
    location *string
    // The login property
    login *string
    // The members_allowed_repository_creation_type property
    members_allowed_repository_creation_type *string
    // The members_can_create_internal_repositories property
    members_can_create_internal_repositories *bool
    // The members_can_create_pages property
    members_can_create_pages *bool
    // The members_can_create_private_pages property
    members_can_create_private_pages *bool
    // The members_can_create_private_repositories property
    members_can_create_private_repositories *bool
    // The members_can_create_public_pages property
    members_can_create_public_pages *bool
    // The members_can_create_public_repositories property
    members_can_create_public_repositories *bool
    // The members_can_create_repositories property
    members_can_create_repositories *bool
    // The members_can_fork_private_repositories property
    members_can_fork_private_repositories *bool
    // The members_url property
    members_url *string
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The owned_private_repos property
    owned_private_repos *int32
    // The plan property
    plan OrganizationFull_planable
    // The private_gists property
    private_gists *int32
    // The public_gists property
    public_gists *int32
    // The public_members_url property
    public_members_url *string
    // The public_repos property
    public_repos *int32
    // The repos_url property
    repos_url *string
    // Whether secret scanning is automatically enabled for new repositories and repositories transferred to thisorganization.This field is only visible to organization owners or members of a team with the security manager role.
    secret_scanning_enabled_for_new_repositories *bool
    // An optional URL string to display to contributors who are blocked from pushing a secret.
    secret_scanning_push_protection_custom_link *string
    // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
    secret_scanning_push_protection_custom_link_enabled *bool
    // Whether secret scanning push protection is automatically enabled for new repositories and repositoriestransferred to this organization.This field is only visible to organization owners or members of a team with the security manager role.
    secret_scanning_push_protection_enabled_for_new_repositories *bool
    // The total_private_repos property
    total_private_repos *int32
    // The twitter_username property
    twitter_username *string
    // The two_factor_requirement_enabled property
    two_factor_requirement_enabled *bool
    // The type property
    typeEscaped *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // The web_commit_signoff_required property
    web_commit_signoff_required *bool
}
// NewOrganizationFull instantiates a new OrganizationFull and sets the default values.
func NewOrganizationFull()(*OrganizationFull) {
    m := &OrganizationFull{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationFullFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationFullFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationFull(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationFull) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurityEnabledForNewRepositories gets the advanced_security_enabled_for_new_repositories property value. Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization.This field is only visible to organization owners or members of a team with the security manager role.
// returns a *bool when successful
func (m *OrganizationFull) GetAdvancedSecurityEnabledForNewRepositories()(*bool) {
    return m.advanced_security_enabled_for_new_repositories
}
// GetArchivedAt gets the archived_at property value. The archived_at property
// returns a *Time when successful
func (m *OrganizationFull) GetArchivedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.archived_at
}
// GetAvatarUrl gets the avatar_url property value. The avatar_url property
// returns a *string when successful
func (m *OrganizationFull) GetAvatarUrl()(*string) {
    return m.avatar_url
}
// GetBillingEmail gets the billing_email property value. The billing_email property
// returns a *string when successful
func (m *OrganizationFull) GetBillingEmail()(*string) {
    return m.billing_email
}
// GetBlog gets the blog property value. The blog property
// returns a *string when successful
func (m *OrganizationFull) GetBlog()(*string) {
    return m.blog
}
// GetCollaborators gets the collaborators property value. The collaborators property
// returns a *int32 when successful
func (m *OrganizationFull) GetCollaborators()(*int32) {
    return m.collaborators
}
// GetCompany gets the company property value. The company property
// returns a *string when successful
func (m *OrganizationFull) GetCompany()(*string) {
    return m.company
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *OrganizationFull) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDefaultRepositoryPermission gets the default_repository_permission property value. The default_repository_permission property
// returns a *string when successful
func (m *OrganizationFull) GetDefaultRepositoryPermission()(*string) {
    return m.default_repository_permission
}
// GetDependabotAlertsEnabledForNewRepositories gets the dependabot_alerts_enabled_for_new_repositories property value. Whether GitHub Advanced Security is automatically enabled for new repositories and repositories transferred tothis organization.This field is only visible to organization owners or members of a team with the security manager role.
// returns a *bool when successful
func (m *OrganizationFull) GetDependabotAlertsEnabledForNewRepositories()(*bool) {
    return m.dependabot_alerts_enabled_for_new_repositories
}
// GetDependabotSecurityUpdatesEnabledForNewRepositories gets the dependabot_security_updates_enabled_for_new_repositories property value. Whether dependabot security updates are automatically enabled for new repositories and repositories transferredto this organization.This field is only visible to organization owners or members of a team with the security manager role.
// returns a *bool when successful
func (m *OrganizationFull) GetDependabotSecurityUpdatesEnabledForNewRepositories()(*bool) {
    return m.dependabot_security_updates_enabled_for_new_repositories
}
// GetDependencyGraphEnabledForNewRepositories gets the dependency_graph_enabled_for_new_repositories property value. Whether dependency graph is automatically enabled for new repositories and repositories transferred to thisorganization.This field is only visible to organization owners or members of a team with the security manager role.
// returns a *bool when successful
func (m *OrganizationFull) GetDependencyGraphEnabledForNewRepositories()(*bool) {
    return m.dependency_graph_enabled_for_new_repositories
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *OrganizationFull) GetDescription()(*string) {
    return m.description
}
// GetDiskUsage gets the disk_usage property value. The disk_usage property
// returns a *int32 when successful
func (m *OrganizationFull) GetDiskUsage()(*int32) {
    return m.disk_usage
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *OrganizationFull) GetEmail()(*string) {
    return m.email
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *OrganizationFull) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationFull) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advanced_security_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurityEnabledForNewRepositories(val)
        }
        return nil
    }
    res["archived_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchivedAt(val)
        }
        return nil
    }
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
    res["billing_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingEmail(val)
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
    res["default_repository_permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRepositoryPermission(val)
        }
        return nil
    }
    res["dependabot_alerts_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotAlertsEnabledForNewRepositories(val)
        }
        return nil
    }
    res["dependabot_security_updates_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotSecurityUpdatesEnabledForNewRepositories(val)
        }
        return nil
    }
    res["dependency_graph_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencyGraphEnabledForNewRepositories(val)
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
    res["members_allowed_repository_creation_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersAllowedRepositoryCreationType(val)
        }
        return nil
    }
    res["members_can_create_internal_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCanCreateInternalRepositories(val)
        }
        return nil
    }
    res["members_can_create_pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCanCreatePages(val)
        }
        return nil
    }
    res["members_can_create_private_pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCanCreatePrivatePages(val)
        }
        return nil
    }
    res["members_can_create_private_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCanCreatePrivateRepositories(val)
        }
        return nil
    }
    res["members_can_create_public_pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCanCreatePublicPages(val)
        }
        return nil
    }
    res["members_can_create_public_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCanCreatePublicRepositories(val)
        }
        return nil
    }
    res["members_can_create_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCanCreateRepositories(val)
        }
        return nil
    }
    res["members_can_fork_private_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCanForkPrivateRepositories(val)
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
        val, err := n.GetObjectValue(CreateOrganizationFull_planFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlan(val.(OrganizationFull_planable))
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
    res["secret_scanning_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningEnabledForNewRepositories(val)
        }
        return nil
    }
    res["secret_scanning_push_protection_custom_link"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtectionCustomLink(val)
        }
        return nil
    }
    res["secret_scanning_push_protection_custom_link_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtectionCustomLinkEnabled(val)
        }
        return nil
    }
    res["secret_scanning_push_protection_enabled_for_new_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtectionEnabledForNewRepositories(val)
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
    res["two_factor_requirement_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwoFactorRequirementEnabled(val)
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
    res["web_commit_signoff_required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebCommitSignoffRequired(val)
        }
        return nil
    }
    return res
}
// GetFollowers gets the followers property value. The followers property
// returns a *int32 when successful
func (m *OrganizationFull) GetFollowers()(*int32) {
    return m.followers
}
// GetFollowing gets the following property value. The following property
// returns a *int32 when successful
func (m *OrganizationFull) GetFollowing()(*int32) {
    return m.following
}
// GetHasOrganizationProjects gets the has_organization_projects property value. The has_organization_projects property
// returns a *bool when successful
func (m *OrganizationFull) GetHasOrganizationProjects()(*bool) {
    return m.has_organization_projects
}
// GetHasRepositoryProjects gets the has_repository_projects property value. The has_repository_projects property
// returns a *bool when successful
func (m *OrganizationFull) GetHasRepositoryProjects()(*bool) {
    return m.has_repository_projects
}
// GetHooksUrl gets the hooks_url property value. The hooks_url property
// returns a *string when successful
func (m *OrganizationFull) GetHooksUrl()(*string) {
    return m.hooks_url
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *OrganizationFull) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *OrganizationFull) GetId()(*int32) {
    return m.id
}
// GetIssuesUrl gets the issues_url property value. The issues_url property
// returns a *string when successful
func (m *OrganizationFull) GetIssuesUrl()(*string) {
    return m.issues_url
}
// GetIsVerified gets the is_verified property value. The is_verified property
// returns a *bool when successful
func (m *OrganizationFull) GetIsVerified()(*bool) {
    return m.is_verified
}
// GetLocation gets the location property value. The location property
// returns a *string when successful
func (m *OrganizationFull) GetLocation()(*string) {
    return m.location
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *OrganizationFull) GetLogin()(*string) {
    return m.login
}
// GetMembersAllowedRepositoryCreationType gets the members_allowed_repository_creation_type property value. The members_allowed_repository_creation_type property
// returns a *string when successful
func (m *OrganizationFull) GetMembersAllowedRepositoryCreationType()(*string) {
    return m.members_allowed_repository_creation_type
}
// GetMembersCanCreateInternalRepositories gets the members_can_create_internal_repositories property value. The members_can_create_internal_repositories property
// returns a *bool when successful
func (m *OrganizationFull) GetMembersCanCreateInternalRepositories()(*bool) {
    return m.members_can_create_internal_repositories
}
// GetMembersCanCreatePages gets the members_can_create_pages property value. The members_can_create_pages property
// returns a *bool when successful
func (m *OrganizationFull) GetMembersCanCreatePages()(*bool) {
    return m.members_can_create_pages
}
// GetMembersCanCreatePrivatePages gets the members_can_create_private_pages property value. The members_can_create_private_pages property
// returns a *bool when successful
func (m *OrganizationFull) GetMembersCanCreatePrivatePages()(*bool) {
    return m.members_can_create_private_pages
}
// GetMembersCanCreatePrivateRepositories gets the members_can_create_private_repositories property value. The members_can_create_private_repositories property
// returns a *bool when successful
func (m *OrganizationFull) GetMembersCanCreatePrivateRepositories()(*bool) {
    return m.members_can_create_private_repositories
}
// GetMembersCanCreatePublicPages gets the members_can_create_public_pages property value. The members_can_create_public_pages property
// returns a *bool when successful
func (m *OrganizationFull) GetMembersCanCreatePublicPages()(*bool) {
    return m.members_can_create_public_pages
}
// GetMembersCanCreatePublicRepositories gets the members_can_create_public_repositories property value. The members_can_create_public_repositories property
// returns a *bool when successful
func (m *OrganizationFull) GetMembersCanCreatePublicRepositories()(*bool) {
    return m.members_can_create_public_repositories
}
// GetMembersCanCreateRepositories gets the members_can_create_repositories property value. The members_can_create_repositories property
// returns a *bool when successful
func (m *OrganizationFull) GetMembersCanCreateRepositories()(*bool) {
    return m.members_can_create_repositories
}
// GetMembersCanForkPrivateRepositories gets the members_can_fork_private_repositories property value. The members_can_fork_private_repositories property
// returns a *bool when successful
func (m *OrganizationFull) GetMembersCanForkPrivateRepositories()(*bool) {
    return m.members_can_fork_private_repositories
}
// GetMembersUrl gets the members_url property value. The members_url property
// returns a *string when successful
func (m *OrganizationFull) GetMembersUrl()(*string) {
    return m.members_url
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *OrganizationFull) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *OrganizationFull) GetNodeId()(*string) {
    return m.node_id
}
// GetOwnedPrivateRepos gets the owned_private_repos property value. The owned_private_repos property
// returns a *int32 when successful
func (m *OrganizationFull) GetOwnedPrivateRepos()(*int32) {
    return m.owned_private_repos
}
// GetPlan gets the plan property value. The plan property
// returns a OrganizationFull_planable when successful
func (m *OrganizationFull) GetPlan()(OrganizationFull_planable) {
    return m.plan
}
// GetPrivateGists gets the private_gists property value. The private_gists property
// returns a *int32 when successful
func (m *OrganizationFull) GetPrivateGists()(*int32) {
    return m.private_gists
}
// GetPublicGists gets the public_gists property value. The public_gists property
// returns a *int32 when successful
func (m *OrganizationFull) GetPublicGists()(*int32) {
    return m.public_gists
}
// GetPublicMembersUrl gets the public_members_url property value. The public_members_url property
// returns a *string when successful
func (m *OrganizationFull) GetPublicMembersUrl()(*string) {
    return m.public_members_url
}
// GetPublicRepos gets the public_repos property value. The public_repos property
// returns a *int32 when successful
func (m *OrganizationFull) GetPublicRepos()(*int32) {
    return m.public_repos
}
// GetReposUrl gets the repos_url property value. The repos_url property
// returns a *string when successful
func (m *OrganizationFull) GetReposUrl()(*string) {
    return m.repos_url
}
// GetSecretScanningEnabledForNewRepositories gets the secret_scanning_enabled_for_new_repositories property value. Whether secret scanning is automatically enabled for new repositories and repositories transferred to thisorganization.This field is only visible to organization owners or members of a team with the security manager role.
// returns a *bool when successful
func (m *OrganizationFull) GetSecretScanningEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_enabled_for_new_repositories
}
// GetSecretScanningPushProtectionCustomLink gets the secret_scanning_push_protection_custom_link property value. An optional URL string to display to contributors who are blocked from pushing a secret.
// returns a *string when successful
func (m *OrganizationFull) GetSecretScanningPushProtectionCustomLink()(*string) {
    return m.secret_scanning_push_protection_custom_link
}
// GetSecretScanningPushProtectionCustomLinkEnabled gets the secret_scanning_push_protection_custom_link_enabled property value. Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
// returns a *bool when successful
func (m *OrganizationFull) GetSecretScanningPushProtectionCustomLinkEnabled()(*bool) {
    return m.secret_scanning_push_protection_custom_link_enabled
}
// GetSecretScanningPushProtectionEnabledForNewRepositories gets the secret_scanning_push_protection_enabled_for_new_repositories property value. Whether secret scanning push protection is automatically enabled for new repositories and repositoriestransferred to this organization.This field is only visible to organization owners or members of a team with the security manager role.
// returns a *bool when successful
func (m *OrganizationFull) GetSecretScanningPushProtectionEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_push_protection_enabled_for_new_repositories
}
// GetTotalPrivateRepos gets the total_private_repos property value. The total_private_repos property
// returns a *int32 when successful
func (m *OrganizationFull) GetTotalPrivateRepos()(*int32) {
    return m.total_private_repos
}
// GetTwitterUsername gets the twitter_username property value. The twitter_username property
// returns a *string when successful
func (m *OrganizationFull) GetTwitterUsername()(*string) {
    return m.twitter_username
}
// GetTwoFactorRequirementEnabled gets the two_factor_requirement_enabled property value. The two_factor_requirement_enabled property
// returns a *bool when successful
func (m *OrganizationFull) GetTwoFactorRequirementEnabled()(*bool) {
    return m.two_factor_requirement_enabled
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *OrganizationFull) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *OrganizationFull) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *OrganizationFull) GetUrl()(*string) {
    return m.url
}
// GetWebCommitSignoffRequired gets the web_commit_signoff_required property value. The web_commit_signoff_required property
// returns a *bool when successful
func (m *OrganizationFull) GetWebCommitSignoffRequired()(*bool) {
    return m.web_commit_signoff_required
}
// Serialize serializes information the current object
func (m *OrganizationFull) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("advanced_security_enabled_for_new_repositories", m.GetAdvancedSecurityEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("archived_at", m.GetArchivedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("avatar_url", m.GetAvatarUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("billing_email", m.GetBillingEmail())
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
        err := writer.WriteStringValue("default_repository_permission", m.GetDefaultRepositoryPermission())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dependabot_alerts_enabled_for_new_repositories", m.GetDependabotAlertsEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dependabot_security_updates_enabled_for_new_repositories", m.GetDependabotSecurityUpdatesEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dependency_graph_enabled_for_new_repositories", m.GetDependencyGraphEnabledForNewRepositories())
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
        err := writer.WriteStringValue("members_allowed_repository_creation_type", m.GetMembersAllowedRepositoryCreationType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("members_can_create_internal_repositories", m.GetMembersCanCreateInternalRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("members_can_create_pages", m.GetMembersCanCreatePages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("members_can_create_private_pages", m.GetMembersCanCreatePrivatePages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("members_can_create_private_repositories", m.GetMembersCanCreatePrivateRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("members_can_create_public_pages", m.GetMembersCanCreatePublicPages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("members_can_create_public_repositories", m.GetMembersCanCreatePublicRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("members_can_create_repositories", m.GetMembersCanCreateRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("members_can_fork_private_repositories", m.GetMembersCanForkPrivateRepositories())
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
        err := writer.WriteBoolValue("secret_scanning_enabled_for_new_repositories", m.GetSecretScanningEnabledForNewRepositories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret_scanning_push_protection_custom_link", m.GetSecretScanningPushProtectionCustomLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("secret_scanning_push_protection_custom_link_enabled", m.GetSecretScanningPushProtectionCustomLinkEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("secret_scanning_push_protection_enabled_for_new_repositories", m.GetSecretScanningPushProtectionEnabledForNewRepositories())
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
        err := writer.WriteBoolValue("two_factor_requirement_enabled", m.GetTwoFactorRequirementEnabled())
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
        err := writer.WriteBoolValue("web_commit_signoff_required", m.GetWebCommitSignoffRequired())
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
func (m *OrganizationFull) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurityEnabledForNewRepositories sets the advanced_security_enabled_for_new_repositories property value. Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization.This field is only visible to organization owners or members of a team with the security manager role.
func (m *OrganizationFull) SetAdvancedSecurityEnabledForNewRepositories(value *bool)() {
    m.advanced_security_enabled_for_new_repositories = value
}
// SetArchivedAt sets the archived_at property value. The archived_at property
func (m *OrganizationFull) SetArchivedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.archived_at = value
}
// SetAvatarUrl sets the avatar_url property value. The avatar_url property
func (m *OrganizationFull) SetAvatarUrl(value *string)() {
    m.avatar_url = value
}
// SetBillingEmail sets the billing_email property value. The billing_email property
func (m *OrganizationFull) SetBillingEmail(value *string)() {
    m.billing_email = value
}
// SetBlog sets the blog property value. The blog property
func (m *OrganizationFull) SetBlog(value *string)() {
    m.blog = value
}
// SetCollaborators sets the collaborators property value. The collaborators property
func (m *OrganizationFull) SetCollaborators(value *int32)() {
    m.collaborators = value
}
// SetCompany sets the company property value. The company property
func (m *OrganizationFull) SetCompany(value *string)() {
    m.company = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *OrganizationFull) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDefaultRepositoryPermission sets the default_repository_permission property value. The default_repository_permission property
func (m *OrganizationFull) SetDefaultRepositoryPermission(value *string)() {
    m.default_repository_permission = value
}
// SetDependabotAlertsEnabledForNewRepositories sets the dependabot_alerts_enabled_for_new_repositories property value. Whether GitHub Advanced Security is automatically enabled for new repositories and repositories transferred tothis organization.This field is only visible to organization owners or members of a team with the security manager role.
func (m *OrganizationFull) SetDependabotAlertsEnabledForNewRepositories(value *bool)() {
    m.dependabot_alerts_enabled_for_new_repositories = value
}
// SetDependabotSecurityUpdatesEnabledForNewRepositories sets the dependabot_security_updates_enabled_for_new_repositories property value. Whether dependabot security updates are automatically enabled for new repositories and repositories transferredto this organization.This field is only visible to organization owners or members of a team with the security manager role.
func (m *OrganizationFull) SetDependabotSecurityUpdatesEnabledForNewRepositories(value *bool)() {
    m.dependabot_security_updates_enabled_for_new_repositories = value
}
// SetDependencyGraphEnabledForNewRepositories sets the dependency_graph_enabled_for_new_repositories property value. Whether dependency graph is automatically enabled for new repositories and repositories transferred to thisorganization.This field is only visible to organization owners or members of a team with the security manager role.
func (m *OrganizationFull) SetDependencyGraphEnabledForNewRepositories(value *bool)() {
    m.dependency_graph_enabled_for_new_repositories = value
}
// SetDescription sets the description property value. The description property
func (m *OrganizationFull) SetDescription(value *string)() {
    m.description = value
}
// SetDiskUsage sets the disk_usage property value. The disk_usage property
func (m *OrganizationFull) SetDiskUsage(value *int32)() {
    m.disk_usage = value
}
// SetEmail sets the email property value. The email property
func (m *OrganizationFull) SetEmail(value *string)() {
    m.email = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *OrganizationFull) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetFollowers sets the followers property value. The followers property
func (m *OrganizationFull) SetFollowers(value *int32)() {
    m.followers = value
}
// SetFollowing sets the following property value. The following property
func (m *OrganizationFull) SetFollowing(value *int32)() {
    m.following = value
}
// SetHasOrganizationProjects sets the has_organization_projects property value. The has_organization_projects property
func (m *OrganizationFull) SetHasOrganizationProjects(value *bool)() {
    m.has_organization_projects = value
}
// SetHasRepositoryProjects sets the has_repository_projects property value. The has_repository_projects property
func (m *OrganizationFull) SetHasRepositoryProjects(value *bool)() {
    m.has_repository_projects = value
}
// SetHooksUrl sets the hooks_url property value. The hooks_url property
func (m *OrganizationFull) SetHooksUrl(value *string)() {
    m.hooks_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *OrganizationFull) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *OrganizationFull) SetId(value *int32)() {
    m.id = value
}
// SetIssuesUrl sets the issues_url property value. The issues_url property
func (m *OrganizationFull) SetIssuesUrl(value *string)() {
    m.issues_url = value
}
// SetIsVerified sets the is_verified property value. The is_verified property
func (m *OrganizationFull) SetIsVerified(value *bool)() {
    m.is_verified = value
}
// SetLocation sets the location property value. The location property
func (m *OrganizationFull) SetLocation(value *string)() {
    m.location = value
}
// SetLogin sets the login property value. The login property
func (m *OrganizationFull) SetLogin(value *string)() {
    m.login = value
}
// SetMembersAllowedRepositoryCreationType sets the members_allowed_repository_creation_type property value. The members_allowed_repository_creation_type property
func (m *OrganizationFull) SetMembersAllowedRepositoryCreationType(value *string)() {
    m.members_allowed_repository_creation_type = value
}
// SetMembersCanCreateInternalRepositories sets the members_can_create_internal_repositories property value. The members_can_create_internal_repositories property
func (m *OrganizationFull) SetMembersCanCreateInternalRepositories(value *bool)() {
    m.members_can_create_internal_repositories = value
}
// SetMembersCanCreatePages sets the members_can_create_pages property value. The members_can_create_pages property
func (m *OrganizationFull) SetMembersCanCreatePages(value *bool)() {
    m.members_can_create_pages = value
}
// SetMembersCanCreatePrivatePages sets the members_can_create_private_pages property value. The members_can_create_private_pages property
func (m *OrganizationFull) SetMembersCanCreatePrivatePages(value *bool)() {
    m.members_can_create_private_pages = value
}
// SetMembersCanCreatePrivateRepositories sets the members_can_create_private_repositories property value. The members_can_create_private_repositories property
func (m *OrganizationFull) SetMembersCanCreatePrivateRepositories(value *bool)() {
    m.members_can_create_private_repositories = value
}
// SetMembersCanCreatePublicPages sets the members_can_create_public_pages property value. The members_can_create_public_pages property
func (m *OrganizationFull) SetMembersCanCreatePublicPages(value *bool)() {
    m.members_can_create_public_pages = value
}
// SetMembersCanCreatePublicRepositories sets the members_can_create_public_repositories property value. The members_can_create_public_repositories property
func (m *OrganizationFull) SetMembersCanCreatePublicRepositories(value *bool)() {
    m.members_can_create_public_repositories = value
}
// SetMembersCanCreateRepositories sets the members_can_create_repositories property value. The members_can_create_repositories property
func (m *OrganizationFull) SetMembersCanCreateRepositories(value *bool)() {
    m.members_can_create_repositories = value
}
// SetMembersCanForkPrivateRepositories sets the members_can_fork_private_repositories property value. The members_can_fork_private_repositories property
func (m *OrganizationFull) SetMembersCanForkPrivateRepositories(value *bool)() {
    m.members_can_fork_private_repositories = value
}
// SetMembersUrl sets the members_url property value. The members_url property
func (m *OrganizationFull) SetMembersUrl(value *string)() {
    m.members_url = value
}
// SetName sets the name property value. The name property
func (m *OrganizationFull) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *OrganizationFull) SetNodeId(value *string)() {
    m.node_id = value
}
// SetOwnedPrivateRepos sets the owned_private_repos property value. The owned_private_repos property
func (m *OrganizationFull) SetOwnedPrivateRepos(value *int32)() {
    m.owned_private_repos = value
}
// SetPlan sets the plan property value. The plan property
func (m *OrganizationFull) SetPlan(value OrganizationFull_planable)() {
    m.plan = value
}
// SetPrivateGists sets the private_gists property value. The private_gists property
func (m *OrganizationFull) SetPrivateGists(value *int32)() {
    m.private_gists = value
}
// SetPublicGists sets the public_gists property value. The public_gists property
func (m *OrganizationFull) SetPublicGists(value *int32)() {
    m.public_gists = value
}
// SetPublicMembersUrl sets the public_members_url property value. The public_members_url property
func (m *OrganizationFull) SetPublicMembersUrl(value *string)() {
    m.public_members_url = value
}
// SetPublicRepos sets the public_repos property value. The public_repos property
func (m *OrganizationFull) SetPublicRepos(value *int32)() {
    m.public_repos = value
}
// SetReposUrl sets the repos_url property value. The repos_url property
func (m *OrganizationFull) SetReposUrl(value *string)() {
    m.repos_url = value
}
// SetSecretScanningEnabledForNewRepositories sets the secret_scanning_enabled_for_new_repositories property value. Whether secret scanning is automatically enabled for new repositories and repositories transferred to thisorganization.This field is only visible to organization owners or members of a team with the security manager role.
func (m *OrganizationFull) SetSecretScanningEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_enabled_for_new_repositories = value
}
// SetSecretScanningPushProtectionCustomLink sets the secret_scanning_push_protection_custom_link property value. An optional URL string to display to contributors who are blocked from pushing a secret.
func (m *OrganizationFull) SetSecretScanningPushProtectionCustomLink(value *string)() {
    m.secret_scanning_push_protection_custom_link = value
}
// SetSecretScanningPushProtectionCustomLinkEnabled sets the secret_scanning_push_protection_custom_link_enabled property value. Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
func (m *OrganizationFull) SetSecretScanningPushProtectionCustomLinkEnabled(value *bool)() {
    m.secret_scanning_push_protection_custom_link_enabled = value
}
// SetSecretScanningPushProtectionEnabledForNewRepositories sets the secret_scanning_push_protection_enabled_for_new_repositories property value. Whether secret scanning push protection is automatically enabled for new repositories and repositoriestransferred to this organization.This field is only visible to organization owners or members of a team with the security manager role.
func (m *OrganizationFull) SetSecretScanningPushProtectionEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_push_protection_enabled_for_new_repositories = value
}
// SetTotalPrivateRepos sets the total_private_repos property value. The total_private_repos property
func (m *OrganizationFull) SetTotalPrivateRepos(value *int32)() {
    m.total_private_repos = value
}
// SetTwitterUsername sets the twitter_username property value. The twitter_username property
func (m *OrganizationFull) SetTwitterUsername(value *string)() {
    m.twitter_username = value
}
// SetTwoFactorRequirementEnabled sets the two_factor_requirement_enabled property value. The two_factor_requirement_enabled property
func (m *OrganizationFull) SetTwoFactorRequirementEnabled(value *bool)() {
    m.two_factor_requirement_enabled = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *OrganizationFull) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *OrganizationFull) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *OrganizationFull) SetUrl(value *string)() {
    m.url = value
}
// SetWebCommitSignoffRequired sets the web_commit_signoff_required property value. The web_commit_signoff_required property
func (m *OrganizationFull) SetWebCommitSignoffRequired(value *bool)() {
    m.web_commit_signoff_required = value
}
type OrganizationFullable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurityEnabledForNewRepositories()(*bool)
    GetArchivedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAvatarUrl()(*string)
    GetBillingEmail()(*string)
    GetBlog()(*string)
    GetCollaborators()(*int32)
    GetCompany()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDefaultRepositoryPermission()(*string)
    GetDependabotAlertsEnabledForNewRepositories()(*bool)
    GetDependabotSecurityUpdatesEnabledForNewRepositories()(*bool)
    GetDependencyGraphEnabledForNewRepositories()(*bool)
    GetDescription()(*string)
    GetDiskUsage()(*int32)
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
    GetMembersAllowedRepositoryCreationType()(*string)
    GetMembersCanCreateInternalRepositories()(*bool)
    GetMembersCanCreatePages()(*bool)
    GetMembersCanCreatePrivatePages()(*bool)
    GetMembersCanCreatePrivateRepositories()(*bool)
    GetMembersCanCreatePublicPages()(*bool)
    GetMembersCanCreatePublicRepositories()(*bool)
    GetMembersCanCreateRepositories()(*bool)
    GetMembersCanForkPrivateRepositories()(*bool)
    GetMembersUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetOwnedPrivateRepos()(*int32)
    GetPlan()(OrganizationFull_planable)
    GetPrivateGists()(*int32)
    GetPublicGists()(*int32)
    GetPublicMembersUrl()(*string)
    GetPublicRepos()(*int32)
    GetReposUrl()(*string)
    GetSecretScanningEnabledForNewRepositories()(*bool)
    GetSecretScanningPushProtectionCustomLink()(*string)
    GetSecretScanningPushProtectionCustomLinkEnabled()(*bool)
    GetSecretScanningPushProtectionEnabledForNewRepositories()(*bool)
    GetTotalPrivateRepos()(*int32)
    GetTwitterUsername()(*string)
    GetTwoFactorRequirementEnabled()(*bool)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetWebCommitSignoffRequired()(*bool)
    SetAdvancedSecurityEnabledForNewRepositories(value *bool)()
    SetArchivedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAvatarUrl(value *string)()
    SetBillingEmail(value *string)()
    SetBlog(value *string)()
    SetCollaborators(value *int32)()
    SetCompany(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDefaultRepositoryPermission(value *string)()
    SetDependabotAlertsEnabledForNewRepositories(value *bool)()
    SetDependabotSecurityUpdatesEnabledForNewRepositories(value *bool)()
    SetDependencyGraphEnabledForNewRepositories(value *bool)()
    SetDescription(value *string)()
    SetDiskUsage(value *int32)()
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
    SetMembersAllowedRepositoryCreationType(value *string)()
    SetMembersCanCreateInternalRepositories(value *bool)()
    SetMembersCanCreatePages(value *bool)()
    SetMembersCanCreatePrivatePages(value *bool)()
    SetMembersCanCreatePrivateRepositories(value *bool)()
    SetMembersCanCreatePublicPages(value *bool)()
    SetMembersCanCreatePublicRepositories(value *bool)()
    SetMembersCanCreateRepositories(value *bool)()
    SetMembersCanForkPrivateRepositories(value *bool)()
    SetMembersUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOwnedPrivateRepos(value *int32)()
    SetPlan(value OrganizationFull_planable)()
    SetPrivateGists(value *int32)()
    SetPublicGists(value *int32)()
    SetPublicMembersUrl(value *string)()
    SetPublicRepos(value *int32)()
    SetReposUrl(value *string)()
    SetSecretScanningEnabledForNewRepositories(value *bool)()
    SetSecretScanningPushProtectionCustomLink(value *string)()
    SetSecretScanningPushProtectionCustomLinkEnabled(value *bool)()
    SetSecretScanningPushProtectionEnabledForNewRepositories(value *bool)()
    SetTotalPrivateRepos(value *int32)()
    SetTwitterUsername(value *string)()
    SetTwoFactorRequirementEnabled(value *bool)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetWebCommitSignoffRequired(value *bool)()
}
