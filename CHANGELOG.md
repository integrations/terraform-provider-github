## (Unreleased)

## 3.0.0 (September 08, 2020)

BREAKING CHANGES:

- `token` becomes optional
- `organization` no longer deprecated
- `individual` and `anonymous` removed
- `owner` inferred from `organization`

BUG FIXES:

- `terraform validate` fails because of missing token (GH-503)
- organization support for various resources (GH-501)

ENHANCEMENTS:

* **New Data Source** `github_organization` (GH-521)



## 2.9.2 (July 14, 2020)

- Adds deprecation of `anonymous` flag for provider configuration ahead of next major release ([#506](https://github.com/terraform-providers/terraform-provider-github/issues/506))
- Adds deprecation of `individual` flag for provider configuration ahead of next major release ([#512](https://github.com/terraform-providers/terraform-provider-github/issues/512))

## 2.9.1 (July 01, 2020)

BUG FIXES:

- Reverts changes introduced in v2.9.0, deferring to the next major release

## 2.9.0 (June 29, 2020)

**NOTE**: This release introduced a provider-level breaking change around `anonymous` use.
See [here](https://github.com/terraform-providers/terraform-provider-github/pull/464#discussion_r427961161) for details and [here](https://github.com/terraform-providers/terraform-provider-github/issues/502#issuecomment-652379322) to discuss a fix.

ENHANCEMENTS:
* Add Ability To Manage Resources For Non-Organization Accounts ([#464](https://github.com/terraform-providers/terraform-provider-github/issues/464))
* resource/github_repository: Add "internal" Visibility Option ([#454](https://github.com/terraform-providers/terraform-provider-github/issues/454))

## 2.8.1 (June 09, 2020)

BUG FIXES:

* resource/github_repository_file: Reduce API requests when looping through commits ([[#466](https://github.com/terraform-providers/terraform-provider-github/issues/466)])
* resource/github_repository: Fix `auto_init` Destroying Repositories ([[#317](https://github.com/terraform-providers/terraform-provider-github/issues/317)])
* resource/github_repository_deploy_key: Fix acceptance test approach ([[#471](https://github.com/terraform-providers/terraform-provider-github/issues/471)])
* resource/github_actions_secret: Fix Case Where Secret Removed Outside Of Terraform ([[#482](https://github.com/terraform-providers/terraform-provider-github/issues/482)])
* Documentation Addition Of `examples/` Directory

## 2.8.0 (May 15, 2020)

BUG FIXES:

* resource/github_branch_protection: Prevent enabling `dismissal_restrictions` in Github console if `dismissal_users` and `dismissal_teams` are not set ([#453](https://github.com/terraform-providers/terraform-provider-github/issues/453))
* resource/github_repository_collaborator: Allow modifying permissions from `maintain` and `triage`  ([#457](https://github.com/terraform-providers/terraform-provider-github/issues/457))
* Documentation Fix for `github_actions_public_key` data-source ([#458](https://github.com/terraform-providers/terraform-provider-github/issues/458))
* Documentation Fix for `github_branch_protection` resource ([#410](https://github.com/terraform-providers/terraform-provider-github/issues/410))
* Documentation Layout Fix for `github_ip_ranges` and `github_membership` data sources ([#423](https://github.com/terraform-providers/terraform-provider-github/issues/423))
* Documentation Fix for `github_repository_file` import ([#443](https://github.com/terraform-providers/terraform-provider-github/issues/443))
* Update `go-github` to `v31.0.0` ([#424](https://github.com/terraform-providers/terraform-provider-github/issues/424))

ENHANCEMENTS:
* **New Data Source** `github_organization_team_sync_groups` ([#400](https://github.com/terraform-providers/terraform-provider-github/issues/400))
* **New Resource** `github_team_sync_group_mapping` ([#400](https://github.com/terraform-providers/terraform-provider-github/issues/400))

## 2.7.0 (May 01, 2020)

BUG FIXES:

* Add Missing Acceptance Test ([#427](https://github.com/terraform-providers/terraform-provider-github/issues/427))

ENHANCEMENTS:

* Add GraphQL Client ([#331](https://github.com/terraform-providers/terraform-provider-github/issues/331))
* **New Data Source** `github_branch` ([#364](https://github.com/terraform-providers/terraform-provider-github/issues/364))
* **New Resource** `github_branch` ([#364](https://github.com/terraform-providers/terraform-provider-github/issues/364))


## 2.6.1 (April 07, 2020)

BUG FIXES:

* Documentation Fix For Option To Manage `Delete Branch on Merge` ([#408](https://github.com/terraform-providers/terraform-provider-github/issues/408))
* Documentation Fix For `github_actions_secret` / `github_actions_public_key` ([#413](https://github.com/terraform-providers/terraform-provider-github/issues/413))

## 2.6.0 (April 03, 2020)

ENHANCEMENTS:

* resource/github_repository: Introduce Option To Manage `Delete Branch on Merge` ([#399](https://github.com/terraform-providers/terraform-provider-github/issues/399))
* resource/github_repository: Configure Repository As Template ([#357](https://github.com/terraform-providers/terraform-provider-github/issues/357))
* **New Data Source** `github_membership` ([#396](https://github.com/terraform-providers/terraform-provider-github/issues/396))

## 2.5.1 (April 02, 2020)

BUG FIXES:

* Fix Broken Link For `github_actions_secret` Documentation ([#405](https://github.com/terraform-providers/terraform-provider-github/issues/405))

## 2.5.0 (March 30, 2020)

ENHANCEMENTS:

* Add `apps` To `github_branch_protection` Restrictions
* **New Data Source** `github_actions_public_key` ([[#362](https://github.com/terraform-providers/terraform-provider-github/issues/362)])
* **New Data Source** `github_actions_secrets` ([[#362](https://github.com/terraform-providers/terraform-provider-github/issues/362)])
* **New Resource:** `github_actions_secret` ([[#362](https://github.com/terraform-providers/terraform-provider-github/issues/362)])

BUG FIXES:

* Prevent Panic From DismissalRestrictions ([[#385](https://github.com/terraform-providers/terraform-provider-github/issues/385)])
* Update `go-github` to `v29.0.3` ([[#369](https://github.com/terraform-providers/terraform-provider-github/issues/369)])
* Update `go` to `1.13` ([[#372](https://github.com/terraform-providers/terraform-provider-github/issues/372)])
* Documentation Fixes For Consistency And Typography


## 2.4.1 (March 05, 2020)

BUG FIXES:

* Updates `go-github` to `v29` to unblock planned feature development ([#342](https://github.com/terraform-providers/terraform-provider-github/issues/342))
* Fixes `insecure_ssl` parameter behaviour for `github_organization_webhook` and  `github_repository_webhook` ([#365](https://github.com/terraform-providers/terraform-provider-github/issues/365))
* Fixes label behaviour to not create new labels when renaming a `github_issue_label` ([#288](https://github.com/terraform-providers/terraform-provider-github/issues/288))

## 2.4.0 (February 26, 2020)

ENHANCEMENTS:

* **New Data Source** `github_release` ([#356](https://github.com/terraform-providers/terraform-provider-github/pull/356))

* **New Resource:** `github_repository_file` ([#318](https://github.com/terraform-providers/terraform-provider-github/pull/318))

## 2.3.2 (February 05, 2020)

BUG FIXES:

* Handle repository 404 for `github_repository_collaborator` resource ([#348](https://github.com/terraform-providers/terraform-provider-github/issues/348))

## 2.3.1 (January 27, 2020)

BUG FIXES:

* Add support for `triage` and `maintain` permissions in some resources ([#303](https://github.com/terraform-providers/terraform-provider-github/issues/303))

## 2.3.0 (January 22, 2020)

BUG FIXES:

* `resource/resource_github_team_membership`: Prevent spurious diffs caused by upstream API change made on 17th January ([#325](https://github.com/terraform-providers/terraform-provider-github/issues/325))

ENHANCEMENTS:

* `resource/repository`: Added functionality to generate a new repository from a Template Repository ([#309](https://github.com/terraform-providers/terraform-provider-github/issues/309))

## 2.2.1 (September 04, 2019)

ENHANCEMENTS:

* dependencies: Updated module `hashicorp/terraform` to `v0.12.7` ([#273](https://github.com/terraform-providers/terraform-provider-github/issues/273))
* `resource/github_branch_protection`: Will now return an error when users are not correctly added ([#158](https://github.com/terraform-providers/terraform-provider-github/issues/158))
* `provider`: Added optional `anonymous` attribute, and made `token` optional ([#255](https://github.com/terraform-providers/terraform-provider-github/issues/255))

BUG FIXES:
* `resource/github_repository`: Allow setting `default_branch` to `master` on creation ([#150](https://github.com/terraform-providers/terraform-provider-github/issues/150))
* `resource/github_team_repository`: Validation of `team_id` ([#253](https://github.com/terraform-providers/terraform-provider-github/issues/253))
* `resource/github_team_membership`: Validation of `team_id` ([#253](https://github.com/terraform-providers/terraform-provider-github/issues/253))
* `resource/github_organization_webhook`: Properly set webhook secret in state ([#251](https://github.com/terraform-providers/terraform-provider-github/issues/251))
* `resource/github_repository_webhook`: Properly set webhook secret in state ([#251](https://github.com/terraform-providers/terraform-provider-github/issues/251))

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
