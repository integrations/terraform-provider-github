data "github_app_token" "this" {
  app_id          = "123456"
  installation_id = "78910"
  pem_file        = file("foo/bar.pem")
}
