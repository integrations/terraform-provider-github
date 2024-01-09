output "github_release_assets_url" {
  description = "Asset URL of a GitHub release"
  value       = data.github_release.by_tag.assets[0].url
}
