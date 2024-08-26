resource "github_user_ssh_signing_key" "example" {
  title = "example title"
  key   = file("~/.ssh/id_rsa.pub")
}
