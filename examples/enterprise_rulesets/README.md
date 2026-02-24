# GitHub Enterprise Ruleset Examples

This directory demonstrates how to configure GitHub Enterprise rulesets using the Terraform GitHub provider.

## Overview

Enterprise rulesets allow you to enforce policies across all organizations in your GitHub Enterprise. The examples showcase all four target types:

- **Branch Target** (`branch_target.tf`) - Branch protection rules with PR requirements, status checks, and commit patterns
- **Tag Target** (`tag_target.tf`) - Tag protection rules with naming patterns and immutability controls
- **Push Target** (`push_target.tf`) - File restrictions, size limits, and content policies (beta feature)
- **Repository Target** (`rulesets.tf`) - Repository management rules for creation, deletion, and naming conventions

## Requirements

- GitHub Enterprise Cloud account
- Personal access token with enterprise admin permissions
- Terraform >= 0.14

## Usage

1. Set your environment variables:

```bash
export TF_VAR_github_token="your_github_token"
export TF_VAR_enterprise_slug="your-enterprise-slug"
```

2. Customize the examples by replacing `"your-enterprise"` with your actual enterprise slug

3. Apply the configuration:

```bash
terraform init
terraform plan
terraform apply
```

## Target Types

Each target type supports different rules:

- **Branch/Tag**: creation, deletion, update, signatures, linear history, PR requirements, status checks
- **Push**: file restrictions, size limits, file extensions, commit patterns
- **Repository**: creation, deletion, transfer, naming patterns, visibility controls

See the individual `.tf` files for detailed examples and available rules.

## Important Notes

- All enterprise rulesets require organization and repository targeting via `conditions`
- The `push` target is currently in beta and subject to change
- Branch and tag targets require `ref_name` conditions
- Repository and push targets do not use `ref_name` conditions
