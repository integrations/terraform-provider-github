package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemWithOrgPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether GitHub Advanced Security is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
    advanced_security_enabled_for_new_repositories *bool
    // Billing email address. This address is not publicized.
    billing_email *string
    // The blog property
    blog *string
    // The company name.
    company *string
    // Whether Dependabot alerts is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
    dependabot_alerts_enabled_for_new_repositories *bool
    // Whether Dependabot security updates is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
    dependabot_security_updates_enabled_for_new_repositories *bool
    // Whether dependency graph is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
    dependency_graph_enabled_for_new_repositories *bool
    // The description of the company. The maximum size is 160 characters.
    description *string
    // The publicly visible email address.
    email *string
    // Whether an organization can use organization projects.
    has_organization_projects *bool
    // Whether repositories that belong to the organization can use repository projects.
    has_repository_projects *bool
    // The location.
    location *string
    // Whether organization members can create internal repositories, which are visible to all enterprise members. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see "[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation.
    members_can_create_internal_repositories *bool
    // Whether organization members can create GitHub Pages sites. Existing published sites will not be impacted.
    members_can_create_pages *bool
    // Whether organization members can create private GitHub Pages sites. Existing published sites will not be impacted.
    members_can_create_private_pages *bool
    // Whether organization members can create private repositories, which are visible to organization members with permission. For more information, see "[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation.
    members_can_create_private_repositories *bool
    // Whether organization members can create public GitHub Pages sites. Existing published sites will not be impacted.
    members_can_create_public_pages *bool
    // Whether organization members can create public repositories, which are visible to anyone. For more information, see "[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation.
    members_can_create_public_repositories *bool
    // Whether of non-admin organization members can create repositories. **Note:** A parameter can override this parameter. See `members_allowed_repository_creation_type` in this table for details.
    members_can_create_repositories *bool
    // Whether organization members can fork private organization repositories.
    members_can_fork_private_repositories *bool
    // The shorthand name of the company.
    name *string
    // Whether secret scanning is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
    secret_scanning_enabled_for_new_repositories *bool
    // If `secret_scanning_push_protection_custom_link_enabled` is true, the URL that will be displayed to contributors who are blocked from pushing a secret.
    secret_scanning_push_protection_custom_link *string
    // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
    secret_scanning_push_protection_custom_link_enabled *bool
    // Whether secret scanning push protection is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
    secret_scanning_push_protection_enabled_for_new_repositories *bool
    // The Twitter username of the company.
    twitter_username *string
    // Whether contributors to organization repositories are required to sign off on commits they make through GitHub's web interface.
    web_commit_signoff_required *bool
}
// NewItemWithOrgPatchRequestBody instantiates a new ItemWithOrgPatchRequestBody and sets the default values.
func NewItemWithOrgPatchRequestBody()(*ItemWithOrgPatchRequestBody) {
    m := &ItemWithOrgPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWithOrgPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemWithOrgPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWithOrgPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemWithOrgPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurityEnabledForNewRepositories gets the advanced_security_enabled_for_new_repositories property value. Whether GitHub Advanced Security is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetAdvancedSecurityEnabledForNewRepositories()(*bool) {
    return m.advanced_security_enabled_for_new_repositories
}
// GetBillingEmail gets the billing_email property value. Billing email address. This address is not publicized.
// returns a *string when successful
func (m *ItemWithOrgPatchRequestBody) GetBillingEmail()(*string) {
    return m.billing_email
}
// GetBlog gets the blog property value. The blog property
// returns a *string when successful
func (m *ItemWithOrgPatchRequestBody) GetBlog()(*string) {
    return m.blog
}
// GetCompany gets the company property value. The company name.
// returns a *string when successful
func (m *ItemWithOrgPatchRequestBody) GetCompany()(*string) {
    return m.company
}
// GetDependabotAlertsEnabledForNewRepositories gets the dependabot_alerts_enabled_for_new_repositories property value. Whether Dependabot alerts is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetDependabotAlertsEnabledForNewRepositories()(*bool) {
    return m.dependabot_alerts_enabled_for_new_repositories
}
// GetDependabotSecurityUpdatesEnabledForNewRepositories gets the dependabot_security_updates_enabled_for_new_repositories property value. Whether Dependabot security updates is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetDependabotSecurityUpdatesEnabledForNewRepositories()(*bool) {
    return m.dependabot_security_updates_enabled_for_new_repositories
}
// GetDependencyGraphEnabledForNewRepositories gets the dependency_graph_enabled_for_new_repositories property value. Whether dependency graph is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetDependencyGraphEnabledForNewRepositories()(*bool) {
    return m.dependency_graph_enabled_for_new_repositories
}
// GetDescription gets the description property value. The description of the company. The maximum size is 160 characters.
// returns a *string when successful
func (m *ItemWithOrgPatchRequestBody) GetDescription()(*string) {
    return m.description
}
// GetEmail gets the email property value. The publicly visible email address.
// returns a *string when successful
func (m *ItemWithOrgPatchRequestBody) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemWithOrgPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetHasOrganizationProjects gets the has_organization_projects property value. Whether an organization can use organization projects.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetHasOrganizationProjects()(*bool) {
    return m.has_organization_projects
}
// GetHasRepositoryProjects gets the has_repository_projects property value. Whether repositories that belong to the organization can use repository projects.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetHasRepositoryProjects()(*bool) {
    return m.has_repository_projects
}
// GetLocation gets the location property value. The location.
// returns a *string when successful
func (m *ItemWithOrgPatchRequestBody) GetLocation()(*string) {
    return m.location
}
// GetMembersCanCreateInternalRepositories gets the members_can_create_internal_repositories property value. Whether organization members can create internal repositories, which are visible to all enterprise members. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see "[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetMembersCanCreateInternalRepositories()(*bool) {
    return m.members_can_create_internal_repositories
}
// GetMembersCanCreatePages gets the members_can_create_pages property value. Whether organization members can create GitHub Pages sites. Existing published sites will not be impacted.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetMembersCanCreatePages()(*bool) {
    return m.members_can_create_pages
}
// GetMembersCanCreatePrivatePages gets the members_can_create_private_pages property value. Whether organization members can create private GitHub Pages sites. Existing published sites will not be impacted.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetMembersCanCreatePrivatePages()(*bool) {
    return m.members_can_create_private_pages
}
// GetMembersCanCreatePrivateRepositories gets the members_can_create_private_repositories property value. Whether organization members can create private repositories, which are visible to organization members with permission. For more information, see "[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetMembersCanCreatePrivateRepositories()(*bool) {
    return m.members_can_create_private_repositories
}
// GetMembersCanCreatePublicPages gets the members_can_create_public_pages property value. Whether organization members can create public GitHub Pages sites. Existing published sites will not be impacted.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetMembersCanCreatePublicPages()(*bool) {
    return m.members_can_create_public_pages
}
// GetMembersCanCreatePublicRepositories gets the members_can_create_public_repositories property value. Whether organization members can create public repositories, which are visible to anyone. For more information, see "[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetMembersCanCreatePublicRepositories()(*bool) {
    return m.members_can_create_public_repositories
}
// GetMembersCanCreateRepositories gets the members_can_create_repositories property value. Whether of non-admin organization members can create repositories. **Note:** A parameter can override this parameter. See `members_allowed_repository_creation_type` in this table for details.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetMembersCanCreateRepositories()(*bool) {
    return m.members_can_create_repositories
}
// GetMembersCanForkPrivateRepositories gets the members_can_fork_private_repositories property value. Whether organization members can fork private organization repositories.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetMembersCanForkPrivateRepositories()(*bool) {
    return m.members_can_fork_private_repositories
}
// GetName gets the name property value. The shorthand name of the company.
// returns a *string when successful
func (m *ItemWithOrgPatchRequestBody) GetName()(*string) {
    return m.name
}
// GetSecretScanningEnabledForNewRepositories gets the secret_scanning_enabled_for_new_repositories property value. Whether secret scanning is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetSecretScanningEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_enabled_for_new_repositories
}
// GetSecretScanningPushProtectionCustomLink gets the secret_scanning_push_protection_custom_link property value. If `secret_scanning_push_protection_custom_link_enabled` is true, the URL that will be displayed to contributors who are blocked from pushing a secret.
// returns a *string when successful
func (m *ItemWithOrgPatchRequestBody) GetSecretScanningPushProtectionCustomLink()(*string) {
    return m.secret_scanning_push_protection_custom_link
}
// GetSecretScanningPushProtectionCustomLinkEnabled gets the secret_scanning_push_protection_custom_link_enabled property value. Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetSecretScanningPushProtectionCustomLinkEnabled()(*bool) {
    return m.secret_scanning_push_protection_custom_link_enabled
}
// GetSecretScanningPushProtectionEnabledForNewRepositories gets the secret_scanning_push_protection_enabled_for_new_repositories property value. Whether secret scanning push protection is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetSecretScanningPushProtectionEnabledForNewRepositories()(*bool) {
    return m.secret_scanning_push_protection_enabled_for_new_repositories
}
// GetTwitterUsername gets the twitter_username property value. The Twitter username of the company.
// returns a *string when successful
func (m *ItemWithOrgPatchRequestBody) GetTwitterUsername()(*string) {
    return m.twitter_username
}
// GetWebCommitSignoffRequired gets the web_commit_signoff_required property value. Whether contributors to organization repositories are required to sign off on commits they make through GitHub's web interface.
// returns a *bool when successful
func (m *ItemWithOrgPatchRequestBody) GetWebCommitSignoffRequired()(*bool) {
    return m.web_commit_signoff_required
}
// Serialize serializes information the current object
func (m *ItemWithOrgPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("advanced_security_enabled_for_new_repositories", m.GetAdvancedSecurityEnabledForNewRepositories())
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
        err := writer.WriteStringValue("company", m.GetCompany())
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
        err := writer.WriteStringValue("email", m.GetEmail())
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
        err := writer.WriteStringValue("location", m.GetLocation())
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
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("twitter_username", m.GetTwitterUsername())
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
func (m *ItemWithOrgPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurityEnabledForNewRepositories sets the advanced_security_enabled_for_new_repositories property value. Whether GitHub Advanced Security is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
func (m *ItemWithOrgPatchRequestBody) SetAdvancedSecurityEnabledForNewRepositories(value *bool)() {
    m.advanced_security_enabled_for_new_repositories = value
}
// SetBillingEmail sets the billing_email property value. Billing email address. This address is not publicized.
func (m *ItemWithOrgPatchRequestBody) SetBillingEmail(value *string)() {
    m.billing_email = value
}
// SetBlog sets the blog property value. The blog property
func (m *ItemWithOrgPatchRequestBody) SetBlog(value *string)() {
    m.blog = value
}
// SetCompany sets the company property value. The company name.
func (m *ItemWithOrgPatchRequestBody) SetCompany(value *string)() {
    m.company = value
}
// SetDependabotAlertsEnabledForNewRepositories sets the dependabot_alerts_enabled_for_new_repositories property value. Whether Dependabot alerts is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
func (m *ItemWithOrgPatchRequestBody) SetDependabotAlertsEnabledForNewRepositories(value *bool)() {
    m.dependabot_alerts_enabled_for_new_repositories = value
}
// SetDependabotSecurityUpdatesEnabledForNewRepositories sets the dependabot_security_updates_enabled_for_new_repositories property value. Whether Dependabot security updates is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
func (m *ItemWithOrgPatchRequestBody) SetDependabotSecurityUpdatesEnabledForNewRepositories(value *bool)() {
    m.dependabot_security_updates_enabled_for_new_repositories = value
}
// SetDependencyGraphEnabledForNewRepositories sets the dependency_graph_enabled_for_new_repositories property value. Whether dependency graph is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
func (m *ItemWithOrgPatchRequestBody) SetDependencyGraphEnabledForNewRepositories(value *bool)() {
    m.dependency_graph_enabled_for_new_repositories = value
}
// SetDescription sets the description property value. The description of the company. The maximum size is 160 characters.
func (m *ItemWithOrgPatchRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetEmail sets the email property value. The publicly visible email address.
func (m *ItemWithOrgPatchRequestBody) SetEmail(value *string)() {
    m.email = value
}
// SetHasOrganizationProjects sets the has_organization_projects property value. Whether an organization can use organization projects.
func (m *ItemWithOrgPatchRequestBody) SetHasOrganizationProjects(value *bool)() {
    m.has_organization_projects = value
}
// SetHasRepositoryProjects sets the has_repository_projects property value. Whether repositories that belong to the organization can use repository projects.
func (m *ItemWithOrgPatchRequestBody) SetHasRepositoryProjects(value *bool)() {
    m.has_repository_projects = value
}
// SetLocation sets the location property value. The location.
func (m *ItemWithOrgPatchRequestBody) SetLocation(value *string)() {
    m.location = value
}
// SetMembersCanCreateInternalRepositories sets the members_can_create_internal_repositories property value. Whether organization members can create internal repositories, which are visible to all enterprise members. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see "[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation.
func (m *ItemWithOrgPatchRequestBody) SetMembersCanCreateInternalRepositories(value *bool)() {
    m.members_can_create_internal_repositories = value
}
// SetMembersCanCreatePages sets the members_can_create_pages property value. Whether organization members can create GitHub Pages sites. Existing published sites will not be impacted.
func (m *ItemWithOrgPatchRequestBody) SetMembersCanCreatePages(value *bool)() {
    m.members_can_create_pages = value
}
// SetMembersCanCreatePrivatePages sets the members_can_create_private_pages property value. Whether organization members can create private GitHub Pages sites. Existing published sites will not be impacted.
func (m *ItemWithOrgPatchRequestBody) SetMembersCanCreatePrivatePages(value *bool)() {
    m.members_can_create_private_pages = value
}
// SetMembersCanCreatePrivateRepositories sets the members_can_create_private_repositories property value. Whether organization members can create private repositories, which are visible to organization members with permission. For more information, see "[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation.
func (m *ItemWithOrgPatchRequestBody) SetMembersCanCreatePrivateRepositories(value *bool)() {
    m.members_can_create_private_repositories = value
}
// SetMembersCanCreatePublicPages sets the members_can_create_public_pages property value. Whether organization members can create public GitHub Pages sites. Existing published sites will not be impacted.
func (m *ItemWithOrgPatchRequestBody) SetMembersCanCreatePublicPages(value *bool)() {
    m.members_can_create_public_pages = value
}
// SetMembersCanCreatePublicRepositories sets the members_can_create_public_repositories property value. Whether organization members can create public repositories, which are visible to anyone. For more information, see "[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)" in the GitHub Help documentation.
func (m *ItemWithOrgPatchRequestBody) SetMembersCanCreatePublicRepositories(value *bool)() {
    m.members_can_create_public_repositories = value
}
// SetMembersCanCreateRepositories sets the members_can_create_repositories property value. Whether of non-admin organization members can create repositories. **Note:** A parameter can override this parameter. See `members_allowed_repository_creation_type` in this table for details.
func (m *ItemWithOrgPatchRequestBody) SetMembersCanCreateRepositories(value *bool)() {
    m.members_can_create_repositories = value
}
// SetMembersCanForkPrivateRepositories sets the members_can_fork_private_repositories property value. Whether organization members can fork private organization repositories.
func (m *ItemWithOrgPatchRequestBody) SetMembersCanForkPrivateRepositories(value *bool)() {
    m.members_can_fork_private_repositories = value
}
// SetName sets the name property value. The shorthand name of the company.
func (m *ItemWithOrgPatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetSecretScanningEnabledForNewRepositories sets the secret_scanning_enabled_for_new_repositories property value. Whether secret scanning is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
func (m *ItemWithOrgPatchRequestBody) SetSecretScanningEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_enabled_for_new_repositories = value
}
// SetSecretScanningPushProtectionCustomLink sets the secret_scanning_push_protection_custom_link property value. If `secret_scanning_push_protection_custom_link_enabled` is true, the URL that will be displayed to contributors who are blocked from pushing a secret.
func (m *ItemWithOrgPatchRequestBody) SetSecretScanningPushProtectionCustomLink(value *string)() {
    m.secret_scanning_push_protection_custom_link = value
}
// SetSecretScanningPushProtectionCustomLinkEnabled sets the secret_scanning_push_protection_custom_link_enabled property value. Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
func (m *ItemWithOrgPatchRequestBody) SetSecretScanningPushProtectionCustomLinkEnabled(value *bool)() {
    m.secret_scanning_push_protection_custom_link_enabled = value
}
// SetSecretScanningPushProtectionEnabledForNewRepositories sets the secret_scanning_push_protection_enabled_for_new_repositories property value. Whether secret scanning push protection is automatically enabled for new repositories.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."You can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.
func (m *ItemWithOrgPatchRequestBody) SetSecretScanningPushProtectionEnabledForNewRepositories(value *bool)() {
    m.secret_scanning_push_protection_enabled_for_new_repositories = value
}
// SetTwitterUsername sets the twitter_username property value. The Twitter username of the company.
func (m *ItemWithOrgPatchRequestBody) SetTwitterUsername(value *string)() {
    m.twitter_username = value
}
// SetWebCommitSignoffRequired sets the web_commit_signoff_required property value. Whether contributors to organization repositories are required to sign off on commits they make through GitHub's web interface.
func (m *ItemWithOrgPatchRequestBody) SetWebCommitSignoffRequired(value *bool)() {
    m.web_commit_signoff_required = value
}
type ItemWithOrgPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurityEnabledForNewRepositories()(*bool)
    GetBillingEmail()(*string)
    GetBlog()(*string)
    GetCompany()(*string)
    GetDependabotAlertsEnabledForNewRepositories()(*bool)
    GetDependabotSecurityUpdatesEnabledForNewRepositories()(*bool)
    GetDependencyGraphEnabledForNewRepositories()(*bool)
    GetDescription()(*string)
    GetEmail()(*string)
    GetHasOrganizationProjects()(*bool)
    GetHasRepositoryProjects()(*bool)
    GetLocation()(*string)
    GetMembersCanCreateInternalRepositories()(*bool)
    GetMembersCanCreatePages()(*bool)
    GetMembersCanCreatePrivatePages()(*bool)
    GetMembersCanCreatePrivateRepositories()(*bool)
    GetMembersCanCreatePublicPages()(*bool)
    GetMembersCanCreatePublicRepositories()(*bool)
    GetMembersCanCreateRepositories()(*bool)
    GetMembersCanForkPrivateRepositories()(*bool)
    GetName()(*string)
    GetSecretScanningEnabledForNewRepositories()(*bool)
    GetSecretScanningPushProtectionCustomLink()(*string)
    GetSecretScanningPushProtectionCustomLinkEnabled()(*bool)
    GetSecretScanningPushProtectionEnabledForNewRepositories()(*bool)
    GetTwitterUsername()(*string)
    GetWebCommitSignoffRequired()(*bool)
    SetAdvancedSecurityEnabledForNewRepositories(value *bool)()
    SetBillingEmail(value *string)()
    SetBlog(value *string)()
    SetCompany(value *string)()
    SetDependabotAlertsEnabledForNewRepositories(value *bool)()
    SetDependabotSecurityUpdatesEnabledForNewRepositories(value *bool)()
    SetDependencyGraphEnabledForNewRepositories(value *bool)()
    SetDescription(value *string)()
    SetEmail(value *string)()
    SetHasOrganizationProjects(value *bool)()
    SetHasRepositoryProjects(value *bool)()
    SetLocation(value *string)()
    SetMembersCanCreateInternalRepositories(value *bool)()
    SetMembersCanCreatePages(value *bool)()
    SetMembersCanCreatePrivatePages(value *bool)()
    SetMembersCanCreatePrivateRepositories(value *bool)()
    SetMembersCanCreatePublicPages(value *bool)()
    SetMembersCanCreatePublicRepositories(value *bool)()
    SetMembersCanCreateRepositories(value *bool)()
    SetMembersCanForkPrivateRepositories(value *bool)()
    SetName(value *string)()
    SetSecretScanningEnabledForNewRepositories(value *bool)()
    SetSecretScanningPushProtectionCustomLink(value *string)()
    SetSecretScanningPushProtectionCustomLinkEnabled(value *bool)()
    SetSecretScanningPushProtectionEnabledForNewRepositories(value *bool)()
    SetTwitterUsername(value *string)()
    SetWebCommitSignoffRequired(value *bool)()
}
