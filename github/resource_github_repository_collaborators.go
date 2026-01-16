package github

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryCollaborators() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryCollaboratorsCreateOrUpdate,
		ReadContext:   resourceGithubRepositoryCollaboratorsRead,
		UpdateContext: resourceGithubRepositoryCollaboratorsCreateOrUpdate,
		DeleteContext: resourceGithubRepositoryCollaboratorsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of users.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "push",
						},
						"username": {
							Type:             schema.TypeString,
							Description:      "(Required) The user to add to the repository as a collaborator.",
							Required:         true,
							DiffSuppressFunc: caseInsensitive(),
						},
					},
				},
			},
			"team": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of teams.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "push",
						},
						"slug": {
							Type:        schema.TypeString,
							Description: "Slug of the team to add to the repository as a collaborator.",
							Optional:    true,
						},
						"team_id": {
							Type:        schema.TypeString,
							Description: "ID of the team to add to the repository as a collaborator.",
							Optional:    true,
							Deprecated:  "Use slug.",
						},
					},
				},
			},
			"ignore_team": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of teams to ignore.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"slug": {
							Type:        schema.TypeString,
							Description: "Slug of the team to add to ignore.",
							Optional:    true,
						},
						"team_id": {
							Type:        schema.TypeString,
							Description: "ID or slug of the team to ignore.",
							Optional:    true,
						},
					},
				},
			},
			"invitation_ids": {
				Type:        schema.TypeMap,
				Description: "Map of usernames to invitation ID for any users added",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},

		CustomizeDiff: customdiff.Sequence(
			// If there was a new user added to the list of collaborators,
			// it's possible a new invitation id will be created in GitHub.
			customdiff.ComputedIf("invitation_ids", func(ctx context.Context, d *schema.ResourceDiff, meta any) bool {
				return d.HasChange("user")
			}),
		),
	}
}

func resourceGithubRepositoryCollaboratorsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	isOrg := meta.IsOrganization

	repoName := d.Get("repository").(string)
	users := d.Get("user").(*schema.Set).List()
	teams := d.Get("team").(*schema.Set).List()
	ignoreTeams := d.Get("ignore_team").(*schema.Set).List()

	inUsers, err := getUserCollaborators(users)
	if err != nil {
		return diag.FromErr(err)
	}

	if checkDuplicateUsers(inUsers) {
		return diag.Errorf("duplicate usernames found in user collaborators")
	}

	inTeams, err := getTeamCollaborators(ctx, meta, teams)
	if err != nil {
		return diag.FromErr(err)
	}

	if checkDuplicateTeams(inTeams) {
		return diag.Errorf("duplicate teams found in team collaborators")
	}

	inIgnoreTeams, err := getTeamIdentities(ctx, meta, ignoreTeams)
	if err != nil {
		return diag.FromErr(err)
	}

	err = updateUserCollaboratorsAndInvites(ctx, client, owner, repoName, inUsers)
	if err != nil {
		return diag.FromErr(err)
	}

	if isOrg {
		err = updateTeamCollaborators(ctx, client, owner, repoName, inTeams, inIgnoreTeams)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(repoName)

	ghInvitations, err := listInvitations(ctx, client, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	invitationIds := make(map[string]string, len(ghInvitations))
	for _, i := range ghInvitations {
		invitationIds[i.login] = strconv.FormatInt(*i.invitationID, 10)
	}

	if err = d.Set("invitation_ids", invitationIds); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCollaboratorsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	isOrg := meta.IsOrganization

	repoName := d.Id()
	ignoreTeams := d.Get("ignore_team").(*schema.Set).List()

	inIgnoreTeams, err := getTeamIdentities(ctx, meta, ignoreTeams)
	if err != nil {
		return diag.FromErr(err)
	}

	ghUsers, err := listUserCollaborators(ctx, client, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("user", ghUsers.flatten())
	if err != nil {
		return diag.FromErr(err)
	}

	if isOrg {
		ghTeams, err := listTeamCollaborators(ctx, client, owner, repoName, inIgnoreTeams)
		if err != nil {
			return diag.FromErr(err)
		}

		err = d.Set("team", ghTeams.flatten())
		if err != nil {
			return diag.FromErr(err)
		}
	}

	ghInvitations, err := listInvitations(ctx, client, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	invitationIds := make(map[string]string, len(ghInvitations))
	for _, i := range ghInvitations {
		invitationIds[i.login] = strconv.FormatInt(*i.invitationID, 10)
	}

	if err = d.Set("invitation_ids", invitationIds); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCollaboratorsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	isOrg := meta.IsOrganization

	repoName := d.Id()
	ignoreTeams := d.Get("ignore_team").(*schema.Set).List()

	inIgnoreTeams, err := getTeamIdentities(ctx, meta, ignoreTeams)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Removing all collaborators from repository %s.", repoName))

	err = updateUserCollaboratorsAndInvites(ctx, client, owner, repoName, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	if isOrg {
		err = updateTeamCollaborators(ctx, client, owner, repoName, nil, inIgnoreTeams)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func listUserCollaborators(ctx context.Context, client *github.Client, owner, repoName string) (userCollaborators, error) {
	col := make([]userCollaborator, 0)

	affiliations := []string{"direct", "outside"}
	for _, affiliation := range affiliations {
		opt := &github.ListCollaboratorsOptions{
			ListOptions: github.ListOptions{
				PerPage: maxPerPage,
			},
			Affiliation: affiliation,
		}

		for {
			collaborators, resp, err := client.Repositories.ListCollaborators(ctx, owner, repoName, opt)
			if err != nil {
				return nil, err
			}

			for _, c := range collaborators {
				col = append(col, userCollaborator{
					userIdentity: userIdentity{
						login: c.GetLogin(),
					},
					permission: getPermission(c.GetRoleName()),
				})
			}

			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}
	return col, nil
}

func listInvitations(ctx context.Context, client *github.Client, owner, repoName string) ([]userCollaborator, error) {
	col := make([]userCollaborator, 0)

	opt := &github.ListOptions{PerPage: maxPerPage}
	for {
		invitations, resp, err := client.Repositories.ListInvitations(ctx, owner, repoName, opt)
		if err != nil {
			return nil, err
		}

		for _, i := range invitations {
			id := i.GetID()

			col = append(col, userCollaborator{
				userIdentity: userIdentity{
					login: i.GetInvitee().GetLogin(),
				},
				permission:   getPermission(i.GetPermissions()),
				invitationID: &id,
			})
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return col, nil
}

func listTeamCollaborators(ctx context.Context, client *github.Client, owner, repoName string, ignoreTeams []teamIdentity) (teamCollaborators, error) {
	col := make([]teamCollaborator, 0)

	opt := &github.ListOptions{
		PerPage: maxPerPage,
	}

	for {
		repoTeams, resp, err := client.Repositories.ListTeams(ctx, owner, repoName, opt)
		if err != nil {
			return nil, err
		}

		for _, t := range repoTeams {
			slug := t.GetSlug()
			if slices.ContainsFunc(ignoreTeams, func(ignore teamIdentity) bool {
				return ignore.slug == slug
			}) {
				continue
			}

			col = append(col, teamCollaborator{
				teamIdentity: teamIdentity{
					slug: slug,
				},
				permission: getPermission(t.GetPermission()),
			})
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return col, nil
}

func updateUserCollaboratorsAndInvites(ctx context.Context, client *github.Client, owner, repoName string, inUsers userCollaborators) error {
	lookup := make(map[string]userCollaborator)
	seen := make(map[string]any)
	remove := make([]string, 0)

	for _, inUser := range inUsers {
		lookup[inUser.login] = inUser
	}

	ghUsers, err := listUserCollaborators(ctx, client, owner, repoName)
	if err != nil {
		return err
	}

	for _, ghUser := range ghUsers {
		inUser, ok := lookup[ghUser.login]
		if ok {
			seen[ghUser.login] = nil

			if ghUser.permission != inUser.permission {
				tflog.Info(ctx, fmt.Sprintf("Updating user %s permission from %s to %s for repo %s.", inUser.login, ghUser.permission, inUser.permission, repoName))
				_, _, err := client.Repositories.AddCollaborator(
					ctx, owner, repoName, inUser.login, &github.RepositoryAddCollaboratorOptions{
						Permission: inUser.permission,
					})
				if err != nil {
					return err
				}
			}
		} else {
			remove = append(remove, ghUser.login)
		}
	}

	ghInvites, err := listInvitations(ctx, client, owner, repoName)
	if err != nil {
		return err
	}

	for _, ghInvite := range ghInvites {
		inInvite, ok := lookup[ghInvite.login]
		if ok {
			seen[ghInvite.login] = nil

			if ghInvite.permission != inInvite.permission {
				tflog.Info(ctx, fmt.Sprintf("Updating invite for user %s permission from %s to %s for repo %s.", inInvite.login, ghInvite.permission, inInvite.permission, repoName))
				_, _, err := client.Repositories.UpdateInvitation(ctx, owner, repoName, *ghInvite.invitationID, inInvite.permission)
				if err != nil {
					return err
				}
			}
		} else {
			tflog.Info(ctx, fmt.Sprintf("Deleting invite for user %s from repo %s.", ghInvite.login, repoName))
			_, err := client.Repositories.DeleteInvitation(ctx, owner, repoName, *ghInvite.invitationID)
			if err != nil {
				return handleArchivedRepoDelete(err, "repository collaborator invitation", ghInvite.login, owner, repoName)
			}
		}
	}

	for _, inUser := range inUsers {
		if _, ok := seen[inUser.login]; ok {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Inviting user %s to repo %s with permission %s.", inUser.login, repoName, inUser.permission))
		_, _, err := client.Repositories.AddCollaborator(ctx, owner, repoName, inUser.login, &github.RepositoryAddCollaboratorOptions{
			Permission: inUser.permission,
		})
		if err != nil {
			return err
		}
	}

	for _, l := range remove {
		tflog.Info(ctx, fmt.Sprintf("Removing user %s from repo %s.", l, repoName))
		_, err := client.Repositories.RemoveCollaborator(ctx, owner, repoName, l)
		if err != nil {
			return handleArchivedRepoDelete(err, "repository collaborator", l, owner, repoName)
		}
	}

	return nil
}

func updateTeamCollaborators(ctx context.Context, client *github.Client, owner, repoName string, inTeams teamCollaborators, ignoreTeams []teamIdentity) error {
	lookup := make(map[string]teamCollaborator)
	seen := make(map[string]any)
	remove := make([]string, 0)

	for _, inTeam := range inTeams {
		lookup[inTeam.slug] = inTeam
	}

	ghTeams, err := listTeamCollaborators(ctx, client, owner, repoName, ignoreTeams)
	if err != nil {
		return err
	}

	for _, ghTeam := range ghTeams {
		inTeam, ok := lookup[ghTeam.slug]
		if ok {
			seen[ghTeam.slug] = nil

			if ghTeam.permission != inTeam.permission {
				tflog.Info(ctx, fmt.Sprintf("Updating team %s permission from %s to %s for repo %s.", inTeam.slug, ghTeam.permission, inTeam.permission, repoName))
				_, err := client.Teams.AddTeamRepoBySlug(ctx, owner, inTeam.slug, owner, repoName, &github.TeamAddTeamRepoOptions{
					Permission: inTeam.permission,
				})
				if err != nil {
					return err
				}
			}
		} else {
			remove = append(remove, ghTeam.slug)
		}
	}

	for _, inTeam := range inTeams {
		if _, ok := seen[inTeam.slug]; ok {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Adding team %s to repo % with permission %s.", inTeam.slug, repoName, inTeam.permission))
		_, err := client.Teams.AddTeamRepoBySlug(ctx, owner, inTeam.slug, owner, repoName, &github.TeamAddTeamRepoOptions{
			Permission: inTeam.permission,
		})
		if err != nil {
			return err
		}
	}

	for _, s := range remove {
		tflog.Info(ctx, fmt.Sprintf("Removing team %s from repo %s.", s, repoName))
		_, err := client.Teams.RemoveTeamRepoBySlug(ctx, owner, s, owner, repoName)
		if err != nil {
			return handleArchivedRepoDelete(err, "team repository access", fmt.Sprintf("team %s", s), owner, repoName)
		}
	}

	return nil
}
