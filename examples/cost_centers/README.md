# Enterprise cost centers

This example creates an enterprise cost center, manages its user, organization,
and repository assignments, and reads the resulting configuration through the
cost center data sources.

The token must have access to the enterprise cost center endpoints. See the
[GitHub REST API documentation](https://docs.github.com/enterprise-cloud@latest/rest/billing/cost-centers)
for current permission requirements.

Set the required variables in a `terraform.tfvars` file, then run:

```shell
terraform init
terraform apply
```
