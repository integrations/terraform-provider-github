# TF Resources & Data Sources

This document is intended to provide a high level overview of the Terraform resources and data sources that are maintained in this repository. It is intended to provide a quick reference for maintainers and contributors to understand the alignment of each resource or data source with our coding standards and internal best practices.

We use the symbols ✅, ⚠️, ❌, and ❓ to indicate the alignment status of each resource or data source with our standards. Resources and data sources that have a `DeprecationMessage` set are marked with (🚫) next to their name.

The overall status of each resource or data source is captured in this document along with the status of ongoing work. This status will be decided by the maintainers and is not only determined by the status of the ongoing work.

## Work in Progress

### Schema Functions (Functions)

- Use context based functions for the resource or data source to support context propagation and cancellation.
- Have separate create and update functions for a resource.

### Logging (Logging)

- Use `tflog` for logging instead of `log.Printf` or `fmt.Printf`.

### Repository ID (Repo ID)

- Add a computed `repository_id` attribute to resources and data sources that have a `repository` attribute to support renaming without breaking references.

### Automation Tests (Tests)

- Use `ConfigStateChecks`, `ConfigPlanChecks` and other modern checks instead of the old check pattern.
- Prefer tests with multiple steps covering a single pattern with different scenarios.

### Automation Test Setup (Test Setup)

- Use the client to setup the resource or data source instead of using TF configuration.

### Documentation (Docs)

- Use a refactored template that auto generates documentation.
- Have automated examples that demonstrate the usage of the resource or data source.
- Have import documentation for both import blocks and the CLI.

## Data Sources

| **Data Source** | **Status** | **Functions** | **Logging** | **Tests** | **Test Setup** | **Docs** |
| --- | --- | --- | --- | --- | --- | --- |
| `github_actions_environment_public_key` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_environment_secrets` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_environment_variables` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_oidc_subject_claim_customization_template` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_public_key` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_registration_token` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_secrets` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_variables` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_public_key` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_registration_token` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_repository_oidc_subject_claim_customization_template` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_secrets` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_variables` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_app` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ✅ |
| `github_app_token` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_branch` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_branch_protection_rules` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_organization_public_key` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_organization_secrets` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_public_key` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_secrets` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_user_public_key` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_user_secrets` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_collaborators` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_dependabot_organization_public_key` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_dependabot_organization_secrets` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_dependabot_public_key` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_dependabot_secrets` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_enterprise` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_external_groups` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_ip_ranges` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_issue_labels` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_membership` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_app_installations` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_custom_properties` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_custom_role` (🚫) | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_external_identities` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_ip_allow_list` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_repository_role` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_repository_roles` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_role` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_role_teams` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_role_users` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_roles` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_security_managers` (🚫) | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_team_sync_groups` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_teams` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_webhooks` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_ref` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_release` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_release_asset` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repositories` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_autolink_references` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_branches` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_custom_properties` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_deploy_keys` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_deployment_branch_policies` (🚫) | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_environment_deployment_policies` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_environments` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_file` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_milestone` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_pages` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_pull_request` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_pull_requests` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_teams` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_webhooks` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_rest_api` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_ssh_keys` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_team` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_tree` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_user` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_user_external_identity` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |
| `github_users` | ⚠️ | ✅ | ❓ | ❓ | ❓ | ❓ |

## Resources

| **Resource** | **Status** | **Functions** | **Logging** | **Repo ID** | **Tests** | **Test Setup** | **Docs** |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `github_actions_environment_secret` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_actions_environment_variable` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_actions_hosted_runner` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_oidc_subject_claim_customization_template` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_permissions` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_secret` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_actions_organization_secret_repositories` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_secret_repository` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_variable` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_actions_organization_variable_repositories` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_variable_repository` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_organization_workflow_permissions` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_repository_access_level` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_repository_oidc_subject_claim_customization_template` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_repository_permissions` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_runner_group` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_actions_secret` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_actions_variable` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_app_installation_repositories` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_app_installation_repository` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_branch` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_branch_default` | ⚠️ | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| `github_branch_protection` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_branch_protection_v3` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_organization_secret` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_organization_secret_repositories` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_secret` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_codespaces_user_secret` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_dependabot_organization_secret` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_dependabot_organization_secret_repositories` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_dependabot_organization_secret_repository` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_dependabot_secret` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_emu_group_mapping` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_enterprise_actions_permissions` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_enterprise_actions_runner_group` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_enterprise_actions_workflow_permissions` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_enterprise_app_installation` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_enterprise_ip_allow_list_entry` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_enterprise_organization` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_enterprise_security_analysis_settings` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_issue` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_issue_label` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_issue_labels` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_membership` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_block` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_custom_properties` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_custom_role` (🚫) | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_project` (🚫) | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_repository_role` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_role` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_role_team` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_role_team_assignment` (🚫) | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_role_user` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_ruleset` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_security_manager` (🚫) | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_settings` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_organization_webhook` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_project_card` (🚫) | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_project_column` (🚫) | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_release` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_repository` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_autolink_reference` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_collaborator` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_collaborators` | ⚠️ | ✅ | ✅ | ✅ | ❌ | ❌ | ✅ |
| `github_repository_custom_property` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_repository_dependabot_security_updates` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_deploy_key` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_deployment_branch_policy` (🚫) | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_environment` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_environment_deployment_policy` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_file` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_milestone` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_pages` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_project` (🚫) | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_pull_request` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_ruleset` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_topics` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_vulnerability_alerts` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_repository_webhook` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_team` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_team_members` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| `github_team_membership` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_team_repository` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_team_settings` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_team_sync_group_mapping` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_user_gpg_key` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_user_invitation_accepter` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_user_ssh_key` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
| `github_workflow_repository_permissions` | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ | ❓ |
