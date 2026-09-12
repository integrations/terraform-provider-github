resource "github_organization_ip_allow_list_entry" "test" {
  ip        = "192.168.1.0/20"
  name      = "My IP Range Name"
  is_active = true
}
