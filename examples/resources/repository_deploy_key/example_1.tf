# Generate an ssh key using provider "hashicorp/tls"
resource "tls_private_key" "example_repository_deploy_key" {
  algorithm = "ED25519"
}

# Add the ssh key as a deploy key
resource "github_repository_deploy_key" "example_repository_deploy_key" {
  title      = "Repository test key"
  repository = "test-repo"
  key        = tls_private_key.example_repository_deploy_key.public_key_openssh
  read_only  = true
}
