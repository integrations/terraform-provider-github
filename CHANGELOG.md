# NOTE: CHANGELOG.md is deprecated

After the release of v4.24.0, please see the [GitHub release notes](https://github.com/integrations/terraform-provider-github/releases) for the provider in order to view the most up-to-date changes.

# 4.24.0 (Apr 28, 2022)

ENHANCEMENTS:
* Support for allow_forking on a repository/update to go-github v42 by @diogopms in https://github.com/integrations/terraform-provider-github/pull/1033
* Upgrade go-github to v43.0.0 by @btkostner in https://github.com/integrations/terraform-provider-github/pull/1087

BUG FIXES:

* Fix go module path by @turkenh in https://github.com/integrations/terraform-provider-github/pull/961
* fix: remove incorrect required schema key on ref data source by @youcandanch in https://github.com/integrations/terraform-provider-github/pull/1109
* Bump Go version for Actions release CI to 1.18 by @kfcampbell in https://github.com/integrations/terraform-provider-github/pull/1134
* build(deps): bump actions/setup-go from 2 to 3 by @dependabot in https://github.com/integrations/terraform-provider-github/pull/1110
* Fix linting issues by @kfcampbell in https://github.com/integrations/terraform-provider-github/pull/1107


# 4.23.0 (Mar 25, 2022)

ENHANCEMENTS:

* Add support for disabling the use of the vulnerability management endpoint by @enieuw in https://github.com/integrations/terraform-provider-github/pull/1022
* Added orgname in github_orgranization attributes by @Kavinraja-G in https://github.com/integrations/terraform-provider-github/pull/1052
* Add a data source for refs. by @youcandanch in https://github.com/integrations/terraform-provider-github/pull/1084
* build(deps): bump actions/checkout from 2 to 3 by @dependabot in https://github.com/integrations/terraform-provider-github/pull/1086

BUG FIXES:

* fix: use pagination to fetch all team members by @carocad in https://github.com/integrations/terraform-provider-github/pull/1092
* docs: fix typos in d/users.html.markdown by @pallxk in https://github.com/integrations/terraform-provider-github/pull/1049

# 4.22.0 (Mar 18, 2022)

ENHANCEMENTS:

* feat: add `tree` data source by @jasonwalsh in https://github.com/integrations/terraform-provider-github/pull/1027
* feat: support for issues using github_issue resource by @ewilde in https://github.com/integrations/terraform-provider-github/pull/1047
* feat: add configurable read_delay_ms by @morremeyer in https://github.com/integrations/terraform-provider-github/pull/1054

## 4.21.0 (Mar 11, 2022)

ENHANCEMENTS:

* Adding BypassPullRequestActorIDs to branch protection by @jtyr in https://github.com/integrations/terraform-provider-github/pull/1030
* Adding suspended_at attribute to github_user data source by @mrobinson-anaplan in https://github.com/integrations/terraform-provider-github/pull/1070
* Documentation: Add id to github_user data dource by @kangaechu in https://github.com/integrations/terraform-provider-github/pull/1061

BUG FIXES:

* fix: use the appropriate ID when trying to import `github_team_members` objects by @bison-brandon in https://github.com/integrations/terraform-provider-github/pull/1074
* Environment ID gets set incorrectly on update by @aceresia-bg in https://github.com/integrations/terraform-provider-github/pull/1058
* Fix whitespace in documentation for branch_protection_v3 by @JCradock in https://github.com/integrations/terraform-provider-github/pull/1059

## 4.20.1 (Mar 3, 2022)

BUG FIXES:

* Remove team from state if deletion failed and it does not exist by @cytopia in https://github.com/integrations/terraform-provider-github/pull/1039
    * Note that this is a behavior change from previous GitHub Terraform provider releases: now, if a GitHub team deletion operation fails and the team does not exist, the team will be automatically removed from state.
* Make data_github_repository work with non-existing repositories by @tobiassjosten in https://github.com/integrations/terraform-provider-github/pull/1031
* Standardize logs by @kfcampbell in https://github.com/integrations/terraform-provider-github/pull/1053

## 4.20.0 (Feb 3, 2022)

ENHANCEMENTS:

* Add new resource `github_team_members` to allow authoritative team management by @stawik-mesa in https://github.com/integrations/terraform-provider-github/pull/975

BUG FIXES:

* test: checkout pull request via sha instead of ref by @jcudit in https://github.com/integrations/terraform-provider-github/pull/1043
* Small CI cleanup by @kfcampbell in https://github.com/integrations/terraform-provider-github/pull/1048

**Full Changelog**: https://github.com/integrations/terraform-provider-github/compare/v4.19.2...v4.20.0


## 4.19.2 (Jan 20, 2022)

BUG FIXES:

- Update `go-github` to v42.0.0 ([#1035](https://github.com/integrations/terraform-provider-github/pull/1035))
- Adjust count requirement of `required_approving_review_count` option for `github_branch_protection` ([#971](https://github.com/integrations/terraform-provider-github/pull/971))
- Add `nil` check for `require_conversation_resolution` field of `github_branch_protection` resource ([#1032](https://github.com/integrations/terraform-provider-github/pull/1032))

## 4.19.1 (Jan 5, 2022)

BUG FIXES:

- Update `go-github` to v41.0.0 ([#993](https://github.com/integrations/terraform-provider-github/pull/993))
- Add `nil` check for `plan` field of `github_organization` data source ([#1016](https://github.com/integrations/terraform-provider-github/pull/1016))


## 4.19.0 (Dec 13, 2021)

ENHANCEMENTS:

- Export `branches` attribute of `github_repository` resource ([[#959](https://github.com/integrations/terraform-provider-github/pull/959)])
- Add `require_conversation_resolution` support for `github_branch_protection` resource ([[#904](https://github.com/integrations/terraform-provider-github/pull/904)])

BUG FIXES:

- Adjust length requirement to `topics` option for `github_repository` ([[#996](https://github.com/integrations/terraform-provider-github/pull/996)])
- Add `required_linear_history` support for `github_branch_protection` resource ([[#935](https://github.com/integrations/terraform-provider-github/pull/935)])


## 4.18.2 (Nov 30, 2021)

BUG FIXES:

- Add length requirement to `name` option for `github_repository` ([[#965](https://github.com/integrations/terraform-provider-github/pull/965)])
- Various documentation fixes 🙇

## 4.18.1 (Nov 22, 2021)

BUG FIXES:

- Add length requirement to `topics` option for `github_repository` ([[#951](https://github.com/integrations/terraform-provider-github/pull/951)])
- Add pagination to `selected_repositories` option for `github_actions_runner_group` ([[#970](https://github.com/integrations/terraform-provider-github/pull/970)])
- Add handling for new `node_id` format introduced to the GitHub GraphQL API (`github_repository`) ([[#914](https://github.com/integrations/terraform-provider-github/pull/914)])

## 4.18.0 (Nov 8, 2021)

ENHANCEMENTS:

- **New Resource:** `github_actions_organization_permissions` ([[#920](https://github.com/integrations/terraform-provider-github/pull/920)])

BUG FIXES:

- Add newline compatbility to GitHub App provider authentication ([[#931](https://github.com/integrations/terraform-provider-github/pull/931)])
- Fix `strict` setting of `required_status_checks` for `github_branch_protection` resource ([[#880](https://github.com/integrations/terraform-provider-github/issues/880)])


## 4.17.0 (Oct 17, 2021)

ENHANCEMENTS:

- **New Resource:** `github_repository_autolink_reference` ([[#924](https://github.com/integrations/terraform-provider-github/pull/924)])
- **New Data Sources** `github_users` ([#900](https://github.com/integrations/terraform-provider-github/pull/900))
- Add `allow_auto_merge` option for `github_repository` ([#923](https://github.com/integrations/terraform-provider-github/pull/923))

BUG FIXES:

- Various documentation fixes 🙇

## 4.16.0 (Oct 5, 2021)

ENHANCEMENTS:

* **New Data Source:** `github_repository_file` ([#896](https://github.com/integrations/terraform-provider-github/pull/896))
- Add `write_delay_ms` provider option [#907](https://github.com/integrations/terraform-provider-github/pull/907))

BUG FIXES:

- Update `go-github` to v39.0.0 ([#905](https://github.com/integrations/terraform-provider-github/pull/905))

## 4.15.1 (Sep 23, 2021)

BUG FIXES:

- Revert suppression of  `etag` changes for `github_repository` resources ([[#910](https://github.com/integrations/terraform-provider-github/issues/910)])

## 4.15.0 (Sep 22, 2021)

ENHANCEMENTS:

- **New Resource:** `github_actions_organization_secret_repositories` ([[#882](https://github.com/integrations/terraform-provider-github/issues/882)])
- **New Resource:** `github_actions_runner_group` ([[#821](https://github.com/integrations/terraform-provider-github/issues/821)])
- Add `require_linear_history` to `github_branch_protection` resource ([[#887](https://github.com/integrations/terraform-provider-github/issues/887)])
- Add `branches` attribute to `github_repository` resource ([[#892](https://github.com/integrations/terraform-provider-github/issues/892)])


BUG FIXES:

- Update documentation for `d/github_ip_ranges` ([#895](https://github.com/integrations/terraform-provider-github/issues/895))
- Update `go-github` to v38 ([#901](https://github.com/integrations/terraform-provider-github/issues/901))
- Suppress `etag` changes for `github_repository` resources ([[#909](https://github.com/integrations/terraform-provider-github/issues/909)])


## 4.14.0 (Sep 2, 2021)

BUG FIXES:

- Adds support for recreating a `github_team_repository` when repository is renamed ([#870](https://github.com/integrations/terraform-provider-github/issues/870))
- Adds logging of configured authentication on provider startup ([#867](https://github.com/integrations/terraform-provider-github/issues/867))
- Update documentation for `github_ip_ranges` data source ([#857](https://github.com/integrations/terraform-provider-github/issues/857))
- Add support for IPv6 addresses returned by `github_ip_ranges` data source ([#883](https://github.com/integrations/terraform-provider-github/issues/883))
- Update `go-github` to v37.0.0 ([#893](https://github.com/integrations/terraform-provider-github/issues/893))

## 4.13.0 (Jul 26, 2021)

BUG FIXES:

- Fix setting `vulnerability_alerts` on private `github_repository` creation ([#768](https://github.com/integrations/terraform-provider-github/issues/768))

ENHANCEMENTS:

- Add `restrict_dismissals` option to `github_branch_protection` resource ([#839](https://github.com/integrations/terraform-provider-github/issues/839))

## 4.12.2 (Jul 12, 2021)

BUG FIXES:

- Update `go-github` to v36.0.0 ([#841](https://github.com/integrations/terraform-provider-github/issues/841))

## 4.12.0 (Jun 18, 2021)

ENHANCEMENTS:

* **New Resource:** `github_actions_environment_secret` ([[#805](https://github.com/integrations/terraform-provider-github/issues/805)])
* **New Resource:** `github_repository_environment` ([[#805](https://github.com/integrations/terraform-provider-github/issues/805)])
* Add `members` field to `github_organization` data source ([[#811](https://github.com/integrations/terraform-provider-github/issues/811)])
* Add `repositories` field to `github_team` data source ([[#791](https://github.com/integrations/terraform-provider-github/issues/791)])
* Add `repositories` field to `github_organization_teams` data source ([[#791](https://github.com/integrations/terraform-provider-github/issues/791)])


BUG FIXES:

- Document incompatibility between `github_app_installation_repository` and GitHub App authentication ([#818](https://github.com/integrations/terraform-provider-github/issues/818))
- Document migration from `hashicorp/terraform-provider-github ([#816](https://github.com/integrations/terraform-provider-github/issues/816))
- Allow users and apps to also be applied to push restrictions for `github_branch_protection` ([#824](https://github.com/integrations/terraform-provider-github/issues/824))


## 4.11.0 (Jun 7, 2021)

BREAKING CHANGES:

- Allow PEM data to be passed directly for GitHub App provider authentication ([#803](https://github.com/integrations/terraform-provider-github/issues/803))

ENHANCEMENTS:

- Add `encrypted_value` field to `github_actions_secret` and `github_actions_organization_secret` resources ([#807](https://github.com/integrations/terraform-provider-github/issues/807))

BUG FIXES:

- Fix error handling when branch does not exist for `github_branch` resource ([#806](https://github.com/integrations/terraform-provider-github/issues/806))

## 4.10.1 (May 25, 2021)

BUG FIXES:

* Improve documentation for provider authentication options ([#801](https://github.com/integrations/terraform-provider-github/issues/801))


## 4.10.0 (May 21, 2021)

ENHANCEMENTS:

* Add GitHub App authentication option to provider ([#753](https://github.com/integrations/terraform-provider-github/issues/753))
* Add Actions and Dependabot IP ranges to `github_ip_ranges` data source ([#784](https://github.com/integrations/terraform-provider-github/issues/784))
* Add additional fields to `github_repository` data source ([#778](https://github.com/integrations/terraform-provider-github/issues/778))

## 4.9.4 (May 11, 2021)

BUG FIXES:

- Add EMU support by allowing `internal` visibility to be configured for `github_repository` ([#781](https://github.com/integrations/terraform-provider-github/issues/781))
- Update `go-github` to 35.1.0 ([#772](https://github.com/integrations/terraform-provider-github/issues/772))

## 4.9.3 (May 7, 2021)

BUG FIXES:

- Mark `slug` as `computed` when `name` is changed for `github_team` ([#757](https://github.com/integrations/terraform-provider-github/issues/757))

## 4.9.2 (April 18, 2021)

BUG FIXES:

- correct visibility for repositories created via a template ([#761](https://github.com/integrations/terraform-provider-github/issues/761))


## 4.9.1 (April 17, 2021)

BUG FIXES:

- Bump Go version to 1.16 for acceptance tests and darwin/arm64 releases
- Update acceptance tests to v2.2.0
- Re-instate releases of darwin/arm64

## 4.9.0 (April 17, 2021)

ENHANCEMENTS:

* **New Data Sources** `github_repository_pull_request` / `github_repository_pull_requests` ([#739](https://github.com/integrations/terraform-provider-github/issues/739))
* **New Resource** `github_repository_pull_request` ([#739](https://github.com/integrations/terraform-provider-github/issues/739))
* Add `repositories` attribute for `github_organization` data source ([#750](https://github.com/integrations/terraform-provider-github/issues/750))
* Add import functionality for `github_actions_organization_secret` ([#745](https://github.com/integrations/terraform-provider-github/issues/745))

BUG FIXES:

- Detect and overwrite value drift for `github_actions_secret` and `github_actions_organization_secret` ([#740](https://github.com/integrations/terraform-provider-github/pull/740))
- Do not destroy repositories when `name` attribute changes ([#699](https://github.com/integrations/terraform-provider-github/pull/699))

## 4.8.0 (April 9, 2021)

BUG FIXES:

-  Deprecate `organization` / `GITHUB_ORGANIZATION` provider configuration options ([#735](https://github.com/integrations/terraform-provider-github/pull/735))

## 4.7.0 (April 9, 2021)

REGRESSIONS:

- new repositories created via a template have a public visibility ([#758](https://github.com/integrations/terraform-provider-github/issues/758))
  - workaround: a subsequent plan / apply will set the visibility to what is configured
  - fix: see v4.9.2

ENHANCEMENTS:

* **New Data Source** `github_organization_teams` ([#725](https://github.com/integrations/terraform-provider-github/issues/725))

BUG FIXES:

- Set visibility on create instead of update for `github_repository` ([#746](https://github.com/integrations/terraform-provider-github/pull/746))
- Various documentation updates

## 4.6.0 (March 23, 2021)

ENHANCEMENTS:

* **New Resource** `github_app_installation_repository` ([#690](https://github.com/integrations/terraform-provider-github/issues/690))

BUG FIXES:

- Fix panic for `github_repository_file` ([#732](https://github.com/integrations/terraform-provider-github/pull/732))
- Improve error messaging for `github_branch` ([#734](https://github.com/integrations/terraform-provider-github/pull/734))
- Improve error messaging for `github_branch_protection` ([#721](https://github.com/integrations/terraform-provider-github/pull/721))
- Fix update operation for `github_default_branch` ([#719](https://github.com/integrations/terraform-provider-github/pull/719))
- Add name validation for `github_actions_organization_secret` ([#714](https://github.com/integrations/terraform-provider-github/pull/714))


## 4.5.2 (March 16, 2021)

BUG FIXES:

- Fix updating `default_branch` on `github_repository` ([#719](https://github.com/integrations/terraform-provider-github/pull/719))


## 4.5.1 (March 3, 2021)

BUG FIXES:

- Fix `github_branch_protection` import by repository node ID and pattern ([#713](https://github.com/integrations/terraform-provider-github/pull/713))
- Add pagination when retrieving team members for `data_source_github_team` ([#702](https://github.com/integrations/terraform-provider-github/pull/702))


## 4.5.0 (February 17, 2021)

ENHANCEMENTS:

- Add ability for `github_team_repository` to accept slug as a valid `team_id` ([#693](https://github.com/integrations/terraform-provider-github/pull/693))

BUG FIXES:

- Add more context to error messaging for `github_branch_protection` ([#691](https://github.com/integrations/terraform-provider-github/pull/691))
- Satisfy linter recommendation for `github_branch_protection` ([#694](https://github.com/integrations/terraform-provider-github/pull/694))

## 4.4.0 (February 5, 2021)

BUG FIXES:

- Add `create_default_maintainer` option to `github_team` ([#527](https://github.com/integrations/terraform-provider-github/pull/527)), ([#104](https://github.com/integrations/terraform-provider-github/pull/104)), ([#130](https://github.com/integrations/terraform-provider-github/pull/130))
- Add diff-suppression option to `repository_collaborator` ([#683](https://github.com/integrations/terraform-provider-github/pull/683))


## 4.3.2 (February 2, 2021)

BUG FIXES:

* Improved detection of repository name for `github_branch_protection` ([#684](https://github.com/integrations/terraform-provider-github/issues/684))
* Reverts error handling in provider configuration ([#685](https://github.com/integrations/terraform-provider-github/issues/685))

## 4.3.1 (January 22, 2021)

REGRESSIONS:

- provider configuration breaks for individual accounts ([#678](https://github.com/integrations/terraform-provider-github/issues/678))

BUG FIXES:

* Send valid payload when editing a repository resource with `github_branch_default` ([#666](https://github.com/integrations/terraform-provider-github/issues/666))
* Add handling to surface errors in provider configuration ([#668](https://github.com/integrations/terraform-provider-github/issues/668))

## 4.3.0 (January 14, 2021)

ENHANCEMENTS:

* **New Resource** `github_branch_protection_v3` ([#642](https://github.com/integrations/terraform-provider-github/issues/642))
* Add `pages` feature to `github_repository` ([#490](https://github.com/integrations/terraform-provider-github/issues/490))


## 4.2.0 (January 07, 2021)

BREAKING CHANGES:

- Project transfer from `terraform-providers` organization to `integrations`
    - See [#652](https://github.com/integrations/terraform-provider-github/issues/652) for migration steps and [#656](https://github.com/integrations/terraform-provider-github/issues/656) for discussion

ENHANCEMENTS:

- Add `allowDeletions` and `allowsForcePushes` to `github_branch_protection` ([#623](https://github.com/integrations/terraform-provider-github/pull/623))
- Add GitHub App actor support to `github_branch_protection` ([#615](https://github.com/integrations/terraform-provider-github/pull/615))

BUG FIXES:

- Allow `required_status_checks` `strict` to be `false` for `github_branch_protection` ([#614](https://github.com/integrations/terraform-provider-github/pull/614))
- Remove `ForceNew` on template-related options for `github_repository` ([#609](https://github.com/integrations/terraform-provider-github/pull/609))
- Fix parsing of input during imports of `github_branch_protection` ([#610](https://github.com/integrations/terraform-provider-github/pull/610))
- `github_repository_file` resource no longer iterates through all commits ([#644](https://github.com/integrations/terraform-provider-github/pull/644))

## 4.1.0 (December 01, 2020)

ENHANCEMENTS:

- Add `github_actions_organization_secret` resource ([#261](https://github.com/integrations/terraform-provider-github/pull/261))
- Add `github_repository_milestone` resource ([#470](https://github.com/integrations/terraform-provider-github/pull/470))
- Add `github_repository_milestone` data source ([#470](https://github.com/integrations/terraform-provider-github/pull/470))
- Add `github_project_card` resource ([#460](https://github.com/integrations/terraform-provider-github/pull/460))
- Add `github_branch_default` resource ([#194](https://github.com/integrations/terraform-provider-github/pull/194))


## 4.0.1 (November 18, 2020)

BUG FIXES:

- `github_team` data source query no longer iterates through a list of teams ([#579](https://github.com/integrations/terraform-provider-github/pull/579))
- `github_repository_file` resource no longer iterates through all commits ([#589](https://github.com/integrations/terraform-provider-github/pull/589))
- fix parsing of `repo:pattern` format during `github_branch_protection` import ([#599](https://github.com/integrations/terraform-provider-github/pull/599))


## 4.0.0 (November 10, 2020)

REGRESSIONS:

- fails parsing of `repo:pattern` format during `github_branch_protection` import ([#597](https://github.com/integrations/terraform-provider-github/issues/597))

BREAKING CHANGES:

- `pattern` replaces `branch` in changes to `github_branch_protection` introduced in `v3.1.0` ([#566](https://github.com/integrations/terraform-provider-github/issues/566))
- `dismissal_restrictions` documented for `github_branch_protection` ([#569](https://github.com/integrations/terraform-provider-github/pull/569))
- project license changes from MPL2 to MIT ([#591](https://github.com/integrations/terraform-provider-github/pull/591))

BUG FIXES:

- `repository_id` for `github_branch_protection` accepts repository name as well as node ID ([#593](https://github.com/integrations/terraform-provider-github/issues/593))

ENHANCEMENTS:

- Add support to get currently authenticated user to `data_source_github_user` ([#261](https://github.com/integrations/terraform-provider-github/pull/261))
- Add importing to `github_organization_webhook` ([#487](https://github.com/integrations/terraform-provider-github/pull/487))


## 3.1.0 (October 12, 2020)

REGRESSIONS:

- undocumented, breaking configuration changes to `github_branch_protection` ([#566](https://github.com/integrations/terraform-provider-github/issues/566))
- slowed performance of `github_branch_protection` ([#567](https://github.com/integrations/terraform-provider-github/issues/567))
- change to default branch breaks refresh of existing `github_repository_file` resources ([#564](https://github.com/integrations/terraform-provider-github/issues/564))

BREAKING CHANGES:

- Deprecate `anonymous` Flag For Provider Configuration ([#506](https://github.com/integrations/terraform-provider-github/issues/506))

BUG FIXES:

- re-instante resources unavailable in the context of an organization ([#501](https://github.com/integrations/terraform-provider-github/issues/501))
- allow overwrite-on-create behaviour for `github_repository_file` resource ([#459](https://github.com/integrations/terraform-provider-github/issues/459))


ENHANCEMENTS:

- update `go-github` to `v32.1.0` ([#475](https://github.com/integrations/terraform-provider-github/issues/475))
- add `vulnerability_alerts` to `github_repository` ([#444](https://github.com/integrations/terraform-provider-github/issues/444))
- add `archive_on_destroy` to `github_repository` ([#432](https://github.com/integrations/terraform-provider-github/issues/432))
- uplift `branch_protection` to GraphQL ([#337](https://github.com/integrations/terraform-provider-github/issues/337))


## 3.0.0 (September 08, 2020)

BREAKING CHANGES:

- `token` becomes optional
- `organization` no longer deprecated
- `individual` and `anonymous` removed
- `owner` inferred from `organization`
- `base_url` provider parameter no longer requires `/api/v3` suffix

BUG FIXES:

- `terraform validate` fails because of missing token ([#503](https://github.com/integrations/terraform-provider-github/issues/503))
- organization support for various resources ([#501](https://github.com/integrations/terraform-provider-github/issues/501))

ENHANCEMENTS:

* **New Data Source** `github_organization` ([#521](https://github.com/integrations/terraform-provider-github/issues/521))


## 2.9.2 (July 14, 2020)

- Adds deprecation of `anonymous` flag for provider configuration ahead of next major release ([#506](https://github.com/integrations/terraform-provider-github/issues/506))
- Adds deprecation of `individual` flag for provider configuration ahead of next major release ([#512](https://github.com/integrations/terraform-provider-github/issues/512))

## 2.9.1 (July 01, 2020)

BUG FIXES:

- Reverts changes introduced in v2.9.0, deferring to the next major release

## 2.9.0 (June 29, 2020)

**NOTE**: This release introduced a provider-level breaking change around `anonymous` use.
See [here](https://github.com/integrations/terraform-provider-github/pull/464#discussion_r427961161) for details and [here](https://github.com/integrations/terraform-provider-github/issues/502#issuecomment-652379322) to discuss a fix.

ENHANCEMENTS:
* Add Ability To Manage Resources For Non-Organization Accounts ([#464](https://github.com/integrations/terraform-provider-github/issues/464))
* resource/github_repository: Add "internal" Visibility Option ([#454](https://github.com/integrations/terraform-provider-github/issues/454))

## 2.8.1 (June 09, 2020)

BUG FIXES:

* resource/github_repository_file: Reduce API requests when looping through commits ([[#466](https://github.com/integrations/terraform-provider-github/issues/466)])
* resource/github_repository: Fix `auto_init` Destroying Repositories ([[#317](https://github.com/integrations/terraform-provider-github/issues/317)])
* resource/github_repository_deploy_key: Fix acceptance test approach ([[#471](https://github.com/integrations/terraform-provider-github/issues/471)])
* resource/github_actions_secret: Fix Case Where Secret Removed Outside Of Terraform ([[#482](https://github.com/integrations/terraform-provider-github/issues/482)])
* Documentation Addition Of `examples/` Directory

## 2.8.0 (May 15, 2020)

BUG FIXES:

* resource/github_branch_protection: Prevent enabling `dismissal_restrictions` in Github console if `dismissal_users` and `dismissal_teams` are not set ([#453](https://github.com/integrations/terraform-provider-github/issues/453))
* resource/github_repository_collaborator: Allow modifying permissions from `maintain` and `triage`  ([#457](https://github.com/integrations/terraform-provider-github/issues/457))
* Documentation Fix for `github_actions_public_key` data-source ([#458](https://github.com/integrations/terraform-provider-github/issues/458))
* Documentation Fix for `github_branch_protection` resource ([#410](https://github.com/integrations/terraform-provider-github/issues/410))
* Documentation Layout Fix for `github_ip_ranges` and `github_membership` data sources ([#423](https://github.com/integrations/terraform-provider-github/issues/423))
* Documentation Fix for `github_repository_file` import ([#443](https://github.com/integrations/terraform-provider-github/issues/443))
* Update `go-github` to `v31.0.0` ([#424](https://github.com/integrations/terraform-provider-github/issues/424))

ENHANCEMENTS:
* **New Data Source** `github_organization_team_sync_groups` ([#400](https://github.com/integrations/terraform-provider-github/issues/400))
* **New Resource** `github_team_sync_group_mapping` ([#400](https://github.com/integrations/terraform-provider-github/issues/400))

## 2.7.0 (May 01, 2020)

BUG FIXES:

* Add Missing Acceptance Test ([#427](https://github.com/integrations/terraform-provider-github/issues/427))

ENHANCEMENTS:

* Add GraphQL Client ([#331](https://github.com/integrations/terraform-provider-github/issues/331))
* **New Data Source** `github_branch` ([#364](https://github.com/integrations/terraform-provider-github/issues/364))
* **New Resource** `github_branch` ([#364](https://github.com/integrations/terraform-provider-github/issues/364))


## 2.6.1 (April 07, 2020)

BUG FIXES:

* Documentation Fix For Option To Manage `Delete Branch on Merge` ([#408](https://github.com/integrations/terraform-provider-github/issues/408))
* Documentation Fix For `github_actions_secret` / `github_actions_public_key` ([#413](https://github.com/integrations/terraform-provider-github/issues/413))

## 2.6.0 (April 03, 2020)

ENHANCEMENTS:

* resource/github_repository: Introduce Option To Manage `Delete Branch on Merge` ([#399](https://github.com/integrations/terraform-provider-github/issues/399))
* resource/github_repository: Configure Repository As Template ([#357](https://github.com/integrations/terraform-provider-github/issues/357))
* **New Data Source** `github_membership` ([#396](https://github.com/integrations/terraform-provider-github/issues/396))

## 2.5.1 (April 02, 2020)

BUG FIXES:

* Fix Broken Link For `github_actions_secret` Documentation ([#405](https://github.com/integrations/terraform-provider-github/issues/405))

## 2.5.0 (March 30, 2020)

REGRESSION:

* `go-github` `v29.03` is incompatible with versions of GitHub Enterprise Server prior to `v2.21.5`. ([[#404](https://github.com/integrations/terraform-provider-github/issues/404)])

ENHANCEMENTS:

* Add `apps` To `github_branch_protection` Restrictions
* **New Data Source** `github_actions_public_key` ([[#362](https://github.com/integrations/terraform-provider-github/issues/362)])
* **New Data Source** `github_actions_secrets` ([[#362](https://github.com/integrations/terraform-provider-github/issues/362)])
* **New Resource:** `github_actions_secret` ([[#362](https://github.com/integrations/terraform-provider-github/issues/362)])

BUG FIXES:

* Prevent Panic From DismissalRestrictions ([[#385](https://github.com/integrations/terraform-provider-github/issues/385)])
* Update `go-github` to `v29.0.3` ([[#369](https://github.com/integrations/terraform-provider-github/issues/369)])
* Update `go` to `1.13` ([[#372](https://github.com/integrations/terraform-provider-github/issues/372)])
* Documentation Fixes For Consistency And Typography


## 2.4.1 (March 05, 2020)

BUG FIXES:

* Updates `go-github` to `v29` to unblock planned feature development ([#342](https://github.com/integrations/terraform-provider-github/issues/342))
* Fixes `insecure_ssl` parameter behaviour for `github_organization_webhook` and  `github_repository_webhook` ([#365](https://github.com/integrations/terraform-provider-github/issues/365))
* Fixes label behaviour to not create new labels when renaming a `github_issue_label` ([#288](https://github.com/integrations/terraform-provider-github/issues/288))

## 2.4.0 (February 26, 2020)

ENHANCEMENTS:

* **New Data Source** `github_release` ([#356](https://github.com/integrations/terraform-provider-github/pull/356))

* **New Resource:** `github_repository_file` ([#318](https://github.com/integrations/terraform-provider-github/pull/318))

## 2.3.2 (February 05, 2020)

BUG FIXES:

* Handle repository 404 for `github_repository_collaborator` resource ([#348](https://github.com/integrations/terraform-provider-github/issues/348))

## 2.3.1 (January 27, 2020)

BUG FIXES:

* Add support for `triage` and `maintain` permissions in some resources ([#303](https://github.com/integrations/terraform-provider-github/issues/303))

## 2.3.0 (January 22, 2020)

BUG FIXES:

* `resource/resource_github_team_membership`: Prevent spurious diffs caused by upstream API change made on 17th January ([#325](https://github.com/integrations/terraform-provider-github/issues/325))

ENHANCEMENTS:

* `resource/repository`: Added functionality to generate a new repository from a Template Repository ([#309](https://github.com/integrations/terraform-provider-github/issues/309))

## 2.2.1 (September 04, 2019)

ENHANCEMENTS:

* dependencies: Updated module `hashicorp/terraform` to `v0.12.7` ([#273](https://github.com/integrations/terraform-provider-github/issues/273))
* `resource/github_branch_protection`: Will now return an error when users are not correctly added ([#158](https://github.com/integrations/terraform-provider-github/issues/158))
* `provider`: Added optional `anonymous` attribute, and made `token` optional ([#255](https://github.com/integrations/terraform-provider-github/issues/255))

BUG FIXES:
* `resource/github_repository`: Allow setting `default_branch` to `master` on creation ([#150](https://github.com/integrations/terraform-provider-github/issues/150))
* `resource/github_team_repository`: Validation of `team_id` ([#253](https://github.com/integrations/terraform-provider-github/issues/253))
* `resource/github_team_membership`: Validation of `team_id` ([#253](https://github.com/integrations/terraform-provider-github/issues/253))
* `resource/github_organization_webhook`: Properly set webhook secret in state ([#251](https://github.com/integrations/terraform-provider-github/issues/251))
* `resource/github_repository_webhook`: Properly set webhook secret in state ([#251](https://github.com/integrations/terraform-provider-github/issues/251))

## 2.2.0 (June 28, 2019)

FEATURES:

* **New Data Source** `github_collaborators` ([#239](https://github.com/integrations/terraform-provider-github/issues/239))

ENHANCEMENTS:

* `provider`: Added optional `individual` attribute, and made `organization` optional ([#242](https://github.com/integrations/terraform-provider-github/issues/242))
* `resource/github_branch_protection`: Added `require_signed_commits` property ([#214](https://github.com/integrations/terraform-provider-github/issues/214))

BUG FIXES:

* `resource/github_membership`: `username` property is now case insensitive ([#241](https://github.com/integrations/terraform-provider-github/issues/241))
* `resource/github_repository`: `has_projects` can now be imported ([#237](https://github.com/integrations/terraform-provider-github/issues/237))
* `resource/github_repository_collaborator`: `username` property is now case insensitive [[#241](https://github.com/integrations/terraform-provider-github/issues/241))
* `resource/github_team_membership`: `username` property is now case insensitive ([#241](https://github.com/integrations/terraform-provider-github/issues/241))


## 2.1.0 (May 15, 2019)

ENHANCEMENTS:

* `resource/github_repository`: Added validation for lowercase topics ([#223](https://github.com/integrations/terraform-provider-github/issues/223))
* `resource/github_organization_webhook`: Added back removed `name` attribute, `name` is now flagged as `Removed` ([#226](https://github.com/integrations/terraform-provider-github/issues/226))
* `resource/github_repository_webhook`: Added back removed `name` attribute, `name` is now flagged as `Removed` ([#226](https://github.com/integrations/terraform-provider-github/issues/226))

## 2.0.0 (May 02, 2019)

* This release adds support for Terraform 0.12 ([#181](https://github.com/integrations/terraform-provider-github/issues/181))

BREAKING CHANGES:

* `resource/github_repository_webhook`: Removed `name` attribute ([#181](https://github.com/integrations/terraform-provider-github/issues/181))
* `resource/github_organization_webhook`: Removed `name` attribute ([#181](https://github.com/integrations/terraform-provider-github/issues/181))

FEATURES:

* **New Resource:** `github_organization_block` ([#181](https://github.com/integrations/terraform-provider-github/issues/181))
* **New Resource:** `github_user_invitation_accepter` ([#161](https://github.com/integrations/terraform-provider-github/issues/161))
* `resource/github_branch_protection`: Added `required_approving_review_count` property ([#181](https://github.com/integrations/terraform-provider-github/issues/181))

BUG FIXES:

* `resource/github_repository`: Prefill `auto_init` during import ([#154](https://github.com/integrations/terraform-provider-github/issues/154))

## 1.3.0 (September 07, 2018)

FEATURES:

* **New Resource:** `github_project_column` ([#139](https://github.com/integrations/terraform-provider-github/issues/139))

ENHANCEMENTS:

* _all resources_: Use `Etag` to save API quota (~ 33%) ([#143](https://github.com/integrations/terraform-provider-github/issues/143))
* _all resources_: Implement & use RateLimitTransport to avoid hitting API rate limits ([#145](https://github.com/integrations/terraform-provider-github/issues/145))

BUG FIXES:

* `resource/github_issue_label`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/integrations/terraform-provider-github/issues/142))
* `resource/github_membership`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/integrations/terraform-provider-github/issues/142))
* `resource/github_repository_deploy_key`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/integrations/terraform-provider-github/issues/142))
* `resource/github_team`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/integrations/terraform-provider-github/issues/142))
* `resource/github_team_membership`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/integrations/terraform-provider-github/issues/142))
* `resource/github_team_repository`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/integrations/terraform-provider-github/issues/142))
* `resource/github_user_gpg_key`: Return genuine errors instead of ignoring them when reading existing resource ([#142](https://github.com/integrations/terraform-provider-github/issues/142))

## 1.2.1 (August 17, 2018)

BUG FIXES:

* `resource/github_repository`: Avoid spurious diff for `topics` ([#138](https://github.com/integrations/terraform-provider-github/issues/138))

## 1.2.0 (August 17, 2018)

FEATURES:

* **New Data Source:** `github_repository` ([#109](https://github.com/integrations/terraform-provider-github/issues/109))
* **New Data Source:** `github_repositories` ([#129](https://github.com/integrations/terraform-provider-github/issues/129))
* **New Resource:** `github_organization_project` ([#111](https://github.com/integrations/terraform-provider-github/issues/111))
* **New Resource:** `github_repository_project` ([#115](https://github.com/integrations/terraform-provider-github/issues/115))
* **New Resource:** `github_user_gpg_key` ([#120](https://github.com/integrations/terraform-provider-github/issues/120))
* **New Resource:** `github_user_ssh_key` ([#119](https://github.com/integrations/terraform-provider-github/issues/119))

ENHANCEMENTS:

* `provider`: Add `insecure` mode ([#48](https://github.com/integrations/terraform-provider-github/issues/48))
* `data-source/github_ip_ranges`: Add importer IPs ([#100](https://github.com/integrations/terraform-provider-github/issues/100))
* `resource/github_issue_label`: Add support for `description` ([#118](https://github.com/integrations/terraform-provider-github/issues/118))
* `resource/github_repository`: Add support for `topics` ([#97](https://github.com/integrations/terraform-provider-github/issues/97))
* `resource/github_team`: Expose `slug` ([#136](https://github.com/integrations/terraform-provider-github/issues/136))
* `resource/github_team_membership`: Make role updatable ([#137](https://github.com/integrations/terraform-provider-github/issues/137))

BUG FIXES:

* `resource/github_*`: Prevent crashing on invalid ID format ([#108](https://github.com/integrations/terraform-provider-github/issues/108))
* `resource/github_organization_webhook`: Avoid spurious diff of `secret` ([#134](https://github.com/integrations/terraform-provider-github/issues/134))
* `resource/github_repository`: Make non-updatable fields `ForceNew` ([#135](https://github.com/integrations/terraform-provider-github/issues/135))
* `resource/github_repository_deploy_key`: Avoid spurious diff of `key` ([#132](https://github.com/integrations/terraform-provider-github/issues/132))
* `resource/github_repository_webhook`: Avoid spurious diff of `secret` ([#133](https://github.com/integrations/terraform-provider-github/issues/133))


## 1.1.0 (May 11, 2018)

FEATURES:

* **New Data Source:** `github_ip_ranges` ([#82](https://github.com/integrations/terraform-provider-github/issues/82))

ENHANCEMENTS:

* `resource/github_repository`: Add support for archiving ([#64](https://github.com/integrations/terraform-provider-github/issues/64))
* `resource/github_repository`: Add `html_url` ([#93](https://github.com/integrations/terraform-provider-github/issues/93))
* `resource/github_repository`: Add `has_projects` ([#92](https://github.com/integrations/terraform-provider-github/issues/92))
* `resource/github_team`: Add `parent_team_id` ([#54](https://github.com/integrations/terraform-provider-github/issues/54))

## 1.0.0 (February 20, 2018)

ENHANCEMENTS:

* `resource/github_branch_protection`: Add support for `require_code_owners_review` ([#51](https://github.com/integrations/terraform-provider-github/issues/51))

## 0.1.2 (February 12, 2018)

BUG FIXES:

* `resource/github_membership`: Fix a crash when bad import input is given ([#72](https://github.com/integrations/terraform-provider-github/issues/72))

## 0.1.1 (July 18, 2017)

BACKWARDS INCOMPATIBILITIES / NOTES:

* `resource/github_branch_protection`: The `include_admin` attributes haven't been working for quite some time due to upstream API changes. These attributes are now deprecated in favor of the new top-level `enforce_admins` attribute. The `include_admin` attributes currently have no affect on the resource, and will yield a `DEPRECATED` notice to the user.

IMPROVEMENTS:

* `resource/github_repository`: Allow updating default_branch ([#23](https://github.com/integrations/terraform-provider-github/issues/23))
* `resource/github_repository`: Add license_template and gitignore_template ([#24](https://github.com/integrations/terraform-provider-github/issues/24))
* `resource/github_repository_webhook`: Add import ([#29](https://github.com/integrations/terraform-provider-github/issues/29))
* `resource/github_branch_protection`: Support enforce_admins ([#26](https://github.com/integrations/terraform-provider-github/issues/26))
* `resource/github_team`: Supports managing a team's LDAP DN in GitHub Enterprise ([#39](https://github.com/integrations/terraform-provider-github/issues/39))

BUG FIXES:

* `resource/github_branch_protection`: Fix crash on nil values ([#26](https://github.com/integrations/terraform-provider-github/issues/26))

## 0.1.0 (June 20, 2017)

FEATURES:

* **New Resource:** `github_repository_deploy_key` [[#15215](https://github.com/integrations/terraform-provider-github/issues/15215)](https://github.com/hashicorp/terraform/pull/15215)

IMPROVEMENTS:

* `resource/github_repository`: Adding merge types ([#1](https://github.com/integrations/terraform-provider-github/issues/1))
* `data-source/github_user` and `data-source/github_team`: Added attributes ([#2](https://github.com/integrations/terraform-provider-github/issues/2))
