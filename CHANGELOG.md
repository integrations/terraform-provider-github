## 0.1.1 (July 18, 2017)

BACKWARDS INCOMPATIBILITIES / NOTES:

* With the `github_branch_protection` resource, the `include_admin` attributes haven't been working for quite some time due to upstream API changes. These attributes are now deprecated in favor of the new top-level `enforce_admins` attribute. The `include_admin` attributes currently have no affect on the resource, and will yield a `DEPRECATED` notice to the user. 

IMPROVEMENTS:

* r/github_repository: Allow updating default_branch ([#23](https://github.com/terraform-providers/terraform-provider-github/issues/23))
* r/github_repository: Add license_template and gitignore_template ([#24](https://github.com/terraform-providers/terraform-provider-github/issues/24))
* r/github_repository_webhook: Add import ([#29](https://github.com/terraform-providers/terraform-provider-github/issues/29))
* r/github_branch_protection: Support enforce_admins ([#26](https://github.com/terraform-providers/terraform-provider-github/issues/26))
* Supports managing a team's LDAP DN in GitHub Enterprise ([#39](https://github.com/terraform-providers/terraform-provider-github/issues/39))

BUG FIXES: 

* r/github_branch_protection: Fix crash on nil values ([#26](https://github.com/terraform-providers/terraform-provider-github/issues/26))

## 0.1.0 (June 20, 2017)

FEATURES:

* **New Resource:** `github_repository_deploy_key` [[#15215](https://github.com/terraform-providers/terraform-provider-github/issues/15215)](https://github.com/hashicorp/terraform/pull/15215)

IMPROVEMENTS:

* Adding merge types to repository resource ([#1](https://github.com/terraform-providers/terraform-provider-github/issues/1))
* added attributes to github_user and github_team data sources ([#2](https://github.com/terraform-providers/terraform-provider-github/issues/2))
