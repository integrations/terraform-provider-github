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
