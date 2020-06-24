# Release Example

This displays retrieval of a GitHub release.

This example will look up a GitHub release available to the specified `owner` organization or a personal account. See https://www.terraform.io/docs/providers/github/index.html for details on configuring [`providers.tf`](./providers.tf) accordingly.

Alternatively, you may use variables passed via command line:

```console
export GITHUB_ORG=
export GITHUB_TOKEN=
export RELEASE_OWNER=
export RELEASE_REPOSITORY=
export RELEASE_TAG=
```
```console
terraform apply \
  -var "organization=${GITHUB_ORG}" \
  -var "github_token=${GITHUB_TOKEN}" \
  -var "owner=${RELEASE_OWNER}" \
  -var "repository=${RELEASE_REPOSITORY}" \
  -var "release_tag=${RELEASE_TAG}"
```
