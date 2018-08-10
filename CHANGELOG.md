## 1.2.0 (Unreleased)

FEATURES:

* **New Data Source:** `github_repository` [GH-109]
* **New Resource:** `github_organization_project` [GH-111]
* **New Resource:** `github_repository_project` [GH-115]
* **New Resource:** `github_user_ssh_key` [GH-119]

ENHANCEMENTS:

* provider: Add `insecure` mode [GH-48]
* data-source/github_ip_ranges: Add importer IPs [GH-100]
* resource/github_repository: Add support for `topics` [GH-97]
* resource/github_issue_label: Add support for `description` [GH-118]

BUG FIXES:

* resource/github_*: Prevent crashing on invalid ID format [GH-108]

## 1.1.0 (May 11, 2018)

FEATURES:

* **New Data Source:** `github_ip_ranges` ([#82](https://github.com/terraform-providers/terraform-provider-github/issues/82))

ENHANCEMENTS:

* resource/github_repository: Add support for archiving ([#64](https://github.com/terraform-providers/terraform-provider-github/issues/64))
* resource/github_repository: Add `html_url` ([#93](https://github.com/terraform-providers/terraform-provider-github/issues/93))
* resource/github_repository: Add `has_projects` ([#92](https://github.com/terraform-providers/terraform-provider-github/issues/92))
* resource/github_team: Add `parent_team_id` ([#54](https://github.com/terraform-providers/terraform-provider-github/issues/54))

## 1.0.0 (February 20, 2018)

ENHANCEMENTS:

* resource/github_branch_protection: Add support for `require_code_owners_review` ([#51](https://github.com/terraform-providers/terraform-provider-github/issues/51))

## 0.1.2 (February 12, 2018)

BUG FIXES:

* resource/github_membership: Fix a crash when bad import input is given ([#72](https://github.com/terraform-providers/terraform-provider-github/issues/72))

## 0.1.1 (July 18, 2017)

BACKWARDS INCOMPATIBILITIES / NOTES:

* resource/github_branch_protection: The `include_admin` attributes haven't been working for quite some time due to upstream API changes. These attributes are now deprecated in favor of the new top-level `enforce_admins` attribute. The `include_admin` attributes currently have no affect on the resource, and will yield a `DEPRECATED` notice to the user. 

IMPROVEMENTS:

* resource/github_repository: Allow updating default_branch ([#23](https://github.com/terraform-providers/terraform-provider-github/issues/23))
* resource/github_repository: Add license_template and gitignore_template ([#24](https://github.com/terraform-providers/terraform-provider-github/issues/24))
* resource/github_repository_webhook: Add import ([#29](https://github.com/terraform-providers/terraform-provider-github/issues/29))
* resource/github_branch_protection: Support enforce_admins ([#26](https://github.com/terraform-providers/terraform-provider-github/issues/26))
* resource/github_team: Supports managing a team's LDAP DN in GitHub Enterprise ([#39](https://github.com/terraform-providers/terraform-provider-github/issues/39))

BUG FIXES: 

* resource/github_branch_protection: Fix crash on nil values ([#26](https://github.com/terraform-providers/terraform-provider-github/issues/26))

## 0.1.0 (June 20, 2017)

FEATURES:

* **New Resource:** `github_repository_deploy_key` [[#15215](https://github.com/terraform-providers/terraform-provider-github/issues/15215)](https://github.com/hashicorp/terraform/pull/15215)

IMPROVEMENTS:

* Adding merge types to repository resource ([#1](https://github.com/terraform-providers/terraform-provider-github/issues/1))
* added attributes to github_user and github_team data sources ([#2](https://github.com/terraform-providers/terraform-provider-github/issues/2))
