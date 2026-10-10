# Enterprise cost centers

This example creates an enterprise cost center, manages its user, organization,
and repository assignments, and reads the resulting configuration through the
cost center data sources.

The token must have access to the enterprise cost center endpoints. See the
[GitHub REST API documentation](https://docs.github.com/enterprise-cloud@latest/rest/billing/cost-centers)
for current permission requirements.

These resources are not available in a released provider version yet. Run
this example with a provider built from the pull request branch:

```shell
go build -o ~/go/bin/terraform-provider-github ../..
export TF_CLI_CONFIG_FILE="$(cd .. && pwd)/dev.tfrc"
terraform apply
```

The `TF_CLI_CONFIG_FILE` setting activates the repository's development
override for `integrations/github`. Terraform should display the
`Provider development overrides are in effect` warning. Set the required
variables in a `terraform.tfvars` file before applying.
