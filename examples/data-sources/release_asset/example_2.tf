data "github_release_asset" "example" {
    repository    = "example-repository"
    owner         = "example-owner"
    asset_id      = 12345
    download_file = true
}
