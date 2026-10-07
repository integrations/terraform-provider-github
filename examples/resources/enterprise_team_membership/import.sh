# When configuration uses team_id:
terraform import github_enterprise_team_membership.member enterprise-slug/12345/username

# When configuration uses team_slug:
terraform import github_enterprise_team_membership.member enterprise-slug/ent:platform/username
