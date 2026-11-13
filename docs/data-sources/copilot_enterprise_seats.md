---
page_title: "github_copilot_enterprise_seats (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list the GitHub Copilot seats billed to an enterprise, across its organizations and enterprise teams.
---

# github_copilot_enterprise_seats (Data Source)

Data source to list the GitHub Copilot seats billed to an enterprise, across its organizations and enterprise teams.

~> **Note:** This data source requires GitHub Enterprise Cloud. Only enterprise owners and billing managers can list Copilot seats; see [List all Copilot seat assignments for an enterprise](https://docs.github.com/enterprise-cloud@latest/rest/copilot/copilot-user-management#list-all-copilot-seat-assignments-for-an-enterprise) for the required token scopes. This API endpoint is in public preview and is subject to change.

## Example Usage

```terraform
data "github_copilot_enterprise_seats" "example" {
  enterprise_slug = "example-co"
}

# Users holding a Copilot seat with no recorded Copilot activity.
output "copilot_users_without_activity" {
  value = distinct([
    for seat in data.github_copilot_enterprise_seats.example.seats : seat.assignee_login
    if seat.last_activity_at == ""
  ])
}
```

<!--
## Schema

### Required

- `enterprise_slug` (String) The slug of the enterprise.

### Read-Only

- `id` (String) The ID of this resource.
- `seats` (List of Object) Copilot seat assignments, sorted by assignee login. A user who has access through several organizations or enterprise teams has one entry for each of them. (see [below for nested schema](#nestedatt--seats))
- `total_seats` (Number) Number of Copilot seats billed to the enterprise. A user who has access through several organizations or enterprise teams is counted once.

<a id="nestedatt--seats"></a>
### Nested Schema for `seats`

Read-Only:

- `assignee_login` (String)
- `assigning_team_slug` (String)
- `created_at` (String)
- `last_activity_at` (String)
- `last_activity_editor` (String)
- `pending_cancellation_date` (String)
- `plan_type` (String)
-->

## Schema

### Required

- `enterprise_slug` (String) The slug of the enterprise.

### Read-Only

- `id` (String) The ID of this resource.
- `seats` (List of Object) Copilot seat assignments, sorted by assignee login. A user who has access through several organizations or enterprise teams has one entry for each of them. (see [below for nested schema](#nestedatt--seats))
- `total_seats` (Number) Number of Copilot seats billed to the enterprise. A user who has access through several organizations or enterprise teams is counted once.

<a id="nestedatt--seats"></a>
### Nested Schema for `seats`

Read-Only:

- `assignee_login` (String) Login of the user the seat is assigned to.
- `assigning_team_slug` (String) Slug of the team or enterprise team through which the seat is granted. Empty when the seat is assigned directly.
- `created_at` (String) Time the assignee was last granted access to Copilot, in RFC 3339 format.
- `last_activity_at` (String) Time of the assignee's last Copilot activity, in RFC 3339 format. Empty when no activity has been recorded.
- `last_activity_editor` (String) Editor used for the assignee's last Copilot activity. Empty when no activity has been recorded.
- `pending_cancellation_date` (String) Date, in `YYYY-MM-DD` format, on which the seat will be cancelled. Empty when no cancellation is pending.
- `plan_type` (String) Copilot plan of the seat, for example `business` or `enterprise`.
