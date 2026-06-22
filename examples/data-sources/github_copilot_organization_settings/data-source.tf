data "github_copilot_organization_settings" "example" {}

output "seat_management" {
  value = data.github_copilot_organization_settings.example.seat_management_setting
}

output "total_seats" {
  value = data.github_copilot_organization_settings.example.seat_breakdown[0].total
}
