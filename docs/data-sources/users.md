---
page_title: "github_users Data Source - terraform-provider-github
description: |-
  Get information about multiple GitHub users.
---

# github_users (Data Source)

Use this data source to retrieve information about multiple GitHub users at once.

## Example Usage

```terraform
# Retrieve information about multiple GitHub users.
data "github_users" "example" {
  usernames = ["example1", "example2", "example3"]
}

output "valid_users" {
  value = "${data.github_users.example.logins}"
}

output "invalid_users" {
  value = "${data.github_users.example.unknown_logins}"
}
```

## Argument Reference

- `usernames` - (Required) List of usernames.

## Attributes Reference

- `node_ids` - list of Node IDs of users that could be found.
- `logins` - list of logins of users that could be found.
- `emails` - list of the user's publicly visible profile email (will be empty string in case if user decided not to show it).
- `unknown_logins` - list of logins without matching user.
