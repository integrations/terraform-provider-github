# Repository `delete_branch_on_merge` Example

This displays configurability of the `delete_branch_on_merge` feature for GitHub repositories.

This example will create a repository in the specified `owner` organization. See https://www.terraform.io/docs/providers/github/index.html for details on configuring [`providers.tf`](./providers.tf) accordingly.

Alternatively, you may use variables passed via command line:

```console
export GITHUB_ORG=
export GITHUB_TOKEN=
```

```console
terraform apply \
  -var "organization=${GITHUB_ORG}" \
  -var "github_token=${GITHUB_TOKEN}"
```
