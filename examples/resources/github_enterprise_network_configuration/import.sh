set -euo pipefail

# Import using <enterprise_slug>/<network_configuration_id>.
terraform import github_enterprise_network_configuration.example my-enterprise/123456789ABCDEF
