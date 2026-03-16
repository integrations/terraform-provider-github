data "github_release" "example" {
    repository  = "example-repository"
    owner       = "example-owner"
    retrieve_by = "latest"
}

data "github_release_asset" "example" {
    repository  = "example-repository"
    owner       = "example-owner"
    asset_id    = data.github_release.example.assets[0].id
}
