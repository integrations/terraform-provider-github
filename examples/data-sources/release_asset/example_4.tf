data "github_release" "example" {
  repository  = "example-repository"
  owner       = "example-owner"
  retrieve_by = "latest"
}

data "github_release_asset" "example" {
  count      = length(data.github_release.example.assets)
  repository = "example-repository"
  owner      = "example-owner"
  asset_id   = data.github_release.example.assets[count.index].id
}
