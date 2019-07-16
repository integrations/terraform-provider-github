## 2.2.1 (Unreleased)

ENHANCEMENTS:

* `resource/github_branch_protection`: will now return an error when users are not correctly added [GH-158]

BUG FIXES:

* `resource/github_organization_webhook`: properly set webhook secret in state [GH-251]
* `resource/github_repository_webhook`: properly set webhook secret in state [GH-251]

## 2.2.0 (June 28, 2019)

FEATURES:

* **New Data Source** `github_collaborators` ([#239](https://github.com/terraform-providers/terraform-provider-github/issues/239))

ENHANCEMENTS:

* `provider`: Added optional `individual` attribute, and made `organization` optional ([#242](https://github.com/terraform-providers/terraform-provider-github/issues/242))
* `resource/github_branch_protection`: Added `require_signed_commits` property [[#214](https://github.com/terraform-providers/terraform-provider-github/issues/214)] 

BUG FIXES:

* `resource/github_membership`: `username` property is now case insensitive ([#241](https://github.com/terraform-providers/terraform-provider-github/issues/241))
* `resource/github_repository`: `has_projects` can now be imported ([#237](https://github.com/terraform-providers/terraform-provider-github/issues/237))
* `resource/github_repository_collaborator`: `username` property is now case insensitive [[#241](https://github.com/terraform-providers/terraform-provider-github/issues/241))
* `resource/github_team_membership`: `username` property is now case insensitive ([#241](https://github.com/terraform-providers/terraform-provider-github/issues/241))
 

## 2.1.0 (May 15, 2019)

ENHANCEMENTS:

* `resource/github_repository`: Added validation for lowercase topics ([#223](https://github.com/terraform-providers/terraform-provider-github/issues/223))
* `resource/github_organization_webhook`: Added back removed `name` attribute, `name` is now flagged as `Removed` ([#226](https://github.com/terraform-providers/terraform-provider-github/issues/226))
* `resource/github_repository_webhook`: Added back removed `name` attribute, `name` is now flagged as `Removed` ([#226](https://github.com/terraform-providers/terraform-provider-github/issues/226))

## 2.0.0 (May 02, 2019)

* This release adds support for Terraform 0.12 ([#181](https://github.com/terraform-providers/terraform-provider-github/issues/181))

BREAKING CHANGES:

* `resource/github_repository_webhook`: Removed `name` attribute ([#181](https://github.com/terraform-providers/terraform-provider-github/issues/181))
* `resource/github_organization_webhook`: Removed `name` attribute ([#181](https://github.com/terraform-providers/terraform-provider-github/issues/181))

FEATURES:

* **New Resource:** `github_organization_block` ([#181](https://github.com/terraform-providers/terraform-provider-github/issues/181))
* **New Resource:** `github_user_invitation_accepter` ([#161](https://github.com/terraform-providers/terraform-provider-github/issues/161))
* `resource/github_branch_protection`: Added `required_approving_review_count` property ([#181](https://github.com/terraform-providers/terraform-provider-github/issues/181))

BUG FIXES:

* `resource/github_repository`: Prefill `auto_init` during import ([#154](https://github.com/terraform-providers/terraform-provider-github/issues/154))

## 1.3.0 (September 07, 2018)

FEATURES:

* **New Resource:** `github_project_column` ([#139](https://github.com/terraform-providers/terraform-provider-github/issues/139))

ENHANCEMENTS:

* _all resources_: Use `Etag` to save API quota (~ 33%) ([#143](https://github.com/terraform-providers/terraform-provider-github/issues/143))
* _all resources_: Implement & use RateLimitTransport to avoid hitting API rate limits ([#145](https://github.com/terraform-providers/terraform-provider-github/issues/145))

BUG FIXES:

* `resource/github_issue_label`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/terraform-providers/terraform-provider-github/issues/142))
* `resource/github_membership`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/terraform-providers/terraform-provider-github/issues/142))
* `resource/github_repository_deploy_key`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/terraform-providers/terraform-provider-github/issues/142))
* `resource/github_team`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/terraform-providers/terraform-provider-github/issues/142))
* `resource/github_team_membership`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/terraform-providers/terraform-provider-github/issues/142))
* `resource/github_team_repository`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/terraform-providers/terraform-provider-github/issues/142))
* `resource/github_user_gpg_key`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/terraform-providers/terraform-provider-github/issues/142))

## 1.2.1 (August 17, 2018)

BUG FIXES:

* `resource/github_repository`: Avoid spurious diff for `topics` ([#138](https://github.com/terraform-providers/terraform-provider-github/issues/138))

## 1.2.0 (August 17, 2018)

FEATURES:

* **New Data Source:** `github_repository` ([#109](https://github.com/terraform-providers/terraform-provider-github/issues/109))
* **New Data Source:** `github_repositories` ([#129](https://github.com/terraform-providers/terraform-provider-github/issues/129))
* **New Resource:** `github_organization_project` ([#111](https://github.com/terraform-providers/terraform-provider-github/issues/111))
* **New Resource:** `github_repository_project` ([#115](https://github.com/terraform-providers/terraform-provider-github/issues/115))
* **New Resource:** `github_user_gpg_key` ([#120](https://github.com/terraform-providers/terraform-provider-github/issues/120))
* **New Resource:** `github_user_ssh_key` ([#119](https://github.com/terraform-providers/terraform-provider-github/issues/119))

ENHANCEMENTS:

* `provider`: Add `insecure` mode ([#48](https://github.com/terraform-providers/terraform-provider-github/issues/48))
* `data-source/github_ip_ranges`: Add importer IPs ([#100](https://github.com/terraform-providers/terraform-provider-github/issues/100))
* `resource/github_issue_label`: Add support for `description` ([#118](https://github.com/terraform-providers/terraform-provider-github/issues/118))
* `resource/github_repository`: Add support for `topics` ([#97](https://github.com/terraform-providers/terraform-provider-github/issues/97))
* `resource/github_team`: Expose `slug` ([#136](https://github.com/terraform-providers/terraform-provider-github/issues/136))
* `resource/github_team_membership`: Make role updatable ([#137](https://github.com/terraform-providers/terraform-provider-github/issues/137))

BUG FIXES:

* `resource/github_*`: Prevent crashing on invalid ID format ([#108](https://github.com/terraform-providers/terraform-provider-github/issues/108))
* `resource/github_organization_webhook`: Avoid spurious diff of `secret` ([#134](https://github.com/terraform-providers/terraform-provider-github/issues/134))
* `resource/github_repository`: Make non-updatable fields `ForceNew` ([#135](https://github.com/terraform-providers/terraform-provider-github/issues/135))
* `resource/github_repository_deploy_key`: Avoid spurious diff of `key` ([#132](https://github.com/terraform-providers/terraform-provider-github/issues/132))
* `resource/github_repository_webhook`: Avoid spurious diff of `secret` ([#133](https://github.com/terraform-providers/terraform-provider-github/issues/133))


## 1.1.0 (May 11, 2018)

FEATURES:

* **New Data Source:** `github_ip_ranges` ([#82](https://github.com/terraform-providers/terraform-provider-github/issues/82))

ENHANCEMENTS:

* `resource/github_repository`: Add support for archiving ([#64](https://github.com/terraform-providers/terraform-provider-github/issues/64))
* `resource/github_repository`: Add `html_url` ([#93](https://github.com/terraform-providers/terraform-provider-github/issues/93))
* `resource/github_repository`: Add `has_projects` ([#92](https://github.com/terraform-providers/terraform-provider-github/issues/92))
* `resource/github_team`: Add `parent_team_id` ([#54](https://github.com/terraform-providers/terraform-provider-github/issues/54))

## 1.0.0 (February 20, 2018)

ENHANCEMENTS:

* `resource/github_branch_protection`: Add support for `require_code_owners_review` ([#51](https://github.com/terraform-providers/terraform-provider-github/issues/51))

## 0.1.2 (February 12, 2018)

BUG FIXES:

* `resource/github_membership`: Fix a crash when bad import input is given ([#72](https://github.com/terraform-providers/terraform-provider-github/issues/72))

## 0.1.1 (July 18, 2017)

BACKWARDS INCOMPATIBILITIES / NOTES:

* `resource/github_branch_protection`: The `include_admin` attributes haven't been working for quite some time due to upstream API changes. These attributes are now deprecated in favor of the new top-level `enforce_admins` attribute. The `include_admin` attributes currently have no affect on the resource, and will yield a `DEPRECATED` notice to the user. 

IMPROVEMENTS:

* `resource/github_repository`: Allow updating default_branch ([#23](https://github.com/terraform-providers/terraform-provider-github/issues/23))
* `resource/github_repository`: Add license_template and gitignore_template ([#24](https://github.com/terraform-providers/terraform-provider-github/issues/24))
* `resource/github_repository_webhook`: Add import ([#29](https://github.com/terraform-providers/terraform-provider-github/issues/29))
* `resource/github_branch_protection`: Support enforce_admins ([#26](https://github.com/terraform-providers/terraform-provider-github/issues/26))
* `resource/github_team`: Supports managing a team's LDAP DN in GitHub Enterprise ([#39](https://github.com/terraform-providers/terraform-provider-github/issues/39))

BUG FIXES: 

* `resource/github_branch_protection`: Fix crash on nil values ([#26](https://github.com/terraform-providers/terraform-provider-github/issues/26))

## 0.1.0 (June 20, 2017)

FEATURES:

* **New Resource:** `github_repository_deploy_key` [[#15215](https://github.com/terraform-providers/terraform-provider-github/issues/15215)](https://github.com/hashicorp/terraform/pull/15215)

IMPROVEMENTS:

* `resource/github_repository`: Adding merge types ([#1](https://github.com/terraform-providers/terraform-provider-github/issues/1))
* `data-source/github_user` and `data-source/github_team`: Added attributes ([#2](https://github.com/terraform-providers/terraform-provider-github/issues/2))
