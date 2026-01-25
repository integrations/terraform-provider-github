package github

import (
	"strconv"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubOrganization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the organization.",
			},
			"ignore_archived_repos": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Whether or not to include archived repos in the repositories list.",
			},
			"orgname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The organization's name as used in URLs and the API.",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GraphQL global node ID for use with the v4 API.",
			},
			"login": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The organization account login.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The organization account description.",
			},
			"plan": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The organization account plan name.",
			},
			"repositories": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of the full names of the repositories in the organization.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"members": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of the organization's members.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Deprecated: "Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version.",
			},
			"users": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list with the members of the organization containing their id, login, email, and role.",
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
			"default_repository_permission": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default permission level members have for organization repositories.",
			},
			"members_can_create_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether non-admin organization members can create repositories.",
			},
			"two_factor_requirement_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether two-factor authentication is required for all members of the organization.",
			},
			"members_allowed_repository_creation_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of repository allowed to be created by members of the organization.",
			},
			"members_can_create_public_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether organization members can create public repositories.",
			},
			"members_can_create_private_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether organization members can create private repositories.",
			},
			"members_can_create_internal_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether organization members can create internal repositories.",
			},
			"members_can_create_pages": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether organization members can create pages sites.",
			},
			"members_can_create_public_pages": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether organization members can create public pages sites.",
			},
			"members_can_create_private_pages": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether organization members can create private pages sites.",
			},
			"members_can_fork_private_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether organization members can fork private repositories.",
			},
			"web_commit_signoff_required": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether organization members must sign all commits.",
			},
			"advanced_security_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether advanced security is enabled for new repositories.",
			},
			"dependabot_alerts_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether Dependabot alerts is automatically enabled for new repositories.",
			},
			"dependabot_security_updates_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether Dependabot security updates is automatically enabled for new repositories.",
			},
			"dependency_graph_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether dependency graph is automatically enabled for new repositories.",
			},
			"secret_scanning_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether secret scanning is automatically enabled for new repositories.",
			},
			"secret_scanning_push_protection_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether secret scanning push protection is automatically enabled for new repositories.",
			},
			"summary_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Exclude the repos, members and other attributes from the returned result.",
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
