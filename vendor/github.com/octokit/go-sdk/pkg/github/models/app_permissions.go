package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppPermissions the permissions granted to the user access token.
type AppPermissions struct {
    // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
    actions *AppPermissions_actions
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
    administration *AppPermissions_administration
    // The level of permission to grant the access token for checks on code.
    checks *AppPermissions_checks
    // The level of permission to grant the access token to create, edit, delete, and list Codespaces.
    codespaces *AppPermissions_codespaces
    // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
    contents *AppPermissions_contents
    // The leve of permission to grant the access token to manage Dependabot secrets.
    dependabot_secrets *AppPermissions_dependabot_secrets
    // The level of permission to grant the access token for deployments and deployment statuses.
    deployments *AppPermissions_deployments
    // The level of permission to grant the access token to manage the email addresses belonging to a user.
    email_addresses *AppPermissions_email_addresses
    // The level of permission to grant the access token for managing repository environments.
    environments *AppPermissions_environments
    // The level of permission to grant the access token to manage the followers belonging to a user.
    followers *AppPermissions_followers
    // The level of permission to grant the access token to manage git SSH keys.
    git_ssh_keys *AppPermissions_git_ssh_keys
    // The level of permission to grant the access token to view and manage GPG keys belonging to a user.
    gpg_keys *AppPermissions_gpg_keys
    // The level of permission to grant the access token to view and manage interaction limits on a repository.
    interaction_limits *AppPermissions_interaction_limits
    // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
    issues *AppPermissions_issues
    // The level of permission to grant the access token for organization teams and members.
    members *AppPermissions_members
    // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
    metadata *AppPermissions_metadata
    // The level of permission to grant the access token to manage access to an organization.
    organization_administration *AppPermissions_organization_administration
    // The level of permission to grant the access token to view and manage announcement banners for an organization.
    organization_announcement_banners *AppPermissions_organization_announcement_banners
    // The level of permission to grant the access token for managing access to GitHub Copilot for members of an organization with a Copilot Business subscription. This property is in beta and is subject to change.
    organization_copilot_seat_management *AppPermissions_organization_copilot_seat_management
    // The level of permission to grant the access token for custom organization roles management.
    organization_custom_org_roles *AppPermissions_organization_custom_org_roles
    // The level of permission to grant the access token for custom property management.
    organization_custom_properties *AppPermissions_organization_custom_properties
    // The level of permission to grant the access token for custom repository roles management.
    organization_custom_roles *AppPermissions_organization_custom_roles
    // The level of permission to grant the access token to view events triggered by an activity in an organization.
    organization_events *AppPermissions_organization_events
    // The level of permission to grant the access token to manage the post-receive hooks for an organization.
    organization_hooks *AppPermissions_organization_hooks
    // The level of permission to grant the access token for organization packages published to GitHub Packages.
    organization_packages *AppPermissions_organization_packages
    // The level of permission to grant the access token for viewing and managing fine-grained personal access tokens that have been approved by an organization.
    organization_personal_access_token_requests *AppPermissions_organization_personal_access_token_requests
    // The level of permission to grant the access token for viewing and managing fine-grained personal access token requests to an organization.
    organization_personal_access_tokens *AppPermissions_organization_personal_access_tokens
    // The level of permission to grant the access token for viewing an organization's plan.
    organization_plan *AppPermissions_organization_plan
    // The level of permission to grant the access token to manage organization projects and projects beta (where available).
    organization_projects *AppPermissions_organization_projects
    // The level of permission to grant the access token to manage organization secrets.
    organization_secrets *AppPermissions_organization_secrets
    // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
    organization_self_hosted_runners *AppPermissions_organization_self_hosted_runners
    // The level of permission to grant the access token to view and manage users blocked by the organization.
    organization_user_blocking *AppPermissions_organization_user_blocking
    // The level of permission to grant the access token for packages published to GitHub Packages.
    packages *AppPermissions_packages
    // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
    pages *AppPermissions_pages
    // The level of permission to grant the access token to manage the profile settings belonging to a user.
    profile *AppPermissions_profile
    // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
    pull_requests *AppPermissions_pull_requests
    // The level of permission to grant the access token to view and edit custom properties for a repository, when allowed by the property.
    repository_custom_properties *AppPermissions_repository_custom_properties
    // The level of permission to grant the access token to manage the post-receive hooks for a repository.
    repository_hooks *AppPermissions_repository_hooks
    // The level of permission to grant the access token to manage repository projects, columns, and cards.
    repository_projects *AppPermissions_repository_projects
    // The level of permission to grant the access token to view and manage secret scanning alerts.
    secret_scanning_alerts *AppPermissions_secret_scanning_alerts
    // The level of permission to grant the access token to manage repository secrets.
    secrets *AppPermissions_secrets
    // The level of permission to grant the access token to view and manage security events like code scanning alerts.
    security_events *AppPermissions_security_events
    // The level of permission to grant the access token to manage just a single file.
    single_file *AppPermissions_single_file
    // The level of permission to grant the access token to list and manage repositories a user is starring.
    starring *AppPermissions_starring
    // The level of permission to grant the access token for commit statuses.
    statuses *AppPermissions_statuses
    // The level of permission to grant the access token to manage team discussions and related comments.
    team_discussions *AppPermissions_team_discussions
    // The level of permission to grant the access token to manage Dependabot alerts.
    vulnerability_alerts *AppPermissions_vulnerability_alerts
    // The level of permission to grant the access token to update GitHub Actions workflow files.
    workflows *AppPermissions_workflows
}
// NewAppPermissions instantiates a new AppPermissions and sets the default values.
func NewAppPermissions()(*AppPermissions) {
    m := &AppPermissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppPermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppPermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppPermissions(), nil
}
// GetActions gets the actions property value. The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
// returns a *AppPermissions_actions when successful
func (m *AppPermissions) GetActions()(*AppPermissions_actions) {
    return m.actions
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppPermissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdministration gets the administration property value. The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
// returns a *AppPermissions_administration when successful
func (m *AppPermissions) GetAdministration()(*AppPermissions_administration) {
    return m.administration
}
// GetChecks gets the checks property value. The level of permission to grant the access token for checks on code.
// returns a *AppPermissions_checks when successful
func (m *AppPermissions) GetChecks()(*AppPermissions_checks) {
    return m.checks
}
// GetCodespaces gets the codespaces property value. The level of permission to grant the access token to create, edit, delete, and list Codespaces.
// returns a *AppPermissions_codespaces when successful
func (m *AppPermissions) GetCodespaces()(*AppPermissions_codespaces) {
    return m.codespaces
}
// GetContents gets the contents property value. The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
// returns a *AppPermissions_contents when successful
func (m *AppPermissions) GetContents()(*AppPermissions_contents) {
    return m.contents
}
// GetDependabotSecrets gets the dependabot_secrets property value. The leve of permission to grant the access token to manage Dependabot secrets.
// returns a *AppPermissions_dependabot_secrets when successful
func (m *AppPermissions) GetDependabotSecrets()(*AppPermissions_dependabot_secrets) {
    return m.dependabot_secrets
}
// GetDeployments gets the deployments property value. The level of permission to grant the access token for deployments and deployment statuses.
// returns a *AppPermissions_deployments when successful
func (m *AppPermissions) GetDeployments()(*AppPermissions_deployments) {
    return m.deployments
}
// GetEmailAddresses gets the email_addresses property value. The level of permission to grant the access token to manage the email addresses belonging to a user.
// returns a *AppPermissions_email_addresses when successful
func (m *AppPermissions) GetEmailAddresses()(*AppPermissions_email_addresses) {
    return m.email_addresses
}
// GetEnvironments gets the environments property value. The level of permission to grant the access token for managing repository environments.
// returns a *AppPermissions_environments when successful
func (m *AppPermissions) GetEnvironments()(*AppPermissions_environments) {
    return m.environments
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppPermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_actions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActions(val.(*AppPermissions_actions))
        }
        return nil
    }
    res["administration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_administration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdministration(val.(*AppPermissions_administration))
        }
        return nil
    }
    res["checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_checks)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecks(val.(*AppPermissions_checks))
        }
        return nil
    }
    res["codespaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_codespaces)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodespaces(val.(*AppPermissions_codespaces))
        }
        return nil
    }
    res["contents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_contents)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContents(val.(*AppPermissions_contents))
        }
        return nil
    }
    res["dependabot_secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_dependabot_secrets)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotSecrets(val.(*AppPermissions_dependabot_secrets))
        }
        return nil
    }
    res["deployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_deployments)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployments(val.(*AppPermissions_deployments))
        }
        return nil
    }
    res["email_addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_email_addresses)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddresses(val.(*AppPermissions_email_addresses))
        }
        return nil
    }
    res["environments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_environments)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironments(val.(*AppPermissions_environments))
        }
        return nil
    }
    res["followers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_followers)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFollowers(val.(*AppPermissions_followers))
        }
        return nil
    }
    res["git_ssh_keys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_git_ssh_keys)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitSshKeys(val.(*AppPermissions_git_ssh_keys))
        }
        return nil
    }
    res["gpg_keys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_gpg_keys)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGpgKeys(val.(*AppPermissions_gpg_keys))
        }
        return nil
    }
    res["interaction_limits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_interaction_limits)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInteractionLimits(val.(*AppPermissions_interaction_limits))
        }
        return nil
    }
    res["issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_issues)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssues(val.(*AppPermissions_issues))
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_members)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembers(val.(*AppPermissions_members))
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_metadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val.(*AppPermissions_metadata))
        }
        return nil
    }
    res["organization_administration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_administration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationAdministration(val.(*AppPermissions_organization_administration))
        }
        return nil
    }
    res["organization_announcement_banners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_announcement_banners)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationAnnouncementBanners(val.(*AppPermissions_organization_announcement_banners))
        }
        return nil
    }
    res["organization_copilot_seat_management"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_copilot_seat_management)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationCopilotSeatManagement(val.(*AppPermissions_organization_copilot_seat_management))
        }
        return nil
    }
    res["organization_custom_org_roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_custom_org_roles)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationCustomOrgRoles(val.(*AppPermissions_organization_custom_org_roles))
        }
        return nil
    }
    res["organization_custom_properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_custom_properties)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationCustomProperties(val.(*AppPermissions_organization_custom_properties))
        }
        return nil
    }
    res["organization_custom_roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_custom_roles)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationCustomRoles(val.(*AppPermissions_organization_custom_roles))
        }
        return nil
    }
    res["organization_events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_events)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationEvents(val.(*AppPermissions_organization_events))
        }
        return nil
    }
    res["organization_hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_hooks)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationHooks(val.(*AppPermissions_organization_hooks))
        }
        return nil
    }
    res["organization_packages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_packages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationPackages(val.(*AppPermissions_organization_packages))
        }
        return nil
    }
    res["organization_personal_access_token_requests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_personal_access_token_requests)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationPersonalAccessTokenRequests(val.(*AppPermissions_organization_personal_access_token_requests))
        }
        return nil
    }
    res["organization_personal_access_tokens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_personal_access_tokens)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationPersonalAccessTokens(val.(*AppPermissions_organization_personal_access_tokens))
        }
        return nil
    }
    res["organization_plan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_plan)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationPlan(val.(*AppPermissions_organization_plan))
        }
        return nil
    }
    res["organization_projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_projects)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationProjects(val.(*AppPermissions_organization_projects))
        }
        return nil
    }
    res["organization_secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_secrets)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationSecrets(val.(*AppPermissions_organization_secrets))
        }
        return nil
    }
    res["organization_self_hosted_runners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_self_hosted_runners)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationSelfHostedRunners(val.(*AppPermissions_organization_self_hosted_runners))
        }
        return nil
    }
    res["organization_user_blocking"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_organization_user_blocking)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationUserBlocking(val.(*AppPermissions_organization_user_blocking))
        }
        return nil
    }
    res["packages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_packages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackages(val.(*AppPermissions_packages))
        }
        return nil
    }
    res["pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_pages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPages(val.(*AppPermissions_pages))
        }
        return nil
    }
    res["profile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_profile)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfile(val.(*AppPermissions_profile))
        }
        return nil
    }
    res["pull_requests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_pull_requests)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequests(val.(*AppPermissions_pull_requests))
        }
        return nil
    }
    res["repository_custom_properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_repository_custom_properties)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryCustomProperties(val.(*AppPermissions_repository_custom_properties))
        }
        return nil
    }
    res["repository_hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_repository_hooks)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryHooks(val.(*AppPermissions_repository_hooks))
        }
        return nil
    }
    res["repository_projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_repository_projects)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryProjects(val.(*AppPermissions_repository_projects))
        }
        return nil
    }
    res["secret_scanning_alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_secret_scanning_alerts)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningAlerts(val.(*AppPermissions_secret_scanning_alerts))
        }
        return nil
    }
    res["secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_secrets)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecrets(val.(*AppPermissions_secrets))
        }
        return nil
    }
    res["security_events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_security_events)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEvents(val.(*AppPermissions_security_events))
        }
        return nil
    }
    res["single_file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_single_file)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleFile(val.(*AppPermissions_single_file))
        }
        return nil
    }
    res["starring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_starring)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStarring(val.(*AppPermissions_starring))
        }
        return nil
    }
    res["statuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_statuses)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatuses(val.(*AppPermissions_statuses))
        }
        return nil
    }
    res["team_discussions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_team_discussions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamDiscussions(val.(*AppPermissions_team_discussions))
        }
        return nil
    }
    res["vulnerability_alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_vulnerability_alerts)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVulnerabilityAlerts(val.(*AppPermissions_vulnerability_alerts))
        }
        return nil
    }
    res["workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppPermissions_workflows)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflows(val.(*AppPermissions_workflows))
        }
        return nil
    }
    return res
}
// GetFollowers gets the followers property value. The level of permission to grant the access token to manage the followers belonging to a user.
// returns a *AppPermissions_followers when successful
func (m *AppPermissions) GetFollowers()(*AppPermissions_followers) {
    return m.followers
}
// GetGitSshKeys gets the git_ssh_keys property value. The level of permission to grant the access token to manage git SSH keys.
// returns a *AppPermissions_git_ssh_keys when successful
func (m *AppPermissions) GetGitSshKeys()(*AppPermissions_git_ssh_keys) {
    return m.git_ssh_keys
}
// GetGpgKeys gets the gpg_keys property value. The level of permission to grant the access token to view and manage GPG keys belonging to a user.
// returns a *AppPermissions_gpg_keys when successful
func (m *AppPermissions) GetGpgKeys()(*AppPermissions_gpg_keys) {
    return m.gpg_keys
}
// GetInteractionLimits gets the interaction_limits property value. The level of permission to grant the access token to view and manage interaction limits on a repository.
// returns a *AppPermissions_interaction_limits when successful
func (m *AppPermissions) GetInteractionLimits()(*AppPermissions_interaction_limits) {
    return m.interaction_limits
}
// GetIssues gets the issues property value. The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
// returns a *AppPermissions_issues when successful
func (m *AppPermissions) GetIssues()(*AppPermissions_issues) {
    return m.issues
}
// GetMembers gets the members property value. The level of permission to grant the access token for organization teams and members.
// returns a *AppPermissions_members when successful
func (m *AppPermissions) GetMembers()(*AppPermissions_members) {
    return m.members
}
// GetMetadata gets the metadata property value. The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
// returns a *AppPermissions_metadata when successful
func (m *AppPermissions) GetMetadata()(*AppPermissions_metadata) {
    return m.metadata
}
// GetOrganizationAdministration gets the organization_administration property value. The level of permission to grant the access token to manage access to an organization.
// returns a *AppPermissions_organization_administration when successful
func (m *AppPermissions) GetOrganizationAdministration()(*AppPermissions_organization_administration) {
    return m.organization_administration
}
// GetOrganizationAnnouncementBanners gets the organization_announcement_banners property value. The level of permission to grant the access token to view and manage announcement banners for an organization.
// returns a *AppPermissions_organization_announcement_banners when successful
func (m *AppPermissions) GetOrganizationAnnouncementBanners()(*AppPermissions_organization_announcement_banners) {
    return m.organization_announcement_banners
}
// GetOrganizationCopilotSeatManagement gets the organization_copilot_seat_management property value. The level of permission to grant the access token for managing access to GitHub Copilot for members of an organization with a Copilot Business subscription. This property is in beta and is subject to change.
// returns a *AppPermissions_organization_copilot_seat_management when successful
func (m *AppPermissions) GetOrganizationCopilotSeatManagement()(*AppPermissions_organization_copilot_seat_management) {
    return m.organization_copilot_seat_management
}
// GetOrganizationCustomOrgRoles gets the organization_custom_org_roles property value. The level of permission to grant the access token for custom organization roles management.
// returns a *AppPermissions_organization_custom_org_roles when successful
func (m *AppPermissions) GetOrganizationCustomOrgRoles()(*AppPermissions_organization_custom_org_roles) {
    return m.organization_custom_org_roles
}
// GetOrganizationCustomProperties gets the organization_custom_properties property value. The level of permission to grant the access token for custom property management.
// returns a *AppPermissions_organization_custom_properties when successful
func (m *AppPermissions) GetOrganizationCustomProperties()(*AppPermissions_organization_custom_properties) {
    return m.organization_custom_properties
}
// GetOrganizationCustomRoles gets the organization_custom_roles property value. The level of permission to grant the access token for custom repository roles management.
// returns a *AppPermissions_organization_custom_roles when successful
func (m *AppPermissions) GetOrganizationCustomRoles()(*AppPermissions_organization_custom_roles) {
    return m.organization_custom_roles
}
// GetOrganizationEvents gets the organization_events property value. The level of permission to grant the access token to view events triggered by an activity in an organization.
// returns a *AppPermissions_organization_events when successful
func (m *AppPermissions) GetOrganizationEvents()(*AppPermissions_organization_events) {
    return m.organization_events
}
// GetOrganizationHooks gets the organization_hooks property value. The level of permission to grant the access token to manage the post-receive hooks for an organization.
// returns a *AppPermissions_organization_hooks when successful
func (m *AppPermissions) GetOrganizationHooks()(*AppPermissions_organization_hooks) {
    return m.organization_hooks
}
// GetOrganizationPackages gets the organization_packages property value. The level of permission to grant the access token for organization packages published to GitHub Packages.
// returns a *AppPermissions_organization_packages when successful
func (m *AppPermissions) GetOrganizationPackages()(*AppPermissions_organization_packages) {
    return m.organization_packages
}
// GetOrganizationPersonalAccessTokenRequests gets the organization_personal_access_token_requests property value. The level of permission to grant the access token for viewing and managing fine-grained personal access tokens that have been approved by an organization.
// returns a *AppPermissions_organization_personal_access_token_requests when successful
func (m *AppPermissions) GetOrganizationPersonalAccessTokenRequests()(*AppPermissions_organization_personal_access_token_requests) {
    return m.organization_personal_access_token_requests
}
// GetOrganizationPersonalAccessTokens gets the organization_personal_access_tokens property value. The level of permission to grant the access token for viewing and managing fine-grained personal access token requests to an organization.
// returns a *AppPermissions_organization_personal_access_tokens when successful
func (m *AppPermissions) GetOrganizationPersonalAccessTokens()(*AppPermissions_organization_personal_access_tokens) {
    return m.organization_personal_access_tokens
}
// GetOrganizationPlan gets the organization_plan property value. The level of permission to grant the access token for viewing an organization's plan.
// returns a *AppPermissions_organization_plan when successful
func (m *AppPermissions) GetOrganizationPlan()(*AppPermissions_organization_plan) {
    return m.organization_plan
}
// GetOrganizationProjects gets the organization_projects property value. The level of permission to grant the access token to manage organization projects and projects beta (where available).
// returns a *AppPermissions_organization_projects when successful
func (m *AppPermissions) GetOrganizationProjects()(*AppPermissions_organization_projects) {
    return m.organization_projects
}
// GetOrganizationSecrets gets the organization_secrets property value. The level of permission to grant the access token to manage organization secrets.
// returns a *AppPermissions_organization_secrets when successful
func (m *AppPermissions) GetOrganizationSecrets()(*AppPermissions_organization_secrets) {
    return m.organization_secrets
}
// GetOrganizationSelfHostedRunners gets the organization_self_hosted_runners property value. The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
// returns a *AppPermissions_organization_self_hosted_runners when successful
func (m *AppPermissions) GetOrganizationSelfHostedRunners()(*AppPermissions_organization_self_hosted_runners) {
    return m.organization_self_hosted_runners
}
// GetOrganizationUserBlocking gets the organization_user_blocking property value. The level of permission to grant the access token to view and manage users blocked by the organization.
// returns a *AppPermissions_organization_user_blocking when successful
func (m *AppPermissions) GetOrganizationUserBlocking()(*AppPermissions_organization_user_blocking) {
    return m.organization_user_blocking
}
// GetPackages gets the packages property value. The level of permission to grant the access token for packages published to GitHub Packages.
// returns a *AppPermissions_packages when successful
func (m *AppPermissions) GetPackages()(*AppPermissions_packages) {
    return m.packages
}
// GetPages gets the pages property value. The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
// returns a *AppPermissions_pages when successful
func (m *AppPermissions) GetPages()(*AppPermissions_pages) {
    return m.pages
}
// GetProfile gets the profile property value. The level of permission to grant the access token to manage the profile settings belonging to a user.
// returns a *AppPermissions_profile when successful
func (m *AppPermissions) GetProfile()(*AppPermissions_profile) {
    return m.profile
}
// GetPullRequests gets the pull_requests property value. The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
// returns a *AppPermissions_pull_requests when successful
func (m *AppPermissions) GetPullRequests()(*AppPermissions_pull_requests) {
    return m.pull_requests
}
// GetRepositoryCustomProperties gets the repository_custom_properties property value. The level of permission to grant the access token to view and edit custom properties for a repository, when allowed by the property.
// returns a *AppPermissions_repository_custom_properties when successful
func (m *AppPermissions) GetRepositoryCustomProperties()(*AppPermissions_repository_custom_properties) {
    return m.repository_custom_properties
}
// GetRepositoryHooks gets the repository_hooks property value. The level of permission to grant the access token to manage the post-receive hooks for a repository.
// returns a *AppPermissions_repository_hooks when successful
func (m *AppPermissions) GetRepositoryHooks()(*AppPermissions_repository_hooks) {
    return m.repository_hooks
}
// GetRepositoryProjects gets the repository_projects property value. The level of permission to grant the access token to manage repository projects, columns, and cards.
// returns a *AppPermissions_repository_projects when successful
func (m *AppPermissions) GetRepositoryProjects()(*AppPermissions_repository_projects) {
    return m.repository_projects
}
// GetSecrets gets the secrets property value. The level of permission to grant the access token to manage repository secrets.
// returns a *AppPermissions_secrets when successful
func (m *AppPermissions) GetSecrets()(*AppPermissions_secrets) {
    return m.secrets
}
// GetSecretScanningAlerts gets the secret_scanning_alerts property value. The level of permission to grant the access token to view and manage secret scanning alerts.
// returns a *AppPermissions_secret_scanning_alerts when successful
func (m *AppPermissions) GetSecretScanningAlerts()(*AppPermissions_secret_scanning_alerts) {
    return m.secret_scanning_alerts
}
// GetSecurityEvents gets the security_events property value. The level of permission to grant the access token to view and manage security events like code scanning alerts.
// returns a *AppPermissions_security_events when successful
func (m *AppPermissions) GetSecurityEvents()(*AppPermissions_security_events) {
    return m.security_events
}
// GetSingleFile gets the single_file property value. The level of permission to grant the access token to manage just a single file.
// returns a *AppPermissions_single_file when successful
func (m *AppPermissions) GetSingleFile()(*AppPermissions_single_file) {
    return m.single_file
}
// GetStarring gets the starring property value. The level of permission to grant the access token to list and manage repositories a user is starring.
// returns a *AppPermissions_starring when successful
func (m *AppPermissions) GetStarring()(*AppPermissions_starring) {
    return m.starring
}
// GetStatuses gets the statuses property value. The level of permission to grant the access token for commit statuses.
// returns a *AppPermissions_statuses when successful
func (m *AppPermissions) GetStatuses()(*AppPermissions_statuses) {
    return m.statuses
}
// GetTeamDiscussions gets the team_discussions property value. The level of permission to grant the access token to manage team discussions and related comments.
// returns a *AppPermissions_team_discussions when successful
func (m *AppPermissions) GetTeamDiscussions()(*AppPermissions_team_discussions) {
    return m.team_discussions
}
// GetVulnerabilityAlerts gets the vulnerability_alerts property value. The level of permission to grant the access token to manage Dependabot alerts.
// returns a *AppPermissions_vulnerability_alerts when successful
func (m *AppPermissions) GetVulnerabilityAlerts()(*AppPermissions_vulnerability_alerts) {
    return m.vulnerability_alerts
}
// GetWorkflows gets the workflows property value. The level of permission to grant the access token to update GitHub Actions workflow files.
// returns a *AppPermissions_workflows when successful
func (m *AppPermissions) GetWorkflows()(*AppPermissions_workflows) {
    return m.workflows
}
// Serialize serializes information the current object
func (m *AppPermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        cast := (*m.GetActions()).String()
        err := writer.WriteStringValue("actions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAdministration() != nil {
        cast := (*m.GetAdministration()).String()
        err := writer.WriteStringValue("administration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetChecks() != nil {
        cast := (*m.GetChecks()).String()
        err := writer.WriteStringValue("checks", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCodespaces() != nil {
        cast := (*m.GetCodespaces()).String()
        err := writer.WriteStringValue("codespaces", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetContents() != nil {
        cast := (*m.GetContents()).String()
        err := writer.WriteStringValue("contents", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependabotSecrets() != nil {
        cast := (*m.GetDependabotSecrets()).String()
        err := writer.WriteStringValue("dependabot_secrets", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeployments() != nil {
        cast := (*m.GetDeployments()).String()
        err := writer.WriteStringValue("deployments", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmailAddresses() != nil {
        cast := (*m.GetEmailAddresses()).String()
        err := writer.WriteStringValue("email_addresses", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnvironments() != nil {
        cast := (*m.GetEnvironments()).String()
        err := writer.WriteStringValue("environments", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFollowers() != nil {
        cast := (*m.GetFollowers()).String()
        err := writer.WriteStringValue("followers", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetGitSshKeys() != nil {
        cast := (*m.GetGitSshKeys()).String()
        err := writer.WriteStringValue("git_ssh_keys", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetGpgKeys() != nil {
        cast := (*m.GetGpgKeys()).String()
        err := writer.WriteStringValue("gpg_keys", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInteractionLimits() != nil {
        cast := (*m.GetInteractionLimits()).String()
        err := writer.WriteStringValue("interaction_limits", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIssues() != nil {
        cast := (*m.GetIssues()).String()
        err := writer.WriteStringValue("issues", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := (*m.GetMembers()).String()
        err := writer.WriteStringValue("members", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := (*m.GetMetadata()).String()
        err := writer.WriteStringValue("metadata", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationAdministration() != nil {
        cast := (*m.GetOrganizationAdministration()).String()
        err := writer.WriteStringValue("organization_administration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationAnnouncementBanners() != nil {
        cast := (*m.GetOrganizationAnnouncementBanners()).String()
        err := writer.WriteStringValue("organization_announcement_banners", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationCopilotSeatManagement() != nil {
        cast := (*m.GetOrganizationCopilotSeatManagement()).String()
        err := writer.WriteStringValue("organization_copilot_seat_management", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationCustomOrgRoles() != nil {
        cast := (*m.GetOrganizationCustomOrgRoles()).String()
        err := writer.WriteStringValue("organization_custom_org_roles", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationCustomProperties() != nil {
        cast := (*m.GetOrganizationCustomProperties()).String()
        err := writer.WriteStringValue("organization_custom_properties", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationCustomRoles() != nil {
        cast := (*m.GetOrganizationCustomRoles()).String()
        err := writer.WriteStringValue("organization_custom_roles", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationEvents() != nil {
        cast := (*m.GetOrganizationEvents()).String()
        err := writer.WriteStringValue("organization_events", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationHooks() != nil {
        cast := (*m.GetOrganizationHooks()).String()
        err := writer.WriteStringValue("organization_hooks", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationPackages() != nil {
        cast := (*m.GetOrganizationPackages()).String()
        err := writer.WriteStringValue("organization_packages", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationPersonalAccessTokens() != nil {
        cast := (*m.GetOrganizationPersonalAccessTokens()).String()
        err := writer.WriteStringValue("organization_personal_access_tokens", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationPersonalAccessTokenRequests() != nil {
        cast := (*m.GetOrganizationPersonalAccessTokenRequests()).String()
        err := writer.WriteStringValue("organization_personal_access_token_requests", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationPlan() != nil {
        cast := (*m.GetOrganizationPlan()).String()
        err := writer.WriteStringValue("organization_plan", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationProjects() != nil {
        cast := (*m.GetOrganizationProjects()).String()
        err := writer.WriteStringValue("organization_projects", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationSecrets() != nil {
        cast := (*m.GetOrganizationSecrets()).String()
        err := writer.WriteStringValue("organization_secrets", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationSelfHostedRunners() != nil {
        cast := (*m.GetOrganizationSelfHostedRunners()).String()
        err := writer.WriteStringValue("organization_self_hosted_runners", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationUserBlocking() != nil {
        cast := (*m.GetOrganizationUserBlocking()).String()
        err := writer.WriteStringValue("organization_user_blocking", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPackages() != nil {
        cast := (*m.GetPackages()).String()
        err := writer.WriteStringValue("packages", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPages() != nil {
        cast := (*m.GetPages()).String()
        err := writer.WriteStringValue("pages", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetProfile() != nil {
        cast := (*m.GetProfile()).String()
        err := writer.WriteStringValue("profile", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPullRequests() != nil {
        cast := (*m.GetPullRequests()).String()
        err := writer.WriteStringValue("pull_requests", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryCustomProperties() != nil {
        cast := (*m.GetRepositoryCustomProperties()).String()
        err := writer.WriteStringValue("repository_custom_properties", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryHooks() != nil {
        cast := (*m.GetRepositoryHooks()).String()
        err := writer.WriteStringValue("repository_hooks", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryProjects() != nil {
        cast := (*m.GetRepositoryProjects()).String()
        err := writer.WriteStringValue("repository_projects", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecrets() != nil {
        cast := (*m.GetSecrets()).String()
        err := writer.WriteStringValue("secrets", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecretScanningAlerts() != nil {
        cast := (*m.GetSecretScanningAlerts()).String()
        err := writer.WriteStringValue("secret_scanning_alerts", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityEvents() != nil {
        cast := (*m.GetSecurityEvents()).String()
        err := writer.WriteStringValue("security_events", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSingleFile() != nil {
        cast := (*m.GetSingleFile()).String()
        err := writer.WriteStringValue("single_file", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStarring() != nil {
        cast := (*m.GetStarring()).String()
        err := writer.WriteStringValue("starring", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatuses() != nil {
        cast := (*m.GetStatuses()).String()
        err := writer.WriteStringValue("statuses", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTeamDiscussions() != nil {
        cast := (*m.GetTeamDiscussions()).String()
        err := writer.WriteStringValue("team_discussions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVulnerabilityAlerts() != nil {
        cast := (*m.GetVulnerabilityAlerts()).String()
        err := writer.WriteStringValue("vulnerability_alerts", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflows() != nil {
        cast := (*m.GetWorkflows()).String()
        err := writer.WriteStringValue("workflows", &cast)
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
// SetActions sets the actions property value. The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
func (m *AppPermissions) SetActions(value *AppPermissions_actions)() {
    m.actions = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppPermissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdministration sets the administration property value. The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
func (m *AppPermissions) SetAdministration(value *AppPermissions_administration)() {
    m.administration = value
}
// SetChecks sets the checks property value. The level of permission to grant the access token for checks on code.
func (m *AppPermissions) SetChecks(value *AppPermissions_checks)() {
    m.checks = value
}
// SetCodespaces sets the codespaces property value. The level of permission to grant the access token to create, edit, delete, and list Codespaces.
func (m *AppPermissions) SetCodespaces(value *AppPermissions_codespaces)() {
    m.codespaces = value
}
// SetContents sets the contents property value. The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
func (m *AppPermissions) SetContents(value *AppPermissions_contents)() {
    m.contents = value
}
// SetDependabotSecrets sets the dependabot_secrets property value. The leve of permission to grant the access token to manage Dependabot secrets.
func (m *AppPermissions) SetDependabotSecrets(value *AppPermissions_dependabot_secrets)() {
    m.dependabot_secrets = value
}
// SetDeployments sets the deployments property value. The level of permission to grant the access token for deployments and deployment statuses.
func (m *AppPermissions) SetDeployments(value *AppPermissions_deployments)() {
    m.deployments = value
}
// SetEmailAddresses sets the email_addresses property value. The level of permission to grant the access token to manage the email addresses belonging to a user.
func (m *AppPermissions) SetEmailAddresses(value *AppPermissions_email_addresses)() {
    m.email_addresses = value
}
// SetEnvironments sets the environments property value. The level of permission to grant the access token for managing repository environments.
func (m *AppPermissions) SetEnvironments(value *AppPermissions_environments)() {
    m.environments = value
}
// SetFollowers sets the followers property value. The level of permission to grant the access token to manage the followers belonging to a user.
func (m *AppPermissions) SetFollowers(value *AppPermissions_followers)() {
    m.followers = value
}
// SetGitSshKeys sets the git_ssh_keys property value. The level of permission to grant the access token to manage git SSH keys.
func (m *AppPermissions) SetGitSshKeys(value *AppPermissions_git_ssh_keys)() {
    m.git_ssh_keys = value
}
// SetGpgKeys sets the gpg_keys property value. The level of permission to grant the access token to view and manage GPG keys belonging to a user.
func (m *AppPermissions) SetGpgKeys(value *AppPermissions_gpg_keys)() {
    m.gpg_keys = value
}
// SetInteractionLimits sets the interaction_limits property value. The level of permission to grant the access token to view and manage interaction limits on a repository.
func (m *AppPermissions) SetInteractionLimits(value *AppPermissions_interaction_limits)() {
    m.interaction_limits = value
}
// SetIssues sets the issues property value. The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
func (m *AppPermissions) SetIssues(value *AppPermissions_issues)() {
    m.issues = value
}
// SetMembers sets the members property value. The level of permission to grant the access token for organization teams and members.
func (m *AppPermissions) SetMembers(value *AppPermissions_members)() {
    m.members = value
}
// SetMetadata sets the metadata property value. The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
func (m *AppPermissions) SetMetadata(value *AppPermissions_metadata)() {
    m.metadata = value
}
// SetOrganizationAdministration sets the organization_administration property value. The level of permission to grant the access token to manage access to an organization.
func (m *AppPermissions) SetOrganizationAdministration(value *AppPermissions_organization_administration)() {
    m.organization_administration = value
}
// SetOrganizationAnnouncementBanners sets the organization_announcement_banners property value. The level of permission to grant the access token to view and manage announcement banners for an organization.
func (m *AppPermissions) SetOrganizationAnnouncementBanners(value *AppPermissions_organization_announcement_banners)() {
    m.organization_announcement_banners = value
}
// SetOrganizationCopilotSeatManagement sets the organization_copilot_seat_management property value. The level of permission to grant the access token for managing access to GitHub Copilot for members of an organization with a Copilot Business subscription. This property is in beta and is subject to change.
func (m *AppPermissions) SetOrganizationCopilotSeatManagement(value *AppPermissions_organization_copilot_seat_management)() {
    m.organization_copilot_seat_management = value
}
// SetOrganizationCustomOrgRoles sets the organization_custom_org_roles property value. The level of permission to grant the access token for custom organization roles management.
func (m *AppPermissions) SetOrganizationCustomOrgRoles(value *AppPermissions_organization_custom_org_roles)() {
    m.organization_custom_org_roles = value
}
// SetOrganizationCustomProperties sets the organization_custom_properties property value. The level of permission to grant the access token for custom property management.
func (m *AppPermissions) SetOrganizationCustomProperties(value *AppPermissions_organization_custom_properties)() {
    m.organization_custom_properties = value
}
// SetOrganizationCustomRoles sets the organization_custom_roles property value. The level of permission to grant the access token for custom repository roles management.
func (m *AppPermissions) SetOrganizationCustomRoles(value *AppPermissions_organization_custom_roles)() {
    m.organization_custom_roles = value
}
// SetOrganizationEvents sets the organization_events property value. The level of permission to grant the access token to view events triggered by an activity in an organization.
func (m *AppPermissions) SetOrganizationEvents(value *AppPermissions_organization_events)() {
    m.organization_events = value
}
// SetOrganizationHooks sets the organization_hooks property value. The level of permission to grant the access token to manage the post-receive hooks for an organization.
func (m *AppPermissions) SetOrganizationHooks(value *AppPermissions_organization_hooks)() {
    m.organization_hooks = value
}
// SetOrganizationPackages sets the organization_packages property value. The level of permission to grant the access token for organization packages published to GitHub Packages.
func (m *AppPermissions) SetOrganizationPackages(value *AppPermissions_organization_packages)() {
    m.organization_packages = value
}
// SetOrganizationPersonalAccessTokenRequests sets the organization_personal_access_token_requests property value. The level of permission to grant the access token for viewing and managing fine-grained personal access tokens that have been approved by an organization.
func (m *AppPermissions) SetOrganizationPersonalAccessTokenRequests(value *AppPermissions_organization_personal_access_token_requests)() {
    m.organization_personal_access_token_requests = value
}
// SetOrganizationPersonalAccessTokens sets the organization_personal_access_tokens property value. The level of permission to grant the access token for viewing and managing fine-grained personal access token requests to an organization.
func (m *AppPermissions) SetOrganizationPersonalAccessTokens(value *AppPermissions_organization_personal_access_tokens)() {
    m.organization_personal_access_tokens = value
}
// SetOrganizationPlan sets the organization_plan property value. The level of permission to grant the access token for viewing an organization's plan.
func (m *AppPermissions) SetOrganizationPlan(value *AppPermissions_organization_plan)() {
    m.organization_plan = value
}
// SetOrganizationProjects sets the organization_projects property value. The level of permission to grant the access token to manage organization projects and projects beta (where available).
func (m *AppPermissions) SetOrganizationProjects(value *AppPermissions_organization_projects)() {
    m.organization_projects = value
}
// SetOrganizationSecrets sets the organization_secrets property value. The level of permission to grant the access token to manage organization secrets.
func (m *AppPermissions) SetOrganizationSecrets(value *AppPermissions_organization_secrets)() {
    m.organization_secrets = value
}
// SetOrganizationSelfHostedRunners sets the organization_self_hosted_runners property value. The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
func (m *AppPermissions) SetOrganizationSelfHostedRunners(value *AppPermissions_organization_self_hosted_runners)() {
    m.organization_self_hosted_runners = value
}
// SetOrganizationUserBlocking sets the organization_user_blocking property value. The level of permission to grant the access token to view and manage users blocked by the organization.
func (m *AppPermissions) SetOrganizationUserBlocking(value *AppPermissions_organization_user_blocking)() {
    m.organization_user_blocking = value
}
// SetPackages sets the packages property value. The level of permission to grant the access token for packages published to GitHub Packages.
func (m *AppPermissions) SetPackages(value *AppPermissions_packages)() {
    m.packages = value
}
// SetPages sets the pages property value. The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
func (m *AppPermissions) SetPages(value *AppPermissions_pages)() {
    m.pages = value
}
// SetProfile sets the profile property value. The level of permission to grant the access token to manage the profile settings belonging to a user.
func (m *AppPermissions) SetProfile(value *AppPermissions_profile)() {
    m.profile = value
}
// SetPullRequests sets the pull_requests property value. The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
func (m *AppPermissions) SetPullRequests(value *AppPermissions_pull_requests)() {
    m.pull_requests = value
}
// SetRepositoryCustomProperties sets the repository_custom_properties property value. The level of permission to grant the access token to view and edit custom properties for a repository, when allowed by the property.
func (m *AppPermissions) SetRepositoryCustomProperties(value *AppPermissions_repository_custom_properties)() {
    m.repository_custom_properties = value
}
// SetRepositoryHooks sets the repository_hooks property value. The level of permission to grant the access token to manage the post-receive hooks for a repository.
func (m *AppPermissions) SetRepositoryHooks(value *AppPermissions_repository_hooks)() {
    m.repository_hooks = value
}
// SetRepositoryProjects sets the repository_projects property value. The level of permission to grant the access token to manage repository projects, columns, and cards.
func (m *AppPermissions) SetRepositoryProjects(value *AppPermissions_repository_projects)() {
    m.repository_projects = value
}
// SetSecrets sets the secrets property value. The level of permission to grant the access token to manage repository secrets.
func (m *AppPermissions) SetSecrets(value *AppPermissions_secrets)() {
    m.secrets = value
}
// SetSecretScanningAlerts sets the secret_scanning_alerts property value. The level of permission to grant the access token to view and manage secret scanning alerts.
func (m *AppPermissions) SetSecretScanningAlerts(value *AppPermissions_secret_scanning_alerts)() {
    m.secret_scanning_alerts = value
}
// SetSecurityEvents sets the security_events property value. The level of permission to grant the access token to view and manage security events like code scanning alerts.
func (m *AppPermissions) SetSecurityEvents(value *AppPermissions_security_events)() {
    m.security_events = value
}
// SetSingleFile sets the single_file property value. The level of permission to grant the access token to manage just a single file.
func (m *AppPermissions) SetSingleFile(value *AppPermissions_single_file)() {
    m.single_file = value
}
// SetStarring sets the starring property value. The level of permission to grant the access token to list and manage repositories a user is starring.
func (m *AppPermissions) SetStarring(value *AppPermissions_starring)() {
    m.starring = value
}
// SetStatuses sets the statuses property value. The level of permission to grant the access token for commit statuses.
func (m *AppPermissions) SetStatuses(value *AppPermissions_statuses)() {
    m.statuses = value
}
// SetTeamDiscussions sets the team_discussions property value. The level of permission to grant the access token to manage team discussions and related comments.
func (m *AppPermissions) SetTeamDiscussions(value *AppPermissions_team_discussions)() {
    m.team_discussions = value
}
// SetVulnerabilityAlerts sets the vulnerability_alerts property value. The level of permission to grant the access token to manage Dependabot alerts.
func (m *AppPermissions) SetVulnerabilityAlerts(value *AppPermissions_vulnerability_alerts)() {
    m.vulnerability_alerts = value
}
// SetWorkflows sets the workflows property value. The level of permission to grant the access token to update GitHub Actions workflow files.
func (m *AppPermissions) SetWorkflows(value *AppPermissions_workflows)() {
    m.workflows = value
}
type AppPermissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()(*AppPermissions_actions)
    GetAdministration()(*AppPermissions_administration)
    GetChecks()(*AppPermissions_checks)
    GetCodespaces()(*AppPermissions_codespaces)
    GetContents()(*AppPermissions_contents)
    GetDependabotSecrets()(*AppPermissions_dependabot_secrets)
    GetDeployments()(*AppPermissions_deployments)
    GetEmailAddresses()(*AppPermissions_email_addresses)
    GetEnvironments()(*AppPermissions_environments)
    GetFollowers()(*AppPermissions_followers)
    GetGitSshKeys()(*AppPermissions_git_ssh_keys)
    GetGpgKeys()(*AppPermissions_gpg_keys)
    GetInteractionLimits()(*AppPermissions_interaction_limits)
    GetIssues()(*AppPermissions_issues)
    GetMembers()(*AppPermissions_members)
    GetMetadata()(*AppPermissions_metadata)
    GetOrganizationAdministration()(*AppPermissions_organization_administration)
    GetOrganizationAnnouncementBanners()(*AppPermissions_organization_announcement_banners)
    GetOrganizationCopilotSeatManagement()(*AppPermissions_organization_copilot_seat_management)
    GetOrganizationCustomOrgRoles()(*AppPermissions_organization_custom_org_roles)
    GetOrganizationCustomProperties()(*AppPermissions_organization_custom_properties)
    GetOrganizationCustomRoles()(*AppPermissions_organization_custom_roles)
    GetOrganizationEvents()(*AppPermissions_organization_events)
    GetOrganizationHooks()(*AppPermissions_organization_hooks)
    GetOrganizationPackages()(*AppPermissions_organization_packages)
    GetOrganizationPersonalAccessTokenRequests()(*AppPermissions_organization_personal_access_token_requests)
    GetOrganizationPersonalAccessTokens()(*AppPermissions_organization_personal_access_tokens)
    GetOrganizationPlan()(*AppPermissions_organization_plan)
    GetOrganizationProjects()(*AppPermissions_organization_projects)
    GetOrganizationSecrets()(*AppPermissions_organization_secrets)
    GetOrganizationSelfHostedRunners()(*AppPermissions_organization_self_hosted_runners)
    GetOrganizationUserBlocking()(*AppPermissions_organization_user_blocking)
    GetPackages()(*AppPermissions_packages)
    GetPages()(*AppPermissions_pages)
    GetProfile()(*AppPermissions_profile)
    GetPullRequests()(*AppPermissions_pull_requests)
    GetRepositoryCustomProperties()(*AppPermissions_repository_custom_properties)
    GetRepositoryHooks()(*AppPermissions_repository_hooks)
    GetRepositoryProjects()(*AppPermissions_repository_projects)
    GetSecrets()(*AppPermissions_secrets)
    GetSecretScanningAlerts()(*AppPermissions_secret_scanning_alerts)
    GetSecurityEvents()(*AppPermissions_security_events)
    GetSingleFile()(*AppPermissions_single_file)
    GetStarring()(*AppPermissions_starring)
    GetStatuses()(*AppPermissions_statuses)
    GetTeamDiscussions()(*AppPermissions_team_discussions)
    GetVulnerabilityAlerts()(*AppPermissions_vulnerability_alerts)
    GetWorkflows()(*AppPermissions_workflows)
    SetActions(value *AppPermissions_actions)()
    SetAdministration(value *AppPermissions_administration)()
    SetChecks(value *AppPermissions_checks)()
    SetCodespaces(value *AppPermissions_codespaces)()
    SetContents(value *AppPermissions_contents)()
    SetDependabotSecrets(value *AppPermissions_dependabot_secrets)()
    SetDeployments(value *AppPermissions_deployments)()
    SetEmailAddresses(value *AppPermissions_email_addresses)()
    SetEnvironments(value *AppPermissions_environments)()
    SetFollowers(value *AppPermissions_followers)()
    SetGitSshKeys(value *AppPermissions_git_ssh_keys)()
    SetGpgKeys(value *AppPermissions_gpg_keys)()
    SetInteractionLimits(value *AppPermissions_interaction_limits)()
    SetIssues(value *AppPermissions_issues)()
    SetMembers(value *AppPermissions_members)()
    SetMetadata(value *AppPermissions_metadata)()
    SetOrganizationAdministration(value *AppPermissions_organization_administration)()
    SetOrganizationAnnouncementBanners(value *AppPermissions_organization_announcement_banners)()
    SetOrganizationCopilotSeatManagement(value *AppPermissions_organization_copilot_seat_management)()
    SetOrganizationCustomOrgRoles(value *AppPermissions_organization_custom_org_roles)()
    SetOrganizationCustomProperties(value *AppPermissions_organization_custom_properties)()
    SetOrganizationCustomRoles(value *AppPermissions_organization_custom_roles)()
    SetOrganizationEvents(value *AppPermissions_organization_events)()
    SetOrganizationHooks(value *AppPermissions_organization_hooks)()
    SetOrganizationPackages(value *AppPermissions_organization_packages)()
    SetOrganizationPersonalAccessTokenRequests(value *AppPermissions_organization_personal_access_token_requests)()
    SetOrganizationPersonalAccessTokens(value *AppPermissions_organization_personal_access_tokens)()
    SetOrganizationPlan(value *AppPermissions_organization_plan)()
    SetOrganizationProjects(value *AppPermissions_organization_projects)()
    SetOrganizationSecrets(value *AppPermissions_organization_secrets)()
    SetOrganizationSelfHostedRunners(value *AppPermissions_organization_self_hosted_runners)()
    SetOrganizationUserBlocking(value *AppPermissions_organization_user_blocking)()
    SetPackages(value *AppPermissions_packages)()
    SetPages(value *AppPermissions_pages)()
    SetProfile(value *AppPermissions_profile)()
    SetPullRequests(value *AppPermissions_pull_requests)()
    SetRepositoryCustomProperties(value *AppPermissions_repository_custom_properties)()
    SetRepositoryHooks(value *AppPermissions_repository_hooks)()
    SetRepositoryProjects(value *AppPermissions_repository_projects)()
    SetSecrets(value *AppPermissions_secrets)()
    SetSecretScanningAlerts(value *AppPermissions_secret_scanning_alerts)()
    SetSecurityEvents(value *AppPermissions_security_events)()
    SetSingleFile(value *AppPermissions_single_file)()
    SetStarring(value *AppPermissions_starring)()
    SetStatuses(value *AppPermissions_statuses)()
    SetTeamDiscussions(value *AppPermissions_team_discussions)()
    SetVulnerabilityAlerts(value *AppPermissions_vulnerability_alerts)()
    SetWorkflows(value *AppPermissions_workflows)()
}
