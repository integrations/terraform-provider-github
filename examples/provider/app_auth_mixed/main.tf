provider "github" {
  owner = var.github_organization
  app_auth {
    id              = "123456"   # the App ID and installation ID are not secret,
    installation_id = "78901234" # so they can be set directly in configuration
    # pem_file is omitted; it falls back to the GITHUB_APP_PEM_FILE variable,
    # keeping the secret out of configuration
  }
}
