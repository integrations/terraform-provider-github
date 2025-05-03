# Organization Security Manager Example

This example demonstrates creating an organization security manager team.

It will:
- Create a team with the specified `team_name` in the specified `owner` organization
- Assign the organization security manager role to the team

The GitHub token must have the `admin:org` scope.

```console
export GITHUB_OWNER=my-organization
export GITHUB_TOKEN=ghp_###
export GITHUB_TEAM_NAME="My Security Manager Team"
```

```console
terraform apply \
  -var "owner=${GITHUB_OWNER}" \
  -var "github_token=${GITHUB_TOKEN}" \
  -var "team_name=${GITHUB_TEAM_NAME}"
```
