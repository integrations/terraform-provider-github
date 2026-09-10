set -euo pipefail

# Import using the network configuration ID. Configure the provider's owner to the organization.
terraform import github_organization_network_configuration.example 123456789ABCDEF
