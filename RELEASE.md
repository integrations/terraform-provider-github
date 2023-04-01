## Release Flow

Since the migration to the [Terraform registry](https://registry.terraform.io/), this repository's maintainers now have
the ability to self-publish Terraform GitHub provider releases. This process uses GitHub Actions
and [`goreleaser`](https://github.com/goreleaser/goreleaser) to build, sign, and upload provider binaries to a GitHub release.

The release flow is as follows:
1. Navigate to the [repository's Releases page](https://github.com/integrations/terraform-provider-github/releases) and click "Draft a new release".
1. Create a new tag that makes sense with the project's semantic versioning.
1. Auto-generate the release notes.
1. Click "Publish release".
1. GitHub Actions will trigger the release workflow which can be
[viewed here](https://github.com/integrations/terraform-provider-github/actions?query=workflow%3Arelease).
After the workflow executes successfully, the GitHub release created in the prior step will
have the relevant assets available for consumption.
1. The new release will show up in https://registry.terraform.io/providers/integrations/github/latest for consumption
by Terraform users.
1. For terraform `0.12.X` users, the new release is available for consumption once it is present in
https://releases.hashicorp.com/terraform-provider-github/.
