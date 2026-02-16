# Release Flow

The release process uses GitHub Actions and [`goreleaser`](https://github.com/goreleaser/goreleaser) to build, sign, and upload provider binaries to a GitHub release. Release are triggered by a tag with the pattern `v*` (e.g. `v1.2.3`); these tags may only be created from the default branch (`main`) or branches that match the pattern `release-v*`.

The release flow is as follows:

[!IMPORTANT]
> In you're planning on releasing a major version, please ensure you've completed the following tasks:
> 
> - Read Hashicorp guidance on [incrementing the major version](https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#example-major-number-increments).
> - Check if there are any outstanding [PRs with breaking changes](https://github.com/integrations/terraform-provider-github/issues?q=state%3Aopen%20label%3A%22Type%3A%20Breaking%20change%22) that could be included in the release.
> - Check that all deprecations have been addressed and removed from the codebase.

1. Navigate to the [repository's Releases page](https://github.com/integrations/terraform-provider-github/releases) and click _Draft a new release_.
1. Create a new [SemVer](https://semver.org/) tag for the release.
1. Select the target as either the default branch (`main`) or a release branch (a branch matching the pattern `release-v*`)
1. Click _Generate release notes_.
1. If this release is from a release branch (unless it really is the latest release) uncheck the _Set as the latest release_ checkbox.
1. Click "Publish release".
1. GitHub Actions will trigger the [release workflow](https://github.com/integrations/terraform-provider-github/actions/workflows/release.yaml).

After the workflow executes successfully, the GitHub release created in the prior step will have the relevant assets available for consumption and the new version will show up in the [Terraform Registry](https://registry.terraform.io/providers/integrations/github/latest).
