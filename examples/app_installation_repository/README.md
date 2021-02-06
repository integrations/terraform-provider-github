# App Installation Example

This example gives an application installation access to a
specific repository in the same organization.

To complete this demo, first [install an application in your
github organization](https://docs.github.com/en/github/customizing-your-github-workflow/installing-an-app-in-your-organization). To use the full scope of this
resource, make sure you install the application only on select repositories
in the organization (instead of all repositories).

This will allow you to use this resource to manage which repositories
the app installation has access to.

After you have installed the application, locate the installation id of the
application by visiting `https://github.com/organizations/{ORG_NAME}/settings/installations`
and configuring the app you'd like to install.
The ID should be located in the URL on the configure page.

This example will create a repository in the specified organization.
It will also add the created repository to the app installation.

Alternatively, you may use variables passed via command line:

```console
export GITHUB_ORG=
export GITHUB_TOKEN=
export INSTALLATION_ID=
```

```console
terraform apply \
  -var "organization=${GITHUB_ORG}" \
  -var "github_token=${GITHUB_TOKEN}" \
  -var "installation_id=${INSTALLATION_ID}" \
```
