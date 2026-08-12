resource "github_enterprise_ip_allow_list_entry" "test" {
  enterprise_slug = "my-enterprise"
  ip              = "192.168.1.0/20"
  name            = "My IP Range Name"
  is_active       = true
}
