output "cost_center" {
  description = "Created cost center"
  value = {
    id                 = github_enterprise_cost_center.example.id
    name               = github_enterprise_cost_center.example.name
    state              = github_enterprise_cost_center.example.state
    azure_subscription = github_enterprise_cost_center.example.azure_subscription
  }
}

output "cost_center_from_data_source" {
  description = "Cost center fetched by the data source, including all assignments"
  value = {
    id            = data.github_enterprise_cost_center.by_id.cost_center_id
    name          = data.github_enterprise_cost_center.by_id.name
    state         = data.github_enterprise_cost_center.by_id.state
    users         = sort(tolist(data.github_enterprise_cost_center.by_id.users))
    organizations = sort(tolist(data.github_enterprise_cost_center.by_id.organizations))
    repositories  = sort(tolist(data.github_enterprise_cost_center.by_id.repositories))
  }
}
