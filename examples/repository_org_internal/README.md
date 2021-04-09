# Repository Visibility with Org, type internal

This demos various repository [visibility settings](https://help.github.com/en/github/administering-a-repository/setting-repository-visibility) for repositories.

This example will create a repository in the specified `owner` organization. See https://www.terraform.io/docs/providers/github/index.html for details on configuring [`providers.tf`](./providers.tf) accordingly.

In order to build the provider for use with this example, see [Building the Provider docs](https://github.com/integrations/terraform-provider-github/blob/master/CONTRIBUTING.md#building-the-provider)

Alternatively, you may use variables passed via command line:

```console
export GITHUB_OWNER=
export GITHUB_TOKEN=
```

```console
terraform apply \
  -var "owner=${GITHUB_OWNER}" \
  -var "github_token=${GITHUB_TOKEN}"
```
