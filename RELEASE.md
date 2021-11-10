## Release Flow

Since the migration to the [Terraform registry](https://registry.terraform.io/), this repository's maintainers now have
the ability to self-publish Terraform GitHub provider releases. This process leverages Github Actions
and [`goreleaser`](https://github.com/goreleaser/goreleaser) to build, sign, and upload provider binaries to a Github release.

The release flow is as follows:
1. Create a CHANGELOG entry for the release using the following format:
    ```markdown
    ## x.y.z (release date)

    ## ENHANCEMENTS:
    ...

    ## BUG FIXES:
    ...
    ```
1. Tag the commit that adds the CHANGELOG entry with the release version and push:
    ```shell
    $ git tag x.y.z
    $ git push origin x.y.z
    ```
1. Github Actions will trigger the release workflow which can be
[viewed here](https://github.com/integrations/terraform-provider-github/actions?query=workflow%3Arelease).
After the workflow executes successfully, the Github release created in the prior step will
have the relevant assets available for consumption.
1. The new release will show up in https://registry.terraform.io/providers/integrations/github/latest for consumption
by terraform `0.13.X` users.
1. For terraform `0.12.X` users, the new release is available for consumption once it is present in
https://releases.hashicorp.com/terraform-provider-github/.
