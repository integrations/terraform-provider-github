locals {
  tenant_namespaces = {
    test01 = { name = "test01" }
    test02 = { name = "test02" }
    test03 = { name = "test03" }
  }
}

resource "github_repository_files" "tenants" {
  repository     = "example"
  branch         = "main"
  commit_message = "chore: sync tenants"
  commit_author  = "Terraform"
  commit_email   = "tf@example.com"

  dynamic "file" {
    for_each = local.tenant_namespaces
    content {
      path    = "tenants/${file.key}.yaml"
      content = yamlencode(file.value)
    }
  }
}
