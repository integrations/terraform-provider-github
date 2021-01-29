# Repository Collaborator

This provides a template for managing [repository collaborators](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/inviting-collaborators-to-a-personal-repository).

This example will also create a repository in the specified `owner` organization. See https://www.terraform.io/docs/providers/github/index.html for details on configuring [`providers.tf`](./providers.tf) accordingly.

Alternatively, you may use variables passed via command line:

```console
export GITHUB_ORGANIZATION=
export GITHUB_TOKEN=
export COLLABORATOR_USERNAME=
export COLLABORATOR_PERMISSION=
```

```console
terraform apply \
  -var "organization=${GITHUB_ORGANIZATION}" \
  -var "github_token=${GITHUB_TOKEN}" \
  -var "username=${COLLABORATOR_USERNAME}" \
  -var "permission=${COLLABORATOR_PERMISSION}"
```
