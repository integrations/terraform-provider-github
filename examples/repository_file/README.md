# Repository File

This provides a template for managing [repository files](https://docs.github.com/en/repositories/working-with-files/managing-files).

This example will also create or update a file in the specified `repository`. See https://www.terraform.io/docs/providers/github/index.html for details on configuring [`providers.tf`](./providers.tf) accordingly.

Alternatively, you may use variables passed via the command line or `auto.tfvars`:

```tfvars
organization = ""
github_token = ""

repository     = ""
file           = ""
content        = ""
branch         = ""
commit_author  = ""
commit_message = ""
commit_email   = ""
```

```console
terraform apply
```
