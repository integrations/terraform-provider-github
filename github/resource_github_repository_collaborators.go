package github

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryCollaborators() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubRepositoryCollaboratorsV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubRepositoryCollaboratorsStateUpgradeV0,
				Version: 0,
			},
		},

		CreateContext: resourceGithubRepositoryCollaboratorsCreate,
		ReadContext:   resourceGithubRepositoryCollaboratorsRead,
		UpdateContext: resourceGithubRepositoryCollaboratorsUpdate,
		DeleteContext: resourceGithubRepositoryCollaboratorsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryCollaboratorsImport,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the repository.",
			},
			"user": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of users.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:             schema.TypeString,
							Description:      "(Required) The user to add to the repository as a collaborator.",
							Required:         true,
							DiffSuppressFunc: caseInsensitive(),
						},
						"permission": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "push",
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
						"team_id": {
							Type:        schema.TypeString,
							Description: "Team ID or slug to add to the repository as a collaborator.",
							Required:    true,
						},
						"permission": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "push",
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
			"ignore_team": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of teams to ignore.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"team_id": {
							Type:        schema.TypeString,
							Description: "ID or slug of the team to ignore.",
							Required:    true,
						},
					},
				},
			},
		},

		CustomizeDiff: customdiff.Sequence(
			diffRepository,
			resourceGithubRepositoryCollaboratorsDiff,
		),
	}
}

func resourceGithubRepositoryCollaboratorsDiff(ctx context.Context, d *schema.ResourceDiff, m any) error {
	if d.HasChange("user") {
		users := d.Get("user").(*schema.Set).List()
		seen := make(map[string]any)

		for _, u := range users {
			user := u.(map[string]any)
			username := user["username"].(string)

			if _, ok := seen[username]; ok {
				return fmt.Errorf("duplicate username %s found in user collaborators", username)
			}
			seen[username] = nil
		}
	}

	if d.HasChange("team") && d.NewValueKnown("team") {
		v, diags := d.GetRawConfigAt(cty.GetAttrPath("team"))
		if diags.HasError() {
			return fmt.Errorf("error reading team config: %v", diags)
		}

		if !v.IsNull() && v.IsKnown() {
			seen := make(map[string]any)
			it := v.ElementIterator()
			for it.Next() {
				_, elem := it.Element()
				val := elem.GetAttr("team_id")
				if val.IsNull() || !val.IsKnown() {
					continue
				}

				teamID := val.AsString()
				if _, ok := seen[teamID]; ok {
					return fmt.Errorf("duplicate team %s found in team collaborators", teamID)
				}
				seen[teamID] = nil
			}
		}
	}

	if len(d.Id()) == 0 {
		return nil
	}

	if d.HasChange("user") {
		if err := d.SetNewComputed("invitation_ids"); err != nil {
			return fmt.Errorf("error setting invitation_ids to computed: %w", err)
		}
	}

	return nil
}

func resourceGithubRepositoryCollaboratorsCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
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

	inTeams, err := getTeamCollaborators(teams)
	if err != nil {
		return diag.FromErr(err)
	}

	inIgnoreTeams, err := getTeamIdentities(ignoreTeams)
	if err != nil {
		return diag.FromErr(err)
	}

	invitations, err := updateUserCollaboratorsAndInvites(ctx, client, owner, repoName, inUsers)
	if err != nil {
		return diag.FromErr(err)
	}

	if isOrg {
		err = updateTeamCollaborators(ctx, client, meta.id, owner, repoName, inTeams, inIgnoreTeams)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	repoID := int(repo.GetID())

	d.SetId(strconv.FormatInt(repo.GetID(), 10))
	if err := d.Set("repository_id", repoID); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("invitation_ids", invitations.flattenInvitations()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCollaboratorsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	isOrg := meta.IsOrganization

	repoName := d.Get("repository").(string)
	teams := d.Get("team").(*schema.Set).List()
	ignoreTeams := d.Get("ignore_team").(*schema.Set).List()

	inTeams, err := getTeamCollaborators(teams)
	if err != nil {
		return diag.FromErr(err)
	}

	inIgnoreTeams, err := getTeamIdentities(ignoreTeams)
	if err != nil {
		return diag.FromErr(err)
	}

	ghUsers, err := listUserCollaborators(ctx, client, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("user", ghUsers.flatten()); err != nil {
		return diag.FromErr(err)
	}

	if isOrg {
		ghTeams, err := listTeamCollaborators(ctx, client, owner, repoName, inTeams, inIgnoreTeams)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("team", ghTeams.flatten()); err != nil {
			return diag.FromErr(err)
		}
	}

	ghInvitations, err := listInvitations(ctx, client, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("invitation_ids", ghInvitations.flattenInvitations()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCollaboratorsUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
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

	inTeams, err := getTeamCollaborators(teams)
	if err != nil {
		return diag.FromErr(err)
	}

	inIgnoreTeams, err := getTeamIdentities(ignoreTeams)
	if err != nil {
		return diag.FromErr(err)
	}

	invitations, err := updateUserCollaboratorsAndInvites(ctx, client, owner, repoName, inUsers)
	if err != nil {
		return diag.FromErr(err)
	}

	if isOrg {
		err := updateTeamCollaborators(ctx, client, meta.id, owner, repoName, inTeams, inIgnoreTeams)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if err = d.Set("invitation_ids", invitations.flattenInvitations()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCollaboratorsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	isOrg := meta.IsOrganization

	repoName := d.Get("repository").(string)
	ignoreTeams := d.Get("ignore_team").(*schema.Set).List()

	inIgnoreTeams, err := getTeamIdentities(ignoreTeams)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Removing all collaborators from repository %s.", repoName))

	_, err = updateUserCollaboratorsAndInvites(ctx, client, owner, repoName, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	if isOrg {
		err = updateTeamCollaborators(ctx, client, meta.id, owner, repoName, nil, inIgnoreTeams)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubRepositoryCollaboratorsImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Id()

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	repoID := int(repo.GetID())

	d.SetId(strconv.FormatInt(repo.GetID(), 10))
	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", repoID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

// getUserCollaborators converts a slice of any type to a slice of userCollaborator.
func getUserCollaborators(col []any) (userCollaborators, error) {
	collaborators := make([]userCollaborator, len(col))

	for i, u := range col {
		m, ok := u.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("input invalid")
		}

		n, ok := m["username"]
		if !ok {
			return nil, fmt.Errorf("username missing")
		}

		username, ok := n.(string)
		if !ok || len(username) == 0 {
			return nil, fmt.Errorf("username invalid")
		}

		p, ok := m["permission"]
		if !ok {
			return nil, fmt.Errorf("permission missing")
		}

		permission, ok := p.(string)
		if !ok || len(permission) == 0 {
			return nil, fmt.Errorf("permission invalid")
		}

		uc := userCollaborator{
			userIdentity: userIdentity{
				login: username,
			},
			permission: permission,
		}

		collaborators[i] = uc
	}

	return collaborators, nil
}

// getTeamCollaborators returns a list of team collaborators represented by the input.
func getTeamCollaborators(col []any) (teamCollaborators, error) {
	collaborators := make([]teamCollaborator, len(col))

	for i, t := range col {
		m, ok := t.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("input invalid")
		}

		id, err := getTeamIdentity(m)
		if err != nil {
			return nil, err
		}

		permission, ok := m["permission"].(string)
		if !ok || len(permission) == 0 {
			return nil, fmt.Errorf("team input must include 'permission'")
		}

		collaborators[i] = teamCollaborator{
			teamIdentity: id,
			permission:   permission,
		}
	}

	return collaborators, nil
}

// getTeamIdentities returns a list of team identities represented by the input.
func getTeamIdentities(col []any) ([]teamIdentity, error) {
	identities := make([]teamIdentity, len(col))

	for i, t := range col {
		id, err := getTeamIdentity(t)
		if err != nil {
			return nil, err
		}
		identities[i] = id
	}

	return identities, nil
}

// getTeamIdentity returns a team identity represented by the input.
func getTeamIdentity(d any) (teamIdentity, error) {
	m, ok := d.(map[string]any)
	if !ok {
		return teamIdentity{}, fmt.Errorf("team input invalid")
	}

	o, ok := m["team_id"]
	if !ok {
		return teamIdentity{}, fmt.Errorf("team input must include 'team_id'")
	}

	id, ok := o.(string)
	if !ok || len(id) == 0 {
		return teamIdentity{}, fmt.Errorf("team_id must be a non-empty string")
	}

	return newLegacyTeamIdentity(id), nil
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

func listInvitations(ctx context.Context, client *github.Client, owner, repoName string) (userCollaborators, error) {
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

func listTeamCollaborators(ctx context.Context, client *github.Client, orgName, repoName string, inTeams teamCollaborators, ignoreTeams []teamIdentity) (teamCollaborators, error) {
	lookup := make(map[string]teamCollaborator)
	ignore := len(ignoreTeams) > 0
	col := make([]teamCollaborator, 0)

	for _, inTeam := range inTeams {
		lookup[inTeam.getTeamID()] = inTeam
	}

	opt := &github.ListOptions{
		PerPage: maxPerPage,
	}

	for {
		repoTeams, resp, err := client.Repositories.ListTeams(ctx, orgName, repoName, opt)
		if err != nil {
			return nil, err
		}

		for _, t := range repoTeams {
			slug := t.GetSlug()
			id := t.GetID()
			if ignore && slices.ContainsFunc(ignoreTeams, func(ignore teamIdentity) bool {
				if s, ok := ignore.getSlugOK(); ok {
					return slug == s
				} else {
					return id == ignore.getID()
				}
			}) {
				continue
			}

			var teamID *string
			if _, ok := lookup[slug]; ok {
				teamID = &slug
			}

			if teamID == nil {
				idStr := strconv.FormatInt(id, 10)
				if _, ok := lookup[idStr]; ok {
					teamID = &idStr
				}
			}

			col = append(col, teamCollaborator{
				teamIdentity: teamIdentity{
					id:     &id,
					slug:   &slug,
					teamID: teamID,
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

func updateUserCollaboratorsAndInvites(ctx context.Context, client *github.Client, owner, repoName string, inUsers userCollaborators) (userCollaborators, error) {
	lookup := make(map[string]userCollaborator)
	seen := make(map[string]any)
	remove := make([]string, 0)

	for _, inUser := range inUsers {
		lookup[inUser.login] = inUser
	}

	ghUsers, err := listUserCollaborators(ctx, client, owner, repoName)
	if err != nil {
		return nil, err
	}

	for _, ghUser := range ghUsers {
		inUser, ok := lookup[ghUser.login]
		if ok {
			seen[ghUser.login] = nil

			if ghUser.permission != inUser.permission {
				tflog.Info(ctx, fmt.Sprintf("Updating user %s permission from %s to %s for repo %s.", inUser.login, ghUser.permission, inUser.permission, repoName))
				_, _, err := client.Repositories.AddCollaborator(ctx, owner, repoName, inUser.login, &github.RepositoryAddCollaboratorOptions{Permission: inUser.permission})
				if err != nil {
					return nil, err
				}
			}
		} else {
			remove = append(remove, ghUser.login)
		}
	}

	ghInvites, err := listInvitations(ctx, client, owner, repoName)
	if err != nil {
		return nil, err
	}

	for _, ghInvite := range ghInvites {
		inInvite, ok := lookup[ghInvite.login]
		if ok {
			seen[ghInvite.login] = nil

			if ghInvite.permission != inInvite.permission {
				tflog.Info(ctx, fmt.Sprintf("Updating invite for user %s permission from %s to %s for repo %s.", inInvite.login, ghInvite.permission, inInvite.permission, repoName))
				_, _, err := client.Repositories.UpdateInvitation(ctx, owner, repoName, *ghInvite.invitationID, inInvite.permission)
				if err != nil {
					return nil, err
				}
			}
		} else {
			tflog.Info(ctx, fmt.Sprintf("Deleting invite for user %s from repo %s.", ghInvite.login, repoName))
			_, err := client.Repositories.DeleteInvitation(ctx, owner, repoName, *ghInvite.invitationID)
			if err != nil {
				return nil, handleArchivedRepoDelete(err, "repository collaborator invitation", ghInvite.login, owner, repoName)
			}
		}
	}

	for _, inUser := range inUsers {
		if _, ok := seen[inUser.login]; ok {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Inviting user %s to repo %s with permission %s.", inUser.login, repoName, inUser.permission))
		inv, _, err := client.Repositories.AddCollaborator(ctx, owner, repoName, inUser.login, &github.RepositoryAddCollaboratorOptions{Permission: inUser.permission})
		if err != nil {
			return nil, err
		}
		inUser.invitationID = inv.ID
		ghInvites = append(ghInvites, inUser)
	}

	for _, l := range remove {
		tflog.Info(ctx, fmt.Sprintf("Removing user %s from repo %s.", l, repoName))
		_, err := client.Repositories.RemoveCollaborator(ctx, owner, repoName, l)
		if err != nil {
			return nil, handleArchivedRepoDelete(err, "repository collaborator", l, owner, repoName)
		}
	}

	return ghInvites, nil
}

func updateTeamCollaborators(ctx context.Context, client *github.Client, orgID int64, orgName, repoName string, inTeams teamCollaborators, ignoreTeams []teamIdentity) error {
	lookup := make(map[string]teamCollaborator)
	seen := make(map[string]any)
	remove := make([]string, 0)

	for _, inTeam := range inTeams {
		lookup[inTeam.getTeamID()] = inTeam
	}

	ghTeams, err := listTeamCollaborators(ctx, client, orgName, repoName, inTeams, ignoreTeams)
	if err != nil {
		return err
	}

	for _, ghTeam := range ghTeams {
		slug := ghTeam.getSlug()
		if teamID, ok := ghTeam.getTeamIDOK(); ok {
			inTeam, ok := lookup[teamID]
			if !ok {
				continue
			}
			seen[teamID] = nil

			if ghTeam.permission != inTeam.permission {
				tflog.Info(ctx, fmt.Sprintf("Updating team %s permission from %s to %s for repo %s.", slug, ghTeam.permission, inTeam.permission, repoName))
				_, err := client.Teams.AddTeamRepoBySlug(ctx, orgName, slug, orgName, repoName, &github.TeamAddTeamRepoOptions{
					Permission: inTeam.permission,
				})
				if err != nil {
					return err
				}
			}
		} else {
			remove = append(remove, slug)
		}
	}

	for _, inTeam := range inTeams {
		teamID := inTeam.getTeamID()
		if _, ok := seen[teamID]; ok {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Adding team %s to repo %s with permission %s.", teamID, repoName, inTeam.permission))
		if slug, ok := inTeam.getSlugOK(); ok {
			_, err := client.Teams.AddTeamRepoBySlug(ctx, orgName, slug, orgName, repoName, &github.TeamAddTeamRepoOptions{Permission: inTeam.permission})
			if err != nil {
				return err
			}
		} else {
			_, err := client.Teams.AddTeamRepoByID(ctx, orgID, inTeam.getID(), orgName, repoName, &github.TeamAddTeamRepoOptions{Permission: inTeam.permission})
			if err != nil {
				return err
			}
		}
	}

	for _, s := range remove {
		tflog.Info(ctx, fmt.Sprintf("Removing team %s from repo %s.", s, repoName))
		_, err := client.Teams.RemoveTeamRepoBySlug(ctx, orgName, s, orgName, repoName)
		if err != nil {
			return handleArchivedRepoDelete(err, "team repository access", fmt.Sprintf("team %s", s), orgName, repoName)
		}
	}

	return nil
}
