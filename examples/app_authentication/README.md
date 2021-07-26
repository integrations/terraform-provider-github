# App Installation Example

This example demonstrates authenticating using a GitHub App.

The example will create a repository in the specified organization.

You may use variables passed via command line:

```console
export GITHUB_OWNER=
export GITHUB_APP_ID=
export GITHUB_APP_INSTALLATION_ID=
export GITHUB_APP_PEM_FILE=
```

```console
terraform apply -var "organization=${GITHUB_ORG}"
```
