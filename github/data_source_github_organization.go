package github

import (
	"strconv"

	"github.com/google/go-github/v74/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubOrganization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ignore_archived_repos": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"orgname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repositories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"members": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Deprecated: "Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version.",
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
			"default_repository_permission": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"members_can_create_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"two_factor_requirement_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"members_allowed_repository_creation_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"members_can_create_public_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"members_can_create_private_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"members_can_create_internal_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"members_can_create_pages": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"members_can_create_public_pages": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"members_can_create_private_pages": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"members_can_fork_private_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"web_commit_signoff_required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"advanced_security_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dependabot_alerts_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dependabot_security_updates_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dependency_graph_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"secret_scanning_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"secret_scanning_push_protection_enabled_for_new_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"summary_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func dataSourceGithubOrganizationRead(d *schema.ResourceData, meta any) error {
	name := d.Get("name").(string)

	client4 := meta.(*Owner).v4client
	client3 := meta.(*Owner).v3client
	ctx := meta.(*Owner).StopContext

	organization, _, err := client3.Organizations.Get(ctx, name)
	if err != nil {
		return err
	}

	var planName string

	if plan := organization.GetPlan(); plan != nil {
		planName = plan.GetName()
	}

	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100, Page: 1},
	}

	summaryOnly := d.Get("summary_only").(bool)
	if !summaryOnly {
		var repoList []string
		var allRepos []*github.Repository

		for {
			repos, resp, err := client3.Repositories.ListByOrg(ctx, name, opts)
			if err != nil {
				return err
			}
			allRepos = append(allRepos, repos...)

			opts.Page = resp.NextPage

			if resp.NextPage == 0 {
				break
			}
		}

		ignoreArchiveRepos := d.Get("ignore_archived_repos").(bool)
		for index := range allRepos {
			repo := allRepos[index]
			if ignoreArchiveRepos && repo.GetArchived() {
				continue
			}

			repoList = append(repoList, repo.GetFullName())
		}

		var query struct {
			Organization struct {
				MembersWithRole struct {
					Edges []struct {
						Role githubv4.String
						Node struct {
							Id    githubv4.String
							Login githubv4.String
							Email githubv4.String
						}
					}
					PageInfo struct {
						EndCursor   githubv4.String
						HasNextPage bool
					}
				} `graphql:"membersWithRole(first: 100, after: $after)"`
			} `graphql:"organization(login: $login)"`
		}
		variables := map[string]any{
			"login": githubv4.String(name),
			"after": (*githubv4.String)(nil),
		}
		var members []string
		var users []map[string]string
		for {
			err := client4.Query(ctx, &query, variables)
			if err != nil {
				return err
			}
			for _, edge := range query.Organization.MembersWithRole.Edges {
				members = append(members, string(edge.Node.Login))
				users = append(users, map[string]string{
					"id":    string(edge.Node.Id),
					"login": string(edge.Node.Login),
					"email": string(edge.Node.Email),
					"role":  string(edge.Role),
				})
			}
			if !query.Organization.MembersWithRole.PageInfo.HasNextPage {
				break
			}
			variables["after"] = githubv4.NewString(query.Organization.MembersWithRole.PageInfo.EndCursor)
		}

		_ = d.Set("repositories", repoList)
		_ = d.Set("members", members)
		_ = d.Set("users", users)
		_ = d.Set("two_factor_requirement_enabled", organization.GetTwoFactorRequirementEnabled())
		_ = d.Set("default_repository_permission", organization.GetDefaultRepoPermission())
		_ = d.Set("members_can_create_repositories", organization.GetMembersCanCreateRepos())
		_ = d.Set("members_allowed_repository_creation_type", organization.GetMembersAllowedRepositoryCreationType())
		_ = d.Set("members_can_create_public_repositories", organization.GetMembersCanCreatePublicRepos())
		_ = d.Set("members_can_create_private_repositories", organization.GetMembersCanCreatePrivateRepos())
		_ = d.Set("members_can_create_internal_repositories", organization.GetMembersCanCreateInternalRepos())
		_ = d.Set("members_can_fork_private_repositories", organization.GetMembersCanCreatePrivateRepos())
		_ = d.Set("web_commit_signoff_required", organization.GetWebCommitSignoffRequired())
		_ = d.Set("members_can_create_pages", organization.GetMembersCanCreatePages())
		_ = d.Set("members_can_create_public_pages", organization.GetMembersCanCreatePublicPages())
		_ = d.Set("members_can_create_private_pages", organization.GetMembersCanCreatePrivatePages())
		_ = d.Set("advanced_security_enabled_for_new_repositories", organization.GetAdvancedSecurityEnabledForNewRepos())
		_ = d.Set("dependabot_alerts_enabled_for_new_repositories", organization.GetDependabotAlertsEnabledForNewRepos())
		_ = d.Set("dependabot_security_updates_enabled_for_new_repositories", organization.GetDependabotSecurityUpdatesEnabledForNewRepos())
		_ = d.Set("dependency_graph_enabled_for_new_repositories", organization.GetDependencyGraphEnabledForNewRepos())
		_ = d.Set("secret_scanning_enabled_for_new_repositories", organization.GetSecretScanningEnabledForNewRepos())
		_ = d.Set("secret_scanning_push_protection_enabled_for_new_repositories", organization.GetSecretScanningPushProtectionEnabledForNewRepos())
	}

	d.SetId(strconv.FormatInt(organization.GetID(), 10))
	_ = d.Set("login", organization.GetLogin())
	_ = d.Set("name", organization.GetName())
	_ = d.Set("orgname", name)
	_ = d.Set("node_id", organization.GetNodeID())
	_ = d.Set("description", organization.GetDescription())
	_ = d.Set("plan", planName)

	return nil
}
