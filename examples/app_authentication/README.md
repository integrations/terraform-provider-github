# App Installation Example

This example demonstrates authenticating using a GitHub App.

The example will create a repository in the specified organization.

## Using a PEM file

```console
export GITHUB_OWNER=
export GITHUB_APP_ID=
export GITHUB_APP_INSTALLATION_ID=
export GITHUB_APP_PEM_FILE=
```

## Using a pre-signed JWT

If you sign the GitHub App JWT externally (e.g., using AWS KMS or HashiCorp Vault),
you can pass the signed JWT directly instead of providing a PEM file.
In this case, `GITHUB_APP_ID` is not required.

```console
export GITHUB_OWNER=
export GITHUB_APP_INSTALLATION_ID=
export GITHUB_APP_JWT=
```

```console
terraform apply -var "organization=${GITHUB_ORG}"
```
